package deref

// DO NOT MODIFY. Generated by nullable-generate.

// Byte returns the value pointed to by np.
// If np is nil, then this function returns 0.
func Byte(np *byte) byte {
	if np == nil {
		return 0
	}
	return *np
}
