package syllables

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_Expressions(t *testing.T) {
	cases := []struct {
		in   *regexp.Regexp
		want string
	}{
		{
			expressionMonosyllabicOne,
			"cia(l|$)|tia|cius|cious|[^aeiou]giu|[aeiouy][^aeiouy]ion|iou|sia$|eous$|[oa]gue$|.[^aeiuoycgltdb]{2,}ed$|.ely$|^jua|uai|eau|^busi$|([aeiouy](b|c|ch|dg|f|g|gh|gn|k|l|lch|ll|lv|m|mm|n|nc|ng|nch|nn|p|r|rc|rn|rs|rv|s|sc|sk|sl|squ|ss|th|v|y|z)ed$)|([aeiouy](b|ch|d|f|gh|gn|k|l|lch|ll|lv|m|mm|n|nch|nn|p|r|rn|rs|rv|s|sc|sk|sl|squ|ss|st|t|th|v|y)es$)",
		},
		{
			expressionMonosyllabicTwo,
			"[aeiouy](b|c|ch|d|dg|f|g|gh|gn|k|l|ll|lv|m|mm|n|nc|ng|nn|p|r|rc|rn|rs|rv|s|sc|sk|sl|squ|ss|st|t|th|v|y|z)e$",
		},
		{
			expressionDoubleSyllabicOne,
			"(([^aeiouy]){1,2}l|[^aeiouy]ie(r|st|t)|[aeiouym]bl|eo|ism|asm|thm|dnt|uity|dea|gean|oa|ua|eings?|[aeiouy]sh?e[rsd])$",
		},
		{
			expressionDoubleSyllabicTwo,
			"[^gq]ua[^auieo]|[aeiou]{3}|^(ia|mc|coa[dglx].)",
		},
		{
			expressionDoubleSyllabicThree,
			"[^aeiou]y[ae]|[^l]lien|riet|dien|iu|io|ii|uen|real|iell|eo[^aeiou]|[aeiou]y[aeiou]",
		},
		{
			expressionDoubleSyllabicFour,
			"[^s]ia",
		},
		{
			expressionSingle,
			"^(un|fore|ware|none?|out|post|sub|pre|pro|dis|side)|(ly|less|some|ful|ers?|ness|cians?|ments?|ettes?|villes?|ships?|sides?|ports?|shires?|tion(ed)?)$",
		},
		{
			expressionDouble,
			"^(above|anti|ante|counter|hyper|afore|agri|infra|intra|inter|over|semi|ultra|under|extra|dia|micro|mega|kilo|pico|nano|macro)|(fully|berry|woman|women)$",
		},
		{
			expressionTriple,
			"(ology|ologist|onomy|onomist)$",
		},
		{
			expressionNonalphabetic,
			"[^a-z]",
		},
	}

	for _, c := range cases {
		got := fmt.Sprint(c.in)
		if got != c.want {
			t.Errorf("Got %q, expected %q", got, c.want)
		}
	}
}
