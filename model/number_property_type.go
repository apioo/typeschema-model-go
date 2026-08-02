package model

// Represents a numeric value, including floating-point and decimal numbers.
type NumberPropertyType struct {
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
}

