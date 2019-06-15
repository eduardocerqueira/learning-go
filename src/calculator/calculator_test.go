package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := 32

	if got := Add(20, 12); got != expected {
		t.Errorf("expected %d but got %d instead", expected, got)
	}
}

func TestMultiply(t *testing.T) {
	expected := 10

	if got := Multiply(5, 2); got != expected {
		t.Errorf("expected %d but got %d instead", expected, got)
	}
}

func TestDivide(t *testing.T) {
	expected := 10
	res := Divide(100, 10)

	assert.Equal(t, expected, res, "result no right")
}
