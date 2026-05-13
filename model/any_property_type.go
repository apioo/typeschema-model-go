package model

// A wildcard property that accepts any valid JSON value (object, array, string, etc.).
type AnyPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
}

