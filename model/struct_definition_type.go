package model

// Represents a fixed-structure object (class/record). It supports inheritance and explicit property definitions.
type StructDefinitionType struct {
    Deprecated bool `json:"deprecated"`
    Description string `json:"description"`
    Type string `json:"type"`
    Base bool `json:"base"`
    Discriminator string `json:"discriminator"`
    Mapping map[string]string `json:"mapping"`
    Parent *ReferencePropertyType `json:"parent"`
    Properties map[string]PropertyType `json:"properties"`
}

