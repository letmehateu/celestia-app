package cmd

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	"github.com/celestiaorg/celestia-app/v6/app"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestOverrideLegacyBlockProp_FlagEnables(t *testing.T) {
	cmd, sctx := newLegacyBlockPropCommand()
	require.NoError(t, cmd.Flags().Set(flagEnableLegacyBlockProp, "true"))

	err := overrideLegacyBlockProp(cmd, log.NewNopLogger())
	require.NoError(t, err)

	require.True(t, sctx.Config.Consensus.EnableLegacyBlockProp)
}

func TestOverrideLegacyBlockProp_FlagDisables(t *testing.T) {
	cmd, sctx := newLegacyBlockPropCommand()
	sctx.Config.Consensus.EnableLegacyBlockProp = true
	require.NoError(t, cmd.Flags().Set(flagEnableLegacyBlockProp, "false"))

	err := overrideLegacyBlockProp(cmd, log.NewNopLogger())
	require.NoError(t, err)

	require.False(t, sctx.Config.Consensus.EnableLegacyBlockProp)
}

func TestOverrideLegacyBlockProp_NoFlagChangePreservesConfig(t *testing.T) {
	cmd, sctx := newLegacyBlockPropCommand()
	sctx.Config.Consensus.EnableLegacyBlockProp = true

	err := overrideLegacyBlockProp(cmd, log.NewNopLogger())
	require.NoError(t, err)

	require.True(t, sctx.Config.Consensus.EnableLegacyBlockProp)
}

func TestOverrideLegacyBlockProp_BypassSkipsOverride(t *testing.T) {
	cmd, sctx := newLegacyBlockPropCommand()
	require.NoError(t, cmd.Flags().Set(flagEnableLegacyBlockProp, "true"))
	require.NoError(t, cmd.Flags().Set(bypassOverridesFlagKey, "true"))

	err := overrideLegacyBlockProp(cmd, log.NewNopLogger())
	require.NoError(t, err)

	require.False(t, sctx.Config.Consensus.EnableLegacyBlockProp)
}

func newLegacyBlockPropCommand() (*cobra.Command, *server.Context) {
	cfg := app.DefaultConsensusConfig()

	cmd := &cobra.Command{
		Use: "test",
	}
	cmd.Flags().Bool(flagEnableLegacyBlockProp, false, "")
	cmd.Flags().Bool(bypassOverridesFlagKey, false, "")

	sctx := server.NewDefaultContext()
	sctx.Config = cfg
	sctx.Logger = log.NewNopLogger()

	ctx := context.WithValue(context.Background(), server.ServerContextKey, sctx)
	cmd.SetContext(ctx)

	return cmd, sctx
}


