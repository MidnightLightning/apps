pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/common/IForwarder.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";
import "@aragon/os/contracts/ens/ENSConstants.sol";

contract Registry is AragonApp, IForwarder, ENSConstants {
    using SafeMath for uint256;

    /// Events
    event RegistrationPeriodStarted(bytes32 root);
    event Registered(address indexed owner, bytes32 indexed username);
    event Deregistered(address indexed owner, bytes32 indexed username);

    /// State
    AbstractENS public ens;
    PublicResolver public resolver;

    bytes32[] public roots;
    mapping(address => bytes32) public ownerToUsername;
    mapping(bytes32 => address) public usernameToOwner;
    mapping(bytes32 => bool) public activeRegPeriod;

    /// ENS
    /* bytes32 internal constant DAONUTS_LABEL = keccak256("daonuts"); */
    /* bytes32 internal constant DAONUTS_LABEL = 0x53bf7a5ae2fa6880bad06201387e90063522a09407b9b95effeb2a65d870dd4c; */
    /* bytes32 internal constant DAONUTS_NODE = keccak256(abi.encodePacked(ETH_TLD_NODE, DAONUTS_LABEL)); */
    bytes32 internal constant DAONUTS_NODE = 0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e;

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

    function initialize(AbstractENS _ens, bytes32 _root) onlyInit public {
        initialized();

        ens = _ens;
        resolver = PublicResolver(ens.resolver(PUBLIC_RESOLVER_NODE));

        // We need ownership to create subnodes
        /* require(ens.owner(DAONUTS_NODE) == address(this), ERROR_REGISTRY_NOT_OWNER); */

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

    function registerSelf(bytes32 _root, bytes32 _username, bytes32[] _proof) external {
        require( activeRegPeriod[_root] == true, NO_ACTIVE_REGISTRATION_PERIOD );
        require( ownerToUsername[msg.sender] == 0x0 && usernameToOwner[_username] == 0x0, REGISTRATION_EXISTS );
        require( validate(_root, msg.sender, _username, _proof), INVALID );
        _register(msg.sender, _username);
    }

    function deregisterSelf() external {
        bytes32 username = ownerToUsername[msg.sender];
        require( username != 0x0, REGISTRATION_NOT_FOUND );
        _deregister(msg.sender, username);
    }

    function _register(address _owner, bytes32 _username) internal {
        ownerToUsername[_owner] = _username;
        usernameToOwner[_username] = _owner;

        bytes32 label = keccak256(_username);
        ens.setSubnodeOwner(DAONUTS_NODE, label, this);
        bytes32 node = keccak256(abi.encodePacked(DAONUTS_NODE, label));
        resolver.setAddr(node, _owner);

        emit Registered(_owner, _username);
    }

    function _deregister(address _owner, bytes32 _username) internal {
        delete ownerToUsername[_owner];
        delete usernameToOwner[_username];

        bytes32 label = keccak256(_username);
        bytes32 node = keccak256(abi.encodePacked(DAONUTS_NODE, label));
        resolver.setAddr(node, address(0));

        emit Deregistered(_owner, _username);
    }

    /**
     * @notice Get length of roots array
     */
    function getRootsCount() public view returns(uint count) {
        return roots.length;
    }

    function validate(
      bytes32 _root, address _owner, bytes32 _username, bytes32[] _proof
    ) public view returns (bool) {
        bytes32 hash = keccak256(_owner, _username);
        return checkProof(_root, _proof, hash);
    }

    function hash(address _owner, bytes32 _username) public view returns (bytes32 hash) {
        hash = keccak256(_owner, _username);
    }

    function hashAddress(address _owner) public view returns (bytes32 hash) {
        hash = keccak256(_owner);
    }

    function hashBytes32(bytes32 _username) public view returns (bytes32 hash) {
        hash = keccak256(_username);
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
        return hasInitialized() && ownerToUsername[_sender] != 0x0;
    }

    function isForwarder() public pure returns (bool) {
        return true;
    }

    function ownsRootNode() public view returns (bool) {
      return ens.owner(DAONUTS_NODE) == address(this);
    }

    /**
     * Transfers ownership of a node to a new address. May only be called by the current
     * owner of the node.
     * @param node The node to transfer ownership of.
     * @param owner The address of the new owner.
     */
     // Can't do this by vote because msg.sender becomes the voting app
    /* function transferRootNode(address owner) auth(TRANSFER_ROOT_NODE) public {
        ens.setOwner(DAONUTS_NODE, owner);
    } */
}
