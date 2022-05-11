package main

import (
	"flag"
	"github.com/morozvol/money_manager_api/internal/config"
	"github.com/morozvol/money_manager_api/pkg/store/sqlstore"
	dbConfig "github.com/morozvol/money_manager_api/pkg/store/sqlstore/config"
	"github.com/morozvol/money_manager_api/pkg/store/sqlstore/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"

	"github.com/morozvol/money_manager_api/internal/server"
)

func main() {
	var configDir, dbConfigName string
	flag.StringVar(&configDir, "config_path", "", "path to config directory")
	flag.StringVar(&dbConfigName, "db_config", "db", "data base config name")
	flag.Parse()
	if configDir == "" {
		log.Fatal("не передан обязательный параметр -config_path")
		return
	}

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	conf, err := config.Init(configDir)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger, _ = zap.NewProduction()
	logger.WithOptions(zap.IncreaseLevel(zapcore.DebugLevel))
	zap.ReplaceGlobals(logger)

	baseConfig, err := dbConfig.GetDataBaseConfig(dbConfigName, configDir)
	if err != nil {
		return
	}

	dataBase, err := db.New(baseConfig, logger)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	store := sqlstore.New(dataBase)
	s := server.New(store, conf)
	err = s.Start()
	if err != nil {
		return
	}
}
