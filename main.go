package main

import (
	"context"
	"flag"

	"github.com/ovn-org/libovsdb/client"
	"github.com/ovn-org/libovsdb/model"
	"github.com/qinqon/ovn-dia/pkg/nbdb"
)

type Context struct {
	context.Context
	nbcli client.Client
}

func main() {
	nbEndpoint := flag.String("nb", "/var/run/ovn/nb.db", "NB endpoint")

	flag.Parse()

	ctx := Context{
		Context: context.Background(),
	}

	nbcli, err := newNBClient(context.Background(), *nbEndpoint)
	if err != nil {
		panic(err)
	}

	ctx.nbcli = nbcli

	/*
		d, err := diagram.New(
			diagram.Label("Kubernetes"),
			diagram.Filename("k8s"),
			diagram.Direction("TB"),
		)

		if err != nil {
			log.Fatal(err)
		}

		ingress := k8s.Network.Ing(diagram.NodeLabel("nginx"))
		svc := k8s.Network.Svc(diagram.NodeLabel("http"))

		d.Connect(ingress, svc)

		g := diagram.NewGroup("pods").Label("Deployment").Connect(svc, k8s.Compute.Pod(diagram.NodeLabel("web server")))

		d.Group(g)

		if err := d.Render(); err != nil {
			log.Fatal(err)
		}
	*/
}

func newNBClient(ctx context.Context, endpoint string) (client.Client, error) {
	ovsNbModel, err := nbdb.FullDatabaseModel()
	if err != nil {
		return nil, err
	}

	return newOVSClient(ctx, ovsNbModel, endpoint)
}

func newOVSClient(ctx context.Context, ovsModel model.ClientDBModel, endpoint string) (client.Client, error) {
	cli, err := client.NewOVSDBClient(ovsModel, client.WithEndpoint(endpoint))

	err = cli.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return cli, nil
}

func logicalSwitches(ctx *Context) ([]nbdb.LogicalSwitch, error) {
	return nil, nil
}
