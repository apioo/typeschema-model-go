package model

// An object with a dynamic set of keys where every value conforms to the same schema.
type MapDefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

