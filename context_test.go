package vidar

import (
	"net/http/httptest"
	"testing"
)

func TestNewContext(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	
	ctx := NewContext(w, req)
	if ctx == nil {
		t.Fatal("NewContext() returned nil")
	}
	
	if ctx.request != req {
		t.Error("Request not set correctly")
	}
	
	if ctx.response == nil {
		t.Error("Response is nil")
	}
}

func TestContextSetGet(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	ctx := NewContext(w, req)
	
	// Test Set and Get
	key := "test_key"
	value := "test_value"
	
	err := ctx.Set(key, value)
	if err != nil {
		t.Errorf("Set() returned error: %v", err)
	}
	
	retrieved, exists := ctx.Get(key)
	if !exists {
		t.Error("Get() returned false for existing key")
	}
	
	if retrieved != value {
		t.Errorf("Expected '%s', got '%s'", value, retrieved)
	}
	
	// Test Get for non-existing key
	_, exists = ctx.Get("nonexistent")
	if exists {
		t.Error("Get() returned true for non-existing key")
	}
}