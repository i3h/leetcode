package main

import "fmt"

func main() {
	p := &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 2,
			},
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 6,
					},
					&Node{
						Val: 7,
						Children: []*Node{
							&Node{
								Val: 11,
								Children: []*Node{
									&Node{
										Val: 14,
									},
								},
							},
						},
					},
				},
			},
			&Node{
				Val: 4,
				Children: []*Node{
					&Node{
						Val: 8,
						Children: []*Node{
							&Node{
								Val: 12,
							},
						},
					},
				},
			},
			&Node{
				Val: 5,
				Children: []*Node{
					&Node{
						Val: 9,
						Children: []*Node{
							&Node{
								Val: 13,
							},
						},
					},
					&Node{
						Val: 10,
					},
				},
			},
		},
	}
	fmt.Println(preorder(p))
}
