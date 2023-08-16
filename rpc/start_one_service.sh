
output_dir="output"
services_dir="service/"
services=$(ls "$services_dir")

service_name="$1"

found=false
for s in $services; do
  if [ "$s" = "$service_name" ]; then
    found=true
    break
  fi
done

if [ "$found" = false ]; then
  echo "Error: Unrecognized service name: $service_name"
  printf 'Available service names:\n%s\n' "$services"
  exit 1
fi

command="$output_dir/bin/$service_name"

# Check if the bootstrap.sh file exists
if [ -f output/bootstrap-"${service_name}".sh ]; then
  command="$output_dir/bootstrap-${service_name}.sh"
fi

if [ ! -f "$command" ]; then
  echo "Error: Service binary not found: $command"
  exit 1
fi

"$command"