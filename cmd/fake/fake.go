package fake

import (
	"logdna/logging"

	"github.com/spf13/cobra"
)

var sublog logging.Logger = logging.Default

// fakeCmd represents the fake command
var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		sublog.Print("cmd run: %s %s", cmd.Parent().Use, cmd.Use)
		return nil
	},
}

// Attach the main ip command to a parent cobra command
func Attach(cmd *cobra.Command, logger logging.Logger) {
	sublog = logger.Child("fake")
	cmd.AddCommand(fakeCmd)
}
