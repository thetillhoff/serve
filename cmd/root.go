package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thetillhoff/serve/pkg/serve"

	"github.com/spf13/viper"
)

var (
	cfgFile string

	verbose   bool
	directory string
	ipaddress string
	port      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "A simple webserver",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err error
		)

		serveEngine := serve.DefaultEngine()
		serveEngine.Verbose = verbose
		serveEngine.Directory = directory
		serveEngine.Ipaddress = ipaddress
		serveEngine.Port = port

		serveEngine.Serve()

		if err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.serve.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Every request will be printed.")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", "./", "Serve another directory. Default is './'.")
	rootCmd.PersistentFlags().StringVarP(&ipaddress, "ip-address", "i", "0.0.0.0", "Bind to specific ip-address. Default is '0.0.0.0'.")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "3000", "Bind to specific port. Default is ':3000'.")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".serve" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".serve")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
