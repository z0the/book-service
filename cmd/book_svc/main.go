package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"books-service/internal/book_svc"
	"books-service/internal/storage"
	http2 "books-service/internal/transport/http"

	"github.com/opentracing/opentracing-go"
)

func main() {
	time.Local = time.UTC
	cfg := GetConfig()

	var logger *zap.Logger
	var err error
	if cfg.IsDev {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatal("failed to initialize logger", err)
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatal("failed to initialize logger", err)
		}
	}
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatal("failed sync logger", err)
		}
	}()
	sugar := logger.Sugar()

	// Init tracing
	defer func(tracing io.Closer) {
		err := tracing.Close()
		if err != nil {
			panic(err)
		}
	}(initTracing(&cfg.JaegerConfig))

	db, err := gorm.Open(postgres.Open(cfg.GetDSN()))
	if err != nil {
		sugar.Fatal(err)
	}
	st := storage.NewStorage(sugar, db)
	svc := book_svc.NewService(&book_svc.ServiceConfig{
		Logger:  sugar,
		Storage: st,
	})
	ctrl := http2.NewController(svc)
	router := mux.NewRouter()
	http2.LoadHTTPRouterWithEndpoints(router, ctrl)
	sugar.Infof("Starting HTTP server on port %s...\n", cfg.HttpPort)
	sugar.Fatal(http.ListenAndServe(cfg.HttpPort, router))
}

func initTracing(cfg *JaegerConfig) io.Closer {
	tracingCfg := config.Configuration{
		ServiceName: cfg.ServiceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: cfg.AgentHost + ":" + cfg.AgentPort,
			LogSpans:           false,
		},
	}
	tracer, closer, err := tracingCfg.NewTracer()
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)

	return closer
}
