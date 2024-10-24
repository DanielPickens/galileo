/*
Copyright © 2024 Daniel Pickens
Author:  Daniel Pickens
Contact: 
*/
package main

import "github.com/danielpickens/galileo/cmd"

var VERSION string = "2.2.1-default"

func main() {
	cmd.Version = VERSION
	cmd.Execute()
}
