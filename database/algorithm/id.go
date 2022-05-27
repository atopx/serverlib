package algorithm

import (
	"crypto/md5"
	"math/big"
)

func ShortMd5(str string) *big.Int {
	uid := new(big.Int)
	if str == "" {
		return uid
	}
	md5sum := md5.Sum([]byte(str))
	buf := make([]byte, 4)
	for i := range buf {
		buf[i] = (md5sum[:4][3-i] ^ md5sum[4:8][3-i]) ^ (md5sum[8:12][3-i] ^ md5sum[12:][3-i])
	}
	uid.SetBytes(buf)
	return uid
}

func ShortMd5ToInt64(str string) int64 {
	return ShortMd5(str).Int64()
}

func ShortMd5ToString(str string) string {
	return ShortMd5(str).String()
}
