package golangnugetinfo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	v, e := GetNugetVersions("Microsoft.Extensions.Logging")
	require.Nil(t, e)
	require.Equal(t, "1.0.0-rc1-final", v[0])
}
