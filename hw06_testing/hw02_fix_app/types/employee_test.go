package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmployee_String(t *testing.T) {
	e := Employee{
		UserID:       1,
		Age:          30,
		Name:         "John Doe",
		DepartmentID: 10,
	}

	expected := "User ID: 1; Age: 30; Name: John Doe; Department ID: 10; "

	require.Equal(t, expected, e.String())
}
