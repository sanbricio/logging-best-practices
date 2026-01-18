#!/bin/bash

# =====================================================
# Script para mostrar el PROBLEMA sin trace-id
# =====================================================
# Simula tráfico real: varios usuarios haciendo cosas
# a la vez. Algunos fallarán aleatoriamente.
#
# La pregunta es: cuando un cliente llama diciendo
# "mi pedido falló", ¿cómo encuentras SU error entre
# todos estos logs?

echo "=========================================="
echo "Simulando tráfico real (10 requests en paralelo)..."
echo "Algunos fallarán aleatoriamente."
echo "=========================================="
echo ""
echo "Mira los logs del servidor..."
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

# Esperar a que terminen
wait

echo ""
echo "=========================================="
echo "PREGUNTAS:"
echo "=========================================="
echo "1. ¿Cuántas requests fallaron?"
echo "2. ¿Qué usuario tuvo el error de base de datos?"
echo "3. Si un cliente te da su order_id, ¿puedes"
echo "   encontrar todos los logs de su petición?"
echo ""
echo "SPOILER: Es imposible saberlo loggando de esta manera"
echo "=========================================="