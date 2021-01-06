/*
Copyright © 2020 I.T. Operational Risk Ltd., Toronto, ON, Canada  <consultant.itor@gmail.com>

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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initializeCmd represents the initialize command
var initializeCmd = &cobra.Command{
	Use:   "initialize",
	Short: "initialize builds the run-time environment for mpm",
	Long: `mpm is a Mac Performance Monitor which analyzes the performance 
	and resource usage of mac computer components, such as CPU, memory, 
	disk, file systems, and communication links. 
	
	initialize is a command to build the run-time environment for mpm. 
	The command sets up directory and file structure. The flag workdir
	allows the configuration parameter wkdir  seatting the working 
	directory to be changed at run time from the settings in the 
	configuration file.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		initializeMpm(args)

		fmt.Println("Application Initialization completed at",
			time.Now().Format("2006-01-02 15:04:05.000000"),
			"\n\n")
	},
}

func init() {
	rootCmd.AddCommand(initializeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	initializeCmd.PersistentFlags().StringP("workdir", "w", "", "Input Run-time Path to Working Directory ")
	viper.BindPFlag("wrkdir", initializeCmd.PersistentFlags().Lookup("workdir"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initializeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initializeMpm(args []string) {
	wd, err := os.Getwd()
	if err != nil {
		er(err)
		return
	}

	if len(args) > 0 {
		if args[0] != "." {
			wd = fmt.Sprintf("%s/%s", wd, args[0])
		}
	}

	// set new configuration variable logging.tstamp to highres (nanosec)
	viper.Set("logging.tstamp", "highres")

	showconfig("Config Variables Inside initializeMpm: ")

	// Write updated configuration values to file
	viper.WriteConfigAs("./writtenconf.toml")

	return

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
