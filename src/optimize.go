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
			cnt--
		} else if isPattern(SET_ZERO, i, commands) {
			optimized = append(optimized, newCommand(SET_ZERO, commands[i]))
			i += 2
		} else if isPattern(ADD_LEFT, i, commands) {
			optimized = append(optimized, newCommand(ADD_LEFT, commands[i]))
			i += 5
		} else if isPattern(ADD_RIGHT, i, commands) {
			optimized = append(optimized, newCommand(ADD_RIGHT, commands[i]))
			i += 5
		} else {
			optimized = append(optimized, commands[i])
		}
		cnt++
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

func newCommand(pattern string, command *Command) *Command {
	return &Command {
		String: pattern,
		Count: 1,
		Row: command.Row,
		Col: command.Col,
	}
}
