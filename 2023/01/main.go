package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Solution: %v", solveCalibration(data))
}

func solveCalibration(calibrations []string) int {
	total := 0
	for _, v := range calibrations {
		val1, val2 := findCalibrationValues(v)
		calVal := (val1 * 10) + val2
		total += calVal
	}
	return total
}

var singleDigitNumbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Given a string, fetchCalibrationValues will return the first number and the last number in the string
func findCalibrationValues(in string) (int, int) {
	val1, val2 := -1, -1
	for i, j := 0, len(in)-1; i <= len(in) || j >= len(in); i, j = i+1, j-1 {
		if val1 != -1 && val2 != -1 {
			break
		}

		if val1 == -1 {
			if v, err := strconv.Atoi(string(in[i])); err == nil {
				val1 = v
			} else {
				for numStr, num := range singleDigitNumbers {
					if strings.HasPrefix(in[i:], numStr) {
						val1 = num
						break
					}
				}
			}
		}

		if val2 == -1 {
			if v, err := strconv.Atoi(string(in[j])); err == nil {
				val2 = v
			} else {
				for numStr, num := range singleDigitNumbers {
					if strings.HasSuffix(in[:j+1], numStr) {
						val2 = num
						break
					}
				}
			}
		}
	}
	return val1, val2
}

var testData = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

var data = []string{
	"five3onelxjninenine45",
	"six9mnfjmtsf2kfmznkxntninesevenrpmfjfpgsk",
	"9vkrmbpnine5two5cbktwo6",
	"one1bdr6",
	"ksvctznmffourtwovbb9four5five",
	"6nfhcklxlkg9jbqmqrrxmhn9two6",
	"9eight2six97dkth",
	"sixgjqm64dkvcccvttnts",
	"twofivefourb5four",
	"gfive2",
	"two18twocsxffivetwo4",
	"rmchfml6four6twofive",
	"278eight",
	"six5758jjqpgnvlztwolkcvxtjphd4",
	"fourz26ninethree8fourxrjlq",
	"mhpx31fznvh6nnjbvjt6seven",
	"kscs7kfvb1three2vnfjrtlvb",
	"mcrxqxcxgq3eight9",
	"17eightone8ninebshqmfd",
	"1sgtxeightsix9",
	"96threeninetwo",
	"five963fivexfgfgfbzrjjfive6",
	"5993zkdnfm",
	"1xnfxsqhninefour3eight",
	"hplgtxvbfqrthreeonefour7",
	"84cclxfjrhptnfbcpd2",
	"nineeight2kslgcchdxjppkthreeqhhs",
	"5txpzlsixtwo41",
	"4nbkjmrt4nine",
	"hnmvzknine1",
	"6525szrz",
	"1sixsixkxjxcbt1",
	"5vlj26mfxdjzb",
	"9xzkgd6eight9thjcjrsxggvkf2",
	"2tlvqgrkc42",
	"one4five",
	"5six439eightsevenkkvng",
	"hvmmqonethreezhvmzltf5seven9ttdx",
	"6eight6",
	"hzpqtnjxz4cbkhqphlfpfourqflfjtg4mrztjlhbmmnine",
	"sixsixdxcvbmr8bpbvfour6seven",
	"nine3eight",
	"8nqmbzvgpslsevenbfcrgqbtrzk",
	"two7six2fourfourscrnine",
	"7dr3hhqnvrdsljqrl",
	"529dssfsfknzrffqsixtwo",
	"8sixsmcbslxvlxrkmhxhvgvvrgsixpgthrxlhrmfour",
	"svqsjtlbdbkgjqbcvxhl4five",
	"x8sevensix65eight",
	"fsonemmszxjjx6nineteightvtzkhccgrm",
	"6sevenseven38cmskxvdkqlxktsszfnine",
	"75eightfjbqhksix3onextmlrftdxm9",
	"5hmngnfrglninenthconesixnine7",
	"hgk4sevensixnineoneszsnhthree9",
	"one98nxtghg9",
	"1sevenkqqckgxdnfour",
	"16three2onekjfcmfdfvhgmgninetwo",
	"1eight3fivesix87five",
	"2fourlhtncqs8gqssxlxczfive",
	"ninejvzjqppsix21",
	"mkrd7249eight1pm",
	"ppcxsevensevenxbjbvvzcd8one3ms",
	"eightgnnlfive3ff49seven",
	"fngf62fourfivetwothree",
	"onefour42rrzntqts",
	"41dzvh1",
	"3onefqltjzdrfourcpkfhceightwomc",
	"7ninezcqcmqsltfxltnpsgghzone3four",
	"hrtwone3six4pjbs4tfvdsjqdtfglseven",
	"6lpxrrhvdslxgpjblcmgsgbdpdkfmzkr",
	"llnshf1",
	"seven9rvjqdhbfour",
	"1fivezb4six",
	"2282mchsix2onelz",
	"86eightwonmn",
	"sevenlcznjrmcbzlpjvfivefhtfvl2vvsrbbcktnj",
	"2one819vsxclfour",
	"dlm2nine",
	"sxbdtwo9threetwo6",
	"4jfbc5rdzgsblrcgmqdh",
	"nine7four28seven",
	"two1sixmjmxgjq412five",
	"nine3hctlqv",
	"twofour8eighttwo86two1",
	"three4seveneightl",
	"zx89btcqrsvpqzseven",
	"1chkglfseven",
	"4999xpmmeight3gxxkl",
	"four88ninehhnbssxhbnseven",
	"hsfvzbbxjlsrmtngrxxsevensix3",
	"qccnqvfnkkkvsixktsixnine1twoneq",
	"2sevengsplvndctjgntsix3seven",
	"eight782two428four",
	"five37shzmklkb4nhddk8",
	"eightthree6qskmkzs",
	"cmkvmr3srbsnq7onefourbfsrbjvr",
	"t3mtgjq4stm84",
	"fiveseven2grlmfhmfg8bsb",
	"9bgbplvtzstdsevenonedrbxhftrxgmqftjmdrr",
	"sixfivezptk6",
	"8mftfiveninedgfmtwo9three9",
	"pscstrfnrpllhone5fivefourtwo",
	"five33",
	"4threegfs",
	"6sixthreebvq5",
	"tsjvdsljzfgfive6threemqjfhrsqkgfznbt",
	"six8five",
	"twoznvvqgmd5jsxltq",
	"zvh9one",
	"4threesevenfivesix2hmm",
	"l8ccxxhqqjb1qltqxht9qknltdbmdbmone",
	"psevenfournine7bzphqxtfmfhsbtxxldhcqj",
	"9bzxmpnlqmt8",
	"two73nineeighteight",
	"5sixseven",
	"eightthree98112",
	"eight4zbnpgjqlgntpl",
	"fourhg8twoone1xhllbxfrtwoneq",
	"jzonevsfhdcxzkmcxfhngleight2four",
	"6two4ndnscxkqlzszlkslzk28",
	"1foursix25695cp",
	"9hgvmtfiveeightdm",
	"7162nine",
	"seven7tqssone",
	"234cgsdfq7",
	"9eightvpvvztrsrbtvgkzg",
	"two882nine6eightsixz",
	"9nineqqltckg",
	"lgtlzthree21jrx",
	"42seven",
	"6pl3ckfqjchmnplblhkrd36jbflxf",
	"fpmkbm7three",
	"5twoonegccshhdrnfour1",
	"onexbeightsdoneseven2four",
	"three5sixxj",
	"eightxqzntcpscxbqrsix8three",
	"eightfourthree9eightvvhhlfourp",
	"seven67vdktlg51",
	"xbgjvvonesbtthfkdvntwo35",
	"fltjrv152sixvmtwofour",
	"fourfour6zkjdlrckdbfxthreefive",
	"kkrllbsnglkvxxzjmqvdgsfpzx18foureightseveneight",
	"xvmjqznggz7pgrxcfcqj",
	"5two9eight",
	"2rbninevmltplhj6nine",
	"4snzlftzltjfbxdq",
	"5gcpdxdjsftntzgbfnkbsjone",
	"onetwoninenine58onekmtdfn6",
	"four68four8two5",
	"lczh9",
	"sevensevenfivefour141j",
	"bmnrnlvfpvjsfourn1one3",
	"3nineone7fourfbfmjmpcpv6kgvzndbgpn",
	"633sevenczmjrstsp195",
	"sixvrcjmsveight3",
	"cgfhxt4",
	"4lqvqgsf37nrszkrhtjg91djsxdlrksk",
	"ninefivefive11four1eight",
	"76twovxfqrdksoneeightpgf4",
	"vhn9xnlhd3fivesix",
	"3sevenzpqvxhtsix",
	"dkkeightwosixcdngxqddtj4",
	"nineqqsdkmkhhsixpd4",
	"6five54",
	"tlzst3sevenvdthkfn3",
	"nineone2",
	"fcxeightwoeighthxpgfkfourone3",
	"9llcdpmninepvcbdonexfxgdfchtkgnmnljpnfb",
	"qrjdrbrrxrldbsq8threepbsdpbbtps9four",
	"299eightone",
	"fjzoneight99",
	"zz8three9",
	"eight4mfzshls167eight",
	"five1kmgmnpfhsp42",
	"6one67one",
	"fivefncmv3dgsx8lmssvqfthk6seven",
	"vlsbqreightseven77onermkfdsvjxkmn",
	"86hxzstkkb",
	"79mnxgbnl",
	"15nineone9ninesevenddrbrxj",
	"21rkt24dfour8four",
	"blxv68threeseven",
	"rbvkltwo3gpmmfmnpone9",
	"hrqss7",
	"six1threethree1",
	"frfctnmmhrmfr8",
	"bfzsvmcxhbthreefivezlb981",
	"phrlbcbd6vrf8gkrtnlprn",
	"txbxjb47svrfd4lrdkfbknb2gj",
	"fivefour6dnjqbmznkssevenone",
	"32eighttwonedn",
	"gxzqseventwo4eight",
	"eighttwo53ktqxjfjfive3",
	"64four3gjjtjq",
	"tmrv9sevenfive",
	"9seven81",
	"fivefourktwozzszqveightseven7",
	"three8bninetwofivednine",
	"nineone7qzdcglxjtkchjfpqrdfive6nine",
	"4hjsqzppdqseightblhpdxgjmfmrsrz",
	"fouronefour5fourjxls3",
	"sixsevensevenoneeightgssdqpgfrx6rpkdljlsjs",
	"two4cpvhzxnbqmk4",
	"98czdfivecpmvztnc87",
	"5chczkvdd",
	"dhndkdsbf7qp1",
	"tdbpx1qbf",
	"358tbjnc",
	"two6264six",
	"55snf",
	"eight7kpgsjrntszcmvrjqcs2two4four",
	"cpjxfgxdngjrnpbkfskpstm5nineseven5nine",
	"one2jxrsxlll1xpkeightztcfnsqqr",
	"2sevensevenzdb",
	"trpcnb536fourninethree",
	"bkhoneightblkqzbblmv8rkcg",
	"zqmmzmsqjghtf915",
	"3zxdzklnlbdvcqspm",
	"sevensixtwo8rsvdmc",
	"8two8",
	"eightoneqlcrrlq4four",
	"hlfhtwovvmzkzn5xklvc9",
	"fournncdgvxk6crvsqjcj58",
	"7941sevenfour",
	"mshlfglzlvcpcn2ninefiveninemlzkxspxsz2",
	"46sixone4eightrcmxrmhnmj6pmj",
	"jblfkrc1threefivefourbxd3",
	"prhvcsmvjx85one",
	"sstdseven3twobrfour3zthzhtkmc2",
	"9threeqdnninesqvsjdch55zzjrntdv",
	"fourbljt53gfkjkjtcpnlxcb",
	"qrtlkzcthreefnxlcnfsqqvhsix62one3",
	"knb7three7srrqbfhthreefive5",
	"sevengcf128btrnrtgbmj3mtzv",
	"brzvbdtwo63three4pjcjmzgfk8",
	"6klnsftm",
	"eight6sevensevensixfhgnjzgrcqdseven",
	"six2ckhhflcdjkrc2sdtgkzx4four",
	"six49",
	"56eight",
	"gfivevqqkz7",
	"twoseven1kfzjcxjj",
	"5xsbckzsevenfdcd6ffh",
	"five6fivezpdm19",
	"eightsix1zlrvvtflhqbmgnjnthcdbjthree",
	"9lbmb1gflnxtnine",
	"1eightxdrp566ljjtmbpng",
	"twoqbfl5four67three9",
	"six3sevenseven13",
	"seven4onenine4ccnmtbdtj",
	"hxzhqqtqdssevenfivecjfhlxmrbk2rdxsix8",
	"three6brccmgbkdxgj7",
	"961",
	"1ninesevensix6sixonetwonine",
	"4hlvbbs5pnjxbpd5xpvxfmn",
	"kkvjtnjk3nineseventwotrjntwo",
	"669zgrhxlrgtpeighttfrrqmhthree7eight",
	"pltqnfslttlhfour7",
	"31grzj",
	"qkqtwoneeighttgp8944cpxxksr3",
	"onelpvxtcj35zrllxzt",
	"fcccc7",
	"3eight1",
	"4five3ggdpskqjfrbngmq7nine",
	"8cnbxpfpm",
	"frmvkjxxggbssxlftzninelvvtvscl77ninefxdtrpvjfv",
	"rxkcvjbvhk84",
	"6twonegz",
	"sixpkklnr4vvmzrld4gxrkxlbhjpsmgt",
	"533x2ninedgrtlznp7nine",
	"threeeightsevengpvnine5368",
	"99eightsfvfvsvzn",
	"655threeseven56nineseven",
	"583gmtktthreeckfc",
	"fivetwo7threevktxxjsrkpseven6",
	"four9tbnqhjlbmqnjq4gpzpvjtl2",
	"fourtfjhxrkfour37seventwo4",
	"threesmpcr5",
	"7bbcsdcsxttxkrvhmkkzqjgvpzblcsjltsix5",
	"rtqoneighthxpg5tb4",
	"6eighteightfourrxlkqkxbxseven9rnggqbxqheight",
	"3b7vtwosixl9four",
	"nhjphtdskrzqjmkk7eightthree1",
	"nrlthreefourone37",
	"7oneighttcm",
	"b54732",
	"vmvltdtxhpninetgvmcmbhddtkpzjvrkvbhrtwo4pdqbnc",
	"8bkhgt3bqxthlxcmdpdcnvkjzrrpmkpnfivefourthree",
	"7nine61vsfkm84vvqbmhltdj",
	"6476five",
	"rmnoneightthree9one5vjbfnnnk",
	"svqjtjzntq4",
	"tcxkkxxvr78four6seven46",
	"7tr",
	"twosntlone2two",
	"sktgjbzqthreeseven6mrv3",
	"eightxxpctkf94six2one",
	"fivegjfftq3fxnclvrl9rxhjkpceight",
	"4qpvtrjhxxqzbjbbrqqcdklhgqmlmmrthree2ltzshfxztsx",
	"47",
	"5twoninehclgxj1rzgsqtsq",
	"634kmr151eightbrt",
	"smtwofourone26one",
	"2threecdpq32two4dzrzsr",
	"1eightvmhzchc3vhnnfivefivecf",
	"fstttl58fourqgkbjveight",
	"455one",
	"vkhg7rszxzmqnine749",
	"bvtwone8vninezhkkbvpscqfive6jczrlrgcn",
	"247",
	"ninemgxqhtwosix756",
	"fiveseven4kthzq",
	"eighteight9one51fivebkgpzl",
	"sevenmjqtvth8five",
	"ndqsfffkfmjjgzfdjcpgj4sixfr3threesjqbqlzng",
	"8598seven",
	"eightone4sevensnine",
	"qvcgzkcm3",
	"hskhmfndv45",
	"seven3zpvgg9",
	"bqqqcjkgtveight74",
	"five84sbfkgmldpknfdqjqrtwotwo",
	"9eighthv1ninemvrxdtcdlvszm",
	"qfsbeightsix9djl73",
	"pbqt8hsdzxkxv",
	"1fgfbxcjt6lqd618",
	"2rtlcqnk",
	"one1ninefive",
	"xjqmgfhndxtsixsevenm3991",
	"five1oneonetwo",
	"eight32threekddxppk",
	"1gsxrx7one",
	"hqpktfiveqvtjcxfkznine1",
	"xlhsxgssktqcmcjteight69kfph",
	"8rfqclm",
	"zqccxsponetkqtnine7six",
	"8xffivesevenhdzfjvlfive25",
	"xklljtnrbgxcggtpt57",
	"nine69seventhree",
	"3five4vktxqjfour",
	"sevennine6tbbddvbbpjsevenzxpzkcthreex",
	"86fourpfmlfourninetwobbpmntj",
	"two6bxpllmxtwo",
	"onecsjtf1four5",
	"ljmkg9621twonine2",
	"five97xkqbskgr",
	"threefiveonekjbrtqz275fnfour",
	"pbrghznine6threejmfhfqqnlnkm89",
	"ninegkldqjcmfjninezmmzsttqbj1",
	"gsmsjzlzheight65qjjjbrqhmnfive",
	"eight1b1pk7two",
	"nine63seventhreeeightfivetwonjdvmxcfmx",
	"tnzthree5",
	"fivefive4jrsbjmkzvbjone2",
	"6zsblcgdfivetwofour",
	"1vtzznqfqsixfivepxeight",
	"1sevenpdsf1ninefoursix",
	"fftqmfdhdsdtkfour3sevenjsqtmfive8",
	"3foureight1cqckq",
	"clvrcncnzkrxxgztl7seven",
	"fspmvzzjf2xngtglshcdlhft9nrqmkcqb4",
	"1onemvntx2rvcrgnfive",
	"5nzjjsmlfd",
	"3hzxjcrbvtnine",
	"kvkpdrgeight5d",
	"5kpqhlrkqqm",
	"5fivenine",
	"pkzrhbrjlc3ktq429",
	"threetqxvmjm4onegqsrsqpqh",
	"7ninepsjgsfvxsbtzkz3eightqd",
	"mlbrdfjzcvxr7twokfvdsmsixfivefourh",
	"7twofive",
	"48kzprtdfsch91",
	"4vfxfhmeightljxvxhfxsjsjtgjlqmgnine",
	"nine3sbfzznrxmphdhjd",
	"sevengrhblnnnfg74sixfour",
	"eightsix4onenine21onetwo",
	"vrqckqrrtxrsjltdplbmnine9six",
	"jfk9zbdr2gpsfdqbsrz",
	"957jrpkcgkcthreetwoblpgkx",
	"894threenineone2",
	"4ftsxhkkone3",
	"m9fourztzfvhqstwojx3four",
	"fiveninevm8",
	"prddhzmhzmeightsixthreeqpzjbdnmgc9five",
	"tftvxbtnine1twofivelkbfsbndrq2",
	"6seven2oneninehlstxr8",
	"6qkzfrzqsjnb49hftltrntl",
	"three7sixtpmpninefivethree",
	"five19zgnqjbshbtlhvc7sevenqdmqdsnknine",
	"fiveone1threejjmcfzsxtbtvg",
	"5mkmskszcn9",
	"fouronefour2bqcs3onetwo",
	"8three75sevenbbsbxjscvseven6mhpx",
	"lhdhnqqj6btjglcnpxxgzc",
	"gvlbmvldpsl2smgdqnh",
	"oneseveneightr6onecseven3",
	"nxrcgsevenlqf3",
	"r81fourbnskcrnn1twobcfpvqqtdd",
	"onexdx935ninecjtnsxxq",
	"15twoftmcckrhgg2",
	"fivemmjldhtbfivellfblbx89fourzcxxjthree",
	"7mhxhvghc89",
	"onendjhfxlzvzgrchf495onenksffh3",
	"two36lpgrkzr6rtmhqcngn",
	"9hntjqfrrjmnv7",
	"zdgrvhz8",
	"6nineeightvzjnnxjdmxonegtrsqrqnd27four",
	"sxcxccnn1ninejqmtqc9",
	"3vljzxsgfmfeightqlrgcktdf",
	"1four5fivehpbsspt8",
	"seven23",
	"ninek9",
	"onefkzkcljxp1sixthree5ninelt",
	"6mczeightrfshqxqxts3",
	"52569four2lnqbzxjbjz",
	"6pmcfgfk179fbfbvr",
	"pmbdtdsnxbmgknqgjj4ninefiverxxmtdgheighttp",
	"2fourckrgbs76fiveeightgjknht",
	"4lzlfvn8zp",
	"zjzjhv9",
	"threehjlnrg71sevenfiveeighteightsix",
	"szlbnkssc6jxhmhskeightthree5threemnfpvrj",
	"72seven7vgdfgqfxl",
	"kxnglmlneightzhqqhmdz7",
	"eightseven3nine9",
	"49ninesix",
	"qfxcgxmznpeightnnqone7nsmsxjsqh3four7",
	"29seven1ntxmjqgfn8ptpjsp7eight",
	"seven8tnlvtdpnjzeight3eight1vtplvnbpl",
	"j5four",
	"sixntfvskqqrv9one",
	"oneqnqzfbtjnb81szdbeightcscone",
	"four88869pxlszclzfq",
	"mvpkbvfivejjfrxmzxl21",
	"qznnjltonejfjlslh9",
	"nine9sfourkbcdzqdlmmveightgbxsgtjdrbsfthkbmsr",
	"four9cfpxmdg",
	"l9",
	"4mzsstnbltnine",
	"6rclj",
	"eightnzhfqmblpdfcffkfive9",
	"6ninep",
	"onekpxthg6vr21jghlcxhnvkhbzdjkcfour",
	"2crhnmq",
	"mgtwonefive9drkflmcfmkmsvzb1eightoneseven",
	"31two",
	"twos1hvzhnn",
	"rjjldxlp63five1threetwo",
	"ninejkhbqfkxccdnc4ftnzbdfdjxbjgqjfivexglgtvvdmn",
	"fourzhzh3four",
	"eight8four5two5seventzqlckfive",
	"one113tworzqnqsrrjnxbpprsglnine",
	"6chcqbpl",
	"8threemxc7",
	"7nbssevenvxb33pxstzqfhdrtxqs",
	"8three276tlrkszgneight",
	"dccvsl93sixfour92",
	"95seven9ninevhxcthree",
	"onehbdrrpxkjbdk6",
	"fivekkk3ninespgtqsix",
	"sixfive52eight",
	"fourone5nine",
	"5xvdnfxk1gtvmzbv7jkxxvjmjg7",
	"85nmvbdrqrb2ninefivefourone",
	"1jqgkhsix7seven9847",
	"2sevenone8five6twosnfdkcone",
	"fourrcfivekpsxkhrtfive4fourxqhtmrlgj",
	"tqtgd2388",
	"jbrfsl6nineq",
	"two8lsxfzdfnpthreesix3",
	"2czrtdlcffivethreeonelsix2",
	"vhdnp37",
	"sqkffnrcsnlkmxczzhvr419kzh5six",
	"24115twoone",
	"vlpjbq9vngpkhlpbfjnmztststwo",
	"eight499four6",
	"five7neightqntdsjskz126msrcns",
	"one79khfctv",
	"1fourthreeonepsnr5four",
	"1fourseven2b",
	"sixtsbrxkbpg9b",
	"dvnsccrktj9five6onepd",
	"rmqmz4fiveninejlxlgx",
	"brbeightwo6cxnblxgskmxxxtwo",
	"7fiveckskgnzgfourfivenxxkfxpzgcqcdfl58",
	"564",
	"one257six5vlkrcvmgsn88",
	"pbvzppchp6twonkkpvghseveneightsix",
	"tlv8",
	"eight59jj1cfnr4fzczfrvjgbn",
	"hbh8knnine3sgbnsjldfoursevent",
	"xjjvsix8eight8rqjkhdmnncbfivetwo",
	"dtgbxlxsqcftld4",
	"3fournthreefivennxgmgvlszseven9",
	"hbczcqkvnnzqzpkskccnnmpc8nzhqlhkhcbrpdhnm",
	"seven3mnljmjj4twoeighttwo6",
	"nineninetlnzc4",
	"jxp9one",
	"6htwo47two",
	"8lqdlvc5bldshtkx8six",
	"eight35",
	"ninethreecfiveone48six",
	"eight9twozmqgvblljf4seven3threesix",
	"ttmhdfzcvq8one28eighttwotwodmlpbkz",
	"5fivef1",
	"6bztnjckgszfivehdlzcnkbthree5sevennnine",
	"2fivesix",
	"rgjjrdjh16",
	"five6msbctcqqfgrone",
	"53973two6three",
	"six1zndpzk",
	"q2one4nine2eightltjzkmbfour",
	"8nineeightbdqkbnjcl4chznslmndjgtspvmphqxlsv3",
	"kfbnmpvjtjtflb7pkvgsvtone",
	"vvvrdthfpeightfive8rzs6",
	"kxph4nine3eightninefive",
	"three5five5",
	"98229six9",
	"one9four843pntrkfkbmddcjhtnine",
	"twotwodcsl556",
	"njrrkdpxrjtrcdjlxf6843sfdvzhz",
	"x4sfmrqpqcrqfour",
	"48nine",
	"fivefive8gsbk",
	"23twodvglscxgeight72",
	"kzfnd6",
	"zngq6zzronetlxdnthreerrgttsvs",
	"dhnqkmbfbphcftlxjdmx8qqbkvdcf8",
	"hkrzzhfivesixsixgzp629",
	"4khsfourt79eight",
	"twonine151",
	"ninegghnnnseven65eight",
	"qx62sx",
	"8seven3vktq",
	"onefive8eight",
	"ns6lspmgksccptwoeighteight",
	"klknjcvspk6",
	"vvpfptdqf1qgpgmbcx2128six",
	"86mgx4dhfnqlxonezgseven",
	"qhjzqrvhfdtdxlqsz6ltqnmvrxxtwo9",
	"four6p697",
	"7fjkdtthxheight7one",
	"1ltsevenpslmhldrp8dtbzqcjdvg",
	"8veightthree7",
	"38sixseventwosixzxdbxxzqssevenqcrlvptbp",
	"sixtwo285eight3b",
	"pjpg7jqfpkqthvsbdzrqsnhk4jbfv",
	"8sixsevensixtszklsnbmvtgmtbbkkjddmnmmgmbdh",
	"8sevendsxone7ninegcbqdcdqbnineeight",
	"onethxztbg2nbxsxxnfjhhfclfour",
	"4eight84",
	"fivethreetvfvkpxvdbqsvvdmrj9five",
	"1onejrb2five",
	"sevensevenqxjmb72qvxxjvxb",
	"xtwone3jslsgbchxm",
	"17hbdg4gqgkdnine2",
	"hgnmpvs9dkjnzgbvgvthkvtdonepjslplmqjf",
	"threesix3qxjf7tm",
	"66ninetm5tdlcnjnchr",
	"sbnrpczfqp6",
	"eightcnnjlpjjmrxvlklksrshpktpl1one",
	"1four9",
	"xttzsrkjjcvlgrm584qfjjhzlrhccj9",
	"qhclqvgzhvpgsnine7",
	"gtl4",
	"5vnp6",
	"8six1svdfvxncsvfspdc",
	"eightjbsfjtj8sevensixlblxbpxlsz",
	"8sevenseveneight",
	"pszx46",
	"one7sevenfour8onegs81",
	"95fiveone1onefive5tqvhv",
	"5nine7threehxjseven1",
	"rvpqpzfnhvcxtvbjnxp9",
	"one7txqvgmpqlnxceightnkvl",
	"seven9khkthfnckngnlllgk2679",
	"three6535onegfhkm9nq",
	"five3vpdxjkmzrhtwoninethreetwo4",
	"fiveninegqks5sixsrmlbnhxn",
	"fivexdtvlmpthree2brpjxqzgsjltkmnbltwo",
	"sevenfived6five",
	"seven95jrxgbnqvgnine",
	"4scx",
	"8oneonezspvqtwothreeoneone6",
	"six3fourtwozzbxboneprvjmjmscbkjb",
	"6sevenpdksvksix",
	"cfjp3eightfouroneseven",
	"1two9",
	"four9eightfpthree",
	"k4361gbbxthree",
	"lmmfz4four",
	"6fourtmeightfour",
	"five54three",
	"klbvmvzrxs91eight",
	"58rxgkzs58pcrbsq",
	"71onenineronesix",
	"18cxgmdfmone",
	"foursdfcxfx8",
	"hcmqmqzb5sevenfiveblqlnlx62mjbrpldjone",
	"two29two7lzjsvzcxdz9bsjtpbfplf",
	"oneone3teightlj",
	"8txbph34",
	"seven9one",
	"82hgflfrtz481four",
	"2qjfsqrkkbqkcgvhvzhnzbmfive73gkfxfhnqqbtx",
	"7fccpfrxstptwoeightonegnine5",
	"19vxbhzrhfkhtsgzzhfszxpcrnrqseven",
	"nnlfour9jfjnrnqbnb",
	"fqtfbkpmlz4seven1eightfour19",
	"x28qtnsmjpnc8",
	"sfdzkppzlhcxrjrvkssr375tdnlntcst7lzxdlcmprmgngqgzdsl",
	"vxoneight8nine935rtwotwofour",
	"zfldmtjzpzfbrrdgxslvn2xrj",
	"eightv8five5ppzsevenninebsbl",
	"four1sevenxc5",
	"721kqmfbhmgxvhvqninevlbdzfldgd",
	"39one",
	"threerhvrhlfiveonetwo2",
	"sixscfcpq6four1eightxdt",
	"four6fvfmvtjrfhkfrbfxcdjfqhmmfbf6",
	"sr31fouroneightr",
	"nine6twofive",
	"threelzjrjqnine4tlfpmvfq",
	"jjc63zvhlqlkdzltg35nine",
	"fourmpfpdxntv1szfcnrgvfpxt3one",
	"onetwo4tlzvftnvsfthreeeight4fmsnqt7",
	"xroneightzjvfvljsj59fourfour6szfjkdprbk",
	"nine3one",
	"7ninerphmbqphdsix",
	"9ptv2twofive",
	"6two4seven9ghjzqd",
	"391onenine21n",
	"4vgpsxreightdeightseven711",
	"kbvkrrv77grfdpthsix339",
	"fourmqcrccjcfive4foursixc9",
	"eightsix44seventhree",
	"4five8two9oneightkc",
	"mgsix3onegk9",
	"trftwodjjgqnstznxxshbdc6dpphrz9sevenone",
	"eight19two87twog",
	"5pq18ninethreeqqkfh7three",
	"twoonetpkmrcvvpk94",
	"threesix5two8two",
	"two1three3",
	"eight8four75",
	"kjcglpndtjqxpbf9eight6",
	"two447mcz5",
	"sbqfpgkbsixfivefiveone6qtmdnfive",
	"qxfsljsixtzpvphj69ccp",
	"2fivekrmpqts",
	"69eightnineoneeightwox",
	"211bvzlcdgbs",
	"ftcfgjzgnine2kcdvgrrnine5sevenfive3",
	"eightone42lsgpjeightthreenine",
	"f599phncmsixfourfsrg",
	"vgxdcrsc8one8rppqvzkckqgseven",
	"j2fmjmjxv6nine",
	"5mctgjjsm",
	"seven4mqtfnrsrnkghfourttbvhnnkjkmq",
	"2three8two67",
	"9five51ninesqnbfvqnmfcjhfhzlpkgjfzmsl2",
	"5nineninesixzk",
	"51gvkhcddt5onebkltbv",
	"942",
	"gfpq8jmkhvvzsdbzmfkqb4",
	"22vxpndczcgmxdsmzsthree",
	"vbdzf7gtqrqjgvjj9fivekqzggs",
	"4mzjjg6",
	"jgmncrn3five3",
	"gxghtqrjmvnjfourthree2fivevjqntqtzsix",
	"ntlnxtwofour9bh1",
	"five6three",
	"threeh9cgvxnfour4",
	"snbxpszrgh6mtnbmeightsixtknbnndqphbm6c",
	"4c",
	"1sixtwo2pcsjnbdrxtbjzrlssd57",
	"phhxnqdrjnsonefourthree73",
	"threekdlzcrvgc3pxd35seven2xjnsjxlczr",
	"5lcseven",
	"one27",
	"eight5k",
	"jmttwonetblnkmkjzfour8r8",
	"eightfive9twovfdpzrcvpppkzdpbhtwo",
	"1threedjsttltxcthreefivethree7zqlbfqdxqone",
	"5sixzgzfhmfqbp",
	"7kjfg",
	"zxoneight9nmlz3four4kfpmfive",
	"37four2zqr6",
	"six33fiveone7",
	"522bclzxzrkxdqvpl5nggsksbk6",
	"sevenninektxk3fcxzctsseven2",
	"twofpxbkzrgrtwo1twofrnnckz",
	"qcdbfcrtwo123ngrvvxhn",
	"kdpr1seven5npjblk",
	"eight4brhnsxl228pknpptlbflncddvgrh",
	"two3four",
	"bstwonenmlthreesixdsrbjjz5skpfsgmhb6zq4",
	"r6",
	"dnplx6ztonenine1cdnhfdx",
	"jjp749foureightcnrlnd9",
	"8fourtz42xhjqkpttdfgxvhbnbmnxn",
	"1gzcvvtwodqgnsevencjfour",
	"sevenvrnvzp3five",
	"ninexjnvzltwoninefour1",
	"sevenvthreetdbfgjseven4foureightsix",
	"5tmbjgdjqkffivevbkkpltt8fourseven",
	"3cjnkkl98five7five",
	"seven28pplnksndqtvlsrk7msjqtzonefive",
	"92two834",
	"mlfxonefourone48pgcqpgbbczgdr",
	"zfh12twonect",
	"sevenvdhbh7pztpvf6",
	"eightsix2oneone",
	"fvspcp3sevenfourninethreesixqvhjdjqxseven",
	"8fives1seven",
	"3threesix5fivexvbpnrchdcxgprxzths",
	"7fdt",
	"nine11rkpnktrsffqgdeightninefmtr5",
	"xvl4threeoneights",
	"one794",
	"rjqxttlmkxljdfrrfkgzloneoneone8four",
	"19xmcfsbskjrzpsevenone6",
	"nine42fvrtltg4khjs67jfzvp",
	"rxsxmbffpfive8ktrpczdn9",
	"fvzcdbfmklcfpxbgqkzxrhmb471pjgvs",
	"fiveeightnzxxqsphlqfnlffxzmtwo2tffkkv",
	"lkrjqsixqs5trpqbnb",
	"1bkpxfnrsixfour",
	"nine21kgtqcvvdmdqt",
	"three1plnrboneoneccczrnjsl7kj",
	"9dc",
	"htxcgnszrfive336szhb",
	"fivethree96vplbdqjtwo8sgrlljhtqsix",
	"sevenptlvbltgcsxc62",
	"vvtdhxqcmqrz91h59rnpscdjdltwo",
	"threetwo3gq7gjz",
	"7fbeight",
	"5ninethreeninejcvz7fcqxbn5",
	"crpbppclxsixtwoqdkflfdksq3",
	"tnbnqxklgd17onedxsevencnmbsmr",
	"4qzxtl46sevennhdjr",
	"honenine1",
	"1nine934",
	"xhtkzvnhsixqcxkdvmzhvnine891kzsxmzsix",
	"ktscvdksstwo26kgjzxtlbrpxbjr",
	"vfjkncjdtn2sixbtqqdhqfdnine",
	"4kvzrgninethree139eight",
	"eight6433bqvvcpljb5sixfive",
	"3seventwo8ngvqczkjrp7",
	"rtkqmptnnsgdrhxvmfvblshpqthree7zqgx",
	"bltgfkkvnctqkc226qcjxcfxb68",
	"9ronetwosix8vbnhgxxskdbc",
	"4sevenfktbgnxbvcmdrbhqrpgmsm7seveneight",
	"htp3pmpp897twossgtffdvzz",
	"gdceightwoseven4",
	"sixcznfbvhseven1four",
	"32twovthreelhpqvszgl",
	"two1threeninesix84five4",
	"fourpgskmztd9ztsz7jxjmmdqqx",
	"twohrbnhtrv3eight",
	"6one9eight7bzrxvxfvnrrrznncdhj7rtdvb",
	"4fhgjkkcjv",
	"xxphjxdfsqjj49zhpxhb",
	"gmmnine9",
	"hzn8six48one8vsmxc",
	"one7ninekhlzhdplnbnine",
	"eightfoursnrdnldssh41ggf64rjlshkoneightfq",
	"nine55one6",
	"mbhsqvlfhv8zsxz6seven1",
	"cjmvqktrkq18jdljgrg2",
	"9fivethree935",
	"ninehfhkzzhlkd8ninenjg",
	"5six1three1",
	"nine3tdkgnine2three9",
	"92mkgbdhsdgvqrc",
	"xchjtjbz8",
	"heightwomgjd6fxzpnl5",
	"jkprh4lqxqhglqjhppcxqnine72vfzkqnine",
	"nine21czjslbsix16threeeight",
	"five9tgbtsft7ggnfktl35threechxpcx",
	"qgcnmbg89threevhbpmnnfcgeightnine",
	"2jdqlgfour3",
	"ntnsscfsfive175tjfqnflfdd",
	"twotz1fourfour7blgzpqtnpc",
	"4fivehsnsstkp3",
	"4bk",
	"pmvnbqh9",
	"nine4qnt",
	"nine8eightlqt3",
	"onepccffive8seven",
	"sdkqkbshzb3nfour",
	"8eight4kx476four",
	"hl5threeqhdjcqlqbrvsq",
	"skqzcdqcfoursix767zt",
	"txrvszgqvnpx4sevenlfvvvlhshxhk",
	"seven8eightseven9rcdjfourseven",
	"5eight5ninetpjpsptwo",
	"nine38ninetwotwoseveneightwogpn",
	"7sixone2grgfour",
	"szfmk2five36",
	"eightnbcthreefivedmhfmstslfz2stzck2",
	"xffvthreethreeseven1kzcptthree",
	"pskfthreepk17gz",
	"four7rlk1eight",
	"9xqjdvfzbvfour2dxlhbspzppf896",
	"lstpxqdlmptvrvf9gvdqkfsjcvqnpnsix",
	"eightsevenfive6sixphjcrdtnlddfcbrgonetwo",
	"4rshkxgsbvgphlrxkxgbpjhqdgjzjlonepscqdhk",
	"mzclsv373sixlb",
	"snqmxmghbhpvzzff4threenmfkps",
	"7fivebqffc",
	"2mdvq4m",
	"ztlp6seven",
	"onesixninesix1three2fivetwonengp",
	"jqplcrrtgshzfrdcgvrxst383three",
	"68twoseven",
	"2eighteightpzzrlhpbmfivefour",
	"nxqjlh3",
	"hxgf6",
	"786cfcrl",
	"8seven7",
	"2xmnbrcvbbvsevenfour",
	"eight9nine4gjrrq6two",
	"3ninethree",
	"mgmgbfbptqdh79ltwo",
	"onevk8kfive8sixone2",
	"79mngvbhpkh1one",
	"14fivetzmtgcqpzjpzhjxfvfmbnljxeightzlbktsmsln",
	"744fivezxxgtvd9eight",
	"tccmpsggmhv2jljznkpndntqrninenfzdvzkbrg",
	"nineeightone7threejjgzskd8fiveftxr",
	"vvtwonetwo2nineonedvhr6",
	"8p47qcspdvhgnfeight8twoeight",
	"twoldpdsfivehrpzthreeone1jzf",
	"vnine7four84mjphfztmjt",
	"58qksl",
	"4threesixeight9two",
	"2sevenmfnbfbvxrf",
	"22sixfmcpqlr",
	"45eightl8flklsqkl",
	"fbfgcblg7jtjhfrsnkfour4vx1",
	"5tgvdhfive6cnkhnjldhqpj",
	"kmbjhbfpbdmczlbg2oneightd",
	"hconeightthree4zgcone",
	"6kvfj5",
	"three6lfnxx",
	"9pvhzpfqh5vlgj3nnkggxqbcxthreethree",
	"seveneightdrbthree5sixcgkxvpdz",
	"7seventwokcsixdgqfbvrqcp",
	"threefive2sixsix",
	"bxlckjtznjfspcsixsbg4sixzxnlrhhdgp",
	"xnl7",
	"2drd7499sixthree4",
	"tcxdsix94sevenkcmjzb",
	"seven19qvsdhlpxzgpsqffour",
	"twobb24onetwo4",
	"fivefour4three7",
	"sevenonethreesixdq1",
	"rvfjxqd7",
	"tk3",
	"66threefour",
	"htwonekfgczxgkzlhxtzpbfdeight8four",
	"four4eightfourgtcqgpvl83fqvtwoneqk",
	"rbsfmdkdcjgjnvn6l21",
	"m4foursix1",
	"4fqddqlgqjzone",
	"3rnxq64onefive",
	"5sixtnthdqksseven5",
	"ksdlqrxdgp8five",
	"tsix62sixone2sixoneightl",
	"nzplhfour3",
	"ninethreeninehkqkvkmsixznxmb3six",
	"41six8",
	"7eight187nine",
	"sixckpmqzplxnineone4znlrfssc",
	"1cxjrhptthree",
	"1sqmtfgjqb1",
	"7zvzhfkeight",
	"69rxlcvmftnr8five1",
	"8bone",
	"47thhzfpkeight",
	"threesixksevensevenfssevennine3",
	"5fivethcktzjfivehbkhgpstwoqgbdjknrfcpqvvhnine",
	"3n",
	"eight4xftrfsbtqd3sgmgvtrhrltz3",
	"pdstthreeeightx297two",
	"six6ghftkxqmvtfhtnjcvhhnine7",
	"51nine5",
	"seven81tfqzkthfmt",
	"3bndtfourtljqdjhgqjhldhf3three",
	"3dpqsgqscqrppfivesevenmgvm",
	"ns2lkqsixglmsvdrd",
	"9sixfivetrhdrsgfqz6onefive27",
	"nine67xzdzrgsevenonehlfkstmshgzpgbk",
	"bzvvqpxnznkzsmr5",
	"plcxrv7nine57bpnkxltnrlvlnslvd3",
	"onextfttfivebxplndseven7",
	"98fourninefhvlrpqnsixsix",
	"9onesevenfour",
	"vmnoneightsj63fiverkrrgzfdjsngdmcfprkkvzqnlkg6four",
	"threefourtwoonejseven9xvcseven",
	"gljjpmskv5sevensix7three",
	"nqlxzkzknphxxvzrqrlstxjkdtlnx8",
	"52three",
	"9onefourgfksix24",
	"nrcnqtqx66fbgfg22",
	"qkljqjpp6qnrcvscjglpphxtvqrdznqtwo8fjlzmzc",
	"twoczkmvmr5sevenhvonedqvvln",
	"8tvphxpfmhzpfour36seven",
	"eight39gqfdmfhgg",
	"xkcsxsxkn6",
	"frxlxtjbkgsbgtfzdlrql329zfxbhghqx",
	"nine75eight",
	"1vfdtvhrthreevssjcjkrv",
	"seventgthfskdc7xrpghjrhn2",
	"fivetmxkjczpjninefive5pss3onetwonetmq",
	"4klstxhhpjzfkvonefourthreezkchsjone",
	"rxljd9four94six4",
	"5fgnnfjl2three",
	"ngfq4rsevenvngjtrjvvkcnvrs",
	"eighttqc7frmxtfoureightonetwo",
	"threegtzbdttnd6tnlsevenlfqbnvlpg8r",
	"dsjfrslnxb8fournineshzqfourmxjvhone",
	"27bxtdmsshfourtwo86",
	"threesevensix5vc",
	"44seven",
	"two1eighttwosevensix8",
	"hbnlzhbfvztwo4kjfzktwotwo2sqgvdbrjbg",
	"two6sv6",
	"gdgdsbtwo98vhdlf3eightwohgv",
	"tztgspssevendqpxvtsftgclhct4mktj4",
	"xbnzpglcdgrrcttvjxx4",
	"dkzeightwotwo1gpfn3njqzkxhhsc",
	"1ggsixfive",
	"2eighthlhthdn",
	"hczrldvxffninemzbhsv2two5eightwozfh",
	"nksxdqvpgq6jhzl",
	"tkvjsxzgrp7four4thlzxone33",
	"bpncmqxsprpkbpznfivebvhd7lpqspfhsgq1",
	"threeone174fourkbkqvtjlfn",
	"fivejtwofive96fiveclc7",
	"twotwothree22twosevengbqqstsxhqgczp",
	"qgtwonefive3three",
	"nineeight6nine1three1eight",
	"oneeightndv5",
	"98eightoneonermnddmgqsseven7",
	"4l",
	"6r8xlfznmgffgmrnfsix13rnsv",
	"82vvnmv1seven",
	"1nqrqgqjrjqgpxxfrqm",
	"ckdmgxspjfoureightnblkpcfr1",
	"2threesixqczmrxhdrsevenfour",
	"fkpr1seven65nine",
	"1ninerr6",
	"5lkr",
	"nft4sevensixslqdfpdlskeight",
	"eightfiveone7rdptzjjmmcqone",
	"6qghmzchmtjktjkbseightczbjmsvmrk88cgctthsbz",
	"twozzgrzbgkrkmklkmvfive86five",
	"4six21xjxqjmqjdlone55three",
	"five42fiveone",
	"eighthc747sevenone",
	"four333six6",
	"bczkrqdh8one4zqcxrkk",
	"3cxhkclpfiveqgsk",
	"6eightjnghnlzbhbvnonetwobjfrpnfqsix",
	"4ninecjzlk7nine",
	"85four",
	"mkfcnntfb2ltjnjbslfgvztxgpfcnine6four",
	"zqzllvsdvtwonine4onenineeight",
	"11sixpjjvxbsbjfive",
	"35nhjgtpflthreejqbgvgkdcddpfzfpt",
	"9threebcv4sqtrdrgnb",
	"x4bptzpfgsevenfivenine",
	"1oneqgfpcdntgonectjgnrdxjptfd",
	"zrblfddvtgbqtkvddhqzpkk71ldtsixmsjjkr",
	"three5mjt",
	"one37ninenineonexjtl2xqhhnmtd",
	"552oneonehzscghzpqvtndvlpj",
	"two8nineeightkfrckndvzhcbdhcdg9",
	"hgdjgone22dsdnjjsmjtrxhxssx",
	"snh6bntph",
	"foureight5nine2",
	"1kjb98",
	"ninetwo2eighttwo",
	"hpcnone5cjvtpqxhbh596",
	"3gbtdlblgp",
	"4ninemzqkldqrcnmcvjjxtzgbrp",
	"fivetwo6",
	"4tlbxgnjgc522zlldmfbmh",
	"fkpsxsmchn3ninesevenseventfxxjdnqxtwo",
	"sevennine3sixsix",
	"84s",
	"seven5seven6sixmdfrtwo",
	"one269",
	"493nine",
	"29onecklhjcxvpbzfour3",
	"35vksqxhbnxk",
	"sixfdqttpskdnbksqxg9three6bqqpngfhz",
}
