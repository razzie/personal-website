package countries

// TypeCurrencyCode for Typer interface
const TypeCurrencyCode string = "countries.CurrencyCode"

// TypeCurrency for Typer interface
const TypeCurrency string = "countries.Currency"

// Currencies. Two codes present, for example CurrencyUSDollar == CurrencyUSD == 840.
const (
	CurrencyUnknown                        CurrencyCode = 0
	CurrencyAfghani                        CurrencyCode = 971
	CurrencyLek                            CurrencyCode = 8
	CurrencyAlgerianDinar                  CurrencyCode = 12
	CurrencyUSDollar                       CurrencyCode = 840
	CurrencyEuro                           CurrencyCode = 978
	CurrencyKwanza                         CurrencyCode = 973
	CurrencyEastCaribbeanDollar            CurrencyCode = 951
	CurrencyArgentinePeso                  CurrencyCode = 32
	CurrencyArmenianDram                   CurrencyCode = 51
	CurrencyArubanFlorin                   CurrencyCode = 533
	CurrencyAustralianDollar               CurrencyCode = 36
	CurrencyAzerbaijanianManat             CurrencyCode = 944
	CurrencyBahamianDollar                 CurrencyCode = 44
	CurrencyBahrainiDinar                  CurrencyCode = 48
	CurrencyTaka                           CurrencyCode = 50
	CurrencyBarbadosDollar                 CurrencyCode = 52
	CurrencyBelarussianRuble               CurrencyCode = 933
	CurrencyBelizeDollar                   CurrencyCode = 84
	CurrencyCFAFrancBCEAO                  CurrencyCode = 952
	CurrencyBermudianDollar                CurrencyCode = 60
	CurrencyNgultrum                       CurrencyCode = 64
	CurrencyIndianRupee                    CurrencyCode = 356
	CurrencyBoliviano                      CurrencyCode = 68
	CurrencyConvertibleMark                CurrencyCode = 977
	CurrencyPula                           CurrencyCode = 72
	CurrencyNorwegianKrone                 CurrencyCode = 578
	CurrencyBrazilianReal                  CurrencyCode = 986
	CurrencyBruneiDollar                   CurrencyCode = 96
	CurrencyBulgarianLev                   CurrencyCode = 975
	CurrencyBurundiFranc                   CurrencyCode = 108
	CurrencyCaboVerdeEscudo                CurrencyCode = 132
	CurrencyRiel                           CurrencyCode = 116
	CurrencyCFAFrancBEAC                   CurrencyCode = 950
	CurrencyCanadianDollar                 CurrencyCode = 124
	CurrencyCaymanIslandsDollar            CurrencyCode = 136
	CurrencyUnidaddeFomento                CurrencyCode = 990
	CurrencyChileanPeso                    CurrencyCode = 152
	CurrencyYuanRenminbi                   CurrencyCode = 156
	CurrencyColombianPeso                  CurrencyCode = 170
	CurrencyUnidaddeValorReal              CurrencyCode = 970
	CurrencyComoroFranc                    CurrencyCode = 174
	CurrencyCongoleseFranc                 CurrencyCode = 976
	CurrencyNewZealandDollar               CurrencyCode = 554
	CurrencyCostaRicanColon                CurrencyCode = 188
	CurrencyKuna                           CurrencyCode = 191
	CurrencyPesoConvertible                CurrencyCode = 931
	CurrencyCubanPeso                      CurrencyCode = 192
	CurrencyNetherlandsAntilleanGuilder    CurrencyCode = 532
	CurrencyCzechKoruna                    CurrencyCode = 203
	CurrencyDanishKrone                    CurrencyCode = 208
	CurrencyDjiboutiFranc                  CurrencyCode = 262
	CurrencyDominicanPeso                  CurrencyCode = 214
	CurrencyEgyptianPound                  CurrencyCode = 818
	CurrencyElSalvadorColon                CurrencyCode = 222
	CurrencyNakfa                          CurrencyCode = 232
	CurrencyEthiopianBirr                  CurrencyCode = 230
	CurrencyFalklandIslandsPound           CurrencyCode = 238
	CurrencyFijiDollar                     CurrencyCode = 242
	CurrencyCFPFranc                       CurrencyCode = 953
	CurrencyDalasi                         CurrencyCode = 270
	CurrencyLari                           CurrencyCode = 981
	CurrencyGhanaCedi                      CurrencyCode = 936
	CurrencyGibraltarPound                 CurrencyCode = 292
	CurrencyQuetzal                        CurrencyCode = 320
	CurrencyPoundSterling                  CurrencyCode = 826
	CurrencyGuineaFranc                    CurrencyCode = 324
	CurrencyGuyanaDollar                   CurrencyCode = 328
	CurrencyGourde                         CurrencyCode = 332
	CurrencyLempira                        CurrencyCode = 340
	CurrencyHongKongDollar                 CurrencyCode = 344
	CurrencyForint                         CurrencyCode = 348
	CurrencyIcelandKrona                   CurrencyCode = 352
	CurrencyRupiah                         CurrencyCode = 360
	CurrencySDR                            CurrencyCode = 960
	CurrencyIranianRial                    CurrencyCode = 364
	CurrencyIraqiDinar                     CurrencyCode = 368
	CurrencyNewIsraeliSheqel               CurrencyCode = 376
	CurrencyJamaicanDollar                 CurrencyCode = 388
	CurrencyYen                            CurrencyCode = 392
	CurrencyJordanianDinar                 CurrencyCode = 400
	CurrencyTenge                          CurrencyCode = 398
	CurrencyKenyanShilling                 CurrencyCode = 404
	CurrencyNorthKoreanWon                 CurrencyCode = 408
	CurrencyWon                            CurrencyCode = 410
	CurrencyKuwaitiDinar                   CurrencyCode = 414
	CurrencySom                            CurrencyCode = 417
	CurrencyKip                            CurrencyCode = 418
	CurrencyLebanesePound                  CurrencyCode = 422
	CurrencyLoti                           CurrencyCode = 426
	CurrencyRand                           CurrencyCode = 710
	CurrencyLiberianDollar                 CurrencyCode = 430
	CurrencyLibyanDinar                    CurrencyCode = 434
	CurrencySwissFranc                     CurrencyCode = 756
	CurrencyPataca                         CurrencyCode = 446
	CurrencyDenar                          CurrencyCode = 807
	CurrencyMalagasyAriary                 CurrencyCode = 969
	CurrencyKwacha                         CurrencyCode = 454
	CurrencyMalaysianRinggit               CurrencyCode = 458
	CurrencyRufiyaa                        CurrencyCode = 462
	CurrencyOuguiya                        CurrencyCode = 929
	CurrencyMauritiusRupee                 CurrencyCode = 480
	CurrencyADBUnitofAccount               CurrencyCode = 965
	CurrencyMexicanPeso                    CurrencyCode = 484
	CurrencyMexicanUnidaddeInversion       CurrencyCode = 979
	CurrencyMexicanUDI                     CurrencyCode = 979
	CurrencyMoldovanLeu                    CurrencyCode = 498
	CurrencyTugrik                         CurrencyCode = 496
	CurrencyMoroccanDirham                 CurrencyCode = 504
	CurrencyMozambiqueMetical              CurrencyCode = 943
	CurrencyKyat                           CurrencyCode = 104
	CurrencyNamibiaDollar                  CurrencyCode = 516
	CurrencyNepaleseRupee                  CurrencyCode = 524
	CurrencyCordobaOro                     CurrencyCode = 558
	CurrencyNaira                          CurrencyCode = 566
	CurrencyRialOmani                      CurrencyCode = 512
	CurrencyPakistanRupee                  CurrencyCode = 586
	CurrencyBalboa                         CurrencyCode = 590
	CurrencyKina                           CurrencyCode = 598
	CurrencyGuarani                        CurrencyCode = 600
	CurrencyNuevoSol                       CurrencyCode = 604
	CurrencyPhilippinePeso                 CurrencyCode = 608
	CurrencyZloty                          CurrencyCode = 985
	CurrencyQatariRial                     CurrencyCode = 634
	CurrencyRomanianLeu                    CurrencyCode = 946
	CurrencyRussianRuble                   CurrencyCode = 643
	CurrencyRwandaFranc                    CurrencyCode = 646
	CurrencySaintHelenaPound               CurrencyCode = 654
	CurrencyTala                           CurrencyCode = 882
	CurrencyDobra                          CurrencyCode = 930
	CurrencySaudiRiyal                     CurrencyCode = 682
	CurrencySerbianDinar                   CurrencyCode = 941
	CurrencySeychellesRupee                CurrencyCode = 690
	CurrencyLeone                          CurrencyCode = 694
	CurrencySingaporeDollar                CurrencyCode = 702
	CurrencySucre                          CurrencyCode = 994
	CurrencySolomonIslandsDollar           CurrencyCode = 90
	CurrencySomaliShilling                 CurrencyCode = 706
	CurrencySouthSudanesePound             CurrencyCode = 728
	CurrencySriLankaRupee                  CurrencyCode = 144
	CurrencySudanesePound                  CurrencyCode = 938
	CurrencySurinamDollar                  CurrencyCode = 968
	CurrencyLilangeni                      CurrencyCode = 748
	CurrencySwedishKrona                   CurrencyCode = 752
	CurrencyWIREuro                        CurrencyCode = 947
	CurrencyWIRFranc                       CurrencyCode = 948
	CurrencySyrianPound                    CurrencyCode = 760
	CurrencyNewTaiwanDollar                CurrencyCode = 901
	CurrencySomoni                         CurrencyCode = 972
	CurrencyTanzanianShilling              CurrencyCode = 834
	CurrencyBaht                           CurrencyCode = 764
	CurrencyPaanga                         CurrencyCode = 776
	CurrencyTrinidadandTobagoDollar        CurrencyCode = 780
	CurrencyTunisianDinar                  CurrencyCode = 788
	CurrencyTurkishLira                    CurrencyCode = 949
	CurrencyTurkmenistanNewManat           CurrencyCode = 934
	CurrencyUgandaShilling                 CurrencyCode = 800
	CurrencyHryvnia                        CurrencyCode = 980
	CurrencyUAEDirham                      CurrencyCode = 784
	CurrencyUSDollarNextday                CurrencyCode = 997
	CurrencyUruguayPesoenUnidadesIndexadas CurrencyCode = 940
	CurrencyUruguayPUI                     CurrencyCode = 940
	CurrencyURUIURUI                       CurrencyCode = 940
	CurrencyPesoUruguayo                   CurrencyCode = 858
	CurrencyUzbekistanSum                  CurrencyCode = 860
	CurrencyVatu                           CurrencyCode = 548
	CurrencyBolivar                        CurrencyCode = 928
	CurrencyBolivarDeprecated              CurrencyCode = 937
	CurrencyDong                           CurrencyCode = 704
	CurrencyYemeniRial                     CurrencyCode = 886
	CurrencyZambianKwacha                  CurrencyCode = 967
	CurrencyZimbabweDollar                 CurrencyCode = 932
	CurrencyYugoslavianDinar               CurrencyCode = 891
	CurrencyNone                           CurrencyCode = 998
)

// Currencies by ISO 4217. Two codes present, for example CurrencyUSDollar == CurrencyUSD == 840.
const (
	CurrencyAFN CurrencyCode = 971
	CurrencyALL CurrencyCode = 8
	CurrencyDZD CurrencyCode = 12
	CurrencyUSD CurrencyCode = 840
	CurrencyEUR CurrencyCode = 978
	CurrencyAOA CurrencyCode = 973
	CurrencyXCD CurrencyCode = 951
	CurrencyARS CurrencyCode = 32
	CurrencyAMD CurrencyCode = 51
	CurrencyAWG CurrencyCode = 533
	CurrencyAUD CurrencyCode = 36
	CurrencyAZN CurrencyCode = 944
	CurrencyBSD CurrencyCode = 44
	CurrencyBHD CurrencyCode = 48
	CurrencyBDT CurrencyCode = 50
	CurrencyBBD CurrencyCode = 52
	CurrencyBYN CurrencyCode = 933
	CurrencyBZD CurrencyCode = 84
	CurrencyXOF CurrencyCode = 952
	CurrencyBMD CurrencyCode = 60
	CurrencyBTN CurrencyCode = 64
	CurrencyINR CurrencyCode = 356
	CurrencyBOB CurrencyCode = 68
	CurrencyBAM CurrencyCode = 977
	CurrencyBWP CurrencyCode = 72
	CurrencyNOK CurrencyCode = 578
	CurrencyBRL CurrencyCode = 986
	CurrencyBND CurrencyCode = 96
	CurrencyBGN CurrencyCode = 975
	CurrencyBIF CurrencyCode = 108
	CurrencyCVE CurrencyCode = 132
	CurrencyKHR CurrencyCode = 116
	CurrencyXAF CurrencyCode = 950
	CurrencyCAD CurrencyCode = 124
	CurrencyKYD CurrencyCode = 136
	CurrencyCLF CurrencyCode = 990
	CurrencyCLP CurrencyCode = 152
	CurrencyCNY CurrencyCode = 156
	CurrencyCOP CurrencyCode = 170
	CurrencyCOU CurrencyCode = 970
	CurrencyKMF CurrencyCode = 174
	CurrencyCDF CurrencyCode = 976
	CurrencyNZD CurrencyCode = 554
	CurrencyCRC CurrencyCode = 188
	CurrencyHRK CurrencyCode = 191
	CurrencyCUC CurrencyCode = 931
	CurrencyCUP CurrencyCode = 192
	CurrencyANG CurrencyCode = 532
	CurrencyCZK CurrencyCode = 203
	CurrencyDKK CurrencyCode = 208
	CurrencyDJF CurrencyCode = 262
	CurrencyDOP CurrencyCode = 214
	CurrencyEGP CurrencyCode = 818
	CurrencySVC CurrencyCode = 222
	CurrencyERN CurrencyCode = 232
	CurrencyETB CurrencyCode = 230
	CurrencyFKP CurrencyCode = 238
	CurrencyFJD CurrencyCode = 242
	CurrencyXPF CurrencyCode = 953
	CurrencyGMD CurrencyCode = 270
	CurrencyGEL CurrencyCode = 981
	CurrencyGHS CurrencyCode = 936
	CurrencyGIP CurrencyCode = 292
	CurrencyGTQ CurrencyCode = 320
	CurrencyGBP CurrencyCode = 826
	CurrencyGNF CurrencyCode = 324
	CurrencyGYD CurrencyCode = 328
	CurrencyHTG CurrencyCode = 332
	CurrencyHNL CurrencyCode = 340
	CurrencyHKD CurrencyCode = 344
	CurrencyHUF CurrencyCode = 348
	CurrencyISK CurrencyCode = 352
	CurrencyIDR CurrencyCode = 360
	CurrencyXDR CurrencyCode = 960
	CurrencyIRR CurrencyCode = 364
	CurrencyIQD CurrencyCode = 368
	CurrencyILS CurrencyCode = 376
	CurrencyJMD CurrencyCode = 388
	CurrencyJPY CurrencyCode = 392
	CurrencyJOD CurrencyCode = 400
	CurrencyKZT CurrencyCode = 398
	CurrencyKES CurrencyCode = 404
	CurrencyKPW CurrencyCode = 408
	CurrencyKRW CurrencyCode = 410
	CurrencyKWD CurrencyCode = 414
	CurrencyKGS CurrencyCode = 417
	CurrencyLAK CurrencyCode = 418
	CurrencyLBP CurrencyCode = 422
	CurrencyLSL CurrencyCode = 426
	CurrencyZAR CurrencyCode = 710
	CurrencyLRD CurrencyCode = 430
	CurrencyLYD CurrencyCode = 434
	CurrencyCHF CurrencyCode = 756
	CurrencyMOP CurrencyCode = 446
	CurrencyMKD CurrencyCode = 807
	CurrencyMGA CurrencyCode = 969
	CurrencyMWK CurrencyCode = 454
	CurrencyMYR CurrencyCode = 458
	CurrencyMVR CurrencyCode = 462
	CurrencyMRU CurrencyCode = 929
	CurrencyMUR CurrencyCode = 480
	CurrencyXUA CurrencyCode = 965
	CurrencyMXN CurrencyCode = 484
	CurrencyMXV CurrencyCode = 979
	CurrencyMDL CurrencyCode = 498
	CurrencyMNT CurrencyCode = 496
	CurrencyMAD CurrencyCode = 504
	CurrencyMZN CurrencyCode = 943
	CurrencyMMK CurrencyCode = 104
	CurrencyNAD CurrencyCode = 516
	CurrencyNPR CurrencyCode = 524
	CurrencyNIO CurrencyCode = 558
	CurrencyNGN CurrencyCode = 566
	CurrencyOMR CurrencyCode = 512
	CurrencyPKR CurrencyCode = 586
	CurrencyPAB CurrencyCode = 590
	CurrencyPGK CurrencyCode = 598
	CurrencyPYG CurrencyCode = 600
	CurrencyPEN CurrencyCode = 604
	CurrencyPHP CurrencyCode = 608
	CurrencyPLN CurrencyCode = 985
	CurrencyQAR CurrencyCode = 634
	CurrencyRON CurrencyCode = 946
	CurrencyRUB CurrencyCode = 643
	CurrencyRWF CurrencyCode = 646
	CurrencySHP CurrencyCode = 654
	CurrencyWST CurrencyCode = 882
	CurrencySTN CurrencyCode = 930
	CurrencySAR CurrencyCode = 682
	CurrencyRSD CurrencyCode = 941
	CurrencySCR CurrencyCode = 690
	CurrencySLL CurrencyCode = 694
	CurrencySGD CurrencyCode = 702
	CurrencyXSU CurrencyCode = 994
	CurrencySBD CurrencyCode = 90
	CurrencySOS CurrencyCode = 706
	CurrencySSP CurrencyCode = 728
	CurrencyLKR CurrencyCode = 144
	CurrencySDG CurrencyCode = 938
	CurrencySRD CurrencyCode = 968
	CurrencySZL CurrencyCode = 748
	CurrencySEK CurrencyCode = 752
	CurrencyCHE CurrencyCode = 947
	CurrencyCHW CurrencyCode = 948
	CurrencySYP CurrencyCode = 760
	CurrencyTWD CurrencyCode = 901
	CurrencyTJS CurrencyCode = 972
	CurrencyTZS CurrencyCode = 834
	CurrencyTHB CurrencyCode = 764
	CurrencyTOP CurrencyCode = 776
	CurrencyTTD CurrencyCode = 780
	CurrencyTND CurrencyCode = 788
	CurrencyTRY CurrencyCode = 949
	CurrencyTMT CurrencyCode = 934
	CurrencyUGX CurrencyCode = 800
	CurrencyUAH CurrencyCode = 980
	CurrencyAED CurrencyCode = 784
	CurrencyUSN CurrencyCode = 997
	CurrencyUYI CurrencyCode = 940
	CurrencyUYU CurrencyCode = 858
	CurrencyUZS CurrencyCode = 860
	CurrencyVUV CurrencyCode = 548
	CurrencyVES CurrencyCode = 928
	CurrencyVEF CurrencyCode = 937
	CurrencyVND CurrencyCode = 704
	CurrencyYER CurrencyCode = 886
	CurrencyZMW CurrencyCode = 967
	CurrencyZWL CurrencyCode = 932
	CurrencyYUD CurrencyCode = 891
	CurrencyNON CurrencyCode = 998
)
