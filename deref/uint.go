package deref

// DO NOT MODIFY. Generated by nullable-generate.

import ()

// Uint returns the value pointed to by np.
// If np is nil, then this function returns 0.
func Uint(np *uint) uint {
	if np == nil {
		return 0
	}
	return *np
}
