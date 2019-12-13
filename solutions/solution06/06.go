package solution06

import (
	"adventofcode/inputs/input06"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func Run(){
	nodes := ReadProblem(input06.Problem)
	diracc := int(0)


	for _,val := range nodes{
		traverse(val,&diracc)
	}
	fmt.Println("Total",diracc)
}


/*******************************************
    G - H       J - K - L
       /           /
COM - B - C - D - E - F
               \
                I
 *******************************************/

func traverse(node *Node,diracc *int){
	parent := node.Parent
	if parent == nil{
		return
	}
	*diracc++
	traverse(node.Parent,diracc)
}

/******************************************
						  YOU
                         /
        G - H       J - K - L
       /           /
COM - B - C - D - E - F
               \
                I - SAN
*******************************************/


type Node struct{
	Name     string
	Children []*Node
	Parent   *Node
}

func (node *Node) AddChildren(child *Node){
	child.Parent = node
	length := len(node.Children)
	if length == 0{
		node.Children = append(make([]*Node,0),child)
	} else{
		node.Children = append(node.Children, child)
	}
}

func ReadProblem(problem string) (result map[string]*Node){
	result = make(map[string]*Node)

	var lines = strings.NewReader(problem)
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		relations := strings.Split(line,")")

		// Create nodes
		for _,name := range relations{
			_, exists := result[name]
			if !exists {
				node := Node{Name:name}
				result[name] = &node
			}
		}

		parent := result[relations[0]]
		child := result[relations[1]]
		parent.AddChildren(child)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Can not scan inputs!")
	}
	return
}
