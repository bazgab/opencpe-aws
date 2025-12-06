## Installation
> TODO: Provide a release version so it can be installed more easily

So far there are three main ways to install OpenCPE: 

- Downloading a tarball from Releases
- Using `go install` to install the module directly (requires go1.25.4)
- Building it yourself after cloning the repository (also requires go1.25.4, ideal for working on development)

### From Releases
TBD 

### Using `go install`

As long as you have the go version described in the [go.mod](https://github.com/bazgab/opencpe/blob/master/go.mod) file, you can install it by simply running:

```sh
go install github.com/bazgab/opencpe-aws
```

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

### Check Installation

Now you can test if openCPE is properly configured by running:

```sh 
opencpe-aws
``` 
