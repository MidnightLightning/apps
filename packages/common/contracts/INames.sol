import "@daonuts/common/contracts/IPublicResolver.sol";

interface INames {
    function resolver() public view returns (IPublicResolver);
    function rootNode() public view returns (bytes32);
    function ownerOfName(string _username) public view returns (address);
    function nameOfOwner(address _owner) public view returns (string);
    function nameNode(string _username) public view returns (bytes32);
    function addrNode(address _addr) public view returns (bytes32);
    function sha3HexAddress(address addr) public pure returns (bytes32 ret);
}
