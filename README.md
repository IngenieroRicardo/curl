# CURL
Una biblioteca ligera para hacer peticiones web en C.  
Compilada usando: `go build -o curl.dll -buildmode=c-shared curl.go`

---

### 📥 Descargar la librería

| Linux | Windows |
| --- | --- |
| `wget https://github.com/IngenieroRicardo/curl/releases/download/1.0/curl.so` | `Invoke-WebRequest https://github.com/IngenieroRicardo/curl/releases/download/1.0/curl.dll -Outcurl ./curl.dll` |
| `wget https://github.com/IngenieroRicardo/curl/releases/download/1.0/curl.h` | `Invoke-WebRequest https://github.com/IngenieroRicardo/curl/releases/download/1.0/curl.h -Outcurl ./curl.h` |

---

### 🛠️ Compilar

| Linux | Windows |
| --- | --- |
| `gcc -o main.bin main.c ./curl.so` | `gcc -o main.exe main.c ./curl.dll` |
| `x86_64-w64-mingw32-gcc -o main.exe main.c ./curl.dll` |  |

---

### 🧪 Ejemplo de escritura y lectura

```c
#include <stdio.h>
#include "curl.h"

int main() {
    char* header = malloc(250);

    sprintf(header, "%s\n%s", 
            Header("Content-Type", "application/json"), 
            HeaderAuthBasic("admin", "s3cr3t")
    );

    // Hacer petición GET con los dos headers
    char* resp = Get("http://localhost:8123/api", header, "{ \"key\": \"value\" }");
    
    if (resp != NULL) {
        printf("Response:\n%s\n", resp);
        free(resp);
    }
    
    // Liberar memoria de los headers individuales
    free(header);
    
    return 0;
}
```


---

## 📚 Documentación de la API

#### Crear Header
- `char* Header(char* key, char* value)`
- `char* HeaderAuthToken(char* token)`
- `char* HeaderAuthBasic(char* user, char* pass)`

#### Hacer Peticiones Web
- `char* Get(char* url, char* headers, char* body)`
- `char* Post(char* url, char* headers, char* body)`
- `char* Put(char* url, char* headers, char* body)`
- `char* Patch(char* url, char* headers, char* body)`
- `char* Delete(char* url, char* headers, char* body)`
- `char* Head(char* url, char* headers, char* body)`
- `char* Options(char* url, char* headers, char* body)`


