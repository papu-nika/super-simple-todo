package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	DynamoClient *dynamodb.Client
	TableName    string
}

func NewAuthHandler() (*AuthHandler, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &AuthHandler{
		DynamoClient: client,
		TableName:    "AppData",
	}, nil
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	// Redirect to Google OAuth
	c.Redirect(http.StatusFound, "https://accounts.google.com/o/oauth2/auth")
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	// Handle Google OAuth callback
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code"})
		return
	}

	// Simulate token exchange and user info retrieval
	userID := "google-user-id" // Replace with actual user ID from Google
	h.saveUser(userID)

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated", "userID": userID})
}

func (h *AuthHandler) saveUser(userID string) error {
	_, err := h.DynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(h.TableName),
		Item: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("USER#%s", userID)},
			"SK": &types.AttributeValueMemberS{Value: "PROFILE"},
		},
	})
	return err
}
