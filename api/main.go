package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"violet.wtf/tunnel/api/appctx"
	"violet.wtf/tunnel/api/controller/dashboardctrl"
	"violet.wtf/tunnel/api/controller/loginctrl"
	"violet.wtf/tunnel/api/service/migrationsvc"
	"violet.wtf/tunnel/api/tutil"
)

func main() {
	connectToDb()
	doMigrations()

	http.HandleFunc("/login", loginctrl.LoginRoute)
	http.HandleFunc("/register", loginctrl.RegisterRoute)
	http.HandleFunc("/tokenExists", loginctrl.TokenExistsRoute)
	http.HandleFunc("/dashboard", dashboardctrl.DashboardRoute)

	fmt.Println("Starting")
	panic(http.ListenAndServe("0.0.0.0:8080", nil))
}

func connectToDb() {
	// postgres://pqgotest:password@localhost/pqgotest?sslmode=disable
	conn, err := sql.Open(
		"postgres",
		"postgres://"+
			os.Getenv("TUNNEL_PG_USER")+":"+
			os.Getenv("TUNNEL_PG_PASSWORD")+"@"+
			os.Getenv("TUNNEL_PG_HOST")+":"+
			os.Getenv("TUNNEL_PG_PORT")+"/"+
			os.Getenv("TUNNEL_PG_DATABASE")+"?sslmode="+
			os.Getenv("TUNNEL_PG_SSL_MODE"),
	)
	tutil.PanicIfErr(err)

	appctx.DB = conn
}

func doMigrations() {
	tutil.PanicIfErr(migrationsvc.EnsureMigrationTableExists())

	paths, err := ioutil.ReadDir("./migrations")
	tutil.PanicIfErr(err)

	currentMigrationVersion, err := migrationsvc.GetVersion()
	tutil.PanicIfErr(err)

	highestMigrationVersion := 0

	for _, path := range paths {
		// 0000-example.sql -> 0000-example
		pathNoExtension := path.Name()[:len(path.Name())-4]

		// 0000-example -> ["0000", "example"]
		parts := strings.Split(pathNoExtension, "-")

		migrationVersion, err := strconv.Atoi(parts[0])
		tutil.PanicIfErr(err)

		if migrationVersion > currentMigrationVersion {
			highestMigrationVersion = migrationVersion

			fileText, err := ioutil.ReadFile("./migrations/" + path.Name())
			tutil.PanicIfErr(err)

			_, err = appctx.DB.Exec(string(fileText))
			tutil.PanicIfErr(err)

			fmt.Println("Running migration:  " + pathNoExtension)
		} else {
			fmt.Println("Skipping migration: " + pathNoExtension)
		}

		if highestMigrationVersion != 0 {
			tutil.PanicIfErr(migrationsvc.SetVersion(highestMigrationVersion))
		}
	}
}
