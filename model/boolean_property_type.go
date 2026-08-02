package model

// Represents a boolean true or false value.
type BooleanPropertyType struct {
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
}

