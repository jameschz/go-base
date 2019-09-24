package base

import "base/gmq/driver"

type MQ struct {
	Driver   *driver.Driver // driver ptr
	NodeName string         // node name
}

type IMQ interface {
	Connect(node string) error
	Close() error
	Shard(sk string) error
	Publish(qName string, qBody string) error
	Consume(qName string, callback func(done chan bool, body []byte)) error
}
