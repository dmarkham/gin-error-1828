package admin

import "fmt"

type IndexController struct{}

func (this *IndexController) Index(ctx *string) {
	fmt.Println(&ctx)
}
