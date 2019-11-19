package addtwonumbers

import (
	"math"
	"strconv"
	"testing"
)

func TestNode2Int(t *testing.T) {
	tests := []struct {
		l    *ListNode
		want int
	}{
		{
			l: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			want: 342,
		},
		{
			l: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			want: 465,
		},
	}

	for idx := range tests {
		tc := tests[idx]
		res := node2Int(tc.l)
		if *res != tc.want {
			t.Errorf("want %d, but returned %d", tc.want, *res)
		}
	}
}

func TestInt2Node(t *testing.T) {
	l := int2Node(807)

	res := ""
	for l != nil {
		res += strconv.Itoa(l.Val)
		l = l.Next
	}

	if res != "708" {
		t.Errorf("want '807', but returned %s", res)
	}
}

func TestFoobar(t *testing.T) {
	t.Error(math.MaxInt64)
}

func TestAddTwoNumbersV2(t *testing.T) {
	tests := []struct {
		l    *ListNode
		want int
	}{
		{
			l: &ListNode{
				Val: 1,
			},
		},
		{
			l: &ListNode{
				Val: 0,
			},
		},
	}
	res := addTwoNumbersV2(tests[0].l, tests[1].l)
	t.Error(res)
}

func TestStringIndex(t *testing.T) {
	s := "bab"
	t.Error(isPalindrome(s))
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
