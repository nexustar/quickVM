#!/usr/bin/bash

DATA_DIR=$(realpath ~/.local/share/quickVM)
VM_NAME=$1

qemu-system-x86_64  \
  -machine accel=kvm,type=q35 \
  -cpu host \
  -m 4G \
  -nographic \
  -device virtio-net-pci,netdev=net0 \
  -netdev user,id=net0,hostfwd=tcp::2222-:22 \
  -drive if=virtio,format=qcow2,file=${DATA_DIR}/vms/${VM_NAME}/${VM_NAME}.qcow2 \
  -drive if=virtio,format=raw,file=${DATA_DIR}/vms/${VM_NAME}/seed.img
