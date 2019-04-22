pragma solidity ^0.4.24;

interface IRegistry {
    function START_REGISTRATION_PERIOD() constant public returns(bytes32);
    function TRANSFER_ROOT_NODE() constant public returns(bytes32);
    function initialize(address _ens, bytes32 _rootNode, bytes32 _root);
}
