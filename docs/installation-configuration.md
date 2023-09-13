---
title: Fortios Installation & Configuration
meta_desc: Information on how to install the Fortios provider.
layout: package
---

## Installation

The Pulumi Fortios provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/fortios`](https://www.npmjs.com/package/@pulumiverse/fortios)
* Python: [`pulumiverse_fortios`](https://pypi.org/project/pulumiverse_fortios/)
* Go: [`github.com/pulumiverse/pulumi-fortios/sdk/go/fortios`](https://pkg.go.dev/github.com/pulumiverse/pulumi-fortios/sdk/go/fortios)
* .NET: [`Pulumiverse.Fortios`](https://www.nuget.org/packages/Pulumiverse.Fortios)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `fortios` provider:

- `fortios:apiKey` (environment: `fortios_API_KEY`) - the API key for `fortios`
- `fortios:region` (environment: `fortios_REGION`) - the region in which to deploy resources

### Provider Binary

The Fortios provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource fortios <version>
```

Replace the version string `<version>` with your desired version.
