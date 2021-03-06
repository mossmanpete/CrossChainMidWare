package config



type ServerConfig struct{
	RpcListenEndPoint map[string]string
	RethinkDbEndPoint map[string]string
	RethinkDbName map[string]string
	AddressTrxDbPath map[string]string
	GetInfoFunctionName map[string]string  // such as getblockchaininfo getinfo choose which can use
	IsTls map[string]bool   // is  https
	IsOldFunctionLevel map[string]bool //is getblock can use second param 2
	SupportCoinType map[string]string
	SourceDataHost map[string]string
	SourceDataPort map[string]string
	SourceDataUserName map[string]string
	SourceDataPassword map[string]string
	DbPathConfig map[string]string
	SafeBlock map[string]int
	MULTISIGVERSION map[string]int
}
var RpcServerConfig = ServerConfig{RpcListenEndPoint:map[string]string{
	"BTC":"0.0.0.0:5444",
	"BTC_TEST":"0.0.0.0:5446",
	"LTC":"0.0.0.0:5445",
	"LTC_TEST":"0.0.0.0:5448",
	"HC":"0.0.0.0:5447",
	"HC_TEST":"0.0.0.0:5449"},
	GetInfoFunctionName:map[string]string{
		"BTC":"getblockchaininfo",
		"BTC_TEST":"getblockchaininfo",
		"LTC":"getblockchaininfo",
		"LTC_TEST":"getblockchaininfo",
		"HC":"getinfo",
		"HC_TEST":"getinfo"},
	IsTls:map[string]bool{
		"BTC":false,
		"BTC_TEST":false,
		"LTC":false,
		"LTC_TEST":false,
		"HC":false,
		"HC_TEST":false},
	IsOldFunctionLevel:map[string]bool{
		"BTC":false,
		"BTC_TEST":false,
		"LTC":false,
		"LTC_TEST":false,
		"HC":true,
		"HC_TEST":true},
	SourceDataHost:map[string]string{
		"BTC":"http://btc_wallet",
		"BTC_TEST":"http://btc_wallet",
		"LTC":"http://ltc_wallet",
		"LTC_TEST":"http://ltc_wallet",
		"HC":"http://hc_wallet",
		"HC_TEST":"http://hc_wallet"},
	SourceDataPort:map[string]string{
		"BTC":"60011",
		"BTC_TEST":"40011",
		"LTC":"60012",
		"LTC_TEST":"40012",
		"HC":"19021",
		"HC_TEST":"39021"},
	SourceDataUserName:map[string]string{
		"BTC":"a",
		"BTC_TEST":"a",
		"LTC":"a",
		"LTC_TEST":"a",
		"HC":"a",
		"HC_TEST":"a"},
	SourceDataPassword:map[string]string{
		"BTC":"b",
		"BTC_TEST":"b",
		"LTC":"b",
		"LTC_TEST":"b",
		"HC":"b",
		"HC_TEST":"b"},
	DbPathConfig:map[string]string{
		"BTC":"/hx/btc_collect_data/",
		"BTC_TEST":"/hx/btc_collect_test_data/",
		"LTC":"/hx/ltc_collect_data/",
		"LTC_TEST":"/hx/ltc_collect_test_data/",
		"HC":"/hx/hc_collect_data/",
		"HC_TEST":"/hx/hc_collect_test_data/"},
	SupportCoinType:map[string]string{
		"BTC":"",
		"BTC_TEST":"",
		"LTC":"",
		"LTC_TEST":"",
		"HC":"",
		"HC_TEST":""},
	SafeBlock:map[string]int{
		"BTC":6,
		"BTC_TEST":2,
		"LTC":6,
		"LTC_TEST":2,
		"HC":6,
		"HC_TEST":6},
	MULTISIGVERSION:map[string]int{
		"BTC":5,
		"BTC_TEST":196,
		"LTC":5,
		"LTC_TEST":196,
		"HC":5,
		"HC_TEST":5}}
