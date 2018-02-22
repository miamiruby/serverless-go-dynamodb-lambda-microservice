package main

import (
        "context"
        "fmt"
        "os"
        "time"
        "encoding/json"

        "github.com/aws/aws-lambda-go/lambda"
        "github.com/aws/aws-lambda-go/events"
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/dynamodb"
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

        "github.com/satori/go.uuid"
)

type Product struct {
        ID          string  `json:"id"`
        Title       string  `json:"title"`
        Description string  `json:"description"`
        Done        bool    `json:"done"`
        Price       int64   `json:"price"`
        CreatedAt   string  `json:"created_at"`
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

func AddProduct(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
        fmt.Println("AddProduct")

        var (
                id = uuid.Must(uuid.NewV4(), nil).String()
                tableName = aws.String(os.Getenv("PRODUCTS_TABLE_NAME"))
        )

        // Initialize product
        product := &Product{
                ID:             id,
                Done:           false,
                Price:          0.00,
                CreatedAt:      time.Now().String(),
        }

        // Parse request body
        json.Unmarshal([]byte(request.Body), product)

        // Write to DynamoDB
        item, _ := dynamodbattribute.MarshalMap(product)
        input := &dynamodb.PutItemInput{
                Item: item,
                TableName: tableName,
        }
        if _, err := ddb.PutItem(input); err != nil {
                return events.APIGatewayProxyResponse{ // Error HTTP response
                        Body: err.Error(),
                        StatusCode: 500,
                }, nil
        } else {
                body, _ := json.Marshal(product)
                return events.APIGatewayProxyResponse{ // Success HTTP response
                        Body: string(body),
                        StatusCode: 200,
                }, nil
        }
}

func main() {
         lambda.Start(AddProduct)
}
