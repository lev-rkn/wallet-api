package handlers

import (
	"encoding/json"
	"net/http"
	"wallet-api/internal/lib/types"
	"wallet-api/internal/middlewares"
	"wallet-api/internal/models"
	"wallet-api/internal/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	walletRouter := router.Group("/wallet")
	walletRouter.Use(middlewares.AuthMiddleware())

	walletRouter.GET("/check-account", checkExistsHandler)
	walletRouter.POST("/deposit", depositHandler)
	walletRouter.GET("/month-history", getMonthHistoryHandler)
	walletRouter.GET("/balance", getWalletBalanceHandler)

	return router
}

func checkExistsHandler(c *gin.Context) {
	userId := c.GetInt(types.KeyUserId)

	wallet, err := storage.FoundWallet(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": types.ErrWalletNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func depositHandler(c *gin.Context) {
	// создаем транзакцию, в которую кладем денежный объем операции
	transaction := &models.Transaction{}
	err := json.NewDecoder(c.Request.Body).Decode(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt(types.KeyUserId)

	wallet, err := storage.Deposit(userId, transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}

func getMonthHistoryHandler(c *gin.Context) {
	userId := c.GetInt(types.KeyUserId)

	totalCount, totalAmount := storage.GetMonthHistory(userId)

	c.JSON(http.StatusOK, gin.H{
		"count":  totalCount,
		"amount": totalAmount,
	})
}

func getWalletBalanceHandler(c *gin.Context) {
	userId := c.GetInt(types.KeyUserId)

	wallet, err := storage.FoundWallet(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": types.ErrWalletNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": wallet.Balance,
	})
}
