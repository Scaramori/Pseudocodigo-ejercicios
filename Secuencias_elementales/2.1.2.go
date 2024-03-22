Ejercico 2.1.2 es 
Ambiente 
	sec: secuencia de caracteres
	v: caracter
	vocales = ("a","e","i","o","u")
	consonates: entero 
Proceso 
	consonates := 0 
	arr(sec), avz(sec,v)

	Repetir 
		Si v NO EN vocales hacer 
			consonantes := consonantes + 1
		fs 
		avz(sec,v)
	Hasta que V = Z
	cerrar(sec) 
	esc("en la secuencia hay" consonates "consonates")
Fin Proceso