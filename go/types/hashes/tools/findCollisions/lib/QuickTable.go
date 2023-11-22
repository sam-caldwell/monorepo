package findCollision

import (
	"github.com/sam-caldwell/monorepo/go/counters"
	"log"
	"time"
)

func NewQuickTable(keySpaceSize, TableSize int) (*QuickTable, []byte) {
	var start time.Time
	var pos int
	generatorStart := time.Now()
	go func() {
		t := time.NewTicker(1 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				progress := 100 * float64(pos) / float64(TableSize)
				log.Printf("pre-populating table.  progress (%d/%d) %3.6f (t/op:%6d) elapsed:%v",
					pos, TableSize, progress, time.Since(start).Microseconds(),
					time.Since(generatorStart).Seconds())
			}
		}
	}()

	table := make(QuickTable)
	c, _ := counters.NewByteCounter(keySpaceSize)

	for i := 0; i < TableSize; i++ {
		start = time.Now()
		table.Store(c.Sha1Bytes())
		_ = c.Increment()
		pos = i
	}
	return &table, c.Bytes()
}

type QuickTable map[byte]QuickTable

func (o *QuickTable) Store(s [20]byte) {
	(*o)[s[0]] = QuickTable{
		s[1]: QuickTable{
			s[2]: QuickTable{
				s[3]: QuickTable{
					s[4]: QuickTable{
						s[5]: QuickTable{
							s[6]: QuickTable{
								s[7]: QuickTable{
									s[8]: QuickTable{
										s[9]: QuickTable{
											s[10]: QuickTable{
												s[11]: QuickTable{
													s[12]: QuickTable{
														s[13]: QuickTable{
															s[14]: QuickTable{
																s[15]: QuickTable{
																	s[16]: QuickTable{
																		s[17]: QuickTable{
																			s[18]: QuickTable{
																				s[19]: nil,
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (o *QuickTable) Lookup(s [20]byte) bool {
	if n1, ok := (*o)[s[0]]; ok {
		if n2, ok := n1[s[1]]; ok {
			if n3, ok := n2[s[2]]; ok {
				if n4, ok := n3[s[3]]; ok {
					if n5, ok := n4[s[4]]; ok {
						if n6, ok := n5[s[5]]; ok {
							if n7, ok := n6[s[6]]; ok {
								if n8, ok := n7[s[7]]; ok {
									if n9, ok := n8[s[8]]; ok {
										if n10, ok := n9[s[9]]; ok {
											if n11, ok := n10[s[10]]; ok {
												if n12, ok := n11[s[11]]; ok {
													if n13, ok := n12[s[12]]; ok {
														if n14, ok := n13[s[13]]; ok {
															if n15, ok := n14[s[14]]; ok {
																if n16, ok := n15[s[15]]; ok {
																	if n17, ok := n16[s[16]]; ok {
																		if n18, ok := n17[s[17]]; ok {
																			if n19, ok := n18[s[18]]; ok {
																				if _, ok := n19[s[19]]; ok {
																					return true
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}
