/*
 * Spinnaker Front50 API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package front50

type EntityTag struct {

	Category string `json:"category,omitempty"`

	MetaClass *MetaClass `json:"metaClass,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Value *interface{} `json:"value,omitempty"`

	ValueType string `json:"valueType,omitempty"`
}
