package util

// we use log for now; should move to a more robust way when commands are finalized.
import (
	"log"
	"os/exec"
)

func Cache(verbose bool, force bool) error {
	flags := []string{}

	if verbose {
		flags = append(flags, "-v")
	}

	if force {
		flags = append(flags, "-f")
	}

	cmd := exec.Command("fc-cache", flags...)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
