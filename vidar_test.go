package vidar

import (
	"testing"
)

func TestNew(t *testing.T) {
	v := New()
	if v == nil {
		t.Fatal("New() returned nil")
	}

	if v.Router == nil {
		t.Error("Router is nil")
	}

	if v.Server == nil {
		t.Error("Server is nil")
	}

	if v.Plugin == nil {
		t.Error("Plugin is nil")
	}

	if v.log == nil {
		t.Error("log is nil")
	}
}

func TestResolveAddress(t *testing.T) {
	v := New()

	// Test default address
	addr, err := v.resolveAddress()
	if err != nil {
		t.Errorf("resolveAddress() returned error: %v", err)
	}
	if addr != DefaultAddress {
		t.Errorf("Expected default address '%s', got '%s'", DefaultAddress, addr)
	}

	// Test with custom address
	addr, err = v.resolveAddress("localhost", "9000")
	if err != nil {
		t.Errorf("resolveAddress() returned error: %v", err)
	}
	if addr != "localhost:9000" {
		t.Errorf("Expected custom address 'localhost:9000', got '%s'", addr)
	}
}

func TestResolveAddressInvalidArgs(t *testing.T) {
	v := New()

	// Test with invalid number of arguments
	addr, err := v.resolveAddress("only-one-arg")
	if err != nil {
		t.Errorf("resolveAddress() returned error: %v", err)
	}
	// Should fall back to default
	if addr != DefaultAddress {
		t.Errorf("Expected fallback to default address '%s', got '%s'", DefaultAddress, addr)
	}
}
