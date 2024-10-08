package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()

    // Enable CORS for all origins
    router.Use(cors.Default())

    // Health check route
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello from Go Backend!",
        })
    })

    // Login route (POST)
    router.POST("/login", func(c *gin.Context) {
        var json struct {
            Username string `json:"username" binding:"required"`
            Password string `json:"password" binding:"required"`
        }

        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if json.Username == "admin" && json.Password == "password" {
            c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
        }
    })

    // Start the server on port 8080
    router.Run() // Default port is 8080
}
