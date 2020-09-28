package gmqbase

import (
	"github.com/jameschz/go-base/lib/gmq/driver"
)

// MQ :
type MQ struct {
	Driver   *gmqdriver.Driver // driver ptr
	NodeName string            // node name
}

// IMQ :
type IMQ interface {
	Connect(node string) error
	Close() error
	Shard(sk string) error
	Publish(qName string, qBody string) error
	Consume(qName string, callback func(done chan bool, body []byte)) error
}
