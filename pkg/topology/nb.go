package topology

import (
	"context"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/generic"
	"github.com/ovn-org/libovsdb/client"
	"github.com/qinqon/ovn-dia/pkg/nbdb"
)

type SwitchPorts map[string]SwitchPort
type RouterPorts map[string]RouterPort
type Switches map[string]Switch
type Routers map[string]Router

type NorthBound struct {
	SwitchPorts SwitchPorts
	RouterPorts RouterPorts
	Switches    Switches
	Routers     Routers
}

type SwitchPort struct {
	Dia *diagram.Node
	NB  *nbdb.LogicalSwitchPort
}

type RouterPort struct {
	Dia *diagram.Node
	NB  *nbdb.LogicalRouterPort
}

type Switch struct {
	Dia *diagram.Node
	NB  *nbdb.LogicalSwitch
}

type Router struct {
	Dia *diagram.Node
	NB  *nbdb.LogicalRouter
}

func newNorthBound() *NorthBound {
	return &NorthBound{
		SwitchPorts: SwitchPorts{},
		RouterPorts: RouterPorts{},
		Switches:    Switches{},
		Routers:     Routers{},
	}
}

func LoadNorthBound(ctx context.Context, cli client.Client) (*NorthBound, error) {
	nb := newNorthBound()
	var err error
	nb.Switches, err = loadSwitches(ctx, cli)
	if err != nil {
		return nil, err
	}
	nb.Routers, err = loadRouters(ctx, cli)
	if err != nil {
		return nil, err
	}
	nb.SwitchPorts, err = loadSwitchPorts(ctx, cli)
	if err != nil {
		return nil, err
	}
	nb.RouterPorts, err = loadRouterPorts(ctx, cli)
	if err != nil {
		return nil, err
	}
	return nb, nil
}

func loadSwitches(ctx context.Context, cli client.Client) (Switches, error) {
	lss := []*nbdb.LogicalSwitch{}
	if err := cli.List(ctx, &lss); err != nil {
		return nil, err
	}
	switches := Switches{}
	for _, ls := range lss {
		switches[ls.UUID] = Switch{
			NB:  ls,
			Dia: generic.Network.Switch(diagram.NodeLabel(ls.Name)),
		}
	}
	return switches, nil
}

func loadRouters(ctx context.Context, cli client.Client) (Routers, error) {
	lrs := []*nbdb.LogicalRouter{}
	if err := cli.List(ctx, &lrs); err != nil {
		return nil, err
	}
	routers := Routers{}
	for _, lr := range lrs {
		routers[lr.UUID] = Router{
			NB:  lr,
			Dia: generic.Network.Router(diagram.NodeLabel(lr.Name)),
		}
	}
	return routers, nil
}

func loadSwitchPorts(ctx context.Context, cli client.Client) (SwitchPorts, error) {
	lsps := []*nbdb.LogicalSwitchPort{}
	if err := cli.List(ctx, &lsps); err != nil {
		return nil, err
	}
	switchPorts := SwitchPorts{}
	for _, lsp := range lsps {
		switchPorts[lsp.UUID] = SwitchPort{
			NB: lsp,
		}
	}
	return switchPorts, nil
}

func loadRouterPorts(ctx context.Context, cli client.Client) (RouterPorts, error) {
	lrps := []*nbdb.LogicalRouterPort{}
	if err := cli.List(ctx, &lrps); err != nil {
		return nil, err
	}
	routerPorts := RouterPorts{}
	for _, lrp := range lrps {
		routerPorts[lrp.UUID] = RouterPort{
			NB: lrp,
		}
	}
	return routerPorts, nil
}
