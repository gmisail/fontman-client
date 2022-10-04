package util

// we use log for now; should move to a more robust way when commands are finalized.
import (
	"log"
	"os/exec"
)

func Cache(verbose bool, force bool) error {
	var verboseFlag string
	var forceFlag string

	if verbose {
		verboseFlag = "-v"
	}
	if force {
		forceFlag += "-f"
	}

	cmd := exec.Command("fc-cache", verboseFlag, forceFlag)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
