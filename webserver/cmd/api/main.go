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
 ______     __  __     ______     __         __         __  __     ______     ______     __  __     ______    
/\  ___\   /\ \_\ \   /\  ___\   /\ \       /\ \       /\ \_\ \   /\  __ \   /\  ___\   /\ \/ /    /\  ___\   
\ \___  \  \ \  __ \  \ \  __\   \ \ \____  \ \ \____  \ \  __ \  \ \  __ \  \ \ \____  \ \  _"-.  \ \___  \  
 \/\_____\  \ \_\ \_\  \ \_____\  \ \_____\  \ \_____\  \ \_\ \_\  \ \_\ \_\  \ \_____\  \ \_\ \_\  \/\_____\ 
  \/_____/   \/_/\/_/   \/_____/   \/_____/   \/_____/   \/_/\/_/   \/_/\/_/   \/_____/   \/_/\/_/   \/_____/ 
                                                                                                               		`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
