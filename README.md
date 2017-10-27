Terraform `go-getter` Provider
==============================

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/EvilSuperstars/terraform-provider-get`

```sh
$ mkdir -p $GOPATH/src/github.com/EvilSuperstars; cd $GOPATH/src/github.com/EvilSuperstars
$ git clone git@github.com:EvilSuperstars/terraform-provider-get
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/EvilSuperstars/terraform-provider-get
$ make build
```

Using The Provider
------------------

See the [documentation](using.md) to get started using the [go-getter](https://github.com/EvilSuperstars/terraform-provider-get) provider.
