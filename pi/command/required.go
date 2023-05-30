package command

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// NoArgs validates args and returns an error if there are any args
func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}

	if cmd.HasSubCommands() {
		return errors.Errorf("\n" + strings.TrimRight(cmd.UsageString(), "\n"))
	}

	return errors.Errorf(
		"%q accepts no arguments.\nSee '%s --help'.\n\nUsage:  %s\n\n%s",
		cmd.CommandPath(),
		cmd.CommandPath(),
		cmd.UseLine(),
		cmd.Short,
	)
}

// RequiresMaxArgs maximum nuber of argument required.
func RequiresMaxArgs(max int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) <= max {
			return nil
		}
		return errors.Errorf(
			"%q requires at most %d %s.\nSee '%s --help'.\n\nUsage:  %s\n\n%s",
			cmd.CommandPath(),
			max,
			pluralize("argument", max),
			cmd.CommandPath(),
			cmd.UseLine(),
			cmd.Short,
		)
	}
}

// RequiresArgsInRange requires argument in between given range
func RequiresArgsInRange(min int, max int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) >= min && len(args) <= max {
			return nil
		}
		return errors.Errorf(
			"%q requires argument in between %d and %d %s.\nSee '%s --help'.\n\nUsage:  %s\n\n%s",
			cmd.CommandPath(),
			min,
			max,
			pluralize("argument", max),
			cmd.CommandPath(),
			cmd.UseLine(),
			cmd.Short,
		)
	}
}

func pluralize(word string, number int) string {
	if number == 1 {
		return word
	}
	return word + "s"
}
