package model

// A property containing a list of items of a consistent type.
type ArrayPropertyType struct {
    Schema *PropertyType `json:"schema"`
    Type string `json:"type"`
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
}

