package syllables

import (
	"testing"
)

func Test_CountIn(t *testing.T) {
	cases := []struct {
		want int
		in   string
	}{
		// Original test cases
		{3, "syllable"},
		{3, "unicorn"},
		{1, "hi"},
		{2, "hihi"},
		{1, "mmmmmmmmmmmmmmmm"},
		{2, "hoopty"},

		// Additional test cases
		{6, "syllables in this phrase"},
		{31, "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."},
		{108, "If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your shell's path."},
		{2, "template"},
		{4, "abalone"},
		{5, "aborigine"},
		{3, "simile"},
		{4, "facsimile"},
		{3, "syncope"},
		{0, ""},
		{1, "yo"},
		{3, "kilogram"},
		{112, "If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your Scientology shell's path. "},
		{115, "If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your Scientology ology shell's path. "},

		// Problematic test cases

		// local nodejs:        284
		// wooorm.com/syllable: 273
		{283, "As of old Phoenician men, to the Tin Isles sailing Straight against the sunset and the edges of the earth Chaunted loud above the storm and the strange sea's wailing Legends of their people and the land that gave them birth Sang aloud to Baal-Peor, sang unto the horned maiden Sang how they should come again with the Brethon treasure laden Sang of all the pride and glory of their hardy enterprise How they found the outer islands, where the unknown stars arise And the rowers down below, rowing hard as they could row Toiling at the stroke and feather through the wet and weary weather Even they forgot their burden in the measure of a song And the merchants and the masters and the bondsmen all together Dreaming of the wondrous islands, brought the gallant ship along So in mighty deeps alone on the chainless breezes blow In my coracle of verses I will sing of lands unknown Flying from the scarlet city where a Lord that knows no pity Mocks the broken people praying round his iron throne Sing about the Hidden Country fresh and full of quiet green Sailing over seas uncharted to a port that none has seen"},

		// local nodejs:        808
		// wooorm.com/syllable: 755
		{805, "My first interview with the manager was curious. He did not ask me to sit down after my twenty–mile walk that morning. He was commonplace in complexion, in features, in manners, and in voice. He was of middle size and of ordinary build. His eyes, of the usual blue, were perhaps remarkably cold, and he cer- tainly could make his glance fall on one as trenchant and heavy as an axe. But even at these times the rest of his person seemed to disclaim the intention. Otherwise there was only an indefinable, faint expression of his lips, something stealthy— a smile—not a smile—I remember it, but I can’t explain. It was unconscious, this smile was, though just after he had said something it got intensified for an instant. It came at the end of his speeches like a seal applied on the words to make the meaning of the commonest phrase appear absolutely inscrut- able. He was a common trader, from his youth up employed in these parts—nothing more. He was obeyed, yet he inspired neither love nor fear, nor even respect. He inspired uneasi- ness. That was it! Uneasiness. Not a definite mistrust—just un- easiness—nothing more. You have no idea how effective such a ... a... . faculty can be. He had no genius for organizing, for initiative, or for order even. That was evident in such things as the deplorable state of the station. He had no learning, and no intelligence. His position had come to him—why? Perhaps be- cause he was never ill ... He had served three terms of three years out there ... Because triumphant health in the general rout of constitutions is a kind of power in itself. When he went home on leave he rioted on a large scale—pompously. Jack ashore—with a difference—in externals only. This one could gather from his casual talk. He originated nothing, he could keep the routine going—that’s all. But he was great. He was great by this little thing that it was impossible to tell what could control such a man. He never gave that secret away. Per- haps there was nothing within him. Such a suspicion made one pause—for out there there were no external checks. Once when various tropical diseases had laid low almost every ‘agent’ in the station, he was heard to say, ‘Men who come out here should have no entrails.’ He sealed the utterance with that smile of his, as though it had been a door opening into a darkness he had in his keeping. You fancied you had seen things—but the seal was on. When annoyed at meal–times by the constant quarrels of the white men about precedence, he ordered an immense round table to be made, for which a spe- cial house had to be built. This was the station’s mess–room. Where he sat was the first place—the rest were nowhere. One felt this to be his unalterable conviction. He was neither civil nor uncivil. He was quiet. He allowed his ‘boy’—an overfed young negro from the coast—to treat the white men, under his very eyes, with provoking insolence."},
	}

	for _, c := range cases {
		got := In(c.in)
		if got != c.want {
			t.Errorf("syllables.In(%q) == %v, expected %v", c.in, got, c.want)
		}
	}
}
