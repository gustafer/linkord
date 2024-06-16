package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gustafer/linkord/internal/handlers"
	"github.com/gustafer/linkord/internal/middleware"
)

func SetupRoutes(r *chi.Mux) {
	// logger := httplog.NewLogger("httplog-example", httplog.Options{
	// 	// JSON:             true,
	// 	LogLevel:         slog.LevelDebug,
	// 	Concise:          true,
	// 	RequestHeaders:   true,
	// 	MessageFieldName: "message",
	// 	// TimeFieldFormat: time.RFC850,
	// 	Tags: map[string]string{
	// 		"version": "v1.0-81aa4244d9fc8076a",
	// 		"env":     "dev",
	// 	},
	// 	QuietDownRoutes: []string{
	// 		"/",
	// 		"/ping",
	// 	},
	// 	QuietDownPeriod: 10 * time.Second,
	// 	// SourceFieldName: "source",
	// })

	r.Use(
		// httplog.RequestLogger(logger),
		middleware.SetHeader("Access-Control-Allow-Credentials", "true"),
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"http://*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{
				"Accept", "Authorization",
				"Content-Type", "X-CSRF-Token", "X-Requested-With", "Origin",
			},
		}))
	r.Get("/", handlers.ProjectInfo)
	r.Get("/health", handlers.Health)

	r.Post("/game", handlers.CreateGame)

	r.Get("/auth/{provider}/callback", handlers.GetAuthCallbackFunction)
	r.Get("/auth/{provider}", handlers.GetAuth)
	r.Get("/logout/{provider}", handlers.GetLogout)

	r.Get("/users", handlers.GetUsers)
	r.Group(ProtectedRoutes)
}

func ProtectedRoutes(r chi.Router) {
	r.Use(middleware.AuthMiddleware)

	r.Get("/user", handlers.GetUser)
	r.Get("/user/game", handlers.PrivateUserInfo)
}
