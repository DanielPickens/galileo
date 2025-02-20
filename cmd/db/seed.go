/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package db

import (
	"github.com/DanielPickens/galileo/pkg/proc"

	"github.com/spf13/cobra"
)

// SeedCmd represents the seed command
var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed database",
	Long:  `Backfill database with seed data.`,
	Run:   execSeedCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execSeedCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.DBSeed()
}
