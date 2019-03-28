pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";

import "@daonuts/token/contracts/Token.sol";
import "@daonuts/registry/contracts/Registry.sol";

contract Tipping is AragonApp {
    using SafeMath for uint256;

    enum ContentType                                       { NONE, COMMENT, POST }

    /* struct Post {
      bytes20                                        author;                    // only use if not already reg
    }

    mapping(uint40 => Post)         public posts; */
    mapping(bytes32 => uint256) public balances;

    /// ACL
    bytes32 constant public NONE = keccak256("NONE");

    /// Errors
    string private constant ERROR_TOKEN_TRANSFER_FROM_REVERTED = "FINANCE_TKN_TRANSFER_FROM_REVERT";
    string private constant USER_NOT_REGISTERED = "USER_NOT_REGISTERED";
    string private constant NOTHING_TO_CLAIM = "NOTHING_TO_CLAIM";

    Token public token;
    Registry public registry;

    event Tip(bytes32 indexed fromName, bytes32 indexed toName, uint amount, ContentType ctype, uint40 cid);
    event Claim(bytes32 indexed toName, uint balance);

    function initialize(Token _token, Registry _registry) onlyInit public {
        initialized();

        token = _token;
        registry = _registry;
    }

    function getOwner(bytes32 _username) public view returns (address) {
      return registry.usernameToOwner(_username);
    }

    function getUsername(address _owner) public view returns (bytes32) {
      return registry.ownerToUsername(_owner);
    }

    function claim(address _owner) external {
        bytes32 toName = registry.ownerToUsername(_owner);
        require( toName != 0x0, USER_NOT_REGISTERED );
        uint256 balance = balances[toName];
        require( balance > 0, NOTHING_TO_CLAIM );
        delete balances[toName];
        require( token.transfer(_owner, balance), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
        emit Claim(toName, balance);
    }

    function tip(bytes32 _toName, uint _amount, ContentType _ctype, uint40 _cid) external {
        address to = registry.usernameToOwner(_toName);
        if(to == 0x0) {
            require( token.transferFrom(msg.sender, this, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
            balances[_toName] = balances[_toName].add(_amount);
        } else {
            require( token.transferFrom(msg.sender, to, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
        }
        bytes32 fromName = registry.ownerToUsername(msg.sender);
        emit Tip(fromName, _toName, _amount, _ctype, _cid);
    }

}
