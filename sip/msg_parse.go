
//line msg_parse.rl:1
// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
)


//line msg_parse.rl:12

//line msg_parse.go:17
const msg_start int = 1
const msg_first_final int = 549
const msg_error int = 0

const msg_en_svalue int = 29
const msg_en_header int = 33
const msg_en_main int = 1


//line msg_parse.rl:13

// ParseMsg turns a a SIP message into a data structure.
func ParseMsg(s string) (msg *Msg, err error) {
	if s == "" {
		return nil, errors.New("Empty SIP message")
	}
	return ParseMsgBytes([]byte(s))
}

// ParseMsg turns a a SIP message byte slice into a data structure.
func ParseMsgBytes(data []byte) (msg *Msg, err error) {
	if data == nil {
		return nil, nil
	}
	msg = new(Msg)
	viap := &msg.Via
	routep := &msg.Route
	rroutep := &msg.RecordRoute
	contactp := &msg.Contact
	cs := 0
	p := 0
	pe := len(data)
	//eof := len(data)
	//stack := make([]int, 2)
	//top := 0
	line := 1
	linep := 0
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	clen := 0
	ctype := ""
	var b1 string
	var hex byte
	var dest *string

	
//line msg_parse.go:65
	{
	cs = msg_start
	}

//line msg_parse.go:70
	{
	var _widec int16
	if p == pe {
		goto _test_eof
	}
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
	case 549:
		goto st_case_549
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
	case 550:
		goto st_case_550
	case 32:
		goto st_case_32
	case 551:
		goto st_case_551
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 552:
		goto st_case_552
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
	case 553:
		goto st_case_553
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
	case 554:
		goto st_case_554
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
	case 555:
		goto st_case_555
	case 245:
		goto st_case_245
	case 246:
		goto st_case_246
	case 247:
		goto st_case_247
	case 248:
		goto st_case_248
	case 249:
		goto st_case_249
	case 250:
		goto st_case_250
	case 251:
		goto st_case_251
	case 252:
		goto st_case_252
	case 253:
		goto st_case_253
	case 254:
		goto st_case_254
	case 255:
		goto st_case_255
	case 256:
		goto st_case_256
	case 257:
		goto st_case_257
	case 258:
		goto st_case_258
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
	case 556:
		goto st_case_556
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
	case 557:
		goto st_case_557
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
	case 558:
		goto st_case_558
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
	case 559:
		goto st_case_559
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
	case 560:
		goto st_case_560
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
	case 561:
		goto st_case_561
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
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line msg_parse.rl:54

			mark = p
		
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line msg_parse.go:1254
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
//line msg_parse.rl:85

			msg.Method = string(data[mark:p])
		
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line msg_parse.go:1301
		if data[p] == 32 {
			goto st0
		}
		goto tr5
tr5:
//line msg_parse.rl:54

			mark = p
		
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line msg_parse.go:1317
		if data[p] == 32 {
			goto tr7
		}
		goto st4
tr7:
//line msg_parse.rl:97

			msg.Request, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line msg_parse.go:1334
		if data[p] == 83 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 73 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 80 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 47 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr12
		}
		goto st0
tr12:
//line msg_parse.rl:89

			msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
		
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line msg_parse.go:1386
		if data[p] == 46 {
			goto st11
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr12
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
tr14:
//line msg_parse.rl:93

			msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
		
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line msg_parse.go:1414
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
tr31:
//line msg_parse.rl:106

			msg.Phrase = string(buf[0:amt])
		
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line msg_parse.go:1433
		if data[p] == 10 {
			goto tr16
		}
		goto st0
tr16:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st549
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
//line msg_parse.go:1451
		goto st0
tr2:
//line msg_parse.rl:54

			mark = p
		
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line msg_parse.go:1464
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
		if p++; p == pe {
			goto _test_eof15
		}
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
		if p++; p == pe {
			goto _test_eof16
		}
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
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr20
		}
		goto st0
tr20:
//line msg_parse.rl:89

			msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
		
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line msg_parse.go:1601
		if data[p] == 46 {
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr20
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr22
		}
		goto st0
tr22:
//line msg_parse.rl:93

			msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
		
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line msg_parse.go:1629
		if data[p] == 32 {
			goto st21
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr22
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr24
		}
		goto st0
tr24:
//line msg_parse.rl:102

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line msg_parse.go:1657
		if 48 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
tr25:
//line msg_parse.rl:102

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line msg_parse.go:1673
		if 48 <= data[p] && data[p] <= 57 {
			goto tr26
		}
		goto st0
tr26:
//line msg_parse.rl:102

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line msg_parse.go:1689
		if data[p] == 32 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch data[p] {
		case 37:
			goto tr29
		case 60:
			goto st0
		case 62:
			goto st0
		case 96:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			switch {
			case data[p] > 8:
				if 10 <= data[p] && data[p] <= 31 {
					goto st0
				}
			default:
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] < 123:
				if 91 <= data[p] && data[p] <= 94 {
					goto st0
				}
			case data[p] > 125:
				if 254 <= data[p] {
					goto st0
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr28
tr28:
//line msg_parse.rl:58

			amt = 0
		
//line msg_parse.rl:62

			buf[amt] = data[p]
			amt++
		
	goto st26
tr30:
//line msg_parse.rl:62

			buf[amt] = data[p]
			amt++
		
	goto st26
tr34:
//line msg_parse.rl:75

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line msg_parse.go:1769
		switch data[p] {
		case 13:
			goto tr31
		case 37:
			goto st27
		case 60:
			goto st0
		case 62:
			goto st0
		case 96:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			switch {
			case data[p] > 8:
				if 10 <= data[p] && data[p] <= 31 {
					goto st0
				}
			default:
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] < 123:
				if 91 <= data[p] && data[p] <= 94 {
					goto st0
				}
			case data[p] > 125:
				if 254 <= data[p] {
					goto st0
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr30
tr29:
//line msg_parse.rl:58

			amt = 0
		
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line msg_parse.go:1822
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr33
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto st0
tr33:
//line msg_parse.rl:71

			hex = unhex(data[p]) * 16
		
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line msg_parse.go:1847
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr34
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr34
			}
		default:
			goto tr34
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr35
		case 269:
			goto tr36
		case 525:
			goto tr37
		}
		if 32 <= _widec && _widec <= 253 {
			goto tr35
		}
		goto st0
tr35:
//line msg_parse.rl:54

			mark = p
		
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line msg_parse.go:1896
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st30
		case 269:
			goto st31
		case 525:
			goto st32
		}
		if 32 <= _widec && _widec <= 253 {
			goto st30
		}
		goto st0
tr36:
//line msg_parse.rl:54

			mark = p
		
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line msg_parse.go:1927
		if data[p] == 10 {
			goto tr41
		}
		goto st0
tr41:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:123

			if dest != nil {
				*dest = string(data[mark:p - 1])
			} else {
				if msg.Headers == nil {
					msg.Headers = Headers{}
				}
				msg.Headers[b1] = string(data[mark:p])
			}
		
//line msg_parse.rl:110

			{goto st33 }
		
	goto st550
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
//line msg_parse.go:1956
		goto st0
tr37:
//line msg_parse.rl:54

			mark = p
		
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line msg_parse.go:1969
		if data[p] == 10 {
			goto tr42
		}
		goto st0
tr42:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:123

			if dest != nil {
				*dest = string(data[mark:p - 1])
			} else {
				if msg.Headers == nil {
					msg.Headers = Headers{}
				}
				msg.Headers[b1] = string(data[mark:p])
			}
		
//line msg_parse.rl:110

			{goto st33 }
		
	goto st551
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
//line msg_parse.go:1998
		switch data[p] {
		case 9:
			goto st30
		case 32:
			goto st30
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 13:
			goto st34
		case 65:
			goto st35
		case 66:
			goto st116
		case 67:
			goto st117
		case 68:
			goto st211
		case 69:
			goto st215
		case 70:
			goto st240
		case 73:
			goto st250
		case 75:
			goto st261
		case 76:
			goto st262
		case 77:
			goto st268
		case 79:
			goto st310
		case 80:
			goto st322
		case 82:
			goto st386
		case 83:
			goto st464
		case 84:
			goto st482
		case 85:
			goto st498
		case 86:
			goto st518
		case 87:
			goto st527
		case 97:
			goto st35
		case 98:
			goto st116
		case 99:
			goto st117
		case 100:
			goto st211
		case 101:
			goto st215
		case 102:
			goto st240
		case 105:
			goto st250
		case 107:
			goto st261
		case 108:
			goto st262
		case 109:
			goto st268
		case 111:
			goto st310
		case 112:
			goto st322
		case 114:
			goto st386
		case 115:
			goto st464
		case 116:
			goto st482
		case 117:
			goto st498
		case 118:
			goto st518
		case 119:
			goto st527
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 10 {
			goto tr62
		}
		goto st0
tr175:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st552
tr62:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:50

			{p++; cs = 552; goto _out }
		
	goto st552
tr70:
//line msg_parse.rl:118

			p--

			{goto st29 }
		
	goto st552
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
//line msg_parse.go:2126
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 9:
			goto tr63
		case 32:
			goto tr63
		case 58:
			goto tr64
		case 67:
			goto st41
		case 76:
			goto st70
		case 85:
			goto st89
		case 99:
			goto st41
		case 108:
			goto st70
		case 117:
			goto st89
		}
		goto st0
tr63:
//line msg_parse.rl:290
dest=&msg.AcceptContact
	goto st36
tr78:
//line msg_parse.rl:289
dest=&msg.Accept
	goto st36
tr97:
//line msg_parse.rl:291
dest=&msg.AcceptEncoding
	goto st36
tr106:
//line msg_parse.rl:292
dest=&msg.AcceptLanguage
	goto st36
tr117:
//line msg_parse.rl:295
dest=&msg.AlertInfo
	goto st36
tr121:
//line msg_parse.rl:293
dest=&msg.Allow
	goto st36
tr130:
//line msg_parse.rl:294
dest=&msg.AllowEvents
	goto st36
tr150:
//line msg_parse.rl:296
dest=&msg.AuthenticationInfo
	goto st36
tr160:
//line msg_parse.rl:297
dest=&msg.Authorization
	goto st36
tr162:
//line msg_parse.rl:314
dest=&msg.ReferredBy
	goto st36
tr197:
//line msg_parse.rl:301
dest=&msg.CallInfo
	goto st36
tr233:
//line msg_parse.rl:298
dest=&msg.ContentDisposition
	goto st36
tr242:
//line msg_parse.rl:300
dest=&msg.ContentEncoding
	goto st36
tr252:
//line msg_parse.rl:299
dest=&msg.ContentLanguage
	goto st36
tr284:
//line msg_parse.rl:302
dest=&msg.Date
	goto st36
tr297:
//line msg_parse.rl:303
dest=&msg.ErrorInfo
	goto st36
tr302:
//line msg_parse.rl:304
dest=&msg.Event
	goto st36
tr339:
//line msg_parse.rl:305
dest=&msg.InReplyTo
	goto st36
tr341:
//line msg_parse.rl:319
dest=&msg.Supported
	goto st36
tr378:
//line msg_parse.rl:307
dest=&msg.MIMEVersion
	goto st36
tr405:
//line msg_parse.rl:308
dest=&msg.Organization
	goto st36
tr444:
//line msg_parse.rl:309
dest=&msg.Priority
	goto st36
tr463:
//line msg_parse.rl:310
dest=&msg.ProxyAuthenticate
	goto st36
tr473:
//line msg_parse.rl:311
dest=&msg.ProxyAuthorization
	goto st36
tr481:
//line msg_parse.rl:312
dest=&msg.ProxyRequire
	goto st36
tr483:
//line msg_parse.rl:313
dest=&msg.ReferTo
	goto st36
tr551:
//line msg_parse.rl:306
dest=&msg.ReplyTo
	goto st36
tr557:
//line msg_parse.rl:315
dest=&msg.Require
	goto st36
tr567:
//line msg_parse.rl:316
dest=&msg.RetryAfter
	goto st36
tr583:
//line msg_parse.rl:318
dest=&msg.Subject
	goto st36
tr591:
//line msg_parse.rl:317
dest=&msg.Server
	goto st36
tr624:
//line msg_parse.rl:320
dest=&msg.Timestamp
	goto st36
tr626:
//line msg_parse.rl:293
dest=&msg.Allow
//line msg_parse.rl:294
dest=&msg.AllowEvents
	goto st36
tr639:
//line msg_parse.rl:321
dest=&msg.Unsupported
	goto st36
tr649:
//line msg_parse.rl:322
dest=&msg.UserAgent
	goto st36
tr671:
//line msg_parse.rl:323
dest=&msg.Warning
	goto st36
tr687:
//line msg_parse.rl:324
dest=&msg.WWWAuthenticate
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line msg_parse.go:2309
		switch data[p] {
		case 9:
			goto st36
		case 32:
			goto st36
		case 58:
			goto st37
		}
		goto st0
tr64:
//line msg_parse.rl:290
dest=&msg.AcceptContact
	goto st37
tr80:
//line msg_parse.rl:289
dest=&msg.Accept
	goto st37
tr98:
//line msg_parse.rl:291
dest=&msg.AcceptEncoding
	goto st37
tr107:
//line msg_parse.rl:292
dest=&msg.AcceptLanguage
	goto st37
tr118:
//line msg_parse.rl:295
dest=&msg.AlertInfo
	goto st37
tr123:
//line msg_parse.rl:293
dest=&msg.Allow
	goto st37
tr131:
//line msg_parse.rl:294
dest=&msg.AllowEvents
	goto st37
tr151:
//line msg_parse.rl:296
dest=&msg.AuthenticationInfo
	goto st37
tr161:
//line msg_parse.rl:297
dest=&msg.Authorization
	goto st37
tr163:
//line msg_parse.rl:314
dest=&msg.ReferredBy
	goto st37
tr198:
//line msg_parse.rl:301
dest=&msg.CallInfo
	goto st37
tr234:
//line msg_parse.rl:298
dest=&msg.ContentDisposition
	goto st37
tr243:
//line msg_parse.rl:300
dest=&msg.ContentEncoding
	goto st37
tr253:
//line msg_parse.rl:299
dest=&msg.ContentLanguage
	goto st37
tr285:
//line msg_parse.rl:302
dest=&msg.Date
	goto st37
tr298:
//line msg_parse.rl:303
dest=&msg.ErrorInfo
	goto st37
tr303:
//line msg_parse.rl:304
dest=&msg.Event
	goto st37
tr340:
//line msg_parse.rl:305
dest=&msg.InReplyTo
	goto st37
tr342:
//line msg_parse.rl:319
dest=&msg.Supported
	goto st37
tr379:
//line msg_parse.rl:307
dest=&msg.MIMEVersion
	goto st37
tr406:
//line msg_parse.rl:308
dest=&msg.Organization
	goto st37
tr445:
//line msg_parse.rl:309
dest=&msg.Priority
	goto st37
tr464:
//line msg_parse.rl:310
dest=&msg.ProxyAuthenticate
	goto st37
tr474:
//line msg_parse.rl:311
dest=&msg.ProxyAuthorization
	goto st37
tr482:
//line msg_parse.rl:312
dest=&msg.ProxyRequire
	goto st37
tr484:
//line msg_parse.rl:313
dest=&msg.ReferTo
	goto st37
tr552:
//line msg_parse.rl:306
dest=&msg.ReplyTo
	goto st37
tr558:
//line msg_parse.rl:315
dest=&msg.Require
	goto st37
tr568:
//line msg_parse.rl:316
dest=&msg.RetryAfter
	goto st37
tr584:
//line msg_parse.rl:318
dest=&msg.Subject
	goto st37
tr592:
//line msg_parse.rl:317
dest=&msg.Server
	goto st37
tr625:
//line msg_parse.rl:320
dest=&msg.Timestamp
	goto st37
tr627:
//line msg_parse.rl:293
dest=&msg.Allow
//line msg_parse.rl:294
dest=&msg.AllowEvents
	goto st37
tr640:
//line msg_parse.rl:321
dest=&msg.Unsupported
	goto st37
tr650:
//line msg_parse.rl:322
dest=&msg.UserAgent
	goto st37
tr672:
//line msg_parse.rl:323
dest=&msg.Warning
	goto st37
tr688:
//line msg_parse.rl:324
dest=&msg.WWWAuthenticate
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line msg_parse.go:2474
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st37
		case 32:
			goto st37
		case 269:
			goto tr70
		case 525:
			goto st38
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr70
			}
		default:
			goto tr70
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 10 {
			goto tr72
		}
		goto st0
tr72:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line msg_parse.go:2519
		switch data[p] {
		case 9:
			goto st40
		case 32:
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 9:
			goto st40
		case 32:
			goto st40
		}
		goto tr70
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 67:
			goto st42
		case 99:
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 69:
			goto st43
		case 101:
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 80:
			goto st44
		case 112:
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 84:
			goto st45
		case 116:
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		switch data[p] {
		case 9:
			goto tr78
		case 32:
			goto tr78
		case 45:
			goto st46
		case 58:
			goto tr80
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 67:
			goto st47
		case 69:
			goto st54
		case 76:
			goto st62
		case 99:
			goto st47
		case 101:
			goto st54
		case 108:
			goto st62
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 79:
			goto st48
		case 111:
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 78:
			goto st49
		case 110:
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 84:
			goto st50
		case 116:
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 65:
			goto st51
		case 97:
			goto st51
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 67:
			goto st52
		case 99:
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 84:
			goto st53
		case 116:
			goto st53
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 9:
			goto tr63
		case 32:
			goto tr63
		case 58:
			goto tr64
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 78:
			goto st55
		case 110:
			goto st55
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		switch data[p] {
		case 67:
			goto st56
		case 99:
			goto st56
		}
		goto st0
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
		switch data[p] {
		case 79:
			goto st57
		case 111:
			goto st57
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 68:
			goto st58
		case 100:
			goto st58
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 73:
			goto st59
		case 105:
			goto st59
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		switch data[p] {
		case 78:
			goto st60
		case 110:
			goto st60
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
		switch data[p] {
		case 71:
			goto st61
		case 103:
			goto st61
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 9:
			goto tr97
		case 32:
			goto tr97
		case 58:
			goto tr98
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 65:
			goto st63
		case 97:
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 78:
			goto st64
		case 110:
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 71:
			goto st65
		case 103:
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 85:
			goto st66
		case 117:
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 65:
			goto st67
		case 97:
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 71:
			goto st68
		case 103:
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 69:
			goto st69
		case 101:
			goto st69
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 9:
			goto tr106
		case 32:
			goto tr106
		case 58:
			goto tr107
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 69:
			goto st71
		case 76:
			goto st79
		case 101:
			goto st71
		case 108:
			goto st79
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 82:
			goto st72
		case 114:
			goto st72
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 84:
			goto st73
		case 116:
			goto st73
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		if data[p] == 45 {
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 73:
			goto st75
		case 105:
			goto st75
		}
		goto st0
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
		switch data[p] {
		case 78:
			goto st76
		case 110:
			goto st76
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 70:
			goto st77
		case 102:
			goto st77
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 79:
			goto st78
		case 111:
			goto st78
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 9:
			goto tr117
		case 32:
			goto tr117
		case 58:
			goto tr118
		}
		goto st0
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
		switch data[p] {
		case 79:
			goto st80
		case 111:
			goto st80
		}
		goto st0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		switch data[p] {
		case 87:
			goto st81
		case 119:
			goto st81
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 9:
			goto tr121
		case 32:
			goto tr121
		case 45:
			goto st82
		case 58:
			goto tr123
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		switch data[p] {
		case 69:
			goto st83
		case 101:
			goto st83
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 86:
			goto st84
		case 118:
			goto st84
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 69:
			goto st85
		case 101:
			goto st85
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 78:
			goto st86
		case 110:
			goto st86
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 84:
			goto st87
		case 116:
			goto st87
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 83:
			goto st88
		case 115:
			goto st88
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 9:
			goto tr130
		case 32:
			goto tr130
		case 58:
			goto tr131
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 84:
			goto st90
		case 116:
			goto st90
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 72:
			goto st91
		case 104:
			goto st91
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		switch data[p] {
		case 69:
			goto st92
		case 79:
			goto st107
		case 101:
			goto st92
		case 111:
			goto st107
		}
		goto st0
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
		switch data[p] {
		case 78:
			goto st93
		case 110:
			goto st93
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 84:
			goto st94
		case 116:
			goto st94
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 73:
			goto st95
		case 105:
			goto st95
		}
		goto st0
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
		switch data[p] {
		case 67:
			goto st96
		case 99:
			goto st96
		}
		goto st0
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
		switch data[p] {
		case 65:
			goto st97
		case 97:
			goto st97
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		switch data[p] {
		case 84:
			goto st98
		case 116:
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 73:
			goto st99
		case 105:
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 79:
			goto st100
		case 111:
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 78:
			goto st101
		case 110:
			goto st101
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		if data[p] == 45 {
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 73:
			goto st103
		case 105:
			goto st103
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 78:
			goto st104
		case 110:
			goto st104
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 70:
			goto st105
		case 102:
			goto st105
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 79:
			goto st106
		case 111:
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 9:
			goto tr150
		case 32:
			goto tr150
		case 58:
			goto tr151
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 82:
			goto st108
		case 114:
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 73:
			goto st109
		case 105:
			goto st109
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 90:
			goto st110
		case 122:
			goto st110
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 65:
			goto st111
		case 97:
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 84:
			goto st112
		case 116:
			goto st112
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 73:
			goto st113
		case 105:
			goto st113
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 79:
			goto st114
		case 111:
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 78:
			goto st115
		case 110:
			goto st115
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 9:
			goto tr160
		case 32:
			goto tr160
		case 58:
			goto tr161
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 9:
			goto tr162
		case 32:
			goto tr162
		case 58:
			goto tr163
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 9:
			goto st118
		case 32:
			goto st118
		case 58:
			goto st119
		case 65:
			goto st126
		case 79:
			goto st142
		case 83:
			goto st198
		case 97:
			goto st126
		case 111:
			goto st142
		case 115:
			goto st198
		}
		goto st0
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
		switch data[p] {
		case 9:
			goto st118
		case 32:
			goto st118
		case 58:
			goto st119
		}
		goto st0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st119
		case 32:
			goto st119
		case 269:
			goto tr170
		case 525:
			goto st123
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr169
		}
		goto st0
tr169:
//line msg_parse.rl:54

			mark = p
		
	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line msg_parse.go:3562
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st120
		case 269:
			goto tr173
		case 525:
			goto tr174
		}
		if 32 <= _widec && _widec <= 253 {
			goto st120
		}
		goto st0
tr170:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:148

			ctype = string(data[mark:p])
		
	goto st121
tr173:
//line msg_parse.rl:148

			ctype = string(data[mark:p])
		
	goto st121
tr189:
//line msg_parse.rl:134

			msg.CallID = string(data[mark:p])
		
	goto st121
tr207:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:138

			*contactp, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *contactp != nil { contactp = &(*contactp).Next }
		
	goto st121
tr210:
//line msg_parse.rl:138

			*contactp, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *contactp != nil { contactp = &(*contactp).Next }
		
	goto st121
tr275:
//line msg_parse.rl:156

			msg.CSeqMethod = string(data[mark:p])
		
	goto st121
tr319:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:164

			msg.From, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr322:
//line msg_parse.rl:164

			msg.From, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr428:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:177

			msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr431:
//line msg_parse.rl:177

			msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr504:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:182

			*rroutep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *rroutep != nil { rroutep = &(*rroutep).Next }
		
	goto st121
tr507:
//line msg_parse.rl:182

			*rroutep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *rroutep != nil { rroutep = &(*rroutep).Next }
		
	goto st121
tr537:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:188

			msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr540:
//line msg_parse.rl:188

			msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr574:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:193

			*routep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *routep != nil { routep = &(*routep).Next }
		
	goto st121
tr577:
//line msg_parse.rl:193

			*routep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *routep != nil { routep = &(*routep).Next }
		
	goto st121
tr608:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:199

			msg.To, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr611:
//line msg_parse.rl:199

			msg.To, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st121
tr655:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:204

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st121
tr658:
//line msg_parse.rl:204

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line msg_parse.go:3767
		if data[p] == 10 {
			goto tr175
		}
		goto st0
tr179:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:148

			ctype = string(data[mark:p])
		
	goto st122
tr174:
//line msg_parse.rl:148

			ctype = string(data[mark:p])
		
	goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line msg_parse.go:3793
		if data[p] == 10 {
			goto tr176
		}
		goto st0
tr176:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st553
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
//line msg_parse.go:3811
		switch data[p] {
		case 9:
			goto st120
		case 32:
			goto st120
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		if data[p] == 10 {
			goto tr177
		}
		goto st0
tr177:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line msg_parse.go:3837
		switch data[p] {
		case 9:
			goto st125
		case 32:
			goto st125
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st125
		case 32:
			goto st125
		case 269:
			goto tr170
		case 525:
			goto tr179
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr169
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 76:
			goto st127
		case 108:
			goto st127
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
		switch data[p] {
		case 76:
			goto st128
		case 108:
			goto st128
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		if data[p] == 45 {
			goto st129
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 73:
			goto st130
		case 105:
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 68:
			goto st131
		case 78:
			goto st139
		case 100:
			goto st131
		case 110:
			goto st139
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 9:
			goto st131
		case 32:
			goto st131
		case 58:
			goto st132
		}
		goto st0
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st132
		case 32:
			goto st132
		case 37:
			goto tr187
		case 60:
			goto tr187
		case 525:
			goto st136
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto tr187
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto tr187
				}
			default:
				goto tr187
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto tr187
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto tr187
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st0
tr187:
//line msg_parse.rl:54

			mark = p
		
	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line msg_parse.go:4012
		switch data[p] {
		case 13:
			goto tr189
		case 37:
			goto st133
		case 60:
			goto st133
		case 64:
			goto st134
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 34:
				if 39 <= data[p] && data[p] <= 43 {
					goto st133
				}
			case data[p] >= 33:
				goto st133
			}
		case data[p] > 58:
			switch {
			case data[p] < 95:
				if 62 <= data[p] && data[p] <= 93 {
					goto st133
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st133
				}
			default:
				goto st133
			}
		default:
			goto st133
		}
		goto st0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 37:
			goto st135
		case 60:
			goto st135
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st135
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st135
				}
			default:
				goto st135
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st135
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st135
				}
			default:
				goto st135
			}
		default:
			goto st135
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 13:
			goto tr189
		case 37:
			goto st135
		case 60:
			goto st135
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st135
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st135
				}
			default:
				goto st135
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st135
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st135
				}
			default:
				goto st135
			}
		default:
			goto st135
		}
		goto st0
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
		if data[p] == 10 {
			goto tr193
		}
		goto st0
tr193:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line msg_parse.go:4154
		switch data[p] {
		case 9:
			goto st138
		case 32:
			goto st138
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 9:
			goto st138
		case 32:
			goto st138
		case 37:
			goto tr187
		case 60:
			goto tr187
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto tr187
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto tr187
				}
			default:
				goto tr187
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto tr187
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto tr187
				}
			default:
				goto tr187
			}
		default:
			goto tr187
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 70:
			goto st140
		case 102:
			goto st140
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 79:
			goto st141
		case 111:
			goto st141
		}
		goto st0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 9:
			goto tr197
		case 32:
			goto tr197
		case 58:
			goto tr198
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 78:
			goto st143
		case 110:
			goto st143
		}
		goto st0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
		switch data[p] {
		case 84:
			goto st144
		case 116:
			goto st144
		}
		goto st0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 65:
			goto st145
		case 69:
			goto st154
		case 97:
			goto st145
		case 101:
			goto st154
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 67:
			goto st146
		case 99:
			goto st146
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 84:
			goto st147
		case 116:
			goto st147
		}
		goto st0
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		switch data[p] {
		case 9:
			goto st147
		case 32:
			goto st147
		case 58:
			goto st148
		}
		goto st0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st148
		case 32:
			goto st148
		case 269:
			goto tr207
		case 525:
			goto st151
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr206
		}
		goto st0
tr206:
//line msg_parse.rl:54

			mark = p
		
	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line msg_parse.go:4361
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st149
		case 269:
			goto tr210
		case 525:
			goto tr211
		}
		if 32 <= _widec && _widec <= 253 {
			goto st149
		}
		goto st0
tr215:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:138

			*contactp, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *contactp != nil { contactp = &(*contactp).Next }
		
	goto st150
tr211:
//line msg_parse.rl:138

			*contactp, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *contactp != nil { contactp = &(*contactp).Next }
		
	goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line msg_parse.go:4406
		if data[p] == 10 {
			goto tr212
		}
		goto st0
tr212:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st554
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
//line msg_parse.go:4424
		switch data[p] {
		case 9:
			goto st149
		case 32:
			goto st149
		}
		goto st0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
		if data[p] == 10 {
			goto tr213
		}
		goto st0
tr213:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line msg_parse.go:4450
		switch data[p] {
		case 9:
			goto st153
		case 32:
			goto st153
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st153
		case 32:
			goto st153
		case 269:
			goto tr207
		case 525:
			goto tr215
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr206
		}
		goto st0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 78:
			goto st155
		case 110:
			goto st155
		}
		goto st0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 84:
			goto st156
		case 116:
			goto st156
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		if data[p] == 45 {
			goto st157
		}
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 68:
			goto st158
		case 69:
			goto st169
		case 76:
			goto st177
		case 84:
			goto st195
		case 100:
			goto st158
		case 101:
			goto st169
		case 108:
			goto st177
		case 116:
			goto st195
		}
		goto st0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 73:
			goto st159
		case 105:
			goto st159
		}
		goto st0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 83:
			goto st160
		case 115:
			goto st160
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 80:
			goto st161
		case 112:
			goto st161
		}
		goto st0
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 79:
			goto st162
		case 111:
			goto st162
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 83:
			goto st163
		case 115:
			goto st163
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 73:
			goto st164
		case 105:
			goto st164
		}
		goto st0
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
		switch data[p] {
		case 84:
			goto st165
		case 116:
			goto st165
		}
		goto st0
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 73:
			goto st166
		case 105:
			goto st166
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 79:
			goto st167
		case 111:
			goto st167
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 78:
			goto st168
		case 110:
			goto st168
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 9:
			goto tr233
		case 32:
			goto tr233
		case 58:
			goto tr234
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 78:
			goto st170
		case 110:
			goto st170
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 67:
			goto st171
		case 99:
			goto st171
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 79:
			goto st172
		case 111:
			goto st172
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 68:
			goto st173
		case 100:
			goto st173
		}
		goto st0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 73:
			goto st174
		case 105:
			goto st174
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 78:
			goto st175
		case 110:
			goto st175
		}
		goto st0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 71:
			goto st176
		case 103:
			goto st176
		}
		goto st0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 9:
			goto tr242
		case 32:
			goto tr242
		case 58:
			goto tr243
		}
		goto st0
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 65:
			goto st178
		case 69:
			goto st185
		case 97:
			goto st178
		case 101:
			goto st185
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 78:
			goto st179
		case 110:
			goto st179
		}
		goto st0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 71:
			goto st180
		case 103:
			goto st180
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 85:
			goto st181
		case 117:
			goto st181
		}
		goto st0
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 65:
			goto st182
		case 97:
			goto st182
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 71:
			goto st183
		case 103:
			goto st183
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 69:
			goto st184
		case 101:
			goto st184
		}
		goto st0
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 9:
			goto tr252
		case 32:
			goto tr252
		case 58:
			goto tr253
		}
		goto st0
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		switch data[p] {
		case 78:
			goto st186
		case 110:
			goto st186
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 71:
			goto st187
		case 103:
			goto st187
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 84:
			goto st188
		case 116:
			goto st188
		}
		goto st0
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 72:
			goto st189
		case 104:
			goto st189
		}
		goto st0
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 9:
			goto st189
		case 32:
			goto st189
		case 58:
			goto st190
		}
		goto st0
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st190
		case 32:
			goto st190
		case 525:
			goto st192
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr259
		}
		goto st0
tr259:
//line msg_parse.rl:274
clen=0
//line msg_parse.rl:144

			clen = clen * 10 + (int(data[p]) - 0x30)
		
	goto st191
tr262:
//line msg_parse.rl:144

			clen = clen * 10 + (int(data[p]) - 0x30)
		
	goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line msg_parse.go:4980
		if data[p] == 13 {
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr262
		}
		goto st0
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		if data[p] == 10 {
			goto tr263
		}
		goto st0
tr263:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line msg_parse.go:5006
		switch data[p] {
		case 9:
			goto st194
		case 32:
			goto st194
		}
		goto st0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 9:
			goto st194
		case 32:
			goto st194
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr259
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 89:
			goto st196
		case 121:
			goto st196
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 80:
			goto st197
		case 112:
			goto st197
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 69:
			goto st118
		case 101:
			goto st118
		}
		goto st0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 69:
			goto st199
		case 101:
			goto st199
		}
		goto st0
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 81:
			goto st200
		case 113:
			goto st200
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 9:
			goto st200
		case 32:
			goto st200
		case 58:
			goto st201
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st201
		case 32:
			goto st201
		case 525:
			goto st208
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr270
		}
		goto st0
tr270:
//line msg_parse.rl:152

			msg.CSeq = msg.CSeq * 10 + (int(data[p]) - 0x30)
		
	goto st202
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
//line msg_parse.go:5138
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st203
		case 32:
			goto st203
		case 525:
			goto st205
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr270
		}
		goto st0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st203
		case 32:
			goto st203
		case 33:
			goto tr274
		case 37:
			goto tr274
		case 39:
			goto tr274
		case 126:
			goto tr274
		case 525:
			goto st205
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr274
				}
			case _widec >= 42:
				goto tr274
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr274
				}
			case _widec >= 65:
				goto tr274
			}
		default:
			goto tr274
		}
		goto st0
tr274:
//line msg_parse.rl:54

			mark = p
		
	goto st204
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
//line msg_parse.go:5220
		switch data[p] {
		case 13:
			goto tr275
		case 33:
			goto st204
		case 37:
			goto st204
		case 39:
			goto st204
		case 126:
			goto st204
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st204
				}
			case data[p] >= 42:
				goto st204
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st204
				}
			case data[p] >= 65:
				goto st204
			}
		default:
			goto st204
		}
		goto st0
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		if data[p] == 10 {
			goto tr277
		}
		goto st0
tr277:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line msg_parse.go:5274
		switch data[p] {
		case 9:
			goto st207
		case 32:
			goto st207
		}
		goto st0
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch data[p] {
		case 9:
			goto st207
		case 32:
			goto st207
		case 33:
			goto tr274
		case 37:
			goto tr274
		case 39:
			goto tr274
		case 126:
			goto tr274
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr274
				}
			case data[p] >= 42:
				goto tr274
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr274
				}
			case data[p] >= 65:
				goto tr274
			}
		default:
			goto tr274
		}
		goto st0
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		if data[p] == 10 {
			goto tr279
		}
		goto st0
tr279:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line msg_parse.go:5342
		switch data[p] {
		case 9:
			goto st210
		case 32:
			goto st210
		}
		goto st0
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		switch data[p] {
		case 9:
			goto st210
		case 32:
			goto st210
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr270
		}
		goto st0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 65:
			goto st212
		case 97:
			goto st212
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 84:
			goto st213
		case 116:
			goto st213
		}
		goto st0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		switch data[p] {
		case 69:
			goto st214
		case 101:
			goto st214
		}
		goto st0
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 9:
			goto tr284
		case 32:
			goto tr284
		case 58:
			goto tr285
		}
		goto st0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 9:
			goto tr242
		case 32:
			goto tr242
		case 58:
			goto tr243
		case 82:
			goto st216
		case 86:
			goto st225
		case 88:
			goto st229
		case 114:
			goto st216
		case 118:
			goto st225
		case 120:
			goto st229
		}
		goto st0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 82:
			goto st217
		case 114:
			goto st217
		}
		goto st0
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch data[p] {
		case 79:
			goto st218
		case 111:
			goto st218
		}
		goto st0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 82:
			goto st219
		case 114:
			goto st219
		}
		goto st0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		if data[p] == 45 {
			goto st220
		}
		goto st0
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 73:
			goto st221
		case 105:
			goto st221
		}
		goto st0
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 78:
			goto st222
		case 110:
			goto st222
		}
		goto st0
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 70:
			goto st223
		case 102:
			goto st223
		}
		goto st0
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 79:
			goto st224
		case 111:
			goto st224
		}
		goto st0
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 9:
			goto tr297
		case 32:
			goto tr297
		case 58:
			goto tr298
		}
		goto st0
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 69:
			goto st226
		case 101:
			goto st226
		}
		goto st0
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
		switch data[p] {
		case 78:
			goto st227
		case 110:
			goto st227
		}
		goto st0
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 84:
			goto st228
		case 116:
			goto st228
		}
		goto st0
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
		switch data[p] {
		case 9:
			goto tr302
		case 32:
			goto tr302
		case 58:
			goto tr303
		}
		goto st0
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
		switch data[p] {
		case 80:
			goto st230
		case 112:
			goto st230
		}
		goto st0
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 73:
			goto st231
		case 105:
			goto st231
		}
		goto st0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 82:
			goto st232
		case 114:
			goto st232
		}
		goto st0
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		switch data[p] {
		case 69:
			goto st233
		case 101:
			goto st233
		}
		goto st0
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
		switch data[p] {
		case 83:
			goto st234
		case 115:
			goto st234
		}
		goto st0
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		switch data[p] {
		case 9:
			goto st234
		case 32:
			goto st234
		case 58:
			goto st235
		}
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st235
		case 32:
			goto st235
		case 525:
			goto st237
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr310
		}
		goto st0
tr310:
//line msg_parse.rl:277
msg.Expires=0
//line msg_parse.rl:160

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
	goto st236
tr312:
//line msg_parse.rl:160

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
	goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line msg_parse.go:5715
		if data[p] == 13 {
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr312
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		if data[p] == 10 {
			goto tr313
		}
		goto st0
tr313:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line msg_parse.go:5741
		switch data[p] {
		case 9:
			goto st239
		case 32:
			goto st239
		}
		goto st0
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		switch data[p] {
		case 9:
			goto st239
		case 32:
			goto st239
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr310
		}
		goto st0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		switch data[p] {
		case 9:
			goto st241
		case 32:
			goto st241
		case 58:
			goto st242
		case 82:
			goto st248
		case 114:
			goto st248
		}
		goto st0
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
		switch data[p] {
		case 9:
			goto st241
		case 32:
			goto st241
		case 58:
			goto st242
		}
		goto st0
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st242
		case 32:
			goto st242
		case 269:
			goto tr319
		case 525:
			goto st245
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr318
		}
		goto st0
tr318:
//line msg_parse.rl:54

			mark = p
		
	goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line msg_parse.go:5833
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st243
		case 269:
			goto tr322
		case 525:
			goto tr323
		}
		if 32 <= _widec && _widec <= 253 {
			goto st243
		}
		goto st0
tr327:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:164

			msg.From, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st244
tr323:
//line msg_parse.rl:164

			msg.From, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st244
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
//line msg_parse.go:5876
		if data[p] == 10 {
			goto tr324
		}
		goto st0
tr324:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st555
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
//line msg_parse.go:5894
		switch data[p] {
		case 9:
			goto st243
		case 32:
			goto st243
		}
		goto st0
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		if data[p] == 10 {
			goto tr325
		}
		goto st0
tr325:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
//line msg_parse.go:5920
		switch data[p] {
		case 9:
			goto st247
		case 32:
			goto st247
		}
		goto st0
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st247
		case 32:
			goto st247
		case 269:
			goto tr319
		case 525:
			goto tr327
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr318
		}
		goto st0
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 79:
			goto st249
		case 111:
			goto st249
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		switch data[p] {
		case 77:
			goto st241
		case 109:
			goto st241
		}
		goto st0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 9:
			goto st131
		case 32:
			goto st131
		case 58:
			goto st132
		case 78:
			goto st251
		case 110:
			goto st251
		}
		goto st0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		if data[p] == 45 {
			goto st252
		}
		goto st0
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 82:
			goto st253
		case 114:
			goto st253
		}
		goto st0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		switch data[p] {
		case 69:
			goto st254
		case 101:
			goto st254
		}
		goto st0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
		switch data[p] {
		case 80:
			goto st255
		case 112:
			goto st255
		}
		goto st0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 76:
			goto st256
		case 108:
			goto st256
		}
		goto st0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 89:
			goto st257
		case 121:
			goto st257
		}
		goto st0
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		if data[p] == 45 {
			goto st258
		}
		goto st0
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 84:
			goto st259
		case 116:
			goto st259
		}
		goto st0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 79:
			goto st260
		case 111:
			goto st260
		}
		goto st0
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 9:
			goto tr339
		case 32:
			goto tr339
		case 58:
			goto tr340
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 9:
			goto tr341
		case 32:
			goto tr341
		case 58:
			goto tr342
		}
		goto st0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 9:
			goto st262
		case 32:
			goto st262
		case 58:
			goto st263
		}
		goto st0
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st263
		case 32:
			goto st263
		case 525:
			goto st265
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr344
		}
		goto st0
tr344:
//line msg_parse.rl:274
clen=0
//line msg_parse.rl:144

			clen = clen * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:277
msg.Expires=0
//line msg_parse.rl:160

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:279
msg.MaxForwards=0
//line msg_parse.rl:169

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:280
msg.MinExpires=0
//line msg_parse.rl:173

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st264
tr346:
//line msg_parse.rl:144

			clen = clen * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:160

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:169

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:173

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st264
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
//line msg_parse.go:6213
		if data[p] == 13 {
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr346
		}
		goto st0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		if data[p] == 10 {
			goto tr347
		}
		goto st0
tr347:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line msg_parse.go:6239
		switch data[p] {
		case 9:
			goto st267
		case 32:
			goto st267
		}
		goto st0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		switch data[p] {
		case 9:
			goto st267
		case 32:
			goto st267
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr344
		}
		goto st0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		switch data[p] {
		case 9:
			goto st147
		case 32:
			goto st147
		case 58:
			goto st148
		case 65:
			goto st269
		case 73:
			goto st285
		case 97:
			goto st269
		case 105:
			goto st285
		}
		goto st0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		switch data[p] {
		case 88:
			goto st270
		case 120:
			goto st270
		}
		goto st0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		if data[p] == 45 {
			goto st271
		}
		goto st0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 70:
			goto st272
		case 102:
			goto st272
		}
		goto st0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		switch data[p] {
		case 79:
			goto st273
		case 111:
			goto st273
		}
		goto st0
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		switch data[p] {
		case 82:
			goto st274
		case 114:
			goto st274
		}
		goto st0
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		switch data[p] {
		case 87:
			goto st275
		case 119:
			goto st275
		}
		goto st0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 65:
			goto st276
		case 97:
			goto st276
		}
		goto st0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 82:
			goto st277
		case 114:
			goto st277
		}
		goto st0
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 68:
			goto st278
		case 100:
			goto st278
		}
		goto st0
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 83:
			goto st279
		case 115:
			goto st279
		}
		goto st0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch data[p] {
		case 9:
			goto st279
		case 32:
			goto st279
		case 58:
			goto st280
		}
		goto st0
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st280
		case 32:
			goto st280
		case 525:
			goto st282
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr362
		}
		goto st0
tr362:
//line msg_parse.rl:279
msg.MaxForwards=0
//line msg_parse.rl:169

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
	goto st281
tr364:
//line msg_parse.rl:169

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
	goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line msg_parse.go:6458
		if data[p] == 13 {
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr364
		}
		goto st0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		if data[p] == 10 {
			goto tr365
		}
		goto st0
tr365:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line msg_parse.go:6484
		switch data[p] {
		case 9:
			goto st284
		case 32:
			goto st284
		}
		goto st0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 9:
			goto st284
		case 32:
			goto st284
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr362
		}
		goto st0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 77:
			goto st286
		case 78:
			goto st296
		case 109:
			goto st286
		case 110:
			goto st296
		}
		goto st0
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		switch data[p] {
		case 69:
			goto st287
		case 101:
			goto st287
		}
		goto st0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		if data[p] == 45 {
			goto st288
		}
		goto st0
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		switch data[p] {
		case 86:
			goto st289
		case 118:
			goto st289
		}
		goto st0
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 69:
			goto st290
		case 101:
			goto st290
		}
		goto st0
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		switch data[p] {
		case 82:
			goto st291
		case 114:
			goto st291
		}
		goto st0
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		switch data[p] {
		case 83:
			goto st292
		case 115:
			goto st292
		}
		goto st0
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		switch data[p] {
		case 73:
			goto st293
		case 105:
			goto st293
		}
		goto st0
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		switch data[p] {
		case 79:
			goto st294
		case 111:
			goto st294
		}
		goto st0
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 78:
			goto st295
		case 110:
			goto st295
		}
		goto st0
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 9:
			goto tr378
		case 32:
			goto tr378
		case 58:
			goto tr379
		}
		goto st0
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		if data[p] == 45 {
			goto st297
		}
		goto st0
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		switch data[p] {
		case 69:
			goto st298
		case 101:
			goto st298
		}
		goto st0
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 88:
			goto st299
		case 120:
			goto st299
		}
		goto st0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 80:
			goto st300
		case 112:
			goto st300
		}
		goto st0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 73:
			goto st301
		case 105:
			goto st301
		}
		goto st0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		switch data[p] {
		case 82:
			goto st302
		case 114:
			goto st302
		}
		goto st0
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 69:
			goto st303
		case 101:
			goto st303
		}
		goto st0
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 83:
			goto st304
		case 115:
			goto st304
		}
		goto st0
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 9:
			goto st304
		case 32:
			goto st304
		case 58:
			goto st305
		}
		goto st0
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st305
		case 32:
			goto st305
		case 525:
			goto st307
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr389
		}
		goto st0
tr389:
//line msg_parse.rl:280
msg.MinExpires=0
//line msg_parse.rl:173

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st306
tr391:
//line msg_parse.rl:173

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st306
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
//line msg_parse.go:6792
		if data[p] == 13 {
			goto st121
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr391
		}
		goto st0
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		if data[p] == 10 {
			goto tr392
		}
		goto st0
tr392:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st308
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
//line msg_parse.go:6818
		switch data[p] {
		case 9:
			goto st309
		case 32:
			goto st309
		}
		goto st0
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		switch data[p] {
		case 9:
			goto st309
		case 32:
			goto st309
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr389
		}
		goto st0
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 9:
			goto tr302
		case 32:
			goto tr302
		case 58:
			goto tr303
		case 82:
			goto st311
		case 114:
			goto st311
		}
		goto st0
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		switch data[p] {
		case 71:
			goto st312
		case 103:
			goto st312
		}
		goto st0
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 65:
			goto st313
		case 97:
			goto st313
		}
		goto st0
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 78:
			goto st314
		case 110:
			goto st314
		}
		goto st0
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 73:
			goto st315
		case 105:
			goto st315
		}
		goto st0
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 90:
			goto st316
		case 122:
			goto st316
		}
		goto st0
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 65:
			goto st317
		case 97:
			goto st317
		}
		goto st0
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		switch data[p] {
		case 84:
			goto st318
		case 116:
			goto st318
		}
		goto st0
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 73:
			goto st319
		case 105:
			goto st319
		}
		goto st0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 79:
			goto st320
		case 111:
			goto st320
		}
		goto st0
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		switch data[p] {
		case 78:
			goto st321
		case 110:
			goto st321
		}
		goto st0
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		switch data[p] {
		case 9:
			goto tr405
		case 32:
			goto tr405
		case 58:
			goto tr406
		}
		goto st0
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 45:
			goto st323
		case 82:
			goto st347
		case 114:
			goto st347
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 65:
			goto st324
		case 97:
			goto st324
		}
		goto st0
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 83:
			goto st325
		case 115:
			goto st325
		}
		goto st0
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		switch data[p] {
		case 83:
			goto st326
		case 115:
			goto st326
		}
		goto st0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 69:
			goto st327
		case 101:
			goto st327
		}
		goto st0
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 82:
			goto st328
		case 114:
			goto st328
		}
		goto st0
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		switch data[p] {
		case 84:
			goto st329
		case 116:
			goto st329
		}
		goto st0
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
		switch data[p] {
		case 69:
			goto st330
		case 101:
			goto st330
		}
		goto st0
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 68:
			goto st331
		case 100:
			goto st331
		}
		goto st0
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		if data[p] == 45 {
			goto st332
		}
		goto st0
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 73:
			goto st333
		case 105:
			goto st333
		}
		goto st0
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 68:
			goto st334
		case 100:
			goto st334
		}
		goto st0
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 69:
			goto st335
		case 101:
			goto st335
		}
		goto st0
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		switch data[p] {
		case 78:
			goto st336
		case 110:
			goto st336
		}
		goto st0
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 84:
			goto st337
		case 116:
			goto st337
		}
		goto st0
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		switch data[p] {
		case 73:
			goto st338
		case 105:
			goto st338
		}
		goto st0
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		switch data[p] {
		case 84:
			goto st339
		case 116:
			goto st339
		}
		goto st0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 89:
			goto st340
		case 121:
			goto st340
		}
		goto st0
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 9:
			goto st340
		case 32:
			goto st340
		case 58:
			goto st341
		}
		goto st0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st341
		case 32:
			goto st341
		case 269:
			goto tr428
		case 525:
			goto st344
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr427
		}
		goto st0
tr427:
//line msg_parse.rl:54

			mark = p
		
	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line msg_parse.go:7259
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st342
		case 269:
			goto tr431
		case 525:
			goto tr432
		}
		if 32 <= _widec && _widec <= 253 {
			goto st342
		}
		goto st0
tr436:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:177

			msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st343
tr432:
//line msg_parse.rl:177

			msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line msg_parse.go:7302
		if data[p] == 10 {
			goto tr433
		}
		goto st0
tr433:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st556
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
//line msg_parse.go:7320
		switch data[p] {
		case 9:
			goto st342
		case 32:
			goto st342
		}
		goto st0
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		if data[p] == 10 {
			goto tr434
		}
		goto st0
tr434:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st345
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
//line msg_parse.go:7346
		switch data[p] {
		case 9:
			goto st346
		case 32:
			goto st346
		}
		goto st0
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st346
		case 32:
			goto st346
		case 269:
			goto tr428
		case 525:
			goto tr436
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr427
		}
		goto st0
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 73:
			goto st348
		case 79:
			goto st354
		case 105:
			goto st348
		case 111:
			goto st354
		}
		goto st0
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 79:
			goto st349
		case 111:
			goto st349
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 82:
			goto st350
		case 114:
			goto st350
		}
		goto st0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 73:
			goto st351
		case 105:
			goto st351
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		switch data[p] {
		case 84:
			goto st352
		case 116:
			goto st352
		}
		goto st0
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		switch data[p] {
		case 89:
			goto st353
		case 121:
			goto st353
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 9:
			goto tr444
		case 32:
			goto tr444
		case 58:
			goto tr445
		}
		goto st0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 88:
			goto st355
		case 120:
			goto st355
		}
		goto st0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 89:
			goto st356
		case 121:
			goto st356
		}
		goto st0
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		if data[p] == 45 {
			goto st357
		}
		goto st0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		switch data[p] {
		case 65:
			goto st358
		case 82:
			goto st379
		case 97:
			goto st358
		case 114:
			goto st379
		}
		goto st0
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 85:
			goto st359
		case 117:
			goto st359
		}
		goto st0
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		switch data[p] {
		case 84:
			goto st360
		case 116:
			goto st360
		}
		goto st0
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 72:
			goto st361
		case 104:
			goto st361
		}
		goto st0
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		switch data[p] {
		case 69:
			goto st362
		case 79:
			goto st370
		case 101:
			goto st362
		case 111:
			goto st370
		}
		goto st0
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 78:
			goto st363
		case 110:
			goto st363
		}
		goto st0
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 84:
			goto st364
		case 116:
			goto st364
		}
		goto st0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 73:
			goto st365
		case 105:
			goto st365
		}
		goto st0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 67:
			goto st366
		case 99:
			goto st366
		}
		goto st0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 65:
			goto st367
		case 97:
			goto st367
		}
		goto st0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 84:
			goto st368
		case 116:
			goto st368
		}
		goto st0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 69:
			goto st369
		case 101:
			goto st369
		}
		goto st0
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		switch data[p] {
		case 9:
			goto tr463
		case 32:
			goto tr463
		case 58:
			goto tr464
		}
		goto st0
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 82:
			goto st371
		case 114:
			goto st371
		}
		goto st0
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 73:
			goto st372
		case 105:
			goto st372
		}
		goto st0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		switch data[p] {
		case 90:
			goto st373
		case 122:
			goto st373
		}
		goto st0
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		switch data[p] {
		case 65:
			goto st374
		case 97:
			goto st374
		}
		goto st0
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
		switch data[p] {
		case 84:
			goto st375
		case 116:
			goto st375
		}
		goto st0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		switch data[p] {
		case 73:
			goto st376
		case 105:
			goto st376
		}
		goto st0
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
		switch data[p] {
		case 79:
			goto st377
		case 111:
			goto st377
		}
		goto st0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 78:
			goto st378
		case 110:
			goto st378
		}
		goto st0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 9:
			goto tr473
		case 32:
			goto tr473
		case 58:
			goto tr474
		}
		goto st0
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 69:
			goto st380
		case 101:
			goto st380
		}
		goto st0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		switch data[p] {
		case 81:
			goto st381
		case 113:
			goto st381
		}
		goto st0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 85:
			goto st382
		case 117:
			goto st382
		}
		goto st0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 73:
			goto st383
		case 105:
			goto st383
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 82:
			goto st384
		case 114:
			goto st384
		}
		goto st0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 69:
			goto st385
		case 101:
			goto st385
		}
		goto st0
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 9:
			goto tr481
		case 32:
			goto tr481
		case 58:
			goto tr482
		}
		goto st0
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 9:
			goto tr483
		case 32:
			goto tr483
		case 58:
			goto tr484
		case 69:
			goto st387
		case 79:
			goto st454
		case 101:
			goto st387
		case 111:
			goto st454
		}
		goto st0
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		switch data[p] {
		case 67:
			goto st388
		case 70:
			goto st404
		case 77:
			goto st415
		case 80:
			goto st434
		case 81:
			goto st440
		case 84:
			goto st445
		case 99:
			goto st388
		case 102:
			goto st404
		case 109:
			goto st415
		case 112:
			goto st434
		case 113:
			goto st440
		case 116:
			goto st445
		}
		goto st0
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		switch data[p] {
		case 79:
			goto st389
		case 111:
			goto st389
		}
		goto st0
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 82:
			goto st390
		case 114:
			goto st390
		}
		goto st0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		switch data[p] {
		case 68:
			goto st391
		case 100:
			goto st391
		}
		goto st0
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		if data[p] == 45 {
			goto st392
		}
		goto st0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 82:
			goto st393
		case 114:
			goto st393
		}
		goto st0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 79:
			goto st394
		case 111:
			goto st394
		}
		goto st0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 85:
			goto st395
		case 117:
			goto st395
		}
		goto st0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 84:
			goto st396
		case 116:
			goto st396
		}
		goto st0
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		switch data[p] {
		case 69:
			goto st397
		case 101:
			goto st397
		}
		goto st0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		switch data[p] {
		case 9:
			goto st397
		case 32:
			goto st397
		case 58:
			goto st398
		}
		goto st0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st398
		case 32:
			goto st398
		case 269:
			goto tr504
		case 525:
			goto st401
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr503
		}
		goto st0
tr503:
//line msg_parse.rl:54

			mark = p
		
	goto st399
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
//line msg_parse.go:8075
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st399
		case 269:
			goto tr507
		case 525:
			goto tr508
		}
		if 32 <= _widec && _widec <= 253 {
			goto st399
		}
		goto st0
tr512:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:182

			*rroutep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *rroutep != nil { rroutep = &(*rroutep).Next }
		
	goto st400
tr508:
//line msg_parse.rl:182

			*rroutep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *rroutep != nil { rroutep = &(*rroutep).Next }
		
	goto st400
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
//line msg_parse.go:8120
		if data[p] == 10 {
			goto tr509
		}
		goto st0
tr509:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st557
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
//line msg_parse.go:8138
		switch data[p] {
		case 9:
			goto st399
		case 32:
			goto st399
		}
		goto st0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if data[p] == 10 {
			goto tr510
		}
		goto st0
tr510:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st402
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
//line msg_parse.go:8164
		switch data[p] {
		case 9:
			goto st403
		case 32:
			goto st403
		}
		goto st0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st403
		case 32:
			goto st403
		case 269:
			goto tr504
		case 525:
			goto tr512
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr503
		}
		goto st0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 69:
			goto st405
		case 101:
			goto st405
		}
		goto st0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 82:
			goto st406
		case 114:
			goto st406
		}
		goto st0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		switch data[p] {
		case 45:
			goto st407
		case 82:
			goto st410
		case 114:
			goto st410
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 84:
			goto st408
		case 116:
			goto st408
		}
		goto st0
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 79:
			goto st409
		case 111:
			goto st409
		}
		goto st0
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		switch data[p] {
		case 9:
			goto tr483
		case 32:
			goto tr483
		case 58:
			goto tr484
		}
		goto st0
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		switch data[p] {
		case 69:
			goto st411
		case 101:
			goto st411
		}
		goto st0
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 68:
			goto st412
		case 100:
			goto st412
		}
		goto st0
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		if data[p] == 45 {
			goto st413
		}
		goto st0
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 66:
			goto st414
		case 98:
			goto st414
		}
		goto st0
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 89:
			goto st116
		case 121:
			goto st116
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 79:
			goto st416
		case 111:
			goto st416
		}
		goto st0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 84:
			goto st417
		case 116:
			goto st417
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 69:
			goto st418
		case 101:
			goto st418
		}
		goto st0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		if data[p] == 45 {
			goto st419
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 80:
			goto st420
		case 112:
			goto st420
		}
		goto st0
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		switch data[p] {
		case 65:
			goto st421
		case 97:
			goto st421
		}
		goto st0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 82:
			goto st422
		case 114:
			goto st422
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 84:
			goto st423
		case 116:
			goto st423
		}
		goto st0
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		switch data[p] {
		case 89:
			goto st424
		case 121:
			goto st424
		}
		goto st0
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		if data[p] == 45 {
			goto st425
		}
		goto st0
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 73:
			goto st426
		case 105:
			goto st426
		}
		goto st0
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		switch data[p] {
		case 68:
			goto st427
		case 100:
			goto st427
		}
		goto st0
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 9:
			goto st427
		case 32:
			goto st427
		case 58:
			goto st428
		}
		goto st0
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st428
		case 32:
			goto st428
		case 269:
			goto tr537
		case 525:
			goto st431
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr536
		}
		goto st0
tr536:
//line msg_parse.rl:54

			mark = p
		
	goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line msg_parse.go:8520
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st429
		case 269:
			goto tr540
		case 525:
			goto tr541
		}
		if 32 <= _widec && _widec <= 253 {
			goto st429
		}
		goto st0
tr545:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:188

			msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st430
tr541:
//line msg_parse.rl:188

			msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line msg_parse.go:8563
		if data[p] == 10 {
			goto tr542
		}
		goto st0
tr542:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st558
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
//line msg_parse.go:8581
		switch data[p] {
		case 9:
			goto st429
		case 32:
			goto st429
		}
		goto st0
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		if data[p] == 10 {
			goto tr543
		}
		goto st0
tr543:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line msg_parse.go:8607
		switch data[p] {
		case 9:
			goto st433
		case 32:
			goto st433
		}
		goto st0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st433
		case 32:
			goto st433
		case 269:
			goto tr537
		case 525:
			goto tr545
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr536
		}
		goto st0
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 76:
			goto st435
		case 108:
			goto st435
		}
		goto st0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 89:
			goto st436
		case 121:
			goto st436
		}
		goto st0
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		if data[p] == 45 {
			goto st437
		}
		goto st0
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 84:
			goto st438
		case 116:
			goto st438
		}
		goto st0
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		switch data[p] {
		case 79:
			goto st439
		case 111:
			goto st439
		}
		goto st0
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 9:
			goto tr551
		case 32:
			goto tr551
		case 58:
			goto tr552
		}
		goto st0
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		switch data[p] {
		case 85:
			goto st441
		case 117:
			goto st441
		}
		goto st0
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		switch data[p] {
		case 73:
			goto st442
		case 105:
			goto st442
		}
		goto st0
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 82:
			goto st443
		case 114:
			goto st443
		}
		goto st0
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		switch data[p] {
		case 69:
			goto st444
		case 101:
			goto st444
		}
		goto st0
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 9:
			goto tr557
		case 32:
			goto tr557
		case 58:
			goto tr558
		}
		goto st0
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 82:
			goto st446
		case 114:
			goto st446
		}
		goto st0
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 89:
			goto st447
		case 121:
			goto st447
		}
		goto st0
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		if data[p] == 45 {
			goto st448
		}
		goto st0
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 65:
			goto st449
		case 97:
			goto st449
		}
		goto st0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 70:
			goto st450
		case 102:
			goto st450
		}
		goto st0
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 84:
			goto st451
		case 116:
			goto st451
		}
		goto st0
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		switch data[p] {
		case 69:
			goto st452
		case 101:
			goto st452
		}
		goto st0
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 82:
			goto st453
		case 114:
			goto st453
		}
		goto st0
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		switch data[p] {
		case 9:
			goto tr567
		case 32:
			goto tr567
		case 58:
			goto tr568
		}
		goto st0
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 85:
			goto st455
		case 117:
			goto st455
		}
		goto st0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		switch data[p] {
		case 84:
			goto st456
		case 116:
			goto st456
		}
		goto st0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 69:
			goto st457
		case 101:
			goto st457
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 9:
			goto st457
		case 32:
			goto st457
		case 58:
			goto st458
		}
		goto st0
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st458
		case 32:
			goto st458
		case 269:
			goto tr574
		case 525:
			goto st461
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr573
		}
		goto st0
tr573:
//line msg_parse.rl:54

			mark = p
		
	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line msg_parse.go:8968
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st459
		case 269:
			goto tr577
		case 525:
			goto tr578
		}
		if 32 <= _widec && _widec <= 253 {
			goto st459
		}
		goto st0
tr582:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:193

			*routep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *routep != nil { routep = &(*routep).Next }
		
	goto st460
tr578:
//line msg_parse.rl:193

			*routep, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
			for *routep != nil { routep = &(*routep).Next }
		
	goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line msg_parse.go:9013
		if data[p] == 10 {
			goto tr579
		}
		goto st0
tr579:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st559
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
//line msg_parse.go:9031
		switch data[p] {
		case 9:
			goto st459
		case 32:
			goto st459
		}
		goto st0
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		if data[p] == 10 {
			goto tr580
		}
		goto st0
tr580:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line msg_parse.go:9057
		switch data[p] {
		case 9:
			goto st463
		case 32:
			goto st463
		}
		goto st0
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st463
		case 32:
			goto st463
		case 269:
			goto tr574
		case 525:
			goto tr582
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr573
		}
		goto st0
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		switch data[p] {
		case 9:
			goto tr583
		case 32:
			goto tr583
		case 58:
			goto tr584
		case 69:
			goto st465
		case 85:
			goto st470
		case 101:
			goto st465
		case 117:
			goto st470
		}
		goto st0
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		switch data[p] {
		case 82:
			goto st466
		case 114:
			goto st466
		}
		goto st0
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 86:
			goto st467
		case 118:
			goto st467
		}
		goto st0
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		switch data[p] {
		case 69:
			goto st468
		case 101:
			goto st468
		}
		goto st0
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
		switch data[p] {
		case 82:
			goto st469
		case 114:
			goto st469
		}
		goto st0
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 9:
			goto tr591
		case 32:
			goto tr591
		case 58:
			goto tr592
		}
		goto st0
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		switch data[p] {
		case 66:
			goto st471
		case 80:
			goto st476
		case 98:
			goto st471
		case 112:
			goto st476
		}
		goto st0
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 74:
			goto st472
		case 106:
			goto st472
		}
		goto st0
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		switch data[p] {
		case 69:
			goto st473
		case 101:
			goto st473
		}
		goto st0
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 67:
			goto st474
		case 99:
			goto st474
		}
		goto st0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 84:
			goto st475
		case 116:
			goto st475
		}
		goto st0
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 9:
			goto tr583
		case 32:
			goto tr583
		case 58:
			goto tr584
		}
		goto st0
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		switch data[p] {
		case 80:
			goto st477
		case 112:
			goto st477
		}
		goto st0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 79:
			goto st478
		case 111:
			goto st478
		}
		goto st0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 82:
			goto st479
		case 114:
			goto st479
		}
		goto st0
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 84:
			goto st480
		case 116:
			goto st480
		}
		goto st0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		switch data[p] {
		case 69:
			goto st481
		case 101:
			goto st481
		}
		goto st0
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 68:
			goto st261
		case 100:
			goto st261
		}
		goto st0
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		switch data[p] {
		case 9:
			goto st483
		case 32:
			goto st483
		case 58:
			goto st484
		case 73:
			goto st490
		case 79:
			goto st483
		case 105:
			goto st490
		case 111:
			goto st483
		}
		goto st0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
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
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st484
		case 32:
			goto st484
		case 269:
			goto tr608
		case 525:
			goto st487
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr607
		}
		goto st0
tr607:
//line msg_parse.rl:54

			mark = p
		
	goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line msg_parse.go:9398
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st485
		case 269:
			goto tr611
		case 525:
			goto tr612
		}
		if 32 <= _widec && _widec <= 253 {
			goto st485
		}
		goto st0
tr616:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:199

			msg.To, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st486
tr612:
//line msg_parse.rl:199

			msg.To, err = ParseAddr(string(data[mark:p]))
			if err != nil { return nil, err }
		
	goto st486
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
//line msg_parse.go:9441
		if data[p] == 10 {
			goto tr613
		}
		goto st0
tr613:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st560
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
//line msg_parse.go:9459
		switch data[p] {
		case 9:
			goto st485
		case 32:
			goto st485
		}
		goto st0
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		if data[p] == 10 {
			goto tr614
		}
		goto st0
tr614:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st488
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
//line msg_parse.go:9485
		switch data[p] {
		case 9:
			goto st489
		case 32:
			goto st489
		}
		goto st0
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st489
		case 32:
			goto st489
		case 269:
			goto tr608
		case 525:
			goto tr616
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr607
		}
		goto st0
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 77:
			goto st491
		case 109:
			goto st491
		}
		goto st0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 69:
			goto st492
		case 101:
			goto st492
		}
		goto st0
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		switch data[p] {
		case 83:
			goto st493
		case 115:
			goto st493
		}
		goto st0
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		switch data[p] {
		case 84:
			goto st494
		case 116:
			goto st494
		}
		goto st0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 65:
			goto st495
		case 97:
			goto st495
		}
		goto st0
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 77:
			goto st496
		case 109:
			goto st496
		}
		goto st0
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		switch data[p] {
		case 80:
			goto st497
		case 112:
			goto st497
		}
		goto st0
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 9:
			goto tr624
		case 32:
			goto tr624
		case 58:
			goto tr625
		}
		goto st0
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 9:
			goto tr626
		case 32:
			goto tr626
		case 58:
			goto tr627
		case 78:
			goto st499
		case 83:
			goto st509
		case 110:
			goto st499
		case 115:
			goto st509
		}
		goto st0
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 83:
			goto st500
		case 115:
			goto st500
		}
		goto st0
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 85:
			goto st501
		case 117:
			goto st501
		}
		goto st0
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 80:
			goto st502
		case 112:
			goto st502
		}
		goto st0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		switch data[p] {
		case 80:
			goto st503
		case 112:
			goto st503
		}
		goto st0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 79:
			goto st504
		case 111:
			goto st504
		}
		goto st0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 82:
			goto st505
		case 114:
			goto st505
		}
		goto st0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 84:
			goto st506
		case 116:
			goto st506
		}
		goto st0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 69:
			goto st507
		case 101:
			goto st507
		}
		goto st0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 68:
			goto st508
		case 100:
			goto st508
		}
		goto st0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		switch data[p] {
		case 9:
			goto tr639
		case 32:
			goto tr639
		case 58:
			goto tr640
		}
		goto st0
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		switch data[p] {
		case 69:
			goto st510
		case 101:
			goto st510
		}
		goto st0
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		switch data[p] {
		case 82:
			goto st511
		case 114:
			goto st511
		}
		goto st0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		if data[p] == 45 {
			goto st512
		}
		goto st0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		switch data[p] {
		case 65:
			goto st513
		case 97:
			goto st513
		}
		goto st0
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		switch data[p] {
		case 71:
			goto st514
		case 103:
			goto st514
		}
		goto st0
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		switch data[p] {
		case 69:
			goto st515
		case 101:
			goto st515
		}
		goto st0
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		switch data[p] {
		case 78:
			goto st516
		case 110:
			goto st516
		}
		goto st0
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		switch data[p] {
		case 84:
			goto st517
		case 116:
			goto st517
		}
		goto st0
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		switch data[p] {
		case 9:
			goto tr649
		case 32:
			goto tr649
		case 58:
			goto tr650
		}
		goto st0
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		switch data[p] {
		case 9:
			goto st519
		case 32:
			goto st519
		case 58:
			goto st520
		case 73:
			goto st526
		case 105:
			goto st526
		}
		goto st0
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 9:
			goto st519
		case 32:
			goto st519
		case 58:
			goto st520
		}
		goto st0
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st520
		case 32:
			goto st520
		case 269:
			goto tr655
		case 525:
			goto st523
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr654
		}
		goto st0
tr654:
//line msg_parse.rl:54

			mark = p
		
	goto st521
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
//line msg_parse.go:9937
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st521
		case 269:
			goto tr658
		case 525:
			goto tr659
		}
		if 32 <= _widec && _widec <= 253 {
			goto st521
		}
		goto st0
tr663:
//line msg_parse.rl:54

			mark = p
		
//line msg_parse.rl:204

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st522
tr659:
//line msg_parse.rl:204

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st522
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
//line msg_parse.go:9982
		if data[p] == 10 {
			goto tr660
		}
		goto st0
tr660:
//line msg_parse.rl:216
 line++; linep = p; 
//line msg_parse.rl:110

			{goto st33 }
		
	goto st561
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
//line msg_parse.go:10000
		switch data[p] {
		case 9:
			goto st521
		case 32:
			goto st521
		}
		goto st0
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		if data[p] == 10 {
			goto tr661
		}
		goto st0
tr661:
//line msg_parse.rl:216
 line++; linep = p; 
	goto st524
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
//line msg_parse.go:10026
		switch data[p] {
		case 9:
			goto st525
		case 32:
			goto st525
		}
		goto st0
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st525
		case 32:
			goto st525
		case 269:
			goto tr655
		case 525:
			goto tr663
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr654
		}
		goto st0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		switch data[p] {
		case 65:
			goto st519
		case 97:
			goto st519
		}
		goto st0
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		switch data[p] {
		case 65:
			goto st528
		case 87:
			goto st534
		case 97:
			goto st528
		case 119:
			goto st534
		}
		goto st0
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		switch data[p] {
		case 82:
			goto st529
		case 114:
			goto st529
		}
		goto st0
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		switch data[p] {
		case 78:
			goto st530
		case 110:
			goto st530
		}
		goto st0
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		switch data[p] {
		case 73:
			goto st531
		case 105:
			goto st531
		}
		goto st0
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 78:
			goto st532
		case 110:
			goto st532
		}
		goto st0
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		switch data[p] {
		case 71:
			goto st533
		case 103:
			goto st533
		}
		goto st0
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 9:
			goto tr671
		case 32:
			goto tr671
		case 58:
			goto tr672
		}
		goto st0
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		switch data[p] {
		case 87:
			goto st535
		case 119:
			goto st535
		}
		goto st0
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		if data[p] == 45 {
			goto st536
		}
		goto st0
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		switch data[p] {
		case 65:
			goto st537
		case 97:
			goto st537
		}
		goto st0
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		switch data[p] {
		case 85:
			goto st538
		case 117:
			goto st538
		}
		goto st0
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		switch data[p] {
		case 84:
			goto st539
		case 116:
			goto st539
		}
		goto st0
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		switch data[p] {
		case 72:
			goto st540
		case 104:
			goto st540
		}
		goto st0
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		switch data[p] {
		case 69:
			goto st541
		case 101:
			goto st541
		}
		goto st0
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		switch data[p] {
		case 78:
			goto st542
		case 110:
			goto st542
		}
		goto st0
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		switch data[p] {
		case 84:
			goto st543
		case 116:
			goto st543
		}
		goto st0
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		switch data[p] {
		case 73:
			goto st544
		case 105:
			goto st544
		}
		goto st0
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		switch data[p] {
		case 67:
			goto st545
		case 99:
			goto st545
		}
		goto st0
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		switch data[p] {
		case 65:
			goto st546
		case 97:
			goto st546
		}
		goto st0
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		switch data[p] {
		case 84:
			goto st547
		case 116:
			goto st547
		}
		goto st0
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		switch data[p] {
		case 69:
			goto st548
		case 101:
			goto st548
		}
		goto st0
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		switch data[p] {
		case 9:
			goto tr687
		case 32:
			goto tr687
		case 58:
			goto tr688
		}
		goto st0
	st_out:
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof549: cs = 549; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof550: cs = 550; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof551: cs = 551; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof552: cs = 552; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof553: cs = 553; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof554: cs = 554; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof
	_test_eof162: cs = 162; goto _test_eof
	_test_eof163: cs = 163; goto _test_eof
	_test_eof164: cs = 164; goto _test_eof
	_test_eof165: cs = 165; goto _test_eof
	_test_eof166: cs = 166; goto _test_eof
	_test_eof167: cs = 167; goto _test_eof
	_test_eof168: cs = 168; goto _test_eof
	_test_eof169: cs = 169; goto _test_eof
	_test_eof170: cs = 170; goto _test_eof
	_test_eof171: cs = 171; goto _test_eof
	_test_eof172: cs = 172; goto _test_eof
	_test_eof173: cs = 173; goto _test_eof
	_test_eof174: cs = 174; goto _test_eof
	_test_eof175: cs = 175; goto _test_eof
	_test_eof176: cs = 176; goto _test_eof
	_test_eof177: cs = 177; goto _test_eof
	_test_eof178: cs = 178; goto _test_eof
	_test_eof179: cs = 179; goto _test_eof
	_test_eof180: cs = 180; goto _test_eof
	_test_eof181: cs = 181; goto _test_eof
	_test_eof182: cs = 182; goto _test_eof
	_test_eof183: cs = 183; goto _test_eof
	_test_eof184: cs = 184; goto _test_eof
	_test_eof185: cs = 185; goto _test_eof
	_test_eof186: cs = 186; goto _test_eof
	_test_eof187: cs = 187; goto _test_eof
	_test_eof188: cs = 188; goto _test_eof
	_test_eof189: cs = 189; goto _test_eof
	_test_eof190: cs = 190; goto _test_eof
	_test_eof191: cs = 191; goto _test_eof
	_test_eof192: cs = 192; goto _test_eof
	_test_eof193: cs = 193; goto _test_eof
	_test_eof194: cs = 194; goto _test_eof
	_test_eof195: cs = 195; goto _test_eof
	_test_eof196: cs = 196; goto _test_eof
	_test_eof197: cs = 197; goto _test_eof
	_test_eof198: cs = 198; goto _test_eof
	_test_eof199: cs = 199; goto _test_eof
	_test_eof200: cs = 200; goto _test_eof
	_test_eof201: cs = 201; goto _test_eof
	_test_eof202: cs = 202; goto _test_eof
	_test_eof203: cs = 203; goto _test_eof
	_test_eof204: cs = 204; goto _test_eof
	_test_eof205: cs = 205; goto _test_eof
	_test_eof206: cs = 206; goto _test_eof
	_test_eof207: cs = 207; goto _test_eof
	_test_eof208: cs = 208; goto _test_eof
	_test_eof209: cs = 209; goto _test_eof
	_test_eof210: cs = 210; goto _test_eof
	_test_eof211: cs = 211; goto _test_eof
	_test_eof212: cs = 212; goto _test_eof
	_test_eof213: cs = 213; goto _test_eof
	_test_eof214: cs = 214; goto _test_eof
	_test_eof215: cs = 215; goto _test_eof
	_test_eof216: cs = 216; goto _test_eof
	_test_eof217: cs = 217; goto _test_eof
	_test_eof218: cs = 218; goto _test_eof
	_test_eof219: cs = 219; goto _test_eof
	_test_eof220: cs = 220; goto _test_eof
	_test_eof221: cs = 221; goto _test_eof
	_test_eof222: cs = 222; goto _test_eof
	_test_eof223: cs = 223; goto _test_eof
	_test_eof224: cs = 224; goto _test_eof
	_test_eof225: cs = 225; goto _test_eof
	_test_eof226: cs = 226; goto _test_eof
	_test_eof227: cs = 227; goto _test_eof
	_test_eof228: cs = 228; goto _test_eof
	_test_eof229: cs = 229; goto _test_eof
	_test_eof230: cs = 230; goto _test_eof
	_test_eof231: cs = 231; goto _test_eof
	_test_eof232: cs = 232; goto _test_eof
	_test_eof233: cs = 233; goto _test_eof
	_test_eof234: cs = 234; goto _test_eof
	_test_eof235: cs = 235; goto _test_eof
	_test_eof236: cs = 236; goto _test_eof
	_test_eof237: cs = 237; goto _test_eof
	_test_eof238: cs = 238; goto _test_eof
	_test_eof239: cs = 239; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
	_test_eof555: cs = 555; goto _test_eof
	_test_eof245: cs = 245; goto _test_eof
	_test_eof246: cs = 246; goto _test_eof
	_test_eof247: cs = 247; goto _test_eof
	_test_eof248: cs = 248; goto _test_eof
	_test_eof249: cs = 249; goto _test_eof
	_test_eof250: cs = 250; goto _test_eof
	_test_eof251: cs = 251; goto _test_eof
	_test_eof252: cs = 252; goto _test_eof
	_test_eof253: cs = 253; goto _test_eof
	_test_eof254: cs = 254; goto _test_eof
	_test_eof255: cs = 255; goto _test_eof
	_test_eof256: cs = 256; goto _test_eof
	_test_eof257: cs = 257; goto _test_eof
	_test_eof258: cs = 258; goto _test_eof
	_test_eof259: cs = 259; goto _test_eof
	_test_eof260: cs = 260; goto _test_eof
	_test_eof261: cs = 261; goto _test_eof
	_test_eof262: cs = 262; goto _test_eof
	_test_eof263: cs = 263; goto _test_eof
	_test_eof264: cs = 264; goto _test_eof
	_test_eof265: cs = 265; goto _test_eof
	_test_eof266: cs = 266; goto _test_eof
	_test_eof267: cs = 267; goto _test_eof
	_test_eof268: cs = 268; goto _test_eof
	_test_eof269: cs = 269; goto _test_eof
	_test_eof270: cs = 270; goto _test_eof
	_test_eof271: cs = 271; goto _test_eof
	_test_eof272: cs = 272; goto _test_eof
	_test_eof273: cs = 273; goto _test_eof
	_test_eof274: cs = 274; goto _test_eof
	_test_eof275: cs = 275; goto _test_eof
	_test_eof276: cs = 276; goto _test_eof
	_test_eof277: cs = 277; goto _test_eof
	_test_eof278: cs = 278; goto _test_eof
	_test_eof279: cs = 279; goto _test_eof
	_test_eof280: cs = 280; goto _test_eof
	_test_eof281: cs = 281; goto _test_eof
	_test_eof282: cs = 282; goto _test_eof
	_test_eof283: cs = 283; goto _test_eof
	_test_eof284: cs = 284; goto _test_eof
	_test_eof285: cs = 285; goto _test_eof
	_test_eof286: cs = 286; goto _test_eof
	_test_eof287: cs = 287; goto _test_eof
	_test_eof288: cs = 288; goto _test_eof
	_test_eof289: cs = 289; goto _test_eof
	_test_eof290: cs = 290; goto _test_eof
	_test_eof291: cs = 291; goto _test_eof
	_test_eof292: cs = 292; goto _test_eof
	_test_eof293: cs = 293; goto _test_eof
	_test_eof294: cs = 294; goto _test_eof
	_test_eof295: cs = 295; goto _test_eof
	_test_eof296: cs = 296; goto _test_eof
	_test_eof297: cs = 297; goto _test_eof
	_test_eof298: cs = 298; goto _test_eof
	_test_eof299: cs = 299; goto _test_eof
	_test_eof300: cs = 300; goto _test_eof
	_test_eof301: cs = 301; goto _test_eof
	_test_eof302: cs = 302; goto _test_eof
	_test_eof303: cs = 303; goto _test_eof
	_test_eof304: cs = 304; goto _test_eof
	_test_eof305: cs = 305; goto _test_eof
	_test_eof306: cs = 306; goto _test_eof
	_test_eof307: cs = 307; goto _test_eof
	_test_eof308: cs = 308; goto _test_eof
	_test_eof309: cs = 309; goto _test_eof
	_test_eof310: cs = 310; goto _test_eof
	_test_eof311: cs = 311; goto _test_eof
	_test_eof312: cs = 312; goto _test_eof
	_test_eof313: cs = 313; goto _test_eof
	_test_eof314: cs = 314; goto _test_eof
	_test_eof315: cs = 315; goto _test_eof
	_test_eof316: cs = 316; goto _test_eof
	_test_eof317: cs = 317; goto _test_eof
	_test_eof318: cs = 318; goto _test_eof
	_test_eof319: cs = 319; goto _test_eof
	_test_eof320: cs = 320; goto _test_eof
	_test_eof321: cs = 321; goto _test_eof
	_test_eof322: cs = 322; goto _test_eof
	_test_eof323: cs = 323; goto _test_eof
	_test_eof324: cs = 324; goto _test_eof
	_test_eof325: cs = 325; goto _test_eof
	_test_eof326: cs = 326; goto _test_eof
	_test_eof327: cs = 327; goto _test_eof
	_test_eof328: cs = 328; goto _test_eof
	_test_eof329: cs = 329; goto _test_eof
	_test_eof330: cs = 330; goto _test_eof
	_test_eof331: cs = 331; goto _test_eof
	_test_eof332: cs = 332; goto _test_eof
	_test_eof333: cs = 333; goto _test_eof
	_test_eof334: cs = 334; goto _test_eof
	_test_eof335: cs = 335; goto _test_eof
	_test_eof336: cs = 336; goto _test_eof
	_test_eof337: cs = 337; goto _test_eof
	_test_eof338: cs = 338; goto _test_eof
	_test_eof339: cs = 339; goto _test_eof
	_test_eof340: cs = 340; goto _test_eof
	_test_eof341: cs = 341; goto _test_eof
	_test_eof342: cs = 342; goto _test_eof
	_test_eof343: cs = 343; goto _test_eof
	_test_eof556: cs = 556; goto _test_eof
	_test_eof344: cs = 344; goto _test_eof
	_test_eof345: cs = 345; goto _test_eof
	_test_eof346: cs = 346; goto _test_eof
	_test_eof347: cs = 347; goto _test_eof
	_test_eof348: cs = 348; goto _test_eof
	_test_eof349: cs = 349; goto _test_eof
	_test_eof350: cs = 350; goto _test_eof
	_test_eof351: cs = 351; goto _test_eof
	_test_eof352: cs = 352; goto _test_eof
	_test_eof353: cs = 353; goto _test_eof
	_test_eof354: cs = 354; goto _test_eof
	_test_eof355: cs = 355; goto _test_eof
	_test_eof356: cs = 356; goto _test_eof
	_test_eof357: cs = 357; goto _test_eof
	_test_eof358: cs = 358; goto _test_eof
	_test_eof359: cs = 359; goto _test_eof
	_test_eof360: cs = 360; goto _test_eof
	_test_eof361: cs = 361; goto _test_eof
	_test_eof362: cs = 362; goto _test_eof
	_test_eof363: cs = 363; goto _test_eof
	_test_eof364: cs = 364; goto _test_eof
	_test_eof365: cs = 365; goto _test_eof
	_test_eof366: cs = 366; goto _test_eof
	_test_eof367: cs = 367; goto _test_eof
	_test_eof368: cs = 368; goto _test_eof
	_test_eof369: cs = 369; goto _test_eof
	_test_eof370: cs = 370; goto _test_eof
	_test_eof371: cs = 371; goto _test_eof
	_test_eof372: cs = 372; goto _test_eof
	_test_eof373: cs = 373; goto _test_eof
	_test_eof374: cs = 374; goto _test_eof
	_test_eof375: cs = 375; goto _test_eof
	_test_eof376: cs = 376; goto _test_eof
	_test_eof377: cs = 377; goto _test_eof
	_test_eof378: cs = 378; goto _test_eof
	_test_eof379: cs = 379; goto _test_eof
	_test_eof380: cs = 380; goto _test_eof
	_test_eof381: cs = 381; goto _test_eof
	_test_eof382: cs = 382; goto _test_eof
	_test_eof383: cs = 383; goto _test_eof
	_test_eof384: cs = 384; goto _test_eof
	_test_eof385: cs = 385; goto _test_eof
	_test_eof386: cs = 386; goto _test_eof
	_test_eof387: cs = 387; goto _test_eof
	_test_eof388: cs = 388; goto _test_eof
	_test_eof389: cs = 389; goto _test_eof
	_test_eof390: cs = 390; goto _test_eof
	_test_eof391: cs = 391; goto _test_eof
	_test_eof392: cs = 392; goto _test_eof
	_test_eof393: cs = 393; goto _test_eof
	_test_eof394: cs = 394; goto _test_eof
	_test_eof395: cs = 395; goto _test_eof
	_test_eof396: cs = 396; goto _test_eof
	_test_eof397: cs = 397; goto _test_eof
	_test_eof398: cs = 398; goto _test_eof
	_test_eof399: cs = 399; goto _test_eof
	_test_eof400: cs = 400; goto _test_eof
	_test_eof557: cs = 557; goto _test_eof
	_test_eof401: cs = 401; goto _test_eof
	_test_eof402: cs = 402; goto _test_eof
	_test_eof403: cs = 403; goto _test_eof
	_test_eof404: cs = 404; goto _test_eof
	_test_eof405: cs = 405; goto _test_eof
	_test_eof406: cs = 406; goto _test_eof
	_test_eof407: cs = 407; goto _test_eof
	_test_eof408: cs = 408; goto _test_eof
	_test_eof409: cs = 409; goto _test_eof
	_test_eof410: cs = 410; goto _test_eof
	_test_eof411: cs = 411; goto _test_eof
	_test_eof412: cs = 412; goto _test_eof
	_test_eof413: cs = 413; goto _test_eof
	_test_eof414: cs = 414; goto _test_eof
	_test_eof415: cs = 415; goto _test_eof
	_test_eof416: cs = 416; goto _test_eof
	_test_eof417: cs = 417; goto _test_eof
	_test_eof418: cs = 418; goto _test_eof
	_test_eof419: cs = 419; goto _test_eof
	_test_eof420: cs = 420; goto _test_eof
	_test_eof421: cs = 421; goto _test_eof
	_test_eof422: cs = 422; goto _test_eof
	_test_eof423: cs = 423; goto _test_eof
	_test_eof424: cs = 424; goto _test_eof
	_test_eof425: cs = 425; goto _test_eof
	_test_eof426: cs = 426; goto _test_eof
	_test_eof427: cs = 427; goto _test_eof
	_test_eof428: cs = 428; goto _test_eof
	_test_eof429: cs = 429; goto _test_eof
	_test_eof430: cs = 430; goto _test_eof
	_test_eof558: cs = 558; goto _test_eof
	_test_eof431: cs = 431; goto _test_eof
	_test_eof432: cs = 432; goto _test_eof
	_test_eof433: cs = 433; goto _test_eof
	_test_eof434: cs = 434; goto _test_eof
	_test_eof435: cs = 435; goto _test_eof
	_test_eof436: cs = 436; goto _test_eof
	_test_eof437: cs = 437; goto _test_eof
	_test_eof438: cs = 438; goto _test_eof
	_test_eof439: cs = 439; goto _test_eof
	_test_eof440: cs = 440; goto _test_eof
	_test_eof441: cs = 441; goto _test_eof
	_test_eof442: cs = 442; goto _test_eof
	_test_eof443: cs = 443; goto _test_eof
	_test_eof444: cs = 444; goto _test_eof
	_test_eof445: cs = 445; goto _test_eof
	_test_eof446: cs = 446; goto _test_eof
	_test_eof447: cs = 447; goto _test_eof
	_test_eof448: cs = 448; goto _test_eof
	_test_eof449: cs = 449; goto _test_eof
	_test_eof450: cs = 450; goto _test_eof
	_test_eof451: cs = 451; goto _test_eof
	_test_eof452: cs = 452; goto _test_eof
	_test_eof453: cs = 453; goto _test_eof
	_test_eof454: cs = 454; goto _test_eof
	_test_eof455: cs = 455; goto _test_eof
	_test_eof456: cs = 456; goto _test_eof
	_test_eof457: cs = 457; goto _test_eof
	_test_eof458: cs = 458; goto _test_eof
	_test_eof459: cs = 459; goto _test_eof
	_test_eof460: cs = 460; goto _test_eof
	_test_eof559: cs = 559; goto _test_eof
	_test_eof461: cs = 461; goto _test_eof
	_test_eof462: cs = 462; goto _test_eof
	_test_eof463: cs = 463; goto _test_eof
	_test_eof464: cs = 464; goto _test_eof
	_test_eof465: cs = 465; goto _test_eof
	_test_eof466: cs = 466; goto _test_eof
	_test_eof467: cs = 467; goto _test_eof
	_test_eof468: cs = 468; goto _test_eof
	_test_eof469: cs = 469; goto _test_eof
	_test_eof470: cs = 470; goto _test_eof
	_test_eof471: cs = 471; goto _test_eof
	_test_eof472: cs = 472; goto _test_eof
	_test_eof473: cs = 473; goto _test_eof
	_test_eof474: cs = 474; goto _test_eof
	_test_eof475: cs = 475; goto _test_eof
	_test_eof476: cs = 476; goto _test_eof
	_test_eof477: cs = 477; goto _test_eof
	_test_eof478: cs = 478; goto _test_eof
	_test_eof479: cs = 479; goto _test_eof
	_test_eof480: cs = 480; goto _test_eof
	_test_eof481: cs = 481; goto _test_eof
	_test_eof482: cs = 482; goto _test_eof
	_test_eof483: cs = 483; goto _test_eof
	_test_eof484: cs = 484; goto _test_eof
	_test_eof485: cs = 485; goto _test_eof
	_test_eof486: cs = 486; goto _test_eof
	_test_eof560: cs = 560; goto _test_eof
	_test_eof487: cs = 487; goto _test_eof
	_test_eof488: cs = 488; goto _test_eof
	_test_eof489: cs = 489; goto _test_eof
	_test_eof490: cs = 490; goto _test_eof
	_test_eof491: cs = 491; goto _test_eof
	_test_eof492: cs = 492; goto _test_eof
	_test_eof493: cs = 493; goto _test_eof
	_test_eof494: cs = 494; goto _test_eof
	_test_eof495: cs = 495; goto _test_eof
	_test_eof496: cs = 496; goto _test_eof
	_test_eof497: cs = 497; goto _test_eof
	_test_eof498: cs = 498; goto _test_eof
	_test_eof499: cs = 499; goto _test_eof
	_test_eof500: cs = 500; goto _test_eof
	_test_eof501: cs = 501; goto _test_eof
	_test_eof502: cs = 502; goto _test_eof
	_test_eof503: cs = 503; goto _test_eof
	_test_eof504: cs = 504; goto _test_eof
	_test_eof505: cs = 505; goto _test_eof
	_test_eof506: cs = 506; goto _test_eof
	_test_eof507: cs = 507; goto _test_eof
	_test_eof508: cs = 508; goto _test_eof
	_test_eof509: cs = 509; goto _test_eof
	_test_eof510: cs = 510; goto _test_eof
	_test_eof511: cs = 511; goto _test_eof
	_test_eof512: cs = 512; goto _test_eof
	_test_eof513: cs = 513; goto _test_eof
	_test_eof514: cs = 514; goto _test_eof
	_test_eof515: cs = 515; goto _test_eof
	_test_eof516: cs = 516; goto _test_eof
	_test_eof517: cs = 517; goto _test_eof
	_test_eof518: cs = 518; goto _test_eof
	_test_eof519: cs = 519; goto _test_eof
	_test_eof520: cs = 520; goto _test_eof
	_test_eof521: cs = 521; goto _test_eof
	_test_eof522: cs = 522; goto _test_eof
	_test_eof561: cs = 561; goto _test_eof
	_test_eof523: cs = 523; goto _test_eof
	_test_eof524: cs = 524; goto _test_eof
	_test_eof525: cs = 525; goto _test_eof
	_test_eof526: cs = 526; goto _test_eof
	_test_eof527: cs = 527; goto _test_eof
	_test_eof528: cs = 528; goto _test_eof
	_test_eof529: cs = 529; goto _test_eof
	_test_eof530: cs = 530; goto _test_eof
	_test_eof531: cs = 531; goto _test_eof
	_test_eof532: cs = 532; goto _test_eof
	_test_eof533: cs = 533; goto _test_eof
	_test_eof534: cs = 534; goto _test_eof
	_test_eof535: cs = 535; goto _test_eof
	_test_eof536: cs = 536; goto _test_eof
	_test_eof537: cs = 537; goto _test_eof
	_test_eof538: cs = 538; goto _test_eof
	_test_eof539: cs = 539; goto _test_eof
	_test_eof540: cs = 540; goto _test_eof
	_test_eof541: cs = 541; goto _test_eof
	_test_eof542: cs = 542; goto _test_eof
	_test_eof543: cs = 543; goto _test_eof
	_test_eof544: cs = 544; goto _test_eof
	_test_eof545: cs = 545; goto _test_eof
	_test_eof546: cs = 546; goto _test_eof
	_test_eof547: cs = 547; goto _test_eof
	_test_eof548: cs = 548; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line msg_parse.rl:341


	if cs < msg_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete SIP message: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in SIP message at line %d offset %d:\n%s", line, p - linep, data))
		}
	}

	if clen > 0 {
		if clen != len(data) - p {
			return nil, errors.New(fmt.Sprintf("Content-Length incorrect: %d != %d", clen, len(data) - p))
		}
		if ctype == sdp.ContentType {
			ms, err := sdp.Parse(string(data[p:len(data)]))
			if err != nil { return nil, err }
			msg.Payload = ms
		} else {
			msg.Payload = &MiscPayload{T: ctype, D: data[p:len(data)]}
		}
	}

	return msg, nil
}
