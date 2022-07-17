// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// Ilm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/xpack/usage/types.ts#L153-L156
type Ilm struct {
	PolicyCount int                   `json:"policy_count"`
	PolicyStats []IlmPolicyStatistics `json:"policy_stats"`
}

// IlmBuilder holds Ilm struct and provides a builder API.
type IlmBuilder struct {
	v *Ilm
}

// NewIlm provides a builder for the Ilm struct.
func NewIlmBuilder() *IlmBuilder {
	r := IlmBuilder{
		&Ilm{},
	}

	return &r
}

// Build finalize the chain and returns the Ilm struct
func (rb *IlmBuilder) Build() Ilm {
	return *rb.v
}

func (rb *IlmBuilder) PolicyCount(policycount int) *IlmBuilder {
	rb.v.PolicyCount = policycount
	return rb
}

func (rb *IlmBuilder) PolicyStats(policy_stats []IlmPolicyStatisticsBuilder) *IlmBuilder {
	tmp := make([]IlmPolicyStatistics, len(policy_stats))
	for _, value := range policy_stats {
		tmp = append(tmp, value.Build())
	}
	rb.v.PolicyStats = tmp
	return rb
}