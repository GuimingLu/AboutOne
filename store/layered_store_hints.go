


package store

type LayeredStoreHint int

const (
	LSH_NO_CACHE LayeredStoreHint = iota
	LSH_MASTER_ONLY
)

func hintsContains(hints []LayeredStoreHint, contains LayeredStoreHint) bool {
	for _, hint := range hints {
		if hint == contains {
			return true
		}
	}
	return false
}
