# C API client for openapi_definition

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project. By using the [OpenAPI spec](https://openapis.org) from a remote server, you can easily generate an API client.

- API version: v0
- Package version: 
- Build package: org.openapitools.codegen.languages.CLibcurlClientCodegen

## Installation
You'll need the `curl 7.58.0` package in order to build the API. To have code formatted nicely, you also need to have uncrustify version 0.67 or later.

# Prerequisites

## Install the `curl 7.58.0` package with the following command on Linux.
```bash
sudo apt remove curl
wget http://curl.haxx.se/download/curl-7.58.0.tar.gz
tar -xvf curl-7.58.0.tar.gz
cd curl-7.58.0/
./configure
make
sudo make install
```
## Install the `uncrustify 0.67` package with the following command on Linux.
```bash
git clone https://github.com/uncrustify/uncrustify.git
cd uncrustify
mkdir build
cd build
cmake ..
make
sudo make install
```

## Compile the sample:
This will compile the generated code and create a library in the build folder which has to be linked to the codes where API will be used.
```bash
mkdir build
cd build
// To install library to specific location, use following commands
cmake -DCMAKE_INSTALL_PREFIX=/pathtolocation ..
// for normal install use following command
cmake ..
make
sudo make install
```
## How to use compiled library
Considering the test/source code which uses the API is written in main.c(respective api include is written and all objects necessary are defined and created)

To compile main.c(considering the file is present in build folder) use following command
-L - location of the library(not required if cmake with normal installation is performed)
-l library name
```bash
gcc main.c -L. -lopenapi_definition -o main
```
Once compiled, you can run it with ``` ./main ```

Note: You don't need to specify includes for models and include folder separately as they are path linked. You just have to import the api.h file in your code, the include linking will work.

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Category | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApiControllerAPI* | [**ApiControllerAPI_version**](docs/ApiControllerAPI.md#ApiControllerAPI_version) | **GET** /api/version | 
*GreetingControllerImplAPI* | [**GreetingControllerImplAPI_greetHuman**](docs/GreetingControllerImplAPI.md#GreetingControllerImplAPI_greetHuman) | **GET** /greet/{name} | 
*GreetingControllerImplAPI* | [**GreetingControllerImplAPI_greetWorld**](docs/GreetingControllerImplAPI.md#GreetingControllerImplAPI_greetWorld) | **GET** /greet | 


## Documentation for Models

 - [error_response_t](docs/error_response.md)
 - [version_response_t](docs/version_response.md)


## Documentation for Authorization

All endpoints do not require authorization.

## Author



