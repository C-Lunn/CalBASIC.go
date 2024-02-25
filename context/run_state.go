package context

type RUN_STATE uint8

const (
	RUN_STATE_NORMAL RUN_STATE = iota
	RUN_STATE_IF
	RUN_STATE_FOR
)

type Run_state interface {
	Get_state() RUN_STATE
}

type Run_state_normal struct {
}

func (r Run_state_normal) Get_state() RUN_STATE {
	return RUN_STATE_NORMAL
}

type Run_state_if struct {
	Condition_pc    int
	Else_pc         int
	End_pc          int
	Result          bool
	Needs_condition bool
}

func (r Run_state_if) Get_state() RUN_STATE {
	return RUN_STATE_IF
}

type Run_state_for struct {
	First_pc int
}
