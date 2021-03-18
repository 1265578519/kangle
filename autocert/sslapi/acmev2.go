package sslapi

import (
	"time"
	//"dbaction"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"path"

	"github.com/eggsampler/acme"
)

type ACMEV2 struct {
	acmeClient acme.Client
	//account    acme.Account
}

func (this *ACMEV2) Init() error {
	c, err := acme.NewClient("https://acme-v02.api.letsencrypt.org/directory")
	if err != nil {
		return err
	}
	this.acmeClient = c
	//	account, err := this.InitAccount()
	//	if err != nil {
	//		return err
	//	}
	//	this.account = this.getAccount(account)
	return nil
}

func (this *ACMEV2) getAccount(acount *AcmeAccount) acme.Account {
	newAccount := acme.Account{}
	newAccount.PrivateKey = acount.Key
	newAccount.URL = acount.URL
	newAccount.Thumbprint = acount.Thumbprint
	return newAccount
}

//获取挑战信息
func (this *ACMEV2) GetChallenges(domain string, dcv_method string, acount *AcmeAccount) (*ChallengeInfo, error) {
	order, err := this.acmeClient.NewOrderDomains(this.getAccount(acount), domain)
	if err != nil {
		return nil, err
	}
	ssl_chall := new(ChallengeInfo)
	orderByte, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	//order 信息保留 v2获取证书时候需要
	ssl_chall.Order = string(orderByte)

	//获取挑战信息
	for _, authURL := range order.Authorizations {
		auth, er := this.acmeClient.FetchAuthorization(this.getAccount(acount), authURL)
		if err != nil {
			return nil, fmt.Errorf("get challenge fetching authorization Url %q: %s", authURL, er.Error())
		}
		//chal := acme.Challenge{}
		//保留挑战信息，下次向acme发送挑战准备.
		switch dcv_method {
		case "DNS_CSR_HASH":
			chal, ok := auth.ChallengeMap[acme.ChallengeTypeDNS01]
			if !ok {
				return nil, fmt.Errorf("autocert: unable to find dns-01 challenge for auth %s, Url: %s", auth.Identifier.Value, authURL)
			}
			ssl_chall.File_path = "_acme-challenge"
			ssl_chall.File_content = acme.EncodeDNS01KeyAuthorization(chal.KeyAuthorization)
			s, err := json.Marshal(chal)
			if err != nil {
				return nil, err
			}
			ssl_chall.Lets_challenge = string(s)
		case "HTTP_CSR_HASH":
			chal, ok := auth.ChallengeMap[acme.ChallengeTypeHTTP01]
			if !ok {
				return nil, fmt.Errorf("autocert: unable to find http-01 challenge for auth %s, Url: %s", auth.Identifier.Value, authURL)
			}
			ssl_chall.File_path = path.Join(".well-known/acme-challenge", chal.Token)
			ssl_chall.File_content = chal.KeyAuthorization
			s, err := json.Marshal(chal)
			if err != nil {
				return nil, err
			}
			ssl_chall.Lets_challenge = string(s)
		default:
			continue
		}

		return ssl_chall, nil
	}
	return nil, fmt.Errorf("get challenge error")
}

//挑战准备好，发送acme服务器 告知准备好了
func (this *ACMEV2) IsReady(challenge string, acount *AcmeAccount) (err error) {

	var chal acme.Challenge
	err = json.Unmarshal([]byte(challenge), &chal)
	if err != nil {
		return
	}
	_, err = this.acmeClient.UpdateChallenge(this.getAccount(acount), chal)
	return
}

//获取证书
func (this *ACMEV2) GetCert(ssl_info *SslParam, acount *AcmeAccount) (err error) {
	csr, priv_key, err := genarateCsr(ssl_info.Domain, GetPasswd())
	if err != nil {
		err = fmt.Errorf("get certificates generate csr:%s", err.Error())
		return
	}
	order := acme.Order{}
	err = json.Unmarshal([]byte(ssl_info.Order), &order)

	if err != nil {
		return
	}
	order, err = this.acmeClient.FinalizeOrder(this.getAccount(acount), order, csr)

	if err != nil {
		err = fmt.Errorf("get certificates finalize order:%s", err.Error())
		return
	}
	certs, err := this.acmeClient.FetchCertificates(this.getAccount(acount), order.Certificate)
	if err != nil {
		err = fmt.Errorf("get certificates:%s", err.Error())
		return
	}
	var pemData []byte
	for _, c := range certs {
		date := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: c.Raw,
		})
		pemData = append(pemData, date...)
	}
	ssl_info.Certificate = string(pemData)
	ssl_info.Private_key = priv_key
	dur, _ := time.ParseDuration("2160h")
	min := time.Now().Add(dur)
	ssl_info.Expire_time = min.Format("2006-01-02 15:04:05")
	return nil
}

//单个用户 多个sslmaster账户
func (this *ACMEV2) InitAccount(user string, account_num int) (map[int]*AcmeAccount, error) {
	accountMap := make(map[int]*AcmeAccount)
	for i := 0; i < account_num; i++ {
		//首先从本地载入
		//fmt.Printf("index:%d\n", i)
		account, err := loadAccount(user, i)
		if err == nil {
			accountMap[i] = account
			continue
		}
		//重新创建
		acmeAccount := newAccount(user, i)
		privatekey, err := acmeAccount.createPrivateFile()
		if err != nil {
			return nil, err
		}
		fmt.Printf("create account %s:%d ...\n", user, i)
		a, err := this.acmeClient.NewAccount(privatekey, false, true)
		if err != nil {
			return nil, err
		}
		acmeAccount.Key = privatekey
		acmeAccount.Thumbprint = a.Thumbprint
		acmeAccount.URL = a.URL
		err = acmeAccount.Save()
		if err != nil {
			return nil, err
		}
		accountMap[i] = acmeAccount
	}

	return accountMap, nil
}
