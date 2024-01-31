// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 276.

package memo_test

import (
	"testing"

	memo "gopl.io/ch9/memo4"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
