[![CircleCI](https://circleci.com/gh/medhoover/gom.svg?style=svg&circle-token=45019dc7f97b86994b79a44e66305018efd9a22f)](https://circleci.com/gh/medhoover/gom) **_Beta_**

**Got a question? / Want to collaborate?** Feel free to open an [issue](https://github.com/medhoover/gom/issues) or reach me at <gmedhoover@gmail.com>

# Introduction

__gom__ is a Powerful commands manager that simplifies complex scripts execution by defining aliases and execution policies. gom reads your commands from the _YAML_ [config file](#config-file), and execute the passed command alias.

# Usage

![imgur](https://i.imgur.com/0Z65Xbq.gif)

### Create a config file
Create a configuration file named `gom.yaml` and add your commands. More information in [config file](#config-file)

### Launch your commands!
Just type :  ```$ gom [Path to Command...] command [Options]```

For more details check the help manual : `$ gom -h`

# Install

If you have [**_GO_**](https://golang.org) installed you just need  to type:  
` $ go get github.com/medhoover/gom `
___

For others, since this project is still in beta, you would have to do a manual install. Just download the binary file from [release page](https://github.com/medhoover/gom/releases),  
 rename it, and place it in one of your path folders.

# Config file

By default, gom looks for `gom.yaml` file in your current directory. But you can manually set the file location by '-f' flag:
` $ gom -f=path/to/file command`

The config file follows the YAML syntax, and contains the name and the commands properties:
```yaml
name: projectName
commands:
  greet:
    morning:
      - echo Bonjour!
      - echo Guten tag!
      - echo Good morning!
    evening: echo Good evening!

```

A command must be in a form of a string value, which gets executed using its YAML path ($ gom greet evening)
An array of strings is considered as a command list, which gets executed sequently ($ gom greet morning)

# TODO

- Add unit tests
- Add config file initialisation
- Add shortcuts (maybe flags) for environment variables
- Add support for parallel execution
- Simplify installation

# Questions and issues

Please feel free to share your questions and issues in [here](https://github.com/medhoover/gom/issues), I will be happy to help.

# Collaborate

Please do :smiley:.
