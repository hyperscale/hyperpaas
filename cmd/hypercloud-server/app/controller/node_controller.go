// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/asdine/storm"
	"github.com/docker/docker/api/types"
	server "github.com/euskadi31/go-server"
	"github.com/euskadi31/go-server/request"
	"github.com/euskadi31/go-server/response"
	"github.com/hyperscale/hypercloud/pkg/hypercloud/docker"
	"github.com/rs/zerolog/log"
)

// NodeController struct
type NodeController struct {
	dockerClient *docker.Client
	db           *storm.DB
	validator    *request.Validator
}

// NewNodeController func
func NewNodeController(dockerClient *docker.Client, db *storm.DB, validator *request.Validator) (*NodeController, error) {
	return &NodeController{
		dockerClient: dockerClient,
		db:           db,
		validator:    validator,
	}, nil
}

// Mount endpoints
func (c NodeController) Mount(r *server.Router) {
	r.HandleFunc("/v1/nodes", c.getNodesHandler).Methods(http.MethodGet)
}

// swagger:route GET /v1/nodes Node getNodesHandler
//
// Get the nodes list
//
//     Responses:
//       200: Node
//
func (c NodeController) getNodesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	nodes, err := c.dockerClient.NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		log.Error().Err(err).Msg("NodeList")

		response.FailureFromError(w, http.StatusInternalServerError, err)

		return
	}

	response.Encode(w, r, http.StatusOK, nodes)
}
