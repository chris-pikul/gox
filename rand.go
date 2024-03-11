package gox

import "crypto/rand"

// RandomCode returns a randomly generated Base62 code which includes only
// alpha-numeric characters. It uses the crypto random reader.
//
// Note: It uses a modulus operator so the resulting characters may not be
// weighted perfectly flat (gaussian). But for low-security purposes this is
// probably fine.
func RandomCode(n uint8) string {
	const alpha = "0123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	prng := make([]byte, n)
	rand.Read(prng)
	for i := range prng {
		prng[i] = byte(alpha[int(prng[i])%len(alpha)])
	}
	return string(prng)
}

// RandomUpperCode returns a randomly generated code using digits and upper-case
// latin characters. It uses the crypto random reader.
//
// Note: It uses a modulus operator so the resulting characters may not be
// weighted perfectly flat (gaussian). But for low-security purposes this is
// probably fine.
func RandomUpperCode(n uint8) string {
	const alpha = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	prng := make([]byte, n)
	rand.Read(prng)
	for i := range prng {
		prng[i] = byte(alpha[int(prng[i])%len(alpha)])
	}
	return string(prng)
}
