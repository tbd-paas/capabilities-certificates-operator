/*
Copyright 2024.

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

package capabilities

import (
	"github.com/spf13/cobra"

	cmdversion "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/version"

	"github.com/tbd-paas/capabilities-certificates-operator/apis/capabilities"
)

// NewCertificatesCapabilitySubCommand creates a new command and adds it to its
// parent command.
func NewCertificatesCapabilitySubCommand(parentCommand *cobra.Command) {
	versionCmd := &cmdversion.VersionSubCommand{
		Name:         "version",
		Description:  "display the version information",
		VersionFunc:  VersionCertificatesCapability,
		SubCommandOf: parentCommand,
	}

	versionCmd.Setup()
}

func VersionCertificatesCapability(v *cmdversion.VersionSubCommand) error {
	apiVersions := make([]string, len(capabilities.CertificatesCapabilityGroupVersions()))

	for i, groupVersion := range capabilities.CertificatesCapabilityGroupVersions() {
		apiVersions[i] = groupVersion.Version
	}

	versionInfo := cmdversion.VersionInfo{
		CLIVersion:  cmdversion.CLIVersion,
		APIVersions: apiVersions,
	}

	return versionInfo.Display()
}
