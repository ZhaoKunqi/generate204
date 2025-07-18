package main
import (
        "log"
        "net/http"
)
func noContentHandler(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNoContent)
}
func main() {
        http.HandleFunc("/", noContentHandler)
        log.Println("start serve at container's 8080 port\n")
        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatalf("could not listen on port 8080: %v\n", err)
        }
}
