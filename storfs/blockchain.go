package ethofs

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	files "github.com/ipfs/go-ipfs-files"

	cid "github.com/ipfs/go-cid"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

var global_pin_counter = uint32(1) // don't start with 0 as we have +1
var firstload = false              //
var is_loading = false             //

// var pinResponseFlag = false
// var pinResponseCount = uint64(10)
var ethClient *ethclient.Client

// func checkPinResponse(pinNumber uint64) {
// 	if pinNumber >= pinResponseCount {
// 		pinResponseFlag = false
// 	}
// }

// var localPinResponseMapping map[string]bool
var localPinResponseMapping = make(map[string]bool)

func markPinResponse(cid string, pinNumber uint32) {
	localPinResponseMapping[cid] = true
}

// type ExpiredContractErrorObject struct {
// 	err error
// 	Err error
// }

func initializeEthClient() error {
	c, err := ethclient.Dial(ipcLocation)
	if err != nil {
		return err
	}
	ethClient = c

	// address := common.HexToAddress("0xD3b80c611999D46895109d75322494F7A49D742F")

	return nil
}

func handleContract(ethofs_contract *Store, contract_i uint32, is_new_contract bool) error {
	// checks against malicios : verify size

	if contract_i > global_pin_counter { // safety check before func exists
		log.Debug("ethofs - ############################## updating", "global_pin_counter", global_pin_counter)
		global_pin_counter = contract_i
	}

	contract, err := ethofs_contract.GetContract(nil, contract_i)
	log.Debug("ethoFS - contract", "contract", contract, "contract_i", contract_i, "timeleft", int64(contract.CreationTimestamp)+int64(contract.DeployedDaysLength)*60*60*24-time.Now().Unix())
	if err != nil {
		log.Debug("ethoFS - Etho Protocol contract connection error (Contract)", "error", err, "number", contract_i)
		return err
	}

	if contract.HostingReplicationFactor == 0 { // in case contract was deleted, example read
		return errors.New("ethoFS - Removed by User")
	}

	timeleft := int64(contract.CreationTimestamp) + int64(contract.DeployedDaysLength)*60*60*24 - time.Now().Unix()

	// check if expired
	if timeleft < 0 {

		removedPin, err := pinRemove(Ipfs, contract.MainContentHash) // Pin data due to insufficient existing providers
		if err != nil {
			log.Info("ethoFS - expired pin removal error", "hash", contract.MainContentHash, "error", err)
			return err
		} else {
			log.Info("ethoFS - expired pin removal successful", "hash", removedPin)
			return errors.New("ethoFS - Expired Contract Error")
		}

	}

	repFactor := uint64(contract.HostingReplicationFactor)

	cid, err := cid.Parse(contract.MainContentHash)
	if err != nil {
		log.Debug("ethoFS - Etho Protocol contract connection error (CID Verification)", "error", err)
		return err
	}

	log.Debug("logINFO_xy", "cid", cid, "err", err, "is_new_contract", is_new_contract, "check", !is_new_contract)

	if err == nil {
		// cids := scanForCids([]byte(contract.MainContentHash))

		// if !is_new_contract {
		resolvedPath := path.IpfsPath(cid) // Request serialized pin list stored on ethoFS // "/ipfs/b" if bad CID

		ctx, cancelCtx := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelCtx()

		resp, err := Ipfs.Unixfs().Get(ctx, resolvedPath) // nil if not found
		log.Debug("logINFO_xx", "resp", resp, "err", err, "resolvedPath", resolvedPath)
		if err != nil {
			if resp == nil {
				localPinResponseMapping[contract.MainContentHash] = true
			}
			log.Debug("ethoFS - bad CID error", "hash", cid, "error", err)
			return err
		}

		// log.Info("logINFO_xx", "resp", resp, "err", err, "resolvedPath", resolvedPath, "resp DagReader", resp.DagReader, "resp dir", resp.dir, "resp size", resp.size())
		upload_size := uint32(0)

		switch resp := resp.(type) {
		case files.File:
			// maybe there an "easier way to get file size?" DIRECTLY VIA DAGREADER
			log.Debug("uuuuuuu ITS A FILE / SOLOHASH", "resp", resp)
			var file files.File
			file = files.ToFile(resp)
			var r io.Reader = file
			if r == nil {
				return errors.New("ethoFS - Empty File Error")
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(r)
			// sub_cids := scanForCids(buf.Bytes())
			// if len(sub_cids) > 0 {
			// 	cids = sub_cids
			// }

			upload_size = uint32(buf.Len())

		case files.Directory:

			log.Debug("uuuuuuu ITS A DIRECTORY / MULTIHASH", "resp", resp)

			_size, err := resp.Size()
			if err != nil {
				return err
			}

			log.Debug("WE NEED TO CALCULATE DIRECTORY SIZE", "test", _size)

			upload_size = uint32(_size)

		}

		// # verify size
		if upload_size > uint32(contract.ContractStorageUsed) {
			return errors.New("ethoFS - Bad Size Contract")
		} else {
			log.Debug("~~~~~~VALIDTY", "WORKED", "TRUE") // check if this size corresponds!
		}

		// log.Info("~~~~~~bufferSIZE", "size", upload_size, "ContractStorageUsed", contract.ContractStorageUsed, "cids", cids) // check if this size corresponds!

		log.Debug("logINFO", "resp", resp, "cid", cid, "err", err)

		// } else {
		// 	log.Info("ADDING ----------------------------", "NEW", "CONTRACT", "cids", cids)
		// }

		pin := contract.MainContentHash

		// for _, pin := range cids {
		log.Debug("ethoFS - pin request detail", "hash", pin, "number", contract_i)

		pinned := pinSearch(pin, localPinMapping)
		if !pinned {
			log.Debug("ethoFS - pin search error", "error", "the requested pin was not found")
		} else {
			log.Debug("ethoFS - data is pinned to local node", "hash", pin)
		}

		providerCount, err := FindProvs(Node, pin)
		log.Debug("ethofs - providerinfo2", "providerCount", providerCount) // check if this size corresponds!

		if err != nil {
			log.Debug("ethoFS - provider search error", "error", err)
			return err
		}

		log.Debug("ethofs - providerinfo", "providers have this file", providerCount, "factor", (repFactor / uint64(2)), "pinned", pinned) // check if this size corresponds!

		if !pinned && providerCount < (repFactor/uint64(2)) {
			log.Debug("ethofs - ADDING PIN***********")                                                   // check if this size corresponds!
			addedPin, err := pinAdd(Ipfs, pin)                                                            // Pin data due to insufficient existing providers
			log.Info("ethofs - ADDING RESULT***********", "addedPin", addedPin, "error", err, "pin", pin) // check if this size corresponds!
			if err != nil {
				log.Debug("ethoFS - pin add error", "hash", pin, "error", err)
				return err
			} else {
				log.Debug("ethoFS - pin added successfully", "hash", addedPin)

				local_ipfs_storage += upload_size

			}

		} else if pinned && providerCount > (repFactor+(repFactor/uint64(2))) {
			removedPin, err := pinRemove(Ipfs, pin) // Pin data due to insufficient existing providers
			if err != nil {
				log.Debug("ethoFS - pin removal error", "hash", pin, "error", err)
				return err
			} else {
				log.Debug("ethoFS - pin removal successful", "hash", removedPin)
			}
		}

		log.Info("ethofs - [SUCCESS]", map[bool]string{true: "[NEW]", false: "[OLD]"}[is_new_contract], contract_i, "SIZE", upload_size/1000, "CID", contract.MainContentHash, "REP_F", contract.HostingReplicationFactor, "TIME", contract.DeployedDaysLength, "TIME_LEFT", timeleft, "PINCOUNT", global_pin_counter, "local_ipfs_storage in [MB]", local_ipfs_storage/1000000)

		// }

	} else {
		// badly formatted cid on blockchain data
		return errors.New("ethoFS - Badly Formatted CID on Blockchain")
	}

	return nil
}

func BlockListener(blockCommunication chan *types.Block, nodeType string) {
	c := ethClient
	ethofs_contract, err := NewStore(contractControllerAddress, c)
	if err != nil {
		// return err
		log.Debug("ETHOFS", "err", err)
	}
	// returnFlag := false
	log.Info("###STARTING BLOCK LISTENER")
	for {
		select {
		case block := <-blockCommunication:
			log.Info("ethoFS - new block received for processing a", "SYNC_STATUS", is_loading == true, "number", block.Header().Number.Int64(), "txs", len(block.Transactions()))

			if block.Header().Number.Int64()%100 == 0 { // garbage collect every 100 blocks ?
				go gc(Node)
			}

			if len(block.Transactions()) > 0 {
				found_new_contracts := false
				for _, transaction := range block.Transactions() {
					recipient := transaction.To()
					if recipient == nil {
						continue
					} else if *recipient == contractControllerAddress {
						found_new_contracts = true
						break
					}
				}

				if found_new_contracts {
					contracts_last_index, err := ethofs_contract.AllHostingObjectsIndex(nil)
					if err != nil {
						// return err
						log.Debug("ETHOFS", "err", err)
					}
					if global_pin_counter < contracts_last_index {
						// if !is_loading {
						// 	log.Info("ethofs - *** [NEW UPLOAD REQUEST] ***", "RECHECK", "NEED A RECHECK", "global_pin_counter", global_pin_counter, "contracts_last_index", contracts_last_index)
						// 	err := loadContracts(ethofs_contract, global_pin_counter+1, contracts_last_index, true)
						// 	if err != nil {
						// 		log.Info("failed loadContracts", err)
						// 		// return err
						// 	}
						// } else {
						// 	log.Info("ETHOFS", "RECHECK", "NEED WAIT....", global_pin_counter, contracts_last_index)
						// }

						if !is_loading {
							log.Info("ethofs - *** [NEW UPLOAD REQUEST] ***", "RECHECK", "NEED A RECHECK", "global_pin_counter", global_pin_counter, "contracts_last_index", contracts_last_index)
							err := loadContracts(ethofs_contract, global_pin_counter+1, contracts_last_index, true)
							if err != nil {
								log.Info("failed loadContracts", err)
								// return err
							}
						}

					}
				}
			}
		}
	}
}

func loadContracts(ethofs_contract *Store, _from uint32, _to uint32, is_new_contract bool) error {
	is_loading = true

	// ITERATE OVER ALL PINS
	for i := uint32(_from); i <= uint32(_to); i++ {
		err := handleContract(ethofs_contract, i, is_new_contract)

		// local_ipfs_storage/1000000

		if err != nil {
			log.Info("ethofs - [FAILED]", "i", i, "err", err)
			// return err
		}

		if i == _to {
			log.Info("ethofs - ************* SYNCED *************", "VERS", map[bool]string{true: "[NEW]", false: "[OLD]"}[is_new_contract], "global_pin_counter", global_pin_counter)
		}
	}

	updateLocalPinMapping(Ipfs)

	is_loading = false
	return nil
}

// ethofs.ethofs_storage_contract.methods.AllHostingObjectsIndex().call().then((total_c) => {
// 	ethofs.ethofs_storage_contract.methods.GetAllHostingContracts(0,total_c).call().then((res) => {
// 		console.log(res)
// 	})
// })

func initEthofsContracts() error { // initates all contracts at load
	// LOADING CONTRACT
	// is_loading = true
	loadstart := time.Now().Unix()

	c := ethClient
	ethofs_contract, err := NewStore(contractControllerAddress, c)
	if err != nil {
		return err
	}

	// GETTING TOTAL COUNT
	contracts_totalcount, err := ethofs_contract.AllHostingObjectsIndex(nil)
	if err != nil {
		return err
	}

	// GETTING FUll INDEX LIST
	all_contract_indices_bc, err := ethofs_contract.GetAllHostingContracts(nil, 0, contracts_totalcount)
	if err != nil {
		return err
	}

	log.Info("############################ Ipfs #####################", "Ipfs", Ipfs)
	// all_contract_indices := make(map[string]bool)
	// all_contract_indices := list.New()
	all_contract_indices := *new([]uint32)
	all_contract_indices = all_contract_indices_bc

	active_contracts := 0
	inactive_contracts := 0

	for _, i := range all_contract_indices_bc {
		if i != 0 { // CAN do this on or off chain..
			log.Info("############################ +++ #####################", "index", i)

			err := handleContract(ethofs_contract, i, false)
			if err != nil {
				log.Info("ethofs - [FAILED]", "i", i, "err", err)
				// return err
				inactive_contracts += 1
			} else {
				active_contracts += 1
			}

		}
	}
	log.Info("############################ all_contract_indices #####################", "all_contract_indices", all_contract_indices, "active_contracts", active_contracts, "inactive_contracts", inactive_contracts)

	log.Info("############################ INIT START #####################", "all_contract_indices_bc", all_contract_indices_bc)

	log.Info("############################ INIT DONE #####################", "contracts_totalcount", uint64(contracts_totalcount), "global_pin_counter", global_pin_counter)
	// log.Info("############################ INIT TIME #####################", "result:", (time.Now().Unix()-loadstart)/60, "minutes", "total")
	log.Info("############################ INIT TIME #####################", "result:", (time.Now().Unix() - loadstart), "seconds", "total")

	updateLocalPinMapping(Ipfs)

	return nil
}
