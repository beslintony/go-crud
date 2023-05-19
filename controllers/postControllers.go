package controllers

import (
	"net/http"

	"github.com/beslintony/go-crud/initializers"
	"github.com/beslintony/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Get the request data from the JSON body
	var postData models.Post
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new Post instance
	post := models.Post{
		Title:  postData.Title,
		Body:   postData.Body,
		Author: postData.Author,
	}

	// Save the post to the database
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	var postData models.Post
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the post from the database
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Update the post fields with the new values if they are non-empty
	if postData.Title != "" {
		post.Title = postData.Title
	}
	if postData.Body != "" {
		post.Body = postData.Body
	}
	if postData.Author != "" {
		post.Author = postData.Author
	}

	// Save the updated post to the database
	result = initializers.DB.Save(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

func DeletePost(c *gin.Context) {
	// Retrieve the post from the database
	var post models.Post
	result := initializers.DB.First(&post, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Delete the post from the database
	result = initializers.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
