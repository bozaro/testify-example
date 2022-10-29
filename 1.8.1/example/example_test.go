package example

import (
	"github.com/bozaro/testify-example/example/helper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimple(t *testing.T) {
	require.Fail(t, "Hello, world!")
}

func TestHelper(t *testing.T) {
	helper.CheckTrace(t)
}
