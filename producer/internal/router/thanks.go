package router

import "regexp"

//nolint:dupword
/**

T. Hanks

HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHHHHHHHHHEHEEEEDDDEDDDDEEPPDDEEEEPDDDEDPZHEHNDEDESZ6S6Y262Scs626666522222scs;6665SYYPYSSSYYS66SSSS6626ZYPZSYYYYYYY2655666666666622226655cc2266c;N6NHH
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEEHEEEEEHHEDPDDDDDDDPPZPHEPDPPNHDDDEPD6SDZDPHHDEYZSZYEYD265P;5;:6c55;;scss2556; ccYYENS52255562s262cs:;:;2c2s:s222566scssss;s22sccccc2csssc222cc2s:6 NDH
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEEEDEEEEEEEDDDDDDDDPPZPDPPPPPDDEEPPDDDEPDZZY6HDYZccZZDDZD2S6PcY s  2:SY25ccsc;s;s. cc55SZSYYYSS6652s::.:sscs;sS6YS665522c22255cc:;;;ssss2222cccccs;;sss NSN
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEDEEEDDDDDDDDEEDDPPPDEEZZYZZP56ZY6S6S66Y6ZYYYEDZYDPSPSD2Z:2 s  : 25c5. 6Z:..:. cs6SSY6S555cccss;.::::;ss52655652:;ss;s.:;;:;....      ;:..;:;:cc;s  NcN
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEEEEDDDEEEDDPPYZZPEEPPYZYZYY6Y56PZYY66cs6565YSDPS6PYZE2Y;6:2:2 .  c2;s;;s2.:..  ;:6ZSSYPc;s;s2s:::;sss;.:s;2c;s;;:scc.: ...      ;:::  . ;::.;.;;:s  N N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEPDPPDDDDDDSYYY6YSY6SSY6Ys25566S6S6PZSS66SS62DZZE5P:5 ;;2  : 22;s:;.;      ;.6P2c5S6Y6c25s.;s:   .css2Z5 .:s22.: .                    . . ;:::. Z N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEDPPPPDPZPEEPPYZ25c526YZs2cc::csS5Y6SSYY625sHDYDSD5E s . .  :.:s:s        ;;6PSYc;c62scc;.      .Z652ss.;s;.:.:                                6 N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHHHHNNNNHEEDDZPDEPEPDPZ.s.;s2;s66cc6c525565s:NDYY5SSD2P s .      ;5          c5cc5Y665Sc::..     :cs5s;: ::;s:                                    : N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHDEDDSZDEEHPE2Sc6.;;s:;s;YS2262DY6sDD5YcY.5:cs2    26:;          2S252225;;. ..      :2; ;Z6;s..                                      . N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEZPZPSZYPDE6Yc5s26S  :...;.6cDZZYSP2Z:2  ;s  ss:s :        : 6D66s;;c:::            ::5s ::.   .                                  . N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHNNDP;5YZDHYPEE26 s:;SSss;.2cs:62SP;6 s    2s56:::s        :.c6s2s2::                 :.                                          . N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNDEYY;c5SDHNNSY;c26666225::2cYE:c;c..  66s2;2.; .      s5266Sc;;c;;.            :                                               N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEHHHHHHHHHHHHHEHHHNNNHNEHHH:sscEHHN:2.sNNPDc26Y.2.:..ss2s66s2;2   .      SN62: ;2ss           .   ;.                                              N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNNNNNNHHHEEHHHHEEHHHHHHHHHHHHNNNNHNYEPPNEDENNsY  ZPZZSP : .222c52YPs2:c :      5cc2s2sss2.          ;.                ..                                  N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNNNNHHHHHHHHHNNNNNHNHHHHHHHHHHHHHHHHHEHHNHNYZHEPPEHEN;6  ss6ZsSssss226S;s.s;5.:      S66Y22sc;s        .c::                                                     N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHPPPPPDDPDEPYPDEHHNNNNHHHHNNNNHHHEHHHHHHHHHHHNEHPDYZDDNNENYDYSZP2YSZSS66scc6c6:;        6656c5::          :6ss:                                                    N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEPZPPZZZYZDZ6YYS665SDDHHNNNNNNNNHHHHHHHHHHHHHHHNHNZD56ZYHNYE6ZPE6YSPsc665S2S.s : .    5s5c;;cYs;          :S;:                 .:                                  D
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEHDEDDPPS52sss. .:;csss2225YSSSZZHHNNNNNNNHNNHHHHHHHHNNHNDDEEHNSPSYZD6Ps2566Ys6 s;2      2265YDcs:s          :Zs;:.             c:      .:   .                        6
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEDEPZYSs ... .;     :         62.:c2NENHNNHENHHEEEHHHNEHEHZPDEPN6YDH5Y6Y2526c6 : .    s;  ;;2Ys2            ;Pc::              2c66YYYSc2;s..   : .                  :
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHNNNNNNNNNNDDS62cc22cc;        ..65ZYHDNHNNNHHHHHHHHNHNHNDNSPYDc6c65Y..s6 s .    csss2c26;c            ;Zc;ss;..       ;YSDYYSS6sc2c  :ssc :
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNNNNNNNNNNNNNNNNNNNNNNHHDEPZ:;....scc2DPHPNHNEHHHHHHEHEHDNZESD6ZZEc2;s26s5cS    55YPs.;c;2.         .;cYs:.      .     s:s:2ssc6S6ZHE66c2 s.;   .
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHHHHHHHHHEHHHNNNNNNNNNNNNNNNNNEHDDYY56 .:;;:NZNHNDHEHHEEHHDHPHPHYDc6ssYZ262Y;c .s:YPcsssc6            :scScs: .        ;:cscsZYPP65  :: :2625  .; .            ..
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHNY26sDNNNNNNNHHHHNNNNNNNZD6PcS ;  s NSNHNHHEHHHNEHDHDHDNSZZZEH6Z:s;2scss;;  s5c5             :.::;;:.      ::csc2ZYEDZZc6YY56ZZ55 .6S :
HHHHHHHHHHHHHHHHHHHHHHHHHEEEEDEHNNNNHHHEHHHEEHHHNNNP    :6NNNNNEHEEEHHHNHNPH5P:6      NSNENHHHEHEHDEPDDHDEDDDEPHSZZE5Sc5;ss2:s              :c;;: ..:;::s65Y62265Y6;; :55SSS6scs5ED662Ss5 .    :.
HHHHHHHHHHHHHHHHHHHHHHHHHEDEEEHENP6DDENHHHHHHDDDDNNN56      sSNNNHHHHHHHHNDNYH6E;S    ; NSNEHHEHEHHHEHEHDHEHEHPDPDHNDNc2.:s6.:               ....:..:sscc5csc26S652scssc;.:ss::      .: :            :.
HHHHHHHHHHHHHHHHHHHHHHHHHHEHHHNYS5;HNNHEHHHHHHNPS665 :        Y5NNNEHENHHHENZHYHcZ.s. 2 D:NEHHHHHHHHHHHHEEEEEHDHDEPDSE2Y5Y;s                  .:. .      ::s;;. :665555YSccccccc2 :;;ss .
HHHHHHHHHHHHHHEHHHHHHHHHHHHHEDH6Y; ZHNNNHHHHNNNDZ:              P6NNDZHEHHHNENYE6D;562PcNYNEHHHHHHHHHHHHHHHHHNEHDHPHYE5Z : .                       ::.:s:            .   s;...:c2.:                    .
HHHHHHHHHHHHHEEHHHHHHEHHHNNNNNNHNs   6PHNNNNH5 6NNNE              HSNNNEHEHHENYEPHSYY6NDNNNHHHHHHHHHHHHHHHHHHHEHDHPH6P5D:s                            .2ss;                .                           .
HHHHHHHHHHHHHHEHHHHHHZZ6655SPDHNNNNNN6 2DZZSSPHNNY .                NNNDNDHHHNPHPHEDHPNHHEHEHHHHHHHHHHHHHHHHHHHNDHZH5Z.c : ...                         ..s.                                            .
HHHHHHHHHHHHHHHHHHHHHNNNNDZ5c       ;YDNNNNNNNNs                      2;NHHEDEDH2SZSNNHHHHHHHHHHHHHHHHHHHHEEHHHHHNDN6DsY :.:                .c;:::. :s:      c6NNNNNNNH55 .                            :
HHHHHHHHHHHHHHHHHHHHHEEHHNNNNNNHDZY:                  ..      .   s;6cHSNNNNHHEN6YPYNHHHHHHHHHHHHHHHHHHHHHHHHHHHENDNYN:2 ;                 ..;:;. .:     6PNNNNNNNNNNNNNNNNNNSY:c                      ;
HHHHHHHHHHHHHHHHHHHHHHHHHEEHHEHHNNNHHPDY55c2c5S2;;:: ... ..   ..    ..Z6EPNNNNNNZPHDNHHHHHHHHHHHHHHHHHHHHHEEHHHHEHDNYHsS .                 :    .:     DHNNNNNDPPPDPHNNNNENNNNNNNNNsS
HHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHEHZZYP66662:::.:::    .:.                    P5NNNNNDHHHHHEHHHHHHHHHHHHHHHHHHHHHHENDN5YsS:2 :               ss6c:    ;HNNNDY::ccZYYS;s 2NNNN;6ZHNNNNNNNNYD .            c
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEHNHHDDSY52cscc:    ;:                          : NPNHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHNDNYDSE:c              . cP5c.   sDNNH55 :25NNNNNNNN2Y ; :DS    NENN6YENcZ 5          H
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEEHNNNNNEPZSSS6252s:::.                      ;   Y:NZNHHHHEHHHHHHHHHHHHHHHHHHHHHHHHDNYD2Y;5 :            :.5H;   cNP2YNS    NNNNEEEHNNNN            Y;HZDDNN:P:Z        N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNEHDDPPS222s:ssccs  ..    :            :5cP6E5DcNPNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHENYEs5:5:s               :  2EPN5 sY5ss   ;NNNNNNNN;5              ND5cHH5N 6      c N
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNHEDDDDDDZ5522s26SD6 :s;::::;;cc522252: ; 6 H6NHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHENDNSE:c : :                2EHNESS6SPZss  :.65S626 c s.;          Y.6:HPSN2N 6    NcH
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHDEPDYSSYZPZZZZ6222c:scscs . ..:;;2c2222;:    2 PsNENEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHENPN2Z s.s.                    ;:        ..c;s;;;;5YY;s          : PsHD6H;D 2  s NSY
HHHHHHHHHHHHHHHHHHHHHHHHNHEEHDDZPSYc:;;262s   s;;        .;.      .         P;NNNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNDNENs6                                       ::s:                P;HP S      NZEZ6
HHHHHHHHHHHHHHHHHHHHNNNEEDEDDDDPDYYYPYY;  ;sssc;;s;;..:.            :.    Y;NNNNHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNEHZN5D .:s                       :.:::.                          5 HP 5 ;  D:NZ666
EDDDEHHDDDDEEHHHHEEDDDDPDEHEEEEDHPPYY6Y6255s:;;. .;s2s .    ;s62. :     P6NPNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHEHHNDN5Z:s:s                         . .                             Y5.2  5 NHYSSY6
HHHDDYYPDPDPPYZPZDDPDYYZPPPZPYYYZ2c25562c5S52s;:.:::s; . .:     csc;s;Y6NNNNHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHENZEYE6S:.                       ::2s:.                          .     NYPS6666S
EEEEHDEDDYZYYSZ52PDPDDEDEYS226Y2225cs::::: s22cc26S5ss;;;::;sc2s5s22NHNNNHNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEHNHNENZP25:                        s:::..                            Z NP6666SSS
HHHEHDEDEPPDDPDYZZZSY2625SY;..:c556csss2Ss    .cYc.;5c:cZS22ss5scsDPNHNHEEEEHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEEHHNENPE52;s..                        .5s   .                        N5DY56SSSS6
HHHEHDEDEHHHNHNPEc:6Yc6  s2SDs    :.::::s: .:::  .2s.:cs;scc22PZNHNNNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEHHHEEEHENZZ6Y6Y:.                         ::::                        . NYS66666YYS
HHHEEEEHHEHEHDEZDDDEHDNcs2cSP2222: .:        :;:.:;;;;2;  .sPPNNNNHEHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHENPPZDZEs:..                                                   . N66SYY66656
HEEDDEEHHEHDEPEPDYSDHYP6Z2;5Y5Scs::. :s:    .. :2.  scY65SPDNNNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHDHZZSP56: .;:                                                  NSSY5565SS6
EEEEEHEHHEHEHEHDHZZDDPESZ555625ssc5:  ...c5s.  sP5;266ZYPPNENHHHHHHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNHNDH665S2c:::;.                                              : NZ5S2cY6YYS
EEEHEHHHHHHEHEHPEYYZZDHZES6cc  :c  :ssc:    ;6PYZ5;ss;56DDNHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNDH6Sss5Y2c;.                                               : NYc565ZYSSS
HEEEEHEHHHHHHEHDHZPYZPEYPSS6Sc;5Zc:;s;ss;c225SZ662;2SYHDHENHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHENPE2c  :2s:                                               . NY6Y55Y6S66
HHHHHHHHHEEEHEHDEDEPE6SSY665S5522;:6Y52c;5Y5;;;c6YZYYYDDHEHHHHEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNDESY25.  .:;:                                            . NYSZc252SS6
HHHHHHHHHHHHHHHHNHNEN2s556SPHS6cs;:ssss2S65565c66YZZEDHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNEHDHPDc;: ..                                               NS6Y6SS6556
HHHHHHHHHHHHHEHHHEHDN56SSZPSS2sSPPE5:5S5c22SZS6SSZDDHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHNDDZDSZ2 . .                                              NSSS56YSSY6
HHHHHHHHHHHHHHHHHHNPEYYYYZPSS66YPYP6cYPS22;c25SYZPDDNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNPDYZSZS5c:.                                              N5YY56YSYY6
HHHHHHHHHHHHHEEHHHHEHEHEEPPZPYSZP2;5SPEY52s6ZPDPZDNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEENPP6S55c::.                                             NcDP25YSYY2
HHHHHHHHHHHHHEEEEHHHHHNHNEDDEDPZZZPY62;6PDEEDDEHNNNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNDHSS2;sss::                                            N;PZ266566c
HHHHHHHHHHHHHHHHHHHHHHHDHEEDDPPYYYZYYYYZZZZPDEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHEHHHHHHEHDHPPYS2s;..                                            N:HPc652SS6
HHHHHHHHHHHHHHHEHHHHHHHDHPYPPPDDDPPPDDDPZDHHHHNHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHEEEHHHHHHHHHHHHEHDEPDZY62s:                                             D HPc622YY5
HHHHHHHHHHHHHHHHHHHEEEHEHPYDEEHDDDEEEEEHHHHHEEHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEEDEPEZPYS552;:                                            P NZ5Y22YY6
HHHHHHHHHHHHHHHHHHHHHEHEEEEEEEEEDEDEEHHHHHHHHHHHHHEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHHHHHHEEEEDEZZZZ62s :;.                                          Z NY6Z55552
HHHHHHHHHHHHHHHHHHHHHEEEEHHHHHHHEHHHHHHHHHHHHHHHHEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHHHHHHHHHHHHEHHHHEEEEEEEEEEEHPDZZYZSS2:::                                           6 NZ566SSS5
HHHHHHHHHHHHHHHHHHHEEEEHHHHHHEEHHHHHHEEEEHHHHHHHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEHDEDEDEPPPPZZZZYYPPPPPPDDEEHHHHEDEHHHEEPPZP62cs;::                                          P NZSZ55ZZS
HHHHHHHHHHHHHHHHHHHHHHHHHEEEEEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNEHDHPEPPYYY6YYYYZZZZZZZPYYYYYYYYZZPDDDDDDEDDPDZZ62;                                            H HYSY6SSS6
HHHHHHHHHHHHHHHHHHHHHHHEHEEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEEHHHEHDHPEZDSSYSHENNNNNNNNNNNNNNHHPDYYSS6S6SYZYYYZPDPP622c:                                          N:EYYY66SS6
HHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHHHEHDHDHZDYZYSHDNNNHHHHHHHHHHHHHHNNNNNNNHHDPYZZPZZYYS665c::..  .                                     N:P666SSYY6
HHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHHHHHEEEEEHEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHPEPDYZZYHENHHHHHHHHHHHHHHHHHHHHHHHNNNNNNNNEDDEDPYS65c;:                                          NsDYSS6666S
HHHHHHHHHHHHHHHHHHHHHHHHHHHEEEEEHEEEEEEHHHHHHEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHPEZPDPHDNHHHHHHHHHHHHHHHHHHHHHHHEHHHHNHNHNDHPZS62ss;.           ;YS62                            NcPYYZSYYY6
HHHHHHHHHHHHHHHHHHHHHHHEEEHHHHHHHEDEHEDEEDDDEEEEEEEEHHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEHEHEHEEEHEHHHHHHHHHHHHHHEHDEDEPDZZYPYPSZc;:.         :s66ZZPZ6c                              NSPZSYSYZP5
HHHHHHHHHHHHHHHHHHHHHHHHHHEEEEDDEEHEDPPPDDDDDDDDDDDDHHEEHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEHHHHNEHDDYZYP66;:        ;YDNHDPSS65s;                                NYPPSZSY6S5
HHHHHHHHHHHHHHHHHHHHHHHHHEEEEEEHHEEPPDHDPPPPDPZPEDDDPPDDEEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHNHNNNNNNNNNNNHE65:.      ;5NNNHZ62sc2c;:                                s NSSY5Yscc26
HHHHHHHHHHHHHHHHHHHHHHHEEEEDDDEPDDDPPPPPPPZZPZZPEDZPPPEDDDHEHHHHHHHHHHHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHEHEHENYPs2c6     :               ;;c;;;::..                                Z NP6Ys2666S2
HHHHHHHHHHHHHHHHHHHHHHHHHEDEEEHPPY6PEPYY6YZZPPPZYZZPDDEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEHHNDHPN5P :                       ;s522c;:                                  NsHZ6YSZ66552
HHHHHHHHHHHHHHHHHDDEEEEEHDDPPDEDDYSZPZYYSYPZZYYYYZZYZZHEHEHHHHEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHEHNDHZE5Zs6.s                      .;:;.                                      NSPZSYZP66YSY
HHHHHHHHHHHHHHHHHEEEEEHEHEHEHPZPEDHScSZPDDPPDZ6SSYSZHEHHHHHHHHEEEEHHHHEEEEHHEEHHHHHEHHHHHHEHHHHHHHEEHNNNHHPZHHNNHNYDs6 .:5.:                                                               NYZZ6Y66ZYDPP
HHHHHHHHHHHHHHHHHEEHHHHEHEEHHHEENZS62YZYYYSYZYYYYPEHNNNNNHHHEDDDZZPPHHNNNHEENHHEEEHEHHHHHHHHHHHHNNNNDEZYNNDEEN2Y6PDN26c5::      .s                                                       5 NPSY6S55ZSPZZ
HHHHHHHHHHHHHHHEEEEEEEEEHDDZYDHZZs:SZEHEPY6S6SSDNNNNNNNNDDPDDDDDZPDPDPYZZZEDDPHEHEHHNNEEHHHHDEZZHNPHDDDPHEHHDN2P .:6 :          cS25.                                                    NcHZSYYZZZ52YSS
HHHHHHHHHHHHHHHEEEEEEHHPPDHPPZYPEPZPPY666YZPPEHHHHNHDDSZPDDPZZPPDEEENH26YYEPZZNNHHHEEEHHEEHHDEPDEN6Z62HNZYEEZDDN2Pc6          ;.c6:s                                                     N6YSYYSSZZSSZYS
HHHHHHHHHHHHHEHHHEEEEEHEHHNPZS5PHPZDEZYPDEEEDENHHDPDEP266YZZYPNEZPYZDDZPDPEDEENNPPEEDEEEDPDDHDHNSEs2EZEE65DEDNYN.6:2 .          sS;2                                                   ; ND56YY6SSS66YYS
HHHHHHHHHHHEEHHHEEEEEEEEHZYZPDDDPPPDEDPDDDDEHHNHHHEEEEDNNP.s56EDPPZYYZPDHDEDDPSSZZEDNH66EEDPNENNSHYYHPPDEHDHZEsZ :              6Hcc                                                   E:NDSYYYYYZYZZZZS
HHHHHHHHHHHHHHHEEHHEHEHDDDEHNZ6YZPPDDPZPDEHHHHNHEHEHNDSSsSZDNDDY:cYZNH6SYYEPDDPPNNYYDPDDHH65NHNNSEPZNNEHHNSY5PsYsY2Y:         ;.5Z::                                                   N6EZYPYYYYSSPZYZ6
HHHHHHHHHHHHHEHHHHHHEHHHHHNDEPZDHDPDEEEEHHHHNHHPSYSPENNE.ss2YSZZEESYPZPZDPHDHHHEED56DPNNYZHNDPNN2S55DPYYscPHSDcY2P              s5sc::                                                 NZYSYY66ZZSSYYYZY
HHHHHHHHHHHHHHHHHEEEEEEHHEHDEEHHEHHEEEEHHHHHHDDDDDPZS65ZDP255SEPPPc22c66DZHDZZ66NHNNHEDEEEDH25ZPEHEE66PEEHSY6EsS.5                c626.                                              S ND6SZZSSZZYSPZYYS
HHHHHHHHHHHHHHHHHHHHHHEEEEEEHHHHHHHEEHHHHHHEEDDDEZ6YDP25s26656PYYY55YSZYHPPZEEPZNHHHHHHHEHEHDHYYDDHNPEYZYPYPSH : 2                   ;                                               NcNP6SYYZZYYSSYYYYZ
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHHHHEHHHHHHHHEEEEDPPPY;6ZY66PS65c266NDDZZZHEHHHHHEHHEEDEDEHNDHENDHYZHN6SSSYP;6 s .                 . .                                               NYZYZZYYYYYYYYYYZYZ
HHHHHHHHHHHHHHHHHHHHHEEHHEEHHHHEHHHEPEDEEEHHHHHDZDHZs2;cccc6NDPY25DYPZPDDDNHHHPDHEEEHHHHHHHNNNZDPE66DEYPS6PHc55Y:6                                                                   NDYYYYYYYYYYYYZYPPP
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHDPPZEEHNHHNNHDZSY6YY6.: .;s6255PYNENHHHHEHHHHEEHHHHHHHNDHDHPEZEYPZPPDZPZYYZZPPEZH6Z                                                               D.ND6SZYZZZZZZYYZYZYY
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEENDDPYEHHHHHDPPPYS2 s66c6DY6522ZYNENHNHHHHHHHHHHHHHHHHHEHDHPEYDSZYYZZSY55YSS6YYZPZDYD6Z.:                                                           NcHZSSZZYYYYZZZZYYYYZ
HHHHHHHHHHHHHHHHHHHHHHHHHHHEHEHHNPSZYHNDDPPZYYZZZY55sss5D6s5PPHDHENHHHHHHHHHHHEEEHEEDEDEPEYPYP6Y55S6SYSSPZPPDPDEPDYPSP2S26c5                                                       NSPP6SZYSSYYSSSSYYYYZ
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHENPZS6SSPPDNPSYYSS2;s22;s:;56DDHDHDHEHHHHHHHHHHEHEEEEDDPDPDSZSY6Y66YYYYYYZYPZPPPPDDDEPDZDPN66c6.;.                                                  ; NDYY66ZZYYZZYYYYPZZZY
HHHHHHHHHHHHHHHHHHHHHEEEHEEEHDNS5YPY6SSZEZ6SP5;c25S5.:ssYSPPHEHEEDHEHHHHHHHHDEDDHEHHHHEHDEEHHHNNNNNNEEEEDDPDPDZDPEPDDHYP6Yc6s2.          .                                       Z NP6SYYZZSSZYYYSSYYZZZ
HHHHHHHHHHHHHHHHHHHHHHHHHHHEHEN6cSZ62PNY;cs6E6;262ss:2HZZS6SDPEDNHEEEEEEEEEEDEEEHHHHHNHNNNNNHNNNNNNNNNNNNNNNHNPESPYZ56262Sc6    .:     .c2c2:                                    NcDSYYPPYZYSPZDDPPZZPPZ
HHHHHHHHHHHHHHHHHHHHHHHHHHHENDNDE5s2622cc555S2ss:sscs6NDHDDPEDHENHHHHHHHHHHHEHHEHHHHEEEEHHEHHNNNNNNNNNNNNNNNNNNNHNPDSY6S25.:  .;    :s    :2s;                                   NSYSPZZZPPPPPZPDYPZYDPP
HHHHHHHHHHHHHHHHHHHHHHHEEEEHNY6ccYE2::.;ccs25s:;;:    . s6HHNNNHNHHEHHEEEEHHHHEEHHDEYZEDEHHNENc665DN2D  5cs6;c2626:::;26;;:;sc    :sc2c;25cs;c;:.                                NDZZZYPZZZYSPZZDSZ2cPYP
HHHHHHHHHHHHHHHHHHHHHHHEHEEENENZZcsss22s;::..::;;:  52DScc25PZDDHHHHHHEHHHHNNNNNHHEEEEHEHHZH2H  ZP2D :  52 .                            . s26P5s:                              6 NHYZZYPZYZYYPPZD2S25Y6P
HHHHHHHHHHHHHHHHHHHHHHHHHEHDEPDZEc;ss:.;c26s ::  .6YNNNNEDPPEPPPPPHEEEPDHHENZPPENNNHNHNNNNDNsZsS2D .    22 s                                  :5s;.                            N5NP56ZYZZYYZZZPSYSYYZ65D
HHHHHHHHHHHHHHHHHHHHHHHHHEHDE662Zs;: s2   ;:     5NHNNHNNNNNNHHHHEHHEHPDED5S:sZP .52NEDPNNHNNNNN:P :s6  cc                                       s.                            N6S655DPZZPPZPZD66DPZPc2Y
HHHHHHHHHHHHHHHHHHHHHHHHHHNDHDNEN2;  ;2s;::.   6HNNNHEHHHHHHHEHHHHEEEEDDHHHNNNDE.sZ6csYS:c  YZ.52H 2 :     ;:s                                    .c:                        : NE66PYDPZZPPPPYP65DZDEYD6
HHHHHHHHHHHHHHHHHHHHEHHEHHNHNYS;;s6s;. :2s:: c6SPPHDNHHHEEEEHHHHHEHHHHHHDDHHEHENHNDPHDNNENPDYZ6D.6.S c  .::2                                         :                       S:ND6SEPPPPPPPPD6SZYDZDDZE5
HHHHHHHHHHHHHHHHHHHHHHHHHHHDNSYcc;;.    .  . 5DDNDDDHEHHEEHHHHHHHHHHHHEHHHHHHHHNDEDDNNHHNNENDNDNSHSN s5Sc6 :                                        .ss2:                    DcEYYYZZYYYYPZYZYSPYDPPPPEY
HHHHHHHHHHHHHHHHHHHHHHHHHEHPDSY6Sss;;:.::: ssSEPPDHEEDDDPEHHHEHHHHHHEEDDPPHEHNHNHHHHHHHHHHEHHNDNYH2Dc6YD;2 . .                                        :Y65s                  NYS5YSSY6666PZ6SYYPYDDZZPDZ
HHHHHHHHHHHHHHHHHHHHEHHHHPZYPYDYZ6S. s5c.ss6ZZPDNHHEDEHEDPYPEHHHEEEEHHEHPDNHHNDDNHHHHHEEHHNNDNs6EN:6SP5Y ;:c;s;;  .:  ::s2;;                             .                   NPYYNEEHZZDZPZSSSSPYZZYYZYY
HHHHHHHHHHHHHHHHHHHHHHHHHHNDDPH65;;ss22c:6PPDDEENDPDDEEDEDDPYYYZDPEDHHDEEEPDPP56EPDDEEDEDHPDPN2Y6Z5P5Z:s2Sscsc::csSZcs25;;:.;ssc.                                          E2E652PZYZ22PZZPPPPPZZZZZYDPD
HHHHHHHHHHHHHHHHHHHHHHHHHEHEN2s;c:::.s2: cZS6PHHNHHDYYYZEDEDZPYPPPEDDDDDDDYZDDZYHDEDDDPDZE;cPEsScYcS . :s5  c2  s;;;;s.:  ::ccc225s;                                       N5DYS6PZSYSSDZPZZYZPYZYYPZHDP
HHHHHHHHHHHHHHHHHHHHHHHEEEHDHYYPDYYc:2552SYYYPDDEDDDDEHEHEEHHHPEZPHPYYPDDE6SY6NHDDNHHHYDPH6YSYYH 5 s:2:s.:  ;;:s    :c .         .c6:                                      NZEPZZDDYPDPPPDPDPPPPDYYZYDPP
HHHHHHHHHHHHHHHHHHHHHHHHHEHEHHNY2SDZZPZSc56YPPPZZDEEDDHHNHHHEDYY.cDPHEEDYZ26ZYEDYYDPEHENYDPE22HN s;22S    s:s;    ;s.:              .:                                     NZDDPPZZSSSS66YYDPDDDEPPYYYYD
HHHHHHHHHHHHHHHHHHHHHHHHHEHDEYYYZDH2;56PHEDDDEHHHDPDEHNHHEHEDDDD5SZZZPNHSYSZSYS622ZSc26D ;56  YD6E.. : .    5S56.:sc::                                                   : NPZZZP6SZYPZPPYYZYDDZZPZHHPDZ
HHHHHHHHHHHHHHHHHHHHHHHHHEHDDZYPDZZDNY2ZEPZDNHHEHEDDEP6ZNHDDYDNNDDSS.;S6YZDHs522Y6PPPPsY  2cscs6  . 56 .Y6PZ   ; ;                                                       5 NDPPYYSYEDEDDDDDZPYZYZPZEDDEY
HHHHHHHHHHHHHHHHHHHHHHHHHHEHHHNHHHNEH5;s;SEDEEEDDENHEDPDHDZZYZDDHHZDSY65ZZYD s;sDZcc6SPN:sS5YD ::;s:DE  : 66PHSD:c                                                       N2NEZPZPDDDPPPPPDDPDYZZZYSYSYZ6
HHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHNEHDDZ5YDDPDEDPY6YYDNNNESPHP6SEDYZ26PPDD6Ys5::::6652;2;c2sc6  ::;:    S52cNNcS .2Z                                                     NPHDPPYZPPDDDPDDDEYPZPYYYSPZYPY
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNHHDDEHHHDYZZPDHNHPZ6YEDYY6YEDYYHDSZ6YNE265S5ZDEYS65.;26.;..22. YY  : DZDDEN::6D .5D..                                                   NPDPDDDDZZZZPPDDZZYZZPYYDPEEPES
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEHDDDDEEHNDYZYDNEDHNNDPSS26SSYYYZNNHNEHSY266Y6Z  ;.:;25::25;s;.Ss:.: NZHD5S5D  c6s6                                                     : NDYZPPPPPPPPYYYYPPPDZZEDEPEEPES
HHHHHHHHHHHHHHHHHHHHHHHHHHHEEHHEEDEDDEHHNNHNNDSZPYYPND2556YZDZ55SSPD66 ;6Y6Z.s  ;:YS :;:66..  5 NDHPNHNNEN;6ZE;s.2   ;                                                 N5NPYZDPPPDDPDDDPPPDYZPZDPDPDDZPD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEDEHEEPSPHP6PHEDY26ZENNDP6SYS66s2 sSYYS   s :6ccs        s Y;NNNEED55DH2S  6P:c;5                                                   NYHDPPDPDPEDPPPPPPPPSYPZDZDPDDZZD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHEHHEHHHEEZ6YYZPDHHNHEP25cYNHHNDE6YPYs2 s5S  . ::cc. ; DYHPHP..Y6EZ55ssNN22YH:2YE;;                                                 NEPPZPDPPPPPPPZZDDPDPPDPDPEDDDPDD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHNHHEEHNNHHHDZPHEHHEHDDDPDDs6YS.s25;s   .;;  :   NZPZ6SY5SY  DYNNZY6YZPYH    sS                                               Y;NEZPYZDPPPDDPPPPDDPDZZPPDPEDEDDDD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEDDDDEYcYEDPPPPHHHNEEEEPP2526c2:2SPSS      . ;.ScDE  : ED;2    YD:.SS5Y:2:;                                                 N2NDPDYZDPDDEEDDEDDDZPYZPPDPDPDPEED
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEDDEHHHHHHHHEEEDZPPENHDPYZDP6YNNNNSPHP  c26Y  c;ZYsc      : 5cEH  ;::c    :cSZ55                                                 NEHDPPSYPPPZPPZZYYEDDEPDYZYSEPPPPZE
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEDEEHNNHHEHHP6YDENEYPEDNHEEZDPEc6DPs5;sZSDEs2ZPSZ66PP  :.YZYZs2 .25    .c2S                                             PcNEZZYPZZZYSSDPPPYYEDPDPDSZS6DPEHSSD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHEEEDEHHDDPDEP6YPZPPHNDD2YHHPP5Sc6HDs5scYS26..s2;2EZZ5;; .:2.:  ::                                                 NPHDZPPPYZDPDDDPPPPPEPPPDEPEZYZYZDSSD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHNHHHEYPPPEDYZYPDHDHYZNN6ZZD5Y..SS25c2SY.2  6:YZDN;s:;  c6.:                                             H6NNDDDDPDYZEPDDPPDDPPPPEDDEZDZZHEPD66P
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHNNNNHHHHHEPDPDEP:s;sHEYPc6 :26c22c6625cZ  6 DD;2:cs2cc56                                               NENDDDEDZP6YDPDDDPEDDDYYEDEHZDPZDPDEPPE
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEPZYZDDDDHHHE6S6ZPDEDc5EESPZP;s ;  EH  NN  2 65:2:2  c5  cY.                                          ssNEDDDDDPZPZPPPZZDDEEDEZPPPEEPDPZDPDEZZE
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEPDHHHHNHEEDHNNPEDENNSSc2EE5SENNN22ZPsc s:5; 665Y:s     :                                             NENHPDDDPDDDDDDDPPPPDDPDDDZZYYZZDDDDDEPPD
HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHHHHHHHHHHHHHHHDDZPNNEHSSSYPDNNYZHESD sNN:.NN5Ds6:.;s   :scc2                                             .sNHEDDDDDPDDDDDPPEEPPDDDDDDDDDDDDDDEDDDDDD
HHHHHHHHHHHHHHHHHHHEHHEHHEEHHHHHHHEEEHHHHHHHHHHHHHHHHHHHHHEEEHHHEEHHNNHE6SYYDHZZHD6ZPENN ;. DSSP2E.s2:PE 2  ;s ..:  .: .                                     NPNEDDDDDDDDDDDDDDEEDDDDDDDDDDDDDPEDEEDDDDD
HHHHHHHHHHHEEHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHDEDDED6YDENNPEc5NNNN565Z6D  NNPN .  ; HN2P      .:::c2;s                                    .NNHDDDDDDDDDDDDDDDDDPDDDDDDDDDEDDDDEDDEDEDD
DHHHHHHHHHHHHHHEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHHHHHHHHHHHPDDDHHHHNHEHPEDDSSs6ZDc2ZP6S  :5  2:Z62Y :;c5S:;                                           EYNEDDDDEDDDDDDDDDDDPDPPPPPPPPDPDDDDPDSYPZEDD
EEEHHEEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHEHEEEEEHHHEEHHHHHHHHHHHHHHHHHHHEEHHHEZPYZPDNNYP56c66SNNPH  c;NEDDYN .c2.:s6  ;s::                                   .2NNHDDDDDDDDDDDDDDDDDZPEDEEPPDDPPPPDPDEZPPZEPD
HHHEEEEHEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEEEEEHHHHEHHEEHHHHHHHHEEHHHHDDDEDHEHNHDD2YYDDHss .YPEENDcs2S2D  sc   .  ...:;s;s                               NHNDDDEDDDPPPPEDDDEEDDPPDPDDDDEEDDDDEEEHDEPPEDE
DEHEEEEHHHHEEHEHHHHHHHHHHHHHHHHHHHHHEEHHHHHHHHHHHHHHHHHHHHHHEEEEEEDEEEHEPDDEZPYZHHDHZD  PD6YcsS5S6EEPN.2  6Y2S;c  s2:::::c                             PYNHDPPPDDDDPPPPEDDDDDPDPDPPPPZZPPPPDDPPPDEEPPDPD
DPDDDHEHEHNEEEEEEEEEEHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHEEEHEEHHHHHHEEEEHEDEEEEHZDPPYYDEs6YZZP6Z:.2c66ZZPH sc2S6.:s5::s2  :; :                            .NHHDPPZPDDDDPPZPDPDDPDPPPPPPDDPPPDDDDDPDPPZPDPEDD
EDPPPPZDDEEDDEEHHHHHHHEEEEEEEEEEEEHHHHHHHHHHHHHHHHHHEEEEEEEEHHEHHHHHHEDDDEHHYP5SEHEHYPEHSZ;2EDEDPD .6Z5Y.   s6 :56::   .  :c   .                     25NNDDPPYPEEHHDDPPEEEEEEDDPPDPDDDDDDZZPPDDPESYPZHDD
EHHHHEEPPDDDDEDEEHHHHEEHHHHEEEEEEHHHHHHEEEEHHHHHHHHHHHHHHHHHEEHEHHHHEEEEHHEHHNNNZPs5HHSZEHSZ;:ZP;5PD  ;cssNNNN:;;2c5csc5.;   :                      YNNHDDDPPPPDPDDDEEEDDPPPPPDDDDDDDDDDDDDDDDEPEPDZSHDE
DEEDDEHEEPPDDPPDDPZDPEEHHHHHHHHHHHHEEEEEEEDEEHHHHHHHHHHHHHEEEEEEEEEEHHEEHHEEDEPEPHZZDDc566DESSHN;c2cSYEHYS65c52Yc5  ;c6S;c  :s                    sNNNHDDDDPPPPDDDDDDDDDDPPPDDDDDEDDDDDDDDDDEDEDEZZDZHDD
DEEDDEEDEDDEEDPEEPZDPEEEEEHHHHEHHNNNHHHHHHHEEHHHEHHHHHEHHHEEEEEEHEDEHHHHDEEHEEDESZYZPP6Z25ZYZYYDc5DPEHS6ZZ66665Ys25Ys;:.SH::  .s. :;            :PHNHDPPPPPDDDDEDDDDDDDDDDDDDDDDDDDEDDDDDDDDDPDZDZZEPHDD
EDEDEDEEEHHEEDDEEDDEEEEDDDPDDEEDPPDDDEEHHHHEDEHHHHHHHEDDEHHHHEEEHHEEDEEEEEPDPDEHEHZPYYPD6ZssHH6Y52PZ6S55YSS6ZPHN22YD: ss5Z2Y;:ss:.  .:        5NNNHZYZYYZPPPPPDDDDDEDDDDDDDDDDPPPPPDDDDDDDDDDPDYPSYDZHDE
EDEPPPPZYZZZZPDPPPPEEEEDDDPDEDDDDEHEDDEDPDHHHHEDDDEENHHEPPDDHEHHEEHHDDDDDEPDEDPP6SZPSYHNENYPcc25SSYS555566EHZZc5SZENYSNNY626. ::sc         cDNNNEDPPPPPPPDPPPPPPPDPDDDDPPDDDDDDDEPPDDDDEDDDDDDEDEYPYSHPE
EDDPDDDPPDDEEDEPPPPPPDEDDPZZPPZPDPPDDHHHHHEEDDEEHHHEHEDDDDDDEEHHDDEENNHHDDZPYZDEDEDEDD66SY2SZZDEEHSSHHNNHNNNNNNNDHS6ss                PNNNNNNDPPPPPZZZPDPDDDDDDPPDPDDDDPDDDEDEEDEDEPPDDDDEDEDDDEHHNSSPYE
DPDDDPPEEPZEEZZPPPDPPZZYYZPDDDPDDZYZZPPZZZZPDDDDPDDDDDDDPPPPPDHHHHDDEEDEDEHHPPDDDHZPHHSYHNHNPPSZs2PZYYYEYP6Y                s26YDENNNNNNEPZYZZZZZZZZPPPPPPPPPPPDDEEDDPPPPDDEEDDDDPPPZEEPDDDEDDDEHPEYYEPE
DZPZPYZPPZZDDPPPPSS66YZYYPPZPS6YPZYZYZZZPPPZYZZPPZYZYYZPPDDDEDPPPPPPYZSY6SPPDPDD6Sc2DDSYYZs2ZYYDcY6YPHc6          scPPNNNNNNNNNNNNHDPYZZZZZZZZZZZZZPZZZZZPZPZZPPPDDDDDPDDPPPPZPPPYYDPDDZPSYPZDPPPYZDPHDE
DZPZPYYZZZZZYZPYYYYZPYZ66666S66SS65SYYZY6YPPPZSZDDPPPPDPZZZZPPPDPPPPDDZPPPPPYYEEPEEEZPYZPH5YYSNNZN:s;s    s;NNNNNNNNNNNNEEDDZZZZYYZPZPPPZZZZPPPZZZZPPPPPPPPDDPPPPPPDPDDHHDDPPZZPPDPEDDEDEPPPZPPPPDDEDEDD
EDHZDZZYYYYZPYYYZPPYS6S66ZPYSYZY6SYSSYZS666SYYZY666SYZDDPDDDDPSYPZEDDDPDPPDDDDHHDEEEDDPDDN:2s;s6        NNNNNNNHDDPZPPPPPPZPZZZPPZZZZZZYZZZYZPPZZZZZPPPPPPPDPDDPPPPPPEHHNHHEEEEEEHEEDDDDDDEDDDDEDHEDDDDD
DPDYPYYPDYYYYYZYYYYPPDDYYYZZZYSS6ZZYYYZZZZZYS65556S652c2566SYYZZPPEDPZPPPPPPPPYY5556SYZP;5 ;;2cZ    NHNNNHDPZPPDPPPPPPPPPPPDPPPPZZZZZZZZZZZYZZZZPPPZPPPZZZZPPDPZPPPHEED5SYYNEPPYZPZEDDDDDPPDDDDPPDPDDDDD
PYZYYPPPDYYDDPPYYYYZZPPZPYYYYYYS6ZPPPPZY6YYZPZYZZZZZPPDPZY666522cc2266YS566666;s66YY2256SZZDHN.;  NNNNDPPPPPPPPDZPZYPPDDDDPPPPPZPPPPZZZPPPZZZZPZZZZZZZZZPPPPPPPPPPPDDDDs522EZDPZZPZDDDDDDDEEEEEPPDDDDDDD
PYYYYPDZZZZYYYYZZPPYYSSSYSSSYZZPPPZSSYZYSYYZPZZZZZPPZZZZZPDPPZZZYY5666SS665655c2ZYZZ66c2c6SD .  NSNHPPPPPZPPPPPPPPZZZZZZPPPPPPPZPPZZZZZPZZZZZPYSYYPDPPPPPPPDPPPPPPPDDYP2665PYHDPPPPEDEEDDDDPDZPDDEDEEDED
PPDPPPPYZPPYYPPYYZZYYZZZZZZPPZZZZPPPZPPPPPZZZPPPDPPPZPEPZZYZDPPPZPPPYYZYSS6YYZYYZY66ZZPHYH.;  NPNNDPPDPPPPPPPPPPPDPPZZZYZZZPPPPPPPPPPPZZZZZZZPZSYSYPZZZSYYZPPPPPPPDEE5Sc5S6S6EPDPZYEPDDPDDDDDPPDPDPEEPDD
PPPPDZPYYZZZZZPZPS6YZZPPPPPPPYYYSYYZZZZPPPZZYZPPPPPPPPDDDDZPPZZZZZPPPPDDEEPPSYSSYY56DEEN;2  NYNNDPPPPPPPDPDDDDDDPPPDZZYYYYZZZZPPPDPPPDPPZZZPPPYSYSSYYZZYYYYSSYYPPDDYP26c25266YSDZDZEPDDDDDDDDEEZPZYEDPPP
PPDZPPPPPZYDDDDPZPDPPZYPPZZZPZYZPZZYSYYZPZSYSYZYYYYYYZPPPPZZYYYZPPZPYZYZPPEEZPZZSYYPPH.s  NPNHDZDPPPDDZZPPPPZPYYZYPDPDDDPPZZYYYYZPZZZYYYZPPDPPY6YZPDPZZYZPPDPSYPPDD2S255266YS56DYDYHDDDDDEDEDEEPPDPDDDDD
DENZPYYZPZZZYYYZZPPZYZZZZZPZZPPZYYYZPZYYYYYZZZPPPPPPPZYYYZPPDDPPZZPPZZSYSSYYYZPDDHDH :  5 NNZYYYPPZZPPPPZZZZYZYSZZZZPPPPZZZPSYYYZZZZZZZZZZY6YDPYYSYPPPPZZZZDDPDDPYZs522YSZY2622Y5HZEDDDZZZYDPDDEEHHEEDDD
ZZPYZZZPPS6DDDDPPDDPPDEZYPPPPZZZPPPZZZZZZPPDEDPPPDDDDPYZZZDPYZDPPPPPZPZPZZ6YSY6Y2Y26  H5NHDYZZPZZZPPPPPPPPZZZZPPPPPPPZPZPPPDPPZYYSZZPPPPYYS6SYS66266SZZPZZZZZPPEP5S2655EPZZc622S2NPDPPPDDDDEDEDDDDDDDDDD
PEHYZZZYZYZYYYZYYZYDEZPZZYSSSYZPDPPPPDEDDDPDDPPZYZDPZZPPZZPPZPPDPPPPYZSYZPYPc52Y:2  . NNZYYYPPPZPPZZYYZZZPZZPPPZPPPDPPPZPPPDPPZYZYZZZPZPYZZZYYYSY6YYZPPDPPPZPDDPZc65565NDDDc6c25sHYHDDPDDEDEDEDEDDDDDDDE
DZPDEYYZPPPZZSSSY52PDYZS6ZPYYS6YZZYZZYSZDDPPPPDDDPPPDDZPZZPPDDDDPPPPZZ6S565Sc6ZE    NPNDYZPPPPPZPZPPPDZPZPZZPPZZZZPPPPPPZZPPZPYSYYZPZZZPZYYYZPYYYSZPPPPPPPPPDEP6Sc222PSHDEE6Ys22sZ5NDPPPPDPDPPPEDDDDDEEE
DYZZPDEDHYYZZPDDDYS656S2255SZYSs.5YS6SSSSYZPDDEDDDDDPPPPPDDDHDDPYYZPZDPP56c62Zs6  P:NHPPPPYYYZZZPPPPPPZZPPZZZPZZZZPPDDDDDDDDZZYSYZZZYYZZZZZZPDPZZYZDPPDDDDPDDDZs55265HZEPDDPDs6226cNPHDDDDDDDDZEPEDDDDDE
6PD666SSPssYY2csc22YY6625s;;:c56652ssccSDY5SZZZYSYPPPPPPZPPPDPEDDDPPSY6ZDEYD;5    NHZYYZYZZPZZZYZZZZZZPPZZPPPDZPSS6666YYSYSSYZPPZYYYZZZPZZZZZZYSSSSYSS65525Y6sc:sssZ6NDPPPPZP;5c2c;6cEZPZYYSSZYDZEDDDDDD
PZPSSPE2SssYY55YPc;22666Y2s252222565222c;s2266YYY622cc2225c5YYDPDZ6Y6Ss2c5;5    NSNHZZZZZPYYZZYYZZZZZZZZZZPPZZYZYYYYYYSS66SSYZZZPPPPZYZZZPPPZYYYZPZYS66SYZZZYc555ccDYEPPPZPPDYPc6555cDSDPPPZZPZDZDPDDDDD
D25YYYD2S52YS66c2::ccc2ssc2c;252cc2222255666Y5c5Y2:ss2S5ss.:sc62PSZZ6ZZZ;s .    NHPPZPPPPPZZPPZZZZPZZZZZPPYZYYPZPPDDDDDDDDPPZZYSYYYYYYYYZPPPPPPDDPZYZPPDDEDZZ5656Y6DYPZDDPPPDDE5Yc25cY6EZPZDPZZZYDZDPDDD
Hc2YZ6Z;c:;2c66YPS62sssc2ssc2cc::;sssccss5YSS6222s::   .;:      Y5EZ.;.:      N2NDYYPPZZZZPPPPZPZZPZZPZZPPZZZYPPPPPPPPPPPPPPZPZZYYYYZPPZPPZZZZZYYYYYZPPPPPYc52526DYDPZZZPZPDDPD5Ys2222cZ6ZYPZZZPYEPDPEDD
DYPYZ566Z6S. :;sc::5S5222csssc556s;. ;sss:...   .;cs;:     .    ;.65sc . .  5 NEPYZZYYZZZZPZPPZPYZPPPPZZPPZZPZDPZZPPPPPPPPPPPPPPPZZYPPPPPPZPPPYYYSYPPDPPZ66c22266HPPPSYYYYYZZYZSZ262552Y5PSYSZYZYPYDPPPD
6SP256Sc2s256662255;;;sccc2s;;s:.       .   .                   . :   ;;    NDDYYYPZPPPPZZDDZPPDPPZZPPPPZYPZPPPPPPPPPPPPPPPDPPZZYYZYZYZZPPPPPZZPZYYYZPPDZ252526PPHDZPSYZYYYSYYYYP6Z6S222cZ6ZYZYZYPYPYDPP
6;s : .25;:ss25: 6Z2cccsc;.:: .   :  :;;;:   :sc2;..              .  .:   N5NEDPPPZZPPPDZZZZZPYZYYZZZYZZZZPZDPDPPPPPPPPDPPPPPPZZZYZZZZPPZZZPPDPPZYYYYPZSSc5256SEPPPSY6SZZZZSYZPPPZDc6252cS2Y6YYY6ZSPYDPP
S::26    55csccssc2. .. .  s6;:::;;::::  .:.                            62NNEPEEDDDDZPZZZZZZPDSY6SYSPYZZZZZZPZPPPPPPDDPPPPPDPPZYZYZZZZPPPPPPPPZZPZPDZYS;22525YSPYYYYYZZYYSSSSSS6SYPs6sc2c5cZS6665ZSDYDPP
5:;YP6Z. ..::2255sss:YY652;6Zcc25cssc.                                  NDZ6PPPZPZDDZZYZPPPPZPZPSYSSDZPZPPPZZZPPPPPZZZYZPPZPZZZZPPPPPPPDPPPZPPZZZZPHZSSs2c225PSSSSYYZPZ6S66SS66556S26;c2c526565S5Y6ZSPYP
Z .  .s .c2    ::.   ::YZ6sSDSY2c56YZ2ss:.                            NDDZZYPP66ZSDPPPDDDDDEZZDDZDSSZSPZDPPPPPPPZZYZPPPDZZYYZPDDDPPPPDPZZYZZYYZDPPPDY5655c2SSDZYYPZYYSYYY6S566655665Sc525522c65S6S5S662Z
62Y  :;  sc:s    s252s.  2S66ss .     ..                            NYNN6SSSZZ66EPPPPDPPPDYZZYPPPDSYZYPZPZPPPPPDZZZZZYYYYYPPZYZZPPDDPPPPPPPPPPPEDEY22s25526DZPY6SSSSYYYYYSS55c222665S26scss2c2c6252652sZ
5.s.:26..:; :    ::  6Yc:26266Y6Y;.                                 NYNHZPZZZZYSNDEEPDZPPDSSPYDDZPYPZZPZPZPZPPPPYZYYYYYZYYZZPPZZZZPPPPPPPPPPPPPPY62scsc2225ZYZZZYYYYYYY6SSS6S25556666c5sc;;cscscs62626cZ
25Y::;s .:;    ;:;:  :.. ;c.   .;s5;.                               NDPPPPZPYYDPHDDEPDZPZZPZPPDPZPZPYYZYPZPPPPZPZPZZDPDDPDPPZZZYZZPPPPPZPPPPZY62cc266666522c226ZZZZYYYYSS666656552522c2c2c2cc2c2ccs65Y6Y
s:c ..:.:;s.:      ::5cSY56c2                                     Y;NPSYSZZPYSZYYYYYYYYP6SPPPZZZPPDDZPZYDPPPPPPPPDZZPZDDDDPPZZZPPPPZZZZZZZYZZPZZPPDEEEHEEEEDPZZZZYYSY6SYYYYSYSS562666SYSS66526565S665656
cs2  :;     .  ;:5cs.s.ZD:  :.:  :;                               NDEPDDPP6YPYHDDDPPDPPDYYZZPZZPZYPPPPYYZZDDPPPPPPZZPZZPZZZZPDPPPZZZZZZZZZZZZPZZPPPPPPPPPDPZPDDZZYYSY6SYYYYSS6S6S55c222252222662266cc526
   ::s2  ..;:: . ssDE6sYZPN5;5Zs.   :                           PcNHYYEDEE6SY6EDDDYYDPPPDDPPDDDDPPPPPPDDDDZZPPPPPPPPDDPPPPDPPPPPPPPPPPPPPPPPPPPPZZYYYPZYZZZYYYS6SSSSSSSSS666656255656sccc222c2c555552526
c:;  ss.:  ;.555c: ;;526S5S::.::;:..:                           NYEPPDEDDDSYZYEDZPZZDPPPDDDDPPDDDPPPPDPPPPPPPPPDZZPZPPPPPPPPPPPPPPPZPPZPZZZPZZZPZZZPZYYYY6666SS6SYYYYYYSSSYSYSS5655552522222265522c65SY6
2;;;;:.s2    ccc:PD2:ZPPPYZYPs.     :                           NEDDYZPPDDPDDDDDYZDPPPDDDDEHPDPPEEDDDDPDDPDPDPPPZPPZPPPPPPDDPPPPPPPPPPPPPPPPPPPPZZZPZYYYYYYYYYYYYZYYYYZZZZZ6S56555525666655556565S6PYYP5
 sc;;::;;  ;.2sc.6S2sSS6Ss;2SSZc:   :                         2.NNDESZZZHEDEDEZZYYPZPPDPDDDDDEPDPPPPPDPDPPDDDDEEDDPPPPPPPPZZPPPPPPPPPPPDPPPZPPPPPPZYYYYSYZZPPPZYYPZYYYYZZZZ565666SS6S665555S6555222666Yc
.    ;s;:55:.  SSYY66;.6Z66   c;ssc:..:                       NYNDDD6YZZDPPDYP56ZYYYPPZZZZDDZPYYPZYYPDYZ66YYPPDDDDPPPPPDPDPPZPZZYSYYZPPPPPPZZYZPPPZZS565566YYYS55666625YSSY6Y25225555c222sscs52666Ys226:
266SS;:::cs56; 52YZPEYS2c2Y56ss    .:.                      ;2NNDPZPYPDPHEEHSZPZEEPPDDPDZZEEZPZYEPHHPDPPZPYYYYZYPPDDDDPDPDZPYYS6YYPDPZZZPPPPPZZPZZYYYSS56Y656SS6SZZZZZZSY6SSYc2c22222556655YSYSPP6Zc6:s;
ss:S6ccc2;.SY2;YY2cssccSYSY6S55ss                           NZNEPDDEPDPPEDZDYZEDEDPPPPDDPDZPYPS6DYDDYZZPZPYZYYYSYYZYZPZYZPPPYYYYZZPDPPPPPZPPPDDDDDDDDPDDDDDDDPDDDPPPPDDYZYZYY5S666666665566665656s2;2scs
:s:62Y666cs6666YZ2s6S6555c2;;;sc5s;;::..                   ;NHNDDEDDYPPPHEDEYZEDEEDDDDPDZPZPYZSSDZDDYYDPDEYP6SYSZYZZZZYYYYYZYZPPPPDDPPPPZYZPPDPPPZZPPPPDDPPZZZPDPZZZZYZSY6656c265655552656625c2s2s2c2s2s
:    6sPP2cS666:.s2c2;:6S22ss;;   .  ::.                   ;NHNDDDPDYZPPHEEEPDYYDPPDDDPPYYDDYPZZPZPPZZDPEHPDZPZZYYZZPPPPPPPDYZY6YSZZZYS6YYYYYYYYS6SZYSSSSZYZZZZYZPZYY56555522sc2c65225566665626c5c2c5:s;
.    ; DPcc552cSY2c26..;ss2:;  :;..    .:.                  :2NNNDDEZDDPZZHEHHDDPPZZZZYZDDPDPEYZPZPPZYPPDDDDZZPPZZZZZPZZZPYZSY65655566SSSSSYYYYYYYYYSSY6SYYYYYYYSSS6S2522666622526625656566266Yc222c5;s;
2  ss: :.c2:.;:SZ222525: :s;s                             Es  c2NNEDPDSYDPEDDDDDDDDDDDEDDDPDYPYYPZEDDDPZZPZZZZPDPPZPYYSYYZSZ55555522225555555256526S6662552556666662222cc2255c252S65566S6YYSZs2;s22sccc2
Ds5  .   ;;ss2css2S:::s.:.. :..::::                      sNNZ2  HZNNZPYZPPEDDDDDDDPPDDPPPDPDYPYYYSPZPDPPZZYYPDZPZZPPZZZZSY6Y6S65666S6652552cccs;scc2s;;.:;s2csssssssscccc22cc2552666622cc55c6.s;;ss:;;;s
Y6P;c::sc;;s:cc;;;sscs2.:::                             65NNEP ;  NNNEPDDPHEEEDDEEDDEEDDDDDEPEYZDPDDPPPPZPPPDDDDPPPDZPYZYPYPYZYYYZSS55csssss;:;;:...::.  ..::;:;:::;;css;;;ccsscs5cssss;;2c:s.:sss;ccs;s
626c526:;    sc26  ;s . :  :;:;..                       NZNNNNNN  .:NNEDDDHEEEEEEEEEEEEEDDDDENPPHPHHPDPDPDPDZZYYPPZPYY6S6S25c2552csc:..:..                     ........::s;;:..;:csss;;ssssscccsss;csccc
6256S5626sc::sss2ssss:s .                              :NHNHNHDNs6  HYNHPPHEHHEEEEEEDEDDDDEHZDSSEPNNDHPPEHPD56SSZPYY6Y25sc:s :  .....:..                                       . :.s;::sss;::;:s;s;s;s;2
6SY25c2566Y22:::s  :;s2.:::                           66NNHHHEPD;6 .  NHHDEDHHDDDPEEEEDDDDDEZDZZEPHEDEPDPPZPYPYY6S6S6S25266Yc2.    .                   Z;                          :.:.. ;:ss:.s:5c2c526
S6S6Sc2c2222625::.:ss;;:s;s:...                       NPNHEEHNNNENSZ  5cNHDZPPPPDDDDDDDDDDDEDEPPPZPZEEZPSYPDPEYPYPYZZPPEZPYZSZ55:                   2NNPHNP                            . . :::.: s.s;s:2
YSYSYYZ666SSS6SYZ55ssss:;                          ;5YNNHEHHHHHNNNDN56  P6HEEEEEEDDDDDPPPPPPZPPPPPZZZZZDYYYY6S6S25sc;;:; .                      s5DN6     cND::                                . ;.s;s;c

*/

func (mr *MediaRouter) tryThanks(text string) bool {
	// Создаем регулярное выражение для поиска ключевых фраз
	regex := regexp.MustCompile(`(?i)\b(thanks|thank you|thx|tq|thanks a lot|thanks a bunch|❤)\b`)

	// Используем FindStringSubmatch для поиска соответствий
	matches := regex.FindStringSubmatch(text)

	// Если найдено соответствие, возвращаем true
	return len(matches) > 0
}
