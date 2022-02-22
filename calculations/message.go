package calculations

func GetMessage(messages ...[]string) string {
	return "este es un mensaje"
}

func mergeMessages(m1, m2 []string) {
	// 1: {"", "", "", "es", "", "mensaje", "", "un", "amigo"},
	// 2: {"este", "", "un", "mensaje", "", "", "amigo"},
	// Mergear:
	//	- Buscar primera palabra en común en ambos arrays
	//		- "mensaje"
	//			- Posición en Slice 1: 5
	//			- Posición en Slice 2: 3
	//	- Inicializar un puntero en "posición 0" para recorrer ambos slices
	//		- Calculo de "posición 0": Pendiente de pensar la lógica
}
