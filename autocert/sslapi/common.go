package sslapi

//"fmt"
//	"common/ssl"
//"config"
//"dbaction"

const (
	//PLATFORM_COMODO = "comodo" //comodo platform

	PLATFORM_ACMEV1     = "acmev1"
	PLATFORM_ACMEV2     = "acmev2"
	DEFAULT_ACCOUNT_DIR = "/../acme_account/" //账户目录
)

var sslProduct SslProduct
var accountDir string
var privEncryptPassword string

type ChallengeInfo struct {
	File_path      string
	File_content   string
	Lets_challenge string
	Order          string
}

const (
	DNS_CSR_HASH  = "DNS_CSR_HASH"
	HTTP_CSR_HASH = "HTTP_CSR_HASH"
)

type SslProduct interface {
	//第一步 获取挑战信息
	GetChallenges(domain string, dcv_method string, sslaccount *AcmeAccount) (*ChallengeInfo, error)
	//第二部 告知准备好
	IsReady(challenge string, acount *AcmeAccount) error
	//获取证书
	GetCert(ssl_info *SslParam, acount *AcmeAccount) error

	//多个账户信息
	InitAccount(user string, account_num int) (map[int]*AcmeAccount, error)
	/**
	初始化
	**/
	Init() error
}

func checkMethod(method string) bool {
	switch method {
	case "DNS_CSR_HASH", "HTTP_CSR_HASH":
		return true
	default:
		return false
	}
}
func InitSsl(account_dir string, priv_passwd string) error {
	privEncryptPassword = priv_passwd
	sslProduct = new(ACMEV2)
	if len(account_dir) == 0 {
		account_dir = DEFAULT_ACCOUNT_DIR
	}
	accountDir = account_dir
	return sslProduct.Init()
}
func GetSslPlatform() SslProduct {
	return sslProduct
}
