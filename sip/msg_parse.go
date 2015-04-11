
//line msg_parse.rl:1
// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
)


//line msg_parse.rl:12

//line msg_parse.rl:13

//line msg_parse.go:19
const msg_start int = 1
const msg_first_final int = 765
const msg_error int = 0

const msg_en_ctype int = 34
const msg_en_via_param int = 68
const msg_en_via int = 103
const msg_en_addr_param int = 153
const msg_en_addr_angled int = 188
const msg_en_addr_uri int = 229
const msg_en_addr int = 256
const msg_en_value int = 263
const msg_en_xheader int = 273
const msg_en_header int = 280
const msg_en_main int = 1


//line msg_parse.rl:14

// ParseMsg turns a SIP message into a data structure.
func ParseMsg(s string) (msg *Msg, err error) {
	if s == "" {
		return nil, errors.New("Empty SIP message")
	}
	return ParseMsgBytes([]byte(s))
}

// ParseMsg turns a SIP message byte slice into a data structure.
func ParseMsgBytes(data []byte) (msg *Msg, err error) {
	if data == nil {
		return nil, nil
	}
	msg = new(Msg)
	viap := &msg.Via
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	clen := 0
	ctype := ""
	var name string
	var hex byte
	var value *string
	var via *Via
	var addrp **Addr
	var addr *Addr

	
//line msg_parse.rl:47
	
//line msg_parse.go:73
	{
	cs = msg_start
	}

//line msg_parse.rl:48
	
//line msg_parse.go:80
	{
	var _widec int16
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 765:
		goto st_case_765
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 766:
		goto st_case_766
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 767:
		goto st_case_767
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 768:
		goto st_case_768
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 769:
		goto st_case_769
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	case 162:
		goto st_case_162
	case 163:
		goto st_case_163
	case 164:
		goto st_case_164
	case 165:
		goto st_case_165
	case 166:
		goto st_case_166
	case 167:
		goto st_case_167
	case 168:
		goto st_case_168
	case 169:
		goto st_case_169
	case 170:
		goto st_case_170
	case 171:
		goto st_case_171
	case 172:
		goto st_case_172
	case 173:
		goto st_case_173
	case 174:
		goto st_case_174
	case 175:
		goto st_case_175
	case 176:
		goto st_case_176
	case 177:
		goto st_case_177
	case 178:
		goto st_case_178
	case 179:
		goto st_case_179
	case 180:
		goto st_case_180
	case 181:
		goto st_case_181
	case 182:
		goto st_case_182
	case 183:
		goto st_case_183
	case 184:
		goto st_case_184
	case 185:
		goto st_case_185
	case 186:
		goto st_case_186
	case 187:
		goto st_case_187
	case 188:
		goto st_case_188
	case 189:
		goto st_case_189
	case 190:
		goto st_case_190
	case 191:
		goto st_case_191
	case 192:
		goto st_case_192
	case 193:
		goto st_case_193
	case 194:
		goto st_case_194
	case 195:
		goto st_case_195
	case 770:
		goto st_case_770
	case 196:
		goto st_case_196
	case 197:
		goto st_case_197
	case 198:
		goto st_case_198
	case 199:
		goto st_case_199
	case 200:
		goto st_case_200
	case 201:
		goto st_case_201
	case 202:
		goto st_case_202
	case 203:
		goto st_case_203
	case 204:
		goto st_case_204
	case 205:
		goto st_case_205
	case 206:
		goto st_case_206
	case 207:
		goto st_case_207
	case 208:
		goto st_case_208
	case 209:
		goto st_case_209
	case 210:
		goto st_case_210
	case 211:
		goto st_case_211
	case 212:
		goto st_case_212
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 225:
		goto st_case_225
	case 226:
		goto st_case_226
	case 227:
		goto st_case_227
	case 228:
		goto st_case_228
	case 229:
		goto st_case_229
	case 230:
		goto st_case_230
	case 231:
		goto st_case_231
	case 232:
		goto st_case_232
	case 233:
		goto st_case_233
	case 234:
		goto st_case_234
	case 771:
		goto st_case_771
	case 235:
		goto st_case_235
	case 236:
		goto st_case_236
	case 237:
		goto st_case_237
	case 238:
		goto st_case_238
	case 239:
		goto st_case_239
	case 240:
		goto st_case_240
	case 241:
		goto st_case_241
	case 242:
		goto st_case_242
	case 243:
		goto st_case_243
	case 244:
		goto st_case_244
	case 245:
		goto st_case_245
	case 246:
		goto st_case_246
	case 772:
		goto st_case_772
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 773:
		goto st_case_773
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 774:
		goto st_case_774
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 775:
		goto st_case_775
	case 776:
		goto st_case_776
	case 777:
		goto st_case_777
	case 778:
		goto st_case_778
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
	case 779:
		goto st_case_779
	case 259:
		goto st_case_259
	case 260:
		goto st_case_260
	case 261:
		goto st_case_261
	case 262:
		goto st_case_262
	case 263:
		goto st_case_263
	case 264:
		goto st_case_264
	case 265:
		goto st_case_265
	case 266:
		goto st_case_266
	case 267:
		goto st_case_267
	case 268:
		goto st_case_268
	case 269:
		goto st_case_269
	case 270:
		goto st_case_270
	case 780:
		goto st_case_780
	case 271:
		goto st_case_271
	case 272:
		goto st_case_272
	case 273:
		goto st_case_273
	case 274:
		goto st_case_274
	case 275:
		goto st_case_275
	case 276:
		goto st_case_276
	case 781:
		goto st_case_781
	case 277:
		goto st_case_277
	case 278:
		goto st_case_278
	case 279:
		goto st_case_279
	case 280:
		goto st_case_280
	case 281:
		goto st_case_281
	case 282:
		goto st_case_282
	case 283:
		goto st_case_283
	case 284:
		goto st_case_284
	case 782:
		goto st_case_782
	case 285:
		goto st_case_285
	case 286:
		goto st_case_286
	case 287:
		goto st_case_287
	case 288:
		goto st_case_288
	case 289:
		goto st_case_289
	case 290:
		goto st_case_290
	case 291:
		goto st_case_291
	case 292:
		goto st_case_292
	case 293:
		goto st_case_293
	case 294:
		goto st_case_294
	case 295:
		goto st_case_295
	case 296:
		goto st_case_296
	case 297:
		goto st_case_297
	case 298:
		goto st_case_298
	case 299:
		goto st_case_299
	case 300:
		goto st_case_300
	case 301:
		goto st_case_301
	case 302:
		goto st_case_302
	case 303:
		goto st_case_303
	case 304:
		goto st_case_304
	case 305:
		goto st_case_305
	case 306:
		goto st_case_306
	case 307:
		goto st_case_307
	case 308:
		goto st_case_308
	case 309:
		goto st_case_309
	case 310:
		goto st_case_310
	case 311:
		goto st_case_311
	case 312:
		goto st_case_312
	case 313:
		goto st_case_313
	case 314:
		goto st_case_314
	case 315:
		goto st_case_315
	case 316:
		goto st_case_316
	case 317:
		goto st_case_317
	case 318:
		goto st_case_318
	case 319:
		goto st_case_319
	case 320:
		goto st_case_320
	case 321:
		goto st_case_321
	case 322:
		goto st_case_322
	case 323:
		goto st_case_323
	case 324:
		goto st_case_324
	case 325:
		goto st_case_325
	case 326:
		goto st_case_326
	case 327:
		goto st_case_327
	case 328:
		goto st_case_328
	case 329:
		goto st_case_329
	case 330:
		goto st_case_330
	case 331:
		goto st_case_331
	case 332:
		goto st_case_332
	case 333:
		goto st_case_333
	case 334:
		goto st_case_334
	case 335:
		goto st_case_335
	case 336:
		goto st_case_336
	case 337:
		goto st_case_337
	case 338:
		goto st_case_338
	case 339:
		goto st_case_339
	case 340:
		goto st_case_340
	case 341:
		goto st_case_341
	case 342:
		goto st_case_342
	case 343:
		goto st_case_343
	case 344:
		goto st_case_344
	case 345:
		goto st_case_345
	case 346:
		goto st_case_346
	case 347:
		goto st_case_347
	case 348:
		goto st_case_348
	case 349:
		goto st_case_349
	case 350:
		goto st_case_350
	case 351:
		goto st_case_351
	case 352:
		goto st_case_352
	case 353:
		goto st_case_353
	case 354:
		goto st_case_354
	case 355:
		goto st_case_355
	case 356:
		goto st_case_356
	case 357:
		goto st_case_357
	case 358:
		goto st_case_358
	case 359:
		goto st_case_359
	case 360:
		goto st_case_360
	case 361:
		goto st_case_361
	case 362:
		goto st_case_362
	case 363:
		goto st_case_363
	case 364:
		goto st_case_364
	case 365:
		goto st_case_365
	case 366:
		goto st_case_366
	case 367:
		goto st_case_367
	case 368:
		goto st_case_368
	case 369:
		goto st_case_369
	case 370:
		goto st_case_370
	case 371:
		goto st_case_371
	case 372:
		goto st_case_372
	case 373:
		goto st_case_373
	case 374:
		goto st_case_374
	case 375:
		goto st_case_375
	case 376:
		goto st_case_376
	case 377:
		goto st_case_377
	case 378:
		goto st_case_378
	case 379:
		goto st_case_379
	case 380:
		goto st_case_380
	case 381:
		goto st_case_381
	case 382:
		goto st_case_382
	case 383:
		goto st_case_383
	case 384:
		goto st_case_384
	case 385:
		goto st_case_385
	case 386:
		goto st_case_386
	case 387:
		goto st_case_387
	case 388:
		goto st_case_388
	case 389:
		goto st_case_389
	case 390:
		goto st_case_390
	case 391:
		goto st_case_391
	case 392:
		goto st_case_392
	case 393:
		goto st_case_393
	case 394:
		goto st_case_394
	case 395:
		goto st_case_395
	case 396:
		goto st_case_396
	case 397:
		goto st_case_397
	case 398:
		goto st_case_398
	case 399:
		goto st_case_399
	case 400:
		goto st_case_400
	case 401:
		goto st_case_401
	case 402:
		goto st_case_402
	case 403:
		goto st_case_403
	case 404:
		goto st_case_404
	case 405:
		goto st_case_405
	case 406:
		goto st_case_406
	case 407:
		goto st_case_407
	case 408:
		goto st_case_408
	case 409:
		goto st_case_409
	case 410:
		goto st_case_410
	case 411:
		goto st_case_411
	case 412:
		goto st_case_412
	case 413:
		goto st_case_413
	case 414:
		goto st_case_414
	case 415:
		goto st_case_415
	case 416:
		goto st_case_416
	case 417:
		goto st_case_417
	case 418:
		goto st_case_418
	case 419:
		goto st_case_419
	case 420:
		goto st_case_420
	case 421:
		goto st_case_421
	case 422:
		goto st_case_422
	case 423:
		goto st_case_423
	case 424:
		goto st_case_424
	case 425:
		goto st_case_425
	case 426:
		goto st_case_426
	case 427:
		goto st_case_427
	case 428:
		goto st_case_428
	case 429:
		goto st_case_429
	case 430:
		goto st_case_430
	case 431:
		goto st_case_431
	case 432:
		goto st_case_432
	case 433:
		goto st_case_433
	case 434:
		goto st_case_434
	case 435:
		goto st_case_435
	case 436:
		goto st_case_436
	case 437:
		goto st_case_437
	case 438:
		goto st_case_438
	case 439:
		goto st_case_439
	case 440:
		goto st_case_440
	case 441:
		goto st_case_441
	case 442:
		goto st_case_442
	case 443:
		goto st_case_443
	case 444:
		goto st_case_444
	case 445:
		goto st_case_445
	case 446:
		goto st_case_446
	case 447:
		goto st_case_447
	case 448:
		goto st_case_448
	case 449:
		goto st_case_449
	case 450:
		goto st_case_450
	case 451:
		goto st_case_451
	case 452:
		goto st_case_452
	case 453:
		goto st_case_453
	case 454:
		goto st_case_454
	case 455:
		goto st_case_455
	case 456:
		goto st_case_456
	case 457:
		goto st_case_457
	case 458:
		goto st_case_458
	case 459:
		goto st_case_459
	case 460:
		goto st_case_460
	case 461:
		goto st_case_461
	case 462:
		goto st_case_462
	case 463:
		goto st_case_463
	case 464:
		goto st_case_464
	case 465:
		goto st_case_465
	case 466:
		goto st_case_466
	case 467:
		goto st_case_467
	case 468:
		goto st_case_468
	case 469:
		goto st_case_469
	case 470:
		goto st_case_470
	case 471:
		goto st_case_471
	case 472:
		goto st_case_472
	case 473:
		goto st_case_473
	case 474:
		goto st_case_474
	case 475:
		goto st_case_475
	case 476:
		goto st_case_476
	case 477:
		goto st_case_477
	case 478:
		goto st_case_478
	case 479:
		goto st_case_479
	case 480:
		goto st_case_480
	case 481:
		goto st_case_481
	case 482:
		goto st_case_482
	case 483:
		goto st_case_483
	case 484:
		goto st_case_484
	case 485:
		goto st_case_485
	case 486:
		goto st_case_486
	case 487:
		goto st_case_487
	case 488:
		goto st_case_488
	case 489:
		goto st_case_489
	case 490:
		goto st_case_490
	case 491:
		goto st_case_491
	case 492:
		goto st_case_492
	case 493:
		goto st_case_493
	case 494:
		goto st_case_494
	case 495:
		goto st_case_495
	case 496:
		goto st_case_496
	case 497:
		goto st_case_497
	case 498:
		goto st_case_498
	case 499:
		goto st_case_499
	case 500:
		goto st_case_500
	case 501:
		goto st_case_501
	case 502:
		goto st_case_502
	case 503:
		goto st_case_503
	case 504:
		goto st_case_504
	case 505:
		goto st_case_505
	case 506:
		goto st_case_506
	case 507:
		goto st_case_507
	case 508:
		goto st_case_508
	case 509:
		goto st_case_509
	case 510:
		goto st_case_510
	case 511:
		goto st_case_511
	case 512:
		goto st_case_512
	case 513:
		goto st_case_513
	case 514:
		goto st_case_514
	case 515:
		goto st_case_515
	case 516:
		goto st_case_516
	case 517:
		goto st_case_517
	case 518:
		goto st_case_518
	case 519:
		goto st_case_519
	case 520:
		goto st_case_520
	case 521:
		goto st_case_521
	case 522:
		goto st_case_522
	case 523:
		goto st_case_523
	case 524:
		goto st_case_524
	case 525:
		goto st_case_525
	case 526:
		goto st_case_526
	case 527:
		goto st_case_527
	case 528:
		goto st_case_528
	case 529:
		goto st_case_529
	case 530:
		goto st_case_530
	case 531:
		goto st_case_531
	case 532:
		goto st_case_532
	case 533:
		goto st_case_533
	case 534:
		goto st_case_534
	case 535:
		goto st_case_535
	case 536:
		goto st_case_536
	case 537:
		goto st_case_537
	case 538:
		goto st_case_538
	case 539:
		goto st_case_539
	case 540:
		goto st_case_540
	case 541:
		goto st_case_541
	case 542:
		goto st_case_542
	case 543:
		goto st_case_543
	case 544:
		goto st_case_544
	case 545:
		goto st_case_545
	case 546:
		goto st_case_546
	case 547:
		goto st_case_547
	case 548:
		goto st_case_548
	case 549:
		goto st_case_549
	case 550:
		goto st_case_550
	case 551:
		goto st_case_551
	case 552:
		goto st_case_552
	case 553:
		goto st_case_553
	case 554:
		goto st_case_554
	case 555:
		goto st_case_555
	case 556:
		goto st_case_556
	case 557:
		goto st_case_557
	case 558:
		goto st_case_558
	case 559:
		goto st_case_559
	case 560:
		goto st_case_560
	case 561:
		goto st_case_561
	case 562:
		goto st_case_562
	case 563:
		goto st_case_563
	case 564:
		goto st_case_564
	case 565:
		goto st_case_565
	case 566:
		goto st_case_566
	case 567:
		goto st_case_567
	case 568:
		goto st_case_568
	case 569:
		goto st_case_569
	case 570:
		goto st_case_570
	case 571:
		goto st_case_571
	case 572:
		goto st_case_572
	case 573:
		goto st_case_573
	case 574:
		goto st_case_574
	case 575:
		goto st_case_575
	case 576:
		goto st_case_576
	case 577:
		goto st_case_577
	case 578:
		goto st_case_578
	case 579:
		goto st_case_579
	case 580:
		goto st_case_580
	case 581:
		goto st_case_581
	case 582:
		goto st_case_582
	case 583:
		goto st_case_583
	case 584:
		goto st_case_584
	case 585:
		goto st_case_585
	case 586:
		goto st_case_586
	case 587:
		goto st_case_587
	case 588:
		goto st_case_588
	case 589:
		goto st_case_589
	case 590:
		goto st_case_590
	case 591:
		goto st_case_591
	case 592:
		goto st_case_592
	case 593:
		goto st_case_593
	case 594:
		goto st_case_594
	case 595:
		goto st_case_595
	case 596:
		goto st_case_596
	case 597:
		goto st_case_597
	case 598:
		goto st_case_598
	case 599:
		goto st_case_599
	case 600:
		goto st_case_600
	case 601:
		goto st_case_601
	case 602:
		goto st_case_602
	case 603:
		goto st_case_603
	case 604:
		goto st_case_604
	case 605:
		goto st_case_605
	case 606:
		goto st_case_606
	case 607:
		goto st_case_607
	case 608:
		goto st_case_608
	case 609:
		goto st_case_609
	case 610:
		goto st_case_610
	case 611:
		goto st_case_611
	case 612:
		goto st_case_612
	case 613:
		goto st_case_613
	case 614:
		goto st_case_614
	case 615:
		goto st_case_615
	case 616:
		goto st_case_616
	case 617:
		goto st_case_617
	case 618:
		goto st_case_618
	case 619:
		goto st_case_619
	case 620:
		goto st_case_620
	case 621:
		goto st_case_621
	case 622:
		goto st_case_622
	case 623:
		goto st_case_623
	case 624:
		goto st_case_624
	case 625:
		goto st_case_625
	case 626:
		goto st_case_626
	case 627:
		goto st_case_627
	case 628:
		goto st_case_628
	case 629:
		goto st_case_629
	case 630:
		goto st_case_630
	case 631:
		goto st_case_631
	case 632:
		goto st_case_632
	case 633:
		goto st_case_633
	case 634:
		goto st_case_634
	case 635:
		goto st_case_635
	case 636:
		goto st_case_636
	case 637:
		goto st_case_637
	case 638:
		goto st_case_638
	case 639:
		goto st_case_639
	case 640:
		goto st_case_640
	case 641:
		goto st_case_641
	case 642:
		goto st_case_642
	case 643:
		goto st_case_643
	case 644:
		goto st_case_644
	case 645:
		goto st_case_645
	case 646:
		goto st_case_646
	case 647:
		goto st_case_647
	case 648:
		goto st_case_648
	case 649:
		goto st_case_649
	case 650:
		goto st_case_650
	case 651:
		goto st_case_651
	case 652:
		goto st_case_652
	case 653:
		goto st_case_653
	case 654:
		goto st_case_654
	case 655:
		goto st_case_655
	case 656:
		goto st_case_656
	case 657:
		goto st_case_657
	case 658:
		goto st_case_658
	case 659:
		goto st_case_659
	case 660:
		goto st_case_660
	case 661:
		goto st_case_661
	case 662:
		goto st_case_662
	case 663:
		goto st_case_663
	case 664:
		goto st_case_664
	case 665:
		goto st_case_665
	case 666:
		goto st_case_666
	case 667:
		goto st_case_667
	case 668:
		goto st_case_668
	case 669:
		goto st_case_669
	case 670:
		goto st_case_670
	case 671:
		goto st_case_671
	case 672:
		goto st_case_672
	case 673:
		goto st_case_673
	case 674:
		goto st_case_674
	case 675:
		goto st_case_675
	case 676:
		goto st_case_676
	case 677:
		goto st_case_677
	case 678:
		goto st_case_678
	case 679:
		goto st_case_679
	case 680:
		goto st_case_680
	case 681:
		goto st_case_681
	case 682:
		goto st_case_682
	case 683:
		goto st_case_683
	case 684:
		goto st_case_684
	case 685:
		goto st_case_685
	case 686:
		goto st_case_686
	case 687:
		goto st_case_687
	case 688:
		goto st_case_688
	case 689:
		goto st_case_689
	case 690:
		goto st_case_690
	case 691:
		goto st_case_691
	case 692:
		goto st_case_692
	case 693:
		goto st_case_693
	case 694:
		goto st_case_694
	case 695:
		goto st_case_695
	case 696:
		goto st_case_696
	case 697:
		goto st_case_697
	case 698:
		goto st_case_698
	case 699:
		goto st_case_699
	case 700:
		goto st_case_700
	case 701:
		goto st_case_701
	case 702:
		goto st_case_702
	case 703:
		goto st_case_703
	case 704:
		goto st_case_704
	case 705:
		goto st_case_705
	case 706:
		goto st_case_706
	case 707:
		goto st_case_707
	case 708:
		goto st_case_708
	case 709:
		goto st_case_709
	case 710:
		goto st_case_710
	case 711:
		goto st_case_711
	case 712:
		goto st_case_712
	case 713:
		goto st_case_713
	case 714:
		goto st_case_714
	case 715:
		goto st_case_715
	case 716:
		goto st_case_716
	case 717:
		goto st_case_717
	case 718:
		goto st_case_718
	case 719:
		goto st_case_719
	case 720:
		goto st_case_720
	case 721:
		goto st_case_721
	case 722:
		goto st_case_722
	case 723:
		goto st_case_723
	case 724:
		goto st_case_724
	case 725:
		goto st_case_725
	case 726:
		goto st_case_726
	case 727:
		goto st_case_727
	case 728:
		goto st_case_728
	case 729:
		goto st_case_729
	case 730:
		goto st_case_730
	case 731:
		goto st_case_731
	case 732:
		goto st_case_732
	case 733:
		goto st_case_733
	case 734:
		goto st_case_734
	case 735:
		goto st_case_735
	case 736:
		goto st_case_736
	case 737:
		goto st_case_737
	case 738:
		goto st_case_738
	case 739:
		goto st_case_739
	case 740:
		goto st_case_740
	case 741:
		goto st_case_741
	case 742:
		goto st_case_742
	case 743:
		goto st_case_743
	case 744:
		goto st_case_744
	case 745:
		goto st_case_745
	case 746:
		goto st_case_746
	case 747:
		goto st_case_747
	case 748:
		goto st_case_748
	case 749:
		goto st_case_749
	case 750:
		goto st_case_750
	case 751:
		goto st_case_751
	case 752:
		goto st_case_752
	case 753:
		goto st_case_753
	case 754:
		goto st_case_754
	case 755:
		goto st_case_755
	case 756:
		goto st_case_756
	case 757:
		goto st_case_757
	case 758:
		goto st_case_758
	case 759:
		goto st_case_759
	case 760:
		goto st_case_760
	case 761:
		goto st_case_761
	case 762:
		goto st_case_762
	case 763:
		goto st_case_763
	case 764:
		goto st_case_764
	}
	goto st_out
	st_case_1:
		switch data[p] {
		case 33:
			goto tr0
		case 37:
			goto tr0
		case 39:
			goto tr0
		case 83:
			goto tr2
		case 126:
			goto tr0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr0
				}
			case data[p] >= 42:
				goto tr0
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr0
				}
			case data[p] >= 65:
				goto tr0
			}
		default:
			goto tr0
		}
		goto st0
tr416:
//line sip.rl:156

	p--

	{goto st273 }

	goto st0
//line msg_parse.go:1696
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line sip.rl:62

	mark = p

	goto st2
	st2:
		p++
	st_case_2:
//line msg_parse.go:1710
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st2
		case 37:
			goto st2
		case 39:
			goto st2
		case 126:
			goto st2
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st2
				}
			case data[p] >= 42:
				goto st2
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st2
				}
			case data[p] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
tr3:
//line sip.rl:94

	msg.Method = string(data[mark:p])

	goto st3
	st3:
		p++
	st_case_3:
//line msg_parse.go:1755
		if data[p] == 32 {
			goto st0
		}
		goto tr5
tr5:
//line sip.rl:62

	mark = p

	goto st4
	st4:
		p++
	st_case_4:
//line msg_parse.go:1769
		if data[p] == 32 {
			goto tr7
		}
		goto st4
tr7:
//line sip.rl:106

	msg.Request, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st5
	st5:
		p++
	st_case_5:
//line msg_parse.go:1784
		if data[p] == 83 {
			goto st6
		}
		goto st0
	st6:
		p++
	st_case_6:
		if data[p] == 73 {
			goto st7
		}
		goto st0
	st7:
		p++
	st_case_7:
		if data[p] == 80 {
			goto st8
		}
		goto st0
	st8:
		p++
	st_case_8:
		if data[p] == 47 {
			goto st9
		}
		goto st0
	st9:
		p++
	st_case_9:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr12
		}
		goto st0
tr12:
//line sip.rl:98

	msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)

	goto st10
	st10:
		p++
	st_case_10:
//line msg_parse.go:1826
		if data[p] == 46 {
			goto st11
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr12
		}
		goto st0
	st11:
		p++
	st_case_11:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
tr14:
//line sip.rl:102

	msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)

	goto st12
	st12:
		p++
	st_case_12:
//line msg_parse.go:1850
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st13
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr14
		}
		goto st0
tr42:
//line sip.rl:115

	msg.Phrase = string(buf[0:amt])

	goto st13
	st13:
		p++
	st_case_13:
//line msg_parse.go:1874
		if data[p] == 10 {
			goto tr16
		}
		goto st0
tr16:
//line sip.rl:248
 {goto st280 } 
	goto st765
	st765:
		p++
	st_case_765:
//line msg_parse.go:1886
		goto st0
tr2:
//line sip.rl:62

	mark = p

	goto st14
	st14:
		p++
	st_case_14:
//line msg_parse.go:1897
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st2
		case 37:
			goto st2
		case 39:
			goto st2
		case 73:
			goto st15
		case 126:
			goto st2
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st2
				}
			case data[p] >= 42:
				goto st2
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st2
				}
			case data[p] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st15:
		p++
	st_case_15:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st2
		case 37:
			goto st2
		case 39:
			goto st2
		case 80:
			goto st16
		case 126:
			goto st2
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st2
				}
			case data[p] >= 42:
				goto st2
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st2
				}
			case data[p] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st16:
		p++
	st_case_16:
		switch data[p] {
		case 32:
			goto tr3
		case 33:
			goto st2
		case 37:
			goto st2
		case 39:
			goto st2
		case 47:
			goto st17
		case 126:
			goto st2
		}
		switch {
		case data[p] < 45:
			if 42 <= data[p] && data[p] <= 43 {
				goto st2
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st2
				}
			case data[p] >= 65:
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st17:
		p++
	st_case_17:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr20
		}
		goto st0
tr20:
//line sip.rl:98

	msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)

	goto st18
	st18:
		p++
	st_case_18:
//line msg_parse.go:2026
		if data[p] == 46 {
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr20
		}
		goto st0
	st19:
		p++
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr22
		}
		goto st0
tr22:
//line sip.rl:102

	msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)

	goto st20
	st20:
		p++
	st_case_20:
//line msg_parse.go:2050
		if data[p] == 32 {
			goto st21
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr22
		}
		goto st0
	st21:
		p++
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr24
		}
		goto st0
tr24:
//line sip.rl:111

	msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)

	goto st22
	st22:
		p++
	st_case_22:
//line msg_parse.go:2074
		if 48 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
tr25:
//line sip.rl:111

	msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)

	goto st23
	st23:
		p++
	st_case_23:
//line msg_parse.go:2088
		if 48 <= data[p] && data[p] <= 57 {
			goto tr26
		}
		goto st0
tr26:
//line sip.rl:111

	msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)

	goto st24
	st24:
		p++
	st_case_24:
//line msg_parse.go:2102
		if data[p] == 32 {
			goto st25
		}
		goto st0
	st25:
		p++
	st_case_25:
		switch data[p] {
		case 9:
			goto tr28
		case 37:
			goto tr29
		case 61:
			goto tr28
		case 95:
			goto tr28
		case 126:
			goto tr28
		}
		switch {
		case data[p] < 192:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 33 {
					goto tr28
				}
			case data[p] > 59:
				switch {
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr28
					}
				case data[p] >= 63:
					goto tr28
				}
			default:
				goto tr28
			}
		case data[p] > 223:
			switch {
			case data[p] < 240:
				if 224 <= data[p] && data[p] <= 239 {
					goto tr31
				}
			case data[p] > 247:
				switch {
				case data[p] > 251:
					if 252 <= data[p] && data[p] <= 253 {
						goto tr34
					}
				case data[p] >= 248:
					goto tr33
				}
			default:
				goto tr32
			}
		default:
			goto tr30
		}
		goto st0
tr28:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st26
tr35:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st26
tr44:
//line sip.rl:88

	hex += unhex(data[p])
	buf[amt] = hex
	amt++

	goto st26
	st26:
		p++
	st_case_26:
//line msg_parse.go:2192
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr35
		case 37:
			goto st27
		case 61:
			goto tr35
		case 95:
			goto tr35
		case 126:
			goto tr35
		case 269:
			goto tr42
		}
		switch {
		case _widec < 192:
			switch {
			case _widec < 36:
				if 32 <= _widec && _widec <= 33 {
					goto tr35
				}
			case _widec > 59:
				switch {
				case _widec > 90:
					if 97 <= _widec && _widec <= 122 {
						goto tr35
					}
				case _widec >= 63:
					goto tr35
				}
			default:
				goto tr35
			}
		case _widec > 223:
			switch {
			case _widec < 240:
				if 224 <= _widec && _widec <= 239 {
					goto tr38
				}
			case _widec > 247:
				switch {
				case _widec > 251:
					if 252 <= _widec && _widec <= 253 {
						goto tr41
					}
				case _widec >= 248:
					goto tr40
				}
			default:
				goto tr39
			}
		default:
			goto tr37
		}
		goto st0
tr29:
//line sip.rl:70

	amt = 0

	goto st27
	st27:
		p++
	st_case_27:
//line msg_parse.go:2264
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr43
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr43
			}
		default:
			goto tr43
		}
		goto st0
tr43:
//line sip.rl:84

	hex = unhex(data[p]) * 16

	goto st28
	st28:
		p++
	st_case_28:
//line msg_parse.go:2287
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr44
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr44
			}
		default:
			goto tr44
		}
		goto st0
tr30:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st29
tr37:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st29
	st29:
		p++
	st_case_29:
//line msg_parse.go:2322
		if 128 <= data[p] && data[p] <= 191 {
			goto tr35
		}
		goto st0
tr31:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st30
tr38:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st30
	st30:
		p++
	st_case_30:
//line msg_parse.go:2348
		if 128 <= data[p] && data[p] <= 191 {
			goto tr37
		}
		goto st0
tr32:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st31
tr39:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st31
	st31:
		p++
	st_case_31:
//line msg_parse.go:2374
		if 128 <= data[p] && data[p] <= 191 {
			goto tr38
		}
		goto st0
tr33:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st32
tr40:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st32
	st32:
		p++
	st_case_32:
//line msg_parse.go:2400
		if 128 <= data[p] && data[p] <= 191 {
			goto tr39
		}
		goto st0
tr34:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st33
tr41:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st33
	st33:
		p++
	st_case_33:
//line msg_parse.go:2426
		if 128 <= data[p] && data[p] <= 191 {
			goto tr40
		}
		goto st0
	st34:
		p++
	st_case_34:
		switch data[p] {
		case 33:
			goto tr45
		case 37:
			goto tr45
		case 39:
			goto tr45
		case 126:
			goto tr45
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr45
				}
			case data[p] >= 42:
				goto tr45
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr45
				}
			case data[p] >= 65:
				goto tr45
			}
		default:
			goto tr45
		}
		goto st0
tr45:
//line sip.rl:62

	mark = p

	goto st35
	st35:
		p++
	st_case_35:
//line msg_parse.go:2476
		switch data[p] {
		case 33:
			goto st35
		case 37:
			goto st35
		case 39:
			goto st35
		case 47:
			goto st36
		case 126:
			goto st35
		}
		switch {
		case data[p] < 45:
			if 42 <= data[p] && data[p] <= 43 {
				goto st35
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st35
				}
			case data[p] >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
	st36:
		p++
	st_case_36:
		switch data[p] {
		case 33:
			goto st37
		case 37:
			goto st37
		case 39:
			goto st37
		case 126:
			goto st37
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st37
				}
			case data[p] >= 42:
				goto st37
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st37
				}
			case data[p] >= 65:
				goto st37
			}
		default:
			goto st37
		}
		goto st0
	st37:
		p++
	st_case_37:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr49
		case 32:
			goto tr49
		case 33:
			goto st37
		case 37:
			goto st37
		case 39:
			goto st37
		case 59:
			goto tr50
		case 126:
			goto st37
		case 269:
			goto tr51
		case 525:
			goto tr52
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st37
				}
			case _widec >= 42:
				goto st37
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st37
				}
			case _widec >= 65:
				goto st37
			}
		default:
			goto st37
		}
		goto st0
tr49:
//line sip.rl:219

	ctype = string(data[mark:p])

	goto st38
	st38:
		p++
	st_case_38:
//line msg_parse.go:2605
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st38
		case 32:
			goto st38
		case 59:
			goto st39
		case 525:
			goto st45
		}
		goto st0
tr50:
//line sip.rl:219

	ctype = string(data[mark:p])

	goto st39
	st39:
		p++
	st_case_39:
//line msg_parse.go:2633
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st39
		case 32:
			goto st39
		case 33:
			goto st40
		case 37:
			goto st40
		case 39:
			goto st40
		case 126:
			goto st40
		case 525:
			goto st65
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st40
				}
			case _widec >= 42:
				goto st40
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st40
				}
			case _widec >= 65:
				goto st40
			}
		default:
			goto st40
		}
		goto st0
	st40:
		p++
	st_case_40:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st41
		case 32:
			goto st41
		case 33:
			goto st40
		case 37:
			goto st40
		case 39:
			goto st40
		case 61:
			goto st42
		case 126:
			goto st40
		case 525:
			goto st62
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st40
				}
			case _widec >= 42:
				goto st40
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st40
				}
			case _widec >= 65:
				goto st40
			}
		default:
			goto st40
		}
		goto st0
	st41:
		p++
	st_case_41:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st41
		case 32:
			goto st41
		case 61:
			goto st42
		case 525:
			goto st62
		}
		goto st0
	st42:
		p++
	st_case_42:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st42
		case 32:
			goto st42
		case 33:
			goto st43
		case 34:
			goto st48
		case 37:
			goto st43
		case 39:
			goto st43
		case 126:
			goto st43
		case 525:
			goto st59
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st43
				}
			case _widec >= 42:
				goto st43
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st43
				}
			case _widec >= 65:
				goto st43
			}
		default:
			goto st43
		}
		goto st0
	st43:
		p++
	st_case_43:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st38
		case 32:
			goto st38
		case 33:
			goto st43
		case 37:
			goto st43
		case 39:
			goto st43
		case 59:
			goto st39
		case 126:
			goto st43
		case 269:
			goto st44
		case 525:
			goto st45
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st43
				}
			case _widec >= 42:
				goto st43
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st43
				}
			case _widec >= 65:
				goto st43
			}
		default:
			goto st43
		}
		goto st0
tr51:
//line sip.rl:219

	ctype = string(data[mark:p])

	goto st44
	st44:
		p++
	st_case_44:
//line msg_parse.go:2865
		if data[p] == 10 {
			goto tr65
		}
		goto st0
tr65:
//line sip.rl:248
 {goto st280 } 
	goto st766
	st766:
		p++
	st_case_766:
//line msg_parse.go:2877
		goto st0
tr52:
//line sip.rl:219

	ctype = string(data[mark:p])

	goto st45
	st45:
		p++
	st_case_45:
//line msg_parse.go:2888
		if data[p] == 10 {
			goto st46
		}
		goto st0
	st46:
		p++
	st_case_46:
		switch data[p] {
		case 9:
			goto st47
		case 32:
			goto st47
		}
		goto st0
	st47:
		p++
	st_case_47:
		switch data[p] {
		case 9:
			goto st47
		case 32:
			goto st47
		case 59:
			goto st39
		}
		goto st0
	st48:
		p++
	st_case_48:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr68
		case 34:
			goto tr69
		case 92:
			goto tr70
		case 525:
			goto tr76
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr71
				}
			case _widec >= 32:
				goto tr68
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr73
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr75
				}
			default:
				goto tr74
			}
		default:
			goto tr72
		}
		goto st0
tr68:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st49
tr77:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st49
	st49:
		p++
	st_case_49:
//line msg_parse.go:2983
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr77
		case 34:
			goto st50
		case 92:
			goto st51
		case 525:
			goto tr85
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr80
				}
			case _widec >= 32:
				goto tr77
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr82
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr84
				}
			default:
				goto tr83
			}
		default:
			goto tr81
		}
		goto st0
tr69:
//line sip.rl:70

	amt = 0

	goto st50
	st50:
		p++
	st_case_50:
//line msg_parse.go:3037
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st38
		case 32:
			goto st38
		case 59:
			goto st39
		case 269:
			goto st44
		case 525:
			goto st45
		}
		goto st0
tr70:
//line sip.rl:70

	amt = 0

	goto st51
	st51:
		p++
	st_case_51:
//line msg_parse.go:3067
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr77
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr77
			}
		default:
			goto tr77
		}
		goto st0
tr71:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st52
tr80:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st52
	st52:
		p++
	st_case_52:
//line msg_parse.go:3102
		if 128 <= data[p] && data[p] <= 191 {
			goto tr77
		}
		goto st0
tr72:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st53
tr81:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st53
	st53:
		p++
	st_case_53:
//line msg_parse.go:3128
		if 128 <= data[p] && data[p] <= 191 {
			goto tr80
		}
		goto st0
tr73:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st54
tr82:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st54
	st54:
		p++
	st_case_54:
//line msg_parse.go:3154
		if 128 <= data[p] && data[p] <= 191 {
			goto tr81
		}
		goto st0
tr74:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st55
tr83:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st55
	st55:
		p++
	st_case_55:
//line msg_parse.go:3180
		if 128 <= data[p] && data[p] <= 191 {
			goto tr82
		}
		goto st0
tr75:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st56
tr84:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st56
	st56:
		p++
	st_case_56:
//line msg_parse.go:3206
		if 128 <= data[p] && data[p] <= 191 {
			goto tr83
		}
		goto st0
tr76:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st57
tr85:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st57
	st57:
		p++
	st_case_57:
//line msg_parse.go:3232
		if data[p] == 10 {
			goto tr86
		}
		goto st0
tr86:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st58
	st58:
		p++
	st_case_58:
//line msg_parse.go:3247
		switch data[p] {
		case 9:
			goto tr77
		case 32:
			goto tr77
		}
		goto st0
	st59:
		p++
	st_case_59:
		if data[p] == 10 {
			goto st60
		}
		goto st0
	st60:
		p++
	st_case_60:
		switch data[p] {
		case 9:
			goto st61
		case 32:
			goto st61
		}
		goto st0
	st61:
		p++
	st_case_61:
		switch data[p] {
		case 9:
			goto st61
		case 32:
			goto st61
		case 33:
			goto st43
		case 34:
			goto st48
		case 37:
			goto st43
		case 39:
			goto st43
		case 126:
			goto st43
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st43
				}
			case data[p] >= 42:
				goto st43
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st43
				}
			case data[p] >= 65:
				goto st43
			}
		default:
			goto st43
		}
		goto st0
	st62:
		p++
	st_case_62:
		if data[p] == 10 {
			goto st63
		}
		goto st0
	st63:
		p++
	st_case_63:
		switch data[p] {
		case 9:
			goto st64
		case 32:
			goto st64
		}
		goto st0
	st64:
		p++
	st_case_64:
		switch data[p] {
		case 9:
			goto st64
		case 32:
			goto st64
		case 61:
			goto st42
		}
		goto st0
	st65:
		p++
	st_case_65:
		if data[p] == 10 {
			goto st66
		}
		goto st0
	st66:
		p++
	st_case_66:
		switch data[p] {
		case 9:
			goto st67
		case 32:
			goto st67
		}
		goto st0
	st67:
		p++
	st_case_67:
		switch data[p] {
		case 9:
			goto st67
		case 32:
			goto st67
		case 33:
			goto st40
		case 37:
			goto st40
		case 39:
			goto st40
		case 126:
			goto st40
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st40
				}
			case data[p] >= 42:
				goto st40
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st40
				}
			case data[p] >= 65:
				goto st40
			}
		default:
			goto st40
		}
		goto st0
	st68:
		p++
	st_case_68:
		switch data[p] {
		case 33:
			goto tr93
		case 37:
			goto tr93
		case 39:
			goto tr93
		case 126:
			goto tr93
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr93
				}
			case data[p] >= 42:
				goto tr93
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr93
				}
			case data[p] >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
tr93:
//line sip.rl:70

	amt = 0

//line sip.rl:62

	mark = p

	goto st69
	st69:
		p++
	st_case_69:
//line msg_parse.go:3449
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr94
		case 32:
			goto tr94
		case 33:
			goto st69
		case 37:
			goto st69
		case 39:
			goto st69
		case 44:
			goto tr96
		case 59:
			goto tr97
		case 61:
			goto tr98
		case 126:
			goto st69
		case 269:
			goto tr99
		case 525:
			goto tr100
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto st69
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st69
				}
			case _widec >= 65:
				goto st69
			}
		default:
			goto st69
		}
		goto st0
tr94:
//line sip.rl:161

	name = string(data[mark:p])

	goto st70
	st70:
		p++
	st_case_70:
//line msg_parse.go:3508
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st70
		case 32:
			goto st70
		case 44:
			goto st71
		case 59:
			goto st75
		case 61:
			goto st79
		case 525:
			goto st100
		}
		goto st0
tr96:
//line sip.rl:161

	name = string(data[mark:p])

	goto st71
	st71:
		p++
	st_case_71:
//line msg_parse.go:3540
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st71
		case 32:
			goto st71
		case 269:
			goto tr106
		case 525:
			goto st72
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
tr106:
//line sip.rl:149

	if via.Params == nil {
		via.Params = Params{}
	}
	via.Params[name] = string(buf[0:amt])

//line sip.rl:123

	*viap = via
	viap = &via.Next
	via = nil

//line sip.rl:119

	via = new(Via)

//line sip.rl:54

	p--


//line sip.rl:250
 {goto st103 } 
	goto st767
tr110:
//line sip.rl:149

	if via.Params == nil {
		via.Params = Params{}
	}
	via.Params[name] = string(buf[0:amt])

//line sip.rl:54

	p--


//line sip.rl:70

	amt = 0

//line sip.rl:251
 {goto st68 } 
	goto st767
tr122:
//line sip.rl:149

	if via.Params == nil {
		via.Params = Params{}
	}
	via.Params[name] = string(buf[0:amt])

//line sip.rl:123

	*viap = via
	viap = &via.Next
	via = nil

//line sip.rl:248
 {goto st280 } 
	goto st767
	st767:
		p++
	st_case_767:
//line msg_parse.go:3633
		goto st0
	st72:
		p++
	st_case_72:
		if data[p] == 10 {
			goto st73
		}
		goto st0
	st73:
		p++
	st_case_73:
		switch data[p] {
		case 9:
			goto st74
		case 32:
			goto st74
		}
		goto st0
	st74:
		p++
	st_case_74:
		switch data[p] {
		case 9:
			goto st74
		case 32:
			goto st74
		}
		goto tr106
tr97:
//line sip.rl:161

	name = string(data[mark:p])

	goto st75
	st75:
		p++
	st_case_75:
//line msg_parse.go:3671
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st75
		case 32:
			goto st75
		case 269:
			goto tr110
		case 525:
			goto st76
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr110
			}
		default:
			goto tr110
		}
		goto st0
	st76:
		p++
	st_case_76:
		if data[p] == 10 {
			goto st77
		}
		goto st0
	st77:
		p++
	st_case_77:
		switch data[p] {
		case 9:
			goto st78
		case 32:
			goto st78
		}
		goto st0
	st78:
		p++
	st_case_78:
		switch data[p] {
		case 9:
			goto st78
		case 32:
			goto st78
		}
		goto tr110
tr98:
//line sip.rl:161

	name = string(data[mark:p])

	goto st79
	st79:
		p++
	st_case_79:
//line msg_parse.go:3734
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st79
		case 32:
			goto st79
		case 33:
			goto tr114
		case 34:
			goto st86
		case 37:
			goto tr114
		case 39:
			goto tr114
		case 93:
			goto tr114
		case 126:
			goto tr114
		case 525:
			goto st97
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr114
				}
			case _widec >= 42:
				goto tr114
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr114
				}
			case _widec >= 65:
				goto tr114
			}
		default:
			goto tr114
		}
		goto st0
tr114:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st80
	st80:
		p++
	st_case_80:
//line msg_parse.go:3795
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st81
		case 32:
			goto st81
		case 33:
			goto tr114
		case 37:
			goto tr114
		case 39:
			goto tr114
		case 44:
			goto st71
		case 59:
			goto st75
		case 93:
			goto tr114
		case 126:
			goto tr114
		case 269:
			goto st85
		case 525:
			goto st82
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr114
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr114
				}
			case _widec >= 65:
				goto tr114
			}
		default:
			goto tr114
		}
		goto st0
	st81:
		p++
	st_case_81:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st81
		case 32:
			goto st81
		case 44:
			goto st71
		case 59:
			goto st75
		case 525:
			goto st82
		}
		goto st0
	st82:
		p++
	st_case_82:
		if data[p] == 10 {
			goto st83
		}
		goto st0
	st83:
		p++
	st_case_83:
		switch data[p] {
		case 9:
			goto st84
		case 32:
			goto st84
		}
		goto st0
	st84:
		p++
	st_case_84:
		switch data[p] {
		case 9:
			goto st84
		case 32:
			goto st84
		case 44:
			goto st71
		case 59:
			goto st75
		}
		goto st0
tr99:
//line sip.rl:161

	name = string(data[mark:p])

	goto st85
	st85:
		p++
	st_case_85:
//line msg_parse.go:3908
		if data[p] == 10 {
			goto tr122
		}
		goto st0
	st86:
		p++
	st_case_86:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr123
		case 34:
			goto tr124
		case 92:
			goto tr125
		case 525:
			goto tr131
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr126
				}
			case _widec >= 32:
				goto tr123
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr128
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr130
				}
			default:
				goto tr129
			}
		default:
			goto tr127
		}
		goto st0
tr123:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st87
tr132:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st87
	st87:
		p++
	st_case_87:
//line msg_parse.go:3981
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr132
		case 34:
			goto st88
		case 92:
			goto st89
		case 525:
			goto tr140
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr135
				}
			case _widec >= 32:
				goto tr132
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr137
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr139
				}
			default:
				goto tr138
			}
		default:
			goto tr136
		}
		goto st0
tr124:
//line sip.rl:70

	amt = 0

	goto st88
	st88:
		p++
	st_case_88:
//line msg_parse.go:4035
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st81
		case 32:
			goto st81
		case 44:
			goto st71
		case 59:
			goto st75
		case 269:
			goto st85
		case 525:
			goto st82
		}
		goto st0
tr125:
//line sip.rl:70

	amt = 0

	goto st89
	st89:
		p++
	st_case_89:
//line msg_parse.go:4067
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr132
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr132
			}
		default:
			goto tr132
		}
		goto st0
tr126:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st90
tr135:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st90
	st90:
		p++
	st_case_90:
//line msg_parse.go:4102
		if 128 <= data[p] && data[p] <= 191 {
			goto tr132
		}
		goto st0
tr127:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st91
tr136:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st91
	st91:
		p++
	st_case_91:
//line msg_parse.go:4128
		if 128 <= data[p] && data[p] <= 191 {
			goto tr135
		}
		goto st0
tr128:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st92
tr137:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st92
	st92:
		p++
	st_case_92:
//line msg_parse.go:4154
		if 128 <= data[p] && data[p] <= 191 {
			goto tr136
		}
		goto st0
tr129:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st93
tr138:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st93
	st93:
		p++
	st_case_93:
//line msg_parse.go:4180
		if 128 <= data[p] && data[p] <= 191 {
			goto tr137
		}
		goto st0
tr130:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st94
tr139:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st94
	st94:
		p++
	st_case_94:
//line msg_parse.go:4206
		if 128 <= data[p] && data[p] <= 191 {
			goto tr138
		}
		goto st0
tr131:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st95
tr140:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st95
	st95:
		p++
	st_case_95:
//line msg_parse.go:4232
		if data[p] == 10 {
			goto tr141
		}
		goto st0
tr141:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st96
	st96:
		p++
	st_case_96:
//line msg_parse.go:4247
		switch data[p] {
		case 9:
			goto tr132
		case 32:
			goto tr132
		}
		goto st0
	st97:
		p++
	st_case_97:
		if data[p] == 10 {
			goto st98
		}
		goto st0
	st98:
		p++
	st_case_98:
		switch data[p] {
		case 9:
			goto st99
		case 32:
			goto st99
		}
		goto st0
	st99:
		p++
	st_case_99:
		switch data[p] {
		case 9:
			goto st99
		case 32:
			goto st99
		case 33:
			goto tr114
		case 34:
			goto st86
		case 37:
			goto tr114
		case 39:
			goto tr114
		case 93:
			goto tr114
		case 126:
			goto tr114
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr114
				}
			case data[p] >= 42:
				goto tr114
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr114
				}
			case data[p] >= 65:
				goto tr114
			}
		default:
			goto tr114
		}
		goto st0
tr100:
//line sip.rl:161

	name = string(data[mark:p])

	goto st100
	st100:
		p++
	st_case_100:
//line msg_parse.go:4325
		if data[p] == 10 {
			goto st101
		}
		goto st0
	st101:
		p++
	st_case_101:
		switch data[p] {
		case 9:
			goto st102
		case 32:
			goto st102
		}
		goto st0
	st102:
		p++
	st_case_102:
		switch data[p] {
		case 9:
			goto st102
		case 32:
			goto st102
		case 44:
			goto st71
		case 59:
			goto st75
		case 61:
			goto st79
		}
		goto st0
	st103:
		p++
	st_case_103:
		switch data[p] {
		case 33:
			goto tr146
		case 37:
			goto tr146
		case 39:
			goto tr146
		case 126:
			goto tr146
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr146
				}
			case data[p] >= 42:
				goto tr146
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr146
				}
			case data[p] >= 65:
				goto tr146
			}
		default:
			goto tr146
		}
		goto st0
tr146:
//line sip.rl:62

	mark = p

	goto st104
	st104:
		p++
	st_case_104:
//line msg_parse.go:4401
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr147
		case 32:
			goto tr147
		case 33:
			goto st104
		case 37:
			goto st104
		case 39:
			goto st104
		case 47:
			goto tr149
		case 126:
			goto st104
		case 525:
			goto tr150
		}
		switch {
		case _widec < 45:
			if 42 <= _widec && _widec <= 43 {
				goto st104
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st104
				}
			case _widec >= 65:
				goto st104
			}
		default:
			goto st104
		}
		goto st0
tr147:
//line sip.rl:129

	via.Protocol = string(data[mark:p])

	goto st105
	st105:
		p++
	st_case_105:
//line msg_parse.go:4454
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st105
		case 32:
			goto st105
		case 47:
			goto st106
		case 525:
			goto st150
		}
		goto st0
tr149:
//line sip.rl:129

	via.Protocol = string(data[mark:p])

	goto st106
	st106:
		p++
	st_case_106:
//line msg_parse.go:4482
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st106
		case 32:
			goto st106
		case 33:
			goto tr154
		case 37:
			goto tr154
		case 39:
			goto tr154
		case 126:
			goto tr154
		case 525:
			goto st147
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr154
				}
			case _widec >= 42:
				goto tr154
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr154
				}
			case _widec >= 65:
				goto tr154
			}
		default:
			goto tr154
		}
		goto st0
tr154:
//line sip.rl:62

	mark = p

	goto st107
	st107:
		p++
	st_case_107:
//line msg_parse.go:4538
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr156
		case 32:
			goto tr156
		case 33:
			goto st107
		case 37:
			goto st107
		case 39:
			goto st107
		case 47:
			goto tr158
		case 126:
			goto st107
		case 525:
			goto tr159
		}
		switch {
		case _widec < 45:
			if 42 <= _widec && _widec <= 43 {
				goto st107
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st107
				}
			case _widec >= 65:
				goto st107
			}
		default:
			goto st107
		}
		goto st0
tr156:
//line sip.rl:133

	via.Version = string(data[mark:p])

	goto st108
	st108:
		p++
	st_case_108:
//line msg_parse.go:4591
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st108
		case 32:
			goto st108
		case 47:
			goto st109
		case 525:
			goto st144
		}
		goto st0
tr158:
//line sip.rl:133

	via.Version = string(data[mark:p])

	goto st109
	st109:
		p++
	st_case_109:
//line msg_parse.go:4619
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st109
		case 32:
			goto st109
		case 33:
			goto tr163
		case 37:
			goto tr163
		case 39:
			goto tr163
		case 126:
			goto tr163
		case 525:
			goto st141
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr163
				}
			case _widec >= 42:
				goto tr163
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr163
				}
			case _widec >= 65:
				goto tr163
			}
		default:
			goto tr163
		}
		goto st0
tr163:
//line sip.rl:62

	mark = p

	goto st110
	st110:
		p++
	st_case_110:
//line msg_parse.go:4675
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr165
		case 32:
			goto tr165
		case 33:
			goto st110
		case 37:
			goto st110
		case 39:
			goto st110
		case 126:
			goto st110
		case 525:
			goto tr167
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st110
				}
			case _widec >= 42:
				goto st110
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st110
				}
			case _widec >= 65:
				goto st110
			}
		default:
			goto st110
		}
		goto st0
tr165:
//line sip.rl:137

	via.Transport = string(data[mark:p])

	goto st111
	st111:
		p++
	st_case_111:
//line msg_parse.go:4731
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st111
		case 32:
			goto st111
		case 91:
			goto st135
		case 525:
			goto st138
		}
		switch {
		case _widec < 48:
			if 45 <= _widec && _widec <= 46 {
				goto tr169
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 97 <= _widec && _widec <= 122 {
					goto tr169
				}
			case _widec >= 65:
				goto tr169
			}
		default:
			goto tr169
		}
		goto st0
tr169:
//line sip.rl:62

	mark = p

	goto st112
	st112:
		p++
	st_case_112:
//line msg_parse.go:4776
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr172
		case 32:
			goto tr172
		case 44:
			goto tr173
		case 58:
			goto tr175
		case 59:
			goto tr176
		case 269:
			goto tr177
		case 525:
			goto tr178
		}
		switch {
		case _widec < 48:
			if 45 <= _widec && _widec <= 46 {
				goto st112
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 97 <= _widec && _widec <= 122 {
					goto st112
				}
			case _widec >= 65:
				goto st112
			}
		default:
			goto st112
		}
		goto st0
tr172:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st113
	st113:
		p++
	st_case_113:
//line msg_parse.go:4827
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st113
		case 32:
			goto st113
		case 44:
			goto st114
		case 58:
			goto st118
		case 59:
			goto st121
		case 525:
			goto st132
		}
		goto st0
tr173:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st114
	st114:
		p++
	st_case_114:
//line msg_parse.go:4859
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st114
		case 32:
			goto st114
		case 269:
			goto tr184
		case 525:
			goto st115
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr184
			}
		default:
			goto tr184
		}
		goto st0
tr184:
//line sip.rl:123

	*viap = via
	viap = &via.Next
	via = nil

//line sip.rl:119

	via = new(Via)

//line sip.rl:54

	p--


//line sip.rl:250
 {goto st103 } 
	goto st768
tr193:
//line sip.rl:54

	p--


//line sip.rl:70

	amt = 0

//line sip.rl:251
 {goto st68 } 
	goto st768
tr199:
//line sip.rl:123

	*viap = via
	viap = &via.Next
	via = nil

//line sip.rl:248
 {goto st280 } 
	goto st768
	st768:
		p++
	st_case_768:
//line msg_parse.go:4931
		goto st0
	st115:
		p++
	st_case_115:
		if data[p] == 10 {
			goto st116
		}
		goto st0
	st116:
		p++
	st_case_116:
		switch data[p] {
		case 9:
			goto st117
		case 32:
			goto st117
		}
		goto st0
	st117:
		p++
	st_case_117:
		switch data[p] {
		case 9:
			goto st117
		case 32:
			goto st117
		}
		goto tr184
tr175:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st118
	st118:
		p++
	st_case_118:
//line msg_parse.go:4969
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st118
		case 32:
			goto st118
		case 525:
			goto st129
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr188
		}
		goto st0
tr188:
//line sip.rl:145

	via.Port = via.Port * 10 + (uint16(data[p]) - 0x30)

	goto st119
	st119:
		p++
	st_case_119:
//line msg_parse.go:4998
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st120
		case 32:
			goto st120
		case 44:
			goto st114
		case 59:
			goto st121
		case 269:
			goto st128
		case 525:
			goto st125
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr188
		}
		goto st0
	st120:
		p++
	st_case_120:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st120
		case 32:
			goto st120
		case 44:
			goto st114
		case 59:
			goto st121
		case 525:
			goto st125
		}
		goto st0
tr176:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st121
	st121:
		p++
	st_case_121:
//line msg_parse.go:5056
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st121
		case 32:
			goto st121
		case 269:
			goto tr193
		case 525:
			goto st122
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr193
			}
		default:
			goto tr193
		}
		goto st0
	st122:
		p++
	st_case_122:
		if data[p] == 10 {
			goto st123
		}
		goto st0
	st123:
		p++
	st_case_123:
		switch data[p] {
		case 9:
			goto st124
		case 32:
			goto st124
		}
		goto st0
	st124:
		p++
	st_case_124:
		switch data[p] {
		case 9:
			goto st124
		case 32:
			goto st124
		}
		goto tr193
	st125:
		p++
	st_case_125:
		if data[p] == 10 {
			goto st126
		}
		goto st0
	st126:
		p++
	st_case_126:
		switch data[p] {
		case 9:
			goto st127
		case 32:
			goto st127
		}
		goto st0
	st127:
		p++
	st_case_127:
		switch data[p] {
		case 9:
			goto st127
		case 32:
			goto st127
		case 44:
			goto st114
		case 59:
			goto st121
		}
		goto st0
tr177:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st128
	st128:
		p++
	st_case_128:
//line msg_parse.go:5150
		if data[p] == 10 {
			goto tr199
		}
		goto st0
	st129:
		p++
	st_case_129:
		if data[p] == 10 {
			goto st130
		}
		goto st0
	st130:
		p++
	st_case_130:
		switch data[p] {
		case 9:
			goto st131
		case 32:
			goto st131
		}
		goto st0
	st131:
		p++
	st_case_131:
		switch data[p] {
		case 9:
			goto st131
		case 32:
			goto st131
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr188
		}
		goto st0
tr178:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st132
	st132:
		p++
	st_case_132:
//line msg_parse.go:5194
		if data[p] == 10 {
			goto st133
		}
		goto st0
	st133:
		p++
	st_case_133:
		switch data[p] {
		case 9:
			goto st134
		case 32:
			goto st134
		}
		goto st0
	st134:
		p++
	st_case_134:
		switch data[p] {
		case 9:
			goto st134
		case 32:
			goto st134
		case 44:
			goto st114
		case 58:
			goto st118
		case 59:
			goto st121
		}
		goto st0
	st135:
		p++
	st_case_135:
		if data[p] == 46 {
			goto tr204
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto tr204
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr204
			}
		default:
			goto tr204
		}
		goto st0
tr204:
//line sip.rl:62

	mark = p

	goto st136
	st136:
		p++
	st_case_136:
//line msg_parse.go:5253
		switch data[p] {
		case 46:
			goto st136
		case 93:
			goto tr206
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st136
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st136
			}
		default:
			goto st136
		}
		goto st0
tr206:
//line sip.rl:141

	via.Host = string(data[mark:p])

	goto st137
	st137:
		p++
	st_case_137:
//line msg_parse.go:5282
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st113
		case 32:
			goto st113
		case 44:
			goto st114
		case 58:
			goto st118
		case 59:
			goto st121
		case 269:
			goto st128
		case 525:
			goto st132
		}
		goto st0
tr167:
//line sip.rl:137

	via.Transport = string(data[mark:p])

	goto st138
	st138:
		p++
	st_case_138:
//line msg_parse.go:5316
		if data[p] == 10 {
			goto st139
		}
		goto st0
	st139:
		p++
	st_case_139:
		switch data[p] {
		case 9:
			goto st140
		case 32:
			goto st140
		}
		goto st0
	st140:
		p++
	st_case_140:
		switch data[p] {
		case 9:
			goto st140
		case 32:
			goto st140
		case 91:
			goto st135
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr169
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr169
				}
			case data[p] >= 65:
				goto tr169
			}
		default:
			goto tr169
		}
		goto st0
	st141:
		p++
	st_case_141:
		if data[p] == 10 {
			goto st142
		}
		goto st0
	st142:
		p++
	st_case_142:
		switch data[p] {
		case 9:
			goto st143
		case 32:
			goto st143
		}
		goto st0
	st143:
		p++
	st_case_143:
		switch data[p] {
		case 9:
			goto st143
		case 32:
			goto st143
		case 33:
			goto tr163
		case 37:
			goto tr163
		case 39:
			goto tr163
		case 126:
			goto tr163
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr163
				}
			case data[p] >= 42:
				goto tr163
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr163
				}
			case data[p] >= 65:
				goto tr163
			}
		default:
			goto tr163
		}
		goto st0
tr159:
//line sip.rl:133

	via.Version = string(data[mark:p])

	goto st144
	st144:
		p++
	st_case_144:
//line msg_parse.go:5426
		if data[p] == 10 {
			goto st145
		}
		goto st0
	st145:
		p++
	st_case_145:
		switch data[p] {
		case 9:
			goto st146
		case 32:
			goto st146
		}
		goto st0
	st146:
		p++
	st_case_146:
		switch data[p] {
		case 9:
			goto st146
		case 32:
			goto st146
		case 47:
			goto st109
		}
		goto st0
	st147:
		p++
	st_case_147:
		if data[p] == 10 {
			goto st148
		}
		goto st0
	st148:
		p++
	st_case_148:
		switch data[p] {
		case 9:
			goto st149
		case 32:
			goto st149
		}
		goto st0
	st149:
		p++
	st_case_149:
		switch data[p] {
		case 9:
			goto st149
		case 32:
			goto st149
		case 33:
			goto tr154
		case 37:
			goto tr154
		case 39:
			goto tr154
		case 126:
			goto tr154
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr154
				}
			case data[p] >= 42:
				goto tr154
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr154
				}
			case data[p] >= 65:
				goto tr154
			}
		default:
			goto tr154
		}
		goto st0
tr150:
//line sip.rl:129

	via.Protocol = string(data[mark:p])

	goto st150
	st150:
		p++
	st_case_150:
//line msg_parse.go:5519
		if data[p] == 10 {
			goto st151
		}
		goto st0
	st151:
		p++
	st_case_151:
		switch data[p] {
		case 9:
			goto st152
		case 32:
			goto st152
		}
		goto st0
	st152:
		p++
	st_case_152:
		switch data[p] {
		case 9:
			goto st152
		case 32:
			goto st152
		case 47:
			goto st106
		}
		goto st0
	st153:
		p++
	st_case_153:
		switch data[p] {
		case 33:
			goto tr217
		case 37:
			goto tr217
		case 39:
			goto tr217
		case 126:
			goto tr217
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr217
				}
			case data[p] >= 42:
				goto tr217
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr217
				}
			case data[p] >= 65:
				goto tr217
			}
		default:
			goto tr217
		}
		goto st0
tr217:
//line sip.rl:70

	amt = 0

//line sip.rl:62

	mark = p

	goto st154
	st154:
		p++
	st_case_154:
//line msg_parse.go:5595
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr218
		case 32:
			goto tr218
		case 33:
			goto st154
		case 37:
			goto st154
		case 39:
			goto st154
		case 44:
			goto tr220
		case 59:
			goto tr221
		case 61:
			goto tr222
		case 126:
			goto st154
		case 269:
			goto tr223
		case 525:
			goto tr224
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto st154
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st154
				}
			case _widec >= 65:
				goto st154
			}
		default:
			goto st154
		}
		goto st0
tr218:
//line sip.rl:161

	name = string(data[mark:p])

	goto st155
	st155:
		p++
	st_case_155:
//line msg_parse.go:5654
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st155
		case 32:
			goto st155
		case 44:
			goto st156
		case 59:
			goto st160
		case 61:
			goto st164
		case 525:
			goto st185
		}
		goto st0
tr220:
//line sip.rl:161

	name = string(data[mark:p])

	goto st156
	st156:
		p++
	st_case_156:
//line msg_parse.go:5686
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st156
		case 32:
			goto st156
		case 269:
			goto tr230
		case 525:
			goto st157
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr230
			}
		default:
			goto tr230
		}
		goto st0
tr230:
//line sip.rl:198

	if addr.Params == nil {
		addr.Params = Params{}
	}
	addr.Params[name] = string(buf[0:amt])

//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st769
tr234:
//line sip.rl:198

	if addr.Params == nil {
		addr.Params = Params{}
	}
	addr.Params[name] = string(buf[0:amt])

//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st769
tr246:
//line sip.rl:198

	if addr.Params == nil {
		addr.Params = Params{}
	}
	addr.Params[name] = string(buf[0:amt])

//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:248
 {goto st280 } 
	goto st769
	st769:
		p++
	st_case_769:
//line msg_parse.go:5771
		goto st0
	st157:
		p++
	st_case_157:
		if data[p] == 10 {
			goto st158
		}
		goto st0
	st158:
		p++
	st_case_158:
		switch data[p] {
		case 9:
			goto st159
		case 32:
			goto st159
		}
		goto st0
	st159:
		p++
	st_case_159:
		switch data[p] {
		case 9:
			goto st159
		case 32:
			goto st159
		}
		goto tr230
tr221:
//line sip.rl:161

	name = string(data[mark:p])

	goto st160
	st160:
		p++
	st_case_160:
//line msg_parse.go:5809
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st160
		case 32:
			goto st160
		case 269:
			goto tr234
		case 525:
			goto st161
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr234
			}
		default:
			goto tr234
		}
		goto st0
	st161:
		p++
	st_case_161:
		if data[p] == 10 {
			goto st162
		}
		goto st0
	st162:
		p++
	st_case_162:
		switch data[p] {
		case 9:
			goto st163
		case 32:
			goto st163
		}
		goto st0
	st163:
		p++
	st_case_163:
		switch data[p] {
		case 9:
			goto st163
		case 32:
			goto st163
		}
		goto tr234
tr222:
//line sip.rl:161

	name = string(data[mark:p])

	goto st164
	st164:
		p++
	st_case_164:
//line msg_parse.go:5872
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st164
		case 32:
			goto st164
		case 33:
			goto tr238
		case 34:
			goto st171
		case 37:
			goto tr238
		case 39:
			goto tr238
		case 93:
			goto tr238
		case 126:
			goto tr238
		case 525:
			goto st182
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr238
				}
			case _widec >= 42:
				goto tr238
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr238
				}
			case _widec >= 65:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st0
tr238:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st165
	st165:
		p++
	st_case_165:
//line msg_parse.go:5933
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st166
		case 32:
			goto st166
		case 33:
			goto tr238
		case 37:
			goto tr238
		case 39:
			goto tr238
		case 44:
			goto st156
		case 59:
			goto st160
		case 93:
			goto tr238
		case 126:
			goto tr238
		case 269:
			goto st170
		case 525:
			goto st167
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr238
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr238
				}
			case _widec >= 65:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st0
	st166:
		p++
	st_case_166:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st166
		case 32:
			goto st166
		case 44:
			goto st156
		case 59:
			goto st160
		case 525:
			goto st167
		}
		goto st0
	st167:
		p++
	st_case_167:
		if data[p] == 10 {
			goto st168
		}
		goto st0
	st168:
		p++
	st_case_168:
		switch data[p] {
		case 9:
			goto st169
		case 32:
			goto st169
		}
		goto st0
	st169:
		p++
	st_case_169:
		switch data[p] {
		case 9:
			goto st169
		case 32:
			goto st169
		case 44:
			goto st156
		case 59:
			goto st160
		}
		goto st0
tr223:
//line sip.rl:161

	name = string(data[mark:p])

	goto st170
	st170:
		p++
	st_case_170:
//line msg_parse.go:6046
		if data[p] == 10 {
			goto tr246
		}
		goto st0
	st171:
		p++
	st_case_171:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr247
		case 34:
			goto tr248
		case 92:
			goto tr249
		case 525:
			goto tr255
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr250
				}
			case _widec >= 32:
				goto tr247
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr252
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr254
				}
			default:
				goto tr253
			}
		default:
			goto tr251
		}
		goto st0
tr247:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st172
tr256:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st172
	st172:
		p++
	st_case_172:
//line msg_parse.go:6119
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr256
		case 34:
			goto st173
		case 92:
			goto st174
		case 525:
			goto tr264
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr259
				}
			case _widec >= 32:
				goto tr256
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr261
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr263
				}
			default:
				goto tr262
			}
		default:
			goto tr260
		}
		goto st0
tr248:
//line sip.rl:70

	amt = 0

	goto st173
	st173:
		p++
	st_case_173:
//line msg_parse.go:6173
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st166
		case 32:
			goto st166
		case 44:
			goto st156
		case 59:
			goto st160
		case 269:
			goto st170
		case 525:
			goto st167
		}
		goto st0
tr249:
//line sip.rl:70

	amt = 0

	goto st174
	st174:
		p++
	st_case_174:
//line msg_parse.go:6205
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr256
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr256
			}
		default:
			goto tr256
		}
		goto st0
tr250:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st175
tr259:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st175
	st175:
		p++
	st_case_175:
//line msg_parse.go:6240
		if 128 <= data[p] && data[p] <= 191 {
			goto tr256
		}
		goto st0
tr251:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st176
tr260:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st176
	st176:
		p++
	st_case_176:
//line msg_parse.go:6266
		if 128 <= data[p] && data[p] <= 191 {
			goto tr259
		}
		goto st0
tr252:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st177
tr261:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st177
	st177:
		p++
	st_case_177:
//line msg_parse.go:6292
		if 128 <= data[p] && data[p] <= 191 {
			goto tr260
		}
		goto st0
tr253:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st178
tr262:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st178
	st178:
		p++
	st_case_178:
//line msg_parse.go:6318
		if 128 <= data[p] && data[p] <= 191 {
			goto tr261
		}
		goto st0
tr254:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st179
tr263:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st179
	st179:
		p++
	st_case_179:
//line msg_parse.go:6344
		if 128 <= data[p] && data[p] <= 191 {
			goto tr262
		}
		goto st0
tr255:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st180
tr264:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st180
	st180:
		p++
	st_case_180:
//line msg_parse.go:6370
		if data[p] == 10 {
			goto tr265
		}
		goto st0
tr265:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st181
	st181:
		p++
	st_case_181:
//line msg_parse.go:6385
		switch data[p] {
		case 9:
			goto tr256
		case 32:
			goto tr256
		}
		goto st0
	st182:
		p++
	st_case_182:
		if data[p] == 10 {
			goto st183
		}
		goto st0
	st183:
		p++
	st_case_183:
		switch data[p] {
		case 9:
			goto st184
		case 32:
			goto st184
		}
		goto st0
	st184:
		p++
	st_case_184:
		switch data[p] {
		case 9:
			goto st184
		case 32:
			goto st184
		case 33:
			goto tr238
		case 34:
			goto st171
		case 37:
			goto tr238
		case 39:
			goto tr238
		case 93:
			goto tr238
		case 126:
			goto tr238
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr238
				}
			case data[p] >= 42:
				goto tr238
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr238
				}
			case data[p] >= 65:
				goto tr238
			}
		default:
			goto tr238
		}
		goto st0
tr224:
//line sip.rl:161

	name = string(data[mark:p])

	goto st185
	st185:
		p++
	st_case_185:
//line msg_parse.go:6463
		if data[p] == 10 {
			goto st186
		}
		goto st0
	st186:
		p++
	st_case_186:
		switch data[p] {
		case 9:
			goto st187
		case 32:
			goto st187
		}
		goto st0
	st187:
		p++
	st_case_187:
		switch data[p] {
		case 9:
			goto st187
		case 32:
			goto st187
		case 44:
			goto st156
		case 59:
			goto st160
		case 61:
			goto st164
		}
		goto st0
	st188:
		p++
	st_case_188:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st189
		case 32:
			goto st189
		case 33:
			goto tr271
		case 34:
			goto tr272
		case 37:
			goto tr271
		case 39:
			goto tr271
		case 60:
			goto st190
		case 126:
			goto tr271
		case 525:
			goto st210
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr271
				}
			case _widec >= 42:
				goto tr271
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr271
				}
			case _widec >= 65:
				goto tr271
			}
		default:
			goto tr271
		}
		goto st0
tr329:
//line sip.rl:181

	addr.Display = string(buf[0:amt])

	goto st189
	st189:
		p++
	st_case_189:
//line msg_parse.go:6556
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st189
		case 32:
			goto st189
		case 60:
			goto st190
		case 525:
			goto st210
		}
		goto st0
tr305:
//line sip.rl:185
{
	end := p
	for end > mark && whitespacec(data[end - 1]) {
		end--
	}
	addr.Display = string(data[mark:end])
}
	goto st190
tr330:
//line sip.rl:181

	addr.Display = string(buf[0:amt])

	goto st190
	st190:
		p++
	st_case_190:
//line msg_parse.go:6594
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr275
			}
		case data[p] >= 65:
			goto tr275
		}
		goto st0
tr275:
//line sip.rl:62

	mark = p

	goto st191
	st191:
		p++
	st_case_191:
//line msg_parse.go:6613
		switch data[p] {
		case 43:
			goto st191
		case 58:
			goto st192
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st191
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st191
				}
			case data[p] >= 65:
				goto st191
			}
		default:
			goto st191
		}
		goto st0
	st192:
		p++
	st_case_192:
		switch data[p] {
		case 33:
			goto st193
		case 61:
			goto st193
		case 93:
			goto st193
		case 95:
			goto st193
		case 126:
			goto st193
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st193
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st193
			}
		default:
			goto st193
		}
		goto st0
	st193:
		p++
	st_case_193:
		switch data[p] {
		case 33:
			goto st193
		case 62:
			goto tr279
		case 93:
			goto st193
		case 95:
			goto st193
		case 126:
			goto st193
		}
		switch {
		case data[p] < 61:
			if 36 <= data[p] && data[p] <= 59 {
				goto st193
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st193
			}
		default:
			goto st193
		}
		goto st0
tr279:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st194
	st194:
		p++
	st_case_194:
//line msg_parse.go:6704
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st194
		case 32:
			goto st194
		case 44:
			goto st195
		case 59:
			goto st199
		case 269:
			goto st203
		case 525:
			goto st204
		}
		goto st0
	st195:
		p++
	st_case_195:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st195
		case 32:
			goto st195
		case 269:
			goto tr285
		case 525:
			goto st196
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr285
			}
		default:
			goto tr285
		}
		goto st0
tr285:
//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st770
tr289:
//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st770
tr293:
//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:248
 {goto st280 } 
	goto st770
	st770:
		p++
	st_case_770:
//line msg_parse.go:6793
		goto st0
	st196:
		p++
	st_case_196:
		if data[p] == 10 {
			goto st197
		}
		goto st0
	st197:
		p++
	st_case_197:
		switch data[p] {
		case 9:
			goto st198
		case 32:
			goto st198
		}
		goto st0
	st198:
		p++
	st_case_198:
		switch data[p] {
		case 9:
			goto st198
		case 32:
			goto st198
		}
		goto tr285
	st199:
		p++
	st_case_199:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st199
		case 32:
			goto st199
		case 269:
			goto tr289
		case 525:
			goto st200
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr289
			}
		default:
			goto tr289
		}
		goto st0
	st200:
		p++
	st_case_200:
		if data[p] == 10 {
			goto st201
		}
		goto st0
	st201:
		p++
	st_case_201:
		switch data[p] {
		case 9:
			goto st202
		case 32:
			goto st202
		}
		goto st0
	st202:
		p++
	st_case_202:
		switch data[p] {
		case 9:
			goto st202
		case 32:
			goto st202
		}
		goto tr289
	st203:
		p++
	st_case_203:
		if data[p] == 10 {
			goto tr293
		}
		goto st0
	st204:
		p++
	st_case_204:
		if data[p] == 10 {
			goto st205
		}
		goto st0
	st205:
		p++
	st_case_205:
		switch data[p] {
		case 9:
			goto st206
		case 32:
			goto st206
		}
		goto st0
	st206:
		p++
	st_case_206:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st206
		case 32:
			goto st206
		case 44:
			goto st195
		case 59:
			goto st199
		case 269:
			goto st203
		case 525:
			goto st207
		}
		goto st0
	st207:
		p++
	st_case_207:
		if data[p] == 10 {
			goto st208
		}
		goto st0
	st208:
		p++
	st_case_208:
		switch data[p] {
		case 9:
			goto st209
		case 32:
			goto st209
		}
		goto st0
	st209:
		p++
	st_case_209:
		switch data[p] {
		case 9:
			goto st209
		case 32:
			goto st209
		case 44:
			goto st195
		case 59:
			goto st199
		}
		goto st0
tr310:
//line sip.rl:185
{
	end := p
	for end > mark && whitespacec(data[end - 1]) {
		end--
	}
	addr.Display = string(data[mark:end])
}
	goto st210
tr331:
//line sip.rl:181

	addr.Display = string(buf[0:amt])

	goto st210
	st210:
		p++
	st_case_210:
//line msg_parse.go:6977
		if data[p] == 10 {
			goto st211
		}
		goto st0
	st211:
		p++
	st_case_211:
		switch data[p] {
		case 9:
			goto st212
		case 32:
			goto st212
		}
		goto st0
	st212:
		p++
	st_case_212:
		switch data[p] {
		case 9:
			goto st212
		case 32:
			goto st212
		case 60:
			goto st190
		}
		goto st0
tr271:
//line sip.rl:62

	mark = p

	goto st213
	st213:
		p++
	st_case_213:
//line msg_parse.go:7013
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st214
		case 32:
			goto st214
		case 33:
			goto st213
		case 37:
			goto st213
		case 39:
			goto st213
		case 126:
			goto st213
		case 525:
			goto st215
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st213
				}
			case _widec >= 42:
				goto st213
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st213
				}
			case _widec >= 65:
				goto st213
			}
		default:
			goto st213
		}
		goto st0
tr304:
//line sip.rl:185
{
	end := p
	for end > mark && whitespacec(data[end - 1]) {
		end--
	}
	addr.Display = string(data[mark:end])
}
	goto st214
	st214:
		p++
	st_case_214:
//line msg_parse.go:7073
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr304
		case 32:
			goto tr304
		case 33:
			goto st213
		case 37:
			goto st213
		case 39:
			goto st213
		case 60:
			goto tr305
		case 126:
			goto st213
		case 525:
			goto tr306
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st213
				}
			case _widec >= 42:
				goto st213
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st213
				}
			case _widec >= 65:
				goto st213
			}
		default:
			goto st213
		}
		goto st0
tr306:
//line sip.rl:185
{
	end := p
	for end > mark && whitespacec(data[end - 1]) {
		end--
	}
	addr.Display = string(data[mark:end])
}
	goto st215
	st215:
		p++
	st_case_215:
//line msg_parse.go:7135
		if data[p] == 10 {
			goto st216
		}
		goto st0
	st216:
		p++
	st_case_216:
		switch data[p] {
		case 9:
			goto st217
		case 32:
			goto st217
		}
		goto st0
tr309:
//line sip.rl:185
{
	end := p
	for end > mark && whitespacec(data[end - 1]) {
		end--
	}
	addr.Display = string(data[mark:end])
}
	goto st217
	st217:
		p++
	st_case_217:
//line msg_parse.go:7163
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr309
		case 32:
			goto tr309
		case 33:
			goto st213
		case 37:
			goto st213
		case 39:
			goto st213
		case 60:
			goto tr305
		case 126:
			goto st213
		case 525:
			goto tr310
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st213
				}
			case _widec >= 42:
				goto st213
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st213
				}
			case _widec >= 65:
				goto st213
			}
		default:
			goto st213
		}
		goto st0
tr272:
//line sip.rl:70

	amt = 0

	goto st218
	st218:
		p++
	st_case_218:
//line msg_parse.go:7221
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr311
		case 34:
			goto tr312
		case 92:
			goto tr313
		case 525:
			goto tr319
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr314
				}
			case _widec >= 32:
				goto tr311
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr316
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr318
				}
			default:
				goto tr317
			}
		default:
			goto tr315
		}
		goto st0
tr311:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st219
tr320:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st219
	st219:
		p++
	st_case_219:
//line msg_parse.go:7287
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr320
		case 34:
			goto st220
		case 92:
			goto st221
		case 525:
			goto tr328
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr323
				}
			case _widec >= 32:
				goto tr320
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr325
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr327
				}
			default:
				goto tr326
			}
		default:
			goto tr324
		}
		goto st0
tr312:
//line sip.rl:70

	amt = 0

	goto st220
	st220:
		p++
	st_case_220:
//line msg_parse.go:7341
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr329
		case 32:
			goto tr329
		case 60:
			goto tr330
		case 525:
			goto tr331
		}
		goto st0
tr313:
//line sip.rl:70

	amt = 0

	goto st221
	st221:
		p++
	st_case_221:
//line msg_parse.go:7369
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr320
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr320
			}
		default:
			goto tr320
		}
		goto st0
tr314:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st222
tr323:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st222
	st222:
		p++
	st_case_222:
//line msg_parse.go:7404
		if 128 <= data[p] && data[p] <= 191 {
			goto tr320
		}
		goto st0
tr315:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st223
tr324:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st223
	st223:
		p++
	st_case_223:
//line msg_parse.go:7430
		if 128 <= data[p] && data[p] <= 191 {
			goto tr323
		}
		goto st0
tr316:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st224
tr325:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st224
	st224:
		p++
	st_case_224:
//line msg_parse.go:7456
		if 128 <= data[p] && data[p] <= 191 {
			goto tr324
		}
		goto st0
tr317:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st225
tr326:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st225
	st225:
		p++
	st_case_225:
//line msg_parse.go:7482
		if 128 <= data[p] && data[p] <= 191 {
			goto tr325
		}
		goto st0
tr318:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st226
tr327:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st226
	st226:
		p++
	st_case_226:
//line msg_parse.go:7508
		if 128 <= data[p] && data[p] <= 191 {
			goto tr326
		}
		goto st0
tr319:
//line sip.rl:70

	amt = 0

//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st227
tr328:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st227
	st227:
		p++
	st_case_227:
//line msg_parse.go:7534
		if data[p] == 10 {
			goto tr332
		}
		goto st0
tr332:
//line sip.rl:74

	buf[amt] = data[p]
	amt++

	goto st228
	st228:
		p++
	st_case_228:
//line msg_parse.go:7549
		switch data[p] {
		case 9:
			goto tr320
		case 32:
			goto tr320
		}
		goto st0
	st229:
		p++
	st_case_229:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st230
			}
		case data[p] >= 65:
			goto st230
		}
		goto st0
	st230:
		p++
	st_case_230:
		switch data[p] {
		case 43:
			goto st230
		case 58:
			goto st231
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st230
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st230
				}
			case data[p] >= 65:
				goto st230
			}
		default:
			goto st230
		}
		goto st0
	st231:
		p++
	st_case_231:
		switch data[p] {
		case 33:
			goto st232
		case 61:
			goto st232
		case 93:
			goto st232
		case 95:
			goto st232
		case 126:
			goto st232
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st232
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st232
			}
		default:
			goto st232
		}
		goto st0
	st232:
		p++
	st_case_232:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr336
		case 32:
			goto tr336
		case 33:
			goto st232
		case 44:
			goto tr337
		case 59:
			goto tr338
		case 61:
			goto st232
		case 93:
			goto st232
		case 95:
			goto st232
		case 126:
			goto st232
		case 269:
			goto tr339
		case 525:
			goto tr340
		}
		switch {
		case _widec < 63:
			if 36 <= _widec && _widec <= 58 {
				goto st232
			}
		case _widec > 91:
			if 97 <= _widec && _widec <= 122 {
				goto st232
			}
		default:
			goto st232
		}
		goto st0
tr336:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st233
	st233:
		p++
	st_case_233:
//line msg_parse.go:7681
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st233
		case 32:
			goto st233
		case 44:
			goto st234
		case 59:
			goto st238
		case 525:
			goto st242
		}
		goto st0
	st234:
		p++
	st_case_234:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st234
		case 32:
			goto st234
		case 269:
			goto tr345
		case 525:
			goto st235
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr345
			}
		default:
			goto tr345
		}
		goto st0
tr345:
//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st771
tr349:
//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st771
	st771:
		p++
	st_case_771:
//line msg_parse.go:7758
		goto st0
	st235:
		p++
	st_case_235:
		if data[p] == 10 {
			goto st236
		}
		goto st0
	st236:
		p++
	st_case_236:
		switch data[p] {
		case 9:
			goto st237
		case 32:
			goto st237
		}
		goto st0
	st237:
		p++
	st_case_237:
		switch data[p] {
		case 9:
			goto st237
		case 32:
			goto st237
		}
		goto tr345
	st238:
		p++
	st_case_238:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st238
		case 32:
			goto st238
		case 269:
			goto tr349
		case 525:
			goto st239
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr349
			}
		default:
			goto tr349
		}
		goto st0
	st239:
		p++
	st_case_239:
		if data[p] == 10 {
			goto st240
		}
		goto st0
	st240:
		p++
	st_case_240:
		switch data[p] {
		case 9:
			goto st241
		case 32:
			goto st241
		}
		goto st0
	st241:
		p++
	st_case_241:
		switch data[p] {
		case 9:
			goto st241
		case 32:
			goto st241
		}
		goto tr349
tr340:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st242
	st242:
		p++
	st_case_242:
//line msg_parse.go:7853
		if data[p] == 10 {
			goto st243
		}
		goto st0
	st243:
		p++
	st_case_243:
		switch data[p] {
		case 9:
			goto st244
		case 32:
			goto st244
		}
		goto st0
	st244:
		p++
	st_case_244:
		switch data[p] {
		case 9:
			goto st244
		case 32:
			goto st244
		case 44:
			goto st234
		case 59:
			goto st238
		}
		goto st0
tr337:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st245
	st245:
		p++
	st_case_245:
//line msg_parse.go:7892
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr355
		case 32:
			goto tr355
		case 33:
			goto tr356
		case 44:
			goto tr337
		case 59:
			goto tr357
		case 60:
			goto tr345
		case 62:
			goto tr345
		case 92:
			goto tr345
		case 94:
			goto tr345
		case 96:
			goto tr345
		case 126:
			goto tr356
		case 269:
			goto tr358
		case 525:
			goto tr359
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr345
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr345
				}
			case _widec >= 36:
				goto tr356
			}
		default:
			goto tr345
		}
		goto st0
tr355:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st246
	st246:
		p++
	st_case_246:
//line msg_parse.go:7956
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st246
		case 32:
			goto st246
		case 44:
			goto st234
		case 59:
			goto tr361
		case 269:
			goto tr345
		case 525:
			goto st247
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr345
			}
		default:
			goto tr345
		}
		goto st0
tr361:
//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st772
	st772:
		p++
	st_case_772:
//line msg_parse.go:8005
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st238
		case 32:
			goto st238
		case 269:
			goto tr349
		case 525:
			goto st239
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr349
			}
		default:
			goto tr349
		}
		goto st0
tr359:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st247
	st247:
		p++
	st_case_247:
//line msg_parse.go:8042
		if data[p] == 10 {
			goto st248
		}
		goto st0
	st248:
		p++
	st_case_248:
		switch data[p] {
		case 9:
			goto st249
		case 32:
			goto st249
		}
		goto st0
	st249:
		p++
	st_case_249:
		switch data[p] {
		case 9:
			goto st249
		case 32:
			goto st249
		case 44:
			goto st234
		case 59:
			goto tr361
		}
		goto tr345
tr356:
//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st773
tr366:
//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st773
	st773:
		p++
	st_case_773:
//line msg_parse.go:8098
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr336
		case 32:
			goto tr336
		case 33:
			goto st232
		case 44:
			goto tr337
		case 59:
			goto tr338
		case 61:
			goto st232
		case 93:
			goto st232
		case 95:
			goto st232
		case 126:
			goto st232
		case 269:
			goto tr339
		case 525:
			goto tr340
		}
		switch {
		case _widec < 63:
			if 36 <= _widec && _widec <= 58 {
				goto st232
			}
		case _widec > 91:
			if 97 <= _widec && _widec <= 122 {
				goto st232
			}
		default:
			goto st232
		}
		goto st0
tr338:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st250
	st250:
		p++
	st_case_250:
//line msg_parse.go:8153
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr365
		case 32:
			goto tr365
		case 33:
			goto tr366
		case 44:
			goto tr367
		case 59:
			goto tr338
		case 60:
			goto tr349
		case 62:
			goto tr349
		case 92:
			goto tr349
		case 94:
			goto tr349
		case 96:
			goto tr349
		case 126:
			goto tr366
		case 269:
			goto tr368
		case 525:
			goto tr369
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr349
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr349
				}
			case _widec >= 36:
				goto tr366
			}
		default:
			goto tr349
		}
		goto st0
tr365:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st251
	st251:
		p++
	st_case_251:
//line msg_parse.go:8217
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st251
		case 32:
			goto st251
		case 44:
			goto tr371
		case 59:
			goto st238
		case 269:
			goto tr349
		case 525:
			goto st252
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr349
			}
		default:
			goto tr349
		}
		goto st0
tr371:
//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st774
	st774:
		p++
	st_case_774:
//line msg_parse.go:8260
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st234
		case 32:
			goto st234
		case 269:
			goto tr345
		case 525:
			goto st235
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr345
			}
		default:
			goto tr345
		}
		goto st0
tr369:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st252
	st252:
		p++
	st_case_252:
//line msg_parse.go:8297
		if data[p] == 10 {
			goto st253
		}
		goto st0
	st253:
		p++
	st_case_253:
		switch data[p] {
		case 9:
			goto st254
		case 32:
			goto st254
		}
		goto st0
	st254:
		p++
	st_case_254:
		switch data[p] {
		case 9:
			goto st254
		case 32:
			goto st254
		case 44:
			goto tr371
		case 59:
			goto st238
		}
		goto tr349
tr367:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st775
	st775:
		p++
	st_case_775:
//line msg_parse.go:8343
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr355
		case 32:
			goto tr355
		case 33:
			goto tr356
		case 44:
			goto tr337
		case 59:
			goto tr357
		case 60:
			goto tr345
		case 62:
			goto tr345
		case 92:
			goto tr345
		case 94:
			goto tr345
		case 96:
			goto tr345
		case 126:
			goto tr356
		case 269:
			goto tr358
		case 525:
			goto tr359
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr345
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr345
				}
			case _widec >= 36:
				goto tr356
			}
		default:
			goto tr345
		}
		goto st0
tr357:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st776
	st776:
		p++
	st_case_776:
//line msg_parse.go:8420
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr365
		case 32:
			goto tr365
		case 33:
			goto tr366
		case 44:
			goto tr367
		case 59:
			goto tr338
		case 60:
			goto tr349
		case 62:
			goto tr349
		case 92:
			goto tr349
		case 94:
			goto tr349
		case 96:
			goto tr349
		case 126:
			goto tr366
		case 269:
			goto tr368
		case 525:
			goto tr369
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr349
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr349
				}
			case _widec >= 36:
				goto tr366
			}
		default:
			goto tr349
		}
		goto st0
tr358:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st777
tr368:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

//line sip.rl:54

	p--


//line sip.rl:245
 {goto st153 } 
	goto st777
	st777:
		p++
	st_case_777:
//line msg_parse.go:8511
		if data[p] == 10 {
			goto tr375
		}
		goto st0
tr375:
//line sip.rl:248
 {goto st280 } 
	goto st778
	st778:
		p++
	st_case_778:
//line msg_parse.go:8523
		goto st0
tr339:
//line sip.rl:193

	addr.Uri, err = ParseURIBytes(data[mark:p])
	if err != nil { return nil, err }

	goto st255
	st255:
		p++
	st_case_255:
//line msg_parse.go:8535
		if data[p] == 10 {
			goto tr375
		}
		goto st0
	st256:
		p++
	st_case_256:
		switch data[p] {
		case 33:
			goto tr376
		case 34:
			goto tr377
		case 37:
			goto tr376
		case 39:
			goto tr376
		case 60:
			goto tr377
		case 126:
			goto tr376
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr376
				}
			case data[p] >= 42:
				goto tr376
			}
		case data[p] > 57:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr378
				}
			case data[p] > 96:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr378
				}
			default:
				goto tr376
			}
		default:
			goto tr376
		}
		goto st0
tr376:
//line sip.rl:62

	mark = p

	goto st257
	st257:
		p++
	st_case_257:
//line msg_parse.go:8593
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st258
		case 32:
			goto st258
		case 33:
			goto st257
		case 37:
			goto st257
		case 39:
			goto st257
		case 126:
			goto st257
		case 525:
			goto st259
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st257
				}
			case _widec >= 42:
				goto st257
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st257
				}
			case _widec >= 65:
				goto st257
			}
		default:
			goto st257
		}
		goto st0
	st258:
		p++
	st_case_258:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st258
		case 32:
			goto st258
		case 33:
			goto st257
		case 37:
			goto st257
		case 39:
			goto st257
		case 60:
			goto tr382
		case 126:
			goto st257
		case 525:
			goto st259
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st257
				}
			case _widec >= 42:
				goto st257
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st257
				}
			case _widec >= 65:
				goto st257
			}
		default:
			goto st257
		}
		goto st0
tr377:
//line sip.rl:177

	addr = new(Addr)

//line sip.rl:54

	p--


//line sip.rl:244
 {goto st188 } 
	goto st779
tr382:
//line sip.rl:177

	addr = new(Addr)

//line sip.rl:66

	p = ( mark) - 1


//line sip.rl:244
 {goto st188 } 
	goto st779
tr386:
//line sip.rl:177

	addr = new(Addr)

//line sip.rl:66

	p = ( mark) - 1


//line sip.rl:246
 {goto st229 } 
	goto st779
	st779:
		p++
	st_case_779:
//line msg_parse.go:8733
		goto st0
	st259:
		p++
	st_case_259:
		if data[p] == 10 {
			goto st260
		}
		goto st0
	st260:
		p++
	st_case_260:
		switch data[p] {
		case 9:
			goto st261
		case 32:
			goto st261
		}
		goto st0
	st261:
		p++
	st_case_261:
		switch data[p] {
		case 9:
			goto st261
		case 32:
			goto st261
		case 33:
			goto st257
		case 37:
			goto st257
		case 39:
			goto st257
		case 60:
			goto tr382
		case 126:
			goto st257
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st257
				}
			case data[p] >= 42:
				goto st257
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st257
				}
			case data[p] >= 65:
				goto st257
			}
		default:
			goto st257
		}
		goto st0
tr378:
//line sip.rl:62

	mark = p

	goto st262
	st262:
		p++
	st_case_262:
//line msg_parse.go:8803
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st258
		case 32:
			goto st258
		case 33:
			goto st257
		case 37:
			goto st257
		case 39:
			goto st257
		case 42:
			goto st257
		case 43:
			goto st262
		case 58:
			goto tr386
		case 126:
			goto st257
		case 525:
			goto st259
		}
		switch {
		case _widec < 65:
			switch {
			case _widec > 46:
				if 48 <= _widec && _widec <= 57 {
					goto st262
				}
			case _widec >= 45:
				goto st262
			}
		case _widec > 90:
			switch {
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto st262
				}
			case _widec >= 95:
				goto st257
			}
		default:
			goto st262
		}
		goto st0
	st263:
		p++
	st_case_263:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr387
		case 269:
			goto tr393
		case 525:
			goto tr394
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr388
				}
			case _widec >= 32:
				goto tr387
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr390
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr392
				}
			default:
				goto tr391
			}
		default:
			goto tr389
		}
		goto st0
tr387:
//line sip.rl:62

	mark = p

	goto st264
	st264:
		p++
	st_case_264:
//line msg_parse.go:8910
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st264
		case 269:
			goto st270
		case 525:
			goto st271
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto st265
				}
			case _widec >= 32:
				goto st264
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto st267
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto st269
				}
			default:
				goto st268
			}
		default:
			goto st266
		}
		goto st0
tr388:
//line sip.rl:62

	mark = p

	goto st265
	st265:
		p++
	st_case_265:
//line msg_parse.go:8962
		if 128 <= data[p] && data[p] <= 191 {
			goto st264
		}
		goto st0
tr389:
//line sip.rl:62

	mark = p

	goto st266
	st266:
		p++
	st_case_266:
//line msg_parse.go:8976
		if 128 <= data[p] && data[p] <= 191 {
			goto st265
		}
		goto st0
tr390:
//line sip.rl:62

	mark = p

	goto st267
	st267:
		p++
	st_case_267:
//line msg_parse.go:8990
		if 128 <= data[p] && data[p] <= 191 {
			goto st266
		}
		goto st0
tr391:
//line sip.rl:62

	mark = p

	goto st268
	st268:
		p++
	st_case_268:
//line msg_parse.go:9004
		if 128 <= data[p] && data[p] <= 191 {
			goto st267
		}
		goto st0
tr392:
//line sip.rl:62

	mark = p

	goto st269
	st269:
		p++
	st_case_269:
//line msg_parse.go:9018
		if 128 <= data[p] && data[p] <= 191 {
			goto st268
		}
		goto st0
tr393:
//line sip.rl:62

	mark = p

	goto st270
	st270:
		p++
	st_case_270:
//line msg_parse.go:9032
		if data[p] == 10 {
			goto tr403
		}
		goto st0
tr403:
//line sip.rl:165
{
	b := data[mark:p - 1]
	if value != nil {
		*value = string(b)
	} else {
		if msg.Headers == nil {
			msg.Headers = Headers{}
		}
		msg.Headers[name] = string(b)
	}
}
//line sip.rl:248
 {goto st280 } 
	goto st780
	st780:
		p++
	st_case_780:
//line msg_parse.go:9056
		goto st0
tr394:
//line sip.rl:62

	mark = p

	goto st271
	st271:
		p++
	st_case_271:
//line msg_parse.go:9067
		if data[p] == 10 {
			goto st272
		}
		goto st0
	st272:
		p++
	st_case_272:
		switch data[p] {
		case 9:
			goto st264
		case 32:
			goto st264
		}
		goto st0
	st273:
		p++
	st_case_273:
		switch data[p] {
		case 33:
			goto st274
		case 37:
			goto st274
		case 39:
			goto st274
		case 126:
			goto st274
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st274
				}
			case data[p] >= 42:
				goto st274
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st0
	st274:
		p++
	st_case_274:
		switch data[p] {
		case 9:
			goto tr406
		case 32:
			goto tr406
		case 33:
			goto st274
		case 37:
			goto st274
		case 39:
			goto st274
		case 58:
			goto tr407
		case 126:
			goto st274
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st274
				}
			case data[p] >= 42:
				goto st274
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st274
				}
			case data[p] >= 65:
				goto st274
			}
		default:
			goto st274
		}
		goto st0
tr406:
//line sip.rl:161

	name = string(data[mark:p])

	goto st275
	st275:
		p++
	st_case_275:
//line msg_parse.go:9169
		switch data[p] {
		case 9:
			goto st275
		case 32:
			goto st275
		case 58:
			goto st276
		}
		goto st0
tr407:
//line sip.rl:161

	name = string(data[mark:p])

	goto st276
	st276:
		p++
	st_case_276:
//line msg_parse.go:9188
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st276
		case 32:
			goto st276
		case 269:
			goto tr410
		case 525:
			goto st277
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr410
			}
		default:
			goto tr410
		}
		goto st0
tr410:
//line sip.rl:523
value=nil
//line sip.rl:54

	p--


//line sip.rl:249
 {goto st263 } 
	goto st781
	st781:
		p++
	st_case_781:
//line msg_parse.go:9229
		goto st0
	st277:
		p++
	st_case_277:
		if data[p] == 10 {
			goto st278
		}
		goto st0
	st278:
		p++
	st_case_278:
		switch data[p] {
		case 9:
			goto st279
		case 32:
			goto st279
		}
		goto st0
	st279:
		p++
	st_case_279:
		switch data[p] {
		case 9:
			goto st279
		case 32:
			goto st279
		}
		goto tr410
	st280:
		p++
	st_case_280:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 33:
			goto tr414
		case 37:
			goto tr414
		case 39:
			goto tr414
		case 126:
			goto tr414
		case 269:
			goto st764
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr414
				}
			case _widec >= 42:
				goto tr414
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr414
				}
			case _widec >= 65:
				goto tr414
			}
		default:
			goto tr414
		}
		goto st0
tr414:
//line sip.rl:62

	mark = p

//line sip.rl:54

	p--


	goto st281
	st281:
		p++
	st_case_281:
//line msg_parse.go:9317
		switch data[p] {
		case 65:
			goto st282
		case 66:
			goto st363
		case 67:
			goto st364
		case 68:
			goto st459
		case 69:
			goto st463
		case 70:
			goto st489
		case 73:
			goto st493
		case 75:
			goto st504
		case 76:
			goto st505
		case 77:
			goto st512
		case 79:
			goto st556
		case 80:
			goto st568
		case 82:
			goto st626
		case 83:
			goto st686
		case 84:
			goto st704
		case 85:
			goto st714
		case 86:
			goto st734
		case 87:
			goto st742
		case 97:
			goto st282
		case 98:
			goto st363
		case 99:
			goto st364
		case 100:
			goto st459
		case 101:
			goto st463
		case 102:
			goto st489
		case 105:
			goto st493
		case 107:
			goto st504
		case 108:
			goto st505
		case 109:
			goto st512
		case 111:
			goto st556
		case 112:
			goto st568
		case 114:
			goto st626
		case 115:
			goto st686
		case 116:
			goto st704
		case 117:
			goto st714
		case 118:
			goto st734
		case 119:
			goto st742
		}
		goto tr416
	st282:
		p++
	st_case_282:
		switch data[p] {
		case 9:
			goto tr435
		case 32:
			goto tr435
		case 58:
			goto tr436
		case 67:
			goto st288
		case 76:
			goto st317
		case 85:
			goto st336
		case 99:
			goto st288
		case 108:
			goto st317
		case 117:
			goto st336
		}
		goto tr416
tr435:
//line sip.rl:442
value=&msg.AcceptContact
	goto st283
tr450:
//line sip.rl:441
value=&msg.Accept
	goto st283
tr469:
//line sip.rl:443
value=&msg.AcceptEncoding
	goto st283
tr478:
//line sip.rl:444
value=&msg.AcceptLanguage
	goto st283
tr489:
//line sip.rl:447
value=&msg.AlertInfo
	goto st283
tr493:
//line sip.rl:445
value=&msg.Allow
	goto st283
tr502:
//line sip.rl:446
value=&msg.AllowEvents
	goto st283
tr522:
//line sip.rl:448
value=&msg.AuthenticationInfo
	goto st283
tr532:
//line sip.rl:449
value=&msg.Authorization
	goto st283
tr534:
//line sip.rl:466
value=&msg.ReferredBy
	goto st283
tr564:
//line sip.rl:453
value=&msg.CallInfo
	goto st283
tr597:
//line sip.rl:450
value=&msg.ContentDisposition
	goto st283
tr606:
//line sip.rl:452
value=&msg.ContentEncoding
	goto st283
tr616:
//line sip.rl:451
value=&msg.ContentLanguage
	goto st283
tr651:
//line sip.rl:454
value=&msg.Date
	goto st283
tr664:
//line sip.rl:455
value=&msg.ErrorInfo
	goto st283
tr669:
//line sip.rl:456
value=&msg.Event
	goto st283
tr698:
//line sip.rl:457
value=&msg.InReplyTo
	goto st283
tr700:
//line sip.rl:471
value=&msg.Supported
	goto st283
tr739:
//line sip.rl:459
value=&msg.MIMEVersion
	goto st283
tr767:
//line sip.rl:460
value=&msg.Organization
	goto st283
tr797:
//line sip.rl:461
value=&msg.Priority
	goto st283
tr816:
//line sip.rl:462
value=&msg.ProxyAuthenticate
	goto st283
tr826:
//line sip.rl:463
value=&msg.ProxyAuthorization
	goto st283
tr834:
//line sip.rl:464
value=&msg.ProxyRequire
	goto st283
tr836:
//line sip.rl:465
value=&msg.ReferTo
	goto st283
tr886:
//line sip.rl:458
value=&msg.ReplyTo
	goto st283
tr892:
//line sip.rl:467
value=&msg.Require
	goto st283
tr902:
//line sip.rl:468
value=&msg.RetryAfter
	goto st283
tr909:
//line sip.rl:470
value=&msg.Subject
	goto st283
tr917:
//line sip.rl:469
value=&msg.Server
	goto st283
tr941:
//line sip.rl:472
value=&msg.Timestamp
	goto st283
tr943:
//line sip.rl:445
value=&msg.Allow
//line sip.rl:446
value=&msg.AllowEvents
	goto st283
tr956:
//line sip.rl:473
value=&msg.Unsupported
	goto st283
tr966:
//line sip.rl:474
value=&msg.UserAgent
	goto st283
tr983:
//line sip.rl:475
value=&msg.Warning
	goto st283
tr999:
//line sip.rl:476
value=&msg.WWWAuthenticate
	goto st283
	st283:
		p++
	st_case_283:
//line msg_parse.go:9570
		switch data[p] {
		case 9:
			goto st283
		case 32:
			goto st283
		case 58:
			goto st284
		}
		goto st0
tr436:
//line sip.rl:442
value=&msg.AcceptContact
	goto st284
tr452:
//line sip.rl:441
value=&msg.Accept
	goto st284
tr470:
//line sip.rl:443
value=&msg.AcceptEncoding
	goto st284
tr479:
//line sip.rl:444
value=&msg.AcceptLanguage
	goto st284
tr490:
//line sip.rl:447
value=&msg.AlertInfo
	goto st284
tr495:
//line sip.rl:445
value=&msg.Allow
	goto st284
tr503:
//line sip.rl:446
value=&msg.AllowEvents
	goto st284
tr523:
//line sip.rl:448
value=&msg.AuthenticationInfo
	goto st284
tr533:
//line sip.rl:449
value=&msg.Authorization
	goto st284
tr535:
//line sip.rl:466
value=&msg.ReferredBy
	goto st284
tr565:
//line sip.rl:453
value=&msg.CallInfo
	goto st284
tr598:
//line sip.rl:450
value=&msg.ContentDisposition
	goto st284
tr607:
//line sip.rl:452
value=&msg.ContentEncoding
	goto st284
tr617:
//line sip.rl:451
value=&msg.ContentLanguage
	goto st284
tr652:
//line sip.rl:454
value=&msg.Date
	goto st284
tr665:
//line sip.rl:455
value=&msg.ErrorInfo
	goto st284
tr670:
//line sip.rl:456
value=&msg.Event
	goto st284
tr699:
//line sip.rl:457
value=&msg.InReplyTo
	goto st284
tr701:
//line sip.rl:471
value=&msg.Supported
	goto st284
tr740:
//line sip.rl:459
value=&msg.MIMEVersion
	goto st284
tr768:
//line sip.rl:460
value=&msg.Organization
	goto st284
tr798:
//line sip.rl:461
value=&msg.Priority
	goto st284
tr817:
//line sip.rl:462
value=&msg.ProxyAuthenticate
	goto st284
tr827:
//line sip.rl:463
value=&msg.ProxyAuthorization
	goto st284
tr835:
//line sip.rl:464
value=&msg.ProxyRequire
	goto st284
tr837:
//line sip.rl:465
value=&msg.ReferTo
	goto st284
tr887:
//line sip.rl:458
value=&msg.ReplyTo
	goto st284
tr893:
//line sip.rl:467
value=&msg.Require
	goto st284
tr903:
//line sip.rl:468
value=&msg.RetryAfter
	goto st284
tr910:
//line sip.rl:470
value=&msg.Subject
	goto st284
tr918:
//line sip.rl:469
value=&msg.Server
	goto st284
tr942:
//line sip.rl:472
value=&msg.Timestamp
	goto st284
tr944:
//line sip.rl:445
value=&msg.Allow
//line sip.rl:446
value=&msg.AllowEvents
	goto st284
tr957:
//line sip.rl:473
value=&msg.Unsupported
	goto st284
tr967:
//line sip.rl:474
value=&msg.UserAgent
	goto st284
tr984:
//line sip.rl:475
value=&msg.Warning
	goto st284
tr1000:
//line sip.rl:476
value=&msg.WWWAuthenticate
	goto st284
	st284:
		p++
	st_case_284:
//line msg_parse.go:9733
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st284
		case 32:
			goto st284
		case 269:
			goto tr442
		case 525:
			goto st285
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr442
			}
		default:
			goto tr442
		}
		goto st0
tr559:
//line sip.rl:248
 {goto st280 } 
	goto st782
tr442:
//line sip.rl:54

	p--


//line sip.rl:249
 {goto st263 } 
	goto st782
tr541:
//line sip.rl:54

	p--


//line sip.rl:247
 {goto st34 } 
	goto st782
tr576:
//line sip.rl:525
value=nil
//line sip.rl:54

	p--


//line sip.rl:243
 {goto st256 } 
	goto st782
tr971:
//line sip.rl:119

	via = new(Via)

//line sip.rl:54

	p--


//line sip.rl:250
 {goto st103 } 
	goto st782
tr1001:
//line sip.rl:58

	{p++; cs = 782; goto _out }

	goto st782
	st782:
		p++
	st_case_782:
//line msg_parse.go:9815
		goto st0
	st285:
		p++
	st_case_285:
		if data[p] == 10 {
			goto st286
		}
		goto st0
	st286:
		p++
	st_case_286:
		switch data[p] {
		case 9:
			goto st287
		case 32:
			goto st287
		}
		goto st0
	st287:
		p++
	st_case_287:
		switch data[p] {
		case 9:
			goto st287
		case 32:
			goto st287
		}
		goto tr442
	st288:
		p++
	st_case_288:
		switch data[p] {
		case 67:
			goto st289
		case 99:
			goto st289
		}
		goto tr416
	st289:
		p++
	st_case_289:
		switch data[p] {
		case 69:
			goto st290
		case 101:
			goto st290
		}
		goto tr416
	st290:
		p++
	st_case_290:
		switch data[p] {
		case 80:
			goto st291
		case 112:
			goto st291
		}
		goto tr416
	st291:
		p++
	st_case_291:
		switch data[p] {
		case 84:
			goto st292
		case 116:
			goto st292
		}
		goto tr416
	st292:
		p++
	st_case_292:
		switch data[p] {
		case 9:
			goto tr450
		case 32:
			goto tr450
		case 45:
			goto st293
		case 58:
			goto tr452
		}
		goto tr416
	st293:
		p++
	st_case_293:
		switch data[p] {
		case 67:
			goto st294
		case 69:
			goto st301
		case 76:
			goto st309
		case 99:
			goto st294
		case 101:
			goto st301
		case 108:
			goto st309
		}
		goto tr416
	st294:
		p++
	st_case_294:
		switch data[p] {
		case 79:
			goto st295
		case 111:
			goto st295
		}
		goto tr416
	st295:
		p++
	st_case_295:
		switch data[p] {
		case 78:
			goto st296
		case 110:
			goto st296
		}
		goto tr416
	st296:
		p++
	st_case_296:
		switch data[p] {
		case 84:
			goto st297
		case 116:
			goto st297
		}
		goto tr416
	st297:
		p++
	st_case_297:
		switch data[p] {
		case 65:
			goto st298
		case 97:
			goto st298
		}
		goto tr416
	st298:
		p++
	st_case_298:
		switch data[p] {
		case 67:
			goto st299
		case 99:
			goto st299
		}
		goto tr416
	st299:
		p++
	st_case_299:
		switch data[p] {
		case 84:
			goto st300
		case 116:
			goto st300
		}
		goto tr416
	st300:
		p++
	st_case_300:
		switch data[p] {
		case 9:
			goto tr435
		case 32:
			goto tr435
		case 58:
			goto tr436
		}
		goto tr416
	st301:
		p++
	st_case_301:
		switch data[p] {
		case 78:
			goto st302
		case 110:
			goto st302
		}
		goto tr416
	st302:
		p++
	st_case_302:
		switch data[p] {
		case 67:
			goto st303
		case 99:
			goto st303
		}
		goto tr416
	st303:
		p++
	st_case_303:
		switch data[p] {
		case 79:
			goto st304
		case 111:
			goto st304
		}
		goto tr416
	st304:
		p++
	st_case_304:
		switch data[p] {
		case 68:
			goto st305
		case 100:
			goto st305
		}
		goto tr416
	st305:
		p++
	st_case_305:
		switch data[p] {
		case 73:
			goto st306
		case 105:
			goto st306
		}
		goto tr416
	st306:
		p++
	st_case_306:
		switch data[p] {
		case 78:
			goto st307
		case 110:
			goto st307
		}
		goto tr416
	st307:
		p++
	st_case_307:
		switch data[p] {
		case 71:
			goto st308
		case 103:
			goto st308
		}
		goto tr416
	st308:
		p++
	st_case_308:
		switch data[p] {
		case 9:
			goto tr469
		case 32:
			goto tr469
		case 58:
			goto tr470
		}
		goto tr416
	st309:
		p++
	st_case_309:
		switch data[p] {
		case 65:
			goto st310
		case 97:
			goto st310
		}
		goto tr416
	st310:
		p++
	st_case_310:
		switch data[p] {
		case 78:
			goto st311
		case 110:
			goto st311
		}
		goto tr416
	st311:
		p++
	st_case_311:
		switch data[p] {
		case 71:
			goto st312
		case 103:
			goto st312
		}
		goto tr416
	st312:
		p++
	st_case_312:
		switch data[p] {
		case 85:
			goto st313
		case 117:
			goto st313
		}
		goto tr416
	st313:
		p++
	st_case_313:
		switch data[p] {
		case 65:
			goto st314
		case 97:
			goto st314
		}
		goto tr416
	st314:
		p++
	st_case_314:
		switch data[p] {
		case 71:
			goto st315
		case 103:
			goto st315
		}
		goto tr416
	st315:
		p++
	st_case_315:
		switch data[p] {
		case 69:
			goto st316
		case 101:
			goto st316
		}
		goto tr416
	st316:
		p++
	st_case_316:
		switch data[p] {
		case 9:
			goto tr478
		case 32:
			goto tr478
		case 58:
			goto tr479
		}
		goto tr416
	st317:
		p++
	st_case_317:
		switch data[p] {
		case 69:
			goto st318
		case 76:
			goto st326
		case 101:
			goto st318
		case 108:
			goto st326
		}
		goto tr416
	st318:
		p++
	st_case_318:
		switch data[p] {
		case 82:
			goto st319
		case 114:
			goto st319
		}
		goto tr416
	st319:
		p++
	st_case_319:
		switch data[p] {
		case 84:
			goto st320
		case 116:
			goto st320
		}
		goto tr416
	st320:
		p++
	st_case_320:
		if data[p] == 45 {
			goto st321
		}
		goto tr416
	st321:
		p++
	st_case_321:
		switch data[p] {
		case 73:
			goto st322
		case 105:
			goto st322
		}
		goto tr416
	st322:
		p++
	st_case_322:
		switch data[p] {
		case 78:
			goto st323
		case 110:
			goto st323
		}
		goto tr416
	st323:
		p++
	st_case_323:
		switch data[p] {
		case 70:
			goto st324
		case 102:
			goto st324
		}
		goto tr416
	st324:
		p++
	st_case_324:
		switch data[p] {
		case 79:
			goto st325
		case 111:
			goto st325
		}
		goto tr416
	st325:
		p++
	st_case_325:
		switch data[p] {
		case 9:
			goto tr489
		case 32:
			goto tr489
		case 58:
			goto tr490
		}
		goto tr416
	st326:
		p++
	st_case_326:
		switch data[p] {
		case 79:
			goto st327
		case 111:
			goto st327
		}
		goto tr416
	st327:
		p++
	st_case_327:
		switch data[p] {
		case 87:
			goto st328
		case 119:
			goto st328
		}
		goto tr416
	st328:
		p++
	st_case_328:
		switch data[p] {
		case 9:
			goto tr493
		case 32:
			goto tr493
		case 45:
			goto st329
		case 58:
			goto tr495
		}
		goto tr416
	st329:
		p++
	st_case_329:
		switch data[p] {
		case 69:
			goto st330
		case 101:
			goto st330
		}
		goto tr416
	st330:
		p++
	st_case_330:
		switch data[p] {
		case 86:
			goto st331
		case 118:
			goto st331
		}
		goto tr416
	st331:
		p++
	st_case_331:
		switch data[p] {
		case 69:
			goto st332
		case 101:
			goto st332
		}
		goto tr416
	st332:
		p++
	st_case_332:
		switch data[p] {
		case 78:
			goto st333
		case 110:
			goto st333
		}
		goto tr416
	st333:
		p++
	st_case_333:
		switch data[p] {
		case 84:
			goto st334
		case 116:
			goto st334
		}
		goto tr416
	st334:
		p++
	st_case_334:
		switch data[p] {
		case 83:
			goto st335
		case 115:
			goto st335
		}
		goto tr416
	st335:
		p++
	st_case_335:
		switch data[p] {
		case 9:
			goto tr502
		case 32:
			goto tr502
		case 58:
			goto tr503
		}
		goto tr416
	st336:
		p++
	st_case_336:
		switch data[p] {
		case 84:
			goto st337
		case 116:
			goto st337
		}
		goto tr416
	st337:
		p++
	st_case_337:
		switch data[p] {
		case 72:
			goto st338
		case 104:
			goto st338
		}
		goto tr416
	st338:
		p++
	st_case_338:
		switch data[p] {
		case 69:
			goto st339
		case 79:
			goto st354
		case 101:
			goto st339
		case 111:
			goto st354
		}
		goto tr416
	st339:
		p++
	st_case_339:
		switch data[p] {
		case 78:
			goto st340
		case 110:
			goto st340
		}
		goto tr416
	st340:
		p++
	st_case_340:
		switch data[p] {
		case 84:
			goto st341
		case 116:
			goto st341
		}
		goto tr416
	st341:
		p++
	st_case_341:
		switch data[p] {
		case 73:
			goto st342
		case 105:
			goto st342
		}
		goto tr416
	st342:
		p++
	st_case_342:
		switch data[p] {
		case 67:
			goto st343
		case 99:
			goto st343
		}
		goto tr416
	st343:
		p++
	st_case_343:
		switch data[p] {
		case 65:
			goto st344
		case 97:
			goto st344
		}
		goto tr416
	st344:
		p++
	st_case_344:
		switch data[p] {
		case 84:
			goto st345
		case 116:
			goto st345
		}
		goto tr416
	st345:
		p++
	st_case_345:
		switch data[p] {
		case 73:
			goto st346
		case 105:
			goto st346
		}
		goto tr416
	st346:
		p++
	st_case_346:
		switch data[p] {
		case 79:
			goto st347
		case 111:
			goto st347
		}
		goto tr416
	st347:
		p++
	st_case_347:
		switch data[p] {
		case 78:
			goto st348
		case 110:
			goto st348
		}
		goto tr416
	st348:
		p++
	st_case_348:
		if data[p] == 45 {
			goto st349
		}
		goto tr416
	st349:
		p++
	st_case_349:
		switch data[p] {
		case 73:
			goto st350
		case 105:
			goto st350
		}
		goto tr416
	st350:
		p++
	st_case_350:
		switch data[p] {
		case 78:
			goto st351
		case 110:
			goto st351
		}
		goto tr416
	st351:
		p++
	st_case_351:
		switch data[p] {
		case 70:
			goto st352
		case 102:
			goto st352
		}
		goto tr416
	st352:
		p++
	st_case_352:
		switch data[p] {
		case 79:
			goto st353
		case 111:
			goto st353
		}
		goto tr416
	st353:
		p++
	st_case_353:
		switch data[p] {
		case 9:
			goto tr522
		case 32:
			goto tr522
		case 58:
			goto tr523
		}
		goto tr416
	st354:
		p++
	st_case_354:
		switch data[p] {
		case 82:
			goto st355
		case 114:
			goto st355
		}
		goto tr416
	st355:
		p++
	st_case_355:
		switch data[p] {
		case 73:
			goto st356
		case 105:
			goto st356
		}
		goto tr416
	st356:
		p++
	st_case_356:
		switch data[p] {
		case 90:
			goto st357
		case 122:
			goto st357
		}
		goto tr416
	st357:
		p++
	st_case_357:
		switch data[p] {
		case 65:
			goto st358
		case 97:
			goto st358
		}
		goto tr416
	st358:
		p++
	st_case_358:
		switch data[p] {
		case 84:
			goto st359
		case 116:
			goto st359
		}
		goto tr416
	st359:
		p++
	st_case_359:
		switch data[p] {
		case 73:
			goto st360
		case 105:
			goto st360
		}
		goto tr416
	st360:
		p++
	st_case_360:
		switch data[p] {
		case 79:
			goto st361
		case 111:
			goto st361
		}
		goto tr416
	st361:
		p++
	st_case_361:
		switch data[p] {
		case 78:
			goto st362
		case 110:
			goto st362
		}
		goto tr416
	st362:
		p++
	st_case_362:
		switch data[p] {
		case 9:
			goto tr532
		case 32:
			goto tr532
		case 58:
			goto tr533
		}
		goto tr416
	st363:
		p++
	st_case_363:
		switch data[p] {
		case 9:
			goto tr534
		case 32:
			goto tr534
		case 58:
			goto tr535
		}
		goto tr416
	st364:
		p++
	st_case_364:
		switch data[p] {
		case 9:
			goto st365
		case 32:
			goto st365
		case 58:
			goto st366
		case 65:
			goto st370
		case 79:
			goto st388
		case 83:
			goto st445
		case 97:
			goto st370
		case 111:
			goto st388
		case 115:
			goto st445
		}
		goto tr416
	st365:
		p++
	st_case_365:
		switch data[p] {
		case 9:
			goto st365
		case 32:
			goto st365
		case 58:
			goto st366
		}
		goto st0
	st366:
		p++
	st_case_366:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st366
		case 32:
			goto st366
		case 269:
			goto tr541
		case 525:
			goto st367
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr541
			}
		default:
			goto tr541
		}
		goto st0
	st367:
		p++
	st_case_367:
		if data[p] == 10 {
			goto st368
		}
		goto st0
	st368:
		p++
	st_case_368:
		switch data[p] {
		case 9:
			goto st369
		case 32:
			goto st369
		}
		goto st0
	st369:
		p++
	st_case_369:
		switch data[p] {
		case 9:
			goto st369
		case 32:
			goto st369
		}
		goto tr541
	st370:
		p++
	st_case_370:
		switch data[p] {
		case 76:
			goto st371
		case 108:
			goto st371
		}
		goto tr416
	st371:
		p++
	st_case_371:
		switch data[p] {
		case 76:
			goto st372
		case 108:
			goto st372
		}
		goto tr416
	st372:
		p++
	st_case_372:
		if data[p] == 45 {
			goto st373
		}
		goto tr416
	st373:
		p++
	st_case_373:
		switch data[p] {
		case 73:
			goto st374
		case 105:
			goto st374
		}
		goto tr416
	st374:
		p++
	st_case_374:
		switch data[p] {
		case 68:
			goto st375
		case 78:
			goto st385
		case 100:
			goto st375
		case 110:
			goto st385
		}
		goto tr416
	st375:
		p++
	st_case_375:
		switch data[p] {
		case 9:
			goto st376
		case 32:
			goto st376
		case 58:
			goto st377
		}
		goto tr416
	st376:
		p++
	st_case_376:
		switch data[p] {
		case 9:
			goto st376
		case 32:
			goto st376
		case 58:
			goto st377
		}
		goto st0
	st377:
		p++
	st_case_377:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st377
		case 32:
			goto st377
		case 37:
			goto tr553
		case 60:
			goto tr553
		case 525:
			goto st382
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto tr553
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto tr553
				}
			default:
				goto tr553
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto tr553
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto tr553
				}
			default:
				goto tr553
			}
		default:
			goto tr553
		}
		goto st0
tr553:
//line sip.rl:62

	mark = p

	goto st378
	st378:
		p++
	st_case_378:
//line msg_parse.go:10867
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 37:
			goto st378
		case 60:
			goto st378
		case 64:
			goto st379
		case 269:
			goto tr557
		}
		switch {
		case _widec < 45:
			switch {
			case _widec > 34:
				if 39 <= _widec && _widec <= 43 {
					goto st378
				}
			case _widec >= 33:
				goto st378
			}
		case _widec > 58:
			switch {
			case _widec < 95:
				if 62 <= _widec && _widec <= 93 {
					goto st378
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto st378
				}
			default:
				goto st378
			}
		default:
			goto st378
		}
		goto st0
	st379:
		p++
	st_case_379:
		switch data[p] {
		case 37:
			goto st380
		case 60:
			goto st380
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st380
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st380
				}
			default:
				goto st380
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st380
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st380
				}
			default:
				goto st380
			}
		default:
			goto st380
		}
		goto st0
	st380:
		p++
	st_case_380:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 37:
			goto st380
		case 60:
			goto st380
		case 269:
			goto tr557
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto st380
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto st380
				}
			default:
				goto st380
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto st380
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto st380
				}
			default:
				goto st380
			}
		default:
			goto st380
		}
		goto st0
tr557:
//line sip.rl:211

	msg.CallID = string(data[mark:p])

	goto st381
tr643:
//line sip.rl:227

	msg.CSeqMethod = string(data[mark:p])

	goto st381
	st381:
		p++
	st_case_381:
//line msg_parse.go:11016
		if data[p] == 10 {
			goto tr559
		}
		goto st0
	st382:
		p++
	st_case_382:
		if data[p] == 10 {
			goto st383
		}
		goto st0
	st383:
		p++
	st_case_383:
		switch data[p] {
		case 9:
			goto st384
		case 32:
			goto st384
		}
		goto st0
	st384:
		p++
	st_case_384:
		switch data[p] {
		case 9:
			goto st384
		case 32:
			goto st384
		case 37:
			goto tr553
		case 60:
			goto tr553
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto tr553
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto tr553
				}
			default:
				goto tr553
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto tr553
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto tr553
				}
			default:
				goto tr553
			}
		default:
			goto tr553
		}
		goto st0
	st385:
		p++
	st_case_385:
		switch data[p] {
		case 70:
			goto st386
		case 102:
			goto st386
		}
		goto tr416
	st386:
		p++
	st_case_386:
		switch data[p] {
		case 79:
			goto st387
		case 111:
			goto st387
		}
		goto tr416
	st387:
		p++
	st_case_387:
		switch data[p] {
		case 9:
			goto tr564
		case 32:
			goto tr564
		case 58:
			goto tr565
		}
		goto tr416
	st388:
		p++
	st_case_388:
		switch data[p] {
		case 78:
			goto st389
		case 110:
			goto st389
		}
		goto tr416
	st389:
		p++
	st_case_389:
		switch data[p] {
		case 84:
			goto st390
		case 116:
			goto st390
		}
		goto tr416
	st390:
		p++
	st_case_390:
		switch data[p] {
		case 65:
			goto st391
		case 69:
			goto st399
		case 97:
			goto st391
		case 101:
			goto st399
		}
		goto tr416
	st391:
		p++
	st_case_391:
		switch data[p] {
		case 67:
			goto st392
		case 99:
			goto st392
		}
		goto tr416
	st392:
		p++
	st_case_392:
		switch data[p] {
		case 84:
			goto st393
		case 116:
			goto st393
		}
		goto tr416
	st393:
		p++
	st_case_393:
		switch data[p] {
		case 9:
			goto tr572
		case 32:
			goto tr572
		case 58:
			goto tr573
		}
		goto tr416
tr572:
//line sip.rl:428
addrp=lastAddr(&msg.Contact)
	goto st394
tr683:
//line sip.rl:429
addrp=lastAddr(&msg.From)
	goto st394
tr788:
//line sip.rl:430
addrp=lastAddr(&msg.PAssertedIdentity)
	goto st394
tr855:
//line sip.rl:431
addrp=lastAddr(&msg.RecordRoute)
	goto st394
tr879:
//line sip.rl:432
addrp=lastAddr(&msg.RemotePartyID)
	goto st394
tr907:
//line sip.rl:433
addrp=lastAddr(&msg.Route)
	goto st394
tr930:
//line sip.rl:434
addrp=lastAddr(&msg.To)
	goto st394
	st394:
		p++
	st_case_394:
//line msg_parse.go:11211
		switch data[p] {
		case 9:
			goto st394
		case 32:
			goto st394
		case 58:
			goto st395
		}
		goto st0
tr573:
//line sip.rl:428
addrp=lastAddr(&msg.Contact)
	goto st395
tr684:
//line sip.rl:429
addrp=lastAddr(&msg.From)
	goto st395
tr789:
//line sip.rl:430
addrp=lastAddr(&msg.PAssertedIdentity)
	goto st395
tr856:
//line sip.rl:431
addrp=lastAddr(&msg.RecordRoute)
	goto st395
tr880:
//line sip.rl:432
addrp=lastAddr(&msg.RemotePartyID)
	goto st395
tr908:
//line sip.rl:433
addrp=lastAddr(&msg.Route)
	goto st395
tr931:
//line sip.rl:434
addrp=lastAddr(&msg.To)
	goto st395
	st395:
		p++
	st_case_395:
//line msg_parse.go:11252
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st395
		case 32:
			goto st395
		case 269:
			goto tr576
		case 525:
			goto st396
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr576
			}
		default:
			goto tr576
		}
		goto st0
	st396:
		p++
	st_case_396:
		if data[p] == 10 {
			goto st397
		}
		goto st0
	st397:
		p++
	st_case_397:
		switch data[p] {
		case 9:
			goto st398
		case 32:
			goto st398
		}
		goto st0
	st398:
		p++
	st_case_398:
		switch data[p] {
		case 9:
			goto st398
		case 32:
			goto st398
		}
		goto tr576
	st399:
		p++
	st_case_399:
		switch data[p] {
		case 78:
			goto st400
		case 110:
			goto st400
		}
		goto tr416
	st400:
		p++
	st_case_400:
		switch data[p] {
		case 84:
			goto st401
		case 116:
			goto st401
		}
		goto tr416
	st401:
		p++
	st_case_401:
		if data[p] == 45 {
			goto st402
		}
		goto tr416
	st402:
		p++
	st_case_402:
		switch data[p] {
		case 68:
			goto st403
		case 69:
			goto st414
		case 76:
			goto st422
		case 84:
			goto st441
		case 100:
			goto st403
		case 101:
			goto st414
		case 108:
			goto st422
		case 116:
			goto st441
		}
		goto tr416
	st403:
		p++
	st_case_403:
		switch data[p] {
		case 73:
			goto st404
		case 105:
			goto st404
		}
		goto tr416
	st404:
		p++
	st_case_404:
		switch data[p] {
		case 83:
			goto st405
		case 115:
			goto st405
		}
		goto tr416
	st405:
		p++
	st_case_405:
		switch data[p] {
		case 80:
			goto st406
		case 112:
			goto st406
		}
		goto tr416
	st406:
		p++
	st_case_406:
		switch data[p] {
		case 79:
			goto st407
		case 111:
			goto st407
		}
		goto tr416
	st407:
		p++
	st_case_407:
		switch data[p] {
		case 83:
			goto st408
		case 115:
			goto st408
		}
		goto tr416
	st408:
		p++
	st_case_408:
		switch data[p] {
		case 73:
			goto st409
		case 105:
			goto st409
		}
		goto tr416
	st409:
		p++
	st_case_409:
		switch data[p] {
		case 84:
			goto st410
		case 116:
			goto st410
		}
		goto tr416
	st410:
		p++
	st_case_410:
		switch data[p] {
		case 73:
			goto st411
		case 105:
			goto st411
		}
		goto tr416
	st411:
		p++
	st_case_411:
		switch data[p] {
		case 79:
			goto st412
		case 111:
			goto st412
		}
		goto tr416
	st412:
		p++
	st_case_412:
		switch data[p] {
		case 78:
			goto st413
		case 110:
			goto st413
		}
		goto tr416
	st413:
		p++
	st_case_413:
		switch data[p] {
		case 9:
			goto tr597
		case 32:
			goto tr597
		case 58:
			goto tr598
		}
		goto tr416
	st414:
		p++
	st_case_414:
		switch data[p] {
		case 78:
			goto st415
		case 110:
			goto st415
		}
		goto tr416
	st415:
		p++
	st_case_415:
		switch data[p] {
		case 67:
			goto st416
		case 99:
			goto st416
		}
		goto tr416
	st416:
		p++
	st_case_416:
		switch data[p] {
		case 79:
			goto st417
		case 111:
			goto st417
		}
		goto tr416
	st417:
		p++
	st_case_417:
		switch data[p] {
		case 68:
			goto st418
		case 100:
			goto st418
		}
		goto tr416
	st418:
		p++
	st_case_418:
		switch data[p] {
		case 73:
			goto st419
		case 105:
			goto st419
		}
		goto tr416
	st419:
		p++
	st_case_419:
		switch data[p] {
		case 78:
			goto st420
		case 110:
			goto st420
		}
		goto tr416
	st420:
		p++
	st_case_420:
		switch data[p] {
		case 71:
			goto st421
		case 103:
			goto st421
		}
		goto tr416
	st421:
		p++
	st_case_421:
		switch data[p] {
		case 9:
			goto tr606
		case 32:
			goto tr606
		case 58:
			goto tr607
		}
		goto tr416
	st422:
		p++
	st_case_422:
		switch data[p] {
		case 65:
			goto st423
		case 69:
			goto st430
		case 97:
			goto st423
		case 101:
			goto st430
		}
		goto tr416
	st423:
		p++
	st_case_423:
		switch data[p] {
		case 78:
			goto st424
		case 110:
			goto st424
		}
		goto tr416
	st424:
		p++
	st_case_424:
		switch data[p] {
		case 71:
			goto st425
		case 103:
			goto st425
		}
		goto tr416
	st425:
		p++
	st_case_425:
		switch data[p] {
		case 85:
			goto st426
		case 117:
			goto st426
		}
		goto tr416
	st426:
		p++
	st_case_426:
		switch data[p] {
		case 65:
			goto st427
		case 97:
			goto st427
		}
		goto tr416
	st427:
		p++
	st_case_427:
		switch data[p] {
		case 71:
			goto st428
		case 103:
			goto st428
		}
		goto tr416
	st428:
		p++
	st_case_428:
		switch data[p] {
		case 69:
			goto st429
		case 101:
			goto st429
		}
		goto tr416
	st429:
		p++
	st_case_429:
		switch data[p] {
		case 9:
			goto tr616
		case 32:
			goto tr616
		case 58:
			goto tr617
		}
		goto tr416
	st430:
		p++
	st_case_430:
		switch data[p] {
		case 78:
			goto st431
		case 110:
			goto st431
		}
		goto tr416
	st431:
		p++
	st_case_431:
		switch data[p] {
		case 71:
			goto st432
		case 103:
			goto st432
		}
		goto tr416
	st432:
		p++
	st_case_432:
		switch data[p] {
		case 84:
			goto st433
		case 116:
			goto st433
		}
		goto tr416
	st433:
		p++
	st_case_433:
		switch data[p] {
		case 72:
			goto st434
		case 104:
			goto st434
		}
		goto tr416
	st434:
		p++
	st_case_434:
		switch data[p] {
		case 9:
			goto st435
		case 32:
			goto st435
		case 58:
			goto st436
		}
		goto tr416
	st435:
		p++
	st_case_435:
		switch data[p] {
		case 9:
			goto st435
		case 32:
			goto st435
		case 58:
			goto st436
		}
		goto st0
	st436:
		p++
	st_case_436:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st436
		case 32:
			goto st436
		case 525:
			goto st438
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr624
		}
		goto st0
tr624:
//line sip.rl:484
clen=0
//line sip.rl:215

	clen = clen * 10 + (int(data[p]) - 0x30)

	goto st437
tr626:
//line sip.rl:215

	clen = clen * 10 + (int(data[p]) - 0x30)

	goto st437
	st437:
		p++
	st_case_437:
//line msg_parse.go:11738
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st381
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr626
		}
		goto st0
	st438:
		p++
	st_case_438:
		if data[p] == 10 {
			goto st439
		}
		goto st0
	st439:
		p++
	st_case_439:
		switch data[p] {
		case 9:
			goto st440
		case 32:
			goto st440
		}
		goto st0
	st440:
		p++
	st_case_440:
		switch data[p] {
		case 9:
			goto st440
		case 32:
			goto st440
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr624
		}
		goto st0
	st441:
		p++
	st_case_441:
		switch data[p] {
		case 89:
			goto st442
		case 121:
			goto st442
		}
		goto tr416
	st442:
		p++
	st_case_442:
		switch data[p] {
		case 80:
			goto st443
		case 112:
			goto st443
		}
		goto tr416
	st443:
		p++
	st_case_443:
		switch data[p] {
		case 69:
			goto st444
		case 101:
			goto st444
		}
		goto tr416
	st444:
		p++
	st_case_444:
		switch data[p] {
		case 9:
			goto st365
		case 32:
			goto st365
		case 58:
			goto st366
		}
		goto tr416
	st445:
		p++
	st_case_445:
		switch data[p] {
		case 69:
			goto st446
		case 101:
			goto st446
		}
		goto tr416
	st446:
		p++
	st_case_446:
		switch data[p] {
		case 81:
			goto st447
		case 113:
			goto st447
		}
		goto tr416
	st447:
		p++
	st_case_447:
		switch data[p] {
		case 9:
			goto st448
		case 32:
			goto st448
		case 58:
			goto st449
		}
		goto tr416
	st448:
		p++
	st_case_448:
		switch data[p] {
		case 9:
			goto st448
		case 32:
			goto st448
		case 58:
			goto st449
		}
		goto st0
	st449:
		p++
	st_case_449:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st449
		case 32:
			goto st449
		case 525:
			goto st456
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr637
		}
		goto st0
tr637:
//line sip.rl:223

	msg.CSeq = msg.CSeq * 10 + (int(data[p]) - 0x30)

	goto st450
	st450:
		p++
	st_case_450:
//line msg_parse.go:11900
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st451
		case 32:
			goto st451
		case 525:
			goto st453
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr637
		}
		goto st0
	st451:
		p++
	st_case_451:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st451
		case 32:
			goto st451
		case 33:
			goto tr641
		case 37:
			goto tr641
		case 39:
			goto tr641
		case 126:
			goto tr641
		case 525:
			goto st453
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr641
				}
			case _widec >= 42:
				goto tr641
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr641
				}
			case _widec >= 65:
				goto tr641
			}
		default:
			goto tr641
		}
		goto st0
tr641:
//line sip.rl:62

	mark = p

	goto st452
	st452:
		p++
	st_case_452:
//line msg_parse.go:11978
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 33:
			goto st452
		case 37:
			goto st452
		case 39:
			goto st452
		case 126:
			goto st452
		case 269:
			goto tr643
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st452
				}
			case _widec >= 42:
				goto st452
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st452
				}
			case _widec >= 65:
				goto st452
			}
		default:
			goto st452
		}
		goto st0
	st453:
		p++
	st_case_453:
		if data[p] == 10 {
			goto st454
		}
		goto st0
	st454:
		p++
	st_case_454:
		switch data[p] {
		case 9:
			goto st455
		case 32:
			goto st455
		}
		goto st0
	st455:
		p++
	st_case_455:
		switch data[p] {
		case 9:
			goto st455
		case 32:
			goto st455
		case 33:
			goto tr641
		case 37:
			goto tr641
		case 39:
			goto tr641
		case 126:
			goto tr641
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr641
				}
			case data[p] >= 42:
				goto tr641
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr641
				}
			case data[p] >= 65:
				goto tr641
			}
		default:
			goto tr641
		}
		goto st0
	st456:
		p++
	st_case_456:
		if data[p] == 10 {
			goto st457
		}
		goto st0
	st457:
		p++
	st_case_457:
		switch data[p] {
		case 9:
			goto st458
		case 32:
			goto st458
		}
		goto st0
	st458:
		p++
	st_case_458:
		switch data[p] {
		case 9:
			goto st458
		case 32:
			goto st458
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr637
		}
		goto st0
	st459:
		p++
	st_case_459:
		switch data[p] {
		case 65:
			goto st460
		case 97:
			goto st460
		}
		goto tr416
	st460:
		p++
	st_case_460:
		switch data[p] {
		case 84:
			goto st461
		case 116:
			goto st461
		}
		goto tr416
	st461:
		p++
	st_case_461:
		switch data[p] {
		case 69:
			goto st462
		case 101:
			goto st462
		}
		goto tr416
	st462:
		p++
	st_case_462:
		switch data[p] {
		case 9:
			goto tr651
		case 32:
			goto tr651
		case 58:
			goto tr652
		}
		goto tr416
	st463:
		p++
	st_case_463:
		switch data[p] {
		case 9:
			goto tr606
		case 32:
			goto tr606
		case 58:
			goto tr607
		case 82:
			goto st464
		case 86:
			goto st473
		case 88:
			goto st477
		case 114:
			goto st464
		case 118:
			goto st473
		case 120:
			goto st477
		}
		goto tr416
	st464:
		p++
	st_case_464:
		switch data[p] {
		case 82:
			goto st465
		case 114:
			goto st465
		}
		goto tr416
	st465:
		p++
	st_case_465:
		switch data[p] {
		case 79:
			goto st466
		case 111:
			goto st466
		}
		goto tr416
	st466:
		p++
	st_case_466:
		switch data[p] {
		case 82:
			goto st467
		case 114:
			goto st467
		}
		goto tr416
	st467:
		p++
	st_case_467:
		if data[p] == 45 {
			goto st468
		}
		goto tr416
	st468:
		p++
	st_case_468:
		switch data[p] {
		case 73:
			goto st469
		case 105:
			goto st469
		}
		goto tr416
	st469:
		p++
	st_case_469:
		switch data[p] {
		case 78:
			goto st470
		case 110:
			goto st470
		}
		goto tr416
	st470:
		p++
	st_case_470:
		switch data[p] {
		case 70:
			goto st471
		case 102:
			goto st471
		}
		goto tr416
	st471:
		p++
	st_case_471:
		switch data[p] {
		case 79:
			goto st472
		case 111:
			goto st472
		}
		goto tr416
	st472:
		p++
	st_case_472:
		switch data[p] {
		case 9:
			goto tr664
		case 32:
			goto tr664
		case 58:
			goto tr665
		}
		goto tr416
	st473:
		p++
	st_case_473:
		switch data[p] {
		case 69:
			goto st474
		case 101:
			goto st474
		}
		goto tr416
	st474:
		p++
	st_case_474:
		switch data[p] {
		case 78:
			goto st475
		case 110:
			goto st475
		}
		goto tr416
	st475:
		p++
	st_case_475:
		switch data[p] {
		case 84:
			goto st476
		case 116:
			goto st476
		}
		goto tr416
	st476:
		p++
	st_case_476:
		switch data[p] {
		case 9:
			goto tr669
		case 32:
			goto tr669
		case 58:
			goto tr670
		}
		goto tr416
	st477:
		p++
	st_case_477:
		switch data[p] {
		case 80:
			goto st478
		case 112:
			goto st478
		}
		goto tr416
	st478:
		p++
	st_case_478:
		switch data[p] {
		case 73:
			goto st479
		case 105:
			goto st479
		}
		goto tr416
	st479:
		p++
	st_case_479:
		switch data[p] {
		case 82:
			goto st480
		case 114:
			goto st480
		}
		goto tr416
	st480:
		p++
	st_case_480:
		switch data[p] {
		case 69:
			goto st481
		case 101:
			goto st481
		}
		goto tr416
	st481:
		p++
	st_case_481:
		switch data[p] {
		case 83:
			goto st482
		case 115:
			goto st482
		}
		goto tr416
	st482:
		p++
	st_case_482:
		switch data[p] {
		case 9:
			goto st483
		case 32:
			goto st483
		case 58:
			goto st484
		}
		goto tr416
	st483:
		p++
	st_case_483:
		switch data[p] {
		case 9:
			goto st483
		case 32:
			goto st483
		case 58:
			goto st484
		}
		goto st0
	st484:
		p++
	st_case_484:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st484
		case 32:
			goto st484
		case 525:
			goto st486
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr678
		}
		goto st0
tr678:
//line sip.rl:486
msg.Expires=0
//line sip.rl:231

	msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)

	goto st485
tr680:
//line sip.rl:231

	msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)

	goto st485
	st485:
		p++
	st_case_485:
//line msg_parse.go:12418
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st381
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr680
		}
		goto st0
	st486:
		p++
	st_case_486:
		if data[p] == 10 {
			goto st487
		}
		goto st0
	st487:
		p++
	st_case_487:
		switch data[p] {
		case 9:
			goto st488
		case 32:
			goto st488
		}
		goto st0
	st488:
		p++
	st_case_488:
		switch data[p] {
		case 9:
			goto st488
		case 32:
			goto st488
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr678
		}
		goto st0
	st489:
		p++
	st_case_489:
		switch data[p] {
		case 9:
			goto tr683
		case 32:
			goto tr683
		case 58:
			goto tr684
		case 82:
			goto st490
		case 114:
			goto st490
		}
		goto tr416
	st490:
		p++
	st_case_490:
		switch data[p] {
		case 79:
			goto st491
		case 111:
			goto st491
		}
		goto tr416
	st491:
		p++
	st_case_491:
		switch data[p] {
		case 77:
			goto st492
		case 109:
			goto st492
		}
		goto tr416
	st492:
		p++
	st_case_492:
		switch data[p] {
		case 9:
			goto tr683
		case 32:
			goto tr683
		case 58:
			goto tr684
		}
		goto tr416
	st493:
		p++
	st_case_493:
		switch data[p] {
		case 9:
			goto st376
		case 32:
			goto st376
		case 58:
			goto st377
		case 78:
			goto st494
		case 110:
			goto st494
		}
		goto tr416
	st494:
		p++
	st_case_494:
		if data[p] == 45 {
			goto st495
		}
		goto tr416
	st495:
		p++
	st_case_495:
		switch data[p] {
		case 82:
			goto st496
		case 114:
			goto st496
		}
		goto tr416
	st496:
		p++
	st_case_496:
		switch data[p] {
		case 69:
			goto st497
		case 101:
			goto st497
		}
		goto tr416
	st497:
		p++
	st_case_497:
		switch data[p] {
		case 80:
			goto st498
		case 112:
			goto st498
		}
		goto tr416
	st498:
		p++
	st_case_498:
		switch data[p] {
		case 76:
			goto st499
		case 108:
			goto st499
		}
		goto tr416
	st499:
		p++
	st_case_499:
		switch data[p] {
		case 89:
			goto st500
		case 121:
			goto st500
		}
		goto tr416
	st500:
		p++
	st_case_500:
		if data[p] == 45 {
			goto st501
		}
		goto tr416
	st501:
		p++
	st_case_501:
		switch data[p] {
		case 84:
			goto st502
		case 116:
			goto st502
		}
		goto tr416
	st502:
		p++
	st_case_502:
		switch data[p] {
		case 79:
			goto st503
		case 111:
			goto st503
		}
		goto tr416
	st503:
		p++
	st_case_503:
		switch data[p] {
		case 9:
			goto tr698
		case 32:
			goto tr698
		case 58:
			goto tr699
		}
		goto tr416
	st504:
		p++
	st_case_504:
		switch data[p] {
		case 9:
			goto tr700
		case 32:
			goto tr700
		case 58:
			goto tr701
		}
		goto tr416
	st505:
		p++
	st_case_505:
		switch data[p] {
		case 9:
			goto st506
		case 32:
			goto st506
		case 58:
			goto st507
		}
		goto tr416
	st506:
		p++
	st_case_506:
		switch data[p] {
		case 9:
			goto st506
		case 32:
			goto st506
		case 58:
			goto st507
		}
		goto st0
	st507:
		p++
	st_case_507:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st507
		case 32:
			goto st507
		case 525:
			goto st509
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr704
		}
		goto st0
tr704:
//line sip.rl:484
clen=0
//line sip.rl:215

	clen = clen * 10 + (int(data[p]) - 0x30)

//line sip.rl:486
msg.Expires=0
//line sip.rl:231

	msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)

//line sip.rl:487
msg.MaxForwards=0
//line sip.rl:235

	msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)

//line sip.rl:488
msg.MinExpires=0
//line sip.rl:239

	msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)

	goto st508
tr706:
//line sip.rl:215

	clen = clen * 10 + (int(data[p]) - 0x30)

//line sip.rl:231

	msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)

//line sip.rl:235

	msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)

//line sip.rl:239

	msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)

	goto st508
	st508:
		p++
	st_case_508:
//line msg_parse.go:12728
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st381
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr706
		}
		goto st0
	st509:
		p++
	st_case_509:
		if data[p] == 10 {
			goto st510
		}
		goto st0
	st510:
		p++
	st_case_510:
		switch data[p] {
		case 9:
			goto st511
		case 32:
			goto st511
		}
		goto st0
	st511:
		p++
	st_case_511:
		switch data[p] {
		case 9:
			goto st511
		case 32:
			goto st511
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr704
		}
		goto st0
	st512:
		p++
	st_case_512:
		switch data[p] {
		case 9:
			goto tr572
		case 32:
			goto tr572
		case 58:
			goto tr573
		case 65:
			goto st513
		case 73:
			goto st530
		case 97:
			goto st513
		case 105:
			goto st530
		}
		goto tr416
	st513:
		p++
	st_case_513:
		switch data[p] {
		case 88:
			goto st514
		case 120:
			goto st514
		}
		goto tr416
	st514:
		p++
	st_case_514:
		if data[p] == 45 {
			goto st515
		}
		goto tr416
	st515:
		p++
	st_case_515:
		switch data[p] {
		case 70:
			goto st516
		case 102:
			goto st516
		}
		goto tr416
	st516:
		p++
	st_case_516:
		switch data[p] {
		case 79:
			goto st517
		case 111:
			goto st517
		}
		goto tr416
	st517:
		p++
	st_case_517:
		switch data[p] {
		case 82:
			goto st518
		case 114:
			goto st518
		}
		goto tr416
	st518:
		p++
	st_case_518:
		switch data[p] {
		case 87:
			goto st519
		case 119:
			goto st519
		}
		goto tr416
	st519:
		p++
	st_case_519:
		switch data[p] {
		case 65:
			goto st520
		case 97:
			goto st520
		}
		goto tr416
	st520:
		p++
	st_case_520:
		switch data[p] {
		case 82:
			goto st521
		case 114:
			goto st521
		}
		goto tr416
	st521:
		p++
	st_case_521:
		switch data[p] {
		case 68:
			goto st522
		case 100:
			goto st522
		}
		goto tr416
	st522:
		p++
	st_case_522:
		switch data[p] {
		case 83:
			goto st523
		case 115:
			goto st523
		}
		goto tr416
	st523:
		p++
	st_case_523:
		switch data[p] {
		case 9:
			goto st524
		case 32:
			goto st524
		case 58:
			goto st525
		}
		goto tr416
	st524:
		p++
	st_case_524:
		switch data[p] {
		case 9:
			goto st524
		case 32:
			goto st524
		case 58:
			goto st525
		}
		goto st0
	st525:
		p++
	st_case_525:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st525
		case 32:
			goto st525
		case 525:
			goto st527
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr723
		}
		goto st0
tr723:
//line sip.rl:487
msg.MaxForwards=0
//line sip.rl:235

	msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)

	goto st526
tr725:
//line sip.rl:235

	msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)

	goto st526
	st526:
		p++
	st_case_526:
//line msg_parse.go:12953
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st381
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr725
		}
		goto st0
	st527:
		p++
	st_case_527:
		if data[p] == 10 {
			goto st528
		}
		goto st0
	st528:
		p++
	st_case_528:
		switch data[p] {
		case 9:
			goto st529
		case 32:
			goto st529
		}
		goto st0
	st529:
		p++
	st_case_529:
		switch data[p] {
		case 9:
			goto st529
		case 32:
			goto st529
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr723
		}
		goto st0
	st530:
		p++
	st_case_530:
		switch data[p] {
		case 77:
			goto st531
		case 78:
			goto st541
		case 109:
			goto st531
		case 110:
			goto st541
		}
		goto tr416
	st531:
		p++
	st_case_531:
		switch data[p] {
		case 69:
			goto st532
		case 101:
			goto st532
		}
		goto tr416
	st532:
		p++
	st_case_532:
		if data[p] == 45 {
			goto st533
		}
		goto tr416
	st533:
		p++
	st_case_533:
		switch data[p] {
		case 86:
			goto st534
		case 118:
			goto st534
		}
		goto tr416
	st534:
		p++
	st_case_534:
		switch data[p] {
		case 69:
			goto st535
		case 101:
			goto st535
		}
		goto tr416
	st535:
		p++
	st_case_535:
		switch data[p] {
		case 82:
			goto st536
		case 114:
			goto st536
		}
		goto tr416
	st536:
		p++
	st_case_536:
		switch data[p] {
		case 83:
			goto st537
		case 115:
			goto st537
		}
		goto tr416
	st537:
		p++
	st_case_537:
		switch data[p] {
		case 73:
			goto st538
		case 105:
			goto st538
		}
		goto tr416
	st538:
		p++
	st_case_538:
		switch data[p] {
		case 79:
			goto st539
		case 111:
			goto st539
		}
		goto tr416
	st539:
		p++
	st_case_539:
		switch data[p] {
		case 78:
			goto st540
		case 110:
			goto st540
		}
		goto tr416
	st540:
		p++
	st_case_540:
		switch data[p] {
		case 9:
			goto tr739
		case 32:
			goto tr739
		case 58:
			goto tr740
		}
		goto tr416
	st541:
		p++
	st_case_541:
		if data[p] == 45 {
			goto st542
		}
		goto tr416
	st542:
		p++
	st_case_542:
		switch data[p] {
		case 69:
			goto st543
		case 101:
			goto st543
		}
		goto tr416
	st543:
		p++
	st_case_543:
		switch data[p] {
		case 88:
			goto st544
		case 120:
			goto st544
		}
		goto tr416
	st544:
		p++
	st_case_544:
		switch data[p] {
		case 80:
			goto st545
		case 112:
			goto st545
		}
		goto tr416
	st545:
		p++
	st_case_545:
		switch data[p] {
		case 73:
			goto st546
		case 105:
			goto st546
		}
		goto tr416
	st546:
		p++
	st_case_546:
		switch data[p] {
		case 82:
			goto st547
		case 114:
			goto st547
		}
		goto tr416
	st547:
		p++
	st_case_547:
		switch data[p] {
		case 69:
			goto st548
		case 101:
			goto st548
		}
		goto tr416
	st548:
		p++
	st_case_548:
		switch data[p] {
		case 83:
			goto st549
		case 115:
			goto st549
		}
		goto tr416
	st549:
		p++
	st_case_549:
		switch data[p] {
		case 9:
			goto st550
		case 32:
			goto st550
		case 58:
			goto st551
		}
		goto tr416
	st550:
		p++
	st_case_550:
		switch data[p] {
		case 9:
			goto st550
		case 32:
			goto st550
		case 58:
			goto st551
		}
		goto st0
	st551:
		p++
	st_case_551:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st551
		case 32:
			goto st551
		case 525:
			goto st553
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr751
		}
		goto st0
tr751:
//line sip.rl:488
msg.MinExpires=0
//line sip.rl:239

	msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)

	goto st552
tr753:
//line sip.rl:239

	msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)

	goto st552
	st552:
		p++
	st_case_552:
//line msg_parse.go:13251
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st381
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr753
		}
		goto st0
	st553:
		p++
	st_case_553:
		if data[p] == 10 {
			goto st554
		}
		goto st0
	st554:
		p++
	st_case_554:
		switch data[p] {
		case 9:
			goto st555
		case 32:
			goto st555
		}
		goto st0
	st555:
		p++
	st_case_555:
		switch data[p] {
		case 9:
			goto st555
		case 32:
			goto st555
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr751
		}
		goto st0
	st556:
		p++
	st_case_556:
		switch data[p] {
		case 9:
			goto tr669
		case 32:
			goto tr669
		case 58:
			goto tr670
		case 82:
			goto st557
		case 114:
			goto st557
		}
		goto tr416
	st557:
		p++
	st_case_557:
		switch data[p] {
		case 71:
			goto st558
		case 103:
			goto st558
		}
		goto tr416
	st558:
		p++
	st_case_558:
		switch data[p] {
		case 65:
			goto st559
		case 97:
			goto st559
		}
		goto tr416
	st559:
		p++
	st_case_559:
		switch data[p] {
		case 78:
			goto st560
		case 110:
			goto st560
		}
		goto tr416
	st560:
		p++
	st_case_560:
		switch data[p] {
		case 73:
			goto st561
		case 105:
			goto st561
		}
		goto tr416
	st561:
		p++
	st_case_561:
		switch data[p] {
		case 90:
			goto st562
		case 122:
			goto st562
		}
		goto tr416
	st562:
		p++
	st_case_562:
		switch data[p] {
		case 65:
			goto st563
		case 97:
			goto st563
		}
		goto tr416
	st563:
		p++
	st_case_563:
		switch data[p] {
		case 84:
			goto st564
		case 116:
			goto st564
		}
		goto tr416
	st564:
		p++
	st_case_564:
		switch data[p] {
		case 73:
			goto st565
		case 105:
			goto st565
		}
		goto tr416
	st565:
		p++
	st_case_565:
		switch data[p] {
		case 79:
			goto st566
		case 111:
			goto st566
		}
		goto tr416
	st566:
		p++
	st_case_566:
		switch data[p] {
		case 78:
			goto st567
		case 110:
			goto st567
		}
		goto tr416
	st567:
		p++
	st_case_567:
		switch data[p] {
		case 9:
			goto tr767
		case 32:
			goto tr767
		case 58:
			goto tr768
		}
		goto tr416
	st568:
		p++
	st_case_568:
		switch data[p] {
		case 45:
			goto st569
		case 82:
			goto st587
		case 114:
			goto st587
		}
		goto tr416
	st569:
		p++
	st_case_569:
		switch data[p] {
		case 65:
			goto st570
		case 97:
			goto st570
		}
		goto tr416
	st570:
		p++
	st_case_570:
		switch data[p] {
		case 83:
			goto st571
		case 115:
			goto st571
		}
		goto tr416
	st571:
		p++
	st_case_571:
		switch data[p] {
		case 83:
			goto st572
		case 115:
			goto st572
		}
		goto tr416
	st572:
		p++
	st_case_572:
		switch data[p] {
		case 69:
			goto st573
		case 101:
			goto st573
		}
		goto tr416
	st573:
		p++
	st_case_573:
		switch data[p] {
		case 82:
			goto st574
		case 114:
			goto st574
		}
		goto tr416
	st574:
		p++
	st_case_574:
		switch data[p] {
		case 84:
			goto st575
		case 116:
			goto st575
		}
		goto tr416
	st575:
		p++
	st_case_575:
		switch data[p] {
		case 69:
			goto st576
		case 101:
			goto st576
		}
		goto tr416
	st576:
		p++
	st_case_576:
		switch data[p] {
		case 68:
			goto st577
		case 100:
			goto st577
		}
		goto tr416
	st577:
		p++
	st_case_577:
		if data[p] == 45 {
			goto st578
		}
		goto tr416
	st578:
		p++
	st_case_578:
		switch data[p] {
		case 73:
			goto st579
		case 105:
			goto st579
		}
		goto tr416
	st579:
		p++
	st_case_579:
		switch data[p] {
		case 68:
			goto st580
		case 100:
			goto st580
		}
		goto tr416
	st580:
		p++
	st_case_580:
		switch data[p] {
		case 69:
			goto st581
		case 101:
			goto st581
		}
		goto tr416
	st581:
		p++
	st_case_581:
		switch data[p] {
		case 78:
			goto st582
		case 110:
			goto st582
		}
		goto tr416
	st582:
		p++
	st_case_582:
		switch data[p] {
		case 84:
			goto st583
		case 116:
			goto st583
		}
		goto tr416
	st583:
		p++
	st_case_583:
		switch data[p] {
		case 73:
			goto st584
		case 105:
			goto st584
		}
		goto tr416
	st584:
		p++
	st_case_584:
		switch data[p] {
		case 84:
			goto st585
		case 116:
			goto st585
		}
		goto tr416
	st585:
		p++
	st_case_585:
		switch data[p] {
		case 89:
			goto st586
		case 121:
			goto st586
		}
		goto tr416
	st586:
		p++
	st_case_586:
		switch data[p] {
		case 9:
			goto tr788
		case 32:
			goto tr788
		case 58:
			goto tr789
		}
		goto tr416
	st587:
		p++
	st_case_587:
		switch data[p] {
		case 73:
			goto st588
		case 79:
			goto st594
		case 105:
			goto st588
		case 111:
			goto st594
		}
		goto tr416
	st588:
		p++
	st_case_588:
		switch data[p] {
		case 79:
			goto st589
		case 111:
			goto st589
		}
		goto tr416
	st589:
		p++
	st_case_589:
		switch data[p] {
		case 82:
			goto st590
		case 114:
			goto st590
		}
		goto tr416
	st590:
		p++
	st_case_590:
		switch data[p] {
		case 73:
			goto st591
		case 105:
			goto st591
		}
		goto tr416
	st591:
		p++
	st_case_591:
		switch data[p] {
		case 84:
			goto st592
		case 116:
			goto st592
		}
		goto tr416
	st592:
		p++
	st_case_592:
		switch data[p] {
		case 89:
			goto st593
		case 121:
			goto st593
		}
		goto tr416
	st593:
		p++
	st_case_593:
		switch data[p] {
		case 9:
			goto tr797
		case 32:
			goto tr797
		case 58:
			goto tr798
		}
		goto tr416
	st594:
		p++
	st_case_594:
		switch data[p] {
		case 88:
			goto st595
		case 120:
			goto st595
		}
		goto tr416
	st595:
		p++
	st_case_595:
		switch data[p] {
		case 89:
			goto st596
		case 121:
			goto st596
		}
		goto tr416
	st596:
		p++
	st_case_596:
		if data[p] == 45 {
			goto st597
		}
		goto tr416
	st597:
		p++
	st_case_597:
		switch data[p] {
		case 65:
			goto st598
		case 82:
			goto st619
		case 97:
			goto st598
		case 114:
			goto st619
		}
		goto tr416
	st598:
		p++
	st_case_598:
		switch data[p] {
		case 85:
			goto st599
		case 117:
			goto st599
		}
		goto tr416
	st599:
		p++
	st_case_599:
		switch data[p] {
		case 84:
			goto st600
		case 116:
			goto st600
		}
		goto tr416
	st600:
		p++
	st_case_600:
		switch data[p] {
		case 72:
			goto st601
		case 104:
			goto st601
		}
		goto tr416
	st601:
		p++
	st_case_601:
		switch data[p] {
		case 69:
			goto st602
		case 79:
			goto st610
		case 101:
			goto st602
		case 111:
			goto st610
		}
		goto tr416
	st602:
		p++
	st_case_602:
		switch data[p] {
		case 78:
			goto st603
		case 110:
			goto st603
		}
		goto tr416
	st603:
		p++
	st_case_603:
		switch data[p] {
		case 84:
			goto st604
		case 116:
			goto st604
		}
		goto tr416
	st604:
		p++
	st_case_604:
		switch data[p] {
		case 73:
			goto st605
		case 105:
			goto st605
		}
		goto tr416
	st605:
		p++
	st_case_605:
		switch data[p] {
		case 67:
			goto st606
		case 99:
			goto st606
		}
		goto tr416
	st606:
		p++
	st_case_606:
		switch data[p] {
		case 65:
			goto st607
		case 97:
			goto st607
		}
		goto tr416
	st607:
		p++
	st_case_607:
		switch data[p] {
		case 84:
			goto st608
		case 116:
			goto st608
		}
		goto tr416
	st608:
		p++
	st_case_608:
		switch data[p] {
		case 69:
			goto st609
		case 101:
			goto st609
		}
		goto tr416
	st609:
		p++
	st_case_609:
		switch data[p] {
		case 9:
			goto tr816
		case 32:
			goto tr816
		case 58:
			goto tr817
		}
		goto tr416
	st610:
		p++
	st_case_610:
		switch data[p] {
		case 82:
			goto st611
		case 114:
			goto st611
		}
		goto tr416
	st611:
		p++
	st_case_611:
		switch data[p] {
		case 73:
			goto st612
		case 105:
			goto st612
		}
		goto tr416
	st612:
		p++
	st_case_612:
		switch data[p] {
		case 90:
			goto st613
		case 122:
			goto st613
		}
		goto tr416
	st613:
		p++
	st_case_613:
		switch data[p] {
		case 65:
			goto st614
		case 97:
			goto st614
		}
		goto tr416
	st614:
		p++
	st_case_614:
		switch data[p] {
		case 84:
			goto st615
		case 116:
			goto st615
		}
		goto tr416
	st615:
		p++
	st_case_615:
		switch data[p] {
		case 73:
			goto st616
		case 105:
			goto st616
		}
		goto tr416
	st616:
		p++
	st_case_616:
		switch data[p] {
		case 79:
			goto st617
		case 111:
			goto st617
		}
		goto tr416
	st617:
		p++
	st_case_617:
		switch data[p] {
		case 78:
			goto st618
		case 110:
			goto st618
		}
		goto tr416
	st618:
		p++
	st_case_618:
		switch data[p] {
		case 9:
			goto tr826
		case 32:
			goto tr826
		case 58:
			goto tr827
		}
		goto tr416
	st619:
		p++
	st_case_619:
		switch data[p] {
		case 69:
			goto st620
		case 101:
			goto st620
		}
		goto tr416
	st620:
		p++
	st_case_620:
		switch data[p] {
		case 81:
			goto st621
		case 113:
			goto st621
		}
		goto tr416
	st621:
		p++
	st_case_621:
		switch data[p] {
		case 85:
			goto st622
		case 117:
			goto st622
		}
		goto tr416
	st622:
		p++
	st_case_622:
		switch data[p] {
		case 73:
			goto st623
		case 105:
			goto st623
		}
		goto tr416
	st623:
		p++
	st_case_623:
		switch data[p] {
		case 82:
			goto st624
		case 114:
			goto st624
		}
		goto tr416
	st624:
		p++
	st_case_624:
		switch data[p] {
		case 69:
			goto st625
		case 101:
			goto st625
		}
		goto tr416
	st625:
		p++
	st_case_625:
		switch data[p] {
		case 9:
			goto tr834
		case 32:
			goto tr834
		case 58:
			goto tr835
		}
		goto tr416
	st626:
		p++
	st_case_626:
		switch data[p] {
		case 9:
			goto tr836
		case 32:
			goto tr836
		case 58:
			goto tr837
		case 69:
			goto st627
		case 79:
			goto st682
		case 101:
			goto st627
		case 111:
			goto st682
		}
		goto tr416
	st627:
		p++
	st_case_627:
		switch data[p] {
		case 67:
			goto st628
		case 70:
			goto st638
		case 77:
			goto st649
		case 80:
			goto st662
		case 81:
			goto st668
		case 84:
			goto st673
		case 99:
			goto st628
		case 102:
			goto st638
		case 109:
			goto st649
		case 112:
			goto st662
		case 113:
			goto st668
		case 116:
			goto st673
		}
		goto tr416
	st628:
		p++
	st_case_628:
		switch data[p] {
		case 79:
			goto st629
		case 111:
			goto st629
		}
		goto tr416
	st629:
		p++
	st_case_629:
		switch data[p] {
		case 82:
			goto st630
		case 114:
			goto st630
		}
		goto tr416
	st630:
		p++
	st_case_630:
		switch data[p] {
		case 68:
			goto st631
		case 100:
			goto st631
		}
		goto tr416
	st631:
		p++
	st_case_631:
		if data[p] == 45 {
			goto st632
		}
		goto tr416
	st632:
		p++
	st_case_632:
		switch data[p] {
		case 82:
			goto st633
		case 114:
			goto st633
		}
		goto tr416
	st633:
		p++
	st_case_633:
		switch data[p] {
		case 79:
			goto st634
		case 111:
			goto st634
		}
		goto tr416
	st634:
		p++
	st_case_634:
		switch data[p] {
		case 85:
			goto st635
		case 117:
			goto st635
		}
		goto tr416
	st635:
		p++
	st_case_635:
		switch data[p] {
		case 84:
			goto st636
		case 116:
			goto st636
		}
		goto tr416
	st636:
		p++
	st_case_636:
		switch data[p] {
		case 69:
			goto st637
		case 101:
			goto st637
		}
		goto tr416
	st637:
		p++
	st_case_637:
		switch data[p] {
		case 9:
			goto tr855
		case 32:
			goto tr855
		case 58:
			goto tr856
		}
		goto tr416
	st638:
		p++
	st_case_638:
		switch data[p] {
		case 69:
			goto st639
		case 101:
			goto st639
		}
		goto tr416
	st639:
		p++
	st_case_639:
		switch data[p] {
		case 82:
			goto st640
		case 114:
			goto st640
		}
		goto tr416
	st640:
		p++
	st_case_640:
		switch data[p] {
		case 45:
			goto st641
		case 82:
			goto st644
		case 114:
			goto st644
		}
		goto tr416
	st641:
		p++
	st_case_641:
		switch data[p] {
		case 84:
			goto st642
		case 116:
			goto st642
		}
		goto tr416
	st642:
		p++
	st_case_642:
		switch data[p] {
		case 79:
			goto st643
		case 111:
			goto st643
		}
		goto tr416
	st643:
		p++
	st_case_643:
		switch data[p] {
		case 9:
			goto tr836
		case 32:
			goto tr836
		case 58:
			goto tr837
		}
		goto tr416
	st644:
		p++
	st_case_644:
		switch data[p] {
		case 69:
			goto st645
		case 101:
			goto st645
		}
		goto tr416
	st645:
		p++
	st_case_645:
		switch data[p] {
		case 68:
			goto st646
		case 100:
			goto st646
		}
		goto tr416
	st646:
		p++
	st_case_646:
		if data[p] == 45 {
			goto st647
		}
		goto tr416
	st647:
		p++
	st_case_647:
		switch data[p] {
		case 66:
			goto st648
		case 98:
			goto st648
		}
		goto tr416
	st648:
		p++
	st_case_648:
		switch data[p] {
		case 89:
			goto st363
		case 121:
			goto st363
		}
		goto tr416
	st649:
		p++
	st_case_649:
		switch data[p] {
		case 79:
			goto st650
		case 111:
			goto st650
		}
		goto tr416
	st650:
		p++
	st_case_650:
		switch data[p] {
		case 84:
			goto st651
		case 116:
			goto st651
		}
		goto tr416
	st651:
		p++
	st_case_651:
		switch data[p] {
		case 69:
			goto st652
		case 101:
			goto st652
		}
		goto tr416
	st652:
		p++
	st_case_652:
		if data[p] == 45 {
			goto st653
		}
		goto tr416
	st653:
		p++
	st_case_653:
		switch data[p] {
		case 80:
			goto st654
		case 112:
			goto st654
		}
		goto tr416
	st654:
		p++
	st_case_654:
		switch data[p] {
		case 65:
			goto st655
		case 97:
			goto st655
		}
		goto tr416
	st655:
		p++
	st_case_655:
		switch data[p] {
		case 82:
			goto st656
		case 114:
			goto st656
		}
		goto tr416
	st656:
		p++
	st_case_656:
		switch data[p] {
		case 84:
			goto st657
		case 116:
			goto st657
		}
		goto tr416
	st657:
		p++
	st_case_657:
		switch data[p] {
		case 89:
			goto st658
		case 121:
			goto st658
		}
		goto tr416
	st658:
		p++
	st_case_658:
		if data[p] == 45 {
			goto st659
		}
		goto tr416
	st659:
		p++
	st_case_659:
		switch data[p] {
		case 73:
			goto st660
		case 105:
			goto st660
		}
		goto tr416
	st660:
		p++
	st_case_660:
		switch data[p] {
		case 68:
			goto st661
		case 100:
			goto st661
		}
		goto tr416
	st661:
		p++
	st_case_661:
		switch data[p] {
		case 9:
			goto tr879
		case 32:
			goto tr879
		case 58:
			goto tr880
		}
		goto tr416
	st662:
		p++
	st_case_662:
		switch data[p] {
		case 76:
			goto st663
		case 108:
			goto st663
		}
		goto tr416
	st663:
		p++
	st_case_663:
		switch data[p] {
		case 89:
			goto st664
		case 121:
			goto st664
		}
		goto tr416
	st664:
		p++
	st_case_664:
		if data[p] == 45 {
			goto st665
		}
		goto tr416
	st665:
		p++
	st_case_665:
		switch data[p] {
		case 84:
			goto st666
		case 116:
			goto st666
		}
		goto tr416
	st666:
		p++
	st_case_666:
		switch data[p] {
		case 79:
			goto st667
		case 111:
			goto st667
		}
		goto tr416
	st667:
		p++
	st_case_667:
		switch data[p] {
		case 9:
			goto tr886
		case 32:
			goto tr886
		case 58:
			goto tr887
		}
		goto tr416
	st668:
		p++
	st_case_668:
		switch data[p] {
		case 85:
			goto st669
		case 117:
			goto st669
		}
		goto tr416
	st669:
		p++
	st_case_669:
		switch data[p] {
		case 73:
			goto st670
		case 105:
			goto st670
		}
		goto tr416
	st670:
		p++
	st_case_670:
		switch data[p] {
		case 82:
			goto st671
		case 114:
			goto st671
		}
		goto tr416
	st671:
		p++
	st_case_671:
		switch data[p] {
		case 69:
			goto st672
		case 101:
			goto st672
		}
		goto tr416
	st672:
		p++
	st_case_672:
		switch data[p] {
		case 9:
			goto tr892
		case 32:
			goto tr892
		case 58:
			goto tr893
		}
		goto tr416
	st673:
		p++
	st_case_673:
		switch data[p] {
		case 82:
			goto st674
		case 114:
			goto st674
		}
		goto tr416
	st674:
		p++
	st_case_674:
		switch data[p] {
		case 89:
			goto st675
		case 121:
			goto st675
		}
		goto tr416
	st675:
		p++
	st_case_675:
		if data[p] == 45 {
			goto st676
		}
		goto tr416
	st676:
		p++
	st_case_676:
		switch data[p] {
		case 65:
			goto st677
		case 97:
			goto st677
		}
		goto tr416
	st677:
		p++
	st_case_677:
		switch data[p] {
		case 70:
			goto st678
		case 102:
			goto st678
		}
		goto tr416
	st678:
		p++
	st_case_678:
		switch data[p] {
		case 84:
			goto st679
		case 116:
			goto st679
		}
		goto tr416
	st679:
		p++
	st_case_679:
		switch data[p] {
		case 69:
			goto st680
		case 101:
			goto st680
		}
		goto tr416
	st680:
		p++
	st_case_680:
		switch data[p] {
		case 82:
			goto st681
		case 114:
			goto st681
		}
		goto tr416
	st681:
		p++
	st_case_681:
		switch data[p] {
		case 9:
			goto tr902
		case 32:
			goto tr902
		case 58:
			goto tr903
		}
		goto tr416
	st682:
		p++
	st_case_682:
		switch data[p] {
		case 85:
			goto st683
		case 117:
			goto st683
		}
		goto tr416
	st683:
		p++
	st_case_683:
		switch data[p] {
		case 84:
			goto st684
		case 116:
			goto st684
		}
		goto tr416
	st684:
		p++
	st_case_684:
		switch data[p] {
		case 69:
			goto st685
		case 101:
			goto st685
		}
		goto tr416
	st685:
		p++
	st_case_685:
		switch data[p] {
		case 9:
			goto tr907
		case 32:
			goto tr907
		case 58:
			goto tr908
		}
		goto tr416
	st686:
		p++
	st_case_686:
		switch data[p] {
		case 9:
			goto tr909
		case 32:
			goto tr909
		case 58:
			goto tr910
		case 69:
			goto st687
		case 85:
			goto st692
		case 101:
			goto st687
		case 117:
			goto st692
		}
		goto tr416
	st687:
		p++
	st_case_687:
		switch data[p] {
		case 82:
			goto st688
		case 114:
			goto st688
		}
		goto tr416
	st688:
		p++
	st_case_688:
		switch data[p] {
		case 86:
			goto st689
		case 118:
			goto st689
		}
		goto tr416
	st689:
		p++
	st_case_689:
		switch data[p] {
		case 69:
			goto st690
		case 101:
			goto st690
		}
		goto tr416
	st690:
		p++
	st_case_690:
		switch data[p] {
		case 82:
			goto st691
		case 114:
			goto st691
		}
		goto tr416
	st691:
		p++
	st_case_691:
		switch data[p] {
		case 9:
			goto tr917
		case 32:
			goto tr917
		case 58:
			goto tr918
		}
		goto tr416
	st692:
		p++
	st_case_692:
		switch data[p] {
		case 66:
			goto st693
		case 80:
			goto st698
		case 98:
			goto st693
		case 112:
			goto st698
		}
		goto tr416
	st693:
		p++
	st_case_693:
		switch data[p] {
		case 74:
			goto st694
		case 106:
			goto st694
		}
		goto tr416
	st694:
		p++
	st_case_694:
		switch data[p] {
		case 69:
			goto st695
		case 101:
			goto st695
		}
		goto tr416
	st695:
		p++
	st_case_695:
		switch data[p] {
		case 67:
			goto st696
		case 99:
			goto st696
		}
		goto tr416
	st696:
		p++
	st_case_696:
		switch data[p] {
		case 84:
			goto st697
		case 116:
			goto st697
		}
		goto tr416
	st697:
		p++
	st_case_697:
		switch data[p] {
		case 9:
			goto tr909
		case 32:
			goto tr909
		case 58:
			goto tr910
		}
		goto tr416
	st698:
		p++
	st_case_698:
		switch data[p] {
		case 80:
			goto st699
		case 112:
			goto st699
		}
		goto tr416
	st699:
		p++
	st_case_699:
		switch data[p] {
		case 79:
			goto st700
		case 111:
			goto st700
		}
		goto tr416
	st700:
		p++
	st_case_700:
		switch data[p] {
		case 82:
			goto st701
		case 114:
			goto st701
		}
		goto tr416
	st701:
		p++
	st_case_701:
		switch data[p] {
		case 84:
			goto st702
		case 116:
			goto st702
		}
		goto tr416
	st702:
		p++
	st_case_702:
		switch data[p] {
		case 69:
			goto st703
		case 101:
			goto st703
		}
		goto tr416
	st703:
		p++
	st_case_703:
		switch data[p] {
		case 68:
			goto st504
		case 100:
			goto st504
		}
		goto tr416
	st704:
		p++
	st_case_704:
		switch data[p] {
		case 9:
			goto tr930
		case 32:
			goto tr930
		case 58:
			goto tr931
		case 73:
			goto st705
		case 79:
			goto st713
		case 105:
			goto st705
		case 111:
			goto st713
		}
		goto tr416
	st705:
		p++
	st_case_705:
		switch data[p] {
		case 77:
			goto st706
		case 109:
			goto st706
		}
		goto tr416
	st706:
		p++
	st_case_706:
		switch data[p] {
		case 69:
			goto st707
		case 101:
			goto st707
		}
		goto tr416
	st707:
		p++
	st_case_707:
		switch data[p] {
		case 83:
			goto st708
		case 115:
			goto st708
		}
		goto tr416
	st708:
		p++
	st_case_708:
		switch data[p] {
		case 84:
			goto st709
		case 116:
			goto st709
		}
		goto tr416
	st709:
		p++
	st_case_709:
		switch data[p] {
		case 65:
			goto st710
		case 97:
			goto st710
		}
		goto tr416
	st710:
		p++
	st_case_710:
		switch data[p] {
		case 77:
			goto st711
		case 109:
			goto st711
		}
		goto tr416
	st711:
		p++
	st_case_711:
		switch data[p] {
		case 80:
			goto st712
		case 112:
			goto st712
		}
		goto tr416
	st712:
		p++
	st_case_712:
		switch data[p] {
		case 9:
			goto tr941
		case 32:
			goto tr941
		case 58:
			goto tr942
		}
		goto tr416
	st713:
		p++
	st_case_713:
		switch data[p] {
		case 9:
			goto tr930
		case 32:
			goto tr930
		case 58:
			goto tr931
		}
		goto tr416
	st714:
		p++
	st_case_714:
		switch data[p] {
		case 9:
			goto tr943
		case 32:
			goto tr943
		case 58:
			goto tr944
		case 78:
			goto st715
		case 83:
			goto st725
		case 110:
			goto st715
		case 115:
			goto st725
		}
		goto tr416
	st715:
		p++
	st_case_715:
		switch data[p] {
		case 83:
			goto st716
		case 115:
			goto st716
		}
		goto tr416
	st716:
		p++
	st_case_716:
		switch data[p] {
		case 85:
			goto st717
		case 117:
			goto st717
		}
		goto tr416
	st717:
		p++
	st_case_717:
		switch data[p] {
		case 80:
			goto st718
		case 112:
			goto st718
		}
		goto tr416
	st718:
		p++
	st_case_718:
		switch data[p] {
		case 80:
			goto st719
		case 112:
			goto st719
		}
		goto tr416
	st719:
		p++
	st_case_719:
		switch data[p] {
		case 79:
			goto st720
		case 111:
			goto st720
		}
		goto tr416
	st720:
		p++
	st_case_720:
		switch data[p] {
		case 82:
			goto st721
		case 114:
			goto st721
		}
		goto tr416
	st721:
		p++
	st_case_721:
		switch data[p] {
		case 84:
			goto st722
		case 116:
			goto st722
		}
		goto tr416
	st722:
		p++
	st_case_722:
		switch data[p] {
		case 69:
			goto st723
		case 101:
			goto st723
		}
		goto tr416
	st723:
		p++
	st_case_723:
		switch data[p] {
		case 68:
			goto st724
		case 100:
			goto st724
		}
		goto tr416
	st724:
		p++
	st_case_724:
		switch data[p] {
		case 9:
			goto tr956
		case 32:
			goto tr956
		case 58:
			goto tr957
		}
		goto tr416
	st725:
		p++
	st_case_725:
		switch data[p] {
		case 69:
			goto st726
		case 101:
			goto st726
		}
		goto tr416
	st726:
		p++
	st_case_726:
		switch data[p] {
		case 82:
			goto st727
		case 114:
			goto st727
		}
		goto tr416
	st727:
		p++
	st_case_727:
		if data[p] == 45 {
			goto st728
		}
		goto tr416
	st728:
		p++
	st_case_728:
		switch data[p] {
		case 65:
			goto st729
		case 97:
			goto st729
		}
		goto tr416
	st729:
		p++
	st_case_729:
		switch data[p] {
		case 71:
			goto st730
		case 103:
			goto st730
		}
		goto tr416
	st730:
		p++
	st_case_730:
		switch data[p] {
		case 69:
			goto st731
		case 101:
			goto st731
		}
		goto tr416
	st731:
		p++
	st_case_731:
		switch data[p] {
		case 78:
			goto st732
		case 110:
			goto st732
		}
		goto tr416
	st732:
		p++
	st_case_732:
		switch data[p] {
		case 84:
			goto st733
		case 116:
			goto st733
		}
		goto tr416
	st733:
		p++
	st_case_733:
		switch data[p] {
		case 9:
			goto tr966
		case 32:
			goto tr966
		case 58:
			goto tr967
		}
		goto tr416
	st734:
		p++
	st_case_734:
		switch data[p] {
		case 9:
			goto st735
		case 32:
			goto st735
		case 58:
			goto st736
		case 73:
			goto st740
		case 105:
			goto st740
		}
		goto tr416
	st735:
		p++
	st_case_735:
		switch data[p] {
		case 9:
			goto st735
		case 32:
			goto st735
		case 58:
			goto st736
		}
		goto st0
	st736:
		p++
	st_case_736:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st736
		case 32:
			goto st736
		case 269:
			goto tr971
		case 525:
			goto st737
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr971
			}
		default:
			goto tr971
		}
		goto st0
	st737:
		p++
	st_case_737:
		if data[p] == 10 {
			goto st738
		}
		goto st0
	st738:
		p++
	st_case_738:
		switch data[p] {
		case 9:
			goto st739
		case 32:
			goto st739
		}
		goto st0
	st739:
		p++
	st_case_739:
		switch data[p] {
		case 9:
			goto st739
		case 32:
			goto st739
		}
		goto tr971
	st740:
		p++
	st_case_740:
		switch data[p] {
		case 65:
			goto st741
		case 97:
			goto st741
		}
		goto tr416
	st741:
		p++
	st_case_741:
		switch data[p] {
		case 9:
			goto st735
		case 32:
			goto st735
		case 58:
			goto st736
		}
		goto tr416
	st742:
		p++
	st_case_742:
		switch data[p] {
		case 65:
			goto st743
		case 87:
			goto st749
		case 97:
			goto st743
		case 119:
			goto st749
		}
		goto tr416
	st743:
		p++
	st_case_743:
		switch data[p] {
		case 82:
			goto st744
		case 114:
			goto st744
		}
		goto tr416
	st744:
		p++
	st_case_744:
		switch data[p] {
		case 78:
			goto st745
		case 110:
			goto st745
		}
		goto tr416
	st745:
		p++
	st_case_745:
		switch data[p] {
		case 73:
			goto st746
		case 105:
			goto st746
		}
		goto tr416
	st746:
		p++
	st_case_746:
		switch data[p] {
		case 78:
			goto st747
		case 110:
			goto st747
		}
		goto tr416
	st747:
		p++
	st_case_747:
		switch data[p] {
		case 71:
			goto st748
		case 103:
			goto st748
		}
		goto tr416
	st748:
		p++
	st_case_748:
		switch data[p] {
		case 9:
			goto tr983
		case 32:
			goto tr983
		case 58:
			goto tr984
		}
		goto tr416
	st749:
		p++
	st_case_749:
		switch data[p] {
		case 87:
			goto st750
		case 119:
			goto st750
		}
		goto tr416
	st750:
		p++
	st_case_750:
		if data[p] == 45 {
			goto st751
		}
		goto tr416
	st751:
		p++
	st_case_751:
		switch data[p] {
		case 65:
			goto st752
		case 97:
			goto st752
		}
		goto tr416
	st752:
		p++
	st_case_752:
		switch data[p] {
		case 85:
			goto st753
		case 117:
			goto st753
		}
		goto tr416
	st753:
		p++
	st_case_753:
		switch data[p] {
		case 84:
			goto st754
		case 116:
			goto st754
		}
		goto tr416
	st754:
		p++
	st_case_754:
		switch data[p] {
		case 72:
			goto st755
		case 104:
			goto st755
		}
		goto tr416
	st755:
		p++
	st_case_755:
		switch data[p] {
		case 69:
			goto st756
		case 101:
			goto st756
		}
		goto tr416
	st756:
		p++
	st_case_756:
		switch data[p] {
		case 78:
			goto st757
		case 110:
			goto st757
		}
		goto tr416
	st757:
		p++
	st_case_757:
		switch data[p] {
		case 84:
			goto st758
		case 116:
			goto st758
		}
		goto tr416
	st758:
		p++
	st_case_758:
		switch data[p] {
		case 73:
			goto st759
		case 105:
			goto st759
		}
		goto tr416
	st759:
		p++
	st_case_759:
		switch data[p] {
		case 67:
			goto st760
		case 99:
			goto st760
		}
		goto tr416
	st760:
		p++
	st_case_760:
		switch data[p] {
		case 65:
			goto st761
		case 97:
			goto st761
		}
		goto tr416
	st761:
		p++
	st_case_761:
		switch data[p] {
		case 84:
			goto st762
		case 116:
			goto st762
		}
		goto tr416
	st762:
		p++
	st_case_762:
		switch data[p] {
		case 69:
			goto st763
		case 101:
			goto st763
		}
		goto tr416
	st763:
		p++
	st_case_763:
		switch data[p] {
		case 9:
			goto tr999
		case 32:
			goto tr999
		case 58:
			goto tr1000
		}
		goto tr416
	st764:
		p++
	st_case_764:
		if data[p] == 10 {
			goto tr1001
		}
		goto st0
	st_out:

	if p == eof {
		switch cs {
		case 281, 282, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 370, 371, 372, 373, 374, 375, 385, 386, 387, 388, 389, 390, 391, 392, 393, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 441, 442, 443, 444, 445, 446, 447, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 503, 504, 505, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 679, 680, 681, 682, 683, 684, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 695, 696, 697, 698, 699, 700, 701, 702, 703, 704, 705, 706, 707, 708, 709, 710, 711, 712, 713, 714, 715, 716, 717, 718, 719, 720, 721, 722, 723, 724, 725, 726, 727, 728, 729, 730, 731, 732, 733, 734, 740, 741, 742, 743, 744, 745, 746, 747, 748, 749, 750, 751, 752, 753, 754, 755, 756, 757, 758, 759, 760, 761, 762, 763:
//line sip.rl:156

	p--

	{goto st273 }

		case 778:
//line sip.rl:205

	*addrp = addr
	addrp = &addr.Next
	addr = nil

//line msg_parse.go:15529
		}
	}

	_out: {}
	}

//line msg_parse.rl:49

	if cs < msg_first_final {
		if p == pe {
			return nil, MsgIncompleteError{data}
		} else {
			return nil, MsgParseError{Msg: data, Offset: p}
		}
	}

	if clen > 0 {
		if clen != len(data) - p {
			return nil, errors.New(fmt.Sprintf("Content-Length incorrect: %d != %d", clen, len(data) - p))
		}
		if ctype == sdp.ContentType {
			ms, err := sdp.Parse(string(data[p:len(data)]))
			if err != nil {
				return nil, err
			}
			msg.Payload = ms
		} else {
			msg.Payload = &MiscPayload{T: ctype, D: data[p:len(data)]}
		}
	}
	return msg, nil
}

func lookAheadWSP(data []byte, p, pe int) bool {
	return p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')
}

func lastAddr(addrp **Addr) **Addr {
	if *addrp == nil {
		return addrp
	} else {
		return lastAddr(&(*addrp).Next)
	}
}
