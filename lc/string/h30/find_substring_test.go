package h30

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	testCases := []struct {
		Input  string
		Words  []string
		Expect []int
	}{
		{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "good"},
			[]int{8},
		},
		{
			"foobarfootool",
			[]string{"foo", "bar"},
			[]int{0, 3},
		},
		{
			"fooabarbfootool",
			[]string{"foo", "bar"},
			[]int{},
		},
		{
			"fooaaaabfootool",
			[]string{"aa", "bb"},
			[]int{},
		},
		{
			"fooaaaaabfootool",
			[]string{"aa", "aa"},
			[]int{3, 4},
		},
		{
			"xjpguhvytyjcknhjqkwelhjqbdgtwxgvgxbdeydxwozidiutuqafxjxaodtkdbfjyiocgtbfhcplmjggbgoarlcgpxssyadyiuapndwxhlitvoayvqzobbuqzpkzpqyzkaqzgmwnyghvvjtszuiawdtxufylvwkhzbhfpfsnmbkjkedlylowqjvkquxmsivrlewakrqysahfgmqhxgfqpbcgxaupkrhvwfviwngrqpwybohaxnsoqvwpxqehkncgvzqtpwkflidoznqwcjksehjdzpkjdmranhtcfejsopgncxjeguymbhpcwbmbpfbcnvhsbqnpftdjsonainoludqtgcwvjyywvhryxepfzuqsjgstthhqmxltbhokfojcvcavgqchmszvyupudykrvvmwedikrroptrmbjojvgkrheibjgnbdknboqjakbpbwgnyrbhmjtfqantjvgmaqcbhulhgowhkeukvxrkhnpznfvwcdldwnedjpkqfjxqnualruvahmcwrxuuafxwubzetmwyvtqkntvhnshwhjsyimujuthoxjuqvqqqmhazayipsqnzbfaktuvpocennadirvadcdeedpvvfixipxujtpajugwhhbsaxsfbvliaadwhmvqbsmmnenxavvhcxbcwwjxtvfuvlqdxlvafhpsnernznxemygiuqfonniiyanxnkzuuoohugvwvsajsirnyydnnnwnplkcwkyqamxvuurrmrafztuauzphmlvdzhfvrflurkpmfidtbgycbuevtufhhakgjrdbwqvqbmciwhbxpcbrwgmscrbjtmsffvgemdupryxphaoxcpobxcvbwwnrkfwscewqjsfcqerzffwjxmmwrhynelgosfiujenvwsxozpogwmrtbeqslqhrbnitsqpevcztxykynaemmvhnbzhnpogqeolyfdccfdxecjcrjidyelnhmvuclduprioylscswaxylbpvkvvqikxuhuytxtkqbapottgrvfphjgetdzjljigrcembzwsczjqsczlygcfpijkmktzvehmgoaknzcqylisnjdlqfshpbsdnndjrkxayykoxogxzqpoascsxubmytsljvuahucisowrccobudsuxuouoqimlaauxwxhqbpkqldsptwjyogviurymclyenueltlcvaollufcnbnmptjzqbycflcjyxnjsynnaealygpljdzzyjyomyrtjvchustnsgctkdgklwwubxvziwouuhpecslxmgmepoxbremcckzdhucqqqmlzcpcwcbilnmabkbtqpxszwvhtzzjslwrnntlsutdjgflsigkyfcxezexydiqrfigudsmalrjtwunfcxdibcmajjbotrfybmtfghftzqpxlcepcjxdmlgvwhjqarqcdlhltoeuettryyhahgfvsnqqucgxtzzykijfwpbcjvujvdjelqadeswawcxpdwpoeyvcqxfzubipetvpjxvpqtqmxpebotpuumxkjelffvwlosczpzrhsjwqycrmvihrugbgkolrjiezcgbtisbadzsbblqytzsqfvyrklitvmvxuyrcqufvvzwyloygnqwsmwjwitrdhobcmugcqnzlnwlykjeaadsmzekhxdlhsojekrjafinseysrjyrjblxbrjkrnvyflhjvasxfbkzhkraustdtfdwymhpzengqwqnxklelvetixvcpphjwkhuzokavxhlwzatjlxxjdqrbnvsccdypltqzdswcbhyaktmxrjgwbzxowqrzvpqgkiipaescoscymovfxebyfbpctgdoxvxidfxdjrfzmkxaavhabiyilpkevpvvksfpzetiwakkkjklgrlhblqnbctyuqtgkawjfrubrenenxpuqcdrptgsyctusmadnyospivhminahewxgzoxvxqtzjntxpymongdvdmknzkudirlhijchbxgkmbjcawsnevkikuvjgspolcyvlacmakymmiqmgibkensqiqbqiqfttdpgfrvfevsqdkelthwzuqpegqvqjakefbmkuhsyfmokwswpbsqwkfatyvjjxvncwzprjhpoteypywhcqxybfaufyfovbbaxcponygdrkeikarmrrmuwnqblvpiwsiuwzkkxqnqctbpusdnlqhhfxkssbapvllskvekmtcqndfhyjujbdtgafauhclenwwaucmiwoyjugqupmfspaarganpcztqxssruebqucbqirkzfsrwsrnardpvclnzfftblusgyvwgnjfudyrvpgwijngnatnfbihmebudwtjlerihrbchjartqzistxyufhikkdpiwauarejjfnsooljglsygpyaxhijrnyalywnsawdfkxtaidgvxgbmhdloougbsipteclezqljnejvjrtgzuvgygvoddrxlgqrjdxititgoeeavxiwrfdroahrdzoqfhhokgygormevsespnpjsscgukzxjopoxyfjedpuxeyswfnoucxmwbvqlwpwmgljestkviesoennjabfeauabpsnljjapwjvochmnngbrvodxribredttvihgthxsssivbwkodniaelyvvzpadkvasejnngfgbqdcmprpczqgmoejptlsdjvxpekdmslniqqufjmhieqwuufjntescbpthbyttjhdbzaiiosssioocvzrqdjugaonbmhxyqczpcixqarkkfaocaftfqnmsbbtqisoyvppxzoqbfclmdzpdgkiyxwqbymtiehjzyyzynrzutnhymwbvimvhkmiiadtekcgjafjpyikrvtqkrthzhcgsqrcquvxhxdsakbrkldbjwttnpcowgvqzotriqorotjqfmhpylthhocxdelcmiulwpdhgtywpkmuwvmugfbqtfpzlcdylxjhnoovkprzzdtvafqjmtbizqhmsmkdlwnykdtusmvrrpnswfbjacrbuaommysxwhyjktdfgzwzqlrmssxtwowqqkfclxchgcqqvwvdxudnhwbarzvnpregclknkowqqniojgtgayvhvyjozebpwhxasjncajuqydghjcplakuxlelkipbgwygrkvvkfqcdvlnenerpplpapcmatogqmnjyiekpwpvrakxpoqgfcxhtcutvicnwrwvbdhtbwovyaupolyunxdizxcvfgiezhbamitnhjkhjfxaqxwfuznuzppgxzkwilxuuskdewkpbhprenwbpkvobmubnfxwqwsmrepvbakejcwqpuregmukaplnuklmjgzamqxpqjualsqdmhjvvefxtskpeybngcpstmilweljwdoimyfhcmgxermlrpyxuqrnebycfmmbpamcyrlceszkllvedwbxmumqwktbyhdojrskidmoxmbizymeupbimnbiawlydoomfgyqmlgjzhuygifcagnmwowykhypyndfvcvhpetolpotztybclpyblwlvuctjhyflwoaajonydhawfbysrytewgztiucrvhdrydthsgixpkvwlwoeujlrpmkzhorcywvwzoftwnsoxoklkbrekcxcrjdyywcwszsupxnlngbmwmxgprmbvkdmthmrdqnyphsehhsuptilhiryzeauqdhjmtdsmqqbakihtcdjxluhtofsufpklwvxryrdrjhrtpyntdyqouxkideeitotrmtlkkqbuxsposchvaamxxyfccknyairmbczovaiuvzjneslguzdsxjwbvjzxsrmvvljqntlitwyxqldlkjfjsbkpnmohfaecnqtblgleelduwjhismtmqgdfurozusbhkwkweyckjihitosldozvuccovqppksxvrjtxhvitdrbwfvjkjkhdmjtkbizodyluietpzbifslbahnmqxuwmfpwjaxzdwkzeqstrweworaqypfrmznagewreuqjqaiwsdrkzvgpnignxnemotmuylmcheozhyvzbmjaksqzcyoclvozocvmnjrwofvvdswhhghtazucziekdulsxjgkszjieefkxcrekaxkatozbtmhnzbmihzdhinnmtzlxsrjtqtvjjwleksukvgucfzlnpbcianhthqoxllhuhuzsotejbanhazwpcyzcoixvanulydhgxganbeydgmminizphatxitsigmvfqdnplnfptdszrgieohvxirwskodqdyxvdkmpzresxyuoeevunsuxjqqthvkmthhxuvotnsoksiayovsboobzfttoofahmhggcucroqdgaeeqbzrppupunkkbpkldtrkymopcgvjgzpwaopsekjaxtlzixkltdxrrliurddzesxfjnzpzipwbcxlcjwvpwmghwabafcgyanjnmymupkxukiwvhtkdhrmdrnfxsmxszihogtixfirpsplzixcrorvigcfyqeqqmxeusoraylprccsnaveqobyueftullmxjstdjndhavacztpzqusevqybwtwhfihodctmpxvpswurpjthfllddlezfcjknsaquvcmsxdmvzemjztqkgtpsarzcalpunhqiledlipgjttsuolgvewpenohnbyjogzyrebeorlxmgshudnpjjgowwxlxxunfwmzapdqgonvuhcrkriubpkzljnlghymdmlfcqvkflfbsjsfbdbculdfwqscatqffdljuiubvbcqlxvmcwqwjvbhmwjmpcrufegbpackdhaoexcgvucgqfncbzqsbjniotkfvmpytspzprflmjrerhgugynhhapxvzcsosqhmhjbzqonaittpznvzaegctezvgrjaksorbsssghuqanhbaeadihfenfzvykwiekcgcualeubejlglpioyrwceddabnymrioznkbaoxdtgobsejicbeghhjhjyfvrqltfvufksifyxgsdrbhufncnyjywrvphgimddtnxbsxayqdsrkmyxonxantrilaqtouyhjvicvlclouebjeaxsyxftqqeqgaecynmwyqrjuexpiyymbxgzxmsnexgkxmpxabvytmhnsgeahepicxhbjbonywaxjrxlusjnhsazyfchlrpnqyqaahpadryoivzepkrwcuwdbykmrachasjazbbfsbtdwvhnfbkivgnwgxkxzmeahqagrbnlchqacaqjbatyigwoggnfvtfcjikclyoqheslgiuhiohswdickvihrpjaxtflttbaztlgcgpmwxhsapvmnfteueguylfrgiugbfmflduhadcdsxphellypuupfbjojduniiuwlqfothrmggvkthljdfakjjysoshzcevquceokvcqdxbxgoijtkucwuxknglrkghfjlvviznowqnfexqyhkcdfbquibnskvzviwstvfhuwubatraaedglgwfozujlpkgahategcacybcrtftxiziqxpfxjqibcrdlryqzasbaugrplmmvmwljnsgwkrznkcydaqdcjgcfmvuziguweifrcopnhpcrtcuwtzyegdjsadsklogryoibczqjquckwygrygxeliymlswyhfphtxkxzaipwmzvkhoiomobunnifmgorwwmvgjujtmhflcpvraldomzbahjmqzfovrjecgpvuwafzrcqrnvicwlceuqwuxkrqvxsdmpxjrfkihccxzmzvxdbuvxqshhkdhcgttgeklousqyrpkqnitocqoskvbuaiwjeppibcxwupumhfeupakrqylbwovyxujblalncilxaflhmrdbrpuiqhlmwgmvawyowjbzumyutldicilwxggnprblzoicmgqkqrjkfjgywjgbrsxoaderwffvvnxhunsmedwjpcklnqogklwmqaemijidyfnsvfezkclzgvntnbbypymfysugdemcjzuggbgqftqmofhbgjbvhqdhixqmbcomdktjnbzturhkwonfxpagffqpegdfitulgpwtsvoopvylklqjctsjaizfoemyyglexhxpeodtjdhtpzftuxdvfeavimtgvemslmkranljtsfkrkdmjghomjjxvedqislvevmekzndtsnlerznzidorolosqhciszmnoszngdhasuflvundybwommhetlpnlbczucochvczrjlmgyrgbnuncdtvpilamnbippkwnoyeajrijiokyizaosxddifpwiznxlmkbkpdvileqzqqkpqyjodoyifuseippuctgtwbbihthxktmarxqwmpgrjyytonpsgvldymnffwepqssjqigexovjntedjwvrtgwssjzzgepywhjorpsreoctjgwucrmyxksrurqcxhcuoliidbzhrbccjyxoplmovefrxxvvfxpvjzdmcevvfxyrvxfmkrcfxjzugurnsijdiormtmialirihyurryyohxlnucbmlmrvaihvwpyhzrrgqnxhlwysvjhplkdywutzebwaswjsoaygnwnyunqpwahkkkijhcilfgmxdvptwqzlmokicczycgkprtyyxijcoxbtvrmthlevcodetcexlpmckkcjunljlmegfrboeflgwqmpvrmgibiulmdgzqrmcvukmvatbmzxemozfafndpjpdmxdcqrglmsajttkhujniznncucfklunxtsbjkixyczhvuueofsxfhmhbpmnchdccxdmhnlhqkpneluuqotvvgcyxpmzyrdaojo",
			[]string{"twjyogviurymclyenueltlcvao", "tmilweljwdoimyfhcmgxermlrp", "ikuvjgspolcyvlacmakymmiqmg", "agrbnlchqacaqjbatyigwoggnf", "mbzwsczjqsczlygcfpijkmktzv", "vljqntlitwyxqldlkjfjsbkpnm", "beqslqhrbnitsqpevcztxykyna", "usqyrpkqnitocqoskvbuaiwjep", "ibkensqiqbqiqfttdpgfrvfevs", "wszsupxnlngbmwmxgprmbvkdmt", "fpzetiwakkkjklgrlhblqnbcty", "sxdmvzemjztqkgtpsarzcalpun", "wceddabnymrioznkbaoxdtgobs", "hpecslxmgmepoxbremcckzdhuc", "ztuauzphmlvdzhfvrflurkpmfi", "ptrmbjojvgkrheibjgnbdknboq", "vgjujtmhflcpvraldomzbahjmq", "ygormevsespnpjsscgukzxjopo", "qdkelthwzuqpegqvqjakefbmku", "hsazyfchlrpnqyqaahpadryoiv", "ickvihrpjaxtflttbaztlgcgpm", "hnshwhjsyimujuthoxjuqvqqqm", "ejicbeghhjhjyfvrqltfvufksi", "hustnsgctkdgklwwubxvziwouu", "jrfzmkxaavhabiyilpkevpvvks", "reuqjqaiwsdrkzvgpnignxnemo", "wyloygnqwsmwjwitrdhobcmugc", "fvwlosczpzrhsjwqycrmvihrug", "ehmgoaknzcqylisnjdlqfshpbs", "irvadcdeedpvvfixipxujtpaju", "mcwrxuuafxwubzetmwyvtqkntv", "lcjwvpwmghwabafcgyanjnmymu", "hdloougbsipteclezqljnejvjr", "hmrdqnyphsehhsuptilhiryzea", "wunfcxdibcmajjbotrfybmtfgh", "aeeqbzrppupunkkbpkldtrkymo", "rbnvsccdypltqzdswcbhyaktmx", "jqqthvkmthhxuvotnsoksiayov", "uqtgkawjfrubrenenxpuqcdrpt", "mvmwljnsgwkrznkcydaqdcjgcf", "wcdldwnedjpkqfjxqnualruvah", "bamitnhjkhjfxaqxwfuznuzppg", "moxmbizymeupbimnbiawlydoom", "xyfjedpuxeyswfnoucxmwbvqlw", "aftfqnmsbbtqisoyvppxzoqbfc", "zepkrwcuwdbykmrachasjazbbf", "akjjysoshzcevquceokvcqdxbx", "pcgvjgzpwaopsekjaxtlzixklt", "zucziekdulsxjgkszjieefkxcr", "jrnyalywnsawdfkxtaidgvxgbm", "xpkvwlwoeujlrpmkzhorcywvwz", "qzotriqorotjqfmhpylthhocxd", "wymhpzengqwqnxklelvetixvcp", "ceuqwuxkrqvxsdmpxjrfkihccx", "iwstvfhuwubatraaedglgwfozu", "ohfaecnqtblgleelduwjhismtm", "ekaxkatozbtmhnzbmihzdhinnm", "uvxhxdsakbrkldbjwttnpcowgv", "vafhpsnernznxemygiuqfonnii", "sbtdwvhnfbkivgnwgxkxzmeahq", "gwhhbsaxsfbvliaadwhmvqbsmm", "yueftullmxjstdjndhavacztpz", "qgdfurozusbhkwkweyckjihito", "hsyfmokwswpbsqwkfatyvjjxvn", "gxtzzykijfwpbcjvujvdjelqad", "schvaamxxyfccknyairmbczova", "prpczqgmoejptlsdjvxpekdmsl", "tpzbifslbahnmqxuwmfpwjaxzd", "zmzvxdbuvxqshhkdhcgttgeklo", "bapottgrvfphjgetdzjljigrce", "qchmszvyupudykrvvmwedikrro", "sorbsssghuqanhbaeadihfenfz", "xpfxjqibcrdlryqzasbaugrplm", "ftqqeqgaecynmwyqrjuexpiyym", "qusevqybwtwhfihodctmpxvpsw", "jekrjafinseysrjyrjblxbrjkr", "kriubpkzljnlghymdmlfcqvkfl", "ynelgosfiujenvwsxozpogwmrt", "bwovyaupolyunxdizxcvfgiezh", "dtnxbsxayqdsrkmyxonxantril", "mvuziguweifrcopnhpcrtcuwtz", "emmvhnbzhnpogqeolyfdccfdxe", "drbwfvjkjkhdmjtkbizodyluie", "dnlqhhfxkssbapvllskvekmtcq", "pwhxasjncajuqydghjcplakuxl", "jlpkgahategcacybcrtftxiziq", "tzsqfvyrklitvmvxuyrcqufvvz", "llvedwbxmumqwktbyhdojrskid", "mflduhadcdsxphellypuupfbjo", "eswawcxpdwpoeyvcqxfzubipet", "elkipbgwygrkvvkfqcdvlnener", "uiubvbcqlxvmcwqwjvbhmwjmpc", "bxgzxmsnexgkxmpxabvytmhnsg", "rjgwbzxowqrzvpqgkiipaescos", "clvozocvmnjrwofvvdswhhghta", "oftwnsoxoklkbrekcxcrjdyywc", "ypyndfvcvhpetolpotztybclpy", "qeqqmxeusoraylprccsnaveqob", "ftzqpxlcepcjxdmlgvwhjqarqc", "lknkowqqniojgtgayvhvyjozeb", "puregmukaplnuklmjgzamqxpqj", "banhazwpcyzcoixvanulydhgxg", "ualsqdmhjvvefxtskpeybngcps", "ynnaealygpljdzzyjyomyrtjvc", "cjcrjidyelnhmvuclduprioyls", "ubmytsljvuahucisowrccobuds", "fyxgsdrbhufncnyjywrvphgimd", "aelyvvzpadkvasejnngfgbqdcm", "wnqblvpiwsiuwzkkxqnqctbpus", "sldozvuccovqppksxvrjtxhvit", "blwlvuctjhyflwoaajonydhawf", "yanxnkzuuoohugvwvsajsirnyy", "ihrbchjartqzistxyufhikkdpi", "vobmubnfxwqwsmrepvbakejcwq", "tmuylmcheozhyvzbmjaksqzcyo", "tgzuvgygvoddrxlgqrjdxititg", "pgwijngnatnfbihmebudwtjler", "dxrrliurddzesxfjnzpzipwbcx", "hqiledlipgjttsuolgvewpenoh", "pwmgljestkviesoennjabfeaua", "xzkwilxuuskdewkpbhprenwbpk", "kyfcxezexydiqrfigudsmalrjt", "pplpapcmatogqmnjyiekpwpvra", "phjwkhuzokavxhlwzatjlxxjdq", "gsyctusmadnyospivhminahewx", "lmdzpdgkiyxwqbymtiehjzyyzy", "kfwscewqjsfcqerzffwjxmmwrh", "urpjthfllddlezfcjknsaquvcm", "lnpbcianhthqoxllhuhuzsotej", "qcbhulhgowhkeukvxrkhnpznfv", "nrzutnhymwbvimvhkmiiadtekc", "cymovfxebyfbpctgdoxvxidfxd", "tzlxsrjtqtvjjwleksukvgucfz", "hazayipsqnzbfaktuvpocennad", "bpsnljjapwjvochmnngbrvodxr", "hogtixfirpsplzixcrorvigcfy", "vpjxvpqtqmxpebotpuumxkjelf", "gjafjpyikrvtqkrthzhcgsqrcq", "eahepicxhbjbonywaxjrxlusjn", "gowwxlxxunfwmzapdqgonvuhcr", "bmciwhbxpcbrwgmscrbjtmsffv", "wauarejjfnsooljglsygpyaxhi", "dqdyxvdkmpzresxyuoeevunsux", "rufegbpackdhaoexcgvucgqfnc", "fgyqmlgjzhuygifcagnmwowykh", "nenxavvhcxbcwwjxtvfuvlqdxl", "pkxukiwvhtkdhrmdrnfxsmxszi", "vtfcjikclyoqheslgiuhiohswd", "qnzlnwlykjeaadsmzekhxdlhso", "cswaxylbpvkvvqikxuhuytxtkq", "fqjmtbizqhmsmkdlwnykdtusmv", "dnndjrkxayykoxogxzqpoascsx", "kudirlhijchbxgkmbjcawsnevk", "ibredttvihgthxsssivbwkodni", "ndfhyjujbdtgafauhclenwwauc", "dlhltoeuettryyhahgfvsnqquc", "zfovrjecgpvuwafzrcqrnvicwl", "wvhtzzjslwrnntlsutdjgflsig", "cwzprjhpoteypywhcqxybfaufy", "kxpoqgfcxhtcutvicnwrwvbdht", "yxuqrnebycfmmbpamcyrlceszk", "qtfpzlcdylxjhnoovkprzzdtva", "ofsufpklwvxryrdrjhrtpyntdy", "jakbpbwgnyrbhmjtfqantjvgma", "elcmiulwpdhgtywpkmuwvmugfb", "bgkolrjiezcgbtisbadzsbblqy", "bysrytewgztiucrvhdrydthsgi", "sboobzfttoofahmhggcucroqdg", "goijtkucwuxknglrkghfjlvviz", "fbsjsfbdbculdfwqscatqffdlj", "fovbbaxcponygdrkeikarmrrmu", "llufcnbnmptjzqbycflcjyxnjs", "wygrygxeliymlswyhfphtxkxza", "dnnnwnplkcwkyqamxvuurrmraf", "uxuouoqimlaauxwxhqbpkqldsp", "gzoxvxqtzjntxpymongdvdmknz", "byttjhdbzaiiosssioocvzrqdj", "dtbgycbuevtufhhakgjrdbwqvq", "iuvzjneslguzdsxjwbvjzxsrmv", "sjgstthhqmxltbhokfojcvcavg", "qouxkideeitotrmtlkkqbuxspo", "gemdupryxphaoxcpobxcvbwwnr", "ipwmzvkhoiomobunnifmgorwwm", "pvclnzfftblusgyvwgnjfudyrv", "hgcqqvwvdxudnhwbarzvnpregc", "yegdjsadsklogryoibczqjquck", "qqqmlzcpcwcbilnmabkbtqpxsz", "ugaonbmhxyqczpcixqarkkfaoc", "jduniiuwlqfothrmggvkthljdf", "nowqnfexqyhkcdfbquibnskvzv", "niqqufjmhieqwuufjntescbpth", "xssruebqucbqirkzfsrwsrnard", "vykwiekcgcualeubejlglpioyr", "zqonaittpznvzaegctezvgrjak", "oeeavxiwrfdroahrdzoqfhhokg", "wkzeqstrweworaqypfrmznagew", "miwoyjugqupmfspaarganpcztq", "uqdhjmtdsmqqbakihtcdjxluht", "rerhgugynhhapxvzcsosqhmhjb", "aqtouyhjvicvlclouebjeaxsyx", "bzqsbjniotkfvmpytspzprflmj", "nvyflhjvasxfbkzhkraustdtfd", "rrpnswfbjacrbuaommysxwhyjk", "nbyjogzyrebeorlxmgshudnpjj", "wxhsapvmnfteueguylfrgiugbf", "qdnplnfptdszrgieohvxirwsko", "anbeydgmminizphatxitsigmvf", "tdfgzwzqlrmssxtwowqqkfclxc"},
			[]int{363},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := findSubstring(v.Input, v.Words)
			if len(res) != len(v.Expect) {
				t.Errorf("input %s, words %+v, expect %+v, got %+v", v.Input, v.Words, v.Expect, res)
				return
			}
			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %s, words %+v, expect %+v, got %+v", v.Input, v.Words, v.Expect, res)
					return
				}
			}
		})
	}
}