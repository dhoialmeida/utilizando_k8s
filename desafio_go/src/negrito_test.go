package main

import "testing"

func TestNegrito(t *testing.T) {
    result := Negrito("Code.education rocks!")
    correto := "<b>Code.education rocks!</b>"
    if result != correto {
       t.Errorf("resultado incorreto: %s obtido, %s seria correto", result, correto)
    }
}
