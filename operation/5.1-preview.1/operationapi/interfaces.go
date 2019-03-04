package operationapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/satori/go.uuid"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetOperation(ctx context.Context, operationold uuid.UUID, organization string, pluginID *uuid.UUID) (result operation.Model, err error)
}

var _ BaseClientAPI = (*operation.BaseClient)(nil)
