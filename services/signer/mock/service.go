// Copyright © 2020 Attestant Limited.
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

package mock

import (
	context "context"

	"github.com/attestantio/dirk/core"
	rules "github.com/attestantio/dirk/rules"
	"github.com/attestantio/dirk/services/checker"
)

// Service is the mock signer service.
type Service struct{}

// New creates a new mock signer service.
func New() *Service {
	return &Service{}
}

// SignGeneric signs generic data.
func (s *Service) SignGeneric(ctx context.Context,
	credentials *checker.Credentials,
	accountName string,
	pubKey []byte,
	data *rules.SignData) (core.Result, []byte) {
	return core.ResultSucceeded, []byte{
		0x90, 0x42, 0xa3, 0x1d, 0xb8, 0x1e, 0x14, 0x65, 0x98, 0xce, 0xd6, 0xe5, 0x6d, 0xff, 0x63, 0x11,
		0xdf, 0xfb, 0x39, 0x52, 0xbc, 0xd0, 0x8f, 0xf9, 0x22, 0x78, 0xad, 0x72, 0x19, 0xb0, 0x69, 0xc9,
		0x86, 0xdb, 0x5d, 0x07, 0x22, 0x01, 0x76, 0xae, 0xd6, 0x1e, 0x6b, 0xe0, 0xc0, 0x52, 0x7f, 0x6d,
		0x0a, 0x16, 0x12, 0x25, 0x62, 0x6e, 0x69, 0xc7, 0xfc, 0x6f, 0xd2, 0xc5, 0x7d, 0x38, 0x99, 0x64,
		0x03, 0xc2, 0x95, 0x70, 0x4b, 0x94, 0xab, 0x7a, 0x36, 0x4c, 0x18, 0x5b, 0x98, 0x34, 0x56, 0xe5,
		0xf9, 0x57, 0x50, 0xd9, 0x0e, 0x92, 0xb1, 0xef, 0x8a, 0x53, 0xd6, 0x3b, 0x3d, 0xf1, 0x91, 0x5a,
	}
}

// SignBeaconAttestation signs a beacon attestation.
func (s *Service) SignBeaconAttestation(ctx context.Context,
	credentials *checker.Credentials,
	accountName string,
	pubKey []byte,
	data *rules.SignBeaconAttestationData) (core.Result, []byte) {
	return core.ResultSucceeded, []byte{
		0x90, 0x42, 0xa3, 0x1d, 0xb8, 0x1e, 0x14, 0x65, 0x98, 0xce, 0xd6, 0xe5, 0x6d, 0xff, 0x63, 0x11,
		0xdf, 0xfb, 0x39, 0x52, 0xbc, 0xd0, 0x8f, 0xf9, 0x22, 0x78, 0xad, 0x72, 0x19, 0xb0, 0x69, 0xc9,
		0x86, 0xdb, 0x5d, 0x07, 0x22, 0x01, 0x76, 0xae, 0xd6, 0x1e, 0x6b, 0xe0, 0xc0, 0x52, 0x7f, 0x6d,
		0x0a, 0x16, 0x12, 0x25, 0x62, 0x6e, 0x69, 0xc7, 0xfc, 0x6f, 0xd2, 0xc5, 0x7d, 0x38, 0x99, 0x64,
		0x03, 0xc2, 0x95, 0x70, 0x4b, 0x94, 0xab, 0x7a, 0x36, 0x4c, 0x18, 0x5b, 0x98, 0x34, 0x56, 0xe5,
		0xf9, 0x57, 0x50, 0xd9, 0x0e, 0x92, 0xb1, 0xef, 0x8a, 0x53, 0xd6, 0x3b, 0x3d, 0xf1, 0x91, 0x5a,
	}
}

// SignBeaconProposal signs a proposal for a beacon block.
func (s *Service) SignBeaconProposal(ctx context.Context,
	credentials *checker.Credentials,
	accountName string,
	pubKey []byte,
	data *rules.SignBeaconProposalData) (core.Result, []byte) {
	return core.ResultSucceeded, []byte{
		0x90, 0x42, 0xa3, 0x1d, 0xb8, 0x1e, 0x14, 0x65, 0x98, 0xce, 0xd6, 0xe5, 0x6d, 0xff, 0x63, 0x11,
		0xdf, 0xfb, 0x39, 0x52, 0xbc, 0xd0, 0x8f, 0xf9, 0x22, 0x78, 0xad, 0x72, 0x19, 0xb0, 0x69, 0xc9,
		0x86, 0xdb, 0x5d, 0x07, 0x22, 0x01, 0x76, 0xae, 0xd6, 0x1e, 0x6b, 0xe0, 0xc0, 0x52, 0x7f, 0x6d,
		0x0a, 0x16, 0x12, 0x25, 0x62, 0x6e, 0x69, 0xc7, 0xfc, 0x6f, 0xd2, 0xc5, 0x7d, 0x38, 0x99, 0x64,
		0x03, 0xc2, 0x95, 0x70, 0x4b, 0x94, 0xab, 0x7a, 0x36, 0x4c, 0x18, 0x5b, 0x98, 0x34, 0x56, 0xe5,
		0xf9, 0x57, 0x50, 0xd9, 0x0e, 0x92, 0xb1, 0xef, 0x8a, 0x53, 0xd6, 0x3b, 0x3d, 0xf1, 0x91, 0x5a,
	}
}
