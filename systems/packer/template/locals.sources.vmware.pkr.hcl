locals {
  # vmware-iso
  vmware_tools_upload_flavor = var.vmware_tools_upload_flavor == null ? (
  var.is_windows ? "windows" : "linux"
  ) : var.vmware_tools_upload_flavor
  vmware_tools_upload_path = var.vmware_tools_upload_path == null ? (
  var.is_windows ? "c:\\vmware-tools.iso" : "/tmp/vmware-tools.iso"
  ) : var.vmware_tools_upload_path
}