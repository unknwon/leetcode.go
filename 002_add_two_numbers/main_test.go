package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_AddTwoNumbers(t *testing.T) {
	Convey("Add two numbers", t, func() {
		l1 := &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 4,
				Next: &ListNode{
					Value: 3,
				},
			},
		}
		l2 := &ListNode{
			Value: 5,
			Next: &ListNode{
				Value: 6,
				Next: &ListNode{
					Value: 4,
				},
			},
		}
		node := AddTwoNumbers(l1, l2)
		So(node.Value, ShouldEqual, 7)
		So(node.Next, ShouldNotBeNil)
		node = node.Next
		So(node.Value, ShouldEqual, 0)
		So(node.Next, ShouldNotBeNil)
		node = node.Next
		So(node.Value, ShouldEqual, 8)
		So(node.Next, ShouldBeNil)
	})
}
