package options

import (
	"testing"

	tests "github.com/sunglim/yh-finance-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetStraddle(t *testing.T) {

	iter := GetStraddle(tests.TestStraddleSymbol)
	success := iter.Next()
	assert.True(t, success)
	assert.Nil(t, iter.Err())
	assert.Equal(t, iter.Meta().UnderlyingSymbol, tests.TestStraddleSymbol)
}
