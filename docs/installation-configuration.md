---
title: Fortios Setup
meta_desc: Information on how to install the Fortios provider.
layout: installation
---

## Installation

The Pulumi Fortios provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/fortios`](https://www.npmjs.com/package/@pulumiverse/fortios)
* Python: [`pulumiverse_fortios`](https://pypi.org/project/pulumiverse_fortios/)
* Go: [`github.com/pulumiverse/pulumi-fortios/sdk/go/fortios`](https://github.com/pulumiverse/pulumi-fortios/sdk/go/fortios)
* .NET: [`Pulumiverse.Fortios`](https://www.nuget.org/packages/Pulumiverse.Fortios)

### Provider Binary

The Fortios provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource fortios v0.1.7
```

Replace the version string with your desired version.
