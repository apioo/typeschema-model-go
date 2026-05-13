package model

// Abstract base for simple value types like strings, numbers, and booleans.
type ScalarPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
}

