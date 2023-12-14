# SOA GoLang 🦫

Proyecto de implementación con GoLang del bus de datos en Service Oriented Architecture, del ramo Arquitectura de Software.

## Environment Variables

Es necesario crear un archivo .env en el directorio principal del proyecto con las siguientes variables, las cuales pueden ser modificadas según gusten.

```dotenv
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=arquisw
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
HIDE_EMPTY_PASSWORD=yes
SOABUS_HOST=soabus
SOABUS_PORT=5000
```

## Docker

Construir docker compose desde el directorio principal del proyecto:

```bash
docker compose build -t arqui
```

Ejecutar servicios:

```bash
docker up
```

Borrar volumen (para reiniciar base de datos):

```bash
docker compose down -v
```

## Probar funcionamiento

```bash
go run ./clients/test/
```