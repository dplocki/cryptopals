package main

import "strings"

func splitQueryParamaters(input string) map[string]string {
	result := make(map[string]string)
	input_split := strings.Split(input, "&")

	for _, v := range input_split {
		v_split := strings.Split(v, "=")
		result[v_split[0]] = v_split[1]
	}

	return result
}

func MainSet2Challenge13() {
	println("ECB cut-and-paste")

	input := "foo=bar&baz=qux&zap=zazzle"
	query_param_map := splitQueryParamaters(input)

	println("{")
	for k, v := range query_param_map {
		println("\t", k, ":", v)
	}
	println("}")
}
