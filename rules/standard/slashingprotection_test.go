// Copyright © 2020, 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package standard_test

import (
	"context"
	"os"
	"testing"

	"github.com/attestantio/dirk/rules"
	standardrules "github.com/attestantio/dirk/rules/standard"
	"github.com/stretchr/testify/require"
)

func TestExportSlashingProtection(t *testing.T) {
	ctx := context.Background()
	base, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(base)
	service, err := standardrules.New(ctx,
		standardrules.WithStoragePath(base),
	)
	require.NoError(t, err)

	// Check that empty store returns a good result.
	export, err := service.ExportSlashingProtection(ctx)
	require.NoError(t, err)
	require.Len(t, export, 0)

	// Check an attestation; will add entry.
	var pubKey [48]byte
	copy(pubKey[:], []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
	})
	result := service.OnSignBeaconAttestation(ctx,
		&rules.ReqMetadata{
			PubKey: pubKey[:],
		},
		&rules.SignBeaconAttestationData{
			Domain: []byte{
				0x01, 0x00, 0x00, 0x00, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			},
			Slot:           0x7fffffffffffffff,
			CommitteeIndex: 0,
			BeaconBlockRoot: []byte{
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			},
			Source: &rules.Checkpoint{
				Epoch: 0x7ffffffffffffffd,
				Root: []byte{
					0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
					0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
				},
			},
			Target: &rules.Checkpoint{
				Epoch: 0x7ffffffffffffffe,
				Root: []byte{
					0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
					0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
				},
			},
		},
	)
	require.Equal(t, rules.APPROVED, result)

	export, err = service.ExportSlashingProtection(ctx)
	require.NoError(t, err)
	require.Len(t, export, 1)
	require.Equal(t, int64(-1), export[pubKey].HighestProposedSlot)
	require.Equal(t, int64(0x7ffffffffffffffd), export[pubKey].HighestAttestedSourceEpoch)
	require.Equal(t, int64(0x7ffffffffffffffe), export[pubKey].HighestAttestedTargetEpoch)

	// Check a proposal; will add entry.
	result = service.OnSignBeaconProposal(ctx,
		&rules.ReqMetadata{
			PubKey: pubKey[:],
		},
		&rules.SignBeaconProposalData{
			Domain: []byte{
				0x00, 0x00, 0x00, 0x00, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			},
			Slot:          0x7fffffffffffffff,
			ProposerIndex: 0,
			ParentRoot: []byte{
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			},
			StateRoot: []byte{
				0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
				0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
			},
			BodyRoot: []byte{
				0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
				0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
			},
		},
	)
	require.Equal(t, rules.APPROVED, result)

	export, err = service.ExportSlashingProtection(ctx)
	require.NoError(t, err)
	require.Len(t, export, 1)
	require.Equal(t, export[pubKey].HighestProposedSlot, int64(0x7fffffffffffffff))
	require.Equal(t, export[pubKey].HighestAttestedSourceEpoch, int64(0x7ffffffffffffffd))
	require.Equal(t, export[pubKey].HighestAttestedTargetEpoch, int64(0x7ffffffffffffffe))
}

func TestImportSlashingProtection(t *testing.T) {
	ctx := context.Background()
	base, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(base)
	service, err := standardrules.New(ctx,
		standardrules.WithStoragePath(base),
	)
	require.NoError(t, err)

	// Import some data.
	protection := make(map[[48]byte]*rules.SlashingProtection)
	key1 := [48]byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
	}
	protection[key1] = &rules.SlashingProtection{
		HighestProposedSlot:        0x7fffffffffffffff,
		HighestAttestedSourceEpoch: -1,
		HighestAttestedTargetEpoch: -1,
	}
	err = service.ImportSlashingProtection(ctx, protection)
	require.NoError(t, err)

	// Ensure the data is present.
	export, err := service.ExportSlashingProtection(ctx)
	require.NoError(t, err)
	require.Len(t, export, 1)
	require.Equal(t, int64(0x7fffffffffffffff), export[key1].HighestProposedSlot)
	require.Equal(t, int64(-1), export[key1].HighestAttestedSourceEpoch)
	require.Equal(t, int64(-1), export[key1].HighestAttestedTargetEpoch)

	// Import some more data.
	protection = make(map[[48]byte]*rules.SlashingProtection)
	key2 := [48]byte{
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
		0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
	}
	protection[key2] = &rules.SlashingProtection{
		HighestProposedSlot:        -1,
		HighestAttestedSourceEpoch: 0x0102030405060708,
		HighestAttestedTargetEpoch: 0x0203040506070809,
	}
	err = service.ImportSlashingProtection(ctx, protection)
	require.NoError(t, err)

	// Ensure the data is present.
	export, err = service.ExportSlashingProtection(ctx)
	require.NoError(t, err)
	require.Len(t, export, 2)
	// Ensure old data still present.
	require.Equal(t, int64(0x7fffffffffffffff), export[key1].HighestProposedSlot)
	require.Equal(t, int64(-1), export[key1].HighestAttestedSourceEpoch)
	require.Equal(t, int64(-1), export[key1].HighestAttestedTargetEpoch)
	// Ensure new data present.
	require.Equal(t, int64(-1), export[key2].HighestProposedSlot)
	require.Equal(t, int64(0x0102030405060708), export[key2].HighestAttestedSourceEpoch)
	require.Equal(t, int64(0x0203040506070809), export[key2].HighestAttestedTargetEpoch)
}
