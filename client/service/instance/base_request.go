package instance

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type baseRequest struct {
	BaseUrl *url.URL
}

var BaseRequest = new(baseRequest)

func (s *baseRequest) SetBaseUrl(baseurl string) {
	s.BaseUrl, _ = url.Parse(baseurl)
}

func (s *baseRequest) UnsetBaseUrl() {
	s.BaseUrl = nil
}

func (s *baseRequest) HttpGet(path string) ([]byte, error) {
	if s.BaseUrl == nil {
		return nil, errors.New("服务未连接")
	}
	var result []byte
	response, err := http.Get(s.BaseUrl.JoinPath(path).String())
	if err == nil {
		result, err = io.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return result, err
}

func (s *baseRequest) HttpGetWithParams(path string, params url.Values) ([]byte, error) {
	if s.BaseUrl == nil {
		return nil, errors.New("服务未连接")
	}
	var result []byte
	response, err := http.Get(s.BaseUrl.JoinPath(path).String() + "?" + params.Encode())
	if err == nil {
		result, err = io.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return result, err
}

func (s *baseRequest) HttpPost(path string, data []byte) ([]byte, error) {
	if s.BaseUrl == nil {
		return nil, errors.New("服务未连接")
	}
	var result []byte
	response, err := http.Post(
		s.BaseUrl.JoinPath(path).String(),
		"application/json",
		bytes.NewBuffer(data))
	if err == nil {
		result, err = io.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return result, err
}

func (s *baseRequest) HttpPostWithParams(path string, params url.Values, data []byte) ([]byte, error) {
	if s.BaseUrl == nil {
		return nil, errors.New("服务未连接")
	}
	var result []byte
	response, err := http.Post(
		s.BaseUrl.JoinPath(path).String()+"?"+params.Encode(),
		"application/json",
		bytes.NewBuffer(data))
	if err == nil {
		result, err = io.ReadAll(response.Body)
		defer response.Body.Close()
	}
	return result, err
}
