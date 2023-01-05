/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "A brief description of your command",
	Long:  `Calculator mod is not zero`,
	Run: func(cmd *cobra.Command, args []string) {
		var oddSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 != 0 {
				oddSum = oddSum + itemp
			}
		}
		fmt.Printf("The odd addition of %s is %d", args, oddSum)
	},
}

func init() {
	addCmd.AddCommand(oddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
