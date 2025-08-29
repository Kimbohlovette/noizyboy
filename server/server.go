package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type RequestPayload struct {
	PushToken string `json:"push_token"`
	Name      string `json:"name"`
}

func main() {

	router := gin.Default()
	expoPushClient := expo.NewPushClient(&expo.ClientConfig{})

	router.GET("/api/v1/make-some-noise", func(c *gin.Context) {
		payload := &RequestPayload{}
		if err := c.ShouldBindBodyWith(payload, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "malformed request body"})
			return
		}

messages := []expo.PushMessage{
			{
				To: []expo.ExponentPushToken{expo.ExponentPushToken(payload.PushToken)},
				Body: "Thanks " + payload.Name + " for subscribing to Noizyboy newsletter. I will give you a weekly digest of all what is happening in Tech. Stay tuned!",
				Data: map[string]string{},
				Sound: "twitter",
				Title: "Noizyboy Weekly updates",
				Priority: expo.HighPriority,Badge: 10,
			},
			{
				To: []expo.ExponentPushToken{expo.ExponentPushToken(payload.PushToken)},
				Body: "Thanks for subscribing to Noizyboy newsletter. I will give you a weekly digest of all what is happening in Tech. Stay tuned!",
				Data: map[string]string{},
				Sound: "twitter",
				Title: "Noizyboy Weekly updates",
				Priority: expo.HighPriority,Badge: 10,
			},{
				To: []expo.ExponentPushToken{expo.ExponentPushToken(payload.PushToken)},
				Body: "Thanks for subscribing to Noizyboy newsletter. I will give you a weekly digest of all what is happening in Tech. Stay tuned!",
				Data: map[string]string{},
				Sound: "twitter",
				Title: "Noizyboy Weekly updates",
				Priority: expo.HighPriority,Badge: 10,
			},{
				To: []expo.ExponentPushToken{expo.ExponentPushToken(payload.PushToken)},
				Body: "Thanks for subscribing to Noizyboy newsletter. I will give you a weekly digest of all what is happening in Tech. Stay tuned!",
				Data: map[string]string{},
				Sound: "twitter",
				Title: "Noizyboy Weekly updates",
				Priority: expo.HighPriority,Badge: 10,
			},
		}

		res, err := expoPushClient.PublishMultiple(messages)
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"message": "malformed request body", "error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"responses": res,
		})
	})

	router.Run(":5000")
}
