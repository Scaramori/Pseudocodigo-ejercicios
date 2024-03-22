/*Suponemos un orden,2 masc siendo mayor y menor y luego 2 fem siendo mayor y menor 
Ejemplo 18|8|22|16:
(18 y 8 = masc, siendo 18 mayor y 8 menor )
(22 y 16 = fem, siendo 22 mayor y 16 menor ) */
Ejercicio 2.1.6 es 
Ambiente
	sec: secuencia de enteros 
	v: enteros 
	total: entero
	masc_mayor,masc_menor: entero 
	fem_mayor,fem_menor: entero
	masc_tot, fem_tot: entero 
	procedimiento inicializar() es 
		masc_mayor,masc_menor:= 0
		fem_mayor,fem_menor:= 0
		total := 0 
	fin_procedimiento 
Proceso 
	arr(sec), avz(sec,v)
	inicializar()

	Mientras NFDS(sec) hacer 
		Para i = 1 hasta 4 hacer 
			segun i hacer 
				=1 : masc_mayor += v 
				=2 : masc_menor += v 
				=3 : fem_mayor += v 
				=4 : fem_menor += V
			fin segun
			avz(sec,v)
		fin_para 
	fin_mientras

	masc_tot := masc_mayor + masc_menor 
	fem_tot := fem_mayor + fem_menor
	mayor_tot := masc_mayor + fem_mayor
	menor_tot := masc_menor + fem_menor
	total :=  masc_tot + fem_tot

	esc("la poblacion total es:" total)
	esc("el porcentaje masculino es de :" masc_tot * 100 / total)
	esc("el porcentaje femenino es de :" fem_tot * 100 / total)
	esc("el porcentaje mayor de edad es:" mayor_tot * 100 / total)
	esc("el porcentaje menor de edad es de:" menor_tot * 100/ total)
Fin_proceso
	
	

