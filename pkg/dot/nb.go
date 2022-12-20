package dot

import (
	"strings"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/qinqon/ovn-dia/pkg/topology"
)

func RenderNorthBound(nb *topology.NorthBound) error {
	d, err := diagram.New(
		diagram.Direction(string(diagram.TopToBottom)),
		diagram.Label("ovn-nbdb"),
		diagram.Filename("ovn-nbdb"),
		diagram.Direction("TB"),
	)
	if err != nil {
		return err
	}

	for _, sw := range nb.Switches {
		d.Add(sw.Dia)
		for _, p := range sw.Ports {
			if p.RouterPort == nil {
				d.Connect(sw.Dia, p.Dia)
				continue
			}
			if p.RouterPort.Owner == nil {
				continue
			}
			d.Connect(sw.Dia, p.RouterPort.Owner.Dia, edgeHeadLabel(strings.Join(p.RouterPort.NB.Networks, ",")))
		}
	}
	for _, rt := range nb.Routers {
		d.Add(rt.Dia)
		for _, p := range rt.Ports {
			if p.Peer == nil {
				continue
			}
			d.Connect(rt.Dia, p.Peer.Dia)
		}
	}

	if err := d.Render(); err != nil {
		return err
	}
	return nil
}

func edgeLabel(label string) diagram.EdgeOption {
	return func(o *diagram.EdgeOptions) {
		o.Label = label
	}
}

func edgeTailLabel(label string) diagram.EdgeOption {
	return func(o *diagram.EdgeOptions) {
		o.Attributes["taillabel"] = label
	}
}

func edgeHeadLabel(label string) diagram.EdgeOption {
	return func(o *diagram.EdgeOptions) {
		o.Attributes["headlabel"] = label
	}
}
