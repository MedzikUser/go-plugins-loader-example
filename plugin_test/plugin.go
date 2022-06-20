package main

import "fmt"

// Plugin name
var PluginName = "test"

// On load function
func OnLoad() {
	fmt.Println("Loading plugin test...")
}

// Main plugin function
func F() {
	fmt.Println("Function from test plugin success executed!")
}
