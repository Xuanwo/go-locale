// +build aix dragonfly freebsd hurd illumos linux nacl netbsd openbsd plan9 solaris zos
// +build !tests

package locale

var detectors = []detector{
	detectViaEnvLanguage,
	detectViaEnvLc,
	detectViaLocaleConf,
}
