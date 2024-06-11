#!/bin/bash

binary_name="lebsite"

binary_path=$(command -v $binary_name)

if [ -z "$binary_path" ]; then
  echo "Binary '$binary_name' not found."
  exit 1
fi

sudo rm $binary_path
sudo rm $binary_name

echo "Binary '$binary_name' deleted."