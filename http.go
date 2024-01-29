package goutil

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// PostJson 请求
func HttpPost(link string, header map[string]string, json []byte) ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	u, _ := url.Parse(link)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

// Get 请求  link：请求url
func HttpGet(link string, header map[string]string) ([]byte, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	//忽略https的证书
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	u, _ := url.Parse(link)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", resp.StatusCode, resp.Status)
	}
	return io.ReadAll(resp.Body)
}
