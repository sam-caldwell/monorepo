locals {
  source_names = [for source in var.sources_enabled : trimprefix(source, "source.")]
}