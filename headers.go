package requests

import (
	"errors"
	"net/http"

	"github.com/gospider007/gson"
)

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.31"
var AcceptLanguage = "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"

// 请求操作========================================================================= start
func DefaultHeaders() http.Header {
	return http.Header{
		"User-Agent":      []string{UserAgent},
		"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"Accept-Encoding": []string{"gzip, deflate, br"},
		"Accept-Language": []string{AcceptLanguage},
	}
}
func (obj *RequestOption) initHeaders() error {
	if obj.Headers == nil {
		obj.Headers = DefaultHeaders()
		return nil
	}
	switch headers := obj.Headers.(type) {
	case http.Header:
		obj.Headers = headers.Clone()
		return nil
	case *gson.Client:
		if !headers.IsObject() {
			return errors.New("new headers error")
		}
		head := http.Header{}
		for kk, vv := range headers.Map() {
			if vv.IsArray() {
				for _, v := range vv.Array() {
					head.Add(kk, v.String())
				}
			} else {
				head.Add(kk, vv.String())
			}
		}
		obj.Headers = head
		return nil
	default:
		jsonData, err := gson.Decode(headers)
		if err != nil {
			return err
		}
		obj.Headers = jsonData
		return obj.initHeaders()
	}
}
