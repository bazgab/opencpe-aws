# OpenCPE - Open Cloud Policy Executor
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
<img width="1159" height="605" alt="openCPEBasicArchitecture" src="https://github.com/user-attachments/assets/c455be7f-57bd-4762-8386-51f0dab7b53d" />


## Installation
> TODO: Provide a release version so it can be installed more easily

#### Development

A development version can be built from source to test changes, first ensuring that your go version matches the version in the [go.mod](https://github.com/bazgab/opencpe/blob/master/go.mod) file.

Start by cloning this repository

```sh
git clone https://github.com/bazgab/opencpe.git && cd opencpe
```
Then go install the project

```sh
go install .
```
**Note**: If your `$GOBIN` is not on path, you will not be able to access the tool from the main command, ensure that by adding the following line to your `~/.bashrc` or `~/.profile`

```sh
export PATH=$PATH:~/go/bin
```

## Getting Started

TBD




