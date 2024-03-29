# Lecture Plan 2021

Each video recorded is linked below.
You may also be interested in this [YouTube playlist](https://www.youtube.com/playlist?list=PLEHv3FhiBSaaQk_RR9TPFnA7Uhgo6GF1F).

| W  | Date  | Day | Ch | Topics                                                                               |
|----|-------|-----|----|--------------------------------------------------------------------------------------|
| 34 | 25.08 | Wed |    | [Course Intro and Tutorials][1]                                                      |
| 34 |       | Wed | 2  | [Introduction to Operating Systems][2]                                               |
| 34 |       | Wed | 4  | [Abstraction: The Process][3]                                                        |
| 34 | 26.08 | Thu | 5  | [Process API][4]                                                                     |
| 34 |       | Thu | 6  | [Mechanism: Limited Direct Execution][5]                                             |
| 35 | 01.09 | Wed |    | [Introduction to Go programming][6]                                                  |
| 35 |       | Wed | 7  | [Scheduling: Introduction][7]                                                        |
| 35 |       | Wed | 8  | [Scheduling: Multi-Level Feedback Queue][8]                                          |
| 35 | 02.09 | Thu | 9  | [Scheduling: Proportional Share][9]                                                  |
| 36 | 08.09 | Wed |    | [Organizing Go Code][10]                                                             |
| 36 |       | Wed | 13 | [Abstraction: Address Spaces][11]                                                    |
| 36 |       | Wed | 14 | [Memory API][12]                                                                     |
| 36 | 09.09 | Thu | 15 | [Mechanism: Address Translation][13]                                                 |
| 36 |       | Thu | 16 | [Segmentation][14]                                                                   |
| 37 | 15.09 | Wed | 17 | [Free-Space Management][15]                                                          |
| 37 |       | Wed | 18 | [Paging: Introduction][16], [Part 2][17]                                             |
| 37 | 16.09 | Thu | 19 | [Paging: Faster Translation (TLBs)][18], [Part 2][19]                                |
| 38 | 22.09 | Wed | 20 | [Paging: Smaller Tables][20], [Part 2][21]                                           |
| 38 |       | Wed | 21 | [Beyond Physical Memory: Mechanisms][22]                                             |
| 38 | 23.09 | Thu | 22 | [Beyond Physical Memory: Policies][23], [Part 2][24]                                 |
| 39 | 29.09 | Wed | 23 | [Complete Virtual Memory Systems][25], [Part 2][26]                                  |
| 39 |       | Wed | 23 | [The Mystery of O(N^2) Matrix Traverse (21:58)][27]                                  |
| 39 |       | Wed | 23 | [Buffer Overflow (17:30)][28]                                                        |
| 39 |       | Wed | 23 | [Spectre & Meltdown (13:44)][29]                                                     |
| 39 |       | Wed | 26 | [Concurrency: Introduction][30], [Part 2][31]                                        |
| 39 | 30.09 | Thu | 27 | [Thread API][32], [Part 2][33]                                                       |
| 40 | 06.10 | Wed |    | [Rob Pike: Go Concurrency Patterns][34]                                              |
| 40 |       | Wed |    | [Live Coding: Shared Integer w/Mutual Exclusion][35]                                 |
| 40 |       | Wed | 28 | [Locks][36], [Part 2][37]                                                            |
| 40 |       | Wed | 29 | [Lock-based Concurrent Data Structures][38], [Part 2][39], [Part 3][40]              |
| 40 | 07.10 | Thu | 30 | [Condition Variables][41], [Part 2][42]                                              |
| 41 | 13.10 | Wed | 31 | [Semaphores][43], [Part 2][44], [Part 3][45]                                         |
| 41 | 14.10 | Thu | 32 | [Common Concurrency Problems][46], [Part 2][47]                                      |
| 42 | 20.10 | Wed | 10 | [Multiprocessor Scheduling][48]                                                      |

| Ch | Topics to watch on your own                                                          |
|----|--------------------------------------------------------------------------------------|
|    | [Network Programming with gRPC in Go][51], Parts: [2][52], [3][53], [4][54], [5][55] |
| 39 | [Files and Directories][49], [Part 2][50]                                            |
| 40 | [File System Implementation][56], [Part 2][57]                                       |
| 48 | [Distributed Systems][58]                                                            |
| 48 | [Live Coding: TCP Echo Client/Server in Go][59], [Part 2][60]                        |
| 49 | [Sun's Network File System][61] [Part 2][62] [Part 3][63]                            |

[1]: https://youtu.be/oORmvjot6wc
[2]: https://youtu.be/UVpbQnaagYE
[3]: https://youtu.be/ok-nbl2wFbM
[4]: https://youtu.be/Ab3rPs3l-5I
[5]: https://youtu.be/32i0xvcYuJo
[6]: https://youtu.be/vqq96BG9aOo
[7]: https://youtu.be/YHK9xqOsQz0
[8]: https://youtu.be/gb93s6kWLLM
[9]: https://youtu.be/jO6wUeTa0lE
[10]: https://youtu.be/cJmYVEx__c8
[11]: https://youtu.be/VZQkKpY8pB8
[12]: https://youtu.be/cPBYxwNgzYU
[13]: https://youtu.be/CZ3KYVV9X08
[14]: https://youtu.be/Riv_PmvEBc0
[15]: https://youtu.be/AbL6Imqr44g
[16]: https://youtu.be/8dUtAVRqKyI
[17]: https://youtu.be/AtqgKOmNwrU
[18]: https://youtu.be/wymc8KWptDo
[19]: https://youtu.be/_FLZplf8JOM
[20]: https://youtu.be/iPIXEMzPq-s
[21]: https://youtu.be/iRfnZVFYTRE
[22]: https://youtu.be/iyDSULxT4hI
[23]: https://youtu.be/dboKNgOpDFo
[24]: https://youtu.be/cNj1IZrizaU
[25]: https://youtu.be/Aw1fkkj6ymQ
[26]: https://youtu.be/q-C2OhlIrlk
[27]: https://youtu.be/rtfHdM6XSV0
[28]: https://youtu.be/1S0aBV-Waeo
[29]: https://youtu.be/I5mRwzVvFGE
[30]: https://youtu.be/enWyVjihK3c
[31]: https://youtu.be/B5MUHjFfV7w
[32]: https://youtu.be/ERS5CHWq5DI
[33]: https://youtu.be/Rcxt7UAit8Q
[34]: https://youtu.be/f6kdp27TYZs
[35]: https://youtu.be/5jkAxITBTVM
[36]: https://youtu.be/AiaWgIreiCY
[37]: https://youtu.be/sCrWRgqzMGA
[38]: https://youtu.be/AVESx9pPmU4
[39]: https://youtu.be/QSVAsOZ6pd0
[40]: https://youtu.be/NUSHciImsbA
[41]: https://youtu.be/FP4vDkFWx3E
[42]: https://youtu.be/vpcKfxtu2yo
[43]: https://youtu.be/RCNOuKZIoog
[44]: https://youtu.be/rKUiBLaQJlg
[45]: https://youtu.be/5VlspGnJJeQ
[46]: https://youtu.be/bpTqDjMcmjY
[47]: https://youtu.be/6w2EfCZM7cg
[48]: https://youtu.be/4SdyybS7q84
[49]: https://youtu.be/GPNZ2Ztgmng
[50]: https://youtu.be/Q9nhu6Yz1KE
[51]: https://youtu.be/qKGNmmNSgHM
[52]: https://youtu.be/v_rd8h-ox_Y
[53]: https://youtu.be/a41gvwO5xKY
[54]: https://youtu.be/NXqJr6s78Gs
[55]: https://youtu.be/1dp1soZbHBs
[56]: https://youtu.be/bbPAbbq4tEw
[57]: https://youtu.be/mB7U5ETiiVs
[58]: https://youtu.be/rQ02ixlgbno
[59]: https://youtu.be/9t8U4Qs0Lts
[60]: https://youtu.be/2dcWPX4yhDM
[61]: https://youtu.be/Y14A5orAzts
[62]: https://youtu.be/UGrvoXk-JpQ
[63]: https://youtu.be/VYg-iWCgZZY
