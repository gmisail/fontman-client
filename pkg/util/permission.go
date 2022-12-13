package util

import "os"

func CheckRoot() bool {
	// os.Geteuid() should be sufficient in detecting if the user running it has root permission
	// SUDO_USER is not reliable as attacker could potentially set the environmental variable
	isRoot := os.Geteuid()
	return isRoot == 0
}
