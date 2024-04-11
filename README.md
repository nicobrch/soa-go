# SOA Go

Esta repo tiene una prueba de lo que sería el proyecto de SOA utilizando Golang.

## Instrucciones de Uso

En primer lugar, se debe iniciar el contenedor del bus de datos, con:

```bash
docker compose build
docker compose up
```

Luego, se puede ejecutar cada uno de los servicios y clientes, utilizando:

```bash
go run <archivo.go>
```