pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/common/IForwarder.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";

import "@daonuts/common/contracts/Names.sol";

contract Registry is AragonApp, IForwarder, Names {
    using SafeMath for uint256;

    /// Events
    event RegistrationPeriodStarted(bytes32 root);
    event Registered(address indexed owner, string username);
    event Deregistered(address indexed owner, string username);
    event RootNodeTransferred(address owner);

    /// State
    AbstractENS public ens;

    bytes32[] public roots;
    mapping(bytes32 => bool) public activeRegPeriod;

    /// ACL
    /* bytes32 constant public START_REGISTRATION_PERIOD = keccak256("START_REGISTRATION_PERIOD"); */
    bytes32 constant public START_REGISTRATION_PERIOD = 0xd31f4ba181fa04f6e556e75747124c08760a53dec98821ac56200ec037aa2bb7;
    /* bytes32 constant public TRANSFER_ROOT_NODE = keccak256("TRANSFER_ROOT_NODE"); */
    bytes32 constant public TRANSFER_ROOT_NODE = 0x371d57b5d5e36ffacd760261ee1986c14c5a44484f6cd32d970413d31467313e;

    // Errors
    string private constant REGISTRATION_EXISTS = "REGISTRATION_EXISTS";
    string private constant NO_ACTIVE_REGISTRATION_PERIOD = "NO_ACTIVE_REGISTRATION_PERIOD";
    string private constant REGISTRATION_PERIOD_EXISTS = "REGISTRATION_PERIOD_EXISTS";
    string private constant REGISTRATION_NOT_FOUND = "REGISTRATION_NOT_FOUND";
    string private constant INVALID = "INVALID";
    string private constant ERROR_CAN_NOT_FORWARD = "REGISTRY_CAN_NOT_FORWARD";
    string private constant ERROR_REGISTRY_NOT_OWNER = "REGISTRY_NOT_OWNER";
    string private constant ERROR_NAME_NOT_SET = "NAME_NOT_SET";
    string private constant ERROR_ADDR_NOT_SET = "ADDR_NOT_SET";

    function initialize(AbstractENS _ens, address altResolver, bytes32 _root) onlyInit public {
        initialized();

        ens = _ens;

        if(altResolver == address(0)) {
          setResolver(ens.resolver(PUBLIC_RESOLVER_NODE));
        } else {
          setResolver(altResolver);
        }

        _addRoot(_root);
    }

    /**
     * @notice Propose a new registration merkle root
     * @param _root New distribution merkle root
     */
    function addRoot(bytes32 _root) auth(START_REGISTRATION_PERIOD) public {
        require( activeRegPeriod[_root] == false, REGISTRATION_PERIOD_EXISTS );
        _addRoot(_root);
    }

    function _addRoot(bytes32 _root) internal {
        roots.push(_root);
        activeRegPeriod[_root] = true;
        emit RegistrationPeriodStarted(_root);
    }

    function registerSelf(bytes32 _root, string _username, bytes32[] _proof) external {
        require( activeRegPeriod[_root] == true, NO_ACTIVE_REGISTRATION_PERIOD );
        require( ownerOfName(_username) == address(0), REGISTRATION_EXISTS );
        require( bytes(nameOfOwner(msg.sender)).length == 0, REGISTRATION_EXISTS );

        require( validate(_root, msg.sender, _username, _proof), INVALID );
        _register(msg.sender, _username);
    }

    function deregisterSelf() external {
        string memory username = nameOfOwner(msg.sender);
        require( bytes(username).length != 0, REGISTRATION_NOT_FOUND );
        _deregister(msg.sender, username);
    }

    function _register(address _owner, string _username) internal {
        bytes32 ownerLabel = sha3HexAddress(_owner);
        ens.setSubnodeOwner(DAONUTS_NODE, ownerLabel, address(this));
        bytes32 ownerNode = keccak256(abi.encodePacked(DAONUTS_NODE, ownerLabel));
        resolver.setName(ownerNode, _username);
        require( keccak256(abi.encodePacked(resolver.name(ownerNode))) == keccak256(abi.encodePacked(_username)), ERROR_NAME_NOT_SET );

        bytes32 usernameLabel = keccak256(_username);
        ens.setSubnodeOwner(DAONUTS_NODE, usernameLabel, address(this));
        bytes32 usernameNode = keccak256(abi.encodePacked(DAONUTS_NODE, usernameLabel));
        resolver.setAddr(usernameNode, _owner);
        require( resolver.addr(usernameNode) == _owner, ERROR_ADDR_NOT_SET );

        emit Registered(_owner, _username);
    }

    function _deregister(address _owner, string _username) internal {
        bytes32 ownerNode = addrNode(_owner);
        if(ens.owner(ownerNode) != address(this)){
            claimOwnerNode(_owner);
        }
        resolver.setName(ownerNode, "");

        bytes32 usernameNode = nameNode(_username);
        if(ens.owner(usernameNode) != address(this)){
            claimNameNode(_username);
        }
        resolver.setAddr(usernameNode, address(0));

        emit Deregistered(_owner, _username);
    }

    function claimOwnerNode(address _owner) internal {                        // claim from a previous Registry contract
        require( ens.owner(DAONUTS_NODE) == address(this), ERROR_REGISTRY_NOT_OWNER );
        bytes32 ownerLabel = sha3HexAddress(_owner);
        ens.setSubnodeOwner(DAONUTS_NODE, ownerLabel, address(this));
    }

    function claimNameNode(string _username) internal {                        // claim from a previous Registry contract
        require( ens.owner(DAONUTS_NODE) == address(this), ERROR_REGISTRY_NOT_OWNER );
        bytes32 usernameLabel = keccak256(_username);
        ens.setSubnodeOwner(DAONUTS_NODE, usernameLabel, address(this));
    }

    /**
     * @notice Get length of roots array
     */
    function getRootsCount() public view returns(uint count) {
        return roots.length;
    }

    function validate(
      bytes32 _root, address _owner, string _username, bytes32[] _proof
    ) public view returns (bool) {
        bytes32 hash = keccak256(_owner, _username);
        return checkProof(_root, _proof, hash);
    }

    function hash(address _owner, string _username) public view returns (bytes32) {
        return keccak256(_owner, _username);
    }

    function hashAddress(address _owner) public view returns (bytes32) {
        return keccak256(_owner);
    }

    function hashString(string _username) public view returns (bytes32) {
        return keccak256(_username);
    }

    function checkProof(bytes32 root, bytes32[] proof, bytes32 hash) public pure returns (bool) {

        for (uint i = 0; i < proof.length; i++) {
            if (hash < proof[i]) {
                hash = keccak256(hash, proof[i]);
            } else {
                hash = keccak256(proof[i], hash);
            }
        }

        return hash == root;
    }

    /**
    * @notice Execute desired action as a registered user
    * @dev IForwarder interface conformance. Forwards any token holder action.
    * @param _evmScript Script being executed
    */
    function forward(bytes _evmScript) public {
        require(canForward(msg.sender, _evmScript), ERROR_CAN_NOT_FORWARD);
        bytes memory input = new bytes(0); // TODO: Consider input for this

        address[] memory blacklist;

        runScript(_evmScript, input, blacklist);
    }

    function canForward(address _sender, bytes) public view returns (bool) {
        // can forward if sender has registered a username
        return hasInitialized() && bytes(resolver.name(addrNode(_sender))).length != 0;
    }

    function isForwarder() public pure returns (bool) {
        return true;
    }

    function ownsRootNode() public view returns (bool) {
        return ens.owner(DAONUTS_NODE) == address(this);
    }

    function rootNodeOwner() public view returns (address) {
        return ens.owner(DAONUTS_NODE);
    }

    /**
    * Transfers ownership of a node to a new address. May only be called by the current
    * owner of the node.
    * @param owner The address of the new owner.
    */
    function transferRootNode(address owner) auth(TRANSFER_ROOT_NODE) public {
        ens.setOwner(DAONUTS_NODE, owner);
        emit RootNodeTransferred(owner);
    }
}
