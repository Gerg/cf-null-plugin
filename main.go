package main

import (
	"code.cloudfoundry.org/cli/plugin"
)

type NullPlugin struct{}

func (c *NullPlugin) Run(cliConnection plugin.CliConnection, args []string) {}

func (c *NullPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "NullPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "null-plugin",
				HelpText: "Plugin that does nothing",
			},
		},
	}
}

func main() {
	plugin.Start(new(NullPlugin))
}
