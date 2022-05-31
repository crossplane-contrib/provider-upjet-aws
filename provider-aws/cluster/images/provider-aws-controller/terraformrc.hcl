provider_installation {
  filesystem_mirror {
    path    = "/terraform/provider-mirror"
    include = ["*/*"]
  }
  direct {
    exclude = ["*/*"]
  }
}
