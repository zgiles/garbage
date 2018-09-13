package garbage

import "io"

type DevZero int

func (z DevZero) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

func (z DevZero) Write(b []byte) (n int, err error) {
	return len(b), nil
}

func NewLimitedZeroReader(size int64) io.Reader {
	z := DevZero(0)
	limiter := io.LimitReader(z, size)
	return limiter
}
