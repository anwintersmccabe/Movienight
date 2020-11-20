package main

import (
    "bufio"
    "io"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	convertMedia()

	const vidDir = "videos"
	const port = 8000

	http.Handle("/", addHeaders(http.FileServer(http.Dir(vidDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", vidDir, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func convertMedia() {
	cmd := exec.Command("python3","mp4_to_hls.py")
	stdout, err := cmd.StdoutPipe(); check(err);
    stderr, err := cmd.StderrPipe(); check(err);
    err = cmd.Start(); check(err);

    go copyOutput(stdout)
    go copyOutput(stderr)
    cmd.Wait()
}

func copyOutput(r io.Reader) {
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}

func check(e error) { if e != nil { panic(e) } }