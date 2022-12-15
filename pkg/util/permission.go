package util

import "os"

// CheckRoot checks if the user is running as a superuser (i.e. with `sudo`).
func CheckRoot() bool {
	// os.Geteuid() should be sufficient in detecting if the user running it has root permission
	// SUDO_USER is not reliable as attacker could potentially set the environmental variable
	return os.Geteuid() == 0
}
