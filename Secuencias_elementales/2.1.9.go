Ejercicio 2.1.9 es 
Ambiente 
	sec: secuencia de enteros
	v: enteros 
	inicio_palabra: booleano
	palabras_i : entero 
Proceso 
	palabras_i := 0
	inicio_palabra := falso 
	arr(sec), avz(sec,v)
	Mientras NFDS(sec) hacer 
		si v= "i" y inicio_palabra = false
			inicio_palabra := true 
			palabras_i:= palabras_i + 1
		fin_si 
		si v = " " entonces 
			inicio_palabra = false 
		fin_si
		avz(sec,v)
	fin_mientras
	cerrar(sec)
	esc("" palabras_i"palabras que inician con i")
Fin_proceso
