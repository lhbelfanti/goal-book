// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"math/rand"
	"sort"
	"testing"

	"gitlab.com/lhbelfanti/goal-book/chapter4/examples/08treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
