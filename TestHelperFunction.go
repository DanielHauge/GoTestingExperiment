package GoTestingExperiment

func contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
