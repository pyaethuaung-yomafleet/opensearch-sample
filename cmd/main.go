package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
	"github.com/yomafleet/opensearch-client-test/cmd/model"
	opensearchPkg "github.com/yomafleet/opensearch-client-test/pkg/interface/persistence/opensearch"
	"github.com/yomafleet/opensearch-client-test/pkg/usecase"
)

func main() {

	port := "8080"

	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // For testing only. Use certificate for validation.
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin", // For testing only. Don't store credentials in code.
		Password:  "admin",
	})
	if err != nil {

		panic(err)
	}

	opensearchProductRepo := opensearchPkg.NewOpenSearchProductRepository(
		client, "products-brand-model-category-index")
	productUsecase := usecase.NewProductUsecase(opensearchProductRepo)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {

			return ctx.Path() == "/health"
		},
		Format: `{"time":"${time_rfc3339_nano}", "method":"${method}",` +
			`"uri":"${uri}","status":${status},"user_agent":"${user_agent}",` +
			`"latency":${latency},"bytes_out":${bytes_out}}` + "\n",
	}))
	e.POST("/products", func(ctx echo.Context) error {

		product := new(model.ProductCreateModel)
		if err := ctx.Bind(product); err != nil {

			return err
		}

		if err := productUsecase.Create(
			product.ConvertToUsecaseModel()); err != nil {

			return err
		}

		return ctx.JSON(http.StatusCreated, nil)
	})
	e.GET("/products", func(ctx echo.Context) error {

		query := ctx.QueryParam("query")

		usecaseProducts, err := productUsecase.Search(query)
		if err != nil {

			return err
		}

		var products []model.ProductGetModel
		for _, v := range usecaseProducts {

			if m := model.ConvertFromUsecaseModel(&v); m != nil {

				products = append(products, *m)
			}
		}

		return ctx.JSON(http.StatusOK, products)
	})

	fmt.Println("notification service is starting")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
