// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
)

// Reference represents a GitHub reference.
type Reference struct {
	Ref    *string    `json:"ref"`
	URL    *string    `json:"url"`
	Object *GitObject `json:"object"`
}

func (r Reference) String() string {
	return Stringify(r)
}

// GitObject represents a Git object.
type GitObject struct {
	Type *string `json:"type"`
	SHA  *string `json:"sha"`
	URL  *string `json:"url"`
}

func (o GitObject) String() string {
	return Stringify(o)
}

// createRefRequest represents the payload for creating a reference.
type createRefRequest struct {
	Ref *string `json:"ref"`
	SHA *string `json:"sha"`
}

// updateRefRequest represents the payload for updating a reference.
type updateRefRequest struct {
	SHA   *string `json:"sha"`
	Force *bool   `json:"force"`
}

// GetRef fetches the Reference object for a given Git ref.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#get-a-reference
func (s *GitService) GetRef(owner string, repo string, ref string) (*Reference, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/refs/%v", owner, repo, ref)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	r := new(Reference)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// ListRefs lists all refs in a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#get-all-references
func (s *GitService) ListRefs(owner string, repo string) ([]*Reference, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/refs", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var rs []*Reference
	resp, err := s.client.Do(req, &rs)
	if err != nil {
		return nil, resp, err
	}

	return rs, resp, err
}

// CreateRef creates a new ref in a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#create-a-reference
func (s *GitService) CreateRef(owner string, repo string, ref *Reference) (*Reference, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/refs", owner, repo)
	req, err := s.client.NewRequest("POST", u, &createRefRequest{
		Ref: String("refs/" + *ref.Ref),
		SHA: ref.Object.SHA,
	})
	if err != nil {
		return nil, nil, err
	}

	r := new(Reference)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// UpdateRef updates an existing ref in a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#update-a-reference
func (s *GitService) UpdateRef(owner string, repo string, ref *Reference, force bool) (*Reference, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/refs/%v", owner, repo, *ref.Ref)
	req, err := s.client.NewRequest("PATCH", u, &updateRefRequest{
		SHA:   ref.Object.SHA,
		Force: &force,
	})
	if err != nil {
		return nil, nil, err
	}

	r := new(Reference)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// DeleteRef deletes a ref from a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#delete-a-reference
func (s *GitService) DeleteRef(owner string, repo string, ref string) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/refs/%v", owner, repo, ref)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
