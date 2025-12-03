package intervals

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	values []int64
	index  int64
}

func New(s string) (*Interval, error) {

	// Clean input
	s = cleanInput(s)
	// Validate format
	if !validateFormat(s) {
		return nil, fmt.Errorf("invalid interval string: %s", s)
	}

	// Parse values
	values, err := parse(s)
	if err != nil {
		return nil, err
	}

	// Sort values
	slices.Sort(values)

	return &Interval{values: values}, nil
}

func cleanInput(s string) string {
	// Remove all whitespace
	return strings.ReplaceAll(s, " ", "")
}

func validateFormat(s string) bool {

	// Regular expression to match valid characters: digits, commas, hyphens
	r, _ := regexp.Compile("^[0-9,-]+$")

	return r.MatchString(s)
}

func parse(s string) ([]int64, error) {

	values := []int64{}
	// Split by comma
	parts := strings.SplitSeq(s, ",")

	for part := range parts {
		// Check if part is a range
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")

			// Validate that we have exactly two bounds
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			// Parse bounds
			lowerBound, err := strconv.ParseInt(bounds[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid integer in range: %s", bounds[0])
			}
			upperBound, err := strconv.ParseInt(bounds[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid integer in range: %s", bounds[1])
			}

			// Validate that the lower bound is less than upper bound
			if lowerBound > upperBound {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			// Append all values in the range
			for i := lowerBound; i <= upperBound; i++ {
				values = append(values, i)
			}
			// Single value
		} else {
			intVal, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid integer in range: %s", part)
			}

			values = append(values, intVal)
		}
	}

	return values, nil
}

func (i *Interval) Next() (int64, error) {
	if int(i.index) >= len(i.values) {
		return 0, fmt.Errorf("no more values")
	}

	val := i.values[i.index]
	i.index++

	return val, nil
}

func (i *Interval) Reset() {
	i.index = 0
}

func (i *Interval) Min() int64 {
	if len(i.values) == 0 {
		return 0
	}
	return i.values[0]
}

func (i *Interval) Max() int64 {
	if len(i.values) == 0 {
		return 0
	}
	return i.values[len(i.values)-1]
}

func (i *Interval) Values() []int64 {
	return i.values
}
