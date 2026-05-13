package model

// Represents a sequence of characters, optionally following a specific format.
type StringPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Default string `json:"default"`
    Format string `json:"format"`
}

