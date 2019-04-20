pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/Kernel.sol";
import "@aragon/os/contracts/acl/ACL.sol";
import "@aragon/os/contracts/apm/Repo.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";

import "@daonuts/distribution/contracts/IDistribution.sol";
import "@daonuts/hamburger/contracts/IHamburger.sol";
import "@daonuts/karma-cap-voting/contracts/IKarmaCapVoting.sol";
import "@daonuts/registry/contracts/IRegistry.sol";
import "@daonuts/token-manager/contracts/ITokenManager.sol";
import "@daonuts/token/contracts/Token.sol";
import "@daonuts/common/contracts/Names.sol";
import "@daonuts/common/contracts/IPublicResolver.sol";

contract AppInstaller02 {

    AbstractENS aragonENS;
    bytes32 apmRootNode;
    AbstractENS ens;
    bytes32 rootNode;
    Names public names;
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

    function install(
      Kernel _dao, ITokenManager _currencyManager, ITokenManager _karmaManager,
      IKarmaCapVoting _voting, bytes32 _regRoot, bytes32 _distRoot) external {

        IRegistry registry = installRegistry(_dao, _regRoot);
        IDistribution distribution = installDistribution(_dao, _currencyManager, _karmaManager, _distRoot);
        IHamburger hamburger = installHamburger(_dao, _currencyManager);
        permissions(_dao, _currencyManager, _karmaManager, _voting, registry, distribution, hamburger);
    }

    function installDistribution(Kernel _dao, ITokenManager _currencyManager, ITokenManager _karmaManager, bytes32 _distRoot) internal returns (IDistribution distribution) {
        bytes32 distributionAppId = namehash(apmRootNode, "daonuts-distribution");
        distribution = IDistribution(_dao.newAppInstance(distributionAppId, latestVersionAppBase(distributionAppId)));
        distribution.initialize(address(names), address(_currencyManager), address(_karmaManager), _distRoot);
    }

    function installHamburger(Kernel _dao, ITokenManager _currencyManager) internal returns (IHamburger hamburger) {
        bytes32 hamburgerAppId = namehash(apmRootNode, "daonuts-hamburger");
        hamburger = IHamburger(_dao.newAppInstance(hamburgerAppId, latestVersionAppBase(hamburgerAppId)));
        hamburger.initialize(address(names), address(_currencyManager));
    }

    function installRegistry(Kernel _dao, bytes32 _regRoot) internal returns (IRegistry registry) {
        bytes32 registryAppId = namehash(apmRootNode, "daonuts-registry");
        registry = IRegistry(_dao.newAppInstance(registryAppId, latestVersionAppBase(registryAppId)));
        registry.initialize(address(ens), rootNode, _regRoot);
    }

    function permissions(
      Kernel _dao, ITokenManager _currencyManager, ITokenManager _karmaManager,
      IKarmaCapVoting _voting, IRegistry _registry, IDistribution _distribution,
      IHamburger _hamburger) internal {
        ACL acl = ACL(_dao.acl());

        acl.createPermission(_voting, _registry, _registry.START_REGISTRATION_PERIOD(), _voting);
        acl.createPermission(_voting, _registry, _registry.TRANSFER_ROOT_NODE(), _voting);

        acl.createPermission(_voting, _distribution, _distribution.START_DISTRIBUTION(), _voting);

        acl.createPermission(_voting, _hamburger, _hamburger.COMMONS_ROLE(), _voting);

        acl.createPermission(_registry, _hamburger, _hamburger.PURCHASE_ASSET_ROLE(), _voting);

        // registered users can create votes
        acl.createPermission(_registry, _voting, _voting.CREATE_VOTES_ROLE(), _voting);

        // allow hamburger to burn currency
        acl.createPermission(_hamburger, _currencyManager, _currencyManager.BURN_ROLE(), _voting);

        // allow voting to mint
        /* acl.grantPermission(_voting, _currencyManager, _currencyManager.MINT_ROLE()); */
        /* acl.grantPermission(_voting, _karmaManager, _karmaManager.MINT_ROLE()); */

        // allow distribution to mint
        acl.grantPermission(_distribution, _currencyManager, _currencyManager.MINT_ROLE());
        acl.grantPermission(_distribution, _karmaManager, _karmaManager.MINT_ROLE());

        // Clean up permissions
        acl.grantPermission(_voting, _dao, _dao.APP_MANAGER_ROLE());
        acl.revokePermission(this, _dao, _dao.APP_MANAGER_ROLE());
        acl.setPermissionManager(_voting, _dao, _dao.APP_MANAGER_ROLE());

        acl.grantPermission(_voting, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.revokePermission(this, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.setPermissionManager(_voting, acl, acl.CREATE_PERMISSIONS_ROLE());
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
