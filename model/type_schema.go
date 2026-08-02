package model

// The root document object containing namespace imports, type definitions, and the root entry point.
type TypeSchema struct {
    Definitions map[string]DefinitionType `json:"definitions"`
    Import map[string]string `json:"import"`
    Root string `json:"root"`
}

