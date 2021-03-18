package sslapi

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
)

/*
密钥加密密码
*/
func GetPasswd() string {
	return privEncryptPassword
	//return "123456"
}

/**
生成证书请求，csr文件　、私钥
commonName 域名
*/
func genarateCsr(commonName string, passwd string) (csr *x509.CertificateRequest, priv_key string, err error) {
	req := &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: commonName,
			/*Country:      []string{"CN"},
			Province:     []string{"江西省"},
			Organization: []string{"南昌帮腾科技有限公司"},
			Locality:     []string{"南昌市"},*/
		},
	}
	priv, err := rsa.GenerateKey(rand.Reader, 2048) //私钥
	if err != nil {
		return nil, "", err
	}

	csr_byte, err := x509.CreateCertificateRequest(rand.Reader, req, priv) //csr请求
	if err != nil {
		return nil, "", err
	}
	csr, err = x509.ParseCertificateRequest(csr_byte)
	if err != nil {
		return nil, "", err
	}
	priv_key, err = encodePriWithPasswd(priv, passwd)
	return
}
func pemEncode(data interface{}) string {
	var pemBlock *pem.Block
	switch key := data.(type) {
	case []byte:
		pemBlock = &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: key}
	case *rsa.PrivateKey:
		pemBlock = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	}

	return string(pem.EncodeToMemory(pemBlock))
}
func encodePriWithPasswd(key *rsa.PrivateKey, passwd string) (string, error) {
	data := x509.MarshalPKCS1PrivateKey(key)
	block, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", data, []byte(passwd), x509.PEMCipherAES128)
	if err != nil {
		return "", err
	}
	return string(pem.EncodeToMemory(block)), nil
}

//DecodePriv 解析私钥文件生成私钥，（RSA，和ECDSA两种私钥格式）
func PemDecode(str, passwd string) (string, error) {
	block, _ := pem.Decode([]byte(str))
	if block == nil {
		return "", fmt.Errorf("pem decode err")
	}
	buf := block.Bytes
	if passwd == "" {
		return string(buf), nil
	}
	buf, err := x509.DecryptPEMBlock(block, []byte(passwd))
	if err != nil {
		return "", err
	}
	priv, err := x509.ParsePKCS1PrivateKey(buf) //解析成RSA私钥
	if err != nil {
		return "", fmt.Errorf("解析成RSA私钥失败[%s]", err.Error())
	}

	return pemEncode(priv), nil
}
func DecryptPrivKey(priv string) (string, error) {
	return PemDecode(priv, GetPasswd())
}
