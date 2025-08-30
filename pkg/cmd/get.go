/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"io"

	"github.com/spf13/cobra"

	"helm.sh/helm/v4/pkg/action"
	"helm.sh/helm/v4/pkg/cli"
	"helm.sh/helm/v4/pkg/cmd/require"
)

var getHelp = `
This command consists of multiple subcommands which can be used to
get extended information about the release, including:

- The values used to generate the release
- The generated manifest file
- The notes provided by the chart of the release
- The hooks associated with the release
- The metadata of the release
`

func newGetCmd(settings *cli.EnvSettings, cfg *action.Configuration, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "download extended information of a named release",
		Long:  getHelp,
		Args:  require.NoArgs,
	}

	cmd.AddCommand(newGetAllCmd(settings, cfg, out))
	cmd.AddCommand(newGetValuesCmd(settings, cfg, out))
	cmd.AddCommand(newGetManifestCmd(settings, cfg, out))
	cmd.AddCommand(newGetHooksCmd(settings, cfg, out))
	cmd.AddCommand(newGetNotesCmd(settings, cfg, out))
	cmd.AddCommand(newGetMetadataCmd(settings, cfg, out))

	return cmd
}
