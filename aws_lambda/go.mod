module github.com/FrancoLiberali/stori_challenge/aws_lambda

go 1.18

require (
	github.com/FrancoLiberali/stori_challenge/app v0.0.1
	github.com/aws/aws-lambda-go v1.45.0
	gorm.io/gorm v1.25.5
)

require (
	github.com/FrancoLiberali/cql v0.1.2 // indirect
	github.com/aws/aws-sdk-go v1.50.0 // indirect
	github.com/elliotchance/pie/v2 v2.8.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.1 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v3 v3.2.0 // indirect
	github.com/mailjet/mailjet-apiv3-go/v4 v4.0.1 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20231219180239-dc181d75b848 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gorm.io/driver/postgres v1.5.4 // indirect
)

replace github.com/FrancoLiberali/stori_challenge/app => ./../app
