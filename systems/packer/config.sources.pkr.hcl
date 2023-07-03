/*
 * systems/packer/config.sources.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This file calculates our build sources.
 */

locals{
  source = {
    list = compact([
      local.fact.source.enabled.hyperv,
      local.fact.source.enabled.parallels,
      local.fact.source.enabled.qemu,
      local.fact.source.enabled.virtualbox,
      local.fact.source.enabled.vmware
    ])
  }
}
locals {
  /*
   * Take our source.list and trim the "source." prefix and store the
   * resulting list.
   */
  source_names = [for source in local.source.list : trimprefix(source, "source.")]
  /*
   * Create a list of pruned sources (above) and return this list if we are dealing
   * with Windows.
   */
  except_windows = local.fact.is_windows ? local.source_names : null
  /*
   * Create a list of pruned sources (above) and return this list if we are dealing
   * with any operating system other than Windows..
   */
  only_windows   = local.fact.is_windows ? null : local.source_names
}