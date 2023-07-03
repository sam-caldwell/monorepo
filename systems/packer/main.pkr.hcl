/*
 * systems/packer/windows10/main.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This file defines the builder and represents
 * where all the things come together to solve
 * our problem.
 */
build {
  /*
   * Define what sources we want to build against.
   * For example parallels, virtualbox, vmware or
   * qemu.  These will be enabled/disabled in the
   * variables.pkrvars.hcl or by passing a source
   * as a file.
   */
  sources = local.source.list
  /*
   * *Nix shell scripts (anything but Windows...cause Redmond is special)
   */
  provisioner "shell" {
    environment_vars = local[var.os_family][var.opsys][var.os_version].environment_vars
    execute_command   = local[var.os_family][var.opsys][var.os_version].execute_command
    expect_disconnect = true
    scripts           = local.scripts
    except            = local.except_windows
  }
#
#  # Windows Updates and scripts
#  provisioner "powershell" {
#    elevated_password = "vagrant"
#    elevated_user     = "vagrant"
#    scripts           = local.scripts
#    except            = local.only_windows
#  }
#  provisioner "windows-restart" {
#    except = local.only_windows
#  }
#  provisioner "windows-update" {
#    search_criteria = "IsInstalled=0"
#    except          = local.only_windows
#  }
#  provisioner "windows-restart" {
#    except = local.only_windows
#  }
#  provisioner "powershell" {
#    elevated_password = "vagrant"
#    elevated_user     = "vagrant"
#    scripts           = [
#      "${path.root}/scripts/windows/cleanup.ps1",
#      "${path.root}/scripts/windows/optimize.ps1"
#    ]
#    except = local.only_windows
#  }
#  # Convert machines to vagrant boxes
#  post-processor "vagrant" {
#    compression_level   = 9
#    keep_input_artifact = local.fact.is_windows
#    output              = join("", [
#      "build/vagrant/boxes/", "{{.Provider}}", ".", var.opsys, ".", var.os_version, ".box"
#    ])
#    vagrantfile_template = local.fact.is_windows ? (var.hyperv_generation == 1 ? "${path.root}/vagrantfile-windows.template" : "${path.root}/vagrantfile-windows-gen2.template") : null
#  }
}
