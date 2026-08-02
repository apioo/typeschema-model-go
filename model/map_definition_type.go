package model

// Represents a key-value map with dynamic key names where all values conform to the same schema.
type MapDefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

