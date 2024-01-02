
# hygen_api_crud
:heavy_check_mark: Check: Repositorio para HYGEN_API_CRUD.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)


### Variables de Entorno
```shell
# parametros de api
HYGEN_API_CRUD_HTTP_PORT=[Puerto de exposición del API]
# paramametros de bd
HYGEN_API_CRUD_PGUSER=[Usuario de BD]
HYGEN_API_CRUD_PGPASS=[Contraseña del usaurio de BD]
HYGEN_API_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
HYGEN_API_CRUD_PGPORT=[Puerto de la BD]
HYGEN_API_CRUD_PGDB=[Nombre de Base de Datos]
HYGEN_API_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con HYGEN_API_CRUD...


### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/hygen_api_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/hygen_api_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
HYGEN_API_CRUD_HTTP_PORT=8080 HYGEN_API_CRUD_SOME_VARIABLE bee run
```

### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/hygen_api_crud

#2. Moverse a la carpeta del repositorio
cd hygen_api_crud

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas
```shell
# En Proceso
```

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI


| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/hygen_api_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/hygen_api_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/hygen_api_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/hygen_api_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/hygen_api_crud/status.svg?ref=refs/heads/master)](https://hubci.portaloas.udistrital.edu.co/udistrital/hygen_api_crud) |

## Esquema de base de datos
* [Esquema de Base de datos](https://drive.google.com/file/d/1TCfrbYWakYtMgJm2iG-gzMKJcqPUVEF2/view?usp=sharing)

## Licencia

This file is part of hygen_api_crud.

hygen_api_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

hygen_api_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with hygen_api_crud. If not, see https://www.gnu.org/licenses/.