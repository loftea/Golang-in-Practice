package main

import "net/http"

const form = `
<html><body>
<form action="/" method="post" name="bar">
	<h1>Statistics</h1>
	<>Computes basic statistics for a given list of numbers</>
	<>Numbers (comma or space-separated):<>
	<input type="text" name="in" />
	<input type="submit" value="Calculate"/>
</form>
</body></html>
`

func main() {
	http.HandleFunc("/calc", CalcServer)
	if err := http.ListenAndServe("9001", nil); err != nil {
		panic(err)
	}
}

func Calcserver(w http.ResponseWriter, r *http.Request) {

}
