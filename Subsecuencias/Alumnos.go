Accion Alumnos es
AMBIENTE
	sec: secuencia de caracter
	v,comision,aa_ingreso_1,aa_ingreso_2: caracter
	cont_alumnos,i: entero

PROCESO

	ARR (sec)
	AVZ (sec,v)
	MIENTRAS nofds (sec) hacer 
		comision := v
		cont_alumnos := 0
		AVZ (sec,v)
		PARA i := 1 a 12 hacer // avanzo toda la info de teoria, practica y horarios
			AVZ(sec,v)
		finPara
		ESC("Se listan a continuacion el nombre y apellido de los estudiantes que ingresaron en el anio 2024: ")
		MIENTRAS v <> "%" hacer // trato alumnos
			cont_alumnos := cont_alumnos + 1 // cuento al alumno
			PARA i := 1 a 4 hacer // salteo dia y mes de ingreso, ya que me interesa el año
				AVZ(sec,v)
				finPara // Al salir del ciclo estoy en el primer caracter del anio
				aa_ingreso_1 := v  // resguardo primer digito del año
				AVZ(sec,v)
				aa_ingreso_2 := v  // resguardo segundo digito del año
				AVZ(sec,v)   // me posiciono en la primera letra del nombre
				SI (aa_ingreso_1 = "2") ^ (aa_ingreso_2 = "4") ENTONCES 
					MIENTRAS v <> "!" hacer // listarlo al alumno
						ESC (v)
						AVZ (sec,v)
					finMientras	
				SINO 
					MIENTRAS v <> "!" hacer 
						AVZ (sec,v)
					finMientras
				fin si
	
				AVZ (sec,v) // salteo el "!"
				
				PARA i := 1 a 4 HACER //salteo legajo
					avz (sec,v)
				finPara 
				
				// aca puedo estar en otro alumno, o directamente en un %, que significa que ya no hay mas alumnos de esa comisión
	
			FinMientras // V <> “%”
	
			// al salir de este mientras, estoy en una nueva comision, o tambien podria estar en un fds 
	
			// informar cantidad de alumnos por comision
	
			ESC("Comision:", comision, "Cantidad de alumnos:", cont_alumnos)
	
		finMientras // NFDS(sec)
	
		cerrar (sec)
	finAccion
	
	
	