package reader

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/varik-08/golang-hw/hw06_testing/hw02_fix_app/types"
)

func TestReadJSON(t *testing.T) {
	data, err := ReadJSON("data.json")

	require.NoError(t, err)

	require.Len(t, data, 2)

	require.IsType(t, []types.Employee{}, data)
}
