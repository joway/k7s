package cmd

import (
	"fmt"
	"github.com/joway/k7s/log"
	"github.com/spf13/cobra"
	"os"
)

var logger = log.NewLogger(log.Debug)
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			logger.Error(cmd.UsageString())
			os.Exit(1)
		}
		pathPattern := args[0]
		fmt.Printf("%s\n", pathPattern)
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	//renderCmd.PersistentFlags().String("foo", "", "A help for foo")
}
