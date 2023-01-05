/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `Calculator add opereation`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		fstatus, _ := cmd.Flags().GetBool("float")

		if fstatus { //if status is true, call addFloat
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func addInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sum = sum + itemp
	}

	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64

	for _, fval := range args {
		//convert string to float64
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}

		sum = sum + ftemp
	}

	fmt.Printf("Sum of floating numbers %s is %f", args, sum)
}
