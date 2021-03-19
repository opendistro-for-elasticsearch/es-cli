# Using the ODFE CLI

This section introduces you to many of the common features and options available in the 
ODFE Command Line Interface (ODFE CLI).


**Topics**
+ [Getting help](./usage.md#getting-help)
+ [Command structure](./usage.md#command-structure)
+ [Specifying parameter values](./usage.md#specifying-parameter-values)
+ [Auto complete](./usage.md#auto-complete)


## Getting help

You can get help with any command when using the ODFE Command Line Interface.
To do so, simply type `--help` or `help` at the end of a command name. 

For example, the following command displays help for the general ODFE CLI options and the available top-level commands. 

```
$ odfe-cli help

odfe-cli is a unified command line interface for managing ODFE clusters

Usage:
  odfe-cli [command]

Available Commands:
  ad          Manage the Anomaly Detection plugin
  completion  Generate completion script for your shell
  curl        Manage Elasticsearch core features
  help        Help about any command
  knn         Manage the k-NN plugin
  profile     Manage a collection of settings and credentials that you can apply to an odfe-cli command

Flags:
  -c, --config string    Configuration file for odfe-cli, default is /Users//.odfe-cli/config.yaml
  -h, --help             Help for odfe-cli
  -p, --profile string   Use a specific profile from your configuration file
  -v, --version          Version for odfe-cli

Use "odfe-cli [command] --help" for more information about a command.

```

The following command displays Anomaly Detection plugin specific commands. 

```
$ odfe-cli ad help

Use the Anomaly Detection commands to create, configure, and manage detectors.

Usage:
  odfe-cli ad [command]

Available Commands:
  create      Create detectors based on JSON files
  delete      Delete detectors based on a list of IDs, names, or name regex patterns
  get         Get detectors based on a list of IDs, names, or name regex patterns
  start       Start detectors based on a list of IDs, names, or name regex patterns
  stop        Stop detectors based on a list of IDs, names, or name regex patterns
  update      Update detectors based on JSON files

Flags:
  -h, --help   Help for Anomaly Detection

Global Flags:
  -c, --config string    Configuration file for odfe-cli, default is /Users/opendistro/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file

Use "odfe-cli ad [command] --help" for more information about a command.

```


## Command structure

The ODFE CLI uses a multipart structure on the command line that must be specified in this order:

1. The base call to the `odfe-cli` program.

1. The top-level *command*, which typically corresponds to OpenDistro plugins supported by the ODFE CLI.

1. The *subcommand* that specifies which operation to perform.

1. General CLI flags or parameters required by the operation.

    ```
    $ odfe-cli <command> <subcommand> [flags and parameters]
    ```

## Specifying parameter values

Many flags used in the ODFE CLI are simple string, numeric values or boolean, such as the path and pretty in the following example. 
Here, path accepts string value while pretty is a boolean flag.

**Note:** you must use quotation marks around strings that include one or more space characters
```
$ odfe-cli curl get --path _cluster/health --pretty
```

## Auto complete
ODFE CLI includes a command-completion feature that enables you to use the Tab key to complete a partially entered command.
This feature isn't automatically installed, you need to configure it manually.

### Configure auto complete

To enable the auto complete feature, you can use the following command and perform operation as mentioned in the help,
based on your operating system
```
odfe-cli completion --help

To enable shell autocompletion:

Bash:

$ source <(odfe-cli completion bash)

# To enable auto completion for commands for each session, execute once:
Linux:
  $ odfe-cli completion bash > /etc/bash_completion.d/odfe-cli
MacOS:
  $ odfe-cli completion bash > /usr/local/etc/bash_completion.d/odfe-cli

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To enable auto completion for commands for each session, execute once:
$ odfe-cli completion zsh > "${fpath[1]}/_odfe-cli"

# You will need to start a new shell for this setup to take effect.

Fish:

$ odfe-cli completion fish | source

# To enable auto completion for commands for each session, execute once:
$ odfe-cli completion fish > ~/.config/fish/completions/odfe-cli.fish

Powershell:

PS> odfe-cli completion powershell | Out-String | Invoke-Expression

# To enable auto completion for commands for each session, execute once:
PS> odfe-cli completion powershell > odfe-cli.ps1
# and source this file from your powershell profile.

Usage:
  odfe-cli completion [ bash | zsh | fish | powershell ]

Flags:
  -h, --help   help for completion

Global Flags:
  -c, --config string    Configuration file for odfe-cli, default is /Users/opendistro/.odfe-cli/config.yaml
  -p, --profile string   Use a specific profile from your configuration file

```
### How to use Auto complete
When you partially enter a command, parameter, or flag, the command-completion feature either automatically
completes your command or displays a suggested list of commands. To prompt command completion, you partially enter a command and press Tab.

The following examples show different ways that you can use command completion:
* Partially enter a command and press Tab to display a suggested list of commands.                                                                                      
```
odfe-cli curl p[tab key]
post    put 
```
* Partially enter a parameter and press Tab to display a suggested list of parameters.

```
odfe-cli curl get --[tab key]
--config          --data            --filter-path     --headers         --help            --output-format=  --path=           --profile         --query-params    
--config=         --data=           --filter-path=    --headers=        --output-format   --path            --pretty          --profile=        --query-params= 
```
* Enter a command and press Tab to display sub commands.
```
odfe-cli curl [tab key]
delete  get     post    put     
```