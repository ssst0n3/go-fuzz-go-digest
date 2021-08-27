package go_fuzz_go_digest

import "github.com/opencontainers/go-digest"

func FuzzParse(b []byte) int {
	digest.Parse(string(b))
	return 0
}
