package main

/*
//#define DEBUG 1

#cgo LDFLAGS: -framework Foundation -framework Foundation -framework Carbon
#include "keyenv.h"
#ifdef DEBUG
#include <stdio.h>
#endif

#define DOWN_SHIFT_AND_OPT 655650
#define LEFT_OPT 58
#define LEFT_SHIFT 56
#define DOWN_CTRL_OPT 786721
#define KEY_SPACE 49



static inline void listen() {
	CGEventMask eventMask = {
#ifdef DEBUG
		CGEventMaskBit(kCGEventKeyDown) |
#endif
		CGEventMaskBit(kCGEventFlagsChanged)
	};

	CFMachPortRef eventTap = CGEventTapCreate(
		kCGSessionEventTap, kCGHeadInsertEventTap, 0, eventMask, CGEventCallback, NULL
	);

	if(!eventTap) {
		fprintf(stderr, "ERROR: Unable to create event tap.\n");
		exit(1);
	}

	CFRunLoopSourceRef runLoopSource = CFMachPortCreateRunLoopSource(kCFAllocatorDefault, eventTap, 0);
	CFRunLoopAddSource(CFRunLoopGetCurrent(), runLoopSource, kCFRunLoopCommonModes);
	CGEventTapEnable(eventTap, true);

	CFRunLoopRun();
}

static inline CGEventRef CGEventCallback(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon) {
	if (
#ifdef DEBUG
		type != kCGEventKeyDown &&
#endif
		type != kCGEventFlagsChanged) { return event; }

	CGKeyCode keyCode = (CGKeyCode) CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);
	CGEventFlags modifiers = CGEventGetFlags(event);

#ifdef DEBUG
		CGKeyCode keycode = (CGKeyCode) CGEventGetIntegerValueField(event, kCGKeyboardEventKeycode);
		fprintf(stderr, ">>> KeyCode: %d Type: %d Flags: %llu\n", keycode, type, modifiers);
#endif

	if ((keyCode == LEFT_SHIFT || keyCode == LEFT_OPT) && modifiers == DOWN_SHIFT_AND_OPT ){
		CGEventRef enterKey = CGEventCreateKeyboardEvent(NULL, (CGKeyCode)KEY_SPACE, true);
		CGEventSetFlags(enterKey, (CGEventFlags)DOWN_CTRL_OPT);
		CGEventSetType(enterKey, kCGEventKeyDown);
		CGEventTapPostEvent(proxy, enterKey);
		CGEventSetType(enterKey, kCGEventKeyUp);
		CGEventTapPostEvent(proxy, enterKey);
		CFRelease(enterKey);
	}
	return event;
}
*/
import "C"

func main() {
	C.listen()
}
