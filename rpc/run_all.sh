BIN_DIR=$(pwd)/output/bin
binaries=$(ls "$BIN_DIR")

for d in "$binaries"; do
    sh start_one_service.sh $d
done