// Code generated by counterfeiter. DO NOT EDIT.
package backendfakes

import (
	"context"
	"net/http"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type FakeImpl struct {
	ClusterConfigStub        func() (*rest.Config, error)
	clusterConfigMutex       sync.RWMutex
	clusterConfigArgsForCall []struct {
	}
	clusterConfigReturns struct {
		result1 *rest.Config
		result2 error
	}
	clusterConfigReturnsOnCall map[int]struct {
		result1 *rest.Config
		result2 error
	}
	ListNodesStub        func(context.Context, *kubernetes.Clientset) (*v1.NodeList, error)
	listNodesMutex       sync.RWMutex
	listNodesArgsForCall []struct {
		arg1 context.Context
		arg2 *kubernetes.Clientset
	}
	listNodesReturns struct {
		result1 *v1.NodeList
		result2 error
	}
	listNodesReturnsOnCall map[int]struct {
		result1 *v1.NodeList
		result2 error
	}
	MarshalStub        func(protoreflect.ProtoMessage) ([]byte, error)
	marshalMutex       sync.RWMutex
	marshalArgsForCall []struct {
		arg1 protoreflect.ProtoMessage
	}
	marshalReturns struct {
		result1 []byte
		result2 error
	}
	marshalReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	NewDiscoveryClientForConfigStub        func(*rest.Config) (*discovery.DiscoveryClient, error)
	newDiscoveryClientForConfigMutex       sync.RWMutex
	newDiscoveryClientForConfigArgsForCall []struct {
		arg1 *rest.Config
	}
	newDiscoveryClientForConfigReturns struct {
		result1 *discovery.DiscoveryClient
		result2 error
	}
	newDiscoveryClientForConfigReturnsOnCall map[int]struct {
		result1 *discovery.DiscoveryClient
		result2 error
	}
	NewForConfigStub        func(*rest.Config) (*kubernetes.Clientset, error)
	newForConfigMutex       sync.RWMutex
	newForConfigArgsForCall []struct {
		arg1 *rest.Config
	}
	newForConfigReturns struct {
		result1 *kubernetes.Clientset
		result2 error
	}
	newForConfigReturnsOnCall map[int]struct {
		result1 *kubernetes.Clientset
		result2 error
	}
	ServerVersionStub        func(*discovery.DiscoveryClient) (*version.Info, error)
	serverVersionMutex       sync.RWMutex
	serverVersionArgsForCall []struct {
		arg1 *discovery.DiscoveryClient
	}
	serverVersionReturns struct {
		result1 *version.Info
		result2 error
	}
	serverVersionReturnsOnCall map[int]struct {
		result1 *version.Info
		result2 error
	}
	WriteStub        func(http.ResponseWriter, []byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) ClusterConfig() (*rest.Config, error) {
	fake.clusterConfigMutex.Lock()
	ret, specificReturn := fake.clusterConfigReturnsOnCall[len(fake.clusterConfigArgsForCall)]
	fake.clusterConfigArgsForCall = append(fake.clusterConfigArgsForCall, struct {
	}{})
	stub := fake.ClusterConfigStub
	fakeReturns := fake.clusterConfigReturns
	fake.recordInvocation("ClusterConfig", []interface{}{})
	fake.clusterConfigMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) ClusterConfigCallCount() int {
	fake.clusterConfigMutex.RLock()
	defer fake.clusterConfigMutex.RUnlock()
	return len(fake.clusterConfigArgsForCall)
}

func (fake *FakeImpl) ClusterConfigCalls(stub func() (*rest.Config, error)) {
	fake.clusterConfigMutex.Lock()
	defer fake.clusterConfigMutex.Unlock()
	fake.ClusterConfigStub = stub
}

func (fake *FakeImpl) ClusterConfigReturns(result1 *rest.Config, result2 error) {
	fake.clusterConfigMutex.Lock()
	defer fake.clusterConfigMutex.Unlock()
	fake.ClusterConfigStub = nil
	fake.clusterConfigReturns = struct {
		result1 *rest.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ClusterConfigReturnsOnCall(i int, result1 *rest.Config, result2 error) {
	fake.clusterConfigMutex.Lock()
	defer fake.clusterConfigMutex.Unlock()
	fake.ClusterConfigStub = nil
	if fake.clusterConfigReturnsOnCall == nil {
		fake.clusterConfigReturnsOnCall = make(map[int]struct {
			result1 *rest.Config
			result2 error
		})
	}
	fake.clusterConfigReturnsOnCall[i] = struct {
		result1 *rest.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ListNodes(arg1 context.Context, arg2 *kubernetes.Clientset) (*v1.NodeList, error) {
	fake.listNodesMutex.Lock()
	ret, specificReturn := fake.listNodesReturnsOnCall[len(fake.listNodesArgsForCall)]
	fake.listNodesArgsForCall = append(fake.listNodesArgsForCall, struct {
		arg1 context.Context
		arg2 *kubernetes.Clientset
	}{arg1, arg2})
	stub := fake.ListNodesStub
	fakeReturns := fake.listNodesReturns
	fake.recordInvocation("ListNodes", []interface{}{arg1, arg2})
	fake.listNodesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) ListNodesCallCount() int {
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	return len(fake.listNodesArgsForCall)
}

func (fake *FakeImpl) ListNodesCalls(stub func(context.Context, *kubernetes.Clientset) (*v1.NodeList, error)) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = stub
}

func (fake *FakeImpl) ListNodesArgsForCall(i int) (context.Context, *kubernetes.Clientset) {
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	argsForCall := fake.listNodesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) ListNodesReturns(result1 *v1.NodeList, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	fake.listNodesReturns = struct {
		result1 *v1.NodeList
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ListNodesReturnsOnCall(i int, result1 *v1.NodeList, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	if fake.listNodesReturnsOnCall == nil {
		fake.listNodesReturnsOnCall = make(map[int]struct {
			result1 *v1.NodeList
			result2 error
		})
	}
	fake.listNodesReturnsOnCall[i] = struct {
		result1 *v1.NodeList
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Marshal(arg1 protoreflect.ProtoMessage) ([]byte, error) {
	fake.marshalMutex.Lock()
	ret, specificReturn := fake.marshalReturnsOnCall[len(fake.marshalArgsForCall)]
	fake.marshalArgsForCall = append(fake.marshalArgsForCall, struct {
		arg1 protoreflect.ProtoMessage
	}{arg1})
	stub := fake.MarshalStub
	fakeReturns := fake.marshalReturns
	fake.recordInvocation("Marshal", []interface{}{arg1})
	fake.marshalMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) MarshalCallCount() int {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	return len(fake.marshalArgsForCall)
}

func (fake *FakeImpl) MarshalCalls(stub func(protoreflect.ProtoMessage) ([]byte, error)) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.MarshalStub = stub
}

func (fake *FakeImpl) MarshalArgsForCall(i int) protoreflect.ProtoMessage {
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	argsForCall := fake.marshalArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) MarshalReturns(result1 []byte, result2 error) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.MarshalStub = nil
	fake.marshalReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) MarshalReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.marshalMutex.Lock()
	defer fake.marshalMutex.Unlock()
	fake.MarshalStub = nil
	if fake.marshalReturnsOnCall == nil {
		fake.marshalReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.marshalReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) NewDiscoveryClientForConfig(arg1 *rest.Config) (*discovery.DiscoveryClient, error) {
	fake.newDiscoveryClientForConfigMutex.Lock()
	ret, specificReturn := fake.newDiscoveryClientForConfigReturnsOnCall[len(fake.newDiscoveryClientForConfigArgsForCall)]
	fake.newDiscoveryClientForConfigArgsForCall = append(fake.newDiscoveryClientForConfigArgsForCall, struct {
		arg1 *rest.Config
	}{arg1})
	stub := fake.NewDiscoveryClientForConfigStub
	fakeReturns := fake.newDiscoveryClientForConfigReturns
	fake.recordInvocation("NewDiscoveryClientForConfig", []interface{}{arg1})
	fake.newDiscoveryClientForConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) NewDiscoveryClientForConfigCallCount() int {
	fake.newDiscoveryClientForConfigMutex.RLock()
	defer fake.newDiscoveryClientForConfigMutex.RUnlock()
	return len(fake.newDiscoveryClientForConfigArgsForCall)
}

func (fake *FakeImpl) NewDiscoveryClientForConfigCalls(stub func(*rest.Config) (*discovery.DiscoveryClient, error)) {
	fake.newDiscoveryClientForConfigMutex.Lock()
	defer fake.newDiscoveryClientForConfigMutex.Unlock()
	fake.NewDiscoveryClientForConfigStub = stub
}

func (fake *FakeImpl) NewDiscoveryClientForConfigArgsForCall(i int) *rest.Config {
	fake.newDiscoveryClientForConfigMutex.RLock()
	defer fake.newDiscoveryClientForConfigMutex.RUnlock()
	argsForCall := fake.newDiscoveryClientForConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) NewDiscoveryClientForConfigReturns(result1 *discovery.DiscoveryClient, result2 error) {
	fake.newDiscoveryClientForConfigMutex.Lock()
	defer fake.newDiscoveryClientForConfigMutex.Unlock()
	fake.NewDiscoveryClientForConfigStub = nil
	fake.newDiscoveryClientForConfigReturns = struct {
		result1 *discovery.DiscoveryClient
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) NewDiscoveryClientForConfigReturnsOnCall(i int, result1 *discovery.DiscoveryClient, result2 error) {
	fake.newDiscoveryClientForConfigMutex.Lock()
	defer fake.newDiscoveryClientForConfigMutex.Unlock()
	fake.NewDiscoveryClientForConfigStub = nil
	if fake.newDiscoveryClientForConfigReturnsOnCall == nil {
		fake.newDiscoveryClientForConfigReturnsOnCall = make(map[int]struct {
			result1 *discovery.DiscoveryClient
			result2 error
		})
	}
	fake.newDiscoveryClientForConfigReturnsOnCall[i] = struct {
		result1 *discovery.DiscoveryClient
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) NewForConfig(arg1 *rest.Config) (*kubernetes.Clientset, error) {
	fake.newForConfigMutex.Lock()
	ret, specificReturn := fake.newForConfigReturnsOnCall[len(fake.newForConfigArgsForCall)]
	fake.newForConfigArgsForCall = append(fake.newForConfigArgsForCall, struct {
		arg1 *rest.Config
	}{arg1})
	stub := fake.NewForConfigStub
	fakeReturns := fake.newForConfigReturns
	fake.recordInvocation("NewForConfig", []interface{}{arg1})
	fake.newForConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) NewForConfigCallCount() int {
	fake.newForConfigMutex.RLock()
	defer fake.newForConfigMutex.RUnlock()
	return len(fake.newForConfigArgsForCall)
}

func (fake *FakeImpl) NewForConfigCalls(stub func(*rest.Config) (*kubernetes.Clientset, error)) {
	fake.newForConfigMutex.Lock()
	defer fake.newForConfigMutex.Unlock()
	fake.NewForConfigStub = stub
}

func (fake *FakeImpl) NewForConfigArgsForCall(i int) *rest.Config {
	fake.newForConfigMutex.RLock()
	defer fake.newForConfigMutex.RUnlock()
	argsForCall := fake.newForConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) NewForConfigReturns(result1 *kubernetes.Clientset, result2 error) {
	fake.newForConfigMutex.Lock()
	defer fake.newForConfigMutex.Unlock()
	fake.NewForConfigStub = nil
	fake.newForConfigReturns = struct {
		result1 *kubernetes.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) NewForConfigReturnsOnCall(i int, result1 *kubernetes.Clientset, result2 error) {
	fake.newForConfigMutex.Lock()
	defer fake.newForConfigMutex.Unlock()
	fake.NewForConfigStub = nil
	if fake.newForConfigReturnsOnCall == nil {
		fake.newForConfigReturnsOnCall = make(map[int]struct {
			result1 *kubernetes.Clientset
			result2 error
		})
	}
	fake.newForConfigReturnsOnCall[i] = struct {
		result1 *kubernetes.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ServerVersion(arg1 *discovery.DiscoveryClient) (*version.Info, error) {
	fake.serverVersionMutex.Lock()
	ret, specificReturn := fake.serverVersionReturnsOnCall[len(fake.serverVersionArgsForCall)]
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct {
		arg1 *discovery.DiscoveryClient
	}{arg1})
	stub := fake.ServerVersionStub
	fakeReturns := fake.serverVersionReturns
	fake.recordInvocation("ServerVersion", []interface{}{arg1})
	fake.serverVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) ServerVersionCallCount() int {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}

func (fake *FakeImpl) ServerVersionCalls(stub func(*discovery.DiscoveryClient) (*version.Info, error)) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = stub
}

func (fake *FakeImpl) ServerVersionArgsForCall(i int) *discovery.DiscoveryClient {
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	argsForCall := fake.serverVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) ServerVersionReturns(result1 *version.Info, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1 *version.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ServerVersionReturnsOnCall(i int, result1 *version.Info, result2 error) {
	fake.serverVersionMutex.Lock()
	defer fake.serverVersionMutex.Unlock()
	fake.ServerVersionStub = nil
	if fake.serverVersionReturnsOnCall == nil {
		fake.serverVersionReturnsOnCall = make(map[int]struct {
			result1 *version.Info
			result2 error
		})
	}
	fake.serverVersionReturnsOnCall[i] = struct {
		result1 *version.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Write(arg1 http.ResponseWriter, arg2 []byte) (int, error) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1, arg2Copy})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeImpl) WriteCalls(stub func(http.ResponseWriter, []byte) (int, error)) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeImpl) WriteArgsForCall(i int) (http.ResponseWriter, []byte) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) WriteReturns(result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clusterConfigMutex.RLock()
	defer fake.clusterConfigMutex.RUnlock()
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	fake.marshalMutex.RLock()
	defer fake.marshalMutex.RUnlock()
	fake.newDiscoveryClientForConfigMutex.RLock()
	defer fake.newDiscoveryClientForConfigMutex.RUnlock()
	fake.newForConfigMutex.RLock()
	defer fake.newForConfigMutex.RUnlock()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
