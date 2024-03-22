ejercicio 2.1.1 
Ambiente 
	sec: secuencia de caracteres
	v: caracter 
	contador_a: entero 
Proceso 
	arr(sec), avz(sec,v)
	contador_a := 0
	Mientras (sec <> "*") hacer 
		si v = "A" entonces
		contador_a := contador_a + 1 
		avz(sec,v)
	fin_mientras 
	esc("La letra A se repite"contador_a "veces")
fin Proceso



