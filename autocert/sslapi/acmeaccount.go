package sslapi

import (
	"crypto/rand"
	"crypto/rsa"

	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type AcmeAccount struct {
	Key         *rsa.PrivateKey `json:"key"`
	URL         string          `json:"url"`
	Thumbprint  string          `json:"thumbprint"`
	user        string
	account_num int
}

func isFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// NewAccount creates a new account for an email address
func loadAccount(user string, account_num int) (*AcmeAccount, error) {
	acc := newAccount(user, account_num)
	return acc.LoadAccountInfo()
}
func newAccount(user string, account_num int) *AcmeAccount {
	return &AcmeAccount{user: user, account_num: account_num}
}
func (a *AcmeAccount) getUserAccountFile() string {
	return fmt.Sprintf("%s%s%s_%d_acme_account.json", filepath.Dir(os.Args[0]), accountDir, a.user, a.account_num)
}
func (a *AcmeAccount) getCommonAccountFile() string {
	return fmt.Sprintf("%s%sacme_account.json", filepath.Dir(os.Args[0]), accountDir)
}
func (a *AcmeAccount) LoadAccountInfo() (*AcmeAccount, error) {
	path := a.getUserAccountFile()

	flag := isFileExist(path)
	if !flag {
		fmt.Printf("\t\tfile:%s not exist\n", path)
		return nil, fmt.Errorf("not exist")
	}
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("file is empty")
	}
	err = json.Unmarshal(fileBytes, a)
	return a, err
}

//创建私钥
func (a *AcmeAccount) createPrivateFile() (*rsa.PrivateKey, error) {
	accountKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return accountKey, nil
}
func (a *AcmeAccount) Save() error {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(
		a.getUserAccountFile(),
		jsonBytes,
		0600,
	)
}

/*

func newAccount(user string) (*AcmeAccount, error) {
	acc := new(AcmeAccount)
	accKeyPath := acc.getPrivateFile(user)
	privKey := new(rsa.PrivateKey)
	if flag := isFileExist(accKeyPath); flag {
		key, err := acc.loadPriKey(user)
		if err != nil {
			return nil, err
		}
		privKey = key
	} else {
		key, err := acc.CreatePrivateFile(user)
		if err != nil {
			return nil, err
		}
		privKey = key
	}
	accountFile := acc.getAccountFile(user)
	if flag := isFileExist(accountFile); !flag {
		return &Lets_Account{User: user, Email: email, key: privKey}, nil
	}

	fileBytes, err := ioutil.ReadFile(accountFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileBytes, acc)
	if err != nil {
		return nil, err
	}

	acc.key = privKey
	return acc, nil
}
func (a *AcmeAccount) loadPriKey(user string) (key *rsa.PrivateKey, err error) {
	path := a.getPrivateFile(user)
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	keyBlock, _ := pem.Decode(keyBytes)
	return x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
}

// Create a private key for your account and register
func (a *AcmeAccount) CreatePrivateFile(user string) (*rsa.PrivateKey, error) {
	accountKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(accountKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	path := a.getPrivateFile(user)
	fmt.Printf("create lets account[%s] private file\n", path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return nil, err
	}
	return accountKey, nil
}

func (a *AcmeAccount) getPrivateFile(user string) string {
	return fmt.Sprintf("%s%s%s_private.pem", filepath.Dir(os.Args[0]), LETS_ACCOUNT_DIR, user)
}
func (a *AcmeAccount) getAccountFile(user string) string {
	return fmt.Sprintf("%s%s%s_account.json", filepath.Dir(os.Args[0]), LETS_ACCOUNT_DIR, user)
}*/

// Save the account to disk
