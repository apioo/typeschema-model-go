package model

// The root object of a TypeSchema document containing imports, definitions, and the entry point.
type TypeSchema struct {
    Definitions map[string]DefinitionType `json:"definitions"`
    Import map[string]string `json:"import"`
    Root string `json:"root"`
}

