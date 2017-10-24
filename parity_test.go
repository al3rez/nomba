package nomba

import (
	"testing"

	"aahframework.org/test.v0/assert"
)

func TestIsEvenInt(t *testing.T) {
	assert.Equal(t, Int(2).IsEven(), true)
	assert.Equal(t, Int(1).IsEven(), false)
}

func TestIsOddInt(t *testing.T) {
	assert.Equal(t, Int(2).IsOdd(), false)
	assert.Equal(t, Int(1).IsOdd(), true)
}

func TestIsEvenInt8(t *testing.T) {
	assert.Equal(t, Int8(2).IsEven(), true)
	assert.Equal(t, Int8(1).IsEven(), false)
}

func TestIsOddInt8(t *testing.T) {
	assert.Equal(t, Int8(2).IsOdd(), false)
	assert.Equal(t, Int8(1).IsOdd(), true)
}

func TestIsEvenInt16(t *testing.T) {
	assert.Equal(t, Int16(2).IsEven(), true)
	assert.Equal(t, Int16(1).IsEven(), false)
}

func TestIsOddInt16(t *testing.T) {
	assert.Equal(t, Int16(2).IsOdd(), false)
	assert.Equal(t, Int16(1).IsOdd(), true)
}

func TestIsEvenInt32(t *testing.T) {
	assert.Equal(t, Int32(2).IsEven(), true)
	assert.Equal(t, Int32(1).IsEven(), false)
}

func TestIsOddInt32(t *testing.T) {
	assert.Equal(t, Int32(2).IsOdd(), false)
	assert.Equal(t, Int32(1).IsOdd(), true)
}

func TestIsEvenInt64(t *testing.T) {
	assert.Equal(t, Int64(2).IsEven(), true)
	assert.Equal(t, Int64(1).IsEven(), false)
}

func TestIsOddInt64(t *testing.T) {
	assert.Equal(t, Int64(2).IsOdd(), false)
	assert.Equal(t, Int64(1).IsOdd(), true)
}
