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

// RegionCapacityListItem schema for a kafka instance type capacity in region
type RegionCapacityListItem struct {
	// kafka instance type
	InstanceType string `json:"instance_type"`
	// list of available Kafka instance sizes that can be created in this region when taking account current capacity and regional limits
	AvailableSizes []string `json:"available_sizes"`
}
