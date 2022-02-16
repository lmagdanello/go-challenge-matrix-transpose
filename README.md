# Web Service challenge: Matriz Transformation

Recebendo como input uma matriz através de um arquivo csv, devemos utilizar o webservice e suas rotas para realizar transformações e transposições.

---

E.g.:

```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 

2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 

3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
  
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 

4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 
---

O arquivo de input pode ter qualquer dimensão, porém o número de células e colunas deve ser igual (quadrado). Todos os valores devem ser inteiros e não deve haver um cabeçalho no arquivo.

O código pode ser testado com o arquivo disponível em: examples/matrix.csv.

---

## Como executar?

Iniciar o web server:

```
go run .
```

Realizar o request:
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```
