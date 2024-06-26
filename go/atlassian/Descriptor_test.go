package Atlassian

import (
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
	"testing"
)

// TestDescriptor verifies the functionality of the Descriptor interface
func TestDescriptor(t *testing.T) {
	var domain Domain
	var d Descriptor
	d = &MockDescriptor{} // initialize with the mock implementation

	// Test Load method
	if err := d.Load("test_file.txt"); err != nil {
		t.Fatalf("Load method returned an error: %v", err)
	}
	if _, err := d.Create(&domain); err != nil {
		t.Fatalf("Create method returned an error: %v", err)
	}
	if _, err := d.Read(&domain, ""); err != nil {
		t.Fatalf("Read method returned an error: %v", err)
	}
	if _, err := d.Update(&domain, ""); err != nil {
		t.Fatalf("Update method returned an error: %v", err)
	}
	if _, err := d.Delete(&domain, ""); err != nil {
		t.Fatalf("Delete method returned an error: %v", err)
	}
	if _, err := d.List(&domain, nil); err != nil {
		t.Fatalf("List method returned an error: %v", err)
	}
}

// MockDescriptor implements the Descriptor interface for testing purposes
type MockDescriptor struct{}

func (m *MockDescriptor) Load(fileName string) error {
	return nil
}

func (m *MockDescriptor) Create(domain *Domain) (*http.Request, error) {
	return nil, nil
}

func (m *MockDescriptor) Read(domain *Domain, n string) (*http.Request, error) {
	return nil, nil
}

func (m *MockDescriptor) Update(domain *Domain, n string) (*http.Request, error) {
	return nil, nil
}

func (m *MockDescriptor) Delete(domain *Domain, n string) (*http.Request, error) {
	// implement logic for Delete method
	return nil, nil
}

func (m *MockDescriptor) List(domain *Domain, n *AtlassianTypes.JqlQuery) (*http.Request, error) {
	return nil, nil
}
