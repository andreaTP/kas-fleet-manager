/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CloudRegionList struct for CloudRegionList
type CloudRegionList struct {
	Kind  string        `json:"kind"`
	Page  int32         `json:"page"`
	Size  int32         `json:"size"`
	Total int32         `json:"total"`
	Items []CloudRegion `json:"items"`
}
