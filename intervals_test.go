package intervals

import (
	"testing"
)

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
