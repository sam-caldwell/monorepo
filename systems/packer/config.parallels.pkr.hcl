locals {
  // Create some parallels-specific configurations...
  parallels = {
    prlctl = {
      windows : [
        ["set", "{{ .Name }}", "--efi-boot", "off"]
      ]
      linux : [
        ["set", "{{ .Name }}", "--3d-accelerate", "off"],
        ["set", "{{ .Name }}", "--videosize", "16"]
      ]
      version_file = null
    }

    tool = {
      flavors = {
        "x86_64" : {
          linux : "lin"
          windows : "win"
        }
        "arm64" : {
          linux : "lin-arm"
          windows : "win-arm"
        }
      }
    }
  }
}