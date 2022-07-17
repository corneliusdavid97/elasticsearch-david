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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// ExtendedStatsBucketAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/pipeline.ts#L155-L157
type ExtendedStatsBucketAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Sigma       *float64             `json:"sigma,omitempty"`
}

// ExtendedStatsBucketAggregationBuilder holds ExtendedStatsBucketAggregation struct and provides a builder API.
type ExtendedStatsBucketAggregationBuilder struct {
	v *ExtendedStatsBucketAggregation
}

// NewExtendedStatsBucketAggregation provides a builder for the ExtendedStatsBucketAggregation struct.
func NewExtendedStatsBucketAggregationBuilder() *ExtendedStatsBucketAggregationBuilder {
	r := ExtendedStatsBucketAggregationBuilder{
		&ExtendedStatsBucketAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedStatsBucketAggregation struct
func (rb *ExtendedStatsBucketAggregationBuilder) Build() ExtendedStatsBucketAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *ExtendedStatsBucketAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *ExtendedStatsBucketAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *ExtendedStatsBucketAggregationBuilder) Format(format string) *ExtendedStatsBucketAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *ExtendedStatsBucketAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *ExtendedStatsBucketAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *ExtendedStatsBucketAggregationBuilder) Meta(meta *MetadataBuilder) *ExtendedStatsBucketAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *ExtendedStatsBucketAggregationBuilder) Name(name string) *ExtendedStatsBucketAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *ExtendedStatsBucketAggregationBuilder) Sigma(sigma float64) *ExtendedStatsBucketAggregationBuilder {
	rb.v.Sigma = &sigma
	return rb
}