package lz

import (
	"fmt"
	"testing"
)

func Test(t *testing.T){
	string := "en N4IgzglgJiBcIBYAOAxAsgWgG4CECmAnBgKoAMIANCAMbRwhIBOeAtgIYDmE1lIbd8Jq07cMARl5IBIAPS8IAOyh4AHnDFiqAFwCeSPPTxY8Crby0QWeMFrYsk6gKwEATAA4xHxwGZSpAGxUUGy2cADaoBZW6h5u3i7+AHQIAOyOVCwA9gCuYHgAGuoIcRk5eQCacP6+2noG8Fm5eGiZxrwKbFgQHCEQmQoAKnVwpAC+FJGW9Z6eKSkuiWLepU2FsGLFyyCNFVUuKbX69DvNrQZUHV09Fv1DR7BjEyBR07FupASJ3gfbZQVFJV+TUqsEcbgQh3qQLyLTaF063V6t2GD3Gk2i6ze/lIiRcmmh/3WqQhBJBCH8BEhxz+sPOIEuiJugxRYwAuqMgA=="
	for {
		jsonString, err := DecompressFromBase64(string[3:])
		fmt.Println(jsonString, err)
	}
}
