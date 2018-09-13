package garbage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	fastaes "github.com/bronze1man/AesCtr"
	"io"
)

func NewReader() io.Reader {
	var key [aes.BlockSize]byte
	var iv [aes.BlockSize]byte
	rand.Read(key[:])
	rand.Read(iv[:])
	a, _ := fastaes.NewCipher(key[:])
	c := cipher.NewCTR(a, iv[:])
	z := DevZero(0)
	reader := &cipher.StreamReader{S: c, R: z}
	return reader
}

func NewLimitedReader(size int64) io.Reader {
	reader := NewReader()
	limiter := io.LimitReader(reader, size)
	return limiter
}
