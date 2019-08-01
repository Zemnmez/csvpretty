package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zemnmez/cardauth/apdu/gen/lib"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile    string
	inputCSV  gen.CSVFile
	inputGo gen.GoFile
	outputFile gen.OutFile
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates definitions for the smart card protocol",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gen.yaml)")



	rootCmd.PersistentFlags().Var(&inputGo, "go", "Input go file")

	rootCmd.PersistentFlags().Var(&inputCSV, "data", "input csv file")

	rootCmd.PersistentFlags().Var(&outputFile, "out", "output file (must be an existing valid go file)")
	rootCmd.MarkPersistentFlagRequired("go")
	rootCmd.MarkPersistentFlagRequired("data")
	rootCmd.MarkPersistentFlagRequired("out")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

		// Search config in home directory with name ".gen" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
