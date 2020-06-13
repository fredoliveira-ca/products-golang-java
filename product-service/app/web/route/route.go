package route

import (
	"net/http"

	controller "github.com/fredoliveira-ca/products-golang-java/product-service/app/web/controller"
)

//LoadRoutes is...
func LoadRoutes() {
	http.HandleFunc("/product", controller.ListProducts)
}
