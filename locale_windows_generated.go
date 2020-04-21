// +build !unit_test

package locale

// osLanguageCode is a mapping from Microsoft Windows language code to language.Tag
// which genereated via internal/cmd/languagecode, data is from microsoft openspecs.
//
// Microsoft will assign 0x1000 to languages that doesn't have LCID, application should
// handle this, and we will return Und instead.
//
// ref:
//   - https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-operatingsystem
//   - https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry
//   - https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/a9eac961-e77d-41a6-90a5-ce1a8b0cdb9c
var osLanguageCode = map[uint32]string{
	0x0036: "af",             // Afrikaans - , supported from Release 7
	0x0436: "af-ZA",          // Afrikaans - South Africa, supported from Release B
	0x001C: "sq",             // Albanian - , supported from Release 7
	0x041C: "sq-AL",          // Albanian - Albania, supported from Release B
	0x0084: "gsw",            // Alsatian - , supported from Release 7
	0x0484: "gsw-FR",         // Alsatian - France, supported from Release V
	0x005E: "am",             // Amharic - , supported from Release 7
	0x045E: "am-ET",          // Amharic - Ethiopia, supported from Release V
	0x0001: "ar",             // Arabic - , supported from Release 7
	0x1401: "ar-DZ",          // Arabic - Algeria, supported from Release B
	0x3C01: "ar-BH",          // Arabic - Bahrain, supported from Release B
	0x0c01: "ar-EG",          // Arabic - Egypt, supported from Release B
	0x0801: "ar-IQ",          // Arabic - Iraq, supported from Release B
	0x2C01: "ar-JO",          // Arabic - Jordan, supported from Release B
	0x3401: "ar-KW",          // Arabic - Kuwait, supported from Release B
	0x3001: "ar-LB",          // Arabic - Lebanon, supported from Release B
	0x1001: "ar-LY",          // Arabic - Libya, supported from Release B
	0x1801: "ar-MA",          // Arabic - Morocco, supported from Release B
	0x2001: "ar-OM",          // Arabic - Oman, supported from Release B
	0x4001: "ar-QA",          // Arabic - Qatar, supported from Release B
	0x0401: "ar-SA",          // Arabic - Saudi Arabia, supported from Release B
	0x2801: "ar-SY",          // Arabic - Syria, supported from Release B
	0x1C01: "ar-TN",          // Arabic - Tunisia, supported from Release B
	0x3801: "ar-AE",          // Arabic - U.A.E., supported from Release B
	0x2401: "ar-YE",          // Arabic - Yemen, supported from Release B
	0x002B: "hy",             // Armenian - , supported from Release 7
	0x042B: "hy-AM",          // Armenian - Armenia, supported from Release C
	0x004D: "as",             // Assamese - , supported from Release 7
	0x044D: "as-IN",          // Assamese - India, supported from Release V
	0x742C: "az-Cyrl",        // Azerbaijani (Cyrillic) - , supported from Windows 7
	0x082C: "az-Cyrl-AZ",     // Azerbaijani (Cyrillic) - Azerbaijan, supported from Release C
	0x002C: "az",             // Azerbaijani (Latin) - , supported from Release 7
	0x782C: "az-Latn",        // Azerbaijani (Latin) - , supported from Windows 7
	0x042C: "az-Latn-AZ",     // Azerbaijani (Latin) - Azerbaijan, supported from Release C
	0x0045: "bn",             // Bangla - , supported from Release 7
	0x0845: "bn-BD",          // Bangla - Bangladesh, supported from Release V
	0x0445: "bn-IN",          // Bangla - India, supported from Release E1
	0x006D: "ba",             // Bashkir - , supported from Release 7
	0x046D: "ba-RU",          // Bashkir - Russia, supported from Release V
	0x002D: "eu",             // Basque - , supported from Release 7
	0x042D: "eu-ES",          // Basque - Spain, supported from Release B
	0x0023: "be",             // Belarusian - , supported from Release 7
	0x0423: "be-BY",          // Belarusian - Belarus, supported from Release B
	0x641A: "bs-Cyrl",        // Bosnian (Cyrillic) - , supported from Windows 7
	0x201A: "bs-Cyrl-BA",     // Bosnian (Cyrillic) - Bosnia and Herzegovina, supported from Release E1
	0x681A: "bs-Latn",        // Bosnian (Latin) - , supported from Windows 7
	0x781A: "bs",             // Bosnian (Latin) - , supported from Release 7
	0x141A: "bs-Latn-BA",     // Bosnian (Latin) - Bosnia and Herzegovina, supported from Release E1
	0x007E: "br",             // Breton - , supported from Release 7
	0x047E: "br-FR",          // Breton - France, supported from Release V
	0x0002: "bg",             // Bulgarian - , supported from Release 7
	0x0402: "bg-BG",          // Bulgarian - Bulgaria, supported from Release B
	0x0055: "my",             // Burmese - , supported from Release 8.1
	0x0455: "my-MM",          // Burmese - Myanmar, supported from Release 8.1
	0x0003: "ca",             // Catalan - , supported from Release 7
	0x0403: "ca-ES",          // Catalan - Spain, supported from Release B
	0x0092: "ku",             // Central Kurdish - , supported from Release 8
	0x7c92: "ku-Arab",        // Central Kurdish - , supported from Release 8
	0x0492: "ku-Arab-IQ",     // Central Kurdish - Iraq, supported from Release 8
	0x005C: "chr",            // Cherokee - , supported from Release 8
	0x7c5C: "chr-Cher",       // Cherokee - , supported from Release 8
	0x045C: "chr-Cher-US",    // Cherokee - United States, supported from Release 8
	0x0004: "zh-Hans",        // Chinese (Simplified) - , supported from Release A
	0x7804: "zh",             // Chinese (Simplified) - , supported from Windows 7
	0x0804: "zh-CN",          // Chinese (Simplified) - People's Republic of China, supported from Release A
	0x1004: "zh-SG",          // Chinese (Simplified) - Singapore, supported from Release A
	0x7C04: "zh-Hant",        // Chinese (Traditional) - , supported from Release A
	0x0C04: "zh-HK",          // Chinese (Traditional) - Hong Kong S.A.R., supported from Release A
	0x1404: "zh-MO",          // Chinese (Traditional) - Macao S.A.R., supported from Release D
	0x0404: "zh-TW",          // Chinese (Traditional) - Taiwan, supported from Release A
	0x0083: "co",             // Corsican - , supported from Release 7
	0x0483: "co-FR",          // Corsican - France, supported from Release V
	0x001A: "hr,",            // Croatian - , supported from Release 7
	0x041A: "hr-HR",          // Croatian - Croatia, supported from Release A
	0x101A: "hr-BA",          // Croatian (Latin) - Bosnia and Herzegovina, supported from Release E1
	0x0005: "cs",             // Czech - , supported from Release 7
	0x0405: "cs-CZ",          // Czech - Czech Republic, supported from Release A
	0x0006: "da",             // Danish - , supported from Release 7
	0x0406: "da-DK",          // Danish - Denmark, supported from Release A
	0x008C: "prs",            // Dari - , supported from Release 7
	0x048C: "prs-AF",         // Dari - Afghanistan, supported from Release V
	0x0065: "dv",             // Divehi - , supported from Release 7
	0x0465: "dv-MV",          // Divehi - Maldives, supported from Release D
	0x0013: "nl",             // Dutch - , supported from Release 7
	0x0813: "nl-BE",          // Dutch - Belgium, supported from Release A
	0x0413: "nl-NL",          // Dutch - Netherlands, supported from Release A
	0x0C51: "dz-BT",          // Dzongkha - Bhutan, supported from Release 10
	0x0009: "en",             // English - , supported from Release 7
	0x0C09: "en-AU",          // English - Australia, supported from Release A
	0x2809: "en-BZ",          // English - Belize, supported from Release B
	0x1009: "en-CA",          // English - Canada, supported from Release A
	0x2409: "en-029",         // English - Caribbean, supported from Release B
	0x3C09: "en-HK",          // English - Hong Kong, supported from Release 8.1
	0x4009: "en-IN",          // English - India, supported from Release V
	0x1809: "en-IE",          // English - Ireland, supported from Release A
	0x2009: "en-JM",          // English - Jamaica, supported from Release B
	0x4409: "en-MY",          // English - Malaysia, supported from Release V
	0x1409: "en-NZ",          // English - New Zealand, supported from Release A
	0x3409: "en-PH",          // English - Republic of the Philippines, supported from Release C
	0x4809: "en-SG",          // English - Singapore, supported from Release V
	0x1C09: "en-ZA",          // English - South Africa, supported from Release B
	0x2c09: "en-TT",          // English - Trinidad and Tobago, supported from Release B
	0x4C09: "en-AE",          // English - United Arab Emirates, supported from Release 10.5
	0x0809: "en-GB",          // English - United Kingdom, supported from Release A
	0x0409: "en-US",          // English - United States, supported from Release A
	0x3009: "en-ZW",          // English - Zimbabwe, supported from Release C
	0x0025: "et",             // Estonian - , supported from Release 7
	0x0425: "et-EE",          // Estonian - Estonia, supported from Release B
	0x0038: "fo",             // Faroese - , supported from Release 7
	0x0438: "fo-FO",          // Faroese - Faroe Islands, supported from Release B
	0x0064: "fil",            // Filipino - , supported from Release 7
	0x0464: "fil-PH",         // Filipino - Philippines, supported from Release E2
	0x000B: "fi",             // Finnish - , supported from Release 7
	0x040B: "fi-FI",          // Finnish - Finland, supported from Release A
	0x000C: "fr",             // French - , supported from Release 7
	0x080C: "fr-BE",          // French - Belgium, supported from Release A
	0x2c0C: "fr-CM",          // French - Cameroon, supported from Release 8.1
	0x0c0C: "fr-CA",          // French - Canada, supported from Release A
	0x240C: "fr-CD",          // French - Congo, DRC, supported from Release 8.1
	0x300C: "fr-CI",          // French - Côte d'Ivoire, supported from Release 8.1
	0x040C: "fr-FR",          // French - France, supported from Release A
	0x3c0C: "fr-HT",          // French - Haiti, supported from Release 8.1
	0x140C: "fr-LU",          // French - Luxembourg, supported from Release A
	0x340C: "fr-ML",          // French - Mali, supported from Release 8.1
	0x380C: "fr-MA",          // French - Morocco, supported from Release 8.1
	0x180C: "fr-MC",          // French - Principality of Monaco, supported from Release A
	0x200C: "fr-RE",          // French - Reunion, supported from Release 8.1
	0x280C: "fr-SN",          // French - Senegal, supported from Release 8.1
	0x100C: "fr-CH",          // French - Switzerland, supported from Release A
	0x0062: "fy",             // Frisian - , supported from Release 7
	0x0462: "fy-NL",          // Frisian - Netherlands, supported from Release E2
	0x0067: "ff",             // Fulah - , supported from Release 8
	0x7C67: "ff-Latn",        // Fulah (Latin) - , supported from Release 8
	0x0867: "ff-Latn-SN",     // Fulah - Senegal, supported from Release 8
	0x0056: "gl",             // Galician - , supported from Release 7
	0x0456: "gl-ES",          // Galician - Spain, supported from Release D
	0x0037: "ka",             // Georgian - , supported from Release 7
	0x0437: "ka-GE",          // Georgian - Georgia, supported from Release C
	0x0007: "de",             // German - , supported from Release 7
	0x0C07: "de-AT",          // German - Austria, supported from Release A
	0x0407: "de-DE",          // German - Germany, supported from Release A
	0x1407: "de-LI",          // German - Liechtenstein, supported from Release B
	0x1007: "de-LU",          // German - Luxembourg, supported from Release B
	0x0807: "de-CH",          // German - Switzerland, supported from Release A
	0x0008: "el",             // Greek - , supported from Release 7
	0x0408: "el-GR",          // Greek - Greece, supported from Release A
	0x006F: "kl",             // Greenlandic - , supported from Release 7
	0x046F: "kl-GL",          // Greenlandic - Greenland, supported from Release V
	0x0074: "gn",             // Guarani - , supported from Release 8.1
	0x0474: "gn-PY",          // Guarani - Paraguay, supported from Release 8.1
	0x0047: "gu",             // Gujarati - , supported from Release 7
	0x0447: "gu-IN",          // Gujarati - India, supported from Release D
	0x0068: "ha",             // Hausa (Latin) - , supported from Release 7
	0x7C68: "ha-Latn",        // Hausa (Latin) - , supported from Windows 7
	0x0468: "ha-Latn-NG",     // Hausa (Latin) - Nigeria, supported from Release V
	0x0075: "haw",            // Hawaiian - , supported from Release 8
	0x0475: "haw-US",         // Hawaiian - United States, supported from Release 8
	0x000D: "he",             // Hebrew - , supported from Release 7
	0x040D: "he-IL",          // Hebrew - Israel, supported from Release B
	0x0039: "hi",             // Hindi - , supported from Release 7
	0x0439: "hi-IN",          // Hindi - India, supported from Release C
	0x000E: "hu",             // Hungarian - , supported from Release 7
	0x040E: "hu-HU",          // Hungarian - Hungary, supported from Release A
	0x000F: "is",             // Icelandic - , supported from Release 7
	0x040F: "is-IS",          // Icelandic - Iceland, supported from Release A
	0x0070: "ig",             // Igbo - , supported from Release 7
	0x0470: "ig-NG",          // Igbo - Nigeria, supported from Release V
	0x0021: "id",             // Indonesian - , supported from Release 7
	0x0421: "id-ID",          // Indonesian - Indonesia, supported from Release B
	0x005D: "iu",             // Inuktitut (Latin) - , supported from Release 7
	0x7C5D: "iu-Latn",        // Inuktitut (Latin) - , supported from Windows 7
	0x085D: "iu-Latn-CA",     // Inuktitut (Latin) - Canada, supported from Release E2
	0x785D: "iu-Cans",        // Inuktitut (Syllabics) - , supported from Windows 7
	0x045d: "iu-Cans-CA",     // Inuktitut (Syllabics) - Canada, supported from Release V
	0x003C: "ga",             // Irish - , supported from Windows 7
	0x083C: "ga-IE",          // Irish - Ireland, supported from Release E2
	0x0010: "it",             // Italian - , supported from Release 7
	0x0410: "it-IT",          // Italian - Italy, supported from Release A
	0x0810: "it-CH",          // Italian - Switzerland, supported from Release A
	0x0011: "ja",             // Japanese - , supported from Release 7
	0x0411: "ja-JP",          // Japanese - Japan, supported from Release A
	0x004B: "kn",             // Kannada - , supported from Release 7
	0x044B: "kn-IN",          // Kannada - India, supported from Release D
	0x0060: "ks",             // Kashmiri - , supported from Release 10
	0x0460: "ks-Arab",        // Kashmiri - Perso-Arabic, supported from Release 10
	0x003F: "kk",             // Kazakh - , supported from Release 7
	0x043F: "kk-KZ",          // Kazakh - Kazakhstan, supported from Release C
	0x0053: "km",             // Khmer - , supported from Release 7
	0x0453: "km-KH",          // Khmer - Cambodia, supported from Release V
	0x0086: "quc",            // K'iche - , supported from Release 10
	0x0486: "quc-Latn-GT",    // K'iche - Guatemala, supported from Release 10
	0x0087: "rw",             // Kinyarwanda - , supported from Release 7
	0x0487: "rw-RW",          // Kinyarwanda - Rwanda, supported from Release V
	0x0041: "sw",             // Kiswahili - , supported from Release 7
	0x0441: "sw-KE",          // Kiswahili - Kenya, supported from Release C
	0x0057: "kok",            // Konkani - , supported from Release 7
	0x0457: "kok-IN",         // Konkani - India, supported from Release C
	0x0012: "ko",             // Korean - , supported from Release 7
	0x0412: "ko-KR",          // Korean - Korea, supported from Release A
	0x0040: "ky",             // Kyrgyz - , supported from Release 7
	0x0440: "ky-KG",          // Kyrgyz - Kyrgyzstan, supported from Release D
	0x0054: "lo",             // Lao - , supported from Release 7
	0x0454: "lo-LA",          // Lao - Lao P.D.R., supported from Release V
	0x0026: "lv",             // Latvian - , supported from Release 7
	0x0426: "lv-LV",          // Latvian - Latvia, supported from Release B
	0x0027: "lt",             // Lithuanian - , supported from Release 7
	0x0427: "lt-LT",          // Lithuanian - Lithuania, supported from Release B
	0x7C2E: "dsb",            // Lower Sorbian - , supported from Windows 7
	0x082E: "dsb-DE",         // Lower Sorbian - Germany, supported from Release V
	0x006E: "lb",             // Luxembourgish - , supported from Release 7
	0x046E: "lb-LU",          // Luxembourgish - Luxembourg, supported from Release E2
	0x002F: "mk",             // Macedonian - , supported from Release 7
	0x042F: "mk-MK",          // Macedonian - North Macedonia, supported from Release C
	0x003E: "ms",             // Malay - , supported from Release 7
	0x083E: "ms-BN",          // Malay - Brunei Darussalam, supported from Release C
	0x043E: "ms-MY",          // Malay - Malaysia, supported from Release C
	0x004C: "ml",             // Malayalam - , supported from Release 7
	0x044C: "ml-IN",          // Malayalam - India, supported from Release E1
	0x003A: "mt",             // Maltese - , supported from Release 7
	0x043A: "mt-MT",          // Maltese - Malta, supported from Release E1
	0x0081: "mi",             // Maori - , supported from Release 7
	0x0481: "mi-NZ",          // Maori - New Zealand, supported from Release E1
	0x007A: "arn",            // Mapudungun - , supported from Release 7
	0x047A: "arn-CL",         // Mapudungun - Chile, supported from Release E2
	0x004E: "mr",             // Marathi - , supported from Release 7
	0x044E: "mr-IN",          // Marathi - India, supported from Release C
	0x007C: "moh",            // Mohawk - , supported from Release 7
	0x047C: "moh-CA",         // Mohawk - Canada, supported from Release E2
	0x0050: "mn",             // Mongolian (Cyrillic) - , supported from Release 7
	0x7850: "mn-Cyrl",        // Mongolian (Cyrillic) - , supported from Windows 7
	0x0450: "mn-MN",          // Mongolian (Cyrillic) - Mongolia, supported from Release D
	0x7C50: "mn-Mong",        // Mongolian (Traditional Mongolian) - , supported from Windows 7
	0x0850: "mn-Mong-CN",     // Mongolian (Traditional Mongolian) - People's Republic of China, supported from Release V
	0x0C50: "mn-Mong-MN",     // Mongolian (Traditional Mongolian) - Mongolia, supported from Windows 7
	0x0061: "ne",             // Nepali - , supported from Release 7
	0x0861: "ne-IN",          // Nepali - India, supported from Release 8.1
	0x0461: "ne-NP",          // Nepali - Nepal, supported from Release E2
	0x0014: "no",             // Norwegian (Bokmal) - , supported from Release 7
	0x7C14: "nb",             // Norwegian (Bokmal) - , supported from Release 7
	0x0414: "nb-NO",          // Norwegian (Bokmal) - Norway, supported from Release A
	0x7814: "nn",             // Norwegian (Nynorsk) - , supported from Release 7
	0x0814: "nn-NO",          // Norwegian (Nynorsk) - Norway, supported from Release A
	0x0082: "oc",             // Occitan - , supported from Release 7
	0x0482: "oc-FR",          // Occitan - France, supported from Release V
	0x0048: "or",             // Odia - , supported from Release 7
	0x0448: "or-IN",          // Odia - India, supported from Release V
	0x0072: "om",             // Oromo - , supported from Release 8.1
	0x0472: "om-ET",          // Oromo - Ethiopia, supported from Release 8.1
	0x0063: "ps",             // Pashto - , supported from Release 7
	0x0463: "ps-AF",          // Pashto - Afghanistan, supported from Release E2
	0x0029: "fa",             // Persian - , supported from Release 7
	0x0429: "fa-IR",          // Persian - Iran, supported from Release B
	0x0015: "pl",             // Polish - , supported from Release 7
	0x0415: "pl-PL",          // Polish - Poland, supported from Release A
	0x0016: "pt",             // Portuguese - , supported from Release 7
	0x0416: "pt-BR",          // Portuguese - Brazil, supported from Release A
	0x0816: "pt-PT",          // Portuguese - Portugal, supported from Release A
	0x05FE: "qps-ploca",      // Pseudo Language - Pseudo locale for east Asian/complex script  localization testing, supported from Release 7
	0x0501: "qps-ploc",       // Pseudo Language - Pseudo locale used for localization testing, supported from Release 7
	0x09FF: "qps-plocm",      // Pseudo Language - Pseudo locale used for localization testing of  mirrored locales, supported from Release 7
	0x0046: "pa",             // Punjabi - , supported from Release 7
	0x7C46: "pa-Arab",        // Punjabi - , supported from Release 8
	0x0446: "pa-IN",          // Punjabi - India, supported from Release D
	0x0846: "pa-Arab-PK",     // Punjabi - Islamic Republic of Pakistan, supported from Release 8
	0x006B: "quz",            // Quechua - , supported from Release 7
	0x046B: "quz-BO",         // Quechua - Bolivia, supported from Release E1
	0x086B: "quz-EC",         // Quechua - Ecuador, supported from Release E1
	0x0C6B: "quz-PE",         // Quechua - Peru, supported from Release E1
	0x0018: "ro",             // Romanian - , supported from Release 7
	0x0818: "ro-MD",          // Romanian - Moldova, supported from Release 8.1
	0x0418: "ro-RO",          // Romanian - Romania, supported from Release A
	0x0017: "rm",             // Romansh - , supported from Release 7
	0x0417: "rm-CH",          // Romansh - Switzerland, supported from Release E2
	0x0019: "ru",             // Russian - , supported from Release 7
	0x0819: "ru-MD",          // Russian - Moldova, supported from Release 10
	0x0419: "ru-RU",          // Russian - Russia, supported from Release A
	0x0085: "sah",            // Sakha - , supported from Release 7
	0x0485: "sah-RU",         // Sakha - Russia, supported from Release V
	0x703B: "smn",            // Sami (Inari) - , supported from Windows 7
	0x243B: "smn-FI",         // Sami (Inari) - Finland, supported from Release E1
	0x7C3B: "smj",            // Sami (Lule) - , supported from Windows 7
	0x103B: "smj-NO",         // Sami (Lule) - Norway, supported from Release E1
	0x143B: "smj-SE",         // Sami (Lule) - Sweden, supported from Release E1
	0x003B: "se",             // Sami (Northern) - , supported from Release 7
	0x0C3B: "se-FI",          // Sami (Northern) - Finland, supported from Release E1
	0x043B: "se-NO",          // Sami (Northern) - Norway, supported from Release E1
	0x083B: "se-SE",          // Sami (Northern) - Sweden, supported from Release E1
	0x743B: "sms",            // Sami (Skolt) - , supported from Windows 7
	0x203B: "sms-FI",         // Sami (Skolt) - Finland, supported from Release E1
	0x783B: "sma",            // Sami (Southern) - , supported from Windows 7
	0x183B: "sma-NO",         // Sami (Southern) - Norway, supported from Release E1
	0x1C3B: "sma-SE",         // Sami (Southern) - Sweden, supported from Release E1
	0x004F: "sa",             // Sanskrit - , supported from Release 7
	0x044F: "sa-IN",          // Sanskrit - India, supported from Release C
	0x0091: "gd",             // Scottish Gaelic - , supported from Windows 7
	0x0491: "gd-GB",          // Scottish Gaelic - United Kingdom, supported from Release 7
	0x6C1A: "sr-Cyrl",        // Serbian (Cyrillic) - , supported from Windows 7
	0x1C1A: "sr-Cyrl-BA",     // Serbian (Cyrillic) - Bosnia and Herzegovina, supported from Release E1
	0x301A: "sr-Cyrl-ME",     // Serbian (Cyrillic) - Montenegro, supported from Release 7
	0x281A: "sr-Cyrl-RS",     // Serbian (Cyrillic) - Serbia, supported from Release 7
	0x0C1A: "sr-Cyrl-CS",     // Serbian (Cyrillic) - Serbia and Montenegro (Former), supported from Release B
	0x701A: "sr-Latn",        // Serbian (Latin) - , supported from Windows 7
	0x7C1A: "sr",             // Serbian (Latin) - , supported from Release 7
	0x181A: "sr-Latn-BA",     // Serbian (Latin) - Bosnia and Herzegovina, supported from Release E1
	0x2c1A: "sr-Latn-ME",     // Serbian (Latin) - Montenegro, supported from Release 7
	0x241A: "sr-Latn-RS",     // Serbian (Latin) - Serbia, supported from Release 7
	0x081A: "sr-Latn-CS",     // Serbian (Latin) - Serbia and Montenegro (Former), supported from Release B
	0x006C: "nso",            // Sesotho sa Leboa - , supported from Release 7
	0x046C: "nso-ZA",         // Sesotho sa Leboa - South Africa, supported from Release E1
	0x0032: "tn",             // Setswana - , supported from Release 7
	0x0832: "tn-BW",          // Setswana - Botswana, supported from Release 8
	0x0432: "tn-ZA",          // Setswana - South Africa, supported from Release E1
	0x0059: "sd",             // Sindhi - , supported from Release 8
	0x7C59: "sd-Arab",        // Sindhi - , supported from Release 8
	0x0859: "sd-Arab-PK",     // Sindhi - Islamic Republic of Pakistan, supported from Release 8
	0x005B: "si",             // Sinhala - , supported from Release 7
	0x045B: "si-LK",          // Sinhala - Sri Lanka, supported from Release V
	0x001B: "sk",             // Slovak - , supported from Release 7
	0x041B: "sk-SK",          // Slovak - Slovakia, supported from Release A
	0x0024: "sl",             // Slovenian - , supported from Release 7
	0x0424: "sl-SI",          // Slovenian - Slovenia, supported from Release A
	0x0077: "so",             // Somali - , supported from Release 8.1
	0x0477: "so-SO",          // Somali - Somalia, supported from Release 8.1
	0x0030: "st",             // Sotho - , supported from Release 8.1
	0x0430: "st-ZA",          // Sotho - South Africa, supported from Release 8.1
	0x000A: "es",             // Spanish - , supported from Release 7
	0x2C0A: "es-AR",          // Spanish - Argentina, supported from Release B
	0x200A: "es-VE",          // Spanish - Bolivarian Republic of Venezuela, supported from Release B
	0x400A: "es-BO",          // Spanish - Bolivia, supported from Release B
	0x340A: "es-CL",          // Spanish - Chile, supported from Release B
	0x240A: "es-CO",          // Spanish - Colombia, supported from Release B
	0x140A: "es-CR",          // Spanish - Costa Rica, supported from Release B
	0x5c0A: "es-CU",          // Spanish - Cuba, supported from Release 10
	0x1c0A: "es-DO",          // Spanish - Dominican Republic, supported from Release B
	0x300A: "es-EC",          // Spanish - Ecuador, supported from Release B
	0x440A: "es-SV",          // Spanish - El Salvador, supported from Release B
	0x100A: "es-GT",          // Spanish - Guatemala, supported from Release B
	0x480A: "es-HN",          // Spanish - Honduras, supported from Release B
	0x580A: "es-419",         // Spanish - Latin America, supported from Release 8.1
	0x080A: "es-MX",          // Spanish - Mexico, supported from Release A
	0x4C0A: "es-NI",          // Spanish - Nicaragua, supported from Release B
	0x180A: "es-PA",          // Spanish - Panama, supported from Release B
	0x3C0A: "es-PY",          // Spanish - Paraguay, supported from Release B
	0x280A: "es-PE",          // Spanish - Peru, supported from Release B
	0x500A: "es-PR",          // Spanish - Puerto Rico, supported from Release B
	0x040A: "es-ES_tradnl",   // Spanish - Spain, supported from Release A
	0x0c0A: "es-ES",          // Spanish - Spain, supported from Release A
	0x540A: "es-US",          // Spanish - United States, supported from Release V
	0x380A: "es-UY",          // Spanish - Uruguay, supported from Release B
	0x001D: "sv",             // Swedish - , supported from Release 7
	0x081D: "sv-FI",          // Swedish - Finland, supported from Release B
	0x041D: "sv-SE",          // Swedish - Sweden, supported from Release A
	0x005A: "syr",            // Syriac - , supported from Release 7
	0x045A: "syr-SY",         // Syriac - Syria, supported from Release D
	0x0028: "tg",             // Tajik (Cyrillic) - , supported from Release 7
	0x7C28: "tg-Cyrl",        // Tajik (Cyrillic) - , supported from Windows 7
	0x0428: "tg-Cyrl-TJ",     // Tajik (Cyrillic) - Tajikistan, supported from Release V
	0x005F: "tzm",            // Tamazight (Latin) - , supported from Release 7
	0x7C5F: "tzm-Latn",       // Tamazight (Latin) - , supported from Windows 7
	0x085F: "tzm-Latn-DZ",    // Tamazight (Latin) - Algeria, supported from Release V
	0x0049: "ta",             // Tamil - , supported from Release 7
	0x0449: "ta-IN",          // Tamil - India, supported from Release C
	0x0849: "ta-LK",          // Tamil - Sri Lanka, supported from Release 8
	0x0044: "tt",             // Tatar - , supported from Release 7
	0x0444: "tt-RU",          // Tatar - Russia, supported from Release D
	0x004A: "te",             // Telugu - , supported from Release 7
	0x044A: "te-IN",          // Telugu - India, supported from Release D
	0x001E: "th",             // Thai - , supported from Release 7
	0x041E: "th-TH",          // Thai - Thailand, supported from Release B
	0x0051: "bo",             // Tibetan - , supported from Release 7
	0x0451: "bo-CN",          // Tibetan - People's Republic of China, supported from Release V
	0x0073: "ti",             // Tigrinya - , supported from Release 8
	0x0873: "ti-ER",          // Tigrinya - Eritrea, supported from Release 8
	0x0473: "ti-ET",          // Tigrinya - Ethiopia, supported from Release 8
	0x0031: "ts",             // Tsonga - , supported from Release 8.1
	0x0431: "ts-ZA",          // Tsonga - South Africa, supported from Release 8.1
	0x001F: "tr",             // Turkish - , supported from Release 7
	0x041F: "tr-TR",          // Turkish - Turkey, supported from Release A
	0x0042: "tk",             // Turkmen - , supported from Release 7
	0x0442: "tk-TM",          // Turkmen - Turkmenistan, supported from Release V
	0x0022: "uk",             // Ukrainian - , supported from Release 7
	0x0422: "uk-UA",          // Ukrainian - Ukraine, supported from Release B
	0x002E: "hsb",            // Upper Sorbian - , supported from Release 7
	0x042E: "hsb-DE",         // Upper Sorbian - Germany, supported from Release V
	0x0020: "ur",             // Urdu - , supported from Release 7
	0x0820: "ur-IN",          // Urdu - India, supported from Release 8.1
	0x0420: "ur-PK",          // Urdu - Islamic Republic of Pakistan, supported from Release C
	0x0080: "ug",             // Uyghur - , supported from Release 7
	0x0480: "ug-CN",          // Uyghur - People's Republic of China, supported from Release V
	0x7843: "uz-Cyrl",        // Uzbek (Cyrillic) - , supported from Windows 7
	0x0843: "uz-Cyrl-UZ",     // Uzbek (Cyrillic) - Uzbekistan, supported from Release C
	0x0043: "uz",             // Uzbek (Latin) - , supported from Release 7
	0x7C43: "uz-Latn",        // Uzbek (Latin) - , supported from Windows 7
	0x0443: "uz-Latn-UZ",     // Uzbek (Latin) - Uzbekistan, supported from Release C
	0x0803: "ca-ES-valencia", // Valencian - Spain, supported from Release 8
	0x0033: "ve",             // Venda - , supported from Release 10
	0x0433: "ve-ZA",          // Venda - South Africa, supported from Release 10
	0x002A: "vi",             // Vietnamese - , supported from Release 7
	0x042A: "vi-VN",          // Vietnamese - Vietnam, supported from Release B
	0x0052: "cy",             // Welsh - , supported from Release 7
	0x0452: "cy-GB",          // Welsh - United Kingdom, supported from Release E1
	0x0088: "wo",             // Wolof - , supported from Release 7
	0x0488: "wo-SN",          // Wolof - Senegal, supported from Release V
	0x0034: "xh",             // Xhosa - , supported from Release 7
	0x0434: "xh-ZA",          // Xhosa - South Africa, supported from Release E1
	0x0078: "ii",             // Yi - , supported from Release 7
	0x0478: "ii-CN",          // Yi - People's Republic of China, supported from Release V
	0x006A: "yo",             // Yoruba - , supported from Release 7
	0x046A: "yo-NG",          // Yoruba - Nigeria, supported from Release V
	0x0035: "zu",             // Zulu - , supported from Release 7
	0x0435: "zu-ZA",          // Zulu - South Africa, supported from Release E1

	0x1000: "Und",
}
