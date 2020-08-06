package gmq

import (
	"go-base/lib/gmq/base"
	"go-base/lib/gmq/driver"
	"go-base/lib/gmq/rabbitmq"
)

// connect by driver
func D(cs string) (imq base.IMQ) {
	// init driver
	driver.Init()
	// get mq driver
	mq_driver := driver.GetDriver(cs)
	if len(mq_driver.Type) == 0 {
		panic("gmq> mq driver error")
	}
	// mq initialize
	switch mq_driver.Type {
	case "rabbitmq":
		rmq := &rabbitmq.RabbitMQ{}
		rmq.Driver = mq_driver
		imq = rmq
	default:
		panic("gmq> unknown driver type")
	}
	return imq
}
