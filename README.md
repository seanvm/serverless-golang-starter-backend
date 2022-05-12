# Serverless Golang Starter Backend

This is a basic template for getting started with Golang and AWS Lambda.

It comes pre-installed with:

- serverless framework
- serverless-offline
- [gorm](https://gorm.io/)

### Project Structure

The `handlers` folder contains the entrypoints for the application as defined in the `serverless.yml`
The `app` folder contains the business logic.

The app structure utilizes Repositories and Services. There is a sample `UserService` (`app/user.go`). The `UserRepository` is injected into the `UserService` via a factory initialization method (`app/factories/factories.go`). The `app/database` is the data persistence layer, and contains the repositories implementations and database connection logic.

This starter template contains a sample `user` handler, a `UserService`, `UserRepository`, and some helper functions.

### Setup

Before running the backend, you will need to install Go if you have not already done so.

You can find a download link [here](https://golang.org/doc/install), or you can install using brew - `brew install golang`

You will also need to install the serverless framework and other npm packages:

```
yarn install
```

### Building and Testing

** Note: You must first configure your database environment variables before running the app **

To build the app:

```
make
```

To run the tests:

```
make test
```

### Debugging Locally

You can run `make start` to run the lambda functions locally (using serverless-offline).

### Deploying

To upload from your local machine:

```
serverless deploy --stage dev
```

To upload specific handler (e.g. user handler)

```
serverless deploy -f user --stage dev
```
