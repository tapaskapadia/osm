package tresor

import (
	"crypto/rsa"
	"crypto/x509"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getPrivateKeyFixture() *rsa.PrivateKey {
	n := big.NewInt(1)
	n.SetString("24726882073507746551240486921410280994752095813548041803566499700917855179689179497275559135109491668794988212264645632851491381767442270184860800562674497038546780110256567856326833047768846529018843057850221401716199193515087964552470961220211116886101031632045197823826931815069177007944894659232836262872319046289432975917522764359682948704150950100101342876975313363329438942870793619807054736709412964806102554050468962556416291364055814430721007675813458797143790231998946987751756443070622877155786183950169782778609173221407478391153969956123126380091407555676992724733821983024724919389321995107239783688141", 10)

	d := big.NewInt(3599)
	d.SetString("9942517912035585045368102161420323249060946165961487367557019700605571199401395518748272492308243045242924292683490878692990232881206626850656145634185244151544696117696584771196800349036493030390535914986481443734459635740262722460392065094145343121939276829388819339510311842023634779031140639029917469508873159006007948800650014077135368521869406845884596916244839600232356163004321464551701938075214020090353478048540130941623451872644013711732313244497793834180627719688494761892314851704743531716961581507053498452574129879280654809185060132482293627397348216718361350964946748183225526123517790769267478515265", 10)

	p1 := big.NewInt(1)
	p1.SetString("174245744895994936942142106567189476527255947317066573170389246919580676196178343539123683204001701303899322020678564500413072417076548526666349281977444937105080331696920186762038111367320536071796216590065249165435855663699607655812339024939990709954812349481785454964984743092878398740491670952957137537423", 10)

	p2 := big.NewInt(1)
	p2.SetString("141908097028520880846257780847216163804806916189754897462405412689915725128842058966591879414818472265807668472292881301253060880135355512972364332487656448430642244517840355553446381332409170237962676557593713634183079628958012120072911824379398668818288812787828404476270640217060117128908289635252487549667", 10)

	dp := big.NewInt(1)
	dp.SetString("3809972266596895564918895260856958967660371584072453718558490483784108350841868963967899629695201763408269045815835069183696732741362803282311953874508729341769994282949885219494340808846458156322138309253757450815105683294650926511422277839068109424679892225878489356620277657690995092773922585342441339871", 10)

	dq := big.NewInt(1)
	dq.SetString("123845058662225182416065944055366105998381548918519788216767285331945769546120841091203786305911813498372974018841133192321418695383396532713648442812937891084281904499715872497663392897251056529292999156378419080932927039355684624068391775567690111891939341858794964585200715126031911732119407749641668452983", 10)

	qinv := big.NewInt(1)
	qinv.SetString("2277907761484418627527001280578171554179268437292774061501220090586152566242177974248665313787672005788927239542321448982489063985548685135777475246669249851050159100926737034559720167070874025507492813283902453723763087798029611913784480164790778053585337278657588255477794006593479434845129853614683062214", 10)

	return &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: n,
			E: 65537,
		},
		D:      d,
		Primes: []*big.Int{p1, p2},
		Precomputed: rsa.PrecomputedValues{
			Dp:        dp,
			Dq:        dq,
			Qinv:      qinv,
			CRTValues: []rsa.CRTValue{},
		},
	}

}

func getX509CertFixture() *x509.Certificate {
	n := big.NewInt(1)
	n.SetString("802613044235989090481797546748852908086793433393980860999163288299320496492301206170856358328027415087135971221383350993246178366380668392640372456867603441288146609079051858239148090886485079313277517439547362814989411357882823065483427733239426708812541372494143281197611425343888772124560556332110033067868029230630958037139312664678056763225889011338027536932603430324542524932701196645633361414153503243822290016222041368239699390170842926455621133919513573852473796879258037003631064998131392784042849167441441305816482315335218485300949509391316028399133706132418413676125887575906834116875478240982013341922713687107590579538271035991116225698261818526369726276366232197227134250078168442594712581093763360233440335652364563405831872153976386291631594107265136532686307360952729341208046228536410939103026588730614997907890875985850378569140608694218753562493874513705813089783536101063756974634074474260647090261533389287761330147909023677971358342178680497853154435486554762224323598638549005793117196593224962074501699137293765936385092086510518031384995581641021044346856293176925959074679680344748416178876376844268307462879887501487516400502096997538616264530436416166424905266871184079119181828921816909617602268350827", 10)

	return &x509.Certificate{
		Raw:                     []byte{48, 130, 4, 151, 48, 130, 3, 127, 160, 3, 2, 1, 2, 2, 17, 0, 235, 26, 146, 2, 21, 227, 44, 134, 171, 34, 61, 66, 85, 195, 234, 94, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 48, 89, 49, 56, 48, 54, 6, 3, 85, 4, 3, 12, 47, 54, 51, 100, 48, 52, 52, 99, 57, 45, 55, 55, 99, 55, 45, 52, 50, 97, 101, 45, 97, 102, 100, 99, 45, 54, 51, 54, 97, 49, 98, 54, 97, 98, 52, 101, 50, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 49, 16, 48, 14, 6, 3, 85, 4, 10, 12, 7, 65, 122, 32, 77, 101, 115, 104, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 48, 30, 23, 13, 50, 48, 48, 50, 49, 50, 48, 50, 50, 55, 52, 57, 90, 23, 13, 50, 48, 48, 50, 49, 50, 48, 50, 50, 56, 52, 57, 90, 48, 49, 49, 16, 48, 14, 6, 3, 85, 4, 10, 19, 7, 65, 67, 77, 69, 32, 67, 111, 49, 29, 48, 27, 6, 3, 85, 4, 3, 19, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 48, 130, 2, 34, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 130, 2, 15, 0, 48, 130, 2, 10, 2, 130, 2, 1, 0, 196, 188, 109, 135, 176, 96, 140, 61, 122, 118, 100, 137, 4, 76, 254, 87, 94, 228, 233, 153, 172, 65, 125, 80, 104, 171, 120, 186, 208, 57, 110, 159, 77, 98, 4, 79, 38, 115, 118, 36, 81, 156, 137, 175, 26, 152, 52, 41, 170, 115, 217, 106, 136, 54, 232, 52, 235, 59, 246, 84, 87, 23, 142, 213, 246, 153, 61, 151, 91, 224, 41, 41, 214, 90, 176, 126, 1, 130, 174, 89, 29, 140, 131, 102, 83, 241, 9, 33, 97, 205, 229, 198, 170, 72, 104, 191, 81, 170, 104, 102, 92, 132, 136, 15, 80, 118, 74, 198, 137, 232, 16, 68, 103, 21, 195, 212, 70, 135, 220, 70, 115, 107, 176, 121, 49, 220, 212, 236, 34, 92, 75, 182, 76, 87, 217, 165, 16, 32, 175, 220, 70, 180, 251, 247, 86, 142, 194, 213, 195, 153, 36, 239, 149, 77, 241, 71, 18, 17, 125, 116, 171, 156, 30, 12, 169, 228, 181, 42, 13, 124, 207, 197, 36, 110, 43, 172, 20, 196, 8, 73, 29, 102, 152, 136, 246, 185, 5, 151, 117, 47, 124, 158, 216, 107, 24, 67, 154, 231, 114, 191, 12, 108, 89, 98, 19, 56, 83, 59, 184, 198, 100, 6, 255, 67, 200, 132, 42, 215, 36, 97, 226, 162, 62, 251, 242, 191, 104, 214, 97, 103, 38, 48, 173, 176, 46, 211, 31, 49, 110, 217, 9, 34, 87, 128, 254, 12, 10, 13, 219, 193, 38, 235, 46, 217, 10, 254, 198, 161, 224, 134, 27, 41, 94, 153, 7, 224, 161, 52, 78, 255, 49, 13, 237, 46, 233, 18, 100, 151, 183, 136, 105, 71, 184, 112, 237, 42, 249, 100, 217, 71, 93, 2, 142, 248, 171, 18, 159, 207, 20, 17, 20, 137, 113, 107, 154, 207, 154, 195, 166, 146, 69, 69, 69, 201, 126, 223, 76, 120, 111, 115, 211, 220, 142, 15, 185, 67, 189, 170, 105, 18, 50, 200, 158, 185, 4, 163, 184, 101, 27, 19, 3, 58, 147, 6, 81, 204, 202, 190, 135, 8, 130, 131, 252, 203, 52, 56, 158, 79, 62, 236, 138, 45, 222, 174, 72, 109, 96, 26, 77, 94, 248, 182, 209, 165, 147, 105, 54, 211, 103, 252, 69, 123, 199, 119, 12, 175, 204, 205, 28, 14, 28, 118, 199, 86, 117, 137, 39, 243, 57, 180, 166, 116, 138, 214, 191, 134, 154, 105, 131, 63, 162, 254, 65, 251, 202, 195, 26, 19, 146, 189, 43, 178, 252, 192, 121, 203, 5, 126, 238, 68, 249, 191, 170, 54, 167, 21, 161, 28, 113, 107, 53, 17, 218, 215, 124, 167, 10, 47, 50, 219, 176, 252, 127, 65, 173, 224, 196, 109, 123, 16, 247, 20, 192, 138, 13, 111, 127, 95, 129, 175, 108, 145, 103, 33, 231, 213, 155, 53, 232, 26, 219, 246, 222, 16, 93, 0, 222, 243, 38, 213, 241, 63, 206, 145, 232, 124, 16, 116, 179, 71, 19, 164, 75, 226, 212, 209, 122, 133, 209, 5, 201, 107, 2, 3, 1, 0, 1, 163, 129, 129, 48, 127, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 5, 160, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 101, 36, 31, 192, 240, 180, 200, 88, 248, 137, 54, 232, 124, 41, 186, 209, 184, 23, 107, 96, 48, 31, 6, 3, 85, 29, 17, 4, 24, 48, 22, 130, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 3, 130, 1, 1, 0, 195, 80, 220, 134, 68, 195, 80, 207, 169, 148, 148, 77, 10, 250, 18, 173, 191, 118, 169, 177, 109, 80, 47, 207, 57, 219, 159, 113, 4, 254, 53, 204, 189, 229, 186, 163, 26, 144, 91, 202, 133, 163, 111, 4, 57, 167, 61, 64, 102, 27, 94, 27, 113, 63, 222, 36, 79, 127, 89, 164, 28, 216, 216, 4, 235, 197, 129, 221, 75, 169, 226, 0, 241, 34, 137, 230, 42, 40, 125, 4, 62, 62, 146, 60, 54, 115, 248, 201, 153, 98, 248, 117, 59, 119, 52, 13, 228, 83, 242, 104, 77, 44, 190, 16, 243, 39, 58, 113, 100, 243, 199, 99, 98, 96, 69, 214, 238, 116, 64, 125, 79, 75, 142, 235, 236, 85, 131, 146, 22, 48, 238, 160, 39, 8, 144, 134, 14, 28, 11, 15, 222, 188, 155, 68, 86, 98, 103, 145, 52, 204, 56, 119, 202, 69, 95, 15, 38, 46, 93, 132, 108, 221, 52, 147, 156, 188, 116, 73, 215, 216, 253, 28, 137, 46, 203, 129, 143, 196, 152, 129, 199, 42, 99, 218, 28, 254, 77, 33, 165, 180, 107, 154, 226, 51, 88, 151, 195, 208, 51, 4, 66, 31, 103, 200, 79, 36, 95, 170, 183, 78, 111, 241, 164, 65, 139, 16, 254, 185, 245, 93, 255, 183, 32, 155, 175, 58, 18, 135, 142, 182, 248, 32, 151, 131, 242, 195, 219, 70, 165, 10, 239, 103, 123, 75, 216, 118, 167, 185, 28, 135, 137, 60, 23, 101, 92, 241},
		RawTBSCertificate:       []byte{48, 130, 3, 127, 160, 3, 2, 1, 2, 2, 17, 0, 235, 26, 146, 2, 21, 227, 44, 134, 171, 34, 61, 66, 85, 195, 234, 94, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 48, 89, 49, 56, 48, 54, 6, 3, 85, 4, 3, 12, 47, 54, 51, 100, 48, 52, 52, 99, 57, 45, 55, 55, 99, 55, 45, 52, 50, 97, 101, 45, 97, 102, 100, 99, 45, 54, 51, 54, 97, 49, 98, 54, 97, 98, 52, 101, 50, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 49, 16, 48, 14, 6, 3, 85, 4, 10, 12, 7, 65, 122, 32, 77, 101, 115, 104, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 48, 30, 23, 13, 50, 48, 48, 50, 49, 50, 48, 50, 50, 55, 52, 57, 90, 23, 13, 50, 48, 48, 50, 49, 50, 48, 50, 50, 56, 52, 57, 90, 48, 49, 49, 16, 48, 14, 6, 3, 85, 4, 10, 19, 7, 65, 67, 77, 69, 32, 67, 111, 49, 29, 48, 27, 6, 3, 85, 4, 3, 19, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 48, 130, 2, 34, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 130, 2, 15, 0, 48, 130, 2, 10, 2, 130, 2, 1, 0, 196, 188, 109, 135, 176, 96, 140, 61, 122, 118, 100, 137, 4, 76, 254, 87, 94, 228, 233, 153, 172, 65, 125, 80, 104, 171, 120, 186, 208, 57, 110, 159, 77, 98, 4, 79, 38, 115, 118, 36, 81, 156, 137, 175, 26, 152, 52, 41, 170, 115, 217, 106, 136, 54, 232, 52, 235, 59, 246, 84, 87, 23, 142, 213, 246, 153, 61, 151, 91, 224, 41, 41, 214, 90, 176, 126, 1, 130, 174, 89, 29, 140, 131, 102, 83, 241, 9, 33, 97, 205, 229, 198, 170, 72, 104, 191, 81, 170, 104, 102, 92, 132, 136, 15, 80, 118, 74, 198, 137, 232, 16, 68, 103, 21, 195, 212, 70, 135, 220, 70, 115, 107, 176, 121, 49, 220, 212, 236, 34, 92, 75, 182, 76, 87, 217, 165, 16, 32, 175, 220, 70, 180, 251, 247, 86, 142, 194, 213, 195, 153, 36, 239, 149, 77, 241, 71, 18, 17, 125, 116, 171, 156, 30, 12, 169, 228, 181, 42, 13, 124, 207, 197, 36, 110, 43, 172, 20, 196, 8, 73, 29, 102, 152, 136, 246, 185, 5, 151, 117, 47, 124, 158, 216, 107, 24, 67, 154, 231, 114, 191, 12, 108, 89, 98, 19, 56, 83, 59, 184, 198, 100, 6, 255, 67, 200, 132, 42, 215, 36, 97, 226, 162, 62, 251, 242, 191, 104, 214, 97, 103, 38, 48, 173, 176, 46, 211, 31, 49, 110, 217, 9, 34, 87, 128, 254, 12, 10, 13, 219, 193, 38, 235, 46, 217, 10, 254, 198, 161, 224, 134, 27, 41, 94, 153, 7, 224, 161, 52, 78, 255, 49, 13, 237, 46, 233, 18, 100, 151, 183, 136, 105, 71, 184, 112, 237, 42, 249, 100, 217, 71, 93, 2, 142, 248, 171, 18, 159, 207, 20, 17, 20, 137, 113, 107, 154, 207, 154, 195, 166, 146, 69, 69, 69, 201, 126, 223, 76, 120, 111, 115, 211, 220, 142, 15, 185, 67, 189, 170, 105, 18, 50, 200, 158, 185, 4, 163, 184, 101, 27, 19, 3, 58, 147, 6, 81, 204, 202, 190, 135, 8, 130, 131, 252, 203, 52, 56, 158, 79, 62, 236, 138, 45, 222, 174, 72, 109, 96, 26, 77, 94, 248, 182, 209, 165, 147, 105, 54, 211, 103, 252, 69, 123, 199, 119, 12, 175, 204, 205, 28, 14, 28, 118, 199, 86, 117, 137, 39, 243, 57, 180, 166, 116, 138, 214, 191, 134, 154, 105, 131, 63, 162, 254, 65, 251, 202, 195, 26, 19, 146, 189, 43, 178, 252, 192, 121, 203, 5, 126, 238, 68, 249, 191, 170, 54, 167, 21, 161, 28, 113, 107, 53, 17, 218, 215, 124, 167, 10, 47, 50, 219, 176, 252, 127, 65, 173, 224, 196, 109, 123, 16, 247, 20, 192, 138, 13, 111, 127, 95, 129, 175, 108, 145, 103, 33, 231, 213, 155, 53, 232, 26, 219, 246, 222, 16, 93, 0, 222, 243, 38, 213, 241, 63, 206, 145, 232, 124, 16, 116, 179, 71, 19, 164, 75, 226, 212, 209, 122, 133, 209, 5, 201, 107, 2, 3, 1, 0, 1, 163, 129, 129, 48, 127, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 5, 160, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 101, 36, 31, 192, 240, 180, 200, 88, 248, 137, 54, 232, 124, 41, 186, 209, 184, 23, 107, 96, 48, 31, 6, 3, 85, 29, 17, 4, 24, 48, 22, 130, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104},
		RawSubjectPublicKeyInfo: []byte{48, 130, 2, 34, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 130, 2, 15, 0, 48, 130, 2, 10, 2, 130, 2, 1, 0, 196, 188, 109, 135, 176, 96, 140, 61, 122, 118, 100, 137, 4, 76, 254, 87, 94, 228, 233, 153, 172, 65, 125, 80, 104, 171, 120, 186, 208, 57, 110, 159, 77, 98, 4, 79, 38, 115, 118, 36, 81, 156, 137, 175, 26, 152, 52, 41, 170, 115, 217, 106, 136, 54, 232, 52, 235, 59, 246, 84, 87, 23, 142, 213, 246, 153, 61, 151, 91, 224, 41, 41, 214, 90, 176, 126, 1, 130, 174, 89, 29, 140, 131, 102, 83, 241, 9, 33, 97, 205, 229, 198, 170, 72, 104, 191, 81, 170, 104, 102, 92, 132, 136, 15, 80, 118, 74, 198, 137, 232, 16, 68, 103, 21, 195, 212, 70, 135, 220, 70, 115, 107, 176, 121, 49, 220, 212, 236, 34, 92, 75, 182, 76, 87, 217, 165, 16, 32, 175, 220, 70, 180, 251, 247, 86, 142, 194, 213, 195, 153, 36, 239, 149, 77, 241, 71, 18, 17, 125, 116, 171, 156, 30, 12, 169, 228, 181, 42, 13, 124, 207, 197, 36, 110, 43, 172, 20, 196, 8, 73, 29, 102, 152, 136, 246, 185, 5, 151, 117, 47, 124, 158, 216, 107, 24, 67, 154, 231, 114, 191, 12, 108, 89, 98, 19, 56, 83, 59, 184, 198, 100, 6, 255, 67, 200, 132, 42, 215, 36, 97, 226, 162, 62, 251, 242, 191, 104, 214, 97, 103, 38, 48, 173, 176, 46, 211, 31, 49, 110, 217, 9, 34, 87, 128, 254, 12, 10, 13, 219, 193, 38, 235, 46, 217, 10, 254, 198, 161, 224, 134, 27, 41, 94, 153, 7, 224, 161, 52, 78, 255, 49, 13, 237, 46, 233, 18, 100, 151, 183, 136, 105, 71, 184, 112, 237, 42, 249, 100, 217, 71, 93, 2, 142, 248, 171, 18, 159, 207, 20, 17, 20, 137, 113, 107, 154, 207, 154, 195, 166, 146, 69, 69, 69, 201, 126, 223, 76, 120, 111, 115, 211, 220, 142, 15, 185, 67, 189, 170, 105, 18, 50, 200, 158, 185, 4, 163, 184, 101, 27, 19, 3, 58, 147, 6, 81, 204, 202, 190, 135, 8, 130, 131, 252, 203, 52, 56, 158, 79, 62, 236, 138, 45, 222, 174, 72, 109, 96, 26, 77, 94, 248, 182, 209, 165, 147, 105, 54, 211, 103, 252, 69, 123, 199, 119, 12, 175, 204, 205, 28, 14, 28, 118, 199, 86, 117, 137, 39, 243, 57, 180, 166, 116, 138, 214, 191, 134, 154, 105, 131, 63, 162, 254, 65, 251, 202, 195, 26, 19, 146, 189, 43, 178, 252, 192, 121, 203, 5, 126, 238, 68, 249, 191, 170, 54, 167, 21, 161, 28, 113, 107, 53, 17, 218, 215, 124, 167, 10, 47, 50, 219, 176, 252, 127, 65, 173, 224, 196, 109, 123, 16, 247, 20, 192, 138, 13, 111, 127, 95, 129, 175, 108, 145, 103, 33, 231, 213, 155, 53, 232, 26, 219, 246, 222, 16, 93, 0, 222, 243, 38, 213, 241, 63, 206, 145, 232, 124, 16, 116, 179, 71, 19, 164, 75, 226, 212, 209, 122, 133, 209, 5, 201, 107, 2, 3, 1, 0, 1},
		RawSubject:              []byte{48, 49, 49, 16, 48, 14, 6, 3, 85, 4, 10, 19, 7, 65, 67, 77, 69, 32, 67, 111, 49, 29, 48, 27, 6, 3, 85, 4, 3, 19, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104},
		RawIssuer:               []byte{48, 89, 49, 56, 48, 54, 6, 3, 85, 4, 3, 12, 47, 54, 51, 100, 48, 52, 52, 99, 57, 45, 55, 55, 99, 55, 45, 52, 50, 97, 101, 45, 97, 102, 100, 99, 45, 54, 51, 54, 97, 49, 98, 54, 97, 98, 52, 101, 50, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 49, 16, 48, 14, 6, 3, 85, 4, 10, 12, 7, 65, 122, 32, 77, 101, 115, 104, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83},
		Signature:               []byte{195, 80, 220, 134, 68, 195, 80, 207, 169, 148, 148, 77, 10, 250, 18, 173, 191, 118, 169, 177, 109, 80, 47, 207, 57, 219, 159, 113, 4, 254, 53, 204, 189, 229, 186, 163, 26, 144, 91, 202, 133, 163, 111, 4, 57, 167, 61, 64, 102, 27, 94, 27, 113, 63, 222, 36, 79, 127, 89, 164, 28, 216, 216, 4, 235, 197, 129, 221, 75, 169, 226, 0, 241, 34, 137, 230, 42, 40, 125, 4, 62, 62, 146, 60, 54, 115, 248, 201, 153, 98, 248, 117, 59, 119, 52, 13, 228, 83, 242, 104, 77, 44, 190, 16, 243, 39, 58, 113, 100, 243, 199, 99, 98, 96, 69, 214, 238, 116, 64, 125, 79, 75, 142, 235, 236, 85, 131, 146, 22, 48, 238, 160, 39, 8, 144, 134, 14, 28, 11, 15, 222, 188, 155, 68, 86, 98, 103, 145, 52, 204, 56, 119, 202, 69, 95, 15, 38, 46, 93, 132, 108, 221, 52, 147, 156, 188, 116, 73, 215, 216, 253, 28, 137, 46, 203, 129, 143, 196, 152, 129, 199, 42, 99, 218, 28, 254, 77, 33, 165, 180, 107, 154, 226, 51, 88, 151, 195, 208, 51, 4, 66, 31, 103, 200, 79, 36, 95, 170, 183, 78, 111, 241, 164, 65, 139, 16, 254, 185, 245, 93, 255, 183, 32, 155, 175, 58, 18, 135, 142, 182, 248, 32, 151, 131, 242, 195, 219, 70, 165, 10, 239, 103, 123, 75, 216, 118, 167, 185, 28, 135, 137, 60, 23, 101, 92, 241},
		SignatureAlgorithm:      4,
		PublicKeyAlgorithm:      1,
		PublicKey: &rsa.PublicKey{
			N: n,
			E: 65537,
		},
	}
}

var _ = Describe("Test Tresor tools", func() {
	Context("Test loading private key from PEM file", func() {
		expectedPrivateKey := getPrivateKeyFixture()
		rsaPrivKey, pemKey, err := privKeyFromFile("sample_private_key.pem")
		It("Should have loaded the private key", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(rsaPrivKey.PublicKey.N.String()).To(Equal(expectedPrivateKey.PublicKey.N.String()))
			Expect(rsaPrivKey.PublicKey.E).To(Equal(expectedPrivateKey.PublicKey.E))
			Expect(rsaPrivKey.D.String()).To(Equal(expectedPrivateKey.D.String()))
			Expect(rsaPrivKey.Primes[0].String()).To(Equal(expectedPrivateKey.Primes[0].String()))
			Expect(rsaPrivKey.Primes[1].String()).To(Equal(expectedPrivateKey.Primes[1].String()))
			Expect(rsaPrivKey.Precomputed.Dp.String()).To(Equal(expectedPrivateKey.Precomputed.Dp.String()))
			Expect(rsaPrivKey.Precomputed.Dq.String()).To(Equal(expectedPrivateKey.Precomputed.Dq.String()))
			Expect(rsaPrivKey.Precomputed.Qinv.String()).To(Equal(expectedPrivateKey.Precomputed.Qinv.String()))
			Expect(rsaPrivKey).To(Equal(expectedPrivateKey))
			Expect([]byte(pemKey)).To(Equal([]byte{48, 130, 4, 189, 2, 1, 0, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 4, 130, 4, 167, 48, 130, 4, 163, 2, 1, 0, 2, 130, 1, 1, 0, 195, 223, 232, 42, 71, 155, 75, 171, 124, 54, 41, 147, 149, 8, 148, 90, 67, 111, 180, 109, 208, 230, 170, 245, 159, 225, 134, 99, 177, 137, 72, 89, 67, 18, 197, 3, 97, 95, 215, 230, 234, 239, 215, 253, 245, 27, 193, 83, 41, 12, 253, 237, 216, 14, 192, 174, 2, 79, 137, 169, 45, 139, 206, 187, 233, 107, 59, 186, 213, 91, 98, 45, 143, 190, 50, 223, 113, 91, 108, 51, 69, 1, 180, 10, 151, 124, 115, 118, 70, 15, 141, 123, 176, 125, 1, 169, 222, 165, 71, 137, 33, 252, 236, 251, 248, 235, 232, 192, 33, 186, 20, 217, 3, 244, 106, 186, 18, 219, 68, 98, 74, 174, 97, 236, 130, 113, 174, 3, 33, 167, 231, 205, 83, 94, 34, 249, 216, 117, 47, 108, 159, 61, 5, 183, 177, 165, 196, 45, 82, 122, 133, 82, 63, 9, 134, 113, 214, 41, 59, 168, 52, 247, 11, 148, 51, 17, 47, 151, 209, 95, 79, 25, 202, 28, 241, 139, 40, 58, 232, 2, 232, 101, 219, 135, 54, 238, 135, 87, 246, 116, 97, 113, 229, 197, 227, 202, 100, 34, 17, 173, 188, 121, 241, 174, 248, 171, 172, 216, 5, 53, 157, 90, 169, 10, 132, 161, 117, 192, 86, 22, 7, 8, 182, 166, 146, 15, 91, 101, 56, 237, 51, 149, 208, 139, 39, 90, 207, 104, 236, 59, 22, 144, 83, 41, 127, 222, 210, 164, 78, 239, 75, 44, 105, 61, 231, 205, 2, 3, 1, 0, 1, 2, 130, 1, 0, 78, 194, 137, 167, 246, 131, 11, 58, 57, 7, 206, 79, 249, 109, 41, 185, 225, 195, 216, 217, 15, 86, 177, 7, 114, 242, 76, 7, 106, 43, 185, 91, 171, 12, 177, 11, 90, 236, 30, 244, 75, 35, 133, 198, 39, 248, 177, 19, 175, 61, 250, 28, 216, 243, 149, 166, 98, 103, 121, 2, 253, 189, 105, 179, 69, 120, 72, 220, 39, 78, 71, 123, 234, 128, 160, 20, 24, 144, 154, 65, 67, 78, 28, 6, 230, 66, 180, 106, 170, 97, 54, 146, 181, 180, 142, 38, 175, 207, 229, 163, 206, 118, 213, 19, 188, 83, 159, 147, 33, 252, 160, 197, 98, 65, 181, 104, 124, 140, 142, 66, 183, 164, 198, 219, 66, 216, 83, 15, 90, 64, 200, 96, 152, 65, 85, 93, 160, 81, 202, 130, 33, 113, 203, 150, 106, 18, 71, 75, 6, 115, 151, 25, 102, 45, 56, 244, 232, 103, 254, 82, 255, 82, 158, 28, 50, 236, 178, 94, 243, 88, 189, 8, 213, 149, 159, 193, 52, 205, 111, 187, 238, 245, 22, 67, 95, 192, 92, 140, 161, 209, 242, 230, 66, 24, 172, 30, 206, 157, 255, 6, 213, 175, 209, 60, 24, 16, 50, 203, 51, 118, 214, 236, 232, 92, 212, 178, 91, 25, 240, 148, 105, 63, 43, 82, 16, 239, 162, 171, 69, 92, 59, 204, 132, 59, 138, 170, 196, 179, 39, 12, 181, 143, 46, 161, 111, 2, 145, 17, 217, 57, 139, 237, 1, 46, 78, 65, 2, 129, 129, 0, 248, 34, 89, 112, 101, 22, 226, 88, 86, 145, 92, 221, 56, 109, 72, 34, 22, 48, 80, 142, 64, 140, 43, 35, 228, 213, 191, 76, 175, 38, 134, 226, 66, 32, 196, 223, 88, 65, 24, 176, 192, 85, 174, 121, 165, 184, 167, 225, 242, 186, 52, 252, 164, 9, 120, 198, 142, 81, 108, 125, 251, 196, 163, 234, 128, 240, 144, 213, 124, 199, 246, 99, 232, 236, 204, 116, 56, 44, 175, 244, 102, 128, 94, 90, 184, 22, 168, 137, 83, 148, 249, 196, 94, 146, 161, 100, 160, 37, 50, 160, 84, 247, 205, 180, 98, 192, 205, 102, 151, 151, 240, 81, 99, 237, 125, 1, 48, 11, 214, 68, 145, 92, 170, 187, 112, 122, 73, 143, 2, 129, 129, 0, 202, 21, 118, 105, 171, 148, 15, 94, 77, 251, 8, 77, 97, 164, 222, 156, 179, 199, 10, 146, 247, 212, 226, 134, 170, 87, 137, 66, 251, 132, 229, 1, 168, 124, 38, 59, 115, 92, 220, 128, 108, 171, 70, 117, 95, 80, 246, 72, 44, 128, 48, 94, 105, 48, 88, 224, 198, 107, 182, 120, 97, 232, 149, 88, 222, 220, 6, 105, 16, 106, 234, 151, 47, 208, 137, 94, 105, 85, 97, 167, 183, 142, 192, 192, 56, 91, 144, 24, 74, 118, 122, 45, 168, 69, 99, 70, 5, 203, 49, 94, 18, 12, 179, 171, 6, 208, 148, 85, 254, 3, 218, 47, 93, 127, 131, 38, 80, 104, 51, 170, 103, 129, 202, 48, 126, 163, 114, 227, 2, 129, 128, 5, 108, 242, 217, 179, 76, 41, 204, 214, 175, 189, 1, 21, 79, 198, 105, 0, 101, 52, 13, 184, 57, 152, 99, 227, 136, 12, 243, 199, 76, 167, 92, 97, 39, 200, 70, 61, 238, 198, 116, 110, 240, 48, 173, 118, 67, 48, 96, 143, 103, 36, 235, 117, 70, 195, 190, 75, 180, 90, 19, 243, 34, 92, 151, 47, 20, 147, 134, 39, 129, 83, 208, 225, 113, 244, 18, 130, 123, 239, 168, 255, 104, 197, 39, 100, 169, 18, 44, 86, 136, 134, 97, 149, 211, 204, 245, 159, 78, 208, 233, 146, 146, 12, 140, 106, 48, 95, 13, 100, 57, 45, 71, 10, 81, 82, 15, 105, 150, 136, 171, 221, 37, 210, 145, 224, 166, 187, 223, 2, 129, 129, 0, 176, 92, 120, 194, 17, 222, 158, 102, 243, 241, 80, 54, 144, 47, 221, 163, 174, 117, 215, 241, 153, 94, 109, 239, 142, 187, 228, 107, 211, 172, 16, 92, 25, 25, 120, 24, 76, 62, 207, 165, 56, 177, 101, 69, 75, 209, 17, 142, 189, 95, 134, 86, 238, 192, 37, 224, 204, 233, 246, 14, 43, 140, 90, 194, 123, 132, 84, 7, 223, 47, 31, 218, 159, 253, 3, 213, 164, 97, 194, 95, 39, 159, 234, 242, 22, 125, 58, 77, 40, 183, 43, 59, 171, 110, 27, 12, 98, 68, 9, 170, 138, 96, 17, 113, 1, 250, 136, 106, 95, 204, 38, 223, 77, 94, 218, 43, 86, 227, 9, 171, 254, 183, 83, 168, 108, 236, 226, 119, 2, 129, 128, 3, 62, 108, 216, 17, 9, 67, 133, 121, 120, 61, 212, 219, 38, 231, 167, 27, 77, 4, 96, 61, 159, 25, 140, 240, 182, 101, 131, 205, 23, 118, 16, 95, 173, 106, 120, 226, 176, 143, 57, 107, 25, 17, 73, 72, 243, 77, 32, 203, 122, 158, 172, 93, 91, 103, 44, 41, 90, 64, 242, 116, 154, 215, 76, 51, 182, 100, 104, 40, 197, 11, 218, 65, 247, 64, 134, 54, 10, 120, 63, 187, 253, 143, 76, 252, 6, 156, 65, 32, 218, 96, 231, 151, 1, 212, 125, 127, 168, 73, 173, 255, 181, 72, 112, 6, 142, 166, 218, 81, 123, 145, 176, 138, 89, 37, 186, 10, 191, 187, 161, 28, 56, 70, 38, 199, 79, 255, 198}))
		})
	})

	Context("Test loading a certificate from PEM file", func() {
		expectedCert := getX509CertFixture()
		x509Cert, pemCert, err := certFromFile("sample_certificate.pem")
		It("Should have loaded the certificate", func() {
			Expect(err).ShouldNot(HaveOccurred())
			Expect(x509Cert.Raw).To(Equal(expectedCert.Raw))
			Expect(x509Cert.RawTBSCertificate).To(Equal(expectedCert.RawTBSCertificate))
			Expect(x509Cert.RawSubjectPublicKeyInfo).To(Equal(expectedCert.RawSubjectPublicKeyInfo))
			Expect(x509Cert.RawSubject).To(Equal(expectedCert.RawSubject))
			Expect(x509Cert.RawIssuer).To(Equal(expectedCert.RawIssuer))
			Expect(x509Cert.Signature).To(Equal(expectedCert.Signature))
			Expect(x509Cert.SignatureAlgorithm).To(Equal(expectedCert.SignatureAlgorithm))
			Expect(x509Cert.PublicKeyAlgorithm).To(Equal(expectedCert.PublicKeyAlgorithm))
			Expect(x509Cert.PublicKey).To(Equal(expectedCert.PublicKey))
			Expect([]byte(pemCert)).To(Equal([]byte{48, 130, 4, 151, 48, 130, 3, 127, 160, 3, 2, 1, 2, 2, 17, 0, 235, 26, 146, 2, 21, 227, 44, 134, 171, 34, 61, 66, 85, 195, 234, 94, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 48, 89, 49, 56, 48, 54, 6, 3, 85, 4, 3, 12, 47, 54, 51, 100, 48, 52, 52, 99, 57, 45, 55, 55, 99, 55, 45, 52, 50, 97, 101, 45, 97, 102, 100, 99, 45, 54, 51, 54, 97, 49, 98, 54, 97, 98, 52, 101, 50, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 49, 16, 48, 14, 6, 3, 85, 4, 10, 12, 7, 65, 122, 32, 77, 101, 115, 104, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 48, 30, 23, 13, 50, 48, 48, 50, 49, 50, 48, 50, 50, 55, 52, 57, 90, 23, 13, 50, 48, 48, 50, 49, 50, 48, 50, 50, 56, 52, 57, 90, 48, 49, 49, 16, 48, 14, 6, 3, 85, 4, 10, 19, 7, 65, 67, 77, 69, 32, 67, 111, 49, 29, 48, 27, 6, 3, 85, 4, 3, 19, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 48, 130, 2, 34, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 130, 2, 15, 0, 48, 130, 2, 10, 2, 130, 2, 1, 0, 196, 188, 109, 135, 176, 96, 140, 61, 122, 118, 100, 137, 4, 76, 254, 87, 94, 228, 233, 153, 172, 65, 125, 80, 104, 171, 120, 186, 208, 57, 110, 159, 77, 98, 4, 79, 38, 115, 118, 36, 81, 156, 137, 175, 26, 152, 52, 41, 170, 115, 217, 106, 136, 54, 232, 52, 235, 59, 246, 84, 87, 23, 142, 213, 246, 153, 61, 151, 91, 224, 41, 41, 214, 90, 176, 126, 1, 130, 174, 89, 29, 140, 131, 102, 83, 241, 9, 33, 97, 205, 229, 198, 170, 72, 104, 191, 81, 170, 104, 102, 92, 132, 136, 15, 80, 118, 74, 198, 137, 232, 16, 68, 103, 21, 195, 212, 70, 135, 220, 70, 115, 107, 176, 121, 49, 220, 212, 236, 34, 92, 75, 182, 76, 87, 217, 165, 16, 32, 175, 220, 70, 180, 251, 247, 86, 142, 194, 213, 195, 153, 36, 239, 149, 77, 241, 71, 18, 17, 125, 116, 171, 156, 30, 12, 169, 228, 181, 42, 13, 124, 207, 197, 36, 110, 43, 172, 20, 196, 8, 73, 29, 102, 152, 136, 246, 185, 5, 151, 117, 47, 124, 158, 216, 107, 24, 67, 154, 231, 114, 191, 12, 108, 89, 98, 19, 56, 83, 59, 184, 198, 100, 6, 255, 67, 200, 132, 42, 215, 36, 97, 226, 162, 62, 251, 242, 191, 104, 214, 97, 103, 38, 48, 173, 176, 46, 211, 31, 49, 110, 217, 9, 34, 87, 128, 254, 12, 10, 13, 219, 193, 38, 235, 46, 217, 10, 254, 198, 161, 224, 134, 27, 41, 94, 153, 7, 224, 161, 52, 78, 255, 49, 13, 237, 46, 233, 18, 100, 151, 183, 136, 105, 71, 184, 112, 237, 42, 249, 100, 217, 71, 93, 2, 142, 248, 171, 18, 159, 207, 20, 17, 20, 137, 113, 107, 154, 207, 154, 195, 166, 146, 69, 69, 69, 201, 126, 223, 76, 120, 111, 115, 211, 220, 142, 15, 185, 67, 189, 170, 105, 18, 50, 200, 158, 185, 4, 163, 184, 101, 27, 19, 3, 58, 147, 6, 81, 204, 202, 190, 135, 8, 130, 131, 252, 203, 52, 56, 158, 79, 62, 236, 138, 45, 222, 174, 72, 109, 96, 26, 77, 94, 248, 182, 209, 165, 147, 105, 54, 211, 103, 252, 69, 123, 199, 119, 12, 175, 204, 205, 28, 14, 28, 118, 199, 86, 117, 137, 39, 243, 57, 180, 166, 116, 138, 214, 191, 134, 154, 105, 131, 63, 162, 254, 65, 251, 202, 195, 26, 19, 146, 189, 43, 178, 252, 192, 121, 203, 5, 126, 238, 68, 249, 191, 170, 54, 167, 21, 161, 28, 113, 107, 53, 17, 218, 215, 124, 167, 10, 47, 50, 219, 176, 252, 127, 65, 173, 224, 196, 109, 123, 16, 247, 20, 192, 138, 13, 111, 127, 95, 129, 175, 108, 145, 103, 33, 231, 213, 155, 53, 232, 26, 219, 246, 222, 16, 93, 0, 222, 243, 38, 213, 241, 63, 206, 145, 232, 124, 16, 116, 179, 71, 19, 164, 75, 226, 212, 209, 122, 133, 209, 5, 201, 107, 2, 3, 1, 0, 1, 163, 129, 129, 48, 127, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 5, 160, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 101, 36, 31, 192, 240, 180, 200, 88, 248, 137, 54, 232, 124, 41, 186, 209, 184, 23, 107, 96, 48, 31, 6, 3, 85, 29, 17, 4, 24, 48, 22, 130, 20, 98, 111, 111, 107, 98, 117, 121, 101, 114, 46, 97, 122, 117, 114, 101, 46, 109, 101, 115, 104, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 3, 130, 1, 1, 0, 195, 80, 220, 134, 68, 195, 80, 207, 169, 148, 148, 77, 10, 250, 18, 173, 191, 118, 169, 177, 109, 80, 47, 207, 57, 219, 159, 113, 4, 254, 53, 204, 189, 229, 186, 163, 26, 144, 91, 202, 133, 163, 111, 4, 57, 167, 61, 64, 102, 27, 94, 27, 113, 63, 222, 36, 79, 127, 89, 164, 28, 216, 216, 4, 235, 197, 129, 221, 75, 169, 226, 0, 241, 34, 137, 230, 42, 40, 125, 4, 62, 62, 146, 60, 54, 115, 248, 201, 153, 98, 248, 117, 59, 119, 52, 13, 228, 83, 242, 104, 77, 44, 190, 16, 243, 39, 58, 113, 100, 243, 199, 99, 98, 96, 69, 214, 238, 116, 64, 125, 79, 75, 142, 235, 236, 85, 131, 146, 22, 48, 238, 160, 39, 8, 144, 134, 14, 28, 11, 15, 222, 188, 155, 68, 86, 98, 103, 145, 52, 204, 56, 119, 202, 69, 95, 15, 38, 46, 93, 132, 108, 221, 52, 147, 156, 188, 116, 73, 215, 216, 253, 28, 137, 46, 203, 129, 143, 196, 152, 129, 199, 42, 99, 218, 28, 254, 77, 33, 165, 180, 107, 154, 226, 51, 88, 151, 195, 208, 51, 4, 66, 31, 103, 200, 79, 36, 95, 170, 183, 78, 111, 241, 164, 65, 139, 16, 254, 185, 245, 93, 255, 183, 32, 155, 175, 58, 18, 135, 142, 182, 248, 32, 151, 131, 242, 195, 219, 70, 165, 10, 239, 103, 123, 75, 216, 118, 167, 185, 28, 135, 137, 60, 23, 101, 92, 241}))
		})
	})
})
