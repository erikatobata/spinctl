/*
 * Spinnaker Front50 API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package front50

type Application struct {

	CreateTs string `json:"createTs,omitempty"`

	Description string `json:"description,omitempty"`

	Email string `json:"email,omitempty"`

	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	MetaClass *MetaClass `json:"metaClass,omitempty"`

	Name string `json:"name,omitempty"`

	UpdateTs string `json:"updateTs,omitempty"`
}