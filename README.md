<br/>
<p align="center">
    <a href="https://github.com/ad/domru/blob/master/LICENSE" target="_blank">
        <img src="https://img.shields.io/github/license/ad/domru" alt="GitHub license">
    </a>
    <a href="https://github.com/ad/domru/actions" target="_blank">
        <img src="https://github.com/ad/domru/workflows/Release%20on%20commit%20or%20tag/badge.svg" alt="GitHub actions status">
    </a>
</p>

**moleus/domru** is a fork of [ad/domru](https://github.com/ad/domru), web server what allows you to control your domofon.

## Run in Docker
```shell
cp example.accounts.json accounts.json
docker build -t moleus/domru:latest .

docker run --name domru --rm -p 8080:18000 -e DOMRU_PORT=18000 -v $(pwd)/accounts.json:/share/domofon/account.json moleus/domru:latest
open http://localhost:8080/login

enter phone number in format 79xxxxxxxxx
choose your address
enter sms code, you will see received token and refresh token

restart docker container

docker run --name domru --rm -p 8080:18000 -e DOMRU_PORT=18000 -v $(pwd)/accounts.json:/share/domofon/account.json moleus/domru:latest

now go to http://localhost:8080
```

## Or Docker Compose
```
docker-compose up
```

And open in browser [http://localhost:8080/snapshot](http://localhost:8080/snapshot)

## Docker image
You can find docker image in packages section of this github repository or pull it using `docker pull ghcr.io/moleus/domru:latest`

## Endpoints and methods

| Endpoint | Method | Description |
| --- | --- | --- |
| `/` | GET | Main interface |
| `/login` | GET/POST | Auth interface |
| `/login/address` | POST | Get address by `phone` and `index` |
| `/sms` | POST | Request sms by `code` |
| `/cameras` | GET | Get list of camera |
| `/door` | GET/POST | Open door by `placeID` and `accessControlID` |
| `/events` | GET | Get list of events |
| `/events/last` | GET | Get last event |
| `/finances` | GET | Get finance info |
| `/operators` | GET | Get operators list |
| `/places` | GET | Get places list |
| `/snapshot` | GET | Get snapshot by `placeID` and `accessControlID` |
| `/stream` | GET | Get link to stream by `cameraID` |


## 🤝&nbsp; Found a bug? Missing a specific feature?

Feel free to **file a new issue** with a respective title and description on the the [moleus/domru](https://github.com/moleus/domru/issues) repository. If you already found a solution to your problem, **we would love to review your pull request**!

## 📘&nbsp; License

Released under the terms of the [MIT License](LICENSE).
