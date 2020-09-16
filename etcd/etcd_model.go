package etcd

import (
	"context"
	"github.com/etcd-io/etcd/clientv3"
	"time"
)

type Etcd struct {
	Cli 		*clientv3.Client
	IsInit		bool
}

func (e *Etcd)Put(k,v string) (*clientv3.PutResponse,error) {
	kv := clientv3.NewKV(e.Cli)
	return kv.Put(context.TODO(),k, v)
}
func (e *Etcd) Get(k string) (string,error) {
	kv := clientv3.NewKV(e.Cli)
	gr,err := kv.Get(context.TODO(), k)
	if len(gr.Kvs) <= 0 {
		return "",err
	}
	return string(gr.Kvs[0].Value),err
}
func (e *Etcd) Delete(k string) (*clientv3.DeleteResponse,error) {
	kv := clientv3.NewKV(e.Cli)
	return kv.Delete(context.TODO(),k)
}
func (e *Etcd)GetWithPrefix(k string) (map[string]string,error) {
	kv := clientv3.NewKV(e.Cli)
	gr,err := kv.Get(context.TODO(),k,clientv3.WithPrefix())
	if len(gr.Kvs) <= 0 {
		return nil,err
	}
	resp := make(map[string]string)
	for _,t := range gr.Kvs {
		resp[string(t.Key)] = string(t.Value)
	}
	return resp,err
}

func (e *Etcd)PutLease(k,v string,ttl int64,end_cn chan<- int) (*clientv3.PutResponse,error) {
	grantResp, err := e.Cli.Lease.Grant(context.TODO(), ttl)
	if err != nil {
		return nil,err
	}
	kv := clientv3.NewKV(e.Cli)
	pr,err := kv.Put(context.TODO(),k,v,clientv3.WithLease(grantResp.ID))
	if err == nil {
		KeepaliveChan, err :=e.Cli.Lease.KeepAlive(context.TODO(),grantResp.ID)
		if err == nil{
			go keepaliveChanShow(KeepaliveChan,ttl,k,end_cn)
		}
	}
	return pr,err
}
func keepaliveChanShow(clk <-chan *clientv3.LeaseKeepAliveResponse,ttl int64,key string,end_cn chan<- int)  {
	timer := time.NewTimer(time.Duration(ttl)*time.Second)
	for {
		select {
		case <-clk:
			timer.Reset(time.Duration(ttl)*time.Second)
			break
		case <-timer.C:
			end_cn<-1
			break
		}
	}
}