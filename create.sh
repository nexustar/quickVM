#!/usr/bin/bash

DATA_DIR=$(realpath ~/.local/share/quickVM)
VM_NAME=$1

mkdir -p ${DATA_DIR}/vms/${VM_NAME}/
qemu-img create -b $(realpath ${DATA_DIR}/template/debian-11-amd64.qcow2) -f qcow2 \
  -F qcow2 ${DATA_DIR}/vms/${VM_NAME}/${VM_NAME}.qcow2 16G

echo "instance-id: ${VM_NAME}
local-hostname: ${VM_NAME}
" > ${DATA_DIR}/vms/${VM_NAME}/metadata.yaml

echo "#cloud-config
ssh_authorized_keys:
  - ssh-rsa xxx
" > ${DATA_DIR}/vms/${VM_NAME}/user-data.yaml

cloud-localds ${DATA_DIR}/vms/${VM_NAME}/seed.img ${DATA_DIR}/vms/${VM_NAME}/user-data.yaml ${DATA_DIR}/vms/${VM_NAME}/metadata.yaml

