package fakewriteseeker

import (
	"errors"
	"io"
)

type fakeWriteSeeker struct {
	w io.Writer
	n int64
}

func NewFakeWriteSeeker(w io.Writer) io.WriteSeeker {
	return &fakeWriteSeeker{w: w}
}

func (w *fakeWriteSeeker) Write(p []byte) (n int, err error) {
	n, err = w.w.Write(p)
	w.n += int64(n)
	return
}

func (w *fakeWriteSeeker) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekCurrent {
		return w.n, nil
	}
	return 0, errors.New("not implemented")
}
