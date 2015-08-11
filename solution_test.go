package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_TwoSum(t *testing.T) {
	Convey("Two sum", t, func() {
		idxs := TwoSum([]int{2, 7, 11, 15}, 9)
		So(idxs[0], ShouldEqual, 1)
		So(idxs[1], ShouldEqual, 2)
	})
}
