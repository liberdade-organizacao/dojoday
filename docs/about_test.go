package docs

/*
Da pasta principal do repositório, execute o comando a seguir:

    go test github.com/liberdade-organizacao/dojoday

para executar os testes aqui
*/

import "testing"
import "fmt"

func TestFuncaoSoma(t *testing.T) {
    dois := Some(1, 1)
    if dois != 2 {
        t.Error("1 + 1 != 2\nTem certeza?")
    }
}

func TestAckermannFunction(t *testing.T) {
    x := Ackermann(4, 1)
    if x != 65533 {
        t.Error(fmt.Sprintf("Algo de errado não está certo... x = %d\n", x))
    }
}

func TestLevenshteinDistance(t *testing.T) {
    x := Levenshtein("kitten", "sitting")
    if x != 3 {
        t.Error("There's something wrong with that Levenshtein dude")
    }
}
