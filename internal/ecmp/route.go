package ecmp

import (
	"context"
	"fmt"

	ovn "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	"github.com/kubeovn/kube-ovn/pkg/ovs"
	ovnnb "github.com/kubeovn/kube-ovn/pkg/ovsdb/ovnnb"
	"github.com/kubeovn/kube-ovn/pkg/util"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func HandleGwRoute(ctx context.Context, gw *ovn.VpcNatGateway, prefix string, rclient client.Client, nbClient ovs.OVNNbClient) error {
	vpc := &ovn.Vpc{}

	err := rclient.Get(ctx, client.ObjectKey{Name: gw.Spec.Vpc}, vpc)
	if err != nil {
		return fmt.Errorf("error getting VPC: %w", err)
	}

	lr := vpc.Name

	c := NewApiInterface(nbClient)
	return c.AddOrUpdateECMPRoute(lr, prefix, gw.Spec.LanIP)
}

type api struct {
	nbClient ovs.OVNNbClient
}

func NewApiInterface(c ovs.OVNNbClient) Interface {
	return &api{nbClient: c}
}

func (c *api) AddOrUpdateECMPRoute(logicalRouterName, cidr string, nextHops ...string) error {
	policy := ovnnb.LogicalRouterStaticRoutePolicyDstIP
	err := c.nbClient.AddLogicalRouterStaticRoute(logicalRouterName, util.MainRouteTable, policy, cidr, nil, nextHops...)
	if err != nil {
		return err
	}
	return nil
}

func (c *api) DeleteECMPRoute(logicalRouterName, cidr string, nextHops ...string) error {
	policy := ovnnb.LogicalRouterStaticRoutePolicyDstIP
	rtb := util.MainRouteTable
	// err := c.nbClient.AddLogicalRouterStaticRoute(logicalRouterName, util.MainRouteTable, policy, cidr, nil, nextHops...)
	for _, nextHop := range nextHops {
		err := c.nbClient.DeleteLogicalRouterStaticRoute(logicalRouterName, &rtb, &policy, cidr, nextHop)
		if err != nil {
			return err
		}
	}

	return nil
}
