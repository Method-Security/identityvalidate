# identityvalidate Documentation

Hello and welcome to the identityvalidate documentation. While we always want to provide the most comprehensive documentation possible, we thought you may find the below sections a helpful place to get started.

- The [Getting Started](./getting-started/basic-usage.md) section provides onboarding material
- The [Development](./development/setup.md) header is the best place to get started on developing on top of and with identityvalidate
- See the [Docs](./docs/index.md) section for a comprehensive rundown of identityvalidate capabilities

# About identityvalidate

identityvalidate provides security operators the tools they need to validate that their security controls and detections are operating as expected. Designed with data-modeling and data-integration needs in mind, identityvalidate can be used on its own as an interactive CLI, orchestrated as part of a broader data pipeline, or leveraged from within the Method Platform.

The number of security-relevant Identity Validation resources that identityvalidate can enumerate are constantly growing. For the most up to date listing, please see the documentation [here](./docs/index.md)

To learn more about identityvalidate, please see the [Documentation site](https://method-security.github.io/identityvalidate/) for the most detailed information.

## Quick Start

### Get identityvalidate

For the full list of available installation options, please see the [Installation](./getting-started/installation.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/identityvalidate`
- `docker run ghcr.io/method-security/identityvalidate`
- Download the latest binary from the [Github Releases](https://github.com/Method-Security/identityvalidate/releases/latest) page
- [Installation documentation](./getting-started/installation.md)


### General Usage

```bash
identityvalidate <resource>
```

#### Examples

```bash
identityvalidate portal azure owa --username myusername@example.com
```


## Contributing

Interested in contributing to identityvalidate? Please see our organization wide [Contribution](https://method-security.github.io/community/contribute/discussions.html) page.

## Want More?

If you're looking for an easy way to tie identityvalidate into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, visit us [here](https://method.security)

## Community

identityvalidate is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](https://github.com/Method-Security) or our organization wide documentation [here](https://method-security.github.io).

Have an idea for a Tool to contribute? Open a Discussion [here](https://github.com/Method-Security/Method-Security.github.io/discussions).
