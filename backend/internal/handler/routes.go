package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var r *chi.Mux

func RegisterRoutes(handler *handler) *chi.Mux {
	r = chi.NewRouter()
	tokenMaker := handler.tokenMaker
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))

	r.Route("/api", func(api chi.Router) {
		api.Route("/listings", func(r chi.Router) {
			r.Get("/", handler.listAllListingAndReviews)
			r.Get("/categories", handler.listCategory)
		})

		api.Route("/users", func(r chi.Router) {
			r.Post("/login", handler.login)
			r.Post("/register", handler.register)
			r.Post("/refresh", handler.refresh)

			r.Group(func(protectedApi chi.Router) {
				protectedApi.Use(GetAuthMiddlewareFunc(tokenMaker))
				protectedApi.Post("/logout", handler.logout)
				protectedApi.Get("/me", handler.getMe)
			})
		})

		api.Route("/rooms", func(r chi.Router) {
			r.Get("/{roomId}", handler.getRoomById)
		})

		api.Group(func(protectedApi chi.Router) {
			protectedApi.Use(GetAuthMiddlewareFunc(tokenMaker))
			protectedApi.Route("/bookings", func(r chi.Router) {
				r.Get("/", handler.listBookings)
				r.Post("/", handler.createBooking)
				r.Delete("/{bookingId}", handler.deleteBooking)
			})
			protectedApi.Route("/wishlist", func(r chi.Router) {
				r.Get("/", handler.getWishListItems)
				r.Get("/ids", handler.listWishListIds)
				r.Post("/", handler.upsertWishList)
			})
			protectedApi.Route("/reviews", func(r chi.Router) {
				r.Post("/", handler.createReview)
				r.Delete("/{reviewId}", handler.deleteReview)
			})
		})
	})

	return r
}

func Start(addr string) error {
	log.Println("connect server to port 3000")
	return http.ListenAndServe(addr, r)
}
