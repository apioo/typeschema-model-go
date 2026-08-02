package model

// Represents a generic placeholder type that is resolved at runtime or via template arguments.
type GenericPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Name string `json:"name"`
}

