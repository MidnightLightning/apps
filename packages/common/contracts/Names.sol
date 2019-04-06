pragma solidity ^0.4.24;

import "@aragon/os/contracts/ens/ENSConstants.sol";
import "@aragon/os/contracts/lib/ens/AbstractENS.sol";
import "@aragon/os/contracts/lib/ens/PublicResolver.sol";

contract Names is ENSConstants {
    /// State
    PublicResolver public resolver;
    /// ENS
    /* bytes32 internal constant DAONUTS_LABEL = keccak256("daonuts"); */
    /* bytes32 internal constant DAONUTS_LABEL = 0x53bf7a5ae2fa6880bad06201387e90063522a09407b9b95effeb2a65d870dd4c; */
    /* bytes32 internal constant DAONUTS_NODE = keccak256(abi.encodePacked(ETH_TLD_NODE, DAONUTS_LABEL)); */
    bytes32 internal constant DAONUTS_NODE = 0xbaa9d81065b9803396ee6ad9faedd650a35f2b9ba9849babde99d4cdbf705a2e;

    function setResolver(address _resolver) internal {
        resolver = PublicResolver(_resolver);
    }

    function ownerOfName(string _username) public view returns (address) {
        return resolver.addr(nameNode(_username));
    }

    function nameOfOwner(address _owner) public view returns (string) {
        return resolver.name(addrNode(_owner));
    }

    function nameNode(string _username) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(DAONUTS_NODE, keccak256(_username)));
    }

    function addrNode(address _addr) public pure returns (bytes32) {
       return keccak256(abi.encodePacked(DAONUTS_NODE, sha3HexAddress(_addr)));
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
