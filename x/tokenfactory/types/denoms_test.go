package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	appparams "github.com/four4two/merlin/v16/app/params"
	"github.com/four4two/merlin/v16/x/tokenfactory/types"
)

func TestDeconstructDenom(t *testing.T) {
	appparams.SetAddressPrefixes()

	for _, tc := range []struct {
		desc             string
		denom            string
		expectedSubdenom string
		err              error
	}{
		{
			desc:  "empty is invalid",
			denom: "",
			err:   types.ErrInvalidDenom,
		},
		{
			desc:             "normal",
			denom:            "factory/fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
			expectedSubdenom: "bitcoin",
		},
		{
			desc:             "multiple slashes in subdenom",
			denom:            "factory/fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin/1",
			expectedSubdenom: "bitcoin/1",
		},
		{
			desc:             "no subdenom",
			denom:            "factory/fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/",
			expectedSubdenom: "",
		},
		{
			desc:  "incorrect prefix",
			denom: "ibc/fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/bitcoin",
			err:   types.ErrInvalidDenom,
		},
		{
			desc:             "subdenom of only slashes",
			denom:            "factory/fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/////",
			expectedSubdenom: "////",
		},
		{
			desc:  "too long name",
			denom: "factory/fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			err:   types.ErrInvalidDenom,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			expectedCreator := "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44"
			creator, subdenom, err := types.DeconstructDenom(tc.denom)
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, expectedCreator, creator)
				require.Equal(t, tc.expectedSubdenom, subdenom)
			}
		})
	}
}

func TestGetTokenDenom(t *testing.T) {
	appparams.SetAddressPrefixes()
	for _, tc := range []struct {
		desc     string
		creator  string
		subdenom string
		valid    bool
	}{
		{
			desc:     "normal",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
			subdenom: "bitcoin",
			valid:    true,
		},
		{
			desc:     "multiple slashes in subdenom",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
			subdenom: "bitcoin/1",
			valid:    true,
		},
		{
			desc:     "no subdenom",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
			subdenom: "",
			valid:    true,
		},
		{
			desc:     "subdenom of only slashes",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
			subdenom: "/////",
			valid:    true,
		},
		{
			desc:     "too long name",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
			subdenom: "adsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsfadsf",
			valid:    false,
		},
		{
			desc:     "subdenom is exactly max length",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44",
			subdenom: "bitcoinfsadfsdfeadfsafwefsefsefsdfsdafasefsf",
			valid:    true,
		},
		{
			desc:     "creator is exactly max length",
			creator:  "fury1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44jhgjhgkhjklhkjhkjhgjhgjgjghelugt",
			subdenom: "bitcoin",
			valid:    true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := types.GetTokenDenom(tc.creator, tc.subdenom)
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
