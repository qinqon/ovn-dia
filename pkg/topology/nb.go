package topology

import (
	"context"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/generic"
	"github.com/ovn-org/libovsdb/client"
	"github.com/qinqon/ovn-dia/pkg/nbdb"
)

type SwitchPorts map[string]*SwitchPort
type RouterPorts map[string]*RouterPort
type Switches map[string]*Switch
type Routers map[string]*Router

type NorthBound struct {
	SwitchPorts      SwitchPorts
	RouterPorts      RouterPorts
	RouterPortByName RouterPorts
	Switches         Switches
	Routers          Routers
	RouterByName     Routers
}

type SwitchPort struct {
	Dia        *diagram.Node
	NB         *nbdb.LogicalSwitchPort
	Owner      *Switch
	RouterPort *RouterPort
}

type RouterPort struct {
	Dia   *diagram.Node
	NB    *nbdb.LogicalRouterPort
	Owner *Router
	Peer  *Router
}

type Switch struct {
	Dia   *diagram.Node
	NB    *nbdb.LogicalSwitch
	Ports SwitchPorts
}

type Router struct {
	Dia   *diagram.Node
	NB    *nbdb.LogicalRouter
	Ports RouterPorts
}

func newNorthBound() *NorthBound {
	return &NorthBound{
		SwitchPorts:      SwitchPorts{},
		RouterPorts:      RouterPorts{},
		RouterPortByName: RouterPorts{},
		Switches:         Switches{},
		Routers:          Routers{},
		RouterByName:     Routers{},
	}
}

func LoadNorthBound(ctx context.Context, cli client.Client) (*NorthBound, error) {
	nb := newNorthBound()
	if err := nb.loadSwitches(ctx, cli); err != nil {
		return nil, err
	}

	if err := nb.loadRouters(ctx, cli); err != nil {
		return nil, err
	}

	if err := nb.loadSwitchPorts(ctx, cli); err != nil {
		return nil, err
	}

	if err := nb.loadRouterPorts(ctx, cli); err != nil {
		return nil, err
	}

	nb.resolveReferences()

	return nb, nil
}

func (nb *NorthBound) loadSwitches(ctx context.Context, cli client.Client) error {
	lss := []*nbdb.LogicalSwitch{}
	if err := cli.List(ctx, &lss); err != nil {
		return err
	}
	for _, ls := range lss {
		nb.Switches[ls.UUID] = &Switch{
			NB:  ls,
			Dia: generic.Network.Switch(diagram.NodeLabel(ls.Name)),
		}
	}
	return nil
}

func (nb *NorthBound) loadRouters(ctx context.Context, cli client.Client) error {
	lrs := []*nbdb.LogicalRouter{}
	if err := cli.List(ctx, &lrs); err != nil {
		return err
	}
	for _, lr := range lrs {
		router := &Router{
			NB:  lr,
			Dia: generic.Network.Router(diagram.NodeLabel(lr.Name)),
		}
		nb.Routers[lr.UUID] = router
		nb.RouterByName[lr.Name] = router
	}
	return nil
}

func (nb *NorthBound) loadSwitchPorts(ctx context.Context, cli client.Client) error {
	lsps := []*nbdb.LogicalSwitchPort{}
	if err := cli.List(ctx, &lsps); err != nil {
		return err
	}
	for _, lsp := range lsps {
		nb.SwitchPorts[lsp.UUID] = &SwitchPort{
			NB: lsp,
		}
	}
	return nil
}

func (nb *NorthBound) loadRouterPorts(ctx context.Context, cli client.Client) error {
	lrps := []*nbdb.LogicalRouterPort{}
	if err := cli.List(ctx, &lrps); err != nil {
		return err
	}
	for _, lrp := range lrps {
		port := &RouterPort{
			NB: lrp,
		}
		nb.RouterPorts[lrp.UUID] = port
		nb.RouterPortByName[lrp.Name] = port
	}
	return nil
}

func (nb *NorthBound) resolveReferences() {
	for _, sw := range nb.Switches {
		sw.Ports = SwitchPorts{}
		for _, portUUID := range sw.NB.Ports {
			port, ok := nb.SwitchPorts[portUUID]
			if ok {
				port.Owner = sw
				sw.Ports[portUUID] = port
			}
		}
	}
	for _, rt := range nb.Routers {
		rt.Ports = RouterPorts{}
		for _, portUUID := range rt.NB.Ports {
			port, ok := nb.RouterPorts[portUUID]
			if ok {
				port.Owner = rt
				rt.Ports[portUUID] = port
			}
		}
	}
	for _, swp := range nb.SwitchPorts {
		if swp.NB.Options == nil {
			continue
		}
		routerPortName, ok := swp.NB.Options["router-port"]
		if !ok {
			continue
		}
		routerPort, ok := nb.RouterPortByName[routerPortName]
		if !ok {
			continue
		}
		swp.RouterPort = routerPort
	}

	for _, rtp := range nb.RouterPorts {
		if rtp.NB.Peer == nil {
			continue
		}
		routerName := *rtp.NB.Peer
		router, ok := nb.Routers[routerName]
		if !ok {
			continue
		}
		rtp.Peer = router
	}
}
