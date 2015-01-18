package main

import (
	"github.com/coopernurse/gorp"
	"github.com/pmylund/go-cache"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"math/rand"
	"net/http"
	"time"
	"os"
)

var (
	config     = configuration{}
	server     Server
	dbmap      *gorp.DbMap
	dataCache  *cache.Cache
	emailQueue chan *Email
)

func init() {
	// Seed the RNG. Only needs doing once at startup.
	rand.Seed(time.Now().UTC().UnixNano())
	dataCache = cache.New(time.Minute, 5*time.Second)
	// Set up the database.
	dbmap = initDatabase()
}

func main() {
	// Default Middleware
	goji.Use(middleware.EnvInit)
	goji.Use(middleware.Recoverer)

	goji.Post("/api/pings", pingHandler)

	// Static Files
	goji.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir(config.StaticLocation))))

	// Serve using the magic of Goji!
	goji.Serve()
}
