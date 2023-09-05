#router
cp ../douyin-server ./router/

#microservices
BIN_DIR="../rpc/output/bin/*"
for d in $BIN_DIR ; do
  SVR_NAME="${d#$BIN_DIR}"
  if [ ! -z "$1" ]; then
    if [ $SVR_NAME != "$1" ]; then
      continue
    fi
  fi
  cp ../rpc/output/bin/$SVR_NAME ./microservices/$SVR_NAME
done
cp ../config/config.yml ./microservices/etcd/

#mysql2redis
cp ../database/dao/douyin.sql ./database/mysql/
# cp ../database/mysql2redis ./database/mysql2rediss
