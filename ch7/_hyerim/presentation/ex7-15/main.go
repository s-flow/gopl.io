package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")

		if !scan.Scan() {
			fmt.Println("Scan end")
			break
		}

		// parse
		s := strings.TrimSpace(scan.Text())
		expr, err := eval.Parse(s)
		if err != nil {
			fmt.Errorf("parse error %s\n", err.Error())
		}

		// setup enviroment
		vars := make(map[eval.Var]bool)
		if err := expr.Check(vars); err != nil {
			fmt.Errorf("expression check error %s\n", err.Error())
		}

		// read variable
		env, err := readEnv(expr, vars)
		if err != nil {
			fmt.Printf("read env error %v\n", err)
		}

		fmt.Printf("result: %g\n", expr.Eval(env))
	}

	if err := scan.Err(); err != nil {
		fmt.Println("read error: %v\n", err)
	}
}

func readEnv(expr eval.Expr, vars map[eval.Var]bool) (eval.Env, error) {
	var env = eval.Env{}
	scan := bufio.NewScanner(os.Stdin)
	for k := range vars {
		for {
			fmt.Printf("\t %s  = ", k)
			if !scan.Scan() {
				return nil, fmt.Errorf("incomplete enviroment")
			}
			v, err := strconv.ParseFloat(scan.Text(), 64)
			if err != nil {
				fmt.Printf("invalid value!! Enter a valid value")
				continue
			}

			env[k] = v
			break
		}
	}
	return env, nil
}
