package pubip

import "testing"

func TestMaster(t *testing.T) {
	m := NewMaster()
	if _, err := m.Address(); err != nil {
		t.Error(err)
	}
}

func TestAddress_parallel(t *testing.T) {
	m := NewMaster()
	m.Parallel = 4
	_, err := m.Address()
	if err != nil {
		t.Error(err)
	}
}

func TestAddress_sequential(t *testing.T) {
	m := NewMaster()
	m.Parallel = 1
	_, err := m.Address()
	if err != nil {
		t.Error(err)
	}
}
