// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Page 278.

package memo_test

import (
	"testing"

	memo "gopl.io/ch9/memo5"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}
