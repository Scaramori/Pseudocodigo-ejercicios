Ejercicio 2.1.10 es 
Ambiente 
	sec: secuencia de enteros
	v: enteros 
	inicio_palabra: booleano
	palabras_contador : entero 
	letra_ingresada: caracter
Proceso 
	esc("ingrese la letra por la que desea discriminar")
	leer(letra_ingresada)
	palabras_i := 0
	inicio_palabra := falso 
	arr(sec), avz(sec,v)
	Mientras NFDS(sec) hacer 
		si v= letra_ingresada y inicio_palabra = false
			inicio_palabra := true 
			palabras_contador:= palabras_contador + 1
		fin_si 
		si v = " " entonces 
			inicio_palabra = false 
		fin_si
		avz(sec,v)
	fin_mientras
	cerrar(sec)
	esc( "hay"palabras_i"palabras que inician con" letra_ingresada)
Fin_proceso
