package before

import (
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	grpcServices "github.com/payfazz/kitx/cmd/grpc"
	httpServices "github.com/payfazz/kitx/cmd/http"

	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var g group.Group

	GRPCAddress := ":8080"
	{
		grpcListener, err := net.Listen("tcp", GRPCAddress)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", ":8080")
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			grpcServices.Init(baseServer, logger)
			reflection.Register(baseServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}

	HTTPAddress := ":9080"
	{
		httpLogger := log.With(logger, "component", "http")
		router := chi.NewRouter()
		router.Handle("/metrics", promhttp.Handler())
		router.Mount("/v1", httpServices.MakeHandler(router, httpLogger))
		g.Add(func() error {
			logger.Log("transport", "debug/HTTP", "addr", HTTPAddress)
			return http.ListenAndServe(HTTPAddress, router)
		}, func(error) {

		})
	}

	logger.Log("exit", g.Run())
}
