package loader

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (l *Loader) Search() [][2]getlin.Vector {
	var err error

	var dst string
	{
		dst = filepath.Join(l.bas, l.rep, "/data/", l.fil)
	}

	var des *os.File
	{
		des, err = os.Open(dst)
		if err != nil {
			panic(err)
		}
	}

	var rdr *csv.Reader
	{
		rdr = csv.NewReader(des)
	}

	var all [][]string
	{
		all, err = rdr.ReadAll()
		if err != nil {
			panic(err)
		}
	}

	var vec [][2]getlin.Vector
	for len(vec) < l.bat {
		var cla []int
		var ind []int
		var sam [2]getlin.Vector

		for len(cla) < 2 {
			var ran int
			{
				ran = rand.Intn(len(all))
			}

			var lab int
			{
				lab = musint(all[ran][0])
			}

			if len(cla) == 0 || len(cla) == 1 && cla[0] != lab {
				ind = append(ind, ran)
				cla = append(cla, lab)
			}
		}

		for i := range sam {
			sam[i] = vector.New(vector.Config{
				Cla: cla[1],
				Inp: musinp(all[ind[i]][1:]),
				Out: []float32{float32(i)},
			})
		}

		{
			vec = append(vec, sam)
		}
	}

	return vec
}

func musinp(str []string) []float32 {
	var max float32
	{
		max = -math.MaxFloat32
	}

	var fea []float32
	for _, x := range str {
		var f32 float32
		{
			f32 = musf32(x)
		}

		if f32 > max {
			max = f32
		}

		{
			fea = append(fea, f32)
		}
	}

	for i := range fea {
		fea[i] /= max
	}

	return fea
}

func musf32(str string) float32 {
	f64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		panic(err)
	}

	return float32(f64)
}

func musint(str string) int {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(num)
}
