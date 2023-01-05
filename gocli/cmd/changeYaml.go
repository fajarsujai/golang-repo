/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// changeYamlCmd represents the changeYaml command
var changeYamlCmd = &cobra.Command{
	Use:   "changeYaml",
	Short: "A brief description of your command",
	Long:  `Change Yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		changeYaml()
	},
}

func init() {
	rootCmd.AddCommand(changeYamlCmd)
}

type repo struct {
	Charts struct {
		Name string `yaml:"name"`
		Repo string `yaml:"repo"`
		Tags string `yaml:"tags"`
	} `yaml:"Charts"`
}

func changeYaml() {
	viper.SetConfigName("config")
	viper.AddConfigPath(currentdir())
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatal(err)
	}

	C := repo{}

	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	fmt.Println(C)

	// Change value in map and marshal back into yaml
	C.Charts.Tags = "assyk banget"

	fmt.Println(C)

	d, err := yaml.Marshal(&C)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(string(d))

	// write to file
	f, err := os.Create("C:/Users/fajar/go/project1/project1/testyaml")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("changed.yaml", d, 0644)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}

func currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}
