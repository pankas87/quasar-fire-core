package quasar_fire_core

import (
	"errors"
	"fmt"
	"strings"
)

func GetMessage(messages ...[]string) (string, error) {
	// TODO: Test this
	if len(messages) < 2 {
		return "", errors.New("You need at least two messages to decode")
	}

	accMessage := messages[0]

	for i := 1; i < len(messages) - 1; i++ {
		mergedMessage, err := MergeMessages(accMessage, messages[i])

		if err != nil {
			return "", err
		}

		accMessage = mergedMessage
	}

	return strings.Join(accMessage[:], " "), nil
}

// Merge:
//	- Search for the first common word in both slices of strings
//	- The loop stops when:
//		- It find a common word in both slices (By looking at the info stored in the hash map)
//		- It gets to the end of the slice
//		- 1 iteration per message, O(n+m), we prevent nested loops
//	- With the common word indexes:
//		Example 1:
//			- Values:
//				m1: {"", "", "", "es", "", "mensaje", "", "un", "amigo"},
// 				m2: {"este", "", "un", "mensaje", "", "", "amigo"},
//
//				m1Index: 0
//              m2Index: m1Index - offest = 0 - 2
//
//			- Matching Word:
//				- "mensaje"
//				- Position at Slice 1: 5
//				- Position at Slice 2: 3
//				- Diff between both matches positions:
//					- (5-3) = 2
//					- (3-5) = -2
// 		Example 2:
//			- Values:
// 				m1: {"este", "es", "un", "mensaje", "", "un", "amigo"},
// 				m2: {"", "este", "", "un", "mensaje", "de", "", "amigo"},
//			- Matching Word:
//				- "este"
//				- Position at Slice 1: 0
//				- Position at Slice 2: 1
//				- Diff between both matches positions:
//					- (0-1) = -1
//					- (1-0) = 0
//		Example 3:
//			- Values:
//				m1: {"", "este", "es", "un", "mensaje", "", "un", "amigo"},
//				m2: {"", "este", "", "", "mensaje", "de", "", ""},
//			- Matching Word:
//				- ""
//				- Position at Slice 1: 1
//				- Position at Slice 2: 1
//				- Diff between both matches positions:
//					- (1-1) = 0
//					- (1-1) = 0
// TODO: Implement different test cases for this function alone
func MergeMessages(messages ...[]string) ([]string, error) {
	var result []string

	if len(messages) != 2 {
		// TODO: Add this test case
		return nil, errors.New("MergeMessages expect exactly two messages to merge")
	}

	commonWordsTracker := make(map[string][]int)
	var commonWordsList []string
	matchingWord := ""

	for mIndex, message := range messages {
		for wIndex, word := range message {
			// If the word is "", continue
			if word == "" {
				continue
			}

			// If the word is not in the tracker map, add it
			if _, found := commonWordsTracker[word]; !found {
				commonWordsTracker[word] = []int{-1,-1}
				commonWordsList = append(commonWordsList, word)
			}

			// If the value at current #mIndex for the current #word in the tracker map is -1
			//	- Store the current #wIndex value at #mIndex for #word in the tracker map
			if val := commonWordsTracker[word][mIndex]; val == -1 {
				commonWordsTracker[word][mIndex] = wIndex
			}

			// If the recently added wIndex completes a matching pair at #word, break the loop
			if commonWordsTracker[word][0] != -1 && commonWordsTracker[word][1] != -1 {
				matchingWord = word
				break
			}
		}
		if matchingWord != "" {
			break
		}
	}

	fmt.Printf("Matching Word: %s\n", matchingWord)
	fmt.Printf("Val: %+v\n", commonWordsTracker[matchingWord])

	// If a matching pair is not found
	if matchingWord == "" {
		return nil, errors.New("A matching pair of words was not found. A matching pair is necessary to calculate the offset")
	} else {
		commonWordPositions := commonWordsTracker[matchingWord]
		offset := commonWordPositions[0] - commonWordPositions[1]
		nonEmptyWordFound := false
		fmt.Printf("Offset: %d\n", offset)

		fmt.Println("Entering the last loop (Hopefully)")
		for m1Index, w1 := range messages[0] {
			m2Index := m1Index - offset
			if m2Index < 0 || m2Index >= len(messages[1]) {
				continue
			}

			w2 := messages[1][m2Index]

			newWord := ""

			// If w1 is empty but w2 is not
			if w1 == "" && w2 != "" {
				newWord = w2
			}

			// If w2 is empty but w1 is not
			if w1 != "" && w2 == "" {
				newWord = w1
			}

			// If w1 and w2 are the same
			if w1 == w2 {
				newWord = w1
			}

			if !nonEmptyWordFound && newWord != ""  {
				nonEmptyWordFound = true
			}

			result = append(result, newWord)


			fmt.Printf("m1Index: %d, m2Index: %d\n", m1Index, m2Index)
			fmt.Printf("w1: %s, w2: %s\n\n", w1, w2)
		}
	}
	return result, nil
}
