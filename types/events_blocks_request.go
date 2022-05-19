// Copyright 2022 Klaytn
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

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// EventsBlocksRequest EventsBlocksRequest is utilized to fetch a sequence of BlockEvents indicating
// which blocks were added and removed from storage to reach the current state.
type EventsBlocksRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	// offset is the offset into the event stream to sync events from. If this field is not
	// populated, we return the limit events backwards from tip. If this is set to 0, we start from
	// the beginning.
	Offset *int64 `json:"offset,omitempty"`
	// limit is the maximum number of events to fetch in one call. The implementation may return <=
	// limit events.
	Limit *int64 `json:"limit,omitempty"`
}
