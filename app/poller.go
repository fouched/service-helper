package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var lastVersion string

func runInstaller(l *os.File, d string) bool {

	for range time.Tick(time.Second * 5) {
		fmt.Println("polling:" + d)
		_, err := fmt.Fprintln(l, "polling")
		if err != nil {
			fmt.Println(err)
			return false
		}

		entries, err := os.ReadDir(d)

		runInstall := false
		for _, e := range entries {
			if strings.HasSuffix(e.Name(), "install.json") {
				runInstall = true
			}
		}

		if runInstall {
			for _, e := range entries {
				if strings.HasSuffix(e.Name(), "exe") {
					lastVersion = e.Name()
				}
			}

			fmt.Println("lastVersion:" + lastVersion)
			process := exec.Command(d + lastVersion)

			e := process.Run()
			if e == nil {
				fmt.Println("running installer")

			} else {
				fmt.Printf("Command failed with exit code %d\n", process.ProcessState.ExitCode())
				fmt.Println(err)
				return false
			}

		} else {
			fmt.Println("Installation not ready, continue polling...")
		}

	}

	return true
}
