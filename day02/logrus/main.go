package main

import (
	"fmt"
	"html"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")

	logger := setupLogger(environment)

	startHTTPServer(logger)
}

func setupLogger(environment string) *logrus.Logger {
	filename := "splunk.log"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("can't create log file %q, error: %v\n", filename, err)
		os.Exit(1)
	}

	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	logger.SetLevel(logrus.InfoLevel)
	// logger.SetReportCaller(true)

	if environment == "production" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
		// Actually not needed, it's the default one
	}

	return logger
}

func startHTTPServer(logger *logrus.Logger) {
	http.HandleFunc("/", logging(logger, helloHandler))

	server := http.Server{
		Addr: "0.0.0.0:5000",
	}

	logger.Infof("Server up! Go to http://127.0.0.1:5000")

	go spamRequestsToSelf(logger)

	server.ListenAndServe()
}

func logging(logger *logrus.Logger, fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		fn(rw, req)
		logger.Infof("%s %s %s %.3f\n", req.RemoteAddr, req.Method, req.URL, float64(time.Since(start).Nanoseconds())/1e6)
	}
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	path := html.EscapeString(req.URL.Path[1:])
	fmt.Fprintf(rw, "Hello, %q", path)
	// message := fmt.Sprintf("induced error for %q", path)
	// http.Error(rw, message, http.StatusInternalServerError)
}

func spamRequestsToSelf(logger *logrus.Logger) {
	time.Sleep(10 * time.Second)
	logger.Info("Request spammer started")

	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz0123456789-", "")
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	wait := 20 + rand.Intn(130)
	for {
		time.Sleep(time.Duration(wait) * time.Millisecond)

		randomString := sample(alphabet, 16)
		req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:5000/"+randomString, nil)
		go client.Do(req)

		wait = 20 + rand.Intn(130)
	}
}

func sample(alphabet []string, k int) string {
	var sb strings.Builder
	for i := 0; i < k; i++ {
		index := rand.Intn(len(alphabet))
		sb.WriteString(alphabet[index])
	}

	return sb.String()
}
