# `getter` Provider

The Terraform [getter](https://github.com/EvilSuperstars/terraform-provider-getter) provider is used to interact with files that can be downloaded from a string URL using a variety of protocols.

This provider requires no configuration.

### Example Usage

```hcl
provider "getter" {}

data "getter_file" "foo" {
  url = "data.dat"
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
