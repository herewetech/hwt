/*
 * Copyright (C) ###__PROJ_AUTHOR__### - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file misc.go
 * @package handler
 * @author ###__PROJ_AUTHOR__###
 * @since ###__TODAY__###
 */

package handler

import (
	"errors"
	"os"
	"svc/handler/response"
	"svc/runtime"
	"svc/utils"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Misc struct {
}

func InitMisc() *Misc {
	h := new(Misc)

	// Load routers
	runtime.Server.All("/", h.index).Name("Index")
	runtime.Server.Get("/routers", h.routers).Name("GetRouters")
	runtime.Server.Get("/swagger.json", h.swaggerJson).Name("GetSwaggerJson")
	runtime.Server.All("/docs/*", swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
	}))

	return h
}

// index

// @Tags Misc
// @Summary Just an empty portal
// @Description Empty JSON envelope
// @ID Index
// @Produce json
// @Success 200 {object} nil
// @Router / [get]
func (h *Misc) index(c *fiber.Ctx) error {
	return c.JSON(utils.WrapResponse(nil))
}

// routers

// @Tags Misc
// @Summary Get HTTP routers
// @Description Route list
// @ID GetRouters
// @Produce json
// @Success 200 {object} nil
// @Router /routers [get]
func (h *Misc) routers(c *fiber.Ctx) error {
	return c.JSON(utils.WrapResponse(runtime.Server.Stack()))
}

func (h *Misc) swaggerJson(c *fiber.Ctx) error {
	content, err := os.ReadFile("./docs/swagger.json")
	if err != nil {
		c.WriteString("{}")

		return err
	}

	c.Write(content)

	return nil
}

/* {{{ *Internal handlers* */
func jwtSuccessHandler(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token parse from context failed")
	}

	claims, ok := u.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("JWT claims type error")
	}

	authType, ok := claims["type"].(string)
	if ok {
		c.Locals("AuthType", authType)
	}

	authEmail, ok := claims["name"].(string)
	if ok {
		c.Locals("AuthEmail", authEmail)
	}

	authID, ok := claims["sub"].(string)
	if ok {
		c.Locals("AuthID", authID)
	}

	return c.Next()
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	resp := utils.WrapResponse(nil)
	resp.Code = response.CodeAuthFailed
	resp.Message = err.Error()
	resp.Status = fiber.StatusUnauthorized

	return c.Status(fiber.StatusUnauthorized).JSON(resp)
}

/* }}} */

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
