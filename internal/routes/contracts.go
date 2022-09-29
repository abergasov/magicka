package routes

import "github.com/gofiber/fiber/v2"

func (a *AppRouter) getBalanceContracts(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"DAI":  "0x6b175474e89094c44da98b954eedeac495271d0f",
		"USDC": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	})
}
