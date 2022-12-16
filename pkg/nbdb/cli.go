package nbdb

import (
	"context"

	"github.com/ovn-org/libovsdb/client"
	"github.com/qinqon/ovn-dia/pkg/ovs"
)

func NewClient(ctx context.Context, endpoint string) (client.Client, error) {
	ovsNbModel, err := FullDatabaseModel()
	if err != nil {
		return nil, err
	}

	return ovs.NewClient(ctx, ovsNbModel, endpoint)
}
