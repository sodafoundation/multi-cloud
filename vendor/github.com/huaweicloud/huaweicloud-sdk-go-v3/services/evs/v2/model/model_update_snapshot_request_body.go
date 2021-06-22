/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type UpdateSnapshotRequestBody struct {
	Snapshot *UpdateSnapshotOption `json:"snapshot"`
}

func (o UpdateSnapshotRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSnapshotRequestBody", string(data)}, " ")
}
