/*
 * clouddriver
 *
 * Cloud read and write operations
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddriver

type ClusterViewModel struct {

	Account string `json:"account,omitempty"`

	LoadBalancers []string `json:"loadBalancers,omitempty"`

	Moniker *Moniker `json:"moniker,omitempty"`

	Name string `json:"name,omitempty"`

	ServerGroups []string `json:"serverGroups,omitempty"`
}