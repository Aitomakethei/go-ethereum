package ethofs

import (
	//"fmt"

	"os"
	"runtime"

	//"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"

	icore "github.com/ipfs/interface-go-ipfs-core"

	"github.com/ipfs/kubo/core"
	//cid "github.com/ipfs/go-cidutil"
)

var localPinMapping map[string]string
var selfNodeID string
var repFactor = uint64(10) // get's overwritten by user input
var BlockHeight = int(0)
var Ipfs icore.CoreAPI
var Node *core.IpfsNode

var gn_space = 70.00
var mn_space = 38.00
var sn_space = 18.00

var gn_memory = 3.0
var mn_memory = 1.5
var sn_memory = 0.75

var local_ipfs_storage = uint32(0) // use this to track how much node has actually stored locally

// var contractControllerAddress = common.HexToAddress("0xc38B47169950D8A28bC77a6Fa7467464f25ADAFc")
// var contractControllerAddress = common.HexToAddress("0xECA1F2F2D8c7Be150eF28B11627b90d9D932fA69")
// var contractControllerAddress = common.HexToAddress("0xc1EdE5dF7703564e7107a357F1eE440fD54d1D2C")
var contractControllerAddress = common.HexToAddress("0xc1170c1Fe7D13D1fe2aEAa7D8e757F258dFfB8D0")

var mainChannelString = "ethoFSPinningChannel_alpha11"
var defaultDataDir string
var ipcLocation string
var isInitialized = false

func IsInitialized() bool {
	return isInitialized
}

func InitializeEthofs(initFlag bool, configFlag bool, nodeType string, blockCommunication chan *types.Block) {

	checkResources(nodeType)

	localPinMapping = make(map[string]string)

	// initalize default locations
	defaultDataDir = node.DefaultDataDir()
	if runtime.GOOS == "linux" {
		ipcLocation = defaultDataDir + "/geth.ipc"
	} else if runtime.GOOS == "windows" {
		//ipcLocation = defaultDataDir + "\\geth.ipc"
		ipcLocation = "\\\\.\\pipe\\geth.ipc"
	} else if runtime.GOOS == "darwin" {
		ipcLocation = defaultDataDir + "/geth.ipc"
	}

	clientErr := initializeEthClient()
	if clientErr != nil {
		os.Exit(0)
	}

	if initFlag {

		log.Info("Starting ethoFS repo initialization")
		err := initializeEthofsNodeRepo(nodeType)
		if err == nil {
			log.Info("ethoFS repo initialization successful")
		} else {
			log.Warn("ethoFS repo initialization failed")
		}
		os.Exit(0)

	} else if configFlag {

		log.Info("Starting ethoFS repo/node configuration")
		err := initializeEthofsNodeConfig(nodeType)
		if err == nil {
			log.Info("ethoFS configuration successful")
		} else {
			log.Warn("ethoFS configuration failed")
		}
		os.Exit(0)

	} else {
		log.Info("Starting ethoFS node initialization", "type", nodeType)
		Ipfs, Node = initializeEthofsNode(nodeType)
		log.Info("Before Launch1")
		go func() {
			err := initEthofsContracts()
			if err != nil {
				log.Debug("ethoFS - error initiating contracts")
			} else {
				log.Debug("ethoFS - contract initiation successful")
			}
		}()

		// Initialize block listener
		go BlockListener(blockCommunication, nodeType)
		// go ContractListener()
	}
}
