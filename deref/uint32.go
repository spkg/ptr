package deref

// DO NOT MODIFY. Generated by nullable-generate.

import ()

// Uint32 returns the value pointed to by np.
// If np is nil, then this function returns 0.
func Uint32(np *uint32) uint32 {
	if np == nil {
		return 0
	}
	return *np
}
