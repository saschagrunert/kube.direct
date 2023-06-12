package function

import (
	"context"
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type defaultImpl struct{}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . impl
type impl interface {
	InClusterConfig() (*rest.Config, error)
	NewForConfig(*rest.Config) (*kubernetes.Clientset, error)
	NewDiscoveryClientForConfig(*rest.Config) (*discovery.DiscoveryClient, error)
	ListNodes(context.Context, *kubernetes.Clientset) (*v1.NodeList, error)
	ServerVersion(*discovery.DiscoveryClient) (*version.Info, error)
	EncodeJSON(*json.Encoder, any) error
}

func (*defaultImpl) InClusterConfig() (*rest.Config, error) {
	return rest.InClusterConfig()
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

func (*defaultImpl) EncodeJSON(e *json.Encoder, a any) error {
	return e.Encode(a)
}
