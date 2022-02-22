package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

import (
	"fmt"
	"logdna/cmd/fake"
	"logdna/cmd/ip"
	"logdna/logging"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var token string
var logger logging.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "logdna",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	err := initConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lvl := "error"
	var debug bool = viper.GetBool("debug")

	if debug {
		lvl = "debug"
	}

	logger = logging.New(lvl, "logdna")
	logger.Print("Debug logging enabled")

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.config/logdna.json)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&token,
		"token",
		"t",
		"",
		"LogDNA Access Token for performing authenticated tasks",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().String("foo-bar", "", "Do The Foobar")

	ip.Attach(rootCmd, logger)
	fake.Attach(rootCmd, logger)
}

func initConfig() error {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.SetEnvPrefix("logdna")
	viper.BindEnv("config")
	viper.BindEnv("token")
	viper.BindEnv("debug")

	viper.AutomaticEnv()

	viper.SetConfigName("logdna.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(path.Join(home, ".config", "logdna"))
	viper.AddConfigPath(path.Join("/etc/logdna/"))
	viper.AddConfigPath(path.Join("."))

	config := viper.GetString("config")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else if config != "" {
		viper.SetConfigFile(config)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("token", rootCmd.Flags().Lookup("token"))
	return nil
}
