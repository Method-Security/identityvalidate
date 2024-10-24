# Portal

The `identityvalidate portal` family of commands provide triggers to help validate security controls

## Azure

Triggers to help validate Azure security controls

### Usage

```bash
identityvalidate portal azure
```

### Help Text

```bash
% identityvalidate portal azure -h
Azure triggers

Usage:
  identityvalidate portal azure [command]

Available Commands:
  owa         Azure OWA portal trigger

Flags:
  -h, --help   help for azure

Global Flags:
  -o, --output string        Output format (signal, json, yaml). Default value is signal (default "signal")
  -f, --output-file string   Path to output file. If blank, will output to STDOUT
  -q, --quiet                Suppress output
  -v, --verbose              Verbose output

Use "identityvalidate portal azure [command] --help" for more information about a command.
  ```