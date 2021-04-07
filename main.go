package main

import (
	"code.cloudfoundry.org/cli/plugin"
)

type NullPlugin struct{}

func (c *NullPlugin) Run(cliConnection plugin.CliConnection, args []string) {}

func (c *NullPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "NullPlugin v7",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 7,
			Minor: 1,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "null-plugin v7",
				HelpText: "Plugin that does nothing on v7",
			},
		},
	}
}

func main() {
	plugin.Start(new(NullPlugin))
}
