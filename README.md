# go-intervals

A tiny helper to parse and use numeric intervals and ranges, like port ranges.  
Parses strings with intervals and ranges, like "1-5, 8, 100-200" and creates an array of integers containing all values.

## Example usage

```
package main

import (
	"fmt"

	intervals "github.com/kjansson/go-intervals"
)

interval, err := intervals.New("5,6,8,10-15")
if err != nil {
    panic()
}



```