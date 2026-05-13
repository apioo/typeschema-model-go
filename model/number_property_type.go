package model

// Represents a floating-point or decimal number.
type NumberPropertyType struct {
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
}

