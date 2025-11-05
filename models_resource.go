package vanta

import (
	"strconv"
	"time"
)

type ListResourcesOption func(m map[string]string)

func ResourceWithPageSize(pageSize uint32) ListResourcesOption {
	return func(m map[string]string) { m["pageSize"] = strconv.Itoa(int(pageSize)) }
}

func ResourceWithPageCursor(pageCursor string) ListResourcesOption {
	return func(m map[string]string) { m["pageCursor"] = pageCursor }
}

func ResourceConnectionID(connectionID string) ListResourcesOption {
	return func(m map[string]string) { m["connectionId"] = connectionID }
}

func ResourceIsInScope(isInScope bool) ListResourcesOption {
	return func(m map[string]string) { m["isInScope"] = strconv.FormatBool(isInScope) }
}

type ListResourcesOutput struct {
	Results ListResourcesResults `json:"results"`
}

type ListResourcesResults struct {
	PageInfo PageInfo   `json:"pageInfo"`
	Data     []Resource `json:"data"`
}

type Resource struct {
	ResponseType string    `json:"responseType"`
	ResourceKind string    `json:"resourceKind"`
	ResourceID   string    `json:"resourceId"`
	ConnectionID string    `json:"connectionId"`
	DisplayName  string    `json:"displayName"`
	Owner        *string   `json:"owner,omitempty"`
	InScope      bool      `json:"inScope"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	Account      *string   `json:"account,omitempty"`
	Region       *string   `json:"region,omitempty"`
}

type UpdateResourceInput struct {
	InScope     bool    `json:"inScope"`
	Description *string `json:"description,omitempty"`
	OwnerId     *string `json:"ownerId,omitempty"`
}
