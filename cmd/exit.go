package cmd

import "os"

func CommandExit() error {
	os.Exit(0)
	return nil
}
