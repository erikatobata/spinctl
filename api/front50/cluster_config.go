/*
 * Spinnaker Front50 API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package front50

type ClusterConfig struct {

	Account string `json:"account,omitempty"`

	Applications []string `json:"applications,omitempty"`

	Detail string `json:"detail,omitempty"`

	MetaClass *MetaClass `json:"metaClass,omitempty"`

	Stack string `json:"stack,omitempty"`
}