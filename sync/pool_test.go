package sync

import (
	"testing"
)

func TestNewPool(t *testing.T) {
	s := NewSync()
	p := s.NewPool()
	_ = p
}

func TestPool_Get_Put(t *testing.T) {
	s := NewSync()
	p := s.NewPool()

	// Put an object
	p.Put("test-value")

	// Get should retrieve it
	val := p.Get()
	if val == nil {
		t.Error("expected to get value from pool")
	}

	// Verify it's the right value
	if str, ok := val.(string); !ok || str != "test-value" {
		t.Errorf("expected 'test-value', got %v", val)
	}
}

func TestPool_Get_Empty(t *testing.T) {
	s := NewSync()
	p := s.NewPool()

	// Get from empty pool should return nil (no New function set)
	val := p.Get()
	if val != nil {
		t.Errorf("expected nil from empty pool, got %v", val)
	}
}

func TestPool_New_Function(t *testing.T) {
	newCalled := false
	s := NewSync()
	p := s.NewPool(WithNew(func() any {
		newCalled = true
		return "new-value"
	}))

	// Get from empty pool should call New function
	val := p.Get()

	if !newCalled {
		t.Error("expected New function to be called")
	}

	if str, ok := val.(string); !ok || str != "new-value" {
		t.Errorf("expected 'new-value', got %v", val)
	}
}

func TestPool_MultipleGetPut(t *testing.T) {
	s := NewSync()
	p := s.NewPool()

	// Put multiple values
	p.Put("value1")
	p.Put("value2")
	p.Put("value3")

	// Get them back (order not guaranteed)
	got := make(map[string]bool)
	for i := 0; i < 3; i++ {
		val := p.Get()
		if val != nil {
			if str, ok := val.(string); ok {
				got[str] = true
			}
		}
	}

	// We should have gotten some of the values
	if len(got) == 0 {
		t.Error("expected to get at least one value from pool")
	}
}

func TestPool_Reuse(t *testing.T) {
	s := NewSync()
	p := s.NewPool()

	// Put, get, put, get cycle
	original := "test-value"
	p.Put(original)

	val1 := p.Get()
	if val1 == nil {
		t.Fatal("expected to get value from pool")
	}

	// Put it back
	p.Put(val1)

	// Get it again
	val2 := p.Get()
	if val2 == nil {
		t.Fatal("expected to get value from pool again")
	}

	// Should be the same value (reused)
	if val1 != val2 {
		t.Error("expected to get the same value back (pool reuse)")
	}
}

func TestPool_DifferentTypes(t *testing.T) {
	s := NewSync()
	p := s.NewPool()

	// Pool can store different types
	p.Put("string")
	p.Put(42)
	p.Put([]int{1, 2, 3})

	// Get them back
	for i := 0; i < 3; i++ {
		val := p.Get()
		if val == nil {
			t.Error("expected to get value from pool")
		}
	}
}

func TestPool_NilValue(t *testing.T) {
	s := NewSync()
	p := s.NewPool()

	// Putting nil is allowed but not useful
	p.Put(nil)

	// Getting from pool with nil might return nil or call New
	_ = p.Get()
}

func TestPool_Concurrent(t *testing.T) {
	s := NewSync()
	p := s.NewPool(WithNew(func() any {
		return &struct{}{}
	}))

	const numGoroutines = 100
	done := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			// Get from pool
			val := p.Get()
			if val == nil {
				t.Error("expected non-nil value")
			}

			// Use it (simulated by brief pause)

			// Put it back
			p.Put(val)

			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}
