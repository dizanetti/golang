package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Partner struct {
	CpfCnpj   int64  `json:"cpfCnpj" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
}

func main() {
	srv := gin.Default()

	/* Ping */
	srv.GET("/ping", ping)
	/* Find Partner by cpfCnpj */
	srv.GET("/partner/:cpfCnpj", getPartner)
	/* Find Bonus of Partner by cpfCnpj and bonusId */
	srv.GET("/partner/bonus", getPartnerBonus)
	/* Insert a new partner */
	srv.POST("/partner/insert", insertPartner)
	/**/
	srv.PATCH("/partner/update", updatePartner)
	/**/
	srv.DELETE("/partner/delete/:cpfCnpj", deletePartner)

	/* Start server on port 8096 */
	srv.Run(":8096")
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func getPartner(ctx *gin.Context) {
	cpfCnpj := ctx.Param("cpfCnpj")

	ctx.String(http.StatusOK, "find partner %s", cpfCnpj)
}

func getPartnerBonus(ctx *gin.Context) {
	cpfCnpj := ctx.Query("cpfCnpj")
	bonusId := ctx.Query("bonusId")

	ctx.String(http.StatusOK, "find partner and bonus ID %s %s", cpfCnpj, bonusId)
}

func insertPartner(ctx *gin.Context) {
	var partner Partner
	if err := ctx.ShouldBindJSON(&partner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": partner.CpfCnpj})
}

func deletePartner(ctx *gin.Context) {
	cpfCnpj := ctx.Param("cpfCnpj")

	ctx.String(http.StatusOK, "partner deleted %s", cpfCnpj)
}

func updatePartner(ctx *gin.Context) {
	var partner Partner
	if err := ctx.ShouldBindJSON(&partner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": partner.CpfCnpj})
}
