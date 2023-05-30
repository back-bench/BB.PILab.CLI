package commands

import (
	"os"

	"github.com/back-bench/BB.PILab.CLI/pi/command"
	initpi "github.com/back-bench/BB.PILab.CLI/pi/command/init"
	"github.com/spf13/cobra"
)

// AddCommands adds all the commands from cli/command to the root command
func AddCommands(cmd *cobra.Command, piCli command.Cli) {
	cmd.AddCommand(
		initpi.PIInitCommand(piCli),
		//hide(env.EnvsCommand(amzcli)),
	)
}

func hide(cmd *cobra.Command) *cobra.Command {
	// If the environment variable with name "QUANTAM_LEGACY_CMD" is not empty,
	if os.Getenv("QUANTAM_LEGACY_CMD") == "" {
		return cmd
	}
	cmdCopy := *cmd
	cmdCopy.Hidden = true
	cmdCopy.Aliases = []string{}
	return &cmdCopy
}
