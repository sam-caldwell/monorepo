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
    environment_vars  = local.fact.is_windows?null:local[var.os_family][var.opsys][var.os_version].environment_vars
    execute_command   = local.fact.is_windows?null:local[var.os_family][var.opsys][var.os_version].execute_command
    expect_disconnect = true
    scripts           = local.scripts
    except            = local.except_windows
  }

  # Windows Updates and scripts
  provisioner "powershell" {
    elevated_password = local.fact.user.password
    elevated_user     = local.fact.user.username
    scripts           = local.scripts
    except            = local.only_windows
  }
  provisioner "windows-restart" {
    except = local.only_windows
  }
  provisioner "windows-update" {
    search_criteria = "IsInstalled=0"
    except          = local.only_windows
  }
  provisioner "windows-restart" {
    except = local.only_windows
  }
  provisioner "powershell" {
    elevated_password = local.fact.user.password
    elevated_user     = local.fact.user.username
    scripts = [
      "${path.root}/scripts/windows/cleanup.ps1",
      "${path.root}/scripts/windows/optimize.ps1"
    ]
    except = local.only_windows
  }
  # Convert machines to vagrant boxes
  post-processor "vagrant" {
    compression_level    = local.vagrant_compression_level
    keep_input_artifact  = local.fact.is_windows
    output               = local.vagrant_output
    vagrantfile_template = local.vagrant_template
  }
}
