/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package start

import (
	"github.com/DanielPickens/galileo/pkg/proc"

	"github.com/spf13/cobra"
)

// WatcherCmd represents the watcher command
var WatcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "Start watcher daemon",
	Long:  `Start watcher daemon.`,
	Run:   execWatcherCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execWatcherCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.StartWatcher()
}
