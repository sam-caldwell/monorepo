locals {

  parallels_prlctl = var.parallels_prlctl == null ? (

  var.is_windows ? [

    ["set", "{{ .Name }}", "--efi-boot", "off"]

  ] : [

    ["set", "{{ .Name }}", "--3d-accelerate", "off"],
    ["set", "{{ .Name }}", "--videosize", "16"]

  ]
  ) : var.parallels_prlctl

}