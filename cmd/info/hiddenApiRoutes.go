/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package info

import (
	"github.com/DanielPickens/galileo/pkg/proc"

	"github.com/spf13/cobra"
)

// HiddenApiRoutesCmd represents the hidden-api-routes command
var HiddenApiRoutesCmd = &cobra.Command{
	Use:   "hidden-api-routes",
	Short: "Print hidden API routes",
	Long:  `Print all hidden API routes table`,
	Run:   execHiddenApiRoutesCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execHiddenApiRoutesCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.PrintHiddenRoutesTable()
}
