package packageManager

import (
	"reflect"
	"testing"
)

func TestDependencyDescriptor(t *testing.T) {
	// Create a sample DependencyDescriptor
	desc := DependencyDescriptor{
		Name:           "test-package",
		Installer:      InstallAs(0), // Assuming InstallAs is an enum or custom type
		Detail:         "test-detail",
		SkipIfDetected: true,
	}

	// Check the fields of DependencyDescriptor
	typ := reflect.TypeOf(desc)
	numFields := typ.NumField()

	// Ensure the expected number of fields
	expectedFields := 4
	if numFields != expectedFields {
		t.Errorf("Expected %d fields, but got %d", expectedFields, numFields)
	}

	// Validate each field
	expectedFieldTypes := map[string]reflect.Kind{
		"Name":           reflect.String,
		"Installer":      reflect.TypeOf(InstallAs(0)).Kind(), // Assuming InstallAs is an enum or custom type
		"Detail":         reflect.String,
		"SkipIfDetected": reflect.Bool,
	}

	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		fieldName := field.Name
		fieldType := field.Type.Kind()

		expectedType, found := expectedFieldTypes[fieldName]
		if !found {
			t.Errorf("Unexpected field name: %s", fieldName)
			continue
		}

		if fieldType != expectedType {
			t.Errorf("Field %s has incorrect type. Expected %v, but got %v", fieldName, expectedType, fieldType)
		}
	}
}
