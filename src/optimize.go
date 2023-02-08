package main

const (
	SET_ZERO = "[-]"
	ADD_LEFT = "[<+>-]"
	ADD_RIGHT = "[>+<-]"
)

func Optimize(commands []*Command) []*Command {
	cnt := 0
	optimized := []*Command{commands[0]}
	i := 1
	for i < len(commands) {
		if optimized[cnt].String == commands[i].String {
			optimized[cnt].Count++
		} else if isPattern(SET_ZERO, i, commands) {
			optimized = append(optimized, &Command {
				String: SET_ZERO,
				Count: 1,
				Row: commands[i].Row,
				Col: commands[i].Col,
			})
			i += 2
			cnt++
		} else if isPattern(ADD_LEFT, i, commands) {
			optimized = append(optimized, &Command {
				String: ADD_LEFT,
				Count: 1,
				Row: commands[i].Row,
				Col: commands[i].Col,
			})
			i += 5
			cnt++
		} else if isPattern(ADD_RIGHT, i, commands) {
			optimized = append(optimized, &Command {
				String: ADD_RIGHT,
				Count: 1,
				Row: commands[i].Row,
				Col: commands[i].Col,
			})
			i += 5
			cnt++
		} else {
			optimized = append(optimized, commands[i])
			cnt++
		}
		i++
	}

	return optimized
}

func isPattern(pattern string, i int, commands []*Command) bool {
	n := len(commands)
	if i > n - len(pattern) {
		return false
	}

	for j := 0; j < len(pattern); j++ {
		if commands[i+j].String != string(pattern[j]) {
			return false
		}
	}

	return true
}
