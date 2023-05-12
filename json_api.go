package main

import (
	"fmt"
	"grpc/config/database"
	"grpc/middlewares"
	"grpc/util"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

type JSONAPIServer struct {
	listenAddr string
	svc        UserService
}

func NewJSONAPIServer(listenAddr string, svc UserService) *JSONAPIServer {
	return &JSONAPIServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

func (s *JSONAPIServer) Run() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	app := Create()

	RouteInit(app)

	if err := Listen(app); err != nil {
		log.Panic(err)
	}
}

func (s *JSONAPIServer) Login(c *fiber.Ctx) error {

	var body struct{
		Email       string `json:"email"`
		Password    string `json:"password"`
		DeviceToken string `json:"device_token"`
	}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	access_token, err := s.svc.Login(c, body.Email, body.Password, body.DeviceToken)
	if err != nil {
		return err
	}

	// resp := models.LoginResponse {
	// 	AccessToken: access_token,
	// }
	// return c.JSON(resp)
	return c.JSON(fiber.Map{
		"access_token": access_token,
	})
}

// #####################################################
func initMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(compress.New(compress.Config{ Level: compress.LevelBestSpeed }))
	app.Use(etag.New())
	app.Use(limiter.New())
	logFile, _ := os.OpenFile("./logs/" + util.GetYmd() + ".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil { log.Fatalf("err openging file: %v", err) }
	// defer logFile.Close()
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}:${port} ${status} - ${method} ${path} ${protocol} ${latency}\nUser-Agent: ${ua}\nerror : ${error}\nRequset: ${body}\nResponse: ${resBody}\n",
		// Format: "[${time}] ${ips}:${host} ${status} - ${method} ${path} ${protocol} ${latency}\nuser agnet : ${ua}\nerror : ${error}\nrequset body : ${body}\nresponse body : ${resBody}\n",
		Output: logFile,
	}))
}

func Create() *fiber.App {
	database.Init()
	app := fiber.New(fiber.Config{
		// Prefork: true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*util.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(util.Error{Status: e.Code, ErrCode: 500, Message: e.Message})
			} else {
				return ctx.Status(500).JSON(util.Error{Status: 500, ErrCode: 500, Message: err.Error()})
			}
		},
	})
	initMiddlewares(app)
	return app
}

func Listen(app * fiber.App) error {
	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(404) })
	return app.Listen(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"),  os.Getenv("SERVER_PORT")))
}

var (
	SECRET =os.Getenv("JWT_SECRET_KEY");
)

func RouteInit(app *fiber.App) {

	// jwt middleware
	jwt := middlewares.NewAuthMiddleware(SECRET)

	// api v1 group
	v1 := app.Group("/api/v1")
	{
		// UserController
		user := v1.Group("user")
		{
			user.Post("/login", Login)
			// user.Get("/profile", jwt, UserController.GetProfile)
		}

	}
}