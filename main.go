package main

import (
	"context"
	"flag"
	"log"

	"github.com/qinqon/ovn-dia/pkg/dot"
	"github.com/qinqon/ovn-dia/pkg/nbdb"
	"github.com/qinqon/ovn-dia/pkg/topology"
)

func main() {
	nbEndpoint := flag.String("nb", "/var/run/ovn/nb.db", "NB endpoint")

	flag.Parse()

	ctx := context.Background()

	nbcli, err := nbdb.NewClient(ctx, *nbEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	nb, err := topology.LoadNorthBound(ctx, nbcli)
	if err != nil {
		log.Fatal(err)
	}

	if err := dot.RenderNorthBound(nb); err != nil {
		log.Fatal(err)
	}
}
