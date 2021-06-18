package m71

// 思路
//	栈操作，读到一个实际文件夹，压栈，读到'.'，什么也不做，读到".."，弹栈，最后从栈底到栈顶输出
func simplifyPath(path string) string {
	stack := make([]string, 0)
	status := 0 // 0 - 未开始、1 - 待读入、2 - 单点、3 - 双点、4 - 读路径
	name := make([]byte, 0, 32)
	for i := 0; i < len(path); i++ {
		v := path[i]
		if v == '/' {
			if status == 3 { // ..
				if len(stack) > 0 {
					stack = stack[0 : len(stack)-1]
				}
			} else if status == 4 {
				stack = append(stack, string(name))
			}
			status = 1       // status == 0 || status == 2 /*.*/
			name = name[0:0] // 待读入状态，已读入内容归0
		} else if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v == '_' || v >= '0' && v <= '9' {
			if status == 1 || status == 2 || status == 3 {
				status = 4
			}
			name = append(name, v)
		} else if v == '.' {
			if status == 1 {
				status = 2
			} else if status == 2 {
				status = 3
			} else if status == 3 {
				status = 4
			} else {
				// 4状态不用变
			}
			name = append(name, v)
		}
	}
	if status == 4 && len(name) > 0 {
		stack = append(stack, string(name))
	} else if status == 3 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
	}

	finalpath := "/"
	for _, v := range stack {
		finalpath += v + "/"
	}
	if len(finalpath) > 1 {
		finalpath = finalpath[0 : len(finalpath)-1]
	}

	return finalpath
}
