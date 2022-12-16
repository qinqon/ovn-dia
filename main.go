package main

import (
	"context"
	"flag"

	"github.com/ovn-org/libovsdb/client"
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

	nbcli, err := nbdb.NewClient(context.Background(), *nbEndpoint)
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

func logicalSwitches(ctx *Context) ([]nbdb.LogicalSwitch, error) {
	return nil, nil
}
