package serve

import (
	"fmt"
	"log"
	"net/http"
)

func (engine Engine) Serve() error {
	var (
		err        error
		fileServer http.Handler
	)

	if engine.Mode == InMemoryOnly || engine.Mode == InMemoryFirst {
		if engine.Verbose {
			log.Println("Amount of InMemoryFiles:", len(engine.InMemoryFiles))
		}
		for path, content := range engine.InMemoryFiles {
			// Adding handlers for in-memory requests (when used as library) to webserver
			http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				if engine.Verbose {
					log.Println("Served InMemoryFile for URL:", r.RequestURI)
				}
				_, err = w.Write([]byte(content))
				if err != nil {
					fmt.Println(err)
				}
				w.WriteHeader(http.StatusOK)
			})
		}
	}

	if engine.Mode == FileOnly || engine.Mode == InMemoryFirst {
		// Creating Webserver for static files
		if engine.Verbose {
			log.Println("Serving directory", engine.Directory)
		}
		fileServer = http.FileServer(http.Dir(engine.Directory))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if engine.Verbose {
				log.Println("Served file for URL:", r.RequestURI)
			}
			fileServer.ServeHTTP(w, r)
		})
	}

	// Starting Webserver
	log.Println("Listening on", engine.Ipaddress+":"+engine.Port, "...")
	err = http.ListenAndServe(engine.Ipaddress+":"+engine.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
