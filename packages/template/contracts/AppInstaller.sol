pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/Kernel.sol";
import "@aragon/os/contracts/apm/Repo.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";

import "@daonuts/common/contracts/Names.sol";
import "@daonuts/token/contracts/Token.sol";
import "@daonuts/token-manager/contracts/TokenManager.sol";
import "@daonuts/karma-cap-voting/contracts/KarmaCapVoting.sol";
import "@daonuts/tipping/contracts/Tipping.sol";
import "@daonuts/registry/contracts/Registry.sol";
import "@daonuts/distribution/contracts/Distribution.sol";
import "@daonuts/hamburger/contracts/Hamburger.sol";


contract AppInstaller {

    AbstractENS aragonENS;
    AbstractENS ens;
    Names names;
    uint64 constant PCT = 10 ** 16;

    event TokenCreated(address token, string name);
    event InstalledApp(address appProxy, bytes32 appId);

    constructor(AbstractENS _aragonENS, AbstractENS _ens, bytes32 _rootNode) {
        aragonENS = _aragonENS;
        ens = _ens;

        address resolver = _ens.resolver(_rootNode);
        require( resolver != address(0), "NO_RESOLVER" );
        names = new Names(resolver, _rootNode);
    }

    function installCurrencyManager(Kernel _dao, bytes32 _tokenManagerAppId, string _name, string _symbol) {
        Token currency = new Token(_name, 18, _symbol, true);
        emit TokenCreated(address(currency), _name);
        TokenManager currencyManager = TokenManager(_dao.newAppInstance(_tokenManagerAppId, latestVersionAppBase(_tokenManagerAppId)));
        currency.changeController(currencyManager);
        currencyManager.initialize(address(currency), true, 0);
        emit InstalledApp(currencyManager, _tokenManagerAppId);
    }

    function installKarmaManager(Kernel _dao, bytes32 _tokenManagerAppId, string _name, string _symbol) {
        Token karma = new Token(_name, 18, _symbol, false);
        emit TokenCreated(address(karma), _name);
        TokenManager karmaManager = TokenManager(_dao.newAppInstance(_tokenManagerAppId, latestVersionAppBase(_tokenManagerAppId)));
        karma.changeController(karmaManager);
        karmaManager.initialize(address(karma), false, 0);
        emit InstalledApp(karmaManager, _tokenManagerAppId);
    }

    function installVoting(Kernel _dao, bytes32 _votingAppId, address currency, address karma) {
        KarmaCapVoting voting = KarmaCapVoting(_dao.newAppInstance(_votingAppId, latestVersionAppBase(_votingAppId)));
        voting.initialize(currency, karma, 50 * PCT, 20 * PCT, 1 days);
        emit InstalledApp(voting, _votingAppId);
    }

    function installTipping(Kernel _dao, bytes32 _tippingAppId, address currency) {
        Tipping tipping = Tipping(_dao.newAppInstance(_tippingAppId, latestVersionAppBase(_tippingAppId)));
        tipping.initialize(address(names), currency);
        emit InstalledApp(tipping, _tippingAppId);
    }

    function installRegistry(Kernel _dao, bytes32 _registryAppId, bytes32 _rootNode, bytes32 _regRoot) {
        Registry registry = Registry(_dao.newAppInstance(_registryAppId, latestVersionAppBase(_registryAppId)));
        registry.initialize(address(ens), _rootNode, _regRoot);
        emit InstalledApp(registry, _registryAppId);
    }

    function installDistribution(Kernel _dao, bytes32 _distributionAppId, address _currencyManager, address _karmaManager, bytes32 _distRoot) {
        Distribution distribution = Distribution(_dao.newAppInstance(_distributionAppId, latestVersionAppBase(_distributionAppId)));
        distribution.initialize(address(names), _currencyManager, _karmaManager, _distRoot);
        emit InstalledApp(distribution, _distributionAppId);
    }

    function installHamburger(Kernel _dao, bytes32 _hamburgerAppId, address _currencyManager) {
        Hamburger hamburger = Hamburger(_dao.newAppInstance(_hamburgerAppId, latestVersionAppBase(_hamburgerAppId)));
        hamburger.initialize(address(names), _currencyManager);
        emit InstalledApp(hamburger, _hamburgerAppId);
    }

    function latestVersionAppBase(bytes32 appId) public view returns (address base) {
        Repo repo = Repo(PublicResolver(aragonENS.resolver(appId)).addr(appId));
        (,base,) = repo.getLatest();

        return base;
    }
}
