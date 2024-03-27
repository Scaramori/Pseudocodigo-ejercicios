/* Dada una secuencia de caracteres que poseen las ventas

la fecha tiene 8 caracteres y un caracter con las unidades vendida
Ejemplo
|1|1|0|3|2|0|2|2|6| siendo el ultimo 6 las unidades

Fila1 sec salida de enteros , con todas las und vendidad, en meses pares
Algoritmo ejercicio1 es */
Ambiente
	sec: secuencia de caracteres
	sal: sec de enteros 
	v: caracter 

	Funcion CarEnt(auxcarent: enteros): enteros 
		segun auxcarent hacer 
			="0" CarEnt = 0
			="1" CarEnt = 1
			="2" CarEnt = 2
			="3" CarEnt = 3
			="4" CarEnt = 4
			="5" CarEnt = 5
			="6" CarEnt = 6
			="7" CarEnt = 7
			="8" CarEnt = 1
			="9" CarEnt = 1
			..
		fin_segun
	fin_funcion

	mes1

Proceso
	arr(sec), avz(sec,v)
	crear(sal)

	Mientras NFDS(sec) hacer 
		para i= 1 hasta 3 hacer
			avz(sec);
		fin_para
		mes1:= CarEnt(v)
		si (mes1 mod 2 = 0) entonces 
			para i = 1 hasta 4 hacer 
				avz(sec)
			fp 
			unidadesvend:= CarEnt(v) 
			esc(sal,unidadesvend)
		sino 
			para i = hasta 5 hacer 
				avz(sec,v)
		fin_si
		avz(sec,v)
	fin_mientras
	cerrar(sal)
fin_proc