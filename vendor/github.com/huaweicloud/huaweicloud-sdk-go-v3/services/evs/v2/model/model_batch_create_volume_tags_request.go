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

// Request Object
type BatchCreateVolumeTagsRequest struct {
	VolumeId string                            `json:"volume_id"`
	Body     *BatchCreateVolumeTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateVolumeTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateVolumeTagsRequest", string(data)}, " ")
}
