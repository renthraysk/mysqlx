package collation

import "strconv"

// Collation is a mysql collation
type Collation uint64

func (c Collation) String() string {
	if str, ok := collationMap[c]; ok {
		return str
	}
	return "Collation(" + strconv.FormatUint(uint64(c), 10) + ")"
}

// IsBinary returns whether this Collation represents the binary collation
func (c Collation) IsBinary() bool {
	return c == Binary
}

const collationName = "big5_chinese_cilatin2_czech_csdec8_swedish_cicp850_general_cilatin1_german1_cihp8_english_cikoi8r_general_cilatin1_swedish_cilatin2_general_ciswe7_swedish_ciascii_general_ciujis_japanese_cisjis_japanese_cicp1251_bulgarian_cilatin1_danish_cihebrew_general_citis620_thai_cieuckr_korean_cilatin7_estonian_cslatin2_hungarian_cikoi8u_general_cicp1251_ukrainian_cigb2312_chinese_cigreek_general_cicp1250_general_cilatin2_croatian_cigbk_chinese_cicp1257_lithuanian_cilatin5_turkish_cilatin1_german2_ciarmscii8_general_ciutf8_general_cicp1250_czech_csucs2_general_cicp866_general_cikeybcs2_general_cimacce_general_cimacroman_general_cicp852_general_cilatin7_general_cilatin7_general_csmacce_bincp1250_croatian_ciutf8mb4_general_ciutf8mb4_binlatin1_binlatin1_general_cilatin1_general_cscp1251_bincp1251_general_cicp1251_general_csmacroman_binutf16_general_ciutf16_binutf16le_general_cicp1256_general_cicp1257_bincp1257_general_ciutf32_general_ciutf32_binutf16le_binbinaryarmscii8_binascii_bincp1250_bincp1256_bincp866_bindec8_bingreek_binhebrew_binhp8_binkeybcs2_binkoi8r_binkoi8u_binutf8_tolower_cilatin2_binlatin5_binlatin7_bincp850_bincp852_binswe7_binutf8_binbig5_bineuckr_bingb2312_bingbk_binsjis_bintis620_binucs2_binujis_bingeostd8_general_cigeostd8_binlatin1_spanish_cicp932_japanese_cicp932_bineucjpms_japanese_cieucjpms_bincp1250_polish_ciutf16_unicode_ciutf16_icelandic_ciutf16_latvian_ciutf16_romanian_ciutf16_slovenian_ciutf16_polish_ciutf16_estonian_ciutf16_spanish_ciutf16_swedish_ciutf16_turkish_ciutf16_czech_ciutf16_danish_ciutf16_lithuanian_ciutf16_slovak_ciutf16_spanish2_ciutf16_roman_ciutf16_persian_ciutf16_esperanto_ciutf16_hungarian_ciutf16_sinhala_ciutf16_german2_ciutf16_croatian_ciutf16_unicode_520_ciutf16_vietnamese_ciucs2_unicode_ciucs2_icelandic_ciucs2_latvian_ciucs2_romanian_ciucs2_slovenian_ciucs2_polish_ciucs2_estonian_ciucs2_spanish_ciucs2_swedish_ciucs2_turkish_ciucs2_czech_ciucs2_danish_ciucs2_lithuanian_ciucs2_slovak_ciucs2_spanish2_ciucs2_roman_ciucs2_persian_ciucs2_esperanto_ciucs2_hungarian_ciucs2_sinhala_ciucs2_german2_ciucs2_croatian_ciucs2_unicode_520_ciucs2_vietnamese_ciucs2_general_mysql500_ciutf32_unicode_ciutf32_icelandic_ciutf32_latvian_ciutf32_romanian_ciutf32_slovenian_ciutf32_polish_ciutf32_estonian_ciutf32_spanish_ciutf32_swedish_ciutf32_turkish_ciutf32_czech_ciutf32_danish_ciutf32_lithuanian_ciutf32_slovak_ciutf32_spanish2_ciutf32_roman_ciutf32_persian_ciutf32_esperanto_ciutf32_hungarian_ciutf32_sinhala_ciutf32_german2_ciutf32_croatian_ciutf32_unicode_520_ciutf32_vietnamese_ciutf8_unicode_ciutf8_icelandic_ciutf8_latvian_ciutf8_romanian_ciutf8_slovenian_ciutf8_polish_ciutf8_estonian_ciutf8_spanish_ciutf8_swedish_ciutf8_turkish_ciutf8_czech_ciutf8_danish_ciutf8_lithuanian_ciutf8_slovak_ciutf8_spanish2_ciutf8_roman_ciutf8_persian_ciutf8_esperanto_ciutf8_hungarian_ciutf8_sinhala_ciutf8_german2_ciutf8_croatian_ciutf8_unicode_520_ciutf8_vietnamese_ciutf8_general_mysql500_ciutf8mb4_unicode_ciutf8mb4_icelandic_ciutf8mb4_latvian_ciutf8mb4_romanian_ciutf8mb4_slovenian_ciutf8mb4_polish_ciutf8mb4_estonian_ciutf8mb4_spanish_ciutf8mb4_swedish_ciutf8mb4_turkish_ciutf8mb4_czech_ciutf8mb4_danish_ciutf8mb4_lithuanian_ciutf8mb4_slovak_ciutf8mb4_spanish2_ciutf8mb4_roman_ciutf8mb4_persian_ciutf8mb4_esperanto_ciutf8mb4_hungarian_ciutf8mb4_sinhala_ciutf8mb4_german2_ciutf8mb4_croatian_ciutf8mb4_unicode_520_ciutf8mb4_vietnamese_cigb18030_chinese_cigb18030_bingb18030_unicode_520_ciutf8mb4_0900_ai_ciutf8mb4_de_pb_0900_ai_ciutf8mb4_is_0900_ai_ciutf8mb4_lv_0900_ai_ciutf8mb4_ro_0900_ai_ciutf8mb4_sl_0900_ai_ciutf8mb4_pl_0900_ai_ciutf8mb4_et_0900_ai_ciutf8mb4_es_0900_ai_ciutf8mb4_sv_0900_ai_ciutf8mb4_tr_0900_ai_ciutf8mb4_cs_0900_ai_ciutf8mb4_da_0900_ai_ciutf8mb4_lt_0900_ai_ciutf8mb4_sk_0900_ai_ciutf8mb4_es_trad_0900_ai_ciutf8mb4_la_0900_ai_ciutf8mb4_eo_0900_ai_ciutf8mb4_hu_0900_ai_ciutf8mb4_hr_0900_ai_ciutf8mb4_vi_0900_ai_ciutf8mb4_0900_as_csutf8mb4_de_pb_0900_as_csutf8mb4_is_0900_as_csutf8mb4_lv_0900_as_csutf8mb4_ro_0900_as_csutf8mb4_sl_0900_as_csutf8mb4_pl_0900_as_csutf8mb4_et_0900_as_csutf8mb4_es_0900_as_csutf8mb4_sv_0900_as_csutf8mb4_tr_0900_as_csutf8mb4_cs_0900_as_csutf8mb4_da_0900_as_csutf8mb4_lt_0900_as_csutf8mb4_sk_0900_as_csutf8mb4_es_trad_0900_as_csutf8mb4_la_0900_as_csutf8mb4_eo_0900_as_csutf8mb4_hu_0900_as_csutf8mb4_hr_0900_as_csutf8mb4_vi_0900_as_csutf8mb4_ja_0900_as_csutf8mb4_ja_0900_as_cs_ksutf8mb4_0900_as_ciutf8mb4_ru_0900_ai_ciutf8mb4_ru_0900_as_cs"

var collationMap = map[Collation]string{
	1:   collationName[0:15],
	2:   collationName[15:30],
	3:   collationName[30:45],
	4:   collationName[45:61],
	5:   collationName[61:78],
	6:   collationName[78:92],
	7:   collationName[92:108],
	8:   collationName[108:125],
	9:   collationName[125:142],
	10:  collationName[142:157],
	11:  collationName[157:173],
	12:  collationName[173:189],
	13:  collationName[189:205],
	14:  collationName[205:224],
	15:  collationName[224:240],
	16:  collationName[240:257],
	18:  collationName[257:271],
	19:  collationName[271:286],
	20:  collationName[286:304],
	21:  collationName[304:323],
	22:  collationName[323:339],
	23:  collationName[339:358],
	24:  collationName[358:375],
	25:  collationName[375:391],
	26:  collationName[391:408],
	27:  collationName[408:426],
	28:  collationName[426:440],
	29:  collationName[440:460],
	30:  collationName[460:477],
	31:  collationName[477:494],
	32:  collationName[494:513],
	33:  collationName[513:528],
	34:  collationName[528:543],
	35:  collationName[543:558],
	36:  collationName[558:574],
	37:  collationName[574:592],
	38:  collationName[592:608],
	39:  collationName[608:627],
	40:  collationName[627:643],
	41:  collationName[643:660],
	42:  collationName[660:677],
	43:  collationName[677:686],
	44:  collationName[686:704],
	45:  collationName[704:722],
	46:  collationName[722:733],
	47:  collationName[733:743],
	48:  collationName[743:760],
	49:  collationName[760:777],
	50:  collationName[777:787],
	51:  collationName[787:804],
	52:  collationName[804:821],
	53:  collationName[821:833],
	54:  collationName[833:849],
	55:  collationName[849:858],
	56:  collationName[858:876],
	57:  collationName[876:893],
	58:  collationName[893:903],
	59:  collationName[903:920],
	60:  collationName[920:936],
	61:  collationName[936:945],
	62:  collationName[945:956],
	63:  collationName[956:962],
	64:  collationName[962:974],
	65:  collationName[974:983],
	66:  collationName[983:993],
	67:  collationName[993:1003],
	68:  collationName[1003:1012],
	69:  collationName[1012:1020],
	70:  collationName[1020:1029],
	71:  collationName[1029:1039],
	72:  collationName[1039:1046],
	73:  collationName[1046:1057],
	74:  collationName[1057:1066],
	75:  collationName[1066:1075],
	76:  collationName[1075:1090],
	77:  collationName[1090:1100],
	78:  collationName[1100:1110],
	79:  collationName[1110:1120],
	80:  collationName[1120:1129],
	81:  collationName[1129:1138],
	82:  collationName[1138:1146],
	83:  collationName[1146:1154],
	84:  collationName[1154:1162],
	85:  collationName[1162:1171],
	86:  collationName[1171:1181],
	87:  collationName[1181:1188],
	88:  collationName[1188:1196],
	89:  collationName[1196:1206],
	90:  collationName[1206:1214],
	91:  collationName[1214:1222],
	92:  collationName[1222:1240],
	93:  collationName[1240:1251],
	94:  collationName[1251:1268],
	95:  collationName[1268:1285],
	96:  collationName[1285:1294],
	97:  collationName[1294:1313],
	98:  collationName[1313:1324],
	99:  collationName[1324:1340],
	101: collationName[1340:1356],
	102: collationName[1356:1374],
	103: collationName[1374:1390],
	104: collationName[1390:1407],
	105: collationName[1407:1425],
	106: collationName[1425:1440],
	107: collationName[1440:1457],
	108: collationName[1457:1473],
	109: collationName[1473:1489],
	110: collationName[1489:1505],
	111: collationName[1505:1519],
	112: collationName[1519:1534],
	113: collationName[1534:1553],
	114: collationName[1553:1568],
	115: collationName[1568:1585],
	116: collationName[1585:1599],
	117: collationName[1599:1615],
	118: collationName[1615:1633],
	119: collationName[1633:1651],
	120: collationName[1651:1667],
	121: collationName[1667:1683],
	122: collationName[1683:1700],
	123: collationName[1700:1720],
	124: collationName[1720:1739],
	128: collationName[1739:1754],
	129: collationName[1754:1771],
	130: collationName[1771:1786],
	131: collationName[1786:1802],
	132: collationName[1802:1819],
	133: collationName[1819:1833],
	134: collationName[1833:1849],
	135: collationName[1849:1864],
	136: collationName[1864:1879],
	137: collationName[1879:1894],
	138: collationName[1894:1907],
	139: collationName[1907:1921],
	140: collationName[1921:1939],
	141: collationName[1939:1953],
	142: collationName[1953:1969],
	143: collationName[1969:1982],
	144: collationName[1982:1997],
	145: collationName[1997:2014],
	146: collationName[2014:2031],
	147: collationName[2031:2046],
	148: collationName[2046:2061],
	149: collationName[2061:2077],
	150: collationName[2077:2096],
	151: collationName[2096:2114],
	159: collationName[2114:2138],
	160: collationName[2138:2154],
	161: collationName[2154:2172],
	162: collationName[2172:2188],
	163: collationName[2188:2205],
	164: collationName[2205:2223],
	165: collationName[2223:2238],
	166: collationName[2238:2255],
	167: collationName[2255:2271],
	168: collationName[2271:2287],
	169: collationName[2287:2303],
	170: collationName[2303:2317],
	171: collationName[2317:2332],
	172: collationName[2332:2351],
	173: collationName[2351:2366],
	174: collationName[2366:2383],
	175: collationName[2383:2397],
	176: collationName[2397:2413],
	177: collationName[2413:2431],
	178: collationName[2431:2449],
	179: collationName[2449:2465],
	180: collationName[2465:2481],
	181: collationName[2481:2498],
	182: collationName[2498:2518],
	183: collationName[2518:2537],
	192: collationName[2537:2552],
	193: collationName[2552:2569],
	194: collationName[2569:2584],
	195: collationName[2584:2600],
	196: collationName[2600:2617],
	197: collationName[2617:2631],
	198: collationName[2631:2647],
	199: collationName[2647:2662],
	200: collationName[2662:2677],
	201: collationName[2677:2692],
	202: collationName[2692:2705],
	203: collationName[2705:2719],
	204: collationName[2719:2737],
	205: collationName[2737:2751],
	206: collationName[2751:2767],
	207: collationName[2767:2780],
	208: collationName[2780:2795],
	209: collationName[2795:2812],
	210: collationName[2812:2829],
	211: collationName[2829:2844],
	212: collationName[2844:2859],
	213: collationName[2859:2875],
	214: collationName[2875:2894],
	215: collationName[2894:2912],
	223: collationName[2912:2936],
	224: collationName[2936:2954],
	225: collationName[2954:2974],
	226: collationName[2974:2992],
	227: collationName[2992:3011],
	228: collationName[3011:3031],
	229: collationName[3031:3048],
	230: collationName[3048:3067],
	231: collationName[3067:3085],
	232: collationName[3085:3103],
	233: collationName[3103:3121],
	234: collationName[3121:3137],
	235: collationName[3137:3154],
	236: collationName[3154:3175],
	237: collationName[3175:3192],
	238: collationName[3192:3211],
	239: collationName[3211:3227],
	240: collationName[3227:3245],
	241: collationName[3245:3265],
	242: collationName[3265:3285],
	243: collationName[3285:3303],
	244: collationName[3303:3321],
	245: collationName[3321:3340],
	246: collationName[3340:3362],
	247: collationName[3362:3383],
	248: collationName[3383:3401],
	249: collationName[3401:3412],
	250: collationName[3412:3434],
	255: collationName[3434:3452],
	256: collationName[3452:3476],
	257: collationName[3476:3497],
	258: collationName[3497:3518],
	259: collationName[3518:3539],
	260: collationName[3539:3560],
	261: collationName[3560:3581],
	262: collationName[3581:3602],
	263: collationName[3602:3623],
	264: collationName[3623:3644],
	265: collationName[3644:3665],
	266: collationName[3665:3686],
	267: collationName[3686:3707],
	268: collationName[3707:3728],
	269: collationName[3728:3749],
	270: collationName[3749:3775],
	271: collationName[3775:3796],
	273: collationName[3796:3817],
	274: collationName[3817:3838],
	275: collationName[3838:3859],
	277: collationName[3859:3880],
	278: collationName[3880:3898],
	279: collationName[3898:3922],
	280: collationName[3922:3943],
	281: collationName[3943:3964],
	282: collationName[3964:3985],
	283: collationName[3985:4006],
	284: collationName[4006:4027],
	285: collationName[4027:4048],
	286: collationName[4048:4069],
	287: collationName[4069:4090],
	288: collationName[4090:4111],
	289: collationName[4111:4132],
	290: collationName[4132:4153],
	291: collationName[4153:4174],
	292: collationName[4174:4195],
	293: collationName[4195:4221],
	294: collationName[4221:4242],
	296: collationName[4242:4263],
	297: collationName[4263:4284],
	298: collationName[4284:4305],
	300: collationName[4305:4326],
	303: collationName[4326:4347],
	304: collationName[4347:4371],
	305: collationName[4371:4389],
	306: collationName[4389:4410],
	307: collationName[4410:4431],
}

// Collations
const (
	Big5ChineseCi         Collation = 1
	Latin2CzechCs         Collation = 2
	Dec8SwedishCi         Collation = 3
	Cp850GeneralCi        Collation = 4
	Latin1German1Ci       Collation = 5
	Hp8EnglishCi          Collation = 6
	Koi8rGeneralCi        Collation = 7
	Latin1SwedishCi       Collation = 8
	Latin2GeneralCi       Collation = 9
	Swe7SwedishCi         Collation = 10
	ASCIIGeneralCi        Collation = 11
	UjisJapaneseCi        Collation = 12
	SjisJapaneseCi        Collation = 13
	Cp1251BulgarianCi     Collation = 14
	Latin1DanishCi        Collation = 15
	HebrewGeneralCi       Collation = 16
	Tis620ThaiCi          Collation = 18
	EuckrKoreanCi         Collation = 19
	Latin7EstonianCs      Collation = 20
	Latin2HungarianCi     Collation = 21
	Koi8uGeneralCi        Collation = 22
	Cp1251UkrainianCi     Collation = 23
	Gb2312ChineseCi       Collation = 24
	GreekGeneralCi        Collation = 25
	Cp1250GeneralCi       Collation = 26
	Latin2CroatianCi      Collation = 27
	GbkChineseCi          Collation = 28
	Cp1257LithuanianCi    Collation = 29
	Latin5TurkishCi       Collation = 30
	Latin1German2Ci       Collation = 31
	Armscii8GeneralCi     Collation = 32
	Utf8GeneralCi         Collation = 33
	Cp1250CzechCs         Collation = 34
	Ucs2GeneralCi         Collation = 35
	Cp866GeneralCi        Collation = 36
	Keybcs2GeneralCi      Collation = 37
	MacceGeneralCi        Collation = 38
	MacromanGeneralCi     Collation = 39
	Cp852GeneralCi        Collation = 40
	Latin7GeneralCi       Collation = 41
	Latin7GeneralCs       Collation = 42
	MacceBin              Collation = 43
	Cp1250CroatianCi      Collation = 44
	Utf8mb4GeneralCi      Collation = 45
	Utf8mb4Bin            Collation = 46
	Latin1Bin             Collation = 47
	Latin1GeneralCi       Collation = 48
	Latin1GeneralCs       Collation = 49
	Cp1251Bin             Collation = 50
	Cp1251GeneralCi       Collation = 51
	Cp1251GeneralCs       Collation = 52
	MacromanBin           Collation = 53
	Utf16GeneralCi        Collation = 54
	Utf16Bin              Collation = 55
	Utf16leGeneralCi      Collation = 56
	Cp1256GeneralCi       Collation = 57
	Cp1257Bin             Collation = 58
	Cp1257GeneralCi       Collation = 59
	Utf32GeneralCi        Collation = 60
	Utf32Bin              Collation = 61
	Utf16leBin            Collation = 62
	Binary                Collation = 63
	Armscii8Bin           Collation = 64
	ASCIIBin              Collation = 65
	Cp1250Bin             Collation = 66
	Cp1256Bin             Collation = 67
	Cp866Bin              Collation = 68
	Dec8Bin               Collation = 69
	GreekBin              Collation = 70
	HebrewBin             Collation = 71
	Hp8Bin                Collation = 72
	Keybcs2Bin            Collation = 73
	Koi8rBin              Collation = 74
	Koi8uBin              Collation = 75
	Utf8TolowerCi         Collation = 76
	Latin2Bin             Collation = 77
	Latin5Bin             Collation = 78
	Latin7Bin             Collation = 79
	Cp850Bin              Collation = 80
	Cp852Bin              Collation = 81
	Swe7Bin               Collation = 82
	Utf8Bin               Collation = 83
	Big5Bin               Collation = 84
	EuckrBin              Collation = 85
	Gb2312Bin             Collation = 86
	GbkBin                Collation = 87
	SjisBin               Collation = 88
	Tis620Bin             Collation = 89
	Ucs2Bin               Collation = 90
	UjisBin               Collation = 91
	Geostd8GeneralCi      Collation = 92
	Geostd8Bin            Collation = 93
	Latin1SpanishCi       Collation = 94
	Cp932JapaneseCi       Collation = 95
	Cp932Bin              Collation = 96
	EucjpmsJapaneseCi     Collation = 97
	EucjpmsBin            Collation = 98
	Cp1250PolishCi        Collation = 99
	Utf16UnicodeCi        Collation = 101
	Utf16IcelandicCi      Collation = 102
	Utf16LatvianCi        Collation = 103
	Utf16RomanianCi       Collation = 104
	Utf16SlovenianCi      Collation = 105
	Utf16PolishCi         Collation = 106
	Utf16EstonianCi       Collation = 107
	Utf16SpanishCi        Collation = 108
	Utf16SwedishCi        Collation = 109
	Utf16TurkishCi        Collation = 110
	Utf16CzechCi          Collation = 111
	Utf16DanishCi         Collation = 112
	Utf16LithuanianCi     Collation = 113
	Utf16SlovakCi         Collation = 114
	Utf16Spanish2Ci       Collation = 115
	Utf16RomanCi          Collation = 116
	Utf16PersianCi        Collation = 117
	Utf16EsperantoCi      Collation = 118
	Utf16HungarianCi      Collation = 119
	Utf16SinhalaCi        Collation = 120
	Utf16German2Ci        Collation = 121
	Utf16CroatianCi       Collation = 122
	Utf16Unicode520Ci     Collation = 123
	Utf16VietnameseCi     Collation = 124
	Ucs2UnicodeCi         Collation = 128
	Ucs2IcelandicCi       Collation = 129
	Ucs2LatvianCi         Collation = 130
	Ucs2RomanianCi        Collation = 131
	Ucs2SlovenianCi       Collation = 132
	Ucs2PolishCi          Collation = 133
	Ucs2EstonianCi        Collation = 134
	Ucs2SpanishCi         Collation = 135
	Ucs2SwedishCi         Collation = 136
	Ucs2TurkishCi         Collation = 137
	Ucs2CzechCi           Collation = 138
	Ucs2DanishCi          Collation = 139
	Ucs2LithuanianCi      Collation = 140
	Ucs2SlovakCi          Collation = 141
	Ucs2Spanish2Ci        Collation = 142
	Ucs2RomanCi           Collation = 143
	Ucs2PersianCi         Collation = 144
	Ucs2EsperantoCi       Collation = 145
	Ucs2HungarianCi       Collation = 146
	Ucs2SinhalaCi         Collation = 147
	Ucs2German2Ci         Collation = 148
	Ucs2CroatianCi        Collation = 149
	Ucs2Unicode520Ci      Collation = 150
	Ucs2VietnameseCi      Collation = 151
	Ucs2GeneralMysql500Ci Collation = 159
	Utf32UnicodeCi        Collation = 160
	Utf32IcelandicCi      Collation = 161
	Utf32LatvianCi        Collation = 162
	Utf32RomanianCi       Collation = 163
	Utf32SlovenianCi      Collation = 164
	Utf32PolishCi         Collation = 165
	Utf32EstonianCi       Collation = 166
	Utf32SpanishCi        Collation = 167
	Utf32SwedishCi        Collation = 168
	Utf32TurkishCi        Collation = 169
	Utf32CzechCi          Collation = 170
	Utf32DanishCi         Collation = 171
	Utf32LithuanianCi     Collation = 172
	Utf32SlovakCi         Collation = 173
	Utf32Spanish2Ci       Collation = 174
	Utf32RomanCi          Collation = 175
	Utf32PersianCi        Collation = 176
	Utf32EsperantoCi      Collation = 177
	Utf32HungarianCi      Collation = 178
	Utf32SinhalaCi        Collation = 179
	Utf32German2Ci        Collation = 180
	Utf32CroatianCi       Collation = 181
	Utf32Unicode520Ci     Collation = 182
	Utf32VietnameseCi     Collation = 183
	Utf8UnicodeCi         Collation = 192
	Utf8IcelandicCi       Collation = 193
	Utf8LatvianCi         Collation = 194
	Utf8RomanianCi        Collation = 195
	Utf8SlovenianCi       Collation = 196
	Utf8PolishCi          Collation = 197
	Utf8EstonianCi        Collation = 198
	Utf8SpanishCi         Collation = 199
	Utf8SwedishCi         Collation = 200
	Utf8TurkishCi         Collation = 201
	Utf8CzechCi           Collation = 202
	Utf8DanishCi          Collation = 203
	Utf8LithuanianCi      Collation = 204
	Utf8SlovakCi          Collation = 205
	Utf8Spanish2Ci        Collation = 206
	Utf8RomanCi           Collation = 207
	Utf8PersianCi         Collation = 208
	Utf8EsperantoCi       Collation = 209
	Utf8HungarianCi       Collation = 210
	Utf8SinhalaCi         Collation = 211
	Utf8German2Ci         Collation = 212
	Utf8CroatianCi        Collation = 213
	Utf8Unicode520Ci      Collation = 214
	Utf8VietnameseCi      Collation = 215
	Utf8GeneralMysql500Ci Collation = 223
	Utf8mb4UnicodeCi      Collation = 224
	Utf8mb4IcelandicCi    Collation = 225
	Utf8mb4LatvianCi      Collation = 226
	Utf8mb4RomanianCi     Collation = 227
	Utf8mb4SlovenianCi    Collation = 228
	Utf8mb4PolishCi       Collation = 229
	Utf8mb4EstonianCi     Collation = 230
	Utf8mb4SpanishCi      Collation = 231
	Utf8mb4SwedishCi      Collation = 232
	Utf8mb4TurkishCi      Collation = 233
	Utf8mb4CzechCi        Collation = 234
	Utf8mb4DanishCi       Collation = 235
	Utf8mb4LithuanianCi   Collation = 236
	Utf8mb4SlovakCi       Collation = 237
	Utf8mb4Spanish2Ci     Collation = 238
	Utf8mb4RomanCi        Collation = 239
	Utf8mb4PersianCi      Collation = 240
	Utf8mb4EsperantoCi    Collation = 241
	Utf8mb4HungarianCi    Collation = 242
	Utf8mb4SinhalaCi      Collation = 243
	Utf8mb4German2Ci      Collation = 244
	Utf8mb4CroatianCi     Collation = 245
	Utf8mb4Unicode520Ci   Collation = 246
	Utf8mb4VietnameseCi   Collation = 247
	Gb18030ChineseCi      Collation = 248
	Gb18030Bin            Collation = 249
	Gb18030Unicode520Ci   Collation = 250
	Utf8mb40900AiCi       Collation = 255
	Utf8mb4DePb0900AiCi   Collation = 256
	Utf8mb4Is0900AiCi     Collation = 257
	Utf8mb4Lv0900AiCi     Collation = 258
	Utf8mb4Ro0900AiCi     Collation = 259
	Utf8mb4Sl0900AiCi     Collation = 260
	Utf8mb4Pl0900AiCi     Collation = 261
	Utf8mb4Et0900AiCi     Collation = 262
	Utf8mb4Es0900AiCi     Collation = 263
	Utf8mb4Sv0900AiCi     Collation = 264
	Utf8mb4Tr0900AiCi     Collation = 265
	Utf8mb4Cs0900AiCi     Collation = 266
	Utf8mb4Da0900AiCi     Collation = 267
	Utf8mb4Lt0900AiCi     Collation = 268
	Utf8mb4Sk0900AiCi     Collation = 269
	Utf8mb4EsTrad0900AiCi Collation = 270
	Utf8mb4La0900AiCi     Collation = 271
	Utf8mb4Eo0900AiCi     Collation = 273
	Utf8mb4Hu0900AiCi     Collation = 274
	Utf8mb4Hr0900AiCi     Collation = 275
	Utf8mb4Vi0900AiCi     Collation = 277
	Utf8mb40900AsCs       Collation = 278
	Utf8mb4DePb0900AsCs   Collation = 279
	Utf8mb4Is0900AsCs     Collation = 280
	Utf8mb4Lv0900AsCs     Collation = 281
	Utf8mb4Ro0900AsCs     Collation = 282
	Utf8mb4Sl0900AsCs     Collation = 283
	Utf8mb4Pl0900AsCs     Collation = 284
	Utf8mb4Et0900AsCs     Collation = 285
	Utf8mb4Es0900AsCs     Collation = 286
	Utf8mb4Sv0900AsCs     Collation = 287
	Utf8mb4Tr0900AsCs     Collation = 288
	Utf8mb4Cs0900AsCs     Collation = 289
	Utf8mb4Da0900AsCs     Collation = 290
	Utf8mb4Lt0900AsCs     Collation = 291
	Utf8mb4Sk0900AsCs     Collation = 292
	Utf8mb4EsTrad0900AsCs Collation = 293
	Utf8mb4La0900AsCs     Collation = 294
	Utf8mb4Eo0900AsCs     Collation = 296
	Utf8mb4Hu0900AsCs     Collation = 297
	Utf8mb4Hr0900AsCs     Collation = 298
	Utf8mb4Vi0900AsCs     Collation = 300
	Utf8mb4Ja0900AsCs     Collation = 303
	Utf8mb4Ja0900AsCsKs   Collation = 304
	Utf8mb40900AsCi       Collation = 305
	Utf8mb4Ru0900AiCi     Collation = 306
	Utf8mb4Ru0900AsCs     Collation = 307
)
