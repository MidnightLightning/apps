pragma solidity ^0.4.24;

import "@aragon/os/contracts/apps/AragonApp.sol";
import "@aragon/os/contracts/lib/math/SafeMath.sol";
import "@aragon/os/contracts/lib/math/SafeMath64.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";
import "@aragon/os/contracts/ens/ENSConstants.sol";

import "@daonuts/token-manager/contracts/TokenManager.sol";
import "@daonuts/token/contracts/IERC20.sol";
import "@daonuts/registry/contracts/Registry.sol";

contract Hamburger is AragonApp, ENSConstants {
    using SafeMath for uint256;
    using SafeMath64 for uint64;

    struct Asset {
        bool    active;
        bool    requireReg;
        address owner;
        uint8   tax;
        uint64  lastPaymentDate;
        uint256 price;
        uint256 balance;
        string  name;
        string  data;
    }

    /// Events
    event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId);
    event Balance(uint256 indexed _tokenId, uint256 _balance);
    event Price(uint256 indexed _tokenId, uint256 _price);
    event Data(uint256 indexed _tokenId, string _data);
    event Tax(uint256 indexed _tokenId, uint8 _tax);

    /// State
    mapping(uint => Asset) public assets;
    uint public assetsCount;
    AbstractENS public ens;
    PublicResolver public resolver;
    TokenManager public currencyManager;
    IERC20 public currency;
    Registry public registry;

    /// ENS
    /* bytes32 internal constant DAONUTS_LABEL = keccak256("daonuts"); */
    /* bytes32 internal constant DAONUTS_LABEL = 0x53bf7a5ae2fa6880bad06201387e90063522a09407b9b95effeb2a65d870dd4c; */
    /* bytes32 internal constant DAONUTS_NODE = keccak256(abi.encodePacked(ETH_TLD_NODE, DAONUTS_LABEL)); */
    bytes32 internal constant DAONUTS_NODE = 0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e;

    /// ACL
    bytes32 constant public PURCHASE_ASSET_ROLE = keccak256("PURCHASE_ASSET_ROLE");
    bytes32 constant public COMMONS_ROLE = keccak256("COMMONS_ROLE");

    // Errors
    string private constant NO_ACTIVE_ASSET = "NO_ACTIVE_ASSET";
    string private constant ACTIVE_ASSET_EXISTS = "ACTIVE_ASSET_EXISTS";
    string private constant NOT_OWNER = "NOT_OWNER";
    string private constant USER_NOT_REGISTERED = "USER_NOT_REGISTERED";
    string private constant INSUFFICIENT_BALANCE = "INSUFFICIENT_BALANCE";
    string private constant ERROR_TOKEN_TRANSFER_FROM_TAX = "FINANCE_TKN_TRANSFER_TAX";
    string private constant ERROR_TOKEN_TRANSFER_FROM_PRICE = "FINANCE_TKN_TRANSFER_PRICE";
    string private constant ERROR_TOKEN_TRANSFER_FROM_REFUND = "FINANCE_TKN_TRANSFER_REFUND";
    string private constant ERROR_TOKEN_TRANSFER_FROM_REVERTED = "FINANCE_TKN_TRANSFER_FROM_REVERT";
    string private constant ERROR_0x0_NOT_ALLOWED = "ERROR_0x0_NOT_ALLOWED";
    string private constant ERROR_INVALID_NAME = "ERROR_INVALID_NAME";
    string private constant ERROR_TM_BURN = "ERROR_TM_BURN";

    /* function initialize(IERC20 _currency, Registry _registry) onlyInit public { */
    function initialize(AbstractENS _ens, TokenManager _currencyManager, Registry _registry) onlyInit public {
        initialized();

        ens = _ens;
        resolver = PublicResolver(ens.resolver(PUBLIC_RESOLVER_NODE));
        currencyManager = _currencyManager;
        currency = _currencyManager.token();
        registry = _registry;
    }

    /**
     * @notice Purchase the banner
     * @param _tokenId Asset tokenId
     * @param _price Set new price
     * @param _data Set new data (if applicable)
     */
    function buy(uint256 _tokenId, uint256 _price, string _data, uint256 _credit) external {
        Asset storage asset = assets[_tokenId];
        if(asset.owner != address(this)) {
          payTax(_tokenId);
        }

        _transferFrom(asset.owner, msg.sender, _tokenId);
        // require min balance of 1 day of tax
        require(_credit > _price.mul(asset.tax).div(100), INSUFFICIENT_BALANCE);
        credit(_tokenId, _credit, true);
        asset.price = _price;
        asset.data = _data;
    }

    /**
     * @notice Pay any tax due
     * @param _tokenId Asset tokenId
     */
    function payTax(uint256 _tokenId) public {
        Asset storage asset = assets[_tokenId];
        uint amount = taxDue(_tokenId);
        if(amount > asset.balance) {
          amount = asset.balance;
          _reclaim(_tokenId);
        }
        if(amount > 0){
          asset.balance = asset.balance.sub(amount);
          currencyManager.burn(address(this), amount);
        }
        asset.lastPaymentDate = getTimestamp64();
    }

    /// @notice Transfer ownership of an NFT -- THE CALLER IS RESPONSIBLE
    ///  TO CONFIRM THAT `_to` IS CAPABLE OF RECEIVING NFTS OR ELSE
    ///  THEY MAY BE PERMANENTLY LOST
    /// @dev Throws unless `msg.sender` is the current owner, an authorized
    ///  operator, or the approved address for this NFT. Throws if `_from` is
    ///  not the current owner. Throws if `_to` is the zero address. Throws if
    ///  `_tokenId` is not a valid NFT.
    /// @param _from The current owner of the NFT
    /// @param _to The new owner
    /// @param _tokenId The NFT to transfer
    function transferFrom(address _from, address _to, uint256 _tokenId) external payable {
        payTax(_tokenId);
        // it does not accept payments in this version
        require(msg.value == 0);

        _transferFrom(_from, _to, _tokenId);
    }

    function _transferFrom(address _from, address _to, uint256 _tokenId) internal {
        Asset storage asset = assets[_tokenId];

        if(asset.requireReg) {
          // require new owner to be registered
          bytes32 username = registry.ownerToUsername(_to);
          require(username != bytes32(0), USER_NOT_REGISTERED);
        }
        require(asset.active == true, NO_ACTIVE_ASSET);
        require(_to != address(0), ERROR_0x0_NOT_ALLOWED);
        require(_from == asset.owner, NOT_OWNER);

        // current owner can transfer
        // if initiated by non-owner
        if(asset.owner != msg.sender) {
          // ...owner must sell at asset.price (Harberger rules)
          require(currency.transferFrom(msg.sender, asset.owner, asset.price), ERROR_TOKEN_TRANSFER_FROM_PRICE);
        }

        _refund(_tokenId);

        asset.owner = _to;

        emit Transfer(_from, _to, _tokenId);
    }

    function _reclaim(uint256 _tokenId) internal {
        Asset storage asset = assets[_tokenId];
        emit Transfer(asset.owner, address(this), _tokenId);
        asset.owner = address(this);
        delete asset.price;
        emit Price(_tokenId, asset.price);
        delete asset.data;
        emit Data(_tokenId, asset.data);
    }

    function _refund(uint256 _tokenId) internal {
        if(asset.balance > 0) return;

        Asset storage asset = assets[_tokenId];
        uint refund = asset.balance;
        asset.balance = 0;
        require(currency.transfer(asset.owner, refund), ERROR_TOKEN_TRANSFER_FROM_REFUND);
    }

    /// @notice Find the owner of an NFT
    /// @dev NFTs assigned to zero address are considered invalid, and queries
    ///  about them do throw.
    /// @param _tokenId The identifier for an NFT
    /// @return The address of the owner of the NFT
    function ownerOf(uint256 _tokenId) external view returns (address) {
        /* TODO owner reverts to `this` or 0x0 if taxDue is greater than balance `!hasSurplusBalance` */
        return assets[_tokenId].owner;
    }

    function taxDue(uint256 _tokenId) public view returns (uint256) {
        Asset storage asset = assets[_tokenId];
        uint dailyTax = asset.price.mul(asset.tax).div(100);
        uint numDays = getTimestamp64().sub(asset.lastPaymentDate).div(1 days);
        return dailyTax.mul(numDays);
    }

    function hasSurplusBalance(uint256 _tokenId) public view returns (bool) {
        Asset storage asset = assets[_tokenId];
        return taxDue(_tokenId) > asset.balance;
    }

    /**
     * @notice Set price of asset (only asset owner)
     * @param _tokenId Asset tokenId
     * @param _price New price
     */
    function setPrice(uint256 _tokenId, uint256 _price) public {
        payTax(_tokenId);
        Asset storage asset = assets[_tokenId];
        require(msg.sender == asset.owner, NOT_OWNER);
        asset.price = _price;
        emit Price(_tokenId, _price);
    }

    /**
     * @notice Set data associated with asset (only asset owner)
     * @param _tokenId Asset tokenId
     * @param _data New data
     */
    function setData(uint256 _tokenId, string _data) public {
        Asset storage asset = assets[_tokenId];
        require(msg.sender == asset.owner, NOT_OWNER);
        asset.data = _data;
        emit Data(_tokenId, _data);
    }

    /**
     * @notice Add credit to an asset
     * @param _tokenId Asset tokenId
     * @param _amount Amount to credit
     * @param _onlyIfSelfOwned Only complete if sender is owner
     */
    function credit(uint256 _tokenId, uint256 _amount, bool _onlyIfSelfOwned) public {
        Asset storage asset = assets[_tokenId];
        if(_onlyIfSelfOwned)
          require(msg.sender == asset.owner, NOT_OWNER);

        require(currency.transferFrom(msg.sender, address(this), _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED);
        asset.balance = asset.balance.add(_amount);
        emit Balance(_tokenId, asset.balance);
    }

    /**
     * @notice Claim refund
     * @param _tokenId Asset tokenId
     * @param _amount Amount to credit
     */
    function debit(uint256 _tokenId, uint256 _amount) public {
        Asset storage asset = assets[_tokenId];
        payTax(_tokenId);
        require(msg.sender == asset.owner, NOT_OWNER);
        require(_amount <= asset.balance, INSUFFICIENT_BALANCE);
        require(currency.transfer(msg.sender, _amount), ERROR_TOKEN_TRANSFER_FROM_REVERTED);
        asset.balance = asset.balance.sub(_amount);
        emit Balance(_tokenId, asset.balance);
    }

    /**
     * @notice Mint a new asset (only commons)
     * @param _name Name
     * @param _tax Tax
     */
    function mint(string _name, uint8 _tax, bool requireReg) auth(COMMONS_ROLE) external {
        require(bytes(_name).length != 0, ERROR_INVALID_NAME);
        uint _tokenId = assetsCount++;
        Asset storage asset = assets[_tokenId];
        require(asset.active == false, ACTIVE_ASSET_EXISTS);
        asset.active = true;
        asset.requireReg = requireReg;
        asset.owner = address(this);
        asset.tax = _tax;
        asset.name = _name;
        emit Transfer(address(0), address(this), _tokenId);
    }

    /**
     * @notice Burn an asset (only commons)
     * @param _tokenId Asset tokenId
     */
    function burn(uint256 _tokenId) auth(COMMONS_ROLE) external {
        Asset storage asset = assets[_tokenId];
        require(asset.active == true, NO_ACTIVE_ASSET);
        payTax(_tokenId);
        _refund(_tokenId);
        // TODO - return balance if currently owned and > 0

        delete assets[_tokenId];
        emit Transfer(asset.owner, address(0), _tokenId);
    }

    /**
     * @notice Set asset tax (only commons)
     * @param _tokenId Asset tokenId
     * @param _tax New tax
     */
    function setTax(uint256 _tokenId, uint8 _tax) auth(COMMONS_ROLE) external {
        Asset storage asset = assets[_tokenId];
        if(asset.owner != address(this))
          payTax(_tokenId);
        require(asset.active == true, NO_ACTIVE_ASSET);
        asset.tax = _tax;
        emit Tax(_tokenId, _tax);
    }
}
