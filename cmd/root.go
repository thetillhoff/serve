/*
Copyright © 2021 Till Hoffmann <till.hoffmann@enforge.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string

	fileServer http.Handler

	verbose   bool
	logfile   string
	port      string
	ipaddress string
	directory string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err error
		)

		// Parsing or setting default for directory
		if directory == "" {
			directory = "./"
		} else {
			directory = path.Clean(directory)
		}

		if verbose {
			log.Println("INF verbose=true")
			// log.Println("INF logfile=" + logfile)
			log.Println("INF port=" + port)
			log.Println("INF ipaddress=" + ipaddress)
			log.Println("INF directory=" + directory)
		}

		// Creating Webserver for static files
		fileServer = http.FileServer(http.Dir(directory))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if verbose {
				log.Println("INF Serving ", r.RequestURI)
			}
			fileServer.ServeHTTP(w, r)
		})

		// Starting Webserver
		fmt.Println("Listening on " + ipaddress + ":" + port + " ...")
		err = http.ListenAndServe(ipaddress+":"+port, nil)
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
	// rootCmd.PersistentFlags().StringVarP(&logfile, "logfile", "l", "", "Output to file instead of stdout.")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "3000", "Bind to specific port. Default is ':3000'.")
	rootCmd.PersistentFlags().StringVarP(&ipaddress, "ip-address", "i", "0.0.0.0", "Bind to specific ip-address. Default is '0.0.0.0'.")
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", "", "Serve another directory. Default is './'.")

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
