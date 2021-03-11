package numbers_test

import (
	"go-service-echo/util/numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomInt64(t *testing.T) {
	ts := []struct {
		expected int64
	}{
		{5110693013082846335},
		{1271216402443795064},
		{1446451073544560071},
		{2381045507947682220},
		{2510080112869722342},
		{132178464212336404},
	}

	for _, tc := range ts {
		assert.NotEqual(t, tc.expected, numbers.RandomInt64())
	}
}

func Test_RandomInt32(t *testing.T) {
	ts := []struct {
		expected int32
	}{
		{311446475},
		{1837464741},
		{1605255085},
		{361368483},
		{317327984},
		{641821864},
	}

	for _, tc := range ts {
		assert.NotEqual(t, tc.expected, numbers.RandomInt32())
	}
}

func Test_Random(t *testing.T) {
	cases := []struct {
		expected int
	}{
		{2677932324900483984},
		{7677636105966346175},
		{7196187788169409176},
		{1821690828470340119},
		{3990467886744853333},
		{6506211577922304348},
	}

	for _, tc := range cases {
		assert.NotEqual(t, tc.expected, numbers.Random())
	}
}

func Test_RandomBetween(t *testing.T) {
	ts := []struct {
		inMin    int
		inMax    int
		expected int
	}{
		{1, 3, 0},
		{2, 4, 1},
		{3, 5, 2},
		{4, 6, 3},
		{5, 7, 4},
		{6, 8, 5},
	}

	for _, tc := range ts {
		assert.NotEqual(t, tc.expected, numbers.RandomBetween(tc.inMin, tc.inMax))
	}
}

func Test_RandomBetweenInt32(t *testing.T) {
	ts := []struct {
		inMin    int32
		inMax    int32
		expected int32
	}{
		{1, 3, 0},
		{2, 4, 1},
		{3, 5, 2},
		{4, 6, 3},
		{5, 7, 4},
		{6, 8, 5},
	}

	for _, tc := range ts {
		assert.NotEqual(t, tc.expected, numbers.RandomBetweenInt32(tc.inMin, tc.inMax))
	}
}

func Test_RandomBetweenInt64(t *testing.T) {
	ts := []struct {
		inMin    int64
		inMax    int64
		expected int64
	}{
		{1, 3, 0},
		{2, 4, 1},
		{3, 5, 2},
		{4, 6, 3},
		{5, 7, 4},
		{6, 8, 5},
	}

	for _, tc := range ts {
		assert.NotEqual(t, tc.expected, numbers.RandomBetweenInt64(tc.inMin, tc.inMax))
	}
}
