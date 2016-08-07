package util

import (
    "crypto/md5"
    "encoding/hex"
)

func GetMD5Hash(text string) string {
  hasher := md5.New()
  hasher.Write([]byte(text))
  enc := hex.EncodeToString(hasher.Sum(nil))
  return enc
}
