/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package cmd

import (
	"os"
	"runtime/debug"
	"strings"

	"github.com/DanielPickens/galileo/pkg/clients/logger"
	"github.com/DanielPickens/galileo/pkg/config"
	"github.com/DanielPickens/galileo/pkg/proc"

	"github.com/spf13/cobra"
)

var Version string
var runtimeInfo, _ = debug.ReadBuildInfo()
var runtimeModuleInfo = strings.Split(runtimeInfo.Main.Path, "/")
var runtimeModuleName = runtimeModuleInfo[len(runtimeModuleInfo)-1]

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   runtimeModuleName,
	Short: runtimeModuleName + " cli",
	Long: `
        Galileo is a CLI tool for managing and deploying applications for redis operators and kubernetes administrators to aws s3 buckets.

` + runtimeModuleName + ` cli`,
}

// Execute adds all requiste commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...

	// Configure cobra
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Set global flags
	rootCmd.PersistentFlags().BoolVarP(&config.DevModeFlag, "dev", "d", false, "Run in development mode")
	rootCmd.PersistentFlags().BoolVarP(&config.EnvModeFlag, "env", "e", false, "Print environment before execution")
	rootCmd.PersistentFlags().StringVarP(&config.LogLevelFlag, "log", "l", "", "Log Level")

	// Initialize app config
	cobra.OnInitialize(initEnv)

	// Register persistent function for all commands
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		execRootPersistentPreRun()
	}
}

func initEnv() {
	if config.DevModeFlag {
		logger.SetDevMode()
	}
	config.OverrideLoggerUsingFlags()
	proc.InitServiceEnv(runtimeModuleName, Version)
}

func execRootPersistentPreRun() {
	logger.Debug("Executing root cmd persistent pre run ...")

	// You can initialize other features here ...
	// this will run before any command, make sure to put only global initializations here
	// to avoid running into nil pointers or undefined variables
	// ...

}
