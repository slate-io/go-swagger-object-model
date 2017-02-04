package models

import (
	"strings"
)

type RootDocument struct {
	swagger             string                 `json:"-"`
	Info                *Info                  `json:"info"`
	Host                string                 `json:"host,omitempty"`
	BasePath            string                 `json:"basePath,omitempty"`
	Schemes             []string               `json:"schemes,omitempty"`
	Consumes            []string               `json:"consumes,omitempty"`
	Produces            []string               `json:"produces,omitempty"`
	Paths               Paths                  `json:"paths"`
	Definitions         Definitions            `json:"definitions,omitempty"`
	Parameters          ParametersDefinitions  `json:"parameters,omitempty"`
	Responses           ResponsesDefinitions   `json:"responses,omitempty"`
	SecurityDefinitions SecurityDefinitions    `json:"securityDefinitions,omitempty"`
	Security            SecurityRequirement    `json:"security,omitempty"`
	Tags                []*Tag                 `json:"tags,omitempty"`
	ExternalDocs        *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Info struct {
	Title          string   `json:"title"`
	Description    string   `json:"description,omitempty"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
	Version        string   `json:"version"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url,omitempty"`
}

type Paths map[string]PathItem

type PathItem map[string]*Operation

type Operation struct {
	Tags         []string               `json:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationId  string                 `json:"operationId,omitempty"`
	Consumes     []string               `json:"consumes,omitempty"`
	Parameters   []*Parameter           `json:"parameters,omitempty"`
	Responses    ResponsesDefinitions   `json:"responses"`
	Schemes      []string               `json:"schemes,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty"`
	Security     interface{}            `json:"security,omitempty"`
}

type ExternalDocumentation struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url"`
}

type ParametersDefinitions map[string]*Parameter

type ResponsesDefinitions map[string]*Response

type Response struct {
	Description string                 `json:"description"`
	Schema      *Schema                `json:"schema,omitempty"`
	Headers     map[string]*Header     `json:"headers,omitempty"`
	Examples    map[string]interface{} `json:"examples,omitempty"`
}

type Header struct {
	Description      string      `json:"description,omitempty"`
	Type             string      `json:"type"`
	Format           string      `json:"format,omitempty"`
	Items            interface{} `json:"items,omitempty"`
	CollectionFormat string      `json:"collectionFormat,omitempty"`
	Default          interface{} `json:"default,omitempty"`
	// TODO: Implement JSON Schema Validation
}

type SecurityDefinitions map[string]*SecurityScheme

type SecurityScheme struct {
	Type             string `json:"type"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name,omitempty"`
	In               string `json:"in,omitempty"`
	Flow             string `json:"flow,omitempty"`
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`
	TokenUrl         string `json:"tokenUrl,omitempty"`
	Scopes           Scopes `json:"scopes,omitempty"`
}

type Parameter struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
}

type Scopes map[string]string

type SecurityRequirement map[string][]string

type Tag struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

type Xml struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

type Definitions map[string]*Schema

type Properties map[string]*Schema

type Schema struct {
	Reference            string                 `json:"$ref,omitempty"`
	Format               string                 `json:"format,omitempty"`
	Title                string                 `json:"title,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Default              interface{}            `json:"default,omitempty"`
	MultipleOf           uint64                 `json:"multipleOf,omitempty"`
	Maximum              int64                  `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                   `json:"exclusiveMaximum,omitempty"`
	Minimum              int64                  `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                   `json:"exclusiveMinimum,omitempty"`
	MaxLength            uint64                 `json:"maxLength,omitempty"`
	MinLength            uint64                 `json:"minLength,omitempty"`
	Pattern              string                 `json:"pattern,omitempty"`
	MaxItems             uint64                 `json:"maxItems,omitempty"`
	MinItems             uint64                 `json:"minItems,omitempty"`
	UniqueItems          bool                   `json:"uniqueItems,omitempty"`
	MaxProperties        uint64                 `json:"maxProperties,omitempty"`
	MinProperties        uint64                 `json:"minProperties,omitempty"`
	Required             []string               `json:"required,omitempty"`
	Enum                 []interface{}          `json:"enum,omitempty"`
	Type                 interface{}            `json:"type,omitempty"`
	Items                interface{}            `json:"items,omitempty"`
	AllOf                interface{}            `json:"allOf,omitempty"`
	Properties           Properties             `json:"properties,omitempty"`
	AdditionalProperties interface{}            `json:"additionalProperties,omitempty"`
	Discriminator        string                 `json:"discriminator,omitempty"`
	ReadOnly             bool                   `json:"readOnly,omitempty"`
	Xml                  *Xml                   `json:"xml,omitempty"`
	ExternalDocs         *ExternalDocumentation `json:"externalDocs,omitempty"`
	Example              interface{}            `json:"example,omitempty"`
}

func (s *Schema) GetReference() string {
	v := strings.Split(s.Reference, "/")
	return v[len(v)-1]
}
