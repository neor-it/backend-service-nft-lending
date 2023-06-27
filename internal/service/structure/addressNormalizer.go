package structure

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func NormalizeAddress(address string) string {
	for len(address) != 42 {
		if strings.HasPrefix(address, "0x0") {
			address = "0x" + address[3:]
		} else {
			break
		}
	}

	return string(common.HexToAddress(strings.ToLower(address)).Hex())
}
