package ecmp

import (
	"fmt"
	"os/exec"
)

type Interface interface {
	AddOrUpdateECMPRoute(lr, prefix string, nextHops ...string) error
	DeleteECMPRoute(lr, prefix string, nextHops ...string) error
}

func NewDefaultInterface() Interface {
	return &ctl{}
}

type ctl struct {
}

func (c *ctl) AddOrUpdateECMPRoute(logicalRouterName, cidr string, nextHops ...string) error {
	for _, nextHop := range nextHops {
		cmd := exec.Command("ovn-nbctl", "--wait=hv",
			"lr-route-add", logicalRouterName, cidr, nextHop, "--ecmp-symmetric-reply=true")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to add/update ECMP route: %s, output: %s", err, string(output))
		}
	}
	return nil
}

func (c *ctl) DeleteECMPRoute(logicalRouterName, cidr string, nextHops ...string) error {
	for _, nextHop := range nextHops {
		cmd := exec.Command("ovn-nbctl", "--wait=hv",
			"lr-route-del", logicalRouterName, cidr, nextHop)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to delete ECMP route: %s, output: %s", err, string(output))
		}
	}
	return nil
}
