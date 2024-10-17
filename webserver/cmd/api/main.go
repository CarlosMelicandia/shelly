package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/weareinit/Opal/internal/handlers"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service on localhost:8000...")

	fmt.Println(`
  ██████  ██░ ██ ▓█████  ██▓     ██▓     ██░ ██  ▄▄▄       ▄████▄   ██ ▄█▀  ██████ 
▒██    ▒ ▓██░ ██▒▓█   ▀ ▓██▒    ▓██▒    ▓██░ ██▒▒████▄    ▒██▀ ▀█   ██▄█▒ ▒██    ▒ 
░ ▓██▄   ▒██▀▀██░▒███   ▒██░    ▒██░    ▒██▀▀██░▒██  ▀█▄  ▒▓█    ▄ ▓███▄░ ░ ▓██▄   
  ▒   ██▒░▓█ ░██ ▒▓█  ▄ ▒██░    ▒██░    ░▓█ ░██ ░██▄▄▄▄██ ▒▓▓▄ ▄██▒▓██ █▄   ▒   ██▒
▒██████▒▒░▓█▒░██▓░▒████▒░██████▒░██████▒░▓█▒░██▓ ▓█   ▓██▒▒ ▓███▀ ░▒██▒ █▄▒██████▒▒
▒ ▒▓▒ ▒ ░ ▒ ░░▒░▒░░ ▒░ ░░ ▒░▓  ░░ ▒░▓  ░ ▒ ░░▒░▒ ▒▒   ▓▒█░░ ░▒ ▒  ░▒ ▒▒ ▓▒▒ ▒▓▒ ▒ ░
░ ░▒  ░ ░ ▒ ░▒░ ░ ░ ░  ░░ ░ ▒  ░░ ░ ▒  ░ ▒ ░▒░ ░  ▒   ▒▒ ░  ░  ▒   ░ ░▒ ▒░░ ░▒  ░ ░
░  ░  ░   ░  ░░ ░   ░     ░ ░     ░ ░    ░  ░░ ░  ░   ▒   ░        ░ ░░ ░ ░  ░  ░  
      ░   ░  ░  ░   ░  ░    ░  ░    ░  ░ ░  ░  ░      ░  ░░ ░      ░  ░         ░  
                                                          ░                        `)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
