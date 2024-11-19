package heap

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNCoolestPresents(t *testing.T) {
	testCases := []struct {
		name             string
		n                int
		body             []Present
		expectedPresents []Present
	}{
		{
			name: "test1",
			n:    2,
			body: []Present{
				Present{5, 1},
				Present{4, 5},
				Present{3, 1},
				Present{5, 2},
			},
			expectedPresents: []Present{
				Present{5, 1},
				Present{5, 2},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := getNCoolestPresents(tc.body, tc.n)
			//fmt.Println(res)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedPresents, res)
		})
	}
}

func TestGrabPresents(t *testing.T) {
	testCases := []struct {
		name             string
		capacity         int
		body             []Present
		expectedPresents []Present
		expectedError    error
	}{
		{
			name:     "test1",
			capacity: 4,
			body: []Present{
				Present{2000, 3},
				Present{3000, 4},
				Present{1500, 1},
			},
			expectedPresents: []Present{
				Present{1500, 1},
				Present{2000, 3},
			},
			expectedError: nil,
		},
		{
			name:     "test2",
			capacity: 50,
			body: []Present{
				{60, 10},
				{100, 20},
				{120, 30},
			},
			expectedPresents: []Present{
				{120, 30},
				{100, 20},
			},
			expectedError: nil,
		},
		{
			name:             "test3",
			capacity:         0,
			body:             []Present{},
			expectedPresents: []Present{},
			expectedError:    nil,
		},
		{
			name:     "test4",
			capacity: 1,
			body: []Present{
				{1000, 3},
				{3000, 4},
				{1500, 2},
			},
			expectedPresents: []Present{},
			expectedError:    nil,
		},
		{
			name:     "test5",
			capacity: -1,
			body: []Present{
				{1000, 3},
				{3000, 4},
				{1500, 2},
			},
			expectedPresents: nil,
			expectedError:    errors.New("capacity is less than 0"),
		},
		{
			name:     "test6",
			capacity: 4,
			body: []Present{
				Present{-2000, 3},
				Present{3000, 4},
				Present{1500, 1},
			},
			expectedPresents: []Present{
				Present{3000, 4},
			},
			expectedError: nil,
		},
		{
			name:     "test6",
			capacity: 4,
			body: []Present{
				Present{-2000, -3},
				Present{3000, 4},
				Present{1500, -1},
			},
			expectedPresents: nil,
			expectedError:    errors.New("invalid input"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := grabPresents(tc.body, tc.capacity)
			assert.Equal(t, tc.expectedPresents, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
