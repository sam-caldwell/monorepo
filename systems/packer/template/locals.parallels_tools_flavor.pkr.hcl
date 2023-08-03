locals{

  parallels_tools_flavor = var.parallels_tools_flavor == null ? (

  var.is_windows ? (

  var.os_arch == "x86_64" ? "win" : "win-arm"

  ) : (

  var.os_arch == "x86_64" ? "lin" : "lin-arm"

  )

  ) : var.parallels_tools_flavor

}