package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func showconfig(statestr string) {
	config := viper.AllSettings()
	fmt.Printf("\n\n" + statestr + "\n")
	for k, v := range config {
		fmt.Printf(statestr+" Config: Key = %s Config Type = %T, Config Value = %s\n", k, v, v)
	}
	fmt.Printf("\n\n")
}

func main() {
	// Declare two flags one integer, and one boolean
	ll := pflag.IntP("linelist", "l", 5, "Number of lines to list")
	debug := pflag.BoolP("debug", "d", false, "Switch debugging off/on (on = true)")

	fmt.Println("Flag default - lines to list has value: ", *ll)
	fmt.Println("Flag default - debugging is set to: ", *debug)

	// Setting Defaults
	// Find home directory and set as default
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set Viper configuration default values
	viper.SetDefault("homedir", home)
	viper.SetDefault("linelist", 10)
	viper.SetDefault("debug", "true")
	viper.SetDefault("wrkdir", ".")

	showconfig("Predefined")

	// Call flag.Parse() after all flags have been defined
	pflag.Parse()

	// Bind the command line flags to the configuration variables
	viper.BindPFlags(pflag.CommandLine)

	showconfig("Modified by flags")

	fmt.Printf("Last values:\n")
	fmt.Printf("Debug: %t\n", viper.GetBool("debug"))
	fmt.Printf("Linelist: %d\n", viper.GetInt("linelist"))

}
