package microservice

import (
	"net/http"
	"os"
	"strings"

	"google.golang.org/appengine"
)

// SetCors ...
// Either server's TLD matches client's TLD,
// or server name is present in env variable CORS_DOMAINS
func SetCors(w http.ResponseWriter, r *http.Request) {
	clientOrigin := r.Header.Get("Origin")
	if "" != clientOrigin {
		clientDomain := TopLevelDomain(clientOrigin)
		serverDomain := TopLevelDomain(r.Header.Get("X-Appengine-Server-Name"))
		corsDomains := strings.Split(os.Getenv("CORS_DOMAINS"), " ")
		if clientDomain == serverDomain ||
			inSlice(clientDomain, corsDomains) ||
			appengine.IsDevAppServer() {

			w.Header().Set("Access-Control-Allow-Origin", clientOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}
	}
}

// TopLevelDomain ...
// Get the TLD from string, works only for Tier-1 TLDs
func TopLevelDomain(n string) string {
	var domain string
	slice := strings.Split(n, ".")
	if len(slice) > 1 {
		last2 := slice[len(slice)-2:]
		domain = strings.Join(last2, ".")
	} else {
		domain = n
	}
	return domain
}

// If string is in slice
func inSlice(s string, l []string) bool {
	for i := range l {
		if l[i] == s {
			return true
		}
	}
	return false
}
