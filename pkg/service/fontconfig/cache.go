package fontconfig

import "os/exec"

func RunCache(verbose bool, force bool) error {
	var flags []string

	if verbose {
		flags = append(flags, "-v")
	}

	if force {
		flags = append(flags, "-f")
	}

	cmd := exec.Command("fc-cache", flags...)
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
