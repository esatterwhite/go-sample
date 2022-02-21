package cmd

import (
	"fmt"
	"logdna/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sublog logging.Logger

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
		sublog = logger.Child("fake")
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(viper.GetString("token"))
		sublog.Trace().Msg("fake")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fakeCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fakeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
