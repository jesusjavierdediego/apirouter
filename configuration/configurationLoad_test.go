package configuration

import (
	"fmt"
	"os"
	"testing"
	"strconv"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConfigurationDevProfile(t *testing.T) {

	Convey("Loading server configuration values ", t, func() {
		os.Setenv("PROFILE", "dev")
		Reload()
		env := os.Getenv("PROFILE")
		So(env, ShouldEqual, "dev")
		os.Remove("PROFILE")
		Reload()
	})
}

func TestConfiguration(t *testing.T) {
	Convey("Reading server configuration values ", t, func() {
		os.Setenv("PROFILE", "release")
		Reload()
		valueK := GlobalConfiguration.Kafka.Bootstrapserver
		valueT := GlobalConfiguration.Kafka.Sessiontimeout
		valueR := GlobalConfiguration.Grpcclient.Gitreaderport

		fmt.Println("valueK: " + valueK)
		fmt.Println("valueT: " + strconv.Itoa(valueT))
		fmt.Println(valueR)

		So(len(valueK), ShouldBeGreaterThan, 0)
		So(valueT, ShouldEqual, 5000)
		So(valueR, ShouldEqual, 50051)
	})
}
