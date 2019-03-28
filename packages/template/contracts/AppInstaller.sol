pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/Kernel.sol";
import "@aragon/os/contracts/acl/ACL.sol";
import "@aragon/os/contracts/apm/APMNamehash.sol";
import "@aragon/os/contracts/apm/Repo.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";
import "@aragon/os/contracts/lib/ens/ENS.sol";

import "@daonuts/token/contracts/Token.sol";
import "@daonuts/token-manager/contracts/TokenManager.sol";
import "@daonuts/distribution/contracts/Distribution.sol";
import "@daonuts/karma-cap-voting/contracts/KarmaCapVoting.sol";
import "@daonuts/tipping/contracts/Tipping.sol";
import "@daonuts/registry/contracts/Registry.sol";
import "@daonuts/hamburger/contracts/Hamburger.sol";

contract AppInstaller is APMNamehash {

    ENS ens;
    uint64 constant PCT = 10 ** 16;
    address constant ANY_ENTITY = address(-1);
    /* function install(Kernel _dao, ENS _ens, Token _token, Token _karma, bytes32 _regRoot, bytes32 _distRoot) external { */
    function install(Kernel _dao, ENS _ens, bytes32 _regRoot, bytes32 _distRoot) external {
        ens = _ens;
        Token token = new Token("App token", 18, "APP", true);
        Token karma = new Token("App karma", 18, "KAR", false);
        (TokenManager tokenManager) = installTokenManager(_dao, token);
        (TokenManager karmaManager) = installKarmaManager(_dao, karma);
        KarmaCapVoting voting = installVoting(_dao, token, karma);
        Registry registry = installRegistry(_dao, _regRoot);
        Distribution distribution = installDistribution(_dao, tokenManager, karmaManager, registry, _distRoot);
        Hamburger hamburger = installHamburger(_dao, token, registry);
        installTipping(_dao, token, registry);
        permissions(_dao, tokenManager, karmaManager, voting, registry, distribution, hamburger);
    }

    function installTokenManager(Kernel _dao, Token _token) internal returns(TokenManager tokenManager) {
        ACL acl = ACL(_dao.acl());
        bytes32 tokenManagerAppId = apmNamehash("daonuts-token-manager");
        tokenManager = TokenManager(_dao.newAppInstance(tokenManagerAppId, latestVersionAppBase(tokenManagerAppId)));
        _token.changeController(tokenManager);
        tokenManager.initialize(_token, true, 0);
        acl.createPermission(this, tokenManager, tokenManager.MINT_ROLE(), this);
        tokenManager.mint(msg.sender, 1); // Give one token to msg.sender
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

    function installVoting(Kernel _dao, Token _token, Token _karma) internal returns (KarmaCapVoting voting) {
        bytes32 votingAppId = apmNamehash("daonuts-karma-cap-voting");
        voting = KarmaCapVoting(_dao.newAppInstance(votingAppId, latestVersionAppBase(votingAppId)));
        voting.initialize(_token, _karma, 50 * PCT, 20 * PCT, 1 days);
    }

    function installRegistry(Kernel _dao, bytes32 _regRoot) internal returns (Registry registry) {
        bytes32 registryAppId = apmNamehash("daonuts-registry");
        registry = Registry(_dao.newAppInstance(registryAppId, latestVersionAppBase(registryAppId)));
        registry.initialize(_regRoot);
    }

    function installDistribution(Kernel _dao, TokenManager _tokenManager, TokenManager _karmaManager, Registry _registry, bytes32 _distRoot) internal returns (Distribution distribution) {
        bytes32 distributionAppId = apmNamehash("daonuts-distribution");
        distribution = Distribution(_dao.newAppInstance(distributionAppId, latestVersionAppBase(distributionAppId)));
        distribution.initialize(_tokenManager, _karmaManager, _registry, _distRoot);
    }

    function installHamburger(Kernel _dao, Token _token, Registry _registry) internal returns (Hamburger hamburger) {
        bytes32 hamburgerAppId = apmNamehash("daonuts-hamburger");
        hamburger = Hamburger(_dao.newAppInstance(hamburgerAppId, latestVersionAppBase(hamburgerAppId)));
        hamburger.initialize(_token, _registry);
    }

    function installTipping(Kernel _dao, Token _token, Registry _registry) internal returns (Tipping tipping) {
        ACL acl = ACL(_dao.acl());
        bytes32 tippingAppId = apmNamehash("daonuts-tipping");
        tipping = Tipping(_dao.newAppInstance(tippingAppId, latestVersionAppBase(tippingAppId)));
        tipping.initialize(_token, _registry);
        acl.createPermission(tipping, tipping, tipping.NONE(), tipping);
    }

    function permissions(
      Kernel _dao, TokenManager _tokenManager, TokenManager _karmaManager,
      KarmaCapVoting _voting, Registry _registry, Distribution _distribution,
      Hamburger _hamburger) internal {
        ACL acl = ACL(_dao.acl());

        acl.createPermission(_voting, _registry, _registry.START_REGISTRATION_PERIOD(), _voting);

        acl.createPermission(_voting, _distribution, _distribution.START_DISTRIBUTION(), _voting);

        acl.createPermission(_voting, _hamburger, _hamburger.COMMONS_ROLE(), _voting);

        acl.createPermission(_registry, _hamburger, _hamburger.PURCHASE_ASSET_ROLE(), _voting);

        // registered users can create votes
        acl.createPermission(_registry, _voting, _voting.CREATE_VOTES_ROLE(), _voting);

        // allow voting to mint
        acl.grantPermission(_voting, _tokenManager, _tokenManager.MINT_ROLE());
        acl.grantPermission(_voting, _karmaManager, _karmaManager.MINT_ROLE());

        // allow distribution to mint
        acl.grantPermission(_distribution, _tokenManager, _tokenManager.MINT_ROLE());
        acl.grantPermission(_distribution, _karmaManager, _karmaManager.MINT_ROLE());
    }

    function latestVersionAppBase(bytes32 appId) public view returns (address base) {
        Repo repo = Repo(PublicResolver(ens.resolver(appId)).addr(appId));
        (,base,) = repo.getLatest();

        return base;
    }

}
