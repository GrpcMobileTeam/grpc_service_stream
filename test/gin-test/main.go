package main

import (
	"github.com/CBDlkl/gin"
	"gin-proto-test/example"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/test", func(context *gin.Context) {
		var input example.HouseContractListInput
		context.Bind(&input)

		var contracts [] *example.HouseContractListDto
		contracts = append(contracts, new(example.HouseContractListDto))
		contracts = append(contracts, new(example.HouseContractListDto))

		output := &example.HouseContractListOutput{TotalRow: 2, ContractListDtos: contracts}

		context.AutoReturn(200, output)
	})

	return router
}

func main() {
	r := setupRouter()
	r.Run()
}
