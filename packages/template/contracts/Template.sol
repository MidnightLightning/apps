/*
 * SPDX-License-Identitifer:    GPL-3.0-or-later
 *
 * This file requires contract dependencies which are licensed as
 * GPL-3.0-or-later, forcing it to also be licensed as such.
 *
 * This is the only file in your project that requires this license and
 * you are free to choose a different license for the rest of the project.
 */

pragma solidity ^0.4.24;

import "@aragon/os/contracts/factory/DAOFactory.sol";
import "@aragon/os/contracts/apm/Repo.sol";
import "@aragon/os/contracts/lib/ens/ENS.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";
import "@aragon/os/contracts/apm/APMNamehash.sol";

/* import "@daonuts/token/contracts/Token.sol"; */

import "./IAppInstaller.sol";

contract TemplateBase is APMNamehash {
    ENS public ens;
    DAOFactory public fac;

    mapping (address => mapping (string => address)) tokenCache;

    /* event DeployToken(address token, address indexed cacheOwner, string tokenSymbol); */
    event DeployInstance(address dao);
    event InstalledApp(address appProxy, bytes32 appId);

    function TemplateBase(DAOFactory _fac, ENS _ens) {
        ens = _ens;

        // If no factory is passed, get it from on-chain bare-kit
        if (address(_fac) == address(0)) {
            bytes32 bareKit = apmNamehash("bare-kit");
            fac = TemplateBase(latestVersionAppBase(bareKit)).fac();
        } else {
            fac = _fac;
        }
    }

    function latestVersionAppBase(bytes32 appId) public view returns (address base) {
        Repo repo = Repo(PublicResolver(ens.resolver(appId)).addr(appId));
        (,base,) = repo.getLatest();

        return base;
    }

    /* function cacheToken(Token _token, address _owner, string _tokenSymbol) internal {
        tokenCache[_owner][_tokenSymbol] = _token;
        emit DeployToken(_token, _owner, _tokenSymbol);
    }

    function popTokenCache(address _owner, string _tokenSymbol) internal returns (Token) {
        require(tokenCache[_owner][_tokenSymbol] != address(0));
        Token token = Token(tokenCache[_owner][_tokenSymbol]);
        delete tokenCache[_owner][_tokenSymbol];

        return token;
    }

    function readTokenCache(address _owner, string _tokenSymbol) public view returns (address) {
        return tokenCache[_owner][_tokenSymbol];
    }

    function newToken(string _tokenName, string _tokenSymbol, bool _transfersEnabled) public {
        Token token = new Token(_tokenName, 18, _tokenSymbol, _transfersEnabled);
        cacheToken(token, msg.sender, _tokenSymbol);
    } */
}

contract Template is TemplateBase {

    function Template(ENS ens) TemplateBase(DAOFactory(0), ens) {
    }

    function newInstance(bytes32 _regRoot, bytes32 _distRoot, IAppInstaller _installer) public {
        Kernel dao = fac.newDAO(this);
        ACL acl = ACL(dao.acl());
        acl.createPermission(this, dao, dao.APP_MANAGER_ROLE(), this);

        /* Token token = popTokenCache(msg.sender, "APP"); */
        /* token.changeController(_installer); */
        /* Token karma = popTokenCache(msg.sender, "KAR"); */
        /* karma.changeController(_installer); */

        // run external installer
        acl.grantPermission(_installer, dao, dao.APP_MANAGER_ROLE());
        acl.grantPermission(_installer, acl, acl.CREATE_PERMISSIONS_ROLE());
        /* _installer.install(dao, ens, token, karma, _regRoot, _distRoot); */
        _installer.install(dao, ens, _regRoot, _distRoot);
        acl.revokePermission(_installer, dao, dao.APP_MANAGER_ROLE());
        acl.revokePermission(_installer, acl, acl.CREATE_PERMISSIONS_ROLE());

        // Clean up permissions
        acl.grantPermission(msg.sender, dao, dao.APP_MANAGER_ROLE());
        acl.revokePermission(this, dao, dao.APP_MANAGER_ROLE());
        acl.setPermissionManager(msg.sender, dao, dao.APP_MANAGER_ROLE());

        acl.grantPermission(msg.sender, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.revokePermission(this, acl, acl.CREATE_PERMISSIONS_ROLE());
        acl.setPermissionManager(msg.sender, acl, acl.CREATE_PERMISSIONS_ROLE());

        DeployInstance(dao);
    }
}
