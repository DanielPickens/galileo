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

// RollbackCmd represents the rollback command
var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback database",
	Long:  `Rollback one database migration.`,
	Run:   execRollbackCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execRollbackCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.DBRollback()
}
