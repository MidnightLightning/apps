pragma solidity ^0.4.24;

interface IPublicResolver {
    function supportsInterface(bytes4 interfaceID) constant returns (bool);
    function addr(bytes32 node) constant returns (address ret);
    function setAddr(bytes32 node, address addr);
    function name(bytes32 node) public constant returns (string ret);
    function setName(bytes32 node, string name);
}
