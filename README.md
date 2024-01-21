# Stori challenge <!-- omit in toc -->

[![CI Status](https://github.com/FrancoLiberali/stori_challenge/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/FrancoLiberali/stori_challenge/actions)
[![CD Status](https://github.com/FrancoLiberali/stori_challenge/actions/workflows/cd.yml/badge.svg?branch=main)](https://github.com/FrancoLiberali/stori_challenge/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/FrancoLiberali/stori_challenge)](https://goreportcard.com/report/github.com/FrancoLiberali/stori_challenge)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=FrancoLiberali_stori_challenge&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=FrancoLiberali_stori_challenge)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=FrancoLiberali_stori_challenge&metric=coverage)](https://sonarcloud.io/summary/new_code?id=FrancoLiberali_stori_challenge)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/FrancoLiberali/stori_challenge)

Coding Challenge for Stori made by Franco Liberali

- [Execution](#execution)
  - [Run on AWS Lambda](#run-on-aws-lambda)
  - [Run locally](#run-locally)
    - [Run with docker](#run-with-docker)
      - [Image from the container registry](#image-from-the-container-registry)
      - [Build the image](#build-the-image)
    - [Run with Go](#run-with-go)
- [Practices used](#practices-used)
  - [Linting](#linting)
  - [Unit tests](#unit-tests)
  - [Integration tests](#integration-tests)
  - [Feature tests](#feature-tests)
  - [Static analyzer](#static-analyzer)
  - [BDD + TDD](#bdd--tdd)
  - [Pull requests](#pull-requests)
  - [CI](#ci)
  - [CD](#cd)
  - [Hexagonal architecture](#hexagonal-architecture)
- [Dependencies](#dependencies)
- [Emails](#emails)
- [Money](#money)
- [Dates](#dates)
- [The challenge](#the-challenge)
  - [Bonus points](#bonus-points)
  - [Delivery and code requirements](#delivery-and-code-requirements)

## Execution

To run the processing you will need a csv file of transactions. These files can be either local or hosted on s3 (publicly accessible). Two local examples can be found in `txns1.csv` and `txns2.csv` and its s3 version `s3://fl-stori-challenge/txns1.csv` and `s3://fl-stori-challenge/txns2.csv`. To run it locally you can:

### Run on AWS Lambda

The developed solution is deployed on AWS Lambda. To execute it performs HTTP request to the function, for example, as follows:

```bash
curl 'https://a7cerhsfswfdffbmpioigjay7q0vzkur.lambda-url.us-east-2.on.aws/' -H 'Content-type: application/json' -d '{ "file": "s3://fl-stori-challenge/txns2.csv", "email": "you@email.com" }'
```

> :warning: Don't forget to replace <you@email.com> with your email address

This version only accepts files hosted on AWS S3 (publicly).

The [CD](#cd) process updates the function each time a commit is made to the main branch.

### Run locally

In the local version it is possible to use both local files and files hosted on AWS S3 (publicly).

#### Run with docker

You can use the pre-built image or build it yourself.

##### Image from the container registry

The [CD](#cd) process updates the image in the container registry each time a commit is made to the main branch.

1. Install docker

With local file:

2. `docker run -e EMAIL_PUBLIC_API_KEY=<public-api-key> -e EMAIL_PRIVATE_API_KEY=<private-api-key> -v $(pwd):/data ghcr.io/francoliberali/stori_challenge:latest -file data/txns2.csv -email you@email.com`

> :warning: To run it locally you will need a [mailjet](mailjet.com) key pair. See [emails](#emails) for details.

With AWS S3 hosted file:

2. `docker run -e EMAIL_PUBLIC_API_KEY=<public-api-key> -e EMAIL_PRIVATE_API_KEY=<private-api-key> ghcr.io/francoliberali/stori_challenge:latest -file s3://fl-stori-challenge/txns2.csv -email you@email.com`

##### Build the image

1. Install docker
2. Clone this repository
3. `docker build -t francoliberali/stori_challenge:latest .`
4. `docker run -e EMAIL_PUBLIC_API_KEY=<public-api-key> -e EMAIL_PRIVATE_API_KEY=<private-api-key> -v $(pwd):/data francoliberali/stori_challenge -file data/txns2.csv -email you@email.com`

#### Run with Go

1. Install Go 1.18+.
2. Clone this repository
3. `go mod download`

Run it with go run:

4. `EMAIL_PUBLIC_API_KEY=<public-api-key> EMAIL_PRIVATE_API_KEY=<private-api-key> go run . -file txns2.csv -email you@email.com`

Or install it and then run it:

4. `go install .`
5. `EMAIL_PUBLIC_API_KEY=<public-api-key> EMAIL_PRIVATE_API_KEY=<private-api-key> stori_challenge -file txns2.csv -email you@email.com`

## Practices used

### Linting

The style of the code is verified using [golangci-lint](https://golangci-lint.run/). The file with the configuration used for the linter can be found at `.golangci.yml`. The linting is executed during the [continuous integration](#ci) process. To run it locally, [install the dependencies](#dependencies) and run:

```bash
make lint
```

### Unit tests

Unit tests are added whenever possible. They are performed within the [TDD](#bdd--tdd) methodology. They can be found in files ending in `_test.go` that accompany the tested files. They are executed during [continuous integration](#ci). To run them locally, run:

```bash
make test_unit
```

To ensure that they are unitary, mocks are used. They are generated using [mockery](https://vektra.github.io/mockery/latest/). To regenerate them, [install the dependencies](#dependencies) and run:

```bash
go generate ./app/...
```

### Integration tests

Integration tests test the correct integration between the different components of the system. They are intermediate between unit tests (where any component other than the one being tested is mocked) and e2e tests (where the system is run as a black box). Here, the file reading component is used, but the email sending is mocked. They are executed during [continuous integration](#ci). To run them locally, run:

```bash
make test_integration
```

### Feature tests

Feature tests (or e2e) are tests that cover the end-to-end system. They are located in the `test_e2e/` folder and are performed under the [BDD](#bdd--tdd) practice. They are executed during [continuous integration](#ci). To run them locally, run:

```bash
EMAIL_PUBLIC_API_KEY=<public-api-key> EMAIL_PRIVATE_API_KEY=<private-api-key> make test_e2e
```

For executing this, you will need to have configured your aws credentials in `~/.aws/credentials`. For details see <https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials>.

### Static analyzer

TA static analysis of the code is performed by [sonarcloud](https://sonarcloud.io/project/overview?id=FrancoLiberali_stori_challenge). It checks for bad smells, security problems, test coverage and other issues. It is displayed in the README of the project. In addition, for a pull request to be accepted, these controls must be passed.

### BDD + TDD

The project, as usual in my work, was carried out following the BDD ([Behaviour-Driven Development](https://cucumber.io/docs/bdd/)) + TDD ([Test-Driven Development](https://martinfowler.com/bliki/TestDrivenDevelopment.html)) process:

![BDD + TDD](https://www.andolasoft.com/blog/wp-content/uploads/2015/05/TDD-vs-BDD.jpg)

In the BDD process, feature tests (or e2e tests) are written in gherkin language and can be found in `test_e2e/features`. The execution of these tests is then automated using the godogs library. In the TDD process, unit tests are written alongside the corresponding code (whenever possible).

### Pull requests

The main branch is protected and new code can only be added via pull requests. They must pass the [continuous integration](#ci) process in order to be merged. In this case, the review is done by myself as I am only one person.

### CI

The continuous integration process is run every time a pull request or commit is made to the main branch. It is based on Github Actions and covers the [linting](#linting), [unit testing](#unit-tests), [integration testing](#integration-tests), [feature testing](#feature-tests) and [static analysis](#static-analyzer) stages.

### CD

The continuous delivery process is executed every time a commit is performed on the main branch and the CI process is successful. It builds the docker image and pushes it to the container registry and builds the AWS Lambda function and deploys it.

### Hexagonal architecture

The code was created following the [hexagonal architecture pattern](https://alistair.cockburn.us/hexagonal-architecture/):

![hexagonal architecture](https://upload.wikimedia.org/wikipedia/commons/thumb/7/75/Hexagonal_Architecture.svg/313px-Hexagonal_Architecture.svg.png)

Following it, the components that interact with external services (io for reading csv files and mailjet for sending mails) are isolated in the `app/adapters` package, while the core business that is in `app/service` only depends on interfaces. This allows for easy testability using mocks and component swapping if necessary, without modifying the business logic.

## Dependencies

Some parts of the development process rely on external dependencies that you will need to install if you want to run them locally:

```bash
make install_dependencies
```

## Emails

Sending of emails is done using [Mailjet](https://www.mailjet.com/).

Some considerations:

1. To run the programme locally, it is necessary to have the API Key to send mails. To avoid creating an account you can use the public api key `0ed12cb0ba8d922820a93ea5242db813` and private api key `0ba6132387f60806f1bf9476eb6e1987` (I only add the api key here to simplify the work of correcting this challenge, obviously in real life the api key would never be shared).
2. The free version of this service is limited in the amount of emails per month, so the sending of emails may start to fail if a lot of testing is done.
3. As I am obviously not the owner of the storicard.com domain, the emails are sent from <franco.liberali@gmail.com>. This may result in emails being marked as Spam. Please check your spam box when testing. In a productive system the domain should be configured to avoid this and phishing.

In the feature tests, [MailSlurp](https://www.mailslurp.com/) is used for the reception of emails.

Some considerations:

1. The API Key is hardcoded into the code. This is an accepted simplification as it is only an account used for testing purposes.

## Money

As transactions are, in general, decimal numbers, the decimal library (<https://github.com/shopspring/decimal>) was used for their representation to avoid loss of information when using floating point types such as float64.

## Dates

Dates are considered to be in UTC, so no time zone transformation is ever performed on them. It is also considered that if these dates do not have a year, it means that they are of the current year, and if the year is to be indicated, the format will be "7/15/2023".

## The challenge

For this challenge you must create a system that processes a file from a mounted directory. The file
will contain a list of debit and credit transactions on an account. Your function should process the file
and send summary information to a user in the form of an email.

An example file is shown below; but create your own file for the challenge. Credit transactions are
indicated with a plus sign like +60.5. Debit transactions are indicated by a minus sign like -20.46

```csv
Id,Date,Transaction
0,7/15,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10
```

We prefer that you code in Python or Golang; but other languages are ok too. Package your code in
one or more Docker images. Include any build or run scripts, Dockerfiles or docker-compose files
needed to build and execute your code.

### Bonus points

1. Save transaction and account info to a database
2. Style the email and include Stori’s logo
3. Package and run code on a cloud platform like AWS. Use AWS Lambda and S3 in lieu of Docker.

### Delivery and code requirements

Your project must meet these requirements:

1. The summary email contains information on the total balance in the account, the number of transactions grouped by month, and the average credit and average debit amounts grouped by month. Using the transactions in the image above as an example, the summary info would be

    Total balance is 39.74

    Number of transactions in July: 2

    Number of transactions in August: 2

    Average debit amount: -15.38

    Average credit amount: 35.25

2. Include the file you create in CSV format.
3. Code is versioned in a git repository. The README.md file should describe the code interface and
instructions on how to execute the code.
