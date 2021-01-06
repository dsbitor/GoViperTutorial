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

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run analyses collected mpm data",
	Long: `mpm is a Mac Performance Monitor which analyzes the performance 
	and resource usage of mac computer components, such as CPU, memory, 
	disk, file systems, and communication links. 
	
	run is a command to analyse mpm collected data. 
	The command uses the environemnt, set up by the initialize command.
	Execute initialize at least once before executing run. The flag 
	called linecount which alters the number of lines per page in 
	run reports.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		initializeMpm(args)

		fmt.Println("MPM Run completed at",
			time.Now().Format("2006-01-02 15:04:05.000000"), "\n\n")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//runCmd.PersistentFlags().String("foo", "", "A help for foo")

	runCmd.PersistentFlags().IntP("linecount", "l", 65, "Number of Lines on a Printed Page ")
	viper.BindPFlag("linect", runCmd.PersistentFlags().Lookup("linecount"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runeMpm(args []string) {
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

	showconfig("Config Variables Inside runMpm: ")

	return

}
