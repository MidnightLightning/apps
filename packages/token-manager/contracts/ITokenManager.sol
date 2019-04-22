pragma solidity ^0.4.24;

interface ITokenManager {
    function MINT_ROLE() constant public returns(bytes32);
    function BURN_ROLE() constant public returns(bytes32);
    function initialize(address _token, bool _transferable, uint256 _maxAccountTokens);
    function mint(address _receiver, uint256 _amount);
}
