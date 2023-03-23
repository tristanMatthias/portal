package model

import "context"

const (
	EventDownloadProgress = "download-progress"
)

type Model struct {
	ctx context.Context
}

func ModelController() *Model {
	return &Model{}
}

func (m *Model) Startup(ctx context.Context) {
	m.ctx = ctx
}
