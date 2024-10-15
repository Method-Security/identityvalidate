<div align="center">
<h1>identityvalidation</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>

identityvalidation provides security operators the tools they need to validate that their security controls and detections are operating as expected. Designed with data-modeling and data-integration needs in mind, identityvalidation can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The number of security-relevant AWS resources that identityvalidation can enumerate are constantly growing. For the most up to date listing, please see the documentation [here](docs-capabilities)

To learn more about identityvalidation, please see the [Documentation site](https://method-security.github.io/identityvalidation/) for the most detailed information.

## Quick Start

### Get identityvalidation

For the full list of available installation options, please see the [Installation](./docs/getting-started/index.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/identityvalidation`
- `docker run ghcr.io/method-security/identityvalidation:0.0.1`
- Download the latest binary from the [Github Releases](releases) page
- [Installation documentation](./docs/getting-started/index.md)

### Authentication


### General Usage

```bash
identityvalidation
```

#### Examples

```bash
identityvalidation
```


## Contributing

Interested in contributing to identityvalidation? Please see our [Contribution](#) page.

## Want More?

If you're looking for an easy way to tie identityvalidation into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, see [https://method.security]

## Community

identityvalidation is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](github-org).

Have an idea for a Tool to contribute? Open a Discussion [here](discussion).

[verify]: https://github.com/Method-Security/identityvalidation/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/identityvalidation/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/identityvalidation
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/identityvalidation
[release]: https://github.com/Method-Security/identityvalidation/releases
[releases]: https://github.com/Method-Security/identityvalidation/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/identityvalidation.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/identityvalidation/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/identityvalidation?logo=docker&label=docker%20pulls%20%2F%20identityvalidation
[docker-pull]: https://hub.docker.com/r/methodsecurity/identityvalidation
[license]: https://github.com/Method-Security/identityvalidation/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[homepage]: https://method.security
[docs-home]: https://method-security.github.io/identityvalidation
[docs-capabilities]: https://method-security.github.io/identityvalidation/docs/index.html
[discussion]: https://github.com/Method-Security/identityvalidation/discussions
[github-org]: https://github.com/Method-Security