package storage

import (
	"context"
	"io"
)

type FS struct {
	StripeMgr interface{}
}

func NewFS() (Storage, error) {
	return &FS{}, nil
}

func (f *FS) Put(ctx context.Context, data io.Reader) (string, int64, error) {
	return "", 0, nil
}

func (f *FS) Get(ctx context.Context, fd string) (io.ReadCloser, error) {
	return nil, nil
}

func (f *FS) Delete(ctx context.Context, fd string) error {
	return nil
}
