locals {
  qemuargs_windows = [
    ["-drive", "file=${path.root}/win_answer_files/virtio-win.iso,media=cdrom,index=3"],
    [
      "-drive",
      "file=${path.root}/../builds/packer-${var.os_name}-${var.os_version}-${var.os_arch}-qemu/{{ .Name }},if=virtio,cache=writeback,discard=ignore,format=qcow2,index=1"
    ],
  ]
  qemuargs_aarch64 = [
    ["-boot", "strict=off"]
  ]

}

locals {
  qemuargs = var.qemuargs == null ? (
  var.is_windows ? local.qemuargs_windows : (
  var.os_arch == "aarch64" ? local.qemuargs_aarch64 : null
  )) : var.qemuargs
}