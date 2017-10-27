# `go-getter` Provider

The Terraform [go-getter](https://github.com/EvilSuperstars/terraform-provider-get) provider is used to interact with files that can be downloaded from a string URL using a variety of protocols.
The provider uses the [go-getter library](https://github.com/hashicorp/go-getter).

This provider requires no configuration.

### Example Usage

```hcl
provider "get" {}

data "get_file" "foo" {
  url = "foo.txt"
}
```

## Data Sources

### `getter_file`

#### Argument Reference

The following arguments are supported:

* `url` - (Required, string) The URL of the file. See [here](https://github.com/hashicorp/go-getter#url-format) for supported URL formats.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `content` - (string) The content of the file.
