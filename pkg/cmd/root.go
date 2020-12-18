package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "test",
		Short: "Test golang containerized build.",
		Long:  `Test golang containerized build.`,
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

// From https://github.com/jeremyhuiskamp/golang-docker-scratch/blob/026ddd6232ecf1f28974f404d6b8247b9d2cf4df/main.go
// main runs some tests to exercise the environment, including
// loading a timezone and verifying a well-known TLS certificate.
func main() {
	fmt.Println("Tests from within scratch container:")
	errors := testTZ()
	errors += testTLS()
	os.Exit(errors)
}

func testTZ() (errors int) {
	shanghai, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("Unable to load timezones: %s\n", err)
		errors++
	} else {
		fmt.Printf("Successfully loaded %q\n", shanghai)
	}
	return
}

func testTLS() (errors int) {
	rsp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Printf("Unable to establish https connection: %s\n", err)
		errors++
	} else {
		rsp.Body.Close()
		fmt.Println("Successfully established https connection")
	}
	return
}
