# `go-getter` Provider

The Terraform [go-getter](https://github.com/hashicorp/go-getter) provider is used to interact with files that can be downloaded from a string URL using a variety of protocols.

This provider requires no configuration.

### Example Usage

```hcl
provider "go_getter" {}

data "go_getter_file" "foo" {
  source = "data.dat"
}
```

## Data Sources

### `go_getter_file`

#### Argument Reference

The following arguments are supported:

* `source` - (Required, string) The source URL of the file. See [here](https://github.com/hashicorp/go-getter#url-format) for supported URL formats.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `body` - (string) The contents of the file.
