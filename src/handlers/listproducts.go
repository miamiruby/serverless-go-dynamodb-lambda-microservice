package main

import (
    "context"
    "fmt"
    "encoding/json"
    "os"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Product struct {
    ID          string  `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Done        bool    `json:"done"`
    Price       int64  `json:"price"`
    CreatedAt   string  `json:"created_at"`
}

type ListProductsResponse struct {
    Products        []Product  `json:"products"`
}

var ddb *dynamodb.DynamoDB
func init() {
    region := os.Getenv("AWS_REGION")
    if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
        Region: &region,
        }); err != nil {
                fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
        } else {
                ddb = dynamodb.New(session) // Create DynamoDB client
        }
}

func ListProducts(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        fmt.Println("ListProducts")

        var (
                 tableName = aws.String(os.Getenv("PRODUCTS_TABLE_NAME"))
        )

        // Read from DynamoDB
        input := &dynamodb.ScanInput{
                TableName: tableName,
        }
        result, _ := ddb.Scan(input)

        // Construct products from response
        var products []Product
        for _, i := range result.Items {
                product := Product{}
                if err := dynamodbattribute.UnmarshalMap(i, &product); err != nil {
                        fmt.Println("Failed to unmarshal")
                        fmt.Println(err)
                }
                products = append(products, product)
        }

        // Success HTTP response
        body, _ := json.Marshal(&ListProductsResponse{
                Products: products,
        })
        return events.APIGatewayProxyResponse{
                Body: string(body),
                StatusCode: 200,
        }, nil
}

func main() {
        lambda.Start(ListProducts)
}
