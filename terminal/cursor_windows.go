package terminal

import (
	"os"
	"syscall"
	"unsafe"
)

func CursorUp(n int) {
	cursorMove(0, n)
}

func CursorDown(n int) {
	cursorMove(0, -1*n)
}

func CursorForward(n int) {
	cursorMove(n, 0)
}

func CursorBack(n int) {
	cursorMove(-1*n, 0)
}

func cursorMove(x int, y int) {
	handle := syscall.Handle(os.Stdout.Fd())

	var csbi consoleScreenBufferInfo
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))

	var cursor Coord
	cursor.x = csbi.cursorPosition.x + Short(x)
	cursor.y = csbi.cursorPosition.y + Short(y)

	procSetConsoleCursorPosition.Call(uintptr(handle), uintptr(*(*int32)(unsafe.Pointer(&cursor))))
}

func CursorNextLine(n int) {
	CursorUp(n)
	CursorHorizontalAbsolute(0)
}

func CursorPreviousLine(n int) {
	CursorDown(n)
	CursorHorizontalAbsolute(0)
}

func CursorHorizontalAbsolute(x int) {
	handle := syscall.Handle(os.Stdout.Fd())

	var csbi consoleScreenBufferInfo
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))

	var cursor Coord
	cursor.x = Short(x)
	cursor.y = csbi.cursorPosition.y

	if csbi.size.x < cursor.x {
		cursor.x = csbi.size.x
	}

	procSetConsoleCursorPosition.Call(uintptr(handle), uintptr(*(*int32)(unsafe.Pointer(&cursor))))
}

func CursorShow() {
	handle := syscall.Handle(os.Stdout.Fd())

	var cci consoleCursorInfo
	procGetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
	cci.visible = 1

	procSetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
}

func CursorHide() {
	handle := syscall.Handle(os.Stdout.Fd())

	var cci consoleCursorInfo
	procGetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
	cci.visible = 0

	procSetConsoleCursorInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&cci)))
}

func CursorLocation() (Coord, error) {
	handle := syscall.Handle(os.Stdout.Fd())

	var csbi consoleScreenBufferInfo
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))

	return csbi.cursorPosition, nil
}

func Size() (Coord, error) {
	handle := syscall.Handle(os.Stdout.Fd())

	var csbi consoleScreenBufferInfo
	procGetConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))

	return csbi.size, nil
}
