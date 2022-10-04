package serve

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

func Serve(settings ...Setting) error {
	var (
		err        error
		fileServer http.Handler

		// Setting defaults
		directory     string            = "./"
		verbose       bool              = true
		ipaddress     string            = "0.0.0.0"
		port          string            = "3000"
		inMemoryFiles map[string]string = map[string]string{}
	)

	for _, setting := range settings {
		switch setting.Type {
		case Directory:
			value, ok := setting.Value.(string)
			if ok {
				directory = path.Clean(value)
			} else {
				// Directory is not a path
				return errors.New("provided setting 'directory' is not a string")
			}
		case Verbose:
			value, ok := setting.Value.(bool)
			if ok {
				verbose = value
			} else {
				// Verbose is not a bool
				return errors.New("provided setting 'verbose' is not a string")
			}
		case IPAddress:
			value, ok := setting.Value.(string)
			if ok {
				ipaddress = value
			} else {
				// IPAddress is not a string
				return errors.New("provided setting 'IPAddress' is not a string")
			}
		case Port:
			value, ok := setting.Value.(string)
			if ok {
				port = value
			} else {
				// Port is not a bool
				return errors.New("provided setting 'Port' is not a string")
			}
		case InMemoryFile:
			value, ok := setting.Value.(map[string]string)
			if ok {
				for key, value := range value {
					inMemoryFiles[key] = value
				}
			} else {
				// InMemoryFile is not a map[string]string
				return errors.New("provided setting 'InMemoryFile' is not a map[string]string")
			}
		}
	}

	if verbose {
		log.Println("INF verbose=true")
		log.Println("INF directory=" + directory)
		log.Println("INF ipaddress=" + ipaddress)
		log.Println("INF port=" + port)
	}

	if len(inMemoryFiles) > 0 {
		// Adding handler for api requests to webserver
		http.HandleFunc("/inmemory", func(w http.ResponseWriter, r *http.Request) {

			path := strings.TrimPrefix(r.URL.Path, "/inmemory")

			value, ok := inMemoryFiles[path]
			if ok {
				_, err = w.Write([]byte(value))
				if err != nil {
					fmt.Println(err)
				}
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		})
	}

	// Creating Webserver for static files
	fileServer = http.FileServer(http.Dir(directory))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if verbose {
			log.Println("INF Serving ", r.RequestURI)
		}
		fileServer.ServeHTTP(w, r)
	})

	// Starting Webserver
	fmt.Println("Listening on " + ipaddress + ":" + port + " ...")
	err = http.ListenAndServe(ipaddress+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
