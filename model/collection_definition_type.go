package model

// Abstract base for definitions that hold multiple values of a single type, such as arrays or maps.
type CollectionDefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

