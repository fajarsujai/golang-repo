/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abdfnx/gosh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cicdModCmd represents the cicdMod command
var cicdModCmd = &cobra.Command{
	Use:   "cicdMod",
	Short: "A brief description of your command",
	Long:  `Ci/Cd Mod`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		fstatus, _ := cmd.Flags().GetBool("auth")

		if fstatus { //if status is true, call addFloat
			authCicd()
		}
	},
}

func init() {
	rootCmd.AddCommand(cicdModCmd)
	cicdModCmd.Flags().BoolP("auth", "a", false, "Auth cicdMod")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cicdModCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cicdModCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Employee struct {
	Registry01Url, Registry01User, Registry01Password string
}

func authCicd() {
	if _, err := os.Stat("C:/Users/fajar/OneDrive/Dokumen/qoin/auth.json"); err == nil {
		// path/to/whatever exists
		viper.SetConfigType("json")
		// viper.AddConfigPath("C:/Users/fajar/OneDrive/Dokumen/qoin")
		viper.AddConfigPath("/mnt/c/Users/fajar/go/project1/project1")
		viper.SetConfigName("auth")

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
		// registry01Url := viper.GetString("REGISTRY01_URL")
		registry01User := viper.GetString("REGISTRY01_USER")
		registry01Password := viper.GetString("REGISTRY01_PASSWORD")
		fmt.Println(viper.GetString("REGISTRY"))
		fmt.Println(viper.GetString("REGISTY_USER"))
		fmt.Println(viper.GetString("REGISTY_PASSWORD"))
		fmt.Println(viper.GetString("GIT_PASSWORD"))
		fmt.Println(viper.GetString("GIT_USER"))
		// sh.Command("echo", "hello\tworld").Command("cut", "-f2").Run()
		// sh.Command("echo", "hello").Run()

		data := Employee{
			Registry01Url:      viper.GetString("REGISTRY01_URL"),
			Registry01User:     registry01User,
			Registry01Password: registry01Password,
		}

		file, _ := json.MarshalIndent(data, "", " ")

		_ = ioutil.WriteFile("test.json", file, 0644)
	} else if errors.Is(err, os.ErrNotExist) {
		// sh.Command("echo", "hello\tworld").Command("cut", "-f2").Run()
		// sh.Command("docker --version").Run()
		gosh.Run("helm repo list")
		// path/to/whatever does *not* exist
		fmt.Println("File not Exists")
	}
}
