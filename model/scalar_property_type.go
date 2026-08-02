package model

// The abstract base type for simple scalar value properties (strings, integers, numbers, booleans).
type ScalarPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
}

