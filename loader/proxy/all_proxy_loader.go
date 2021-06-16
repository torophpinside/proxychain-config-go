package loader_proxy

import (
	config_connection "proxychain-config-go/config/connection"
)

type allProxyLoader struct {
	conn config_connection.ConnectionInterface
	torConn config_connection.ConnectionInterface
}

func NewAllProxyLoader(c config_connection.ConnectionInterface, t config_connection.ConnectionInterface) ProxyLoaderInterface {
	return &allProxyLoader{
		conn: c,
		torConn: t,
	}
}

func (alp *allProxyLoader) Load() ([]string, error) {
	var ull []string
	pl := NewAltProxyLoader(alp.torConn)
	ul, err := pl.Load()
	if err != nil {
		return nil, err
	}
	ull = append(ull, ul...)

	pl = NewNovaProxyLoader(alp.torConn)
	ul, err = pl.Load()
	if err != nil {
		return nil, err
	}
	ull = append(ull, ul...)

	//pl = NewProxyListPlusProxyLoader(alp.conn)
	//ul, err = pl.Load()
	//if err != nil {
	//	return nil, err
	//}
	//ull = append(ull, ul...)

	pl = NewFreeProxyListProxyLoader(alp.conn)
	ul, err = pl.Load()
	if err != nil {
		return nil, err
	}
	ull = append(ull, ul...)

	pl = NewSpysoneProxyLoader(alp.conn)
	ul, err = pl.Load()
	if err != nil {
		return nil, err
	}
	ull = append(ull, ul...)

	return ull, nil
}
