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

	if !validateFormat(s) {
		return nil, fmt.Errorf("invalid interval string: %s", s)
	}

	values, err := parse(s)
	if err != nil {
		return nil, err
	}

	slices.Sort(values)

	return &Interval{values: values}, nil
}

func validateFormat(s string) bool {

	r, _ := regexp.Compile("^[0-9,-]+$")

	return r.MatchString(s)
}

func parse(s string) ([]int64, error) {

	values := []int64{}
	parts := strings.Split(s, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")

			lowerBound, err := strconv.ParseInt(bounds[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid integer in range: %s", bounds[0])
			}
			upperBound, err := strconv.ParseInt(bounds[1], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid integer in range: %s", bounds[1])
			}

			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			if lowerBound > upperBound {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			for i := lowerBound; i <= upperBound; i++ {
				values = append(values, i)
			}
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
