package etcd_pulumi_backend

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/config"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClientInterface interface {
	Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error)
	Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error)
	Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error)
	Close() error
}

type EtcdBackend struct {
	client EtcdClientInterface
	prefix string
}

func NewEtcdBackend(client EtcdClientInterface, prefix string) *EtcdBackend {
	return &EtcdBackend{
		client: client,
		prefix: prefix,
	}
}

func (b *EtcdBackend) GetStackCrypter(stackName tokens.QName) (config.Crypter, error) {
	InfoLogger.Printf("Getting stack crypter for %s", stackName)
	return config.NewSymmetricCrypter(nil)
}

func (b *EtcdBackend) GetStack(ctx context.Context, stackName tokens.QName) ([]byte, error) {
	key := fmt.Sprintf("%s/%s", b.prefix, stackName)
	InfoLogger.Printf("Getting stack %s", key)

	resp, err := b.client.Get(ctx, key)
	if err != nil {
		ErrorLogger.Printf("Error getting stack %s: %v", key, err)
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		WarnLogger.Printf("Stack %s not found", key)
		return nil, nil
	}
	DebugLogger.Printf("Retrieved stack %s", key)
	return resp.Kvs[0].Value, nil
}

func (b *EtcdBackend) SetStack(ctx context.Context, stackName tokens.QName, snapshot []byte) error {
	key := fmt.Sprintf("%s/%s", b.prefix, stackName)
	InfoLogger.Printf("Setting stack %s", key)

	_, err := b.client.Put(ctx, key, string(snapshot))
	if err != nil {
		ErrorLogger.Printf("Error setting stack %s: %v", key, err)
		return err
	}
	DebugLogger.Printf("Set stack %s", key)
	return nil
}

func (b *EtcdBackend) RemoveStack(ctx context.Context, stackName tokens.QName) error {
	key := fmt.Sprintf("%s/%s", b.prefix, stackName)
	InfoLogger.Printf("Removing stack %s", key)

	_, err := b.client.Delete(ctx, key)
	if err != nil {
		ErrorLogger.Printf("Error removing stack %s: %v", key, err)
		return err
	}
	DebugLogger.Printf("Removed stack %s", key)
	return nil
}

func (b *EtcdBackend) Close() error {
	InfoLogger.Println("Closing etcd backend")
	return b.client.Close()
}
