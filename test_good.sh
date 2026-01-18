#!/bin/bash

# =====================================================
# Script de pruebas - Versión BUENA con trace-id
# =====================================================

echo "=========================================="
echo "Simulando el MISMO tráfico pero con trace-id..."
echo "=========================================="
echo ""

# Simular varios usuarios concurrentes
for i in {1..3}; do
    curl -s http://localhost:8080/users/$i > /dev/null &
done

for i in {1..3}; do
    curl -s -X POST http://localhost:8080/orders \
        -H "Content-Type: application/json" \
        -d "{\"product_id\": \"PROD-$i\", \"quantity\": $i}" > /dev/null &
done

for i in {1..2}; do
    curl -s http://localhost:8080/products > /dev/null &
done

curl -s http://localhost:8080/users/99 > /dev/null &
curl -s http://localhost:8080/users/100 > /dev/null &

wait

echo ""
echo "=========================================="
echo "AHORA SÍ PUEDES:"
echo "=========================================="
echo "1. Filtrar por trace_id para ver UNA request completa"
echo "2. Saber qué producto tuvo el error de pago"
echo "3. Saber qué usuario tuvo el timeout de DB"
echo "4. Dar al cliente su trace_id del header X-Trace-ID"
echo "=========================================="
echo ""
echo "Ejemplo: grep 'trace_id=abc123' en tus logs"
echo "=========================================="