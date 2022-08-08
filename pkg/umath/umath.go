/*
A package of math utilities similar to those found in math libraries
in other languages.
*/
package umath

/* Returns the smaller of (a, b). */
func Min[N int | int32 | int64 | uint | uint32 | uint64](a, b N) N {
	if a < b {
		return a
	}
	return b
}

/* Returns the larger of (a, b). */
func Max[N int | int32 | int64 | uint | uint32 | uint64](a, b N) N {
	if a < b {
		return a
	}
	return b
}
