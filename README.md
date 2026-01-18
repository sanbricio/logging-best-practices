# 🪵 Logging Best Practices en Go

Ejemplo práctico de cómo implementar **logs que realmente sirven en producción** usando Go, Gin y Zap.

## El problema

Cuando tienes múltiples requests concurrentes, los logs básicos son inútiles:

```
ERROR: Database connection timeout
ERROR: Order validation error payment.ID
ERROR: Order can not be processed
```

¿Cuál es de qué request? Imposible saberlo.

## La solución

Añadir un **Trace ID** único a cada request:

```json
{"level":"error","trace_id":"a1b2c3d4","user_id":"100","msg":"database error","error":"connection timeout"}
{"level":"error","trace_id":"e5f6g7h8","product_id":"PROD-1","msg":"payment failed","error":"gateway timeout"}
```

Ahora puedes filtrar por `trace_id` y ver exactamente qué pasó en cada request.

## Stack

- **[Gin](https://github.com/gin-gonic/gin)** - Framework web
- **[Zap](https://github.com/uber-go/zap)** - Logger estructurado de Uber
- **Middleware personalizado** - Genera y propaga el Trace ID

## Estructura

```
.
├── main.go
├── pkg/
│   ├── logger/
│   │   └── logger.go      # Configuración de Zap + helpers
│   └── middleware/
│       └── logging.go     # Middleware que genera el Trace ID, y recuperan ante panics 
└── internal/
    └── handlers/
        └── handlers.go    # Handlers de ejemplo
```

## Ejecutar

```bash
# Desarrollo (logs en consola con colores)
go run main.go

# Producción (logs en JSON)
ENV=production go run main.go
```

## Probar

```bash
# Request simple
curl http://localhost:8080/users/123

# Crear orden
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"product_id": "PROD-1", "quantity": 5}'

# Propagar trace-id desde otro servicio
curl http://localhost:8080/users/123 \
  -H "X-Trace-ID: mi-trace-id-externo"
```

## Conceptos clave

### 1. Trace ID
Identificador único por request. Se genera automáticamente o se propaga si viene en el header `X-Trace-ID`.

### 2. Structured Logging
Logs en formato JSON con campos tipados, no strings concatenados:

```go
// ❌ Mal
log.Printf("Error fetching user %s: %v", userID, err)

// ✅ Bien
logger.Error(ctx, "database error",
    zap.String("user_id", userID),
    zap.Error(err),
)
```

### 3. Contexto
El trace-id viaja en el `context.Context` de Go, disponible en cualquier capa de tu aplicación.

## Video

Este repositorio acompaña al video **"Logs BUENOS vs MALOS en Go"** donde comparo ambas implementaciones lado a lado.

📺 [Ver video](#) https://www.youtube.com/watch?v=AWR3yZtZtgo

## Siguiente nivel

Si tienes múltiples microservicios y necesitas saber **dónde** se va el tiempo (no solo cuánto), el siguiente paso es **OpenTelemetry** para tracing distribuido.

## Licencia

MIT
