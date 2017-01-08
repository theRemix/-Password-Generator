# Password-Checker
SSH 2017 Hackathon Project

## Running this project

```sh
yarn install
npm start
```

## Release

```sh
npm run release
npm run docker-compile-bin
docker build -t theremix/password-generator:latest .

```
### Running release

```sh
docker run -it --rm --name password-generator -p 9090:80 theremix/password-generator
```
will run the server listening on port 9090
