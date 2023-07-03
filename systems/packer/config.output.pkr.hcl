/*
 * systems/packer/config.output.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This file calculates our output_directory
 */
locals {
  output_directory = join("", ["build/packer/", "{{.Provider}}"])
}