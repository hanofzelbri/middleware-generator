package cmd

import (
	"fmt"
	"os"

	"github.com/hanofzelbri/middleware-generator/interfaces"

	"github.com/spf13/cobra"
)

var (
	options = interfaces.Options{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dmgmori-logging-generator",
	Short: "Generates logging middleware for golang interface",
	Long: `This golang generator can be used to generate a logging
middleware with the zerolog logging library for an provided interface.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o, err := interfaces.SetupConfig(options)
		if err != nil {
			return err
		}

		i, err := interfaces.BuildInterface(o)
		if err != nil {
			return err
		}

		return interfaces.PrintInterface(i, options.Output)
	},
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
	rootCmd.PersistentFlags().StringVarP(&options.Query, "interface", "i", "", "Interface definition to generate logging middleware for.")
	rootCmd.MarkPersistentFlagRequired("interface")

	rootCmd.PersistentFlags().StringVarP(&options.Output, "output", "o", "", "Output file. If empty StdOut is used")
	rootCmd.PersistentFlags().StringVarP(&options.Wrapper, "wrapper", "w", "", "Wrapper definition for implementation of middleware interface.")
	rootCmd.PersistentFlags().StringVarP(&options.MiddlewareFunctionName, "middlewareFunctionName", "f", "WithMiddleware", "Function name for middleware")
	rootCmd.PersistentFlags().StringVarP(&options.EmptyFunctionParamNamePrefix, "emptyFunctionParamNamePrefix", "p", "param", "If there is no function parameter name provided this prefix will be used")
	rootCmd.PersistentFlags().StringVarP(&options.EmptyFunctionReturnParamNamePrefix, "emptyFunctionReturnParamNamePrefix", "r", "ret", "If there is no function parameter return name provided this prefix will be used")
}
