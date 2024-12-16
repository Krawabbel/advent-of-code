package main

import (
	"aoc/internal/lib"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Instruction struct {
	power                  bool
	x0, x1, y0, y1, z0, z1 int
}

type Input struct {
	instrs []Instruction
}

var re = regexp.MustCompile(`(\bon\b|\boff\b) x=([|-]?[0-9]*)\.\.([|-]?[0-9]*),y=([|-]?[0-9]*)\.\.([|-]?[0-9]*),z=([|-]?[0-9]*)\.\.([|-]?[0-9]*)`)

func preprocess(data []byte) Input {
	lines := strings.Split(string(data), "\n")
	instrs := make([]Instruction, len(lines))
	for i, line := range lines {
		m := re.FindStringSubmatch(line)
		pwr := false
		switch m[1] {
		case "on":
			pwr = true
		case "off":
			pwr = false
		default:
			lib.Panicf("unexpected power %s", m[1])
		}
		instrs[i] = Instruction{
			power: pwr,
			x0:    lib.MustToInt(m[2]),
			x1:    lib.MustToInt(m[3]),
			y0:    lib.MustToInt(m[4]),
			y1:    lib.MustToInt(m[5]),
			z0:    lib.MustToInt(m[6]),
			z1:    lib.MustToInt(m[7]),
		}
	}
	return Input{instrs: instrs}
}

func main() {
	path := os.Args[1]

	blob, err := os.ReadFile(path)
	lib.Must(err)

	input := preprocess(blob)

	part1(input)

	part2(input)
}

type Point struct {
	x, y, z int
}

func part1(input Input) {
	lb := func(i int) int { return max(-50, i) }
	ub := func(i int) int { return min(50, i) }
	on := lib.MakeSet[Point]()
	for _, instr := range input.instrs {
		f := on.Delete
		if instr.power {
			f = on.Insert
		}
		for x := lb(instr.x0); x <= ub(instr.x1); x++ {
			for y := lb(instr.y0); y <= ub(instr.y1); y++ {
				for z := lb(instr.z0); z <= ub(instr.z1); z++ {
					p := Point{x: x, y: y, z: z}
					f(p)
				}
			}
		}

	}
	fmt.Println("SOLUTION TO PART 1:", on.Size())
}

type Box struct {
	lb, ub Point
}

func (b Box) IsBox() bool {
	return b.lb.x <= b.ub.x && b.lb.y <= b.ub.y && b.lb.z <= b.ub.z
}

func (b Box) Area() int {
	return (b.lb.x - b.ub.x) * (b.lb.y - b.ub.y) * (b.lb.z - b.ub.z)
}

func (b Box) String() string {
	return fmt.Sprintf("[%d,%d]x[%d,%d]x[%d,%d]", b.lb.x, b.ub.x, b.lb.y, b.ub.y, b.lb.z, b.ub.z)
}

func (b Box) Contains(sub Box) bool {
	return b.lb.x <= sub.lb.x &&
		b.lb.y <= sub.lb.y &&
		b.lb.z <= sub.lb.z &&
		b.ub.x >= sub.ub.x &&
		b.ub.y >= sub.ub.y &&
		b.ub.z >= sub.ub.z
}

func PointMin(p1, p2 Point) Point {
	return Point{
		x: min(p1.x, p2.x),
		y: min(p1.y, p2.z),
		z: min(p1.y, p2.z),
	}
}

func PointMax(p1, p2 Point) Point {
	return Point{
		x: max(p1.x, p2.x),
		y: max(p1.y, p2.z),
		z: max(p1.y, p2.z),
	}
}

func Intersect(b1, b2 Box) Box {
	return Box{lb: PointMax(b1.lb, b2.ub), ub: PointMin(b1.ub, b2.ub)}
}

func (b Box) BisectX(cx int) (Box, Box, Box) {

	b1 := Box{
		lb: b.lb,
		ub: Point{cx - 1, b.ub.y, b.ub.z},
	}

	b2 := Box{
		lb: Point{cx, b.lb.y, b.lb.z},
		ub: Point{cx, b.ub.y, b.ub.z},
	}

	b3 := Box{
		lb: Point{cx + 1, b.lb.y, b.lb.z},
		ub: b.ub,
	}

	return b1, b2, b3
}

func (b Box) BisectY(cy int) (Box, Box, Box) {

	b1 := Box{
		lb: b.lb,
		ub: Point{b.ub.x, cy - 1, b.ub.z},
	}

	b2 := Box{
		lb: Point{b.lb.x, cy, b.lb.z},
		ub: Point{b.ub.x, cy, b.ub.z},
	}

	b3 := Box{
		lb: Point{b.lb.x, cy + 1, b.lb.z},
		ub: b.ub,
	}

	return b1, b2, b3
}

func (b Box) BisectZ(cz int) (Box, Box, Box) {

	b1 := Box{
		lb: b.lb,
		ub: Point{b.ub.x, b.ub.y, cz - 1},
	}

	b2 := Box{
		lb: Point{b.lb.x, b.lb.y, cz},
		ub: Point{b.ub.x, b.ub.y, cz},
	}

	b3 := Box{
		lb: Point{b.lb.x, b.lb.y, cz + 1},
		ub: b.ub,
	}

	return b1, b2, b3
}

func turn_on(on lib.Set[Box], area Box) {

	isPartitioned := true

	for _, box := range on.Slice() {

		if box.Contains(next) {
			isPartitioned = false
			break
		}

		cut := Intersect(box, area)
		if cut.IsBox() {
			areas.Insert(next.Bisect()...)
			isPartitioned = false
		}

	}

	// no subset, no intersection -> add to partition
	if isPartitioned {
		on.Insert(area)
	}

}

func turn_off(on lib.Set[Box], area Box) {

	for _, box := range on.Slice() {

		if area.Contains(box) {
			on.Delete(box)
		}

		cut := Intersect(box, area)
		if cut.IsBox() {
			on.Delete(box)
			on.Insert(box.Bisect()...)
		}

	}

}

func part2(input Input) {
	on := lib.MakeSet[Box]()

	for _, instr := range input.instrs {
		lb := Point{x: instr.x0, y: instr.y0, z: instr.z0}
		ub := Point{x: instr.x1, y: instr.y1, z: instr.z1}
		area := Box{lb: lb, ub: ub}

		if instr.power {
			turn_on(on, area)
		} else {
			turn_off(on, area)
		}

		fmt.Println(on)
		lib.MustPressEnter()
	}

	sol := 0
	for box := range on.Elements {
		sol += box.Area()
	}

	fmt.Println("SOLUTION TO PART 2:", sol)
}
