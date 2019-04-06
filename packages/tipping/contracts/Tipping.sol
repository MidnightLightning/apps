pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";

import "@daonuts/token/contracts/IERC20.sol";
import "@daonuts/common/contracts/Names.sol";

contract Tipping is AragonApp, Names {
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
    IERC20 public currency;

    /// ENS
    /* bytes32 internal constant DAONUTS_LABEL = keccak256("daonuts"); */
    /* bytes32 internal constant DAONUTS_LABEL = 0x53bf7a5ae2fa6880bad06201387e90063522a09407b9b95effeb2a65d870dd4c; */
    /* bytes32 internal constant DAONUTS_NODE = keccak256(abi.encodePacked(ETH_TLD_NODE, DAONUTS_LABEL)); */
    bytes32 internal constant DAONUTS_NODE = 0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e;

    event Tip(string fromName, string toName, uint amount, ContentType ctype, uint40 cid);
    event Claim(string toName, uint balance);

    function initialize(AbstractENS _ens, IERC20 _currency) onlyInit public {
        initialized();

        ens = _ens;
        setResolver(ens.resolver(PUBLIC_RESOLVER_NODE));
        currency = _currency;
    }

    function claim(address _owner) external {
        string memory toName = nameOfOwner(_owner);
        require( bytes(toName).length != 0, USER_NOT_REGISTERED );
        bytes32 nameHash = keccak256(toName);
        uint256 balance = balances[nameHash];
        require( balance > 0, NOTHING_TO_CLAIM );
        delete balances[nameHash];
        require( currency.transfer(_owner, balance), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
        emit Claim(toName, balance);
    }

    function tip(string _toName, uint _amount, ContentType _ctype, uint40 _cid) external {
        address to = ownerOfName(_toName);
        if(to == 0x0) {
            bytes32 nameHash = keccak256(_toName);
            require( currency.transferFrom(msg.sender, this, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
            balances[nameHash] = balances[nameHash].add(_amount);
        } else {
            require( currency.transferFrom(msg.sender, to, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED );
        }
        string memory fromName = nameOfOwner(msg.sender);
        emit Tip(fromName, _toName, _amount, _ctype, _cid);
    }

}
