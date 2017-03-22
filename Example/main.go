package main

import (
	"my/argparser"
	"os"
	"log"
	"fmt"
)
func onChildVerbOpt(_ *argparser.String){
	fmt.Println("Yee")
}

func onChildVerb(verb *argparser.Verb){
	original, ok := verb.Options[0].(*argparser.String);
	if ok == true {
		fmt.Println(original.Value())
	}
}

func onMainVerbOpt(str *argparser.Bool){
	fmt.Println(str.Value())
}

func main() {

	parser := argparser.NewParser(&argparser.Verb{

			Options: []argparser.Option{
				&argparser.Bool{
					LongName: "mainVerbOption",
					Function: onMainVerbOpt,
				},
			},

			ChildVerbs: []*argparser.Verb{
				{
					Name: "childVerb",
					Function: onChildVerb,

					Options: []argparser.Option{
						&argparser.String{
							LongName: "childVerbOption",
							Function: onChildVerbOpt,
						},
					},

					ChildVerbs: []*argparser.Verb{
						{
							Name: "childchildVerb",
						},
					},
				},
			},
		},
	)

	err := parser.Parse(os.Args[1:])
	if err != nil{
		log.Fatal(err)
	}
}
