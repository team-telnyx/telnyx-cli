/*
Copyright © Telnyx LLC
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/team-telnyx/telnyx-cli/config"
	"gopkg.in/yaml.v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the configuration files for the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := config.ConfigExists()
		if res {
			fmt.Println("Already configured!")
			os.Exit(0)
		} else if err != nil {
			cobra.CheckErr(err)
		}

		defaultConfig, err := yaml.Marshal(config.DefaultConfig())
		if err != nil {
			cobra.CheckErr(err)
		}

		if err := os.Mkdir(config.DefaultConfigFolder(), os.ModePerm); err != nil {
			cobra.CheckErr(err)
		}

		if err := os.Mkdir(path.Join(config.DefaultConfigFolder(), config.CacheFolder), os.ModePerm); err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("Creating configuration file in %s\n", config.DefaultConfigPath())

		f, err := os.Create(config.DefaultConfigPath())
		if err != nil {
			cobra.CheckErr(err)
		}

		defer f.Close()

		_, err = f.Write(defaultConfig)
		if err != nil {
			cobra.CheckErr(err)
		}

		f.Sync()

		fmt.Println("Done!")
		fmt.Println()
		fmt.Println("Note: To use 'telnyx-cli get service-diff', configure your GitHub token:")
		fmt.Println("  1. Create token at: https://github.com/settings/tokens (needs 'repo' scope)")
		fmt.Println("  2. Edit ~/.telnyx-cli/config.yaml and add:")
		fmt.Println("     github_token: your_token_here")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
