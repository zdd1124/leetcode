package Num_451_FrequencySort

import (
	"fmt"
	"testing"
)

func TestNum_451_FrequencySort(t *testing.T) {

	e := []struct {
		expected interface{}
	}{
		{},
	}
	for _, v := range e {
		o := 0
		if o != v.expected {
			t.Error(fmt.Sprintf("%v %v", v, o))
		}
	}
}
