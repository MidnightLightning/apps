pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/Kernel.sol";
import "@aragon/os/contracts/acl/ACL.sol";
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


contract PermissionSetter {

    function setPermissions(
      Kernel _dao, TokenManager _currencyManager, TokenManager _karmaManager,
      KarmaCapVoting _voting, Registry _registry, Distribution _distribution,
      Hamburger _hamburger, Tipping _tipping) {
        ACL acl = ACL(_dao.acl());

        acl.createPermission(_tipping, _tipping, _tipping.NONE(), _tipping);

        acl.createPermission(this, _currencyManager, _currencyManager.MINT_ROLE(), this);
        _currencyManager.mint(msg.sender, 1); // Give one currency to msg.sender

        acl.createPermission(this, _karmaManager, _karmaManager.MINT_ROLE(), this);
        _karmaManager.mint(msg.sender, 1); // Give one karma to msg.sender

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
        acl.grantPermission(_voting, _currencyManager, _currencyManager.MINT_ROLE());
        acl.grantPermission(_voting, _karmaManager, _karmaManager.MINT_ROLE());

        // allow distribution to mint
        acl.grantPermission(_distribution, _currencyManager, _currencyManager.MINT_ROLE());
        acl.grantPermission(_distribution, _karmaManager, _karmaManager.MINT_ROLE());

        // Clean up permissions
        /* acl.grantPermission(msg.sender, _dao, _dao.APP_MANAGER_ROLE()); */
        /* acl.revokePermission(this, _dao, _dao.APP_MANAGER_ROLE()); */
        /* acl.setPermissionManager(msg.sender, _dao, _dao.APP_MANAGER_ROLE()); */

        /* acl.grantPermission(msg.sender, acl, acl.CREATE_PERMISSIONS_ROLE()); */
        /* acl.revokePermission(this, acl, acl.CREATE_PERMISSIONS_ROLE()); */
        /* acl.setPermissionManager(msg.sender, acl, acl.CREATE_PERMISSIONS_ROLE()); */
    }
}
