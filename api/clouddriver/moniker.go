/*
 * clouddriver
 *
 * Cloud read and write operations
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddriver

type Moniker struct {

	App string `json:"app,omitempty"`

	Cluster string `json:"cluster,omitempty"`

	Detail string `json:"detail,omitempty"`

	Sequence int32 `json:"sequence,omitempty"`

	Stack string `json:"stack,omitempty"`
}