package cmd

import (
	log "github.com/sirupsen/logrus"
)

func writeLog(sc int, target, method string, errorMsg string) {
	if sc != 200 {
		if sc == 401 || sc == 403 {
			logUnauth(sc, target, method, errorMsg)
		} else if sc == 301 || sc == 302 {
			logRedirect(sc, target, method)
		} else if sc == 0 {
			logBad(sc, target, method)
		} else if sc == 404 {
			logNotFound(sc, target, method, errorMsg)
		} else if sc == 1 {
			logDangerous(target, method)
		} else {
			logManual(sc, target, method, errorMsg)
		}
	} else {
		logAccessible(sc, target, method)
	}
}

func logAccessible(status int, target, method string) {
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Print("Endpoint accessible!")
}

func logDangerous(target, method string) {
	log.WithFields(log.Fields{
		"Status": "skipped",
		"Target": target,
		"Method": method,
	}).Warn("Endpoint skipped due to dangerous keyword.")
}

func logManual(status int, target, method, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "Manual testing may be required."
	}
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Warn(errorMsg)
}

func logNotFound(status int, target, method, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "Endpoint not found."
	}
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Error(errorMsg)
}

func logRedirect(status int, target, method string) {
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Error("Redirect detected. This likely requires authentication.")
}

func logBad(status int, target, method string) {
	log.WithFields(log.Fields{
		"Status": "N/A",
		"Target": target,
		"Method": method,
	}).Warn("Bad request (could not reach the target).")
}

func logUnauth(status int, target, method, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "Unauthorized."
	}
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Error(errorMsg)
}
