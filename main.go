package main

type Schema struct {
	Swagger             string                 `json:"swagger"`
	Info                *Info                  `json:"info"`
	Host                string                 `json:"host,omitempty"`
	BasePath            string                 `json:"basePath,omitempty"`
	Schemes             []string               `json:"schemes,omitempty"`
	Consumes            []string               `json:"consumes,omitempty"`
	Produces            []string               `json:"produces,omitempty"`
	Paths               Paths                  `json:"paths"`
	Definitions         interface{}            `json:"definitions,omitempty"`
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

type Paths map[string]*PathItem

type PathItem struct {
	Get        *Operation `json:"get,omitempty"`
	Put        *Operation `json:"put,omitempty"`
	Post       *Operation `json:"post,omitempty"`
	Delete     *Operation `json:"delete,omitempty"`
	Options    *Operation `json:"options,omitempty"`
	Head       *Operation `json:"head,omitempty"`
	Patch      *Operation `json:"patch,omitempty"`
	Parameters *Operation `json:"parameters,omitempty"`
}

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

type Parameter interface{}

type ParametersDefinitions map[string]*Parameter

type ResponsesDefinitions map[string]*Response

type Response struct {
	Description string                 `json:"description"`
	Schema      interface{}            `json:"schema,omitempty"`
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

type Scopes map[string]string

type SecurityRequirement map[string][]string

type Tag struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

// TODO: main entry point should NOT be required. Need to refer to the
// documentation on how to structure sub packages.
func main() {}
