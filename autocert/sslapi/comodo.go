package sslapi

/*import (
	"common/myutils"
	"common/ssl"
	"comodo"
	"dbaction"
	"fmt"
	"strings"
	"time"
)

type ComodoSsl struct {
}

func (c ComodoSsl) GetChallenges(domain string, dcv_method string, let_ac ...*Lets_Account) (ssl_chall ChallengeInfo, err error) {
	sha256 := "%s\ncomodoca.com"
	md5 := ".well-known/pki-validation/%s.txt"
	if dcv_method == "CNAME_CSR_HASH" {
		sha256 = "%s"
		md5 = "%s"
	}

	csr, priv_key, err := genarateCsr(domain, GetPasswd())
	if err != nil {
		return
	}
	filepath := fmt.Sprintf(md5, strings.ToUpper(myutils.GetMd5(csr)))
	file_content := fmt.Sprintf(sha256, strings.ToUpper(myutils.GetSha256(csr)))
	ssl_chall = ChallengeInfo{File_content: file_content, File_path: filepath}
	ssl_chall.Csr = csr
	ssl_chall.Private_key = priv_key

	return
}
func (c ComodoSsl) IsReady(ssl_info *dbaction.SslParam, let_ac ...*Lets_Account) (err error) {
	como := comodo.GetComodo()
	by_return, err := como.BuySsl(ssl_info.Product, ssl_info.Days, ssl_info.Csr, ssl_info.Dcv_method)
	if err != nil {
		return
	}
	mem := by_return.CertificateStatus
	status := ssl.SslInProcess
	if by_return.ErrorCode < 0 {
		//购买失败
		mem = by_return.ErrorMessage
		status = ssl.SslFailed
		//return fmt.Errorf("buy error[%s]", by_return.ErrorMessage)
	}
	//购买成功
	err = dbaction.UpdateSslOrder(ssl_info.Id, ssl_info.User, status, by_return.OrderNumber, mem)
	if err != nil {
		return err
	}
	if by_return.ErrorCode < 0 {
		return fmt.Errorf(by_return.ErrorMessage)
	}
	return nil
}
func (c ComodoSsl) GetCer(ssl_id int64, uid string, order string, csr string, let_ac ...*Lets_Account) (err error) {
	if order == "" {
		return fmt.Errorf("order is empty")
	}
	como := comodo.GetComodo()
	cet, err := como.GetSslInfo(order)
	if err != nil {
		return fmt.Errorf("get comodo ssl err[%s]\n", err.Error())
	}
	if cet.ErrorCode == 0 {
		//正在验证
		return fmt.Errorf("check in process,status[%s]\n", cet.CertificateStatus)
	}
	if cet.ErrorCode < 0 {
		//验证失败
		err := dbaction.UpdateSslMem(ssl_id, uid, ssl.SslFailed, cet.ErrorMessage)
		if err != nil {
			fmt.Printf("update mem and status err[%s]\n", err.Error())
		}
		return fmt.Errorf("check failed,mess[%s]", cet.ErrorMessage)
	}
	//验证成功
	tm := time.Unix(int64(myutils.Intval(cet.NotAfter)), 0)
	err = dbaction.UpdateSslCerti(ssl_id, uid, ssl.SslSuccess, cet.Certificate, tm.Format("2006-01-02 03:04:05 PM"), "")
	if err != nil {
		fmt.Printf("update mem and status err[%s]\n", err.Error())
	}
	return nil
}
func (c ComodoSsl) InitAccount(args ...string) (interface{}, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("comodo init error")
	}
	comodo_name := args[0]
	comodo_passwd := args[1]
	if comodo_name == "" {
		return nil, fmt.Errorf("comodo name is empty")
	}
	if comodo_passwd == "" {
		return nil, fmt.Errorf("comodo passwd is empty")
	}
	comodo.InitComodo(comodo_name, comodo_passwd)
	fmt.Println(comodo.Comodo_manage)
	return nil, nil
}
func (c ComodoSsl) RenewSLL(ssl_info *dbaction.SslParam, let_ac ...*Lets_Account) (*ssl.SslInfo, error) {
	return nil, nil
}*/
