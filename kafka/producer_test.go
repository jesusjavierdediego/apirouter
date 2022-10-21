package kafka

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSendEventMessageToTopic(t *testing.T) {
	Convey("Send Kafka message ", t, func() {
		err := SendMessageToTopic("Hello", "gitoperator-in")
		So(err, ShouldBeNil)
	})
}