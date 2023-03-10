package main

import (
	"github.com/rs/cors"
	"net/http"
)

func CorsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPut,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhost:4000",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Authorization",
		},
		OptionsPassthrough: true,
		ExposedHeaders: []string{
			"Access-Control-Allow-Origin",
			"Content-Type",
		},
		Debug: false,
	})
	return c
}
