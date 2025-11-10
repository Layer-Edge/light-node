package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MerkleTreeStorageMetaData contains all meta data concerning the MerkleTreeStorage contract.
var MerkleTreeStorageMetaData = &bind.MetaData{
	ABI: `[{"inputs":[{"internalType":"address","name":"target","type":"address"}],"name":"AddressEmptyCode","type":"error"},{"inputs":[{"internalType":"address","name":"implementation","type":"address"}],"name":"ERC1967InvalidImplementation","type":"error"},{"inputs":[],"name":"ERC1967NonPayable","type":"error"},{"inputs":[],"name":"FailedCall","type":"error"},{"inputs":[],"name":"InvalidInitialization","type":"error"},{"inputs":[],"name":"NotInitializing","type":"error"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"OwnableInvalidOwner","type":"error"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"OwnableUnauthorizedAccount","type":"error"},{"inputs":[],"name":"UUPSUnauthorizedCallContext","type":"error"},{"inputs":[{"internalType":"bytes32","name":"slot","type":"bytes32"}],"name":"UUPSUnsupportedProxiableUUID","type":"error"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint64","name":"version","type":"uint64"}],"name":"Initialized","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"merkleRoot","type":"bytes32"},{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":false,"internalType":"bytes32[]","name":"leaves","type":"bytes32[]"}],"name":"TreeCreated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"implementation","type":"address"}],"name":"Upgraded","type":"event"},{"inputs":[],"name":"UPGRADE_INTERFACE_VERSION","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getAllRoots","outputs":[{"internalType":"bytes32[]","name":"","type":"bytes32[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"merkleRoot","type":"bytes32"}],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getTreeCount","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"merkleRoot","type":"bytes32"}],"name":"getTreeInfo","outputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"leafCount","type":"uint256"},{"internalType":"uint256","name":"createdAt","type":"uint256"},{"internalType":"bytes32[]","name":"leaves","type":"bytes32[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"initialOwner","type":"address"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"proxiableUUID","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"merkleRoot","type":"bytes32"},{"internalType":"bytes32[]","name":"leaves","type":"bytes32[]"}],"name":"storeTree","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"merkleRoot","type":"bytes32"}],"name":"treeExists","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newImplementation","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"upgradeToAndCall","outputs":[],"stateMutability":"payable","type":"function"}]`,
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156041575f5ffd5b506080516123656100685f395f8181610b7c01528181610bd10152610d8b01526123655ff3fe6080604052600436106100c1575f3560e01c8063ad3cb1cc1161007e578063c4d66de811610058578063c4d66de814610244578063deb931a21461026c578063f26f68bf146102a8578063f2fde38b146102e4576100c1565b8063ad3cb1cc146101b4578063bacdb394146101de578063bada8fc41461021a576100c1565b80634f1ef286146100c55780635028f70b146100e157806352d1902d1461010b578063715018a61461013557806384ab4db11461014b5780638da5cb5b1461018a575b5f5ffd5b6100df60048036038101906100da9190611665565b61030c565b005b3480156100ec575f5ffd5b506100f561032b565b60405161010291906116d7565b60405180910390f35b348015610116575f5ffd5b5061011f610337565b60405161012c9190611708565b60405180910390f35b348015610140575f5ffd5b50610149610368565b005b348015610156575f5ffd5b50610171600480360381019061016c919061174b565b61037b565b60405161018194939291906118a0565b60405180910390f35b348015610195575f5ffd5b5061019e610547565b6040516101ab91906118ea565b60405180910390f35b3480156101bf575f5ffd5b506101c861057c565b6040516101d59190611955565b60405180910390f35b3480156101e9575f5ffd5b5061020460048036038101906101ff91906119d2565b6105b5565b6040516102119190611708565b60405180910390f35b348015610225575f5ffd5b5061022e610818565b60405161023b9190611ae6565b60405180910390f35b34801561024f575f5ffd5b5061026a60048036038101906102659190611b06565b61086e565b005b348015610277575f5ffd5b50610292600480360381019061028d919061174b565b6109ef565b60405161029f91906118ea565b60405180910390f35b3480156102b3575f5ffd5b506102ce60048036038101906102c9919061174b565b610ace565b6040516102db9190611b4b565b60405180910390f35b3480156102ef575f5ffd5b5061030a60048036038101906103059190611b06565b610af6565b005b610314610b7a565b61031d82610c60565b6103278282610c6b565b5050565b5f600180549050905090565b5f610340610d89565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b905090565b610370610e10565b6103795f610e97565b565b5f5f5f6060845f5f1b81036103c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bc90611bae565b60405180910390fd5b5f5f8781526020019081526020015f206002015f9054906101000a900460ff16610424576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041b90611c16565b60405180910390fd5b5f5f5f8881526020019081526020015f209050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816001015482600301548360040180805480602002602001604051908101604052809291908181526020015f905b82821015610530578382905f5260205f200180546104a590611c61565b80601f01602080910402602001604051908101604052809291908181526020018280546104d190611c61565b801561051c5780601f106104f35761010080835404028352916020019161051c565b820191905f5260205f20905b8154815290600101906020018083116104ff57829003601f168201915b505050505081526020019060010190610488565b505050509050955095509550955050509193509193565b5f5f610551610f68565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b5f835f5f1b81036105fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f290611bae565b60405180910390fd5b5f8484905011610640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063790611d01565b60405180910390fd5b5f5f8681526020019081526020015f206002015f9054906101000a900460ff16156106a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069790611d8f565b60405180910390fd5b600185908060018154018082558091505060019003905f5260205f20015f90919091909150556040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018585905081526020016001151581526020014281526020018585906107129190611e5e565b8152505f5f8781526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055506060820151816003015560808201518160040190805190602001906107b89291906113ec565b509050503373ffffffffffffffffffffffffffffffffffffffff16857fa56aae2dc75f499b69586042d0b4778cf971ee75f803dedb6b6a568502f7a0478686604051610805929190611faa565b60405180910390a3849150509392505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561086457602002820191905f5260205f20905b815481526020019060010190808311610850575b5050505050905090565b5f610877610f8f565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff161480156108bf5750825b90505f60018367ffffffffffffffff161480156108f257505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015610900575080155b15610937576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315610984576001855f0160086101000a81548160ff0219169083151502179055505b61098d86610fa2565b83156109e7575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516109de9190612021565b60405180910390a15b505050505050565b5f815f5f1b8103610a35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2c90611bae565b60405180910390fd5b5f5f8481526020019081526020015f206002015f9054906101000a900460ff16610a94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8b90611c16565b60405180910390fd5b5f5f8481526020019081526020015f205f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915050919050565b5f5f5f8381526020019081526020015f206002015f9054906101000a900460ff169050919050565b610afe610e10565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b6e575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610b6591906118ea565b60405180910390fd5b610b7781610e97565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff161480610c2757507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610c0e610fb6565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610c5e576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610c68610e10565b50565b8173ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610cd357506040513d601f19601f82011682018060405250810190610cd0919061204e565b60015b610d1457816040517f4c9c8ce3000000000000000000000000000000000000000000000000000000008152600401610d0b91906118ea565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b8114610d7a57806040517faa1d49a4000000000000000000000000000000000000000000000000000000008152600401610d719190611708565b60405180910390fd5b610d848383611009565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610e0e576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610e1861107b565b73ffffffffffffffffffffffffffffffffffffffff16610e36610547565b73ffffffffffffffffffffffffffffffffffffffff1614610e9557610e5961107b565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401610e8c91906118ea565b60405180910390fd5b565b5f610ea0610f68565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f5f610f99611082565b90508091505090565b610faa6110ab565b610fb3816110eb565b50565b5f610fe27f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61116f565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61101282611178565b8173ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a25f8151111561106e576110688282611241565b50611077565b6110766112c1565b5b5050565b5f33905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005f1b905090565b6110b36112fd565b6110e9576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6110f36110ab565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611163575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161115a91906118ea565b60405180910390fd5b61116c81610e97565b50565b5f819050919050565b5f8173ffffffffffffffffffffffffffffffffffffffff163b036111d357806040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526004016111ca91906118ea565b60405180910390fd5b806111ff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61116f565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f5f8473ffffffffffffffffffffffffffffffffffffffff168460405161126a91906120b3565b5f60405180830381855af49150503d805f81146112a2576040519150601f19603f3d011682016040523d82523d5f602084013e6112a7565b606091505b50915091506112b785838361131b565b9250505092915050565b5f3411156112fb576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f611306610f8f565b5f0160089054906101000a900460ff16905090565b6060826113305761132b826113a8565b6113a0565b5f825114801561135657505f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561139857836040517f9996b31500000000000000000000000000000000000000000000000000000000815260040161138f91906118ea565b60405180910390fd5b8190506113a1565b5b9392505050565b5f815111156113ba5780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b828054828255905f5260205f20908101928215611432579160200282015b828111156114315782518290816114219190612260565b509160200191906001019061140a565b5b50905061143f9190611443565b5090565b5b80821115611462575f81816114599190611466565b50600101611444565b5090565b50805461147290611c61565b5f825580601f1061148357506114a0565b601f0160209004905f5260205f209081019061149f91906114a3565b5b50565b5b808211156114ba575f815f9055506001016114a4565b5090565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6114f8826114cf565b9050919050565b611508816114ee565b8114611512575f5ffd5b50565b5f81359050611523816114ff565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61157782611531565b810181811067ffffffffffffffff8211171561159657611595611541565b5b80604052505050565b5f6115a86114be565b90506115b4828261156e565b919050565b5f67ffffffffffffffff8211156115d3576115d2611541565b5b6115dc82611531565b9050602081019050919050565b828183375f83830152505050565b5f611609611604846115b9565b61159f565b9050828152602081018484840111156116255761162461152d565b5b6116308482856115e9565b509392505050565b5f82601f83011261164c5761164b611529565b5b813561165c8482602086016115f7565b91505092915050565b5f5f6040838503121561167b5761167a6114c7565b5b5f61168885828601611515565b925050602083013567ffffffffffffffff8111156116a9576116a86114cb565b5b6116b585828601611638565b9150509250929050565b5f819050919050565b6116d1816116bf565b82525050565b5f6020820190506116ea5f8301846116c8565b92915050565b5f819050919050565b611702816116f0565b82525050565b5f60208201905061171b5f8301846116f9565b92915050565b61172a816116f0565b8114611734575f5ffd5b50565b5f8135905061174581611721565b92915050565b5f602082840312156117605761175f6114c7565b5b5f61176d84828501611737565b91505092915050565b61177f816114ee565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6117e0826117ae565b6117ea81856117b8565b93506117fa8185602086016117c8565b61180381611531565b840191505092915050565b5f61181983836117d6565b905092915050565b5f602082019050919050565b5f61183782611785565b611841818561178f565b9350836020820285016118538561179f565b805f5b8581101561188e578484038952815161186f858261180e565b945061187a83611821565b925060208a01995050600181019050611856565b50829750879550505050505092915050565b5f6080820190506118b35f830187611776565b6118c060208301866116c8565b6118cd60408301856116c8565b81810360608301526118df818461182d565b905095945050505050565b5f6020820190506118fd5f830184611776565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f61192782611903565b611931818561190d565b93506119418185602086016117c8565b61194a81611531565b840191505092915050565b5f6020820190508181035f83015261196d818461191d565b905092915050565b5f5ffd5b5f5ffd5b5f5f83601f84011261199257611991611529565b5b8235905067ffffffffffffffff8111156119af576119ae611975565b5b6020830191508360208202830111156119cb576119ca611979565b5b9250929050565b5f5f5f604084860312156119e9576119e86114c7565b5b5f6119f686828701611737565b935050602084013567ffffffffffffffff811115611a1757611a166114cb565b5b611a238682870161197d565b92509250509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611a61816116f0565b82525050565b5f611a728383611a58565b60208301905092915050565b5f602082019050919050565b5f611a9482611a2f565b611a9e8185611a39565b9350611aa983611a49565b805f5b83811015611ad9578151611ac08882611a67565b9750611acb83611a7e565b925050600181019050611aac565b5085935050505092915050565b5f6020820190508181035f830152611afe8184611a8a565b905092915050565b5f60208284031215611b1b57611b1a6114c7565b5b5f611b2884828501611515565b91505092915050565b5f8115159050919050565b611b4581611b31565b82525050565b5f602082019050611b5e5f830184611b3c565b92915050565b7f496e76616c6964206d65726b6c6520726f6f74000000000000000000000000005f82015250565b5f611b9860138361190d565b9150611ba382611b64565b602082019050919050565b5f6020820190508181035f830152611bc581611b8c565b9050919050565b7f5472656520646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f611c0060138361190d565b9150611c0b82611bcc565b602082019050919050565b5f6020820190508181035f830152611c2d81611bf4565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611c7857607f821691505b602082108103611c8b57611c8a611c34565b5b50919050565b7f4c65616620636f756e74206d7573742062652067726561746572207468616e205f8201527f3000000000000000000000000000000000000000000000000000000000000000602082015250565b5f611ceb60218361190d565b9150611cf682611c91565b604082019050919050565b5f6020820190508181035f830152611d1881611cdf565b9050919050565b7f547265652077697468207468697320726f6f7420616c726561647920657869735f8201527f7473000000000000000000000000000000000000000000000000000000000000602082015250565b5f611d7960228361190d565b9150611d8482611d1f565b604082019050919050565b5f6020820190508181035f830152611da681611d6d565b9050919050565b5f67ffffffffffffffff821115611dc757611dc6611541565b5b602082029050602081019050919050565b5f611dea611de584611dad565b61159f565b90508083825260208201905060208402830185811115611e0d57611e0c611979565b5b835b81811015611e5457803567ffffffffffffffff811115611e3257611e31611529565b5b808601611e3f8982611638565b85526020850194505050602081019050611e0f565b5050509392505050565b5f611e6a368484611dd8565b905092915050565b5f819050919050565b5f611e8683856117b8565b9350611e938385846115e9565b611e9c83611531565b840190509392505050565b5f611eb3848484611e7b565b90509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112611ee457611ee3611ec4565b5b83810192508235915060208301925067ffffffffffffffff821115611f0c57611f0b611ebc565b5b600182023603831315611f2257611f21611ec0565b5b509250929050565b5f602082019050919050565b5f611f41838561178f565b935083602084028501611f5384611e72565b805f5b87811015611f98578484038952611f6d8284611ec8565b611f78868284611ea7565b9550611f8384611f2a565b935060208b019a505050600181019050611f56565b50829750879450505050509392505050565b5f6020820190508181035f830152611fc3818486611f36565b90509392505050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f61200b61200661200184611fcc565b611fe8565b611fd5565b9050919050565b61201b81611ff1565b82525050565b5f6020820190506120345f830184612012565b92915050565b5f8151905061204881611721565b92915050565b5f60208284031215612063576120626114c7565b5b5f6120708482850161203a565b91505092915050565b5f81905092915050565b5f61208d826117ae565b6120978185612079565b93506120a78185602086016117c8565b80840191505092915050565b5f6120be8284612083565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121257fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826120ea565b61212f86836120ea565b95508019841693508086168417925050509392505050565b5f61216161215c612157846116bf565b611fe8565b6116bf565b9050919050565b5f819050919050565b61217a83612147565b61218e61218682612168565b8484546120f6565b825550505050565b5f5f905090565b6121a5612196565b6121b0818484612171565b505050565b5b818110156121d3576121c85f8261219d565b6001810190506121b6565b5050565b601f821115612218576121e9816120c9565b6121f2846120db565b81016020851015612201578190505b61221561220d856120db565b8301826121b5565b50505b505050565b5f82821c905092915050565b5f6122385f198460080261221d565b1980831691505092915050565b5f6122508383612229565b9150826002028217905092915050565b612269826117ae565b67ffffffffffffffff81111561228257612281611541565b5b61228c8254611c61565b6122978282856121d7565b5f60209050601f8311600181146122c8575f84156122b6578287015190505b6122c08582612245565b865550612327565b601f1984166122d6866120c9565b5f5b828110156122fd578489015182556001820191506020850194506020810190506122d8565b8683101561231a5784890151612316601f891682612229565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220c76576c980988dee3cee159b39d0b1788e263c11e0d256507de9c423c1057df264736f6c634300081e0033",
}

// MerkleTreeStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeStorageMetaData.ABI instead.
var MerkleTreeStorageABI = MerkleTreeStorageMetaData.ABI

// MerkleTreeStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeStorageMetaData.Bin instead.
var MerkleTreeStorageBin = MerkleTreeStorageMetaData.Bin

// DeployMerkleTreeStorage deploys a new Ethereum contract, binding an instance of MerkleTreeStorage to it.
func DeployMerkleTreeStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTreeStorage, error) {
	parsed, err := MerkleTreeStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTreeStorage{MerkleTreeStorageCaller: MerkleTreeStorageCaller{contract: contract}, MerkleTreeStorageTransactor: MerkleTreeStorageTransactor{contract: contract}, MerkleTreeStorageFilterer: MerkleTreeStorageFilterer{contract: contract}}, nil
}

// MerkleTreeStorage is an auto generated Go binding around an Ethereum contract.
type MerkleTreeStorage struct {
	MerkleTreeStorageCaller     // Read-only binding to the contract
	MerkleTreeStorageTransactor // Write-only binding to the contract
	MerkleTreeStorageFilterer   // Log filterer for contract events
}

// MerkleTreeStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeStorageSession struct {
	Contract     *MerkleTreeStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MerkleTreeStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeStorageCallerSession struct {
	Contract *MerkleTreeStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MerkleTreeStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeStorageTransactorSession struct {
	Contract     *MerkleTreeStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MerkleTreeStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeStorageRaw struct {
	Contract *MerkleTreeStorage // Generic contract binding to access the raw methods on
}

// MerkleTreeStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeStorageCallerRaw struct {
	Contract *MerkleTreeStorageCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeStorageTransactorRaw struct {
	Contract *MerkleTreeStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTreeStorage creates a new instance of MerkleTreeStorage, bound to a specific deployed contract.
func NewMerkleTreeStorage(address common.Address, backend bind.ContractBackend) (*MerkleTreeStorage, error) {
	contract, err := bindMerkleTreeStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorage{MerkleTreeStorageCaller: MerkleTreeStorageCaller{contract: contract}, MerkleTreeStorageTransactor: MerkleTreeStorageTransactor{contract: contract}, MerkleTreeStorageFilterer: MerkleTreeStorageFilterer{contract: contract}}, nil
}

// NewMerkleTreeStorageCaller creates a new read-only instance of MerkleTreeStorage, bound to a specific deployed contract.
func NewMerkleTreeStorageCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeStorageCaller, error) {
	contract, err := bindMerkleTreeStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageCaller{contract: contract}, nil
}

// NewMerkleTreeStorageTransactor creates a new write-only instance of MerkleTreeStorage, bound to a specific deployed contract.
func NewMerkleTreeStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeStorageTransactor, error) {
	contract, err := bindMerkleTreeStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageTransactor{contract: contract}, nil
}

// NewMerkleTreeStorageFilterer creates a new log filterer instance of MerkleTreeStorage, bound to a specific deployed contract.
func NewMerkleTreeStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeStorageFilterer, error) {
	contract, err := bindMerkleTreeStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageFilterer{contract: contract}, nil
}

// bindMerkleTreeStorage binds a generic wrapper to an already deployed contract.
func bindMerkleTreeStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleTreeStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeStorage *MerkleTreeStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeStorage.Contract.MerkleTreeStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeStorage *MerkleTreeStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.MerkleTreeStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeStorage *MerkleTreeStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.MerkleTreeStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTreeStorage *MerkleTreeStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTreeStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTreeStorage *MerkleTreeStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTreeStorage *MerkleTreeStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MerkleTreeStorage *MerkleTreeStorageSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MerkleTreeStorage.Contract.UPGRADEINTERFACEVERSION(&_MerkleTreeStorage.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MerkleTreeStorage.Contract.UPGRADEINTERFACEVERSION(&_MerkleTreeStorage.CallOpts)
}

// GetAllRoots is a free data retrieval call binding the contract method 0xbada8fc4.
//
// Solidity: function getAllRoots() view returns(bytes32[])
func (_MerkleTreeStorage *MerkleTreeStorageCaller) GetAllRoots(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "getAllRoots")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllRoots is a free data retrieval call binding the contract method 0xbada8fc4.
//
// Solidity: function getAllRoots() view returns(bytes32[])
func (_MerkleTreeStorage *MerkleTreeStorageSession) GetAllRoots() ([][32]byte, error) {
	return _MerkleTreeStorage.Contract.GetAllRoots(&_MerkleTreeStorage.CallOpts)
}

// GetAllRoots is a free data retrieval call binding the contract method 0xbada8fc4.
//
// Solidity: function getAllRoots() view returns(bytes32[])
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) GetAllRoots() ([][32]byte, error) {
	return _MerkleTreeStorage.Contract.GetAllRoots(&_MerkleTreeStorage.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 merkleRoot) view returns(address)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) GetOwner(opts *bind.CallOpts, merkleRoot [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "getOwner", merkleRoot)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 merkleRoot) view returns(address)
func (_MerkleTreeStorage *MerkleTreeStorageSession) GetOwner(merkleRoot [32]byte) (common.Address, error) {
	return _MerkleTreeStorage.Contract.GetOwner(&_MerkleTreeStorage.CallOpts, merkleRoot)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 merkleRoot) view returns(address)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) GetOwner(merkleRoot [32]byte) (common.Address, error) {
	return _MerkleTreeStorage.Contract.GetOwner(&_MerkleTreeStorage.CallOpts, merkleRoot)
}

// GetTreeCount is a free data retrieval call binding the contract method 0x5028f70b.
//
// Solidity: function getTreeCount() view returns(uint256)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) GetTreeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "getTreeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTreeCount is a free data retrieval call binding the contract method 0x5028f70b.
//
// Solidity: function getTreeCount() view returns(uint256)
func (_MerkleTreeStorage *MerkleTreeStorageSession) GetTreeCount() (*big.Int, error) {
	return _MerkleTreeStorage.Contract.GetTreeCount(&_MerkleTreeStorage.CallOpts)
}

// GetTreeCount is a free data retrieval call binding the contract method 0x5028f70b.
//
// Solidity: function getTreeCount() view returns(uint256)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) GetTreeCount() (*big.Int, error) {
	return _MerkleTreeStorage.Contract.GetTreeCount(&_MerkleTreeStorage.CallOpts)
}

// GetTreeInfo is a free data retrieval call binding the contract method 0x84ab4db1.
//
// Solidity: function getTreeInfo(bytes32 merkleRoot) view returns(address owner, uint256 leafCount, uint256 createdAt, bytes[] leaves)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) GetTreeInfo(opts *bind.CallOpts, merkleRoot [32]byte) (struct {
	Owner     common.Address
	LeafCount *big.Int
	CreatedAt *big.Int
	Leaves    [][]byte
}, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "getTreeInfo", merkleRoot)

	outstruct := new(struct {
		Owner     common.Address
		LeafCount *big.Int
		CreatedAt *big.Int
		Leaves    [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LeafCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Leaves = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetTreeInfo is a free data retrieval call binding the contract method 0x84ab4db1.
//
// Solidity: function getTreeInfo(bytes32 merkleRoot) view returns(address owner, uint256 leafCount, uint256 createdAt, bytes[] leaves)
func (_MerkleTreeStorage *MerkleTreeStorageSession) GetTreeInfo(merkleRoot [32]byte) (struct {
	Owner     common.Address
	LeafCount *big.Int
	CreatedAt *big.Int
	Leaves    [][]byte
}, error) {
	return _MerkleTreeStorage.Contract.GetTreeInfo(&_MerkleTreeStorage.CallOpts, merkleRoot)
}

// GetTreeInfo is a free data retrieval call binding the contract method 0x84ab4db1.
//
// Solidity: function getTreeInfo(bytes32 merkleRoot) view returns(address owner, uint256 leafCount, uint256 createdAt, bytes[] leaves)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) GetTreeInfo(merkleRoot [32]byte) (struct {
	Owner     common.Address
	LeafCount *big.Int
	CreatedAt *big.Int
	Leaves    [][]byte
}, error) {
	return _MerkleTreeStorage.Contract.GetTreeInfo(&_MerkleTreeStorage.CallOpts, merkleRoot)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerkleTreeStorage *MerkleTreeStorageSession) Owner() (common.Address, error) {
	return _MerkleTreeStorage.Contract.Owner(&_MerkleTreeStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) Owner() (common.Address, error) {
	return _MerkleTreeStorage.Contract.Owner(&_MerkleTreeStorage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MerkleTreeStorage *MerkleTreeStorageSession) ProxiableUUID() ([32]byte, error) {
	return _MerkleTreeStorage.Contract.ProxiableUUID(&_MerkleTreeStorage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MerkleTreeStorage.Contract.ProxiableUUID(&_MerkleTreeStorage.CallOpts)
}

// TreeExists is a free data retrieval call binding the contract method 0xf26f68bf.
//
// Solidity: function treeExists(bytes32 merkleRoot) view returns(bool)
func (_MerkleTreeStorage *MerkleTreeStorageCaller) TreeExists(opts *bind.CallOpts, merkleRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _MerkleTreeStorage.contract.Call(opts, &out, "treeExists", merkleRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TreeExists is a free data retrieval call binding the contract method 0xf26f68bf.
//
// Solidity: function treeExists(bytes32 merkleRoot) view returns(bool)
func (_MerkleTreeStorage *MerkleTreeStorageSession) TreeExists(merkleRoot [32]byte) (bool, error) {
	return _MerkleTreeStorage.Contract.TreeExists(&_MerkleTreeStorage.CallOpts, merkleRoot)
}

// TreeExists is a free data retrieval call binding the contract method 0xf26f68bf.
//
// Solidity: function treeExists(bytes32 merkleRoot) view returns(bool)
func (_MerkleTreeStorage *MerkleTreeStorageCallerSession) TreeExists(merkleRoot [32]byte) (bool, error) {
	return _MerkleTreeStorage.Contract.TreeExists(&_MerkleTreeStorage.CallOpts, merkleRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _MerkleTreeStorage.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MerkleTreeStorage *MerkleTreeStorageSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.Initialize(&_MerkleTreeStorage.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.Initialize(&_MerkleTreeStorage.TransactOpts, initialOwner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTreeStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerkleTreeStorage *MerkleTreeStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.RenounceOwnership(&_MerkleTreeStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.RenounceOwnership(&_MerkleTreeStorage.TransactOpts)
}

// StoreTree is a paid mutator transaction binding the contract method 0xbacdb394.
//
// Solidity: function storeTree(bytes32 merkleRoot, bytes[] leaves) returns(bytes32)
func (_MerkleTreeStorage *MerkleTreeStorageTransactor) StoreTree(opts *bind.TransactOpts, merkleRoot [32]byte, leaves [][32]byte) (*types.Transaction, error) {
	return _MerkleTreeStorage.contract.Transact(opts, "storeTree", merkleRoot, leaves)
}

// StoreTree is a paid mutator transaction binding the contract method 0xbacdb394.
//
// Solidity: function storeTree(bytes32 merkleRoot, bytes[] leaves) returns(bytes32)
func (_MerkleTreeStorage *MerkleTreeStorageSession) StoreTree(merkleRoot [32]byte, leaves [][32]byte) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.StoreTree(&_MerkleTreeStorage.TransactOpts, merkleRoot, leaves)
}

// StoreTree is a paid mutator transaction binding the contract method 0xbacdb394.
//
// Solidity: function storeTree(bytes32 merkleRoot, bytes[] leaves) returns(bytes32)
func (_MerkleTreeStorage *MerkleTreeStorageTransactorSession) StoreTree(merkleRoot [32]byte, leaves [][32]byte) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.StoreTree(&_MerkleTreeStorage.TransactOpts, merkleRoot, leaves)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MerkleTreeStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerkleTreeStorage *MerkleTreeStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.TransferOwnership(&_MerkleTreeStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.TransferOwnership(&_MerkleTreeStorage.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MerkleTreeStorage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MerkleTreeStorage *MerkleTreeStorageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.UpgradeToAndCall(&_MerkleTreeStorage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MerkleTreeStorage *MerkleTreeStorageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MerkleTreeStorage.Contract.UpgradeToAndCall(&_MerkleTreeStorage.TransactOpts, newImplementation, data)
}

// MerkleTreeStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MerkleTreeStorage contract.
type MerkleTreeStorageInitializedIterator struct {
	Event *MerkleTreeStorageInitialized // Event containing the contract specifics and raw log

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
func (it *MerkleTreeStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeStorageInitialized)
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
		it.Event = new(MerkleTreeStorageInitialized)
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
func (it *MerkleTreeStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeStorageInitialized represents a Initialized event raised by the MerkleTreeStorage contract.
type MerkleTreeStorageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*MerkleTreeStorageInitializedIterator, error) {

	logs, sub, err := _MerkleTreeStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageInitializedIterator{contract: _MerkleTreeStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MerkleTreeStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _MerkleTreeStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeStorageInitialized)
				if err := _MerkleTreeStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) ParseInitialized(log types.Log) (*MerkleTreeStorageInitialized, error) {
	event := new(MerkleTreeStorageInitialized)
	if err := _MerkleTreeStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleTreeStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerkleTreeStorage contract.
type MerkleTreeStorageOwnershipTransferredIterator struct {
	Event *MerkleTreeStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerkleTreeStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeStorageOwnershipTransferred)
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
		it.Event = new(MerkleTreeStorageOwnershipTransferred)
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
func (it *MerkleTreeStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeStorageOwnershipTransferred represents a OwnershipTransferred event raised by the MerkleTreeStorage contract.
type MerkleTreeStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerkleTreeStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerkleTreeStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageOwnershipTransferredIterator{contract: _MerkleTreeStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerkleTreeStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerkleTreeStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeStorageOwnershipTransferred)
				if err := _MerkleTreeStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) ParseOwnershipTransferred(log types.Log) (*MerkleTreeStorageOwnershipTransferred, error) {
	event := new(MerkleTreeStorageOwnershipTransferred)
	if err := _MerkleTreeStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleTreeStorageTreeCreatedIterator is returned from FilterTreeCreated and is used to iterate over the raw logs and unpacked data for TreeCreated events raised by the MerkleTreeStorage contract.
type MerkleTreeStorageTreeCreatedIterator struct {
	Event *MerkleTreeStorageTreeCreated // Event containing the contract specifics and raw log

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
func (it *MerkleTreeStorageTreeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeStorageTreeCreated)
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
		it.Event = new(MerkleTreeStorageTreeCreated)
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
func (it *MerkleTreeStorageTreeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeStorageTreeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeStorageTreeCreated represents a TreeCreated event raised by the MerkleTreeStorage contract.
type MerkleTreeStorageTreeCreated struct {
	MerkleRoot [32]byte
	Owner      common.Address
	Leaves     [][]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTreeCreated is a free log retrieval operation binding the contract event 0xa56aae2dc75f499b69586042d0b4778cf971ee75f803dedb6b6a568502f7a047.
//
// Solidity: event TreeCreated(bytes32 indexed merkleRoot, address indexed owner, bytes[] leaves)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) FilterTreeCreated(opts *bind.FilterOpts, merkleRoot [][32]byte, owner []common.Address) (*MerkleTreeStorageTreeCreatedIterator, error) {

	var merkleRootRule []interface{}
	for _, merkleRootItem := range merkleRoot {
		merkleRootRule = append(merkleRootRule, merkleRootItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MerkleTreeStorage.contract.FilterLogs(opts, "TreeCreated", merkleRootRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageTreeCreatedIterator{contract: _MerkleTreeStorage.contract, event: "TreeCreated", logs: logs, sub: sub}, nil
}

// WatchTreeCreated is a free log subscription operation binding the contract event 0xa56aae2dc75f499b69586042d0b4778cf971ee75f803dedb6b6a568502f7a047.
//
// Solidity: event TreeCreated(bytes32 indexed merkleRoot, address indexed owner, bytes[] leaves)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) WatchTreeCreated(opts *bind.WatchOpts, sink chan<- *MerkleTreeStorageTreeCreated, merkleRoot [][32]byte, owner []common.Address) (event.Subscription, error) {

	var merkleRootRule []interface{}
	for _, merkleRootItem := range merkleRoot {
		merkleRootRule = append(merkleRootRule, merkleRootItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MerkleTreeStorage.contract.WatchLogs(opts, "TreeCreated", merkleRootRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeStorageTreeCreated)
				if err := _MerkleTreeStorage.contract.UnpackLog(event, "TreeCreated", log); err != nil {
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

// ParseTreeCreated is a log parse operation binding the contract event 0xa56aae2dc75f499b69586042d0b4778cf971ee75f803dedb6b6a568502f7a047.
//
// Solidity: event TreeCreated(bytes32 indexed merkleRoot, address indexed owner, bytes[] leaves)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) ParseTreeCreated(log types.Log) (*MerkleTreeStorageTreeCreated, error) {
	event := new(MerkleTreeStorageTreeCreated)
	if err := _MerkleTreeStorage.contract.UnpackLog(event, "TreeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleTreeStorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MerkleTreeStorage contract.
type MerkleTreeStorageUpgradedIterator struct {
	Event *MerkleTreeStorageUpgraded // Event containing the contract specifics and raw log

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
func (it *MerkleTreeStorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleTreeStorageUpgraded)
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
		it.Event = new(MerkleTreeStorageUpgraded)
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
func (it *MerkleTreeStorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleTreeStorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleTreeStorageUpgraded represents a Upgraded event raised by the MerkleTreeStorage contract.
type MerkleTreeStorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MerkleTreeStorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MerkleTreeStorage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeStorageUpgradedIterator{contract: _MerkleTreeStorage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MerkleTreeStorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MerkleTreeStorage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleTreeStorageUpgraded)
				if err := _MerkleTreeStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MerkleTreeStorage *MerkleTreeStorageFilterer) ParseUpgraded(log types.Log) (*MerkleTreeStorageUpgraded, error) {
	event := new(MerkleTreeStorageUpgraded)
	if err := _MerkleTreeStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
