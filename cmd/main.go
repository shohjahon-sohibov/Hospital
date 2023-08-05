package main

import (
	"context"
	// "fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"freelance/clinic_queue/api"
	"freelance/clinic_queue/api/handlers"
	"freelance/clinic_queue/config"
	"freelance/clinic_queue/storage/mongodb"

	"github.com/saidamir98/udevs_pkg/logger"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	mongoString := "mongodb://localhost/27017"

	log.Println("MAIN TEST PRINT " + mongoString)
	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoString))
	if err != nil {
		log.Fatal("error to connect to mongo database", logger.Error(err))
	}

	defer func(mongoConn *mongo.Client, ctx context.Context) {
		err := mongoConn.Disconnect(ctx)
		if err != nil {
			return
		}
	}(mongoConn, context.Background())

	if err := mongoConn.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Cannot connect to database error -> ", logger.Error(err))
	}
	connDB := mongoConn.Database(cfg.MongoDatabase)

	store := mongodb.NewStoragePg(connDB)

	logger.Any("Connected to MongoDB in ", logger.Any("Server: ", cfg.MongoHost))

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	h := handlers.NewHandler(cfg, log, store)

	r := api.SetUpRouter(h, cfg)

	r.Run(cfg.HTTPPort)

}
