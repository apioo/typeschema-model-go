package model

// Represents a whole number.
type IntegerPropertyType struct {
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
}

