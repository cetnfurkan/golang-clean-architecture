package main

import (
	"embed"
	"golang-clean-architecture/client/rabbitmq"
	"golang-clean-architecture/config"
	"golang-clean-architecture/database"
	"golang-clean-architecture/server"
	"golang-clean-architecture/target/ent"
	"os"
	"os/signal"
	"syscall"

	"github.com/cetnfurkan/core/cache"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
)

var (
	grpcEntryName = "userService"

	//go:embed boot.yaml
	boot []byte

	//go:embed target/grpc/user
	docsFS embed.FS

	//go:embed target/grpc/user
	staticFS embed.FS
)

func init() {
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.DocsEntryType, grpcEntryName, &docsFS)
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.SWEntryType, grpcEntryName, &docsFS)
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.StaticFileHandlerEntryType, grpcEntryName, &staticFS)
}

// @title xebula/xebula-crypto-service Go API
// @version 1.0
// @description This is a sample server for XArch Go API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://support.xebula.io
// @contact.email support@xebula.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	var (
		cfg        = config.Get()
		cache      = cache.NewRedisCache(&cfg.Redis)
		db         = database.NewPostgresDatabase(&cfg.PGSQL).Get().(*ent.Client)
		echoServer = server.NewEchoServer(&cfg, db, cache)
		grpcServer = server.NewGRPCServer(nil, db, cache, grpcEntryName, boot)
		rabbitmMQ  = rabbitmq.NewRabbitMQ(&cfg, db, cache)
	)

	echoServer.Start()
	grpcServer.Start()
	rabbitmMQ.Consume()

	waitSignal()
}

// waitSignal waits for signals from OS.
// It will block the calling goroutine until receiving signals.
func waitSignal() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)
	<-signalChannel
}
