package ctxio

import (
	"context"
	"io"
)

func WrapIoReader(r io.Reader) Reader {
	return &wrapReader{r: r}
}

type wrapReader struct {
	r io.Reader
}

var _ Reader = (*wrapReader)(nil)

// Read implements Reader.
func (c *wrapReader) Read(ctx context.Context, p []byte) (n int, err error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return c.r.Read(p)
}

func WrapIoReaderAt(r io.ReaderAt) ReaderAt {
	return &wrapReaderAt{r: r}
}

type wrapReaderAt struct {
	r io.ReaderAt
}

// Read implements Reader.
func (c *wrapReaderAt) ReadAt(ctx context.Context, p []byte, off int64) (n int, err error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return c.r.ReadAt(p, off)
}

func WrapIoWriter(w io.Writer) Writer {
	return &wrapWriter{w: w}
}

type wrapWriter struct {
	w io.Writer
}

var _ Writer = (*wrapWriter)(nil)

// Write implements Writer.
func (c *wrapWriter) Write(ctx context.Context, p []byte) (n int, err error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return c.w.Write(p)
}

func WrapIoWriterAt(w io.WriterAt) WriterAt {
	return &wrapWriterAt{w: w}
}

type wrapWriterAt struct {
	w io.WriterAt
}

var _ Writer = (*wrapWriter)(nil)

// Write implements Writer.
func (c *wrapWriterAt) WriteAt(ctx context.Context, p []byte, off int64) (n int, err error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return c.w.WriteAt(p, off)
}

func WrapIoWriterTo(w io.WriterTo) WriterTo {
	return &wrapWriterTo{w: w}
}

type wrapWriterTo struct {
	w io.WriterTo
}

var _ WriterTo = (*wrapWriterTo)(nil)

// Write implements Writer.
func (c *wrapWriterTo) WriteTo(ctx context.Context, w Writer) (int64, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return c.w.WriteTo(IoWriter(ctx, w))
}

func WrapIoReadCloser(r io.ReadCloser) ReadCloser {
	return &wrapReadCloser{r: r}
}

type wrapReadCloser struct {
	r io.ReadCloser
}

var _ Reader = (*wrapReadCloser)(nil)

// Read implements Reader.
func (c *wrapReadCloser) Read(ctx context.Context, p []byte) (n int, err error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return c.r.Read(p)
}

// Close implements ReadCloser.
func (c *wrapReadCloser) Close(ctx context.Context) error {
	return c.r.Close()
}

type wrapReadWriter struct {
	Reader
	Writer
}

func WrapReadWriter(rw io.ReadWriter) ReadWriter {
	return wrapReadWriter{
		Reader: WrapIoReader(rw),
		Writer: WrapIoWriter(rw),
	}
}
