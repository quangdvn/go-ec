package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input = 1
	// 	ouput = 2
	// )
	// actual := AddOne(input)
	// if actual != ouput {
	// 	t.Errorf("AddOne(%d), input %d, actual %d", input, ouput, actual)
	// }
	assert.Equal(t, AddOne(1), 2, "AddOne(1) should be 2")
	assert.NotEqual(t, 2, 3)
	assert.Nil(t, nil, nil)
}

func TestAddTwo(t *testing.T) {
	assert.Equal(t, AddTwo(1), 3, "AddTwo(1) should be 3")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3) // Stop the rest of the code if this fails
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Executing")
}
