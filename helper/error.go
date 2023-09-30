package helper

func PanicIfEror(err error) {
	if err != nil {
		panic(err)

	}
}
