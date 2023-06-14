package function

import (
	"context"
	"errors"
	"function/backendfakes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/version"
)

var errTest = errors.New("test")

func TestRun(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name    string
		prepare func(mock *backendfakes.FakeImpl) (method string)
		assert  func(*http.Response, string)
	}{
		{
			name: "success",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.ListNodesReturns(&v1.NodeList{Items: []v1.Node{{}, {}}}, nil)
				mock.ServerVersionReturns(&version.Info{GitVersion: "1.2.3"}, nil)
				mock.MarshalReturns([]byte("test"), nil)
				mock.WriteCalls(func(w http.ResponseWriter, b []byte) (int, error) {
					return w.Write(b)
				})

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusOK, res.StatusCode)
				assert.Contains(t, body, "test")
			},
		},
		{
			name: "failure write response",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.ListNodesReturns(&v1.NodeList{Items: []v1.Node{{}}}, nil)
				mock.ServerVersionReturns(&version.Info{}, nil)
				mock.WriteReturns(0, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure marshal data",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.ListNodesReturns(&v1.NodeList{Items: []v1.Node{{}}}, nil)
				mock.ServerVersionReturns(&version.Info{}, nil)
				mock.MarshalReturns(nil, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure get server version",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.ListNodesReturns(&v1.NodeList{Items: []v1.Node{{}}}, nil)
				mock.ServerVersionReturns(nil, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure node list empty",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.ListNodesReturns(&v1.NodeList{Items: []v1.Node{}}, nil)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, "node list is empty")
			},
		},
		{
			name: "failure list nodes",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.ListNodesReturns(nil, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure create discovery client",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.NewDiscoveryClientForConfigReturns(nil, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure create client",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.NewForConfigReturns(nil, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure get in cluster config",
			prepare: func(mock *backendfakes.FakeImpl) string {
				mock.InClusterConfigReturns(nil, errTest)

				return http.MethodGet
			},
			assert: func(res *http.Response, body string) {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
				assert.Contains(t, body, errTest.Error())
			},
		},
		{
			name: "failure wrong method",
			prepare: func(*backendfakes.FakeImpl) string {
				return http.MethodPost
			},
			assert: func(res *http.Response, _ string) {
				assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
			},
		},
	} {
		testPrepare := tc.prepare
		testAssert := tc.assert

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mock := &backendfakes.FakeImpl{}
			method := testPrepare(mock)
			req := httptest.NewRequest(method, "/", http.NoBody)

			sut := NewHandler()
			sut.impl = mock

			res := httptest.NewRecorder()
			sut.Handle(context.Background(), res, req)

			resp := res.Result()
			body, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			testAssert(resp, string(body))
			resp.Body.Close()
		})
	}
}
