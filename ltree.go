package gox

import "strings"

// LTree is a type of string which features dot-delimitated portions declaring
// a scope in order from widest, to narrowest. They are good for searching scoped
// tags that can be held as human-readable text.
//
// The LTree segments should be case-insensitive leaning towards lower-case
// favored.
type LTree string

func (t LTree) String() string {
	return string(t)
}

func (t LTree) Segments() []string {
	return SplitStringByRune(string(t), '.')
}

func (t LTree) Prefix(segments ...string) LTree {
	return NewLTree(JoinSlices(segments, t.Segments())...)
}

func (t LTree) Postfix(segments ...string) LTree {
	return NewLTree(JoinSlices(t.Segments(), segments)...)
}

// Match checks if the given query string matches against this L-Tree. This is
// a very basic implementation and allows only for the '*' operator to use as
// a wildcard for a whole segment. Otherwise, each part is matched in order.
func (t LTree) Match(str string) bool {
	if len(t) == 0 {
		return str == "*"
	}

	segsQ := SplitStringByRune(strings.ToLower(str), '.')
	segsT := t.Segments()

	return MatchLTreeSegments(segsT, segsQ)
}

// NewLTree joins the given segments with the "." deliminator to form a new
// L-Tree string. This normalizes the string to lowercase
func NewLTree(segments ...string) LTree {
	return LTree(strings.ToLower(strings.Join(segments, ".")))
}

// MatchLTreeSegments checks if the given L-Tree segments match against the
// given query segments. This is a low-level function for custom iterators.
func MatchLTreeSegments(ltree []string, query []string) bool {
	if len(query) > len(ltree) {
		return false
	}

	for i, q := range query {
		if q != "*" && ltree[i] != q {
			return false
		}
	}
	return true
}
