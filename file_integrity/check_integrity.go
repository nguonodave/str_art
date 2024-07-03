package file_integrity

import (
	"crypto/sha256"
	"encoding/hex"
)

// FileAltered returns true if the any of the banner files in use is corrupted.
// A file is consdered corrupted if it's contents are altered (deleted or added content)
func FileAltered(file_bytes []byte, original_hash string) bool {
	hash_object := sha256.New()

	hash_object.Write(file_bytes)

	hash_sum := hash_object.Sum(nil)

	hash_string := hex.EncodeToString(hash_sum)

	if hash_string == original_hash {
		return false
	}

	return hash_string != original_hash
}
