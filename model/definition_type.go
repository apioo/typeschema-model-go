package model

// The abstract base type for all schema definitions. It provides common metadata such as descriptions and deprecation status.
type DefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
}

