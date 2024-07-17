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

package commands

import (
	"github.com/spf13/cobra"

	// common imports for subcommands
	cmdgenerate "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/generate"
	cmdinit "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/init"
	cmdversion "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/version"

	// specific imports for workloads
	generatecertificates "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/generate/certificates"
	initcertificates "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/init/certificates"
	versioncertificates "github.com/tbd-paas/capabilities-certificates-operator/cmd/certificatesctl/commands/version/certificates"
	// +kubebuilder:scaffold:operator-builder:subcommands:imports
)

// CertificatesctlCommand represents the base command when called without any subcommands.
type CertificatesctlCommand struct {
	*cobra.Command
}

// NewCertificatesctlCommand returns an instance of the CertificatesctlCommand.
func NewCertificatesctlCommand() *CertificatesctlCommand {
	c := &CertificatesctlCommand{
		Command: &cobra.Command{
			Use:   "certificatesctl",
			Short: "Manage the certificates capability",
			Long:  "Manage the certificates capability",
		},
	}

	c.addSubCommands()

	return c
}

// Run represents the main entry point into the command
// This is called by main.main() to execute the root command.
func (c *CertificatesctlCommand) Run() {
	cobra.CheckErr(c.Execute())
}

func (c *CertificatesctlCommand) newInitSubCommand() {
	parentCommand := cmdinit.GetParent(c.Command)
	_ = parentCommand

	// add the init subcommands
	initcertificates.NewTrustManagerSubCommand(parentCommand)
	initcertificates.NewCertManagerSubCommand(parentCommand)
	// +kubebuilder:scaffold:operator-builder:subcommands:init
}

func (c *CertificatesctlCommand) newGenerateSubCommand() {
	parentCommand := cmdgenerate.GetParent(c.Command)
	_ = parentCommand

	// add the generate subcommands
	generatecertificates.NewTrustManagerSubCommand(parentCommand)
	generatecertificates.NewCertManagerSubCommand(parentCommand)
	// +kubebuilder:scaffold:operator-builder:subcommands:generate
}

func (c *CertificatesctlCommand) newVersionSubCommand() {
	parentCommand := cmdversion.GetParent(c.Command)
	_ = parentCommand

	// add the version subcommands
	versioncertificates.NewTrustManagerSubCommand(parentCommand)
	versioncertificates.NewCertManagerSubCommand(parentCommand)
	// +kubebuilder:scaffold:operator-builder:subcommands:version
}

// addSubCommands adds any additional subCommands to the root command.
func (c *CertificatesctlCommand) addSubCommands() {
	c.newInitSubCommand()
	c.newGenerateSubCommand()
	c.newVersionSubCommand()
}
