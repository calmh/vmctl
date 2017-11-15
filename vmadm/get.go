package vmadm

import (
	"encoding/json"
	"time"
)

type VM struct {
	ZoneName          string    `json:"zonename"`
	Autoboot          bool      `json:"autoboot"`
	Brand             string    `json:"brand"`
	LimitPriv         string    `json:"limit_priv"`
	V                 int       `json:"v"`
	Created           time.Time `json:"create_timestamp"`
	ImageUUID         string    `json:"image_uuid"`
	CPUShares         int       `json:"cpu_shares" name:"cpu-shares"`
	MaxLWPs           int       `json:"max_lwps"`
	MaxMsgIDs         int       `json:"max_msg_ids"`
	MaxSemIDs         int       `json:"max_sem_ids"`
	MaxShmIDs         int       `json:"max_shm_ids"`
	MaxShmMemory      int       `json:"max_shm_memory"`
	ZFSIOPriority     int       `json:"zfs_io_priority" name:"zfs-io-priority"`
	MaxPhysicalMemory int       `json:"max_physical_memory"`
	MaxLockedMemory   int       `json:"max_locked_memory"`
	MaxSwap           int       `json:"max_swap"`
	BillingID         string    `json:"billing_id"`
	TmpFS             int       `json:"tmpfs"`
	Resolvers         []string  `json:"resolvers"`
	Alias             string    `json:"alias"`
	HostName          string    `json:"hostname"`
	DNSDomain         string    `json:"dns_domain" name:"dns-domain"`
	OwnerUUID         string    `json:"owner_uuid"`
	NICs              []NIC     `json:"nics"`
	/*
	  "filesystems": [
	    {
	      "source": "/zones/home",
	      "target": "/home",
	      "type": "lofs"
	    }
	  ],
	  "uuid": "61316833-2323-4d07-a4be-4a25d76f63a6",
	  "zone_state": "running",
	  "zonepath": "/zones/61316833-2323-4d07-a4be-4a25d76f63a6",
	  "zoneid": 3,
	  "last_modified": "2016-09-15T07:59:12.000Z",
	  "firewall_enabled": false,
	  "server_uuid": "080020ff-ffff-ffff-ffff-00144fe7390a",
	  "platform_buildstamp": "20160915T211220Z",
	  "state": "running",
	  "boot_timestamp": "2016-09-21T18:27:41.000Z",
	  "pid": 7110,
	  "customer_metadata": {},
	  "internal_metadata": {},
	  "routes": {},
	  "tags": {},
	  "quota": 10,
	  "zfs_root_compression": "lz4",
	  "zfs_root_recsize": 131072,
	  "zfs_filesystem": "zones/61316833-2323-4d07-a4be-4a25d76f63a6",
	  "zpool": "zones",
	  "snapshots": []
	}*/
}

type NIC struct {
	Interface      string   `json:"interface"`
	MAC            string   `json:"mac"`
	Tag            string   `json:"nic_tag"`
	Gateway        string   `json:"gateway"`
	Gateways       []string `json:"gateways"`
	Netmask        string   `json:"netmask"`
	IP             string   `json:"ip"`
	IPs            []string `json:"ips"`
	AllowIPSpooing bool     `json:"allow_ip_spoofing" name:"allow-ip-spoofing"`
	Primary        bool     `json:"primary"`
}

func Get(uuid string) (VM, error) {
	bs, err := vmadmGet(uuid)
	if err != nil {
		return VM{}, err
	}

	var vm VM
	if err := json.Unmarshal(bs, &vm); err != nil {
		return VM{}, err
	}

	return vm, nil
}
