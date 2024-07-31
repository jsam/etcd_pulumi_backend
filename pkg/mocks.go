package etcd_pulumi_backend

import (
	"context"

	"github.com/stretchr/testify/mock"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// MockEtcdClient implements EtcdClientInterface for testing
type MockEtcdClient struct {
	mock.Mock
}

func (m *MockEtcdClient) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	args := m.Called(ctx, key, opts)
	return args.Get(0).(*clientv3.GetResponse), args.Error(1)
}

func (m *MockEtcdClient) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	args := m.Called(ctx, key, val, opts)
	return args.Get(0).(*clientv3.PutResponse), args.Error(1)
}

func (m *MockEtcdClient) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	args := m.Called(ctx, key, opts)
	return args.Get(0).(clientv3.WatchChan)
}

func (m *MockEtcdClient) Grant(ctx context.Context, ttl int64) (*clientv3.LeaseGrantResponse, error) {
	args := m.Called(ctx, ttl)
	return args.Get(0).(*clientv3.LeaseGrantResponse), args.Error(1)
}

func (m *MockEtcdClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockEtcdClient) Txn(ctx context.Context) clientv3.Txn {
	args := m.Called(ctx)
	return args.Get(0).(clientv3.Txn)
}

func (m *MockEtcdClient) Revoke(ctx context.Context, id clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*clientv3.LeaseRevokeResponse), args.Error(1)
}

func (m *MockEtcdClient) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	args := m.Called(ctx, key, opts)
	return args.Get(0).(*clientv3.DeleteResponse), args.Error(1)
}

func (m *MockEtcdClient) KeepAlive(ctx context.Context, id clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(<-chan *clientv3.LeaseKeepAliveResponse), args.Error(1)
}

// MockTxn is a mock of the etcd Txn
type MockTxn struct {
	mock.Mock
}

func (m *MockTxn) If(cs ...clientv3.Cmp) clientv3.Txn {
	args := m.Called(cs)
	return args.Get(0).(clientv3.Txn)
}

func (m *MockTxn) Then(ops ...clientv3.Op) clientv3.Txn {
	args := m.Called(ops)
	return args.Get(0).(clientv3.Txn)
}

func (m *MockTxn) Else(ops ...clientv3.Op) clientv3.Txn {
	args := m.Called(ops)
	return args.Get(0).(clientv3.Txn)
}

func (m *MockTxn) Commit() (*clientv3.TxnResponse, error) {
	args := m.Called()
	return args.Get(0).(*clientv3.TxnResponse), args.Error(1)
}
