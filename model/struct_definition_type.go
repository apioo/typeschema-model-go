package model

// Represents an object with a fixed set of properties (such as a class or record). Supports inheritance and explicit property typing.
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

