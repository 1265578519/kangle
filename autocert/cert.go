package main

import (
	"autocert/sslapi"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var sslapiAccount *sslapi.AcmeAccount

func loadCertAccount() error {
	fmt.Printf("load cert account\n")

	os.Mkdir(cert_base_dir, 0755)

	err := sslapi.InitSsl("/../etc/", "kangle")
	if err != nil {
		return err
	}
	account, err := sslapi.GetSslPlatform().InitAccount("kangle", 1)
	if err != nil {
		return err
	}
	var ok bool
	sslapiAccount, ok = account[0]
	if !ok {
		return errors.New("cann't find ssl account")
	}
	return nil
}
func saveCert(ssl_info *sslapi.SslParam) error {
	priv_key, err := sslapi.DecryptPrivKey(ssl_info.Private_key)
	if err != nil {
		return err
	}
	cert_file := cert_base_dir + ssl_info.Domain + ".crt"
	priv_file := cert_base_dir + ssl_info.Domain + ".key"
	err = ioutil.WriteFile(cert_file, []byte(ssl_info.Certificate), 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(priv_file, []byte(priv_key), 0644)
	if err == nil {
		addCertIndex(ssl_info.Domain, ssl_info.Expire_time)
	}
	return err
}
func deleteCert(domain string) bool {
	cert_file := cert_base_dir + domain + ".crt"
	priv_file := cert_base_dir + domain + ".key"
	os.Remove(cert_file)
	os.Remove(priv_file)
	return deleteCertIndex(domain)
}
func autoCertDomain(domain string) error {
	ci, err := sslapi.GetSslPlatform().GetChallenges(domain, sslapi.HTTP_CSR_HASH, sslapiAccount)
	if err != nil {
		return err
	}
	//fmt.Println(ci)
	challenge := addChallenge(domain, ci)
	defer removeChallenge(challenge)
	err = sslapi.GetSslPlatform().IsReady(ci.Lets_challenge, sslapiAccount)
	if err != nil {
		fmt.Printf("call ready error\n")
		return err
	}
	result := waitChallengeTimeout(challenge, 30)
	if !result {
		return errors.New("timeout for challenge")
	}
	ssl_info := &sslapi.SslParam{Domain: domain, Order: ci.Order}
	err = sslapi.GetSslPlatform().GetCert(ssl_info, sslapiAccount)
	if err != nil {
		return err
	}
	return saveCert(ssl_info)
}
