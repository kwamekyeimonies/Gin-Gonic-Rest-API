package helper

func Error_Log(err error) {

	if err != nil {
		// log.Fatal(err.Error())
		// return
		panic(err)
	}
}
