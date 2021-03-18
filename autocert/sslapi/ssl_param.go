package sslapi

type SslParam struct {
	Id          int64
	Domain      string
	User        string
	Product     int
	Days        int
	Csr         string
	Private_key string
	Certificate string
	Status      int
	Dcv_method  string
	Create_time string
	Expire_time string
	Ordernum    string
	Mem         string
	Challenge   string
	Order       string
	SslAccount  int
	CertUrl     string
}
