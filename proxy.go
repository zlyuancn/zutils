package zutils

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

var Proxy = proxyCli{}

type proxyCli struct{}

type ISocks5Proxy interface {
	Dial(network, addr string) (c net.Conn, err error)
	DialContext(ctx context.Context, network, address string) (net.Conn, error)

	SetProxy(transport *http.Transport)
	DisableProxy(transport *http.Transport)
}

type Socks5Proxy struct {
	dialFn        func(network, addr string) (c net.Conn, err error)
	dialContextFn func(ctx context.Context, network, address string) (net.Conn, error)
}

func (s *Socks5Proxy) Dial(network, addr string) (c net.Conn, err error) {
	return s.dialFn(network, addr)
}
func (s *Socks5Proxy) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return s.dialContextFn(ctx, network, address)
}

func (s *Socks5Proxy) SetProxy(transport *http.Transport) {
	transport.Dial = s.dialFn
	transport.DialContext = s.dialContextFn
}
func (s *Socks5Proxy) DisableProxy(transport *http.Transport) {
	transport.Dial = nil
	transport.DialContext = nil
	transport.Proxy = nil
}

/*
创建一个socks5代理

	address 代理地址. 支持socks5, socks5h. 示例: socks5://127.0.0.1:1080 socks5://user:pwd@127.0.0.1:1080
*/
func (proxyCli) NewSocks5Proxy(address string) (ISocks5Proxy, error) {
	// 解析地址
	u, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("address无法解析: %v", err)
	}

	scheme := strings.ToLower(u.Scheme)
	switch scheme {
	case "socks5", "socks5h":
		var auth *proxy.Auth
		if u.User != nil {
			auth = &proxy.Auth{User: u.User.Username()}
			pwd, ok := u.User.Password()
			if ok {
				auth.Password = pwd
			}
		}

		dialer, err := proxy.SOCKS5("tcp", u.Host, auth, nil)
		if err != nil {
			return nil, fmt.Errorf("sock5.dialer生成失败: %v", err)
		}
		sp := &Socks5Proxy{}
		sp.dialFn = dialer.Dial

		if d, ok := dialer.(proxy.ContextDialer); ok {
			sp.dialContextFn = d.DialContext
			return sp, nil
		}
		sp.dialContextFn = func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.Dial(network, address)
		}
		return sp, nil
	}
	return nil, fmt.Errorf("address的scheme不支持: %s", scheme)
}

type IHttpProxy interface {
	SetProxy(transport *http.Transport)
	DisableProxy(transport *http.Transport)
}

type HttpProxy struct {
	setProxy func(transport *http.Transport)
}

func (h *HttpProxy) SetProxy(transport *http.Transport) {
	h.setProxy(transport)
}
func (h *HttpProxy) DisableProxy(transport *http.Transport) {
	transport.Dial = nil
	transport.DialContext = nil
	transport.Proxy = nil
}

/*
创建一个http代理

	address 代理地址. 支持 http, https, socks5, socks5h. 示例: https://127.0.0.1:1080 https://user:pwd@127.0.0.1:1080
*/
func (p proxyCli) NewHttpProxy(address string) (IHttpProxy, error) {
	// 解析地址
	u, err := url.Parse(address)
	if err != nil {
		return nil, fmt.Errorf("address无法解析: %v", err)
	}

	scheme := strings.ToLower(u.Scheme)
	switch scheme {
	case "http", "https":
		setProxy := func(transport *http.Transport) {
			transport.Proxy = func(request *http.Request) (*url.URL, error) {
				return u, nil
			}
		}
		return &HttpProxy{setProxy: setProxy}, nil
	case "socks5", "socks5h":
		s5, err := p.NewSocks5Proxy(address)
		if err != nil {
			return nil, err
		}
		return &HttpProxy{setProxy: s5.SetProxy}, nil
	}
	return nil, fmt.Errorf("address的scheme不支持: %s", scheme)
}
