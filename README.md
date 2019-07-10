# Go API client for dadapush-client

DaDaPush: Real-time Notifications App Send real-time notifications through our API without coding and maintaining your own app for iOS or Android devices.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.dadapush.com](https://www.dadapush.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./dadapush-client"
```

## Documentation for API Endpoints

All URIs are relative to *https://www.dadapush.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DaDaPushMessageApi* | [**CreateMessage**](docs/DaDaPushMessageApi.md#createmessage) | **Post** /api/v1/message | push Message to a Channel
*DaDaPushMessageApi* | [**DeleteMessage**](docs/DaDaPushMessageApi.md#deletemessage) | **Delete** /api/v1/message/{messageId} | delete a Channel Message
*DaDaPushMessageApi* | [**GetMessage**](docs/DaDaPushMessageApi.md#getmessage) | **Get** /api/v1/message/{messageId} | get a Channel Message
*DaDaPushMessageApi* | [**GetMessages**](docs/DaDaPushMessageApi.md#getmessages) | **Get** /api/v1/messages | get Message List


## Documentation For Models

 - [Action](docs/Action.md)
 - [MessageObject](docs/MessageObject.md)
 - [MessagePushRequest](docs/MessagePushRequest.md)
 - [MessagePushResponse](docs/MessagePushResponse.md)
 - [PageResponseOfMessageObject](docs/PageResponseOfMessageObject.md)
 - [Result](docs/Result.md)
 - [ResultOfMessageObject](docs/ResultOfMessageObject.md)
 - [ResultOfMessagePushResponse](docs/ResultOfMessagePushResponse.md)
 - [ResultOfPageResponseOfMessageObject](docs/ResultOfPageResponseOfMessageObject.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Author

contacts@dadapush.com

