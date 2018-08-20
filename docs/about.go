package docs

/*
Há várias outras funcionalidades na linguagem, como suporte para programação
concorrente, uma vasta biblioteca padrão e ferramentas para trabalhar com
repositórios git, mas aqui temos tudo que precisaremos para lidar com
o desafio de hoje!
*/

// Apenas o clássico: some dois números
func Some(a int, b int) int {
    return a + b
}

// Implementação recursiva da função de Ackermann
func Ackermann(m, n int) int {
    var r int

    if m == 0 {
        r = n+1
    } else if (m > 0) && (n == 0) {
        r = Ackermann(m-1, 1)
    } else {
        r = Ackermann(m-1, Ackermann(m, n-1))
    }

    return r
}

// Descobre o valor mínimo em um vetor de inteiros
// Note que esta é a única função que não é exportada
func min(x []int) int {
    r := 2147483647 // não sei se é o maior inteiro: depende da implementação
    for i := 0; i < len(x); i++ {
        if x[i] < r {
            r = x[i]
        }
    }
    return r
}

// Calcula a distância de Levenshtein entre duas strings
func Levenshtein(s, t string) int {
    cost := 0

    if len(s) == 0 {
        return len(t)
    } else if len(t) == 0 {
        return len(s)
    }

    if s[len(s)-1] != t[len(t)-1] {
        cost = 1
    }

    maybe := make([]int, 0)
    maybe = append(maybe, Levenshtein(s[0:len(s)-1], t)+1)
    maybe = append(maybe, Levenshtein(s, t[0:len(t)-1])+1)
    maybe = append(maybe, Levenshtein(s[0:len(s)-1], t[0:len(t)-1])+cost)
    return min(maybe)
}

// Constrói dois vetores a partir de um mapa: um com as suas chaves,
// outro com os valores associados a tais chaves
func KeysAndValues(inlet map[string]string) ([]string, []string) {
    keys := make([]string, 0)
    values := make([]string, 0)

    for key, value := range inlet {
        keys = append(keys, key)
        values = append(values, value)
    }

    return keys, values
}

//
func MapFromKeysAndValues(keys, values []string) map[string]string {
    outlet := make(map[string]string)

    for i := 0; i < len(keys); i++ {
        outlet[keys[i]] = values[i]
    }

    return outlet
}
