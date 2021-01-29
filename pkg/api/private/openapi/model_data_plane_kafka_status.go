/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DataPlaneKafkaStatus Schema of the status object for a Kafka cluster
type DataPlaneKafkaStatus struct {
	// The status conditions of a Kafka cluster
	Conditions []DataPlaneClusterUpdateStatusRequestConditions `json:"conditions,omitempty"`
	Capacity   DataPlaneKafkaStatusCapacity                    `json:"capacity,omitempty"`
	Versions   DataPlaneKafkaStatusVersions                    `json:"versions,omitempty"`
}
