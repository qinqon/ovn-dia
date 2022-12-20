package dot

import (
	"github.com/blushft/go-diagrams/diagram"
	"github.com/qinqon/ovn-dia/pkg/topology"
)

func RenderNorthBound(nb *topology.NorthBound) error {
	d, err := diagram.New(
		diagram.Label("ovn-nbdb"),
		diagram.Filename("ovn-nbdb"),
		diagram.Direction("TB"),
	)
	if err != nil {
		return err
	}

	for _, sw := range nb.Switches {
		d.Add(sw.Dia)
	}
	for _, rt := range nb.Routers {
		d.Add(rt.Dia)
	}

	/*
		svc := k8s.Network.Svc(diagram.NodeLabel("http"))

		d.Connect(ingress, svc)

		g := diagram.NewGroup("pods").Label("Deployment").Connect(svc, k8s.Compute.Pod(diagram.NodeLabel("web server")))

		d.Group(g)
	*/

	if err := d.Render(); err != nil {
		return err
	}
	return nil
}
