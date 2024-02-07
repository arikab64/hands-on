package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton_Trivial(t *testing.T) {
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	if singleton1 != singleton2 {
		t.Error("Expected same instance")
	}

	singleton1.AddOne()
	assert.Equal(t, int64(1), singleton2.GetCount())

	singleton2.AddOne()
	assert.Equal(t, int64(2), singleton1.GetCount())
}
