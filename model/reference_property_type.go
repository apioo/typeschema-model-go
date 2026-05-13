package model

// A reference to a defined type in the global 'definitions' map.
type ReferencePropertyType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Nullable bool `json:"nullable"`
    Type string `json:"type"`
    Target string `json:"target"`
    Template map[string]string `json:"template"`
}

