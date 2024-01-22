module github.com/FrancoLiberali/stori_challenge/aws_lambda

go 1.18

require (
	github.com/FrancoLiberali/stori_challenge/app v0.0.1
	github.com/aws/aws-lambda-go v1.45.0
)

require (
	github.com/FrancoLiberali/cql v0.1.2 // indirect
	github.com/aws/aws-sdk-go v1.50.0 // indirect
	github.com/elliotchance/pie/v2 v2.8.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v3 v3.2.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v4 v4.0.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20231219180239-dc181d75b848 // indirect
	gorm.io/gorm v1.25.5 // indirect
)

replace github.com/FrancoLiberali/stori_challenge/app => ./../app
