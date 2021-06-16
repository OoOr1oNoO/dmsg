//+build windows

package dmsgpty

import (
	"golang.org/x/sys/windows"
)

// getSize gets windows terminal size
func getSize() (*windows.Coord, error) {
	var bufInfo windows.ConsoleScreenBufferInfo
	c, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return nil, err
	}
	if err = windows.GetConsoleScreenBufferInfo(c, &bufInfo); err != nil {
		return nil, err
	}
	return &windows.Coord{
		X: bufInfo.Window.Right - bufInfo.Window.Left + 1,
		Y: bufInfo.Window.Bottom - bufInfo.Window.Top + 1,
	}, nil
}

// Start starts the pty.
func (sc *PtyClient) Start(name string, arg ...string) error {
	sz, err := getSize()
	if err != nil {
		sc.log.WithError(err).Warn("failed to obtain terminal size")
	}
	return sc.StartWithSize(name, arg, sz)
}

// StartWithSize starts the pty with a specified size.
func (sc *PtyClient) StartWithSize(name string, arg []string, c *windows.Coord) error {
	return sc.call("Start", &CommandReq{Name: name, Arg: arg, Size: c}, &empty)
}

// SetPtySize sets the pty size.
func (sc *PtyClient) SetPtySize(size *windows.Coord) error {
	return sc.call("SetPtySize", size, &empty)
}
