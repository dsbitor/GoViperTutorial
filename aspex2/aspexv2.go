//Example 2 - aspexv2.go
/* Code to:
    set default configuration values
	read viper configuraion from TOML file aspexv2.toml
	update / add configuration values from environment variables
	show values at each stage of program
*/

package main

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func showconfig(statestr string) {
	config := viper.AllSettings()
	fmt.Printf("\n\n" + statestr + "\n")
	for k, v := range config {
		fmt.Printf(statestr+" Config: Key = %s Config Value = %v\n", k, v)
	}
	fmt.Printf("\n\n")
}

func main() {
	// Environment Vars vary depending upon development and system settings
	// If you want to look at available Environment Variables Un-comment the
	// code block below and list environment variables using base os package
	/*
		envVars := os.Environ()
		for _, val := range envVars {
			fmt.Println("Environment Variable: ", val)
		}
	*/

	// Setting Defaults
	// Find home directory and set as default
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetDefault("homedir", home)
	viper.SetDefault("devparms.debug", "false")
	viper.SetDefault("wrkdir", ".")
	viper.SetDefault("datadir", home)
	viper.SetDefault("softlicencetype", "MIT")
	viper.SetDefault("support", "itoprisk@gmail.com")
	viper.SetDefault("logging.filename", "aspex.log")
	viper.SetDefault("logging.dir", home)

	showconfig("Predefined")

	// Adding Configuration File Default Changes & Additional Conf Parms
	// Set config file path including file name and extension
	viper.SetConfigFile("./aspexv2.toml")
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file, %s", err)
	}
	// Confirm which config file is used
	fmt.Printf("Using configuration file: %s\n", viper.ConfigFileUsed())

	showconfig("Configuration File Applied")

}
