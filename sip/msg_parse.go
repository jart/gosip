
//line msg_parse.rl:1
// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
	"strings"
)


//line msg_parse.rl:13

//line msg_parse.go:18
const msg_start int = 1
const msg_first_final int = 634
const msg_error int = 0

const msg_en_avalue int = 34
const msg_en_svalue int = 129
const msg_en_xheader int = 138
const msg_en_header int = 145
const msg_en_main int = 1


//line msg_parse.rl:14

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
	fromp := &msg.From
	top := &msg.To
	viap := &msg.Via
	routep := &msg.Route
	rroutep := &msg.RecordRoute
	contactp := &msg.Contact
	paidp := &msg.PAssertedIdentity
	rpidp := &msg.RemotePartyID
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	line := 1
	linep := 0
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	clen := 0
	ctype := ""
	var name string
	var hex byte
	var value *string
	var addr *Addr
	var addrpp ***Addr

	
//line msg_parse.go:72
	{
	cs = msg_start
	}

//line msg_parse.go:77
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
	case 634:
		goto st_case_634
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
	case 635:
		goto st_case_635
	case 63:
		goto st_case_63
	case 636:
		goto st_case_636
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
	case 637:
		goto st_case_637
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
	case 638:
		goto st_case_638
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
	case 639:
		goto st_case_639
	case 107:
		goto st_case_107
	case 640:
		goto st_case_640
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 641:
		goto st_case_641
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
	case 642:
		goto st_case_642
	case 137:
		goto st_case_137
	case 643:
		goto st_case_643
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 644:
		goto st_case_644
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
	case 645:
		goto st_case_645
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
	case 646:
		goto st_case_646
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
	case 647:
		goto st_case_647
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
tr242:
//line msg_parse.rl:144

			p--

			{goto st138 }
		
//line msg_parse.rl:138

			p--

			{goto st138 }
		
	goto st0
tr274:
//line msg_parse.rl:138

			p--

			{goto st138 }
		
	goto st0
tr463:
//line msg_parse.rl:144

			p--

			{goto st138 }
		
	goto st0
//line msg_parse.go:1448
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line msg_parse.rl:59

			mark = p
		
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
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
//line msg_parse.rl:99

			msg.Method = string(data[mark:p])
		
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line msg_parse.go:1511
		if data[p] == 32 {
			goto st0
		}
		goto tr5
tr5:
//line msg_parse.rl:59

			mark = p
		
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line msg_parse.go:1527
		if data[p] == 32 {
			goto tr7
		}
		goto st4
tr7:
//line msg_parse.rl:111

			msg.Request, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line msg_parse.go:1544
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
//line msg_parse.rl:103

			msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
		
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line msg_parse.go:1596
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
//line msg_parse.rl:107

			msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
		
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line msg_parse.go:1624
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
tr36:
//line msg_parse.rl:120

			msg.Phrase = string(buf[0:amt])
		
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line msg_parse.go:1643
		if data[p] == 10 {
			goto tr16
		}
		goto st0
tr16:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st634
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
//line msg_parse.go:1661
		goto st0
tr2:
//line msg_parse.rl:59

			mark = p
		
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line msg_parse.go:1674
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
//line msg_parse.rl:103

			msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
		
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line msg_parse.go:1811
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
//line msg_parse.rl:107

			msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
		
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line msg_parse.go:1839
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
//line msg_parse.rl:116

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line msg_parse.go:1867
		if 48 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
tr25:
//line msg_parse.rl:116

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line msg_parse.go:1883
		if 48 <= data[p] && data[p] <= 57 {
			goto tr26
		}
		goto st0
tr26:
//line msg_parse.rl:116

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line msg_parse.go:1899
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
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st26
tr35:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st26
tr44:
//line msg_parse.rl:89

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line msg_parse.go:1993
		switch data[p] {
		case 9:
			goto tr35
		case 13:
			goto tr36
		case 37:
			goto st27
		case 61:
			goto tr35
		case 95:
			goto tr35
		case 126:
			goto tr35
		}
		switch {
		case data[p] < 192:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 33 {
					goto tr35
				}
			case data[p] > 59:
				switch {
				case data[p] > 90:
					if 97 <= data[p] && data[p] <= 122 {
						goto tr35
					}
				case data[p] >= 63:
					goto tr35
				}
			default:
				goto tr35
			}
		case data[p] > 223:
			switch {
			case data[p] < 240:
				if 224 <= data[p] && data[p] <= 239 {
					goto tr39
				}
			case data[p] > 247:
				switch {
				case data[p] > 251:
					if 252 <= data[p] && data[p] <= 253 {
						goto tr42
					}
				case data[p] >= 248:
					goto tr41
				}
			default:
				goto tr40
			}
		default:
			goto tr38
		}
		goto st0
tr29:
//line msg_parse.rl:63

			amt = 0
		
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line msg_parse.go:2060
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
//line msg_parse.rl:85

			hex = unhex(data[p]) * 16
		
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line msg_parse.go:2085
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
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st29
tr38:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line msg_parse.go:2122
		if 128 <= data[p] && data[p] <= 191 {
			goto tr35
		}
		goto st0
tr31:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st30
tr39:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line msg_parse.go:2150
		if 128 <= data[p] && data[p] <= 191 {
			goto tr38
		}
		goto st0
tr32:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st31
tr40:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line msg_parse.go:2178
		if 128 <= data[p] && data[p] <= 191 {
			goto tr39
		}
		goto st0
tr33:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st32
tr41:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line msg_parse.go:2206
		if 128 <= data[p] && data[p] <= 191 {
			goto tr40
		}
		goto st0
tr34:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st33
tr42:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line msg_parse.go:2234
		if 128 <= data[p] && data[p] <= 191 {
			goto tr41
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr45
		case 32:
			goto tr45
		case 33:
			goto tr46
		case 34:
			goto tr47
		case 37:
			goto tr46
		case 39:
			goto tr46
		case 60:
			goto tr48
		case 126:
			goto tr46
		case 525:
			goto tr50
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr46
				}
			case _widec >= 42:
				goto tr46
			}
		case _widec > 57:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 90 {
					goto tr49
				}
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr49
				}
			default:
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr45:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line msg_parse.go:2317
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st35
		case 32:
			goto st35
		case 34:
			goto st36
		case 60:
			goto st39
		case 525:
			goto st98
		}
		goto st0
tr47:
//line msg_parse.rl:174

			addr = new(Addr)
		
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line msg_parse.go:2349
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr55
		case 34:
			goto tr56
		case 92:
			goto tr57
		case 525:
			goto tr63
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr58
				}
			case _widec >= 32:
				goto tr55
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr60
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr62
				}
			default:
				goto tr61
			}
		default:
			goto tr59
		}
		goto st0
tr55:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st37
tr64:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line msg_parse.go:2417
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr64
		case 34:
			goto st38
		case 92:
			goto st116
		case 525:
			goto st122
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr67
				}
			case _widec >= 32:
				goto tr64
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr69
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr71
				}
			default:
				goto tr70
			}
		default:
			goto tr68
		}
		goto st0
tr56:
//line msg_parse.rl:63

			amt = 0
		
	goto st38
tr73:
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line msg_parse.go:2479
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr73
		case 32:
			goto tr73
		case 60:
			goto tr74
		case 525:
			goto tr75
		}
		goto st0
tr48:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st39
tr74:
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st39
tr94:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line msg_parse.go:2534
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr76
			}
		case data[p] >= 65:
			goto tr76
		}
		goto st0
tr76:
//line msg_parse.rl:59

			mark = p
		
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line msg_parse.go:2555
		switch data[p] {
		case 43:
			goto st40
		case 58:
			goto st41
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st40
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st40
				}
			case data[p] >= 65:
				goto st40
			}
		default:
			goto st40
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 33:
			goto st42
		case 61:
			goto st42
		case 95:
			goto st42
		case 126:
			goto st42
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 33:
			goto st42
		case 62:
			goto tr80
		case 95:
			goto st42
		case 126:
			goto st42
		}
		switch {
		case data[p] < 61:
			if 36 <= data[p] && data[p] <= 59 {
				goto st42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st42
			}
		default:
			goto st42
		}
		goto st0
tr80:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st43
tr81:
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line msg_parse.go:2659
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr81
		case 32:
			goto tr81
		case 44:
			goto tr82
		case 59:
			goto st57
		case 269:
			goto tr84
		case 525:
			goto tr85
		}
		goto st0
tr86:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st44
tr82:
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st44
tr117:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st44
tr130:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line msg_parse.go:2752
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr86
		case 32:
			goto tr86
		case 33:
			goto tr46
		case 34:
			goto tr47
		case 37:
			goto tr46
		case 39:
			goto tr46
		case 60:
			goto tr48
		case 126:
			goto tr46
		case 525:
			goto tr88
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr46
				}
			case _widec >= 42:
				goto tr46
			}
		case _widec > 57:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 90 {
					goto tr87
				}
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr87
				}
			default:
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr90:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st45
tr46:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st45
tr93:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line msg_parse.go:2846
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st46
		case 32:
			goto st46
		case 33:
			goto tr90
		case 37:
			goto tr90
		case 39:
			goto tr90
		case 126:
			goto tr90
		case 525:
			goto st47
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr90
				}
			case _widec >= 42:
				goto tr90
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr90
				}
			case _widec >= 65:
				goto tr90
			}
		default:
			goto tr90
		}
		goto st0
tr92:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line msg_parse.go:2909
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr92
		case 32:
			goto tr92
		case 33:
			goto tr93
		case 37:
			goto tr93
		case 39:
			goto tr93
		case 60:
			goto tr94
		case 126:
			goto tr93
		case 525:
			goto tr95
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr93
				}
			case _widec >= 42:
				goto tr93
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr93
				}
			case _widec >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
tr95:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line msg_parse.go:2974
		if data[p] == 10 {
			goto tr96
		}
		goto st0
tr96:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line msg_parse.go:2988
		switch data[p] {
		case 9:
			goto st49
		case 32:
			goto st49
		}
		goto st0
tr98:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line msg_parse.go:3012
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr98
		case 32:
			goto tr98
		case 33:
			goto tr93
		case 37:
			goto tr93
		case 39:
			goto tr93
		case 60:
			goto tr94
		case 126:
			goto tr93
		case 525:
			goto tr99
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr93
				}
			case _widec >= 42:
				goto tr93
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr93
				}
			case _widec >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
tr201:
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st50
tr99:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line msg_parse.go:3083
		if data[p] == 10 {
			goto tr100
		}
		goto st0
tr100:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line msg_parse.go:3097
		switch data[p] {
		case 9:
			goto st52
		case 32:
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 9:
			goto st52
		case 32:
			goto st52
		case 60:
			goto st39
		}
		goto st0
tr102:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st53
tr87:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line msg_parse.go:3150
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st46
		case 32:
			goto st46
		case 33:
			goto tr90
		case 37:
			goto tr90
		case 39:
			goto tr90
		case 42:
			goto tr90
		case 43:
			goto tr102
		case 58:
			goto st54
		case 126:
			goto tr90
		case 525:
			goto st47
		}
		switch {
		case _widec < 65:
			switch {
			case _widec > 46:
				if 48 <= _widec && _widec <= 57 {
					goto tr102
				}
			case _widec >= 45:
				goto tr102
			}
		case _widec > 90:
			switch {
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr102
				}
			case _widec >= 95:
				goto tr90
			}
		default:
			goto tr102
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 33:
			goto st55
		case 61:
			goto st55
		case 95:
			goto st55
		case 126:
			goto st55
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 58 {
				goto st55
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st55
			}
		default:
			goto st55
		}
		goto st0
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr105
		case 32:
			goto tr105
		case 33:
			goto st55
		case 44:
			goto tr106
		case 59:
			goto tr107
		case 61:
			goto st55
		case 95:
			goto st55
		case 126:
			goto st55
		case 269:
			goto tr108
		case 525:
			goto tr109
		}
		switch {
		case _widec < 63:
			if 36 <= _widec && _widec <= 58 {
				goto st55
			}
		case _widec > 90:
			if 97 <= _widec && _widec <= 122 {
				goto st55
			}
		default:
			goto st55
		}
		goto st0
tr105:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st56
tr128:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line msg_parse.go:3317
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st56
		case 32:
			goto st56
		case 44:
			goto st44
		case 59:
			goto st57
		case 525:
			goto st91
		}
		goto st0
tr107:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st57
tr118:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st57
tr131:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line msg_parse.go:3372
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st57
		case 32:
			goto st57
		case 33:
			goto tr113
		case 37:
			goto tr113
		case 39:
			goto tr113
		case 126:
			goto tr113
		case 525:
			goto st88
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr113
				}
			case _widec >= 42:
				goto tr113
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr113
				}
			case _widec >= 65:
				goto tr113
			}
		default:
			goto tr113
		}
		goto st0
tr113:
//line msg_parse.rl:59

			mark = p
		
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line msg_parse.go:3430
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr115
		case 32:
			goto tr115
		case 33:
			goto st58
		case 37:
			goto st58
		case 39:
			goto st58
		case 44:
			goto tr117
		case 59:
			goto tr118
		case 61:
			goto tr119
		case 126:
			goto st58
		case 269:
			goto tr120
		case 525:
			goto tr121
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto st58
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st58
				}
			case _widec >= 65:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
tr115:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line msg_parse.go:3507
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st59
		case 32:
			goto st59
		case 44:
			goto st44
		case 59:
			goto st57
		case 61:
			goto st60
		case 525:
			goto st84
		}
		goto st0
tr119:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line msg_parse.go:3541
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st60
		case 32:
			goto st60
		case 33:
			goto tr125
		case 34:
			goto st65
		case 37:
			goto tr125
		case 39:
			goto tr125
		case 93:
			goto tr125
		case 126:
			goto tr125
		case 525:
			goto st78
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr125
				}
			case _widec >= 42:
				goto tr125
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr125
				}
			case _widec >= 65:
				goto tr125
			}
		default:
			goto tr125
		}
		goto st0
tr125:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st61
tr129:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line msg_parse.go:3615
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr128
		case 32:
			goto tr128
		case 33:
			goto tr129
		case 37:
			goto tr129
		case 39:
			goto tr129
		case 44:
			goto tr130
		case 59:
			goto tr131
		case 93:
			goto tr129
		case 126:
			goto tr129
		case 269:
			goto tr132
		case 525:
			goto tr133
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr129
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr129
				}
			case _widec >= 65:
				goto tr129
			}
		default:
			goto tr129
		}
		goto st0
tr84:
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st62
tr108:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st62
tr120:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st62
tr132:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line msg_parse.go:3737
		if data[p] == 10 {
			goto tr134
		}
		goto st0
tr134:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st635
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
//line msg_parse.go:3755
		goto st0
tr197:
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st63
tr109:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st63
tr133:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line msg_parse.go:3807
		if data[p] == 10 {
			goto tr135
		}
		goto st0
tr135:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st636
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
//line msg_parse.go:3825
		switch data[p] {
		case 9:
			goto st64
		case 32:
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 9:
			goto st64
		case 32:
			goto st64
		case 44:
			goto st44
		case 59:
			goto st57
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr137
		case 34:
			goto tr138
		case 92:
			goto tr139
		case 525:
			goto tr145
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr140
				}
			case _widec >= 32:
				goto tr137
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr142
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr144
				}
			default:
				goto tr143
			}
		default:
			goto tr141
		}
		goto st0
tr137:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st66
tr146:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line msg_parse.go:3921
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr146
		case 34:
			goto st67
		case 92:
			goto st70
		case 525:
			goto st76
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr149
				}
			case _widec >= 32:
				goto tr146
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr151
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr153
				}
			default:
				goto tr152
			}
		default:
			goto tr150
		}
		goto st0
tr138:
//line msg_parse.rl:63

			amt = 0
		
	goto st67
tr155:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line msg_parse.go:3995
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr155
		case 32:
			goto tr155
		case 44:
			goto tr130
		case 59:
			goto tr131
		case 269:
			goto tr132
		case 525:
			goto tr156
		}
		goto st0
tr156:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line msg_parse.go:4041
		if data[p] == 10 {
			goto tr157
		}
		goto st0
tr157:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st637
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
//line msg_parse.go:4059
		switch data[p] {
		case 9:
			goto st69
		case 32:
			goto st69
		}
		goto st0
tr158:
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line msg_parse.go:4090
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr158
		case 32:
			goto tr158
		case 44:
			goto tr130
		case 59:
			goto tr131
		case 269:
			goto tr132
		case 525:
			goto tr133
		}
		goto st0
tr139:
//line msg_parse.rl:63

			amt = 0
		
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line msg_parse.go:4124
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr146
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr146
			}
		default:
			goto tr146
		}
		goto st0
tr140:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st71
tr149:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line msg_parse.go:4161
		if 128 <= data[p] && data[p] <= 191 {
			goto tr146
		}
		goto st0
tr141:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st72
tr150:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line msg_parse.go:4189
		if 128 <= data[p] && data[p] <= 191 {
			goto tr149
		}
		goto st0
tr142:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st73
tr151:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line msg_parse.go:4217
		if 128 <= data[p] && data[p] <= 191 {
			goto tr150
		}
		goto st0
tr143:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st74
tr152:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line msg_parse.go:4245
		if 128 <= data[p] && data[p] <= 191 {
			goto tr151
		}
		goto st0
tr144:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st75
tr153:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line msg_parse.go:4273
		if 128 <= data[p] && data[p] <= 191 {
			goto tr152
		}
		goto st0
tr145:
//line msg_parse.rl:63

			amt = 0
		
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line msg_parse.go:4289
		if data[p] == 10 {
			goto tr159
		}
		goto st0
tr159:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line msg_parse.go:4303
		switch data[p] {
		case 9:
			goto tr146
		case 32:
			goto tr146
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		if data[p] == 10 {
			goto tr160
		}
		goto st0
tr160:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line msg_parse.go:4329
		switch data[p] {
		case 9:
			goto st80
		case 32:
			goto st80
		}
		goto st0
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st80
		case 32:
			goto st80
		case 33:
			goto tr125
		case 34:
			goto st65
		case 37:
			goto tr125
		case 39:
			goto tr125
		case 93:
			goto tr125
		case 126:
			goto tr125
		case 525:
			goto st81
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr125
				}
			case _widec >= 42:
				goto tr125
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr125
				}
			case _widec >= 65:
				goto tr125
			}
		default:
			goto tr125
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		if data[p] == 10 {
			goto tr163
		}
		goto st0
tr163:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line msg_parse.go:4410
		switch data[p] {
		case 9:
			goto st83
		case 32:
			goto st83
		}
		goto st0
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
		switch data[p] {
		case 9:
			goto st83
		case 32:
			goto st83
		case 34:
			goto st65
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 10 {
			goto tr165
		}
		goto st0
tr165:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line msg_parse.go:4450
		switch data[p] {
		case 9:
			goto st86
		case 32:
			goto st86
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		switch data[p] {
		case 9:
			goto st86
		case 32:
			goto st86
		case 44:
			goto st44
		case 59:
			goto st57
		case 61:
			goto st60
		}
		goto st0
tr121:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
//line msg_parse.rl:187

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line msg_parse.go:4503
		if data[p] == 10 {
			goto tr167
		}
		goto st0
tr167:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st638
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
//line msg_parse.go:4521
		switch data[p] {
		case 9:
			goto st86
		case 32:
			goto st86
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		if data[p] == 10 {
			goto tr168
		}
		goto st0
tr168:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line msg_parse.go:4547
		switch data[p] {
		case 9:
			goto st90
		case 32:
			goto st90
		}
		goto st0
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
		switch data[p] {
		case 9:
			goto st90
		case 32:
			goto st90
		case 33:
			goto tr113
		case 37:
			goto tr113
		case 39:
			goto tr113
		case 126:
			goto tr113
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr113
				}
			case data[p] >= 42:
				goto tr113
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr113
				}
			case data[p] >= 65:
				goto tr113
			}
		default:
			goto tr113
		}
		goto st0
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
		if data[p] == 10 {
			goto tr170
		}
		goto st0
tr170:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line msg_parse.go:4615
		switch data[p] {
		case 9:
			goto st64
		case 32:
			goto st64
		}
		goto st0
tr106:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line msg_parse.go:4644
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr171
		case 32:
			goto tr171
		case 33:
			goto tr172
		case 34:
			goto tr47
		case 37:
			goto tr172
		case 39:
			goto tr172
		case 44:
			goto tr106
		case 47:
			goto st55
		case 59:
			goto tr107
		case 60:
			goto tr48
		case 95:
			goto tr172
		case 96:
			goto tr46
		case 126:
			goto tr172
		case 269:
			goto tr108
		case 525:
			goto tr174
		}
		switch {
		case _widec < 58:
			switch {
			case _widec > 41:
				if 42 <= _widec && _widec <= 57 {
					goto tr172
				}
			case _widec >= 36:
				goto st55
			}
		case _widec > 61:
			switch {
			case _widec < 65:
				if 63 <= _widec && _widec <= 64 {
					goto st55
				}
			case _widec > 90:
				if 97 <= _widec && _widec <= 122 {
					goto tr173
				}
			default:
				goto tr173
			}
		default:
			goto st55
		}
		goto st0
tr175:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st94
tr171:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st94
tr207:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line msg_parse.go:4786
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr175
		case 32:
			goto tr175
		case 33:
			goto tr46
		case 34:
			goto tr47
		case 37:
			goto tr46
		case 39:
			goto tr46
		case 44:
			goto st44
		case 59:
			goto st57
		case 60:
			goto tr48
		case 126:
			goto tr46
		case 525:
			goto tr176
		}
		switch {
		case _widec < 65:
			switch {
			case _widec > 46:
				if 48 <= _widec && _widec <= 57 {
					goto tr46
				}
			case _widec >= 42:
				goto tr46
			}
		case _widec > 90:
			switch {
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr87
				}
			case _widec >= 95:
				goto tr46
			}
		default:
			goto tr87
		}
		goto st0
tr176:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line msg_parse.go:4860
		if data[p] == 10 {
			goto tr177
		}
		goto st0
tr177:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line msg_parse.go:4874
		switch data[p] {
		case 9:
			goto st97
		case 32:
			goto st97
		}
		goto st0
tr179:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line msg_parse.go:4901
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr179
		case 32:
			goto tr179
		case 33:
			goto tr46
		case 34:
			goto tr47
		case 37:
			goto tr46
		case 39:
			goto tr46
		case 44:
			goto st44
		case 59:
			goto st57
		case 60:
			goto tr48
		case 126:
			goto tr46
		case 525:
			goto tr50
		}
		switch {
		case _widec < 65:
			switch {
			case _widec > 46:
				if 48 <= _widec && _widec <= 57 {
					goto tr46
				}
			case _widec >= 42:
				goto tr46
			}
		case _widec > 90:
			switch {
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr87
				}
			case _widec >= 95:
				goto tr46
			}
		default:
			goto tr87
		}
		goto st0
tr50:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line msg_parse.go:4975
		if data[p] == 10 {
			goto tr180
		}
		goto st0
tr180:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line msg_parse.go:4989
		switch data[p] {
		case 9:
			goto st100
		case 32:
			goto st100
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 9:
			goto st100
		case 32:
			goto st100
		case 34:
			goto st36
		case 60:
			goto st39
		}
		goto st0
tr183:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st101
tr172:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st101
tr173:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line msg_parse.go:5059
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr182
		case 32:
			goto tr182
		case 33:
			goto tr183
		case 37:
			goto tr183
		case 39:
			goto tr183
		case 44:
			goto tr106
		case 47:
			goto st55
		case 58:
			goto st55
		case 59:
			goto tr107
		case 61:
			goto st55
		case 96:
			goto tr90
		case 126:
			goto tr183
		case 269:
			goto tr108
		case 525:
			goto tr184
		}
		switch {
		case _widec < 63:
			switch {
			case _widec > 41:
				if 42 <= _widec && _widec <= 57 {
					goto tr183
				}
			case _widec >= 36:
				goto st55
			}
		case _widec > 64:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr183
				}
			case _widec >= 65:
				goto tr183
			}
		default:
			goto st55
		}
		goto st0
tr185:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st102
tr182:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st102
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
//line msg_parse.go:5152
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr185
		case 32:
			goto tr185
		case 33:
			goto tr93
		case 37:
			goto tr93
		case 39:
			goto tr93
		case 44:
			goto st44
		case 59:
			goto st57
		case 60:
			goto tr94
		case 126:
			goto tr93
		case 525:
			goto tr186
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr93
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr93
				}
			case _widec >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
tr186:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line msg_parse.go:5216
		if data[p] == 10 {
			goto tr187
		}
		goto st0
tr187:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line msg_parse.go:5230
		switch data[p] {
		case 9:
			goto st105
		case 32:
			goto st105
		}
		goto st0
tr189:
//line msg_parse.rl:72

			buf[amt] = ' '
			amt++
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line msg_parse.go:5254
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr189
		case 32:
			goto tr189
		case 33:
			goto tr93
		case 37:
			goto tr93
		case 39:
			goto tr93
		case 44:
			goto st44
		case 59:
			goto st57
		case 60:
			goto tr94
		case 126:
			goto tr93
		case 525:
			goto tr99
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr93
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr93
				}
			case _widec >= 65:
				goto tr93
			}
		default:
			goto tr93
		}
		goto st0
tr184:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line msg_parse.go:5323
		if data[p] == 10 {
			goto tr190
		}
		goto st0
tr190:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st639
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
//line msg_parse.go:5341
		switch data[p] {
		case 9:
			goto st105
		case 32:
			goto st105
		}
		goto st0
tr174:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st107
tr210:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line msg_parse.go:5410
		if data[p] == 10 {
			goto tr191
		}
		goto st0
tr191:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st640
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
//line msg_parse.go:5428
		switch data[p] {
		case 9:
			goto st97
		case 32:
			goto st97
		}
		goto st0
tr88:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line msg_parse.go:5455
		if data[p] == 10 {
			goto tr192
		}
		goto st0
tr192:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line msg_parse.go:5469
		switch data[p] {
		case 9:
			goto st110
		case 32:
			goto st110
		}
		goto st0
tr194:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line msg_parse.go:5496
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr194
		case 32:
			goto tr194
		case 33:
			goto tr46
		case 34:
			goto tr47
		case 37:
			goto tr46
		case 39:
			goto tr46
		case 60:
			goto tr48
		case 126:
			goto tr46
		case 525:
			goto tr50
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr46
				}
			case _widec >= 42:
				goto tr46
			}
		case _widec > 57:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 90 {
					goto tr87
				}
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr87
				}
			default:
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr85:
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line msg_parse.go:5567
		if data[p] == 10 {
			goto tr195
		}
		goto st0
tr195:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st641
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
//line msg_parse.go:5585
		switch data[p] {
		case 9:
			goto st112
		case 32:
			goto st112
		}
		goto st0
tr196:
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line msg_parse.go:5609
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr196
		case 32:
			goto tr196
		case 44:
			goto tr82
		case 59:
			goto st57
		case 269:
			goto tr84
		case 525:
			goto tr197
		}
		goto st0
tr75:
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line msg_parse.go:5643
		if data[p] == 10 {
			goto tr198
		}
		goto st0
tr198:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line msg_parse.go:5657
		switch data[p] {
		case 9:
			goto st115
		case 32:
			goto st115
		}
		goto st0
tr200:
//line msg_parse.rl:178

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line msg_parse.go:5676
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr200
		case 32:
			goto tr200
		case 60:
			goto tr74
		case 525:
			goto tr201
		}
		goto st0
tr57:
//line msg_parse.rl:63

			amt = 0
		
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line msg_parse.go:5706
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr64
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr64
			}
		default:
			goto tr64
		}
		goto st0
tr58:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st117
tr67:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line msg_parse.go:5743
		if 128 <= data[p] && data[p] <= 191 {
			goto tr64
		}
		goto st0
tr59:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st118
tr68:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line msg_parse.go:5771
		if 128 <= data[p] && data[p] <= 191 {
			goto tr67
		}
		goto st0
tr60:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st119
tr69:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line msg_parse.go:5799
		if 128 <= data[p] && data[p] <= 191 {
			goto tr68
		}
		goto st0
tr61:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st120
tr70:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line msg_parse.go:5827
		if 128 <= data[p] && data[p] <= 191 {
			goto tr69
		}
		goto st0
tr62:
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st121
tr71:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line msg_parse.go:5855
		if 128 <= data[p] && data[p] <= 191 {
			goto tr70
		}
		goto st0
tr63:
//line msg_parse.rl:63

			amt = 0
		
	goto st122
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
//line msg_parse.go:5871
		if data[p] == 10 {
			goto tr202
		}
		goto st0
tr202:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st123
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
//line msg_parse.go:5885
		switch data[p] {
		case 9:
			goto tr64
		case 32:
			goto tr64
		}
		goto st0
tr203:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st124
tr49:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line msg_parse.go:5924
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st46
		case 32:
			goto st46
		case 33:
			goto tr90
		case 37:
			goto tr90
		case 39:
			goto tr90
		case 42:
			goto tr90
		case 43:
			goto tr203
		case 58:
			goto st125
		case 126:
			goto tr90
		case 525:
			goto st47
		}
		switch {
		case _widec < 65:
			switch {
			case _widec > 46:
				if 48 <= _widec && _widec <= 57 {
					goto tr203
				}
			case _widec >= 45:
				goto tr203
			}
		case _widec > 90:
			switch {
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto tr203
				}
			case _widec >= 95:
				goto tr90
			}
		default:
			goto tr203
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 33:
			goto st126
		case 61:
			goto st126
		case 95:
			goto st126
		case 126:
			goto st126
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 58 {
				goto st126
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st126
			}
		default:
			goto st126
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr105
		case 32:
			goto tr105
		case 33:
			goto st126
		case 44:
			goto tr206
		case 59:
			goto tr107
		case 61:
			goto st126
		case 95:
			goto st126
		case 126:
			goto st126
		case 269:
			goto tr108
		case 525:
			goto tr109
		}
		switch {
		case _widec < 63:
			if 36 <= _widec && _widec <= 58 {
				goto st126
			}
		case _widec > 90:
			if 97 <= _widec && _widec <= 122 {
				goto st126
			}
		default:
			goto st126
		}
		goto st0
tr206:
//line msg_parse.rl:182

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:165

			// TODO(jart): Why does this fire multiple times?
			if addr != nil {
				**addrpp = addr
				*addrpp = &addr.Next
				addr = nil
			}
		
	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line msg_parse.go:6073
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr207
		case 32:
			goto tr207
		case 33:
			goto tr208
		case 34:
			goto tr47
		case 37:
			goto tr208
		case 39:
			goto tr208
		case 44:
			goto tr206
		case 47:
			goto st126
		case 59:
			goto tr107
		case 60:
			goto tr48
		case 95:
			goto tr208
		case 96:
			goto tr46
		case 126:
			goto tr208
		case 269:
			goto tr108
		case 525:
			goto tr210
		}
		switch {
		case _widec < 58:
			switch {
			case _widec > 41:
				if 42 <= _widec && _widec <= 57 {
					goto tr208
				}
			case _widec >= 36:
				goto st126
			}
		case _widec > 61:
			switch {
			case _widec < 65:
				if 63 <= _widec && _widec <= 64 {
					goto st126
				}
			case _widec > 90:
				if 97 <= _widec && _widec <= 122 {
					goto tr209
				}
			default:
				goto tr209
			}
		default:
			goto st126
		}
		goto st0
tr211:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st128
tr208:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st128
tr209:
//line msg_parse.rl:174

			addr = new(Addr)
		
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:63

			amt = 0
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line msg_parse.go:6186
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr182
		case 32:
			goto tr182
		case 33:
			goto tr211
		case 37:
			goto tr211
		case 39:
			goto tr211
		case 44:
			goto tr206
		case 47:
			goto st126
		case 58:
			goto st126
		case 59:
			goto tr107
		case 61:
			goto st126
		case 96:
			goto tr90
		case 126:
			goto tr211
		case 269:
			goto tr108
		case 525:
			goto tr184
		}
		switch {
		case _widec < 63:
			switch {
			case _widec > 41:
				if 42 <= _widec && _widec <= 57 {
					goto tr211
				}
			case _widec >= 36:
				goto st126
			}
		case _widec > 64:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr211
				}
			case _widec >= 65:
				goto tr211
			}
		default:
			goto st126
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr212
		case 32:
			goto tr212
		case 269:
			goto tr219
		case 525:
			goto tr220
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr214
				}
			case _widec >= 33:
				goto tr213
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr216
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr218
				}
			default:
				goto tr217
			}
		default:
			goto tr215
		}
		goto st0
tr212:
//line msg_parse.rl:59

			mark = p
		
	goto st130
tr222:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st130
tr213:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
//line msg_parse.go:6325
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st130
		case 32:
			goto st130
		case 269:
			goto st136
		case 525:
			goto st137
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr223
				}
			case _widec >= 33:
				goto tr222
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr225
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr227
				}
			default:
				goto tr226
			}
		default:
			goto tr224
		}
		goto st0
tr223:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st131
tr214:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line msg_parse.go:6393
		if 128 <= data[p] && data[p] <= 191 {
			goto tr222
		}
		goto st0
tr224:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st132
tr215:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line msg_parse.go:6421
		if 128 <= data[p] && data[p] <= 191 {
			goto tr223
		}
		goto st0
tr225:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st133
tr216:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line msg_parse.go:6449
		if 128 <= data[p] && data[p] <= 191 {
			goto tr224
		}
		goto st0
tr226:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st134
tr217:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line msg_parse.go:6477
		if 128 <= data[p] && data[p] <= 191 {
			goto tr225
		}
		goto st0
tr227:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st135
tr218:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st135
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
//line msg_parse.go:6505
		if 128 <= data[p] && data[p] <= 191 {
			goto tr226
		}
		goto st0
tr219:
//line msg_parse.rl:59

			mark = p
		
	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line msg_parse.go:6521
		if data[p] == 10 {
			goto tr230
		}
		goto st0
tr230:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:153
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
//line msg_parse.rl:129

			{goto st145 }
		
	goto st642
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
//line msg_parse.go:6551
		goto st0
tr220:
//line msg_parse.rl:59

			mark = p
		
	goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line msg_parse.go:6564
		if data[p] == 10 {
			goto tr231
		}
		goto st0
tr231:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:153
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
//line msg_parse.rl:129

			{goto st145 }
		
	goto st643
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
//line msg_parse.go:6594
		switch data[p] {
		case 9:
			goto st130
		case 32:
			goto st130
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 33:
			goto tr232
		case 37:
			goto tr232
		case 39:
			goto tr232
		case 126:
			goto tr232
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr232
				}
			case data[p] >= 42:
				goto tr232
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr232
				}
			case data[p] >= 65:
				goto tr232
			}
		default:
			goto tr232
		}
		goto st0
tr232:
//line msg_parse.rl:59

			mark = p
		
	goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line msg_parse.go:6651
		switch data[p] {
		case 9:
			goto tr233
		case 32:
			goto tr233
		case 33:
			goto st139
		case 37:
			goto st139
		case 39:
			goto st139
		case 58:
			goto tr235
		case 126:
			goto st139
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st139
				}
			case data[p] >= 42:
				goto st139
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st139
				}
			case data[p] >= 65:
				goto st139
			}
		default:
			goto st139
		}
		goto st0
tr233:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
	goto st140
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
//line msg_parse.go:6702
		switch data[p] {
		case 9:
			goto st140
		case 32:
			goto st140
		case 58:
			goto st141
		}
		goto st0
tr235:
//line msg_parse.rl:149

			name = string(data[mark:p])
		
	goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line msg_parse.go:6723
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st141
		case 32:
			goto st141
		case 269:
			goto tr238
		case 525:
			goto st142
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr238
			}
		default:
			goto tr238
		}
		goto st0
tr238:
//line msg_parse.rl:373
value=nil
//line msg_parse.rl:133

			p--

			{goto st129 }
		
	goto st644
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
//line msg_parse.go:6765
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		if data[p] == 10 {
			goto tr240
		}
		goto st0
tr240:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line msg_parse.go:6785
		switch data[p] {
		case 9:
			goto st144
		case 32:
			goto st144
		}
		goto st0
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
		switch data[p] {
		case 9:
			goto st144
		case 32:
			goto st144
		}
		goto tr238
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 13:
			goto st146
		case 65:
			goto tr244
		case 66:
			goto tr245
		case 67:
			goto tr246
		case 68:
			goto tr247
		case 69:
			goto tr248
		case 70:
			goto tr249
		case 73:
			goto tr250
		case 75:
			goto tr251
		case 76:
			goto st372
		case 77:
			goto tr253
		case 79:
			goto tr254
		case 80:
			goto tr255
		case 82:
			goto tr256
		case 83:
			goto tr257
		case 84:
			goto tr258
		case 85:
			goto tr259
		case 86:
			goto st598
		case 87:
			goto tr261
		case 97:
			goto tr244
		case 98:
			goto tr245
		case 99:
			goto tr246
		case 100:
			goto tr247
		case 101:
			goto tr248
		case 102:
			goto tr249
		case 105:
			goto tr250
		case 107:
			goto tr251
		case 108:
			goto st372
		case 109:
			goto tr253
		case 111:
			goto tr254
		case 112:
			goto tr255
		case 114:
			goto tr256
		case 115:
			goto tr257
		case 116:
			goto tr258
		case 117:
			goto tr259
		case 118:
			goto st598
		case 119:
			goto tr261
		}
		goto tr242
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		if data[p] == 10 {
			goto tr262
		}
		goto st0
tr387:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st645
tr262:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:55

			{p++; cs = 645; goto _out }
		
	goto st645
tr270:
//line msg_parse.rl:133

			p--

			{goto st129 }
		
	goto st645
tr421:
//line msg_parse.rl:124

			p--

			{goto st34 }
		
	goto st645
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
//line msg_parse.go:6933
		goto st0
tr244:
//line msg_parse.rl:59

			mark = p
		
	goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line msg_parse.go:6946
		switch data[p] {
		case 9:
			goto tr263
		case 32:
			goto tr263
		case 58:
			goto tr264
		case 67:
			goto st153
		case 76:
			goto st182
		case 85:
			goto st201
		case 99:
			goto st153
		case 108:
			goto st182
		case 117:
			goto st201
		}
		goto st0
tr263:
//line msg_parse.rl:334
value=&msg.AcceptContact
	goto st148
tr279:
//line msg_parse.rl:333
value=&msg.Accept
	goto st148
tr298:
//line msg_parse.rl:335
value=&msg.AcceptEncoding
	goto st148
tr307:
//line msg_parse.rl:336
value=&msg.AcceptLanguage
	goto st148
tr318:
//line msg_parse.rl:339
value=&msg.AlertInfo
	goto st148
tr322:
//line msg_parse.rl:337
value=&msg.Allow
	goto st148
tr331:
//line msg_parse.rl:338
value=&msg.AllowEvents
	goto st148
tr351:
//line msg_parse.rl:340
value=&msg.AuthenticationInfo
	goto st148
tr361:
//line msg_parse.rl:341
value=&msg.Authorization
	goto st148
tr363:
//line msg_parse.rl:358
value=&msg.ReferredBy
	goto st148
tr409:
//line msg_parse.rl:345
value=&msg.CallInfo
	goto st148
tr442:
//line msg_parse.rl:342
value=&msg.ContentDisposition
	goto st148
tr451:
//line msg_parse.rl:344
value=&msg.ContentEncoding
	goto st148
tr461:
//line msg_parse.rl:343
value=&msg.ContentLanguage
	goto st148
tr494:
//line msg_parse.rl:346
value=&msg.Date
	goto st148
tr507:
//line msg_parse.rl:347
value=&msg.ErrorInfo
	goto st148
tr512:
//line msg_parse.rl:348
value=&msg.Event
	goto st148
tr540:
//line msg_parse.rl:349
value=&msg.InReplyTo
	goto st148
tr542:
//line msg_parse.rl:363
value=&msg.Supported
	goto st148
tr579:
//line msg_parse.rl:351
value=&msg.MIMEVersion
	goto st148
tr606:
//line msg_parse.rl:352
value=&msg.Organization
	goto st148
tr636:
//line msg_parse.rl:353
value=&msg.Priority
	goto st148
tr655:
//line msg_parse.rl:354
value=&msg.ProxyAuthenticate
	goto st148
tr665:
//line msg_parse.rl:355
value=&msg.ProxyAuthorization
	goto st148
tr673:
//line msg_parse.rl:356
value=&msg.ProxyRequire
	goto st148
tr675:
//line msg_parse.rl:357
value=&msg.ReferTo
	goto st148
tr726:
//line msg_parse.rl:350
value=&msg.ReplyTo
	goto st148
tr732:
//line msg_parse.rl:359
value=&msg.Require
	goto st148
tr742:
//line msg_parse.rl:360
value=&msg.RetryAfter
	goto st148
tr749:
//line msg_parse.rl:362
value=&msg.Subject
	goto st148
tr757:
//line msg_parse.rl:361
value=&msg.Server
	goto st148
tr782:
//line msg_parse.rl:364
value=&msg.Timestamp
	goto st148
tr784:
//line msg_parse.rl:337
value=&msg.Allow
//line msg_parse.rl:338
value=&msg.AllowEvents
	goto st148
tr797:
//line msg_parse.rl:365
value=&msg.Unsupported
	goto st148
tr807:
//line msg_parse.rl:366
value=&msg.UserAgent
	goto st148
tr840:
//line msg_parse.rl:367
value=&msg.Warning
	goto st148
tr856:
//line msg_parse.rl:368
value=&msg.WWWAuthenticate
	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line msg_parse.go:7123
		switch data[p] {
		case 9:
			goto st148
		case 32:
			goto st148
		case 58:
			goto st149
		}
		goto st0
tr264:
//line msg_parse.rl:334
value=&msg.AcceptContact
	goto st149
tr281:
//line msg_parse.rl:333
value=&msg.Accept
	goto st149
tr299:
//line msg_parse.rl:335
value=&msg.AcceptEncoding
	goto st149
tr308:
//line msg_parse.rl:336
value=&msg.AcceptLanguage
	goto st149
tr319:
//line msg_parse.rl:339
value=&msg.AlertInfo
	goto st149
tr324:
//line msg_parse.rl:337
value=&msg.Allow
	goto st149
tr332:
//line msg_parse.rl:338
value=&msg.AllowEvents
	goto st149
tr352:
//line msg_parse.rl:340
value=&msg.AuthenticationInfo
	goto st149
tr362:
//line msg_parse.rl:341
value=&msg.Authorization
	goto st149
tr364:
//line msg_parse.rl:358
value=&msg.ReferredBy
	goto st149
tr410:
//line msg_parse.rl:345
value=&msg.CallInfo
	goto st149
tr443:
//line msg_parse.rl:342
value=&msg.ContentDisposition
	goto st149
tr452:
//line msg_parse.rl:344
value=&msg.ContentEncoding
	goto st149
tr462:
//line msg_parse.rl:343
value=&msg.ContentLanguage
	goto st149
tr495:
//line msg_parse.rl:346
value=&msg.Date
	goto st149
tr508:
//line msg_parse.rl:347
value=&msg.ErrorInfo
	goto st149
tr513:
//line msg_parse.rl:348
value=&msg.Event
	goto st149
tr541:
//line msg_parse.rl:349
value=&msg.InReplyTo
	goto st149
tr543:
//line msg_parse.rl:363
value=&msg.Supported
	goto st149
tr580:
//line msg_parse.rl:351
value=&msg.MIMEVersion
	goto st149
tr607:
//line msg_parse.rl:352
value=&msg.Organization
	goto st149
tr637:
//line msg_parse.rl:353
value=&msg.Priority
	goto st149
tr656:
//line msg_parse.rl:354
value=&msg.ProxyAuthenticate
	goto st149
tr666:
//line msg_parse.rl:355
value=&msg.ProxyAuthorization
	goto st149
tr674:
//line msg_parse.rl:356
value=&msg.ProxyRequire
	goto st149
tr676:
//line msg_parse.rl:357
value=&msg.ReferTo
	goto st149
tr727:
//line msg_parse.rl:350
value=&msg.ReplyTo
	goto st149
tr733:
//line msg_parse.rl:359
value=&msg.Require
	goto st149
tr743:
//line msg_parse.rl:360
value=&msg.RetryAfter
	goto st149
tr750:
//line msg_parse.rl:362
value=&msg.Subject
	goto st149
tr758:
//line msg_parse.rl:361
value=&msg.Server
	goto st149
tr783:
//line msg_parse.rl:364
value=&msg.Timestamp
	goto st149
tr785:
//line msg_parse.rl:337
value=&msg.Allow
//line msg_parse.rl:338
value=&msg.AllowEvents
	goto st149
tr798:
//line msg_parse.rl:365
value=&msg.Unsupported
	goto st149
tr808:
//line msg_parse.rl:366
value=&msg.UserAgent
	goto st149
tr841:
//line msg_parse.rl:367
value=&msg.Warning
	goto st149
tr857:
//line msg_parse.rl:368
value=&msg.WWWAuthenticate
	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line msg_parse.go:7288
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
		case 32:
			goto st149
		case 269:
			goto tr270
		case 525:
			goto st150
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr270
			}
		default:
			goto tr270
		}
		goto st0
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
		if data[p] == 10 {
			goto tr272
		}
		goto st0
tr272:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line msg_parse.go:7333
		switch data[p] {
		case 9:
			goto st152
		case 32:
			goto st152
		}
		goto st0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 9:
			goto st152
		case 32:
			goto st152
		}
		goto tr270
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 67:
			goto st154
		case 99:
			goto st154
		}
		goto tr274
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 69:
			goto st155
		case 101:
			goto st155
		}
		goto tr274
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 80:
			goto st156
		case 112:
			goto st156
		}
		goto tr274
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 84:
			goto st157
		case 116:
			goto st157
		}
		goto tr274
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 9:
			goto tr279
		case 32:
			goto tr279
		case 45:
			goto st158
		case 58:
			goto tr281
		}
		goto st0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 67:
			goto st159
		case 69:
			goto st166
		case 76:
			goto st174
		case 99:
			goto st159
		case 101:
			goto st166
		case 108:
			goto st174
		}
		goto tr274
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 79:
			goto st160
		case 111:
			goto st160
		}
		goto tr274
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 78:
			goto st161
		case 110:
			goto st161
		}
		goto tr274
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		switch data[p] {
		case 84:
			goto st162
		case 116:
			goto st162
		}
		goto tr274
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 65:
			goto st163
		case 97:
			goto st163
		}
		goto tr274
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 67:
			goto st164
		case 99:
			goto st164
		}
		goto tr274
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
		goto tr274
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
		switch data[p] {
		case 9:
			goto tr263
		case 32:
			goto tr263
		case 58:
			goto tr264
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 78:
			goto st167
		case 110:
			goto st167
		}
		goto tr274
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		switch data[p] {
		case 67:
			goto st168
		case 99:
			goto st168
		}
		goto tr274
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 79:
			goto st169
		case 111:
			goto st169
		}
		goto tr274
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 68:
			goto st170
		case 100:
			goto st170
		}
		goto tr274
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 73:
			goto st171
		case 105:
			goto st171
		}
		goto tr274
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 78:
			goto st172
		case 110:
			goto st172
		}
		goto tr274
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 71:
			goto st173
		case 103:
			goto st173
		}
		goto tr274
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 9:
			goto tr298
		case 32:
			goto tr298
		case 58:
			goto tr299
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 65:
			goto st175
		case 97:
			goto st175
		}
		goto tr274
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 78:
			goto st176
		case 110:
			goto st176
		}
		goto tr274
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 71:
			goto st177
		case 103:
			goto st177
		}
		goto tr274
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
		switch data[p] {
		case 85:
			goto st178
		case 117:
			goto st178
		}
		goto tr274
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 65:
			goto st179
		case 97:
			goto st179
		}
		goto tr274
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
		goto tr274
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 69:
			goto st181
		case 101:
			goto st181
		}
		goto tr274
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
		switch data[p] {
		case 9:
			goto tr307
		case 32:
			goto tr307
		case 58:
			goto tr308
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		switch data[p] {
		case 69:
			goto st183
		case 76:
			goto st191
		case 101:
			goto st183
		case 108:
			goto st191
		}
		goto tr274
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 82:
			goto st184
		case 114:
			goto st184
		}
		goto tr274
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
		switch data[p] {
		case 84:
			goto st185
		case 116:
			goto st185
		}
		goto tr274
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
		if data[p] == 45 {
			goto st186
		}
		goto tr274
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 73:
			goto st187
		case 105:
			goto st187
		}
		goto tr274
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
		switch data[p] {
		case 78:
			goto st188
		case 110:
			goto st188
		}
		goto tr274
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		switch data[p] {
		case 70:
			goto st189
		case 102:
			goto st189
		}
		goto tr274
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
		switch data[p] {
		case 79:
			goto st190
		case 111:
			goto st190
		}
		goto tr274
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
		switch data[p] {
		case 9:
			goto tr318
		case 32:
			goto tr318
		case 58:
			goto tr319
		}
		goto st0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 79:
			goto st192
		case 111:
			goto st192
		}
		goto tr274
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
		switch data[p] {
		case 87:
			goto st193
		case 119:
			goto st193
		}
		goto tr274
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
		switch data[p] {
		case 9:
			goto tr322
		case 32:
			goto tr322
		case 45:
			goto st194
		case 58:
			goto tr324
		}
		goto st0
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
		switch data[p] {
		case 69:
			goto st195
		case 101:
			goto st195
		}
		goto tr274
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 86:
			goto st196
		case 118:
			goto st196
		}
		goto tr274
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 69:
			goto st197
		case 101:
			goto st197
		}
		goto tr274
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 78:
			goto st198
		case 110:
			goto st198
		}
		goto tr274
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 84:
			goto st199
		case 116:
			goto st199
		}
		goto tr274
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 83:
			goto st200
		case 115:
			goto st200
		}
		goto tr274
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 9:
			goto tr331
		case 32:
			goto tr331
		case 58:
			goto tr332
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		switch data[p] {
		case 84:
			goto st202
		case 116:
			goto st202
		}
		goto tr274
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 72:
			goto st203
		case 104:
			goto st203
		}
		goto tr274
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 69:
			goto st204
		case 79:
			goto st219
		case 101:
			goto st204
		case 111:
			goto st219
		}
		goto tr274
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 78:
			goto st205
		case 110:
			goto st205
		}
		goto tr274
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		switch data[p] {
		case 84:
			goto st206
		case 116:
			goto st206
		}
		goto tr274
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		switch data[p] {
		case 73:
			goto st207
		case 105:
			goto st207
		}
		goto tr274
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		switch data[p] {
		case 67:
			goto st208
		case 99:
			goto st208
		}
		goto tr274
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		switch data[p] {
		case 65:
			goto st209
		case 97:
			goto st209
		}
		goto tr274
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
		switch data[p] {
		case 84:
			goto st210
		case 116:
			goto st210
		}
		goto tr274
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
		switch data[p] {
		case 73:
			goto st211
		case 105:
			goto st211
		}
		goto tr274
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 79:
			goto st212
		case 111:
			goto st212
		}
		goto tr274
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 78:
			goto st213
		case 110:
			goto st213
		}
		goto tr274
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		if data[p] == 45 {
			goto st214
		}
		goto tr274
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		switch data[p] {
		case 73:
			goto st215
		case 105:
			goto st215
		}
		goto tr274
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 78:
			goto st216
		case 110:
			goto st216
		}
		goto tr274
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 70:
			goto st217
		case 102:
			goto st217
		}
		goto tr274
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
		goto tr274
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 9:
			goto tr351
		case 32:
			goto tr351
		case 58:
			goto tr352
		}
		goto st0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 82:
			goto st220
		case 114:
			goto st220
		}
		goto tr274
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
		goto tr274
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
		switch data[p] {
		case 90:
			goto st222
		case 122:
			goto st222
		}
		goto tr274
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
		switch data[p] {
		case 65:
			goto st223
		case 97:
			goto st223
		}
		goto tr274
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
		switch data[p] {
		case 84:
			goto st224
		case 116:
			goto st224
		}
		goto tr274
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
		switch data[p] {
		case 73:
			goto st225
		case 105:
			goto st225
		}
		goto tr274
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
		switch data[p] {
		case 79:
			goto st226
		case 111:
			goto st226
		}
		goto tr274
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
		goto tr274
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
		switch data[p] {
		case 9:
			goto tr361
		case 32:
			goto tr361
		case 58:
			goto tr362
		}
		goto st0
tr245:
//line msg_parse.rl:59

			mark = p
		
	goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line msg_parse.go:8296
		switch data[p] {
		case 9:
			goto tr363
		case 32:
			goto tr363
		case 58:
			goto tr364
		}
		goto st0
tr246:
//line msg_parse.rl:59

			mark = p
		
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line msg_parse.go:8317
		switch data[p] {
		case 9:
			goto st230
		case 32:
			goto st230
		case 58:
			goto st231
		case 65:
			goto st243
		case 79:
			goto st259
		case 83:
			goto st314
		case 97:
			goto st243
		case 111:
			goto st259
		case 115:
			goto st314
		}
		goto tr242
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
		switch data[p] {
		case 9:
			goto st230
		case 32:
			goto st230
		case 58:
			goto st231
		}
		goto st0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st231
		case 32:
			goto st231
		case 269:
			goto tr376
		case 525:
			goto st240
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr371
				}
			case _widec >= 33:
				goto tr370
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr373
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr375
				}
			default:
				goto tr374
			}
		default:
			goto tr372
		}
		goto st0
tr379:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st232
tr370:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st232
	st232:
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
//line msg_parse.go:8425
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st232
		case 32:
			goto st232
		case 269:
			goto tr385
		case 525:
			goto tr386
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr380
				}
			case _widec >= 33:
				goto tr379
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr382
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr384
				}
			default:
				goto tr383
			}
		default:
			goto tr381
		}
		goto st0
tr380:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st233
tr371:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line msg_parse.go:8493
		if 128 <= data[p] && data[p] <= 191 {
			goto tr379
		}
		goto st0
tr381:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st234
tr372:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st234
	st234:
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
//line msg_parse.go:8521
		if 128 <= data[p] && data[p] <= 191 {
			goto tr380
		}
		goto st0
tr382:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st235
tr373:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st235
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
//line msg_parse.go:8549
		if 128 <= data[p] && data[p] <= 191 {
			goto tr381
		}
		goto st0
tr383:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st236
tr374:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st236
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
//line msg_parse.go:8577
		if 128 <= data[p] && data[p] <= 191 {
			goto tr382
		}
		goto st0
tr384:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st237
tr375:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st237
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
//line msg_parse.go:8605
		if 128 <= data[p] && data[p] <= 191 {
			goto tr383
		}
		goto st0
tr376:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:202

			ctype = string(data[mark:p])
		
	goto st238
tr385:
//line msg_parse.rl:202

			ctype = string(data[mark:p])
		
	goto st238
tr401:
//line msg_parse.rl:194

			msg.CallID = string(data[mark:p])
		
	goto st238
tr485:
//line msg_parse.rl:210

			msg.CSeqMethod = string(data[mark:p])
		
	goto st238
tr818:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:226

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st238
tr827:
//line msg_parse.rl:226

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st238
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
//line msg_parse.go:8663
		if data[p] == 10 {
			goto tr387
		}
		goto st0
tr391:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:202

			ctype = string(data[mark:p])
		
	goto st239
tr386:
//line msg_parse.rl:202

			ctype = string(data[mark:p])
		
	goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line msg_parse.go:8689
		if data[p] == 10 {
			goto tr388
		}
		goto st0
tr388:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st646
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
//line msg_parse.go:8707
		switch data[p] {
		case 9:
			goto st232
		case 32:
			goto st232
		}
		goto st0
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
		if data[p] == 10 {
			goto tr389
		}
		goto st0
tr389:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
//line msg_parse.go:8733
		switch data[p] {
		case 9:
			goto st242
		case 32:
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
			goto tr376
		case 525:
			goto tr391
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr371
				}
			case _widec >= 33:
				goto tr370
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr373
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr375
				}
			default:
				goto tr374
			}
		default:
			goto tr372
		}
		goto st0
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 76:
			goto st244
		case 108:
			goto st244
		}
		goto tr242
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
		switch data[p] {
		case 76:
			goto st245
		case 108:
			goto st245
		}
		goto tr242
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
		if data[p] == 45 {
			goto st246
		}
		goto tr242
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
		switch data[p] {
		case 73:
			goto st247
		case 105:
			goto st247
		}
		goto tr242
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
		switch data[p] {
		case 68:
			goto st248
		case 78:
			goto st256
		case 100:
			goto st248
		case 110:
			goto st256
		}
		goto tr242
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 9:
			goto st248
		case 32:
			goto st248
		case 58:
			goto st249
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st249
		case 32:
			goto st249
		case 37:
			goto tr399
		case 60:
			goto tr399
		case 525:
			goto st253
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto tr399
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto tr399
				}
			default:
				goto tr399
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto tr399
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto tr399
				}
			default:
				goto tr399
			}
		default:
			goto tr399
		}
		goto st0
tr399:
//line msg_parse.rl:59

			mark = p
		
	goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
//line msg_parse.go:8931
		switch data[p] {
		case 13:
			goto tr401
		case 37:
			goto st250
		case 60:
			goto st250
		case 64:
			goto st251
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 34:
				if 39 <= data[p] && data[p] <= 43 {
					goto st250
				}
			case data[p] >= 33:
				goto st250
			}
		case data[p] > 58:
			switch {
			case data[p] < 95:
				if 62 <= data[p] && data[p] <= 93 {
					goto st250
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st250
				}
			default:
				goto st250
			}
		default:
			goto st250
		}
		goto st0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 37:
			goto st252
		case 60:
			goto st252
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st252
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st252
				}
			default:
				goto st252
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st252
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st252
				}
			default:
				goto st252
			}
		default:
			goto st252
		}
		goto st0
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
		switch data[p] {
		case 13:
			goto tr401
		case 37:
			goto st252
		case 60:
			goto st252
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st252
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st252
				}
			default:
				goto st252
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st252
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st252
				}
			default:
				goto st252
			}
		default:
			goto st252
		}
		goto st0
	st253:
		if p++; p == pe {
			goto _test_eof253
		}
	st_case_253:
		if data[p] == 10 {
			goto tr405
		}
		goto st0
tr405:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
//line msg_parse.go:9073
		switch data[p] {
		case 9:
			goto st255
		case 32:
			goto st255
		}
		goto st0
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
		switch data[p] {
		case 9:
			goto st255
		case 32:
			goto st255
		case 37:
			goto tr399
		case 60:
			goto tr399
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto tr399
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto tr399
				}
			default:
				goto tr399
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto tr399
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto tr399
				}
			default:
				goto tr399
			}
		default:
			goto tr399
		}
		goto st0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
		switch data[p] {
		case 70:
			goto st257
		case 102:
			goto st257
		}
		goto tr274
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
		switch data[p] {
		case 79:
			goto st258
		case 111:
			goto st258
		}
		goto tr274
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		switch data[p] {
		case 9:
			goto tr409
		case 32:
			goto tr409
		case 58:
			goto tr410
		}
		goto st0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		switch data[p] {
		case 78:
			goto st260
		case 110:
			goto st260
		}
		goto tr242
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
		switch data[p] {
		case 84:
			goto st261
		case 116:
			goto st261
		}
		goto tr242
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 65:
			goto st262
		case 69:
			goto st270
		case 97:
			goto st262
		case 101:
			goto st270
		}
		goto tr242
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 67:
			goto st263
		case 99:
			goto st263
		}
		goto tr274
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 84:
			goto st264
		case 116:
			goto st264
		}
		goto tr274
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 9:
			goto tr417
		case 32:
			goto tr417
		case 58:
			goto tr418
		}
		goto st0
tr417:
//line msg_parse.rl:314
addrpp=&contactp
	goto st265
tr525:
//line msg_parse.rl:315
addrpp=&fromp
	goto st265
tr627:
//line msg_parse.rl:316
addrpp=&paidp
	goto st265
tr694:
//line msg_parse.rl:317
addrpp=&rroutep
	goto st265
tr719:
//line msg_parse.rl:318
addrpp=&rpidp
	goto st265
tr747:
//line msg_parse.rl:319
addrpp=&routep
	goto st265
tr771:
//line msg_parse.rl:320
addrpp=&top
	goto st265
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
//line msg_parse.go:9276
		switch data[p] {
		case 9:
			goto st265
		case 32:
			goto st265
		case 58:
			goto st266
		}
		goto st0
tr418:
//line msg_parse.rl:314
addrpp=&contactp
	goto st266
tr526:
//line msg_parse.rl:315
addrpp=&fromp
	goto st266
tr628:
//line msg_parse.rl:316
addrpp=&paidp
	goto st266
tr695:
//line msg_parse.rl:317
addrpp=&rroutep
	goto st266
tr720:
//line msg_parse.rl:318
addrpp=&rpidp
	goto st266
tr748:
//line msg_parse.rl:319
addrpp=&routep
	goto st266
tr772:
//line msg_parse.rl:320
addrpp=&top
	goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line msg_parse.go:9319
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st266
		case 32:
			goto st266
		case 269:
			goto tr421
		case 525:
			goto st267
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr421
			}
		default:
			goto tr421
		}
		goto st0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		if data[p] == 10 {
			goto tr423
		}
		goto st0
tr423:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st268
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
//line msg_parse.go:9364
		switch data[p] {
		case 9:
			goto st269
		case 32:
			goto st269
		}
		goto st0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		switch data[p] {
		case 9:
			goto st269
		case 32:
			goto st269
		}
		goto tr421
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 78:
			goto st271
		case 110:
			goto st271
		}
		goto tr242
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 84:
			goto st272
		case 116:
			goto st272
		}
		goto tr242
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		if data[p] == 45 {
			goto st273
		}
		goto tr242
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
		switch data[p] {
		case 68:
			goto st274
		case 69:
			goto st285
		case 76:
			goto st293
		case 84:
			goto st311
		case 100:
			goto st274
		case 101:
			goto st285
		case 108:
			goto st293
		case 116:
			goto st311
		}
		goto tr242
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
		switch data[p] {
		case 73:
			goto st275
		case 105:
			goto st275
		}
		goto tr274
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 83:
			goto st276
		case 115:
			goto st276
		}
		goto tr274
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 80:
			goto st277
		case 112:
			goto st277
		}
		goto tr274
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 79:
			goto st278
		case 111:
			goto st278
		}
		goto tr274
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
		goto tr274
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		switch data[p] {
		case 73:
			goto st280
		case 105:
			goto st280
		}
		goto tr274
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		switch data[p] {
		case 84:
			goto st281
		case 116:
			goto st281
		}
		goto tr274
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 73:
			goto st282
		case 105:
			goto st282
		}
		goto tr274
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 79:
			goto st283
		case 111:
			goto st283
		}
		goto tr274
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		switch data[p] {
		case 78:
			goto st284
		case 110:
			goto st284
		}
		goto tr274
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 9:
			goto tr442
		case 32:
			goto tr442
		case 58:
			goto tr443
		}
		goto st0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 78:
			goto st286
		case 110:
			goto st286
		}
		goto tr274
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		switch data[p] {
		case 67:
			goto st287
		case 99:
			goto st287
		}
		goto tr274
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		switch data[p] {
		case 79:
			goto st288
		case 111:
			goto st288
		}
		goto tr274
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		switch data[p] {
		case 68:
			goto st289
		case 100:
			goto st289
		}
		goto tr274
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
		switch data[p] {
		case 73:
			goto st290
		case 105:
			goto st290
		}
		goto tr274
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		switch data[p] {
		case 78:
			goto st291
		case 110:
			goto st291
		}
		goto tr274
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		switch data[p] {
		case 71:
			goto st292
		case 103:
			goto st292
		}
		goto tr274
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
		switch data[p] {
		case 9:
			goto tr451
		case 32:
			goto tr451
		case 58:
			goto tr452
		}
		goto st0
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		switch data[p] {
		case 65:
			goto st294
		case 69:
			goto st301
		case 97:
			goto st294
		case 101:
			goto st301
		}
		goto tr242
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
		goto tr274
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 71:
			goto st296
		case 103:
			goto st296
		}
		goto tr274
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 85:
			goto st297
		case 117:
			goto st297
		}
		goto tr274
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		switch data[p] {
		case 65:
			goto st298
		case 97:
			goto st298
		}
		goto tr274
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 71:
			goto st299
		case 103:
			goto st299
		}
		goto tr274
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 69:
			goto st300
		case 101:
			goto st300
		}
		goto tr274
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		switch data[p] {
		case 9:
			goto tr461
		case 32:
			goto tr461
		case 58:
			goto tr462
		}
		goto st0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		switch data[p] {
		case 78:
			goto st302
		case 110:
			goto st302
		}
		goto tr463
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 71:
			goto st303
		case 103:
			goto st303
		}
		goto tr463
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 84:
			goto st304
		case 116:
			goto st304
		}
		goto tr463
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 72:
			goto st305
		case 104:
			goto st305
		}
		goto tr463
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		switch data[p] {
		case 9:
			goto st305
		case 32:
			goto st305
		case 58:
			goto st306
		}
		goto st0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st306
		case 32:
			goto st306
		case 525:
			goto st308
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr469
		}
		goto st0
tr469:
//line msg_parse.rl:324
clen=0
//line msg_parse.rl:198

			clen = clen * 10 + (int(data[p]) - 0x30)
		
	goto st307
tr472:
//line msg_parse.rl:198

			clen = clen * 10 + (int(data[p]) - 0x30)
		
	goto st307
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
//line msg_parse.go:9880
		if data[p] == 13 {
			goto st238
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr472
		}
		goto st0
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		if data[p] == 10 {
			goto tr473
		}
		goto st0
tr473:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st309
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
//line msg_parse.go:9906
		switch data[p] {
		case 9:
			goto st310
		case 32:
			goto st310
		}
		goto st0
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 9:
			goto st310
		case 32:
			goto st310
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr469
		}
		goto st0
	st311:
		if p++; p == pe {
			goto _test_eof311
		}
	st_case_311:
		switch data[p] {
		case 89:
			goto st312
		case 121:
			goto st312
		}
		goto tr463
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 80:
			goto st313
		case 112:
			goto st313
		}
		goto tr463
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 69:
			goto st230
		case 101:
			goto st230
		}
		goto tr463
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 69:
			goto st315
		case 101:
			goto st315
		}
		goto tr463
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 81:
			goto st316
		case 113:
			goto st316
		}
		goto tr463
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
		switch data[p] {
		case 9:
			goto st316
		case 32:
			goto st316
		case 58:
			goto st317
		}
		goto st0
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st317
		case 32:
			goto st317
		case 525:
			goto st324
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr480
		}
		goto st0
tr480:
//line msg_parse.rl:206

			msg.CSeq = msg.CSeq * 10 + (int(data[p]) - 0x30)
		
	goto st318
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
//line msg_parse.go:10038
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st319
		case 32:
			goto st319
		case 525:
			goto st321
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr480
		}
		goto st0
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st319
		case 32:
			goto st319
		case 33:
			goto tr484
		case 37:
			goto tr484
		case 39:
			goto tr484
		case 126:
			goto tr484
		case 525:
			goto st321
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr484
				}
			case _widec >= 42:
				goto tr484
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr484
				}
			case _widec >= 65:
				goto tr484
			}
		default:
			goto tr484
		}
		goto st0
tr484:
//line msg_parse.rl:59

			mark = p
		
	goto st320
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
//line msg_parse.go:10120
		switch data[p] {
		case 13:
			goto tr485
		case 33:
			goto st320
		case 37:
			goto st320
		case 39:
			goto st320
		case 126:
			goto st320
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st320
				}
			case data[p] >= 42:
				goto st320
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st320
				}
			case data[p] >= 65:
				goto st320
			}
		default:
			goto st320
		}
		goto st0
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		if data[p] == 10 {
			goto tr487
		}
		goto st0
tr487:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st322
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
//line msg_parse.go:10174
		switch data[p] {
		case 9:
			goto st323
		case 32:
			goto st323
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 9:
			goto st323
		case 32:
			goto st323
		case 33:
			goto tr484
		case 37:
			goto tr484
		case 39:
			goto tr484
		case 126:
			goto tr484
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr484
				}
			case data[p] >= 42:
				goto tr484
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr484
				}
			case data[p] >= 65:
				goto tr484
			}
		default:
			goto tr484
		}
		goto st0
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		if data[p] == 10 {
			goto tr489
		}
		goto st0
tr489:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st325
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
//line msg_parse.go:10242
		switch data[p] {
		case 9:
			goto st326
		case 32:
			goto st326
		}
		goto st0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 9:
			goto st326
		case 32:
			goto st326
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr480
		}
		goto st0
tr247:
//line msg_parse.rl:59

			mark = p
		
	goto st327
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
//line msg_parse.go:10276
		switch data[p] {
		case 65:
			goto st328
		case 97:
			goto st328
		}
		goto tr274
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
		goto tr274
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
		goto tr274
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 9:
			goto tr494
		case 32:
			goto tr494
		case 58:
			goto tr495
		}
		goto st0
tr248:
//line msg_parse.rl:59

			mark = p
		
	goto st331
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
//line msg_parse.go:10333
		switch data[p] {
		case 9:
			goto tr451
		case 32:
			goto tr451
		case 58:
			goto tr452
		case 82:
			goto st332
		case 86:
			goto st341
		case 88:
			goto st345
		case 114:
			goto st332
		case 118:
			goto st341
		case 120:
			goto st345
		}
		goto tr463
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 82:
			goto st333
		case 114:
			goto st333
		}
		goto tr274
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 79:
			goto st334
		case 111:
			goto st334
		}
		goto tr274
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 82:
			goto st335
		case 114:
			goto st335
		}
		goto tr274
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
		if data[p] == 45 {
			goto st336
		}
		goto tr274
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 73:
			goto st337
		case 105:
			goto st337
		}
		goto tr274
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		switch data[p] {
		case 78:
			goto st338
		case 110:
			goto st338
		}
		goto tr274
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
	st_case_338:
		switch data[p] {
		case 70:
			goto st339
		case 102:
			goto st339
		}
		goto tr274
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 79:
			goto st340
		case 111:
			goto st340
		}
		goto tr274
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 9:
			goto tr507
		case 32:
			goto tr507
		case 58:
			goto tr508
		}
		goto st0
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		switch data[p] {
		case 69:
			goto st342
		case 101:
			goto st342
		}
		goto tr274
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 78:
			goto st343
		case 110:
			goto st343
		}
		goto tr274
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		switch data[p] {
		case 84:
			goto st344
		case 116:
			goto st344
		}
		goto tr274
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		switch data[p] {
		case 9:
			goto tr512
		case 32:
			goto tr512
		case 58:
			goto tr513
		}
		goto st0
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		switch data[p] {
		case 80:
			goto st346
		case 112:
			goto st346
		}
		goto tr463
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		switch data[p] {
		case 73:
			goto st347
		case 105:
			goto st347
		}
		goto tr463
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 82:
			goto st348
		case 114:
			goto st348
		}
		goto tr463
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 69:
			goto st349
		case 101:
			goto st349
		}
		goto tr463
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 83:
			goto st350
		case 115:
			goto st350
		}
		goto tr463
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 9:
			goto st350
		case 32:
			goto st350
		case 58:
			goto st351
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st351
		case 32:
			goto st351
		case 525:
			goto st353
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr520
		}
		goto st0
tr520:
//line msg_parse.rl:327
msg.Expires=0
//line msg_parse.rl:214

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
	goto st352
tr522:
//line msg_parse.rl:214

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
	goto st352
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
//line msg_parse.go:10629
		if data[p] == 13 {
			goto st238
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr522
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		if data[p] == 10 {
			goto tr523
		}
		goto st0
tr523:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st354
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
//line msg_parse.go:10655
		switch data[p] {
		case 9:
			goto st355
		case 32:
			goto st355
		}
		goto st0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 9:
			goto st355
		case 32:
			goto st355
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr520
		}
		goto st0
tr249:
//line msg_parse.rl:59

			mark = p
		
	goto st356
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
//line msg_parse.go:10689
		switch data[p] {
		case 9:
			goto tr525
		case 32:
			goto tr525
		case 58:
			goto tr526
		case 82:
			goto st357
		case 114:
			goto st357
		}
		goto st0
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		switch data[p] {
		case 79:
			goto st358
		case 111:
			goto st358
		}
		goto tr274
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 77:
			goto st359
		case 109:
			goto st359
		}
		goto tr274
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		switch data[p] {
		case 9:
			goto tr525
		case 32:
			goto tr525
		case 58:
			goto tr526
		}
		goto st0
tr250:
//line msg_parse.rl:59

			mark = p
		
	goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line msg_parse.go:10752
		switch data[p] {
		case 9:
			goto st248
		case 32:
			goto st248
		case 58:
			goto st249
		case 78:
			goto st361
		case 110:
			goto st361
		}
		goto tr274
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		if data[p] == 45 {
			goto st362
		}
		goto tr274
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
		switch data[p] {
		case 82:
			goto st363
		case 114:
			goto st363
		}
		goto tr274
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
		switch data[p] {
		case 69:
			goto st364
		case 101:
			goto st364
		}
		goto tr274
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 80:
			goto st365
		case 112:
			goto st365
		}
		goto tr274
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 76:
			goto st366
		case 108:
			goto st366
		}
		goto tr274
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 89:
			goto st367
		case 121:
			goto st367
		}
		goto tr274
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 45 {
			goto st368
		}
		goto tr274
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 84:
			goto st369
		case 116:
			goto st369
		}
		goto tr274
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
		switch data[p] {
		case 79:
			goto st370
		case 111:
			goto st370
		}
		goto tr274
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 9:
			goto tr540
		case 32:
			goto tr540
		case 58:
			goto tr541
		}
		goto st0
tr251:
//line msg_parse.rl:59

			mark = p
		
	goto st371
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
//line msg_parse.go:10893
		switch data[p] {
		case 9:
			goto tr542
		case 32:
			goto tr542
		case 58:
			goto tr543
		}
		goto st0
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		switch data[p] {
		case 9:
			goto st372
		case 32:
			goto st372
		case 58:
			goto st373
		}
		goto st0
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st373
		case 32:
			goto st373
		case 525:
			goto st375
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr545
		}
		goto st0
tr545:
//line msg_parse.rl:324
clen=0
//line msg_parse.rl:198

			clen = clen * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:327
msg.Expires=0
//line msg_parse.rl:214

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:328
msg.MaxForwards=0
//line msg_parse.rl:218

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:329
msg.MinExpires=0
//line msg_parse.rl:222

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st374
tr547:
//line msg_parse.rl:198

			clen = clen * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:214

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:218

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:222

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
//line msg_parse.go:10990
		if data[p] == 13 {
			goto st238
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr547
		}
		goto st0
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
		if data[p] == 10 {
			goto tr548
		}
		goto st0
tr548:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st376
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
//line msg_parse.go:11016
		switch data[p] {
		case 9:
			goto st377
		case 32:
			goto st377
		}
		goto st0
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		switch data[p] {
		case 9:
			goto st377
		case 32:
			goto st377
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr545
		}
		goto st0
tr253:
//line msg_parse.rl:59

			mark = p
		
	goto st378
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
//line msg_parse.go:11050
		switch data[p] {
		case 9:
			goto tr417
		case 32:
			goto tr417
		case 58:
			goto tr418
		case 65:
			goto st379
		case 73:
			goto st395
		case 97:
			goto st379
		case 105:
			goto st395
		}
		goto tr242
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 88:
			goto st380
		case 120:
			goto st380
		}
		goto tr463
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		if data[p] == 45 {
			goto st381
		}
		goto tr463
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 70:
			goto st382
		case 102:
			goto st382
		}
		goto tr463
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 79:
			goto st383
		case 111:
			goto st383
		}
		goto tr463
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
		goto tr463
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 87:
			goto st385
		case 119:
			goto st385
		}
		goto tr463
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 65:
			goto st386
		case 97:
			goto st386
		}
		goto tr463
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 82:
			goto st387
		case 114:
			goto st387
		}
		goto tr463
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
		switch data[p] {
		case 68:
			goto st388
		case 100:
			goto st388
		}
		goto tr463
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		switch data[p] {
		case 83:
			goto st389
		case 115:
			goto st389
		}
		goto tr463
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 9:
			goto st389
		case 32:
			goto st389
		case 58:
			goto st390
		}
		goto st0
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st390
		case 32:
			goto st390
		case 525:
			goto st392
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr563
		}
		goto st0
tr563:
//line msg_parse.rl:328
msg.MaxForwards=0
//line msg_parse.rl:218

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
	goto st391
tr565:
//line msg_parse.rl:218

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
	goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line msg_parse.go:11242
		if data[p] == 13 {
			goto st238
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr565
		}
		goto st0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		if data[p] == 10 {
			goto tr566
		}
		goto st0
tr566:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st393
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
//line msg_parse.go:11268
		switch data[p] {
		case 9:
			goto st394
		case 32:
			goto st394
		}
		goto st0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		switch data[p] {
		case 9:
			goto st394
		case 32:
			goto st394
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr563
		}
		goto st0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 77:
			goto st396
		case 78:
			goto st406
		case 109:
			goto st396
		case 110:
			goto st406
		}
		goto tr242
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
		goto tr274
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		if data[p] == 45 {
			goto st398
		}
		goto tr274
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 86:
			goto st399
		case 118:
			goto st399
		}
		goto tr274
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 69:
			goto st400
		case 101:
			goto st400
		}
		goto tr274
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 82:
			goto st401
		case 114:
			goto st401
		}
		goto tr274
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		switch data[p] {
		case 83:
			goto st402
		case 115:
			goto st402
		}
		goto tr274
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
	st_case_402:
		switch data[p] {
		case 73:
			goto st403
		case 105:
			goto st403
		}
		goto tr274
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 79:
			goto st404
		case 111:
			goto st404
		}
		goto tr274
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 78:
			goto st405
		case 110:
			goto st405
		}
		goto tr274
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 9:
			goto tr579
		case 32:
			goto tr579
		case 58:
			goto tr580
		}
		goto st0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		if data[p] == 45 {
			goto st407
		}
		goto tr463
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 69:
			goto st408
		case 101:
			goto st408
		}
		goto tr463
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 88:
			goto st409
		case 120:
			goto st409
		}
		goto tr463
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		switch data[p] {
		case 80:
			goto st410
		case 112:
			goto st410
		}
		goto tr463
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
		switch data[p] {
		case 73:
			goto st411
		case 105:
			goto st411
		}
		goto tr463
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 82:
			goto st412
		case 114:
			goto st412
		}
		goto tr463
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		switch data[p] {
		case 69:
			goto st413
		case 101:
			goto st413
		}
		goto tr463
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
		switch data[p] {
		case 83:
			goto st414
		case 115:
			goto st414
		}
		goto tr463
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 9:
			goto st414
		case 32:
			goto st414
		case 58:
			goto st415
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st415
		case 32:
			goto st415
		case 525:
			goto st417
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr590
		}
		goto st0
tr590:
//line msg_parse.rl:329
msg.MinExpires=0
//line msg_parse.rl:222

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st416
tr592:
//line msg_parse.rl:222

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
//line msg_parse.go:11576
		if data[p] == 13 {
			goto st238
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr592
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		if data[p] == 10 {
			goto tr593
		}
		goto st0
tr593:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st418
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
//line msg_parse.go:11602
		switch data[p] {
		case 9:
			goto st419
		case 32:
			goto st419
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 9:
			goto st419
		case 32:
			goto st419
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr590
		}
		goto st0
tr254:
//line msg_parse.rl:59

			mark = p
		
	goto st420
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
//line msg_parse.go:11636
		switch data[p] {
		case 9:
			goto tr512
		case 32:
			goto tr512
		case 58:
			goto tr513
		case 82:
			goto st421
		case 114:
			goto st421
		}
		goto st0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 71:
			goto st422
		case 103:
			goto st422
		}
		goto tr274
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 65:
			goto st423
		case 97:
			goto st423
		}
		goto tr274
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		switch data[p] {
		case 78:
			goto st424
		case 110:
			goto st424
		}
		goto tr274
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 73:
			goto st425
		case 105:
			goto st425
		}
		goto tr274
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 90:
			goto st426
		case 122:
			goto st426
		}
		goto tr274
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
		switch data[p] {
		case 65:
			goto st427
		case 97:
			goto st427
		}
		goto tr274
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 84:
			goto st428
		case 116:
			goto st428
		}
		goto tr274
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 73:
			goto st429
		case 105:
			goto st429
		}
		goto tr274
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
		switch data[p] {
		case 79:
			goto st430
		case 111:
			goto st430
		}
		goto tr274
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
		switch data[p] {
		case 78:
			goto st431
		case 110:
			goto st431
		}
		goto tr274
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 9:
			goto tr606
		case 32:
			goto tr606
		case 58:
			goto tr607
		}
		goto st0
tr255:
//line msg_parse.rl:59

			mark = p
		
	goto st432
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
//line msg_parse.go:11795
		switch data[p] {
		case 45:
			goto st433
		case 82:
			goto st451
		case 114:
			goto st451
		}
		goto tr274
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 65:
			goto st434
		case 97:
			goto st434
		}
		goto tr274
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 83:
			goto st435
		case 115:
			goto st435
		}
		goto tr274
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 83:
			goto st436
		case 115:
			goto st436
		}
		goto tr274
	st436:
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		switch data[p] {
		case 69:
			goto st437
		case 101:
			goto st437
		}
		goto tr274
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 82:
			goto st438
		case 114:
			goto st438
		}
		goto tr274
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		switch data[p] {
		case 84:
			goto st439
		case 116:
			goto st439
		}
		goto tr274
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
		switch data[p] {
		case 69:
			goto st440
		case 101:
			goto st440
		}
		goto tr274
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
		switch data[p] {
		case 68:
			goto st441
		case 100:
			goto st441
		}
		goto tr274
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		if data[p] == 45 {
			goto st442
		}
		goto tr274
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 73:
			goto st443
		case 105:
			goto st443
		}
		goto tr274
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
		switch data[p] {
		case 68:
			goto st444
		case 100:
			goto st444
		}
		goto tr274
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
		switch data[p] {
		case 69:
			goto st445
		case 101:
			goto st445
		}
		goto tr274
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 78:
			goto st446
		case 110:
			goto st446
		}
		goto tr274
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 84:
			goto st447
		case 116:
			goto st447
		}
		goto tr274
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
		switch data[p] {
		case 73:
			goto st448
		case 105:
			goto st448
		}
		goto tr274
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 84:
			goto st449
		case 116:
			goto st449
		}
		goto tr274
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 89:
			goto st450
		case 121:
			goto st450
		}
		goto tr274
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 9:
			goto tr627
		case 32:
			goto tr627
		case 58:
			goto tr628
		}
		goto st0
	st451:
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		switch data[p] {
		case 73:
			goto st452
		case 79:
			goto st458
		case 105:
			goto st452
		case 111:
			goto st458
		}
		goto tr274
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 79:
			goto st453
		case 111:
			goto st453
		}
		goto tr274
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		switch data[p] {
		case 82:
			goto st454
		case 114:
			goto st454
		}
		goto tr274
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 73:
			goto st455
		case 105:
			goto st455
		}
		goto tr274
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
		goto tr274
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 89:
			goto st457
		case 121:
			goto st457
		}
		goto tr274
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 9:
			goto tr636
		case 32:
			goto tr636
		case 58:
			goto tr637
		}
		goto st0
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		switch data[p] {
		case 88:
			goto st459
		case 120:
			goto st459
		}
		goto tr274
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 89:
			goto st460
		case 121:
			goto st460
		}
		goto tr274
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		if data[p] == 45 {
			goto st461
		}
		goto tr274
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 65:
			goto st462
		case 82:
			goto st483
		case 97:
			goto st462
		case 114:
			goto st483
		}
		goto tr274
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
		switch data[p] {
		case 85:
			goto st463
		case 117:
			goto st463
		}
		goto tr274
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
		switch data[p] {
		case 84:
			goto st464
		case 116:
			goto st464
		}
		goto tr274
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		switch data[p] {
		case 72:
			goto st465
		case 104:
			goto st465
		}
		goto tr274
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		switch data[p] {
		case 69:
			goto st466
		case 79:
			goto st474
		case 101:
			goto st466
		case 111:
			goto st474
		}
		goto tr274
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 78:
			goto st467
		case 110:
			goto st467
		}
		goto tr274
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		switch data[p] {
		case 84:
			goto st468
		case 116:
			goto st468
		}
		goto tr274
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
		switch data[p] {
		case 73:
			goto st469
		case 105:
			goto st469
		}
		goto tr274
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 67:
			goto st470
		case 99:
			goto st470
		}
		goto tr274
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		switch data[p] {
		case 65:
			goto st471
		case 97:
			goto st471
		}
		goto tr274
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 84:
			goto st472
		case 116:
			goto st472
		}
		goto tr274
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
		goto tr274
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 9:
			goto tr655
		case 32:
			goto tr655
		case 58:
			goto tr656
		}
		goto st0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 82:
			goto st475
		case 114:
			goto st475
		}
		goto tr274
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 73:
			goto st476
		case 105:
			goto st476
		}
		goto tr274
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		switch data[p] {
		case 90:
			goto st477
		case 122:
			goto st477
		}
		goto tr274
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 65:
			goto st478
		case 97:
			goto st478
		}
		goto tr274
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 84:
			goto st479
		case 116:
			goto st479
		}
		goto tr274
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 73:
			goto st480
		case 105:
			goto st480
		}
		goto tr274
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		switch data[p] {
		case 79:
			goto st481
		case 111:
			goto st481
		}
		goto tr274
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 78:
			goto st482
		case 110:
			goto st482
		}
		goto tr274
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		switch data[p] {
		case 9:
			goto tr665
		case 32:
			goto tr665
		case 58:
			goto tr666
		}
		goto st0
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
		switch data[p] {
		case 69:
			goto st484
		case 101:
			goto st484
		}
		goto tr274
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		switch data[p] {
		case 81:
			goto st485
		case 113:
			goto st485
		}
		goto tr274
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		switch data[p] {
		case 85:
			goto st486
		case 117:
			goto st486
		}
		goto tr274
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		switch data[p] {
		case 73:
			goto st487
		case 105:
			goto st487
		}
		goto tr274
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		switch data[p] {
		case 82:
			goto st488
		case 114:
			goto st488
		}
		goto tr274
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		switch data[p] {
		case 69:
			goto st489
		case 101:
			goto st489
		}
		goto tr274
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		switch data[p] {
		case 9:
			goto tr673
		case 32:
			goto tr673
		case 58:
			goto tr674
		}
		goto st0
tr256:
//line msg_parse.rl:59

			mark = p
		
	goto st490
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
//line msg_parse.go:12516
		switch data[p] {
		case 9:
			goto tr675
		case 32:
			goto tr675
		case 58:
			goto tr676
		case 69:
			goto st491
		case 79:
			goto st546
		case 101:
			goto st491
		case 111:
			goto st546
		}
		goto tr274
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 67:
			goto st492
		case 70:
			goto st502
		case 77:
			goto st513
		case 80:
			goto st526
		case 81:
			goto st532
		case 84:
			goto st537
		case 99:
			goto st492
		case 102:
			goto st502
		case 109:
			goto st513
		case 112:
			goto st526
		case 113:
			goto st532
		case 116:
			goto st537
		}
		goto tr274
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		switch data[p] {
		case 79:
			goto st493
		case 111:
			goto st493
		}
		goto tr274
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		switch data[p] {
		case 82:
			goto st494
		case 114:
			goto st494
		}
		goto tr274
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 68:
			goto st495
		case 100:
			goto st495
		}
		goto tr274
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		if data[p] == 45 {
			goto st496
		}
		goto tr274
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		switch data[p] {
		case 82:
			goto st497
		case 114:
			goto st497
		}
		goto tr274
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 79:
			goto st498
		case 111:
			goto st498
		}
		goto tr274
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 85:
			goto st499
		case 117:
			goto st499
		}
		goto tr274
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 84:
			goto st500
		case 116:
			goto st500
		}
		goto tr274
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		switch data[p] {
		case 69:
			goto st501
		case 101:
			goto st501
		}
		goto tr274
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 9:
			goto tr694
		case 32:
			goto tr694
		case 58:
			goto tr695
		}
		goto st0
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		switch data[p] {
		case 69:
			goto st503
		case 101:
			goto st503
		}
		goto tr274
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		switch data[p] {
		case 82:
			goto st504
		case 114:
			goto st504
		}
		goto tr274
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 45:
			goto st505
		case 82:
			goto st508
		case 114:
			goto st508
		}
		goto tr274
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
		goto tr274
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 79:
			goto st507
		case 111:
			goto st507
		}
		goto tr274
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 9:
			goto tr675
		case 32:
			goto tr675
		case 58:
			goto tr676
		}
		goto st0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		switch data[p] {
		case 69:
			goto st509
		case 101:
			goto st509
		}
		goto tr274
	st509:
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		switch data[p] {
		case 68:
			goto st510
		case 100:
			goto st510
		}
		goto tr274
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		if data[p] == 45 {
			goto st511
		}
		goto tr274
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		switch data[p] {
		case 66:
			goto st512
		case 98:
			goto st512
		}
		goto tr274
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		switch data[p] {
		case 89:
			goto st228
		case 121:
			goto st228
		}
		goto tr274
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		switch data[p] {
		case 79:
			goto st514
		case 111:
			goto st514
		}
		goto tr274
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		switch data[p] {
		case 84:
			goto st515
		case 116:
			goto st515
		}
		goto tr274
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		switch data[p] {
		case 69:
			goto st516
		case 101:
			goto st516
		}
		goto tr274
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		if data[p] == 45 {
			goto st517
		}
		goto tr274
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		switch data[p] {
		case 80:
			goto st518
		case 112:
			goto st518
		}
		goto tr274
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		switch data[p] {
		case 65:
			goto st519
		case 97:
			goto st519
		}
		goto tr274
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 82:
			goto st520
		case 114:
			goto st520
		}
		goto tr274
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		switch data[p] {
		case 84:
			goto st521
		case 116:
			goto st521
		}
		goto tr274
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 89:
			goto st522
		case 121:
			goto st522
		}
		goto tr274
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		if data[p] == 45 {
			goto st523
		}
		goto tr274
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
		switch data[p] {
		case 73:
			goto st524
		case 105:
			goto st524
		}
		goto tr274
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
		switch data[p] {
		case 68:
			goto st525
		case 100:
			goto st525
		}
		goto tr274
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		switch data[p] {
		case 9:
			goto tr719
		case 32:
			goto tr719
		case 58:
			goto tr720
		}
		goto st0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		switch data[p] {
		case 76:
			goto st527
		case 108:
			goto st527
		}
		goto tr274
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		switch data[p] {
		case 89:
			goto st528
		case 121:
			goto st528
		}
		goto tr274
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		if data[p] == 45 {
			goto st529
		}
		goto tr274
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
		switch data[p] {
		case 84:
			goto st530
		case 116:
			goto st530
		}
		goto tr274
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
		switch data[p] {
		case 79:
			goto st531
		case 111:
			goto st531
		}
		goto tr274
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 9:
			goto tr726
		case 32:
			goto tr726
		case 58:
			goto tr727
		}
		goto st0
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		switch data[p] {
		case 85:
			goto st533
		case 117:
			goto st533
		}
		goto tr274
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 73:
			goto st534
		case 105:
			goto st534
		}
		goto tr274
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		switch data[p] {
		case 82:
			goto st535
		case 114:
			goto st535
		}
		goto tr274
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		switch data[p] {
		case 69:
			goto st536
		case 101:
			goto st536
		}
		goto tr274
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		switch data[p] {
		case 9:
			goto tr732
		case 32:
			goto tr732
		case 58:
			goto tr733
		}
		goto st0
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		switch data[p] {
		case 82:
			goto st538
		case 114:
			goto st538
		}
		goto tr274
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		switch data[p] {
		case 89:
			goto st539
		case 121:
			goto st539
		}
		goto tr274
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		if data[p] == 45 {
			goto st540
		}
		goto tr274
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
		switch data[p] {
		case 65:
			goto st541
		case 97:
			goto st541
		}
		goto tr274
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		switch data[p] {
		case 70:
			goto st542
		case 102:
			goto st542
		}
		goto tr274
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
		goto tr274
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		switch data[p] {
		case 69:
			goto st544
		case 101:
			goto st544
		}
		goto tr274
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		switch data[p] {
		case 82:
			goto st545
		case 114:
			goto st545
		}
		goto tr274
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		switch data[p] {
		case 9:
			goto tr742
		case 32:
			goto tr742
		case 58:
			goto tr743
		}
		goto st0
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		switch data[p] {
		case 85:
			goto st547
		case 117:
			goto st547
		}
		goto tr274
	st547:
		if p++; p == pe {
			goto _test_eof547
		}
	st_case_547:
		switch data[p] {
		case 84:
			goto st548
		case 116:
			goto st548
		}
		goto tr274
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		switch data[p] {
		case 69:
			goto st549
		case 101:
			goto st549
		}
		goto tr274
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		switch data[p] {
		case 9:
			goto tr747
		case 32:
			goto tr747
		case 58:
			goto tr748
		}
		goto st0
tr257:
//line msg_parse.rl:59

			mark = p
		
	goto st550
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
//line msg_parse.go:13271
		switch data[p] {
		case 9:
			goto tr749
		case 32:
			goto tr749
		case 58:
			goto tr750
		case 69:
			goto st551
		case 85:
			goto st556
		case 101:
			goto st551
		case 117:
			goto st556
		}
		goto st0
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		switch data[p] {
		case 82:
			goto st552
		case 114:
			goto st552
		}
		goto tr274
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		switch data[p] {
		case 86:
			goto st553
		case 118:
			goto st553
		}
		goto tr274
	st553:
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		switch data[p] {
		case 69:
			goto st554
		case 101:
			goto st554
		}
		goto tr274
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		switch data[p] {
		case 82:
			goto st555
		case 114:
			goto st555
		}
		goto tr274
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		switch data[p] {
		case 9:
			goto tr757
		case 32:
			goto tr757
		case 58:
			goto tr758
		}
		goto st0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		switch data[p] {
		case 66:
			goto st557
		case 80:
			goto st562
		case 98:
			goto st557
		case 112:
			goto st562
		}
		goto tr274
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		switch data[p] {
		case 74:
			goto st558
		case 106:
			goto st558
		}
		goto tr274
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		switch data[p] {
		case 69:
			goto st559
		case 101:
			goto st559
		}
		goto tr274
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		switch data[p] {
		case 67:
			goto st560
		case 99:
			goto st560
		}
		goto tr274
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		switch data[p] {
		case 84:
			goto st561
		case 116:
			goto st561
		}
		goto tr274
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		switch data[p] {
		case 9:
			goto tr749
		case 32:
			goto tr749
		case 58:
			goto tr750
		}
		goto st0
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		switch data[p] {
		case 80:
			goto st563
		case 112:
			goto st563
		}
		goto tr274
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 79:
			goto st564
		case 111:
			goto st564
		}
		goto tr274
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
		switch data[p] {
		case 82:
			goto st565
		case 114:
			goto st565
		}
		goto tr274
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		switch data[p] {
		case 84:
			goto st566
		case 116:
			goto st566
		}
		goto tr274
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		switch data[p] {
		case 69:
			goto st567
		case 101:
			goto st567
		}
		goto tr274
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		switch data[p] {
		case 68:
			goto st371
		case 100:
			goto st371
		}
		goto tr274
tr258:
//line msg_parse.rl:59

			mark = p
		
	goto st568
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
//line msg_parse.go:13512
		switch data[p] {
		case 9:
			goto tr771
		case 32:
			goto tr771
		case 58:
			goto tr772
		case 73:
			goto st569
		case 79:
			goto st577
		case 105:
			goto st569
		case 111:
			goto st577
		}
		goto tr274
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		switch data[p] {
		case 77:
			goto st570
		case 109:
			goto st570
		}
		goto tr274
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 69:
			goto st571
		case 101:
			goto st571
		}
		goto tr274
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		switch data[p] {
		case 83:
			goto st572
		case 115:
			goto st572
		}
		goto tr274
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		switch data[p] {
		case 84:
			goto st573
		case 116:
			goto st573
		}
		goto tr274
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		switch data[p] {
		case 65:
			goto st574
		case 97:
			goto st574
		}
		goto tr274
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 77:
			goto st575
		case 109:
			goto st575
		}
		goto tr274
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		switch data[p] {
		case 80:
			goto st576
		case 112:
			goto st576
		}
		goto tr274
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 9:
			goto tr782
		case 32:
			goto tr782
		case 58:
			goto tr783
		}
		goto st0
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		switch data[p] {
		case 9:
			goto tr771
		case 32:
			goto tr771
		case 58:
			goto tr772
		}
		goto st0
tr259:
//line msg_parse.rl:59

			mark = p
		
	goto st578
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
//line msg_parse.go:13653
		switch data[p] {
		case 9:
			goto tr784
		case 32:
			goto tr784
		case 58:
			goto tr785
		case 78:
			goto st579
		case 83:
			goto st589
		case 110:
			goto st579
		case 115:
			goto st589
		}
		goto st0
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		switch data[p] {
		case 83:
			goto st580
		case 115:
			goto st580
		}
		goto tr274
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		switch data[p] {
		case 85:
			goto st581
		case 117:
			goto st581
		}
		goto tr274
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		switch data[p] {
		case 80:
			goto st582
		case 112:
			goto st582
		}
		goto tr274
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 80:
			goto st583
		case 112:
			goto st583
		}
		goto tr274
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		switch data[p] {
		case 79:
			goto st584
		case 111:
			goto st584
		}
		goto tr274
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 82:
			goto st585
		case 114:
			goto st585
		}
		goto tr274
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 84:
			goto st586
		case 116:
			goto st586
		}
		goto tr274
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		switch data[p] {
		case 69:
			goto st587
		case 101:
			goto st587
		}
		goto tr274
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		switch data[p] {
		case 68:
			goto st588
		case 100:
			goto st588
		}
		goto tr274
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		switch data[p] {
		case 9:
			goto tr797
		case 32:
			goto tr797
		case 58:
			goto tr798
		}
		goto st0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		switch data[p] {
		case 69:
			goto st590
		case 101:
			goto st590
		}
		goto tr274
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		switch data[p] {
		case 82:
			goto st591
		case 114:
			goto st591
		}
		goto tr274
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		if data[p] == 45 {
			goto st592
		}
		goto tr274
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		switch data[p] {
		case 65:
			goto st593
		case 97:
			goto st593
		}
		goto tr274
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		switch data[p] {
		case 71:
			goto st594
		case 103:
			goto st594
		}
		goto tr274
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		switch data[p] {
		case 69:
			goto st595
		case 101:
			goto st595
		}
		goto tr274
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 78:
			goto st596
		case 110:
			goto st596
		}
		goto tr274
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		switch data[p] {
		case 84:
			goto st597
		case 116:
			goto st597
		}
		goto tr274
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		switch data[p] {
		case 9:
			goto tr807
		case 32:
			goto tr807
		case 58:
			goto tr808
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		switch data[p] {
		case 9:
			goto st599
		case 32:
			goto st599
		case 58:
			goto st600
		case 73:
			goto st611
		case 105:
			goto st611
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 9:
			goto st599
		case 32:
			goto st599
		case 58:
			goto st600
		}
		goto st0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st600
		case 32:
			goto st600
		case 269:
			goto tr818
		case 525:
			goto st608
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr813
				}
			case _widec >= 33:
				goto tr812
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr815
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr817
				}
			default:
				goto tr816
			}
		default:
			goto tr814
		}
		goto st0
tr821:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st601
tr812:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st601
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
//line msg_parse.go:14004
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st601
		case 32:
			goto st601
		case 269:
			goto tr827
		case 525:
			goto tr828
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr822
				}
			case _widec >= 33:
				goto tr821
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr824
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr826
				}
			default:
				goto tr825
			}
		default:
			goto tr823
		}
		goto st0
tr822:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st602
tr813:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st602
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
//line msg_parse.go:14072
		if 128 <= data[p] && data[p] <= 191 {
			goto tr821
		}
		goto st0
tr823:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st603
tr814:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st603
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
//line msg_parse.go:14100
		if 128 <= data[p] && data[p] <= 191 {
			goto tr822
		}
		goto st0
tr824:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st604
tr815:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st604
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
//line msg_parse.go:14128
		if 128 <= data[p] && data[p] <= 191 {
			goto tr823
		}
		goto st0
tr825:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st605
tr816:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st605
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
//line msg_parse.go:14156
		if 128 <= data[p] && data[p] <= 191 {
			goto tr824
		}
		goto st0
tr826:
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st606
tr817:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:67

			buf[amt] = data[p]
			amt++
		
	goto st606
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
//line msg_parse.go:14184
		if 128 <= data[p] && data[p] <= 191 {
			goto tr825
		}
		goto st0
tr832:
//line msg_parse.rl:59

			mark = p
		
//line msg_parse.rl:226

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st607
tr828:
//line msg_parse.rl:226

			*viap, err = ParseVia(string(data[mark:p]))
			if err != nil { return nil, err }
			for *viap != nil { viap = &(*viap).Next }
		
	goto st607
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
//line msg_parse.go:14214
		if data[p] == 10 {
			goto tr829
		}
		goto st0
tr829:
//line msg_parse.rl:238
 line++; linep = p; 
//line msg_parse.rl:129

			{goto st145 }
		
	goto st647
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
//line msg_parse.go:14232
		switch data[p] {
		case 9:
			goto st601
		case 32:
			goto st601
		}
		goto st0
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		if data[p] == 10 {
			goto tr830
		}
		goto st0
tr830:
//line msg_parse.rl:238
 line++; linep = p; 
	goto st609
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
//line msg_parse.go:14258
		switch data[p] {
		case 9:
			goto st610
		case 32:
			goto st610
		}
		goto st0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st610
		case 32:
			goto st610
		case 269:
			goto tr818
		case 525:
			goto tr832
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr813
				}
			case _widec >= 33:
				goto tr812
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr815
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr817
				}
			default:
				goto tr816
			}
		default:
			goto tr814
		}
		goto st0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 65:
			goto st599
		case 97:
			goto st599
		}
		goto tr463
tr261:
//line msg_parse.rl:59

			mark = p
		
	goto st612
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
//line msg_parse.go:14338
		switch data[p] {
		case 65:
			goto st613
		case 87:
			goto st619
		case 97:
			goto st613
		case 119:
			goto st619
		}
		goto tr274
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		switch data[p] {
		case 82:
			goto st614
		case 114:
			goto st614
		}
		goto tr274
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		switch data[p] {
		case 78:
			goto st615
		case 110:
			goto st615
		}
		goto tr274
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		switch data[p] {
		case 73:
			goto st616
		case 105:
			goto st616
		}
		goto tr274
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		switch data[p] {
		case 78:
			goto st617
		case 110:
			goto st617
		}
		goto tr274
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		switch data[p] {
		case 71:
			goto st618
		case 103:
			goto st618
		}
		goto tr274
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		switch data[p] {
		case 9:
			goto tr840
		case 32:
			goto tr840
		case 58:
			goto tr841
		}
		goto st0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 87:
			goto st620
		case 119:
			goto st620
		}
		goto tr274
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		if data[p] == 45 {
			goto st621
		}
		goto tr274
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		switch data[p] {
		case 65:
			goto st622
		case 97:
			goto st622
		}
		goto tr274
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 85:
			goto st623
		case 117:
			goto st623
		}
		goto tr274
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		switch data[p] {
		case 84:
			goto st624
		case 116:
			goto st624
		}
		goto tr274
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		switch data[p] {
		case 72:
			goto st625
		case 104:
			goto st625
		}
		goto tr274
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
		switch data[p] {
		case 69:
			goto st626
		case 101:
			goto st626
		}
		goto tr274
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
		switch data[p] {
		case 78:
			goto st627
		case 110:
			goto st627
		}
		goto tr274
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
		switch data[p] {
		case 84:
			goto st628
		case 116:
			goto st628
		}
		goto tr274
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		switch data[p] {
		case 73:
			goto st629
		case 105:
			goto st629
		}
		goto tr274
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		switch data[p] {
		case 67:
			goto st630
		case 99:
			goto st630
		}
		goto tr274
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		switch data[p] {
		case 65:
			goto st631
		case 97:
			goto st631
		}
		goto tr274
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		switch data[p] {
		case 84:
			goto st632
		case 116:
			goto st632
		}
		goto tr274
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		switch data[p] {
		case 69:
			goto st633
		case 101:
			goto st633
		}
		goto tr274
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		switch data[p] {
		case 9:
			goto tr856
		case 32:
			goto tr856
		case 58:
			goto tr857
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
	_test_eof634: cs = 634; goto _test_eof
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
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
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
	_test_eof635: cs = 635; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof636: cs = 636; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof637: cs = 637; goto _test_eof
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
	_test_eof638: cs = 638; goto _test_eof
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
	_test_eof639: cs = 639; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof640: cs = 640; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof641: cs = 641; goto _test_eof
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
	_test_eof642: cs = 642; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof643: cs = 643; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof644: cs = 644; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof645: cs = 645; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
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
	_test_eof646: cs = 646; goto _test_eof
	_test_eof240: cs = 240; goto _test_eof
	_test_eof241: cs = 241; goto _test_eof
	_test_eof242: cs = 242; goto _test_eof
	_test_eof243: cs = 243; goto _test_eof
	_test_eof244: cs = 244; goto _test_eof
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
	_test_eof549: cs = 549; goto _test_eof
	_test_eof550: cs = 550; goto _test_eof
	_test_eof551: cs = 551; goto _test_eof
	_test_eof552: cs = 552; goto _test_eof
	_test_eof553: cs = 553; goto _test_eof
	_test_eof554: cs = 554; goto _test_eof
	_test_eof555: cs = 555; goto _test_eof
	_test_eof556: cs = 556; goto _test_eof
	_test_eof557: cs = 557; goto _test_eof
	_test_eof558: cs = 558; goto _test_eof
	_test_eof559: cs = 559; goto _test_eof
	_test_eof560: cs = 560; goto _test_eof
	_test_eof561: cs = 561; goto _test_eof
	_test_eof562: cs = 562; goto _test_eof
	_test_eof563: cs = 563; goto _test_eof
	_test_eof564: cs = 564; goto _test_eof
	_test_eof565: cs = 565; goto _test_eof
	_test_eof566: cs = 566; goto _test_eof
	_test_eof567: cs = 567; goto _test_eof
	_test_eof568: cs = 568; goto _test_eof
	_test_eof569: cs = 569; goto _test_eof
	_test_eof570: cs = 570; goto _test_eof
	_test_eof571: cs = 571; goto _test_eof
	_test_eof572: cs = 572; goto _test_eof
	_test_eof573: cs = 573; goto _test_eof
	_test_eof574: cs = 574; goto _test_eof
	_test_eof575: cs = 575; goto _test_eof
	_test_eof576: cs = 576; goto _test_eof
	_test_eof577: cs = 577; goto _test_eof
	_test_eof578: cs = 578; goto _test_eof
	_test_eof579: cs = 579; goto _test_eof
	_test_eof580: cs = 580; goto _test_eof
	_test_eof581: cs = 581; goto _test_eof
	_test_eof582: cs = 582; goto _test_eof
	_test_eof583: cs = 583; goto _test_eof
	_test_eof584: cs = 584; goto _test_eof
	_test_eof585: cs = 585; goto _test_eof
	_test_eof586: cs = 586; goto _test_eof
	_test_eof587: cs = 587; goto _test_eof
	_test_eof588: cs = 588; goto _test_eof
	_test_eof589: cs = 589; goto _test_eof
	_test_eof590: cs = 590; goto _test_eof
	_test_eof591: cs = 591; goto _test_eof
	_test_eof592: cs = 592; goto _test_eof
	_test_eof593: cs = 593; goto _test_eof
	_test_eof594: cs = 594; goto _test_eof
	_test_eof595: cs = 595; goto _test_eof
	_test_eof596: cs = 596; goto _test_eof
	_test_eof597: cs = 597; goto _test_eof
	_test_eof598: cs = 598; goto _test_eof
	_test_eof599: cs = 599; goto _test_eof
	_test_eof600: cs = 600; goto _test_eof
	_test_eof601: cs = 601; goto _test_eof
	_test_eof602: cs = 602; goto _test_eof
	_test_eof603: cs = 603; goto _test_eof
	_test_eof604: cs = 604; goto _test_eof
	_test_eof605: cs = 605; goto _test_eof
	_test_eof606: cs = 606; goto _test_eof
	_test_eof607: cs = 607; goto _test_eof
	_test_eof647: cs = 647; goto _test_eof
	_test_eof608: cs = 608; goto _test_eof
	_test_eof609: cs = 609; goto _test_eof
	_test_eof610: cs = 610; goto _test_eof
	_test_eof611: cs = 611; goto _test_eof
	_test_eof612: cs = 612; goto _test_eof
	_test_eof613: cs = 613; goto _test_eof
	_test_eof614: cs = 614; goto _test_eof
	_test_eof615: cs = 615; goto _test_eof
	_test_eof616: cs = 616; goto _test_eof
	_test_eof617: cs = 617; goto _test_eof
	_test_eof618: cs = 618; goto _test_eof
	_test_eof619: cs = 619; goto _test_eof
	_test_eof620: cs = 620; goto _test_eof
	_test_eof621: cs = 621; goto _test_eof
	_test_eof622: cs = 622; goto _test_eof
	_test_eof623: cs = 623; goto _test_eof
	_test_eof624: cs = 624; goto _test_eof
	_test_eof625: cs = 625; goto _test_eof
	_test_eof626: cs = 626; goto _test_eof
	_test_eof627: cs = 627; goto _test_eof
	_test_eof628: cs = 628; goto _test_eof
	_test_eof629: cs = 629; goto _test_eof
	_test_eof630: cs = 630; goto _test_eof
	_test_eof631: cs = 631; goto _test_eof
	_test_eof632: cs = 632; goto _test_eof
	_test_eof633: cs = 633; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 153, 154, 155, 156, 158, 159, 160, 161, 162, 163, 164, 166, 167, 168, 169, 170, 171, 172, 174, 175, 176, 177, 178, 179, 180, 182, 183, 184, 185, 186, 187, 188, 189, 191, 192, 194, 195, 196, 197, 198, 199, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 219, 220, 221, 222, 223, 224, 225, 226, 256, 257, 262, 263, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 285, 286, 287, 288, 289, 290, 291, 294, 295, 296, 297, 298, 299, 327, 328, 329, 332, 333, 334, 335, 336, 337, 338, 339, 341, 342, 343, 357, 358, 360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 396, 397, 398, 399, 400, 401, 402, 403, 404, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 432, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 451, 452, 453, 454, 455, 456, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 474, 475, 476, 477, 478, 479, 480, 481, 483, 484, 485, 486, 487, 488, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 502, 503, 504, 505, 506, 508, 509, 510, 511, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 524, 526, 527, 528, 529, 530, 532, 533, 534, 535, 537, 538, 539, 540, 541, 542, 543, 544, 546, 547, 548, 551, 552, 553, 554, 556, 557, 558, 559, 560, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 579, 580, 581, 582, 583, 584, 585, 586, 587, 589, 590, 591, 592, 593, 594, 595, 596, 612, 613, 614, 615, 616, 617, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632:
//line msg_parse.rl:138

			p--

			{goto st138 }
		
		case 301, 302, 303, 304, 311, 312, 313, 314, 315, 331, 345, 346, 347, 348, 349, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 406, 407, 408, 409, 410, 411, 412, 413, 611:
//line msg_parse.rl:144

			p--

			{goto st138 }
		
		case 145, 229, 243, 244, 245, 246, 247, 259, 260, 261, 270, 271, 272, 273, 293, 378, 395:
//line msg_parse.rl:144

			p--

			{goto st138 }
		
//line msg_parse.rl:138

			p--

			{goto st138 }
		
//line msg_parse.go:15281
		}
	}

	_out: {}
	}

//line msg_parse.rl:387


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
