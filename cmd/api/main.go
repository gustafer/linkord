package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gustafer/linkord/configs"
	"github.com/gustafer/linkord/internal/auth"
	"github.com/gustafer/linkord/internal/database"
	"github.com/gustafer/linkord/internal/routes"
)

func main() {
	r := chi.NewRouter()
	routes.SetupRoutes(r)

	auth.NewAuth()
	err := database.Migrate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`
     __    _       __                  __
    / /   (_)___  / /______  _________/ /
   / /   / / __ \/ //_/ __ \/ ___/ __  / 
  / /___/ / / / / ,< / /_/ / /  / /_/ /  
 /_____/_/_/ /_/_/|_|\____/_/   \__,_/ 

 ----------------------------------------
	`)
	port := configs.LoadPort()
	fmt.Printf("Server running on \033[35mhttp://%v\033[0m\n", port)
	http.ListenAndServe(port, r)
}
