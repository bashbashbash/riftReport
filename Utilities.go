package main

func printIferr(err error) {
	if err != nil {
		print(err)
	}
}
