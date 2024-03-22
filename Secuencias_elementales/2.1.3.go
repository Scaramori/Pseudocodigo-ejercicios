Ejercico 2.1.3 es 
Ambiente 
	sec,sal: secuencia de caracter
	v: caracter 
Proceso 
	arr(sec), avz(sec,v)
	crear(sal)
	mientras NFDS(sec) hacer 
		Si v <> "$" entonces 
			esc(sal,v)
		fs
	fin_mientras
	cerrar(sal)
fin proceso 