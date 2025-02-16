package helper_test

import (
	"cnyes-stock-news/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	assert := assert.New(t)
	assert.True(helper.Filter("hello world", []string{}, []string{}, ""))
	assert.True(helper.Filter("hello world", []string{"hello"}, []string{}, ""))
	assert.True(helper.Filter("hello world", []string{"hello"}, []string{"foo"}, ""))
	assert.True(helper.Filter("hello world", []string{"hello"}, []string{}, "^hello"))
	assert.True(helper.Filter("hello world", []string{"hello"}, []string{"foo"}, "^hello"))
	assert.True(helper.Filter("hello world", []string{}, []string{"foo"}, ""))
	assert.True(helper.Filter("hello world", []string{}, []string{"foo"}, "^hello"))
	assert.True(helper.Filter("hello world", []string{}, []string{}, "^hello"))

	assert.False(helper.Filter("hello world", []string{"foo"}, []string{}, ""))
	assert.False(helper.Filter("hello world", []string{}, []string{"hello"}, ""))
	assert.False(helper.Filter("hello world", []string{}, []string{}, "^foo"))

	// Invalid regular expression pattern return false
	assert.True(helper.Filter("hello world", []string{"("}, []string{}, ""))
	assert.True(helper.Filter("hello world", []string{}, []string{"("}, ""))
	assert.True(helper.Filter("hello world", []string{}, []string{}, "?"))
}
