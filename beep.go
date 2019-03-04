// Package allows playing sounds on Windows.
//
// Copyright (C) 2019 Tad Vizbaras. All Rights Reserved.

// +build windows

package beep

import (
	"syscall"
)

const (
	SimpleSound  = 0xFFFFFFFF
	ErrorSound   = 0x00000010
	StopSound    = 0x00000010
	WarningSound = 0x00000030
)

var (
	kernel32, _ = syscall.LoadLibrary("kernel32.dll")
	beep32, _   = syscall.GetProcAddress(kernel32, "Beep")

	user32, _    = syscall.LoadLibrary("user32.dll")
	msgbeep32, _ = syscall.GetProcAddress(user32, "MessageBeep")
)

var note2freq = map[string]int{
	`c`: 261,
	`d`: 293,
	`e`: 329,
	`f`: 349,
	`g`: 392,
	`a`: 440,
	`b`: 493,
}

func Beep(freq, duration int) (err error) {
	_, _, errno := syscall.Syscall(
		uintptr(beep32),
		uintptr(2),
		uintptr(freq),
		uintptr(duration),
		0,
	)
	if errno != 0 {
		err = errno
	}
	return
}

func Alert() {
	Beep(note2freq[`e`], 500)
}

func Error() {
	Play(ErrorSound)
}

func Stop() {
	Play(StopSound)
}

func Warning() {
	Play(WarningSound)
}

func Play(typ int) (err error) {
	_, _, errno := syscall.Syscall(
		uintptr(msgbeep32),
		uintptr(1),
		uintptr(typ),
		0,
		0,
	)
	if errno != 0 {
		err = errno
	}
	return
}
