// +build !solaris

package vmadm

func vmadmList() ([]byte, error) {
	bs := []byte(`30dbeb90-4b1d-45a1-b8ad-a1fe93068b5a:LX:1024:running:dl3
61316833-2323-4d07-a4be-4a25d76f63a6:OS:1024:running:mailbk
721161d5-c982-4506-e477-82c2f3ac6549:OS:1024:running:ddclient
84b4bcea-8c5f-4d70-88c2-36a274e3ac8e:LX:1024:running:unifi4
c5ceefa2-ed42-4da1-be99-f80b1945351b:LX:1024:running:plex
e7f024b7-54f7-4840-9dc9-8079678868d6:OS:1024:running:env
26e4d675-d3cf-6d90-ba6d-e33a8a79a1f4:KVM:4096:running:buildslave
9407eae8-0b0e-4348-bd2d-c9f5a4dfc493:OS:16384:running:zlogin4
`)
	return bs, nil
}

func vmadmGet(uuid string) ([]byte, error) {
	bs := []byte(`{
  "zonename": "61316833-2323-4d07-a4be-4a25d76f63a6",
  "autoboot": true,
  "brand": "joyent",
  "limit_priv": "default",
  "v": 1,
  "create_timestamp": "2015-03-12T14:06:04.736Z",
  "image_uuid": "c02a2044-c1bd-11e4-bd8c-dfc1db8b0182",
  "cpu_shares": 100,
  "max_lwps": 2000,
  "max_msg_ids": 4096,
  "max_sem_ids": 4096,
  "max_shm_ids": 4096,
  "max_shm_memory": 1024,
  "zfs_io_priority": 100,
  "max_physical_memory": 1024,
  "max_locked_memory": 1024,
  "max_swap": 1024,
  "billing_id": "00000000-0000-0000-0000-000000000000",
  "tmpfs": 1024,
  "resolvers": [
    "8.8.8.8",
    "8.8.4.4"
  ],
  "alias": "mailbk",
  "hostname": "mailbk.nym.se",
  "dns_domain": "nym.se",
  "owner_uuid": "c8ce395b-7508-4101-aea6-dc2a66e48451",
  "nics": [
    {
      "interface": "net0",
      "mac": "c2:33:ad:fc:4e:e4",
      "nic_tag": "internal",
      "gateway": "172.16.32.1",
      "gateways": [
        "172.16.32.1"
      ],
      "netmask": "255.255.255.0",
      "ip": "172.16.32.28",
      "ips": [
        "172.16.32.28/24",
        "2001:7a0:8003:32::28/64"
      ],
      "allow_ip_spoofing": true,
      "primary": true
    }
  ],
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
}
`)
	return bs, nil
}
