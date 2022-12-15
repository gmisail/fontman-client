package fontconfig

import "os/exec"

func ListAll() (string, error) {
	cmd := exec.Command("fc-list", ":", "family", "style", "file")
	output, outputErr := cmd.Output()

	if outputErr != nil {
		return "", outputErr
	}

	return string(output), nil
}
