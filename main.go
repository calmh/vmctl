package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"text/tabwriter"
)

type TerseVM struct {
	UUID  string
	Type  string
	RAM   int
	State string
	Alias string
}

func list() ([]TerseVM, error) {
	// root@anto:/u/vm # vmadm list -p
	// 30dbeb90-4b1d-45a1-b8ad-a1fe93068b5a:LX:1024:running:dl3
	// 61316833-2323-4d07-a4be-4a25d76f63a6:OS:1024:running:mailbk
	bs, err := exec.Command("vmadm", "list", "-p").CombinedOutput()
	if err != nil {
		return nil, err
	}

	var res []TerseVM
	br := bufio.NewScanner(bytes.NewBuffer(bs))
	for br.Scan() {
		fields := bytes.Split(br.Bytes(), []byte(":"))
		if len(fields) != 5 {
			continue
		}
		ram, _ := strconv.Atoi(string(fields[2]))
		res = append(res, TerseVM{
			UUID:  string(fields[0]),
			Type:  string(fields[1]),
			RAM:   ram,
			State: string(fields[3]),
			Alias: string(fields[4]),
		})
	}

	return res, nil
}

func main() {
	cmdList()
}

func cmdList() {
	res, err := list()
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
