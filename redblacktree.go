package main

import "fmt"

type TreeNode struct {
	data int
	leftChild *TreeNode
	rightChild *TreeNode
	color Color //1,代表红色，0代表黑色
}
type Color int
const(
	Black Color=0
	Red Color =1
)
func main() {
	data:=[]int{5,3,6,2,4,7,8,1}
	root:=createTree(data)
	zhongXuBianLiTree(root)
	qianXuBianLiTree(root)
}
func createTree(datas []int)*TreeNode{
	if len(datas)==0{
		return nil
	}
	root:=&TreeNode{data:datas[0],color:0}
	for _,data:=range datas[1:]{
		ok:=insertTree(data,root)
		//如果插入新元素破坏了红黑树规则，则进行调整
		if ok{
			_,tmpRoot,_:=tiaoZhengTree(root,nil,nil,nil,false,false)
			if tmpRoot!=nil{
				root=tmpRoot
			}
			root.color=Black
		}

	}
	return root
}
//返回值，说明新插入的值是否出现双红
func insertTree(data int,node *TreeNode)bool{
	if node==nil{
		return false
	}
	var ok bool
	if data>=node.data{
		ok=insertTree(data,node.rightChild)
		if node.rightChild==nil{
			node.rightChild=&TreeNode{data:data,color:1}
			if node.color==Red&&node.rightChild.color==Red{
				ok=true
			}
		}
	}else{
		ok=insertTree(data,node.leftChild)
		if node.leftChild==nil{
			node.leftChild=&TreeNode{data:data,color:1}
			if node.color==Red&&node.leftChild.color==Red{
				ok=true
			}
		}
	}
	return ok
}
func tiaoZhengTree(node *TreeNode,brotherNode *TreeNode,parentNode *TreeNode,parentOfParent *TreeNode,nodeIsLeft bool,parentIsLeft bool)(color Color,root *TreeNode,status bool){
	if node==nil{
		return Black,nil,false
	}
	root=nil

	leftColor,leftRoot,leftStatus:=tiaoZhengTree(node.leftChild,node.rightChild,node,parentNode,true,nodeIsLeft)
	if leftStatus{
		return leftColor,leftRoot,leftStatus
	}
	rightColor,rightRoot,rightStatus:=tiaoZhengTree(node.rightChild,node.leftChild,node,parentNode,false,nodeIsLeft)
	if rightStatus{
		return rightColor,rightRoot,rightStatus
	}
	//连续两个红色
	if node.color==Red&&leftColor==Red{
		//如果叔父节点是黑色,//进行旋转
		if brotherNode==nil||brotherNode.color==Black{
			//LL型
			if nodeIsLeft{
				//如果父节点的父节点为nil时，则说明parent为root节点
				if parentOfParent==nil{
					root=node
				}else{
					//否则父节点不是root节点
					if parentIsLeft{
						parentOfParent.leftChild=node
					}else{
						parentOfParent.rightChild=node
					}
				}
				tmpRight:=node.rightChild
				node.rightChild=parentNode
				parentNode.leftChild=nil
				if tmpRight!=nil{

					//将tmpRight挂在node.right的最左节点
					tmp:=node.rightChild
					for tmp.leftChild!=nil{
						tmp=tmp.leftChild
					}
					tmp.leftChild=tmpRight
				}
				//将node节点调成黑色，其左右孩子为红色
				node.color=Black
				if node.leftChild!=nil{
					node.leftChild.color=Red
				}
				if node.rightChild!=nil{
					node.rightChild.color=Red
				}
				return node.color,root,true

			}else{
				//否则是RL型
				//如果parentOfParent为nil，则说明parent为root
				if parentOfParent==nil{
					root=node.leftChild
				}else{
					//否则，parent不是父节点
					if parentIsLeft{
						parentOfParent.leftChild=node.leftChild
					}else {
						parentOfParent.rightChild = node.leftChild
					}
				}
				kNode:=node.leftChild
				tmpRight:=kNode.rightChild
				tmpLeft:=kNode.leftChild

				kNode.rightChild=parentNode
				kNode.leftChild=node
				parentNode.rightChild=nil
				node.leftChild=nil

				if tmpLeft!=nil{
					tmp:=kNode.leftChild
					for tmp.rightChild!=nil{
						tmp=tmp.rightChild
					}
					tmp.rightChild=tmpLeft
				}
				if tmpRight!=nil{
					tmp:=kNode.rightChild
					for tmp.leftChild!=nil{
						tmp=tmp.leftChild
					}
					tmp.leftChild=tmpRight
				}
				//将node节点调成黑色，其左右孩子为红色
				kNode.color=Black
				if kNode.leftChild!=nil{
					kNode.leftChild.color=Red
				}
				if kNode.rightChild!=nil{
					kNode.rightChild.color=Red
				}
				return kNode.color,root,true
			}

		}else{
			//否则进行父祖换色调整
			parentNode.color=Red
			if parentNode.leftChild!=nil{
				parentNode.leftChild.color=Black
			}
			if parentNode.rightChild!=nil{
				parentNode.rightChild.color=Black
			}
			return node.color,root,false
		}
	}
	if node.color==Red && rightColor==Red{
		//如果叔父节点是黑色,//进行旋转
		if brotherNode==nil||brotherNode.color==Black{
			//LR型
			if nodeIsLeft{
				//如果parentOfParent为nil，则说明parent为root
				if parentOfParent==nil{
					root=node.rightChild
				}else{
					//否则，parent不是父节点
					if parentIsLeft{
						parentOfParent.leftChild=node.rightChild
					}else{
						parentOfParent.rightChild=node.rightChild
					}

				}
				kNode:=node.rightChild
				tmpRight:=kNode.rightChild
				tmpLeft:=kNode.leftChild

				kNode.leftChild=node
				kNode.rightChild=parentNode
				parentNode.leftChild=nil
				node.rightChild=nil

				if tmpLeft!=nil{
					tmp:=kNode.leftChild
					for tmp.rightChild!=nil{
						tmp=tmp.rightChild
					}
					tmp.rightChild=tmpLeft
				}
				if tmpRight!=nil{
					tmp:=kNode.rightChild
					for tmp.leftChild!=nil{
						tmp=tmp.leftChild
					}
					tmp.leftChild=tmpRight
				}
				//将node节点调成黑色，其左右孩子为红色
				kNode.color=Black
				if kNode.leftChild!=nil{
					kNode.leftChild.color=Red
				}
				if kNode.rightChild!=nil{
					kNode.rightChild.color=Red
				}
				return kNode.color,root,true

			}else{
				//RR型
				//如果父节点的父节点为nil时，则说明parent为root节点
				if parentOfParent==nil{
					root=node
				}else{
					//否则父节点不是root节点
					if parentIsLeft{
						parentOfParent.leftChild=node
					}else{
						parentOfParent.rightChild=node
					}
				}
				tmpLeft:=node.leftChild
				node.leftChild=parentNode
				parentNode.rightChild=nil
				if tmpLeft!=nil{
					//将tmpRight挂在node.right的最左节点
					tmp:=node.leftChild
					for tmp.rightChild!=nil{
						tmp=tmp.rightChild
					}
					tmp.rightChild=tmpLeft
				}
				//将node节点调成黑色，其左右孩子为红色
				node.color=Black
				if node.leftChild!=nil{
					node.leftChild.color=Red
				}
				if node.rightChild!=nil{
					node.rightChild.color=Red
				}
				return node.color,root,true
			}

		}else{
			//否则进行父祖换色调整
			node.color=Red
			if node.leftChild!=nil{
				node.leftChild.color=Black
			}
			if node.rightChild!=nil{
				node.rightChild.color=Black
			}
			return node.color,root,false
		}
	}
	return node.color,root,false
}
func qianXuBianLiTree(root *TreeNode){
	if root==nil{
		return
	}
	fmt.Println(root.data)
	qianXuBianLiTree(root.leftChild)
	qianXuBianLiTree(root.rightChild)
}
func zhongXuBianLiTree(root *TreeNode){
	if root==nil{
		return
	}
	zhongXuBianLiTree(root.leftChild)
	fmt.Println(root.data," ",root.color)
	zhongXuBianLiTree(root.rightChild)
}
func houXuBianLiTree(root *TreeNode){

	if root==nil{
		return
	}
	houXuBianLiTree(root.leftChild)
	houXuBianLiTree(root.rightChild)
	fmt.Println(root.data)

}

