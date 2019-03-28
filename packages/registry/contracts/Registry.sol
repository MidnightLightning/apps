pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/common/IForwarder.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";

contract Registry is AragonApp, IForwarder {
    using SafeMath for uint256;

    /// Events
    event RegistrationPeriodStarted(bytes32 root);
    event Registered(address indexed owner, bytes32 indexed username);
    event Deregistered(address indexed owner, bytes32 indexed username);

    /// State
    bytes32[] public roots;
    mapping(address => bytes32) public ownerToUsername;
    mapping(bytes32 => address) public usernameToOwner;
    mapping(bytes32 => bool) public activeRegPeriod;

    /// ACL
    bytes32 constant public START_REGISTRATION_PERIOD = keccak256("START_REGISTRATION_PERIOD");

    // Errors
    string private constant REGISTRATION_EXISTS = "REGISTRATION_EXISTS";
    string private constant NO_ACTIVE_REGISTRATION_PERIOD = "NO_ACTIVE_REGISTRATION_PERIOD";
    string private constant REGISTRATION_PERIOD_EXISTS = "REGISTRATION_PERIOD_EXISTS";
    string private constant REGISTRATION_NOT_FOUND = "REGISTRATION_NOT_FOUND";
    string private constant INVALID = "INVALID";
    string private constant ERROR_CAN_NOT_FORWARD = "REGISTRY_CAN_NOT_FORWARD";

    function initialize(bytes32 _root) onlyInit public {
        initialized();
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
        emit Registered(_owner, _username);
    }

    function _deregister(address _owner, bytes32 _username) internal {
        delete ownerToUsername[_owner];
        delete usernameToOwner[_username];
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
}
