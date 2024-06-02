# Reddis commands

```shell
valkey-cli ./redis.conf

valkey-cli -h redis-cluster-0.redis -p 7777

valkey-cli --cluster create redis-cluster-0.redis:7777 redis-cluster-1.redis:7777 redis-cluster-2.redis:7777 --cluster-replicas 0 --cluster-yes

valkey-cli --cluster add-node redis-cluster-3.redis:7777 redis-cluster-0.redis:7777

valkey-cli -h redis-cluster-0.redis -p 7777 cluster nodes

valkey-cli -h redis-cluster-0.redis -p 7777 CLUSTER FORGET b037b5db5bfed3c359a35a2cbf470874778c9980

valkey-cli --cluster check redis-cluster-0.redis:7777

valkey-cli --cluster reshard redis-cluster-0.redis:7777 \
	--cluster-from 92385b2ea26a1a27c786cbea34a75f155f87f762 \
	--cluster-to a987d7a94052c6e9426c96ebcd896c2003923eb4 \
	--cluster-slots 3276 \
	--cluster-yes

valkey-cli --cluster fix redis-cluster-0.redis:7777

valkey-cli --cluster rebalance redis-cluster-0.redis:7777 --cluster-use-empty-masters
```


```redis.conf
port 7777
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
appendonly yes 
```