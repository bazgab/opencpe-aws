# OpenCPE - Open Cloud Policy Executor (AWS)
**_IMPORTANT_: Work in Progress - not ready for usage**

OpenCPE aims to be an all-in-one tool for Cloud resource control focused on notification capabilities with heavily opinionated defaults, aiming at ease of use and maintenance. 

Among the features that differ OpenCPE from other monitoring/notification systems, it includes:
- Pre-defined sensible policies, avoiding a learning curve by requiring users to write their own policies
- All-in-one functionality: logging, mailing and all other actions are not required to install any additional plugins/packages
- Simple configuration, being able to parse all required configuration options by a single JSON file
- Ability to run multiple accounts with specific differing environment needs for each
- Specific edge-case manipulation of policies by including tag-based "ignore" capabilities defined in the configuration file
- Focus on Notification capabilities, for when deleting a resource is a "last-case scenario" and manually having a back and forth with your colleagues is ineffective 

## Architecture Overview
<img width="1095" height="577" alt="OpenCPEBasicArchitecturev2" src="https://github.com/user-attachments/assets/fa02ac6a-53c2-4407-9a4e-e151ff9d51e3" />

## Installation
> TODO: Provide a release version so it can be installed more easily
### From Release
TBD 

### From Source
A development version can be built from source to test changes, first ensuring that your go version matches the version in the [go.mod](https://github.com/bazgab/opencpe/blob/master/go.mod) file.

Start by cloning this repository

```sh
git clone https://github.com/bazgab/opencpe-aws.git && cd opencpe-aws
```
Then go install the project

```sh
go build && go install
```
**Note**: If your `$GOBIN` is not on path, you will not be able to access the tool from the main command, ensure that by adding the following line to your `~/.bashrc` or `~/.profile`

```sh
export PATH=$PATH:~/go/bin
```

## More Information

Extensive information on how to use OpenCPE can be found on the docs folder of this repository. For further questions please do not hesitate to open an issue. 

### Content
- [Getting Started](https://github.com/bazgab/opencpe/blob/master/docs/Getting_Started.md)
- [Installation](https://github.com/bazgab/opencpe/blob/master/docs/Installation.md)
- [Configuration](https://github.com/bazgab/opencpe/blob/master/docs/Configuration.md)
- [Policy Reference](https://github.com/bazgab/opencpe/blob/master/docs/Policy_Reference.md)








