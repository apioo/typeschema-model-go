package model

// Represents a property containing a list of items that share the same schema.
type ArrayPropertyType struct {
    Schema *PropertyType `json:"schema"`
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
}

