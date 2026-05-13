package model

// A property containing a map of dynamic keys to a consistent value type.
type MapPropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Schema *PropertyType `json:"schema"`
}

