pragma solidity ^0.4.24;

import "@aragon/os/contracts/kernel/IKernel.sol";

interface IAppInstaller {
    /* function install(address dao, address ens, address token, address karma, bytes32 regRoot, bytes32 distRoot) external; */
    function install(address dao, address ens, bytes32 regRoot, bytes32 distRoot) external;
}
