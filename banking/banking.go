package banking

import (
	aqbanking "github.com/umsatz/go-aqbanking"
)

/*
#cgo LDFLAGS: -laqbanking
#cgo LDFLAGS: -lgwenhywfar
#cgo darwin CFLAGS: -I/usr/local/include/gwenhywfar4
#cgo darwin CFLAGS: -I/usr/local/include/aqbanking5
#cgo linux CFLAGS: -I/usr/include/gwenhywfar4
#cgo linux CFLAGS: -I/usr/include/aqbanking5
#include <gwenhywfar/cgui.h>
#include <aqbanking/abgui.h>
#include <aqhbci/aqhbci.h>
#include <gwenhywfar/gwenhywfar.h>
*/
import "C"

func Init(verbose, debug bool) (*aqbanking.AQBanking, error) {
	aq, err := aqbanking.DefaultAQBanking()
	if err != nil {
		return nil, err
	}

	logLevel := C.GWEN_LoggerLevel_Critical
	if verbose {
		logLevel = C.GWEN_LoggerLevel_Error
	}
	if debug {
		logLevel = C.GWEN_LoggerLevel_Warning
	}

	C.GWEN_Logger_SetLevel(C.CString(C.GWEN_LOGDOMAIN), C.GWEN_LOGGER_LEVEL(logLevel))
	C.GWEN_Logger_SetLevel(C.CString(C.AQBANKING_LOGDOMAIN), C.GWEN_LOGGER_LEVEL(logLevel))
	C.GWEN_Logger_SetLevel(C.CString(C.AQHBCI_LOGDOMAIN), C.GWEN_LOGGER_LEVEL(logLevel))

	return aq, nil
}
