package sslapi

/*
import (
	"crypto/x509"
	"dbaction"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/ericchiang/letsencrypt"
)

type ACMEV1 struct {
	client *letsencrypt.Client
	//account *AcmeAccount
}

func (l *ACMEV1) Init() error {
	cli, err := letsencrypt.NewClient("https://acme-v01.api.letsencrypt.org/directory")
	if err != nil {
		return err
	}
	l.client = cli
	//l.account, err = l.InitAccount()
	return nil
}

//获取挑战信息
func (l *ACMEV1) GetChallenges(domain string, dcv_method string, sslaccount *AcmeAccount) (*ChallengeInfo, error) {
	if domain == "" {
		return nil, fmt.Errorf("domain is empty")
	}

	auth, _, err := l.client.NewAuthorization(sslaccount.Key, "dns", domain)
	if err != nil {
		return nil, err
	}
	chals := auth.Challenges
	if len(chals) == 0 {
		return nil, fmt.Errorf("no supported challenge combinations")
	}
	http_flag := false
	var filepath, file_content string
	var challenge_by []byte
	for _, val := range chals {
		switch dcv_method {
		case "DNS_CSR_HASH":
			filepath, file_content, err = val.DNS(sslaccount.Key)
		case "HTTP_CSR_HASH":
			filepath, file_content, err = val.HTTP(sslaccount.Key)
		default:
			return nil, fmt.Errorf("challenge %s is not supprt", dcv_method)
		}
		if err != nil {
			continue
		}
		if filepath == "" || file_content == "" {
			return nil, fmt.Errorf("get challenge info failed")
		}
		filepath = strings.TrimLeft(filepath, "/")
		http_flag = true
		//保留挑战信息，下次向acme发送挑战准备.
		challenge_by, err = json.Marshal(val)
		if err != nil {
			return nil, err

		}
		break
	}
	if !http_flag {

		return nil, fmt.Errorf("can not find http challenge")
	}
	return &ChallengeInfo{File_content: file_content, File_path: filepath, Lets_challenge: string(challenge_by), Order: ""}, nil
}

//挑战准备好，发送服务器
func (l *ACMEV1) IsReady(ssl_info *dbaction.SslParam, sslaccount *AcmeAccount) (err error) {
	var chal letsencrypt.Challenge
	err = json.Unmarshal([]byte(ssl_info.Challenge), &chal)
	if err != nil {
		return
	}
	return l.client.ChallengeReady(sslaccount.Key, chal)
}

//获取证书
type CsrKey struct {
	Certificate string
	Private     string
}

func (l *ACMEV1) GetCer(ssl_info *dbaction.SslParam, sslaccount *AcmeAccount) (err error) {

	// load the account's private key and the certificate request
	//TODO 获取私钥
	csr_re, priv_key, err := genarateCsr(ssl_info.Domain, GetPasswd())
	if err != nil {
		return
	}

	// ask for a new certifiate
	cert, err := l.client.NewCertificate(sslaccount.Key, csr_re)
	if err != nil {
		return
	}
	if !cert.IsAvailable() {
		return fmt.Errorf("Expected certificate to be available in CertificateResponse,after %s", cert.RetryAfter)
	}
	if cert.Issuer == "" {
		return fmt.Errorf("Expected issuer to be non empty.")
	}
	//fmt.Printf("get cert %v\n", cert)
	pemBundle, err := l.client.Bundle(cert)
	if err != nil {
		return err
	}
	ssl_info.Certificate = string(pemBundle)
	ssl_info.Private_key = priv_key
	return nil
}
func loadCSR(csr string) (*x509.CertificateRequest, error) {
	block, _ := pem.Decode([]byte(csr))
	if block == nil {
		return nil, fmt.Errorf("pem decode: no key found")
	}
	return x509.ParseCertificateRequest(block.Bytes)
}

//向letsencrypt 创建的private，并注册
func (l *ACMEV1) InitAccount(user string, account_num int) (map[int]*AcmeAccount, error) {
	accountMap := make(map[int]*AcmeAccount)
	for i := 0; i < account_num; i++ {
		//首先从本地载入
		account, err := loadAccount(user, i)
		if err == nil {
			//return account, nil
			accountMap[i] = account
			continue
		}
		//重新创建
		acmeAccount := newAccount(user, i)
		privatekey, err := acmeAccount.createPrivateFile()
		if err != nil {
			return nil, err
		}
		_, err = l.client.NewRegistration(privatekey)
		if err != nil {
			return nil, err
		}
		acmeAccount.Key = privatekey
		err = acmeAccount.Save()
		if err != nil {
			return nil, err
		}
		accountMap[i] = account
	}
	return accountMap, nil
}
*/
