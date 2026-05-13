package model

// The base abstract type for all property definitions within a struct or collection.
type PropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
}

