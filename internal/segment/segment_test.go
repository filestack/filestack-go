package segment

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilePart_Split(t *testing.T) {

	type filePartTest struct {
		part      Segment
		splitSize int64
		result    []Segment
	}

	tests := []filePartTest{
		{
			part: Segment{
				Offset: 0,
				Length: 8,
			},
			splitSize: 8,
			result: []Segment{
				{Offset: 0, Length: 8},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 8,
			},
			splitSize: 16,
			result: []Segment{
				{Offset: 0, Length: 8},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 10,
			},
			splitSize: 8,
			result: []Segment{
				{Offset: 0, Length: 8},
				{Offset: 8, Length: 2},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 16,
			},
			splitSize: 8,
			result: []Segment{
				{Offset: 0, Length: 8},
				{Offset: 8, Length: 8},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 20,
			},
			splitSize: 8,
			result: []Segment{
				{Offset: 0, Length: 8},
				{Offset: 8, Length: 8},
				{Offset: 16, Length: 4},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 50237794,
			},
			splitSize: 8388608,
			result: []Segment{
				{Offset: 0, Length: 8388608},
				{Offset: 8388608, Length: 8388608},
				{Offset: 16777216, Length: 8388608},
				{Offset: 25165824, Length: 8388608},
				{Offset: 33554432, Length: 8388608},
				{Offset: 41943040, Length: 8294754},
			},
		},
		{
			part: Segment{
				Offset: 16,
				Length: 16,
			},
			splitSize: 6,
			result: []Segment{
				{Offset: 16, Length: 6},
				{Offset: 22, Length: 6},
				{Offset: 28, Length: 4},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 15261194,
			},
			splitSize: 8388608,
			result: []Segment{
				{Offset: 0, Length: 8388608},
				{Offset: 8388608, Length: 6872586},
			},
		},
		{
			part: Segment{
				Offset: 0,
				Length: 8388608,
			},
			splitSize: 5242880,
			result: []Segment{
				{Offset: 0, Length: 5242880},
				{Offset: 5242880, Length: 3145728},
			},
		},
		{
			part: Segment{
				Offset: 8388608,
				Length: 6872586,
			},
			splitSize: 5242880,
			result: []Segment{
				{Offset: 8388608, Length: 5242880},
				{Offset: 13631488, Length: 1629706},
			},
		},
	}

	for testNum, data := range tests {
		t.Run(fmt.Sprintf("Test %v", testNum), func(t *testing.T) {
			result := data.part.Split(data.splitSize)
			assert.Equal(t, data.result, result)
		})
	}
}

func TestFilePart_Bytes(t *testing.T) {

	file := []byte("AAAAAAAAAABBBBBBBBBBCCCCCCCCCCDD")
	reader := bytes.NewReader(file)
	part := newSegment(reader, 0, int64(len(file)))
	parts := part.Split(6)

	expectedSlices := []string{
		"AAAAAA",
		"AAAABB",
		"BBBBBB",
		"BBCCCC",
		"CCCCCC",
		"DD",
	}

	for i, expected := range expectedSlices {
		bytes, err := parts[i].Bytes()
		if err != nil {
			t.Fail()
		}
		assert.Equal(t, []byte(expected), bytes)
	}
}
