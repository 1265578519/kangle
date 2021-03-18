package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type CertIndex struct {
	domain      string
	expire_time string
}

const (
	INDEX_PREFIX = "<!--\t"
)

var cert_index_file string
var cert_index map[string]*CertIndex
var cert_base_dir string
var kangle_base_dir string

func sortKey() []string {
	lst := []string{}
	for k, _ := range cert_index {
		lst = append(lst, k)
	}
	sort.Strings(lst)
	return lst
}
func printCertIndex() {
	keys := sortKey()
	fmt.Printf("domain\texpire\n")
	fmt.Printf("--------------------------------------------\n")
	for _, k := range keys {
		ci, _ := cert_index[k]
		fmt.Printf("%s\t%s\n", ci.domain, ci.expire_time)
	}
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("total %d\n", len(keys))
}
func loadCertIndex() {
	kangle_base_dir = filepath.Dir(os.Args[0]) + "/../"
	cert_base_dir = kangle_base_dir + "etc/ssl/"
	cert_index_file = kangle_base_dir + "ext/autocert.xml"
	cert_index = make(map[string]*CertIndex, 0)
	fp, err := os.OpenFile(cert_index_file, os.O_RDONLY, 0)
	if err == nil {
		defer fp.Close()
		br := bufio.NewReader(fp)
		for {
			line, _, err := br.ReadLine()
			if err != nil {
				break
			}

			str := string(line)
			//fmt.Printf("%s\n", str)
			if strings.HasPrefix(str, INDEX_PREFIX) {
				strs := strings.Split(str, "\t")
				if len(strs) != 4 {
					continue
				}
				ci := &CertIndex{domain: strs[1], expire_time: strs[2]}
				cert_index[ci.domain] = ci
			}
		}
	}
}
func saveCertIndex() error {
	tmp_cert_index_file := filepath.Dir(os.Args[0]) + "/autocert.xml.tmp"
	fp, err := os.OpenFile(tmp_cert_index_file, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()
	fp.WriteString("<!--#start 1 -->\n<config>\n")
	keys := sortKey()
	for _, k := range keys {
		ci, _ := cert_index[k]
		_, err = fp.WriteString(INDEX_PREFIX)
		if err != nil {
			return err
		}
		_, err = fp.WriteString(ci.domain +
			"\t" + ci.expire_time + "\t-->\n" +
			"<ssl domain=\"" +
			ci.domain + "\" certificate=\"-etc/ssl/" +
			ci.domain + ".crt\" certificate_key=\"-etc/ssl/" +
			ci.domain + ".key\"/>\n")
		if err != nil {
			return err
		}
	}
	_, err = fp.WriteString("</config>")
	if err != nil {
		return err
	}
	fp.Close()
	fp = nil
	os.Remove(cert_index_file)
	return os.Rename(tmp_cert_index_file, cert_index_file)
}

func addCertIndex(domain string, expire_time string) {
	ci := &CertIndex{domain: domain, expire_time: expire_time}
	cert_index[domain] = ci
}
func deleteCertIndex(domain string) bool {
	_, ok := cert_index[domain]
	if ok {
		delete(cert_index, domain)
	}
	return ok
}
