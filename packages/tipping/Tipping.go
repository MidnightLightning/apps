// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tipping

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TippingABI is the input ABI used to generate the binding from.
const TippingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NONE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_toName\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_ctype\",\"type\":\"uint8\"},{\"name\":\"_cid\",\"type\":\"uint40\"}],\"name\":\"tip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getUsername\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_username\",\"type\":\"bytes32\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"fromName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"toName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ctype\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint40\"}],\"name\":\"Tip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"toName\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"}]"

// TippingBin is the compiled bytecode used for deploying new contracts.
const TippingBin = `60806040526200001d62000023640100000000026401000000009004565b62000309565b60006200003e6200015e640100000000026401000000009004565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a4544000000000000000081525090151562000121576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015620000e5578082015181840152602081019050620000c8565b50505050905090810190601f168015620001135780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506200015c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff620001ab640100000000026401000000009004565b565b6000620001a67febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e60010260001916620002f76401000000000262002323176401000000009004565b905090565b6000620001c66200015e640100000000026401000000009004565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a45440000000000000000815250901515620002a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200026d57808201518184015260208101905062000250565b50505050905090810190601f1680156200029b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50620002f4817febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e60010260001916620003026401000000000262002480179091906401000000009004565b50565b600081549050919050565b8082555050565b6124b380620003196000396000f300608060405260043610610112576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630803fac0146101175780631e83409a146101465780632914b9bd1461018957806332f0a3b514610232578063485cc955146102895780637b103999146102ec5780637e7db6e11461034357806380afdea81461039e57806383525394146103d15780638909aa3f146104045780638b3dd749146104495780639d4941d814610474578063a1658fad146104b7578063a479e50814610563578063b3cd155d146105ba578063ce43c03214610613578063d4aae0c414610672578063de4796ed146106c9578063deb931a2146106f8578063fc0c546a14610769575b600080fd5b34801561012357600080fd5b5061012c6107c0565b604051808215151515815260200191505060405180910390f35b34801561015257600080fd5b50610187600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107eb565b005b34801561019557600080fd5b506101f0600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d11565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561023e57600080fd5b50610247610e30565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561029557600080fd5b506102ea600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610edd565b005b3480156102f857600080fd5b50610301611055565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034f57600080fd5b50610384600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061107b565b604051808215151515815260200191505060405180910390f35b3480156103aa57600080fd5b506103b3611086565b60405180826000191660001916815260200191505060405180910390f35b3480156103dd57600080fd5b506103e66110bd565b60405180826000191660001916815260200191505060405180910390f35b34801561041057600080fd5b5061043360048036038101908080356000191690602001909291905050506110f6565b6040518082815260200191505060405180910390f35b34801561045557600080fd5b5061045e61110e565b6040518082815260200191505060405180910390f35b34801561048057600080fd5b506104b5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611145565b005b3480156104c357600080fd5b50610549600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611577565b604051808215151515815260200191505060405180910390f35b34801561056f57600080fd5b5061057861177e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105c657600080fd5b50610611600480360381019080803560001916906020019092919080359060200190929190803560ff169060200190929190803564ffffffffff169060200190929190505050611899565b005b34801561061f57600080fd5b50610654600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611f77565b60405180826000191660001916815260200191505060405180910390f35b34801561067e57600080fd5b50610687612078565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106d557600080fd5b506106de6120af565b604051808215151515815260200191505060405180910390f35b34801561070457600080fd5b5061072760048036038101908080356000191690602001909291905050506120e0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561077557600080fd5b5061077e6121bd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806107cb61110e565b9050600081141580156107e55750806107e26121e3565b10155b91505090565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663730ecf34846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156108ab57600080fd5b505af11580156108bf573d6000803e3d6000fd5b505050506040513d60208110156108d557600080fd5b810190808051906020019092919050505091506000600102826000191614156040805190810160405280601381526020017f555345525f4e4f545f52454749535445524544000000000000000000000000008152509015156109d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561099757808201518184015260208101905061097c565b50505050905090810190601f1680156109c45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000808360001916600019168152602001908152602001600020549050600081116040805190810160405280601081526020017f4e4f5448494e475f544f5f434c41494d00000000000000000000000000000000815250901515610ad2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a97578082015181840152602081019050610a7c565b50505050905090810190601f168015610ac45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600080836000191660001916815260200190815260200160002060009055600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610bb657600080fd5b505af1158015610bca573d6000803e3d6000fd5b505050506040513d6020811015610be057600080fd5b81019080805190602001909291905050506040805190810160405280602081526020017f46494e414e43455f544b4e5f5452414e534645525f46524f4d5f524556455254815250901515610ccf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c94578082015181840152602081019050610c79565b50505050905090810190601f168015610cc15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5081600019167fac86c2f8a32db75c8fd0ea87c8be8c73c16136f1f4ee544bc56031c6a12d9528826040518082815260200191505060405180910390a2505050565b6000610d1b61177e565b73ffffffffffffffffffffffffffffffffffffffff166304bf2a7f836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610da2578082015181840152602081019050610d87565b50505050905090810190601f168015610dcf5780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610dee57600080fd5b505af1158015610e02573d6000803e3d6000fd5b505050506040513d6020811015610e1857600080fd5b81019080805190602001909291905050509050919050565b6000610e3a612078565b73ffffffffffffffffffffffffffffffffffffffff166332f0a3b56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610e9d57600080fd5b505af1158015610eb1573d6000803e3d6000fd5b505050506040513d6020811015610ec757600080fd5b8101908080519060200190929190505050905090565b6000610ee761110e565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a45440000000000000000815250901515610fc6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f8b578082015181840152602081019050610f70565b50505050905090810190601f168015610fb85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610fcf6121eb565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019050919050565b60006110b87fd625496217aa6a3453eecb9c3489dc5a53e6c67b444329ea2b2cbc9ff547639b60010260001916612318565b905090565b60405180807f4e4f4e45000000000000000000000000000000000000000000000000000000008152506004019050604051809103902081565b60006020528060005260406000206000915090505481565b60006111407febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e60010260001916612323565b905090565b6000806111518361107b565b6040805190810160405280601281526020017f5245434f5645525f444953414c4c4f574544000000000000000000000000000081525090151561122f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156111f45780820151818401526020810190506111d9565b50505050905090810190601f1680156112215780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50611238610e30565b91506112438261232e565b6040805190810160405280601a81526020017f5245434f5645525f5641554c545f4e4f545f434f4e5452414354000000000000815250901515611321576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156112e65780820151818401526020810190506112cb565b50505050905090810190601f1680156113135780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156113ba578173ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f193505050501580156113b4573d6000803e3d6000fd5b50611572565b8273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561145557600080fd5b505af1158015611469573d6000803e3d6000fd5b505050506040513d602081101561147f57600080fd5b810190808051906020019092919050505090508273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561153557600080fd5b505af1158015611549573d6000803e3d6000fd5b505050506040513d602081101561155f57600080fd5b8101908080519060200190929190505050505b505050565b600080606060006115866107c0565b15156115955760009350611774565b61159d612078565b9250600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156115dd5760009350611774565b602085510290508491508082528273ffffffffffffffffffffffffffffffffffffffff1663fdef9106883089866040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001836000191660001916815260200180602001828103825283818151815260200191508051906020019080838360005b838110156116e75780820151818401526020810190506116cc565b50505050905090810190601f1680156117145780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561173657600080fd5b505af115801561174a573d6000803e3d6000fd5b505050506040513d602081101561176057600080fd5b810190808051906020019092919050505093505b5050509392505050565b600080611789612078565b73ffffffffffffffffffffffffffffffffffffffff1663be00bbd87fd6f028ca0e8edb4a8c9757ca4fdccab25fa1e0317da1188108f7d2dee14902fb6001027fddbcfd564f642ab5627cf68b9b7d374fb4f8a36e941a75d89c87998cef03bd616001026040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808360001916600019168152602001826000191660001916815260200192505050602060405180830381600087803b15801561185557600080fd5b505af1158015611869573d6000803e3d6000fd5b505050506040513d602081101561187f57600080fd5b810190808051906020019092919050505090508091505090565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347079892876040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561193557600080fd5b505af1158015611949573d6000803e3d6000fd5b505050506040513d602081101561195f57600080fd5b8101908080519060200190929190505050915060008273ffffffffffffffffffffffffffffffffffffffff161415611bf457600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330886040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611a8a57600080fd5b505af1158015611a9e573d6000803e3d6000fd5b505050506040513d6020811015611ab457600080fd5b81019080805190602001909291905050506040805190810160405280602081526020017f46494e414e43455f544b4e5f5452414e534645525f46524f4d5f524556455254815250901515611ba3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611b68578082015181840152602081019050611b4d565b50505050905090810190601f168015611b955780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50611bd18560008089600019166000191681526020019081526020016000205461238090919063ffffffff16565b600080886000191660001916815260200190815260200160002081905550611e08565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3384886040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611ced57600080fd5b505af1158015611d01573d6000803e3d6000fd5b505050506040513d6020811015611d1757600080fd5b81019080805190602001909291905050506040805190810160405280602081526020017f46494e414e43455f544b4e5f5452414e534645525f46524f4d5f524556455254815250901515611e06576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611dcb578082015181840152602081019050611db0565b50505050905090810190601f168015611df85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663730ecf34336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611ec557600080fd5b505af1158015611ed9573d6000803e3d6000fd5b505050506040513d6020811015611eef57600080fd5b81019080805190602001909291905050509050856000191681600019167f7387928d33732bf7c5ed6a8b97a2dd3a487d5d8c80ff16d0e50c926baeb1e72587878760405180848152602001836002811115611f4657fe5b60ff1681526020018264ffffffffff1664ffffffffff168152602001935050505060405180910390a3505050505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663730ecf34836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561203657600080fd5b505af115801561204a573d6000803e3d6000fd5b505050506040513d602081101561206057600080fd5b81019080805190602001909291905050509050919050565b60006120aa7f4172f0f7d2289153072b0a6ca36959e0cbe2efc3afe50fc81636caa96338137b60010260001916612475565b905090565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6120da61110e565b14905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166347079892836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561217b57600080fd5b505af115801561218f573d6000803e3d6000fd5b505050506040513d60208110156121a557600080fd5b81019080805190602001909291905050509050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600043905090565b60006121f561110e565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a454400000000000000008152509015156122d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561229957808201518184015260208101905061227e565b50505050905090810190601f1680156122c65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506123166122e06121e3565b7febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e6001026000191661248090919063ffffffff16565b565b600081549050919050565b600081549050919050565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561236f576000915061237a565b823b90506000811191505b50919050565b6000808284019050838110156040805190810160405280601181526020017f4d4154485f4144445f4f564552464c4f5700000000000000000000000000000081525090151561246a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561242f578082015181840152602081019050612414565b50505050905090810190601f16801561245c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b508091505092915050565b600081549050919050565b80825550505600a165627a7a72305820126e613029490e0cf8e5492ef4155ae8ddbe4fd3c371f9f55ae808868bd614980029`

// DeployTipping deploys a new Ethereum contract, binding an instance of Tipping to it.
func DeployTipping(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tipping, error) {
	parsed, err := abi.JSON(strings.NewReader(TippingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TippingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tipping{TippingCaller: TippingCaller{contract: contract}, TippingTransactor: TippingTransactor{contract: contract}, TippingFilterer: TippingFilterer{contract: contract}}, nil
}

// Tipping is an auto generated Go binding around an Ethereum contract.
type Tipping struct {
	TippingCaller     // Read-only binding to the contract
	TippingTransactor // Write-only binding to the contract
	TippingFilterer   // Log filterer for contract events
}

// TippingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TippingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TippingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TippingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TippingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TippingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TippingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TippingSession struct {
	Contract     *Tipping          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TippingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TippingCallerSession struct {
	Contract *TippingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TippingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TippingTransactorSession struct {
	Contract     *TippingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TippingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TippingRaw struct {
	Contract *Tipping // Generic contract binding to access the raw methods on
}

// TippingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TippingCallerRaw struct {
	Contract *TippingCaller // Generic read-only contract binding to access the raw methods on
}

// TippingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TippingTransactorRaw struct {
	Contract *TippingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipping creates a new instance of Tipping, bound to a specific deployed contract.
func NewTipping(address common.Address, backend bind.ContractBackend) (*Tipping, error) {
	contract, err := bindTipping(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tipping{TippingCaller: TippingCaller{contract: contract}, TippingTransactor: TippingTransactor{contract: contract}, TippingFilterer: TippingFilterer{contract: contract}}, nil
}

// NewTippingCaller creates a new read-only instance of Tipping, bound to a specific deployed contract.
func NewTippingCaller(address common.Address, caller bind.ContractCaller) (*TippingCaller, error) {
	contract, err := bindTipping(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TippingCaller{contract: contract}, nil
}

// NewTippingTransactor creates a new write-only instance of Tipping, bound to a specific deployed contract.
func NewTippingTransactor(address common.Address, transactor bind.ContractTransactor) (*TippingTransactor, error) {
	contract, err := bindTipping(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TippingTransactor{contract: contract}, nil
}

// NewTippingFilterer creates a new log filterer instance of Tipping, bound to a specific deployed contract.
func NewTippingFilterer(address common.Address, filterer bind.ContractFilterer) (*TippingFilterer, error) {
	contract, err := bindTipping(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TippingFilterer{contract: contract}, nil
}

// bindTipping binds a generic wrapper to an already deployed contract.
func bindTipping(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TippingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tipping *TippingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tipping.Contract.TippingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tipping *TippingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tipping.Contract.TippingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tipping *TippingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tipping.Contract.TippingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tipping *TippingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tipping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tipping *TippingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tipping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tipping *TippingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tipping.Contract.contract.Transact(opts, method, params...)
}

// NONE is a free data retrieval call binding the contract method 0x83525394.
//
// Solidity: function NONE() constant returns(bytes32)
func (_Tipping *TippingCaller) NONE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "NONE")
	return *ret0, err
}

// NONE is a free data retrieval call binding the contract method 0x83525394.
//
// Solidity: function NONE() constant returns(bytes32)
func (_Tipping *TippingSession) NONE() ([32]byte, error) {
	return _Tipping.Contract.NONE(&_Tipping.CallOpts)
}

// NONE is a free data retrieval call binding the contract method 0x83525394.
//
// Solidity: function NONE() constant returns(bytes32)
func (_Tipping *TippingCallerSession) NONE() ([32]byte, error) {
	return _Tipping.Contract.NONE(&_Tipping.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) constant returns(bool)
func (_Tipping *TippingCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "allowRecoverability", token)
	return *ret0, err
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) constant returns(bool)
func (_Tipping *TippingSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Tipping.Contract.AllowRecoverability(&_Tipping.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) constant returns(bool)
func (_Tipping *TippingCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Tipping.Contract.AllowRecoverability(&_Tipping.CallOpts, token)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() constant returns(bytes32)
func (_Tipping *TippingCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "appId")
	return *ret0, err
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() constant returns(bytes32)
func (_Tipping *TippingSession) AppId() ([32]byte, error) {
	return _Tipping.Contract.AppId(&_Tipping.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() constant returns(bytes32)
func (_Tipping *TippingCallerSession) AppId() ([32]byte, error) {
	return _Tipping.Contract.AppId(&_Tipping.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x8909aa3f.
//
// Solidity: function balances(bytes32 ) constant returns(uint256)
func (_Tipping *TippingCaller) Balances(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x8909aa3f.
//
// Solidity: function balances(bytes32 ) constant returns(uint256)
func (_Tipping *TippingSession) Balances(arg0 [32]byte) (*big.Int, error) {
	return _Tipping.Contract.Balances(&_Tipping.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x8909aa3f.
//
// Solidity: function balances(bytes32 ) constant returns(uint256)
func (_Tipping *TippingCallerSession) Balances(arg0 [32]byte) (*big.Int, error) {
	return _Tipping.Contract.Balances(&_Tipping.CallOpts, arg0)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) constant returns(bool)
func (_Tipping *TippingCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "canPerform", _sender, _role, _params)
	return *ret0, err
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) constant returns(bool)
func (_Tipping *TippingSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Tipping.Contract.CanPerform(&_Tipping.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) constant returns(bool)
func (_Tipping *TippingCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Tipping.Contract.CanPerform(&_Tipping.CallOpts, _sender, _role, _params)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) constant returns(address)
func (_Tipping *TippingCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "getEVMScriptExecutor", _script)
	return *ret0, err
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) constant returns(address)
func (_Tipping *TippingSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Tipping.Contract.GetEVMScriptExecutor(&_Tipping.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) constant returns(address)
func (_Tipping *TippingCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Tipping.Contract.GetEVMScriptExecutor(&_Tipping.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() constant returns(address)
func (_Tipping *TippingCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "getEVMScriptRegistry")
	return *ret0, err
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() constant returns(address)
func (_Tipping *TippingSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Tipping.Contract.GetEVMScriptRegistry(&_Tipping.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() constant returns(address)
func (_Tipping *TippingCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Tipping.Contract.GetEVMScriptRegistry(&_Tipping.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() constant returns(uint256)
func (_Tipping *TippingCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "getInitializationBlock")
	return *ret0, err
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() constant returns(uint256)
func (_Tipping *TippingSession) GetInitializationBlock() (*big.Int, error) {
	return _Tipping.Contract.GetInitializationBlock(&_Tipping.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() constant returns(uint256)
func (_Tipping *TippingCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _Tipping.Contract.GetInitializationBlock(&_Tipping.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _username) constant returns(address)
func (_Tipping *TippingCaller) GetOwner(opts *bind.CallOpts, _username [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "getOwner", _username)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _username) constant returns(address)
func (_Tipping *TippingSession) GetOwner(_username [32]byte) (common.Address, error) {
	return _Tipping.Contract.GetOwner(&_Tipping.CallOpts, _username)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _username) constant returns(address)
func (_Tipping *TippingCallerSession) GetOwner(_username [32]byte) (common.Address, error) {
	return _Tipping.Contract.GetOwner(&_Tipping.CallOpts, _username)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() constant returns(address)
func (_Tipping *TippingCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "getRecoveryVault")
	return *ret0, err
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() constant returns(address)
func (_Tipping *TippingSession) GetRecoveryVault() (common.Address, error) {
	return _Tipping.Contract.GetRecoveryVault(&_Tipping.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() constant returns(address)
func (_Tipping *TippingCallerSession) GetRecoveryVault() (common.Address, error) {
	return _Tipping.Contract.GetRecoveryVault(&_Tipping.CallOpts)
}

// GetUsername is a free data retrieval call binding the contract method 0xce43c032.
//
// Solidity: function getUsername(address _owner) constant returns(bytes32)
func (_Tipping *TippingCaller) GetUsername(opts *bind.CallOpts, _owner common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "getUsername", _owner)
	return *ret0, err
}

// GetUsername is a free data retrieval call binding the contract method 0xce43c032.
//
// Solidity: function getUsername(address _owner) constant returns(bytes32)
func (_Tipping *TippingSession) GetUsername(_owner common.Address) ([32]byte, error) {
	return _Tipping.Contract.GetUsername(&_Tipping.CallOpts, _owner)
}

// GetUsername is a free data retrieval call binding the contract method 0xce43c032.
//
// Solidity: function getUsername(address _owner) constant returns(bytes32)
func (_Tipping *TippingCallerSession) GetUsername(_owner common.Address) ([32]byte, error) {
	return _Tipping.Contract.GetUsername(&_Tipping.CallOpts, _owner)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() constant returns(bool)
func (_Tipping *TippingCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "hasInitialized")
	return *ret0, err
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() constant returns(bool)
func (_Tipping *TippingSession) HasInitialized() (bool, error) {
	return _Tipping.Contract.HasInitialized(&_Tipping.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() constant returns(bool)
func (_Tipping *TippingCallerSession) HasInitialized() (bool, error) {
	return _Tipping.Contract.HasInitialized(&_Tipping.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() constant returns(bool)
func (_Tipping *TippingCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "isPetrified")
	return *ret0, err
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() constant returns(bool)
func (_Tipping *TippingSession) IsPetrified() (bool, error) {
	return _Tipping.Contract.IsPetrified(&_Tipping.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() constant returns(bool)
func (_Tipping *TippingCallerSession) IsPetrified() (bool, error) {
	return _Tipping.Contract.IsPetrified(&_Tipping.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() constant returns(address)
func (_Tipping *TippingCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "kernel")
	return *ret0, err
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() constant returns(address)
func (_Tipping *TippingSession) Kernel() (common.Address, error) {
	return _Tipping.Contract.Kernel(&_Tipping.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() constant returns(address)
func (_Tipping *TippingCallerSession) Kernel() (common.Address, error) {
	return _Tipping.Contract.Kernel(&_Tipping.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Tipping *TippingCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Tipping *TippingSession) Registry() (common.Address, error) {
	return _Tipping.Contract.Registry(&_Tipping.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Tipping *TippingCallerSession) Registry() (common.Address, error) {
	return _Tipping.Contract.Registry(&_Tipping.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Tipping *TippingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Tipping *TippingSession) Token() (common.Address, error) {
	return _Tipping.Contract.Token(&_Tipping.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_Tipping *TippingCallerSession) Token() (common.Address, error) {
	return _Tipping.Contract.Token(&_Tipping.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _owner) returns()
func (_Tipping *TippingTransactor) Claim(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Tipping.contract.Transact(opts, "claim", _owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _owner) returns()
func (_Tipping *TippingSession) Claim(_owner common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.Claim(&_Tipping.TransactOpts, _owner)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _owner) returns()
func (_Tipping *TippingTransactorSession) Claim(_owner common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.Claim(&_Tipping.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address _registry) returns()
func (_Tipping *TippingTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _registry common.Address) (*types.Transaction, error) {
	return _Tipping.contract.Transact(opts, "initialize", _token, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address _registry) returns()
func (_Tipping *TippingSession) Initialize(_token common.Address, _registry common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.Initialize(&_Tipping.TransactOpts, _token, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token, address _registry) returns()
func (_Tipping *TippingTransactorSession) Initialize(_token common.Address, _registry common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.Initialize(&_Tipping.TransactOpts, _token, _registry)
}

// Tip is a paid mutator transaction binding the contract method 0xb3cd155d.
//
// Solidity: function tip(bytes32 _toName, uint256 _amount, uint8 _ctype, uint40 _cid) returns()
func (_Tipping *TippingTransactor) Tip(opts *bind.TransactOpts, _toName [32]byte, _amount *big.Int, _ctype uint8, _cid *big.Int) (*types.Transaction, error) {
	return _Tipping.contract.Transact(opts, "tip", _toName, _amount, _ctype, _cid)
}

// Tip is a paid mutator transaction binding the contract method 0xb3cd155d.
//
// Solidity: function tip(bytes32 _toName, uint256 _amount, uint8 _ctype, uint40 _cid) returns()
func (_Tipping *TippingSession) Tip(_toName [32]byte, _amount *big.Int, _ctype uint8, _cid *big.Int) (*types.Transaction, error) {
	return _Tipping.Contract.Tip(&_Tipping.TransactOpts, _toName, _amount, _ctype, _cid)
}

// Tip is a paid mutator transaction binding the contract method 0xb3cd155d.
//
// Solidity: function tip(bytes32 _toName, uint256 _amount, uint8 _ctype, uint40 _cid) returns()
func (_Tipping *TippingTransactorSession) Tip(_toName [32]byte, _amount *big.Int, _ctype uint8, _cid *big.Int) (*types.Transaction, error) {
	return _Tipping.Contract.Tip(&_Tipping.TransactOpts, _toName, _amount, _ctype, _cid)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Tipping *TippingTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Tipping.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Tipping *TippingSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.TransferToVault(&_Tipping.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Tipping *TippingTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.TransferToVault(&_Tipping.TransactOpts, _token)
}

// TippingClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Tipping contract.
type TippingClaimIterator struct {
	Event *TippingClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TippingClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TippingClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TippingClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TippingClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TippingClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TippingClaim represents a Claim event raised by the Tipping contract.
type TippingClaim struct {
	ToName  [32]byte
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0xac86c2f8a32db75c8fd0ea87c8be8c73c16136f1f4ee544bc56031c6a12d9528.
//
// Solidity: event Claim(bytes32 indexed toName, uint256 balance)
func (_Tipping *TippingFilterer) FilterClaim(opts *bind.FilterOpts, toName [][32]byte) (*TippingClaimIterator, error) {

	var toNameRule []interface{}
	for _, toNameItem := range toName {
		toNameRule = append(toNameRule, toNameItem)
	}

	logs, sub, err := _Tipping.contract.FilterLogs(opts, "Claim", toNameRule)
	if err != nil {
		return nil, err
	}
	return &TippingClaimIterator{contract: _Tipping.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0xac86c2f8a32db75c8fd0ea87c8be8c73c16136f1f4ee544bc56031c6a12d9528.
//
// Solidity: event Claim(bytes32 indexed toName, uint256 balance)
func (_Tipping *TippingFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *TippingClaim, toName [][32]byte) (event.Subscription, error) {

	var toNameRule []interface{}
	for _, toNameItem := range toName {
		toNameRule = append(toNameRule, toNameItem)
	}

	logs, sub, err := _Tipping.contract.WatchLogs(opts, "Claim", toNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TippingClaim)
				if err := _Tipping.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TippingScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the Tipping contract.
type TippingScriptResultIterator struct {
	Event *TippingScriptResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TippingScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TippingScriptResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TippingScriptResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TippingScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TippingScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TippingScriptResult represents a ScriptResult event raised by the Tipping contract.
type TippingScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Tipping *TippingFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*TippingScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Tipping.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &TippingScriptResultIterator{contract: _Tipping.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Tipping *TippingFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *TippingScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Tipping.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TippingScriptResult)
				if err := _Tipping.contract.UnpackLog(event, "ScriptResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TippingTipIterator is returned from FilterTip and is used to iterate over the raw logs and unpacked data for Tip events raised by the Tipping contract.
type TippingTipIterator struct {
	Event *TippingTip // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TippingTipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TippingTip)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TippingTip)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TippingTipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TippingTipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TippingTip represents a Tip event raised by the Tipping contract.
type TippingTip struct {
	FromName [32]byte
	ToName   [32]byte
	Amount   *big.Int
	Ctype    uint8
	Cid      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTip is a free log retrieval operation binding the contract event 0x7387928d33732bf7c5ed6a8b97a2dd3a487d5d8c80ff16d0e50c926baeb1e725.
//
// Solidity: event Tip(bytes32 indexed fromName, bytes32 indexed toName, uint256 amount, uint8 ctype, uint40 cid)
func (_Tipping *TippingFilterer) FilterTip(opts *bind.FilterOpts, fromName [][32]byte, toName [][32]byte) (*TippingTipIterator, error) {

	var fromNameRule []interface{}
	for _, fromNameItem := range fromName {
		fromNameRule = append(fromNameRule, fromNameItem)
	}
	var toNameRule []interface{}
	for _, toNameItem := range toName {
		toNameRule = append(toNameRule, toNameItem)
	}

	logs, sub, err := _Tipping.contract.FilterLogs(opts, "Tip", fromNameRule, toNameRule)
	if err != nil {
		return nil, err
	}
	return &TippingTipIterator{contract: _Tipping.contract, event: "Tip", logs: logs, sub: sub}, nil
}

// WatchTip is a free log subscription operation binding the contract event 0x7387928d33732bf7c5ed6a8b97a2dd3a487d5d8c80ff16d0e50c926baeb1e725.
//
// Solidity: event Tip(bytes32 indexed fromName, bytes32 indexed toName, uint256 amount, uint8 ctype, uint40 cid)
func (_Tipping *TippingFilterer) WatchTip(opts *bind.WatchOpts, sink chan<- *TippingTip, fromName [][32]byte, toName [][32]byte) (event.Subscription, error) {

	var fromNameRule []interface{}
	for _, fromNameItem := range fromName {
		fromNameRule = append(fromNameRule, fromNameItem)
	}
	var toNameRule []interface{}
	for _, toNameItem := range toName {
		toNameRule = append(toNameRule, toNameItem)
	}

	logs, sub, err := _Tipping.contract.WatchLogs(opts, "Tip", fromNameRule, toNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TippingTip)
				if err := _Tipping.contract.UnpackLog(event, "Tip", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
