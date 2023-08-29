package client

import (
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config generate.yaml schema/microcks-openapi-v1.7.yaml
