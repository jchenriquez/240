package search2d

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Matrix  [][]int `json:"matrix"`
	Targets []int   `json:"targets"`
	Output  []bool  `json:"output"`
}

func TestSearchMatrix(tst *testing.T) {
	f, err := os.Open("./test.json")

	if err != nil {
		tst.Error(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	var tests map[string]Test

	for {
		err := decoder.Decode(&tests)
		if err == nil {
			for name, test := range tests {
				tst.Run(name, func(st *testing.T) {
					for i := 0; i < len(test.Targets); i++ {
						if SearchMatrix(test.Matrix, test.Targets[i]) != test.Output[i] {
							st.Errorf("target %d expected %v\n", test.Targets[i], test.Output[i])
						}
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			tst.Error(err)
		}
	}
}
