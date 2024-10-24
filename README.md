<div align="center">
<h1>identityvalidate</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>

identityvalidate provides security operators the tools they need to validate that their security controls and detections are operating as expected. Designed with data-modeling and data-integration needs in mind, identityvalidate can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The number of security-relevant AWS resources that identityvalidate can enumerate are constantly growing. For the most up to date listing, please see the documentation [here](docs-capabilities)

To learn more about identityvalidate, please see the [Documentation site](https://method-security.github.io/identityvalidate/) for the most detailed information.

## Quick Start

### Get identityvalidate

For the full list of available installation options, please see the [Installation](./docs/getting-started/index.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/identityvalidate`
- `docker run ghcr.io/method-security/identityvalidate:0.0.1`
- Download the latest binary from the [Github Releases](releases) page
- [Installation documentation](./docs/getting-started/index.md)

### Authentication


### General Usage

```bash
identityvalidate <resourse>
```

#### Examples

```bash
identityvalidate portal azure owa --username myusername@example.com
```


## Contributing

Interested in contributing to identityvalidate? Please see our [Contribution](#) page.

## Want More?

If you're looking for an easy way to tie identityvalidate into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, see [https://method.security]

## Community

identityvalidate is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](github-org).

Have an idea for a Tool to contribute? Open a Discussion [here](discussion).

[verify]: https://github.com/Method-Security/identityvalidate/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/identityvalidate/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/identityvalidate
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/identityvalidate
[release]: https://github.com/Method-Security/identityvalidate/releases
[releases]: https://github.com/Method-Security/identityvalidate/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/identityvalidate.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/identityvalidate/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/identityvalidate?logo=docker&label=docker%20pulls%20%2F%20identityvalidate
[docker-pull]: https://hub.docker.com/r/methodsecurity/identityvalidate
[license]: https://github.com/Method-Security/identityvalidate/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[homepage]: https://method.security
[docs-home]: https://method-security.github.io/identityvalidate
[docs-capabilities]: https://method-security.github.io/identityvalidate/docs/index.html
[discussion]: https://github.com/Method-Security/identityvalidate/discussions
[github-org]: https://github.com/Method-Security