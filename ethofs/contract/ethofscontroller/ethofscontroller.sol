// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// TO_DO
// ADD ORACLE PRICE VIA CONTRACT DIRECT CALL !!!!!!!!
// DELETE FILE IF EXPIRED -- expiration dict?
// HOSTING IDENTIFIER - add for deletion -- not needed if we have unsorted indexes
// create cost for deletion so contract isn't SPAMMED by bad actors ??
// use exact bytesize for cid ?
// how to handle marked for deletion ?



// using identifier as any point ? or could we also just use fileID?

contract EthoFSController {
	uint32 public contractsTotalCount; // counts all contracts
	uint64 public contractsTotalSize; // adds up all contract sizes together
	uint128 public hostingCost_perMB_perYEAR; // static price per mb per year ( should be dynamically set by oracle in USDT ratio )

	address public owner; // owner of ethofs controller
	address public operator; // owner of ethofs controller
	address public collection_address; // Address to collect all Storage Payments

	uint16 public MaximumStorageTime; // maximum time in days for storage contracts to ever run for
	uint8 public MaximumDeletionCount; // each person who creates a new contract should never pay for deletion of more than 10 contracts
	uint8 public MinimumReplicationFactor; // each person who creates a new contract should never pay for deletion of more than 10 contracts

	uint32 public fake_day_cycle;
	uint32 public day_cycle;

	// uint16 public DAILY_SHIFT_THRESHHOLD;

	uint32 public controllerCreationTime; // time reference used for deletion dict

	// mapping(uint16 => ExpiryTracker) public ExpirationDict; // dict to track contracts for deletion
	mapping(uint16 => uint32[]) public ExpirationDict; // dict to track contracts for deletion
	mapping(uint16 => uint32) public ExpirationDict_Indexer; // dict to track contracts for deletion
	// mapping (uint16 => DeleteObject) public ExpirationDict;
	// mapping(uint => bool) public allActiveHostingIndices;
	mapping(uint32 => uint32) public fileID_to_index_allActiveHostingIndices; // mapping global fileID to list index
	uint32[] free_indices_allActiveHostingIndices; // global list carrying all freed indices
	uint32[] allActiveHostingIndices; // global list carrying all [ACTIVE] hosting indices
	// uint[] public allInactiveHostingIndices;
	uint16 last_date_check_ExpirationDict;
	uint32 last_index_check_ExpirationDict;

	event OnFundCollection(address indexed sender, uint256 amount);
	event OnNewContractHost(address newHost);
	event OnCreateContract(uint32 contract_id);
	event OnRemoveContract(uint32 contract_id);
	event OnExtendContract(uint32 contract_id, uint32 extendedDaysLength);
	event OnContractExpired(uint32 contract_id);
	event OnChangedOwnership(uint32 contract_id);

	// struct ExpiryTracker {
	// 	uint8 hostingReplicationFactor; // how many times files should be replicated
	// 	uint16 deployedDaysLength;
	// 	uint32[] active_indices; // 4294967295 files
	// }

	// struct DeleteObject {
	// 	uint32 local_index;
	// 	uint32[] all_indices;
	// }

	struct HostingObject {
		uint8 hostingReplicationFactor; // how many times files should be replicated
		uint16 deployedDaysLength;
		uint32 contractStorageUsed;
		uint32 creationTimestamp;
		// uint128 contractPayAmount; // use this to track that uploaded file is not a spoofed file size upload ?
		string mainContentHash;
		string hostingContractName;
		address ownerAddress;
	}

	struct UserStorageAccount {
		uint32 numHostingObjects; //4294967295, 16 bit should also work 65536
		uint64 totalHostingSize;
		uint32[] active_indices; // 4294967295 files
		uint32[] free_indices; // 4294967295 files
		// uint32 active_index; // 4294967295 files
		mapping(uint32 => uint32) fileID_to_Index;
		// mapping(uint32 => uint32) active_indices;
	}

	// mapping (uint32 => UserStorageAccount) public users;
	mapping (uint160 => UserStorageAccount) public users;

	uint32 public allHostingObjectsIndex;
	mapping (uint32 => HostingObject) allHostingObjects;

	// MAIN FUNCTIONS
	function addContracts(string[] memory CID_LIST,string[] memory NAME_LIST,uint8[] memory REPLICATION_LIST,uint32[] memory SIZE_LIST,uint16[] memory DURATION_LIST) public payable returns ( uint32[] memory ) {
		// check integrity
		// require(CID_LIST.length < 2, "Malformed Array sizes");
		require(CID_LIST.length < 256, "Exceeded Maximum Size");
		//success_ids = new uint8[](0);
		uint32[] memory success_ids = new uint32[](0);

		require(CID_LIST.length == NAME_LIST.length, "Malformed Array sizes");
		for (uint8 i=0; i<CID_LIST.length; i++) {
			//success_ids.push(addContract(CID_LIST[i],NAME_LIST[i],REPLICATION_LIST[i],SIZE_LIST[i],DURATION_LIST[i]));
			success_ids[i] = addContract(CID_LIST[i],NAME_LIST[i],REPLICATION_LIST[i],SIZE_LIST[i],DURATION_LIST[i]);
		}
		return success_ids;
	}

	function addContract(string memory CID, string memory ContractName, uint8 ReplicationFactor, uint32 ContractSize, uint16 ContractDuration) public payable returns (uint32 fileID) { // pay to add a file to storage. [CID,CONTRACT_NAME,size in BYTES,DURATION in days]
		// handleExpiredContracts(); // scrub old contracts
		if (MaximumDeletionCount > 0) {
			wipeContracts(MaximumDeletionCount); // always make users delete at least 2 contracts?
		}

		// require(addcount > 0, "Must Define Contract Count"); // used to calc cost for multi add correctly
		require(ReplicationFactor >= MinimumReplicationFactor, "ReplicationFactor Minimum"); 
		require(ContractDuration > 0, "Bad Contract Duration");
		require(ContractDuration <= MaximumStorageTime, "Bad Storage Duration");
		require(ContractSize > 0, "Content Size Too Small");
		require(ContractSize <= 4294967295, "Content Size Too Big");

		uint256 _required_amt = 0; //GetHostingCost(ContractSize,ContractDuration,ReplicationFactor);
		require(msg.value >= _required_amt, "Did not Pay Enough for Contract Creation");
		require(bytes(CID).length > 0, "Bad CID"); // maybe use exact bytesize for cid ?
		UserStorageAccount storage c = users[uint160(msg.sender)];

		if (c.numHostingObjects == 0) {
			emit OnNewContractHost(msg.sender);
		}

		// added size to user and global vars
		contractsTotalSize += ContractSize;
		contractsTotalCount++;
		c.totalHostingSize += ContractSize;
		c.numHostingObjects++;
		
		allHostingObjectsIndex++;

		if (c.free_indices.length > 0) { // if we have a "free slot" let's reuse it
			uint32 free_index = c.free_indices[c.free_indices.length-1]; // temp placeholder for freeindex
			c.fileID_to_Index[allHostingObjectsIndex] = free_index;
			c.active_indices[free_index] = allHostingObjectsIndex;

			c.free_indices.pop(); // pop free_index

		} else {
			c.fileID_to_Index[allHostingObjectsIndex] = uint32(c.active_indices.length);
			c.active_indices.push(allHostingObjectsIndex);
		}
		// c.active_index ++;
		// c.active_indices[c.active_index] = allHostingObjectsIndex;


		if (free_indices_allActiveHostingIndices.length > 0) { // if we have a "free slot" let's reuse it
			uint32 free_index = free_indices_allActiveHostingIndices[free_indices_allActiveHostingIndices.length-1]; // temp placeholder for freeindex
			fileID_to_index_allActiveHostingIndices[allHostingObjectsIndex] = free_index;
			allActiveHostingIndices[free_index] = allHostingObjectsIndex;

			free_indices_allActiveHostingIndices.pop(); // pop free_index
		} else {
			fileID_to_index_allActiveHostingIndices[allHostingObjectsIndex] = uint32(allActiveHostingIndices.length);
			allActiveHostingIndices.push(allHostingObjectsIndex);
		}
		




		// uint expiry_day = ((block.timestamp-controllerCreationTime)/day_cycle)+ContractDuration;
		uint16 expiry_day = uint16((block.timestamp-controllerCreationTime)/day_cycle)+ContractDuration;

		if (ExpirationDict[expiry_day].length == 0) { 
			ExpirationDict[expiry_day] = new uint32[](0);
			// ExpirationDict[expiry_day] = [allHostingObjectsIndex];
			ExpirationDict_Indexer[expiry_day] = 0;
			// ExpirationDict[expiry_day] = DeleteObject({local_index:0,all_indices:uint32[]});
			// ExpirationDict[expiry_day] = DeleteObject({local_index:0,all_indices:[allHostingObjectsIndex]});
		} 
		// } else {
			
		// }
		ExpirationDict[expiry_day].push(allHostingObjectsIndex);
		// ExpirationDict[expiry_day].push(allHostingObjectsIndex);

		emit OnCreateContract(allHostingObjectsIndex);

		// allHostingObjects[allHostingObjectsIndex] = HostingObject({hostingReplicationFactor:ReplicationFactor, mainContentHash: CID,hostingContractName : ContractName, contractStorageUsed : ContractSize, deployedDaysLength: ContractDuration,contractPayAmount : msg.value,creationTimestamp:block.timestamp, ownerAddress:msg.sender});
		allHostingObjects[allHostingObjectsIndex] = HostingObject({hostingReplicationFactor:ReplicationFactor, mainContentHash: CID,hostingContractName : ContractName, contractStorageUsed : ContractSize, deployedDaysLength: ContractDuration,creationTimestamp:uint32(block.timestamp), ownerAddress:msg.sender});

		payable(collection_address).transfer(_required_amt);

		return fileID;
	}

	function extendContract(uint32 fileID, uint16 extendDaysLength) public payable returns (bool success) { // uses unique fileID index to extend a file
		// handleExpiredContracts(); // scrub old contracts

		require(extendDaysLength > 0, "Bad Contract Extension Length");
		require((allHostingObjects[fileID].deployedDaysLength + extendDaysLength) <= MaximumStorageTime, "Bad Storage Duration");
		// check if contract really exists ?

		uint160 userID = uint160(msg.sender);

		require(users[userID].numHostingObjects > 0, "No Hosted Files"); // maybe use exact bytesize for cid ?

		uint256 _required_amt = GetHostingCost(allHostingObjects[fileID].contractStorageUsed,extendDaysLength,allHostingObjects[fileID].hostingReplicationFactor);
		require(msg.value >= _required_amt, "Did not Pay Enough for Contract Extension");

		allHostingObjects[fileID].deployedDaysLength += extendDaysLength;

		emit OnExtendContract(fileID,extendDaysLength);

		return true;

	}

	// function scrubContracts() internal { // this contract is merely for manual scrubbing by admin. users only mark deletion, they never execute the deletion.
	// 	for (uint32 i=0; i<allActiveHostingIndices.length; i++) {
	// 		uint32 fileID = allActiveHostingIndices[i];
	// 		if (block.timestamp-allHostingObjects[fileID].creationTimestamp > 60*60*24*allHostingObjects[fileID].deployedDaysLength) {
	// 			InternalDelete(fileID);
	// 			emit OnContractExpired(fileID);
	// 		}
	// 	}
	// }


	// DELETE FUNCTIONS
	function deleteContract(uint32 fileID) public returns (bool success) { // user delete of his own contracts/files
		require(allHostingObjects[fileID].ownerAddress == msg.sender, "Only the Owner can delete Contract");
		return InternalDelete(fileID);
	}

	function changeContractOwnership(address newOwner, uint32 fileID) public returns (bool success) { // user delete of his own contracts/files
		// return InternalDelete(msg.sender,fileID);
		require(allHostingObjects[fileID].ownerAddress == msg.sender, "Only the Owner can change Ownership");
		allHostingObjects[fileID].ownerAddress = newOwner;
		emit OnChangedOwnership(fileID);
		return true;
	}

	function InternalDelete(uint32 fileID) internal returns (bool success) { // will delete a file via CID, and will also take care of cleaning up of memory objects, so always delete files with this function.
		require(allHostingObjects[fileID].ownerAddress != address(0), "Contract doesn't exist"); // maybe use exact bytesize for cid ?

		uint160 userID = uint160(allHostingObjects[fileID].ownerAddress);

		require(users[userID].numHostingObjects > 0, "No Hosted Files"); // maybe use exact bytesize for cid ?

		contractsTotalCount--;
		contractsTotalSize -= allHostingObjects[fileID].contractStorageUsed; // remove size from global size tracker

		emit OnRemoveContract(fileID);

		if ( users[userID].numHostingObjects - 1 == 0 ) { // also clean up / remove user entirely if he has no more remaining files...
			delete users[userID];  // remove entire userobj
		} else {
			users[userID].totalHostingSize -= allHostingObjects[fileID].contractStorageUsed; // remove size from local user obj size tracker
			users[userID].numHostingObjects--;

			// for (uint32 i=0; i < users[userID].active_indices.length; i++) {
			// 	if (users[userID].active_indices[i] == fileID) {
			// 		users[userID].active_indices[i] = users[userID].active_indices[users[userID].active_indices.length - 1];
			// 		users[userID].active_indices.pop();
			// 		break;
			// 	}
			// }
			users[userID].active_indices[users[userID].fileID_to_Index[fileID]] = 0;
			users[userID].free_indices.push(users[userID].fileID_to_Index[fileID]);
			delete users[userID].fileID_to_Index[fileID];
			// delete users[userID].active_indices[fileID];

		}


		// delete from global list indexer
		allActiveHostingIndices[fileID_to_index_allActiveHostingIndices[fileID]] = 0; // set this slot temporarily to 0 to "deactivate"
		free_indices_allActiveHostingIndices.push(fileID_to_index_allActiveHostingIndices[fileID]); // add free slot to list so we can reuse it
		delete fileID_to_index_allActiveHostingIndices[fileID]; // reference not needed anymore after we passed it to freeslots

		delete allHostingObjects[fileID]; // remove global hosting obj struct
		return true;
	}

	// VIEW FUNCTIONS
	function GetActiveUserContracts(address userID) public view returns (uint32[] memory) { // view user statistics [number of files,total size stored]
		return users[uint160(userID)].active_indices;
	}

	// function GetInactiveUserContracts(address userID) public view returns (uint32[] memory) { // view user statistics [number of files,total size stored]
	// 	return users[uint160(userID)].inactive_indices;
	// }
	
	// function GetAllHostingContracts() public view returns (uint32[] memory) { // view user statistics [number of files,total size stored]
	// 	return allActiveHostingIndices;
	// }

	// function GetAllHostingContracts(uint32 bound_a, uint32 bound_b) public view returns (uint32[] memory) { // view user statistics [number of files,total size stored]
    //     uint32 _size = bound_b-bound_a;
	// 	uint32 limit = bound_b;
	// 	if (_size > allActiveHostingIndices.length-bound_a) {
	// 		_size = uint32(allActiveHostingIndices.length)-bound_a;
	// 		limit = uint32(allActiveHostingIndices.length);
	// 	}
	// 	uint32[] memory memoryArray = new uint32[](_size);
    //     for (uint32 j=bound_a; j<limit; j++) {
	// 		// if (j < allActiveHostingIndices.length) {
	// 		memoryArray[j] = allActiveHostingIndices[j];
	// 		// }
    //     }
    //     return memoryArray;
	// }

	function GetAllHostingContracts(uint32 bound_a, uint32 bound_b) public view returns (uint32[] memory) { // view user statistics [number of files,total size stored]
        uint32 _size = (bound_b+1)-bound_a;
		uint32 limit = (bound_b+1);
		if (_size > allActiveHostingIndices.length-bound_a) {
			_size = uint32(allActiveHostingIndices.length)-bound_a;
			limit = uint32(allActiveHostingIndices.length);
		}
		uint32[] memory memoryArray = new uint32[](_size);
        for (uint32 j=bound_a; j<limit; j++) {
			// if (j < allActiveHostingIndices.length) {

			memoryArray[j-bound_a] = allActiveHostingIndices[j];
			// }
        }
        return memoryArray;
	}

	// mapping(uint16 => uint32[]) public ExpirationDict;
	// function GetExpirationDict() public view returns (uint32[] memory) { // view user statistics [number of files,total size stored]
	// function GetExpirationDict() public view returns (uint[] memory) { // view user statistics [number of files,total size stored]
	
    //     uint[] memory memoryArray = new uint[](16);
    //     for (uint16 j=0; j< ExpirationDict[1].length; j++) {
    //         memoryArray[j] = ExpirationDict[1][j];
    //     }
    //     return memoryArray;
	// }
	
	function GetUserStats(address userID) public view returns (uint32 numHostingObjects,uint64 totalHostingSize) { // view user statistics [number of files,total size stored]
		return (users[uint160(userID)].numHostingObjects,users[uint160(userID)].totalHostingSize);
	}

	function GetContract(uint32 fileID) public view returns (HostingObject memory) { // view specific file via ID, 0-X, ressource friendly retrieval
		return allHostingObjects[fileID];
	}

	function GetUserID(address _userID) public pure returns (uint160 userID) { // returns userID converted from address
		return uint160(_userID);
	}
	
	function GetContractDurationLeft(uint32 fileID) public view returns (uint32 timeleft) { // time left in SECONDS!
		// return uint32(block.timestamp-allHostingObjects[fileID].creationTimestamp);
		return uint32((allHostingObjects[fileID].deployedDaysLength * day_cycle)-(block.timestamp-allHostingObjects[fileID].creationTimestamp));
	}

	// function authorizedDelete(uint[] memory expired_array) internal { // time left in SECONDS!
		// for (uint32 i=0; i< expired_array.length; i++) {
			// InternalDelete(expired_array[i]);
			// emit OnContractExpired(expired_array[i]);
			
			// delete ExpirationDict[i][j];
		// }
		
	// }
	
	function wipeContracts(uint16 maxDelete) public returns (uint32 del_count) {
		uint16 total_deletion_count = 0;
		uint16 today_date = uint16((uint32(block.timestamp)-controllerCreationTime)/fake_day_cycle)+1; // exact days since contract start -- exact todays day since contract start

		if (contractsTotalCount > 0) { // only look up if there is active contracts
			uint16 start_date = uint16(last_date_check_ExpirationDict); // required to avoid overwrite?
			for (uint16 d=start_date; d<=today_date; d++) { // check the days that have passed if any expirations happened

				if (ExpirationDict[d].length > ExpirationDict_Indexer[d]) { // only check if there is indices that (we haven't handled yet)

					uint32 start_index = uint32(ExpirationDict_Indexer[d]);
					uint32 end_index = uint32(start_index+maxDelete-total_deletion_count);
					if (end_index >= ExpirationDict[d].length) {
						end_index = uint32(ExpirationDict[d].length);
					}

					for (uint32 j=start_index; j<end_index; j++) {
						
						uint32 expiredID = uint32(ExpirationDict[d][j]);
						bool is_expired = uint32(block.timestamp) >= (allHostingObjects[expiredID].creationTimestamp + (allHostingObjects[expiredID].deployedDaysLength * fake_day_cycle));
				
						if (is_expired) { // C expired

							InternalDelete(expiredID);
							emit OnContractExpired(expiredID);

							total_deletion_count ++;
							ExpirationDict_Indexer[d] ++;

							if (total_deletion_count == maxDelete) { // clean exit once we removed entire threshold
								return total_deletion_count;
							}
						} else { // there is nothing more to check
							return total_deletion_count;
						}
					}
				} else {
					if (today_date > d + 1) { // only increase index if we are 2 days ahead in todays date. // never set below 2 days behind...
						for (uint16 e=last_date_check_ExpirationDict; e<=d; e++) {
							if (ExpirationDict[e].length > 0 ) {
								delete ExpirationDict[e];
							}
						}
						// delete ExpirationDict[d];
						last_date_check_ExpirationDict = d + 1; // set this day as last day ( max 2 days behind for start checking )
						// last_date_check_ExpirationDict = d; // move to next day
					} else {
						return total_deletion_count; // make sure we return here so we don't handle days that haven't happened yet
					}
				}
			}
		}
		return total_deletion_count;
	}

	// UTILITY FUNCTIONS

    function GetHostingCost(uint32 ContractSize, uint16 ContractDuration, uint8 ReplicationFactor) public view returns (uint128 total_cost) {
		require(ReplicationFactor >= MinimumReplicationFactor, "ReplicationFactor Minimum"); 
        return (ContractSize * (hostingCost_perMB_perYEAR / 1048576 / 365) * ContractDuration * ReplicationFactor/8);
    }

	// ADMIN FUNCTIONS

    function SetAccountCollectionAddress(address set) public restricted {
        collection_address = set;
    }

    function SetHostingContractCost(uint128 set) public restricted {
        hostingCost_perMB_perYEAR = set;
    }
	
    function SetMaximumStorageTime(uint16 set) public restricted {
        MaximumStorageTime = set;
    }
	
    function SetMaximumDeletionCount(uint8 set) public restricted {
        MaximumDeletionCount = set;
    }

    function SetMinimumReplicationFactor(uint8 set) public restricted {
        MinimumReplicationFactor = set;
    }

    // function SetDailyShiftThreshhold(uint16 set) public restricted {
    //     DAILY_SHIFT_THRESHHOLD = set;
    // }

    function SetOperator(address newoperator) public restricted {
        operator = newoperator;
    }
	
    function force_collect() public restricted returns (bool success) {
		payable(collection_address).transfer(payable(address(this)).balance);
        emit OnFundCollection(collection_address, uint256(payable(address(this)).balance));
        return true;
    }

	// CONTRACT
	
	// VIEW EXPIRATION DICT......

    modifier restricted {
        require(msg.sender == owner || msg.sender == operator, "This function is restricted to owner");
        _;
    }

    constructor() {
        owner = msg.sender;

		// createTime = block.timestamp;
		controllerCreationTime = uint32(block.timestamp);

		last_date_check_ExpirationDict = 1;
		last_index_check_ExpirationDict = 0;
		day_cycle = 60*60*24;
		// day_cycle = 15;
		fake_day_cycle = 10*1;

		// SET DEFAULT VALUES
		// SetMaximumStorageFactor(360.25*10);
		SetMaximumStorageTime(3603);
		SetMaximumDeletionCount(0);
		SetMinimumReplicationFactor(3);
		// SetDailyShiftThreshhold(30);
		SetAccountCollectionAddress(0x6668685991423444373955509583400957353773);
		hostingCost_perMB_perYEAR = 1000000000000000000;

		SetOperator(0x6668685991423444373955509583400957353773);
		// minimumCost_perMB_perYEAR = 1000000000000000; // 0,000001
    }
}

