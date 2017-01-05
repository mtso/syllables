package syllables

import (
	"regexp"
)

var (
	expressionMonosyllabicOne = regexp.MustCompile(
		"cia(l|$)|" +
		"tia|" +
		"cius|" +
		"cious|" +
		"[^aeiou]giu|" +
		"[aeiouy][^aeiouy]ion|" +
		"iou|" +
		"sia$|" +
		"eous$|" +
		"[oa]gue$|" +
		".[^aeiuoycgltdb]{2,}ed$|" +
		".ely$|" +
		"^jua|" +
		"uai|" +
		"eau|" +
		"^busi$|" +
		"(" +
			"[aeiouy]" +
			"(" +
				"b|" +
				"c|" +
				"ch|" +
				"dg|" +
				"f|" +
				"g|" +
				"gh|" +
				"gn|" +
				"k|" +
				"l|" +
				"lch|" +
				"ll|" +
				"lv|" +
				"m|" +
				"mm|" +
				"n|" +
				"nc|" +
				"ng|" +
				"nch|" +
				"nn|" +
				"p|" +
				"r|" +
				"rc|" +
				"rn|" +
				"rs|" +
				"rv|" +
				"s|" +
				"sc|" +
				"sk|" +
				"sl|" +
				"squ|" +
				"ss|" +
				"th|" +
				"v|" +
				"y|" +
				"z" +
			")" +
			"ed$" +
		")|" +
		"(" +
			"[aeiouy]" +
			"(" +
				"b|" +
				"ch|" +
				"d|" +
				"f|" +
				"gh|" +
				"gn|" +
				"k|" +
				"l|" +
				"lch|" +
				"ll|" +
				"lv|" +
				"m|" +
				"mm|" +
				"n|" +
				"nch|" +
				"nn|" +
				"p|" +
				"r|" +
				"rn|" +
				"rs|" +
				"rv|" +
				"s|" +
				"sc|" +
				"sk|" +
				"sl|" +
				"squ|" +
				"ss|" +
				"st|" +
				"t|" +
				"th|" +
				"v|" +
				"y" +
			")" +
			"es$" +
		")",
	)

	expressionMonosyllabicTwo = regexp.MustCompile(
		"[aeiouy]" +
		"(" +
			"b|" +
			"c|" +
			"ch|" +
			"d|" +
			"dg|" +
			"f|" +
			"g|" +
			"gh|" +
			"gn|" +
			"k|" +
			"l|" +
			"ll|" +
			"lv|" +
			"m|" +
			"mm|" +
			"n|" +
			"nc|" +
			"ng|" +
			"nn|" +
			"p|" +
			"r|" +
			"rc|" +
			"rn|" +
			"rs|" +
			"rv|" +
			"s|" +
			"sc|" +
			"sk|" +
			"sl|" +
			"squ|" +
			"ss|" +
			"st|" +
			"t|" +
			"th|" +
			"v|" +
			"y|" +
			"z" +
		")" +
		"e$",
	)

	expressionDoubleSyllabicOne = regexp.MustCompile(
		"(" +
			"(" +
				"[^aeiouy]" +
			// Remove unsupported backreference
			// Replace with {1,2} instead (one or two repeated consanants)
			// Original with backreference: ")\\2l|" +
			"){1,2}l|" +
			"[^aeiouy]ie" +
			"(" +
				"r|" +
				"st|" +
				"t" +
			")|" +
			"[aeiouym]bl|" +
			"eo|" +
			"ism|" +
			"asm|" +
			"thm|" +
			"dnt|" +
			"uity|" +
			"dea|" +
			"gean|" +
			"oa|" +
			"ua|" +
			"eings?|" +
			"[aeiouy]sh?e[rsd]" +
		")$",
	)

	expressionDoubleSyllabicTwo = regexp.MustCompile(
		"[^gq]ua[^auieo]|" +
		"[aeiou]{3}|" +
		"^(" +
			"ia|" +
			"mc|" +
			"coa[dglx]." +
		")",
	)

	expressionDoubleSyllabicThree = regexp.MustCompile(
		"[^aeiou]y[ae]|" +
		"[^l]lien|" +
		"riet|" +
		"dien|" +
		"iu|" +
		"io|" +
		"ii|" +
		"uen|" +
		"real|" +
		"iell|" +
		"eo[^aeiou]|" +
		"[aeiou]y[aeiou]",
	)

	expressionDoubleSyllabicFour = regexp.MustCompile(
		"[^s]ia",
	)

	expressionSingle = regexp.MustCompile(
		"^" +
		"(" +
			"un|" +
			"fore|" +
			"ware|" +
			"none?|" +
			"out|" +
			"post|" +
			"sub|" +
			"pre|" +
			"pro|" +
			"dis|" +
			"side" +
		")" +
		"|" +
		"(" +
			"ly|" +
			"less|" +
			"some|" +
			"ful|" +
			"ers?|" +
			"ness|" +
			"cians?|" +
			"ments?|" +
			"ettes?|" +
			"villes?|" +
			"ships?|" +
			"sides?|" +
			"ports?|" +
			"shires?|" +
			"tion(ed)?" +
		")",// +
		//"$",
	)

	expressionDouble = regexp.MustCompile(
		"^" +
		"(" +
			"above|" +
			"anti|" +
			"ante|" +
			"counter|" +
			"hyper|" +
			"afore|" +
			"agri|" +
			"infra|" +
			"intra|" +
			"inter|" +
			"over|" +
			"semi|" +
			"ultra|" +
			"under|" +
			"extra|" +
			"dia|" +
			"micro|" +
			"mega|" +
			"kilo|" +
			"pico|" +
			"nano|" +
			"macro" +
		")" +
		"|" +
		"(" +
			"fully|" +
			"berry|" +
			"woman|" +
			"women" +
		")",// +
		// "$",
	)

	expressionTriple = regexp.MustCompile(
		"(ology|ologist|onomy|onomist)", //$",
	)

	expressionNonalphabetic = regexp.MustCompile(
		"[^a-z]",
	)

	consanants = regexp.MustCompile(
		"[^aeiouy]+",
	)
)
