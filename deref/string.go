package deref

// DO NOT MODIFY. Generated by nullable-generate.

import ()

// String returns the value pointed to by sp.
// If sp is nil, then this function returns "".
func String(sp *string) string {
	if sp == nil {
		return ""
	}
	return *sp
}
