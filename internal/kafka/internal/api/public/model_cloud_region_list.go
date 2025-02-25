/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.14.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// CloudRegionList struct for CloudRegionList
type CloudRegionList struct {
	Kind  string        `json:"kind"`
	Page  int32         `json:"page"`
	Size  int32         `json:"size"`
	Total int32         `json:"total"`
	Items []CloudRegion `json:"items"`
}
