/*
 * Open DAM API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package opendamclient
// Version version data of an asset
type Version struct {
	// The version number
	Number float32 `json:"number"`
	// A point in time represented as milliseconds from the Epoch (UTC)
	Timestamp float32 `json:"timestamp"`
	JobId string `json:"job_id"`
}
