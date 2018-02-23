// This package implements the Levenshtein algorithm for computing the
// similarity between two strings. The Calc function returns the Levenshtein
// distance and Similarity degree.
package levenshtein

import (
	"math"
	"strings"
)

// The Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// http://en.wikipedia.org/wiki/Levenshtein_distance
func Calc(a, b string) (int, float64) {
	if strings.Compare(a, b) == 0 {
		return 0, 1.0
	}
	if len([]rune(a)) > len([]rune(b)) {
		a, b = b, a
	}
	ra, rb := []rune(a), []rune(b)
	la, lb := len(ra), len(rb)
	max := math.Max(float64(la), float64(lb))
	for la > 0 && ra[la-1] == rb[lb-1] {
		la--
		lb--
	}
	offset := 0
	for offset < la && ra[offset] == rb[offset] {
		offset++
	}
	la -= offset
	lb -= offset
	if la == 0 || lb == 1 {
		return lb, 1.0 - float64(lb)/max
	}
	x := 0
	var y, d0, d1, d2, d3, dd, dy, ay, bx0, bx1, bx2, bx3 int
	var vector []int
	for y = 0; y < la; y++ {
		vector = append(vector, y+1, int(ra[offset+y]))
	}
	for (x + 3) < lb {
		d0, d1, d2, d3 = x, x+1, x+2, x+3
		bx0, bx1, bx2, bx3 = int(rb[offset+d0]), int(rb[offset+d1]), int(rb[offset+d2]), int(rb[offset+d3])
		x += 4
		dd = x
		for y = 0; y < len(vector); y += 2 {
			dy, ay = vector[y], vector[y+1]
			d0 = min(dy, d0, d1, bx0, ay)
			d1 = min(d0, d1, d2, bx1, ay)
			d2 = min(d1, d2, d3, bx2, ay)
			dd = min(d2, d3, dd, bx3, ay)
			vector[y] = dd
			d3, d2, d1, d0 = d2, d1, d0, dy
		}
	}
	for x < lb {
		d0 = x
		bx0 = int(rb[offset+d0])
		x++
		dd = x
		for y = 0; y < len(vector); y += 2 {
			dy = vector[y]
			dd = min(dy, d0, dd, bx0, vector[y+1])
			vector[y] = dd
			d0 = dy
		}
	}
	return dd, 1.0 - float64(dd)/max
}
func min(d0, d1, d2, bx, ay int) int {
	if d0 < d1 || d2 < d1 {
		if d0 > d2 {
			return d2 + 1
		}
		return d0 + 1
	} else {
		if bx == ay {
			return d1
		}
		return d1 + 1
	}
}
