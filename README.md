# spacecore-template

The `spacecore-template` plugin is a built using hashicorp/go-plugin framework to provide seamless integration with the Vimana. This plugin is a starting point of building a spacecore with Vimana.

You should be able to build this project and run the plugin from Vimana CLI.

Vimana runs this spacecore as a plugin in a separate goroutine and communicates with it over gRPC. Spacecore developers would only ned to implement `Start()`, `Stop()`, `Status()`, `Logs()` methods to get started. Just provide how to start, stop, get status and logs of your spacecore.


## Installation

To install the `spacecore-template` plugin, follow these steps:

1. Ensure that you have Go installed on your machine.
2. Build from source and install the plugin:

```shell
go mod tidy
go build -v -o bin/spaceco
```

If you try to run the binary, you will get an error message that says the binary is a go plugin and cannot be executed directly without a client. Vimana will act as a gRPC client that connects to gRPC server running in the plugin and executes the methods.


## Integration with Vimana CLI

The `spacecore-template` plugin is designed to be executed by the Vimana CLI. To scaffold new SpaceCore projects using this plugin as a template, follow these steps:

1. Install the Vimana CLI on your machine.
2. Run the following command to create a new project using the `spacecore-template` plugin:

```shell

export PLUGIN_PATH="/path/to/spacecore-template/bin/spacecods"
vimana plugin $PLUGIN_PATH start
vimana plugin $PLUGIN_PATH status
```

3. The Vimana CLI will use the `spacecore-template` plugin to generate the project structure and files based on the template.

That's it! You can now use the `spacecore-template` plugin with Vimana CLI to easily scaffold new Spacecore projects.


## Examples

Here are a few examples to help you get started with the `spacecore-template` plugin:

- Start the spacecore:

```shell
vimana plugin $PLUGIN_PATH start
```

- Stop the spacecore:

```shell
vimana plugin $PLUGIN_PATH stop
```

- Get status of the spacecore:

```shell
vimana plugin $PLUGIN_PATH status
```

- Get logs of the spacecore:

```shell
vimana plugin $PLUGIN_PATH logs
```


## Contributing

Contributions to the `spacecore-template` plugin are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request on the [GitHub repository](https://github.com/vistara-labs/spacecore-template).

## License

The `spacecore-template` plugin is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute it as per the terms of the license.