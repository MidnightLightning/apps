pragma solidity ^0.4.24;

contract ITipping {
    function NONE() constant public returns(bytes32);
    function initialize(address _names, address _currency);
}
