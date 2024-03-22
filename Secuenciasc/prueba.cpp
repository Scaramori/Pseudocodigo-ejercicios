#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main() {
    ifstream archivo("secuenciabc.txt");
    if (!archivo.is_open()) {
        cout << "No se pudo abrir el archivo secuenciabc.txt" << endl;
        return 1;
    }

    string sec;
    char v;
    int contador_a = 0;

    // Leer la secuencia desde el archivo
    archivo >> sec;

    // Mientras la secuencia no sea '*', continuar
    while (sec != "*") {
        // Leer el próximo caracter
        v = sec[0];

        // Si el caracter es 'A' (mayúscula o minúscula), incrementar el contador
        if (v == 'A' || v == 'a') {
            contador_a++;
        }

        // Avanzar la secuencia eliminando el primer caracter
        sec = sec.substr(1);

        // Leer la próxima secuencia desde el archivo
        archivo >> sec;
    }

    // Mostrar el resultado
    cout << "La letra 'A' se repite " << contador_a << " veces." << endl;

    archivo.close();

    return 0;
}
