package types

type Plugin struct {
	Name   string // Plugin name
	OnLoad func() // Function has be executed even plugin loaded
	F      func() // Main plugin function
}
