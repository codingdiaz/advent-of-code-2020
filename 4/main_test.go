package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_checkYears(t *testing.T) {
	assert.True(t, checkYears("1994", 1990, 2000))
	assert.False(t, checkYears("1994", 1998, 2000))
}


func Test_checkHeights(t *testing.T) {
	assert.True(t, checkHeight("150cm"))
	assert.True(t, checkHeight("60in"))
	assert.False(t, checkHeight("100cm"))
}

func Test_checkHair(t *testing.T) {
	assert.True(t, checkHair("#123456"))
	assert.False(t, checkHair("#12345"))
	assert.False(t, checkHair("#12347g"))
	assert.True(t, checkHair("#12347f"))
}