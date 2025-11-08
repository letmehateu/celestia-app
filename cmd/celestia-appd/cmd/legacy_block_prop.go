package cmd

import (
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
)

const flagEnableLegacyBlockProp = "enable-legacy-block-prop"

// overrideLegacyBlockProp updates the in-memory consensus configuration when
// the --enable-legacy-block-prop flag is explicitly provided. Without the flag,
// the existing value from config.toml is preserved.
func overrideLegacyBlockProp(cmd *cobra.Command, logger log.Logger) error {
	// Allow operators to bypass overrides entirely.
	if bypass, err := cmd.Flags().GetBool(bypassOverridesFlagKey); err == nil && bypass {
		logger.Info("Bypassing legacy block propagation override due to flag")
		return nil
	}

	// If the flag is not registered on this command, there's nothing to do.
	if cmd.Flags().Lookup(flagEnableLegacyBlockProp) == nil {
		return nil
	}

	// Only override the configuration if the operator explicitly set the flag.
	if !cmd.Flags().Changed(flagEnableLegacyBlockProp) {
		return nil
	}

	enable, err := cmd.Flags().GetBool(flagEnableLegacyBlockProp)
	if err != nil {
		return err
	}

	sctx := server.GetServerContextFromCmd(cmd)
	if sctx == nil || sctx.Config == nil {
		return nil
	}

	sctx.Config.Consensus.EnableLegacyBlockProp = enable

	if enable {
		logger.Info("Enabled legacy block propagation via flag")
	} else {
		logger.Info("Disabled legacy block propagation via flag")
	}

	return nil
}

