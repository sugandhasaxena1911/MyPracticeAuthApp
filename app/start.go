package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/helpers/logger"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/adapters/handlers"
	"github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/adapters/repositoryDB"
	coreservice "github.com/sugandhasaxena1911/MyPracticeAuthApp/internal/core/service"
)

var DBClient *sql.DB

func init() {
	// Load env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Cannot load env variables ")
	}
	// sanity check

	if os.Getenv("DBNAME") == "" || os.Getenv("DBHOST") == "" || os.Getenv("DBPASSWORD") == "" || os.Getenv("DBUSER") == "" || os.Getenv("DBPORT") == "" || os.
		Getenv("SECRETKEY") == "" {
		logger.Error("Env variables are not defined")
	}
}

func init() {

	// take db connection
	logger.Info("taking db connection")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")
	dbhost := os.Getenv("DBHOST")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbhost, dbport, dbname)
	var err error
	DBClient, err = sql.Open("mysql", connstr)
	if err != nil {
		logger.Error(err.Error())
		panic("Cannot take db connection ")
	}

	err = DBClient.Ping()
	if err != nil {
		log.Panicln("Cannot ping DB", err)
	}
	logger.Info("Db ping successful ")
	log.Println(DBClient)
	DBClient.SetConnMaxLifetime(time.Minute * 3)
	DBClient.SetMaxOpenConns(10)
	DBClient.SetMaxIdleConns(10)
}
func Start() {
	log.Println("Register routes ", DBClient)
	router := mux.NewRouter()
	usrservice := coreservice.NewUserCoreService(repositoryDB.NewUserRespositoryDB(DBClient))
	userhandler := handlers.UserHandler{Usrservice: usrservice}
	router.HandleFunc("/auth", handlers.GetTestAuth).Methods(http.MethodGet)
	router.HandleFunc("/auth/register", userhandler.PostUser).Methods(http.MethodPost)

	loginservice := handlers.AuthHandler{Authservice: coreservice.NewLoginCoreService(repositoryDB.NewAuthRepositoryDB(DBClient))}
	router.HandleFunc("/login", loginservice.FetchLogindetails).Methods(http.MethodGet)

	http.ListenAndServe("localhost:8000", router)

}
