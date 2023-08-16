# # 检查 kitex 命令是否存在于 PATH 环境变量中
# if ! command -v kitex &>/dev/null; then
#   echo "错误：kitex 命令未在 PATH 中找到，尝试安装......"
#   go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
#   go install github.com/cloudwego/thriftgo@latest
# fi

# mkdir -p rpc
# cd rpc

mkdir -p idl
mkdir -p kitex_gen
kitex -module "douyin-server-rpc" -I idl/ idl/"$1".thrift

mkdir -p service/"$1"
cd service/"$1" && kitex -module "douyin-server-rpc" -service "$1" -use douyin-server-rpc/kitex_gen/ -I ../../idl/ ../../idl/"$1".thrift

go mod tidy