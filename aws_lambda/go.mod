module github.com/FrancoLiberali/stori_challenge/aws_lambda

go 1.18

require (
	github.com/FrancoLiberali/stori_challenge/app v0.0.1
	github.com/aws/aws-lambda-go v1.45.0
)

require (
	github.com/aws/aws-sdk-go v1.50.0 // indirect
	github.com/elliotchance/pie/v2 v2.8.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v3 v3.2.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v4 v4.0.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20220321173239-a90fa8a75705 // indirect
)

replace github.com/FrancoLiberali/stori_challenge/app => ./../app
