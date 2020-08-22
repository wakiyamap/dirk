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

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Service) setupReadyMetrics() error {
	s.ready = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "dirk",
		Name:      "ready",
		Help:      "1 if Dirk is ready to serve requests, otherwise 0.",
	})
	if err := prometheus.Register(s.ready); err != nil {
		return err
	}

	return nil
}

// Ready is called when the service is ready to serve requests, or when it stops being so.
func (s *Service) Ready(ready bool) {
	if ready {
		s.ready.Set(1)
	} else {
		s.ready.Set(0)
	}
}
