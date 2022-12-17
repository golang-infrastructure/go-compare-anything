package compare_anything

import (
	"github.com/golang-infrastructure/go-compare-anything/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCastToComparable(t *testing.T) {
	o := &test.ComparableTest{}
	toComparable, err := CastToComparable(o)
	assert.Nil(t, err)
	assert.NotNil(t, toComparable)
}

func TestIsComparable(t *testing.T) {
	o := &test.ComparableTest{}
	r := IsComparable(o)
	assert.True(t, r)
}
