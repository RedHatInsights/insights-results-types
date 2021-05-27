// Copyright 2020, 2021  Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package types contains declaration of various data types (usually structures)
// used across the whole CCX data pipeline.
package types

// OrgID represents organization ID
type OrgID uint32

// ClusterName represents name of cluster in format c8590f31-e97e-4b85-b506-c45ce1911a12
type ClusterName string

// UserID represents type for user id
type UserID string

// ClusterReport represents cluster report
type ClusterReport string

// Timestamp represents any timestamp in a form gathered from database
// TODO: need to be improved
type Timestamp string

// UserVote is a type for user's vote
type UserVote int

// RequestID is used to store the request ID supplied in input Kafka records as
// a unique identifier of payloads. Empty string represents a missing request ID.
type RequestID string

// ErrorKey represents type for error key
type ErrorKey string

// KafkaOffset type for kafka offset
type KafkaOffset int64

// DBDriver type for db driver enum
type DBDriver int

//SchemaVersion is just a constant integer for now, max value 255. If we one day
//need more versions, better consider upgrading to semantic versioning.
type SchemaVersion uint8
