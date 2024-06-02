package ctxio

import (
	"context"
	"io"
)

func IoReaderAt(ctx context.Context, r ReaderAt) io.ReaderAt {
	return &contextReaderAt{ctx: ctx, r: r}
}

type contextReader struct {
	ctx context.Context
	r   Reader
}

func (r *contextReader) Read(p []byte) (n int, err error) {
	if r.ctx.Err() != nil {
		return 0, r.ctx.Err()
	}

	return r.r.Read(r.ctx, p)
}

func IoReader(ctx context.Context, r Reader) io.Reader {
	return &contextReader{ctx: ctx, r: r}
}

type contextReaderAt struct {
	ctx context.Context
	r   ReaderAt
}

func (c *contextReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	if c.ctx.Err() != nil {
		return 0, c.ctx.Err()
	}

	return c.r.ReadAt(c.ctx, p, off)
}

func IoWriter(ctx context.Context, w Writer) io.Writer {
	return &contextWriter{ctx: ctx, w: w}
}

type contextWriter struct {
	ctx context.Context
	w   Writer
}

func (c *contextWriter) Write(p []byte) (n int, err error) {
	if c.ctx.Err() != nil {
		return 0, c.ctx.Err()
	}

	return c.w.Write(c.ctx, p)
}
