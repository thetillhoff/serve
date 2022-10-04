package serve

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
)

func Serve(settings ...Setting) error {
	var (
		err        error
		fileServer http.Handler

		// Setting defaults
		directory string = "./"
		verbose   bool   = true
		ipaddress        = "0.0.0.0"
		port             = "3000"
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
		}
	}

	if verbose {
		log.Println("INF verbose=true")
		log.Println("INF directory=" + directory)
		log.Println("INF ipaddress=" + ipaddress)
		log.Println("INF port=" + port)
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
