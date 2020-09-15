# hello-frog

## About this plugin
This plugin is a template and a functioning example for a basic JFrog CLI plugin. 
This README shows the expected structure of your plugin's README.

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin-install hello-frog`

Installing a specific version:

`$ jfrog plugin-install hello-frog@version`

## Usage
### Commands
* hello
    - arguments:
        - addressee - The name of the person you would like to greet.
    - flags:
        - shout: Makes output uppercase **[Default: false]**
        - repeat: Greets multiple times **[Default: 1]**

### Environment variables
* HELLO_FROG_GREET_PREFIX - Adds a prefix to every greet **[Default: A new greet from your plugin template: ]**

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).
