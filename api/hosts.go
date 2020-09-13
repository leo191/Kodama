package api

import (
	"os"
)

func readConfig(hostfile string) {
	hostconfig, err := os.Open(hostfile)
}
