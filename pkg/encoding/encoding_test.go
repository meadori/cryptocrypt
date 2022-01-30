package encoding

import (
	"bytes"
	"testing"

	"github.com/meadori/cryptocrypt/pkg/encoding/hex"
)

func b(str string) []byte {
	return []byte(str)
}

var testCases = []struct {
	input []byte
}{
	{b("")},
	{b("e\x05\xe8\x96\x14W\xf0\xbe\xab\xd1vO\xa44\x0b\xa1|\xc5\x7f\x02\x93")},
	{b("\x0b\xf9\xb9\xfe\x82>\x92H\xdc \x004\x8fT\x10\xd9V\xc3\xca\xf6\xb9\xb2\x19*\xc1\xece\xe0\xab-\x98r\x1e1\x85\xe5_G\xa1T")},
	{b("\x1cC\x9a\x03\xc2\xeasX\xc6\x01\x16\xf3\xb3\xe4\x84\x87\xef\x9d$\x93\x04,")},
	{b("\x9d\x1e\x9a\xb0N5\xa8\x05\x922\n\x83\xa2Z")},
	{b("p\x976\x83\xe7\x89\x10\xca\xf5H5\xc7\xc1\xa3i9w\x9d\xe2\xcf\xd6\x8e\xa0\xd8O\xa3BF\xc8\xc9\xf2O\xde\x18aTM\xa72\x81\x84\xb3")},
	{b("\x86;3_s\xb9")},
	{b("\xc9G\xa4\xae\x04w\xa6C\xf2\xafe\xe8T\x1d9;\xc6\xf1\xa8\xecU\xa7\x86\x8c1\x897\xef\x82\xb1@*\x18v")},
	{b("A\x7f\x8d\nR\x0c\r\xe3\xdc\x85\xcbWE\x0e\xfa)\xc7L\xd0\xaa7'\xbd\x10/\x7f\x15\xafM+Se\xaa\xa5\xac(\x87~\xa95")},
	{b("X")},
	{b("~\x97\xdeS\xf9\x14m\xff;\xb7\xb5\xd2\x89`\xca\xecw\xa4}")},
	{b("\x7f\xb6\x80\xcfzt\xc7n\xab\xaf(@\x84\x05\x99\x18\xbe\xb3\xf3\xdb\xcb\xe9\xec\xee\xc9\x9c{\xa0\x14|V\x9b\x81f\xe8|")},
	{b("\x10\x11\x88\xf5\xde")},
	{b("\nO\x1a\x98\xe7\xa2\xff\xad]\x05\x9d\x00\x864\x8bL\x99x\xe6t\xf5\xa3\xd8Y\x91.\x87=~O\x1c\xa0%\xcc\xf76_-&7\x11i\xa7\x853\xb5o\xa7\x1e\x98\xd4\xa2\x12\x00\x0c\xa91")},
	{b("\xa4L\xfc\xec[\xc6\x98\xee\n~\xb4^\xd3")},
	{b("(\x15p\xac\x18\x03\xd5\xe9H\xe4\x17\xe6\xc9y\xebs+d+!\xd1\t\xa9V-\xe9v\xf8")},
	{b("\xfbg\x88{z\x08\xbbs\xa6\x90x!e\x13\xf5\x1b\xdd\xa6\xe9")},
	{b("\x01\x9a\x9e\x8c\xef\xbb\xa3Be\xb4\xc8*")},
	{b("\xb1j\x7f\xbc\xb5")},
	{b("\x80LU\xde\x1d9\xad\xda\x9dH\x91\xe3JPsJfI\x83Lx\x12\xa8\xed\xc8\xa0\x0cw\xf6\x1a3\xdc\xe1\xa8\xba.\xd5\xab\xe3\x1c\xf1{")},
	{b("\xeb\\\xe5CvU\xd7P&O\x9esDJ\x87\x05\x8fZ\x8bh\xbd\xac\x85m.\xf3\xaf\x90\xe0d&\xcb\x0f=\x86\xed\xc6[\x89E\x92\t\xfe\x1de\xf3\xa9\n\x86\x8a\x05eZ\x18\xb8\x96\x03\x87\xb1\xf1\x14\x10\x7f\xcb")},
	{b("\x01o\xe1&[\xac\x14>\xf9\x14N\xbdZ\xa3X\xd8\xb3")},
	{b("?*")},
	{b("\x10\x0b#%\xdf\xe2\x8d2\xfb\xbe8\xb0n\xfa")},
	{b("\xdaP")},
	{b("&d\xf58\xd2q\x03\xfb^\x19Il\xd5\xd9\xe1  k\x9f\x97qj\x81 \x18J\xcc\xbe\xab\xf3SK>\xd1%\x91\xce\xbeo\xe5\x00_\xa7\x83}4\xdf\xc1X\xf0\xef?<\xb4\xa3,~^\x00\x0f \xab\xed\xb8")},
	{b("l\x0b\xc2\x90Aoz=\xfeF'Dn\xe9\xf2\x16\xf9!\xac\x80\xe0\xa6\xde\x06Z+=\n\xb8;j")},
	{b("/\x8bU\xd5 J!\\\xdf\xd5[")},
	{b("\x81\xaa\x8b\xe2-|7\xb3\x0e\xd2")},
	{b("]p\x85")},
	{b("X\xc6\x98aj")},
	{b("N\xde\x8a\xb8\xa9\xe2\x06\x0b#\xf7\xf6\x93\x83F\xb2")},
	{b("X\xd3\xc4\x86\xb8\xd5\x15")},
	{b("\xc1\xd8\xbf\x88\xf2\x8d-,\x11\xc0:\x1d\x96\xe2\xd2\xb2X\xf1r9\xfe\xc9\xb2\x1a&Q\x1b\xa7\x8f\xb5\xad:\xe1Ap\xccI\xaa=\xff#F,")},
	{b("\xaa\xdd0\xe6\x90&\x9bI\xd6n|{\x18No\xc8\xa36\xda")},
	{b("\xc2\x87\x18\x08\x89F\xf8\x05\xd5Xm\x7fC\xed6\xda\xad\xedm\x0f\xf8\xbe\xbe\xf4\x07E\x12\xc53\\\xcd\x02R\x108\x90\xc9\x07'")},
	{b("\xcc$f\xfd\xa0L\xf2>\x80\xf5")},
	{b("\xb9\xf5\x96\xb26\xfeP\x0f\xda^\xd5l\xc2\x978;\n\xd3\xcf\xc9\xa0r\xa9J\xfd\xbe\x0b\xa6\xf3\xb7\xc21\xbfX\xf5Oy)\xca~\xc5x\t\xa1D\xcb\xb2\nG\xbd\x18T\x02\xfe\x82e\xa7\t\xdc\xbc`\x8d\xf7\xa63S")},
	{b("gt\xa0,\x1f\xac\x02\xe2\xd6\xc0o\xdf/N_E\xa0?\xa7\xd9\x80F\x0c\xd7\xb6\xfd\xd5\xca\x9f:p)")},
	{b("Gd")},
	{b("o\xc1\xea\x94\xe7&\xdf\xda~\x05\x94\x05\x93")},
	{b("\xaf\xf7:\xc0\xf9\xddmT=\xbe\xff\xbc\xa8\x07\xbb\x0b\x12\xf2\xe6\x93e\x0f\x85\xa8o\xb9")},
	{b("{\x80\xac\xd1_\x88Q:B$X\xb9\xcd\x00\xf8q\xc4 \xf4\xf2m@\x83:\x9e\xabc\xf8c{\xbd\xdf")},
	{b("\xe0\xa1T\x7f\xbb\xcfy\xbe\xb7<\x1c\x14\xf6\xe1C\xa4\x16\xd9)2\xdb\xc7\x9f\x1dM\xa2\x04H\xf8\x96\rRu\x18\x0fR\xed\x9c\xab\xd9\xf4\xcf\x1e\x91w\n\xa7\x98")},
	{b("P\x94z\xa9y\xad\xaf\xde(\x96\xdaV7U\xf1~\xc1\xae\xf7\xf2\x1b\xdb$?)!\x89\xba\x0b\xce\x99e\x03\x82i\x1f\xc9\xbf\xea1$\x00\xa1d\xb6\xc8\xd4c\tw:")},
	{b("\xaa_g\xc9\x82\x0c\xf6\x1a\x9e\xbf\x9b\xb2\xd8Q\xd3\xc6o\xcb\xc4\xb1\x05\xe49G=*\xb5\xd1\xf8\xe1\xb8\xeb\xc7\x1b\xb5\xf0\x95\x96\xe6\xa8\xda\xae#J\x7f\x1el\xb8s5Lo\xdb>\x9d\x9c\xf8\xa9N\x87\xa6\x8c\xcf")},
	{b("\xfa\xbe\xa7\xad\xf3\x14\x19\x89\xe2\x16\x80\x1b\x89\x91k\xfc\xc6H\xab\x90\xa4\xbc$\xcfjI\xff\x8co\xc2(\xc2\xbbk3\x11\xe4\xda\x93\x92\xd3\x1b^*\xf1x\xe0R\x0fEw\x10c\xeal\xf27\xf0`\x81\x9f\xa0;")},
	{b("qy\xe3D\xe2E\xb2mJ\xa0\xc9\xa2\xfd\xd9\x11\x12G\xf8\xf4\xe5\xc9\x8e@\xd4*\x08\x1al\xe4f\x95\x03\xa0\xd2\x83\x9e\xdbf4\x89\x87\xc7\xaf\x12<U8l&\x14Oo\x06A\x98~")},
	{b("O\xfd_\xd2}\x8dC\x83\xe6BPj\xb9\x80\x9d\x0e\xbby\x85W\x1bl\rA\xb4\xb0`\xd5\x00\xbc\xdbt\x9e:\xc4\x0f$\x9c}\n\xd7\x9e\x80\xb78\xb6")},
	{b("\xe2\xf9\x97k\xf2\xa7\xffP\xb4\xfb\x11\xa9\x89\x81\x95OR\xc2\xaf~{\xe1\xfd\x08]\xe1\x83\xae\xd6\x96=\x9bt-s\xe9\x84\xbe\x05\x025\xb6+rd")},
	{b(">\x18\x16\xe0")},
	{b("\xc2\xbeC-A\xdbj\x05\x0f/\xd8\xcfS\xc3S\x8c\x02\x83\x95\x91s\x018!\xb8\xd9\xb2\xfd\xc3\xa8\xfb\xb0#8\xf0\x9eW\x0cgo\x07\xdc\\\x10yy\x85\xa80\x9b\xab]\xcab\xbc2")},
	{b("q\x068s\x0e\xb5\xd1")},
	{b("\x91N\xe2\xd1\x17\xe2Kd\xdcNN\x00")},
	{b("\x94\xd9\x8dc\x0b\xbb\x08\xe8\x8d\xf01\xcd/\xbc\xd4\xddK\xee\x06\x19\x18 \xa9\x9c\xedz\xdb\xe9\xb6\xbd4\xb8\x90]\xa7\x16!C\x85-\x9a\x97\xb1I,{\xd9\xa59^\xe5n\x9dN\x9c6\xbc\xad\xa3J\xbbsN\xb7\xf4\xd3")},
	{b("")},
	{b("\xba\xe5\xaf\x91\tW';\x9e\xeal\xc2<\x80\xcfH,\xe9#I\n\xee\xd1\xcd~>-\xabt\xc0A\x95\xda^n\xa4\x98\x86\x1e")},
	{b("\xbeP\x17\x90i\r8|W")},
	{b("v8\xcd\x7f:\xa2Q\xdd\xf9 #\x85[\xce\xb6\xb5\xf8\xbb\xd1\xca\xd5\xe9\xf4\xe6\xa9")},
	{b("\x15\x91V")},
	{b("\x94\xcf0[+OJIN\xb1pg\xf3\x87d=\xb5\xc7\xb9\xd0\nd\xd8JE1j\x80\xe9\x83R\x87\xe9]\xd0<\xb0\x0b\x1c\x06\x88\xff\xee\xed;i\xed\xb7\x93\xf7vW\xfe,&\xde\xba\xa5\x1d*\x83Ff\x85\xf8G")},
	{b("\xd3#\xad\x10\x17\x90l\xe9x\xf3")},
	{b("\xb0\xa5\xff\xbf\\\xe0\xcf@\xef\xa5b\x8d")},
	{b("\x1c\x18\xb0\x06\x9d\x8e\xe7\xf54ii8=\x03J\x14\xcbw\\\x9d\xfc\xa4\x03\xce\xb3\xb9#\x1d\xec\xc6k\x0b\xdb\xe8\x99\x81\xbeVw\x8eP\xc4\xa9+*\xb6\x9fD\xa9d\x7f\xb3\x9b\xa7\xd2\xee\xa5\xf08")},
	{b("\xb9\xa4\xd7m\xe4\xbfJY\xbe\xbb\xb05yTU\xd5\x0c*\xcb\xac\x0e\x9f\x96}\xc1y\x81vEB&\x9b\xf8U\x06\xfa\xa1>\xa3\xa0\xfe{\xfd\xc0W\x06-\xe6E\xc4\xddYP&\xc1E\xb1\x1e\x17<\xcb")},
	{b("&oC\xa3cuK\xdd\x81\xee\xfe\x8e\xdd\xa5\xcc\xec\x04\xeb(\xc1\xe5\xaa\x05|>\xcdn5g\x10m\x93\x8cU]?\x10\xaf\x00\xa8\x91\xad\xf4\xcabrY2\x88D\x1e\xca")},
	{b("\x84\xdf\x89\xc7\xd4f\x82Q\x00\x85HB\x15 U\xfb\xcc\xe7@\xfd\xf1\xaa\xed\x9dVp\xf7S7u\x03\xff]\xbe\x897\x04\\{\xf0J)\xfa\xc6\x9d\xb6\xdc\x88\x10\xb1d\xb9y\xc9J\x86a\xc6")},
	{b("\xf80&rJ\xb6\x15b\xd2@\x8cI\xb3Ck/\xa5\x82wn\xd03dO\xd2\xb5\xbf\x018\xb9\xd9\x024\xe0\xcd\xc0\x97\xa0\x02\xd8\x13\xd1\xcb\n\xdcp\xbfn\xc2\xf8\x15+\x88\x97\x82\xbb\xd6\xb6Sy\xe6\xda\x1d\xcb")},
	{b("3\xbc\x96\xaaI\x07\x16\x13\xf5\xcfU\xfc")},
	{b("\xf1\x98\xe7YaiL\x1b3\xc1\x10\xcd\x9a_")},
	{b(" \xa7p\x01\xbe'\xa9\xf2\x11\x16\xe2\xf5\xc7'\x0f)\xa5d*\xfb\x9d\xb1\xae*\xfes\x173DH\x85\x9eT\x1aW\xa10\xe2t\x1c\xb1\xa3[pX\xc8\xe2\x842W~=")},
	{b("+\x83+]\x08\xf8\xabi hqS\x0f\\y\x02\x9b60\xb7X\xd8\xa4\x9bn%/!U\xa9)$~\x1a\xe5:~\xa2\xef\xc28\tR!!\xb9")},
	{b("\x87\xf2\x86\xf9o&i\xb3g\xc5\x96\xa8>\xed\xb6\xa6\xee\xde\\FV\xba\xb2\x05\xec6\xadG\xfa\x85\xaa\x86\xba;\xa8")},
	{b("\x02\xd9\xa7\xba\xda\xbc\xbf\x85\xa8\xaccx>\tc\x15=\xcd\x16\xf1\xa3\x14\xde\xd1\n\x04\xa9\xc7\xca,E5-\x07")},
	{b("\xd8\xc4\xb3\x8d,\xe4X\x1d\xf0\x95\xab\x88'\xef\x9aw\x139\xffT\x99}RN\xf5m\x9aS\x80EW*g\xce\xc5\xa2\xe6}\x95\xd0\xfe\x87\xed\xb5\xef\xf6{\xb6\x11\x9fKa\xdf\xf5k")},
	{b("\x0fe\x05\xcc\x06\xef\xfeP=X\xa9\xa1}\xab&\x82\xcb?g")},
	{b("L\xfd\xec\xb4\x0cJ7H\xb8\xb1,\x87\xbb\x15\xbd[\xef\xaf]\x99\xcf\xe4\xb4\xb1\xa6\x9e")},
	{b("1\xab\x02\xae\xea\xfe\x11\xba\xba\xbepf\xd1\xcc\xc6\xb8d\x1f\xf3<\\\xe0\x1f\xb1\x06\xb8\xc9\xbfZ\xd4\xf6\xec\x9f\x9f\xfcx\xe9O\xf3\xf6\xb4vb\xdb\xe4\xa2:6\xa1\xd2\x04\x90\xff\x92n1]\x8dS\xc3b\xce\na")},
	{b("\x08\x82Y\xd5\xfa^Tj\xd0\x16b\xd4A>J\xa5\x92\x0c/#\xe9\x16\xb2{M/\xec\x01\xd1m\xcc\x05\xa9\xb92X\xe5\x96D1\xd4\\\xb5\x1aN#A\xc0tZm\xcbK\x03Ik\x11>")},
	{b("\x91\x14\xea\xe0\\\xfb\xd1\x87\xa9\xf97F\xee{\x91\xc0\xfbf\x10n\xdf\x82\xcd\xc5\xb1\xa8\x18:p\x9a\x12\x02\xd3\xac\x81\x9b")},
	{b("\x1b\xae\x92m\xc9")},
	{b("&\xc9\x11\x1c\x1a\xee\xbep-L;\xe8\x17eb\xaft\xd0\x99q\xb8")},
	{b("\xa5P7Q{W\x15i\x12L")},
	{b("R\xfa\xc3\xde\xdd\xf0\x9bZ\xc39\x07\xe8\x96\xb5\x95\xc7\xbbM\xce\xaa\xfe2C}C\x94\xfcz\xbd\xf9\xa0\x9f\xe1\xea\x97q>^\x16\xce-\xa4ex\xb8\xfc6\xce\x10]+\xc7\x160\x06\xecE\xb2\x0f\x1a_\xbf\n\x1b\xea\x9a")},
	{b("\xd3\x88\xec\x06\xfabf>")},
	{b("_@\xe4\x08\xb5\xd0\xd2(")},
	{b("$x\xe1\xd6\xf1\xf5bS-G\xf0[\x12\xba\xf77\xe6h\xa6\xa9\xab\xf9\xa0\xf0V\xb7\xe7\xc2")},
	{b(";Y\xe1\xf8\xc6\x99s%`\xce\xf8\x049L\xb1\xb4\xc2\xc9\xbb\xdb\xac\xf7\xea\xdc\xda-\xe4\xfb")},
	{b("0\x87\x8f\x95\xa7\xb5\xc9\x1d\xb0T\x1f")},
	{b("\x05\xbf\x05Y\x0c\x8dW\xf8\x9fs\x973\x90U(\xd1U\x14\xda\x9e8u_4\x16\xc5\x85b\xcaO")},
}

func TestHex(t *testing.T) {
	for _, test := range testCases {
		output := hex.Encode(test.input)
		input, _ := hex.Decode(output)
		if bytes.Compare(test.input, input) != 0 {
			t.Fatalf("HEX(UNHEX('%v')) != '%v'", test.input, test.input)
		}
	}
}