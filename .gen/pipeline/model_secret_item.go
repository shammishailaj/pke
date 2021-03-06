/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
import (
	"time"
)

type SecretItem struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Version int32 `json:"version,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Values map[string]interface{} `json:"values,omitempty"`
}
