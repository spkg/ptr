package deref

// DO NOT MODIFY. Generated by nullable-generate.

import ()

// Uint16 returns the value pointed to by np.
// If np is nil, then this function returns 0.
func Uint16(np *uint16) uint16 {
	if np == nil {
		return 0
	}
	return *np
}
