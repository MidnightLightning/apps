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
const TippingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"names\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_names\",\"type\":\"address\"},{\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_toName\",\"type\":\"string\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_ctype\",\"type\":\"uint8\"},{\"name\":\"_cid\",\"type\":\"uint40\"}],\"name\":\"tip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NONE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"fromName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"toName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ctype\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint40\"}],\"name\":\"Tip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"toName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"}]"

// TippingBin is the compiled bytecode used for deploying new contracts.
const TippingBin = `60806040526200001d62000023640100000000026401000000009004565b62000309565b60006200003e6200015e640100000000026401000000009004565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a4544000000000000000081525090151562000121576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015620000e5578082015181840152602081019050620000c8565b50505050905090810190601f168015620001135780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506200015c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff620001ab640100000000026401000000009004565b565b6000620001a67febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e60010260001916620002f7640100000000026200237a176401000000009004565b905090565b6000620001c66200015e640100000000026401000000009004565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a45440000000000000000815250901515620002a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200026d57808201518184015260208101905062000250565b50505050905090810190601f1680156200029b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50620002f4817febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e6001026000191662000302640100000000026200269c179091906401000000009004565b50565b600081549050919050565b8082555050565b61275680620003196000396000f3006080604052600436106100fc576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063056da048146101015780630803fac0146101585780631e83409a146101875780632914b9bd146101ca57806332f0a3b514610273578063485cc955146102ca5780634c6e58fb1461032d5780637e7db6e11461039057806380afdea8146103eb578063835253941461041e5780638909aa3f146104515780638b3dd749146104965780639d4941d8146104c1578063a1658fad14610504578063a479e508146105b0578063d4aae0c414610607578063de4796ed1461065e578063e5a6b10f1461068d575b600080fd5b34801561010d57600080fd5b506101166106e4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016457600080fd5b5061016d61070a565b604051808215151515815260200191505060405180910390f35b34801561019357600080fd5b506101c8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610735565b005b3480156101d657600080fd5b50610231600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d7b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561027f57600080fd5b50610288610e9a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102d657600080fd5b5061032b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f47565b005b34801561033957600080fd5b5061038e60048036038101908080359060200190820180359060200191909192939192939080359060200190929190803560ff169060200190929190803564ffffffffff1690602001909291905050506110bf565b005b34801561039c57600080fd5b506103d1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506118ad565b604051808215151515815260200191505060405180910390f35b3480156103f757600080fd5b506104006118b8565b60405180826000191660001916815260200191505060405180910390f35b34801561042a57600080fd5b506104336118ef565b60405180826000191660001916815260200191505060405180910390f35b34801561045d57600080fd5b506104806004803603810190808035600019169060200190929190505050611928565b6040518082815260200191505060405180910390f35b3480156104a257600080fd5b506104ab611940565b6040518082815260200191505060405180910390f35b3480156104cd57600080fd5b50610502600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611977565b005b34801561051057600080fd5b50610596600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803560001916906020019092919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050611d95565b604051808215151515815260200191505060405180910390f35b3480156105bc57600080fd5b506105c5611f9c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561061357600080fd5b5061061c6120b7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561066a57600080fd5b506106736120ee565b604051808215151515815260200191505060405180910390f35b34801561069957600080fd5b506106a261211f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080610715611940565b90506000811415801561072f57508061072c612145565b10155b91505090565b6060600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637cb7acf7856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b1580156107f757600080fd5b505af115801561080b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561083557600080fd5b81019080805164010000000081111561084d57600080fd5b8281019050602081018481111561086357600080fd5b815185600182028301116401000000008211171561088057600080fd5b505092919050505092506000835114156040805190810160405280601381526020017f555345525f4e4f545f524547495354455245440000000000000000000000000081525090151561096e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610933578082015181840152602081019050610918565b50505050905090810190601f1680156109605780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50826040518082805190602001908083835b6020831015156109a55780518252602082019150602081019050602083039250610980565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902091506000808360001916600019168152602001908152602001600020549050600081116040805190810160405280601081526020017f4e4f5448494e475f544f5f434c41494d00000000000000000000000000000000815250901515610ad3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a98578082015181840152602081019050610a7d565b50505050905090810190601f168015610ac55780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600080836000191660001916815260200190815260200160002060009055600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610bb757600080fd5b505af1158015610bcb573d6000803e3d6000fd5b505050506040513d6020811015610be157600080fd5b81019080805190602001909291905050506040805190810160405280602081526020017f46494e414e43455f544b4e5f5452414e534645525f46524f4d5f524556455254815250901515610cd0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610c95578082015181840152602081019050610c7a565b50505050905090810190601f168015610cc25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b507f990b6cd9495ee903483cf910ae467cbd572f7f4ae1fbc9e81bd623a1377ca2af83826040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610d3a578082015181840152602081019050610d1f565b50505050905090810190601f168015610d675780820380516001836020036101000a031916815260200191505b50935050505060405180910390a150505050565b6000610d85611f9c565b73ffffffffffffffffffffffffffffffffffffffff166304bf2a7f836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e0c578082015181840152602081019050610df1565b50505050905090810190601f168015610e395780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015610e5857600080fd5b505af1158015610e6c573d6000803e3d6000fd5b505050506040513d6020811015610e8257600080fd5b81019080805190602001909291905050509050919050565b6000610ea46120b7565b73ffffffffffffffffffffffffffffffffffffffff166332f0a3b56040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610f0757600080fd5b505af1158015610f1b573d6000803e3d6000fd5b505050506040513d6020811015610f3157600080fd5b8101908080519060200190929190505050905090565b6000610f51611940565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a45440000000000000000815250901515611030576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610ff5578082015181840152602081019050610fda565b50505050905090810190601f1680156110225780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061103961214d565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000806060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637f87374989896040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825284848281815260200192508082843782019150509350505050602060405180830381600087803b15801561116f57600080fd5b505af1158015611183573d6000803e3d6000fd5b505050506040513d602081101561119957600080fd5b8101908080519060200190929190505050925060008373ffffffffffffffffffffffffffffffffffffffff16141561144d5787876040518083838082843782019150509250505060405180910390209150600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330896040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156112e357600080fd5b505af11580156112f7573d6000803e3d6000fd5b505050506040513d602081101561130d57600080fd5b81019080805190602001909291905050506040805190810160405280602081526020017f46494e414e43455f544b4e5f5452414e534645525f46524f4d5f5245564552548152509015156113fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156113c15780820151818401526020810190506113a6565b50505050905090810190601f1680156113ee5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061142a8660008085600019166000191681526020019081526020016000205461227a90919063ffffffff16565b600080846000191660001916815260200190815260200160002081905550611661565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3385896040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561154657600080fd5b505af115801561155a573d6000803e3d6000fd5b505050506040513d602081101561157057600080fd5b81019080805190602001909291905050506040805190810160405280602081526020017f46494e414e43455f544b4e5f5452414e534645525f46524f4d5f52455645525481525090151561165f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611624578082015181840152602081019050611609565b50505050905090810190601f1680156116515780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637cb7acf7336040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561171e57600080fd5b505af1158015611732573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561175c57600080fd5b81019080805164010000000081111561177457600080fd5b8281019050602081018481111561178a57600080fd5b81518560018202830111640100000000821117156117a757600080fd5b505092919050505090507f2bd10d551f42f2d84ed70a813ececdf3969e370789ab544ea13b401fbbdc32c58189898989896040518080602001806020018681526020018560028111156117f657fe5b60ff1681526020018464ffffffffff1664ffffffffff168152602001838103835289818151815260200191508051906020019080838360005b8381101561184a57808201518184015260208101905061182f565b50505050905090810190601f1680156118775780820380516001836020036101000a031916815260200191505b50838103825288888281815260200192508082843782019150509850505050505050505060405180910390a15050505050505050565b600060019050919050565b60006118ea7fd625496217aa6a3453eecb9c3489dc5a53e6c67b444329ea2b2cbc9ff547639b6001026000191661236f565b905090565b60405180807f4e4f4e45000000000000000000000000000000000000000000000000000000008152506004019050604051809103902081565b60006020528060005260406000206000915090505481565b60006119727febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e6001026000191661237a565b905090565b6000806000611985846118ad565b6040805190810160405280601281526020017f5245434f5645525f444953414c4c4f5745440000000000000000000000000000815250901515611a63576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611a28578082015181840152602081019050611a0d565b50505050905090810190601f168015611a555780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50611a6c610e9a565b9250611a7783612385565b6040805190810160405280601a81526020017f5245434f5645525f5641554c545f4e4f545f434f4e5452414354000000000000815250901515611b55576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611b1a578082015181840152602081019050611aff565b50505050905090810190601f168015611b475780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415611bf1573073ffffffffffffffffffffffffffffffffffffffff163191508273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015611beb573d6000803e3d6000fd5b50611d2a565b839050611c1d308273ffffffffffffffffffffffffffffffffffffffff166123d790919063ffffffff16565b9150611c4a83838373ffffffffffffffffffffffffffffffffffffffff166125b19092919063ffffffff16565b6040805190810160405280601d81526020017f5245434f5645525f544f4b454e5f5452414e534645525f4641494c4544000000815250901515611d28576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611ced578082015181840152602081019050611cd2565b50505050905090810190601f168015611d1a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02846040518082815260200191505060405180910390a350505050565b60008060606000611da461070a565b1515611db35760009350611f92565b611dbb6120b7565b9250600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611dfb5760009350611f92565b602085510290508491508082528273ffffffffffffffffffffffffffffffffffffffff1663fdef9106883089866040518563ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001836000191660001916815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611f05578082015181840152602081019050611eea565b50505050905090810190601f168015611f325780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015611f5457600080fd5b505af1158015611f68573d6000803e3d6000fd5b505050506040513d6020811015611f7e57600080fd5b810190808051906020019092919050505093505b5050509392505050565b600080611fa76120b7565b73ffffffffffffffffffffffffffffffffffffffff1663be00bbd87fd6f028ca0e8edb4a8c9757ca4fdccab25fa1e0317da1188108f7d2dee14902fb6001027fddbcfd564f642ab5627cf68b9b7d374fb4f8a36e941a75d89c87998cef03bd616001026040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808360001916600019168152602001826000191660001916815260200192505050602060405180830381600087803b15801561207357600080fd5b505af1158015612087573d6000803e3d6000fd5b505050506040513d602081101561209d57600080fd5b810190808051906020019092919050505090508091505090565b60006120e97f4172f0f7d2289153072b0a6ca36959e0cbe2efc3afe50fc81636caa96338137b60010260001916612691565b905090565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff612119611940565b14905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600043905090565b6000612157611940565b146040805190810160405280601881526020017f494e49545f414c52454144595f494e495449414c495a45440000000000000000815250901515612236576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156121fb5780820151818401526020810190506121e0565b50505050905090810190601f1680156122285780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50612278612242612145565b7febb05b386a8d34882b8711d156f463690983dc47815980fb82aeeff1aa43579e6001026000191661269c90919063ffffffff16565b565b6000808284019050838110156040805190810160405280601181526020017f4d4154485f4144445f4f564552464c4f57000000000000000000000000000000815250901515612364576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561232957808201518184015260208101905061230e565b50505050905090810190601f1680156123565780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b508091505092915050565b600081549050919050565b600081549050919050565b600080600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156123c657600091506123d1565b823b90506000811191505b50919050565b600060606000808573ffffffffffffffffffffffffffffffffffffffff166370a0823190507c01000000000000000000000000000000000000000000000000000000000285604051602401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505092506124c186846126a3565b91509150816040805190810160405280601c81526020017f534146455f4552435f32305f42414c414e43455f5245564552544544000000008152509015156125a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561256957808201518184015260208101905061254e565b50505050905090810190601f1680156125965780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5080935050505092915050565b6000606063a9059cbb7c0100000000000000000000000000000000000000000000000000000000028484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905061268785826126d8565b9150509392505050565b600081549050919050565b8082555050565b6000806000806040516020818751602089018a5afa925060008311156126c857805191505b5081819350935050509250929050565b6000806040516020818551602087016000895af1600081111561271e573d6000811461270b57602081146127145761271c565b6001935061271c565b600183511493505b505b505080915050929150505600a165627a7a72305820fbcb951308d85b8fd25bda5255429d0ea3d1369df35b759f597cbd8f1154c7450029`

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

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(address)
func (_Tipping *TippingCaller) Currency(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "currency")
	return *ret0, err
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(address)
func (_Tipping *TippingSession) Currency() (common.Address, error) {
	return _Tipping.Contract.Currency(&_Tipping.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(address)
func (_Tipping *TippingCallerSession) Currency() (common.Address, error) {
	return _Tipping.Contract.Currency(&_Tipping.CallOpts)
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

// Names is a free data retrieval call binding the contract method 0x056da048.
//
// Solidity: function names() constant returns(address)
func (_Tipping *TippingCaller) Names(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tipping.contract.Call(opts, out, "names")
	return *ret0, err
}

// Names is a free data retrieval call binding the contract method 0x056da048.
//
// Solidity: function names() constant returns(address)
func (_Tipping *TippingSession) Names() (common.Address, error) {
	return _Tipping.Contract.Names(&_Tipping.CallOpts)
}

// Names is a free data retrieval call binding the contract method 0x056da048.
//
// Solidity: function names() constant returns(address)
func (_Tipping *TippingCallerSession) Names() (common.Address, error) {
	return _Tipping.Contract.Names(&_Tipping.CallOpts)
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
// Solidity: function initialize(address _names, address _currency) returns()
func (_Tipping *TippingTransactor) Initialize(opts *bind.TransactOpts, _names common.Address, _currency common.Address) (*types.Transaction, error) {
	return _Tipping.contract.Transact(opts, "initialize", _names, _currency)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _names, address _currency) returns()
func (_Tipping *TippingSession) Initialize(_names common.Address, _currency common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.Initialize(&_Tipping.TransactOpts, _names, _currency)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _names, address _currency) returns()
func (_Tipping *TippingTransactorSession) Initialize(_names common.Address, _currency common.Address) (*types.Transaction, error) {
	return _Tipping.Contract.Initialize(&_Tipping.TransactOpts, _names, _currency)
}

// Tip is a paid mutator transaction binding the contract method 0x4c6e58fb.
//
// Solidity: function tip(string _toName, uint256 _amount, uint8 _ctype, uint40 _cid) returns()
func (_Tipping *TippingTransactor) Tip(opts *bind.TransactOpts, _toName string, _amount *big.Int, _ctype uint8, _cid *big.Int) (*types.Transaction, error) {
	return _Tipping.contract.Transact(opts, "tip", _toName, _amount, _ctype, _cid)
}

// Tip is a paid mutator transaction binding the contract method 0x4c6e58fb.
//
// Solidity: function tip(string _toName, uint256 _amount, uint8 _ctype, uint40 _cid) returns()
func (_Tipping *TippingSession) Tip(_toName string, _amount *big.Int, _ctype uint8, _cid *big.Int) (*types.Transaction, error) {
	return _Tipping.Contract.Tip(&_Tipping.TransactOpts, _toName, _amount, _ctype, _cid)
}

// Tip is a paid mutator transaction binding the contract method 0x4c6e58fb.
//
// Solidity: function tip(string _toName, uint256 _amount, uint8 _ctype, uint40 _cid) returns()
func (_Tipping *TippingTransactorSession) Tip(_toName string, _amount *big.Int, _ctype uint8, _cid *big.Int) (*types.Transaction, error) {
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
	ToName  string
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x990b6cd9495ee903483cf910ae467cbd572f7f4ae1fbc9e81bd623a1377ca2af.
//
// Solidity: event Claim(string toName, uint256 balance)
func (_Tipping *TippingFilterer) FilterClaim(opts *bind.FilterOpts) (*TippingClaimIterator, error) {

	logs, sub, err := _Tipping.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &TippingClaimIterator{contract: _Tipping.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x990b6cd9495ee903483cf910ae467cbd572f7f4ae1fbc9e81bd623a1377ca2af.
//
// Solidity: event Claim(string toName, uint256 balance)
func (_Tipping *TippingFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *TippingClaim) (event.Subscription, error) {

	logs, sub, err := _Tipping.contract.WatchLogs(opts, "Claim")
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

// TippingRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the Tipping contract.
type TippingRecoverToVaultIterator struct {
	Event *TippingRecoverToVault // Event containing the contract specifics and raw log

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
func (it *TippingRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TippingRecoverToVault)
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
		it.Event = new(TippingRecoverToVault)
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
func (it *TippingRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TippingRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TippingRecoverToVault represents a RecoverToVault event raised by the Tipping contract.
type TippingRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Tipping *TippingFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*TippingRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tipping.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TippingRecoverToVaultIterator{contract: _Tipping.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Tipping *TippingFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *TippingRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tipping.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TippingRecoverToVault)
				if err := _Tipping.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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
	FromName string
	ToName   string
	Amount   *big.Int
	Ctype    uint8
	Cid      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTip is a free log retrieval operation binding the contract event 0x2bd10d551f42f2d84ed70a813ececdf3969e370789ab544ea13b401fbbdc32c5.
//
// Solidity: event Tip(string fromName, string toName, uint256 amount, uint8 ctype, uint40 cid)
func (_Tipping *TippingFilterer) FilterTip(opts *bind.FilterOpts) (*TippingTipIterator, error) {

	logs, sub, err := _Tipping.contract.FilterLogs(opts, "Tip")
	if err != nil {
		return nil, err
	}
	return &TippingTipIterator{contract: _Tipping.contract, event: "Tip", logs: logs, sub: sub}, nil
}

// WatchTip is a free log subscription operation binding the contract event 0x2bd10d551f42f2d84ed70a813ececdf3969e370789ab544ea13b401fbbdc32c5.
//
// Solidity: event Tip(string fromName, string toName, uint256 amount, uint8 ctype, uint40 cid)
func (_Tipping *TippingFilterer) WatchTip(opts *bind.WatchOpts, sink chan<- *TippingTip) (event.Subscription, error) {

	logs, sub, err := _Tipping.contract.WatchLogs(opts, "Tip")
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
