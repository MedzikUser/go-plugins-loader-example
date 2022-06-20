package main

import (
	"fmt"
	"os"

	"github.com/MedzikUser/go-plugins-loader-example/plugin"
)

func main() {
	pluginsDirectory := "plugins"

	// find all plugins file
	files := readFilesFromDir(pluginsDirectory)

	// load plugins
	plugins := plugin.LoadPlugins(files)

	// execute all plugins function
	for _, plugin := range plugins {
		// execute plugin function
		plugin.F()
	}
}

func readFilesFromDir(path string) []string {
	outputDirRead, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		panic(err)
	}

	var plugins_files []string

	// one by one, add the plugin path to the `plugins_files` variable
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]
		plugins_files = append(plugins_files, fmt.Sprintf("%s/%s", path, outputFileHere.Name()))
	}

	return plugins_files
}
