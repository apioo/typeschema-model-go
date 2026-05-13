package model

// The base abstract type for all schema definitions. It provides metadata common to all types such as descriptions and deprecation status.
type DefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
}

