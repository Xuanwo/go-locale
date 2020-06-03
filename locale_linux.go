// +build !unit_test

package locale

var detectors = []detector{
	detectViaEnvLanguage,
	detectViaEnvLc,
	detectViaLocaleConf,
}
