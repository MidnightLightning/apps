pragma solidity ^0.4.24;

interface IHamburger {
    function PURCHASE_ASSET_ROLE() constant public returns(bytes32);
    function COMMONS_ROLE() constant public returns(bytes32);
    function initialize(address _names, address _currencyManager);
}
