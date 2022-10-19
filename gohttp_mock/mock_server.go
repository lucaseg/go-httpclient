package gohttp_mock

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enable      bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = false
}

func FlushMocks() {
	mockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, responseBody string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(responseBody)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}

func GetMock(method, url, responseBody string) *Mock {
	if !mockupServer.enable {
		return nil
	}
	key := mockupServer.getMockKey(method, url, responseBody)
	mock := mockupServer.mocks[key]
	if mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New("no mock match"),
	}
}
