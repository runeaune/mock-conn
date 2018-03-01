package mock_conn

import (
	"github.com/free/concurrent-writer/concurrent"
	"io"
)

const bufferSize = 1 << 16

// MockConn facilitates testing by providing two connected ReadWriteClosers
// each of which can be used in place of a net.Conn
type Conn struct {
	Server *End
	Client *End
}

func NewConn() *Conn {
	// A connection consists of two pipes:
	// Client      |      Server
	//   writes   ===>  reads
	//    reads  <===   writes

	serverRead, clientWrite := io.Pipe()
	clientRead, serverWrite := io.Pipe()

	return &Conn{
		Server: &End{
			Reader:         serverRead,
			Writer:         serverWrite,
			BufferedWriter: concurrent.NewWriterSize(serverWrite, bufferSize),
		},
		Client: &End{
			Reader:         clientRead,
			Writer:         clientWrite,
			BufferedWriter: concurrent.NewWriterSize(clientWrite, bufferSize),
		},
	}
}

func (c *Conn) Close() error {
	if err := c.Server.Close(); err != nil {
		return err
	}
	if err := c.Client.Close(); err != nil {
		return err
	}
	return nil
}
