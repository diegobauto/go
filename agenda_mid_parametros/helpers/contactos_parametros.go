package helpers

import (
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/agenda_mid_parametros/models"
)

func ListarContactosParametros() (contactosParametros models.ContactosParametros, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{"funcion": "ListarContactosParametros", "err": err, "status": "500"}
			panic(outputError)
		}
	}()

	var parametros []models.Parametro
	var contactos []models.Contacto

	url := "parametro?query=TipoParametroId__CodigoAbreviacion:C&limit=0"
	if err := GetRequestNew("UrlCrudParametros", url, &parametros); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	urlContactos := "contacto?limit=0"
	if err := GetRequestNew("UrlCrudAgenda", urlContactos, &contactos); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	contactosParametros.Parametros = parametros
	contactosParametros.Contactos = contactos

	return contactosParametros, outputError
}
