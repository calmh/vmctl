package main

import "github.com/kastelo/vmctl/vmadm"

type completer struct {
	uuids   map[string]struct{}
	aliases map[string]string // alias -> uuid
}

func newCompleter(vms []vmadm.TerseVM) *completer {
	c := &completer{
		uuids:   make(map[string]struct{}),
		aliases: make(map[string]string),
	}
	for _, vm := range vms {
		c.uuids[vm.UUID] = struct{}{}
		c.aliases[vm.Alias] = vm.UUID // TODO: duplicates
	}
	return c
}

func (c *completer) complete(line string) string {
	return line
}
