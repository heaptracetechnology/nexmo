# Nexmo as a microservice
An OMG service for Nexmo, it allows to send SMS.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
<!-- [![Build Status](https://travis-ci.com/heaptracetechnology/microservice-firebase.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-firebase)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-firebase/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-firebase)
[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-web.svg)](https://golangci.com) -->

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Send SMS
```sh
$ omg run send -a from=<SENDER_PHONE_NUMBER> -a to=<RECEIVER_PHONE_NUMBER> -a text=<MESSAGE_TEXT> -e API_KEY=<API_KEY> -e API_SECRET=<API_SECRET>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-nexmo .
```
### RUN
```
docker run -p 3000:3000 microservice-nexmo
```
