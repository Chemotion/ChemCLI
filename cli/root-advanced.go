package cli

import (
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

var advancedCmdTable = make(cmdTable)

// Backbone for system-related commands
var advancedRootCmd = &cobra.Command{
	Use:       "advanced",
	Short:     "Perform advanced actions related to system and " + nameCLI,
	Args:      cobra.NoArgs,
	ValidArgs: maps.Keys(advancedCmdTable),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("selected-instance").Changed {
			zboth.Warn().Msgf("You used `-i` flag. It has no effect on `%s advanced` command.", commandForCLI)
		} // RETURN
		isInteractive(true)
		acceptedOpts := []string{"info - about this system", "update - " + nameCLI}
		advancedCmdTable["info - about this system"] = infoAdvancedRootCmd.Run
		advancedCmdTable["update - "+nameCLI] = updateSelfAdvancedRootCmd.Run
		if existingFile(conf.ConfigFileUsed()) {
			acceptedOpts = append(acceptedOpts, "uninstall")
			advancedCmdTable["uninstall"] = uninstallAdvancedRootCmd.Run
		}
		if cmd.CalledAs() == "advanced" {
			acceptedOpts = append(acceptedOpts, "exit")
		} else {
			acceptedOpts = append(acceptedOpts, []string{"back", "exit"}...)
			advancedCmdTable["back"] = cmd.Run
		}
		advancedCmdTable[selectOpt(acceptedOpts, "")](cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(advancedRootCmd)
}
