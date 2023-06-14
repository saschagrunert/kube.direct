package function

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type defaultImpl struct{}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . impl
type impl interface {
	ClusterConfig() (*rest.Config, error)
	NewForConfig(*rest.Config) (*kubernetes.Clientset, error)
	NewDiscoveryClientForConfig(*rest.Config) (*discovery.DiscoveryClient, error)
	ListNodes(context.Context, *kubernetes.Clientset) (*v1.NodeList, error)
	ServerVersion(*discovery.DiscoveryClient) (*version.Info, error)
	Marshal(proto.Message) ([]byte, error)
	Write(http.ResponseWriter, []byte) (int, error)
}

func (*defaultImpl) ClusterConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
}

func (*defaultImpl) NewForConfig(c *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(c)
}

func (*defaultImpl) NewDiscoveryClientForConfig(c *rest.Config) (*discovery.DiscoveryClient, error) {
	return discovery.NewDiscoveryClientForConfig(c)
}

func (*defaultImpl) ListNodes(ctx context.Context, c *kubernetes.Clientset) (*v1.NodeList, error) {
	return c.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
}

func (*defaultImpl) ServerVersion(d *discovery.DiscoveryClient) (*version.Info, error) {
	return d.ServerVersion()
}

func (*defaultImpl) Marshal(m proto.Message) ([]byte, error) {
	return proto.Marshal(m)
}

func (*defaultImpl) Write(w http.ResponseWriter, b []byte) (int, error) {
	return w.Write(b)
}
