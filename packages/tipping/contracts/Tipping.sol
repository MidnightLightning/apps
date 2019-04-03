pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";

import "@daonuts/token/contracts/IERC20.sol";
import "@daonuts/registry/contracts/Registry.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";
import "@aragon/os/contracts/ens/ENSConstants.sol";

contract Tipping is AragonApp, ENSConstants {
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

    AbstractENS public ens;
    PublicResolver public resolver;
    IERC20 public currency;
    Registry public registry;

    /// ENS
    /* bytes32 internal constant DAONUTS_LABEL = keccak256("daonuts"); */
    /* bytes32 internal constant DAONUTS_LABEL = 0x53bf7a5ae2fa6880bad06201387e90063522a09407b9b95effeb2a65d870dd4c; */
    /* bytes32 internal constant DAONUTS_NODE = keccak256(abi.encodePacked(ETH_TLD_NODE, DAONUTS_LABEL)); */
    bytes32 internal constant DAONUTS_NODE = 0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e;

    event Tip(bytes32 indexed fromName, bytes32 indexed toName, uint amount, ContentType ctype, uint40 cid);
    event Claim(bytes32 indexed toName, uint balance);

    function initialize(AbstractENS _ens, IERC20 _currency, Registry _registry) onlyInit public {
        initialized();

        ens = _ens;
        resolver = PublicResolver(ens.resolver(PUBLIC_RESOLVER_NODE));
        currency = _currency;
        registry = _registry;
    }

    function getOwner(bytes32 _username) public view returns (address) {
        bytes32 node = usernameNode(_username);
        return resolver.addr(node);
        /* return registry.usernameToOwner(_username); */
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
        require( currency.transfer(_owner, balance), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
        emit Claim(toName, balance);
    }

    function tip(bytes32 _toName, uint _amount, ContentType _ctype, uint40 _cid) external {
        /* address to = registry.usernameToOwner(_toName); */
        bytes32 node = usernameNode(_toName);
        address to = resolver.addr(node);
        if(to == 0x0) {
            require( currency.transferFrom(msg.sender, this, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
            balances[_toName] = balances[_toName].add(_amount);
        } else {
            require( currency.transferFrom(msg.sender, to, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
        }
        bytes32 fromName = registry.ownerToUsername(msg.sender);
        emit Tip(fromName, _toName, _amount, _ctype, _cid);
    }

    function usernameNode(bytes32 _username) public view returns (bytes32 node) {
        bytes32 label = keccak256(_username);
        node = keccak256(abi.encodePacked(DAONUTS_NODE, label));
    }

}
