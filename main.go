package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kastelo/structedit"
	"github.com/kastelo/vmctl/vmadm"
)

func main() {
	cmdGet("")
}

func cmdGet(uuid string) {
	vm, err := vmadm.Get(uuid)
	if err != nil {
		fmt.Println("list:", err)
		return
	}
	structedit.Print(vm)
}

func cmdList() {
	res, err := vmadm.List()
	if err != nil {
		fmt.Println("list:", err)
		return
	}

	tw := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	fmt.Fprintf(tw, "UUID\tType\t%s\tState\tAlias\n", fmt.Sprintf("% 6s", "RAM"))
	for _, entry := range res {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", entry.UUID, entry.Type, fmt.Sprintf("% 6d", entry.RAM), entry.State, entry.Alias)
	}
	tw.Flush()
}
