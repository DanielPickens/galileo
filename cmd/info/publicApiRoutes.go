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

// PublicApiRoutesCmd represents the public-api-routes command
var PublicApiRoutesCmd = &cobra.Command{
	Use:   "public-api-routes",
	Short: "Print public API routes",
	Long:  `Print all public API routes table`,
	Run:   execPublicApiRoutesCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execPublicApiRoutesCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.PrintPublicRoutesTable()
}
