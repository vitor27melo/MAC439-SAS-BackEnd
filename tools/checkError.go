package tools

func CheckError(err error) {
	if err != nil {
		print(err)
	}
}
