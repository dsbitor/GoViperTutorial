package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func showconfig(statestr string) {
	config := viper.AllSettings()
	fmt.Printf("\n\n" + statestr + "\n")
	for k, v := range config {
		fmt.Printf(statestr+" : Config Key = %s Config Value = %v\n", k, v)
	}
	fmt.Printf("\n\n")
}

func showcEnvVar() {
	fmt.Println("Environment Variables available: ")
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], " = ", pair[1])
	}
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

	// Adding changes from Environmental Variables
	// Enable Viper to read Environment Variables
	viper.AutomaticEnv()
	// Bind  each name used for a configuration variable to environment variable
	viper.BindEnv("wrkdir", "PWD")
	// Overwrite homedir with environment variable if HOME is defined
	viper.BindEnv("homedir", "HOME")
	viper.BindEnv("hostname", "HOSTNAME") // Note hostname may not change
	// on my Mac HOSTNAME is not available as an Environment Variable to
	// Go programs by default
	// See 'Environment Variables available using function showcEnvVar() below

	showconfig("Environment Variables Applied")

	// Output available Environment Variables
	showcEnvVar()

	// Check if a particular key is set print key value if avail
	// Example of viper.IsSet(keyName)
	if !viper.IsSet("wrkdir") {
		fmt.Printf("Missing Working Directory Environment Variable: \n")
		log.Fatal("missing PWD")
	} else {
		fmt.Printf("After Environment Variable Applied Working Directory is: %v \n", viper.Get("wrkdir"))
	}

}
