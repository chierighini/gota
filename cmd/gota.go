package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "gota",
		Short: "Gota serena bot, o segundo e nao unico",
		Long:  `esgurme`,
	}
)

// has an error type return value
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnFinalize(wat)
}

func initConfig() {
	fmt.Print("begin")
}

func wat() {
	fmt.Print("end")
}
