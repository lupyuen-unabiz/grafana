// +build go1.6
// +build !darwin

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"runtime"
)

func init() {
	//  Listen for HTTP requests.
	http.HandleFunc("/_ah/start", startHandler)
	http.HandleFunc("/_ah/stop", stopHandler)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a new App Engine context from the request.
	ctx := appengine.NewContext(r)
	Infof(ctx, "handler: %v, header: %v", r.URL.Path, r.Header)
	// if r.URL.Path == "/_ah/remote_api/" { return }  //  Let remote api handle this.
	// if r.URL.Path == "/_ah/push-handlers/" { handlePush(ctx, w, r); return }  //  PubSub push.
	// if r.URL.Path == "/fetch" { fetchRecords(ctx, w, r); return }  //  Fetch records.
	// if r.URL.Path == "/update" { updateRecords(ctx, w, r); return; }  //  Update records.
	http.NotFound(w, r)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	//  /_ah/start is called when AppEngine starts up.
	ctx := appengine.NewContext(r)
	Infof(ctx, "startHandler: %v", r.URL.Path)
	w.WriteHeader(204)
	result := "OK"
	fmt.Fprint(w, "Completed startHandler: %v", result)
	Infof(ctx, "Completed startHandler: %v, %v", r.URL.Path, result)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	//  /_ah/stop is called when AppEngine shuts down.
	ctx := appengine.NewContext(r)
	// ctx := appengine.NewContext(r)
	Infof(ctx, "stopHandler: %v", r.URL.Path)
	w.WriteHeader(204)
	fmt.Fprint(w, "Completed stopHandler")
	Infof(ctx, "Completed stopHandler: %v", r.URL.Path)
}


func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "ok")
}

/* func Debugf(ctx context.Context, format string, args ...interface{}) {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { fmt.Printf(format, args...); fmt.Printf("\n"); return }
	log.Debugf(ctx, format, args...) } */

func Infof(ctx context.Context, format string, args ...interface{}) {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { fmt.Printf(format, args...); fmt.Printf("\n"); return }
	log.Infof(ctx, format, args...) }

func Warningf(ctx context.Context, format string, args ...interface{}) {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { fmt.Printf(format, args...); fmt.Printf("\n"); return }
	log.Warningf(ctx, format, args...) }

func Errorf(ctx context.Context, format string, args ...interface{}) {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { fmt.Printf(format, args...); fmt.Printf("\n"); return }
	log.Errorf(ctx, format, args...) }
