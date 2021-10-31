package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"mocerize-api/pkg/config"
	"mocerize-api/pkg/firebase"
	"mocerize-api/pkg/firebase/auth"
	"mocerize-api/pkg/firebase/database"
	"mocerize-api/pkg/firebase/storage"
	"mocerize-api/pkg/routes/api/v1"
)

func main() {

	ctx := context.Background()

	// init config
	if err := config.Load(); err != nil {
		log.Fatalln("Config Load Error ", err)
	}
	log.Println("Environment config loading successfully")

	// init firebase
	if err := firebase.Setup(ctx); err != nil {
		log.Fatalln("Firebase Error", err)
	}
	log.Println("Firebase Setup initializing successfully")

	// init firestore
	if err := firestore.Setup(ctx); err != nil {
		log.Fatalln("Firestore Error ", err)
	}
	log.Println("Firestore Client initializing successfully")

	// init cloud storage
	if err := storage.Setup(ctx); err != nil {
		log.Fatalln("Fire storage ", err)
	}
	log.Println("Cloud Storage initializing successfully")

	//dsnap, err := firestore.Client().Collection("test").Doc("IvwqTa1UK4kBY0j1l1ML").Get(ctx)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//m := dsnap.Data()
	//fmt.Printf("Document data: %#v\n", m["names"])

	// init auth
	if err := auth.Setup(ctx); err != nil {
		log.Fatalln("Firebase Auth Error ", err)
	}
	log.Println("Firebase Auth Client initializing successfully")

	userRecord, err := auth.Client().GetUser(ctx, "PyJtmAkOt0YyFc5yQ4wUpgT4zuB2")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(userRecord.DisplayName)
	//
	//a := auth2.UserToUpdate{}
	//a.DisplayName("Atilla Pehlivan")
	//
	//
	//u, err := auth.Client().CustomToken(ctx,"PyJtmAkOt0YyFc5yQ4wUpgT4zuB2")
	// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJmaXJlYmFzZS1hZG1pbnNkay1xbXFvMkBtb2NrZXJpemUuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLCJhdWQiOiJodHRwczovL2lkZW50aXR5dG9vbGtpdC5nb29nbGVhcGlzLmNvbS9nb29nbGUuaWRlbnRpdHkuaWRlbnRpdHl0b29sa2l0LnYxLklkZW50aXR5VG9vbGtpdCIsImV4cCI6MTYzNTAyMjAwMCwiaWF0IjoxNjM1MDE4NDAwLCJzdWIiOiJmaXJlYmFzZS1hZG1pbnNkay1xbXFvMkBtb2NrZXJpemUuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLCJ1aWQiOiJQeUp0bUFrT3QwWXlGYzV5UTR3VXBnVDR6dUIyIiwiY2xhaW1zIjp7ImF0aSI6ImF0aSJ9fQ.JGQymNOOOwahM2SWmPpIVI6B5AKvp83idEl1dywCzAHHLIdAMffEIPTJSrjEnmOI3n4HWFDBoSAKqoHBv2xO6DGhJs4ew4wBP_jOaC8lAuw9OuGrXgaI6EXIx7wVBmj_qnZXiLF7Gxq9tf3P2-fkfn_Ug2FovXYT3LTpy8pFgYG_LH4ZboDazh7VlBqjMwlxHCNGFdmt2iUs8N5UMd85JkfDJEAZEmIAhkmozozM8t6inLw-ye5xLnXwpMG-zju8j-3db3P5FO7oa0diI7qW1aWzvFJfWclZLanF9XOc6N5YEuj4VAcyyzsrIW39NcmoYZ8DvsYF_-YEgz5h-dvfgA
	//log.Println(u)
	//
	//fmt.Println(u.DisplayName,u.Email)

	// init application
	app := fiber.New(fiber.Config{
		AppName:      "Mockerize",
		ServerHeader: "mockerize.com",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	// middlewares
	app.Use(recover.New())
	//app.Use(middleware.Auth)
	app.Use(logger.New())

	// monitoring
	app.Get("/monitoring", monitor.New())

	// setup /api/v1/* routes
	route.SetupApiV1(app)

	//404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	defer firestore.Client().Close()

	// run server
	log.Fatal(app.Listen(":" + config.Get("APP_PORT")))

}
