package server

func isMutant(sample []string) bool {
	count := 0
	var dp int = 0
	var ds int = 0
	for i := 0; i < len(sample); i++ {
		h := 0
		v := 0
		for j := 0; j < len(sample)-1; j++ {

			// Chequear horizontal
			if sample[i][j] == sample[i][j+1] {
				h = h + 1
			} else {
				h = 0
			}
			// Chequear vertical
			if sample[j][i] == sample[j+1][i] {
				v = v + 1
			} else {
				v = 0
			}
			//Chequear diagonal principal
			if i == j {
				if sample[i][j] == sample[i+1][j+1] {
					dp = dp + 1
				}
			}
			//Chequear diagonal secundaria
			if (i + j) == 5 {
				if sample[i-1][j+1] == sample[i][j] {
					ds = ds + 1
				}
			}

			if h == 3 {
				count = count + 1
			}
			if v == 3 {
				count = count + 1
			}
		}

	}

	if dp == 3 {
		count = count + 1
	}
	if ds == 3 {
		count = count + 1
	}
	if count > 1 {
		return true
	}

	return false
}
