# 生成用户客户端代码 /root/code/gomall/rpc_gen
cwgo client --type RPC --server_name user --module github.com/LXJ0000/go-kitex/rpc_gen --I ../idl/ --idl ../idl/user.proto
# 生成用户服务端代码 /root/code/gomall/app/user 直接使用 rpc_gen 下的代码 避免耦合
cwgo server --type RPC --server_name user --module github.com/LXJ0000/go-kitex/app/user --I ../../idl --idl ../../idl/user.proto --pass "-use github.com/LXJ0000/go-kitex/rpc_gen"


cwgo client --type RPC --service user --module github.com/LXJ0000/gomall/rpc_gen --I ../idl/ --idl ../idl/user.proto
