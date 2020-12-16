package app

import (
	"MyApp/internal/app/config"
	"MyApp/internal/app/pkg/ent"
	"MyApp/internal/app/pkg/ent/migrate"
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitEntOrm() (*ent.Client, func(), error) {
	cfg := config.C
	dsn := cfg.MySQL.DSN()
	db, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}
	cleanFunc := func() {
		err := db.Close()
		if err != nil {
			log.Printf("Ent orm closed error:%v", err.Error())
		}
	}
	err = db.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return nil, nil, err
	}
	return db, cleanFunc, nil
}
