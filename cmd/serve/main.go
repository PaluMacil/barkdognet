package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/PaluMacil/barkdognet/configuration"
	"github.com/PaluMacil/barkdognet/dist/ui"
	"github.com/PaluMacil/barkdognet/site"
	"github.com/PaluMacil/barkdognet/templates"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configProvider := configuration.DefaultProvider{}
	config, err := configProvider.Config()
	if err != nil {
		log.Fatalf("loading configuration: %v", err)
	}

	tmplSrc := templates.NewTemplateSource(config.Site.LiveTemplates)
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if tmpl, err := tmplSrc.Get(""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			layout := site.Layout{
				Site:        config.Site,
				Title:       "Home",
				Author:      "Dan Wolf",
				Keywords:    "",
				Description: "",
			}
			data := struct {
				Layout site.Layout
				word   string
			}{layout, "blah"}
			if err := tmpl.Execute(w, data); err != nil { // replace nil with actual data
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		if tmpl, err := tmplSrc.Get("admin"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			layout := site.Layout{
				Site:        config.Site,
				Title:       "Admin",
				Author:      "Dan Wolf",
				Keywords:    "",
				Description: "",
			}
			data := struct {
				Layout site.Layout
				word   string
			}{layout, "blah"}
			if err := tmpl.Execute(w, data); err != nil { // replace nil with actual data
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	var staticFS fs.FS
	if config.Site.LiveTemplates {
		staticFS = os.DirFS("dist/ui")
	} else {
		staticFS = ui.EmbeddedStaticFS
	}
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServerFS(staticFS)))

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Site.ListenAddr, config.Site.ListenPort),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	log.Printf("Now serving %s", config.Site.BaseURL)
	<-stop

	log.Println("Shutting down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
