package intervals

import (
	"testing"
)

func TestCleanInput(t *testing.T) {

	input := " 1, 2 ,3 -5 "
	cleaned := cleanInput(input)
	expected := "1,2,3-5"

	if cleaned != expected {
		t.Errorf("Expected '%s', got '%s'", expected, cleaned)
	}
}

func TestValidate(t *testing.T) {

	_, err := New("1,2,3-5")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = New("1,2,a-5")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

}

func TestValuesLength(t *testing.T) {

	i, _ := New("1,2,10-50")

	if len(i.values) != 43 {
		t.Errorf("Expected length 43, got %d", len(i.values))
	}

}

func TestNext(t *testing.T) {

	i, _ := New("1,2,10-12")

	expectedValues := []int64{1, 2, 10, 11, 12}

	for _, expected := range expectedValues {
		value, err := i.Next()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if value != expected {
			t.Errorf("Expected value %d, got %d", expected, value)
		}
	}

	_, err := i.Next()
	if err == nil {
		t.Errorf("Expected error for exhausted interval, got nil")
	}
}

func TestReset(t *testing.T) {

	i, _ := New("1,2,10-12")

	// Exhaust the interval
	for range i.values {
		i.Next()
	}

	i.Reset()

	value, err := i.Next()
	if err != nil {
		t.Errorf("Expected no error after reset, got %v", err)
	}
	if value != 1 {
		t.Errorf("Expected value 1 after reset, got %d", value)
	}
}

func TestMin(t *testing.T) {

	i, _ := New("5,10-15,20")

	min := i.Min()
	if min != 5 {
		t.Errorf("Expected min 5, got %d", min)
	}
}

func TestMax(t *testing.T) {

	i, _ := New("5,10-15,20")
	max := i.Max()
	if max != 20 {
		t.Errorf("Expected max 20, got %d", max)
	}
}

func TestValues(t *testing.T) {

	i, _ := New("1,3-5,7")

	expectedValues := []int64{1, 3, 4, 5, 7}

	v := i.Values()
	if len(v) != len(expectedValues) {
		t.Errorf("Expected length %d, got %d", len(expectedValues), len(v))
	}

	for idx, expected := range expectedValues {
		if v[idx] != expected {
			t.Errorf("At index %d, expected %d, got %d", idx, expected, v[idx])
		}
	}
}
