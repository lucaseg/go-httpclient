package gohttp

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
)

var (
	MockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enable      bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	MockupServer.enable = true
}

func StopMockServer() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	MockupServer.enable = false
}

func FlushMocks() {
	MockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	key := MockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	MockupServer.mocks[key] = &mock
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

func (m *mockServer) getMock(method, url, responseBody string) *Mock {
	if !m.enable {
		return nil
	}
	key := MockupServer.getMockKey(method, url, responseBody)
	mock := m.mocks[key]
	if mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New("no mock match"),
	}
}
