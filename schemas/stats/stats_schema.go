package stats

import (
	_ "embed"
	"encoding/json"
	"fmt"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

//go:embed stats_schema.json
var statsSchema []byte

func StatsSchema() *jsonschema.Schema {
	s, err := compileSchemaFromBytes(statsSchema)
	if err != nil {
		panic(err)
	}
	return s
}

func compileSchemaFromBytes(schemaBytes []byte) (*jsonschema.Schema, error) {
	comp := jsonschema.NewCompiler()

	// Optionally: enable content / format assertions if desired
	// comp.AssertFormat()
	// comp.AssertContent()

	// Unmarshal schema JSON into a Go value (map / slice / primitive)
	var doc any
	if err := json.Unmarshal(schemaBytes, &doc); err != nil {
		return nil, fmt.Errorf("unmarshal schema JSON: %w", err)
	}

	// Pick some pseudo-URL (or identifier) under which to register the schema.
	const schemaID = "memory://schema.json"

	// Register the document
	if err := comp.AddResource(schemaID, doc); err != nil {
		return nil, fmt.Errorf("AddResource failed: %w", err)
	}

	// Compile the schema
	sch, err := comp.Compile(schemaID)
	if err != nil {
		return nil, fmt.Errorf("Compile failed: %w", err)
	}

	return sch, nil
}

/*
// validateAgainstSchema validates dataBytes (JSON) against a compiled schema
func validateAgainstSchema(schema *jsonschema.Schema, dataBytes []byte) error {
	var inst interface{}
	if err := json.Unmarshal(dataBytes, &inst); err != nil {
		return fmt.Errorf("invalid JSON data: %w", err)
	}
	if err := schema.Validate(inst); err != nil {
		// Validation errors are of type *jsonschema.ValidationError
		return fmt.Errorf("schema validation failed: %w", err)
	}
	return nil
}
*/
