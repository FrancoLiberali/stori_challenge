# Stori challenge <!-- omit in toc -->

[![Build Status](https://github.com/FrancoLiberali/stori_challenge/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/FrancoLiberali/stori_challenge/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/FrancoLiberali/stori_challenge)](https://goreportcard.com/report/github.com/FrancoLiberali/stori_challenge)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=FrancoLiberali_stori_challenge&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=FrancoLiberali_stori_challenge)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=FrancoLiberali_stori_challenge&metric=coverage)](https://sonarcloud.io/summary/new_code?id=FrancoLiberali_stori_challenge)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/FrancoLiberali/stori_challenge)

Coding Challenge for Stori made by Franco Liberali

- [Execution](#execution)
- [Practices used](#practices-used)
  - [Linting](#linting)
  - [Unit tests](#unit-tests)
  - [Feature tests](#feature-tests)
  - [Coverage](#coverage)
  - [BDD + TDD](#bdd--tdd)
  - [Pull requests](#pull-requests)
  - [CI](#ci)
- [Dependencies](#dependencies)
- [Emails](#emails)
- [Money](#money)
- [Dates](#dates)
- [The challenge](#the-challenge)
  - [Bonus points](#bonus-points)
  - [Delivery and code requirements](#delivery-and-code-requirements)

## Execution

To run the processing you will need a csv file of transactions. Two examples can be found in txns1.csv and txns2.csv. To run it locally you can:

Run it with go run:

```bash
MAILGUN_API_KEY=<api-key> go run . -file txns2.csv -email you@email.com
```

> :warning: To run it locally you will need a [mailgun](https://www.mailgun.com/) api key. See [emails](#emails) for details.

> :warning: For this to work, your email must be in the list of accepted recipients. Contact [Franco Liberali](mailto:franco.liberali@gmail.com) to be added to the list. See [emails](#emails) for details.

Install it and then run it:

```bash
go install .
MAILGUN_API_KEY=<api-key> stori_challenge -file txns2.csv -email you@email.com
```

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

or directly:

```bash
go test ./...
```

To ensure that they are unitary, mocks are used. They are generated using [mockery](https://vektra.github.io/mockery/latest/). To regenerate them, [install the dependencies](#dependencies) and run:

```bash
go generate ./...
```

### Feature tests

Feature tests (or e2e) are tests that cover the end-to-end system. They are located in the `test_e2e/` folder and are performed under the [BDD](#bdd--tdd) practice. They are executed during [continuous integration](#ci). To run them locally, run:

```bash
make test_e2e
```

### Coverage

The coverage produced by the tests is measured using sonarcloud. It is displayed in the README of the project. In addition, for a pull request to be accepted, the minimum coverage on the new code must be at least 80%.

### BDD + TDD

The project, as usual in my work, was carried out following the BDD ([Behaviour-Driven Development](https://cucumber.io/docs/bdd/)) + TDD ([Test-Driven Development](https://martinfowler.com/bliki/TestDrivenDevelopment.html)) process:

![BDD + TDD](https://www.andolasoft.com/blog/wp-content/uploads/2015/05/TDD-vs-BDD.jpg)

In the BDD process, feature tests (or e2e tests) are written in gherkin language and can be found in `test_e2e/features`. The execution of these tests is then automated using the godogs library. In the TDD process, unit tests are written alongside the corresponding code (whenever possible).

### Pull requests

The main branch is protected and new code can only be added via pull requests. They must pass the [continuous integration](#ci) process in order to be merged. In this case, the review is done by myself as I am only one person.

### CI

The continuous integration process is run every time a pull request or commit is made to the main branch. It is based on Github Actions and covers the [linting](#linting) and [unit testing](#unit-tests) stages.

## Dependencies

Some parts of the development process rely on external dependencies that you will need to install if you want to run them locally:

```bash
make install_dependencies
```

## Emails

Sending of emails is done using [MailSlurp](https://www.mailslurp.com/). Also the reception of emails in the e2e tests is done with this service.

Some considerations:

1. The API Key is hardcoded into the code in both tests and transaction processing. This is a clear security flaw but it was decided to do so in order to avoid that if correctors want to do a local test they need to create an account on that service. For a productive system this would be implemented with environment variables and secrets.
2. The free version of this service is limited to 100 emails per month, so the sending of emails may start to fail if a lot of testing is done.
3. As I am obviously not the owner of the storicard.com domain, the emails are sent from a domain provided by MailSlurp but in a productive system the domain should be configured to avoid phishing.

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
