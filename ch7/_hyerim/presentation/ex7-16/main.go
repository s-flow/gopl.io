package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	eval "gopl.io/ch7/_hyerim/presentation/ex7-14"
)

func main() {
	http.HandleFunc("/calcurator", calcurator)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func printResult(w io.Writer, f float64) {
	fmt.Fprintf(w, "result : %g", f)
}

func calcurator(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	s := r.Form.Get("expr")
	if s == "" {
		http.Error(w, "expr empty", http.StatusBadRequest)
		return
	}
	fmt.Printf("s: %s\n", s)
	expr, err := eval.Parse(s)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}

	var env = eval.Env{}
	for k := range vars {
		v := r.Form.Get(k.String())
		if v == "" {
			http.Error(w, k.String()+" is empty"+err.Error(), http.StatusBadRequest)
		}
		env[k], err = strconv.ParseFloat(v, 64)
		if err != nil {
			http.Error(w, k.String()+" bad value"+err.Error(), http.StatusBadRequest)
		}
	}
	printResult(w, expr.Eval(env))
}

//http://localhost:8000/calcurator?expr=x%2By&x=1&y=10
