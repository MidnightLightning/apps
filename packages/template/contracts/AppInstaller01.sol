pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/Kernel.sol";
import "@aragon/os/contracts/acl/ACL.sol";
import "@aragon/os/contracts/apm/Repo.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";

/* import "@daonuts/distribution/contracts/IDistribution.sol"; */
/* import "@daonuts/hamburger/contracts/IHamburger.sol"; */
import "@daonuts/karma-cap-voting/contracts/IKarmaCapVoting.sol";
/* import "@daonuts/registry/contracts/IRegistry.sol"; */
import "@daonuts/tipping/contracts/ITipping.sol";
import "@daonuts/token-manager/contracts/ITokenManager.sol";
import "@daonuts/token/contracts/Token.sol";
import "@daonuts/common/contracts/Names.sol";
import "@daonuts/common/contracts/IPublicResolver.sol";

contract AppInstaller01 {

    AbstractENS aragonENS;
    bytes32 apmRootNode;
    AbstractENS ens;
    bytes32 rootNode;
    Names public names;
    Token public currency;
    Token public karma;
    uint64 constant PCT = 10 ** 16;

    constructor(AbstractENS _aragonENS, bytes32 _apmRootNode, AbstractENS _ens, bytes32 _rootNode) {
        aragonENS = _aragonENS;
        apmRootNode = _apmRootNode;
        ens = _ens;
        rootNode = _rootNode;

        address resolver = _ens.resolver(rootNode);
        require( resolver != address(0), "NO_RESOLVER" );

        names = new Names(resolver, rootNode);
    }

    function createTokens() external {
        currency = new Token("Currency", 18, "NUTS", true);
        karma = new Token("Karma", 18, "KARMA", false);
    }

    function install(Kernel _dao, address _nextInstaller) external {
        (ITokenManager currencyManager) = installCurrencyManager(_dao);
        (ITokenManager karmaManager) = installKarmaManager(_dao);
        IKarmaCapVoting voting = installVoting(_dao);
        installTipping(_dao);
        permissions(_dao, currencyManager, karmaManager, voting, _nextInstaller);
    }

    function installVoting(Kernel _dao) internal returns (IKarmaCapVoting voting) {
        bytes32 votingAppId = namehash(apmRootNode, "daonuts-karma-cap-voting");
        voting = IKarmaCapVoting(_dao.newAppInstance(votingAppId, latestVersionAppBase(votingAppId)));
        voting.initialize(address(currency), address(karma), 50 * PCT, 20 * PCT, 1 days);
    }

    function installTipping(Kernel _dao) internal returns (ITipping tipping) {
        ACL acl = ACL(_dao.acl());
        bytes32 tippingAppId = namehash(apmRootNode, "daonuts-tipping");
        tipping = ITipping(_dao.newAppInstance(tippingAppId, latestVersionAppBase(tippingAppId)));
        tipping.initialize(address(names), address(currency));
        acl.createPermission(tipping, tipping, tipping.NONE(), tipping);
    }

    function installCurrencyManager(Kernel _dao) internal returns(ITokenManager currencyManager) {
        ACL acl = ACL(_dao.acl());
        bytes32 currencyManagerAppId = namehash(apmRootNode, "daonuts-token-manager");
        currencyManager = ITokenManager(_dao.newAppInstance(currencyManagerAppId, latestVersionAppBase(currencyManagerAppId)));
        currency.changeController(currencyManager);
        currencyManager.initialize(address(currency), true, 0);
    }

    function installKarmaManager(Kernel _dao) internal returns(ITokenManager karmaManager) {
        ACL acl = ACL(_dao.acl());
        bytes32 karmaManagerAppId = namehash(apmRootNode, "daonuts-token-manager");
        karmaManager = ITokenManager(_dao.newAppInstance(karmaManagerAppId, latestVersionAppBase(karmaManagerAppId)));
        karma.changeController(karmaManager);
        karmaManager.initialize(address(karma), false, 0);
    }

    function permissions(
      Kernel _dao, ITokenManager _currencyManager, ITokenManager _karmaManager,
      IKarmaCapVoting _voting, address _nextInstaller) internal {
        ACL acl = ACL(_dao.acl());

        acl.createPermission(this, _currencyManager, _currencyManager.MINT_ROLE(), this);
        _currencyManager.mint(msg.sender, 1); // Give one currency to msg.sender

        acl.createPermission(this, _karmaManager, _karmaManager.MINT_ROLE(), this);
        _karmaManager.mint(msg.sender, 1); // Give one karma to msg.sender

        /* acl.createPermission(_voting, _registry, _registry.START_REGISTRATION_PERIOD(), _voting); */
        /* acl.createPermission(_voting, _registry, _registry.TRANSFER_ROOT_NODE(), _voting); */

        /* acl.createPermission(_voting, _distribution, _distribution.START_DISTRIBUTION(), _voting); */

        /* acl.createPermission(_voting, _hamburger, _hamburger.COMMONS_ROLE(), _voting); */

        /* acl.createPermission(_registry, _hamburger, _hamburger.PURCHASE_ASSET_ROLE(), _voting); */

        // registered users can create votes
        /* acl.createPermission(_registry, _voting, _voting.CREATE_VOTES_ROLE(), _voting); */

        // allow hamburger to burn currency
        /* acl.createPermission(_hamburger, _currencyManager, _currencyManager.BURN_ROLE(), _voting); */

        // allow voting to mint
        acl.grantPermission(_voting, _currencyManager, _currencyManager.MINT_ROLE());
        acl.grantPermission(_voting, _karmaManager, _karmaManager.MINT_ROLE());

        // allow distribution to mint
        /* acl.grantPermission(_distribution, _currencyManager, _currencyManager.MINT_ROLE()); */
        /* acl.grantPermission(_distribution, _karmaManager, _karmaManager.MINT_ROLE()); */

        /* acl.grantPermission(_nextInstaller, _dao, _dao.APP_MANAGER_ROLE());
        acl.revokePermission(this, _dao, _dao.APP_MANAGER_ROLE());
        acl.setPermissionManager(_nextInstaller, _dao, _dao.APP_MANAGER_ROLE());
        acl.grantPermission(_nextInstaller, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.revokePermission(this, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.setPermissionManager(_nextInstaller, acl, acl.CREATE_PERMISSIONS_ROLE()); */

        // Clean up permissions
        /* acl.grantPermission(msg.sender, _dao, _dao.APP_MANAGER_ROLE());
        acl.revokePermission(this, _dao, _dao.APP_MANAGER_ROLE());
        acl.setPermissionManager(msg.sender, _dao, _dao.APP_MANAGER_ROLE());

        acl.grantPermission(msg.sender, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.revokePermission(this, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.setPermissionManager(msg.sender, acl, acl.CREATE_PERMISSIONS_ROLE()); */
    }

    function latestVersionAppBase(bytes32 appId) public view returns (address base) {
        Repo repo = Repo(IPublicResolver(aragonENS.resolver(appId)).addr(appId));
        (,base,) = repo.getLatest();

        return base;
    }

    function namehash(bytes32 node, string name) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(node, keccak256(bytes(name))));
    }

}
