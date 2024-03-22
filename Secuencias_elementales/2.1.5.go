//Se dispone de una secuencia de numeros de socios (Sec de enteros) 
//grabar en una secuencia de salida los socios con numero par
ejercicio 2.1.5 
ambiente 
	sec,sal: secuencia de enteros
	v: entero 
Proceso 
	Arr(sec), avz(sec,v)
	crear(sal)
	Mientras NFDS(sec) hacer 
		si v mod 2 = 0 entonces 
			Esc(sal,v)
		fs 
		avz(sec,sal)
	fin_mientras
	cerrar(sal)
Fin_proceso 
