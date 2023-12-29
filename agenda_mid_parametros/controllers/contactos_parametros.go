package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/agenda_mid_parametros/helpers"
)

type ContactosParametrosController struct {
	beego.Controller
}

func (c *ContactosParametrosController) URLMapping() {

}

// GetAll ...
// @Title Get All
// @Description get Obtener agenda y parametros
// @Success 200 {object} []models.AgendaParametros
// @Failure 400 bad request
// @Failure 500 Internal server error
// @router / [get]
func (c *ContactosParametrosController) GetAll() {
	defer helpers.ErrorController(c.Controller, "ContactosParametrosController")

	if v, err := helpers.ListarContactosParametros(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Listado consultado con exito", "Data": v}
	} else {
		panic(err)
	}
	c.ServeJSON()
}
