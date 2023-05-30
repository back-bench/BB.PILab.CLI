package main

import (
	"fmt"

	"github.com/back-bench/BB.PILab.CLI/pi/command"
	"github.com/back-bench/BB.PILab.CLI/pi/command/commands"
	"github.com/spf13/cobra"
)

func newCliCommand(piCli *command.PiCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pi",
		Short: "PI lab CLI services.",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			//need to do som validation
			if len(args) == 0 {
				fmt.Println("Need help pi")
				//return command.ShowHelp(amzCli.Err())(cmd, args)
			}
			return fmt.Errorf("invalid command", args[0])
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			//run validation
			return isSupported()
		},
	}
	commands.AddCommands(cmd, piCli)
	return cmd
}

func isSupported() error {
	//return errors.New("Pre Run is running")
	return nil
}

func runPi(piCli *command.PiCli) error {
	cmd := newCliCommand(piCli)
	return cmd.Execute()
}
