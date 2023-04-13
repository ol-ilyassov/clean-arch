package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"ol-ilyassov/clean_arch/services/contact/internal/useCase"
)

func init() {
	viper.SetDefault("HTTP_PORT", 34777)
	viper.SetDefault("HTTP_HOST", "127.0.0.1")
	viper.SetDefault("IS_PRODUCTION", "false")
}

// @title slurm contact service on clean architecture
// @version 1.0
// @description contact service on clean architecture
// @license.name kolyadkons

// @contact.name API Support
// @contact.email kolyadkons@gmail.com

// @BasePath /

type Delivery struct {
	ucContact useCase.Contact
	ucGroup   useCase.Group
	router    *gin.Engine

	options Options
}

type Options struct{} // optional. Further Dev: requests timeout, etc...

func New(ucContact useCase.Contact, ucGroup useCase.Group, options Options) *Delivery {
	// Here: could be check nil values for ucContact, and ucGroup.

	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}
	d.SetOptions(options)
	d.router = d.initRouter()
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf("%s:%d", viper.GetString("HTTP_HOST"), uint16(viper.GetUint("HTTP_PORT"))))
}

// Write authorization method
func checkAuth(c *gin.Context) {
	c.Next()
}
