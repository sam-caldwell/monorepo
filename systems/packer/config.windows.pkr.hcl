locals {
  /*
   * Windows-specific configuration for windows 10.
   */
  //os_family
  windows = {
    //opsys
    windows = {
      //os_version
      10 = {
        arch             = "x86_64"
        boot_command     = []
        boot_wait        = "60s"
        communicator     = "winrm"
        environment_vars = null
        execute_command  = null
        floppy_files     = null
        guest_os_type    = {
          parallels = "win-10"
          vbox      = "Windows10_64"
          vmware    = "windows9srv-64"
        }
        iso          = "https://software-static.download.prss.microsoft.com/dbazure/988969d5-f34g-4e03-ac9d-1f9786c66750/19045.2006.220908-0225.22h2_release_svc_refresh_CLIENTENTERPRISEEVAL_OEMRET_x64FRE_en-us.iso"
        iso_checksum = "ef7312733a9f5d7d51cfa04ac497671995674ca5e1058d5164d6028f0938d668"
        shutdown     = {
          command = "shutdown /s /t 10 /f /d p:4:1 /c Packer_Provisioning_Shutdown",
          timeout = "10s"
        }
      }
    }
  }
}