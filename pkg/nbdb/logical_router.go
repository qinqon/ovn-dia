// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import "github.com/ovn-org/libovsdb/model"

// LogicalRouter defines an object in Logical_Router table
type LogicalRouter struct {
	UUID              string            `ovsdb:"_uuid"`
	Copp              *string           `ovsdb:"copp"`
	Enabled           *bool             `ovsdb:"enabled"`
	ExternalIDs       map[string]string `ovsdb:"external_ids"`
	LoadBalancer      []string          `ovsdb:"load_balancer"`
	LoadBalancerGroup []string          `ovsdb:"load_balancer_group"`
	Name              string            `ovsdb:"name"`
	Nat               []string          `ovsdb:"nat"`
	Options           map[string]string `ovsdb:"options"`
	Policies          []string          `ovsdb:"policies"`
	Ports             []string          `ovsdb:"ports"`
	StaticRoutes      []string          `ovsdb:"static_routes"`
}

func copyLogicalRouterCopp(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalLogicalRouterCopp(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func copyLogicalRouterEnabled(a *bool) *bool {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalLogicalRouterEnabled(a, b *bool) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func copyLogicalRouterExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLogicalRouterExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func copyLogicalRouterLoadBalancer(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterLoadBalancer(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyLogicalRouterLoadBalancerGroup(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterLoadBalancerGroup(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyLogicalRouterNat(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterNat(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyLogicalRouterOptions(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLogicalRouterOptions(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func copyLogicalRouterPolicies(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterPolicies(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyLogicalRouterPorts(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterPorts(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func copyLogicalRouterStaticRoutes(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterStaticRoutes(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *LogicalRouter) DeepCopyInto(b *LogicalRouter) {
	*b = *a
	b.Copp = copyLogicalRouterCopp(a.Copp)
	b.Enabled = copyLogicalRouterEnabled(a.Enabled)
	b.ExternalIDs = copyLogicalRouterExternalIDs(a.ExternalIDs)
	b.LoadBalancer = copyLogicalRouterLoadBalancer(a.LoadBalancer)
	b.LoadBalancerGroup = copyLogicalRouterLoadBalancerGroup(a.LoadBalancerGroup)
	b.Nat = copyLogicalRouterNat(a.Nat)
	b.Options = copyLogicalRouterOptions(a.Options)
	b.Policies = copyLogicalRouterPolicies(a.Policies)
	b.Ports = copyLogicalRouterPorts(a.Ports)
	b.StaticRoutes = copyLogicalRouterStaticRoutes(a.StaticRoutes)
}

func (a *LogicalRouter) DeepCopy() *LogicalRouter {
	b := new(LogicalRouter)
	a.DeepCopyInto(b)
	return b
}

func (a *LogicalRouter) CloneModelInto(b model.Model) {
	c := b.(*LogicalRouter)
	a.DeepCopyInto(c)
}

func (a *LogicalRouter) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *LogicalRouter) Equals(b *LogicalRouter) bool {
	return a.UUID == b.UUID &&
		equalLogicalRouterCopp(a.Copp, b.Copp) &&
		equalLogicalRouterEnabled(a.Enabled, b.Enabled) &&
		equalLogicalRouterExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		equalLogicalRouterLoadBalancer(a.LoadBalancer, b.LoadBalancer) &&
		equalLogicalRouterLoadBalancerGroup(a.LoadBalancerGroup, b.LoadBalancerGroup) &&
		a.Name == b.Name &&
		equalLogicalRouterNat(a.Nat, b.Nat) &&
		equalLogicalRouterOptions(a.Options, b.Options) &&
		equalLogicalRouterPolicies(a.Policies, b.Policies) &&
		equalLogicalRouterPorts(a.Ports, b.Ports) &&
		equalLogicalRouterStaticRoutes(a.StaticRoutes, b.StaticRoutes)
}

func (a *LogicalRouter) EqualsModel(b model.Model) bool {
	c := b.(*LogicalRouter)
	return a.Equals(c)
}

var _ model.CloneableModel = &LogicalRouter{}
var _ model.ComparableModel = &LogicalRouter{}
