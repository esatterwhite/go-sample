package fake

import (
	"fmt"
	"logdna/logging"

	"github.com/spf13/cobra"
)

// RealCommand Internal cobra wrapper
type RealCommand struct {
	cobra.Command
}

var sublog logging.Logger = logging.Default

// The real command is a subcommand of fake
var realCmd = &cobra.Command{
	Use:   "real",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		sublog.Print("cmd run: %s %s", cmd.Parent().Use, cmd.Use)
	},
}

// Attach to parent
func Attach(parent *cobra.Command, logger logging.Logger) {
	sublog = logger.Child(fmt.Sprintf("%s real", parent.Use))
	parent.AddCommand(realCmd)
}
