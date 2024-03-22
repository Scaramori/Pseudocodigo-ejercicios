Ejercicio 2.1.8 es 
Ambiente 
	sec,sal: secuencia de enteros 
	v : entero 
	deseados = (0,1,2,3)

	Funcion CarEnt(AuxCarEnt: enteros): caracter
            Segun AuxCarEnt hacer
                =0: CarEnt:= "0"
                =1: CarEnt:= "1"
                =2: CarEnt:= "2"
                =3: CarEnt:= "3"
                =4: CarEnt:= "4"
                =5: CarEnt:= "5"
                =6: CarEnt:= "6"
                =7: CarEnt:= "7"
                =8: CarEnt:= "8"
                =9: CarEnt:= "9"
            FinSegun
    FinFuncion

Proceso 
	arr(sec), avz(sec)

	Mientras NFDS(sec) hacer 
		si (v % 10) en deseados entonces // mod 
			CarEnt(V)
			Esc(sal,V)
			esc(sal,"-")
		fs
		avz(sec,v)
	fin_mientras
	cerrar(sal)
Fin_proceso



