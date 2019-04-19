pragma solidity ^0.4.24;

import "@daonuts/common/contracts/INames.sol";

contract Names {
    /// State
    IPublicResolver public resolver;
    bytes32 public rootNode;

    constructor(IPublicResolver _resolver, bytes32 _rootNode) {

        require(_rootNode != bytes32(0), "NO_ROOT_NODE");

        resolver = _resolver;
        rootNode = _rootNode;
    }

    function ownerOfName(string _username) public view returns (address) {
        return resolver.addr(nameNode(_username));
    }

    function nameOfOwner(address _owner) public view returns (string) {
        return resolver.name(addrNode(_owner));
    }

    function nameNode(string _username) public view returns (bytes32) {
        return keccak256(abi.encodePacked(rootNode, keccak256(_username)));
    }

    function addrNode(address _addr) public view returns (bytes32) {
       return keccak256(abi.encodePacked(rootNode, sha3HexAddress(_addr)));
    }

    /**
     * @dev An optimised function to compute the sha3 of the lower-case
     *      hexadecimal representation of an Ethereum address.
     * @param addr The address to hash
     * @return The SHA3 hash of the lower-case hexadecimal encoding of the
     *         input address.
     */
    function sha3HexAddress(address addr) public pure returns (bytes32 ret) {
        addr;
        ret; // Stop warning us about unused variables
        assembly {
            let lookup := 0x3031323334353637383961626364656600000000000000000000000000000000

            for { let i := 40 } gt(i, 0) { } {
                i := sub(i, 1)
                mstore8(i, byte(and(addr, 0xf), lookup))
                addr := div(addr, 0x10)
                i := sub(i, 1)
                mstore8(i, byte(and(addr, 0xf), lookup))
                addr := div(addr, 0x10)
            }

            ret := keccak256(0, 40)
        }
    }

}
