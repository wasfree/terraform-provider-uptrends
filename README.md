# Terraform Provider for Uptrends

Version 0.x of the Uptrends Provider requires Terraform 0.15.x and later, but 1.x is recommended.

* [Terraform Website](https://www.terraform.io)
* [Uptrends Website](https://www.uptrends.com)
* [Uptrends API Documentation](https://www.uptrends.de/support/kb/api)
* [Uptrends Provider Documentation](https://registry.terraform.io/providers/wasfree/uptrends/latest/docs)
* [Uptrends Provider Usage Examples](https://github.com/wasfree/terraform-provider-uptrends/tree/master/examples)

## Usage Example

> When using the Uptrends Provider with Terraform 0.15 and later, the recommended approach is to declare Provider versions in the root module Terraform configuration, using a `required_providers` block as per the following example. For previous versions, please continue to pin the version within the provider block.

```hcl
# We strongly recommend using the required_providers block to set the
# Uptrends Provider source and version being used
terraform {
  required_providers {
    uptrends = {
      source = "wasfree/uptrends"
      version = "=0.2.0"
    }
  }
}

# Configure the Uptrends Provider
provider "uptrends" {

  # More information on the authentication methods supported by
  # the Uptrends Provider can be found here:
  # https://registry.terraform.io/providers/wasfree/uptrends/latest/docs

  # username = "..."
  # password = "..."
}

resource "uptrends_monitor_web" "https" {
  name                       = "example-https-monitor"
  type                       = "Https"
  url                        = "https://example.org/"

  check_interval             = 10
  alert_on_load_time_limit_1 = true
  load_time_limit_1          = 3000
  alert_on_load_time_limit_2 = true
  load_time_limit_2          = 7000

  alert_on_min_bytes         = true
  min_bytes                  = 1024

  expected_http_status_code = 200

  user_agent                = "chrome83_android"
  auth_type                 = "Basic"
  username                  = "user"
  password                  = "secret"

  request_headers {
    name  = "X-Test1-Header"
    value = "true"
  }

  match_pattern {
    pattern     = "example"
    is_positive = true
  }

  selected_checkpoints {
    regions           = ["Asia", "1005"]
    checkpoints       = ["Salzburg", "1"]
    exclude_locations = ["Vancouver"]
  }
}
```

Further [usage documentation is available on the Terraform website](https://www.terraform.io/docs/providers/uptrends/index.html).

## Developer Requirements

* [Terraform](https://www.terraform.io/downloads.html) version 0.15.x + (but 1.x is recommended)
* [Go](https://golang.org/doc/install) version 1.17.x (to build the provider plugin)

### On Windows

If you're on Windows you'll also need:
* [Git Bash for Windows](https://git-scm.com/download/win)
* [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm)

For *GNU32 Make*, make sure its bin path is added to PATH environment variable.*

For *Git Bash for Windows*, at the step of "Adjusting your PATH environment", please choose "Use Git and optional Unix tools from Windows Command Prompt".*

Or install via [Chocolatey](https://chocolatey.org/install) (`Git Bash for Windows` must be installed per steps above)
```powershell
choco install make golang terraform -y
refreshenv
```

You must run `Developing the Provider` commands in `bash` because `sh` scrips are invoked as part of these.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.16+ is **required**). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

First clone the repository to: `$GOPATH/src/github.com/wasfree/terraform-provider-uptrends`

```sh
$ mkdir -p $GOPATH/src/github.com/hashicorp; cd $GOPATH/src/github.com/hashicorp
$ git clone git@github.com:wasfree/terraform-provider-uptrends
$ cd $GOPATH/src/github.com/wasfree/terraform-provider-uptrends
```

Once inside the provider directory, you can run `make tools` to install the dependent tooling required to compile the provider.

At this point you can compile the provider by running `make build`, which will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-uptrends
...
```

You can also cross-compile if necessary:

```sh
GOOS=windows GOARCH=amd64 make build
```

In order to run the `Unit Tests` for the provider, you can run:

```sh
$ make test
```

The majority of tests in the provider are `Acceptance Tests` - which provisions real resources in Uptrends. It's possible to run the entire acceptance test suite by running `make testacc` - however it's likely you'll want to run a subset, which you can do using a prefix, by running:

```sh
make testacc
```

The following Environment Variables must be set in your shell prior to running acceptance tests:

- `UPTRENDS_USERNAME`
- `UPTRENDS_PASSWORD`

**Note:** Acceptance tests create real resources in Uptrends.

---

## Developer: Using the locally compiled Uptrends Provider binary

When using Terraform 0.15 and later, after successfully compiling the Uptrends Provider, you must [instruct Terraform to use your locally compiled provider binary](https://www.terraform.io/docs/commands/cli-config.html#development-overrides-for-provider-developers) instead of the official binary from the Terraform Registry.

For example, add the following to `~/.terraformrc` for a provider binary located in `/home/developer/go/bin`:

```hcl
provider_installation {

  # Use /home/developer/go/bin as an overridden package directory
  # for the wasfree/uptrends provider. This disables the version and checksum
  # verifications for this provider and forces Terraform to look for the
  # uptrends provider plugin in the given directory.
  dev_overrides {
    "wasfree/uptrends" = "/home/developer/go/bin"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

---
