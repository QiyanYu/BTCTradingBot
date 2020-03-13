package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

//GetSignature : for encrypt signature
func GetSignature(context string) string {
	secretKey := "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"
	// apiKey := "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A"
	// stdin := "c8db56825ae71d6d79447849e617115f4a920fa2acdcab2b053c4b2838bd6b71"

	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(context))
	return hex.EncodeToString(h.Sum(nil))
}
