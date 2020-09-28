package gmq

import (
	"github.com/jameschz/go-base/lib/gmq/base"
	"github.com/jameschz/go-base/lib/gmq/driver"
	"github.com/jameschz/go-base/lib/gmq/rabbitmq"
)

// D : connect by driver
func D(cs string) (imq gmqbase.IMQ) {
	// init driver
	gmqdriver.Init()
	// get mq driver
	mqDriver := gmqdriver.GetDriver(cs)
	if len(mqDriver.Type) == 0 {
		panic("gmq> mq driver error")
	}
	// mq initialize
	switch mqDriver.Type {
	case "rabbitmq":
		rmq := &gmqrabbitmq.RabbitMQ{}
		rmq.Driver = mqDriver
		imq = rmq
	default:
		panic("gmq> unknown driver type")
	}
	return imq
}
