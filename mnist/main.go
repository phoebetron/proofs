package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/module/multic"
	"github.com/phoebetron/getlin/module/single"
	"github.com/phoebetron/getlin/stat32"
	"github.com/phoebetron/proofs/mnist/loader"
)

const (
	bas = "/Users/xh3b4sd/project/phoebetron/proofs/"
	tes = "mnist_test.csv"
	tra = "mnist_train.csv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var ldr getlin.Loader
	{
		ldr = loader.New(loader.Config{
			Bas: bas,
			Fil: tra,
			Git: true,
			Org: "phoebetron",
			Rep: "mnist",
		})
	}

	{
		ldr.Create()
	}

	var con single.Config
	{
		con = single.Config{
			Cla: 1000,
			Fre: 0.10,
			Inp: 784,
		}
	}

	var mod getlin.Module
	{
		mod = multic.New(multic.Config{
			Sin: []getlin.Module{
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
				single.New(con),
			},
		})
	}

	var met getlin.Metric
	for i := 1; i <= 100; i++ {
		var bat [][2]getlin.Vector
		{
			bat = ldr.Search()
		}

		for _, x := range bat {
			mod.Update(x[0])
			mod.Update(x[1])
		}

		{
			met = mod.Verify(bat)
		}

		fmt.Printf(
			"epo %4d        mae %4.3f        cla %4d        prd %4d\n",
			i,
			met.Get().Err().Mae(),
			bat[0][1].Cla(),
			stat32.Argmax(mod.Search(bat[0][1]).Out()),
		)
	}

	{
		prints(mod, ldr)
	}
}

func prints(mod getlin.Module, ldr getlin.Loader) {
	var bat [][2]getlin.Vector
	{
		bat = ldr.Search()[:10]
	}

	for _, x := range bat {
		var p []float32
		{
			p = mod.Search(x[1]).Out()
		}

		var c int
		var r []float32
		{
			c = x[1].Cla()
			r = x[1].Inp()
		}

		for i := 0; i < 784; i++ {
			if r[i] <= 0.5 {
				fmt.Printf(" ")
			}
			if r[i] > 0.5 {
				fmt.Printf("#")
			}
			if i%28 == 0 {
				fmt.Printf("\n")
			}
		}
		fmt.Printf("\n")
		fmt.Printf("cla %v\n", c)
		fmt.Printf("prd %v\n", p)
		fmt.Printf("\n")
	}
}
