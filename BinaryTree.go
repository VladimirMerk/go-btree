package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

type Circle struct {
	p   image.Point
	r   int
}

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red  color.Color = color.RGBA{200, 30, 30, 255}
var width, height int = 500, 500

func main() {
	// Пока не знаю что такое defer но это работает
	defer elapsed("total")()
	var orig []int
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//treeLength := 10
	//fmt.Println("treeLength", treeLength)
	//var tree *Tree = nil;
	//for i := 0; i <= treeLength; i++ {
	//	val := r1.Intn(treeLength)
	//	orig = append(orig, val)
	//	if (tree == nil) {
	//		tree = &Tree{nil,val,nil}
	//	} else {
	//		tree.addNode(&Tree{nil,val,nil})
	//	}
	//}

	var tree Tree
	list := []int{9,8,6,5,4,3,1,2,7,10,0}
	for i := 0; i < len(list); i++ {
		if i == 0 {
			tree = Tree{nil,list[i],nil}
		} else {
			tree.addNode(&Tree{nil,list[i],nil})
		}
	}

	fmt.Println("tree min value", tree.getMinValue())
	fmt.Println("tree max value", tree.getMaxValue())

	found := tree.findValue(6)
	if found == nil {
		fmt.Println("find 6 is nil")
	} else {
		fmt.Println("find 6 is", found)
	}
	fmt.Println("Original sequence", orig)
	fmt.Println("Bypass preorder", tree.bypassPreorder())
	fmt.Println("Bypass inorder", tree.bypassInorder())
	fmt.Println("Bypass postorder", tree.bypassPostorder())
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func (tree *Tree) addNode(node *Tree) {
	if (node.Value < tree.Value) {
		if (tree.Left == nil) {
			tree.Left = node
		} else {
			tree.Left.addNode(node)
		}
	} else {
		if (tree.Right == nil) {
			tree.Right = node
		} else {
			tree.Right.addNode(node)
		}
	}
}

// Получение минимального значения дерева
func (tree *Tree) getMinValue() int {
	if (tree.Left == nil) {
		return tree.Value
	} else {
		return tree.Left.getMinValue()
	}
}

// Получение максимального значения дерева
func (tree *Tree) getMaxValue() int {
	if (tree.Right == nil) {
		return tree.Value
	} else {
		return tree.Right.getMinValue()
	}
}

func (tree *Tree) findValue(findValue int) *int {
	if tree.Value == findValue {
		return &tree.Value
	} else if tree.Value > findValue && tree.Left != nil {
		return tree.Left.findValue(findValue)
	} else if tree.Value < findValue && tree.Right != nil {
		return tree.Right.findValue(findValue)
	} else {
		return nil
	}
}

// Прямой обход (NLR)
// Проверяем, не является ли текущий узел пустым или null.
// Показываем поле данных корня (или текущего узла).
// Обходим левое поддерево рекурсивно, вызвав функцию прямого обхода.
// Обходим правое поддерево рекурсивно, вызвав функцию прямого обхода
// https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/Sorted_binary_tree_preorder.svg/330px-Sorted_binary_tree_preorder.svg.png
func (node *Tree) bypassPreorder() []int {
	var result []int
	result = append(result, node.Value)
	// Сначала выполняется обход левых ветвей от рута
	if node.Left != nil {
		result = append(result, node.Left.bypassPreorder()...)
	}
	// Затем выполняется обход правых ветвей от рута
	if node.Right != nil {
		result = append(result, node.Right.bypassPreorder()...)
	}
	return result
}

// Симметричный обход (Центрированный обход) (LNR)
// Проверяем, не является ли текущий узел пустым или null.
// Обходим левое поддерево рекурсивно, вызвав функцию центрированного обхода.
// Показываем поле данных корня (или текущего узла).
// Обходим правое поддерево рекурсивно, вызвав функцию центрированного обхода
// https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Sorted_binary_tree_inorder.svg/330px-Sorted_binary_tree_inorder.svg.png
func (node *Tree) bypassInorder() []int {
	var result []int
	if node.Left != nil {
		result = append(result, node.Left.bypassInorder()...)
	}
	if node.Right != nil {
		result = append(result, node.Value)
		result = append(result, node.Right.bypassInorder()...)
	} else {
		result = append(result, node.Value)
	}
	return result
}

// Обратный обход (LRN)
// Проверяем, не является ли текущий узел пустым или null.
// Обходим левое поддерево рекурсивно, вызвав функцию обратного обхода.
// Обходим правое поддерево рекурсивно, вызвав функцию обратного обхода.
// Показываем поле данных корня (или текущего узла).
func (node *Tree) bypassPostorder() []int {
	// ЕСТЬ КОСЯК С ПОСЛЕДОВАТЕЛЬНОСТЬЮ
	// 9,8,6,5,4,3,1,2,7,10,0
	var result []int
	if node.Left == nil && node.Right == nil {
		result = append(result, node.Value)
	}
	if node.Left != nil {
		result = append(result, node.Left.bypassPostorder()...)
	}
	if node.Right != nil {
		result = append(result, node.Right.bypassPostorder()...)
		if (node != nil) {
			result = append(result, node.Value)
		}
	}
	return result
}

func drawTree() {
	file, err := os.Create("someimage.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)
	// или draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)
	mask := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(mask, mask.Bounds(), &image.Uniform{red}, image.ZP, draw.Src)
	draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{width / 2, height / 2}, 20}, image.ZP, draw.Over)
	png.Encode(file, img)
}

func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
