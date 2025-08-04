package vidar

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewPlugin(t *testing.T) {
	plugin := NewPlugin()
	if plugin == nil {
		t.Fatal("NewPlugin() returned nil")
	}

	if plugin.rings == nil {
		t.Error("Plugin rings slice is nil")
	}

	if len(plugin.rings) != 0 {
		t.Errorf("Expected empty rings slice, got length %d", len(plugin.rings))
	}
}

func TestPluginAppend(t *testing.T) {
	plugin := NewPlugin()

	// Test middleware function
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Test", "middleware")
			next.ServeHTTP(w, r)
		})
	}

	plugin.Append(middleware)

	if len(plugin.rings) != 1 {
		t.Errorf("Expected 1 middleware, got %d", len(plugin.rings))
	}
}

func TestPluginExtend(t *testing.T) {
	plugin1 := NewPlugin()
	plugin2 := NewPlugin()

	middleware1 := func(next http.Handler) http.Handler {
		return next
	}
	middleware2 := func(next http.Handler) http.Handler {
		return next
	}

	plugin1.Append(middleware1)
	plugin2.Append(middleware2)

	extended := plugin1.Extend(*plugin2)

	if len(extended.rings) != 2 {
		t.Errorf("Expected 2 middleware in extended plugin, got %d", len(extended.rings))
	}
}

func TestContextFunc(t *testing.T) {
	handler := ContextFunc(func(c *Context) {
		c.response.WriteHeader(http.StatusOK)
		c.response.Write([]byte("test"))
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if w.Body.String() != "test" {
		t.Errorf("Expected body 'test', got '%s'", w.Body.String())
	}
}
