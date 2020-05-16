package types

const (
	JS_LEFT = 0
	JS_RIGHT = 1
	JS_UP = 2
	JS_DOWN = 3
	JS_START = 4
)

type Joystick struct {
	Button int `json:"button"`
}
