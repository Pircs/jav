/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"time"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jav",
	Short: "javbus spider written by go",
	Long:  `javbus spider written by go`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		get(1)
		for {
			time.Sleep(time.Second)
			if int(comindex) == curindex {
				fmt.Println("抓取完成")
				return
			}
		}
	},
	//Run: func(cmd *cobra.Command, args []string) {
	//	var chanbuffer = make(chan int, parallel)
	//	getdetail("https://www.javbus.com/CESD-833", "CESD-833", chanbuffer)
	//},
}