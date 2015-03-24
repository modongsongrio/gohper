package types

import (
	"testing"

	"github.com/cosiner/gohper/lib/test"
)

func TestTrieTree(t *testing.T) {
	tt := test.WrapTest(t)
	tree := new(TrieTree)
	tree.AddPath("abcde", 123)
	tree.AddPath("bcdef", 234)
	tree.AddPath("efghi", 456)
	tree.AddPath("fghij", 789)
	tt.AssertNNil(tree.Match("efghi"))
	tt.AssertEq(tree.Match("abcde").(int), 123)
	tt.AssertEq(tree.Match("bcdef").(int), 234)
	tt.AssertEq(tree.Match("efghi").(int), 456)
	tt.AssertEq(tree.Match("fghij").(int), 789)
	tt.AssertEq(tree.Match("fghia"), nil)
	tt.AssertEq(tree.Match("fasdahia"), nil)
	tt.AssertEq(tree.Match("csaghia"), nil)
}

var s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
var tree = func() *TrieTree {
	t := new(TrieTree)
	for i := 0; i < len(s); i++ {
		t.AddPath(s[i:]+s[:i], true)
	}
	return t
}()

func BenchmarkTrieTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if tree.Match(s).(bool) != true {
			b.Fail()
		}
	}
}