package collation

import "strconv"

// Collation is a mysql collation
type Collation uint64

// IsBinary returns whether this Collation represents the binary collation
func (c Collation) IsBinary() bool {
	return c == Binary
}

func (c Collation) String() string {
	if n, ok := collationToName[c]; ok {
		return n
	}
	return strconv.FormatUint(uint64(c), 10)
}

const (
	Big5ChineseCi         Collation = 1
	Latin2CzechCs                   = 2
	Dec8SwedishCi                   = 3
	Cp850GeneralCi                  = 4
	Latin1German1Ci                 = 5
	Hp8EnglishCi                    = 6
	Koi8rGeneralCi                  = 7
	Latin1SwedishCi                 = 8
	Latin2GeneralCi                 = 9
	Swe7SwedishCi                   = 10
	ASCIIGeneralCi                  = 11
	UjisJapaneseCi                  = 12
	SjisJapaneseCi                  = 13
	Cp1251BulgarianCi               = 14
	Latin1DanishCi                  = 15
	HebrewGeneralCi                 = 16
	Tis620ThaiCi                    = 18
	EuckrKoreanCi                   = 19
	Latin7EstonianCs                = 20
	Latin2HungarianCi               = 21
	Koi8uGeneralCi                  = 22
	Cp1251UkrainianCi               = 23
	Gb2312ChineseCi                 = 24
	GreekGeneralCi                  = 25
	Cp1250GeneralCi                 = 26
	Latin2CroatianCi                = 27
	GbkChineseCi                    = 28
	Cp1257LithuanianCi              = 29
	Latin5TurkishCi                 = 30
	Latin1German2Ci                 = 31
	Armscii8GeneralCi               = 32
	Utf8GeneralCi                   = 33
	Cp1250CzechCs                   = 34
	Ucs2GeneralCi                   = 35
	Cp866GeneralCi                  = 36
	Keybcs2GeneralCi                = 37
	MacceGeneralCi                  = 38
	MacromanGeneralCi               = 39
	Cp852GeneralCi                  = 40
	Latin7GeneralCi                 = 41
	Latin7GeneralCs                 = 42
	MacceBin                        = 43
	Cp1250CroatianCi                = 44
	Utf8mb4GeneralCi                = 45
	Utf8mb4Bin                      = 46
	Latin1Bin                       = 47
	Latin1GeneralCi                 = 48
	Latin1GeneralCs                 = 49
	Cp1251Bin                       = 50
	Cp1251GeneralCi                 = 51
	Cp1251GeneralCs                 = 52
	MacromanBin                     = 53
	Utf16GeneralCi                  = 54
	Utf16Bin                        = 55
	Utf16leGeneralCi                = 56
	Cp1256GeneralCi                 = 57
	Cp1257Bin                       = 58
	Cp1257GeneralCi                 = 59
	Utf32GeneralCi                  = 60
	Utf32Bin                        = 61
	Utf16leBin                      = 62
	Binary                          = 63
	Armscii8Bin                     = 64
	ASCIIBin                        = 65
	Cp1250Bin                       = 66
	Cp1256Bin                       = 67
	Cp866Bin                        = 68
	Dec8Bin                         = 69
	GreekBin                        = 70
	HebrewBin                       = 71
	Hp8Bin                          = 72
	Keybcs2Bin                      = 73
	Koi8rBin                        = 74
	Koi8uBin                        = 75
	Utf8TolowerCi                   = 76
	Latin2Bin                       = 77
	Latin5Bin                       = 78
	Latin7Bin                       = 79
	Cp850Bin                        = 80
	Cp852Bin                        = 81
	Swe7Bin                         = 82
	Utf8Bin                         = 83
	Big5Bin                         = 84
	EuckrBin                        = 85
	Gb2312Bin                       = 86
	GbkBin                          = 87
	SjisBin                         = 88
	Tis620Bin                       = 89
	Ucs2Bin                         = 90
	UjisBin                         = 91
	Geostd8GeneralCi                = 92
	Geostd8Bin                      = 93
	Latin1SpanishCi                 = 94
	Cp932JapaneseCi                 = 95
	Cp932Bin                        = 96
	EucjpmsJapaneseCi               = 97
	EucjpmsBin                      = 98
	Cp1250PolishCi                  = 99
	Utf16UnicodeCi                  = 101
	Utf16IcelandicCi                = 102
	Utf16LatvianCi                  = 103
	Utf16RomanianCi                 = 104
	Utf16SlovenianCi                = 105
	Utf16PolishCi                   = 106
	Utf16EstonianCi                 = 107
	Utf16SpanishCi                  = 108
	Utf16SwedishCi                  = 109
	Utf16TurkishCi                  = 110
	Utf16CzechCi                    = 111
	Utf16DanishCi                   = 112
	Utf16LithuanianCi               = 113
	Utf16SlovakCi                   = 114
	Utf16Spanish2Ci                 = 115
	Utf16RomanCi                    = 116
	Utf16PersianCi                  = 117
	Utf16EsperantoCi                = 118
	Utf16HungarianCi                = 119
	Utf16SinhalaCi                  = 120
	Utf16German2Ci                  = 121
	Utf16CroatianCi                 = 122
	Utf16Unicode520Ci               = 123
	Utf16VietnameseCi               = 124
	Ucs2UnicodeCi                   = 128
	Ucs2IcelandicCi                 = 129
	Ucs2LatvianCi                   = 130
	Ucs2RomanianCi                  = 131
	Ucs2SlovenianCi                 = 132
	Ucs2PolishCi                    = 133
	Ucs2EstonianCi                  = 134
	Ucs2SpanishCi                   = 135
	Ucs2SwedishCi                   = 136
	Ucs2TurkishCi                   = 137
	Ucs2CzechCi                     = 138
	Ucs2DanishCi                    = 139
	Ucs2LithuanianCi                = 140
	Ucs2SlovakCi                    = 141
	Ucs2Spanish2Ci                  = 142
	Ucs2RomanCi                     = 143
	Ucs2PersianCi                   = 144
	Ucs2EsperantoCi                 = 145
	Ucs2HungarianCi                 = 146
	Ucs2SinhalaCi                   = 147
	Ucs2German2Ci                   = 148
	Ucs2CroatianCi                  = 149
	Ucs2Unicode520Ci                = 150
	Ucs2VietnameseCi                = 151
	Ucs2GeneralMysql500Ci           = 159
	Utf32UnicodeCi                  = 160
	Utf32IcelandicCi                = 161
	Utf32LatvianCi                  = 162
	Utf32RomanianCi                 = 163
	Utf32SlovenianCi                = 164
	Utf32PolishCi                   = 165
	Utf32EstonianCi                 = 166
	Utf32SpanishCi                  = 167
	Utf32SwedishCi                  = 168
	Utf32TurkishCi                  = 169
	Utf32CzechCi                    = 170
	Utf32DanishCi                   = 171
	Utf32LithuanianCi               = 172
	Utf32SlovakCi                   = 173
	Utf32Spanish2Ci                 = 174
	Utf32RomanCi                    = 175
	Utf32PersianCi                  = 176
	Utf32EsperantoCi                = 177
	Utf32HungarianCi                = 178
	Utf32SinhalaCi                  = 179
	Utf32German2Ci                  = 180
	Utf32CroatianCi                 = 181
	Utf32Unicode520Ci               = 182
	Utf32VietnameseCi               = 183
	Utf8UnicodeCi                   = 192
	Utf8IcelandicCi                 = 193
	Utf8LatvianCi                   = 194
	Utf8RomanianCi                  = 195
	Utf8SlovenianCi                 = 196
	Utf8PolishCi                    = 197
	Utf8EstonianCi                  = 198
	Utf8SpanishCi                   = 199
	Utf8SwedishCi                   = 200
	Utf8TurkishCi                   = 201
	Utf8CzechCi                     = 202
	Utf8DanishCi                    = 203
	Utf8LithuanianCi                = 204
	Utf8SlovakCi                    = 205
	Utf8Spanish2Ci                  = 206
	Utf8RomanCi                     = 207
	Utf8PersianCi                   = 208
	Utf8EsperantoCi                 = 209
	Utf8HungarianCi                 = 210
	Utf8SinhalaCi                   = 211
	Utf8German2Ci                   = 212
	Utf8CroatianCi                  = 213
	Utf8Unicode520Ci                = 214
	Utf8VietnameseCi                = 215
	Utf8GeneralMysql500Ci           = 223
	Utf8mb4UnicodeCi                = 224
	Utf8mb4IcelandicCi              = 225
	Utf8mb4LatvianCi                = 226
	Utf8mb4RomanianCi               = 227
	Utf8mb4SlovenianCi              = 228
	Utf8mb4PolishCi                 = 229
	Utf8mb4EstonianCi               = 230
	Utf8mb4SpanishCi                = 231
	Utf8mb4SwedishCi                = 232
	Utf8mb4TurkishCi                = 233
	Utf8mb4CzechCi                  = 234
	Utf8mb4DanishCi                 = 235
	Utf8mb4LithuanianCi             = 236
	Utf8mb4SlovakCi                 = 237
	Utf8mb4Spanish2Ci               = 238
	Utf8mb4RomanCi                  = 239
	Utf8mb4PersianCi                = 240
	Utf8mb4EsperantoCi              = 241
	Utf8mb4HungarianCi              = 242
	Utf8mb4SinhalaCi                = 243
	Utf8mb4German2Ci                = 244
	Utf8mb4CroatianCi               = 245
	Utf8mb4Unicode520Ci             = 246
	Utf8mb4VietnameseCi             = 247
	Gb18030ChineseCi                = 248
	Gb18030Bin                      = 249
	Gb18030Unicode520Ci             = 250
	Utf8mb40900AiCi                 = 255
	Utf8mb4DePb0900AiCi             = 256
	Utf8mb4Is0900AiCi               = 257
	Utf8mb4Lv0900AiCi               = 258
	Utf8mb4Ro0900AiCi               = 259
	Utf8mb4Sl0900AiCi               = 260
	Utf8mb4Pl0900AiCi               = 261
	Utf8mb4Et0900AiCi               = 262
	Utf8mb4Es0900AiCi               = 263
	Utf8mb4Sv0900AiCi               = 264
	Utf8mb4Tr0900AiCi               = 265
	Utf8mb4Cs0900AiCi               = 266
	Utf8mb4Da0900AiCi               = 267
	Utf8mb4Lt0900AiCi               = 268
	Utf8mb4Sk0900AiCi               = 269
	Utf8mb4EsTrad0900AiCi           = 270
	Utf8mb4La0900AiCi               = 271
	Utf8mb4Eo0900AiCi               = 273
	Utf8mb4Hu0900AiCi               = 274
	Utf8mb4Hr0900AiCi               = 275
	Utf8mb4Vi0900AiCi               = 277
	Utf8mb40900AsCs                 = 278
	Utf8mb4DePb0900AsCs             = 279
	Utf8mb4Is0900AsCs               = 280
	Utf8mb4Lv0900AsCs               = 281
	Utf8mb4Ro0900AsCs               = 282
	Utf8mb4Sl0900AsCs               = 283
	Utf8mb4Pl0900AsCs               = 284
	Utf8mb4Et0900AsCs               = 285
	Utf8mb4Es0900AsCs               = 286
	Utf8mb4Sv0900AsCs               = 287
	Utf8mb4Tr0900AsCs               = 288
	Utf8mb4Cs0900AsCs               = 289
	Utf8mb4Da0900AsCs               = 290
	Utf8mb4Lt0900AsCs               = 291
	Utf8mb4Sk0900AsCs               = 292
	Utf8mb4EsTrad0900AsCs           = 293
	Utf8mb4La0900AsCs               = 294
	Utf8mb4Eo0900AsCs               = 296
	Utf8mb4Hu0900AsCs               = 297
	Utf8mb4Hr0900AsCs               = 298
	Utf8mb4Vi0900AsCs               = 300
	Utf8mb4Ja0900AsCs               = 303
	Utf8mb4Ja0900AsCsKs             = 304
	Utf8mb40900AsCi                 = 305
	Utf8mb4Ru0900AiCi               = 306
	Utf8mb4Ru0900AsCs               = 307
)

var nameToCollation = map[string]Collation{
	"big5_chinese_ci":            1,
	"latin2_czech_cs":            2,
	"dec8_swedish_ci":            3,
	"cp850_general_ci":           4,
	"latin1_german1_ci":          5,
	"hp8_english_ci":             6,
	"koi8r_general_ci":           7,
	"latin1_swedish_ci":          8,
	"latin2_general_ci":          9,
	"swe7_swedish_ci":            10,
	"ascii_general_ci":           11,
	"ujis_japanese_ci":           12,
	"sjis_japanese_ci":           13,
	"cp1251_bulgarian_ci":        14,
	"latin1_danish_ci":           15,
	"hebrew_general_ci":          16,
	"tis620_thai_ci":             18,
	"euckr_korean_ci":            19,
	"latin7_estonian_cs":         20,
	"latin2_hungarian_ci":        21,
	"koi8u_general_ci":           22,
	"cp1251_ukrainian_ci":        23,
	"gb2312_chinese_ci":          24,
	"greek_general_ci":           25,
	"cp1250_general_ci":          26,
	"latin2_croatian_ci":         27,
	"gbk_chinese_ci":             28,
	"cp1257_lithuanian_ci":       29,
	"latin5_turkish_ci":          30,
	"latin1_german2_ci":          31,
	"armscii8_general_ci":        32,
	"utf8_general_ci":            33,
	"cp1250_czech_cs":            34,
	"ucs2_general_ci":            35,
	"cp866_general_ci":           36,
	"keybcs2_general_ci":         37,
	"macce_general_ci":           38,
	"macroman_general_ci":        39,
	"cp852_general_ci":           40,
	"latin7_general_ci":          41,
	"latin7_general_cs":          42,
	"macce_bin":                  43,
	"cp1250_croatian_ci":         44,
	"utf8mb4_general_ci":         45,
	"utf8mb4_bin":                46,
	"latin1_bin":                 47,
	"latin1_general_ci":          48,
	"latin1_general_cs":          49,
	"cp1251_bin":                 50,
	"cp1251_general_ci":          51,
	"cp1251_general_cs":          52,
	"macroman_bin":               53,
	"utf16_general_ci":           54,
	"utf16_bin":                  55,
	"utf16le_general_ci":         56,
	"cp1256_general_ci":          57,
	"cp1257_bin":                 58,
	"cp1257_general_ci":          59,
	"utf32_general_ci":           60,
	"utf32_bin":                  61,
	"utf16le_bin":                62,
	"binary":                     63,
	"armscii8_bin":               64,
	"ascii_bin":                  65,
	"cp1250_bin":                 66,
	"cp1256_bin":                 67,
	"cp866_bin":                  68,
	"dec8_bin":                   69,
	"greek_bin":                  70,
	"hebrew_bin":                 71,
	"hp8_bin":                    72,
	"keybcs2_bin":                73,
	"koi8r_bin":                  74,
	"koi8u_bin":                  75,
	"utf8_tolower_ci":            76,
	"latin2_bin":                 77,
	"latin5_bin":                 78,
	"latin7_bin":                 79,
	"cp850_bin":                  80,
	"cp852_bin":                  81,
	"swe7_bin":                   82,
	"utf8_bin":                   83,
	"big5_bin":                   84,
	"euckr_bin":                  85,
	"gb2312_bin":                 86,
	"gbk_bin":                    87,
	"sjis_bin":                   88,
	"tis620_bin":                 89,
	"ucs2_bin":                   90,
	"ujis_bin":                   91,
	"geostd8_general_ci":         92,
	"geostd8_bin":                93,
	"latin1_spanish_ci":          94,
	"cp932_japanese_ci":          95,
	"cp932_bin":                  96,
	"eucjpms_japanese_ci":        97,
	"eucjpms_bin":                98,
	"cp1250_polish_ci":           99,
	"utf16_unicode_ci":           101,
	"utf16_icelandic_ci":         102,
	"utf16_latvian_ci":           103,
	"utf16_romanian_ci":          104,
	"utf16_slovenian_ci":         105,
	"utf16_polish_ci":            106,
	"utf16_estonian_ci":          107,
	"utf16_spanish_ci":           108,
	"utf16_swedish_ci":           109,
	"utf16_turkish_ci":           110,
	"utf16_czech_ci":             111,
	"utf16_danish_ci":            112,
	"utf16_lithuanian_ci":        113,
	"utf16_slovak_ci":            114,
	"utf16_spanish2_ci":          115,
	"utf16_roman_ci":             116,
	"utf16_persian_ci":           117,
	"utf16_esperanto_ci":         118,
	"utf16_hungarian_ci":         119,
	"utf16_sinhala_ci":           120,
	"utf16_german2_ci":           121,
	"utf16_croatian_ci":          122,
	"utf16_unicode_520_ci":       123,
	"utf16_vietnamese_ci":        124,
	"ucs2_unicode_ci":            128,
	"ucs2_icelandic_ci":          129,
	"ucs2_latvian_ci":            130,
	"ucs2_romanian_ci":           131,
	"ucs2_slovenian_ci":          132,
	"ucs2_polish_ci":             133,
	"ucs2_estonian_ci":           134,
	"ucs2_spanish_ci":            135,
	"ucs2_swedish_ci":            136,
	"ucs2_turkish_ci":            137,
	"ucs2_czech_ci":              138,
	"ucs2_danish_ci":             139,
	"ucs2_lithuanian_ci":         140,
	"ucs2_slovak_ci":             141,
	"ucs2_spanish2_ci":           142,
	"ucs2_roman_ci":              143,
	"ucs2_persian_ci":            144,
	"ucs2_esperanto_ci":          145,
	"ucs2_hungarian_ci":          146,
	"ucs2_sinhala_ci":            147,
	"ucs2_german2_ci":            148,
	"ucs2_croatian_ci":           149,
	"ucs2_unicode_520_ci":        150,
	"ucs2_vietnamese_ci":         151,
	"ucs2_general_mysql500_ci":   159,
	"utf32_unicode_ci":           160,
	"utf32_icelandic_ci":         161,
	"utf32_latvian_ci":           162,
	"utf32_romanian_ci":          163,
	"utf32_slovenian_ci":         164,
	"utf32_polish_ci":            165,
	"utf32_estonian_ci":          166,
	"utf32_spanish_ci":           167,
	"utf32_swedish_ci":           168,
	"utf32_turkish_ci":           169,
	"utf32_czech_ci":             170,
	"utf32_danish_ci":            171,
	"utf32_lithuanian_ci":        172,
	"utf32_slovak_ci":            173,
	"utf32_spanish2_ci":          174,
	"utf32_roman_ci":             175,
	"utf32_persian_ci":           176,
	"utf32_esperanto_ci":         177,
	"utf32_hungarian_ci":         178,
	"utf32_sinhala_ci":           179,
	"utf32_german2_ci":           180,
	"utf32_croatian_ci":          181,
	"utf32_unicode_520_ci":       182,
	"utf32_vietnamese_ci":        183,
	"utf8_unicode_ci":            192,
	"utf8_icelandic_ci":          193,
	"utf8_latvian_ci":            194,
	"utf8_romanian_ci":           195,
	"utf8_slovenian_ci":          196,
	"utf8_polish_ci":             197,
	"utf8_estonian_ci":           198,
	"utf8_spanish_ci":            199,
	"utf8_swedish_ci":            200,
	"utf8_turkish_ci":            201,
	"utf8_czech_ci":              202,
	"utf8_danish_ci":             203,
	"utf8_lithuanian_ci":         204,
	"utf8_slovak_ci":             205,
	"utf8_spanish2_ci":           206,
	"utf8_roman_ci":              207,
	"utf8_persian_ci":            208,
	"utf8_esperanto_ci":          209,
	"utf8_hungarian_ci":          210,
	"utf8_sinhala_ci":            211,
	"utf8_german2_ci":            212,
	"utf8_croatian_ci":           213,
	"utf8_unicode_520_ci":        214,
	"utf8_vietnamese_ci":         215,
	"utf8_general_mysql500_ci":   223,
	"utf8mb4_unicode_ci":         224,
	"utf8mb4_icelandic_ci":       225,
	"utf8mb4_latvian_ci":         226,
	"utf8mb4_romanian_ci":        227,
	"utf8mb4_slovenian_ci":       228,
	"utf8mb4_polish_ci":          229,
	"utf8mb4_estonian_ci":        230,
	"utf8mb4_spanish_ci":         231,
	"utf8mb4_swedish_ci":         232,
	"utf8mb4_turkish_ci":         233,
	"utf8mb4_czech_ci":           234,
	"utf8mb4_danish_ci":          235,
	"utf8mb4_lithuanian_ci":      236,
	"utf8mb4_slovak_ci":          237,
	"utf8mb4_spanish2_ci":        238,
	"utf8mb4_roman_ci":           239,
	"utf8mb4_persian_ci":         240,
	"utf8mb4_esperanto_ci":       241,
	"utf8mb4_hungarian_ci":       242,
	"utf8mb4_sinhala_ci":         243,
	"utf8mb4_german2_ci":         244,
	"utf8mb4_croatian_ci":        245,
	"utf8mb4_unicode_520_ci":     246,
	"utf8mb4_vietnamese_ci":      247,
	"gb18030_chinese_ci":         248,
	"gb18030_bin":                249,
	"gb18030_unicode_520_ci":     250,
	"utf8mb4_0900_ai_ci":         255,
	"utf8mb4_de_pb_0900_ai_ci":   256,
	"utf8mb4_is_0900_ai_ci":      257,
	"utf8mb4_lv_0900_ai_ci":      258,
	"utf8mb4_ro_0900_ai_ci":      259,
	"utf8mb4_sl_0900_ai_ci":      260,
	"utf8mb4_pl_0900_ai_ci":      261,
	"utf8mb4_et_0900_ai_ci":      262,
	"utf8mb4_es_0900_ai_ci":      263,
	"utf8mb4_sv_0900_ai_ci":      264,
	"utf8mb4_tr_0900_ai_ci":      265,
	"utf8mb4_cs_0900_ai_ci":      266,
	"utf8mb4_da_0900_ai_ci":      267,
	"utf8mb4_lt_0900_ai_ci":      268,
	"utf8mb4_sk_0900_ai_ci":      269,
	"utf8mb4_es_trad_0900_ai_ci": 270,
	"utf8mb4_la_0900_ai_ci":      271,
	"utf8mb4_eo_0900_ai_ci":      273,
	"utf8mb4_hu_0900_ai_ci":      274,
	"utf8mb4_hr_0900_ai_ci":      275,
	"utf8mb4_vi_0900_ai_ci":      277,
	"utf8mb4_0900_as_cs":         278,
	"utf8mb4_de_pb_0900_as_cs":   279,
	"utf8mb4_is_0900_as_cs":      280,
	"utf8mb4_lv_0900_as_cs":      281,
	"utf8mb4_ro_0900_as_cs":      282,
	"utf8mb4_sl_0900_as_cs":      283,
	"utf8mb4_pl_0900_as_cs":      284,
	"utf8mb4_et_0900_as_cs":      285,
	"utf8mb4_es_0900_as_cs":      286,
	"utf8mb4_sv_0900_as_cs":      287,
	"utf8mb4_tr_0900_as_cs":      288,
	"utf8mb4_cs_0900_as_cs":      289,
	"utf8mb4_da_0900_as_cs":      290,
	"utf8mb4_lt_0900_as_cs":      291,
	"utf8mb4_sk_0900_as_cs":      292,
	"utf8mb4_es_trad_0900_as_cs": 293,
	"utf8mb4_la_0900_as_cs":      294,
	"utf8mb4_eo_0900_as_cs":      296,
	"utf8mb4_hu_0900_as_cs":      297,
	"utf8mb4_hr_0900_as_cs":      298,
	"utf8mb4_vi_0900_as_cs":      300,
	"utf8mb4_ja_0900_as_cs":      303,
	"utf8mb4_ja_0900_as_cs_ks":   304,
	"utf8mb4_0900_as_ci":         305,
	"utf8mb4_ru_0900_ai_ci":      306,
	"utf8mb4_ru_0900_as_cs":      307,
}

var collationToName = map[Collation]string{
	1:   "big5_chinese_ci",
	2:   "latin2_czech_cs",
	3:   "dec8_swedish_ci",
	4:   "cp850_general_ci",
	5:   "latin1_german1_ci",
	6:   "hp8_english_ci",
	7:   "koi8r_general_ci",
	8:   "latin1_swedish_ci",
	9:   "latin2_general_ci",
	10:  "swe7_swedish_ci",
	11:  "ascii_general_ci",
	12:  "ujis_japanese_ci",
	13:  "sjis_japanese_ci",
	14:  "cp1251_bulgarian_ci",
	15:  "latin1_danish_ci",
	16:  "hebrew_general_ci",
	18:  "tis620_thai_ci",
	19:  "euckr_korean_ci",
	20:  "latin7_estonian_cs",
	21:  "latin2_hungarian_ci",
	22:  "koi8u_general_ci",
	23:  "cp1251_ukrainian_ci",
	24:  "gb2312_chinese_ci",
	25:  "greek_general_ci",
	26:  "cp1250_general_ci",
	27:  "latin2_croatian_ci",
	28:  "gbk_chinese_ci",
	29:  "cp1257_lithuanian_ci",
	30:  "latin5_turkish_ci",
	31:  "latin1_german2_ci",
	32:  "armscii8_general_ci",
	33:  "utf8_general_ci",
	34:  "cp1250_czech_cs",
	35:  "ucs2_general_ci",
	36:  "cp866_general_ci",
	37:  "keybcs2_general_ci",
	38:  "macce_general_ci",
	39:  "macroman_general_ci",
	40:  "cp852_general_ci",
	41:  "latin7_general_ci",
	42:  "latin7_general_cs",
	43:  "macce_bin",
	44:  "cp1250_croatian_ci",
	45:  "utf8mb4_general_ci",
	46:  "utf8mb4_bin",
	47:  "latin1_bin",
	48:  "latin1_general_ci",
	49:  "latin1_general_cs",
	50:  "cp1251_bin",
	51:  "cp1251_general_ci",
	52:  "cp1251_general_cs",
	53:  "macroman_bin",
	54:  "utf16_general_ci",
	55:  "utf16_bin",
	56:  "utf16le_general_ci",
	57:  "cp1256_general_ci",
	58:  "cp1257_bin",
	59:  "cp1257_general_ci",
	60:  "utf32_general_ci",
	61:  "utf32_bin",
	62:  "utf16le_bin",
	63:  "binary",
	64:  "armscii8_bin",
	65:  "ascii_bin",
	66:  "cp1250_bin",
	67:  "cp1256_bin",
	68:  "cp866_bin",
	69:  "dec8_bin",
	70:  "greek_bin",
	71:  "hebrew_bin",
	72:  "hp8_bin",
	73:  "keybcs2_bin",
	74:  "koi8r_bin",
	75:  "koi8u_bin",
	76:  "utf8_tolower_ci",
	77:  "latin2_bin",
	78:  "latin5_bin",
	79:  "latin7_bin",
	80:  "cp850_bin",
	81:  "cp852_bin",
	82:  "swe7_bin",
	83:  "utf8_bin",
	84:  "big5_bin",
	85:  "euckr_bin",
	86:  "gb2312_bin",
	87:  "gbk_bin",
	88:  "sjis_bin",
	89:  "tis620_bin",
	90:  "ucs2_bin",
	91:  "ujis_bin",
	92:  "geostd8_general_ci",
	93:  "geostd8_bin",
	94:  "latin1_spanish_ci",
	95:  "cp932_japanese_ci",
	96:  "cp932_bin",
	97:  "eucjpms_japanese_ci",
	98:  "eucjpms_bin",
	99:  "cp1250_polish_ci",
	101: "utf16_unicode_ci",
	102: "utf16_icelandic_ci",
	103: "utf16_latvian_ci",
	104: "utf16_romanian_ci",
	105: "utf16_slovenian_ci",
	106: "utf16_polish_ci",
	107: "utf16_estonian_ci",
	108: "utf16_spanish_ci",
	109: "utf16_swedish_ci",
	110: "utf16_turkish_ci",
	111: "utf16_czech_ci",
	112: "utf16_danish_ci",
	113: "utf16_lithuanian_ci",
	114: "utf16_slovak_ci",
	115: "utf16_spanish2_ci",
	116: "utf16_roman_ci",
	117: "utf16_persian_ci",
	118: "utf16_esperanto_ci",
	119: "utf16_hungarian_ci",
	120: "utf16_sinhala_ci",
	121: "utf16_german2_ci",
	122: "utf16_croatian_ci",
	123: "utf16_unicode_520_ci",
	124: "utf16_vietnamese_ci",
	128: "ucs2_unicode_ci",
	129: "ucs2_icelandic_ci",
	130: "ucs2_latvian_ci",
	131: "ucs2_romanian_ci",
	132: "ucs2_slovenian_ci",
	133: "ucs2_polish_ci",
	134: "ucs2_estonian_ci",
	135: "ucs2_spanish_ci",
	136: "ucs2_swedish_ci",
	137: "ucs2_turkish_ci",
	138: "ucs2_czech_ci",
	139: "ucs2_danish_ci",
	140: "ucs2_lithuanian_ci",
	141: "ucs2_slovak_ci",
	142: "ucs2_spanish2_ci",
	143: "ucs2_roman_ci",
	144: "ucs2_persian_ci",
	145: "ucs2_esperanto_ci",
	146: "ucs2_hungarian_ci",
	147: "ucs2_sinhala_ci",
	148: "ucs2_german2_ci",
	149: "ucs2_croatian_ci",
	150: "ucs2_unicode_520_ci",
	151: "ucs2_vietnamese_ci",
	159: "ucs2_general_mysql500_ci",
	160: "utf32_unicode_ci",
	161: "utf32_icelandic_ci",
	162: "utf32_latvian_ci",
	163: "utf32_romanian_ci",
	164: "utf32_slovenian_ci",
	165: "utf32_polish_ci",
	166: "utf32_estonian_ci",
	167: "utf32_spanish_ci",
	168: "utf32_swedish_ci",
	169: "utf32_turkish_ci",
	170: "utf32_czech_ci",
	171: "utf32_danish_ci",
	172: "utf32_lithuanian_ci",
	173: "utf32_slovak_ci",
	174: "utf32_spanish2_ci",
	175: "utf32_roman_ci",
	176: "utf32_persian_ci",
	177: "utf32_esperanto_ci",
	178: "utf32_hungarian_ci",
	179: "utf32_sinhala_ci",
	180: "utf32_german2_ci",
	181: "utf32_croatian_ci",
	182: "utf32_unicode_520_ci",
	183: "utf32_vietnamese_ci",
	192: "utf8_unicode_ci",
	193: "utf8_icelandic_ci",
	194: "utf8_latvian_ci",
	195: "utf8_romanian_ci",
	196: "utf8_slovenian_ci",
	197: "utf8_polish_ci",
	198: "utf8_estonian_ci",
	199: "utf8_spanish_ci",
	200: "utf8_swedish_ci",
	201: "utf8_turkish_ci",
	202: "utf8_czech_ci",
	203: "utf8_danish_ci",
	204: "utf8_lithuanian_ci",
	205: "utf8_slovak_ci",
	206: "utf8_spanish2_ci",
	207: "utf8_roman_ci",
	208: "utf8_persian_ci",
	209: "utf8_esperanto_ci",
	210: "utf8_hungarian_ci",
	211: "utf8_sinhala_ci",
	212: "utf8_german2_ci",
	213: "utf8_croatian_ci",
	214: "utf8_unicode_520_ci",
	215: "utf8_vietnamese_ci",
	223: "utf8_general_mysql500_ci",
	224: "utf8mb4_unicode_ci",
	225: "utf8mb4_icelandic_ci",
	226: "utf8mb4_latvian_ci",
	227: "utf8mb4_romanian_ci",
	228: "utf8mb4_slovenian_ci",
	229: "utf8mb4_polish_ci",
	230: "utf8mb4_estonian_ci",
	231: "utf8mb4_spanish_ci",
	232: "utf8mb4_swedish_ci",
	233: "utf8mb4_turkish_ci",
	234: "utf8mb4_czech_ci",
	235: "utf8mb4_danish_ci",
	236: "utf8mb4_lithuanian_ci",
	237: "utf8mb4_slovak_ci",
	238: "utf8mb4_spanish2_ci",
	239: "utf8mb4_roman_ci",
	240: "utf8mb4_persian_ci",
	241: "utf8mb4_esperanto_ci",
	242: "utf8mb4_hungarian_ci",
	243: "utf8mb4_sinhala_ci",
	244: "utf8mb4_german2_ci",
	245: "utf8mb4_croatian_ci",
	246: "utf8mb4_unicode_520_ci",
	247: "utf8mb4_vietnamese_ci",
	248: "gb18030_chinese_ci",
	249: "gb18030_bin",
	250: "gb18030_unicode_520_ci",
	255: "utf8mb4_0900_ai_ci",
	256: "utf8mb4_de_pb_0900_ai_ci",
	257: "utf8mb4_is_0900_ai_ci",
	258: "utf8mb4_lv_0900_ai_ci",
	259: "utf8mb4_ro_0900_ai_ci",
	260: "utf8mb4_sl_0900_ai_ci",
	261: "utf8mb4_pl_0900_ai_ci",
	262: "utf8mb4_et_0900_ai_ci",
	263: "utf8mb4_es_0900_ai_ci",
	264: "utf8mb4_sv_0900_ai_ci",
	265: "utf8mb4_tr_0900_ai_ci",
	266: "utf8mb4_cs_0900_ai_ci",
	267: "utf8mb4_da_0900_ai_ci",
	268: "utf8mb4_lt_0900_ai_ci",
	269: "utf8mb4_sk_0900_ai_ci",
	270: "utf8mb4_es_trad_0900_ai_ci",
	271: "utf8mb4_la_0900_ai_ci",
	273: "utf8mb4_eo_0900_ai_ci",
	274: "utf8mb4_hu_0900_ai_ci",
	275: "utf8mb4_hr_0900_ai_ci",
	277: "utf8mb4_vi_0900_ai_ci",
	278: "utf8mb4_0900_as_cs",
	279: "utf8mb4_de_pb_0900_as_cs",
	280: "utf8mb4_is_0900_as_cs",
	281: "utf8mb4_lv_0900_as_cs",
	282: "utf8mb4_ro_0900_as_cs",
	283: "utf8mb4_sl_0900_as_cs",
	284: "utf8mb4_pl_0900_as_cs",
	285: "utf8mb4_et_0900_as_cs",
	286: "utf8mb4_es_0900_as_cs",
	287: "utf8mb4_sv_0900_as_cs",
	288: "utf8mb4_tr_0900_as_cs",
	289: "utf8mb4_cs_0900_as_cs",
	290: "utf8mb4_da_0900_as_cs",
	291: "utf8mb4_lt_0900_as_cs",
	292: "utf8mb4_sk_0900_as_cs",
	293: "utf8mb4_es_trad_0900_as_cs",
	294: "utf8mb4_la_0900_as_cs",
	296: "utf8mb4_eo_0900_as_cs",
	297: "utf8mb4_hu_0900_as_cs",
	298: "utf8mb4_hr_0900_as_cs",
	300: "utf8mb4_vi_0900_as_cs",
	303: "utf8mb4_ja_0900_as_cs",
	304: "utf8mb4_ja_0900_as_cs_ks",
	305: "utf8mb4_0900_as_ci",
	306: "utf8mb4_ru_0900_ai_ci",
	307: "utf8mb4_ru_0900_as_cs",
}
