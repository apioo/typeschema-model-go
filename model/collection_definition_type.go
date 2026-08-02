package model

// The abstract base type for collection definitions that contain multiple elements of a uniform type.
type CollectionDefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

