# Serverless REST Product Microservice in Go and Lambda

REST Product microservice application written in Go using DynamoDB and Amazon Lambda.

## Prerequisites

- [Node.js & NPM](https://github.com/creationix/nvm)
- [Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/installation/): `npm install -g serverless`
- [Go](https://golang.org/dl/)
- [dep](https://github.com/golang/dep): `brew install dep && brew upgrade dep`

## Quick Start

0. Clone the repo

```
git clone git@github.com:miamiruby/serverless-go-dynamodb-lambda-microservice.git
cd serverless-go-dynamodb-lambda-microservice
```

1. Install Go dependencies

```
dep ensure
```

2. Compile functions as individual binaries for deployment package:

```
./scripts/build.sh
```

> You need to perform this compilation step before deploying.

3. Deploy!

```
serverless deploy
```

> You can perform steps 2 and 3 simultaneously by running `./scripts/deploy.sh`.

4. To Create a Record:

> Use curl to create product:

```
> curl -X POST -H "Content-Type:application/json" https://<hash>.execute-api.<region>.amazonaws.com/dev/products --data '{
            "title": "The Hitchhikers Guide to the Galaxy",
            "description": "42",
            "done": false,
            "price": 10.00,
            "created_at": "2018-01-23 08:48:21.211887436 +0000 UTC m=+0.045616262"
 }'

```

5. List All Records:

> Use curl to list products:

```
> curl https://<hash>.execute-api.<region>.amazonaws.com/dev/products
{
    "products": [
        {
            "id": "d2e38e20-2e73-1e24-9390-1747cf5d19b5i",
            "title": "The Hitchhikers Guide to the Galaxy",
            "description": "42",
            "done": false,
            "price": 1.00,
            "created_at": "2018-01-23 08:48:21.211887436 +0000 UTC m=+0.045616262"
        }
    ]
}
```

6. Delete Record:

> Use curl to delete a product:

```
> curl -X DELETE https://<hash>.execute-api.<region>.amazonaws.com/dev/products/d2e38e20-2e73-1e24-9390-1747cf5d19b5i
```

