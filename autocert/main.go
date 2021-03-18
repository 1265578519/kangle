package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	LISTEN_ADDR = "127.0.0.1:19981"
	//LISTEN_ADDR = ":80"
)

const (
	VERSION = "1.0.0"
)

var (
	//kangle_path    *string = flag.String("k", "", "kangle path")
	domains        *string = flag.String("a", "", "add domains(domain1,domain2,..)")
	remove_domains *string = flag.String("d", "", "delete domains(domain1,domain2,...)")
	list           *bool   = flag.Bool("l", false, "list domains")
	//renew          *int    = flag.Int("r", -1, "renew domain days")
)

func autoCertDomains(domains []string) {
	err := startServer()
	if err != nil {
		fmt.Printf("cann't listen [%s]\n", LISTEN_ADDR)
		os.Exit(1)
	}
	err = prepareKangle()
	if err != nil {
		fmt.Printf("prepare kangle error [%s]\n", err.Error())
		os.Exit(1)
	}
	time.Sleep(time.Second)
	defer cleanKangle()
	err = loadCertAccount()
	if err != nil {
		fmt.Printf("load account error [%s]\n", err.Error())
		os.Exit(1)
		return
	}
	any_success := false
	for _, domain := range domains {
		fmt.Printf("auto cert for domain = [%s]\n", domain)
		err = autoCertDomain(domain)
		if err != nil {
			fmt.Printf("auto cert domain [%s] error [%s]\n", domain, err.Error())
		} else {
			any_success = true
		}
	}
	if any_success {
		err := saveCertIndex()
		if err != nil {
			fmt.Printf("save cert index failed. error=[%s]\n", err.Error())
			os.Exit(1)
		}
	}
}
func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Auto certificate for kangle web server version: %s\nUsage: %s [OPTIONS]\r\n", VERSION, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	loadCertIndex()
	if *list {
		printCertIndex()
		return
	}
	if len(*remove_domains) > 0 {
		defer reloadKangle()
		domain_array := strings.Split(*remove_domains, ",")
		any_success := false
		for _, domain := range domain_array {
			result := deleteCert(domain)
			if result {
				any_success = true
			}
			fmt.Printf("remove domain [%s] ", domain)
			if result {
				fmt.Printf("success\n")
			} else {
				fmt.Printf("failed\n")
			}
		}
		if any_success {
			err := saveCertIndex()
			if err != nil {
				fmt.Printf("save cert index failed. error=[%s]\n", err.Error())
				os.Exit(1)
			}
		}
		return
	}
	if len(*domains) > 0 {
		add_domains := strings.Split(*domains, ",")
		autoCertDomains(add_domains)
		return
	}
}
