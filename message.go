package quasar_fire_core


func GetMessage(messages ...[]string) string {
	return "este es un mensaje"
	// use reduce to iteratively merge the arrays, using the mergeMessages function
}

func mergeMessages(m1, m2 []string) []string {
	// 1: {"", "", "", "es", "", "mensaje", "", "un", "amigo"},
	// 2: {"este", "", "un", "mensaje", "", "", "amigo"},

	// We get and store the length of the smallest slice, to use it as a stop point of every loop
	lenghtM1 := len(m1)
	lenghtM2 := len(m2)
	var minLength int

	if lenghtM1 <= lenghtM2 {
		minLength = lenghtM1
	} else {
		minLength = lenghtM2
	}
	// Mergear:
	//	- Buscar primera palabra en común en ambos arrays
	//		- 1 solo loop, almacenar uno o dos hash tables de "seen", así hacemos un solo loop
	//		- El loop para cuando:
	//			- Consigue una palabra en común en ambos slices (De acuerdo a la información de la hash table)
	//			- Llega al final del slice más pequeño
	//		- 1 sola iteración, complejidad O(n), evitamos los loops anidados
	//		- El resultado de esto debe ser la posición de la palabra en común en cada slice
	//		- "mensaje"
	//			- Posición en Slice 1: 5
	//			- Posición en Slice 2: 3
	//
	// TODO: Initialize the hash tables
	for i := 0; i <= minLength; i++ {

	}

	//	- Inicializar un puntero en "posición 0" para recorrer ambos slices
	//	- El puntero para el array donde la palabra en común esté más adelante, se calcula en base a un offset que es la diferencia entre la posición mayor y la menor
}
