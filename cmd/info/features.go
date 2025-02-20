/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package info

import (
	"github.com/DanielPickens/galileo/pkg/config"

	"github.com/spf13/cobra"
)

// FeaturesCmd represents the features command
var FeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "Print service features configuration",
	Long:  `Print service features configuration based on environment`,
	Run:   execFeaturesCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execFeaturesCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	config.Env.PrintServiceFeatures()
}
