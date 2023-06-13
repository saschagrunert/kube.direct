package function

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// Handle is the entrypoint of the go function.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	NewHandler().Handle(ctx, res, req)
}

// Data is the data returned to the API user.
type Data struct {
	Nodes                   int    `json:"nodes,omitempty"`
	KubernetesVersion       string `json:"kubernetes_version,omitempty"`
	OSImage                 string `json:"os_image,omitempty"`
	KernelVersion           string `json:"kernel_version,omitempty"`
	ContainerRuntimeVersion string `json:"container_runtime_version,omitempty"`
}

type Handler struct {
	impl
}

func NewHandler() *Handler {
	return &Handler{
		impl: &defaultImpl{},
	}
}

// Handle the HTTP request.
func (h *Handler) Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	log.Printf("Received request: %+v", req)

	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	config, err := h.InClusterConfig()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	client, err := h.NewForConfig(config)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	discoveryClient, err := h.NewDiscoveryClientForConfig(config)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	nodeList, err := h.ListNodes(ctx, client)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	nodes := nodeList.Items

	if len(nodes) == 0 {
		http.Error(res, "node list is empty", http.StatusInternalServerError)
		return
	}

	version, err := h.ServerVersion(discoveryClient)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	nodeInfo := nodes[0].Status.NodeInfo
	data := Data{
		Nodes:                   len(nodes),
		KubernetesVersion:       version.String(),
		OSImage:                 nodeInfo.OSImage,
		KernelVersion:           nodeInfo.KernelVersion,
		ContainerRuntimeVersion: nodeInfo.ContainerRuntimeVersion,
	}

	res.Header().Set("Content-Type", "application/json")
	if err := h.EncodeJSON(json.NewEncoder(res), &data); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
