/*
 * Kayenta API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package kayenta

type ExpectedArtifact struct {

	BoundArtifact *Artifact `json:"boundArtifact,omitempty"`

	DefaultArtifact *Artifact `json:"defaultArtifact,omitempty"`

	Id string `json:"id,omitempty"`

	MatchArtifact *Artifact `json:"matchArtifact,omitempty"`

	UseDefaultArtifact bool `json:"useDefaultArtifact,omitempty"`

	UsePriorArtifact bool `json:"usePriorArtifact,omitempty"`
}