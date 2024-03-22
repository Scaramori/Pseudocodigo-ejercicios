Ejercico 2.1.11 es 
Ambiente 
	sec: secuencia de caracteres 
	v: caracter
	palabras_contador: entero 
	palabras_long: entero
Proceso
	palabras_contador := 0
	palabras_long := 0 

	Mientras NFDS(sec) hacer 
		Si v = " " entonces 
			si palabras_long = 4 entonces 
				palabras_contador += 1
			fs	
			palabras_long := 0
		fs
		si v != " " entonces 
			palabras_long += 1 
		fs
		avz(sec,v)
	fin_mientras
	cerrar(sec)
	esc("Hay " + palabras_contador + " palabras de 4 caracteres")
Fin_proceso