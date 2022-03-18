package main

import (
	metrics3 "cloudNativeLearn/homework/module10/2/metrics"
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	metrics3.Register()
	mux := http.NewServeMux()
	//Part1
	mux.HandleFunc("/readHeader", readHeader)
	//Part2
	mux.HandleFunc("/readOSVer", readOSVer)
	//Part3
	mux.HandleFunc("/serverLog", serverLog)
	//Part4
	mux.HandleFunc("/healthz", healthz)

	mux.HandleFunc("/metrics", metrics)

	//metrics prometheus use grafana
	mux.HandleFunc("/metrics2", metrics2)

	server := &http.Server{
		Addr:    ":7000",
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGALRM, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	<-c
	gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(gracefulCtx); nil != err {
		log.Fatalf("server shutdown failed, err: %v\n", err)
	} else {
		log.Printf("gracefully stopped\n")
	}

}

func metrics2(w http.ResponseWriter, r *http.Request) {
	timer := metrics3.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))

}

func reg(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: namespace, Name: "execution_latency_seconds", Help: help, Buckets: prometheus.ExponentialBuckets(0.001, 2, 15)}, []string{"step"})
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

func metrics(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}
