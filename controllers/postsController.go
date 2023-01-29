package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dan-kc/go-rest-api/packages/initializers"
	"github.com/dan-kc/go-rest-api/packages/models"
)

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&posts)
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post

	result := initializers.DB.Find(&post, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&post)
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	initializers.DB.Create(&post)
	return c.Status(201).JSON(&post)
}

func UpdatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	id := c.Params("id")

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	initializers.DB.Where("id = ?", id).Updates(&post)
	return c.Status(200).JSON(&post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post

	result := initializers.DB.Delete(&post, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
