package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	//Part1
	http.HandleFunc("/readHeader", readHeader)
	//Part2
	http.HandleFunc("/readOSVer", readOSVer)
	//Part3
	http.HandleFunc("/serverLog", serverLog)
	//Part4
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func serverLog(w http.ResponseWriter, r *http.Request) {
	host, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("clientIp ip: %s, port is: %s return Code: %d", host, port, http.StatusOK)
	w.Header().Add("clientIp", host)
	w.WriteHeader(http.StatusOK)
}

func readOSVer(w http.ResponseWriter, r *http.Request) {
	ver := os.Getenv("VERSION")
	if ver == "" {
		ver = "unknown"
	}
	w.Header().Add("OSVersion", ver)
	w.WriteHeader(http.StatusOK)
}

func readHeader(w http.ResponseWriter, r *http.Request) {
	value := r.Header.Get("test")
	w.Header().Add("test", value)
	w.WriteHeader(http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
