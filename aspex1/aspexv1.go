//Example 1 - aspexv1.go
// Code to read viper configuraion from TOML file aspexv1.toml
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// Set config file path including file name and extension
	viper.SetConfigFile("./aspexv1.toml")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// Confirm config file used
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	debug := viper.Get("devparms.debug") // returns string
	fmt.Printf("devparms.debug Value: %v, Type: %T\n", debug, debug)

	logfilename := viper.Get("logging.filename") // returns string
	fmt.Printf("logging.filename Value: %v, Type: %T\n", logfilename, logfilename)

	// Check if a particular key is set print if avail
	if !viper.IsSet("title") {
		log.Fatal("missing title")
	} else {
		fmt.Printf("Configuration File title: %v \n", viper.Get("title"))
	}
}
