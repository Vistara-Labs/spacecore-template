# Create a container bundle
mkdir -p mycontainer/rootfs
docker export $(docker create alpine) | tar -C mycontainer/rootfs -xvf -

# Generate a default config.json
cd mycontainer
runc spec

# Run the container
sudo runc run mycontainer
