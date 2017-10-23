package nomba

import (
	"testing"

	"aahframework.org/test.v0/assert"
)

func TestAbsInt(t *testing.T) {
	assert.Equal(t, Abs(Int(-100)), Int(100))
	assert.Equal(t, Abs(Int(100)), Int(100))
}

func TestAbsInt8(t *testing.T) {
	assert.Equal(t, Abs(Int8(-100)), Int8(100))
	assert.Equal(t, Abs(Int8(100)), Int8(100))
}

func TestAbsInt16(t *testing.T) {
	assert.Equal(t, Abs(Int16(-100)), Int16(100))
	assert.Equal(t, Abs(Int16(100)), Int16(100))
}

func TestAbsInt32(t *testing.T) {
	assert.Equal(t, Abs(Int32(-100)), Int32(100))
	assert.Equal(t, Abs(Int32(100)), Int32(100))
}

func TestAbsInt64(t *testing.T) {
	assert.Equal(t, Abs(Int64(-100)), Int64(100))
	assert.Equal(t, Abs(Int64(100)), Int64(100))
}

func TestAbsFloat32(t *testing.T) {
	assert.Equal(t, Abs(Float32(-100)), Float32(100))
	assert.Equal(t, Abs(Float32(100)), Float32(100))
}

func TestAbsFloat64(t *testing.T) {
	assert.Equal(t, Abs(Float64(-100)), Float64(100))
	assert.Equal(t, Abs(Float64(100)), Float64(100))
}
