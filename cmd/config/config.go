package config

import (
	"gin-quickly-template/config"
	"github.com/spf13/cobra"
	"os"
)

var (
	configYml string
	forceGen  bool
	StartCmd  = &cobra.Command{
		Use:     "config",
		Short:   "Generate config file",
		Example: "app config -p config/config.yaml.example -f",
		Run: func(cmd *cobra.Command, args []string) {
			println("Generate config...")
			err := load()
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "path", "p", "./config.yaml.example", "Generate config in provided path")
	StartCmd.PersistentFlags().BoolVarP(&forceGen, "force", "f", false, "Force generate config in provided path")
}

func load() error {
	return config.GenConfig(configYml, forceGen)
}
