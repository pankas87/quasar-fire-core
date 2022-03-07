package quasar_fire_core

import "fmt"

func GetMessage(messages ...[]string) string {
	return "este es un mensaje"
	// use reduce to iteratively merge the arrays, using the mergeMessages function
}

// TODO: Implement different test cases for this function alone
func MergeMessages(m1, m2 []string) []string {
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
	//		- Palabra: "mensaje"
	//			- Posición en Slice 1: 5
	//			- Posición en Slice 2: 3
	//		- Diff entre ambas ocurrencias de la palabra:
	//			- (5-3) = 2
	//
	//
	commonWordsTracker := make(map[string][]int)

	for i := 0; i < minLength; i++ {
		// Get the words
		words := []string{
			m1[i],
			m2[i],
		}

		for position, word := range words {
			// If the word is "", continue
			if word == "" {
				continue
			}

			// If the word is not in the tracker map, add it
			if _, found := commonWordsTracker[word]; !found {
				commonWordsTracker[word] = []int{-1,-1}
			}


			// If the value at current #position for the current #word in the tracker map is -1
			//		- Store the current #i value at #position for #word in the tracker map
			if val := commonWordsTracker[word][position]; val == -1 {
				commonWordsTracker[word][position] = i
			}
		}
	}

	// TODO: Iterate the tracker map and use the first matching pair of words found
	//var matchingWord string
	for val, key := range commonWordsTracker {
		fmt.Printf("val: %+v\n", val)
		fmt.Printf("key: %+v\n", key)
	}

	//	- Inicializar un puntero en "posición 0" para recorrer ambos slices
	//	- El puntero para el array donde la palabra en común esté más adelante, se calcula en base a un offset que es la diferencia entre la posición mayor y la menor
	return []string{"este", "es", "un", "mensaje"}
}
