package testdata

/*
#include <stdint.h>
typedef unsigned char custom_uchar_t;

char            *ncp = 0;
char            *cp = "test";
char             ca[6] = {'t', 'e', 's', 't', '2', '\0'};
unsigned char    uca[6] = {'t', 'e', 's', 't', '3', '\0'};
signed char      sca[6] = {'t', 'e', 's', 't', '4', '\0'};
uint8_t          ui8ta[6] = {'t', 'e', 's', 't', '5', '\0'};
custom_uchar_t   tuca[6] = {'t', 'e', 's', 't', '6', '\0'};
*/
import "C"

// GetCgoNullCharPointer returns a null char pointer via cgo.  This is only
// used for tests.
func GetCgoNullCharPointer() interface{} {
	return C.ncp
}

// GetCgoCharPointer returns a char pointer via cgo.  This is only used for
// tests.
func GetCgoCharPointer() interface{} {
	return C.cp
}

// GetCgoCharArray returns a char array via cgo and the array's len and cap.
// This is only used for tests.
func GetCgoCharArray() (interface{}, int, int) {
	return C.ca, len(C.ca), cap(C.ca)
}

// GetCgoUnsignedCharArray returns an unsigned char array via cgo and the
// array's len and cap.  This is only used for tests.
func GetCgoUnsignedCharArray() (interface{}, int, int) {
	return C.uca, len(C.uca), cap(C.uca)
}

// GetCgoSignedCharArray returns a signed char array via cgo and the array's len
// and cap.  This is only used for tests.
func GetCgoSignedCharArray() (interface{}, int, int) {
	return C.sca, len(C.sca), cap(C.sca)
}

// GetCgoUint8tArray returns a uint8_t array via cgo and the array's len and
// cap.  This is only used for tests.
func GetCgoUint8tArray() (interface{}, int, int) {
	return C.ui8ta, len(C.ui8ta), cap(C.ui8ta)
}

// GetCgoTypdefedUnsignedCharArray returns a typedefed unsigned char array via
// cgo and the array's len and cap.  This is only used for tests.
func GetCgoTypdefedUnsignedCharArray() (interface{}, int, int) {
	return C.tuca, len(C.tuca), cap(C.tuca)
}
