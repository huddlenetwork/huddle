package app

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func TestSimAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewDesmosApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{},
		DefaultNodeHome, 0, MakeTestEncodingConfig(), simapp.EmptyAppOptions{}, wasm.EnableAllProposals,
	)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewDesmosApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{},
		DefaultNodeHome, 0, MakeTestEncodingConfig(), simapp.EmptyAppOptions{}, wasm.EnableAllProposals,
	)
	_, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}
