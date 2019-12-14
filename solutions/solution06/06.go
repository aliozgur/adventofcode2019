package solution06

import (
	"adventofcode/inputs/input06"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func Run(){

	fmt.Println("Read problem...")
	nodes := ReadProblem(input06.Problem)
	fmt.Printf("Found %d nodes\n", len(nodes))
	fmt.Println(strings.Repeat("-",80))


	// Part 1
	fmt.Println("Solving Part 1...")
	diracc := int(0)
	for _,val := range nodes{
		traverse(val,&diracc)
	}
	fmt.Println("Total direct and indirect orbits:",diracc)
	fmt.Println("Part 1 DONE")
	fmt.Println(strings.Repeat("-",80))


	// Part 2
	fmt.Println("Solving Part 2...")

	root := nodes["COM"]
	first := nodes["YOU"].Parent.Name
	second := nodes["SAN"].Parent.Name
	minDistance := findDist(root,first,second)

	fmt.Println("Minimum transfers required:",minDistance)
	fmt.Println("Part 2 DONE")
	fmt.Println(strings.Repeat("-",80))

}

type Node struct{
	Name 		string
	Children  	[]*Node
	Parent    	*Node
	Righ 	  	*Node
	Left		*Node
}

func (node Node) String() string {
	parentName := ""
	if node.Parent != nil{
		parentName = node.Parent.Name
	}

	return fmt.Sprintf("[ Name: %s, Parent: %s, State: %d, Children Count:%d ]", node.Name, parentName, len(node.Children))
}

func (node *Node) AddChild(child *Node){
	if child == nil{
		return
	}

	child.Parent = node

	var length = len(node.Children)
	if length == 0{
		node.Children = append(make([]*Node,0),child)
	} else{
		node.Children = append(node.Children, child)
	}

	if length == 0{
		node.Righ = child
	} else if length == 1 {
		node.Left = child
	}
}

/*******************************************
    G - H       J - K - L
       /           /
COM - B - C - D - E - F
               \
                I
 *******************************************/

/*
traverse all direct and indirect routes for all nodes in a tree starting from this node
*/
func traverse(node *Node,diracc *int){
	parent := node.Parent
	if parent == nil{
		return
	}
	*diracc++
	traverse(node.Parent,diracc)
}

/*
	Find minimum distance between nodes with values first and second
*/
func findDist(root *Node,first string,second string) int {

	lca := findLCA(root,first,second)
	d1 := heightFrom(lca,first)
	d2 := heightFrom(lca,second)
	return d1 + d2
}

/*
Find lowest common ancestor of nodes with values first and second
*/
func findLCA(root *Node,first string, second string) *Node{
	if root == nil{
		return nil
	}

	if root.Name == first || root.Name == second{
		return root
	}

	left := findLCA(root.Left,first, second)
	right := findLCA(root.Righ,first,second)

	if left != nil && right != nil {
		return root
	}

	if left == nil{
		return right
	} else {
		return left
	}
}
/*
Find height from root to a node with nodeName
*/
func heightFrom(root *Node,nodeName string) int {
	if root == nil || root.Name == nodeName{
		return 0
	}
	var queue []*Node
	queue = append(queue, root) // Enqueue
	queue = append(queue, nil)
	height := 0
	for ;len(queue)>0;{
		temp := queue[0] // get first element
		queue = queue[1:] // dequeue

		if temp == nil{
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
			height++
		} else{
			if temp.Name == nodeName{
				return height
			}

			if temp.Left != nil{
				queue = append(queue, temp.Left)
			}
			if temp.Righ != nil{
				queue = append(queue, temp.Righ)
			}

		}
	}
	return height
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
		parent.AddChild(child)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Can not scan inputs!")
	}
	return
}
