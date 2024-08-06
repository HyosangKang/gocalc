package gocalc_test

import (
	"gocalc"
	"testing"
)

func TestIntegral(t *testing.T) {
	mat := [][]gocalc.Real{
		{SimpleReal(1), SimpleReal(2)},
		{SimpleReal(3), SimpleReal(4)},
	}
	t.Log(gocalc.Det(mat))
	mat = [][]gocalc.Real{
		{SimpleReal(1), SimpleReal(2), SimpleReal(3)},
		{SimpleReal(4), SimpleReal(5), SimpleReal(6)},
		{SimpleReal(7), SimpleReal(8), SimpleReal(9)},
	}
	t.Log(gocalc.Det(mat))
}
