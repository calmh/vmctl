package vmadm

import (
	"bufio"
	"bytes"
	"strconv"
)

type TerseVM struct {
	UUID  string
	Type  string
	RAM   int
	State string
	Alias string
}

func List() ([]TerseVM, error) {
	bs, err := vmadmList()
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
