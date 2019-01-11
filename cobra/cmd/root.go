// Copyright © 2015 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
rm -rf rm -rf $GOPATH/src/github.com/wataash/tesgo/tesCobra/
cobra init github.com/wataash/tesgo/tesCobra

~/.cobra.yaml:
https://github.com/spf13/cobra/tree/master/cobra

author: Wataru Ashihara <wataash0607@gmail.com>
year: 2019
# license: MIT
license:
  header: This file is part of {{ .appName }}.
  text: |
    {{ .copyright }}

    This is my license. There are many like it, but this one is mine.
    My license is my best friend. It is my life. I must master it as I must
    master my life.
*/

package cmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile, userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

// Execute executes the root command.
func Execute() {
	fmt.Println("rootCmd.Execute()")
	rootCmd.Execute()
}

var _author *string
var _ = _author

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	_author = rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// --------------------------------------------------
	tmpFlag := rootCmd.PersistentFlags().Lookup("author")
	_ = tmpFlag
	// *author = "foo"
	// _, _ = fmt.Fprintf(os.Stderr, "author %p -> *author %s\n", author, *author)
	// _, _ = fmt.Fprintf(os.Stderr, "&tmp.Value %p -> tmp.Value %p %s\n", &tmp.Value, tmp.Value, tmp.Value)
	// author                               0xc4200a5130 -> *author foo
	// &tmp.Value 0xc4201e28f0 -> tmp.Value 0xc4200a5130 :          foo
	//
	// --------------------------------------------------
	// try: comment out
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	// should be moved to add.go, init.go
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		// reads ~/.cobra.yaml, but ~/.cobra
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	// with HCL
	// yaml: license.header can be got as string

	tmpAll := viper.AllKeys()                  // {"author", "useviper", "license"}
	tmpAG := viper.Get("author")               // interface "NAME HERE <EMAIL ADDRESS>"
	tmpA := viper.GetString("author")        // "NAME HERE <EMAIL ADDRESS>"
	tmpV := viper.GetString("useViper")        // "true"
	tmpVB := viper.GetBool("useViper")         // true
	tmpL := viper.GetString("license")         // "apache"
	tmpLH := viper.GetString("license.header") // ""
	tmpLT := viper.GetString("license.text")   // ""
	tmpLG := viper.Get("license")              // interface{}
	tmpLHG := viper.Get("license.header")      // interface nil
	tmpLTG := viper.Get("license.text")        // interface nil
	tmpLS := viper.Sub("license")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	tmpAll = viper.AllKeys()                  // + "year"
	tmpAG = viper.Get("author")               // interface{"Wataru Ashihara <wataash0607@gmail.com>"}
	tmpA = viper.GetString("author")        // Wataru Ashihara <wataash0607@gmail.com>
	tmpV = viper.GetString("useViper")        // unchanged
	tmpVB = viper.GetBool("useViper")         // unchanged
	tmpL = viper.GetString("license")         // ""
	tmpLH = viper.GetString("license.header") // "This file is part of {{ .appName }}."
	tmpLT = viper.GetString("license.text")   // "{{ .copyright }} ... "
	tmpLG = viper.Get("license")              // interface map
	tmpLHG = viper.Get("license.header")      // interface nil
	tmpLTG = viper.Get("license.text")        // interface nil
	tmpLS = viper.Sub("license")

	_ = tmpAll
	_ = tmpAG
	_ = tmpA
	_ = tmpV
	_ = tmpVB
	_ = tmpL
	_ = tmpLH
	_ = tmpLT
	_ = tmpLG
	_ = tmpLHG
	_ = tmpLTG
	_ = tmpLS
}
