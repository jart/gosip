
//line msg_parse.rl:1
package sip

import (
//  "bytes"
  "errors"
  "fmt"
)


//line msg_parse.rl:10

//line msg_parse.go:15
const msg_start int = 1
const msg_first_final int = 810
const msg_error int = 0

const msg_en_main int = 1


//line msg_parse.rl:11

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
  line := 1
  linep := 0
  buf := make([]byte, len(data))
  amt := 0
  mark := 0
  clen := 0
//  var b1 string
  var hex byte

  
//line msg_parse.go:56
	{
	cs = msg_start
	}

//line msg_parse.go:61
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
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 810:
		goto st_case_810
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
	case 765:
		goto st_case_765
	case 766:
		goto st_case_766
	case 767:
		goto st_case_767
	case 768:
		goto st_case_768
	case 769:
		goto st_case_769
	case 770:
		goto st_case_770
	case 771:
		goto st_case_771
	case 772:
		goto st_case_772
	case 773:
		goto st_case_773
	case 774:
		goto st_case_774
	case 775:
		goto st_case_775
	case 776:
		goto st_case_776
	case 777:
		goto st_case_777
	case 778:
		goto st_case_778
	case 779:
		goto st_case_779
	case 780:
		goto st_case_780
	case 781:
		goto st_case_781
	case 782:
		goto st_case_782
	case 783:
		goto st_case_783
	case 784:
		goto st_case_784
	case 785:
		goto st_case_785
	case 786:
		goto st_case_786
	case 787:
		goto st_case_787
	case 788:
		goto st_case_788
	case 789:
		goto st_case_789
	case 790:
		goto st_case_790
	case 791:
		goto st_case_791
	case 792:
		goto st_case_792
	case 793:
		goto st_case_793
	case 794:
		goto st_case_794
	case 795:
		goto st_case_795
	case 796:
		goto st_case_796
	case 797:
		goto st_case_797
	case 798:
		goto st_case_798
	case 799:
		goto st_case_799
	case 800:
		goto st_case_800
	case 801:
		goto st_case_801
	case 802:
		goto st_case_802
	case 803:
		goto st_case_803
	case 804:
		goto st_case_804
	case 805:
		goto st_case_805
	case 806:
		goto st_case_806
	case 807:
		goto st_case_807
	case 808:
		goto st_case_808
	case 809:
		goto st_case_809
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
//line msg_parse.rl:47

      mark = p
    
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line msg_parse.go:1743
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
//line msg_parse.rl:78

      msg.Method = string(data[mark:p])
    
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line msg_parse.go:1790
		if data[p] == 32 {
			goto st0
		}
		goto tr5
tr5:
//line msg_parse.rl:47

      mark = p
    
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line msg_parse.go:1806
		if data[p] == 32 {
			goto tr7
		}
		goto st4
tr7:
//line msg_parse.rl:90

      msg.Request, err = ParseURIBytes(data[mark:p])
      if err != nil { return nil, err }
    
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line msg_parse.go:1823
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
//line msg_parse.rl:82

      msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
    
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line msg_parse.go:1875
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
//line msg_parse.rl:86

      msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
    
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line msg_parse.go:1903
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr14
		}
		goto st0
tr44:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:118

      msg.AcceptContact = string(data[mark:p])
    
	goto st13
tr47:
//line msg_parse.rl:118

      msg.AcceptContact = string(data[mark:p])
    
	goto st13
tr53:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:272

      msg.ReferredBy = string(data[mark:p])
    
	goto st13
tr56:
//line msg_parse.rl:272

      msg.ReferredBy = string(data[mark:p])
    
	goto st13
tr66:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:176

      msg.ContentType = string(data[mark:p])
    
	goto st13
tr69:
//line msg_parse.rl:176

      msg.ContentType = string(data[mark:p])
    
	goto st13
tr78:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:192

      msg.Date = string(data[mark:p])
    
	goto st13
tr81:
//line msg_parse.rl:192

      msg.Date = string(data[mark:p])
    
	goto st13
tr91:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:172

      msg.ContentEncoding = string(data[mark:p])
    
	goto st13
tr94:
//line msg_parse.rl:172

      msg.ContentEncoding = string(data[mark:p])
    
	goto st13
tr102:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:208

      msg.From, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr105:
//line msg_parse.rl:208

      msg.From, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr113:
//line msg_parse.rl:150

      msg.CallID = string(data[mark:p])
    
	goto st13
tr131:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:213

      msg.InReplyTo = string(data[mark:p])
    
	goto st13
tr134:
//line msg_parse.rl:213

      msg.InReplyTo = string(data[mark:p])
    
	goto st13
tr140:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:303

      msg.Supported = string(data[mark:p])
    
	goto st13
tr143:
//line msg_parse.rl:303

      msg.Supported = string(data[mark:p])
    
	goto st13
tr158:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:154

      *contactp, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *contactp != nil { contactp = &(*contactp).Next }
    
	goto st13
tr161:
//line msg_parse.rl:154

      *contactp, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *contactp != nil { contactp = &(*contactp).Next }
    
	goto st13
tr169:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:200

      msg.Event = string(data[mark:p])
    
	goto st13
tr172:
//line msg_parse.rl:200

      msg.Event = string(data[mark:p])
    
	goto st13
tr197:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:241

      msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr200:
//line msg_parse.rl:241

      msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr209:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:268

      msg.ReferTo = string(data[mark:p])
    
	goto st13
tr212:
//line msg_parse.rl:268

      msg.ReferTo = string(data[mark:p])
    
	goto st13
tr221:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:299

      msg.Subject = string(data[mark:p])
    
	goto st13
tr224:
//line msg_parse.rl:299

      msg.Subject = string(data[mark:p])
    
	goto st13
tr232:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:311

      msg.To, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr235:
//line msg_parse.rl:311

      msg.To, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr244:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st13
tr247:
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st13
tr255:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:324

      *viap, err = ParseVia(string(data[mark:p]))
      if err != nil { return nil, err }
      for *viap != nil { viap = &(*viap).Next }
    
	goto st13
tr258:
//line msg_parse.rl:324

      *viap, err = ParseVia(string(data[mark:p]))
      if err != nil { return nil, err }
      for *viap != nil { viap = &(*viap).Next }
    
	goto st13
tr271:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:330

      msg.Warning = string(data[mark:p])
    
	goto st13
tr274:
//line msg_parse.rl:330

      msg.Warning = string(data[mark:p])
    
	goto st13
tr298:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:334

      msg.WWWAuthenticate = string(data[mark:p])
    
	goto st13
tr301:
//line msg_parse.rl:334

      msg.WWWAuthenticate = string(data[mark:p])
    
	goto st13
tr328:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:316

      msg.Unsupported = string(data[mark:p])
    
	goto st13
tr331:
//line msg_parse.rl:316

      msg.Unsupported = string(data[mark:p])
    
	goto st13
tr349:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:320

      msg.UserAgent = string(data[mark:p])
    
	goto st13
tr352:
//line msg_parse.rl:320

      msg.UserAgent = string(data[mark:p])
    
	goto st13
tr373:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:307

      msg.Timestamp = string(data[mark:p])
    
	goto st13
tr376:
//line msg_parse.rl:307

      msg.Timestamp = string(data[mark:p])
    
	goto st13
tr394:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:295

      msg.Server = string(data[mark:p])
    
	goto st13
tr397:
//line msg_parse.rl:295

      msg.Server = string(data[mark:p])
    
	goto st13
tr436:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:262

      *rroutep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *rroutep != nil { rroutep = &(*rroutep).Next }
    
	goto st13
tr439:
//line msg_parse.rl:262

      *rroutep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *rroutep != nil { rroutep = &(*rroutep).Next }
    
	goto st13
tr470:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:276

      msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr473:
//line msg_parse.rl:276

      msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st13
tr488:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:229

      msg.ReplyTo = string(data[mark:p])
    
	goto st13
tr491:
//line msg_parse.rl:229

      msg.ReplyTo = string(data[mark:p])
    
	goto st13
tr505:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:281

      msg.Require = string(data[mark:p])
    
	goto st13
tr508:
//line msg_parse.rl:281

      msg.Require = string(data[mark:p])
    
	goto st13
tr526:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:285

      msg.RetryAfter = string(data[mark:p])
    
	goto st13
tr529:
//line msg_parse.rl:285

      msg.RetryAfter = string(data[mark:p])
    
	goto st13
tr542:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:289

      *routep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *routep != nil { routep = &(*routep).Next }
    
	goto st13
tr545:
//line msg_parse.rl:289

      *routep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *routep != nil { routep = &(*routep).Next }
    
	goto st13
tr566:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:246

      msg.Priority = string(data[mark:p])
    
	goto st13
tr569:
//line msg_parse.rl:246

      msg.Priority = string(data[mark:p])
    
	goto st13
tr596:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:250

      msg.ProxyAuthenticate = string(data[mark:p])
    
	goto st13
tr599:
//line msg_parse.rl:250

      msg.ProxyAuthenticate = string(data[mark:p])
    
	goto st13
tr617:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:254

      msg.ProxyAuthorization = string(data[mark:p])
    
	goto st13
tr620:
//line msg_parse.rl:254

      msg.ProxyAuthorization = string(data[mark:p])
    
	goto st13
tr636:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:258

      msg.ProxyRequire = string(data[mark:p])
    
	goto st13
tr639:
//line msg_parse.rl:258

      msg.ProxyRequire = string(data[mark:p])
    
	goto st13
tr663:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:237

      msg.Organization = string(data[mark:p])
    
	goto st13
tr666:
//line msg_parse.rl:237

      msg.Organization = string(data[mark:p])
    
	goto st13
tr707:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:233

      msg.MIMEVersion = string(data[mark:p])
    
	goto st13
tr710:
//line msg_parse.rl:233

      msg.MIMEVersion = string(data[mark:p])
    
	goto st13
tr758:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:196

      msg.ErrorInfo = string(data[mark:p])
    
	goto st13
tr761:
//line msg_parse.rl:196

      msg.ErrorInfo = string(data[mark:p])
    
	goto st13
tr798:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:188

      msg.CallInfo = string(data[mark:p])
    
	goto st13
tr801:
//line msg_parse.rl:188

      msg.CallInfo = string(data[mark:p])
    
	goto st13
tr833:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:160

      msg.ContentDisposition = string(data[mark:p])
    
	goto st13
tr836:
//line msg_parse.rl:160

      msg.ContentDisposition = string(data[mark:p])
    
	goto st13
tr860:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:164

      msg.ContentLanguage = string(data[mark:p])
    
	goto st13
tr863:
//line msg_parse.rl:164

      msg.ContentLanguage = string(data[mark:p])
    
	goto st13
tr889:
//line msg_parse.rl:184

      msg.CSeqMethod = string(data[mark:p])
    
	goto st13
tr912:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:114

      msg.Accept = string(data[mark:p])
    
	goto st13
tr915:
//line msg_parse.rl:114

      msg.Accept = string(data[mark:p])
    
	goto st13
tr940:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:122

      msg.AcceptEncoding = string(data[mark:p])
    
	goto st13
tr943:
//line msg_parse.rl:122

      msg.AcceptEncoding = string(data[mark:p])
    
	goto st13
tr960:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:126

      msg.AcceptLanguage = string(data[mark:p])
    
	goto st13
tr963:
//line msg_parse.rl:126

      msg.AcceptLanguage = string(data[mark:p])
    
	goto st13
tr982:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:138

      msg.AlertInfo = string(data[mark:p])
    
	goto st13
tr985:
//line msg_parse.rl:138

      msg.AlertInfo = string(data[mark:p])
    
	goto st13
tr999:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
	goto st13
tr1002:
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
	goto st13
tr1018:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st13
tr1021:
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st13
tr1049:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:142

      msg.AuthenticationInfo = string(data[mark:p])
    
	goto st13
tr1052:
//line msg_parse.rl:142

      msg.AuthenticationInfo = string(data[mark:p])
    
	goto st13
tr1070:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:146

      msg.Authorization = string(data[mark:p])
    
	goto st13
tr1073:
//line msg_parse.rl:146

      msg.Authorization = string(data[mark:p])
    
	goto st13
tr1094:
//line msg_parse.rl:99

      msg.Phrase = string(buf[0:amt])
    
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line msg_parse.go:2702
		if data[p] == 10 {
			goto tr16
		}
		goto st0
tr16:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line msg_parse.go:2716
		switch data[p] {
		case 13:
			goto st15
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 10 {
			goto tr36
		}
		goto st0
tr36:
//line msg_parse.rl:344
 line++; linep = p; 
//line msg_parse.rl:43

      {p++; cs = 810; goto _out }
    
	goto st810
	st810:
		if p++; p == pe {
			goto _test_eof810
		}
	st_case_810:
//line msg_parse.go:2816
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch data[p] {
		case 9:
			goto st17
		case 32:
			goto st17
		case 58:
			goto st18
		case 67:
			goto st663
		case 76:
			goto st713
		case 85:
			goto st754
		case 99:
			goto st663
		case 108:
			goto st713
		case 117:
			goto st754
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 9:
			goto st17
		case 32:
			goto st17
		case 58:
			goto st18
		}
		goto st0
tr42:
//line msg_parse.rl:47

      mark = p
    
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line msg_parse.go:2869
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr42
		case 32:
			goto tr42
		case 269:
			goto tr44
		case 525:
			goto tr45
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr43
		}
		goto st0
tr43:
//line msg_parse.rl:47

      mark = p
    
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line msg_parse.go:2902
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st19
		case 269:
			goto tr47
		case 525:
			goto tr48
		}
		if 32 <= _widec && _widec <= 253 {
			goto st19
		}
		goto st0
tr902:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:118

      msg.AcceptContact = string(data[mark:p])
    
	goto st20
tr48:
//line msg_parse.rl:118

      msg.AcceptContact = string(data[mark:p])
    
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line msg_parse.go:2943
		if data[p] == 10 {
			goto tr49
		}
		goto st0
tr49:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line msg_parse.go:2957
		switch data[p] {
		case 9:
			goto st19
		case 13:
			goto st15
		case 32:
			goto st19
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 9:
			goto st22
		case 32:
			goto st22
		case 58:
			goto st23
		}
		goto st0
tr51:
//line msg_parse.rl:47

      mark = p
    
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line msg_parse.go:3064
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr51
		case 32:
			goto tr51
		case 269:
			goto tr53
		case 525:
			goto tr54
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr52
		}
		goto st0
tr52:
//line msg_parse.rl:47

      mark = p
    
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line msg_parse.go:3097
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st24
		case 269:
			goto tr56
		case 525:
			goto tr57
		}
		if 32 <= _widec && _widec <= 253 {
			goto st24
		}
		goto st0
tr898:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:272

      msg.ReferredBy = string(data[mark:p])
    
	goto st25
tr57:
//line msg_parse.rl:272

      msg.ReferredBy = string(data[mark:p])
    
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line msg_parse.go:3138
		if data[p] == 10 {
			goto tr58
		}
		goto st0
tr58:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line msg_parse.go:3152
		switch data[p] {
		case 9:
			goto st24
		case 13:
			goto st15
		case 32:
			goto st24
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 58:
			goto st29
		case 65:
			goto st567
		case 79:
			goto st582
		case 83:
			goto st644
		case 97:
			goto st567
		case 111:
			goto st582
		case 115:
			goto st644
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 58:
			goto st29
		}
		goto st0
tr64:
//line msg_parse.rl:47

      mark = p
    
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line msg_parse.go:3285
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
		case 32:
			goto tr64
		case 269:
			goto tr66
		case 525:
			goto tr67
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr65
		}
		goto st0
tr65:
//line msg_parse.rl:47

      mark = p
    
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line msg_parse.go:3318
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
			goto tr69
		case 525:
			goto tr70
		}
		if 32 <= _widec && _widec <= 253 {
			goto st30
		}
		goto st0
tr787:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:176

      msg.ContentType = string(data[mark:p])
    
	goto st31
tr70:
//line msg_parse.rl:176

      msg.ContentType = string(data[mark:p])
    
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line msg_parse.go:3359
		if data[p] == 10 {
			goto tr71
		}
		goto st0
tr71:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line msg_parse.go:3373
		switch data[p] {
		case 9:
			goto st30
		case 13:
			goto st15
		case 32:
			goto st30
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 65:
			goto st34
		case 97:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 84:
			goto st35
		case 116:
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 69:
			goto st36
		case 101:
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 9:
			goto st36
		case 32:
			goto st36
		case 58:
			goto st37
		}
		goto st0
tr76:
//line msg_parse.rl:47

      mark = p
    
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line msg_parse.go:3516
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr76
		case 32:
			goto tr76
		case 269:
			goto tr78
		case 525:
			goto tr79
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr77
		}
		goto st0
tr77:
//line msg_parse.rl:47

      mark = p
    
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line msg_parse.go:3549
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st38
		case 269:
			goto tr81
		case 525:
			goto tr82
		}
		if 32 <= _widec && _widec <= 253 {
			goto st38
		}
		goto st0
tr783:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:192

      msg.Date = string(data[mark:p])
    
	goto st39
tr82:
//line msg_parse.rl:192

      msg.Date = string(data[mark:p])
    
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line msg_parse.go:3590
		if data[p] == 10 {
			goto tr83
		}
		goto st0
tr83:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line msg_parse.go:3604
		switch data[p] {
		case 9:
			goto st38
		case 13:
			goto st15
		case 32:
			goto st38
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 9:
			goto st42
		case 32:
			goto st42
		case 58:
			goto st43
		case 82:
			goto st531
		case 86:
			goto st547
		case 88:
			goto st550
		case 114:
			goto st531
		case 118:
			goto st547
		case 120:
			goto st550
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 9:
			goto st42
		case 32:
			goto st42
		case 58:
			goto st43
		}
		goto st0
tr89:
//line msg_parse.rl:47

      mark = p
    
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line msg_parse.go:3737
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr89
		case 32:
			goto tr89
		case 269:
			goto tr91
		case 525:
			goto tr92
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr90
		}
		goto st0
tr90:
//line msg_parse.rl:47

      mark = p
    
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line msg_parse.go:3770
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st44
		case 269:
			goto tr94
		case 525:
			goto tr95
		}
		if 32 <= _widec && _widec <= 253 {
			goto st44
		}
		goto st0
tr746:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:172

      msg.ContentEncoding = string(data[mark:p])
    
	goto st45
tr95:
//line msg_parse.rl:172

      msg.ContentEncoding = string(data[mark:p])
    
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line msg_parse.go:3811
		if data[p] == 10 {
			goto tr96
		}
		goto st0
tr96:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line msg_parse.go:3825
		switch data[p] {
		case 9:
			goto st44
		case 13:
			goto st15
		case 32:
			goto st44
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 9:
			goto st48
		case 32:
			goto st48
		case 58:
			goto st49
		case 82:
			goto st526
		case 114:
			goto st526
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 9:
			goto st48
		case 32:
			goto st48
		case 58:
			goto st49
		}
		goto st0
tr100:
//line msg_parse.rl:47

      mark = p
    
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line msg_parse.go:3950
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr100
		case 32:
			goto tr100
		case 269:
			goto tr102
		case 525:
			goto tr103
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr101
		}
		goto st0
tr101:
//line msg_parse.rl:47

      mark = p
    
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line msg_parse.go:3983
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st50
		case 269:
			goto tr105
		case 525:
			goto tr106
		}
		if 32 <= _widec && _widec <= 253 {
			goto st50
		}
		goto st0
tr741:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:208

      msg.From, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st51
tr106:
//line msg_parse.rl:208

      msg.From, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line msg_parse.go:4026
		if data[p] == 10 {
			goto tr107
		}
		goto st0
tr107:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line msg_parse.go:4040
		switch data[p] {
		case 9:
			goto st50
		case 13:
			goto st15
		case 32:
			goto st50
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
		switch data[p] {
		case 9:
			goto st54
		case 32:
			goto st54
		case 58:
			goto st55
		case 78:
			goto st62
		case 110:
			goto st62
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 9:
			goto st54
		case 32:
			goto st54
		case 58:
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
			goto st55
		case 32:
			goto st55
		case 37:
			goto tr111
		case 60:
			goto tr111
		case 525:
			goto st59
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto tr111
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto tr111
				}
			default:
				goto tr111
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto tr111
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto tr111
				}
			default:
				goto tr111
			}
		default:
			goto tr111
		}
		goto st0
tr111:
//line msg_parse.rl:47

      mark = p
    
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line msg_parse.go:4220
		switch data[p] {
		case 13:
			goto tr113
		case 37:
			goto st56
		case 60:
			goto st56
		case 64:
			goto st57
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 34:
				if 39 <= data[p] && data[p] <= 43 {
					goto st56
				}
			case data[p] >= 33:
				goto st56
			}
		case data[p] > 58:
			switch {
			case data[p] < 95:
				if 62 <= data[p] && data[p] <= 93 {
					goto st56
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st56
				}
			default:
				goto st56
			}
		default:
			goto st56
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 37:
			goto st58
		case 60:
			goto st58
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st58
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st58
				}
			default:
				goto st58
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st58
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st58
				}
			default:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 13:
			goto tr113
		case 37:
			goto st58
		case 60:
			goto st58
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st58
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st58
				}
			default:
				goto st58
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st58
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st58
				}
			default:
				goto st58
			}
		default:
			goto st58
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 10 {
			goto tr117
		}
		goto st0
tr117:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line msg_parse.go:4362
		switch data[p] {
		case 9:
			goto st61
		case 32:
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
			goto st61
		case 32:
			goto st61
		case 37:
			goto tr111
		case 60:
			goto tr111
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto tr111
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto tr111
				}
			default:
				goto tr111
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto tr111
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto tr111
				}
			default:
				goto tr111
			}
		default:
			goto tr111
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 45 {
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 82:
			goto st64
		case 114:
			goto st64
		}
		goto st0
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
		switch data[p] {
		case 69:
			goto st65
		case 101:
			goto st65
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch data[p] {
		case 80:
			goto st66
		case 112:
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 76:
			goto st67
		case 108:
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 89:
			goto st68
		case 121:
			goto st68
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		if data[p] == 45 {
			goto st69
		}
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		switch data[p] {
		case 84:
			goto st70
		case 116:
			goto st70
		}
		goto st0
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
		switch data[p] {
		case 79:
			goto st71
		case 111:
			goto st71
		}
		goto st0
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
		switch data[p] {
		case 9:
			goto st71
		case 32:
			goto st71
		case 58:
			goto st72
		}
		goto st0
tr129:
//line msg_parse.rl:47

      mark = p
    
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line msg_parse.go:4543
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr129
		case 32:
			goto tr129
		case 269:
			goto tr131
		case 525:
			goto tr132
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr130
		}
		goto st0
tr130:
//line msg_parse.rl:47

      mark = p
    
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line msg_parse.go:4576
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st73
		case 269:
			goto tr134
		case 525:
			goto tr135
		}
		if 32 <= _widec && _widec <= 253 {
			goto st73
		}
		goto st0
tr737:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:213

      msg.InReplyTo = string(data[mark:p])
    
	goto st74
tr135:
//line msg_parse.rl:213

      msg.InReplyTo = string(data[mark:p])
    
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line msg_parse.go:4617
		if data[p] == 10 {
			goto tr136
		}
		goto st0
tr136:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line msg_parse.go:4631
		switch data[p] {
		case 9:
			goto st73
		case 13:
			goto st15
		case 32:
			goto st73
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 9:
			goto st76
		case 32:
			goto st76
		case 58:
			goto st77
		}
		goto st0
tr138:
//line msg_parse.rl:47

      mark = p
    
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line msg_parse.go:4738
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr138
		case 32:
			goto tr138
		case 269:
			goto tr140
		case 525:
			goto tr141
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr139
		}
		goto st0
tr139:
//line msg_parse.rl:47

      mark = p
    
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line msg_parse.go:4771
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st78
		case 269:
			goto tr143
		case 525:
			goto tr144
		}
		if 32 <= _widec && _widec <= 253 {
			goto st78
		}
		goto st0
tr733:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:303

      msg.Supported = string(data[mark:p])
    
	goto st79
tr144:
//line msg_parse.rl:303

      msg.Supported = string(data[mark:p])
    
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line msg_parse.go:4812
		if data[p] == 10 {
			goto tr145
		}
		goto st0
tr145:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line msg_parse.go:4826
		switch data[p] {
		case 9:
			goto st78
		case 13:
			goto st15
		case 32:
			goto st78
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		switch data[p] {
		case 9:
			goto st81
		case 32:
			goto st81
		case 58:
			goto st82
		}
		goto st0
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st82
		case 32:
			goto st82
		case 525:
			goto st84
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr147
		}
		goto st0
tr147:
//line msg_parse.rl:168

      clen = clen * 10 + (int(data[p]) - 0x30)
    
//line msg_parse.rl:204

      msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
    
//line msg_parse.rl:217

      msg.MaxForwards = 0
    
//line msg_parse.rl:221

      msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
    
//line msg_parse.rl:225

      msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
    
	goto st83
tr149:
//line msg_parse.rl:168

      clen = clen * 10 + (int(data[p]) - 0x30)
    
//line msg_parse.rl:204

      msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
    
//line msg_parse.rl:221

      msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
    
//line msg_parse.rl:225

      msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
    
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line msg_parse.go:4991
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr149
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		if data[p] == 10 {
			goto tr150
		}
		goto st0
tr150:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line msg_parse.go:5017
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
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr147
		}
		goto st0
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
		switch data[p] {
		case 9:
			goto st88
		case 32:
			goto st88
		case 58:
			goto st89
		case 65:
			goto st469
		case 73:
			goto st485
		case 97:
			goto st469
		case 105:
			goto st485
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 9:
			goto st88
		case 32:
			goto st88
		case 58:
			goto st89
		}
		goto st0
tr156:
//line msg_parse.rl:47

      mark = p
    
	goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line msg_parse.go:5087
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr156
		case 32:
			goto tr156
		case 269:
			goto tr158
		case 525:
			goto tr159
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr157
		}
		goto st0
tr157:
//line msg_parse.rl:47

      mark = p
    
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line msg_parse.go:5120
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st90
		case 269:
			goto tr161
		case 525:
			goto tr162
		}
		if 32 <= _widec && _widec <= 253 {
			goto st90
		}
		goto st0
tr676:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:154

      *contactp, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *contactp != nil { contactp = &(*contactp).Next }
    
	goto st91
tr162:
//line msg_parse.rl:154

      *contactp, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *contactp != nil { contactp = &(*contactp).Next }
    
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line msg_parse.go:5165
		if data[p] == 10 {
			goto tr163
		}
		goto st0
tr163:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line msg_parse.go:5179
		switch data[p] {
		case 9:
			goto st90
		case 13:
			goto st15
		case 32:
			goto st90
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
		switch data[p] {
		case 9:
			goto st94
		case 32:
			goto st94
		case 58:
			goto st95
		case 82:
			goto st448
		case 114:
			goto st448
		}
		goto st0
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
		switch data[p] {
		case 9:
			goto st94
		case 32:
			goto st94
		case 58:
			goto st95
		}
		goto st0
tr167:
//line msg_parse.rl:47

      mark = p
    
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line msg_parse.go:5304
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr167
		case 32:
			goto tr167
		case 269:
			goto tr169
		case 525:
			goto tr170
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr168
		}
		goto st0
tr168:
//line msg_parse.rl:47

      mark = p
    
	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line msg_parse.go:5337
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st96
		case 269:
			goto tr172
		case 525:
			goto tr173
		}
		if 32 <= _widec && _widec <= 253 {
			goto st96
		}
		goto st0
tr649:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:200

      msg.Event = string(data[mark:p])
    
	goto st97
tr173:
//line msg_parse.rl:200

      msg.Event = string(data[mark:p])
    
	goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line msg_parse.go:5378
		if data[p] == 10 {
			goto tr174
		}
		goto st0
tr174:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line msg_parse.go:5392
		switch data[p] {
		case 9:
			goto st96
		case 13:
			goto st15
		case 32:
			goto st96
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
		switch data[p] {
		case 45:
			goto st100
		case 82:
			goto st378
		case 114:
			goto st378
		}
		goto st0
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
		switch data[p] {
		case 65:
			goto st101
		case 97:
			goto st101
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 83:
			goto st102
		case 115:
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 83:
			goto st103
		case 115:
			goto st103
		}
		goto st0
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
		switch data[p] {
		case 69:
			goto st104
		case 101:
			goto st104
		}
		goto st0
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
		switch data[p] {
		case 82:
			goto st105
		case 114:
			goto st105
		}
		goto st0
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
		switch data[p] {
		case 84:
			goto st106
		case 116:
			goto st106
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 69:
			goto st107
		case 101:
			goto st107
		}
		goto st0
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
		switch data[p] {
		case 68:
			goto st108
		case 100:
			goto st108
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		if data[p] == 45 {
			goto st109
		}
		goto st0
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
		switch data[p] {
		case 73:
			goto st110
		case 105:
			goto st110
		}
		goto st0
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
		switch data[p] {
		case 68:
			goto st111
		case 100:
			goto st111
		}
		goto st0
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
		switch data[p] {
		case 69:
			goto st112
		case 101:
			goto st112
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 78:
			goto st113
		case 110:
			goto st113
		}
		goto st0
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
		switch data[p] {
		case 84:
			goto st114
		case 116:
			goto st114
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 73:
			goto st115
		case 105:
			goto st115
		}
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		switch data[p] {
		case 84:
			goto st116
		case 116:
			goto st116
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 89:
			goto st117
		case 121:
			goto st117
		}
		goto st0
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
		switch data[p] {
		case 9:
			goto st117
		case 32:
			goto st117
		case 58:
			goto st118
		}
		goto st0
tr195:
//line msg_parse.rl:47

      mark = p
    
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line msg_parse.go:5714
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr195
		case 32:
			goto tr195
		case 269:
			goto tr197
		case 525:
			goto tr198
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr196
		}
		goto st0
tr196:
//line msg_parse.rl:47

      mark = p
    
	goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line msg_parse.go:5747
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
		case 269:
			goto tr200
		case 525:
			goto tr201
		}
		if 32 <= _widec && _widec <= 253 {
			goto st119
		}
		goto st0
tr555:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:241

      msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st120
tr201:
//line msg_parse.rl:241

      msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line msg_parse.go:5790
		if data[p] == 10 {
			goto tr202
		}
		goto st0
tr202:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line msg_parse.go:5804
		switch data[p] {
		case 9:
			goto st119
		case 13:
			goto st15
		case 32:
			goto st119
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 9:
			goto st123
		case 32:
			goto st123
		case 58:
			goto st124
		case 69:
			goto st275
		case 79:
			goto st364
		case 101:
			goto st275
		case 111:
			goto st364
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 9:
			goto st123
		case 32:
			goto st123
		case 58:
			goto st124
		}
		goto st0
tr207:
//line msg_parse.rl:47

      mark = p
    
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line msg_parse.go:5933
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
		case 269:
			goto tr209
		case 525:
			goto tr210
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr208
		}
		goto st0
tr208:
//line msg_parse.rl:47

      mark = p
    
	goto st125
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
//line msg_parse.go:5966
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
		case 269:
			goto tr212
		case 525:
			goto tr213
		}
		if 32 <= _widec && _widec <= 253 {
			goto st125
		}
		goto st0
tr417:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:268

      msg.ReferTo = string(data[mark:p])
    
	goto st126
tr213:
//line msg_parse.rl:268

      msg.ReferTo = string(data[mark:p])
    
	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line msg_parse.go:6007
		if data[p] == 10 {
			goto tr214
		}
		goto st0
tr214:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line msg_parse.go:6021
		switch data[p] {
		case 9:
			goto st125
		case 13:
			goto st15
		case 32:
			goto st125
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
		switch data[p] {
		case 9:
			goto st129
		case 32:
			goto st129
		case 58:
			goto st130
		case 69:
			goto st249
		case 85:
			goto st261
		case 101:
			goto st249
		case 117:
			goto st261
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 9:
			goto st129
		case 32:
			goto st129
		case 58:
			goto st130
		}
		goto st0
tr219:
//line msg_parse.rl:47

      mark = p
    
	goto st130
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
//line msg_parse.go:6150
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr219
		case 32:
			goto tr219
		case 269:
			goto tr221
		case 525:
			goto tr222
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr220
		}
		goto st0
tr220:
//line msg_parse.rl:47

      mark = p
    
	goto st131
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
//line msg_parse.go:6183
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st131
		case 269:
			goto tr224
		case 525:
			goto tr225
		}
		if 32 <= _widec && _widec <= 253 {
			goto st131
		}
		goto st0
tr386:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:299

      msg.Subject = string(data[mark:p])
    
	goto st132
tr225:
//line msg_parse.rl:299

      msg.Subject = string(data[mark:p])
    
	goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line msg_parse.go:6224
		if data[p] == 10 {
			goto tr226
		}
		goto st0
tr226:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line msg_parse.go:6238
		switch data[p] {
		case 9:
			goto st131
		case 13:
			goto st15
		case 32:
			goto st131
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
		switch data[p] {
		case 9:
			goto st135
		case 32:
			goto st135
		case 58:
			goto st136
		case 73:
			goto st231
		case 79:
			goto st135
		case 105:
			goto st231
		case 111:
			goto st135
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 9:
			goto st135
		case 32:
			goto st135
		case 58:
			goto st136
		}
		goto st0
tr230:
//line msg_parse.rl:47

      mark = p
    
	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line msg_parse.go:6367
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr230
		case 32:
			goto tr230
		case 269:
			goto tr232
		case 525:
			goto tr233
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr231
		}
		goto st0
tr231:
//line msg_parse.rl:47

      mark = p
    
	goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line msg_parse.go:6400
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st137
		case 269:
			goto tr235
		case 525:
			goto tr236
		}
		if 32 <= _widec && _widec <= 253 {
			goto st137
		}
		goto st0
tr362:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:311

      msg.To, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st138
tr236:
//line msg_parse.rl:311

      msg.To, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line msg_parse.go:6443
		if data[p] == 10 {
			goto tr237
		}
		goto st0
tr237:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line msg_parse.go:6457
		switch data[p] {
		case 9:
			goto st137
		case 13:
			goto st15
		case 32:
			goto st137
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 9:
			goto st141
		case 32:
			goto st141
		case 58:
			goto st142
		case 78:
			goto st195
		case 83:
			goto st212
		case 110:
			goto st195
		case 115:
			goto st212
		}
		goto st0
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		switch data[p] {
		case 9:
			goto st141
		case 32:
			goto st141
		case 58:
			goto st142
		}
		goto st0
tr242:
//line msg_parse.rl:47

      mark = p
    
	goto st142
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
//line msg_parse.go:6586
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr242
		case 32:
			goto tr242
		case 269:
			goto tr244
		case 525:
			goto tr245
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr243
		}
		goto st0
tr243:
//line msg_parse.rl:47

      mark = p
    
	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line msg_parse.go:6619
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st143
		case 269:
			goto tr247
		case 525:
			goto tr248
		}
		if 32 <= _widec && _widec <= 253 {
			goto st143
		}
		goto st0
tr315:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st144
tr248:
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line msg_parse.go:6668
		if data[p] == 10 {
			goto tr249
		}
		goto st0
tr249:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st145
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
//line msg_parse.go:6682
		switch data[p] {
		case 9:
			goto st143
		case 13:
			goto st15
		case 32:
			goto st143
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
		switch data[p] {
		case 9:
			goto st147
		case 32:
			goto st147
		case 58:
			goto st148
		case 73:
			goto st191
		case 105:
			goto st191
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
tr253:
//line msg_parse.rl:47

      mark = p
    
	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line msg_parse.go:6807
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr253
		case 32:
			goto tr253
		case 269:
			goto tr255
		case 525:
			goto tr256
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr254
		}
		goto st0
tr254:
//line msg_parse.rl:47

      mark = p
    
	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line msg_parse.go:6840
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
			goto tr258
		case 525:
			goto tr259
		}
		if 32 <= _widec && _widec <= 253 {
			goto st149
		}
		goto st0
tr311:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:324

      *viap, err = ParseVia(string(data[mark:p]))
      if err != nil { return nil, err }
      for *viap != nil { viap = &(*viap).Next }
    
	goto st150
tr259:
//line msg_parse.rl:324

      *viap, err = ParseVia(string(data[mark:p]))
      if err != nil { return nil, err }
      for *viap != nil { viap = &(*viap).Next }
    
	goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line msg_parse.go:6885
		if data[p] == 10 {
			goto tr260
		}
		goto st0
tr260:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line msg_parse.go:6899
		switch data[p] {
		case 9:
			goto st149
		case 13:
			goto st15
		case 32:
			goto st149
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
		switch data[p] {
		case 65:
			goto st153
		case 87:
			goto st166
		case 97:
			goto st153
		case 119:
			goto st166
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 82:
			goto st154
		case 114:
			goto st154
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
		case 73:
			goto st156
		case 105:
			goto st156
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 78:
			goto st157
		case 110:
			goto st157
		}
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		switch data[p] {
		case 71:
			goto st158
		case 103:
			goto st158
		}
		goto st0
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
		switch data[p] {
		case 9:
			goto st158
		case 32:
			goto st158
		case 58:
			goto st159
		}
		goto st0
tr269:
//line msg_parse.rl:47

      mark = p
    
	goto st159
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
//line msg_parse.go:7082
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr269
		case 32:
			goto tr269
		case 269:
			goto tr271
		case 525:
			goto tr272
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr270
		}
		goto st0
tr270:
//line msg_parse.rl:47

      mark = p
    
	goto st160
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
//line msg_parse.go:7115
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st160
		case 269:
			goto tr274
		case 525:
			goto tr275
		}
		if 32 <= _widec && _widec <= 253 {
			goto st160
		}
		goto st0
tr280:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:330

      msg.Warning = string(data[mark:p])
    
	goto st161
tr275:
//line msg_parse.rl:330

      msg.Warning = string(data[mark:p])
    
	goto st161
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
//line msg_parse.go:7156
		if data[p] == 10 {
			goto tr276
		}
		goto st0
tr276:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st162
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
//line msg_parse.go:7170
		switch data[p] {
		case 9:
			goto st160
		case 13:
			goto st15
		case 32:
			goto st160
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr272:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:330

      msg.Warning = string(data[mark:p])
    
	goto st163
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
//line msg_parse.go:7267
		if data[p] == 10 {
			goto tr277
		}
		goto st0
tr277:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st164
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
//line msg_parse.go:7281
		switch data[p] {
		case 9:
			goto st165
		case 13:
			goto st15
		case 32:
			goto st165
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr279:
//line msg_parse.rl:47

      mark = p
    
	goto st165
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
//line msg_parse.go:7374
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr279
		case 32:
			goto tr279
		case 269:
			goto tr271
		case 525:
			goto tr280
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr270
		}
		goto st0
	st166:
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		switch data[p] {
		case 87:
			goto st167
		case 119:
			goto st167
		}
		goto st0
	st167:
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 45 {
			goto st168
		}
		goto st0
	st168:
		if p++; p == pe {
			goto _test_eof168
		}
	st_case_168:
		switch data[p] {
		case 65:
			goto st169
		case 97:
			goto st169
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
	st_case_169:
		switch data[p] {
		case 85:
			goto st170
		case 117:
			goto st170
		}
		goto st0
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
		switch data[p] {
		case 84:
			goto st171
		case 116:
			goto st171
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		switch data[p] {
		case 72:
			goto st172
		case 104:
			goto st172
		}
		goto st0
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
		switch data[p] {
		case 69:
			goto st173
		case 101:
			goto st173
		}
		goto st0
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
		switch data[p] {
		case 78:
			goto st174
		case 110:
			goto st174
		}
		goto st0
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
		switch data[p] {
		case 84:
			goto st175
		case 116:
			goto st175
		}
		goto st0
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
		switch data[p] {
		case 73:
			goto st176
		case 105:
			goto st176
		}
		goto st0
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
		switch data[p] {
		case 67:
			goto st177
		case 99:
			goto st177
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
		case 97:
			goto st178
		}
		goto st0
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
		switch data[p] {
		case 84:
			goto st179
		case 116:
			goto st179
		}
		goto st0
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
		switch data[p] {
		case 69:
			goto st180
		case 101:
			goto st180
		}
		goto st0
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
		switch data[p] {
		case 9:
			goto st180
		case 32:
			goto st180
		case 58:
			goto st181
		}
		goto st0
tr296:
//line msg_parse.rl:47

      mark = p
    
	goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line msg_parse.go:7586
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr296
		case 32:
			goto tr296
		case 269:
			goto tr298
		case 525:
			goto tr299
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr297
		}
		goto st0
tr297:
//line msg_parse.rl:47

      mark = p
    
	goto st182
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
//line msg_parse.go:7619
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st182
		case 269:
			goto tr301
		case 525:
			goto tr302
		}
		if 32 <= _widec && _widec <= 253 {
			goto st182
		}
		goto st0
tr307:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:334

      msg.WWWAuthenticate = string(data[mark:p])
    
	goto st183
tr302:
//line msg_parse.rl:334

      msg.WWWAuthenticate = string(data[mark:p])
    
	goto st183
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
//line msg_parse.go:7660
		if data[p] == 10 {
			goto tr303
		}
		goto st0
tr303:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st184
	st184:
		if p++; p == pe {
			goto _test_eof184
		}
	st_case_184:
//line msg_parse.go:7674
		switch data[p] {
		case 9:
			goto st182
		case 13:
			goto st15
		case 32:
			goto st182
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr299:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:334

      msg.WWWAuthenticate = string(data[mark:p])
    
	goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line msg_parse.go:7771
		if data[p] == 10 {
			goto tr304
		}
		goto st0
tr304:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st186
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
//line msg_parse.go:7785
		switch data[p] {
		case 9:
			goto st187
		case 13:
			goto st15
		case 32:
			goto st187
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr306:
//line msg_parse.rl:47

      mark = p
    
	goto st187
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
	st_case_187:
//line msg_parse.go:7878
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr306
		case 32:
			goto tr306
		case 269:
			goto tr298
		case 525:
			goto tr307
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr297
		}
		goto st0
tr256:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:324

      *viap, err = ParseVia(string(data[mark:p]))
      if err != nil { return nil, err }
      for *viap != nil { viap = &(*viap).Next }
    
	goto st188
	st188:
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
//line msg_parse.go:7917
		if data[p] == 10 {
			goto tr308
		}
		goto st0
tr308:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line msg_parse.go:7931
		switch data[p] {
		case 9:
			goto st190
		case 13:
			goto st15
		case 32:
			goto st190
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr310:
//line msg_parse.rl:47

      mark = p
    
	goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line msg_parse.go:8024
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr310
		case 32:
			goto tr310
		case 269:
			goto tr255
		case 525:
			goto tr311
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr254
		}
		goto st0
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
		switch data[p] {
		case 65:
			goto st147
		case 97:
			goto st147
		}
		goto st0
tr245:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st192
	st192:
		if p++; p == pe {
			goto _test_eof192
		}
	st_case_192:
//line msg_parse.go:8077
		if data[p] == 10 {
			goto tr312
		}
		goto st0
tr312:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st193
	st193:
		if p++; p == pe {
			goto _test_eof193
		}
	st_case_193:
//line msg_parse.go:8091
		switch data[p] {
		case 9:
			goto st194
		case 13:
			goto st15
		case 32:
			goto st194
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr314:
//line msg_parse.rl:47

      mark = p
    
	goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line msg_parse.go:8184
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr314
		case 32:
			goto tr314
		case 269:
			goto tr244
		case 525:
			goto tr315
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr243
		}
		goto st0
	st195:
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		switch data[p] {
		case 83:
			goto st196
		case 115:
			goto st196
		}
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		switch data[p] {
		case 85:
			goto st197
		case 117:
			goto st197
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 80:
			goto st198
		case 112:
			goto st198
		}
		goto st0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 80:
			goto st199
		case 112:
			goto st199
		}
		goto st0
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		switch data[p] {
		case 79:
			goto st200
		case 111:
			goto st200
		}
		goto st0
	st200:
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		switch data[p] {
		case 82:
			goto st201
		case 114:
			goto st201
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
		goto st0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 69:
			goto st203
		case 101:
			goto st203
		}
		goto st0
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		switch data[p] {
		case 68:
			goto st204
		case 100:
			goto st204
		}
		goto st0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		switch data[p] {
		case 9:
			goto st204
		case 32:
			goto st204
		case 58:
			goto st205
		}
		goto st0
tr326:
//line msg_parse.rl:47

      mark = p
    
	goto st205
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
//line msg_parse.go:8339
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr326
		case 32:
			goto tr326
		case 269:
			goto tr328
		case 525:
			goto tr329
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr327
		}
		goto st0
tr327:
//line msg_parse.rl:47

      mark = p
    
	goto st206
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
//line msg_parse.go:8372
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st206
		case 269:
			goto tr331
		case 525:
			goto tr332
		}
		if 32 <= _widec && _widec <= 253 {
			goto st206
		}
		goto st0
tr337:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:316

      msg.Unsupported = string(data[mark:p])
    
	goto st207
tr332:
//line msg_parse.rl:316

      msg.Unsupported = string(data[mark:p])
    
	goto st207
	st207:
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
//line msg_parse.go:8413
		if data[p] == 10 {
			goto tr333
		}
		goto st0
tr333:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st208
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
//line msg_parse.go:8427
		switch data[p] {
		case 9:
			goto st206
		case 13:
			goto st15
		case 32:
			goto st206
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr329:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:316

      msg.Unsupported = string(data[mark:p])
    
	goto st209
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
	st_case_209:
//line msg_parse.go:8524
		if data[p] == 10 {
			goto tr334
		}
		goto st0
tr334:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line msg_parse.go:8538
		switch data[p] {
		case 9:
			goto st211
		case 13:
			goto st15
		case 32:
			goto st211
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr336:
//line msg_parse.rl:47

      mark = p
    
	goto st211
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
//line msg_parse.go:8631
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr336
		case 32:
			goto tr336
		case 269:
			goto tr328
		case 525:
			goto tr337
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr327
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
	st_case_212:
		switch data[p] {
		case 69:
			goto st213
		case 101:
			goto st213
		}
		goto st0
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
		switch data[p] {
		case 82:
			goto st214
		case 114:
			goto st214
		}
		goto st0
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
		if data[p] == 45 {
			goto st215
		}
		goto st0
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
		switch data[p] {
		case 65:
			goto st216
		case 97:
			goto st216
		}
		goto st0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 71:
			goto st217
		case 103:
			goto st217
		}
		goto st0
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
		switch data[p] {
		case 69:
			goto st218
		case 101:
			goto st218
		}
		goto st0
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
		switch data[p] {
		case 78:
			goto st219
		case 110:
			goto st219
		}
		goto st0
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
		switch data[p] {
		case 84:
			goto st220
		case 116:
			goto st220
		}
		goto st0
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
		switch data[p] {
		case 9:
			goto st220
		case 32:
			goto st220
		case 58:
			goto st221
		}
		goto st0
tr347:
//line msg_parse.rl:47

      mark = p
    
	goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line msg_parse.go:8771
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr347
		case 32:
			goto tr347
		case 269:
			goto tr349
		case 525:
			goto tr350
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr348
		}
		goto st0
tr348:
//line msg_parse.rl:47

      mark = p
    
	goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
//line msg_parse.go:8804
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st222
		case 269:
			goto tr352
		case 525:
			goto tr353
		}
		if 32 <= _widec && _widec <= 253 {
			goto st222
		}
		goto st0
tr358:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:320

      msg.UserAgent = string(data[mark:p])
    
	goto st223
tr353:
//line msg_parse.rl:320

      msg.UserAgent = string(data[mark:p])
    
	goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line msg_parse.go:8845
		if data[p] == 10 {
			goto tr354
		}
		goto st0
tr354:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
//line msg_parse.go:8859
		switch data[p] {
		case 9:
			goto st222
		case 13:
			goto st15
		case 32:
			goto st222
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr350:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:320

      msg.UserAgent = string(data[mark:p])
    
	goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line msg_parse.go:8956
		if data[p] == 10 {
			goto tr355
		}
		goto st0
tr355:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
//line msg_parse.go:8970
		switch data[p] {
		case 9:
			goto st227
		case 13:
			goto st15
		case 32:
			goto st227
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr357:
//line msg_parse.rl:47

      mark = p
    
	goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
//line msg_parse.go:9063
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr357
		case 32:
			goto tr357
		case 269:
			goto tr349
		case 525:
			goto tr358
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr348
		}
		goto st0
tr233:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:311

      msg.To, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line msg_parse.go:9101
		if data[p] == 10 {
			goto tr359
		}
		goto st0
tr359:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st229
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
	st_case_229:
//line msg_parse.go:9115
		switch data[p] {
		case 9:
			goto st230
		case 13:
			goto st15
		case 32:
			goto st230
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr361:
//line msg_parse.rl:47

      mark = p
    
	goto st230
	st230:
		if p++; p == pe {
			goto _test_eof230
		}
	st_case_230:
//line msg_parse.go:9208
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr361
		case 32:
			goto tr361
		case 269:
			goto tr232
		case 525:
			goto tr362
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr231
		}
		goto st0
	st231:
		if p++; p == pe {
			goto _test_eof231
		}
	st_case_231:
		switch data[p] {
		case 77:
			goto st232
		case 109:
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
		case 84:
			goto st235
		case 116:
			goto st235
		}
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		switch data[p] {
		case 65:
			goto st236
		case 97:
			goto st236
		}
		goto st0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 77:
			goto st237
		case 109:
			goto st237
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 80:
			goto st238
		case 112:
			goto st238
		}
		goto st0
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		switch data[p] {
		case 9:
			goto st238
		case 32:
			goto st238
		case 58:
			goto st239
		}
		goto st0
tr371:
//line msg_parse.rl:47

      mark = p
    
	goto st239
	st239:
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
//line msg_parse.go:9339
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr371
		case 32:
			goto tr371
		case 269:
			goto tr373
		case 525:
			goto tr374
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr372
		}
		goto st0
tr372:
//line msg_parse.rl:47

      mark = p
    
	goto st240
	st240:
		if p++; p == pe {
			goto _test_eof240
		}
	st_case_240:
//line msg_parse.go:9372
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st240
		case 269:
			goto tr376
		case 525:
			goto tr377
		}
		if 32 <= _widec && _widec <= 253 {
			goto st240
		}
		goto st0
tr382:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:307

      msg.Timestamp = string(data[mark:p])
    
	goto st241
tr377:
//line msg_parse.rl:307

      msg.Timestamp = string(data[mark:p])
    
	goto st241
	st241:
		if p++; p == pe {
			goto _test_eof241
		}
	st_case_241:
//line msg_parse.go:9413
		if data[p] == 10 {
			goto tr378
		}
		goto st0
tr378:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st242
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
//line msg_parse.go:9427
		switch data[p] {
		case 9:
			goto st240
		case 13:
			goto st15
		case 32:
			goto st240
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr374:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:307

      msg.Timestamp = string(data[mark:p])
    
	goto st243
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
//line msg_parse.go:9524
		if data[p] == 10 {
			goto tr379
		}
		goto st0
tr379:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st244
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
	st_case_244:
//line msg_parse.go:9538
		switch data[p] {
		case 9:
			goto st245
		case 13:
			goto st15
		case 32:
			goto st245
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr381:
//line msg_parse.rl:47

      mark = p
    
	goto st245
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
//line msg_parse.go:9631
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr381
		case 32:
			goto tr381
		case 269:
			goto tr373
		case 525:
			goto tr382
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr372
		}
		goto st0
tr222:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:299

      msg.Subject = string(data[mark:p])
    
	goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
//line msg_parse.go:9668
		if data[p] == 10 {
			goto tr383
		}
		goto st0
tr383:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line msg_parse.go:9682
		switch data[p] {
		case 9:
			goto st248
		case 13:
			goto st15
		case 32:
			goto st248
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr385:
//line msg_parse.rl:47

      mark = p
    
	goto st248
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
//line msg_parse.go:9775
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr385
		case 32:
			goto tr385
		case 269:
			goto tr221
		case 525:
			goto tr386
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr220
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
	st_case_249:
		switch data[p] {
		case 82:
			goto st250
		case 114:
			goto st250
		}
		goto st0
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
		switch data[p] {
		case 86:
			goto st251
		case 118:
			goto st251
		}
		goto st0
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
		switch data[p] {
		case 69:
			goto st252
		case 101:
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
		case 9:
			goto st253
		case 32:
			goto st253
		case 58:
			goto st254
		}
		goto st0
tr392:
//line msg_parse.rl:47

      mark = p
    
	goto st254
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
	st_case_254:
//line msg_parse.go:9870
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr392
		case 32:
			goto tr392
		case 269:
			goto tr394
		case 525:
			goto tr395
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr393
		}
		goto st0
tr393:
//line msg_parse.rl:47

      mark = p
    
	goto st255
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
//line msg_parse.go:9903
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st255
		case 269:
			goto tr397
		case 525:
			goto tr398
		}
		if 32 <= _widec && _widec <= 253 {
			goto st255
		}
		goto st0
tr403:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:295

      msg.Server = string(data[mark:p])
    
	goto st256
tr398:
//line msg_parse.rl:295

      msg.Server = string(data[mark:p])
    
	goto st256
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
	st_case_256:
//line msg_parse.go:9944
		if data[p] == 10 {
			goto tr399
		}
		goto st0
tr399:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st257
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
//line msg_parse.go:9958
		switch data[p] {
		case 9:
			goto st255
		case 13:
			goto st15
		case 32:
			goto st255
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr395:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:295

      msg.Server = string(data[mark:p])
    
	goto st258
	st258:
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
//line msg_parse.go:10055
		if data[p] == 10 {
			goto tr400
		}
		goto st0
tr400:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st259
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
//line msg_parse.go:10069
		switch data[p] {
		case 9:
			goto st260
		case 13:
			goto st15
		case 32:
			goto st260
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr402:
//line msg_parse.rl:47

      mark = p
    
	goto st260
	st260:
		if p++; p == pe {
			goto _test_eof260
		}
	st_case_260:
//line msg_parse.go:10162
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr402
		case 32:
			goto tr402
		case 269:
			goto tr394
		case 525:
			goto tr403
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr393
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
	st_case_261:
		switch data[p] {
		case 66:
			goto st262
		case 80:
			goto st266
		case 98:
			goto st262
		case 112:
			goto st266
		}
		goto st0
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
		switch data[p] {
		case 74:
			goto st263
		case 106:
			goto st263
		}
		goto st0
	st263:
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		switch data[p] {
		case 69:
			goto st264
		case 101:
			goto st264
		}
		goto st0
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
		switch data[p] {
		case 67:
			goto st265
		case 99:
			goto st265
		}
		goto st0
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
		switch data[p] {
		case 84:
			goto st129
		case 116:
			goto st129
		}
		goto st0
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
		switch data[p] {
		case 80:
			goto st267
		case 112:
			goto st267
		}
		goto st0
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
		switch data[p] {
		case 79:
			goto st268
		case 111:
			goto st268
		}
		goto st0
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
		switch data[p] {
		case 82:
			goto st269
		case 114:
			goto st269
		}
		goto st0
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
		switch data[p] {
		case 84:
			goto st270
		case 116:
			goto st270
		}
		goto st0
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
		switch data[p] {
		case 69:
			goto st271
		case 101:
			goto st271
		}
		goto st0
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
		switch data[p] {
		case 68:
			goto st76
		case 100:
			goto st76
		}
		goto st0
tr210:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:268

      msg.ReferTo = string(data[mark:p])
    
	goto st272
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
//line msg_parse.go:10335
		if data[p] == 10 {
			goto tr414
		}
		goto st0
tr414:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st273
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
	st_case_273:
//line msg_parse.go:10349
		switch data[p] {
		case 9:
			goto st274
		case 13:
			goto st15
		case 32:
			goto st274
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr416:
//line msg_parse.rl:47

      mark = p
    
	goto st274
	st274:
		if p++; p == pe {
			goto _test_eof274
		}
	st_case_274:
//line msg_parse.go:10442
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr416
		case 32:
			goto tr416
		case 269:
			goto tr209
		case 525:
			goto tr417
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr208
		}
		goto st0
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
		switch data[p] {
		case 67:
			goto st276
		case 70:
			goto st293
		case 77:
			goto st303
		case 80:
			goto st323
		case 81:
			goto st336
		case 84:
			goto st348
		case 99:
			goto st276
		case 102:
			goto st293
		case 109:
			goto st303
		case 112:
			goto st323
		case 113:
			goto st336
		case 116:
			goto st348
		}
		goto st0
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
		switch data[p] {
		case 79:
			goto st277
		case 111:
			goto st277
		}
		goto st0
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		switch data[p] {
		case 82:
			goto st278
		case 114:
			goto st278
		}
		goto st0
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 68:
			goto st279
		case 100:
			goto st279
		}
		goto st0
	st279:
		if p++; p == pe {
			goto _test_eof279
		}
	st_case_279:
		if data[p] == 45 {
			goto st280
		}
		goto st0
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		switch data[p] {
		case 82:
			goto st281
		case 114:
			goto st281
		}
		goto st0
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
		switch data[p] {
		case 79:
			goto st282
		case 111:
			goto st282
		}
		goto st0
	st282:
		if p++; p == pe {
			goto _test_eof282
		}
	st_case_282:
		switch data[p] {
		case 85:
			goto st283
		case 117:
			goto st283
		}
		goto st0
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
		switch data[p] {
		case 84:
			goto st284
		case 116:
			goto st284
		}
		goto st0
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
		switch data[p] {
		case 69:
			goto st285
		case 101:
			goto st285
		}
		goto st0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		switch data[p] {
		case 9:
			goto st285
		case 32:
			goto st285
		case 58:
			goto st286
		}
		goto st0
tr434:
//line msg_parse.rl:47

      mark = p
    
	goto st286
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
//line msg_parse.go:10626
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr434
		case 32:
			goto tr434
		case 269:
			goto tr436
		case 525:
			goto tr437
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr435
		}
		goto st0
tr435:
//line msg_parse.rl:47

      mark = p
    
	goto st287
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
//line msg_parse.go:10659
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st287
		case 269:
			goto tr439
		case 525:
			goto tr440
		}
		if 32 <= _widec && _widec <= 253 {
			goto st287
		}
		goto st0
tr445:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:262

      *rroutep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *rroutep != nil { rroutep = &(*rroutep).Next }
    
	goto st288
tr440:
//line msg_parse.rl:262

      *rroutep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *rroutep != nil { rroutep = &(*rroutep).Next }
    
	goto st288
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
//line msg_parse.go:10704
		if data[p] == 10 {
			goto tr441
		}
		goto st0
tr441:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st289
	st289:
		if p++; p == pe {
			goto _test_eof289
		}
	st_case_289:
//line msg_parse.go:10718
		switch data[p] {
		case 9:
			goto st287
		case 13:
			goto st15
		case 32:
			goto st287
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr437:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:262

      *rroutep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *rroutep != nil { rroutep = &(*rroutep).Next }
    
	goto st290
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
//line msg_parse.go:10817
		if data[p] == 10 {
			goto tr442
		}
		goto st0
tr442:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st291
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
//line msg_parse.go:10831
		switch data[p] {
		case 9:
			goto st292
		case 13:
			goto st15
		case 32:
			goto st292
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr444:
//line msg_parse.rl:47

      mark = p
    
	goto st292
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
	st_case_292:
//line msg_parse.go:10924
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr444
		case 32:
			goto tr444
		case 269:
			goto tr436
		case 525:
			goto tr445
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr435
		}
		goto st0
	st293:
		if p++; p == pe {
			goto _test_eof293
		}
	st_case_293:
		switch data[p] {
		case 69:
			goto st294
		case 101:
			goto st294
		}
		goto st0
	st294:
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 82:
			goto st295
		case 114:
			goto st295
		}
		goto st0
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 45:
			goto st296
		case 82:
			goto st298
		case 114:
			goto st298
		}
		goto st0
	st296:
		if p++; p == pe {
			goto _test_eof296
		}
	st_case_296:
		switch data[p] {
		case 84:
			goto st297
		case 116:
			goto st297
		}
		goto st0
	st297:
		if p++; p == pe {
			goto _test_eof297
		}
	st_case_297:
		switch data[p] {
		case 79:
			goto st123
		case 111:
			goto st123
		}
		goto st0
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 69:
			goto st299
		case 101:
			goto st299
		}
		goto st0
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 68:
			goto st300
		case 100:
			goto st300
		}
		goto st0
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
	st_case_300:
		if data[p] == 45 {
			goto st301
		}
		goto st0
	st301:
		if p++; p == pe {
			goto _test_eof301
		}
	st_case_301:
		switch data[p] {
		case 66:
			goto st302
		case 98:
			goto st302
		}
		goto st0
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 89:
			goto st22
		case 121:
			goto st22
		}
		goto st0
	st303:
		if p++; p == pe {
			goto _test_eof303
		}
	st_case_303:
		switch data[p] {
		case 79:
			goto st304
		case 111:
			goto st304
		}
		goto st0
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 84:
			goto st305
		case 116:
			goto st305
		}
		goto st0
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		switch data[p] {
		case 69:
			goto st306
		case 101:
			goto st306
		}
		goto st0
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		if data[p] == 45 {
			goto st307
		}
		goto st0
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		switch data[p] {
		case 80:
			goto st308
		case 112:
			goto st308
		}
		goto st0
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
	st_case_308:
		switch data[p] {
		case 65:
			goto st309
		case 97:
			goto st309
		}
		goto st0
	st309:
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		switch data[p] {
		case 82:
			goto st310
		case 114:
			goto st310
		}
		goto st0
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 84:
			goto st311
		case 116:
			goto st311
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
		goto st0
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		if data[p] == 45 {
			goto st313
		}
		goto st0
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 73:
			goto st314
		case 105:
			goto st314
		}
		goto st0
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 68:
			goto st315
		case 100:
			goto st315
		}
		goto st0
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 9:
			goto st315
		case 32:
			goto st315
		case 58:
			goto st316
		}
		goto st0
tr468:
//line msg_parse.rl:47

      mark = p
    
	goto st316
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
	st_case_316:
//line msg_parse.go:11228
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr468
		case 32:
			goto tr468
		case 269:
			goto tr470
		case 525:
			goto tr471
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr469
		}
		goto st0
tr469:
//line msg_parse.rl:47

      mark = p
    
	goto st317
	st317:
		if p++; p == pe {
			goto _test_eof317
		}
	st_case_317:
//line msg_parse.go:11261
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
		case 269:
			goto tr473
		case 525:
			goto tr474
		}
		if 32 <= _widec && _widec <= 253 {
			goto st317
		}
		goto st0
tr479:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:276

      msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st318
tr474:
//line msg_parse.rl:276

      msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st318
	st318:
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
//line msg_parse.go:11304
		if data[p] == 10 {
			goto tr475
		}
		goto st0
tr475:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st319
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
//line msg_parse.go:11318
		switch data[p] {
		case 9:
			goto st317
		case 13:
			goto st15
		case 32:
			goto st317
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr471:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:276

      msg.RemotePartyID, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st320
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
//line msg_parse.go:11416
		if data[p] == 10 {
			goto tr476
		}
		goto st0
tr476:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st321
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
//line msg_parse.go:11430
		switch data[p] {
		case 9:
			goto st322
		case 13:
			goto st15
		case 32:
			goto st322
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr478:
//line msg_parse.rl:47

      mark = p
    
	goto st322
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
//line msg_parse.go:11523
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr478
		case 32:
			goto tr478
		case 269:
			goto tr470
		case 525:
			goto tr479
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr469
		}
		goto st0
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 76:
			goto st324
		case 108:
			goto st324
		}
		goto st0
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 89:
			goto st325
		case 121:
			goto st325
		}
		goto st0
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
	st_case_325:
		if data[p] == 45 {
			goto st326
		}
		goto st0
	st326:
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 84:
			goto st327
		case 116:
			goto st327
		}
		goto st0
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 79:
			goto st328
		case 111:
			goto st328
		}
		goto st0
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
	st_case_328:
		switch data[p] {
		case 9:
			goto st328
		case 32:
			goto st328
		case 58:
			goto st329
		}
		goto st0
tr486:
//line msg_parse.rl:47

      mark = p
    
	goto st329
	st329:
		if p++; p == pe {
			goto _test_eof329
		}
	st_case_329:
//line msg_parse.go:11627
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr486
		case 32:
			goto tr486
		case 269:
			goto tr488
		case 525:
			goto tr489
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr487
		}
		goto st0
tr487:
//line msg_parse.rl:47

      mark = p
    
	goto st330
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
//line msg_parse.go:11660
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st330
		case 269:
			goto tr491
		case 525:
			goto tr492
		}
		if 32 <= _widec && _widec <= 253 {
			goto st330
		}
		goto st0
tr497:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:229

      msg.ReplyTo = string(data[mark:p])
    
	goto st331
tr492:
//line msg_parse.rl:229

      msg.ReplyTo = string(data[mark:p])
    
	goto st331
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
//line msg_parse.go:11701
		if data[p] == 10 {
			goto tr493
		}
		goto st0
tr493:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st332
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
//line msg_parse.go:11715
		switch data[p] {
		case 9:
			goto st330
		case 13:
			goto st15
		case 32:
			goto st330
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr489:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:229

      msg.ReplyTo = string(data[mark:p])
    
	goto st333
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
//line msg_parse.go:11812
		if data[p] == 10 {
			goto tr494
		}
		goto st0
tr494:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st334
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
//line msg_parse.go:11826
		switch data[p] {
		case 9:
			goto st335
		case 13:
			goto st15
		case 32:
			goto st335
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr496:
//line msg_parse.rl:47

      mark = p
    
	goto st335
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
	st_case_335:
//line msg_parse.go:11919
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr496
		case 32:
			goto tr496
		case 269:
			goto tr488
		case 525:
			goto tr497
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr487
		}
		goto st0
	st336:
		if p++; p == pe {
			goto _test_eof336
		}
	st_case_336:
		switch data[p] {
		case 85:
			goto st337
		case 117:
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
		case 82:
			goto st339
		case 114:
			goto st339
		}
		goto st0
	st339:
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 69:
			goto st340
		case 101:
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
tr503:
//line msg_parse.rl:47

      mark = p
    
	goto st341
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
//line msg_parse.go:12014
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr503
		case 32:
			goto tr503
		case 269:
			goto tr505
		case 525:
			goto tr506
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr504
		}
		goto st0
tr504:
//line msg_parse.rl:47

      mark = p
    
	goto st342
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
//line msg_parse.go:12047
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
			goto tr508
		case 525:
			goto tr509
		}
		if 32 <= _widec && _widec <= 253 {
			goto st342
		}
		goto st0
tr514:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:281

      msg.Require = string(data[mark:p])
    
	goto st343
tr509:
//line msg_parse.rl:281

      msg.Require = string(data[mark:p])
    
	goto st343
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
//line msg_parse.go:12088
		if data[p] == 10 {
			goto tr510
		}
		goto st0
tr510:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st344
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
//line msg_parse.go:12102
		switch data[p] {
		case 9:
			goto st342
		case 13:
			goto st15
		case 32:
			goto st342
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr506:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:281

      msg.Require = string(data[mark:p])
    
	goto st345
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
//line msg_parse.go:12199
		if data[p] == 10 {
			goto tr511
		}
		goto st0
tr511:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st346
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
//line msg_parse.go:12213
		switch data[p] {
		case 9:
			goto st347
		case 13:
			goto st15
		case 32:
			goto st347
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr513:
//line msg_parse.rl:47

      mark = p
    
	goto st347
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
//line msg_parse.go:12306
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr513
		case 32:
			goto tr513
		case 269:
			goto tr505
		case 525:
			goto tr514
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr504
		}
		goto st0
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		switch data[p] {
		case 82:
			goto st349
		case 114:
			goto st349
		}
		goto st0
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 89:
			goto st350
		case 121:
			goto st350
		}
		goto st0
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		if data[p] == 45 {
			goto st351
		}
		goto st0
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		switch data[p] {
		case 65:
			goto st352
		case 97:
			goto st352
		}
		goto st0
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		switch data[p] {
		case 70:
			goto st353
		case 102:
			goto st353
		}
		goto st0
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
	st_case_353:
		switch data[p] {
		case 84:
			goto st354
		case 116:
			goto st354
		}
		goto st0
	st354:
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 69:
			goto st355
		case 101:
			goto st355
		}
		goto st0
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 82:
			goto st356
		case 114:
			goto st356
		}
		goto st0
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 9:
			goto st356
		case 32:
			goto st356
		case 58:
			goto st357
		}
		goto st0
tr524:
//line msg_parse.rl:47

      mark = p
    
	goto st357
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
//line msg_parse.go:12446
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr524
		case 32:
			goto tr524
		case 269:
			goto tr526
		case 525:
			goto tr527
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr525
		}
		goto st0
tr525:
//line msg_parse.rl:47

      mark = p
    
	goto st358
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
//line msg_parse.go:12479
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st358
		case 269:
			goto tr529
		case 525:
			goto tr530
		}
		if 32 <= _widec && _widec <= 253 {
			goto st358
		}
		goto st0
tr535:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:285

      msg.RetryAfter = string(data[mark:p])
    
	goto st359
tr530:
//line msg_parse.rl:285

      msg.RetryAfter = string(data[mark:p])
    
	goto st359
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
//line msg_parse.go:12520
		if data[p] == 10 {
			goto tr531
		}
		goto st0
tr531:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st360
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
//line msg_parse.go:12534
		switch data[p] {
		case 9:
			goto st358
		case 13:
			goto st15
		case 32:
			goto st358
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr527:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:285

      msg.RetryAfter = string(data[mark:p])
    
	goto st361
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
//line msg_parse.go:12631
		if data[p] == 10 {
			goto tr532
		}
		goto st0
tr532:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st362
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
	st_case_362:
//line msg_parse.go:12645
		switch data[p] {
		case 9:
			goto st363
		case 13:
			goto st15
		case 32:
			goto st363
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr534:
//line msg_parse.rl:47

      mark = p
    
	goto st363
	st363:
		if p++; p == pe {
			goto _test_eof363
		}
	st_case_363:
//line msg_parse.go:12738
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr534
		case 32:
			goto tr534
		case 269:
			goto tr526
		case 525:
			goto tr535
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr525
		}
		goto st0
	st364:
		if p++; p == pe {
			goto _test_eof364
		}
	st_case_364:
		switch data[p] {
		case 85:
			goto st365
		case 117:
			goto st365
		}
		goto st0
	st365:
		if p++; p == pe {
			goto _test_eof365
		}
	st_case_365:
		switch data[p] {
		case 84:
			goto st366
		case 116:
			goto st366
		}
		goto st0
	st366:
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		switch data[p] {
		case 69:
			goto st367
		case 101:
			goto st367
		}
		goto st0
	st367:
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		switch data[p] {
		case 9:
			goto st367
		case 32:
			goto st367
		case 58:
			goto st368
		}
		goto st0
tr540:
//line msg_parse.rl:47

      mark = p
    
	goto st368
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
//line msg_parse.go:12821
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr540
		case 32:
			goto tr540
		case 269:
			goto tr542
		case 525:
			goto tr543
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr541
		}
		goto st0
tr541:
//line msg_parse.rl:47

      mark = p
    
	goto st369
	st369:
		if p++; p == pe {
			goto _test_eof369
		}
	st_case_369:
//line msg_parse.go:12854
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st369
		case 269:
			goto tr545
		case 525:
			goto tr546
		}
		if 32 <= _widec && _widec <= 253 {
			goto st369
		}
		goto st0
tr551:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:289

      *routep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *routep != nil { routep = &(*routep).Next }
    
	goto st370
tr546:
//line msg_parse.rl:289

      *routep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *routep != nil { routep = &(*routep).Next }
    
	goto st370
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
//line msg_parse.go:12899
		if data[p] == 10 {
			goto tr547
		}
		goto st0
tr547:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st371
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
//line msg_parse.go:12913
		switch data[p] {
		case 9:
			goto st369
		case 13:
			goto st15
		case 32:
			goto st369
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr543:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:289

      *routep, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *routep != nil { routep = &(*routep).Next }
    
	goto st372
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
//line msg_parse.go:13012
		if data[p] == 10 {
			goto tr548
		}
		goto st0
tr548:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st373
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
//line msg_parse.go:13026
		switch data[p] {
		case 9:
			goto st374
		case 13:
			goto st15
		case 32:
			goto st374
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr550:
//line msg_parse.rl:47

      mark = p
    
	goto st374
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
	st_case_374:
//line msg_parse.go:13119
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr550
		case 32:
			goto tr550
		case 269:
			goto tr542
		case 525:
			goto tr551
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr541
		}
		goto st0
tr198:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:241

      msg.PAssertedIdentity, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st375
	st375:
		if p++; p == pe {
			goto _test_eof375
		}
	st_case_375:
//line msg_parse.go:13157
		if data[p] == 10 {
			goto tr552
		}
		goto st0
tr552:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st376
	st376:
		if p++; p == pe {
			goto _test_eof376
		}
	st_case_376:
//line msg_parse.go:13171
		switch data[p] {
		case 9:
			goto st377
		case 13:
			goto st15
		case 32:
			goto st377
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr554:
//line msg_parse.rl:47

      mark = p
    
	goto st377
	st377:
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
//line msg_parse.go:13264
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr554
		case 32:
			goto tr554
		case 269:
			goto tr197
		case 525:
			goto tr555
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr196
		}
		goto st0
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
		switch data[p] {
		case 73:
			goto st379
		case 79:
			goto st392
		case 105:
			goto st379
		case 111:
			goto st392
		}
		goto st0
	st379:
		if p++; p == pe {
			goto _test_eof379
		}
	st_case_379:
		switch data[p] {
		case 79:
			goto st380
		case 111:
			goto st380
		}
		goto st0
	st380:
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		switch data[p] {
		case 82:
			goto st381
		case 114:
			goto st381
		}
		goto st0
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
		switch data[p] {
		case 73:
			goto st382
		case 105:
			goto st382
		}
		goto st0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		switch data[p] {
		case 84:
			goto st383
		case 116:
			goto st383
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 89:
			goto st384
		case 121:
			goto st384
		}
		goto st0
	st384:
		if p++; p == pe {
			goto _test_eof384
		}
	st_case_384:
		switch data[p] {
		case 9:
			goto st384
		case 32:
			goto st384
		case 58:
			goto st385
		}
		goto st0
tr564:
//line msg_parse.rl:47

      mark = p
    
	goto st385
	st385:
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
//line msg_parse.go:13387
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr564
		case 32:
			goto tr564
		case 269:
			goto tr566
		case 525:
			goto tr567
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr565
		}
		goto st0
tr565:
//line msg_parse.rl:47

      mark = p
    
	goto st386
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
//line msg_parse.go:13420
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st386
		case 269:
			goto tr569
		case 525:
			goto tr570
		}
		if 32 <= _widec && _widec <= 253 {
			goto st386
		}
		goto st0
tr575:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:246

      msg.Priority = string(data[mark:p])
    
	goto st387
tr570:
//line msg_parse.rl:246

      msg.Priority = string(data[mark:p])
    
	goto st387
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
	st_case_387:
//line msg_parse.go:13461
		if data[p] == 10 {
			goto tr571
		}
		goto st0
tr571:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st388
	st388:
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
//line msg_parse.go:13475
		switch data[p] {
		case 9:
			goto st386
		case 13:
			goto st15
		case 32:
			goto st386
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr567:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:246

      msg.Priority = string(data[mark:p])
    
	goto st389
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
//line msg_parse.go:13572
		if data[p] == 10 {
			goto tr572
		}
		goto st0
tr572:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st390
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
	st_case_390:
//line msg_parse.go:13586
		switch data[p] {
		case 9:
			goto st391
		case 13:
			goto st15
		case 32:
			goto st391
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr574:
//line msg_parse.rl:47

      mark = p
    
	goto st391
	st391:
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
//line msg_parse.go:13679
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr574
		case 32:
			goto tr574
		case 269:
			goto tr566
		case 525:
			goto tr575
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr565
		}
		goto st0
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 88:
			goto st393
		case 120:
			goto st393
		}
		goto st0
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
	st_case_393:
		switch data[p] {
		case 89:
			goto st394
		case 121:
			goto st394
		}
		goto st0
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
		if data[p] == 45 {
			goto st395
		}
		goto st0
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
		switch data[p] {
		case 65:
			goto st396
		case 82:
			goto st431
		case 97:
			goto st396
		case 114:
			goto st431
		}
		goto st0
	st396:
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		switch data[p] {
		case 85:
			goto st397
		case 117:
			goto st397
		}
		goto st0
	st397:
		if p++; p == pe {
			goto _test_eof397
		}
	st_case_397:
		switch data[p] {
		case 84:
			goto st398
		case 116:
			goto st398
		}
		goto st0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 72:
			goto st399
		case 104:
			goto st399
		}
		goto st0
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 69:
			goto st400
		case 79:
			goto st415
		case 101:
			goto st400
		case 111:
			goto st415
		}
		goto st0
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 78:
			goto st401
		case 110:
			goto st401
		}
		goto st0
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		switch data[p] {
		case 84:
			goto st402
		case 116:
			goto st402
		}
		goto st0
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
		goto st0
	st403:
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 67:
			goto st404
		case 99:
			goto st404
		}
		goto st0
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 65:
			goto st405
		case 97:
			goto st405
		}
		goto st0
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 84:
			goto st406
		case 116:
			goto st406
		}
		goto st0
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		switch data[p] {
		case 69:
			goto st407
		case 101:
			goto st407
		}
		goto st0
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 9:
			goto st407
		case 32:
			goto st407
		case 58:
			goto st408
		}
		goto st0
tr594:
//line msg_parse.rl:47

      mark = p
    
	goto st408
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
//line msg_parse.go:13911
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr594
		case 32:
			goto tr594
		case 269:
			goto tr596
		case 525:
			goto tr597
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr595
		}
		goto st0
tr595:
//line msg_parse.rl:47

      mark = p
    
	goto st409
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
//line msg_parse.go:13944
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st409
		case 269:
			goto tr599
		case 525:
			goto tr600
		}
		if 32 <= _widec && _widec <= 253 {
			goto st409
		}
		goto st0
tr605:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:250

      msg.ProxyAuthenticate = string(data[mark:p])
    
	goto st410
tr600:
//line msg_parse.rl:250

      msg.ProxyAuthenticate = string(data[mark:p])
    
	goto st410
	st410:
		if p++; p == pe {
			goto _test_eof410
		}
	st_case_410:
//line msg_parse.go:13985
		if data[p] == 10 {
			goto tr601
		}
		goto st0
tr601:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st411
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
//line msg_parse.go:13999
		switch data[p] {
		case 9:
			goto st409
		case 13:
			goto st15
		case 32:
			goto st409
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr597:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:250

      msg.ProxyAuthenticate = string(data[mark:p])
    
	goto st412
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
//line msg_parse.go:14096
		if data[p] == 10 {
			goto tr602
		}
		goto st0
tr602:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st413
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
	st_case_413:
//line msg_parse.go:14110
		switch data[p] {
		case 9:
			goto st414
		case 13:
			goto st15
		case 32:
			goto st414
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr604:
//line msg_parse.rl:47

      mark = p
    
	goto st414
	st414:
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
//line msg_parse.go:14203
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr604
		case 32:
			goto tr604
		case 269:
			goto tr596
		case 525:
			goto tr605
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr595
		}
		goto st0
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 82:
			goto st416
		case 114:
			goto st416
		}
		goto st0
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 73:
			goto st417
		case 105:
			goto st417
		}
		goto st0
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 90:
			goto st418
		case 122:
			goto st418
		}
		goto st0
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		switch data[p] {
		case 65:
			goto st419
		case 97:
			goto st419
		}
		goto st0
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 84:
			goto st420
		case 116:
			goto st420
		}
		goto st0
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		switch data[p] {
		case 73:
			goto st421
		case 105:
			goto st421
		}
		goto st0
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
	st_case_421:
		switch data[p] {
		case 79:
			goto st422
		case 111:
			goto st422
		}
		goto st0
	st422:
		if p++; p == pe {
			goto _test_eof422
		}
	st_case_422:
		switch data[p] {
		case 78:
			goto st423
		case 110:
			goto st423
		}
		goto st0
	st423:
		if p++; p == pe {
			goto _test_eof423
		}
	st_case_423:
		switch data[p] {
		case 9:
			goto st423
		case 32:
			goto st423
		case 58:
			goto st424
		}
		goto st0
tr615:
//line msg_parse.rl:47

      mark = p
    
	goto st424
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
//line msg_parse.go:14346
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr615
		case 32:
			goto tr615
		case 269:
			goto tr617
		case 525:
			goto tr618
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr616
		}
		goto st0
tr616:
//line msg_parse.rl:47

      mark = p
    
	goto st425
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
//line msg_parse.go:14379
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st425
		case 269:
			goto tr620
		case 525:
			goto tr621
		}
		if 32 <= _widec && _widec <= 253 {
			goto st425
		}
		goto st0
tr626:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:254

      msg.ProxyAuthorization = string(data[mark:p])
    
	goto st426
tr621:
//line msg_parse.rl:254

      msg.ProxyAuthorization = string(data[mark:p])
    
	goto st426
	st426:
		if p++; p == pe {
			goto _test_eof426
		}
	st_case_426:
//line msg_parse.go:14420
		if data[p] == 10 {
			goto tr622
		}
		goto st0
tr622:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st427
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
//line msg_parse.go:14434
		switch data[p] {
		case 9:
			goto st425
		case 13:
			goto st15
		case 32:
			goto st425
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr618:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:254

      msg.ProxyAuthorization = string(data[mark:p])
    
	goto st428
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
//line msg_parse.go:14531
		if data[p] == 10 {
			goto tr623
		}
		goto st0
tr623:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st429
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
	st_case_429:
//line msg_parse.go:14545
		switch data[p] {
		case 9:
			goto st430
		case 13:
			goto st15
		case 32:
			goto st430
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr625:
//line msg_parse.rl:47

      mark = p
    
	goto st430
	st430:
		if p++; p == pe {
			goto _test_eof430
		}
	st_case_430:
//line msg_parse.go:14638
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr625
		case 32:
			goto tr625
		case 269:
			goto tr617
		case 525:
			goto tr626
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr616
		}
		goto st0
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 69:
			goto st432
		case 101:
			goto st432
		}
		goto st0
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 81:
			goto st433
		case 113:
			goto st433
		}
		goto st0
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 85:
			goto st434
		case 117:
			goto st434
		}
		goto st0
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
	st_case_434:
		switch data[p] {
		case 73:
			goto st435
		case 105:
			goto st435
		}
		goto st0
	st435:
		if p++; p == pe {
			goto _test_eof435
		}
	st_case_435:
		switch data[p] {
		case 82:
			goto st436
		case 114:
			goto st436
		}
		goto st0
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
		goto st0
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
		switch data[p] {
		case 9:
			goto st437
		case 32:
			goto st437
		case 58:
			goto st438
		}
		goto st0
tr634:
//line msg_parse.rl:47

      mark = p
    
	goto st438
	st438:
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
//line msg_parse.go:14757
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr634
		case 32:
			goto tr634
		case 269:
			goto tr636
		case 525:
			goto tr637
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr635
		}
		goto st0
tr635:
//line msg_parse.rl:47

      mark = p
    
	goto st439
	st439:
		if p++; p == pe {
			goto _test_eof439
		}
	st_case_439:
//line msg_parse.go:14790
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st439
		case 269:
			goto tr639
		case 525:
			goto tr640
		}
		if 32 <= _widec && _widec <= 253 {
			goto st439
		}
		goto st0
tr645:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:258

      msg.ProxyRequire = string(data[mark:p])
    
	goto st440
tr640:
//line msg_parse.rl:258

      msg.ProxyRequire = string(data[mark:p])
    
	goto st440
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
	st_case_440:
//line msg_parse.go:14831
		if data[p] == 10 {
			goto tr641
		}
		goto st0
tr641:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st441
	st441:
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
//line msg_parse.go:14845
		switch data[p] {
		case 9:
			goto st439
		case 13:
			goto st15
		case 32:
			goto st439
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr637:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:258

      msg.ProxyRequire = string(data[mark:p])
    
	goto st442
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
//line msg_parse.go:14942
		if data[p] == 10 {
			goto tr642
		}
		goto st0
tr642:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st443
	st443:
		if p++; p == pe {
			goto _test_eof443
		}
	st_case_443:
//line msg_parse.go:14956
		switch data[p] {
		case 9:
			goto st444
		case 13:
			goto st15
		case 32:
			goto st444
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr644:
//line msg_parse.rl:47

      mark = p
    
	goto st444
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
	st_case_444:
//line msg_parse.go:15049
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr644
		case 32:
			goto tr644
		case 269:
			goto tr636
		case 525:
			goto tr645
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr635
		}
		goto st0
tr170:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:200

      msg.Event = string(data[mark:p])
    
	goto st445
	st445:
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
//line msg_parse.go:15086
		if data[p] == 10 {
			goto tr646
		}
		goto st0
tr646:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st446
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
//line msg_parse.go:15100
		switch data[p] {
		case 9:
			goto st447
		case 13:
			goto st15
		case 32:
			goto st447
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr648:
//line msg_parse.rl:47

      mark = p
    
	goto st447
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
	st_case_447:
//line msg_parse.go:15193
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr648
		case 32:
			goto tr648
		case 269:
			goto tr169
		case 525:
			goto tr649
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr168
		}
		goto st0
	st448:
		if p++; p == pe {
			goto _test_eof448
		}
	st_case_448:
		switch data[p] {
		case 71:
			goto st449
		case 103:
			goto st449
		}
		goto st0
	st449:
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		switch data[p] {
		case 65:
			goto st450
		case 97:
			goto st450
		}
		goto st0
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
		switch data[p] {
		case 78:
			goto st451
		case 110:
			goto st451
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
		case 105:
			goto st452
		}
		goto st0
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
		switch data[p] {
		case 90:
			goto st453
		case 122:
			goto st453
		}
		goto st0
	st453:
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		switch data[p] {
		case 65:
			goto st454
		case 97:
			goto st454
		}
		goto st0
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 84:
			goto st455
		case 116:
			goto st455
		}
		goto st0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
	st_case_455:
		switch data[p] {
		case 73:
			goto st456
		case 105:
			goto st456
		}
		goto st0
	st456:
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		switch data[p] {
		case 79:
			goto st457
		case 111:
			goto st457
		}
		goto st0
	st457:
		if p++; p == pe {
			goto _test_eof457
		}
	st_case_457:
		switch data[p] {
		case 78:
			goto st458
		case 110:
			goto st458
		}
		goto st0
	st458:
		if p++; p == pe {
			goto _test_eof458
		}
	st_case_458:
		switch data[p] {
		case 9:
			goto st458
		case 32:
			goto st458
		case 58:
			goto st459
		}
		goto st0
tr661:
//line msg_parse.rl:47

      mark = p
    
	goto st459
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
//line msg_parse.go:15360
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr661
		case 32:
			goto tr661
		case 269:
			goto tr663
		case 525:
			goto tr664
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr662
		}
		goto st0
tr662:
//line msg_parse.rl:47

      mark = p
    
	goto st460
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
//line msg_parse.go:15393
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st460
		case 269:
			goto tr666
		case 525:
			goto tr667
		}
		if 32 <= _widec && _widec <= 253 {
			goto st460
		}
		goto st0
tr672:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:237

      msg.Organization = string(data[mark:p])
    
	goto st461
tr667:
//line msg_parse.rl:237

      msg.Organization = string(data[mark:p])
    
	goto st461
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
//line msg_parse.go:15434
		if data[p] == 10 {
			goto tr668
		}
		goto st0
tr668:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st462
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
	st_case_462:
//line msg_parse.go:15448
		switch data[p] {
		case 9:
			goto st460
		case 13:
			goto st15
		case 32:
			goto st460
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr664:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:237

      msg.Organization = string(data[mark:p])
    
	goto st463
	st463:
		if p++; p == pe {
			goto _test_eof463
		}
	st_case_463:
//line msg_parse.go:15545
		if data[p] == 10 {
			goto tr669
		}
		goto st0
tr669:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st464
	st464:
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
//line msg_parse.go:15559
		switch data[p] {
		case 9:
			goto st465
		case 13:
			goto st15
		case 32:
			goto st465
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr671:
//line msg_parse.rl:47

      mark = p
    
	goto st465
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
//line msg_parse.go:15652
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr671
		case 32:
			goto tr671
		case 269:
			goto tr663
		case 525:
			goto tr672
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr662
		}
		goto st0
tr159:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:154

      *contactp, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
      for *contactp != nil { contactp = &(*contactp).Next }
    
	goto st466
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
//line msg_parse.go:15691
		if data[p] == 10 {
			goto tr673
		}
		goto st0
tr673:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st467
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
//line msg_parse.go:15705
		switch data[p] {
		case 9:
			goto st468
		case 13:
			goto st15
		case 32:
			goto st468
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr675:
//line msg_parse.rl:47

      mark = p
    
	goto st468
	st468:
		if p++; p == pe {
			goto _test_eof468
		}
	st_case_468:
//line msg_parse.go:15798
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr675
		case 32:
			goto tr675
		case 269:
			goto tr158
		case 525:
			goto tr676
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr157
		}
		goto st0
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 88:
			goto st470
		case 120:
			goto st470
		}
		goto st0
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		if data[p] == 45 {
			goto st471
		}
		goto st0
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 70:
			goto st472
		case 102:
			goto st472
		}
		goto st0
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
	st_case_472:
		switch data[p] {
		case 79:
			goto st473
		case 111:
			goto st473
		}
		goto st0
	st473:
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 82:
			goto st474
		case 114:
			goto st474
		}
		goto st0
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 87:
			goto st475
		case 119:
			goto st475
		}
		goto st0
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 65:
			goto st476
		case 97:
			goto st476
		}
		goto st0
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
	st_case_476:
		switch data[p] {
		case 82:
			goto st477
		case 114:
			goto st477
		}
		goto st0
	st477:
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 68:
			goto st478
		case 100:
			goto st478
		}
		goto st0
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 83:
			goto st479
		case 115:
			goto st479
		}
		goto st0
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 9:
			goto st479
		case 32:
			goto st479
		case 58:
			goto st480
		}
		goto st0
	st480:
		if p++; p == pe {
			goto _test_eof480
		}
	st_case_480:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st480
		case 32:
			goto st480
		case 525:
			goto st482
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr688
		}
		goto st0
tr688:
//line msg_parse.rl:217

      msg.MaxForwards = 0
    
//line msg_parse.rl:221

      msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
    
	goto st481
tr690:
//line msg_parse.rl:221

      msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
    
	goto st481
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
//line msg_parse.go:15996
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr690
		}
		goto st0
	st482:
		if p++; p == pe {
			goto _test_eof482
		}
	st_case_482:
		if data[p] == 10 {
			goto tr691
		}
		goto st0
tr691:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st483
	st483:
		if p++; p == pe {
			goto _test_eof483
		}
	st_case_483:
//line msg_parse.go:16022
		switch data[p] {
		case 9:
			goto st484
		case 32:
			goto st484
		}
		goto st0
	st484:
		if p++; p == pe {
			goto _test_eof484
		}
	st_case_484:
		switch data[p] {
		case 9:
			goto st484
		case 32:
			goto st484
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr688
		}
		goto st0
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
		switch data[p] {
		case 77:
			goto st486
		case 78:
			goto st503
		case 109:
			goto st486
		case 110:
			goto st503
		}
		goto st0
	st486:
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		switch data[p] {
		case 69:
			goto st487
		case 101:
			goto st487
		}
		goto st0
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		if data[p] == 45 {
			goto st488
		}
		goto st0
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
	st_case_488:
		switch data[p] {
		case 86:
			goto st489
		case 118:
			goto st489
		}
		goto st0
	st489:
		if p++; p == pe {
			goto _test_eof489
		}
	st_case_489:
		switch data[p] {
		case 69:
			goto st490
		case 101:
			goto st490
		}
		goto st0
	st490:
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 82:
			goto st491
		case 114:
			goto st491
		}
		goto st0
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 83:
			goto st492
		case 115:
			goto st492
		}
		goto st0
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
	st_case_492:
		switch data[p] {
		case 73:
			goto st493
		case 105:
			goto st493
		}
		goto st0
	st493:
		if p++; p == pe {
			goto _test_eof493
		}
	st_case_493:
		switch data[p] {
		case 79:
			goto st494
		case 111:
			goto st494
		}
		goto st0
	st494:
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		switch data[p] {
		case 78:
			goto st495
		case 110:
			goto st495
		}
		goto st0
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 9:
			goto st495
		case 32:
			goto st495
		case 58:
			goto st496
		}
		goto st0
tr705:
//line msg_parse.rl:47

      mark = p
    
	goto st496
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
//line msg_parse.go:16191
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr705
		case 32:
			goto tr705
		case 269:
			goto tr707
		case 525:
			goto tr708
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr706
		}
		goto st0
tr706:
//line msg_parse.rl:47

      mark = p
    
	goto st497
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
//line msg_parse.go:16224
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st497
		case 269:
			goto tr710
		case 525:
			goto tr711
		}
		if 32 <= _widec && _widec <= 253 {
			goto st497
		}
		goto st0
tr716:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:233

      msg.MIMEVersion = string(data[mark:p])
    
	goto st498
tr711:
//line msg_parse.rl:233

      msg.MIMEVersion = string(data[mark:p])
    
	goto st498
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
//line msg_parse.go:16265
		if data[p] == 10 {
			goto tr712
		}
		goto st0
tr712:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st499
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
//line msg_parse.go:16279
		switch data[p] {
		case 9:
			goto st497
		case 13:
			goto st15
		case 32:
			goto st497
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr708:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:233

      msg.MIMEVersion = string(data[mark:p])
    
	goto st500
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
//line msg_parse.go:16376
		if data[p] == 10 {
			goto tr713
		}
		goto st0
tr713:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st501
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
//line msg_parse.go:16390
		switch data[p] {
		case 9:
			goto st502
		case 13:
			goto st15
		case 32:
			goto st502
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr715:
//line msg_parse.rl:47

      mark = p
    
	goto st502
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
//line msg_parse.go:16483
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr715
		case 32:
			goto tr715
		case 269:
			goto tr707
		case 525:
			goto tr716
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr706
		}
		goto st0
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
	st_case_503:
		if data[p] == 45 {
			goto st504
		}
		goto st0
	st504:
		if p++; p == pe {
			goto _test_eof504
		}
	st_case_504:
		switch data[p] {
		case 69:
			goto st505
		case 101:
			goto st505
		}
		goto st0
	st505:
		if p++; p == pe {
			goto _test_eof505
		}
	st_case_505:
		switch data[p] {
		case 88:
			goto st506
		case 120:
			goto st506
		}
		goto st0
	st506:
		if p++; p == pe {
			goto _test_eof506
		}
	st_case_506:
		switch data[p] {
		case 80:
			goto st507
		case 112:
			goto st507
		}
		goto st0
	st507:
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		switch data[p] {
		case 73:
			goto st508
		case 105:
			goto st508
		}
		goto st0
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
		switch data[p] {
		case 82:
			goto st509
		case 114:
			goto st509
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
		case 83:
			goto st511
		case 115:
			goto st511
		}
		goto st0
	st511:
		if p++; p == pe {
			goto _test_eof511
		}
	st_case_511:
		switch data[p] {
		case 9:
			goto st511
		case 32:
			goto st511
		case 58:
			goto st512
		}
		goto st0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
	st_case_512:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st512
		case 32:
			goto st512
		case 525:
			goto st514
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr726
		}
		goto st0
tr726:
//line msg_parse.rl:225

      msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
    
	goto st513
	st513:
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
//line msg_parse.go:16647
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr726
		}
		goto st0
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		if data[p] == 10 {
			goto tr728
		}
		goto st0
tr728:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st515
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
//line msg_parse.go:16673
		switch data[p] {
		case 9:
			goto st516
		case 32:
			goto st516
		}
		goto st0
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		switch data[p] {
		case 9:
			goto st516
		case 32:
			goto st516
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr726
		}
		goto st0
tr141:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:303

      msg.Supported = string(data[mark:p])
    
	goto st517
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
//line msg_parse.go:16711
		if data[p] == 10 {
			goto tr730
		}
		goto st0
tr730:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st518
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
//line msg_parse.go:16725
		switch data[p] {
		case 9:
			goto st519
		case 13:
			goto st15
		case 32:
			goto st519
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr732:
//line msg_parse.rl:47

      mark = p
    
	goto st519
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
//line msg_parse.go:16818
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr732
		case 32:
			goto tr732
		case 269:
			goto tr140
		case 525:
			goto tr733
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr139
		}
		goto st0
tr132:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:213

      msg.InReplyTo = string(data[mark:p])
    
	goto st520
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
//line msg_parse.go:16855
		if data[p] == 10 {
			goto tr734
		}
		goto st0
tr734:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st521
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
//line msg_parse.go:16869
		switch data[p] {
		case 9:
			goto st522
		case 13:
			goto st15
		case 32:
			goto st522
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr736:
//line msg_parse.rl:47

      mark = p
    
	goto st522
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
//line msg_parse.go:16962
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr736
		case 32:
			goto tr736
		case 269:
			goto tr131
		case 525:
			goto tr737
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr130
		}
		goto st0
tr103:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:208

      msg.From, err = ParseAddr(string(data[mark:p]))
      if err != nil { return nil, err }
    
	goto st523
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
	st_case_523:
//line msg_parse.go:17000
		if data[p] == 10 {
			goto tr738
		}
		goto st0
tr738:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st524
	st524:
		if p++; p == pe {
			goto _test_eof524
		}
	st_case_524:
//line msg_parse.go:17014
		switch data[p] {
		case 9:
			goto st525
		case 13:
			goto st15
		case 32:
			goto st525
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr740:
//line msg_parse.rl:47

      mark = p
    
	goto st525
	st525:
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
//line msg_parse.go:17107
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr740
		case 32:
			goto tr740
		case 269:
			goto tr102
		case 525:
			goto tr741
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr101
		}
		goto st0
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
		switch data[p] {
		case 79:
			goto st527
		case 111:
			goto st527
		}
		goto st0
	st527:
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		switch data[p] {
		case 77:
			goto st48
		case 109:
			goto st48
		}
		goto st0
tr92:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:172

      msg.ContentEncoding = string(data[mark:p])
    
	goto st528
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
//line msg_parse.go:17168
		if data[p] == 10 {
			goto tr743
		}
		goto st0
tr743:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st529
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
	st_case_529:
//line msg_parse.go:17182
		switch data[p] {
		case 9:
			goto st530
		case 13:
			goto st15
		case 32:
			goto st530
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr745:
//line msg_parse.rl:47

      mark = p
    
	goto st530
	st530:
		if p++; p == pe {
			goto _test_eof530
		}
	st_case_530:
//line msg_parse.go:17275
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr745
		case 32:
			goto tr745
		case 269:
			goto tr91
		case 525:
			goto tr746
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr90
		}
		goto st0
	st531:
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 82:
			goto st532
		case 114:
			goto st532
		}
		goto st0
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		switch data[p] {
		case 79:
			goto st533
		case 111:
			goto st533
		}
		goto st0
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 82:
			goto st534
		case 114:
			goto st534
		}
		goto st0
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		if data[p] == 45 {
			goto st535
		}
		goto st0
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		switch data[p] {
		case 73:
			goto st536
		case 105:
			goto st536
		}
		goto st0
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		switch data[p] {
		case 78:
			goto st537
		case 110:
			goto st537
		}
		goto st0
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		switch data[p] {
		case 70:
			goto st538
		case 102:
			goto st538
		}
		goto st0
	st538:
		if p++; p == pe {
			goto _test_eof538
		}
	st_case_538:
		switch data[p] {
		case 79:
			goto st539
		case 111:
			goto st539
		}
		goto st0
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		switch data[p] {
		case 9:
			goto st539
		case 32:
			goto st539
		case 58:
			goto st540
		}
		goto st0
tr756:
//line msg_parse.rl:47

      mark = p
    
	goto st540
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
	st_case_540:
//line msg_parse.go:17415
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr756
		case 32:
			goto tr756
		case 269:
			goto tr758
		case 525:
			goto tr759
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr757
		}
		goto st0
tr757:
//line msg_parse.rl:47

      mark = p
    
	goto st541
	st541:
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
//line msg_parse.go:17448
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st541
		case 269:
			goto tr761
		case 525:
			goto tr762
		}
		if 32 <= _widec && _widec <= 253 {
			goto st541
		}
		goto st0
tr767:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:196

      msg.ErrorInfo = string(data[mark:p])
    
	goto st542
tr762:
//line msg_parse.rl:196

      msg.ErrorInfo = string(data[mark:p])
    
	goto st542
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
//line msg_parse.go:17489
		if data[p] == 10 {
			goto tr763
		}
		goto st0
tr763:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st543
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
//line msg_parse.go:17503
		switch data[p] {
		case 9:
			goto st541
		case 13:
			goto st15
		case 32:
			goto st541
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr759:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:196

      msg.ErrorInfo = string(data[mark:p])
    
	goto st544
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
//line msg_parse.go:17600
		if data[p] == 10 {
			goto tr764
		}
		goto st0
tr764:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st545
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
//line msg_parse.go:17614
		switch data[p] {
		case 9:
			goto st546
		case 13:
			goto st15
		case 32:
			goto st546
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr766:
//line msg_parse.rl:47

      mark = p
    
	goto st546
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
//line msg_parse.go:17707
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr766
		case 32:
			goto tr766
		case 269:
			goto tr758
		case 525:
			goto tr767
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr757
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
		case 78:
			goto st549
		case 110:
			goto st549
		}
		goto st0
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
	st_case_549:
		switch data[p] {
		case 84:
			goto st94
		case 116:
			goto st94
		}
		goto st0
	st550:
		if p++; p == pe {
			goto _test_eof550
		}
	st_case_550:
		switch data[p] {
		case 80:
			goto st551
		case 112:
			goto st551
		}
		goto st0
	st551:
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		switch data[p] {
		case 73:
			goto st552
		case 105:
			goto st552
		}
		goto st0
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
		switch data[p] {
		case 82:
			goto st553
		case 114:
			goto st553
		}
		goto st0
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
		goto st0
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		switch data[p] {
		case 83:
			goto st555
		case 115:
			goto st555
		}
		goto st0
	st555:
		if p++; p == pe {
			goto _test_eof555
		}
	st_case_555:
		switch data[p] {
		case 9:
			goto st555
		case 32:
			goto st555
		case 58:
			goto st556
		}
		goto st0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
	st_case_556:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st556
		case 32:
			goto st556
		case 525:
			goto st558
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr776
		}
		goto st0
tr776:
//line msg_parse.rl:204

      msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
    
	goto st557
	st557:
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
//line msg_parse.go:17874
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr776
		}
		goto st0
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		if data[p] == 10 {
			goto tr778
		}
		goto st0
tr778:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st559
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
//line msg_parse.go:17900
		switch data[p] {
		case 9:
			goto st560
		case 32:
			goto st560
		}
		goto st0
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		switch data[p] {
		case 9:
			goto st560
		case 32:
			goto st560
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr776
		}
		goto st0
tr79:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:192

      msg.Date = string(data[mark:p])
    
	goto st561
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
//line msg_parse.go:17938
		if data[p] == 10 {
			goto tr780
		}
		goto st0
tr780:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st562
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
//line msg_parse.go:17952
		switch data[p] {
		case 9:
			goto st563
		case 13:
			goto st15
		case 32:
			goto st563
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr782:
//line msg_parse.rl:47

      mark = p
    
	goto st563
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
//line msg_parse.go:18045
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr782
		case 32:
			goto tr782
		case 269:
			goto tr78
		case 525:
			goto tr783
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr77
		}
		goto st0
tr67:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:176

      msg.ContentType = string(data[mark:p])
    
	goto st564
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
//line msg_parse.go:18082
		if data[p] == 10 {
			goto tr784
		}
		goto st0
tr784:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st565
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
//line msg_parse.go:18096
		switch data[p] {
		case 9:
			goto st566
		case 13:
			goto st15
		case 32:
			goto st566
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr786:
//line msg_parse.rl:47

      mark = p
    
	goto st566
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
//line msg_parse.go:18189
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr786
		case 32:
			goto tr786
		case 269:
			goto tr66
		case 525:
			goto tr787
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr65
		}
		goto st0
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
	st_case_567:
		switch data[p] {
		case 76:
			goto st568
		case 108:
			goto st568
		}
		goto st0
	st568:
		if p++; p == pe {
			goto _test_eof568
		}
	st_case_568:
		switch data[p] {
		case 76:
			goto st569
		case 108:
			goto st569
		}
		goto st0
	st569:
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		if data[p] == 45 {
			goto st570
		}
		goto st0
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 73:
			goto st571
		case 105:
			goto st571
		}
		goto st0
	st571:
		if p++; p == pe {
			goto _test_eof571
		}
	st_case_571:
		switch data[p] {
		case 68:
			goto st54
		case 78:
			goto st572
		case 100:
			goto st54
		case 110:
			goto st572
		}
		goto st0
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		switch data[p] {
		case 70:
			goto st573
		case 102:
			goto st573
		}
		goto st0
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		switch data[p] {
		case 79:
			goto st574
		case 111:
			goto st574
		}
		goto st0
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 9:
			goto st574
		case 32:
			goto st574
		case 58:
			goto st575
		}
		goto st0
tr796:
//line msg_parse.rl:47

      mark = p
    
	goto st575
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
//line msg_parse.go:18321
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr796
		case 32:
			goto tr796
		case 269:
			goto tr798
		case 525:
			goto tr799
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr797
		}
		goto st0
tr797:
//line msg_parse.rl:47

      mark = p
    
	goto st576
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
//line msg_parse.go:18354
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st576
		case 269:
			goto tr801
		case 525:
			goto tr802
		}
		if 32 <= _widec && _widec <= 253 {
			goto st576
		}
		goto st0
tr807:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:188

      msg.CallInfo = string(data[mark:p])
    
	goto st577
tr802:
//line msg_parse.rl:188

      msg.CallInfo = string(data[mark:p])
    
	goto st577
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
//line msg_parse.go:18395
		if data[p] == 10 {
			goto tr803
		}
		goto st0
tr803:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st578
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
//line msg_parse.go:18409
		switch data[p] {
		case 9:
			goto st576
		case 13:
			goto st15
		case 32:
			goto st576
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr799:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:188

      msg.CallInfo = string(data[mark:p])
    
	goto st579
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
//line msg_parse.go:18506
		if data[p] == 10 {
			goto tr804
		}
		goto st0
tr804:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st580
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
//line msg_parse.go:18520
		switch data[p] {
		case 9:
			goto st581
		case 13:
			goto st15
		case 32:
			goto st581
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr806:
//line msg_parse.rl:47

      mark = p
    
	goto st581
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
//line msg_parse.go:18613
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr806
		case 32:
			goto tr806
		case 269:
			goto tr798
		case 525:
			goto tr807
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr797
		}
		goto st0
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 78:
			goto st583
		case 110:
			goto st583
		}
		goto st0
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		switch data[p] {
		case 84:
			goto st584
		case 116:
			goto st584
		}
		goto st0
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 65:
			goto st585
		case 69:
			goto st587
		case 97:
			goto st585
		case 101:
			goto st587
		}
		goto st0
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 67:
			goto st586
		case 99:
			goto st586
		}
		goto st0
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
	st_case_586:
		switch data[p] {
		case 84:
			goto st88
		case 116:
			goto st88
		}
		goto st0
	st587:
		if p++; p == pe {
			goto _test_eof587
		}
	st_case_587:
		switch data[p] {
		case 78:
			goto st588
		case 110:
			goto st588
		}
		goto st0
	st588:
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		switch data[p] {
		case 84:
			goto st589
		case 116:
			goto st589
		}
		goto st0
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		if data[p] == 45 {
			goto st590
		}
		goto st0
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		switch data[p] {
		case 68:
			goto st591
		case 69:
			goto st609
		case 76:
			goto st616
		case 84:
			goto st641
		case 100:
			goto st591
		case 101:
			goto st609
		case 108:
			goto st616
		case 116:
			goto st641
		}
		goto st0
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		switch data[p] {
		case 73:
			goto st592
		case 105:
			goto st592
		}
		goto st0
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		switch data[p] {
		case 83:
			goto st593
		case 115:
			goto st593
		}
		goto st0
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
	st_case_593:
		switch data[p] {
		case 80:
			goto st594
		case 112:
			goto st594
		}
		goto st0
	st594:
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		switch data[p] {
		case 79:
			goto st595
		case 111:
			goto st595
		}
		goto st0
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 83:
			goto st596
		case 115:
			goto st596
		}
		goto st0
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		switch data[p] {
		case 73:
			goto st597
		case 105:
			goto st597
		}
		goto st0
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
	st_case_597:
		switch data[p] {
		case 84:
			goto st598
		case 116:
			goto st598
		}
		goto st0
	st598:
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		switch data[p] {
		case 73:
			goto st599
		case 105:
			goto st599
		}
		goto st0
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 79:
			goto st600
		case 111:
			goto st600
		}
		goto st0
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		switch data[p] {
		case 78:
			goto st601
		case 110:
			goto st601
		}
		goto st0
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
	st_case_601:
		switch data[p] {
		case 9:
			goto st601
		case 32:
			goto st601
		case 58:
			goto st602
		}
		goto st0
tr831:
//line msg_parse.rl:47

      mark = p
    
	goto st602
	st602:
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
//line msg_parse.go:18901
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr831
		case 32:
			goto tr831
		case 269:
			goto tr833
		case 525:
			goto tr834
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr832
		}
		goto st0
tr832:
//line msg_parse.rl:47

      mark = p
    
	goto st603
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
//line msg_parse.go:18934
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st603
		case 269:
			goto tr836
		case 525:
			goto tr837
		}
		if 32 <= _widec && _widec <= 253 {
			goto st603
		}
		goto st0
tr842:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:160

      msg.ContentDisposition = string(data[mark:p])
    
	goto st604
tr837:
//line msg_parse.rl:160

      msg.ContentDisposition = string(data[mark:p])
    
	goto st604
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
//line msg_parse.go:18975
		if data[p] == 10 {
			goto tr838
		}
		goto st0
tr838:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st605
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
//line msg_parse.go:18989
		switch data[p] {
		case 9:
			goto st603
		case 13:
			goto st15
		case 32:
			goto st603
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr834:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:160

      msg.ContentDisposition = string(data[mark:p])
    
	goto st606
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
//line msg_parse.go:19086
		if data[p] == 10 {
			goto tr839
		}
		goto st0
tr839:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st607
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
//line msg_parse.go:19100
		switch data[p] {
		case 9:
			goto st608
		case 13:
			goto st15
		case 32:
			goto st608
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr841:
//line msg_parse.rl:47

      mark = p
    
	goto st608
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
//line msg_parse.go:19193
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr841
		case 32:
			goto tr841
		case 269:
			goto tr833
		case 525:
			goto tr842
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr832
		}
		goto st0
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
	st_case_609:
		switch data[p] {
		case 78:
			goto st610
		case 110:
			goto st610
		}
		goto st0
	st610:
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		switch data[p] {
		case 67:
			goto st611
		case 99:
			goto st611
		}
		goto st0
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 79:
			goto st612
		case 111:
			goto st612
		}
		goto st0
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		switch data[p] {
		case 68:
			goto st613
		case 100:
			goto st613
		}
		goto st0
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		switch data[p] {
		case 73:
			goto st614
		case 105:
			goto st614
		}
		goto st0
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
		goto st0
	st615:
		if p++; p == pe {
			goto _test_eof615
		}
	st_case_615:
		switch data[p] {
		case 71:
			goto st42
		case 103:
			goto st42
		}
		goto st0
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		switch data[p] {
		case 65:
			goto st617
		case 69:
			goto st631
		case 97:
			goto st617
		case 101:
			goto st631
		}
		goto st0
	st617:
		if p++; p == pe {
			goto _test_eof617
		}
	st_case_617:
		switch data[p] {
		case 78:
			goto st618
		case 110:
			goto st618
		}
		goto st0
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
	st_case_618:
		switch data[p] {
		case 71:
			goto st619
		case 103:
			goto st619
		}
		goto st0
	st619:
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 85:
			goto st620
		case 117:
			goto st620
		}
		goto st0
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		switch data[p] {
		case 65:
			goto st621
		case 97:
			goto st621
		}
		goto st0
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		switch data[p] {
		case 71:
			goto st622
		case 103:
			goto st622
		}
		goto st0
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 69:
			goto st623
		case 101:
			goto st623
		}
		goto st0
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		switch data[p] {
		case 9:
			goto st623
		case 32:
			goto st623
		case 58:
			goto st624
		}
		goto st0
tr858:
//line msg_parse.rl:47

      mark = p
    
	goto st624
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
//line msg_parse.go:19412
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr858
		case 32:
			goto tr858
		case 269:
			goto tr860
		case 525:
			goto tr861
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr859
		}
		goto st0
tr859:
//line msg_parse.rl:47

      mark = p
    
	goto st625
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
	st_case_625:
//line msg_parse.go:19445
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st625
		case 269:
			goto tr863
		case 525:
			goto tr864
		}
		if 32 <= _widec && _widec <= 253 {
			goto st625
		}
		goto st0
tr869:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:164

      msg.ContentLanguage = string(data[mark:p])
    
	goto st626
tr864:
//line msg_parse.rl:164

      msg.ContentLanguage = string(data[mark:p])
    
	goto st626
	st626:
		if p++; p == pe {
			goto _test_eof626
		}
	st_case_626:
//line msg_parse.go:19486
		if data[p] == 10 {
			goto tr865
		}
		goto st0
tr865:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st627
	st627:
		if p++; p == pe {
			goto _test_eof627
		}
	st_case_627:
//line msg_parse.go:19500
		switch data[p] {
		case 9:
			goto st625
		case 13:
			goto st15
		case 32:
			goto st625
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr861:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:164

      msg.ContentLanguage = string(data[mark:p])
    
	goto st628
	st628:
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
//line msg_parse.go:19597
		if data[p] == 10 {
			goto tr866
		}
		goto st0
tr866:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st629
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
//line msg_parse.go:19611
		switch data[p] {
		case 9:
			goto st630
		case 13:
			goto st15
		case 32:
			goto st630
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr868:
//line msg_parse.rl:47

      mark = p
    
	goto st630
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
//line msg_parse.go:19704
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr868
		case 32:
			goto tr868
		case 269:
			goto tr860
		case 525:
			goto tr869
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr859
		}
		goto st0
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		switch data[p] {
		case 78:
			goto st632
		case 110:
			goto st632
		}
		goto st0
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		switch data[p] {
		case 71:
			goto st633
		case 103:
			goto st633
		}
		goto st0
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		switch data[p] {
		case 84:
			goto st634
		case 116:
			goto st634
		}
		goto st0
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		switch data[p] {
		case 72:
			goto st635
		case 104:
			goto st635
		}
		goto st0
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		switch data[p] {
		case 9:
			goto st635
		case 32:
			goto st635
		case 58:
			goto st636
		}
		goto st0
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st636
		case 32:
			goto st636
		case 525:
			goto st638
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr875
		}
		goto st0
tr875:
//line msg_parse.rl:168

      clen = clen * 10 + (int(data[p]) - 0x30)
    
	goto st637
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
	st_case_637:
//line msg_parse.go:19823
		if data[p] == 13 {
			goto st13
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr875
		}
		goto st0
	st638:
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		if data[p] == 10 {
			goto tr877
		}
		goto st0
tr877:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st639
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
//line msg_parse.go:19849
		switch data[p] {
		case 9:
			goto st640
		case 32:
			goto st640
		}
		goto st0
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
	st_case_640:
		switch data[p] {
		case 9:
			goto st640
		case 32:
			goto st640
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr875
		}
		goto st0
	st641:
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		switch data[p] {
		case 89:
			goto st642
		case 121:
			goto st642
		}
		goto st0
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		switch data[p] {
		case 80:
			goto st643
		case 112:
			goto st643
		}
		goto st0
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
	st_case_643:
		switch data[p] {
		case 69:
			goto st28
		case 101:
			goto st28
		}
		goto st0
	st644:
		if p++; p == pe {
			goto _test_eof644
		}
	st_case_644:
		switch data[p] {
		case 69:
			goto st645
		case 101:
			goto st645
		}
		goto st0
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		switch data[p] {
		case 81:
			goto st646
		case 113:
			goto st646
		}
		goto st0
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		switch data[p] {
		case 9:
			goto st646
		case 32:
			goto st646
		case 58:
			goto st647
		}
		goto st0
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st647
		case 32:
			goto st647
		case 525:
			goto st654
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr884
		}
		goto st0
tr884:
//line msg_parse.rl:180

      msg.CSeq = msg.CSeq * 10 + (int(data[p]) - 0x30)
    
	goto st648
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
//line msg_parse.go:19981
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st649
		case 32:
			goto st649
		case 525:
			goto st651
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr884
		}
		goto st0
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st649
		case 32:
			goto st649
		case 33:
			goto tr888
		case 37:
			goto tr888
		case 39:
			goto tr888
		case 126:
			goto tr888
		case 525:
			goto st651
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr888
				}
			case _widec >= 42:
				goto tr888
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr888
				}
			case _widec >= 65:
				goto tr888
			}
		default:
			goto tr888
		}
		goto st0
tr888:
//line msg_parse.rl:47

      mark = p
    
	goto st650
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
//line msg_parse.go:20063
		switch data[p] {
		case 13:
			goto tr889
		case 33:
			goto st650
		case 37:
			goto st650
		case 39:
			goto st650
		case 126:
			goto st650
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st650
				}
			case data[p] >= 42:
				goto st650
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st650
				}
			case data[p] >= 65:
				goto st650
			}
		default:
			goto st650
		}
		goto st0
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		if data[p] == 10 {
			goto tr891
		}
		goto st0
tr891:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st652
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
//line msg_parse.go:20117
		switch data[p] {
		case 9:
			goto st653
		case 32:
			goto st653
		}
		goto st0
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 9:
			goto st653
		case 32:
			goto st653
		case 33:
			goto tr888
		case 37:
			goto tr888
		case 39:
			goto tr888
		case 126:
			goto tr888
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr888
				}
			case data[p] >= 42:
				goto tr888
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr888
				}
			case data[p] >= 65:
				goto tr888
			}
		default:
			goto tr888
		}
		goto st0
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		if data[p] == 10 {
			goto tr893
		}
		goto st0
tr893:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st655
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
//line msg_parse.go:20185
		switch data[p] {
		case 9:
			goto st656
		case 32:
			goto st656
		}
		goto st0
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		switch data[p] {
		case 9:
			goto st656
		case 32:
			goto st656
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr884
		}
		goto st0
tr54:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:272

      msg.ReferredBy = string(data[mark:p])
    
	goto st657
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
//line msg_parse.go:20223
		if data[p] == 10 {
			goto tr895
		}
		goto st0
tr895:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st658
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
//line msg_parse.go:20237
		switch data[p] {
		case 9:
			goto st659
		case 13:
			goto st15
		case 32:
			goto st659
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr897:
//line msg_parse.rl:47

      mark = p
    
	goto st659
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
//line msg_parse.go:20330
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr897
		case 32:
			goto tr897
		case 269:
			goto tr53
		case 525:
			goto tr898
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr52
		}
		goto st0
tr45:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:118

      msg.AcceptContact = string(data[mark:p])
    
	goto st660
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
//line msg_parse.go:20367
		if data[p] == 10 {
			goto tr899
		}
		goto st0
tr899:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st661
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
	st_case_661:
//line msg_parse.go:20381
		switch data[p] {
		case 9:
			goto st662
		case 13:
			goto st15
		case 32:
			goto st662
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr901:
//line msg_parse.rl:47

      mark = p
    
	goto st662
	st662:
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
//line msg_parse.go:20474
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr901
		case 32:
			goto tr901
		case 269:
			goto tr44
		case 525:
			goto tr902
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr43
		}
		goto st0
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		switch data[p] {
		case 67:
			goto st664
		case 99:
			goto st664
		}
		goto st0
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		switch data[p] {
		case 69:
			goto st665
		case 101:
			goto st665
		}
		goto st0
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		switch data[p] {
		case 80:
			goto st666
		case 112:
			goto st666
		}
		goto st0
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		switch data[p] {
		case 84:
			goto st667
		case 116:
			goto st667
		}
		goto st0
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
	st_case_667:
		switch data[p] {
		case 9:
			goto st668
		case 32:
			goto st668
		case 45:
			goto st676
		case 58:
			goto st669
		}
		goto st0
	st668:
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		switch data[p] {
		case 9:
			goto st668
		case 32:
			goto st668
		case 58:
			goto st669
		}
		goto st0
tr910:
//line msg_parse.rl:47

      mark = p
    
	goto st669
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
//line msg_parse.go:20585
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr910
		case 32:
			goto tr910
		case 269:
			goto tr912
		case 525:
			goto tr913
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr911
		}
		goto st0
tr911:
//line msg_parse.rl:47

      mark = p
    
	goto st670
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
//line msg_parse.go:20618
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st670
		case 269:
			goto tr915
		case 525:
			goto tr916
		}
		if 32 <= _widec && _widec <= 253 {
			goto st670
		}
		goto st0
tr921:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:114

      msg.Accept = string(data[mark:p])
    
	goto st671
tr916:
//line msg_parse.rl:114

      msg.Accept = string(data[mark:p])
    
	goto st671
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
//line msg_parse.go:20659
		if data[p] == 10 {
			goto tr917
		}
		goto st0
tr917:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st672
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
	st_case_672:
//line msg_parse.go:20673
		switch data[p] {
		case 9:
			goto st670
		case 13:
			goto st15
		case 32:
			goto st670
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr913:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:114

      msg.Accept = string(data[mark:p])
    
	goto st673
	st673:
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
//line msg_parse.go:20770
		if data[p] == 10 {
			goto tr918
		}
		goto st0
tr918:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st674
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
//line msg_parse.go:20784
		switch data[p] {
		case 9:
			goto st675
		case 13:
			goto st15
		case 32:
			goto st675
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr920:
//line msg_parse.rl:47

      mark = p
    
	goto st675
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
//line msg_parse.go:20877
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr920
		case 32:
			goto tr920
		case 269:
			goto tr912
		case 525:
			goto tr921
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr911
		}
		goto st0
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		switch data[p] {
		case 67:
			goto st677
		case 69:
			goto st683
		case 76:
			goto st698
		case 99:
			goto st677
		case 101:
			goto st683
		case 108:
			goto st698
		}
		goto st0
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 79:
			goto st678
		case 111:
			goto st678
		}
		goto st0
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		switch data[p] {
		case 78:
			goto st679
		case 110:
			goto st679
		}
		goto st0
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		switch data[p] {
		case 84:
			goto st680
		case 116:
			goto st680
		}
		goto st0
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		switch data[p] {
		case 65:
			goto st681
		case 97:
			goto st681
		}
		goto st0
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
	st_case_681:
		switch data[p] {
		case 67:
			goto st682
		case 99:
			goto st682
		}
		goto st0
	st682:
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		switch data[p] {
		case 84:
			goto st17
		case 116:
			goto st17
		}
		goto st0
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		switch data[p] {
		case 78:
			goto st684
		case 110:
			goto st684
		}
		goto st0
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		switch data[p] {
		case 67:
			goto st685
		case 99:
			goto st685
		}
		goto st0
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
	st_case_685:
		switch data[p] {
		case 79:
			goto st686
		case 111:
			goto st686
		}
		goto st0
	st686:
		if p++; p == pe {
			goto _test_eof686
		}
	st_case_686:
		switch data[p] {
		case 68:
			goto st687
		case 100:
			goto st687
		}
		goto st0
	st687:
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		switch data[p] {
		case 73:
			goto st688
		case 105:
			goto st688
		}
		goto st0
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 78:
			goto st689
		case 110:
			goto st689
		}
		goto st0
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		switch data[p] {
		case 71:
			goto st690
		case 103:
			goto st690
		}
		goto st0
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		switch data[p] {
		case 9:
			goto st690
		case 32:
			goto st690
		case 58:
			goto st691
		}
		goto st0
tr938:
//line msg_parse.rl:47

      mark = p
    
	goto st691
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
	st_case_691:
//line msg_parse.go:21100
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr938
		case 32:
			goto tr938
		case 269:
			goto tr940
		case 525:
			goto tr941
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr939
		}
		goto st0
tr939:
//line msg_parse.rl:47

      mark = p
    
	goto st692
	st692:
		if p++; p == pe {
			goto _test_eof692
		}
	st_case_692:
//line msg_parse.go:21133
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st692
		case 269:
			goto tr943
		case 525:
			goto tr944
		}
		if 32 <= _widec && _widec <= 253 {
			goto st692
		}
		goto st0
tr949:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:122

      msg.AcceptEncoding = string(data[mark:p])
    
	goto st693
tr944:
//line msg_parse.rl:122

      msg.AcceptEncoding = string(data[mark:p])
    
	goto st693
	st693:
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
//line msg_parse.go:21174
		if data[p] == 10 {
			goto tr945
		}
		goto st0
tr945:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st694
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
//line msg_parse.go:21188
		switch data[p] {
		case 9:
			goto st692
		case 13:
			goto st15
		case 32:
			goto st692
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr941:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:122

      msg.AcceptEncoding = string(data[mark:p])
    
	goto st695
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
//line msg_parse.go:21285
		if data[p] == 10 {
			goto tr946
		}
		goto st0
tr946:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st696
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
//line msg_parse.go:21299
		switch data[p] {
		case 9:
			goto st697
		case 13:
			goto st15
		case 32:
			goto st697
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr948:
//line msg_parse.rl:47

      mark = p
    
	goto st697
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
	st_case_697:
//line msg_parse.go:21392
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr948
		case 32:
			goto tr948
		case 269:
			goto tr940
		case 525:
			goto tr949
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr939
		}
		goto st0
	st698:
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		switch data[p] {
		case 65:
			goto st699
		case 97:
			goto st699
		}
		goto st0
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		switch data[p] {
		case 78:
			goto st700
		case 110:
			goto st700
		}
		goto st0
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		switch data[p] {
		case 71:
			goto st701
		case 103:
			goto st701
		}
		goto st0
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 85:
			goto st702
		case 117:
			goto st702
		}
		goto st0
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		switch data[p] {
		case 65:
			goto st703
		case 97:
			goto st703
		}
		goto st0
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		switch data[p] {
		case 71:
			goto st704
		case 103:
			goto st704
		}
		goto st0
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
	st_case_704:
		switch data[p] {
		case 69:
			goto st705
		case 101:
			goto st705
		}
		goto st0
	st705:
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		switch data[p] {
		case 9:
			goto st705
		case 32:
			goto st705
		case 58:
			goto st706
		}
		goto st0
tr958:
//line msg_parse.rl:47

      mark = p
    
	goto st706
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
//line msg_parse.go:21523
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr958
		case 32:
			goto tr958
		case 269:
			goto tr960
		case 525:
			goto tr961
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr959
		}
		goto st0
tr959:
//line msg_parse.rl:47

      mark = p
    
	goto st707
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
//line msg_parse.go:21556
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st707
		case 269:
			goto tr963
		case 525:
			goto tr964
		}
		if 32 <= _widec && _widec <= 253 {
			goto st707
		}
		goto st0
tr969:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:126

      msg.AcceptLanguage = string(data[mark:p])
    
	goto st708
tr964:
//line msg_parse.rl:126

      msg.AcceptLanguage = string(data[mark:p])
    
	goto st708
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
//line msg_parse.go:21597
		if data[p] == 10 {
			goto tr965
		}
		goto st0
tr965:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st709
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
//line msg_parse.go:21611
		switch data[p] {
		case 9:
			goto st707
		case 13:
			goto st15
		case 32:
			goto st707
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr961:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:126

      msg.AcceptLanguage = string(data[mark:p])
    
	goto st710
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
//line msg_parse.go:21708
		if data[p] == 10 {
			goto tr966
		}
		goto st0
tr966:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st711
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
//line msg_parse.go:21722
		switch data[p] {
		case 9:
			goto st712
		case 13:
			goto st15
		case 32:
			goto st712
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr968:
//line msg_parse.rl:47

      mark = p
    
	goto st712
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
	st_case_712:
//line msg_parse.go:21815
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr968
		case 32:
			goto tr968
		case 269:
			goto tr960
		case 525:
			goto tr969
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr959
		}
		goto st0
	st713:
		if p++; p == pe {
			goto _test_eof713
		}
	st_case_713:
		switch data[p] {
		case 69:
			goto st714
		case 76:
			goto st729
		case 101:
			goto st714
		case 108:
			goto st729
		}
		goto st0
	st714:
		if p++; p == pe {
			goto _test_eof714
		}
	st_case_714:
		switch data[p] {
		case 82:
			goto st715
		case 114:
			goto st715
		}
		goto st0
	st715:
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		switch data[p] {
		case 84:
			goto st716
		case 116:
			goto st716
		}
		goto st0
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		if data[p] == 45 {
			goto st717
		}
		goto st0
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		switch data[p] {
		case 73:
			goto st718
		case 105:
			goto st718
		}
		goto st0
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		switch data[p] {
		case 78:
			goto st719
		case 110:
			goto st719
		}
		goto st0
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		switch data[p] {
		case 70:
			goto st720
		case 102:
			goto st720
		}
		goto st0
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		switch data[p] {
		case 79:
			goto st721
		case 111:
			goto st721
		}
		goto st0
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		switch data[p] {
		case 9:
			goto st721
		case 32:
			goto st721
		case 58:
			goto st722
		}
		goto st0
tr980:
//line msg_parse.rl:47

      mark = p
    
	goto st722
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
//line msg_parse.go:21959
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr980
		case 32:
			goto tr980
		case 269:
			goto tr982
		case 525:
			goto tr983
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr981
		}
		goto st0
tr981:
//line msg_parse.rl:47

      mark = p
    
	goto st723
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
//line msg_parse.go:21992
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st723
		case 269:
			goto tr985
		case 525:
			goto tr986
		}
		if 32 <= _widec && _widec <= 253 {
			goto st723
		}
		goto st0
tr991:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:138

      msg.AlertInfo = string(data[mark:p])
    
	goto st724
tr986:
//line msg_parse.rl:138

      msg.AlertInfo = string(data[mark:p])
    
	goto st724
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
	st_case_724:
//line msg_parse.go:22033
		if data[p] == 10 {
			goto tr987
		}
		goto st0
tr987:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st725
	st725:
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
//line msg_parse.go:22047
		switch data[p] {
		case 9:
			goto st723
		case 13:
			goto st15
		case 32:
			goto st723
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr983:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:138

      msg.AlertInfo = string(data[mark:p])
    
	goto st726
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
//line msg_parse.go:22144
		if data[p] == 10 {
			goto tr988
		}
		goto st0
tr988:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st727
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
//line msg_parse.go:22158
		switch data[p] {
		case 9:
			goto st728
		case 13:
			goto st15
		case 32:
			goto st728
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr990:
//line msg_parse.rl:47

      mark = p
    
	goto st728
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
//line msg_parse.go:22251
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr990
		case 32:
			goto tr990
		case 269:
			goto tr982
		case 525:
			goto tr991
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr981
		}
		goto st0
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		switch data[p] {
		case 79:
			goto st730
		case 111:
			goto st730
		}
		goto st0
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		switch data[p] {
		case 87:
			goto st731
		case 119:
			goto st731
		}
		goto st0
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		switch data[p] {
		case 9:
			goto st732
		case 32:
			goto st732
		case 45:
			goto st740
		case 58:
			goto st733
		}
		goto st0
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		switch data[p] {
		case 9:
			goto st732
		case 32:
			goto st732
		case 58:
			goto st733
		}
		goto st0
tr997:
//line msg_parse.rl:47

      mark = p
    
	goto st733
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
	st_case_733:
//line msg_parse.go:22338
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr997
		case 32:
			goto tr997
		case 269:
			goto tr999
		case 525:
			goto tr1000
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr998
		}
		goto st0
tr998:
//line msg_parse.rl:47

      mark = p
    
	goto st734
	st734:
		if p++; p == pe {
			goto _test_eof734
		}
	st_case_734:
//line msg_parse.go:22371
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st734
		case 269:
			goto tr1002
		case 525:
			goto tr1003
		}
		if 32 <= _widec && _widec <= 253 {
			goto st734
		}
		goto st0
tr1008:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
	goto st735
tr1003:
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
	goto st735
	st735:
		if p++; p == pe {
			goto _test_eof735
		}
	st_case_735:
//line msg_parse.go:22412
		if data[p] == 10 {
			goto tr1004
		}
		goto st0
tr1004:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st736
	st736:
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
//line msg_parse.go:22426
		switch data[p] {
		case 9:
			goto st734
		case 13:
			goto st15
		case 32:
			goto st734
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1000:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:130

      msg.Allow = string(data[mark:p])
    
	goto st737
	st737:
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
//line msg_parse.go:22523
		if data[p] == 10 {
			goto tr1005
		}
		goto st0
tr1005:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st738
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
//line msg_parse.go:22537
		switch data[p] {
		case 9:
			goto st739
		case 13:
			goto st15
		case 32:
			goto st739
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1007:
//line msg_parse.rl:47

      mark = p
    
	goto st739
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
//line msg_parse.go:22630
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1007
		case 32:
			goto tr1007
		case 269:
			goto tr999
		case 525:
			goto tr1008
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr998
		}
		goto st0
	st740:
		if p++; p == pe {
			goto _test_eof740
		}
	st_case_740:
		switch data[p] {
		case 69:
			goto st741
		case 101:
			goto st741
		}
		goto st0
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
	st_case_741:
		switch data[p] {
		case 86:
			goto st742
		case 118:
			goto st742
		}
		goto st0
	st742:
		if p++; p == pe {
			goto _test_eof742
		}
	st_case_742:
		switch data[p] {
		case 69:
			goto st743
		case 101:
			goto st743
		}
		goto st0
	st743:
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
		switch data[p] {
		case 78:
			goto st744
		case 110:
			goto st744
		}
		goto st0
	st744:
		if p++; p == pe {
			goto _test_eof744
		}
	st_case_744:
		switch data[p] {
		case 84:
			goto st745
		case 116:
			goto st745
		}
		goto st0
	st745:
		if p++; p == pe {
			goto _test_eof745
		}
	st_case_745:
		switch data[p] {
		case 83:
			goto st746
		case 115:
			goto st746
		}
		goto st0
	st746:
		if p++; p == pe {
			goto _test_eof746
		}
	st_case_746:
		switch data[p] {
		case 9:
			goto st746
		case 32:
			goto st746
		case 58:
			goto st747
		}
		goto st0
tr1016:
//line msg_parse.rl:47

      mark = p
    
	goto st747
	st747:
		if p++; p == pe {
			goto _test_eof747
		}
	st_case_747:
//line msg_parse.go:22749
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1016
		case 32:
			goto tr1016
		case 269:
			goto tr1018
		case 525:
			goto tr1019
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr1017
		}
		goto st0
tr1017:
//line msg_parse.rl:47

      mark = p
    
	goto st748
	st748:
		if p++; p == pe {
			goto _test_eof748
		}
	st_case_748:
//line msg_parse.go:22782
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st748
		case 269:
			goto tr1021
		case 525:
			goto tr1022
		}
		if 32 <= _widec && _widec <= 253 {
			goto st748
		}
		goto st0
tr1027:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st749
tr1022:
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st749
	st749:
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
//line msg_parse.go:22823
		if data[p] == 10 {
			goto tr1023
		}
		goto st0
tr1023:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st750
	st750:
		if p++; p == pe {
			goto _test_eof750
		}
	st_case_750:
//line msg_parse.go:22837
		switch data[p] {
		case 9:
			goto st748
		case 13:
			goto st15
		case 32:
			goto st748
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1019:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:134

      msg.AllowEvents = string(data[mark:p])
    
	goto st751
	st751:
		if p++; p == pe {
			goto _test_eof751
		}
	st_case_751:
//line msg_parse.go:22934
		if data[p] == 10 {
			goto tr1024
		}
		goto st0
tr1024:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st752
	st752:
		if p++; p == pe {
			goto _test_eof752
		}
	st_case_752:
//line msg_parse.go:22948
		switch data[p] {
		case 9:
			goto st753
		case 13:
			goto st15
		case 32:
			goto st753
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1026:
//line msg_parse.rl:47

      mark = p
    
	goto st753
	st753:
		if p++; p == pe {
			goto _test_eof753
		}
	st_case_753:
//line msg_parse.go:23041
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1026
		case 32:
			goto tr1026
		case 269:
			goto tr1018
		case 525:
			goto tr1027
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr1017
		}
		goto st0
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		switch data[p] {
		case 84:
			goto st755
		case 116:
			goto st755
		}
		goto st0
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		switch data[p] {
		case 72:
			goto st756
		case 104:
			goto st756
		}
		goto st0
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		switch data[p] {
		case 69:
			goto st757
		case 79:
			goto st779
		case 101:
			goto st757
		case 111:
			goto st779
		}
		goto st0
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		switch data[p] {
		case 78:
			goto st758
		case 110:
			goto st758
		}
		goto st0
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		switch data[p] {
		case 84:
			goto st759
		case 116:
			goto st759
		}
		goto st0
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		switch data[p] {
		case 73:
			goto st760
		case 105:
			goto st760
		}
		goto st0
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		switch data[p] {
		case 67:
			goto st761
		case 99:
			goto st761
		}
		goto st0
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		switch data[p] {
		case 65:
			goto st762
		case 97:
			goto st762
		}
		goto st0
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		switch data[p] {
		case 84:
			goto st763
		case 116:
			goto st763
		}
		goto st0
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
	st_case_763:
		switch data[p] {
		case 73:
			goto st764
		case 105:
			goto st764
		}
		goto st0
	st764:
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		switch data[p] {
		case 79:
			goto st765
		case 111:
			goto st765
		}
		goto st0
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
		switch data[p] {
		case 78:
			goto st766
		case 110:
			goto st766
		}
		goto st0
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
		if data[p] == 45 {
			goto st767
		}
		goto st0
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
		switch data[p] {
		case 73:
			goto st768
		case 105:
			goto st768
		}
		goto st0
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
		switch data[p] {
		case 78:
			goto st769
		case 110:
			goto st769
		}
		goto st0
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
		switch data[p] {
		case 70:
			goto st770
		case 102:
			goto st770
		}
		goto st0
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
		switch data[p] {
		case 79:
			goto st771
		case 111:
			goto st771
		}
		goto st0
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
		switch data[p] {
		case 9:
			goto st771
		case 32:
			goto st771
		case 58:
			goto st772
		}
		goto st0
tr1047:
//line msg_parse.rl:47

      mark = p
    
	goto st772
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
//line msg_parse.go:23293
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1047
		case 32:
			goto tr1047
		case 269:
			goto tr1049
		case 525:
			goto tr1050
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr1048
		}
		goto st0
tr1048:
//line msg_parse.rl:47

      mark = p
    
	goto st773
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
//line msg_parse.go:23326
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st773
		case 269:
			goto tr1052
		case 525:
			goto tr1053
		}
		if 32 <= _widec && _widec <= 253 {
			goto st773
		}
		goto st0
tr1058:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:142

      msg.AuthenticationInfo = string(data[mark:p])
    
	goto st774
tr1053:
//line msg_parse.rl:142

      msg.AuthenticationInfo = string(data[mark:p])
    
	goto st774
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
//line msg_parse.go:23367
		if data[p] == 10 {
			goto tr1054
		}
		goto st0
tr1054:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st775
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
//line msg_parse.go:23381
		switch data[p] {
		case 9:
			goto st773
		case 13:
			goto st15
		case 32:
			goto st773
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1050:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:142

      msg.AuthenticationInfo = string(data[mark:p])
    
	goto st776
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
//line msg_parse.go:23478
		if data[p] == 10 {
			goto tr1055
		}
		goto st0
tr1055:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st777
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
//line msg_parse.go:23492
		switch data[p] {
		case 9:
			goto st778
		case 13:
			goto st15
		case 32:
			goto st778
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1057:
//line msg_parse.rl:47

      mark = p
    
	goto st778
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
//line msg_parse.go:23585
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1057
		case 32:
			goto tr1057
		case 269:
			goto tr1049
		case 525:
			goto tr1058
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr1048
		}
		goto st0
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
		switch data[p] {
		case 82:
			goto st780
		case 114:
			goto st780
		}
		goto st0
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
		switch data[p] {
		case 73:
			goto st781
		case 105:
			goto st781
		}
		goto st0
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
		switch data[p] {
		case 90:
			goto st782
		case 122:
			goto st782
		}
		goto st0
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
		switch data[p] {
		case 65:
			goto st783
		case 97:
			goto st783
		}
		goto st0
	st783:
		if p++; p == pe {
			goto _test_eof783
		}
	st_case_783:
		switch data[p] {
		case 84:
			goto st784
		case 116:
			goto st784
		}
		goto st0
	st784:
		if p++; p == pe {
			goto _test_eof784
		}
	st_case_784:
		switch data[p] {
		case 73:
			goto st785
		case 105:
			goto st785
		}
		goto st0
	st785:
		if p++; p == pe {
			goto _test_eof785
		}
	st_case_785:
		switch data[p] {
		case 79:
			goto st786
		case 111:
			goto st786
		}
		goto st0
	st786:
		if p++; p == pe {
			goto _test_eof786
		}
	st_case_786:
		switch data[p] {
		case 78:
			goto st787
		case 110:
			goto st787
		}
		goto st0
	st787:
		if p++; p == pe {
			goto _test_eof787
		}
	st_case_787:
		switch data[p] {
		case 9:
			goto st787
		case 32:
			goto st787
		case 58:
			goto st788
		}
		goto st0
tr1068:
//line msg_parse.rl:47

      mark = p
    
	goto st788
	st788:
		if p++; p == pe {
			goto _test_eof788
		}
	st_case_788:
//line msg_parse.go:23728
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1068
		case 32:
			goto tr1068
		case 269:
			goto tr1070
		case 525:
			goto tr1071
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr1069
		}
		goto st0
tr1069:
//line msg_parse.rl:47

      mark = p
    
	goto st789
	st789:
		if p++; p == pe {
			goto _test_eof789
		}
	st_case_789:
//line msg_parse.go:23761
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st789
		case 269:
			goto tr1073
		case 525:
			goto tr1074
		}
		if 32 <= _widec && _widec <= 253 {
			goto st789
		}
		goto st0
tr1079:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:146

      msg.Authorization = string(data[mark:p])
    
	goto st790
tr1074:
//line msg_parse.rl:146

      msg.Authorization = string(data[mark:p])
    
	goto st790
	st790:
		if p++; p == pe {
			goto _test_eof790
		}
	st_case_790:
//line msg_parse.go:23802
		if data[p] == 10 {
			goto tr1075
		}
		goto st0
tr1075:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st791
	st791:
		if p++; p == pe {
			goto _test_eof791
		}
	st_case_791:
//line msg_parse.go:23816
		switch data[p] {
		case 9:
			goto st789
		case 13:
			goto st15
		case 32:
			goto st789
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1071:
//line msg_parse.rl:47

      mark = p
    
//line msg_parse.rl:146

      msg.Authorization = string(data[mark:p])
    
	goto st792
	st792:
		if p++; p == pe {
			goto _test_eof792
		}
	st_case_792:
//line msg_parse.go:23913
		if data[p] == 10 {
			goto tr1076
		}
		goto st0
tr1076:
//line msg_parse.rl:344
 line++; linep = p; 
	goto st793
	st793:
		if p++; p == pe {
			goto _test_eof793
		}
	st_case_793:
//line msg_parse.go:23927
		switch data[p] {
		case 9:
			goto st794
		case 13:
			goto st15
		case 32:
			goto st794
		case 65:
			goto st16
		case 66:
			goto st22
		case 67:
			goto st27
		case 68:
			goto st33
		case 69:
			goto st41
		case 70:
			goto st47
		case 73:
			goto st53
		case 75:
			goto st76
		case 76:
			goto st81
		case 77:
			goto st87
		case 79:
			goto st93
		case 80:
			goto st99
		case 82:
			goto st122
		case 83:
			goto st128
		case 84:
			goto st134
		case 85:
			goto st140
		case 86:
			goto st146
		case 87:
			goto st152
		case 97:
			goto st16
		case 98:
			goto st22
		case 99:
			goto st27
		case 100:
			goto st33
		case 101:
			goto st41
		case 102:
			goto st47
		case 105:
			goto st53
		case 107:
			goto st76
		case 108:
			goto st81
		case 109:
			goto st87
		case 111:
			goto st93
		case 112:
			goto st99
		case 114:
			goto st122
		case 115:
			goto st128
		case 116:
			goto st134
		case 117:
			goto st140
		case 118:
			goto st146
		case 119:
			goto st152
		}
		goto st0
tr1078:
//line msg_parse.rl:47

      mark = p
    
	goto st794
	st794:
		if p++; p == pe {
			goto _test_eof794
		}
	st_case_794:
//line msg_parse.go:24020
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr1078
		case 32:
			goto tr1078
		case 269:
			goto tr1070
		case 525:
			goto tr1079
		}
		if 33 <= _widec && _widec <= 253 {
			goto tr1069
		}
		goto st0
tr2:
//line msg_parse.rl:47

      mark = p
    
	goto st795
	st795:
		if p++; p == pe {
			goto _test_eof795
		}
	st_case_795:
//line msg_parse.go:24053
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
			goto st796
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
	st796:
		if p++; p == pe {
			goto _test_eof796
		}
	st_case_796:
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
			goto st797
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
	st797:
		if p++; p == pe {
			goto _test_eof797
		}
	st_case_797:
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
			goto st798
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
	st798:
		if p++; p == pe {
			goto _test_eof798
		}
	st_case_798:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1083
		}
		goto st0
tr1083:
//line msg_parse.rl:82

      msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
    
	goto st799
	st799:
		if p++; p == pe {
			goto _test_eof799
		}
	st_case_799:
//line msg_parse.go:24190
		if data[p] == 46 {
			goto st800
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1083
		}
		goto st0
	st800:
		if p++; p == pe {
			goto _test_eof800
		}
	st_case_800:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1085
		}
		goto st0
tr1085:
//line msg_parse.rl:86

      msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
    
	goto st801
	st801:
		if p++; p == pe {
			goto _test_eof801
		}
	st_case_801:
//line msg_parse.go:24218
		if data[p] == 32 {
			goto st802
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1085
		}
		goto st0
	st802:
		if p++; p == pe {
			goto _test_eof802
		}
	st_case_802:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1087
		}
		goto st0
tr1087:
//line msg_parse.rl:95

      msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
    
	goto st803
	st803:
		if p++; p == pe {
			goto _test_eof803
		}
	st_case_803:
//line msg_parse.go:24246
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1088
		}
		goto st0
tr1088:
//line msg_parse.rl:95

      msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
    
	goto st804
	st804:
		if p++; p == pe {
			goto _test_eof804
		}
	st_case_804:
//line msg_parse.go:24262
		if 48 <= data[p] && data[p] <= 57 {
			goto tr1089
		}
		goto st0
tr1089:
//line msg_parse.rl:95

      msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
    
	goto st805
	st805:
		if p++; p == pe {
			goto _test_eof805
		}
	st_case_805:
//line msg_parse.go:24278
		if data[p] == 32 {
			goto st806
		}
		goto st0
	st806:
		if p++; p == pe {
			goto _test_eof806
		}
	st_case_806:
		switch data[p] {
		case 37:
			goto tr1092
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
		goto tr1091
tr1091:
//line msg_parse.rl:51

      amt = 0
    
//line msg_parse.rl:55

      buf[amt] = data[p]
      amt++
    
	goto st807
tr1093:
//line msg_parse.rl:55

      buf[amt] = data[p]
      amt++
    
	goto st807
tr1097:
//line msg_parse.rl:68

      hex += unhex(data[p])
      buf[amt] = hex
      amt++
    
	goto st807
	st807:
		if p++; p == pe {
			goto _test_eof807
		}
	st_case_807:
//line msg_parse.go:24358
		switch data[p] {
		case 13:
			goto tr1094
		case 37:
			goto st808
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
		goto tr1093
tr1092:
//line msg_parse.rl:51

      amt = 0
    
	goto st808
	st808:
		if p++; p == pe {
			goto _test_eof808
		}
	st_case_808:
//line msg_parse.go:24411
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1096
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr1096
			}
		default:
			goto tr1096
		}
		goto st0
tr1096:
//line msg_parse.rl:64

      hex = unhex(data[p]) * 16
    
	goto st809
	st809:
		if p++; p == pe {
			goto _test_eof809
		}
	st_case_809:
//line msg_parse.go:24436
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1097
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr1097
			}
		default:
			goto tr1097
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
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof810: cs = 810; goto _test_eof
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
	_test_eof634: cs = 634; goto _test_eof
	_test_eof635: cs = 635; goto _test_eof
	_test_eof636: cs = 636; goto _test_eof
	_test_eof637: cs = 637; goto _test_eof
	_test_eof638: cs = 638; goto _test_eof
	_test_eof639: cs = 639; goto _test_eof
	_test_eof640: cs = 640; goto _test_eof
	_test_eof641: cs = 641; goto _test_eof
	_test_eof642: cs = 642; goto _test_eof
	_test_eof643: cs = 643; goto _test_eof
	_test_eof644: cs = 644; goto _test_eof
	_test_eof645: cs = 645; goto _test_eof
	_test_eof646: cs = 646; goto _test_eof
	_test_eof647: cs = 647; goto _test_eof
	_test_eof648: cs = 648; goto _test_eof
	_test_eof649: cs = 649; goto _test_eof
	_test_eof650: cs = 650; goto _test_eof
	_test_eof651: cs = 651; goto _test_eof
	_test_eof652: cs = 652; goto _test_eof
	_test_eof653: cs = 653; goto _test_eof
	_test_eof654: cs = 654; goto _test_eof
	_test_eof655: cs = 655; goto _test_eof
	_test_eof656: cs = 656; goto _test_eof
	_test_eof657: cs = 657; goto _test_eof
	_test_eof658: cs = 658; goto _test_eof
	_test_eof659: cs = 659; goto _test_eof
	_test_eof660: cs = 660; goto _test_eof
	_test_eof661: cs = 661; goto _test_eof
	_test_eof662: cs = 662; goto _test_eof
	_test_eof663: cs = 663; goto _test_eof
	_test_eof664: cs = 664; goto _test_eof
	_test_eof665: cs = 665; goto _test_eof
	_test_eof666: cs = 666; goto _test_eof
	_test_eof667: cs = 667; goto _test_eof
	_test_eof668: cs = 668; goto _test_eof
	_test_eof669: cs = 669; goto _test_eof
	_test_eof670: cs = 670; goto _test_eof
	_test_eof671: cs = 671; goto _test_eof
	_test_eof672: cs = 672; goto _test_eof
	_test_eof673: cs = 673; goto _test_eof
	_test_eof674: cs = 674; goto _test_eof
	_test_eof675: cs = 675; goto _test_eof
	_test_eof676: cs = 676; goto _test_eof
	_test_eof677: cs = 677; goto _test_eof
	_test_eof678: cs = 678; goto _test_eof
	_test_eof679: cs = 679; goto _test_eof
	_test_eof680: cs = 680; goto _test_eof
	_test_eof681: cs = 681; goto _test_eof
	_test_eof682: cs = 682; goto _test_eof
	_test_eof683: cs = 683; goto _test_eof
	_test_eof684: cs = 684; goto _test_eof
	_test_eof685: cs = 685; goto _test_eof
	_test_eof686: cs = 686; goto _test_eof
	_test_eof687: cs = 687; goto _test_eof
	_test_eof688: cs = 688; goto _test_eof
	_test_eof689: cs = 689; goto _test_eof
	_test_eof690: cs = 690; goto _test_eof
	_test_eof691: cs = 691; goto _test_eof
	_test_eof692: cs = 692; goto _test_eof
	_test_eof693: cs = 693; goto _test_eof
	_test_eof694: cs = 694; goto _test_eof
	_test_eof695: cs = 695; goto _test_eof
	_test_eof696: cs = 696; goto _test_eof
	_test_eof697: cs = 697; goto _test_eof
	_test_eof698: cs = 698; goto _test_eof
	_test_eof699: cs = 699; goto _test_eof
	_test_eof700: cs = 700; goto _test_eof
	_test_eof701: cs = 701; goto _test_eof
	_test_eof702: cs = 702; goto _test_eof
	_test_eof703: cs = 703; goto _test_eof
	_test_eof704: cs = 704; goto _test_eof
	_test_eof705: cs = 705; goto _test_eof
	_test_eof706: cs = 706; goto _test_eof
	_test_eof707: cs = 707; goto _test_eof
	_test_eof708: cs = 708; goto _test_eof
	_test_eof709: cs = 709; goto _test_eof
	_test_eof710: cs = 710; goto _test_eof
	_test_eof711: cs = 711; goto _test_eof
	_test_eof712: cs = 712; goto _test_eof
	_test_eof713: cs = 713; goto _test_eof
	_test_eof714: cs = 714; goto _test_eof
	_test_eof715: cs = 715; goto _test_eof
	_test_eof716: cs = 716; goto _test_eof
	_test_eof717: cs = 717; goto _test_eof
	_test_eof718: cs = 718; goto _test_eof
	_test_eof719: cs = 719; goto _test_eof
	_test_eof720: cs = 720; goto _test_eof
	_test_eof721: cs = 721; goto _test_eof
	_test_eof722: cs = 722; goto _test_eof
	_test_eof723: cs = 723; goto _test_eof
	_test_eof724: cs = 724; goto _test_eof
	_test_eof725: cs = 725; goto _test_eof
	_test_eof726: cs = 726; goto _test_eof
	_test_eof727: cs = 727; goto _test_eof
	_test_eof728: cs = 728; goto _test_eof
	_test_eof729: cs = 729; goto _test_eof
	_test_eof730: cs = 730; goto _test_eof
	_test_eof731: cs = 731; goto _test_eof
	_test_eof732: cs = 732; goto _test_eof
	_test_eof733: cs = 733; goto _test_eof
	_test_eof734: cs = 734; goto _test_eof
	_test_eof735: cs = 735; goto _test_eof
	_test_eof736: cs = 736; goto _test_eof
	_test_eof737: cs = 737; goto _test_eof
	_test_eof738: cs = 738; goto _test_eof
	_test_eof739: cs = 739; goto _test_eof
	_test_eof740: cs = 740; goto _test_eof
	_test_eof741: cs = 741; goto _test_eof
	_test_eof742: cs = 742; goto _test_eof
	_test_eof743: cs = 743; goto _test_eof
	_test_eof744: cs = 744; goto _test_eof
	_test_eof745: cs = 745; goto _test_eof
	_test_eof746: cs = 746; goto _test_eof
	_test_eof747: cs = 747; goto _test_eof
	_test_eof748: cs = 748; goto _test_eof
	_test_eof749: cs = 749; goto _test_eof
	_test_eof750: cs = 750; goto _test_eof
	_test_eof751: cs = 751; goto _test_eof
	_test_eof752: cs = 752; goto _test_eof
	_test_eof753: cs = 753; goto _test_eof
	_test_eof754: cs = 754; goto _test_eof
	_test_eof755: cs = 755; goto _test_eof
	_test_eof756: cs = 756; goto _test_eof
	_test_eof757: cs = 757; goto _test_eof
	_test_eof758: cs = 758; goto _test_eof
	_test_eof759: cs = 759; goto _test_eof
	_test_eof760: cs = 760; goto _test_eof
	_test_eof761: cs = 761; goto _test_eof
	_test_eof762: cs = 762; goto _test_eof
	_test_eof763: cs = 763; goto _test_eof
	_test_eof764: cs = 764; goto _test_eof
	_test_eof765: cs = 765; goto _test_eof
	_test_eof766: cs = 766; goto _test_eof
	_test_eof767: cs = 767; goto _test_eof
	_test_eof768: cs = 768; goto _test_eof
	_test_eof769: cs = 769; goto _test_eof
	_test_eof770: cs = 770; goto _test_eof
	_test_eof771: cs = 771; goto _test_eof
	_test_eof772: cs = 772; goto _test_eof
	_test_eof773: cs = 773; goto _test_eof
	_test_eof774: cs = 774; goto _test_eof
	_test_eof775: cs = 775; goto _test_eof
	_test_eof776: cs = 776; goto _test_eof
	_test_eof777: cs = 777; goto _test_eof
	_test_eof778: cs = 778; goto _test_eof
	_test_eof779: cs = 779; goto _test_eof
	_test_eof780: cs = 780; goto _test_eof
	_test_eof781: cs = 781; goto _test_eof
	_test_eof782: cs = 782; goto _test_eof
	_test_eof783: cs = 783; goto _test_eof
	_test_eof784: cs = 784; goto _test_eof
	_test_eof785: cs = 785; goto _test_eof
	_test_eof786: cs = 786; goto _test_eof
	_test_eof787: cs = 787; goto _test_eof
	_test_eof788: cs = 788; goto _test_eof
	_test_eof789: cs = 789; goto _test_eof
	_test_eof790: cs = 790; goto _test_eof
	_test_eof791: cs = 791; goto _test_eof
	_test_eof792: cs = 792; goto _test_eof
	_test_eof793: cs = 793; goto _test_eof
	_test_eof794: cs = 794; goto _test_eof
	_test_eof795: cs = 795; goto _test_eof
	_test_eof796: cs = 796; goto _test_eof
	_test_eof797: cs = 797; goto _test_eof
	_test_eof798: cs = 798; goto _test_eof
	_test_eof799: cs = 799; goto _test_eof
	_test_eof800: cs = 800; goto _test_eof
	_test_eof801: cs = 801; goto _test_eof
	_test_eof802: cs = 802; goto _test_eof
	_test_eof803: cs = 803; goto _test_eof
	_test_eof804: cs = 804; goto _test_eof
	_test_eof805: cs = 805; goto _test_eof
	_test_eof806: cs = 806; goto _test_eof
	_test_eof807: cs = 807; goto _test_eof
	_test_eof808: cs = 808; goto _test_eof
	_test_eof809: cs = 809; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line msg_parse.rl:466


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
    msg.Payload = string(data[p:len(data)])
  }

  return msg, nil
}
