package service

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type BaseRequest struct {
	BaseUrl *url.URL
}

func NewBaseRequest(baseurl string) *BaseRequest {
	request := new(BaseRequest)
	request.BaseUrl, _ = url.Parse(baseurl)
	return request
}

func (s *BaseRequest) HttpGet(path string) ([]byte, error) {
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

func (s *BaseRequest) HttpGetWithParams(path string, params url.Values) ([]byte, error) {
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

func (s *BaseRequest) HttpPost(path string, data []byte) ([]byte, error) {
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

func (s *BaseRequest) HttpPostWithParams(path string, params url.Values, data []byte) ([]byte, error) {
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
