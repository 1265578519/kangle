package main

import (
	"autocert/sslapi"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type Challenge struct {
	domain string
	ci     *sslapi.ChallengeInfo
	ch     chan bool
}

var challengs map[string]*Challenge
var lock sync.Mutex

func index(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	challenge, ok := challengs[r.URL.Path]
	lock.Unlock()
	if !ok {
		fmt.Printf("%s %s%s 404\n", r.Method, r.URL.Host, r.URL.Path)
		str := "404 page not found"
		w.Write([]byte(str))
		return
	}
	remote_addr := r.RemoteAddr
	host, _, _ := net.SplitHostPort(remote_addr)
	if strings.HasPrefix(host, "127.0.0.") {
		remote_addr = r.Header.Get("X-Real-Ip")
	} else {
		remote_addr = host
	}
	fmt.Printf("challenge domain [%s] from ip=[%s]\n", challenge.domain, remote_addr)
	w.Write([]byte(challenge.ci.File_content))
	challenge.ch <- true
	return
}
func removeChallenge(challenge *Challenge) {
	lock.Lock()
	delete(challengs, "/"+challenge.ci.File_path)
	defer lock.Unlock()
	close(challenge.ch)
}
func addChallenge(domain string, ci *sslapi.ChallengeInfo) *Challenge {
	challenge := &Challenge{domain: domain, ci: ci, ch: make(chan bool, 16)}
	lock.Lock()
	defer lock.Unlock()
	challengs["/"+ci.File_path] = challenge
	return challenge
}
func waitChallengeTimeout(challenge *Challenge, tmo int) bool {
	fmt.Printf("wait domain [%s] challenge...\n", challenge.domain)
	select {
	case <-challenge.ch:
		return true
	case <-time.After(time.Second * time.Duration(tmo)):
		return false
	}
}
func startAddress(addr string) error {
	srv := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		err = srv.Serve(ln)
		if err != nil {
			os.Exit(1)
		}
	}()
	return nil
}
func startServer() error {
	challengs = make(map[string]*Challenge, 0)
	http.HandleFunc("/", index)
	err := startAddress(":80")
	if err != nil {
		fmt.Printf("try to listen :80 failed.\n")
	}
	return startAddress(LISTEN_ADDR)
}
