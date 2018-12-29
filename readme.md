## Introduction
A restful api example which Integrate gin, mysql and gorm, based on go.

## Features
You can run this app just with a docker environment.

## Installation
Get docker first: [https://www.docker.com](https://www.docker.com).

```bash
git clone https://github.com/bougieL/golearning.git
cd golearning
make dev
```

## API Documentation
### scripts
<!-- #### `make image`
Build base image, just excute it at the first time.

#### `make vendors`
Get the project Dependencies from internet. -->

#### `make dev`
Run the app at development mode, this mode run with livereload function.

#### `make build`
Build the project, the output file is at `./build/golearning`

#### `make release`
Build the app and than run it!

### config
* [http://localhost:8001](http://localhost:8001) go service
* [http://localhost:8002](http://localhost:8002) swagger
* [http://localhost:3307](http://localhost:3307) mysql