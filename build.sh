#!/bin/bash


# build command
go build -o lebsite main.go

if [ ! -f "./lebsite" ]; then
  echo "Binary Not Found. Make sure you build it first with 'go build'."
  exit 1
fi

# Move binary to /usr/local/bin
sudo mv ./lebsite /usr/local/bin/

# Change binary permission into executable
sudo chmod +x /usr/local/bin/lebsite

echo "lebsite was installed and added to /usr/local/bin. type 'lebsite' to use that."