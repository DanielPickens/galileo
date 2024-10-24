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

// ProtectedApiRoutesCmd represents the protected-api-routes command
var ProtectedApiRoutesCmd = &cobra.Command{
	Use:   "protected-api-routes",
	Short: "Print protected API routes",
	Long:  `Print all protected API routes table`,
	Run:   execProtectedApiRoutesCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execProtectedApiRoutesCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.PrintProtectedRoutesTable()
}
