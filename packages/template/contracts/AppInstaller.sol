pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/Kernel.sol";
import "@aragon/os/contracts/acl/ACL.sol";
import "@aragon/os/contracts/apm/APMNamehash.sol";
import "@aragon/os/contracts/apm/Repo.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";
import "@aragon/os/contracts/lib/ens/ENS.sol";

import "@daonuts/distribution/contracts/Distribution.sol";
import "@daonuts/hamburger/contracts/Hamburger.sol";
import "@daonuts/karma-cap-voting/contracts/KarmaCapVoting.sol";
import "@daonuts/registry/contracts/Registry.sol";
import "@daonuts/tipping/contracts/Tipping.sol";
import "@daonuts/token/contracts/Token.sol";
import "@daonuts/token-manager/contracts/TokenManager.sol";

contract AppInstaller is APMNamehash {

    ENS ens;
    uint64 constant PCT = 10 ** 16;
    address constant ANY_ENTITY = address(-1);
    /* function install(Kernel _dao, ENS _ens, Token _currency, Token _karma, bytes32 _regRoot, bytes32 _distRoot) external { */
    function install(Kernel _dao, ENS _ens, bytes32 _regRoot, bytes32 _distRoot) external {
        ens = _ens;
        Token currency = new Token("Currency", 18, "NUTS", true);
        Token karma = new Token("Karma", 18, "KARM", false);
        (TokenManager currencyManager) = installCurrencyManager(_dao, currency);
        (TokenManager karmaManager) = installKarmaManager(_dao, karma);
        KarmaCapVoting voting = installVoting(_dao, currency, karma);
        Registry registry = installRegistry(_dao, _regRoot);
        Distribution distribution = installDistribution(_dao, currencyManager, karmaManager, registry, _distRoot);
        Hamburger hamburger = installHamburger(_dao, currencyManager, registry);
        installTipping(_dao, currency, registry);
        permissions(_dao, currencyManager, karmaManager, voting, registry, distribution, hamburger);
    }

    function installDistribution(Kernel _dao, TokenManager _currencyManager, TokenManager _karmaManager, Registry _registry, bytes32 _distRoot) internal returns (Distribution distribution) {
        bytes32 distributionAppId = apmNamehash("daonuts-distribution");
        distribution = Distribution(_dao.newAppInstance(distributionAppId, latestVersionAppBase(distributionAppId)));
        distribution.initialize(_currencyManager, _karmaManager, _registry, _distRoot);
    }

    function installHamburger(Kernel _dao, TokenManager _currencyManager, Registry _registry) internal returns (Hamburger hamburger) {
        bytes32 hamburgerAppId = apmNamehash("daonuts-hamburger");
        hamburger = Hamburger(_dao.newAppInstance(hamburgerAppId, latestVersionAppBase(hamburgerAppId)));
        hamburger.initialize(_currencyManager, _registry);
    }

    function installVoting(Kernel _dao, Token _currency, Token _karma) internal returns (KarmaCapVoting voting) {
        bytes32 votingAppId = apmNamehash("daonuts-karma-cap-voting");
        voting = KarmaCapVoting(_dao.newAppInstance(votingAppId, latestVersionAppBase(votingAppId)));
        voting.initialize(_currency, _karma, 50 * PCT, 20 * PCT, 1 days);
    }

    function installRegistry(Kernel _dao, bytes32 _regRoot) internal returns (Registry registry) {
        bytes32 registryAppId = apmNamehash("daonuts-registry");
        registry = Registry(_dao.newAppInstance(registryAppId, latestVersionAppBase(registryAppId)));
        registry.initialize(_regRoot);
    }

    function installTipping(Kernel _dao, Token _currency, Registry _registry) internal returns (Tipping tipping) {
        ACL acl = ACL(_dao.acl());
        bytes32 tippingAppId = apmNamehash("daonuts-tipping");
        tipping = Tipping(_dao.newAppInstance(tippingAppId, latestVersionAppBase(tippingAppId)));
        tipping.initialize(_currency, _registry);
        acl.createPermission(tipping, tipping, tipping.NONE(), tipping);
    }

    function installCurrencyManager(Kernel _dao, Token _currency) internal returns(TokenManager currencyManager) {
        ACL acl = ACL(_dao.acl());
        bytes32 currencyManagerAppId = apmNamehash("daonuts-token-manager");
        currencyManager = TokenManager(_dao.newAppInstance(currencyManagerAppId, latestVersionAppBase(currencyManagerAppId)));
        _currency.changeController(currencyManager);
        currencyManager.initialize(_currency, true, 0);
        acl.createPermission(this, currencyManager, currencyManager.MINT_ROLE(), this);
        currencyManager.mint(msg.sender, 1); // Give one currency to msg.sender
    }

    function installKarmaManager(Kernel _dao, Token _karma) internal returns(TokenManager karmaManager) {
        ACL acl = ACL(_dao.acl());
        bytes32 karmaManagerAppId = apmNamehash("daonuts-token-manager");
        karmaManager = TokenManager(_dao.newAppInstance(karmaManagerAppId, latestVersionAppBase(karmaManagerAppId)));
        _karma.changeController(karmaManager);
        karmaManager.initialize(_karma, false, 0);
        acl.createPermission(this, karmaManager, karmaManager.MINT_ROLE(), this);
        karmaManager.mint(msg.sender, 1); // Give one karma to msg.sender
    }

    function permissions(
      Kernel _dao, TokenManager _currencyManager, TokenManager _karmaManager,
      KarmaCapVoting _voting, Registry _registry, Distribution _distribution,
      Hamburger _hamburger) internal {
        ACL acl = ACL(_dao.acl());

        acl.createPermission(_voting, _registry, _registry.START_REGISTRATION_PERIOD(), _voting);

        acl.createPermission(_voting, _distribution, _distribution.START_DISTRIBUTION(), _voting);

        acl.createPermission(_voting, _hamburger, _hamburger.COMMONS_ROLE(), _voting);

        acl.createPermission(_registry, _hamburger, _hamburger.PURCHASE_ASSET_ROLE(), _voting);

        // registered users can create votes
        acl.createPermission(_registry, _voting, _voting.CREATE_VOTES_ROLE(), _voting);

        // allow hamburger to burn currency
        acl.createPermission(_hamburger, _currencyManager, _currencyManager.BURN_ROLE(), _voting);

        // allow voting to mint
        acl.grantPermission(_voting, _currencyManager, _currencyManager.MINT_ROLE());
        acl.grantPermission(_voting, _karmaManager, _karmaManager.MINT_ROLE());

        // allow distribution to mint
        acl.grantPermission(_distribution, _currencyManager, _currencyManager.MINT_ROLE());
        acl.grantPermission(_distribution, _karmaManager, _karmaManager.MINT_ROLE());
    }

    function latestVersionAppBase(bytes32 appId) public view returns (address base) {
        Repo repo = Repo(PublicResolver(ens.resolver(appId)).addr(appId));
        (,base,) = repo.getLatest();

        return base;
    }

}
