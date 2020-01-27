package gcr

// GreatestCommonDivisor returns the largest common factor of two positive integers
func GreatestCommonDivisor(p, q uint) uint {
	if q == 0 {
		return p
	}
	r := p % q
	return GreatestCommonDivisor(q, r)
}
