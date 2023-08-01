# https://www.packer.io/docs/templates/hcl_templates/blocks/build
build {
  sources = var.sources_enabled

  # Linux Shell scripts
  provisioner "shell" {
    environment_vars = local.linux_environment_variables
    execute_command = "echo 'vagrant' | {{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
    expect_disconnect = true
    scripts           = local.scripts
    except            = var.is_windows ? local.source_names : null
  }

  # Windows Updates and scripts
  provisioner "powershell" {
    elevated_password = "vagrant"
    elevated_user     = "vagrant"
    scripts           = local.scripts
    except            = var.is_windows ? null : local.source_names
  }
  provisioner "windows-restart" {
    except = var.is_windows ? null : local.source_names
  }
  provisioner "windows-update" {
    search_criteria = "IsInstalled=0"
    except          = var.is_windows ? null : local.source_names
  }
  provisioner "windows-restart" {
    except = var.is_windows ? null : local.source_names
  }
  provisioner "powershell" {
    elevated_password = "vagrant"
    elevated_user     = "vagrant"
    scripts = [
      "${path.root}/scripts/windows/cleanup.ps1",
      "${path.root}/scripts/windows/optimize.ps1"
    ]
    except = var.is_windows ? null : local.source_names
  }

  # Convert machines to vagrant boxes
  post-processor "vagrant" {
    compression_level    = 9
    keep_input_artifact  = var.is_windows
    output               = "${path.root}/../builds/${var.vm_name}-${var.os_name}-${var.os_version}-${var.os_arch}.{{ .Provider }}.box"
    vagrantfile_template = var.is_windows ? "${path.root}/vagrant.template/windows-gen2.template" : null
  }
}