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

// A metadata detail record for a specific image. Multiple detail records may map a single catalog image.
type ImageDetail struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	// Full docker-pullalbe tag string referencing the image
	Fulltag string `json:"fulltag,omitempty"`
	// Full docker-pullable digest string including the registry url and repository necessary get the image
	Fulldigest string `json:"fulldigest,omitempty"`
	UserId string `json:"userId,omitempty"`
	ImageId string `json:"imageId,omitempty"`
	Registry string `json:"registry,omitempty"`
	Repo string `json:"repo,omitempty"`
	Dockerfile string `json:"dockerfile,omitempty"`
	// The parent Anchore Image record to which this detail maps
	ImageDigest string `json:"imageDigest,omitempty"`
}
