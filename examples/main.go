package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jsam/etcd_pulumi_backend"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// Create etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer client.Close()

	// Create etcd backend
	backend := etcd_pulumi_backend.NewEtcdBackend(client, "pulumi-states")

	// Use the backend in your Pulumi program
	err = pulumi.RunErr(func(ctx *pulumi.Context) error {
		// Your Pulumi program here
		return nil
	}, pulumi.WithBackend(backend))

	if err != nil {
		log.Fatalf("Error running Pulumi program: %v", err)
	}

	// Example of using the backend directly
	stackName := "mystack"
	snapshot := []byte(`{"resources": []}`)

	err = backend.SetStack(context.Background(), stackName, snapshot)
	if err != nil {
		log.Fatalf("Error setting stack: %v", err)
	}

	retrievedSnapshot, err := backend.GetStack(context.Background(), stackName)
	if err != nil {
		log.Fatalf("Error getting stack: %v", err)
	}

	fmt.Printf("Retrieved snapshot: %s\n", string(retrievedSnapshot))
}
