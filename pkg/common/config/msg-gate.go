package config

type GateConfig struct {
	TcpSvrCfg TcpSvr `toml:"tcp_svr"`
	WsSvrCfg  WsSvr  `toml:"ws_svr"`
}

type TcpSvr struct {
	Port string
}

type WsSvr struct {
	Port       string
	MaxConnNum int `toml:"max_conn_num"`
	MaxMsgLen  int `toml:"max_msg_len"`
	Timeout    int `toml:"timeout"`
}
