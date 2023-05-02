package main

import (
	"fmt"
	"github.com/colinmarc/hdfs/v2"
)

func main() {

	rootCABundle := "/srv/hops/super_crypto/hdfs/hops_root_ca.pem"
	clientCertificate := "/srv/hops/super_crypto/hdfs/hdfs_certificate_bundle.pem"
	clientKey := "/srv/hops/super_crypto/hdfs/hdfs_priv.pem"

	hdfsOptions := hdfs.ClientOptions{
		Addresses: []string{"rpc.namenode.service.consul:8020"},
		TLS:       true,
		User:      "gibson_meb10000",
	}

	hdfsOptions.RootCABundle = rootCABundle
	hdfsOptions.ClientKey = clientKey
	hdfsOptions.ClientCertificate = clientCertificate

	client, err := hdfs.NewClient(hdfsOptions)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	bytes, err := client.ReadFile("/_test/copytoremote.txt")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	fmt.Sprintf(string(bytes))
	client.Close()
}
