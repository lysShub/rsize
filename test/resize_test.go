package test_test

import (
	"testing"

	"github.com/lysShub/rsize"
	"github.com/lysShub/rsize/test"
	"github.com/stretchr/testify/require"
)

func TestResize(t *testing.T) {
	for _, test := range test.Tests {
		size := rsize.Size(test.Value)
		require.Equal(t, test.Size, size, test.Name)
	}
}
