package ovs

import (
	"context"

	"github.com/ovn-org/libovsdb/client"
	"github.com/ovn-org/libovsdb/model"
)

func NewClient(ctx context.Context, ovsModel model.ClientDBModel, endpoint string) (client.Client, error) {
	cli, err := client.NewOVSDBClient(ovsModel, client.WithEndpoint(endpoint))

	err = cli.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
