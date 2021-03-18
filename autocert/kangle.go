package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

const (
	FORWARD_AUTOCERT_XML = "ext/forward_autocert.xml"
)

func prepareKangle() error {
	str := "<!--#start 999 -->\n<config>\n" +
		"\t<server name='autocert' proto='http' host='127.0.0.1' port='19981' life_time='5' />\n" +
		"\t\t<request>\n" +
		"\t\t\t<table name='BEGIN'>\n" +
		"\t\t\t\t<chain  action='server:autocert' >\n" +
		"\t\t\t\t\t<acl_path  path='/.well-known/acme-challenge/*'></acl_path>\n" +
		"\t\t\t\t\t<mark_flag  x_real_ip='1' ></mark_flag>\n" +
		"\t\t\t\t</chain>\n" +
		"\t\t</table>\n" +
		"\t</request>\n" +
		"</config>"
	//try to mkdir ext
	os.Mkdir(kangle_base_dir+"ext", 0755)
	err := ioutil.WriteFile(kangle_base_dir+FORWARD_AUTOCERT_XML, []byte(str), 0644)
	if err != nil {
		return err
	}
	return reloadKangle()
}

func reloadKangle() error {
	kangle := kangle_base_dir + "bin/kangle"
	if runtime.GOOS == "windows" {
		kangle += ".exe"
	}
	return exec.Command(kangle, "-r").Run()
}
func cleanKangle() error {
	err := os.Remove(kangle_base_dir + FORWARD_AUTOCERT_XML)
	if err != nil {
		return err
	}
	return reloadKangle()
}
