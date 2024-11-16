package model

// Represents an array which contains a dynamic list of values of the same type
type ArrayDefinitionType struct {
    Schema *PropertyType `json:"schema"`
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
}

