package controller

import (
	"go-api/usecase"
	"net/http"
    "strings"
    "go-api/model" 
    "strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUserCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUserCase: usecase,
	}
}

    //GET

func (p *productController) GetProducts(ctx *gin.Context) {
    products, err := p.productUserCase.GetProducts()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return // Importante adicionar return para não continuar
    }

    ctx.JSON(http.StatusOK, products)
}


//delete

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id") // Assumindo rota como /products/:id

	err := p.productUserCase.DeleteProduct(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete product",
			"details": err.Error(),
		})
		return
	}

	// resposta do JSON para o delete
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Produto deletado com sucesso",
		
	})

    // http://localhost:8000/products/ID (ID ira ser o numero produto desejado para deletar)
}


    //---POST---
func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.productUserCase.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar o produto", "causa": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Produto criado com sucesso",
		"ID":      product.ID,
		"name":    product.Name,
		"price":   product.Price,
		"estoque": product.Estoque,
		"categoria": product.Categoria,
		"descrição": product.Descricao,
	})
}



    //---update---

func(p *productController)PutProduct(ctx *gin.Context){
    var product model.Product

    
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	product.ID, _ = strconv.Atoi(id) // converte string para int

	err := p.productUserCase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}


ctx.JSON(http.StatusCreated, gin.H{
    "ID": product.ID,
     "name":    product.Name,
      "price":   product.Price,
    }) //"productID": product.ID,
}


