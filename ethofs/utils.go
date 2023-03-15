package ethofs

import (
	"strings"
	"unsafe"

	cid "github.com/ipfs/go-cidutil"
)

func ByteSliceToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func scanForCids(data []byte) []string {
	var cidArray []string
	i, j, _, pin := cid.ScanForCid(data)
	for i != j {
		cidArray = append(cidArray, pin)
		if j >= (len(data) - 1) {
			return cidArray
		}

		// 68, dann 32 bis ende von value
		// replication := data[68:(68 + 32)]
		// replication := data[(i - 128):(i - 128 + 32)]
		// log.Info("lllllllllllllllllg", "data0", data, "replication", replication, "replication2", uint32(replication))
		data = data[j:(len(data) - 1)]
		// log.Info("lllllllllllllllll", "data", data, "i", i, "j", j, "pin", pin)
		i, j, _, pin = cid.ScanForCid(data)
	}
	return cidArray
}
