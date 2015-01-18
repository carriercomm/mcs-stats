package main

import (
	"github.com/coopernurse/gorp"
	"github.com/pmylund/go-cache"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"encoding/json"
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

	goji.Post("/api/pings", HandlePing)

	// Static Files
	goji.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir(config.StaticLocation))))

	// Serve using the magic of Goji!
	goji.Serve()
}

func HandlePing(c web.C, w http.Responsewriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	ping := &Ping{}
	err := decoder.Decode(ping)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write("Malformed Body")
		return
	}
	ping.Timestamp = time.Now()
	ping.NextPing = time.Now().Add(15*time.Minute)
	// Begin validation of different fields.
	if ping.Timestamp
}
