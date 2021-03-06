// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	server "github.com/euskadi31/go-server"
	"github.com/euskadi31/go-server/response"
	"github.com/gorilla/mux"
	"github.com/hyperscale/hypercloud/pkg/hypercloud/docker"
	"github.com/rs/zerolog/log"
)

var acceptedLabels = map[string]bool{
	"country":    true,
	"datacenter": true,
	"room":       true,
	"rack":       true,
	"block":      true,
	"position":   true,
}

// TopologyController struct
type TopologyController struct {
	dockerClient *docker.Client
}

// NewTopologyController func
func NewTopologyController(dockerClient *docker.Client) (*TopologyController, error) {
	return &TopologyController{
		dockerClient: dockerClient,
	}, nil
}

// Mount endpoints
func (c TopologyController) Mount(r *server.Router) {
	r.HandleFunc("/v1/topologies/{name:[a-z]+}", c.GetTopologiesHandler).Methods(http.MethodGet)
}

// GetTopologiesHandler endpoint
func (c TopologyController) GetTopologiesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)

	label := params["name"]

	if _, ok := acceptedLabels[label]; !ok {
		response.FailureFromError(w, http.StatusNotFound, fmt.Errorf("URL %s not found", r.URL.Path))

		log.Error().Msgf("Topology %s is not valid", label)

		return
	}

	nodes, err := c.dockerClient.NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		response.FailureFromError(w, http.StatusInternalServerError, err)

		log.Error().Err(err).Msg("docker.NodeList")

		return
	}

	log.Debug().Msgf("Nodes founded: %d", len(nodes))

	valueMap := map[string]bool{}

	for _, node := range nodes {
		value, ok := node.Spec.Labels[label]
		if !ok {
			continue
		}

		if _, ok := valueMap[value]; ok {
			continue
		}

		valueMap[value] = true
	}

	list := []string{}

	for value := range valueMap {
		list = append(list, value)
	}

	response.Encode(w, r, http.StatusOK, list)
}
