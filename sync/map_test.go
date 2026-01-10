package sync

import (
	"testing"
)

func TestNewMap(t *testing.T) {
	s := NewSync()
	m := s.NewMap()
	_ = m
}

func TestMap_Store_Load(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Test Store and Load - verifies Load() bug fix
	m.Store("key1", "value1")

	val, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if val != "value1" {
		t.Errorf("expected 'value1', got %v", val)
	}
}

func TestMap_Load_Missing(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Test loading non-existent key
	val, ok := m.Load("nonexistent")
	if ok {
		t.Errorf("expected key to not exist, but got value: %v", val)
	}
}

func TestMap_LoadOrStore_New(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Test LoadOrStore with new key - verifies LoadOrStore() bug fix
	val, loaded := m.LoadOrStore("key1", "value1")
	if loaded {
		t.Error("expected loaded to be false for new key")
	}
	if val != "value1" {
		t.Errorf("expected 'value1', got %v", val)
	}

	// Verify it was stored
	storedVal, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist after LoadOrStore")
	}
	if storedVal != "value1" {
		t.Errorf("expected 'value1', got %v", storedVal)
	}
}

func TestMap_LoadOrStore_Existing(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store initial value
	m.Store("key1", "value1")

	// Test LoadOrStore with existing key
	val, loaded := m.LoadOrStore("key1", "value2")
	if !loaded {
		t.Error("expected loaded to be true for existing key")
	}
	if val != "value1" {
		t.Errorf("expected existing value 'value1', got %v", val)
	}

	// Verify original value wasn't changed
	storedVal, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if storedVal != "value1" {
		t.Errorf("expected 'value1', got %v", storedVal)
	}
}

func TestMap_LoadAndDelete(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store a value
	m.Store("key1", "value1")

	// Test LoadAndDelete - verifies LoadAndDelete() bug fix
	val, loaded := m.LoadAndDelete("key1")
	if !loaded {
		t.Error("expected loaded to be true")
	}
	if val != "value1" {
		t.Errorf("expected 'value1', got %v", val)
	}

	// Verify it was deleted
	_, ok := m.Load("key1")
	if ok {
		t.Error("expected key to be deleted")
	}
}

func TestMap_Delete(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store a value
	m.Store("key1", "value1")

	// Delete it
	m.Delete("key1")

	// Verify it was deleted
	_, ok := m.Load("key1")
	if ok {
		t.Error("expected key to be deleted")
	}
}

func TestMap_Range(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store some values
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	// Test Range - verifies Range() bug fix
	seen := make(map[any]any)
	m.Range(func(key, value any) bool {
		seen[key] = value
		return true // continue iteration
	})

	// Verify all keys were seen
	if len(seen) != 3 {
		t.Errorf("expected to see 3 keys, saw %d", len(seen))
	}

	if seen["key1"] != "value1" {
		t.Errorf("expected value1 for key1, got %v", seen["key1"])
	}
	if seen["key2"] != "value2" {
		t.Errorf("expected value2 for key2, got %v", seen["key2"])
	}
	if seen["key3"] != "value3" {
		t.Errorf("expected value3 for key3, got %v", seen["key3"])
	}
}

func TestMap_Range_StopEarly(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store some values
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	// Test stopping Range early
	count := 0
	m.Range(func(key, value any) bool {
		count++
		return count < 2 // stop after 1 iteration
	})

	if count != 2 {
		t.Errorf("expected Range to iterate 2 times, got %d", count)
	}
}

func TestMap_Swap(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store initial value
	m.Store("key1", "value1")

	// Test Swap - verifies Swap() bug fix
	oldVal, loaded := m.Swap("key1", "value2")
	if !loaded {
		t.Error("expected loaded to be true")
	}
	if oldVal != "value1" {
		t.Errorf("expected old value 'value1', got %v", oldVal)
	}

	// Verify new value was stored
	newVal, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if newVal != "value2" {
		t.Errorf("expected 'value2', got %v", newVal)
	}
}

func TestMap_Swap_NewKey(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Test Swap with new key
	oldVal, loaded := m.Swap("key1", "value1")
	if loaded {
		t.Error("expected loaded to be false for new key")
	}
	if oldVal != nil {
		t.Errorf("expected nil old value, got %v", oldVal)
	}

	// Verify new value was stored
	newVal, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if newVal != "value1" {
		t.Errorf("expected 'value1', got %v", newVal)
	}
}

func TestMap_CompareAndSwap(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store initial value
	m.Store("key1", "value1")

	// Test CompareAndSwap with correct old value
	swapped := m.CompareAndSwap("key1", "value1", "value2")
	if !swapped {
		t.Error("expected CompareAndSwap to succeed")
	}

	// Verify new value
	val, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if val != "value2" {
		t.Errorf("expected 'value2', got %v", val)
	}

	// Test CompareAndSwap with incorrect old value
	swapped = m.CompareAndSwap("key1", "value1", "value3")
	if swapped {
		t.Error("expected CompareAndSwap to fail with wrong old value")
	}

	// Verify value unchanged
	val, ok = m.Load("key1")
	if !ok {
		t.Fatal("expected key to exist")
	}
	if val != "value2" {
		t.Errorf("expected 'value2' (unchanged), got %v", val)
	}
}

func TestMap_CompareAndDelete(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store initial value
	m.Store("key1", "value1")

	// Test CompareAndDelete with correct old value
	deleted := m.CompareAndDelete("key1", "value1")
	if !deleted {
		t.Error("expected CompareAndDelete to succeed")
	}

	// Verify key was deleted
	_, ok := m.Load("key1")
	if ok {
		t.Error("expected key to be deleted")
	}

	// Store again
	m.Store("key1", "value2")

	// Test CompareAndDelete with incorrect old value
	deleted = m.CompareAndDelete("key1", "value1")
	if deleted {
		t.Error("expected CompareAndDelete to fail with wrong old value")
	}

	// Verify key still exists
	val, ok := m.Load("key1")
	if !ok {
		t.Fatal("expected key to still exist")
	}
	if val != "value2" {
		t.Errorf("expected 'value2', got %v", val)
	}
}

func TestMap_Clear(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	// Store some values
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	m.Store("key3", "value3")

	// Clear the map
	m.Clear()

	// Verify all keys are gone
	count := 0
	m.Range(func(key, value any) bool {
		count++
		return true
	})

	if count != 0 {
		t.Errorf("expected map to be empty, but found %d entries", count)
	}
}

func TestMap_ConcurrentReadWrite(t *testing.T) {
	s := NewSync()
	m := s.NewMap()

	done := make(chan bool)

	// Writer goroutine
	go func() {
		for i := 0; i < 100; i++ {
			m.Store(i, i*2)
		}
		done <- true
	}()

	// Reader goroutine
	go func() {
		for i := 0; i < 100; i++ {
			m.Load(i)
		}
		done <- true
	}()

	// Wait for both
	<-done
	<-done

	// Verify some values
	val, ok := m.Load(50)
	if !ok {
		t.Error("expected key 50 to exist")
	}
	if val != 100 {
		t.Errorf("expected value 100 for key 50, got %v", val)
	}
}

func TestMap_Operations_TableDriven(t *testing.T) {
	tests := []struct {
		name  string
		setup func(Map)
		op    func(Map)
		check func(Map, *testing.T)
	}{
		{
			name:  "store and load",
			setup: func(m Map) {},
			op:    func(m Map) { m.Store("k1", "v1") },
			check: func(m Map, t *testing.T) {
				val, ok := m.Load("k1")
				if !ok || val != "v1" {
					t.Errorf("expected v1, got %v (exists: %v)", val, ok)
				}
			},
		},
		{
			name: "delete existing",
			setup: func(m Map) {
				m.Store("k1", "v1")
			},
			op: func(m Map) { m.Delete("k1") },
			check: func(m Map, t *testing.T) {
				_, ok := m.Load("k1")
				if ok {
					t.Error("expected key to be deleted")
				}
			},
		},
		{
			name:  "delete non-existent",
			setup: func(m Map) {},
			op:    func(m Map) { m.Delete("k1") },
			check: func(m Map, t *testing.T) {
				_, ok := m.Load("k1")
				if ok {
					t.Error("expected key to not exist")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSync()
			m := s.NewMap()
			tt.setup(m)
			tt.op(m)
			tt.check(m, t)
		})
	}
}
