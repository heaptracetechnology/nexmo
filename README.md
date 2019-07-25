# _Nexmo_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/nexmo.svg?branch=master)](https://travis-ci.com/omg-services/nexmo)
[![codecov](https://codecov.io/gh/omg-services/nexmo/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/nexmo)


An OMG service for Nexmo, it allows to send SMS.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Send SMS
```coffee
>>> nexmo send from:'senderPhoneNumber' to:'receiverPhoneNumber' message:'messageText'
{"message-count": "messageCount","messages": ["messageDetails"]}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Send SMS
```shell
$ omg run send -a from=<SENDER_PHONE_NUMBER> -a to=<RECEIVER_PHONE_NUMBER> -a text=<MESSAGE_TEXT> -e API_KEY=<API_KEY> -e API_SECRET=<API_SECRET>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/nexmo/blob/master/LICENSE).
