# Simple CRUD con GO
Este proyecto consiste en una API REST para la gestion de contactos

Instalación de GORM y Gorilla Mux

```bash
export PATH=$PATH:/usr/local/go/bin
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go mod init github.com/rbartolome/go_demo_fallabella
go run main.go
```
Estructura del proyecto
```
.
├── controllers/
  ├── ...
├── models/
  └── ...
├── utils/
  └── ...
└── main.go
```
Los datos de los contactos serán:
* Nombre
* Edad
* Teléfono
* Dirección
* E-mail
* Descripción