package model

// The abstract base type for properties that define inline collections (maps or arrays).
type CollectionPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

