package model

// A placeholder for a type that will be specified at runtime or through template arguments.
type GenericPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Name string `json:"name"`
}

