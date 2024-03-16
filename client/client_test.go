package client_test

import (
	"testing"

	c "github.com/alexfalkowski/go-service/client"
	"github.com/alexfalkowski/go-service/retry"
	"github.com/alexfalkowski/go-service/telemetry/tracer"
	v1c "github.com/alexfalkowski/standort/client/v1/config"
	v1 "github.com/alexfalkowski/standort/client/v1/transport/grpc"
	v2c "github.com/alexfalkowski/standort/client/v2/config"
	v2 "github.com/alexfalkowski/standort/client/v2/transport/grpc"
	. "github.com/smartystreets/goconvey/convey" //nolint:revive
	"go.uber.org/fx/fxtest"
)

func TestV1Client(t *testing.T) {
	Convey("Given I have the correct parameters", t, func() {
		lc := fxtest.NewLifecycle(t)

		params := v1.ServiceClientParams{
			Lifecycle:    lc,
			ClientConfig: &v1c.Config{Config: c.Config{Host: "localhost", Retry: retry.Config{Timeout: "1s"}}},
			Tracer:       tracer.NewNoopTracer("test"),
		}

		Convey("When I create a new client", func() {
			client, err := v1.NewServiceClient(params)
			So(err, ShouldBeNil)

			lc.RequireStart()

			Convey("Then I should have a valid client", func() {
				So(client, ShouldNotBeNil)
			})

			lc.RequireStop()
		})
	})
}

func TestV2Client(t *testing.T) {
	Convey("Given I have the correct parameters", t, func() {
		lc := fxtest.NewLifecycle(t)

		params := v2.ServiceClientParams{
			Lifecycle:    lc,
			ClientConfig: &v2c.Config{Config: c.Config{Host: "localhost", Retry: retry.Config{Timeout: "1s"}}},
			Tracer:       tracer.NewNoopTracer("test"),
		}

		Convey("When I create a new client", func() {
			client, err := v2.NewServiceClient(params)
			So(err, ShouldBeNil)

			lc.RequireStart()

			Convey("Then I should have a valid client", func() {
				So(client, ShouldNotBeNil)
			})

			lc.RequireStop()
		})
	})
}
