package model

// Represents an ordered list of elements where every item conforms to the same schema.
type ArrayDefinitionType struct {
    Schema *PropertyType `json:"schema"`
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
}

