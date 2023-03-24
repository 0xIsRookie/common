package http

import (
	"context"
	"net"
	"net/http"
	stURL "net/url"
	"strings"
	"sync"
	"time"
)

// SetDefaultHeaders 设置默认headers
func SetDefaultHeaders() map[string]string {
	return map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
	}
}

type Dial func(string, string) (net.Conn, error)

// SelectIp 选择请求IP
func SelectIp(ip string) Dial {
	if ip == "" {
		return nil
	}
	return func(netw, addr string) (net.Conn, error) {
		lAddr, err := net.ResolveTCPAddr(netw, ip+":0")
		if err != nil {
			return nil, err
		}
		// 被请求的地址
		rAddr, err := net.ResolveTCPAddr(netw, addr)
		if err != nil {
			return nil, err
		}
		conn, err := net.DialTCP(netw, lAddr, rAddr)
		if err != nil {
			return nil, err
		}
		deadline := time.Now().Add(35 * time.Second)
		conn.SetDeadline(deadline)
		return conn, nil
	}
}

var sendIp = ""
var number = 0
var maxNumber = 20
var dnsLookLock sync.Mutex

// request 封装了 HTTP 请求，可以通过指定请求类型、URL、数据、请求头、超时时间和请求 IP 等参数来发送 HTTP 请求，
// 并返回响应结果和错误信息。
//
// 参数：
//   - _type：请求类型，如 GET、POST 等
//   - url：请求 URL 地址
//   - data：请求数据，通常用于 POST 请求，如果不需要数据则传空字符串
//   - headers：请求头信息，可选参数，如果不需要特别指定则传 nil
//   - timeout：请求超时时间，单位为秒
//   - ip：请求的 IP 地址，可选参数，如果不需要指定则传空字符串
//
// 返回值：
//   - *http.Response：HTTP 响应结果
//   - error：请求过程中的错误信息，如果没有错误则为 nil
func request(_type, url, data string, headers map[string]string, timeout int, ip string) (*http.Response, error) {
	domain, _ := stURL.Parse(url)
	if number == 0 {
		dnsLookLock.Lock()
		if number == 0 {
			iprecords, _ := net.LookupIP(domain.Hostname())
			for _, ip := range iprecords {
				sendIp = ip.String()
				break
			}
		}
		number += 1
		dnsLookLock.Unlock()
	}

	number += 1
	if number > maxNumber {
		number = 0
	}
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				lAddr, err := net.ResolveTCPAddr(network, ip+":0")
				if err != nil {
					return nil, err
				}
				// 被请求的地址
				tt := sendIp + ":" + strings.Split(addr, ":")[1]
				rAddr, err := net.ResolveTCPAddr(network, tt)
				//log.Println(addr, tt)
				//rAddr, err := net.ResolveTCPAddr(network, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(network, lAddr, rAddr)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(35 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
	req, _ := http.NewRequest(strings.ToUpper(_type), url, strings.NewReader(data))

	if headers == nil {
		headers = SetDefaultHeaders()
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	return client.Do(req)
}

// Post 发送 HTTP POST 请求
//
// 参数：
//   - url: 请求的目标 URL
//   - data: 请求携带的数据
//   - headers: 请求头部信息
//   - timeout: 请求的超时时间，单位为秒
//
// 返回值：
//   - *http.Response: HTTP 响应体
//   - error: 请求过程中的错误信息，如果没有错误，则为 nil
func Post(url, data string, headers map[string]string, timeout int) (*http.Response, error) {
	return request("POST", url, data, headers, timeout, "")
}

// Get 发送 HTTP Get 请求
//
// 参数：
//   - url: 请求的目标 URL
//   - headers: 请求头部信息
//   - timeout: 请求的超时时间，单位为秒
//
// 返回值：
//   - *http.Response: HTTP 响应体
//   - error: 请求过程中的错误信息，如果没有错误，则为 nil
func Get(url string, headers map[string]string, timeout int) (*http.Response, error) {
	return request("GET", url, "", headers, timeout, "")
}

// GetWithLocalIP 发送一个 GET 请求到指定的 URL，使用指定的 IP 地址
//
// 参数：
//   - url: 目标 URL
//   - headers: HTTP 请求头
//   - timeout: 请求超时时间（秒）
//   - ip: 使用的 IP 地址
//
// 返回值：
//   - *http.Response: HTTP 响应
//   - error: 错误信息，如果没有错误，则为 nil
func GetWithLocalIP(url string, headers map[string]string, timeout int, ip string) (*http.Response, error) {
	return request("GET", url, "", headers, timeout, ip)
}

// PostWithLocalIP 使用指定的本地IP地址向指定URL发送POST请求
//
// 参数:
//   - url: string类型，请求的目标URL地址
//   - data: string类型，POST请求要发送的数据
//   - headers: map[string]string类型，请求中携带的HTTP头信息，如果为nil则使用默认头信息
//   - timeout: int类型，请求的超时时间（秒）
//   - ip: string类型，发送请求使用的本地IP地址
//
// 返回值：
//   - *http.Response: HTTP响应对象指针
//   - error: 错误信息对象，如果没有错误则为nil
func PostWithLocalIP(url, data string, headers map[string]string, timeout int, ip string) (*http.Response, error) {
	return request("POST", url, data, headers, timeout, ip)
}
