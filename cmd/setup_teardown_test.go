package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	if _, err := os.Stat("LICENSE"); err == nil {
		os.Remove("LICENSE")
	}

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	if _, err := os.Stat("LICENSE"); err == nil {
		os.Remove("LICENSE")
	}

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}
