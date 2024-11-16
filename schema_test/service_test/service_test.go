package schema_test

import (
	"os"
	"testing"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

func TestService_Schema(t *testing.T) {
	schemaFile := "../../schema/service/service.schema.json"
	instanceFile := "./service_instance.json"

	c := jsonschema.NewCompiler()
	sch, err := c.Compile(schemaFile)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(instanceFile)
	if err != nil {
		t.Fatal(err)
	}
	inst, err := jsonschema.UnmarshalJSON(f)
	if err != nil {
		t.Fatal(err)
	}

	err = sch.Validate(inst)
	if err != nil {
		t.Error("Schema validation failed: ", err)
	}

}
