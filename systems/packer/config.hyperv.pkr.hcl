locals {
  // Create some hyperv-specific configurations...
  hyperv = {
    /*
     * By default we always use gen 2 hyperv vms...
     * ...but we try to avoid hyperv generally.
     */
    generation = 2
  }
}