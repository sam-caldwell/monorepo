keybox
======

```text
 kbxutil --verbose ~/.gnupg/pubring.kbx --help
BEGIN-RECORD: 0
Length: 32
Type:   Header
Version: 1
Flags:   0002 (openpgp)
created-at: 1691555578
last-maint: 1691555578
END-RECORD


BEGIN-RECORD: 1
Length: 1945
Type:   OpenPGP
Version: 1
Blob-Flags: 0000
Data-Offset: 126
Data-Length: 1799
Unhashed: 20
Key-Count: 2
Key-Info-Length: 28
Key-Fpr[0]: 65A6E0F4356ACEFCEAFD50E567B706124DFAABFF
Key-Kid-Off[0]: 32
Key-Kid[0]: 67B706124DFAABFF
Key-Flags[0]: 0000
Key-Fpr[1]: 3EB7AADA68E1A9BC5412B09385E098664FA683F3
Key-Kid-Off[1]: 60
Key-Kid[1]: 85E098664FA683F3
Key-Flags[1]: 0000
Serial-No: none
Uid-Count: 1
Uid-Info-Length: 12
Uid-Off[0]: 542
Uid-Len[0]: 35
Uid[0]: "Sam Caldwell <mail@samcaldwell.net>"
Uid-Flags[0]: 0000
Uid-Validity[0]: 0
Sig-Count: 2
Sig-Info-Length: 4
Sig-Expire[0-1]: [not checked]
Ownertrust: 0
All-Validity: 0
Recheck-After: 0
Latest-Timestamp: 0
Created-At: 1713308579
Reserved-Space: 0
Checksum: c239031eb0be21b96e48d6bfb0b5b6437ebd89b4 [valid]
END-RECORD


BEGIN-RECORD: 2
Length: 1945
Type:   OpenPGP
Version: 1
Blob-Flags: 0000
Data-Offset: 126
Data-Length: 1799
Unhashed: 20
Key-Count: 2
Key-Info-Length: 28
Key-Fpr[0]: 7D3961E58C175A9AC2503C98AD1B3C6062CD8118
Key-Kid-Off[0]: 32
Key-Kid[0]: AD1B3C6062CD8118
Key-Flags[0]: 0000
Key-Fpr[1]: 12CC3E6EED8B97B78B9FE1412EEB730EB2FC8027
Key-Kid-Off[1]: 60
Key-Kid[1]: 2EEB730EB2FC8027
Key-Flags[1]: 0000
Serial-No: none
Uid-Count: 1
Uid-Info-Length: 12
Uid-Off[0]: 542
Uid-Len[0]: 35
Uid[0]: "Sam Caldwell <mail@samcaldwell.net>"
Uid-Flags[0]: 0000
Uid-Validity[0]: 0
Sig-Count: 2
Sig-Info-Length: 4
Sig-Expire[0-1]: [not checked]
Ownertrust: 0
All-Validity: 0
Recheck-After: 0
Latest-Timestamp: 0
Created-At: 1716589528
Reserved-Space: 0
Checksum: 703746624b1aaf7ac6eb87b492736f46a6cc359b [valid]
END-RECORD


BEGIN-RECORD: 3
Length: 849
Type:   OpenPGP
Version: 1
Blob-Flags: 0000
Data-Offset: 110
Data-Length: 719
Unhashed: 20
Key-Count: 1
Key-Info-Length: 28
Key-Fpr[0]: E8780D3B1C24348632AB4BF11737E9CE68965A60
Key-Kid-Off[0]: 32
Key-Kid[0]: 1737E9CE68965A60
Key-Flags[0]: 0000
Serial-No: none
Uid-Count: 2
Uid-Info-Length: 12
Uid-Off[0]: 275
Uid-Len[0]: 35
Uid[0]: "Sam Caldwell <mail@samcaldwell.net>"
Uid-Flags[0]: 0000
Uid-Validity[0]: 0
Uid-Off[1]: 553
Uid-Len[1]: 37
Uid[1]: "Sam Caldwell <github@samcaldwell.net>"
Uid-Flags[1]: 0000
Uid-Validity[1]: 0
Sig-Count: 2
Sig-Info-Length: 4
Sig-Expire[0-1]: [not checked]
Ownertrust: 0
All-Validity: 0
Recheck-After: 0
Latest-Timestamp: 0
Created-At: 1716659663
Reserved-Space: 0
Checksum: 1859198ecfe8827a0a409349bda4225350bcad6d [valid]
END-RECORD
```

```text
hexdump -C ~/.gnupg/pubring.kbx
00000000  00 00 00 20 01 01 00 02  4b 42 58 66 00 00 00 00  |... ....KBXf....|
00000010  64 d3 16 fa 64 d3 16 fa  00 00 00 00 00 00 00 00  |d...d...........|
00000020  00 00 07 99 02 01 00 00  00 00 00 7e 00 00 07 07  |...........~....|
00000030  00 02 00 1c 65 a6 e0 f4  35 6a ce fc ea fd 50 e5  |....e...5j....P.|
00000040  67 b7 06 12 4d fa ab ff  00 00 00 20 00 00 00 00  |g...M...... ....|
00000050  3e b7 aa da 68 e1 a9 bc  54 12 b0 93 85 e0 98 66  |>...h...T......f|
00000060  4f a6 83 f3 00 00 00 3c  00 00 00 00 00 00 00 01  |O......<........|
00000070  00 0c 00 00 02 1e 00 00  00 23 00 00 00 00 00 02  |.........#......|
00000080  00 04 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000090  00 00 00 00 00 00 66 1f  03 a3 00 00 00 00 99 01  |......f.........|
000000a0  8d 04 66 1f 03 97 01 0c  00 99 e3 30 39 0a e0 43  |..f........09..C|
000000b0  4a ae 8e 13 27 59 82 36  d3 a9 f5 66 3e 2b 03 bc  |J...'Y.6...f>+..|
000000c0  90 55 40 33 f4 6d 13 71  50 2a 74 48 92 a8 b6 d2  |.U@3.m.qP*tH....|
000000d0  0e b2 83 a7 21 f2 0a cf  a2 b1 83 ff b5 c8 32 48  |....!.........2H|
000000e0  62 1d f1 60 d8 2c 88 22  38 cc e7 1b 3b 6e 97 5c  |b..`.,."8...;n.\|
000000f0  15 f0 63 d1 1c e9 99 c2  64 bb 69 6e 54 6f 5c b0  |..c.....d.inTo\.|
00000100  c6 78 df 5e fa 69 7c 9b  e6 94 5a 02 3c 3f 4d 1a  |.x.^.i|...Z.<?M.|
00000110  16 74 65 4b a0 4b 04 46  1d bf 2f 6e 24 fc 8d 1b  |.teK.K.F../n$...|
00000120  64 93 06 78 62 f4 04 d2  91 55 4b f9 4a f2 b0 d9  |d..xb....UK.J...|
00000130  d9 54 0d 83 c1 7d a2 35  0c 26 e4 b2 ea 5d f2 67  |.T...}.5.&...].g|
00000140  87 da e9 b2 c1 d0 90 ea  d2 ab 75 43 93 8e c2 46  |..........uC...F|
00000150  7f 89 24 04 19 bb 6f 43  83 5a 7c b0 ce 91 d3 0b  |..$...oC.Z|.....|
00000160  37 3a 15 4d 3e 63 c8 ce  90 d3 8e 75 91 5f f8 53  |7:.M>c.....u._.S|
00000170  0c 46 a1 36 99 65 49 0b  7b 90 dc a2 f4 44 32 32  |.F.6.eI.{....D22|
00000180  75 c1 2c 3c a2 31 62 c5  fb 79 2f b2 e1 84 e4 9f  |u.,<.1b..y/.....|
00000190  ea 7b 7e 57 3f 25 8e 4c  6b 31 93 d2 b9 c0 07 bf  |.{~W?%.Lk1......|
000001a0  5b 53 df a9 87 de 72 1e  6f 24 da f6 8d 08 3e 30  |[S....r.o$....>0|
000001b0  fb bd d0 3f 75 f3 6b 2a  30 6a 96 32 a8 b8 47 4c  |...?u.k*0j.2..GL|
000001c0  43 d3 9b 6e 80 55 f3 93  9d 7d 31 0d 9a 90 ce 20  |C..n.U...}1.... |
000001d0  5f 71 01 b9 b7 89 57 b6  83 96 0e da 10 cd 55 fe  |_q....W.......U.|
000001e0  ec d6 a1 a5 96 46 10 9b  91 69 c4 1b 41 35 df 4a  |.....F...i..A5.J|
000001f0  40 3e 42 be 26 db 92 e3  76 52 ac e1 30 51 98 a9  |@>B.&...vR..0Q..|
00000200  9e ba 7d ca 3f 8c 13 f7  8a 28 de 8e 5d 21 50 63  |..}.?....(..]!Pc|
00000210  96 d9 07 51 80 ee 85 1f  be 04 85 f7 c2 b2 d2 8c  |...Q............|
00000220  e0 6f 5e 6a c6 86 bc a4  7b 00 11 01 00 01 b0 0c  |.o^j....{.......|
00000230  00 00 67 70 67 01 00 00  00 00 00 00 b4 23 53 61  |..gpg........#Sa|
00000240  6d 20 43 61 6c 64 77 65  6c 6c 20 3c 6d 61 69 6c  |m Caldwell <mail|
00000250  40 73 61 6d 63 61 6c 64  77 65 6c 6c 2e 6e 65 74  |@samcaldwell.net|
00000260  3e b0 0c 00 00 67 70 67  02 00 00 00 00 00 00 89  |>....gpg........|
00000270  01 d4 04 13 01 0a 00 3e  16 21 04 65 a6 e0 f4 35  |.......>.!.e...5|
00000280  6a ce fc ea fd 50 e5 67  b7 06 12 4d fa ab ff 05  |j....P.g...M....|
00000290  02 66 1f 03 97 02 1b 03  05 09 03 c2 67 00 05 0b  |.f..........g...|
000002a0  09 08 07 02 06 15 0a 09  08 0b 02 04 16 02 03 01  |................|
000002b0  02 1e 01 02 17 80 00 0a  09 10 67 b7 06 12 4d fa  |..........g...M.|
000002c0  ab ff 63 5a 0b ff 58 7e  43 a7 52 5a b2 21 f9 8a  |..cZ..X~C.RZ.!..|
000002d0  c6 8a 5d ed 6d e6 fc 05  b7 aa c4 b4 ca 95 58 7a  |..].m.........Xz|
000002e0  1b 98 03 29 d3 82 10 61  1a bf 93 1b 6e d6 31 e6  |...)...a....n.1.|
000002f0  be 1c a2 7e f7 72 51 41  12 1e 0e b7 9a 70 d4 a7  |...~.rQA.....p..|
00000300  e3 89 cd 33 bc df 60 dc  29 88 52 e4 3a f5 96 fb  |...3..`.).R.:...|
00000310  89 7f ff c8 55 e5 5a 8b  05 8b 61 e2 ff 09 6c 7b  |....U.Z...a...l{|
00000320  7c 50 b0 44 5b e6 59 35  4d a3 b5 d8 92 e7 ca 7b  ||P.D[.Y5M......{|
00000330  55 52 d9 b8 5e 05 a4 93  e7 2e e7 0c cf 7c 50 62  |UR..^........|Pb|
00000340  49 75 bb de c5 bb dd 72  0c 53 3d e5 ba 79 81 3d  |Iu.....r.S=..y.=|
00000350  18 e0 32 fc f1 15 ac e1  d5 1f 70 26 67 56 63 25  |..2.......p&gVc%|
00000360  4d d3 67 25 93 d1 6f 38  48 ae 03 a6 37 1b 83 c0  |M.g%..o8H...7...|
00000370  84 d6 e0 d9 1f d8 db b2  9f bd 94 ca d4 f2 e2 62  |...............b|
00000380  9e e5 13 11 27 7a 04 37  46 29 57 bd c0 15 30 af  |....'z.7F)W...0.|
00000390  82 da 93 75 e2 68 c2 ea  7f 00 06 da d3 bb b7 11  |...u.h..........|
000003a0  17 94 07 28 51 8f 5b bd  b3 16 38 8c c0 b4 49 8f  |...(Q.[...8...I.|
000003b0  9a cd 94 f5 09 9b 0f 3f  ed f9 30 f6 55 6c 3d 49  |.......?..0.Ul=I|
000003c0  ce fd 14 d7 ad 07 0f f1  47 0b 08 8a d2 62 db df  |........G....b..|
000003d0  26 5d 97 f4 3a c4 10 00  ed 7d 37 71 0b b7 e0 59  |&]..:....}7q...Y|
000003e0  90 da 82 b4 96 46 df fb  f0 6a 58 fd e1 fe 0e e7  |.....F...jX.....|
000003f0  08 07 b7 78 2f d4 8d 56  a1 7e 5c 8b 2f 12 2d 46  |...x/..V.~\./.-F|
00000400  e5 53 11 aa 12 db 05 c3  77 a3 d9 a5 10 37 5e 93  |.S......w....7^.|
00000410  eb d9 04 c5 c2 39 80 a3  34 1e b0 d6 c3 71 56 40  |.....9..4....qV@|
00000420  2a 58 fc 9d 60 d7 b6 2d  2f 4e 3f dc 37 70 87 db  |*X..`..-/N?.7p..|
00000430  d4 9c e7 5d cb 2c f8 8a  bd 51 9b e5 a6 06 97 e8  |...].,...Q......|
00000440  ed 95 ad 37 ef 20 b0 06  00 00 67 70 67 00 b9 01  |...7. ....gpg...|
00000450  8d 04 66 1f 03 97 01 0c  00 d0 0a 0d 34 32 89 ec  |..f.........42..|
00000460  40 1d 35 23 e6 37 85 d1  c7 04 19 c7 f8 36 88 e9  |@.5#.7.......6..|
00000470  f4 1b 64 ec a9 6d a5 ce  ea 87 ab 09 20 04 dd 84  |..d..m...... ...|
00000480  b5 49 6d 0f f6 19 c5 5e  d4 c5 c5 3a d3 26 36 39  |.Im....^...:.&69|
00000490  8c ff 6c 20 95 df 06 e3  b5 43 2b eb dc 25 21 53  |..l .....C+..%!S|
000004a0  e5 88 fc 31 cb a0 d7 71  d0 73 10 1f 5e 8c 2b 60  |...1...q.s..^.+`|
000004b0  a0 1e 77 94 da 69 5e 2b  50 9e 17 ef 6c 55 9e 65  |..w..i^+P...lU.e|
000004c0  f4 43 3a 29 12 7d 67 67  a7 b0 62 d5 9a 62 07 c4  |.C:).}gg..b..b..|
000004d0  c9 fb 7d 72 2d 6b 21 20  9c 75 2d 98 54 b2 dd 40  |..}r-k! .u-.T..@|
000004e0  ec 4e 53 b3 c2 1c 63 4e  7e 0c 46 a5 39 e5 0c 7a  |.NS...cN~.F.9..z|
000004f0  4a 89 3d a0 71 77 f8 3b  ed b8 5b e1 c4 60 2b 63  |J.=.qw.;..[..`+c|
00000500  35 a7 80 9f e4 bc 05 be  a4 5d c9 9b 46 4b 48 77  |5........]..FKHw|
00000510  dd 6c b9 c9 d2 81 8a a8  41 d8 1f b3 cd 94 a9 fe  |.l......A.......|
00000520  97 cf 09 c6 0a 47 7e 19  05 5d d1 8d fe b1 6e f7  |.....G~..]....n.|
00000530  ec f1 19 68 00 89 e7 4c  e2 20 1a e3 55 c5 41 da  |...h...L. ..U.A.|
00000540  48 fb 3a d6 1e af 54 4e  e5 30 b2 8f cd ce 06 9c  |H.:...TN.0......|
00000550  1d f2 0b 60 ab c9 00 eb  4b 83 dd 13 31 a3 0c 77  |...`....K...1..w|
00000560  77 58 58 0f d4 d9 76 26  29 bd bb b4 92 53 ac 43  |wXX...v&)....S.C|
00000570  32 fa 6b 36 22 14 6c 6d  6d 92 f3 f9 f8 91 af a1  |2.k6".lmm.......|
00000580  18 8f 98 59 fb bc f1 2f  0d dd d0 1e 89 25 c4 09  |...Y.../.....%..|
00000590  35 db 93 5e 29 10 97 58  10 3c e0 ac db 22 f3 d0  |5..^)..X.<..."..|
000005a0  e2 43 cc a0 1e 2e 9f 5a  82 11 4c cb 29 77 2d 31  |.C.....Z..L.)w-1|
000005b0  23 15 52 65 aa e8 f9 01  5e ef 2b 5e 13 49 28 6d  |#.Re....^.+^.I(m|
000005c0  ad 31 4d 0c 13 a9 9f 3a  14 28 cd 6f 7b 28 5a 5e  |.1M....:.(.o{(Z^|
000005d0  e8 32 0b 58 d5 00 40 cd  f5 00 11 01 00 01 89 01  |.2.X..@.........|
000005e0  bc 04 18 01 0a 00 26 16  21 04 65 a6 e0 f4 35 6a  |......&.!.e...5j|
000005f0  ce fc ea fd 50 e5 67 b7  06 12 4d fa ab ff 05 02  |....P.g...M.....|
00000600  66 1f 03 97 02 1b 0c 05  09 03 c2 67 00 00 0a 09  |f..........g....|
00000610  10 67 b7 06 12 4d fa ab  ff 79 1b 0b ff 77 8e 03  |.g...M...y...w..|
00000620  44 16 1e ab 54 d4 56 79  dd e4 1b 05 39 28 86 8b  |D...T.Vy....9(..|
00000630  21 3c df 90 d4 ae 97 47  36 bf 7f 57 53 4a 2f 02  |!<.....G6..WSJ/.|
00000640  90 da a9 36 4e 0d 1d 3e  ce 36 bc 98 52 aa dc a5  |...6N..>.6..R...|
00000650  cd 18 1d a3 f2 e6 28 21  c9 4f b6 00 65 b7 1d 6b  |......(!.O..e..k|
00000660  d4 52 49 71 41 4b 18 80  47 c2 4f 2f 79 64 31 a6  |.RIqAK..G.O/yd1.|
00000670  41 f6 a7 1b 97 d9 64 15  c8 bc 05 b8 65 48 22 33  |A.....d.....eH"3|
00000680  eb ac d9 fc ad ac 26 6d  4a e8 8e 4d 4b fd 83 e4  |......&mJ..MK...|
00000690  ca b3 4c d7 0d 78 c7 35  2f 98 0a ae a8 d9 66 32  |..L..x.5/.....f2|
000006a0  d6 07 af c6 90 bd 6d fb  d1 da 23 a9 72 e6 dd 7a  |......m...#.r..z|
000006b0  71 59 97 b7 34 90 6a 89  c9 ac e7 f2 8e d7 d7 e2  |qY..4.j.........|
000006c0  dd 66 5a 37 8b 83 93 1c  1c 37 de 2c 58 f8 39 23  |.fZ7.....7.,X.9#|
000006d0  da 66 69 c1 83 d7 d9 85  77 19 61 23 0b bb f9 66  |.fi.....w.a#...f|
000006e0  7f f0 24 4f f9 72 c5 84  2d 17 79 17 bc 97 8d 01  |..$O.r..-.y.....|
000006f0  af ef 32 22 62 ac 8d c2  1f 75 84 03 13 3a 26 57  |..2"b....u...:&W|
00000700  20 bd 76 11 6f 55 b2 ec  80 84 ad 34 8a 45 77 ef  | .v.oU.....4.Ew.|
00000710  ac f5 c3 49 82 1c d4 dd  2a a4 4a ef fb fb fe e3  |...I....*.J.....|
00000720  68 d6 ef eb f5 6d bc 6e  6c f9 0a 06 7f 18 b5 5a  |h....m.nl......Z|
00000730  5a 56 05 2d bd ca 26 d9  e5 79 cd e8 a4 1b 00 a8  |ZV.-..&..y......|
00000740  c3 fe fa b5 6c c3 3d d1  88 f0 cf e4 71 44 1b 6f  |....l.=.....qD.o|
00000750  40 7a dc 9c 2f 53 25 1c  0c 9c 4e 0d ce 7c c9 b6  |@z../S%...N..|..|
00000760  b4 ff 4b 1b c5 01 d3 50  b5 18 67 b9 85 29 40 67  |..K....P..g..)@g|
00000770  cc f9 de fd c6 50 dc 88  98 2d 54 0f d9 c0 bf ab  |.....P...-T.....|
00000780  43 cc 65 98 2c 8b c9 a5  51 b6 30 0f d3 f5 17 8e  |C.e.,...Q.0.....|
00000790  fd 69 75 8b 5e 95 66 78  fa c0 68 2e e0 b0 06 00  |.iu.^.fx..h.....|
000007a0  00 67 70 67 00 c2 39 03  1e b0 be 21 b9 6e 48 d6  |.gpg..9....!.nH.|
000007b0  bf b0 b5 b6 43 7e bd 89  b4 00 00 07 99 02 01 00  |....C~..........|
000007c0  00 00 00 00 7e 00 00 07  07 00 02 00 1c 7d 39 61  |....~........}9a|
000007d0  e5 8c 17 5a 9a c2 50 3c  98 ad 1b 3c 60 62 cd 81  |...Z..P<...<`b..|
000007e0  18 00 00 00 20 00 00 00  00 12 cc 3e 6e ed 8b 97  |.... ......>n...|
000007f0  b7 8b 9f e1 41 2e eb 73  0e b2 fc 80 27 00 00 00  |....A..s....'...|
00000800  3c 00 00 00 00 00 00 00  01 00 0c 00 00 02 1e 00  |<...............|
00000810  00 00 23 00 00 00 00 00  02 00 04 00 00 00 00 00  |..#.............|
00000820  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 66  |...............f|
00000830  51 13 d8 00 00 00 00 99  01 8d 04 66 51 13 bf 01  |Q..........fQ...|
00000840  0c 00 cb a7 84 40 7d 3b  be aa bb df 09 17 2e d1  |.....@};........|
00000850  54 d9 00 49 7e d7 06 34  9f 9e 34 83 3d 61 75 e3  |T..I~..4..4.=au.|
00000860  20 bd 21 a5 3d be 86 d1  af 04 1d ee 02 c0 bb d2  | .!.=...........|
00000870  62 f7 48 9a 23 78 40 c9  08 94 12 ab 67 08 b6 dc  |b.H.#x@.....g...|
00000880  9b e4 56 59 a9 fe 67 ca  fe e4 cf 47 57 db 81 f0  |..VY..g....GW...|
00000890  a7 33 9b 5a b3 17 48 70  99 e9 d5 49 85 5d 44 ac  |.3.Z..Hp...I.]D.|
000008a0  9e e5 c3 73 68 46 c6 87  6d ab 1f 5f 35 62 05 38  |...shF..m.._5b.8|
000008b0  4b 90 02 be 46 b1 ab 40  d7 f6 2e 21 f6 d1 8e 4c  |K...F..@...!...L|
000008c0  74 5c 21 f6 69 fd f4 2d  e6 88 32 da e6 99 18 43  |t\!.i..-..2....C|
000008d0  37 e5 ae 68 40 32 44 3a  a5 5f 96 e1 dd 59 9f 4c  |7..h@2D:._...Y.L|
000008e0  95 9b b9 7a ec fe f3 e9  ab ce df ff a9 70 9c 3a  |...z.........p.:|
000008f0  be 95 8d 3e 82 98 d7 44  a2 d7 cb 70 ad c1 3c 25  |...>...D...p..<%|
00000900  c5 7e d3 8f 0d ec 40 6c  65 c2 ed 63 f1 b6 d3 b9  |.~....@le..c....|
00000910  28 ef 4b 54 96 90 e6 fc  74 63 ec 87 ad e2 25 3a  |(.KT....tc....%:|
00000920  8a 94 76 1d a7 c8 c7 ec  37 47 18 d9 01 24 07 a7  |..v.....7G...$..|
00000930  36 ca 4d 60 af 5f 53 b5  6a 27 20 9e 39 ef 30 cb  |6.M`._S.j' .9.0.|
00000940  9a 7b e8 4d 68 d1 fa ba  50 61 c3 9d 7c 5e b4 39  |.{.Mh...Pa..|^.9|
00000950  c6 0c cb e9 97 45 93 0d  0e 17 2a 73 c5 5f f4 18  |.....E....*s._..|
00000960  5a 13 e6 63 2b fe bc 30  84 d0 7c a3 cf 23 9a db  |Z..c+..0..|..#..|
00000970  e5 e3 f6 97 ca 81 3a 84  b7 db c0 20 72 fe 9e bc  |......:.... r...|
00000980  13 94 1a f5 17 23 8d 6b  d0 3a f3 6e 5b d5 b7 79  |.....#.k.:.n[..y|
00000990  0e bf ca 12 49 8d 92 d7  1c ee 6f f0 fe c7 94 ac  |....I.....o.....|
000009a0  15 a4 26 2b cb 6e 0f 77  8b e7 ec f2 8f b4 d4 88  |..&+.n.w........|
000009b0  b8 a2 f7 96 93 d6 dc 1b  ed a8 ad fd 5c 1f a8 14  |............\...|
000009c0  64 2f 00 11 01 00 01 b0  0c 00 00 67 70 67 01 00  |d/.........gpg..|
000009d0  00 00 00 00 00 b4 23 53  61 6d 20 43 61 6c 64 77  |......#Sam Caldw|
000009e0  65 6c 6c 20 3c 6d 61 69  6c 40 73 61 6d 63 61 6c  |ell <mail@samcal|
000009f0  64 77 65 6c 6c 2e 6e 65  74 3e b0 0c 00 00 67 70  |dwell.net>....gp|
00000a00  67 02 00 00 00 00 00 00  89 01 d4 04 13 01 0a 00  |g...............|
00000a10  3e 16 21 04 7d 39 61 e5  8c 17 5a 9a c2 50 3c 98  |>.!.}9a...Z..P<.|
00000a20  ad 1b 3c 60 62 cd 81 18  05 02 66 51 13 bf 02 1b  |..<`b.....fQ....|
00000a30  03 05 09 03 c2 67 00 05  0b 09 08 07 02 06 15 0a  |.....g..........|
00000a40  09 08 0b 02 04 16 02 03  01 02 1e 01 02 17 80 00  |................|
00000a50  0a 09 10 ad 1b 3c 60 62  cd 81 18 99 24 0b fe 28  |.....<`b....$..(|
00000a60  a3 01 0c cc 8d aa 0f 62  56 d1 0e 3c 77 bf 49 45  |.......bV..<w.IE|
00000a70  ba 5d 84 98 57 a7 4d 22  e7 29 d2 90 9d e4 d8 d4  |.]..W.M".)......|
00000a80  b7 4d b5 95 61 5c 27 99  c0 c6 d4 4c 45 02 80 4a  |.M..a\'....LE..J|
00000a90  f6 56 09 82 d3 50 e3 d6  92 73 a5 c7 f8 b9 dd 4b  |.V...P...s.....K|
00000aa0  48 13 b7 5d de 93 77 9d  2b a5 cf 24 29 59 f7 52  |H..]..w.+..$)Y.R|
00000ab0  db 84 9b 0d e5 b0 63 af  e0 4c d8 de 1a f4 14 10  |......c..L......|
00000ac0  f2 7d ad 09 f2 31 7d de  19 82 14 b4 8c f4 bd 76  |.}...1}........v|
00000ad0  b6 ce 1d 6a 25 06 f7 3d  77 99 f7 7e cb 38 96 5f  |...j%..=w..~.8._|
00000ae0  02 3d 4d a2 c1 2a e4 17  ad dc e3 65 2b 6c 81 af  |.=M..*.....e+l..|
00000af0  5c 84 72 8d 38 60 fe 09  a0 1a 7b 4e 07 1b 23 2f  |\.r.8`....{N..#/|
00000b00  f1 e6 df da 0b e2 60 f8  da 97 a4 25 a2 ee 44 f3  |......`....%..D.|
00000b10  8c 99 72 dd ad 78 44 f7  13 23 0f 60 d0 7a 4f 45  |..r..xD..#.`.zOE|
00000b20  6a 4f 3b bb 75 41 51 a0  27 d7 b3 23 3f b4 df 41  |jO;.uAQ.'..#?..A|
00000b30  32 0f 89 f6 60 21 26 73  71 5b 8c bd 34 21 24 e7  |2...`!&sq[..4!$.|
00000b40  8a bc 8e 5c 23 5f a2 d2  3d 67 d6 6c d7 23 fd cd  |...\#_..=g.l.#..|
00000b50  ad ed d8 29 6c 23 0c c7  92 75 38 e5 af 23 2f 7a  |...)l#...u8..#/z|
00000b60  93 41 06 51 63 15 39 4f  ac 15 ec 39 a3 cb a7 49  |.A.Qc.9O...9...I|
00000b70  56 f7 9d 87 fb 5e a9 19  a8 12 e1 03 57 d4 62 1e  |V....^......W.b.|
00000b80  06 7e dd 33 bf 98 bd d1  08 3f b7 6d bf 7b 83 07  |.~.3.....?.m.{..|
00000b90  19 79 61 cc f5 b9 f6 8f  fd 49 e7 4d 76 70 3e 71  |.ya......I.Mvp>q|
00000ba0  9a f3 21 5f 55 a3 6e 1d  23 95 76 25 b2 fd b0 c3  |..!_U.n.#.v%....|
00000bb0  45 df 63 24 3c a9 d6 bd  48 6e b8 a9 d5 33 7b e8  |E.c$<...Hn...3{.|
00000bc0  51 d6 c1 73 e7 46 d0 48  3e a5 2f c7 a6 99 61 01  |Q..s.F.H>./...a.|
00000bd0  4a ed c2 62 a5 39 86 b7  70 36 79 5c 9e 98 f0 b0  |J..b.9..p6y\....|
00000be0  06 00 00 67 70 67 00 b9  01 8d 04 66 51 13 bf 01  |...gpg.....fQ...|
00000bf0  0c 00 b5 12 c5 21 d9 34  36 e2 eb dc 1a 6f 8b 87  |.....!.46....o..|
00000c00  27 d3 33 6e 33 57 06 53  da b0 ed fe f6 a2 fb 27  |'.3n3W.S.......'|
00000c10  da ec 63 79 69 1d 2c 75  e7 c9 6b b4 4a a3 1d 2c  |..cyi.,u..k.J..,|
00000c20  7d f2 f7 1f e3 2e 1c e7  ca 60 21 63 77 63 20 05  |}........`!cwc .|
00000c30  80 45 29 3e 18 84 80 cc  db 01 52 31 1c 4d a2 70  |.E)>......R1.M.p|
00000c40  7a 96 f7 13 fd c1 ec 22  7e b4 9c e2 59 76 ef bf  |z......"~...Yv..|
00000c50  9c a6 77 e9 ec 2e f5 ba  20 98 64 c5 ac f5 f2 62  |..w..... .d....b|
00000c60  10 63 e1 95 3b 4e 7c 65  f9 b0 3b 51 b1 c5 2d 24  |.c..;N|e..;Q..-$|
00000c70  69 45 e8 28 c0 2d a4 43  fd 17 f1 81 2d e9 dd 6b  |iE.(.-.C....-..k|
00000c80  ad 32 ca d9 7c 03 15 4c  36 09 b5 61 b7 4b 0d 2f  |.2..|..L6..a.K./|
00000c90  52 c9 4e 53 87 13 77 66  59 40 33 3d ac f4 1b a2  |R.NS..wfY@3=....|
00000ca0  4c 87 a8 7c 41 1c 13 c8  73 5a 02 f5 c6 85 bc c4  |L..|A...sZ......|
00000cb0  d6 ed 2d 30 7e a2 61 cb  97 fa 32 6f b7 24 8b 89  |..-0~.a...2o.$..|
00000cc0  99 ca ec 3a 56 e0 8e c8  19 fe 9d 35 63 0c 2f 90  |...:V......5c./.|
00000cd0  7c a4 71 16 90 41 29 6c  b2 e7 3b be 5c 7d 11 ba  ||.q..A)l..;.\}..|
00000ce0  ad f1 01 02 c9 84 a8 1f  bb 27 83 82 b5 77 04 4d  |.........'...w.M|
00000cf0  bb 5a a7 f6 dc 31 37 4a  69 df ef dd bf c9 5b 23  |.Z...17Ji.....[#|
00000d00  95 5d 70 23 a4 59 ba 84  c8 4a 8d bc 78 d3 a3 75  |.]p#.Y...J..x..u|
00000d10  9c f3 00 30 7c bd bf ed  13 a0 25 3f b0 fc 82 a8  |...0|.....%?....|
00000d20  97 ef 27 fd 33 48 ae 98  81 11 83 45 cb 33 8b 31  |..'.3H.....E.3.1|
00000d30  e7 34 2f 25 b5 7e c1 e2  b1 d9 7a 49 42 bd 5d d8  |.4/%.~....zIB.].|
00000d40  ee 66 38 d2 0f e2 36 9e  ac 63 2d af 1e a0 0d 9a  |.f8...6..c-.....|
00000d50  d9 b6 37 49 6a da cb 7b  37 08 ec e4 35 b7 34 73  |..7Ij..{7...5.4s|
00000d60  f8 7c a3 d4 dd 23 c5 c0  f1 f8 58 bd a0 0d cb 3d  |.|...#....X....=|
00000d70  45 df 00 11 01 00 01 89  01 bc 04 18 01 0a 00 26  |E..............&|
00000d80  16 21 04 7d 39 61 e5 8c  17 5a 9a c2 50 3c 98 ad  |.!.}9a...Z..P<..|
00000d90  1b 3c 60 62 cd 81 18 05  02 66 51 13 bf 02 1b 0c  |.<`b.....fQ.....|
00000da0  05 09 03 c2 67 00 00 0a  09 10 ad 1b 3c 60 62 cd  |....g.......<`b.|
00000db0  81 18 23 af 0c 00 bf 1d  83 a1 16 6e 5e e5 90 ad  |..#........n^...|
00000dc0  7a 74 7e 23 36 95 5e bd  db e1 92 93 cd cb 8b 2b  |zt~#6.^........+|
00000dd0  e6 60 4f b5 ae 67 a5 aa  65 b5 43 de 92 57 d9 1d  |.`O..g..e.C..W..|
00000de0  57 53 13 44 c6 90 4d a8  99 af e9 4f 7d 34 94 bc  |WS.D..M....O}4..|
00000df0  ce 57 7b 2e 9b 46 e1 e9  00 e9 0e ad cd 3e 5d b6  |.W{..F.......>].|
00000e00  d1 17 a3 32 27 76 23 01  bb 48 68 51 82 6b d9 1f  |...2'v#..HhQ.k..|
00000e10  c1 0f 01 2a 78 f9 12 bd  7d d0 da 5e cd f0 f2 c6  |...*x...}..^....|
00000e20  7a a6 9f b3 65 1c 86 3f  e4 24 2d b3 02 91 03 6f  |z...e..?.$-....o|
00000e30  a4 73 34 ca 27 6a b9 4a  7c 6e 78 bb 1e e6 41 7e  |.s4.'j.J|nx...A~|
00000e40  31 6a e3 82 75 3d c7 97  6a b6 12 b0 86 79 8b af  |1j..u=..j....y..|
00000e50  d6 42 9a d7 ee f2 90 d7  05 f4 5d 27 a5 49 d1 99  |.B........]'.I..|
00000e60  a1 7b a5 e2 d1 a4 fd 19  1a d6 2c 63 9e d6 3f bb  |.{........,c..?.|
00000e70  b0 28 07 a9 a1 dc 3e 0a  14 5c 46 c6 09 6a 83 8e  |.(....>..\F..j..|
00000e80  9f cf eb 1b 26 e7 a4 bf  15 fc 0f e1 4c 06 b4 da  |....&.......L...|
00000e90  19 ab d4 88 66 02 eb f6  3f c9 ca 26 91 2a 01 ee  |....f...?..&.*..|
00000ea0  a9 90 d3 5d 23 9b a1 bf  aa 5a 62 d9 45 cb c7 6b  |...]#....Zb.E..k|
00000eb0  94 96 67 20 a5 42 0a 53  dd 0b 10 1f b5 52 fc 32  |..g .B.S.....R.2|
00000ec0  de 20 b3 71 57 91 51 e9  d7 ac eb e8 67 98 e3 54  |. .qW.Q.....g..T|
00000ed0  fd ee b3 e0 c6 c2 35 66  9e 67 c4 64 3b 21 00 f2  |......5f.g.d;!..|
00000ee0  eb ee 00 b1 35 11 48 9e  ac 98 20 d2 73 5b e5 2c  |....5.H... .s[.,|
00000ef0  f9 67 37 91 a7 f6 22 bb  ec b8 6d 17 7a 7a 7d 2e  |.g7..."...m.zz}.|
00000f00  00 18 03 72 5a 12 4f a8  c1 e1 99 9c cb 6c 61 b9  |...rZ.O......la.|
00000f10  3e 33 63 26 15 37 9a b1  72 4a ca ce a4 f2 f8 cc  |>3c&.7..rJ......|
00000f20  32 60 e7 cd fc c4 af 2a  0c 31 75 df 5b fb 2f ac  |2`.....*.1u.[./.|
00000f30  f5 2d 7b 73 a7 67 b0 06  00 00 67 70 67 00 70 37  |.-{s.g....gpg.p7|
00000f40  46 62 4b 1a af 7a c6 eb  87 b4 92 73 6f 46 a6 cc  |FbK..z.....soF..|
00000f50  35 9b 00 00 03 51 02 01  00 00 00 00 00 6e 00 00  |5....Q.......n..|
00000f60  02 cf 00 01 00 1c e8 78  0d 3b 1c 24 34 86 32 ab  |.......x.;.$4.2.|
00000f70  4b f1 17 37 e9 ce 68 96  5a 60 00 00 00 20 00 00  |K..7..h.Z`... ..|
00000f80  00 00 00 00 00 02 00 0c  00 00 01 13 00 00 00 23  |...............#|
00000f90  00 00 00 00 00 00 02 29  00 00 00 25 00 00 00 00  |.......)...%....|
00000fa0  00 02 00 04 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000fb0  00 00 00 00 00 00 00 00  66 52 25 cf 00 00 00 00  |........fR%.....|
00000fc0  98 93 04 66 51 14 80 13  05 2b 81 04 00 23 04 23  |...fQ....+...#.#|
00000fd0  04 00 56 33 14 68 b7 41  b6 69 b4 2c 61 97 8a 48  |..V3.h.A.i.,a..H|
00000fe0  5a fb 6a 52 6d 08 87 ee  6c d8 58 aa dd 15 eb 85  |Z.jRm...l.X.....|
00000ff0  9a ea 48 ec 55 8b 09 1e  dd da 7a c0 e4 52 c7 54  |..H.U.....z..R.T|
00001000  f1 8a 9b e7 a7 ba e2 d0  7a 94 32 e4 06 0a f4 48  |........z.2....H|
00001010  c8 53 cf 00 a5 86 59 5a  5e 7a 0a c6 3d 7b f4 5c  |.S....YZ^z..={.\|
00001020  13 d3 43 4c 24 3f 1b 5d  4a dd 72 89 61 ff 7d ad  |..CL$?.]J.r.a.}.|
00001030  90 1e a1 6e d9 76 bf 16  35 65 93 81 5e a3 56 e2  |...n.v..5e..^.V.|
00001040  69 f7 2b 72 42 c8 4d 16  01 26 94 6c 22 1e b5 2f  |i.+rB.M..&.l"../|
00001050  0f 8d 20 09 63 b0 0c 00  00 67 70 67 01 00 00 00  |.. .c....gpg....|
00001060  00 00 00 b4 23 53 61 6d  20 43 61 6c 64 77 65 6c  |....#Sam Caldwel|
00001070  6c 20 3c 6d 61 69 6c 40  73 61 6d 63 61 6c 64 77  |l <mail@samcaldw|
00001080  65 6c 6c 2e 6e 65 74 3e  b0 0c 00 00 67 70 67 02  |ell.net>....gpg.|
00001090  00 00 00 00 00 00 88 d9  04 13 13 0a 00 3e 16 21  |.............>.!|
000010a0  04 e8 78 0d 3b 1c 24 34  86 32 ab 4b f1 17 37 e9  |..x.;.$4.2.K..7.|
000010b0  ce 68 96 5a 60 05 02 66  51 14 80 02 1b 21 05 09  |.h.Z`..fQ....!..|
000010c0  12 cc 03 00 05 0b 09 08  07 02 06 15 0a 09 08 0b  |................|
000010d0  02 04 16 02 03 01 02 1e  01 02 17 80 00 0a 09 10  |................|
000010e0  17 37 e9 ce 68 96 5a 60  e0 31 02 09 01 fc a4 a2  |.7..h.Z`.1......|
000010f0  46 ab 42 c5 ea da ec 55  22 2d 1a f7 de 58 99 4f  |F.B....U"-...X.O|
00001100  0b c7 64 49 a8 5c f5 b8  12 3c 1e c7 e0 46 ad 63  |..dI.\...<...F.c|
00001110  a0 4d 0e fa c9 08 d4 c3  74 6d 01 0d 69 27 7a ad  |.M......tm..i'z.|
00001120  a9 18 99 9e 97 52 3c 8a  78 94 e2 dc bb ab 02 08  |.....R<.x.......|
00001130  8a d6 e2 e2 33 d6 94 8f  eb 54 9d 99 61 3f ec 6a  |....3....T..a?.j|
00001140  c3 f9 28 0f 70 c9 d1 6e  c5 61 dd 01 10 46 b7 01  |..(.p..n.a...F..|
00001150  e1 10 a6 73 86 ec 44 e4  01 98 9a f9 b5 9f df d8  |...s..D.........|
00001160  c7 9d 9a e1 3f 94 34 f5  25 53 e9 18 a8 e4 65 86  |....?.4.%S....e.|
00001170  eb b0 06 00 03 67 70 67  00 b4 25 53 61 6d 20 43  |.....gpg..%Sam C|
00001180  61 6c 64 77 65 6c 6c 20  3c 67 69 74 68 75 62 40  |aldwell <github@|
00001190  73 61 6d 63 61 6c 64 77  65 6c 6c 2e 6e 65 74 3e  |samcaldwell.net>|
000011a0  b0 0c 00 00 67 70 67 02  00 00 00 00 00 00 88 d7  |....gpg.........|
000011b0  04 13 13 0a 00 3c 16 21  04 e8 78 0d 3b 1c 24 34  |.....<.!..x.;.$4|
000011c0  86 32 ab 4b f1 17 37 e9  ce 68 96 5a 60 05 02 66  |.2.K..7..h.Z`..f|
000011d0  52 25 62 02 1b 21 05 09  12 cc 03 00 04 0b 09 08  |R%b..!..........|
000011e0  07 04 15 0a 09 08 05 16  02 03 01 00 02 1e 01 02  |................|
000011f0  17 80 00 0a 09 10 17 37  e9 ce 68 96 5a 60 5a fc  |.......7..h.Z`Z.|
00001200  02 09 01 00 23 80 50 f6  76 70 d2 d3 8c 06 d9 3c  |....#.P.vp.....<|
00001210  14 4f 64 d9 f1 af c9 c7  9a 81 c9 a1 d6 cd 1c 51  |.Od............Q|
00001220  a7 cf e4 fc 56 de d5 c8  b5 44 a9 85 66 3b fc 02  |....V....D..f;..|
00001230  9c 91 19 32 4f 17 03 5c  fb bf 39 8a 64 42 99 76  |...2O..\..9.dB.v|
00001240  9c 82 ef c1 02 08 b8 a0  3a c9 99 7a f9 30 3f b5  |........:..z.0?.|
00001250  04 db 6e d4 8f ba 7f 35  ea 53 df c4 6e 10 c7 22  |..n....5.S..n.."|
00001260  c2 79 0a 5b 99 85 a2 12  aa a2 32 57 d9 f3 e7 dd  |.y.[......2W....|
00001270  20 14 21 35 ab fa 14 2a  ad d3 aa 5f b6 37 e6 11  | .!5...*..._.7..|
00001280  ad 75 81 54 cf 9e a2 b0  06 00 03 67 70 67 00 18  |.u.T.......gpg..|
00001290  59 19 8e cf e8 82 7a 0a  40 93 49 bd a4 22 53 50  |Y.....z.@.I.."SP|
000012a0  bc ad 6d                                          |..m|
000012a3
```
