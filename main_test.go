package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.True(t, true, "This is good")

}
