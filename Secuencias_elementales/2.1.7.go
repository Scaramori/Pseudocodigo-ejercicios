//11 digitos cuit 
Ejercicio 2.1.7 es 
Ambiente 
	sec,sal: secuencia de enteros 
	v : entero 
	deseados = (0,1,2,3)
Proceso 
	arr(sec), avz(sec)

	Mientras NFDS(sec) hacer 
		si (v % 10) en deseados entonces // mod 
			Esc(sal,V)
		fs
		avz(sec,v)
	fin_mientras
	cerrar(sal)
Fin_proceso