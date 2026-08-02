package model

// Represents a property containing a key-value map where all values share the same schema.
type MapPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

