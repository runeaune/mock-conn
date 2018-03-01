package mock_conn

import (
	"github.com/free/concurrent-writer/concurrent"
	"io"
	"net"
	"time"
)

// End is one 'end' of a simulated connection.
type End struct {
	Reader         *io.PipeReader
	Writer         *io.PipeWriter
	BufferedWriter *concurrent.Writer
}

func (c End) Close() error {
	if err := c.Writer.Close(); err != nil {
		return err
	}
	if err := c.Reader.Close(); err != nil {
		return err
	}
	return nil
}

func (e End) Read(data []byte) (n int, err error) {
	n, err = e.Reader.Read(data)
	return n, err
}
func (e End) Write(data []byte) (n int, err error) {
	n, err = e.BufferedWriter.Write(data)
	// TODO(runeaune): Move the flushing to a single dedicated background thread.
	go e.BufferedWriter.Flush()
	return n, err
}

func (e End) LocalAddr() net.Addr {
	return Addr{
		NetworkString: "tcp",
		AddrString:    "127.0.0.1",
	}
}

func (e End) RemoteAddr() net.Addr {
	return Addr{
		NetworkString: "tcp",
		AddrString:    "127.0.0.1",
	}
}

func (e End) SetDeadline(t time.Time) error      { return nil }
func (e End) SetReadDeadline(t time.Time) error  { return nil }
func (e End) SetWriteDeadline(t time.Time) error { return nil }
