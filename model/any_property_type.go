package model

// Represents a wildcard property that accepts any valid JSON value (object, array, string, number, boolean, or null).
type AnyPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
}

