pragma solidity ^0.4.24;

interface IDistribution {
    function START_DISTRIBUTION() constant public returns(bytes32);
    function initialize(address _names, address _currencyManager, address _karmaManager, bytes32 _root);
}
