package api

import (
	"database/sql"
	db "my-bank-app/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR NGN"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, account)
}
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	// This pulls the ID from the URL (e.g., /accounts/1)
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// This calls your sqlc store to find the account in Postgres
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		// If the ID doesn't exist in the database
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Success! Send the account data back
	ctx.JSON(http.StatusOK, account)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
