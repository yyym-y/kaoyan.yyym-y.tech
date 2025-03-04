package middle

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	srcByte := []byte(str)
	md5New := md5.New()
	md5Bytes := md5New.Sum(srcByte)
	md5String := hex.EncodeToString(md5Bytes)
	print(md5String)
	return md5String
}
