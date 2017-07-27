package main

import "../../AllenDang/w32" // I had putted the w32 package inside the project source folder -
import "syscall"
import "unsafe"

func MakeIntResource(id uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}
func WndProc(hWnd w32.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
	default:
		return w32.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}
func WinMain() int {
	hInstance := w32.GetModuleHandle("")
	lpszClassName := syscall.StringToUTF16Ptr("WNDclass")
	var wcex w32.WNDCLASSEX
	wcex.Size = uint32(unsafe.Sizeof(wcex))
	wcex.Style = w32.CS_HREDRAW | w32.CS_VREDRAW
	wcex.WndProc = syscall.NewCallback(WndProc)
	wcex.ClsExtra = 0
	wcex.WndExtra = 0
	wcex.Instance = hInstance
	wcex.Icon = w32.LoadIcon(hInstance, MakeIntResource(w32.IDI_APPLICATION))
	wcex.Cursor = w32.LoadCursor(0, MakeIntResource(w32.IDC_ARROW))
	wcex.Background = w32.COLOR_WINDOW + 11
	wcex.MenuName = nil

	wcex.ClassName = lpszClassName
	wcex.IconSm = w32.LoadIcon(hInstance, MakeIntResource(w32.IDI_APPLICATION))
	w32.RegisterClassEx(&wcex)
	hWnd := w32.CreateWindowEx(
		0, lpszClassName, syscall.StringToUTF16Ptr("Simple Go Window!"),
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT, 400, 400, 0, 0, hInstance, nil)
	w32.ShowWindow(hWnd, w32.SW_SHOWDEFAULT)
	w32.UpdateWindow(hWnd)
	var msg w32.MSG
	for {
		if w32.GetMessage(&msg, 0, 0, 0) == 0 {
			break
		}
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
	return int(msg.WParam)
}
func main() {
	WinMain()
	return
}
