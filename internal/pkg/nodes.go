package pkg

import (
	"encoding/json"
	"strings"
)

type NodesOutput struct {
	Data []struct {
		Cpu            float64 `json:"cpu"`
		Id             string  `json:"id"`
		Uptime         int     `json:"uptime"`
		SslFingerprint string  `json:"ssl_fingerprint"`
		Disk           int64   `json:"disk"`
		Maxcpu         int     `json:"maxcpu"`
		Node           string  `json:"node"`
		Level          string  `json:"level"`
		Status         string  `json:"status"`
		Mem            int     `json:"mem"`
		Maxmem         int64   `json:"maxmem"`
		Type           string  `json:"type"`
		Maxdisk        int64   `json:"maxdisk"`
	} `json:"data"`
}

func GetNodeIds(data []byte) []string {
	var ni []string
	nodes := NodesOutput{}
	err := json.Unmarshal(data, &nodes)
	CheckErr(err)

	for _, v := range nodes.Data {
		ni = append(ni, strings.Replace(v.Id, "node/", "", -1))
	}

	return ni
}
