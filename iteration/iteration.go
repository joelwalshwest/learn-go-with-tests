package iteration

func Repeat(charactor string, count int64) string {
	var repeated string
	for i := 0; i < int(count); i++ {
		repeated += charactor
	}

	return repeated
}
