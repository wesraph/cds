package main

import (
	"github.com/spf13/cobra"

	"github.com/ovh/cds/cli"
	"github.com/ovh/cds/sdk"
)

var projectMetadataCmd = cli.Command{
	Name:  "metadata",
	Short: "Manage CDS project metadata",
}

func projectMetadata() *cobra.Command {
	return cli.NewCommand(projectMetadataCmd, nil, []*cobra.Command{
		cli.NewCommand(projectMetadataSetCmd, projectSetMetadataRun, nil, withAllCommandModifiers()...),
		cli.NewListCommand(projectMetadataListCmd, projectListMetadataRun, nil, withAllCommandModifiers()...),
		cli.NewCommand(projectMetadataDeleteCmd, projectDeleteMetadataRun, nil, withAllCommandModifiers()...),
	})
}

var projectMetadataSetCmd = cli.Command{
	Name:  "set",
	Short: "Add a new metadata on project. cds project metadata set <projectKey> <key=value> [key=value [...]]",
	Ctx: []cli.Arg{
		{Name: _ProjectKey},
	},
	Args: []cli.Arg{
		{Name: "key-name"},
		{Name: "key-value"},
	},
}

func projectSetMetadataRun(v cli.Values) error {
	key := &sdk.ProjectKey{
		Key: sdk.Key{
			Name: v["key-name"],
			Type: v["key-value"],
		},
	}
	return client.ProjectMetadataSet(v[_ProjectKey], key)
}

var projectMetadataListCmd = cli.Command{
	Name:  "list",
	Short: "List CDS project metadata",
	Args: []cli.Arg{
		{Name: _ProjectKey},
	},
}

func projectListMetadataRun(v cli.Values) (cli.ListResult, error) {
	metadatas, err := client.ProjectMetadataList(v[_ProjectKey])
	if err != nil {
		return nil, err
	}
	return cli.AsListResult(metadatas), nil
}

var projectMetadataDeleteCmd = cli.Command{
	Name:  "delete",
	Short: "Delete CDS project metadata",
	Ctx: []cli.Arg{
		{Name: _ProjectKey},
	},
	Args: []cli.Arg{
		{Name: "key-name"},
	},
}

func projectDeleteMetadataRun(v cli.Values) error {
	return client.ProjectMetadataDelete(v[_ProjectKey], v["key-name"])
}
