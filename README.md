# Lector de ADN para saber si eres mutante

## Funcionamiento
- Para saber si eres mutante se encuentrarán más de una secuencia de cuatro letras
iguales, de forma oblicua, horizontal o vertical.

- La aplicación funciona de la siguiente manera:

1. Deberás hacer un POST a 
```
ec2-18-118-132-37.us-east-2.compute.amazonaws.com:3000/mutant 
```
con un json que contenga la secuencia de ADN a revisar.

2. Puedes ejecutar el siguiente comando 
```
docker-compose -f docker-compose.yaml up
```
3. También puedes clonar el repositorio, luego:
```
go run .
```

## Respuestas
- La aplicacionr recibe la secuencia enviada y:
    - Responde con un header "HTTP 200-OK", estado que también retornará si al revisar secuencia del humano se determina que es *Mutante*.
    - Responde con un header "UNVALID DNA STRING SEQUENCE", y estado "HTTP 403-FORBIDDEN" si alguna letra secuencia de ADN es diferente a A C T G.
    - Responde con un header "HTTP 403-FORBIDDEN", estado que también retornará si al revisar secuencia del humano se determina que NO es *Mutante*.

