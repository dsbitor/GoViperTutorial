/*
Copyright © 2020 I.T. Operational Risk Ltd., Toronto,
ON, Canada  <consultant.itor@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
// ITOR edited / added appropriate descriptions  for Short: and Long: below
var rootCmd = &cobra.Command{
	Use:   "mpm1",
	Short: "An example of Viper configuration management",
	Long: `Run mpm with various flag settings such as:
	--trace true
	-t false
	`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root cmd & sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setConfigDefaults() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.SetDefault("homedir", home)
	viper.SetDefault("wrkdir", ".")
	viper.SetDefault("datadir", home+"/tmpr")
	viper.SetDefault("app.softlicencetype", "MIT")
	viper.SetDefault("support", "itoprisk@gmail.com")
	viper.SetDefault("logging.filename", "aspex.log")
	viper.SetDefault("logging.dir", home+"/tmpr")
	viper.SetDefault("devparms.debug", true)
	viper.SetDefault("devparms.trace", true)
}

func showconfig(statestr string) {
	config := viper.AllSettings()
	fmt.Printf("\n\n" + statestr + "\n")
	for k, v := range config {
		fmt.Printf(statestr+" : Config Key = %s Config Value = %v\n", k, v)
	}
	fmt.Printf("\n\n")
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// ITOR this flag controls TRACE messages
	rootCmd.PersistentFlags().BoolP("trace", "t", true, "Switches on/off Trace messages in mpm")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mpm1.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Setting Configuration Defaults
	setConfigDefaults()
	showconfig("Default Configuration ")

	// Configuration File defined by flag?
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "mpm1" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("mpm1")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Failed to read configuration file:", viper.ConfigFileUsed())
	}

	showconfig("Conf after reading mpm1.toml")

	// Handle Environment Variables

	// Bind each viper name configuration var to an environment variable
	viper.BindEnv("wrkdir", "PWD")
	// Overwrite homedir with environment variable if HOME is defined
	viper.BindEnv("homedir", "HOME")
	viper.BindEnv("tmpdir", "TMPDIR")

	// Initialize environment variables that match config variables
	viper.AutomaticEnv()

	showconfig("Conf - Environment vars applied")

}
