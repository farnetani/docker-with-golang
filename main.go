package main

import(
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func main(){
	var logrusLog = logrus.New()
	logrusLog.Infof("Connecting....")
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://dailystaruser:dailystaruser@localhost:9432/dailystardb?sslmode=disable")
	if err != nil {
		logrusLog.Errorf("Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	logrusLog.Infof("Connection ok")
	dbpool.Close()
}