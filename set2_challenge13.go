package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func fromQueryParameters(input string) map[string]string {
	result := make(map[string]string)
	input_split := strings.Split(input, "&")

	for _, value := range input_split {
		valueSplit := strings.Split(value, "=")
		result[valueSplit[0]] = valueSplit[1]
	}

	return result
}

func toQueryParameters(input map[string]string) string {
	result := strings.Builder{}
	index := 0

	for key, value := range input {
		result.WriteString(key)
		result.WriteRune('=')
		result.WriteString(value)
		index++

		if index != len(input) {
			result.WriteRune('&')
		}
	}

	return result.String()
}

func profileFor(email string) string {
	result := make(map[string]string)

	result["email"] = email
	result["uid"] = strconv.Itoa(rand.Int())
	result["role"] = "user"

	return toQueryParameters(result)
}

func MainSet2Challenge13() {
	println("ECB cut-and-paste")

	input := "foo=bar&baz=qux&zap=zazzle"
	query_param_map := fromQueryParameters(input)

	println("{")
	for k, v := range query_param_map {
		println("\t", k, ":", v)
	}
	println("}")

	println(toQueryParameters(query_param_map))

	println(profileFor("foo@bar.com"))
}
