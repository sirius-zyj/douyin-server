#gateway
cp ../douyin-server ./gateway/

#microservices
# BIN_DIR="../rpc/output/bin/*"
# for d in $BIN_DIR ; do
#   SVR_NAME="${d#$BIN_DIR}"
#   if [ ! -z "$1" ]; then
#     if [ $SVR_NAME != "$1" ]; then
#       continue
#     fi
#   fi
#   cp ../rpc/output/bin/$SVR_NAME ./microservices/$SVR_NAME
# done

#mysql2redis
# cp ../database/mysql2redis ./database/mysql2rediss
