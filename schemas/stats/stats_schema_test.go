package stats

import "testing"

func TestSchema(t *testing.T) {
	_, err := compileSchemaFromBytes(statsSchema)
	if err != nil {
		t.Errorf("stats.compileSchemaFromBytes() error: (%s)", err.Error())
	}
}
