# Stori challenge <!-- omit in toc -->

Coding Challenge for Stori made by Franco Liberali

- [Practices used](#practices-used)
  - [Linting](#linting)
  - [CI](#ci)
- [Dependencies](#dependencies)
- [The challenge](#the-challenge)
  - [Bonus points](#bonus-points)
  - [Delivery and code requirements](#delivery-and-code-requirements)

## Practices used

### Linting

The style of the code is verified using [golangci-lint](https://golangci-lint.run/). The file with the configuration used for the linter can be found at `.golangci.yml`. The linting is executed during the [continuous integration](#ci) process. To run it locally, [install the dependencies](#dependencies) and run:

```bash
make lint
```

### CI

The continuous integration process is run every time a pull request or commit is made to the main branch. It is based on Github Actions and covers the [linting](#linting) stages.

## Dependencies

Some parts of the development process rely on external dependencies that you will need to install if you want to run them locally:

```bash
make install_dependencies
```

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
