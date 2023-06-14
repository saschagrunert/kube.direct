package function

import (
	"context"
	"function/api"
	"log"
	"net/http"
)

// Handle is the entrypoint of the go function.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	NewHandler().Handle(ctx, res, req)
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

	config, err := h.ClusterConfig()
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
	data, err := h.Marshal(&api.Data{
		Nodes:                   uint32(len(nodes)),
		KubernetesVersion:       version.String(),
		OsImage:                 nodeInfo.OSImage,
		KernelVersion:           nodeInfo.KernelVersion,
		ContainerRuntimeVersion: nodeInfo.ContainerRuntimeVersion,
	})
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := h.Write(res, data); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}
