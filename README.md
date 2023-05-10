# godesde0
Repositorio de todo el curso de Go
## Crear un proyecto Go
- Inicializar
    ```
    $ go mod init
    ```
- Crear main.go (Hola Mundo)
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hola Mundo")
    }
    ```
- Ejecutar el proyecto
    ```
    $ go run .
    Hola Mundo
    ```
- Compilar el proyecto (Esto generará un ejecutable llamado godesde0)
    ```
    $ go build main.go  // Esto Generará un ejecutable llamado: main
    $ go build .        //                                    : godesde0
    ```
- Ejecutar el proyecto desde el ejecutable
    ```
    $ ./godesde0
    Hola Mundo
    ```

