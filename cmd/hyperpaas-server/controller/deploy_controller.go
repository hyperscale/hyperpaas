// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/asdine/storm"
	"github.com/docker/docker/client"
	"github.com/euskadi31/go-server"
	"github.com/hyperscale/hyperpaas/database/entity"
	"github.com/hyperscale/hyperpaas/provider/vcs/bitbucket"
)

// PushEvent struct
// @TODO move this in provider
type PushEvent struct {
	Commit string `json:"commit"`
}

// DeployController struct
type DeployController struct {
	dockerClient *client.Client
	db           *storm.DB
	validator    *server.Validator
}

// NewDeployController func
func NewDeployController(dockerClient *client.Client, db *storm.DB, validator *server.Validator) (*DeployController, error) {
	if err := db.Init(&entity.Application{}); err != nil {
		return nil, err
	}

	return &DeployController{
		dockerClient: dockerClient,
		db:           db,
		validator:    validator,
	}, nil
}

// Mount endpoints
func (c DeployController) Mount(r *server.Router) {
	r.AddRouteFunc("/v1/deploy/{id:[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}}/hooks/{provider:[a-z]+}", c.PostDeployHookHandler).Methods(http.MethodPost)
	//r.AddRouteFunc("/v1/deploy/{id:[0-9A-Fa-f\\-]+}/hooks/{provider:[a-z]+}", c.PostDeployHookHandler).Methods(http.MethodPost)
}

// PostDeployHookHandler endpoint
func (c DeployController) PostDeployHookHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	params := mux.Vars(r)

	id := params["id"]
	provider := params["provider"]

	log.Debug().Msgf("ID: %s", id)
	log.Debug().Msgf("Provider: %s", provider)

	var event *PushEvent
	var err error

	switch provider {
	case "bitbucket":
		event, err = c.parseBitbucketEvent(r)
	case "github":
		event, err = c.parseGithubEvent(r)
	case "gitlab":
		event, err = c.parseGitlabEvent(r)
	default:
		server.FailureFromError(w, http.StatusNotFound, fmt.Errorf("URL %s not found", r.URL.Path))

		return
	}

	if err != nil {
		server.FailureFromError(w, http.StatusBadRequest, err)

		return
	}

	server.JSON(w, http.StatusOK, event)
}

//@TODO move this in provider
func (c DeployController) parseBitbucketEvent(r *http.Request) (*PushEvent, error) {
	var event bitbucket.PushEvent

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, err
	}

	return &PushEvent{
		Commit: event.Data.Push.Changes[0].New.Target.Hash,
	}, nil
}

//@TODO move this in provider
func (c DeployController) parseGithubEvent(r *http.Request) (*PushEvent, error) {
	var event bitbucket.PushEvent

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, err
	}

	return &PushEvent{}, nil
}

//@TODO move this in provider
func (c DeployController) parseGitlabEvent(r *http.Request) (*PushEvent, error) {
	var event bitbucket.PushEvent

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, err
	}

	return &PushEvent{}, nil
}