// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line msg_parse.rl:1
// -*-go-*-

package sip

import (
	"fmt"
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

// ParseMsg turns a SIP message byte slice into a data structure.
//
//line msg_parse.rl:14
func ParseMsg(data []byte) (msg *Msg, err error) {
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

//line msg_parse.rl:39

//line msg_parse.go:65
	{
		cs = msg_start
	}

//line msg_parse.rl:40

//line msg_parse.go:72
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
//line sip.rl:158

		p--

		{
			goto st273
		}

		goto st0
//line msg_parse.go:1691
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr0:
//line sip.rl:67

		mark = p

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line msg_parse.go:1707
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
//line sip.rl:99

		msg.Method = string(data[mark:p])

		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line msg_parse.go:1754
		if data[p] == 32 {
			goto st0
		}
		goto tr5
	tr5:
//line sip.rl:67

		mark = p

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line msg_parse.go:1770
		if data[p] == 32 {
			goto tr7
		}
		goto st4
	tr7:
//line sip.rl:111

		msg.Request, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line msg_parse.go:1787
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
//line sip.rl:103

		msg.VersionMajor = msg.VersionMajor*10 + (data[p] - 0x30)

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line msg_parse.go:1839
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
//line sip.rl:107

		msg.VersionMinor = msg.VersionMinor*10 + (data[p] - 0x30)

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line msg_parse.go:1867
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:120

		msg.Phrase = string(buf[0:amt])

		goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line msg_parse.go:1893
		if data[p] == 10 {
			goto tr16
		}
		goto st0
	tr16:
//line sip.rl:244
		{
			goto st280
		}
		goto st765
	st765:
		if p++; p == pe {
			goto _test_eof765
		}
	st_case_765:
//line msg_parse.go:1907
		goto st0
	tr2:
//line sip.rl:67

		mark = p

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line msg_parse.go:1920
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
//line sip.rl:103

		msg.VersionMajor = msg.VersionMajor*10 + (data[p] - 0x30)

		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line msg_parse.go:2057
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
//line sip.rl:107

		msg.VersionMinor = msg.VersionMinor*10 + (data[p] - 0x30)

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line msg_parse.go:2085
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
//line sip.rl:116

		msg.Status = msg.Status*10 + (int(data[p]) - 0x30)

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line msg_parse.go:2113
		if 48 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
	tr25:
//line sip.rl:116

		msg.Status = msg.Status*10 + (int(data[p]) - 0x30)

		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line msg_parse.go:2129
		if 48 <= data[p] && data[p] <= 57 {
			goto tr26
		}
		goto st0
	tr26:
//line sip.rl:116

		msg.Status = msg.Status*10 + (int(data[p]) - 0x30)

		goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line msg_parse.go:2145
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st26
	tr35:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st26
	tr44:
//line sip.rl:93

		hex += unhex(data[p])
		buf[amt] = hex
		amt++

		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line msg_parse.go:2239
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line msg_parse.go:2313
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
//line sip.rl:89

		hex = unhex(data[p]) * 16

		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line msg_parse.go:2338
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st29
	tr37:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line msg_parse.go:2375
		if 128 <= data[p] && data[p] <= 191 {
			goto tr35
		}
		goto st0
	tr31:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st30
	tr38:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line msg_parse.go:2403
		if 128 <= data[p] && data[p] <= 191 {
			goto tr37
		}
		goto st0
	tr32:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st31
	tr39:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line msg_parse.go:2431
		if 128 <= data[p] && data[p] <= 191 {
			goto tr38
		}
		goto st0
	tr33:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st32
	tr40:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line msg_parse.go:2459
		if 128 <= data[p] && data[p] <= 191 {
			goto tr39
		}
		goto st0
	tr34:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st33
	tr41:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line msg_parse.go:2487
		if 128 <= data[p] && data[p] <= 191 {
			goto tr40
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
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
//line sip.rl:67

		mark = p

		goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line msg_parse.go:2541
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
		if p++; p == pe {
			goto _test_eof36
		}
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
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:215

		ctype = string(data[mark:p])

		goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line msg_parse.go:2676
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:215

		ctype = string(data[mark:p])

		goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line msg_parse.go:2706
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:215

		ctype = string(data[mark:p])

		goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line msg_parse.go:2948
		if data[p] == 10 {
			goto tr65
		}
		goto st0
	tr65:
//line sip.rl:244
		{
			goto st280
		}
		goto st766
	st766:
		if p++; p == pe {
			goto _test_eof766
		}
	st_case_766:
//line msg_parse.go:2962
		goto st0
	tr52:
//line sip.rl:215

		ctype = string(data[mark:p])

		goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line msg_parse.go:2975
		if data[p] == 10 {
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 9:
			goto st47
		case 32:
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
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
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st49
	tr77:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line msg_parse.go:3078
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line msg_parse.go:3134
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line msg_parse.go:3166
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st52
	tr80:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line msg_parse.go:3203
		if 128 <= data[p] && data[p] <= 191 {
			goto tr77
		}
		goto st0
	tr72:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st53
	tr81:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line msg_parse.go:3231
		if 128 <= data[p] && data[p] <= 191 {
			goto tr80
		}
		goto st0
	tr73:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st54
	tr82:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line msg_parse.go:3259
		if 128 <= data[p] && data[p] <= 191 {
			goto tr81
		}
		goto st0
	tr74:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st55
	tr83:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line msg_parse.go:3287
		if 128 <= data[p] && data[p] <= 191 {
			goto tr82
		}
		goto st0
	tr75:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st56
	tr84:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line msg_parse.go:3315
		if 128 <= data[p] && data[p] <= 191 {
			goto tr83
		}
		goto st0
	tr76:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st57
	tr85:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line msg_parse.go:3343
		if data[p] == 10 {
			goto tr86
		}
		goto st0
	tr86:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line msg_parse.go:3360
		switch data[p] {
		case 9:
			goto tr77
		case 32:
			goto tr77
		}
		goto st0
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
		if data[p] == 10 {
			goto st60
		}
		goto st0
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
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
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		if data[p] == 10 {
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
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
		case 61:
			goto st42
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		if data[p] == 10 {
			goto st66
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 9:
			goto st67
		case 32:
			goto st67
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
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
		if p++; p == pe {
			goto _test_eof68
		}
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
//line sip.rl:75

		amt = 0

//line sip.rl:67

		mark = p

		goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line msg_parse.go:3584
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line msg_parse.go:3645
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line msg_parse.go:3679
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:154

		via.Param = &Param{name, string(buf[0:amt]), via.Param}

//line sip.rl:128

		*viap = via
		viap = &via.Next
		via = nil

//line sip.rl:124

		via = new(Via)

//line sip.rl:59

		p--

//line sip.rl:246
		{
			goto st103
		}
		goto st767
	tr110:
//line sip.rl:154

		via.Param = &Param{name, string(buf[0:amt]), via.Param}

//line sip.rl:59

		p--

//line sip.rl:75

		amt = 0

//line sip.rl:247
		{
			goto st68
		}
		goto st767
	tr122:
//line sip.rl:154

		via.Param = &Param{name, string(buf[0:amt]), via.Param}

//line sip.rl:128

		*viap = via
		viap = &via.Next
		via = nil

//line sip.rl:244
		{
			goto st280
		}
		goto st767
	st767:
		if p++; p == pe {
			goto _test_eof767
		}
	st_case_767:
//line msg_parse.go:3765
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		if data[p] == 10 {
			goto st73
		}
		goto st0
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
		switch data[p] {
		case 9:
			goto st74
		case 32:
			goto st74
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 9:
			goto st74
		case 32:
			goto st74
		}
		goto tr106
	tr97:
//line sip.rl:163

		name = string(data[mark:p])

		goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line msg_parse.go:3811
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		if data[p] == 10 {
			goto st77
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 9:
			goto st78
		case 32:
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
			goto st78
		case 32:
			goto st78
		}
		goto tr110
	tr98:
//line sip.rl:163

		name = string(data[mark:p])

		goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line msg_parse.go:3882
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line msg_parse.go:3945
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
		if data[p] == 10 {
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
			goto st84
		case 32:
			goto st84
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st85
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
//line msg_parse.go:4068
		if data[p] == 10 {
			goto tr122
		}
		goto st0
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st87
	tr132:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line msg_parse.go:4145
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st88
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
//line msg_parse.go:4201
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st89
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
//line msg_parse.go:4235
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st90
	tr135:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line msg_parse.go:4272
		if 128 <= data[p] && data[p] <= 191 {
			goto tr132
		}
		goto st0
	tr127:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st91
	tr136:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line msg_parse.go:4300
		if 128 <= data[p] && data[p] <= 191 {
			goto tr135
		}
		goto st0
	tr128:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st92
	tr137:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line msg_parse.go:4328
		if 128 <= data[p] && data[p] <= 191 {
			goto tr136
		}
		goto st0
	tr129:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st93
	tr138:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line msg_parse.go:4356
		if 128 <= data[p] && data[p] <= 191 {
			goto tr137
		}
		goto st0
	tr130:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st94
	tr139:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line msg_parse.go:4384
		if 128 <= data[p] && data[p] <= 191 {
			goto tr138
		}
		goto st0
	tr131:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st95
	tr140:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line msg_parse.go:4412
		if data[p] == 10 {
			goto tr141
		}
		goto st0
	tr141:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line msg_parse.go:4429
		switch data[p] {
		case 9:
			goto tr132
		case 32:
			goto tr132
		}
		goto st0
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
		if data[p] == 10 {
			goto st98
		}
		goto st0
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
		switch data[p] {
		case 9:
			goto st99
		case 32:
			goto st99
		}
		goto st0
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line msg_parse.go:4515
		if data[p] == 10 {
			goto st101
		}
		goto st0
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
		switch data[p] {
		case 9:
			goto st102
		case 32:
			goto st102
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
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
		if p++; p == pe {
			goto _test_eof103
		}
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
//line sip.rl:67

		mark = p

		goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line msg_parse.go:4599
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:134

		via.Protocol = string(data[mark:p])

		goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line msg_parse.go:4654
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:134

		via.Protocol = string(data[mark:p])

		goto st106
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
//line msg_parse.go:4684
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line msg_parse.go:4742
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:138

		via.Version = string(data[mark:p])

		goto st108
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
//line msg_parse.go:4797
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:138

		via.Version = string(data[mark:p])

		goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line msg_parse.go:4827
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line msg_parse.go:4885
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:142

		via.Transport = string(data[mark:p])

		goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line msg_parse.go:4943
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st112
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
//line msg_parse.go:4990
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line msg_parse.go:5043
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st114
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
//line msg_parse.go:5077
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:128

		*viap = via
		viap = &via.Next
		via = nil

//line sip.rl:124

		via = new(Via)

//line sip.rl:59

		p--

//line sip.rl:246
		{
			goto st103
		}
		goto st768
	tr193:
//line sip.rl:59

		p--

//line sip.rl:75

		amt = 0

//line sip.rl:247
		{
			goto st68
		}
		goto st768
	tr199:
//line sip.rl:128

		*viap = via
		viap = &via.Next
		via = nil

//line sip.rl:244
		{
			goto st280
		}
		goto st768
	st768:
		if p++; p == pe {
			goto _test_eof768
		}
	st_case_768:
//line msg_parse.go:5151
		goto st0
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
		if data[p] == 10 {
			goto st116
		}
		goto st0
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
		switch data[p] {
		case 9:
			goto st117
		case 32:
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
		}
		goto tr184
	tr175:
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line msg_parse.go:5197
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:150

		via.Port = via.Port*10 + (uint16(data[p]) - 0x30)

		goto st119
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
//line msg_parse.go:5228
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line msg_parse.go:5290
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		if data[p] == 10 {
			goto st123
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 9:
			goto st124
		case 32:
			goto st124
		}
		goto st0
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
		switch data[p] {
		case 9:
			goto st124
		case 32:
			goto st124
		}
		goto tr193
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		if data[p] == 10 {
			goto st126
		}
		goto st0
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
		switch data[p] {
		case 9:
			goto st127
		case 32:
			goto st127
		}
		goto st0
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
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
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line msg_parse.go:5398
		if data[p] == 10 {
			goto tr199
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		if data[p] == 10 {
			goto st130
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 9:
			goto st131
		case 32:
			goto st131
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
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr188
		}
		goto st0
	tr178:
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line msg_parse.go:5450
		if data[p] == 10 {
			goto st133
		}
		goto st0
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
		switch data[p] {
		case 9:
			goto st134
		case 32:
			goto st134
		}
		goto st0
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
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
		if p++; p == pe {
			goto _test_eof135
		}
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
//line sip.rl:67

		mark = p

		goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line msg_parse.go:5517
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
//line sip.rl:146

		via.Host = string(data[mark:p])

		goto st137
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
//line msg_parse.go:5548
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:142

		via.Transport = string(data[mark:p])

		goto st138
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
//line msg_parse.go:5584
		if data[p] == 10 {
			goto st139
		}
		goto st0
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
		switch data[p] {
		case 9:
			goto st140
		case 32:
			goto st140
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
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
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
		if data[p] == 10 {
			goto st142
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 9:
			goto st143
		case 32:
			goto st143
		}
		goto st0
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
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
//line sip.rl:138

		via.Version = string(data[mark:p])

		goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line msg_parse.go:5706
		if data[p] == 10 {
			goto st145
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 9:
			goto st146
		case 32:
			goto st146
		}
		goto st0
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
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
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
		if data[p] == 10 {
			goto st148
		}
		goto st0
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
		switch data[p] {
		case 9:
			goto st149
		case 32:
			goto st149
		}
		goto st0
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
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
//line sip.rl:134

		via.Protocol = string(data[mark:p])

		goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line msg_parse.go:5811
		if data[p] == 10 {
			goto st151
		}
		goto st0
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
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
		case 47:
			goto st106
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
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
//line sip.rl:75

		amt = 0

//line sip.rl:67

		mark = p

		goto st154
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
//line msg_parse.go:5895
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st155
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
//line msg_parse.go:5956
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st156
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
//line msg_parse.go:5990
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:197

		addr.Param = &Param{name, string(buf[0:amt]), addr.Param}

//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st769
	tr234:
//line sip.rl:197

		addr.Param = &Param{name, string(buf[0:amt]), addr.Param}

//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st769
	tr246:
//line sip.rl:197

		addr.Param = &Param{name, string(buf[0:amt]), addr.Param}

//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:244
		{
			goto st280
		}
		goto st769
	st769:
		if p++; p == pe {
			goto _test_eof769
		}
	st_case_769:
//line msg_parse.go:6068
		goto st0
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
		if data[p] == 10 {
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
			goto st159
		case 32:
			goto st159
		}
		goto st0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 9:
			goto st159
		case 32:
			goto st159
		}
		goto tr230
	tr221:
//line sip.rl:163

		name = string(data[mark:p])

		goto st160
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
//line msg_parse.go:6114
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
		if data[p] == 10 {
			goto st162
		}
		goto st0
	st162:
		if p++; p == pe {
			goto _test_eof162
		}
	st_case_162:
		switch data[p] {
		case 9:
			goto st163
		case 32:
			goto st163
		}
		goto st0
	st163:
		if p++; p == pe {
			goto _test_eof163
		}
	st_case_163:
		switch data[p] {
		case 9:
			goto st163
		case 32:
			goto st163
		}
		goto tr234
	tr222:
//line sip.rl:163

		name = string(data[mark:p])

		goto st164
	st164:
		if p++; p == pe {
			goto _test_eof164
		}
	st_case_164:
//line msg_parse.go:6185
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st165
	st165:
		if p++; p == pe {
			goto _test_eof165
		}
	st_case_165:
//line msg_parse.go:6248
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof166
		}
	st_case_166:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof167
		}
	st_case_167:
		if data[p] == 10 {
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
			goto st169
		case 32:
			goto st169
		}
		goto st0
	st169:
		if p++; p == pe {
			goto _test_eof169
		}
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st170
	st170:
		if p++; p == pe {
			goto _test_eof170
		}
	st_case_170:
//line msg_parse.go:6371
		if data[p] == 10 {
			goto tr246
		}
		goto st0
	st171:
		if p++; p == pe {
			goto _test_eof171
		}
	st_case_171:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st172
	tr256:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st172
	st172:
		if p++; p == pe {
			goto _test_eof172
		}
	st_case_172:
//line msg_parse.go:6448
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st173
	st173:
		if p++; p == pe {
			goto _test_eof173
		}
	st_case_173:
//line msg_parse.go:6504
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st174
	st174:
		if p++; p == pe {
			goto _test_eof174
		}
	st_case_174:
//line msg_parse.go:6538
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st175
	tr259:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st175
	st175:
		if p++; p == pe {
			goto _test_eof175
		}
	st_case_175:
//line msg_parse.go:6575
		if 128 <= data[p] && data[p] <= 191 {
			goto tr256
		}
		goto st0
	tr251:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st176
	tr260:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st176
	st176:
		if p++; p == pe {
			goto _test_eof176
		}
	st_case_176:
//line msg_parse.go:6603
		if 128 <= data[p] && data[p] <= 191 {
			goto tr259
		}
		goto st0
	tr252:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st177
	tr261:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st177
	st177:
		if p++; p == pe {
			goto _test_eof177
		}
	st_case_177:
//line msg_parse.go:6631
		if 128 <= data[p] && data[p] <= 191 {
			goto tr260
		}
		goto st0
	tr253:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st178
	tr262:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st178
	st178:
		if p++; p == pe {
			goto _test_eof178
		}
	st_case_178:
//line msg_parse.go:6659
		if 128 <= data[p] && data[p] <= 191 {
			goto tr261
		}
		goto st0
	tr254:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st179
	tr263:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st179
	st179:
		if p++; p == pe {
			goto _test_eof179
		}
	st_case_179:
//line msg_parse.go:6687
		if 128 <= data[p] && data[p] <= 191 {
			goto tr262
		}
		goto st0
	tr255:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st180
	tr264:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st180
	st180:
		if p++; p == pe {
			goto _test_eof180
		}
	st_case_180:
//line msg_parse.go:6715
		if data[p] == 10 {
			goto tr265
		}
		goto st0
	tr265:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st181
	st181:
		if p++; p == pe {
			goto _test_eof181
		}
	st_case_181:
//line msg_parse.go:6732
		switch data[p] {
		case 9:
			goto tr256
		case 32:
			goto tr256
		}
		goto st0
	st182:
		if p++; p == pe {
			goto _test_eof182
		}
	st_case_182:
		if data[p] == 10 {
			goto st183
		}
		goto st0
	st183:
		if p++; p == pe {
			goto _test_eof183
		}
	st_case_183:
		switch data[p] {
		case 9:
			goto st184
		case 32:
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st185
	st185:
		if p++; p == pe {
			goto _test_eof185
		}
	st_case_185:
//line msg_parse.go:6818
		if data[p] == 10 {
			goto st186
		}
		goto st0
	st186:
		if p++; p == pe {
			goto _test_eof186
		}
	st_case_186:
		switch data[p] {
		case 9:
			goto st187
		case 32:
			goto st187
		}
		goto st0
	st187:
		if p++; p == pe {
			goto _test_eof187
		}
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
		if p++; p == pe {
			goto _test_eof188
		}
	st_case_188:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:180

		addr.Display = string(buf[0:amt])

		goto st189
	st189:
		if p++; p == pe {
			goto _test_eof189
		}
	st_case_189:
//line msg_parse.go:6919
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:184
		{
			end := p
			for end > mark && whitespacec(data[end-1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
		goto st190
	tr330:
//line sip.rl:180

		addr.Display = string(buf[0:amt])

		goto st190
	st190:
		if p++; p == pe {
			goto _test_eof190
		}
	st_case_190:
//line msg_parse.go:6959
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
//line sip.rl:67

		mark = p

		goto st191
	st191:
		if p++; p == pe {
			goto _test_eof191
		}
	st_case_191:
//line msg_parse.go:6980
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
		if p++; p == pe {
			goto _test_eof192
		}
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
		if p++; p == pe {
			goto _test_eof193
		}
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st194
	st194:
		if p++; p == pe {
			goto _test_eof194
		}
	st_case_194:
//line msg_parse.go:7077
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof195
		}
	st_case_195:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st770
	tr289:
//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st770
	tr293:
//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:244
		{
			goto st280
		}
		goto st770
	st770:
		if p++; p == pe {
			goto _test_eof770
		}
	st_case_770:
//line msg_parse.go:7170
		goto st0
	st196:
		if p++; p == pe {
			goto _test_eof196
		}
	st_case_196:
		if data[p] == 10 {
			goto st197
		}
		goto st0
	st197:
		if p++; p == pe {
			goto _test_eof197
		}
	st_case_197:
		switch data[p] {
		case 9:
			goto st198
		case 32:
			goto st198
		}
		goto st0
	st198:
		if p++; p == pe {
			goto _test_eof198
		}
	st_case_198:
		switch data[p] {
		case 9:
			goto st198
		case 32:
			goto st198
		}
		goto tr285
	st199:
		if p++; p == pe {
			goto _test_eof199
		}
	st_case_199:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof200
		}
	st_case_200:
		if data[p] == 10 {
			goto st201
		}
		goto st0
	st201:
		if p++; p == pe {
			goto _test_eof201
		}
	st_case_201:
		switch data[p] {
		case 9:
			goto st202
		case 32:
			goto st202
		}
		goto st0
	st202:
		if p++; p == pe {
			goto _test_eof202
		}
	st_case_202:
		switch data[p] {
		case 9:
			goto st202
		case 32:
			goto st202
		}
		goto tr289
	st203:
		if p++; p == pe {
			goto _test_eof203
		}
	st_case_203:
		if data[p] == 10 {
			goto tr293
		}
		goto st0
	st204:
		if p++; p == pe {
			goto _test_eof204
		}
	st_case_204:
		if data[p] == 10 {
			goto st205
		}
		goto st0
	st205:
		if p++; p == pe {
			goto _test_eof205
		}
	st_case_205:
		switch data[p] {
		case 9:
			goto st206
		case 32:
			goto st206
		}
		goto st0
	st206:
		if p++; p == pe {
			goto _test_eof206
		}
	st_case_206:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof207
		}
	st_case_207:
		if data[p] == 10 {
			goto st208
		}
		goto st0
	st208:
		if p++; p == pe {
			goto _test_eof208
		}
	st_case_208:
		switch data[p] {
		case 9:
			goto st209
		case 32:
			goto st209
		}
		goto st0
	st209:
		if p++; p == pe {
			goto _test_eof209
		}
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
//line sip.rl:184
		{
			end := p
			for end > mark && whitespacec(data[end-1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
		goto st210
	tr331:
//line sip.rl:180

		addr.Display = string(buf[0:amt])

		goto st210
	st210:
		if p++; p == pe {
			goto _test_eof210
		}
	st_case_210:
//line msg_parse.go:7384
		if data[p] == 10 {
			goto st211
		}
		goto st0
	st211:
		if p++; p == pe {
			goto _test_eof211
		}
	st_case_211:
		switch data[p] {
		case 9:
			goto st212
		case 32:
			goto st212
		}
		goto st0
	st212:
		if p++; p == pe {
			goto _test_eof212
		}
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
//line sip.rl:67

		mark = p

		goto st213
	st213:
		if p++; p == pe {
			goto _test_eof213
		}
	st_case_213:
//line msg_parse.go:7426
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:184
		{
			end := p
			for end > mark && whitespacec(data[end-1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
		goto st214
	st214:
		if p++; p == pe {
			goto _test_eof214
		}
	st_case_214:
//line msg_parse.go:7488
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:184
		{
			end := p
			for end > mark && whitespacec(data[end-1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
		goto st215
	st215:
		if p++; p == pe {
			goto _test_eof215
		}
	st_case_215:
//line msg_parse.go:7552
		if data[p] == 10 {
			goto st216
		}
		goto st0
	st216:
		if p++; p == pe {
			goto _test_eof216
		}
	st_case_216:
		switch data[p] {
		case 9:
			goto st217
		case 32:
			goto st217
		}
		goto st0
	tr309:
//line sip.rl:184
		{
			end := p
			for end > mark && whitespacec(data[end-1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
		goto st217
	st217:
		if p++; p == pe {
			goto _test_eof217
		}
	st_case_217:
//line msg_parse.go:7584
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st218
	st218:
		if p++; p == pe {
			goto _test_eof218
		}
	st_case_218:
//line msg_parse.go:7644
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st219
	tr320:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st219
	st219:
		if p++; p == pe {
			goto _test_eof219
		}
	st_case_219:
//line msg_parse.go:7712
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st220
	st220:
		if p++; p == pe {
			goto _test_eof220
		}
	st_case_220:
//line msg_parse.go:7768
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:75

		amt = 0

		goto st221
	st221:
		if p++; p == pe {
			goto _test_eof221
		}
	st_case_221:
//line msg_parse.go:7798
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
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st222
	tr323:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st222
	st222:
		if p++; p == pe {
			goto _test_eof222
		}
	st_case_222:
//line msg_parse.go:7835
		if 128 <= data[p] && data[p] <= 191 {
			goto tr320
		}
		goto st0
	tr315:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st223
	tr324:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st223
	st223:
		if p++; p == pe {
			goto _test_eof223
		}
	st_case_223:
//line msg_parse.go:7863
		if 128 <= data[p] && data[p] <= 191 {
			goto tr323
		}
		goto st0
	tr316:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st224
	tr325:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st224
	st224:
		if p++; p == pe {
			goto _test_eof224
		}
	st_case_224:
//line msg_parse.go:7891
		if 128 <= data[p] && data[p] <= 191 {
			goto tr324
		}
		goto st0
	tr317:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st225
	tr326:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st225
	st225:
		if p++; p == pe {
			goto _test_eof225
		}
	st_case_225:
//line msg_parse.go:7919
		if 128 <= data[p] && data[p] <= 191 {
			goto tr325
		}
		goto st0
	tr318:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st226
	tr327:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st226
	st226:
		if p++; p == pe {
			goto _test_eof226
		}
	st_case_226:
//line msg_parse.go:7947
		if 128 <= data[p] && data[p] <= 191 {
			goto tr326
		}
		goto st0
	tr319:
//line sip.rl:75

		amt = 0

//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st227
	tr328:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st227
	st227:
		if p++; p == pe {
			goto _test_eof227
		}
	st_case_227:
//line msg_parse.go:7975
		if data[p] == 10 {
			goto tr332
		}
		goto st0
	tr332:
//line sip.rl:79

		buf[amt] = data[p]
		amt++

		goto st228
	st228:
		if p++; p == pe {
			goto _test_eof228
		}
	st_case_228:
//line msg_parse.go:7992
		switch data[p] {
		case 9:
			goto tr320
		case 32:
			goto tr320
		}
		goto st0
	st229:
		if p++; p == pe {
			goto _test_eof229
		}
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
		if p++; p == pe {
			goto _test_eof230
		}
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
		if p++; p == pe {
			goto _test_eof231
		}
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
		if p++; p == pe {
			goto _test_eof232
		}
	st_case_232:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st233
	st233:
		if p++; p == pe {
			goto _test_eof233
		}
	st_case_233:
//line msg_parse.go:8134
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof234
		}
	st_case_234:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st771
	tr349:
//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st771
	st771:
		if p++; p == pe {
			goto _test_eof771
		}
	st_case_771:
//line msg_parse.go:8215
		goto st0
	st235:
		if p++; p == pe {
			goto _test_eof235
		}
	st_case_235:
		if data[p] == 10 {
			goto st236
		}
		goto st0
	st236:
		if p++; p == pe {
			goto _test_eof236
		}
	st_case_236:
		switch data[p] {
		case 9:
			goto st237
		case 32:
			goto st237
		}
		goto st0
	st237:
		if p++; p == pe {
			goto _test_eof237
		}
	st_case_237:
		switch data[p] {
		case 9:
			goto st237
		case 32:
			goto st237
		}
		goto tr345
	st238:
		if p++; p == pe {
			goto _test_eof238
		}
	st_case_238:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof239
		}
	st_case_239:
		if data[p] == 10 {
			goto st240
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
		}
		goto tr349
	tr340:
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st242
	st242:
		if p++; p == pe {
			goto _test_eof242
		}
	st_case_242:
//line msg_parse.go:8326
		if data[p] == 10 {
			goto st243
		}
		goto st0
	st243:
		if p++; p == pe {
			goto _test_eof243
		}
	st_case_243:
		switch data[p] {
		case 9:
			goto st244
		case 32:
			goto st244
		}
		goto st0
	st244:
		if p++; p == pe {
			goto _test_eof244
		}
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st245
	st245:
		if p++; p == pe {
			goto _test_eof245
		}
	st_case_245:
//line msg_parse.go:8371
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st246
	st246:
		if p++; p == pe {
			goto _test_eof246
		}
	st_case_246:
//line msg_parse.go:8437
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st772
	st772:
		if p++; p == pe {
			goto _test_eof772
		}
	st_case_772:
//line msg_parse.go:8488
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st247
	st247:
		if p++; p == pe {
			goto _test_eof247
		}
	st_case_247:
//line msg_parse.go:8527
		if data[p] == 10 {
			goto st248
		}
		goto st0
	st248:
		if p++; p == pe {
			goto _test_eof248
		}
	st_case_248:
		switch data[p] {
		case 9:
			goto st249
		case 32:
			goto st249
		}
		goto st0
	st249:
		if p++; p == pe {
			goto _test_eof249
		}
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
//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st773
	tr366:
//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st773
	st773:
		if p++; p == pe {
			goto _test_eof773
		}
	st_case_773:
//line msg_parse.go:8589
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st250
	st250:
		if p++; p == pe {
			goto _test_eof250
		}
	st_case_250:
//line msg_parse.go:8646
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st251
	st251:
		if p++; p == pe {
			goto _test_eof251
		}
	st_case_251:
//line msg_parse.go:8712
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st774
	st774:
		if p++; p == pe {
			goto _test_eof774
		}
	st_case_774:
//line msg_parse.go:8757
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st252
	st252:
		if p++; p == pe {
			goto _test_eof252
		}
	st_case_252:
//line msg_parse.go:8796
		if data[p] == 10 {
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
			goto st254
		case 32:
			goto st254
		}
		goto st0
	st254:
		if p++; p == pe {
			goto _test_eof254
		}
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st775
	st775:
		if p++; p == pe {
			goto _test_eof775
		}
	st_case_775:
//line msg_parse.go:8848
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st776
	st776:
		if p++; p == pe {
			goto _test_eof776
		}
	st_case_776:
//line msg_parse.go:8927
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

//line sip.rl:201

		*addrp = addr
		addrp = &addr.Next
		addr = nil

//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st777
	tr368:
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

//line sip.rl:59

		p--

//line sip.rl:241
		{
			goto st153
		}
		goto st777
	st777:
		if p++; p == pe {
			goto _test_eof777
		}
	st_case_777:
//line msg_parse.go:9020
		if data[p] == 10 {
			goto tr375
		}
		goto st0
	tr375:
//line sip.rl:244
		{
			goto st280
		}
		goto st778
	st778:
		if p++; p == pe {
			goto _test_eof778
		}
	st_case_778:
//line msg_parse.go:9034
		goto st0
	tr339:
//line sip.rl:192

		addr.Uri, err = ParseURI(data[mark:p])
		if err != nil {
			return nil, err
		}

		goto st255
	st255:
		if p++; p == pe {
			goto _test_eof255
		}
	st_case_255:
//line msg_parse.go:9048
		if data[p] == 10 {
			goto tr375
		}
		goto st0
	st256:
		if p++; p == pe {
			goto _test_eof256
		}
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
//line sip.rl:67

		mark = p

		goto st257
	st257:
		if p++; p == pe {
			goto _test_eof257
		}
	st_case_257:
//line msg_parse.go:9110
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof258
		}
	st_case_258:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:176

		addr = new(Addr)

//line sip.rl:59

		p--

//line sip.rl:240
		{
			goto st188
		}
		goto st779
	tr382:
//line sip.rl:176

		addr = new(Addr)

//line sip.rl:71

		p = (mark) - 1

//line sip.rl:240
		{
			goto st188
		}
		goto st779
	tr386:
//line sip.rl:176

		addr = new(Addr)

//line sip.rl:71

		p = (mark) - 1

//line sip.rl:242
		{
			goto st229
		}
		goto st779
	st779:
		if p++; p == pe {
			goto _test_eof779
		}
	st_case_779:
//line msg_parse.go:9254
		goto st0
	st259:
		if p++; p == pe {
			goto _test_eof259
		}
	st_case_259:
		if data[p] == 10 {
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
			goto st261
		case 32:
			goto st261
		}
		goto st0
	st261:
		if p++; p == pe {
			goto _test_eof261
		}
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
//line sip.rl:67

		mark = p

		goto st262
	st262:
		if p++; p == pe {
			goto _test_eof262
		}
	st_case_262:
//line msg_parse.go:9332
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof263
		}
	st_case_263:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st264
	st264:
		if p++; p == pe {
			goto _test_eof264
		}
	st_case_264:
//line msg_parse.go:9443
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st265
	st265:
		if p++; p == pe {
			goto _test_eof265
		}
	st_case_265:
//line msg_parse.go:9497
		if 128 <= data[p] && data[p] <= 191 {
			goto st264
		}
		goto st0
	tr389:
//line sip.rl:67

		mark = p

		goto st266
	st266:
		if p++; p == pe {
			goto _test_eof266
		}
	st_case_266:
//line msg_parse.go:9513
		if 128 <= data[p] && data[p] <= 191 {
			goto st265
		}
		goto st0
	tr390:
//line sip.rl:67

		mark = p

		goto st267
	st267:
		if p++; p == pe {
			goto _test_eof267
		}
	st_case_267:
//line msg_parse.go:9529
		if 128 <= data[p] && data[p] <= 191 {
			goto st266
		}
		goto st0
	tr391:
//line sip.rl:67

		mark = p

		goto st268
	st268:
		if p++; p == pe {
			goto _test_eof268
		}
	st_case_268:
//line msg_parse.go:9545
		if 128 <= data[p] && data[p] <= 191 {
			goto st267
		}
		goto st0
	tr392:
//line sip.rl:67

		mark = p

		goto st269
	st269:
		if p++; p == pe {
			goto _test_eof269
		}
	st_case_269:
//line msg_parse.go:9561
		if 128 <= data[p] && data[p] <= 191 {
			goto st268
		}
		goto st0
	tr393:
//line sip.rl:67

		mark = p

		goto st270
	st270:
		if p++; p == pe {
			goto _test_eof270
		}
	st_case_270:
//line msg_parse.go:9577
		if data[p] == 10 {
			goto tr403
		}
		goto st0
	tr403:
//line sip.rl:167
		{
			b := data[mark : p-1]
			if value != nil {
				*value = string(b)
			} else {
				msg.XHeader = &XHeader{name, b, msg.XHeader}
			}
		}
//line sip.rl:244
		{
			goto st280
		}
		goto st780
	st780:
		if p++; p == pe {
			goto _test_eof780
		}
	st_case_780:
//line msg_parse.go:9600
		goto st0
	tr394:
//line sip.rl:67

		mark = p

		goto st271
	st271:
		if p++; p == pe {
			goto _test_eof271
		}
	st_case_271:
//line msg_parse.go:9613
		if data[p] == 10 {
			goto st272
		}
		goto st0
	st272:
		if p++; p == pe {
			goto _test_eof272
		}
	st_case_272:
		switch data[p] {
		case 9:
			goto st264
		case 32:
			goto st264
		}
		goto st0
	st273:
		if p++; p == pe {
			goto _test_eof273
		}
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
		if p++; p == pe {
			goto _test_eof274
		}
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st275
	st275:
		if p++; p == pe {
			goto _test_eof275
		}
	st_case_275:
//line msg_parse.go:9723
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
//line sip.rl:163

		name = string(data[mark:p])

		goto st276
	st276:
		if p++; p == pe {
			goto _test_eof276
		}
	st_case_276:
//line msg_parse.go:9744
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:519
		value = nil
//line sip.rl:59

		p--

//line sip.rl:245
		{
			goto st263
		}
		goto st781
	st781:
		if p++; p == pe {
			goto _test_eof781
		}
	st_case_781:
//line msg_parse.go:9787
		goto st0
	st277:
		if p++; p == pe {
			goto _test_eof277
		}
	st_case_277:
		if data[p] == 10 {
			goto st278
		}
		goto st0
	st278:
		if p++; p == pe {
			goto _test_eof278
		}
	st_case_278:
		switch data[p] {
		case 9:
			goto st279
		case 32:
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
		}
		goto tr410
	st280:
		if p++; p == pe {
			goto _test_eof280
		}
	st_case_280:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

//line sip.rl:59

		p--

		goto st281
	st281:
		if p++; p == pe {
			goto _test_eof281
		}
	st_case_281:
//line msg_parse.go:9885
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
		if p++; p == pe {
			goto _test_eof282
		}
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
//line sip.rl:438
		value = &msg.AcceptContact
		goto st283
	tr450:
//line sip.rl:437
		value = &msg.Accept
		goto st283
	tr469:
//line sip.rl:439
		value = &msg.AcceptEncoding
		goto st283
	tr478:
//line sip.rl:440
		value = &msg.AcceptLanguage
		goto st283
	tr489:
//line sip.rl:443
		value = &msg.AlertInfo
		goto st283
	tr493:
//line sip.rl:441
		value = &msg.Allow
		goto st283
	tr502:
//line sip.rl:442
		value = &msg.AllowEvents
		goto st283
	tr522:
//line sip.rl:444
		value = &msg.AuthenticationInfo
		goto st283
	tr532:
//line sip.rl:445
		value = &msg.Authorization
		goto st283
	tr534:
//line sip.rl:462
		value = &msg.ReferredBy
		goto st283
	tr564:
//line sip.rl:449
		value = &msg.CallInfo
		goto st283
	tr597:
//line sip.rl:446
		value = &msg.ContentDisposition
		goto st283
	tr606:
//line sip.rl:448
		value = &msg.ContentEncoding
		goto st283
	tr616:
//line sip.rl:447
		value = &msg.ContentLanguage
		goto st283
	tr651:
//line sip.rl:450
		value = &msg.Date
		goto st283
	tr664:
//line sip.rl:451
		value = &msg.ErrorInfo
		goto st283
	tr669:
//line sip.rl:452
		value = &msg.Event
		goto st283
	tr698:
//line sip.rl:453
		value = &msg.InReplyTo
		goto st283
	tr700:
//line sip.rl:467
		value = &msg.Supported
		goto st283
	tr739:
//line sip.rl:455
		value = &msg.MIMEVersion
		goto st283
	tr767:
//line sip.rl:456
		value = &msg.Organization
		goto st283
	tr797:
//line sip.rl:457
		value = &msg.Priority
		goto st283
	tr816:
//line sip.rl:458
		value = &msg.ProxyAuthenticate
		goto st283
	tr826:
//line sip.rl:459
		value = &msg.ProxyAuthorization
		goto st283
	tr834:
//line sip.rl:460
		value = &msg.ProxyRequire
		goto st283
	tr836:
//line sip.rl:461
		value = &msg.ReferTo
		goto st283
	tr886:
//line sip.rl:454
		value = &msg.ReplyTo
		goto st283
	tr892:
//line sip.rl:463
		value = &msg.Require
		goto st283
	tr902:
//line sip.rl:464
		value = &msg.RetryAfter
		goto st283
	tr909:
//line sip.rl:466
		value = &msg.Subject
		goto st283
	tr917:
//line sip.rl:465
		value = &msg.Server
		goto st283
	tr941:
//line sip.rl:468
		value = &msg.Timestamp
		goto st283
	tr943:
//line sip.rl:441
		value = &msg.Allow
//line sip.rl:442
		value = &msg.AllowEvents
		goto st283
	tr956:
//line sip.rl:469
		value = &msg.Unsupported
		goto st283
	tr966:
//line sip.rl:470
		value = &msg.UserAgent
		goto st283
	tr983:
//line sip.rl:471
		value = &msg.Warning
		goto st283
	tr999:
//line sip.rl:472
		value = &msg.WWWAuthenticate
		goto st283
	st283:
		if p++; p == pe {
			goto _test_eof283
		}
	st_case_283:
//line msg_parse.go:10142
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
//line sip.rl:438
		value = &msg.AcceptContact
		goto st284
	tr452:
//line sip.rl:437
		value = &msg.Accept
		goto st284
	tr470:
//line sip.rl:439
		value = &msg.AcceptEncoding
		goto st284
	tr479:
//line sip.rl:440
		value = &msg.AcceptLanguage
		goto st284
	tr490:
//line sip.rl:443
		value = &msg.AlertInfo
		goto st284
	tr495:
//line sip.rl:441
		value = &msg.Allow
		goto st284
	tr503:
//line sip.rl:442
		value = &msg.AllowEvents
		goto st284
	tr523:
//line sip.rl:444
		value = &msg.AuthenticationInfo
		goto st284
	tr533:
//line sip.rl:445
		value = &msg.Authorization
		goto st284
	tr535:
//line sip.rl:462
		value = &msg.ReferredBy
		goto st284
	tr565:
//line sip.rl:449
		value = &msg.CallInfo
		goto st284
	tr598:
//line sip.rl:446
		value = &msg.ContentDisposition
		goto st284
	tr607:
//line sip.rl:448
		value = &msg.ContentEncoding
		goto st284
	tr617:
//line sip.rl:447
		value = &msg.ContentLanguage
		goto st284
	tr652:
//line sip.rl:450
		value = &msg.Date
		goto st284
	tr665:
//line sip.rl:451
		value = &msg.ErrorInfo
		goto st284
	tr670:
//line sip.rl:452
		value = &msg.Event
		goto st284
	tr699:
//line sip.rl:453
		value = &msg.InReplyTo
		goto st284
	tr701:
//line sip.rl:467
		value = &msg.Supported
		goto st284
	tr740:
//line sip.rl:455
		value = &msg.MIMEVersion
		goto st284
	tr768:
//line sip.rl:456
		value = &msg.Organization
		goto st284
	tr798:
//line sip.rl:457
		value = &msg.Priority
		goto st284
	tr817:
//line sip.rl:458
		value = &msg.ProxyAuthenticate
		goto st284
	tr827:
//line sip.rl:459
		value = &msg.ProxyAuthorization
		goto st284
	tr835:
//line sip.rl:460
		value = &msg.ProxyRequire
		goto st284
	tr837:
//line sip.rl:461
		value = &msg.ReferTo
		goto st284
	tr887:
//line sip.rl:454
		value = &msg.ReplyTo
		goto st284
	tr893:
//line sip.rl:463
		value = &msg.Require
		goto st284
	tr903:
//line sip.rl:464
		value = &msg.RetryAfter
		goto st284
	tr910:
//line sip.rl:466
		value = &msg.Subject
		goto st284
	tr918:
//line sip.rl:465
		value = &msg.Server
		goto st284
	tr942:
//line sip.rl:468
		value = &msg.Timestamp
		goto st284
	tr944:
//line sip.rl:441
		value = &msg.Allow
//line sip.rl:442
		value = &msg.AllowEvents
		goto st284
	tr957:
//line sip.rl:469
		value = &msg.Unsupported
		goto st284
	tr967:
//line sip.rl:470
		value = &msg.UserAgent
		goto st284
	tr984:
//line sip.rl:471
		value = &msg.Warning
		goto st284
	tr1000:
//line sip.rl:472
		value = &msg.WWWAuthenticate
		goto st284
	st284:
		if p++; p == pe {
			goto _test_eof284
		}
	st_case_284:
//line msg_parse.go:10307
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:244
		{
			goto st280
		}
		goto st782
	tr442:
//line sip.rl:59

		p--

//line sip.rl:245
		{
			goto st263
		}
		goto st782
	tr541:
//line sip.rl:59

		p--

//line sip.rl:243
		{
			goto st34
		}
		goto st782
	tr576:
//line sip.rl:521
		value = nil
//line sip.rl:59

		p--

//line sip.rl:239
		{
			goto st256
		}
		goto st782
	tr971:
//line sip.rl:124

		via = new(Via)

//line sip.rl:59

		p--

//line sip.rl:246
		{
			goto st103
		}
		goto st782
	tr1001:
//line sip.rl:63

		{
			p++
			cs = 782
			goto _out
		}

		goto st782
	st782:
		if p++; p == pe {
			goto _test_eof782
		}
	st_case_782:
//line msg_parse.go:10391
		goto st0
	st285:
		if p++; p == pe {
			goto _test_eof285
		}
	st_case_285:
		if data[p] == 10 {
			goto st286
		}
		goto st0
	st286:
		if p++; p == pe {
			goto _test_eof286
		}
	st_case_286:
		switch data[p] {
		case 9:
			goto st287
		case 32:
			goto st287
		}
		goto st0
	st287:
		if p++; p == pe {
			goto _test_eof287
		}
	st_case_287:
		switch data[p] {
		case 9:
			goto st287
		case 32:
			goto st287
		}
		goto tr442
	st288:
		if p++; p == pe {
			goto _test_eof288
		}
	st_case_288:
		switch data[p] {
		case 67:
			goto st289
		case 99:
			goto st289
		}
		goto tr416
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
		goto tr416
	st290:
		if p++; p == pe {
			goto _test_eof290
		}
	st_case_290:
		switch data[p] {
		case 80:
			goto st291
		case 112:
			goto st291
		}
		goto tr416
	st291:
		if p++; p == pe {
			goto _test_eof291
		}
	st_case_291:
		switch data[p] {
		case 84:
			goto st292
		case 116:
			goto st292
		}
		goto tr416
	st292:
		if p++; p == pe {
			goto _test_eof292
		}
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
		if p++; p == pe {
			goto _test_eof293
		}
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
		if p++; p == pe {
			goto _test_eof294
		}
	st_case_294:
		switch data[p] {
		case 79:
			goto st295
		case 111:
			goto st295
		}
		goto tr416
	st295:
		if p++; p == pe {
			goto _test_eof295
		}
	st_case_295:
		switch data[p] {
		case 78:
			goto st296
		case 110:
			goto st296
		}
		goto tr416
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
		goto tr416
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
		goto tr416
	st298:
		if p++; p == pe {
			goto _test_eof298
		}
	st_case_298:
		switch data[p] {
		case 67:
			goto st299
		case 99:
			goto st299
		}
		goto tr416
	st299:
		if p++; p == pe {
			goto _test_eof299
		}
	st_case_299:
		switch data[p] {
		case 84:
			goto st300
		case 116:
			goto st300
		}
		goto tr416
	st300:
		if p++; p == pe {
			goto _test_eof300
		}
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
		goto tr416
	st302:
		if p++; p == pe {
			goto _test_eof302
		}
	st_case_302:
		switch data[p] {
		case 67:
			goto st303
		case 99:
			goto st303
		}
		goto tr416
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
		goto tr416
	st304:
		if p++; p == pe {
			goto _test_eof304
		}
	st_case_304:
		switch data[p] {
		case 68:
			goto st305
		case 100:
			goto st305
		}
		goto tr416
	st305:
		if p++; p == pe {
			goto _test_eof305
		}
	st_case_305:
		switch data[p] {
		case 73:
			goto st306
		case 105:
			goto st306
		}
		goto tr416
	st306:
		if p++; p == pe {
			goto _test_eof306
		}
	st_case_306:
		switch data[p] {
		case 78:
			goto st307
		case 110:
			goto st307
		}
		goto tr416
	st307:
		if p++; p == pe {
			goto _test_eof307
		}
	st_case_307:
		switch data[p] {
		case 71:
			goto st308
		case 103:
			goto st308
		}
		goto tr416
	st308:
		if p++; p == pe {
			goto _test_eof308
		}
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
		if p++; p == pe {
			goto _test_eof309
		}
	st_case_309:
		switch data[p] {
		case 65:
			goto st310
		case 97:
			goto st310
		}
		goto tr416
	st310:
		if p++; p == pe {
			goto _test_eof310
		}
	st_case_310:
		switch data[p] {
		case 78:
			goto st311
		case 110:
			goto st311
		}
		goto tr416
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
		goto tr416
	st312:
		if p++; p == pe {
			goto _test_eof312
		}
	st_case_312:
		switch data[p] {
		case 85:
			goto st313
		case 117:
			goto st313
		}
		goto tr416
	st313:
		if p++; p == pe {
			goto _test_eof313
		}
	st_case_313:
		switch data[p] {
		case 65:
			goto st314
		case 97:
			goto st314
		}
		goto tr416
	st314:
		if p++; p == pe {
			goto _test_eof314
		}
	st_case_314:
		switch data[p] {
		case 71:
			goto st315
		case 103:
			goto st315
		}
		goto tr416
	st315:
		if p++; p == pe {
			goto _test_eof315
		}
	st_case_315:
		switch data[p] {
		case 69:
			goto st316
		case 101:
			goto st316
		}
		goto tr416
	st316:
		if p++; p == pe {
			goto _test_eof316
		}
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
		if p++; p == pe {
			goto _test_eof317
		}
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
		if p++; p == pe {
			goto _test_eof318
		}
	st_case_318:
		switch data[p] {
		case 82:
			goto st319
		case 114:
			goto st319
		}
		goto tr416
	st319:
		if p++; p == pe {
			goto _test_eof319
		}
	st_case_319:
		switch data[p] {
		case 84:
			goto st320
		case 116:
			goto st320
		}
		goto tr416
	st320:
		if p++; p == pe {
			goto _test_eof320
		}
	st_case_320:
		if data[p] == 45 {
			goto st321
		}
		goto tr416
	st321:
		if p++; p == pe {
			goto _test_eof321
		}
	st_case_321:
		switch data[p] {
		case 73:
			goto st322
		case 105:
			goto st322
		}
		goto tr416
	st322:
		if p++; p == pe {
			goto _test_eof322
		}
	st_case_322:
		switch data[p] {
		case 78:
			goto st323
		case 110:
			goto st323
		}
		goto tr416
	st323:
		if p++; p == pe {
			goto _test_eof323
		}
	st_case_323:
		switch data[p] {
		case 70:
			goto st324
		case 102:
			goto st324
		}
		goto tr416
	st324:
		if p++; p == pe {
			goto _test_eof324
		}
	st_case_324:
		switch data[p] {
		case 79:
			goto st325
		case 111:
			goto st325
		}
		goto tr416
	st325:
		if p++; p == pe {
			goto _test_eof325
		}
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
		if p++; p == pe {
			goto _test_eof326
		}
	st_case_326:
		switch data[p] {
		case 79:
			goto st327
		case 111:
			goto st327
		}
		goto tr416
	st327:
		if p++; p == pe {
			goto _test_eof327
		}
	st_case_327:
		switch data[p] {
		case 87:
			goto st328
		case 119:
			goto st328
		}
		goto tr416
	st328:
		if p++; p == pe {
			goto _test_eof328
		}
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
		goto tr416
	st330:
		if p++; p == pe {
			goto _test_eof330
		}
	st_case_330:
		switch data[p] {
		case 86:
			goto st331
		case 118:
			goto st331
		}
		goto tr416
	st331:
		if p++; p == pe {
			goto _test_eof331
		}
	st_case_331:
		switch data[p] {
		case 69:
			goto st332
		case 101:
			goto st332
		}
		goto tr416
	st332:
		if p++; p == pe {
			goto _test_eof332
		}
	st_case_332:
		switch data[p] {
		case 78:
			goto st333
		case 110:
			goto st333
		}
		goto tr416
	st333:
		if p++; p == pe {
			goto _test_eof333
		}
	st_case_333:
		switch data[p] {
		case 84:
			goto st334
		case 116:
			goto st334
		}
		goto tr416
	st334:
		if p++; p == pe {
			goto _test_eof334
		}
	st_case_334:
		switch data[p] {
		case 83:
			goto st335
		case 115:
			goto st335
		}
		goto tr416
	st335:
		if p++; p == pe {
			goto _test_eof335
		}
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
		goto tr416
	st337:
		if p++; p == pe {
			goto _test_eof337
		}
	st_case_337:
		switch data[p] {
		case 72:
			goto st338
		case 104:
			goto st338
		}
		goto tr416
	st338:
		if p++; p == pe {
			goto _test_eof338
		}
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
		if p++; p == pe {
			goto _test_eof339
		}
	st_case_339:
		switch data[p] {
		case 78:
			goto st340
		case 110:
			goto st340
		}
		goto tr416
	st340:
		if p++; p == pe {
			goto _test_eof340
		}
	st_case_340:
		switch data[p] {
		case 84:
			goto st341
		case 116:
			goto st341
		}
		goto tr416
	st341:
		if p++; p == pe {
			goto _test_eof341
		}
	st_case_341:
		switch data[p] {
		case 73:
			goto st342
		case 105:
			goto st342
		}
		goto tr416
	st342:
		if p++; p == pe {
			goto _test_eof342
		}
	st_case_342:
		switch data[p] {
		case 67:
			goto st343
		case 99:
			goto st343
		}
		goto tr416
	st343:
		if p++; p == pe {
			goto _test_eof343
		}
	st_case_343:
		switch data[p] {
		case 65:
			goto st344
		case 97:
			goto st344
		}
		goto tr416
	st344:
		if p++; p == pe {
			goto _test_eof344
		}
	st_case_344:
		switch data[p] {
		case 84:
			goto st345
		case 116:
			goto st345
		}
		goto tr416
	st345:
		if p++; p == pe {
			goto _test_eof345
		}
	st_case_345:
		switch data[p] {
		case 73:
			goto st346
		case 105:
			goto st346
		}
		goto tr416
	st346:
		if p++; p == pe {
			goto _test_eof346
		}
	st_case_346:
		switch data[p] {
		case 79:
			goto st347
		case 111:
			goto st347
		}
		goto tr416
	st347:
		if p++; p == pe {
			goto _test_eof347
		}
	st_case_347:
		switch data[p] {
		case 78:
			goto st348
		case 110:
			goto st348
		}
		goto tr416
	st348:
		if p++; p == pe {
			goto _test_eof348
		}
	st_case_348:
		if data[p] == 45 {
			goto st349
		}
		goto tr416
	st349:
		if p++; p == pe {
			goto _test_eof349
		}
	st_case_349:
		switch data[p] {
		case 73:
			goto st350
		case 105:
			goto st350
		}
		goto tr416
	st350:
		if p++; p == pe {
			goto _test_eof350
		}
	st_case_350:
		switch data[p] {
		case 78:
			goto st351
		case 110:
			goto st351
		}
		goto tr416
	st351:
		if p++; p == pe {
			goto _test_eof351
		}
	st_case_351:
		switch data[p] {
		case 70:
			goto st352
		case 102:
			goto st352
		}
		goto tr416
	st352:
		if p++; p == pe {
			goto _test_eof352
		}
	st_case_352:
		switch data[p] {
		case 79:
			goto st353
		case 111:
			goto st353
		}
		goto tr416
	st353:
		if p++; p == pe {
			goto _test_eof353
		}
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
		if p++; p == pe {
			goto _test_eof354
		}
	st_case_354:
		switch data[p] {
		case 82:
			goto st355
		case 114:
			goto st355
		}
		goto tr416
	st355:
		if p++; p == pe {
			goto _test_eof355
		}
	st_case_355:
		switch data[p] {
		case 73:
			goto st356
		case 105:
			goto st356
		}
		goto tr416
	st356:
		if p++; p == pe {
			goto _test_eof356
		}
	st_case_356:
		switch data[p] {
		case 90:
			goto st357
		case 122:
			goto st357
		}
		goto tr416
	st357:
		if p++; p == pe {
			goto _test_eof357
		}
	st_case_357:
		switch data[p] {
		case 65:
			goto st358
		case 97:
			goto st358
		}
		goto tr416
	st358:
		if p++; p == pe {
			goto _test_eof358
		}
	st_case_358:
		switch data[p] {
		case 84:
			goto st359
		case 116:
			goto st359
		}
		goto tr416
	st359:
		if p++; p == pe {
			goto _test_eof359
		}
	st_case_359:
		switch data[p] {
		case 73:
			goto st360
		case 105:
			goto st360
		}
		goto tr416
	st360:
		if p++; p == pe {
			goto _test_eof360
		}
	st_case_360:
		switch data[p] {
		case 79:
			goto st361
		case 111:
			goto st361
		}
		goto tr416
	st361:
		if p++; p == pe {
			goto _test_eof361
		}
	st_case_361:
		switch data[p] {
		case 78:
			goto st362
		case 110:
			goto st362
		}
		goto tr416
	st362:
		if p++; p == pe {
			goto _test_eof362
		}
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
		if p++; p == pe {
			goto _test_eof363
		}
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
		if p++; p == pe {
			goto _test_eof364
		}
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
		if p++; p == pe {
			goto _test_eof365
		}
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
		if p++; p == pe {
			goto _test_eof366
		}
	st_case_366:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof367
		}
	st_case_367:
		if data[p] == 10 {
			goto st368
		}
		goto st0
	st368:
		if p++; p == pe {
			goto _test_eof368
		}
	st_case_368:
		switch data[p] {
		case 9:
			goto st369
		case 32:
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
			goto st369
		case 32:
			goto st369
		}
		goto tr541
	st370:
		if p++; p == pe {
			goto _test_eof370
		}
	st_case_370:
		switch data[p] {
		case 76:
			goto st371
		case 108:
			goto st371
		}
		goto tr416
	st371:
		if p++; p == pe {
			goto _test_eof371
		}
	st_case_371:
		switch data[p] {
		case 76:
			goto st372
		case 108:
			goto st372
		}
		goto tr416
	st372:
		if p++; p == pe {
			goto _test_eof372
		}
	st_case_372:
		if data[p] == 45 {
			goto st373
		}
		goto tr416
	st373:
		if p++; p == pe {
			goto _test_eof373
		}
	st_case_373:
		switch data[p] {
		case 73:
			goto st374
		case 105:
			goto st374
		}
		goto tr416
	st374:
		if p++; p == pe {
			goto _test_eof374
		}
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
		if p++; p == pe {
			goto _test_eof375
		}
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
		if p++; p == pe {
			goto _test_eof376
		}
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
		if p++; p == pe {
			goto _test_eof377
		}
	st_case_377:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st378
	st378:
		if p++; p == pe {
			goto _test_eof378
		}
	st_case_378:
//line msg_parse.go:11631
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof379
		}
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
		if p++; p == pe {
			goto _test_eof380
		}
	st_case_380:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:207

		msg.CallID = string(data[mark:p])

		goto st381
	tr643:
//line sip.rl:223

		msg.CSeqMethod = string(data[mark:p])

		goto st381
	st381:
		if p++; p == pe {
			goto _test_eof381
		}
	st_case_381:
//line msg_parse.go:11786
		if data[p] == 10 {
			goto tr559
		}
		goto st0
	st382:
		if p++; p == pe {
			goto _test_eof382
		}
	st_case_382:
		if data[p] == 10 {
			goto st383
		}
		goto st0
	st383:
		if p++; p == pe {
			goto _test_eof383
		}
	st_case_383:
		switch data[p] {
		case 9:
			goto st384
		case 32:
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
		if p++; p == pe {
			goto _test_eof385
		}
	st_case_385:
		switch data[p] {
		case 70:
			goto st386
		case 102:
			goto st386
		}
		goto tr416
	st386:
		if p++; p == pe {
			goto _test_eof386
		}
	st_case_386:
		switch data[p] {
		case 79:
			goto st387
		case 111:
			goto st387
		}
		goto tr416
	st387:
		if p++; p == pe {
			goto _test_eof387
		}
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
		if p++; p == pe {
			goto _test_eof388
		}
	st_case_388:
		switch data[p] {
		case 78:
			goto st389
		case 110:
			goto st389
		}
		goto tr416
	st389:
		if p++; p == pe {
			goto _test_eof389
		}
	st_case_389:
		switch data[p] {
		case 84:
			goto st390
		case 116:
			goto st390
		}
		goto tr416
	st390:
		if p++; p == pe {
			goto _test_eof390
		}
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
		if p++; p == pe {
			goto _test_eof391
		}
	st_case_391:
		switch data[p] {
		case 67:
			goto st392
		case 99:
			goto st392
		}
		goto tr416
	st392:
		if p++; p == pe {
			goto _test_eof392
		}
	st_case_392:
		switch data[p] {
		case 84:
			goto st393
		case 116:
			goto st393
		}
		goto tr416
	st393:
		if p++; p == pe {
			goto _test_eof393
		}
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
//line sip.rl:424
		addrp = lastAddr(&msg.Contact)
		goto st394
	tr683:
//line sip.rl:425
		addrp = lastAddr(&msg.From)
		goto st394
	tr788:
//line sip.rl:426
		addrp = lastAddr(&msg.PAssertedIdentity)
		goto st394
	tr855:
//line sip.rl:427
		addrp = lastAddr(&msg.RecordRoute)
		goto st394
	tr879:
//line sip.rl:428
		addrp = lastAddr(&msg.RemotePartyID)
		goto st394
	tr907:
//line sip.rl:429
		addrp = lastAddr(&msg.Route)
		goto st394
	tr930:
//line sip.rl:430
		addrp = lastAddr(&msg.To)
		goto st394
	st394:
		if p++; p == pe {
			goto _test_eof394
		}
	st_case_394:
//line msg_parse.go:12007
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
//line sip.rl:424
		addrp = lastAddr(&msg.Contact)
		goto st395
	tr684:
//line sip.rl:425
		addrp = lastAddr(&msg.From)
		goto st395
	tr789:
//line sip.rl:426
		addrp = lastAddr(&msg.PAssertedIdentity)
		goto st395
	tr856:
//line sip.rl:427
		addrp = lastAddr(&msg.RecordRoute)
		goto st395
	tr880:
//line sip.rl:428
		addrp = lastAddr(&msg.RemotePartyID)
		goto st395
	tr908:
//line sip.rl:429
		addrp = lastAddr(&msg.Route)
		goto st395
	tr931:
//line sip.rl:430
		addrp = lastAddr(&msg.To)
		goto st395
	st395:
		if p++; p == pe {
			goto _test_eof395
		}
	st_case_395:
//line msg_parse.go:12050
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof396
		}
	st_case_396:
		if data[p] == 10 {
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
			goto st398
		case 32:
			goto st398
		}
		goto st0
	st398:
		if p++; p == pe {
			goto _test_eof398
		}
	st_case_398:
		switch data[p] {
		case 9:
			goto st398
		case 32:
			goto st398
		}
		goto tr576
	st399:
		if p++; p == pe {
			goto _test_eof399
		}
	st_case_399:
		switch data[p] {
		case 78:
			goto st400
		case 110:
			goto st400
		}
		goto tr416
	st400:
		if p++; p == pe {
			goto _test_eof400
		}
	st_case_400:
		switch data[p] {
		case 84:
			goto st401
		case 116:
			goto st401
		}
		goto tr416
	st401:
		if p++; p == pe {
			goto _test_eof401
		}
	st_case_401:
		if data[p] == 45 {
			goto st402
		}
		goto tr416
	st402:
		if p++; p == pe {
			goto _test_eof402
		}
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
		if p++; p == pe {
			goto _test_eof403
		}
	st_case_403:
		switch data[p] {
		case 73:
			goto st404
		case 105:
			goto st404
		}
		goto tr416
	st404:
		if p++; p == pe {
			goto _test_eof404
		}
	st_case_404:
		switch data[p] {
		case 83:
			goto st405
		case 115:
			goto st405
		}
		goto tr416
	st405:
		if p++; p == pe {
			goto _test_eof405
		}
	st_case_405:
		switch data[p] {
		case 80:
			goto st406
		case 112:
			goto st406
		}
		goto tr416
	st406:
		if p++; p == pe {
			goto _test_eof406
		}
	st_case_406:
		switch data[p] {
		case 79:
			goto st407
		case 111:
			goto st407
		}
		goto tr416
	st407:
		if p++; p == pe {
			goto _test_eof407
		}
	st_case_407:
		switch data[p] {
		case 83:
			goto st408
		case 115:
			goto st408
		}
		goto tr416
	st408:
		if p++; p == pe {
			goto _test_eof408
		}
	st_case_408:
		switch data[p] {
		case 73:
			goto st409
		case 105:
			goto st409
		}
		goto tr416
	st409:
		if p++; p == pe {
			goto _test_eof409
		}
	st_case_409:
		switch data[p] {
		case 84:
			goto st410
		case 116:
			goto st410
		}
		goto tr416
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
		goto tr416
	st411:
		if p++; p == pe {
			goto _test_eof411
		}
	st_case_411:
		switch data[p] {
		case 79:
			goto st412
		case 111:
			goto st412
		}
		goto tr416
	st412:
		if p++; p == pe {
			goto _test_eof412
		}
	st_case_412:
		switch data[p] {
		case 78:
			goto st413
		case 110:
			goto st413
		}
		goto tr416
	st413:
		if p++; p == pe {
			goto _test_eof413
		}
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
		if p++; p == pe {
			goto _test_eof414
		}
	st_case_414:
		switch data[p] {
		case 78:
			goto st415
		case 110:
			goto st415
		}
		goto tr416
	st415:
		if p++; p == pe {
			goto _test_eof415
		}
	st_case_415:
		switch data[p] {
		case 67:
			goto st416
		case 99:
			goto st416
		}
		goto tr416
	st416:
		if p++; p == pe {
			goto _test_eof416
		}
	st_case_416:
		switch data[p] {
		case 79:
			goto st417
		case 111:
			goto st417
		}
		goto tr416
	st417:
		if p++; p == pe {
			goto _test_eof417
		}
	st_case_417:
		switch data[p] {
		case 68:
			goto st418
		case 100:
			goto st418
		}
		goto tr416
	st418:
		if p++; p == pe {
			goto _test_eof418
		}
	st_case_418:
		switch data[p] {
		case 73:
			goto st419
		case 105:
			goto st419
		}
		goto tr416
	st419:
		if p++; p == pe {
			goto _test_eof419
		}
	st_case_419:
		switch data[p] {
		case 78:
			goto st420
		case 110:
			goto st420
		}
		goto tr416
	st420:
		if p++; p == pe {
			goto _test_eof420
		}
	st_case_420:
		switch data[p] {
		case 71:
			goto st421
		case 103:
			goto st421
		}
		goto tr416
	st421:
		if p++; p == pe {
			goto _test_eof421
		}
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
		if p++; p == pe {
			goto _test_eof422
		}
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
		goto tr416
	st424:
		if p++; p == pe {
			goto _test_eof424
		}
	st_case_424:
		switch data[p] {
		case 71:
			goto st425
		case 103:
			goto st425
		}
		goto tr416
	st425:
		if p++; p == pe {
			goto _test_eof425
		}
	st_case_425:
		switch data[p] {
		case 85:
			goto st426
		case 117:
			goto st426
		}
		goto tr416
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
		goto tr416
	st427:
		if p++; p == pe {
			goto _test_eof427
		}
	st_case_427:
		switch data[p] {
		case 71:
			goto st428
		case 103:
			goto st428
		}
		goto tr416
	st428:
		if p++; p == pe {
			goto _test_eof428
		}
	st_case_428:
		switch data[p] {
		case 69:
			goto st429
		case 101:
			goto st429
		}
		goto tr416
	st429:
		if p++; p == pe {
			goto _test_eof429
		}
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
		goto tr416
	st431:
		if p++; p == pe {
			goto _test_eof431
		}
	st_case_431:
		switch data[p] {
		case 71:
			goto st432
		case 103:
			goto st432
		}
		goto tr416
	st432:
		if p++; p == pe {
			goto _test_eof432
		}
	st_case_432:
		switch data[p] {
		case 84:
			goto st433
		case 116:
			goto st433
		}
		goto tr416
	st433:
		if p++; p == pe {
			goto _test_eof433
		}
	st_case_433:
		switch data[p] {
		case 72:
			goto st434
		case 104:
			goto st434
		}
		goto tr416
	st434:
		if p++; p == pe {
			goto _test_eof434
		}
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
		if p++; p == pe {
			goto _test_eof435
		}
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
		if p++; p == pe {
			goto _test_eof436
		}
	st_case_436:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:480
		clen = 0
//line sip.rl:211

		clen = clen*10 + (int(data[p]) - 0x30)

		goto st437
	tr626:
//line sip.rl:211

		clen = clen*10 + (int(data[p]) - 0x30)

		goto st437
	st437:
		if p++; p == pe {
			goto _test_eof437
		}
	st_case_437:
//line msg_parse.go:12620
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof438
		}
	st_case_438:
		if data[p] == 10 {
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
			goto st440
		case 32:
			goto st440
		}
		goto st0
	st440:
		if p++; p == pe {
			goto _test_eof440
		}
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
		if p++; p == pe {
			goto _test_eof441
		}
	st_case_441:
		switch data[p] {
		case 89:
			goto st442
		case 121:
			goto st442
		}
		goto tr416
	st442:
		if p++; p == pe {
			goto _test_eof442
		}
	st_case_442:
		switch data[p] {
		case 80:
			goto st443
		case 112:
			goto st443
		}
		goto tr416
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
		goto tr416
	st444:
		if p++; p == pe {
			goto _test_eof444
		}
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
		if p++; p == pe {
			goto _test_eof445
		}
	st_case_445:
		switch data[p] {
		case 69:
			goto st446
		case 101:
			goto st446
		}
		goto tr416
	st446:
		if p++; p == pe {
			goto _test_eof446
		}
	st_case_446:
		switch data[p] {
		case 81:
			goto st447
		case 113:
			goto st447
		}
		goto tr416
	st447:
		if p++; p == pe {
			goto _test_eof447
		}
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
		if p++; p == pe {
			goto _test_eof448
		}
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
		if p++; p == pe {
			goto _test_eof449
		}
	st_case_449:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:219

		msg.CSeq = msg.CSeq*10 + (int(data[p]) - 0x30)

		goto st450
	st450:
		if p++; p == pe {
			goto _test_eof450
		}
	st_case_450:
//line msg_parse.go:12808
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof451
		}
	st_case_451:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:67

		mark = p

		goto st452
	st452:
		if p++; p == pe {
			goto _test_eof452
		}
	st_case_452:
//line msg_parse.go:12890
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof453
		}
	st_case_453:
		if data[p] == 10 {
			goto st454
		}
		goto st0
	st454:
		if p++; p == pe {
			goto _test_eof454
		}
	st_case_454:
		switch data[p] {
		case 9:
			goto st455
		case 32:
			goto st455
		}
		goto st0
	st455:
		if p++; p == pe {
			goto _test_eof455
		}
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
		if p++; p == pe {
			goto _test_eof456
		}
	st_case_456:
		if data[p] == 10 {
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
			goto st458
		case 32:
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
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr637
		}
		goto st0
	st459:
		if p++; p == pe {
			goto _test_eof459
		}
	st_case_459:
		switch data[p] {
		case 65:
			goto st460
		case 97:
			goto st460
		}
		goto tr416
	st460:
		if p++; p == pe {
			goto _test_eof460
		}
	st_case_460:
		switch data[p] {
		case 84:
			goto st461
		case 116:
			goto st461
		}
		goto tr416
	st461:
		if p++; p == pe {
			goto _test_eof461
		}
	st_case_461:
		switch data[p] {
		case 69:
			goto st462
		case 101:
			goto st462
		}
		goto tr416
	st462:
		if p++; p == pe {
			goto _test_eof462
		}
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
		if p++; p == pe {
			goto _test_eof463
		}
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
		if p++; p == pe {
			goto _test_eof464
		}
	st_case_464:
		switch data[p] {
		case 82:
			goto st465
		case 114:
			goto st465
		}
		goto tr416
	st465:
		if p++; p == pe {
			goto _test_eof465
		}
	st_case_465:
		switch data[p] {
		case 79:
			goto st466
		case 111:
			goto st466
		}
		goto tr416
	st466:
		if p++; p == pe {
			goto _test_eof466
		}
	st_case_466:
		switch data[p] {
		case 82:
			goto st467
		case 114:
			goto st467
		}
		goto tr416
	st467:
		if p++; p == pe {
			goto _test_eof467
		}
	st_case_467:
		if data[p] == 45 {
			goto st468
		}
		goto tr416
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
		goto tr416
	st469:
		if p++; p == pe {
			goto _test_eof469
		}
	st_case_469:
		switch data[p] {
		case 78:
			goto st470
		case 110:
			goto st470
		}
		goto tr416
	st470:
		if p++; p == pe {
			goto _test_eof470
		}
	st_case_470:
		switch data[p] {
		case 70:
			goto st471
		case 102:
			goto st471
		}
		goto tr416
	st471:
		if p++; p == pe {
			goto _test_eof471
		}
	st_case_471:
		switch data[p] {
		case 79:
			goto st472
		case 111:
			goto st472
		}
		goto tr416
	st472:
		if p++; p == pe {
			goto _test_eof472
		}
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
		if p++; p == pe {
			goto _test_eof473
		}
	st_case_473:
		switch data[p] {
		case 69:
			goto st474
		case 101:
			goto st474
		}
		goto tr416
	st474:
		if p++; p == pe {
			goto _test_eof474
		}
	st_case_474:
		switch data[p] {
		case 78:
			goto st475
		case 110:
			goto st475
		}
		goto tr416
	st475:
		if p++; p == pe {
			goto _test_eof475
		}
	st_case_475:
		switch data[p] {
		case 84:
			goto st476
		case 116:
			goto st476
		}
		goto tr416
	st476:
		if p++; p == pe {
			goto _test_eof476
		}
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
		if p++; p == pe {
			goto _test_eof477
		}
	st_case_477:
		switch data[p] {
		case 80:
			goto st478
		case 112:
			goto st478
		}
		goto tr416
	st478:
		if p++; p == pe {
			goto _test_eof478
		}
	st_case_478:
		switch data[p] {
		case 73:
			goto st479
		case 105:
			goto st479
		}
		goto tr416
	st479:
		if p++; p == pe {
			goto _test_eof479
		}
	st_case_479:
		switch data[p] {
		case 82:
			goto st480
		case 114:
			goto st480
		}
		goto tr416
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
		goto tr416
	st481:
		if p++; p == pe {
			goto _test_eof481
		}
	st_case_481:
		switch data[p] {
		case 83:
			goto st482
		case 115:
			goto st482
		}
		goto tr416
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
		}
		goto tr416
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
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:482
		msg.Expires = 0
//line sip.rl:227

		msg.Expires = msg.Expires*10 + (int(data[p]) - 0x30)

		goto st485
	tr680:
//line sip.rl:227

		msg.Expires = msg.Expires*10 + (int(data[p]) - 0x30)

		goto st485
	st485:
		if p++; p == pe {
			goto _test_eof485
		}
	st_case_485:
//line msg_parse.go:13396
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof486
		}
	st_case_486:
		if data[p] == 10 {
			goto st487
		}
		goto st0
	st487:
		if p++; p == pe {
			goto _test_eof487
		}
	st_case_487:
		switch data[p] {
		case 9:
			goto st488
		case 32:
			goto st488
		}
		goto st0
	st488:
		if p++; p == pe {
			goto _test_eof488
		}
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
		if p++; p == pe {
			goto _test_eof489
		}
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
		if p++; p == pe {
			goto _test_eof490
		}
	st_case_490:
		switch data[p] {
		case 79:
			goto st491
		case 111:
			goto st491
		}
		goto tr416
	st491:
		if p++; p == pe {
			goto _test_eof491
		}
	st_case_491:
		switch data[p] {
		case 77:
			goto st492
		case 109:
			goto st492
		}
		goto tr416
	st492:
		if p++; p == pe {
			goto _test_eof492
		}
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
		if p++; p == pe {
			goto _test_eof493
		}
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
		if p++; p == pe {
			goto _test_eof494
		}
	st_case_494:
		if data[p] == 45 {
			goto st495
		}
		goto tr416
	st495:
		if p++; p == pe {
			goto _test_eof495
		}
	st_case_495:
		switch data[p] {
		case 82:
			goto st496
		case 114:
			goto st496
		}
		goto tr416
	st496:
		if p++; p == pe {
			goto _test_eof496
		}
	st_case_496:
		switch data[p] {
		case 69:
			goto st497
		case 101:
			goto st497
		}
		goto tr416
	st497:
		if p++; p == pe {
			goto _test_eof497
		}
	st_case_497:
		switch data[p] {
		case 80:
			goto st498
		case 112:
			goto st498
		}
		goto tr416
	st498:
		if p++; p == pe {
			goto _test_eof498
		}
	st_case_498:
		switch data[p] {
		case 76:
			goto st499
		case 108:
			goto st499
		}
		goto tr416
	st499:
		if p++; p == pe {
			goto _test_eof499
		}
	st_case_499:
		switch data[p] {
		case 89:
			goto st500
		case 121:
			goto st500
		}
		goto tr416
	st500:
		if p++; p == pe {
			goto _test_eof500
		}
	st_case_500:
		if data[p] == 45 {
			goto st501
		}
		goto tr416
	st501:
		if p++; p == pe {
			goto _test_eof501
		}
	st_case_501:
		switch data[p] {
		case 84:
			goto st502
		case 116:
			goto st502
		}
		goto tr416
	st502:
		if p++; p == pe {
			goto _test_eof502
		}
	st_case_502:
		switch data[p] {
		case 79:
			goto st503
		case 111:
			goto st503
		}
		goto tr416
	st503:
		if p++; p == pe {
			goto _test_eof503
		}
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
		if p++; p == pe {
			goto _test_eof504
		}
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
		if p++; p == pe {
			goto _test_eof505
		}
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
		if p++; p == pe {
			goto _test_eof506
		}
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
		if p++; p == pe {
			goto _test_eof507
		}
	st_case_507:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:480
		clen = 0
//line sip.rl:211

		clen = clen*10 + (int(data[p]) - 0x30)

//line sip.rl:482
		msg.Expires = 0
//line sip.rl:227

		msg.Expires = msg.Expires*10 + (int(data[p]) - 0x30)

//line sip.rl:483
		msg.MaxForwards = 0
//line sip.rl:231

		msg.MaxForwards = msg.MaxForwards*10 + (int(data[p]) - 0x30)

//line sip.rl:484
		msg.MinExpires = 0
//line sip.rl:235

		msg.MinExpires = msg.MinExpires*10 + (int(data[p]) - 0x30)

		goto st508
	tr706:
//line sip.rl:211

		clen = clen*10 + (int(data[p]) - 0x30)

//line sip.rl:227

		msg.Expires = msg.Expires*10 + (int(data[p]) - 0x30)

//line sip.rl:231

		msg.MaxForwards = msg.MaxForwards*10 + (int(data[p]) - 0x30)

//line sip.rl:235

		msg.MinExpires = msg.MinExpires*10 + (int(data[p]) - 0x30)

		goto st508
	st508:
		if p++; p == pe {
			goto _test_eof508
		}
	st_case_508:
//line msg_parse.go:13752
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof509
		}
	st_case_509:
		if data[p] == 10 {
			goto st510
		}
		goto st0
	st510:
		if p++; p == pe {
			goto _test_eof510
		}
	st_case_510:
		switch data[p] {
		case 9:
			goto st511
		case 32:
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
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr704
		}
		goto st0
	st512:
		if p++; p == pe {
			goto _test_eof512
		}
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
		if p++; p == pe {
			goto _test_eof513
		}
	st_case_513:
		switch data[p] {
		case 88:
			goto st514
		case 120:
			goto st514
		}
		goto tr416
	st514:
		if p++; p == pe {
			goto _test_eof514
		}
	st_case_514:
		if data[p] == 45 {
			goto st515
		}
		goto tr416
	st515:
		if p++; p == pe {
			goto _test_eof515
		}
	st_case_515:
		switch data[p] {
		case 70:
			goto st516
		case 102:
			goto st516
		}
		goto tr416
	st516:
		if p++; p == pe {
			goto _test_eof516
		}
	st_case_516:
		switch data[p] {
		case 79:
			goto st517
		case 111:
			goto st517
		}
		goto tr416
	st517:
		if p++; p == pe {
			goto _test_eof517
		}
	st_case_517:
		switch data[p] {
		case 82:
			goto st518
		case 114:
			goto st518
		}
		goto tr416
	st518:
		if p++; p == pe {
			goto _test_eof518
		}
	st_case_518:
		switch data[p] {
		case 87:
			goto st519
		case 119:
			goto st519
		}
		goto tr416
	st519:
		if p++; p == pe {
			goto _test_eof519
		}
	st_case_519:
		switch data[p] {
		case 65:
			goto st520
		case 97:
			goto st520
		}
		goto tr416
	st520:
		if p++; p == pe {
			goto _test_eof520
		}
	st_case_520:
		switch data[p] {
		case 82:
			goto st521
		case 114:
			goto st521
		}
		goto tr416
	st521:
		if p++; p == pe {
			goto _test_eof521
		}
	st_case_521:
		switch data[p] {
		case 68:
			goto st522
		case 100:
			goto st522
		}
		goto tr416
	st522:
		if p++; p == pe {
			goto _test_eof522
		}
	st_case_522:
		switch data[p] {
		case 83:
			goto st523
		case 115:
			goto st523
		}
		goto tr416
	st523:
		if p++; p == pe {
			goto _test_eof523
		}
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
		if p++; p == pe {
			goto _test_eof524
		}
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
		if p++; p == pe {
			goto _test_eof525
		}
	st_case_525:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:483
		msg.MaxForwards = 0
//line sip.rl:231

		msg.MaxForwards = msg.MaxForwards*10 + (int(data[p]) - 0x30)

		goto st526
	tr725:
//line sip.rl:231

		msg.MaxForwards = msg.MaxForwards*10 + (int(data[p]) - 0x30)

		goto st526
	st526:
		if p++; p == pe {
			goto _test_eof526
		}
	st_case_526:
//line msg_parse.go:14013
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof527
		}
	st_case_527:
		if data[p] == 10 {
			goto st528
		}
		goto st0
	st528:
		if p++; p == pe {
			goto _test_eof528
		}
	st_case_528:
		switch data[p] {
		case 9:
			goto st529
		case 32:
			goto st529
		}
		goto st0
	st529:
		if p++; p == pe {
			goto _test_eof529
		}
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
		if p++; p == pe {
			goto _test_eof530
		}
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
		if p++; p == pe {
			goto _test_eof531
		}
	st_case_531:
		switch data[p] {
		case 69:
			goto st532
		case 101:
			goto st532
		}
		goto tr416
	st532:
		if p++; p == pe {
			goto _test_eof532
		}
	st_case_532:
		if data[p] == 45 {
			goto st533
		}
		goto tr416
	st533:
		if p++; p == pe {
			goto _test_eof533
		}
	st_case_533:
		switch data[p] {
		case 86:
			goto st534
		case 118:
			goto st534
		}
		goto tr416
	st534:
		if p++; p == pe {
			goto _test_eof534
		}
	st_case_534:
		switch data[p] {
		case 69:
			goto st535
		case 101:
			goto st535
		}
		goto tr416
	st535:
		if p++; p == pe {
			goto _test_eof535
		}
	st_case_535:
		switch data[p] {
		case 82:
			goto st536
		case 114:
			goto st536
		}
		goto tr416
	st536:
		if p++; p == pe {
			goto _test_eof536
		}
	st_case_536:
		switch data[p] {
		case 83:
			goto st537
		case 115:
			goto st537
		}
		goto tr416
	st537:
		if p++; p == pe {
			goto _test_eof537
		}
	st_case_537:
		switch data[p] {
		case 73:
			goto st538
		case 105:
			goto st538
		}
		goto tr416
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
		goto tr416
	st539:
		if p++; p == pe {
			goto _test_eof539
		}
	st_case_539:
		switch data[p] {
		case 78:
			goto st540
		case 110:
			goto st540
		}
		goto tr416
	st540:
		if p++; p == pe {
			goto _test_eof540
		}
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
		if p++; p == pe {
			goto _test_eof541
		}
	st_case_541:
		if data[p] == 45 {
			goto st542
		}
		goto tr416
	st542:
		if p++; p == pe {
			goto _test_eof542
		}
	st_case_542:
		switch data[p] {
		case 69:
			goto st543
		case 101:
			goto st543
		}
		goto tr416
	st543:
		if p++; p == pe {
			goto _test_eof543
		}
	st_case_543:
		switch data[p] {
		case 88:
			goto st544
		case 120:
			goto st544
		}
		goto tr416
	st544:
		if p++; p == pe {
			goto _test_eof544
		}
	st_case_544:
		switch data[p] {
		case 80:
			goto st545
		case 112:
			goto st545
		}
		goto tr416
	st545:
		if p++; p == pe {
			goto _test_eof545
		}
	st_case_545:
		switch data[p] {
		case 73:
			goto st546
		case 105:
			goto st546
		}
		goto tr416
	st546:
		if p++; p == pe {
			goto _test_eof546
		}
	st_case_546:
		switch data[p] {
		case 82:
			goto st547
		case 114:
			goto st547
		}
		goto tr416
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
		goto tr416
	st548:
		if p++; p == pe {
			goto _test_eof548
		}
	st_case_548:
		switch data[p] {
		case 83:
			goto st549
		case 115:
			goto st549
		}
		goto tr416
	st549:
		if p++; p == pe {
			goto _test_eof549
		}
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
		if p++; p == pe {
			goto _test_eof550
		}
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
		if p++; p == pe {
			goto _test_eof551
		}
	st_case_551:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
//line sip.rl:484
		msg.MinExpires = 0
//line sip.rl:235

		msg.MinExpires = msg.MinExpires*10 + (int(data[p]) - 0x30)

		goto st552
	tr753:
//line sip.rl:235

		msg.MinExpires = msg.MinExpires*10 + (int(data[p]) - 0x30)

		goto st552
	st552:
		if p++; p == pe {
			goto _test_eof552
		}
	st_case_552:
//line msg_parse.go:14363
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof553
		}
	st_case_553:
		if data[p] == 10 {
			goto st554
		}
		goto st0
	st554:
		if p++; p == pe {
			goto _test_eof554
		}
	st_case_554:
		switch data[p] {
		case 9:
			goto st555
		case 32:
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
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr751
		}
		goto st0
	st556:
		if p++; p == pe {
			goto _test_eof556
		}
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
		if p++; p == pe {
			goto _test_eof557
		}
	st_case_557:
		switch data[p] {
		case 71:
			goto st558
		case 103:
			goto st558
		}
		goto tr416
	st558:
		if p++; p == pe {
			goto _test_eof558
		}
	st_case_558:
		switch data[p] {
		case 65:
			goto st559
		case 97:
			goto st559
		}
		goto tr416
	st559:
		if p++; p == pe {
			goto _test_eof559
		}
	st_case_559:
		switch data[p] {
		case 78:
			goto st560
		case 110:
			goto st560
		}
		goto tr416
	st560:
		if p++; p == pe {
			goto _test_eof560
		}
	st_case_560:
		switch data[p] {
		case 73:
			goto st561
		case 105:
			goto st561
		}
		goto tr416
	st561:
		if p++; p == pe {
			goto _test_eof561
		}
	st_case_561:
		switch data[p] {
		case 90:
			goto st562
		case 122:
			goto st562
		}
		goto tr416
	st562:
		if p++; p == pe {
			goto _test_eof562
		}
	st_case_562:
		switch data[p] {
		case 65:
			goto st563
		case 97:
			goto st563
		}
		goto tr416
	st563:
		if p++; p == pe {
			goto _test_eof563
		}
	st_case_563:
		switch data[p] {
		case 84:
			goto st564
		case 116:
			goto st564
		}
		goto tr416
	st564:
		if p++; p == pe {
			goto _test_eof564
		}
	st_case_564:
		switch data[p] {
		case 73:
			goto st565
		case 105:
			goto st565
		}
		goto tr416
	st565:
		if p++; p == pe {
			goto _test_eof565
		}
	st_case_565:
		switch data[p] {
		case 79:
			goto st566
		case 111:
			goto st566
		}
		goto tr416
	st566:
		if p++; p == pe {
			goto _test_eof566
		}
	st_case_566:
		switch data[p] {
		case 78:
			goto st567
		case 110:
			goto st567
		}
		goto tr416
	st567:
		if p++; p == pe {
			goto _test_eof567
		}
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
		if p++; p == pe {
			goto _test_eof568
		}
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
		if p++; p == pe {
			goto _test_eof569
		}
	st_case_569:
		switch data[p] {
		case 65:
			goto st570
		case 97:
			goto st570
		}
		goto tr416
	st570:
		if p++; p == pe {
			goto _test_eof570
		}
	st_case_570:
		switch data[p] {
		case 83:
			goto st571
		case 115:
			goto st571
		}
		goto tr416
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
		goto tr416
	st572:
		if p++; p == pe {
			goto _test_eof572
		}
	st_case_572:
		switch data[p] {
		case 69:
			goto st573
		case 101:
			goto st573
		}
		goto tr416
	st573:
		if p++; p == pe {
			goto _test_eof573
		}
	st_case_573:
		switch data[p] {
		case 82:
			goto st574
		case 114:
			goto st574
		}
		goto tr416
	st574:
		if p++; p == pe {
			goto _test_eof574
		}
	st_case_574:
		switch data[p] {
		case 84:
			goto st575
		case 116:
			goto st575
		}
		goto tr416
	st575:
		if p++; p == pe {
			goto _test_eof575
		}
	st_case_575:
		switch data[p] {
		case 69:
			goto st576
		case 101:
			goto st576
		}
		goto tr416
	st576:
		if p++; p == pe {
			goto _test_eof576
		}
	st_case_576:
		switch data[p] {
		case 68:
			goto st577
		case 100:
			goto st577
		}
		goto tr416
	st577:
		if p++; p == pe {
			goto _test_eof577
		}
	st_case_577:
		if data[p] == 45 {
			goto st578
		}
		goto tr416
	st578:
		if p++; p == pe {
			goto _test_eof578
		}
	st_case_578:
		switch data[p] {
		case 73:
			goto st579
		case 105:
			goto st579
		}
		goto tr416
	st579:
		if p++; p == pe {
			goto _test_eof579
		}
	st_case_579:
		switch data[p] {
		case 68:
			goto st580
		case 100:
			goto st580
		}
		goto tr416
	st580:
		if p++; p == pe {
			goto _test_eof580
		}
	st_case_580:
		switch data[p] {
		case 69:
			goto st581
		case 101:
			goto st581
		}
		goto tr416
	st581:
		if p++; p == pe {
			goto _test_eof581
		}
	st_case_581:
		switch data[p] {
		case 78:
			goto st582
		case 110:
			goto st582
		}
		goto tr416
	st582:
		if p++; p == pe {
			goto _test_eof582
		}
	st_case_582:
		switch data[p] {
		case 84:
			goto st583
		case 116:
			goto st583
		}
		goto tr416
	st583:
		if p++; p == pe {
			goto _test_eof583
		}
	st_case_583:
		switch data[p] {
		case 73:
			goto st584
		case 105:
			goto st584
		}
		goto tr416
	st584:
		if p++; p == pe {
			goto _test_eof584
		}
	st_case_584:
		switch data[p] {
		case 84:
			goto st585
		case 116:
			goto st585
		}
		goto tr416
	st585:
		if p++; p == pe {
			goto _test_eof585
		}
	st_case_585:
		switch data[p] {
		case 89:
			goto st586
		case 121:
			goto st586
		}
		goto tr416
	st586:
		if p++; p == pe {
			goto _test_eof586
		}
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
		if p++; p == pe {
			goto _test_eof587
		}
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
		if p++; p == pe {
			goto _test_eof588
		}
	st_case_588:
		switch data[p] {
		case 79:
			goto st589
		case 111:
			goto st589
		}
		goto tr416
	st589:
		if p++; p == pe {
			goto _test_eof589
		}
	st_case_589:
		switch data[p] {
		case 82:
			goto st590
		case 114:
			goto st590
		}
		goto tr416
	st590:
		if p++; p == pe {
			goto _test_eof590
		}
	st_case_590:
		switch data[p] {
		case 73:
			goto st591
		case 105:
			goto st591
		}
		goto tr416
	st591:
		if p++; p == pe {
			goto _test_eof591
		}
	st_case_591:
		switch data[p] {
		case 84:
			goto st592
		case 116:
			goto st592
		}
		goto tr416
	st592:
		if p++; p == pe {
			goto _test_eof592
		}
	st_case_592:
		switch data[p] {
		case 89:
			goto st593
		case 121:
			goto st593
		}
		goto tr416
	st593:
		if p++; p == pe {
			goto _test_eof593
		}
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
		if p++; p == pe {
			goto _test_eof594
		}
	st_case_594:
		switch data[p] {
		case 88:
			goto st595
		case 120:
			goto st595
		}
		goto tr416
	st595:
		if p++; p == pe {
			goto _test_eof595
		}
	st_case_595:
		switch data[p] {
		case 89:
			goto st596
		case 121:
			goto st596
		}
		goto tr416
	st596:
		if p++; p == pe {
			goto _test_eof596
		}
	st_case_596:
		if data[p] == 45 {
			goto st597
		}
		goto tr416
	st597:
		if p++; p == pe {
			goto _test_eof597
		}
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
		if p++; p == pe {
			goto _test_eof598
		}
	st_case_598:
		switch data[p] {
		case 85:
			goto st599
		case 117:
			goto st599
		}
		goto tr416
	st599:
		if p++; p == pe {
			goto _test_eof599
		}
	st_case_599:
		switch data[p] {
		case 84:
			goto st600
		case 116:
			goto st600
		}
		goto tr416
	st600:
		if p++; p == pe {
			goto _test_eof600
		}
	st_case_600:
		switch data[p] {
		case 72:
			goto st601
		case 104:
			goto st601
		}
		goto tr416
	st601:
		if p++; p == pe {
			goto _test_eof601
		}
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
		if p++; p == pe {
			goto _test_eof602
		}
	st_case_602:
		switch data[p] {
		case 78:
			goto st603
		case 110:
			goto st603
		}
		goto tr416
	st603:
		if p++; p == pe {
			goto _test_eof603
		}
	st_case_603:
		switch data[p] {
		case 84:
			goto st604
		case 116:
			goto st604
		}
		goto tr416
	st604:
		if p++; p == pe {
			goto _test_eof604
		}
	st_case_604:
		switch data[p] {
		case 73:
			goto st605
		case 105:
			goto st605
		}
		goto tr416
	st605:
		if p++; p == pe {
			goto _test_eof605
		}
	st_case_605:
		switch data[p] {
		case 67:
			goto st606
		case 99:
			goto st606
		}
		goto tr416
	st606:
		if p++; p == pe {
			goto _test_eof606
		}
	st_case_606:
		switch data[p] {
		case 65:
			goto st607
		case 97:
			goto st607
		}
		goto tr416
	st607:
		if p++; p == pe {
			goto _test_eof607
		}
	st_case_607:
		switch data[p] {
		case 84:
			goto st608
		case 116:
			goto st608
		}
		goto tr416
	st608:
		if p++; p == pe {
			goto _test_eof608
		}
	st_case_608:
		switch data[p] {
		case 69:
			goto st609
		case 101:
			goto st609
		}
		goto tr416
	st609:
		if p++; p == pe {
			goto _test_eof609
		}
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
		if p++; p == pe {
			goto _test_eof610
		}
	st_case_610:
		switch data[p] {
		case 82:
			goto st611
		case 114:
			goto st611
		}
		goto tr416
	st611:
		if p++; p == pe {
			goto _test_eof611
		}
	st_case_611:
		switch data[p] {
		case 73:
			goto st612
		case 105:
			goto st612
		}
		goto tr416
	st612:
		if p++; p == pe {
			goto _test_eof612
		}
	st_case_612:
		switch data[p] {
		case 90:
			goto st613
		case 122:
			goto st613
		}
		goto tr416
	st613:
		if p++; p == pe {
			goto _test_eof613
		}
	st_case_613:
		switch data[p] {
		case 65:
			goto st614
		case 97:
			goto st614
		}
		goto tr416
	st614:
		if p++; p == pe {
			goto _test_eof614
		}
	st_case_614:
		switch data[p] {
		case 84:
			goto st615
		case 116:
			goto st615
		}
		goto tr416
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
		goto tr416
	st616:
		if p++; p == pe {
			goto _test_eof616
		}
	st_case_616:
		switch data[p] {
		case 79:
			goto st617
		case 111:
			goto st617
		}
		goto tr416
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
		goto tr416
	st618:
		if p++; p == pe {
			goto _test_eof618
		}
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
		if p++; p == pe {
			goto _test_eof619
		}
	st_case_619:
		switch data[p] {
		case 69:
			goto st620
		case 101:
			goto st620
		}
		goto tr416
	st620:
		if p++; p == pe {
			goto _test_eof620
		}
	st_case_620:
		switch data[p] {
		case 81:
			goto st621
		case 113:
			goto st621
		}
		goto tr416
	st621:
		if p++; p == pe {
			goto _test_eof621
		}
	st_case_621:
		switch data[p] {
		case 85:
			goto st622
		case 117:
			goto st622
		}
		goto tr416
	st622:
		if p++; p == pe {
			goto _test_eof622
		}
	st_case_622:
		switch data[p] {
		case 73:
			goto st623
		case 105:
			goto st623
		}
		goto tr416
	st623:
		if p++; p == pe {
			goto _test_eof623
		}
	st_case_623:
		switch data[p] {
		case 82:
			goto st624
		case 114:
			goto st624
		}
		goto tr416
	st624:
		if p++; p == pe {
			goto _test_eof624
		}
	st_case_624:
		switch data[p] {
		case 69:
			goto st625
		case 101:
			goto st625
		}
		goto tr416
	st625:
		if p++; p == pe {
			goto _test_eof625
		}
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
		if p++; p == pe {
			goto _test_eof626
		}
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
		if p++; p == pe {
			goto _test_eof627
		}
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
		if p++; p == pe {
			goto _test_eof628
		}
	st_case_628:
		switch data[p] {
		case 79:
			goto st629
		case 111:
			goto st629
		}
		goto tr416
	st629:
		if p++; p == pe {
			goto _test_eof629
		}
	st_case_629:
		switch data[p] {
		case 82:
			goto st630
		case 114:
			goto st630
		}
		goto tr416
	st630:
		if p++; p == pe {
			goto _test_eof630
		}
	st_case_630:
		switch data[p] {
		case 68:
			goto st631
		case 100:
			goto st631
		}
		goto tr416
	st631:
		if p++; p == pe {
			goto _test_eof631
		}
	st_case_631:
		if data[p] == 45 {
			goto st632
		}
		goto tr416
	st632:
		if p++; p == pe {
			goto _test_eof632
		}
	st_case_632:
		switch data[p] {
		case 82:
			goto st633
		case 114:
			goto st633
		}
		goto tr416
	st633:
		if p++; p == pe {
			goto _test_eof633
		}
	st_case_633:
		switch data[p] {
		case 79:
			goto st634
		case 111:
			goto st634
		}
		goto tr416
	st634:
		if p++; p == pe {
			goto _test_eof634
		}
	st_case_634:
		switch data[p] {
		case 85:
			goto st635
		case 117:
			goto st635
		}
		goto tr416
	st635:
		if p++; p == pe {
			goto _test_eof635
		}
	st_case_635:
		switch data[p] {
		case 84:
			goto st636
		case 116:
			goto st636
		}
		goto tr416
	st636:
		if p++; p == pe {
			goto _test_eof636
		}
	st_case_636:
		switch data[p] {
		case 69:
			goto st637
		case 101:
			goto st637
		}
		goto tr416
	st637:
		if p++; p == pe {
			goto _test_eof637
		}
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
		if p++; p == pe {
			goto _test_eof638
		}
	st_case_638:
		switch data[p] {
		case 69:
			goto st639
		case 101:
			goto st639
		}
		goto tr416
	st639:
		if p++; p == pe {
			goto _test_eof639
		}
	st_case_639:
		switch data[p] {
		case 82:
			goto st640
		case 114:
			goto st640
		}
		goto tr416
	st640:
		if p++; p == pe {
			goto _test_eof640
		}
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
		if p++; p == pe {
			goto _test_eof641
		}
	st_case_641:
		switch data[p] {
		case 84:
			goto st642
		case 116:
			goto st642
		}
		goto tr416
	st642:
		if p++; p == pe {
			goto _test_eof642
		}
	st_case_642:
		switch data[p] {
		case 79:
			goto st643
		case 111:
			goto st643
		}
		goto tr416
	st643:
		if p++; p == pe {
			goto _test_eof643
		}
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
		goto tr416
	st645:
		if p++; p == pe {
			goto _test_eof645
		}
	st_case_645:
		switch data[p] {
		case 68:
			goto st646
		case 100:
			goto st646
		}
		goto tr416
	st646:
		if p++; p == pe {
			goto _test_eof646
		}
	st_case_646:
		if data[p] == 45 {
			goto st647
		}
		goto tr416
	st647:
		if p++; p == pe {
			goto _test_eof647
		}
	st_case_647:
		switch data[p] {
		case 66:
			goto st648
		case 98:
			goto st648
		}
		goto tr416
	st648:
		if p++; p == pe {
			goto _test_eof648
		}
	st_case_648:
		switch data[p] {
		case 89:
			goto st363
		case 121:
			goto st363
		}
		goto tr416
	st649:
		if p++; p == pe {
			goto _test_eof649
		}
	st_case_649:
		switch data[p] {
		case 79:
			goto st650
		case 111:
			goto st650
		}
		goto tr416
	st650:
		if p++; p == pe {
			goto _test_eof650
		}
	st_case_650:
		switch data[p] {
		case 84:
			goto st651
		case 116:
			goto st651
		}
		goto tr416
	st651:
		if p++; p == pe {
			goto _test_eof651
		}
	st_case_651:
		switch data[p] {
		case 69:
			goto st652
		case 101:
			goto st652
		}
		goto tr416
	st652:
		if p++; p == pe {
			goto _test_eof652
		}
	st_case_652:
		if data[p] == 45 {
			goto st653
		}
		goto tr416
	st653:
		if p++; p == pe {
			goto _test_eof653
		}
	st_case_653:
		switch data[p] {
		case 80:
			goto st654
		case 112:
			goto st654
		}
		goto tr416
	st654:
		if p++; p == pe {
			goto _test_eof654
		}
	st_case_654:
		switch data[p] {
		case 65:
			goto st655
		case 97:
			goto st655
		}
		goto tr416
	st655:
		if p++; p == pe {
			goto _test_eof655
		}
	st_case_655:
		switch data[p] {
		case 82:
			goto st656
		case 114:
			goto st656
		}
		goto tr416
	st656:
		if p++; p == pe {
			goto _test_eof656
		}
	st_case_656:
		switch data[p] {
		case 84:
			goto st657
		case 116:
			goto st657
		}
		goto tr416
	st657:
		if p++; p == pe {
			goto _test_eof657
		}
	st_case_657:
		switch data[p] {
		case 89:
			goto st658
		case 121:
			goto st658
		}
		goto tr416
	st658:
		if p++; p == pe {
			goto _test_eof658
		}
	st_case_658:
		if data[p] == 45 {
			goto st659
		}
		goto tr416
	st659:
		if p++; p == pe {
			goto _test_eof659
		}
	st_case_659:
		switch data[p] {
		case 73:
			goto st660
		case 105:
			goto st660
		}
		goto tr416
	st660:
		if p++; p == pe {
			goto _test_eof660
		}
	st_case_660:
		switch data[p] {
		case 68:
			goto st661
		case 100:
			goto st661
		}
		goto tr416
	st661:
		if p++; p == pe {
			goto _test_eof661
		}
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
		if p++; p == pe {
			goto _test_eof662
		}
	st_case_662:
		switch data[p] {
		case 76:
			goto st663
		case 108:
			goto st663
		}
		goto tr416
	st663:
		if p++; p == pe {
			goto _test_eof663
		}
	st_case_663:
		switch data[p] {
		case 89:
			goto st664
		case 121:
			goto st664
		}
		goto tr416
	st664:
		if p++; p == pe {
			goto _test_eof664
		}
	st_case_664:
		if data[p] == 45 {
			goto st665
		}
		goto tr416
	st665:
		if p++; p == pe {
			goto _test_eof665
		}
	st_case_665:
		switch data[p] {
		case 84:
			goto st666
		case 116:
			goto st666
		}
		goto tr416
	st666:
		if p++; p == pe {
			goto _test_eof666
		}
	st_case_666:
		switch data[p] {
		case 79:
			goto st667
		case 111:
			goto st667
		}
		goto tr416
	st667:
		if p++; p == pe {
			goto _test_eof667
		}
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
		if p++; p == pe {
			goto _test_eof668
		}
	st_case_668:
		switch data[p] {
		case 85:
			goto st669
		case 117:
			goto st669
		}
		goto tr416
	st669:
		if p++; p == pe {
			goto _test_eof669
		}
	st_case_669:
		switch data[p] {
		case 73:
			goto st670
		case 105:
			goto st670
		}
		goto tr416
	st670:
		if p++; p == pe {
			goto _test_eof670
		}
	st_case_670:
		switch data[p] {
		case 82:
			goto st671
		case 114:
			goto st671
		}
		goto tr416
	st671:
		if p++; p == pe {
			goto _test_eof671
		}
	st_case_671:
		switch data[p] {
		case 69:
			goto st672
		case 101:
			goto st672
		}
		goto tr416
	st672:
		if p++; p == pe {
			goto _test_eof672
		}
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
		if p++; p == pe {
			goto _test_eof673
		}
	st_case_673:
		switch data[p] {
		case 82:
			goto st674
		case 114:
			goto st674
		}
		goto tr416
	st674:
		if p++; p == pe {
			goto _test_eof674
		}
	st_case_674:
		switch data[p] {
		case 89:
			goto st675
		case 121:
			goto st675
		}
		goto tr416
	st675:
		if p++; p == pe {
			goto _test_eof675
		}
	st_case_675:
		if data[p] == 45 {
			goto st676
		}
		goto tr416
	st676:
		if p++; p == pe {
			goto _test_eof676
		}
	st_case_676:
		switch data[p] {
		case 65:
			goto st677
		case 97:
			goto st677
		}
		goto tr416
	st677:
		if p++; p == pe {
			goto _test_eof677
		}
	st_case_677:
		switch data[p] {
		case 70:
			goto st678
		case 102:
			goto st678
		}
		goto tr416
	st678:
		if p++; p == pe {
			goto _test_eof678
		}
	st_case_678:
		switch data[p] {
		case 84:
			goto st679
		case 116:
			goto st679
		}
		goto tr416
	st679:
		if p++; p == pe {
			goto _test_eof679
		}
	st_case_679:
		switch data[p] {
		case 69:
			goto st680
		case 101:
			goto st680
		}
		goto tr416
	st680:
		if p++; p == pe {
			goto _test_eof680
		}
	st_case_680:
		switch data[p] {
		case 82:
			goto st681
		case 114:
			goto st681
		}
		goto tr416
	st681:
		if p++; p == pe {
			goto _test_eof681
		}
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
		if p++; p == pe {
			goto _test_eof682
		}
	st_case_682:
		switch data[p] {
		case 85:
			goto st683
		case 117:
			goto st683
		}
		goto tr416
	st683:
		if p++; p == pe {
			goto _test_eof683
		}
	st_case_683:
		switch data[p] {
		case 84:
			goto st684
		case 116:
			goto st684
		}
		goto tr416
	st684:
		if p++; p == pe {
			goto _test_eof684
		}
	st_case_684:
		switch data[p] {
		case 69:
			goto st685
		case 101:
			goto st685
		}
		goto tr416
	st685:
		if p++; p == pe {
			goto _test_eof685
		}
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
		if p++; p == pe {
			goto _test_eof686
		}
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
		if p++; p == pe {
			goto _test_eof687
		}
	st_case_687:
		switch data[p] {
		case 82:
			goto st688
		case 114:
			goto st688
		}
		goto tr416
	st688:
		if p++; p == pe {
			goto _test_eof688
		}
	st_case_688:
		switch data[p] {
		case 86:
			goto st689
		case 118:
			goto st689
		}
		goto tr416
	st689:
		if p++; p == pe {
			goto _test_eof689
		}
	st_case_689:
		switch data[p] {
		case 69:
			goto st690
		case 101:
			goto st690
		}
		goto tr416
	st690:
		if p++; p == pe {
			goto _test_eof690
		}
	st_case_690:
		switch data[p] {
		case 82:
			goto st691
		case 114:
			goto st691
		}
		goto tr416
	st691:
		if p++; p == pe {
			goto _test_eof691
		}
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
		if p++; p == pe {
			goto _test_eof692
		}
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
		if p++; p == pe {
			goto _test_eof693
		}
	st_case_693:
		switch data[p] {
		case 74:
			goto st694
		case 106:
			goto st694
		}
		goto tr416
	st694:
		if p++; p == pe {
			goto _test_eof694
		}
	st_case_694:
		switch data[p] {
		case 69:
			goto st695
		case 101:
			goto st695
		}
		goto tr416
	st695:
		if p++; p == pe {
			goto _test_eof695
		}
	st_case_695:
		switch data[p] {
		case 67:
			goto st696
		case 99:
			goto st696
		}
		goto tr416
	st696:
		if p++; p == pe {
			goto _test_eof696
		}
	st_case_696:
		switch data[p] {
		case 84:
			goto st697
		case 116:
			goto st697
		}
		goto tr416
	st697:
		if p++; p == pe {
			goto _test_eof697
		}
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
		if p++; p == pe {
			goto _test_eof698
		}
	st_case_698:
		switch data[p] {
		case 80:
			goto st699
		case 112:
			goto st699
		}
		goto tr416
	st699:
		if p++; p == pe {
			goto _test_eof699
		}
	st_case_699:
		switch data[p] {
		case 79:
			goto st700
		case 111:
			goto st700
		}
		goto tr416
	st700:
		if p++; p == pe {
			goto _test_eof700
		}
	st_case_700:
		switch data[p] {
		case 82:
			goto st701
		case 114:
			goto st701
		}
		goto tr416
	st701:
		if p++; p == pe {
			goto _test_eof701
		}
	st_case_701:
		switch data[p] {
		case 84:
			goto st702
		case 116:
			goto st702
		}
		goto tr416
	st702:
		if p++; p == pe {
			goto _test_eof702
		}
	st_case_702:
		switch data[p] {
		case 69:
			goto st703
		case 101:
			goto st703
		}
		goto tr416
	st703:
		if p++; p == pe {
			goto _test_eof703
		}
	st_case_703:
		switch data[p] {
		case 68:
			goto st504
		case 100:
			goto st504
		}
		goto tr416
	st704:
		if p++; p == pe {
			goto _test_eof704
		}
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
		if p++; p == pe {
			goto _test_eof705
		}
	st_case_705:
		switch data[p] {
		case 77:
			goto st706
		case 109:
			goto st706
		}
		goto tr416
	st706:
		if p++; p == pe {
			goto _test_eof706
		}
	st_case_706:
		switch data[p] {
		case 69:
			goto st707
		case 101:
			goto st707
		}
		goto tr416
	st707:
		if p++; p == pe {
			goto _test_eof707
		}
	st_case_707:
		switch data[p] {
		case 83:
			goto st708
		case 115:
			goto st708
		}
		goto tr416
	st708:
		if p++; p == pe {
			goto _test_eof708
		}
	st_case_708:
		switch data[p] {
		case 84:
			goto st709
		case 116:
			goto st709
		}
		goto tr416
	st709:
		if p++; p == pe {
			goto _test_eof709
		}
	st_case_709:
		switch data[p] {
		case 65:
			goto st710
		case 97:
			goto st710
		}
		goto tr416
	st710:
		if p++; p == pe {
			goto _test_eof710
		}
	st_case_710:
		switch data[p] {
		case 77:
			goto st711
		case 109:
			goto st711
		}
		goto tr416
	st711:
		if p++; p == pe {
			goto _test_eof711
		}
	st_case_711:
		switch data[p] {
		case 80:
			goto st712
		case 112:
			goto st712
		}
		goto tr416
	st712:
		if p++; p == pe {
			goto _test_eof712
		}
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
		if p++; p == pe {
			goto _test_eof713
		}
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
		if p++; p == pe {
			goto _test_eof714
		}
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
		if p++; p == pe {
			goto _test_eof715
		}
	st_case_715:
		switch data[p] {
		case 83:
			goto st716
		case 115:
			goto st716
		}
		goto tr416
	st716:
		if p++; p == pe {
			goto _test_eof716
		}
	st_case_716:
		switch data[p] {
		case 85:
			goto st717
		case 117:
			goto st717
		}
		goto tr416
	st717:
		if p++; p == pe {
			goto _test_eof717
		}
	st_case_717:
		switch data[p] {
		case 80:
			goto st718
		case 112:
			goto st718
		}
		goto tr416
	st718:
		if p++; p == pe {
			goto _test_eof718
		}
	st_case_718:
		switch data[p] {
		case 80:
			goto st719
		case 112:
			goto st719
		}
		goto tr416
	st719:
		if p++; p == pe {
			goto _test_eof719
		}
	st_case_719:
		switch data[p] {
		case 79:
			goto st720
		case 111:
			goto st720
		}
		goto tr416
	st720:
		if p++; p == pe {
			goto _test_eof720
		}
	st_case_720:
		switch data[p] {
		case 82:
			goto st721
		case 114:
			goto st721
		}
		goto tr416
	st721:
		if p++; p == pe {
			goto _test_eof721
		}
	st_case_721:
		switch data[p] {
		case 84:
			goto st722
		case 116:
			goto st722
		}
		goto tr416
	st722:
		if p++; p == pe {
			goto _test_eof722
		}
	st_case_722:
		switch data[p] {
		case 69:
			goto st723
		case 101:
			goto st723
		}
		goto tr416
	st723:
		if p++; p == pe {
			goto _test_eof723
		}
	st_case_723:
		switch data[p] {
		case 68:
			goto st724
		case 100:
			goto st724
		}
		goto tr416
	st724:
		if p++; p == pe {
			goto _test_eof724
		}
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
		if p++; p == pe {
			goto _test_eof725
		}
	st_case_725:
		switch data[p] {
		case 69:
			goto st726
		case 101:
			goto st726
		}
		goto tr416
	st726:
		if p++; p == pe {
			goto _test_eof726
		}
	st_case_726:
		switch data[p] {
		case 82:
			goto st727
		case 114:
			goto st727
		}
		goto tr416
	st727:
		if p++; p == pe {
			goto _test_eof727
		}
	st_case_727:
		if data[p] == 45 {
			goto st728
		}
		goto tr416
	st728:
		if p++; p == pe {
			goto _test_eof728
		}
	st_case_728:
		switch data[p] {
		case 65:
			goto st729
		case 97:
			goto st729
		}
		goto tr416
	st729:
		if p++; p == pe {
			goto _test_eof729
		}
	st_case_729:
		switch data[p] {
		case 71:
			goto st730
		case 103:
			goto st730
		}
		goto tr416
	st730:
		if p++; p == pe {
			goto _test_eof730
		}
	st_case_730:
		switch data[p] {
		case 69:
			goto st731
		case 101:
			goto st731
		}
		goto tr416
	st731:
		if p++; p == pe {
			goto _test_eof731
		}
	st_case_731:
		switch data[p] {
		case 78:
			goto st732
		case 110:
			goto st732
		}
		goto tr416
	st732:
		if p++; p == pe {
			goto _test_eof732
		}
	st_case_732:
		switch data[p] {
		case 84:
			goto st733
		case 116:
			goto st733
		}
		goto tr416
	st733:
		if p++; p == pe {
			goto _test_eof733
		}
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
		if p++; p == pe {
			goto _test_eof734
		}
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
		if p++; p == pe {
			goto _test_eof735
		}
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
		if p++; p == pe {
			goto _test_eof736
		}
	st_case_736:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if lookAheadWSP(data, p, pe) {
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
		if p++; p == pe {
			goto _test_eof737
		}
	st_case_737:
		if data[p] == 10 {
			goto st738
		}
		goto st0
	st738:
		if p++; p == pe {
			goto _test_eof738
		}
	st_case_738:
		switch data[p] {
		case 9:
			goto st739
		case 32:
			goto st739
		}
		goto st0
	st739:
		if p++; p == pe {
			goto _test_eof739
		}
	st_case_739:
		switch data[p] {
		case 9:
			goto st739
		case 32:
			goto st739
		}
		goto tr971
	st740:
		if p++; p == pe {
			goto _test_eof740
		}
	st_case_740:
		switch data[p] {
		case 65:
			goto st741
		case 97:
			goto st741
		}
		goto tr416
	st741:
		if p++; p == pe {
			goto _test_eof741
		}
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
		if p++; p == pe {
			goto _test_eof742
		}
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
		if p++; p == pe {
			goto _test_eof743
		}
	st_case_743:
		switch data[p] {
		case 82:
			goto st744
		case 114:
			goto st744
		}
		goto tr416
	st744:
		if p++; p == pe {
			goto _test_eof744
		}
	st_case_744:
		switch data[p] {
		case 78:
			goto st745
		case 110:
			goto st745
		}
		goto tr416
	st745:
		if p++; p == pe {
			goto _test_eof745
		}
	st_case_745:
		switch data[p] {
		case 73:
			goto st746
		case 105:
			goto st746
		}
		goto tr416
	st746:
		if p++; p == pe {
			goto _test_eof746
		}
	st_case_746:
		switch data[p] {
		case 78:
			goto st747
		case 110:
			goto st747
		}
		goto tr416
	st747:
		if p++; p == pe {
			goto _test_eof747
		}
	st_case_747:
		switch data[p] {
		case 71:
			goto st748
		case 103:
			goto st748
		}
		goto tr416
	st748:
		if p++; p == pe {
			goto _test_eof748
		}
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
		if p++; p == pe {
			goto _test_eof749
		}
	st_case_749:
		switch data[p] {
		case 87:
			goto st750
		case 119:
			goto st750
		}
		goto tr416
	st750:
		if p++; p == pe {
			goto _test_eof750
		}
	st_case_750:
		if data[p] == 45 {
			goto st751
		}
		goto tr416
	st751:
		if p++; p == pe {
			goto _test_eof751
		}
	st_case_751:
		switch data[p] {
		case 65:
			goto st752
		case 97:
			goto st752
		}
		goto tr416
	st752:
		if p++; p == pe {
			goto _test_eof752
		}
	st_case_752:
		switch data[p] {
		case 85:
			goto st753
		case 117:
			goto st753
		}
		goto tr416
	st753:
		if p++; p == pe {
			goto _test_eof753
		}
	st_case_753:
		switch data[p] {
		case 84:
			goto st754
		case 116:
			goto st754
		}
		goto tr416
	st754:
		if p++; p == pe {
			goto _test_eof754
		}
	st_case_754:
		switch data[p] {
		case 72:
			goto st755
		case 104:
			goto st755
		}
		goto tr416
	st755:
		if p++; p == pe {
			goto _test_eof755
		}
	st_case_755:
		switch data[p] {
		case 69:
			goto st756
		case 101:
			goto st756
		}
		goto tr416
	st756:
		if p++; p == pe {
			goto _test_eof756
		}
	st_case_756:
		switch data[p] {
		case 78:
			goto st757
		case 110:
			goto st757
		}
		goto tr416
	st757:
		if p++; p == pe {
			goto _test_eof757
		}
	st_case_757:
		switch data[p] {
		case 84:
			goto st758
		case 116:
			goto st758
		}
		goto tr416
	st758:
		if p++; p == pe {
			goto _test_eof758
		}
	st_case_758:
		switch data[p] {
		case 73:
			goto st759
		case 105:
			goto st759
		}
		goto tr416
	st759:
		if p++; p == pe {
			goto _test_eof759
		}
	st_case_759:
		switch data[p] {
		case 67:
			goto st760
		case 99:
			goto st760
		}
		goto tr416
	st760:
		if p++; p == pe {
			goto _test_eof760
		}
	st_case_760:
		switch data[p] {
		case 65:
			goto st761
		case 97:
			goto st761
		}
		goto tr416
	st761:
		if p++; p == pe {
			goto _test_eof761
		}
	st_case_761:
		switch data[p] {
		case 84:
			goto st762
		case 116:
			goto st762
		}
		goto tr416
	st762:
		if p++; p == pe {
			goto _test_eof762
		}
	st_case_762:
		switch data[p] {
		case 69:
			goto st763
		case 101:
			goto st763
		}
		goto tr416
	st763:
		if p++; p == pe {
			goto _test_eof763
		}
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
		if p++; p == pe {
			goto _test_eof764
		}
	st_case_764:
		if data[p] == 10 {
			goto tr1001
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof765:
		cs = 765
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof766:
		cs = 766
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof767:
		cs = 767
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof75:
		cs = 75
		goto _test_eof
	_test_eof76:
		cs = 76
		goto _test_eof
	_test_eof77:
		cs = 77
		goto _test_eof
	_test_eof78:
		cs = 78
		goto _test_eof
	_test_eof79:
		cs = 79
		goto _test_eof
	_test_eof80:
		cs = 80
		goto _test_eof
	_test_eof81:
		cs = 81
		goto _test_eof
	_test_eof82:
		cs = 82
		goto _test_eof
	_test_eof83:
		cs = 83
		goto _test_eof
	_test_eof84:
		cs = 84
		goto _test_eof
	_test_eof85:
		cs = 85
		goto _test_eof
	_test_eof86:
		cs = 86
		goto _test_eof
	_test_eof87:
		cs = 87
		goto _test_eof
	_test_eof88:
		cs = 88
		goto _test_eof
	_test_eof89:
		cs = 89
		goto _test_eof
	_test_eof90:
		cs = 90
		goto _test_eof
	_test_eof91:
		cs = 91
		goto _test_eof
	_test_eof92:
		cs = 92
		goto _test_eof
	_test_eof93:
		cs = 93
		goto _test_eof
	_test_eof94:
		cs = 94
		goto _test_eof
	_test_eof95:
		cs = 95
		goto _test_eof
	_test_eof96:
		cs = 96
		goto _test_eof
	_test_eof97:
		cs = 97
		goto _test_eof
	_test_eof98:
		cs = 98
		goto _test_eof
	_test_eof99:
		cs = 99
		goto _test_eof
	_test_eof100:
		cs = 100
		goto _test_eof
	_test_eof101:
		cs = 101
		goto _test_eof
	_test_eof102:
		cs = 102
		goto _test_eof
	_test_eof103:
		cs = 103
		goto _test_eof
	_test_eof104:
		cs = 104
		goto _test_eof
	_test_eof105:
		cs = 105
		goto _test_eof
	_test_eof106:
		cs = 106
		goto _test_eof
	_test_eof107:
		cs = 107
		goto _test_eof
	_test_eof108:
		cs = 108
		goto _test_eof
	_test_eof109:
		cs = 109
		goto _test_eof
	_test_eof110:
		cs = 110
		goto _test_eof
	_test_eof111:
		cs = 111
		goto _test_eof
	_test_eof112:
		cs = 112
		goto _test_eof
	_test_eof113:
		cs = 113
		goto _test_eof
	_test_eof114:
		cs = 114
		goto _test_eof
	_test_eof768:
		cs = 768
		goto _test_eof
	_test_eof115:
		cs = 115
		goto _test_eof
	_test_eof116:
		cs = 116
		goto _test_eof
	_test_eof117:
		cs = 117
		goto _test_eof
	_test_eof118:
		cs = 118
		goto _test_eof
	_test_eof119:
		cs = 119
		goto _test_eof
	_test_eof120:
		cs = 120
		goto _test_eof
	_test_eof121:
		cs = 121
		goto _test_eof
	_test_eof122:
		cs = 122
		goto _test_eof
	_test_eof123:
		cs = 123
		goto _test_eof
	_test_eof124:
		cs = 124
		goto _test_eof
	_test_eof125:
		cs = 125
		goto _test_eof
	_test_eof126:
		cs = 126
		goto _test_eof
	_test_eof127:
		cs = 127
		goto _test_eof
	_test_eof128:
		cs = 128
		goto _test_eof
	_test_eof129:
		cs = 129
		goto _test_eof
	_test_eof130:
		cs = 130
		goto _test_eof
	_test_eof131:
		cs = 131
		goto _test_eof
	_test_eof132:
		cs = 132
		goto _test_eof
	_test_eof133:
		cs = 133
		goto _test_eof
	_test_eof134:
		cs = 134
		goto _test_eof
	_test_eof135:
		cs = 135
		goto _test_eof
	_test_eof136:
		cs = 136
		goto _test_eof
	_test_eof137:
		cs = 137
		goto _test_eof
	_test_eof138:
		cs = 138
		goto _test_eof
	_test_eof139:
		cs = 139
		goto _test_eof
	_test_eof140:
		cs = 140
		goto _test_eof
	_test_eof141:
		cs = 141
		goto _test_eof
	_test_eof142:
		cs = 142
		goto _test_eof
	_test_eof143:
		cs = 143
		goto _test_eof
	_test_eof144:
		cs = 144
		goto _test_eof
	_test_eof145:
		cs = 145
		goto _test_eof
	_test_eof146:
		cs = 146
		goto _test_eof
	_test_eof147:
		cs = 147
		goto _test_eof
	_test_eof148:
		cs = 148
		goto _test_eof
	_test_eof149:
		cs = 149
		goto _test_eof
	_test_eof150:
		cs = 150
		goto _test_eof
	_test_eof151:
		cs = 151
		goto _test_eof
	_test_eof152:
		cs = 152
		goto _test_eof
	_test_eof153:
		cs = 153
		goto _test_eof
	_test_eof154:
		cs = 154
		goto _test_eof
	_test_eof155:
		cs = 155
		goto _test_eof
	_test_eof156:
		cs = 156
		goto _test_eof
	_test_eof769:
		cs = 769
		goto _test_eof
	_test_eof157:
		cs = 157
		goto _test_eof
	_test_eof158:
		cs = 158
		goto _test_eof
	_test_eof159:
		cs = 159
		goto _test_eof
	_test_eof160:
		cs = 160
		goto _test_eof
	_test_eof161:
		cs = 161
		goto _test_eof
	_test_eof162:
		cs = 162
		goto _test_eof
	_test_eof163:
		cs = 163
		goto _test_eof
	_test_eof164:
		cs = 164
		goto _test_eof
	_test_eof165:
		cs = 165
		goto _test_eof
	_test_eof166:
		cs = 166
		goto _test_eof
	_test_eof167:
		cs = 167
		goto _test_eof
	_test_eof168:
		cs = 168
		goto _test_eof
	_test_eof169:
		cs = 169
		goto _test_eof
	_test_eof170:
		cs = 170
		goto _test_eof
	_test_eof171:
		cs = 171
		goto _test_eof
	_test_eof172:
		cs = 172
		goto _test_eof
	_test_eof173:
		cs = 173
		goto _test_eof
	_test_eof174:
		cs = 174
		goto _test_eof
	_test_eof175:
		cs = 175
		goto _test_eof
	_test_eof176:
		cs = 176
		goto _test_eof
	_test_eof177:
		cs = 177
		goto _test_eof
	_test_eof178:
		cs = 178
		goto _test_eof
	_test_eof179:
		cs = 179
		goto _test_eof
	_test_eof180:
		cs = 180
		goto _test_eof
	_test_eof181:
		cs = 181
		goto _test_eof
	_test_eof182:
		cs = 182
		goto _test_eof
	_test_eof183:
		cs = 183
		goto _test_eof
	_test_eof184:
		cs = 184
		goto _test_eof
	_test_eof185:
		cs = 185
		goto _test_eof
	_test_eof186:
		cs = 186
		goto _test_eof
	_test_eof187:
		cs = 187
		goto _test_eof
	_test_eof188:
		cs = 188
		goto _test_eof
	_test_eof189:
		cs = 189
		goto _test_eof
	_test_eof190:
		cs = 190
		goto _test_eof
	_test_eof191:
		cs = 191
		goto _test_eof
	_test_eof192:
		cs = 192
		goto _test_eof
	_test_eof193:
		cs = 193
		goto _test_eof
	_test_eof194:
		cs = 194
		goto _test_eof
	_test_eof195:
		cs = 195
		goto _test_eof
	_test_eof770:
		cs = 770
		goto _test_eof
	_test_eof196:
		cs = 196
		goto _test_eof
	_test_eof197:
		cs = 197
		goto _test_eof
	_test_eof198:
		cs = 198
		goto _test_eof
	_test_eof199:
		cs = 199
		goto _test_eof
	_test_eof200:
		cs = 200
		goto _test_eof
	_test_eof201:
		cs = 201
		goto _test_eof
	_test_eof202:
		cs = 202
		goto _test_eof
	_test_eof203:
		cs = 203
		goto _test_eof
	_test_eof204:
		cs = 204
		goto _test_eof
	_test_eof205:
		cs = 205
		goto _test_eof
	_test_eof206:
		cs = 206
		goto _test_eof
	_test_eof207:
		cs = 207
		goto _test_eof
	_test_eof208:
		cs = 208
		goto _test_eof
	_test_eof209:
		cs = 209
		goto _test_eof
	_test_eof210:
		cs = 210
		goto _test_eof
	_test_eof211:
		cs = 211
		goto _test_eof
	_test_eof212:
		cs = 212
		goto _test_eof
	_test_eof213:
		cs = 213
		goto _test_eof
	_test_eof214:
		cs = 214
		goto _test_eof
	_test_eof215:
		cs = 215
		goto _test_eof
	_test_eof216:
		cs = 216
		goto _test_eof
	_test_eof217:
		cs = 217
		goto _test_eof
	_test_eof218:
		cs = 218
		goto _test_eof
	_test_eof219:
		cs = 219
		goto _test_eof
	_test_eof220:
		cs = 220
		goto _test_eof
	_test_eof221:
		cs = 221
		goto _test_eof
	_test_eof222:
		cs = 222
		goto _test_eof
	_test_eof223:
		cs = 223
		goto _test_eof
	_test_eof224:
		cs = 224
		goto _test_eof
	_test_eof225:
		cs = 225
		goto _test_eof
	_test_eof226:
		cs = 226
		goto _test_eof
	_test_eof227:
		cs = 227
		goto _test_eof
	_test_eof228:
		cs = 228
		goto _test_eof
	_test_eof229:
		cs = 229
		goto _test_eof
	_test_eof230:
		cs = 230
		goto _test_eof
	_test_eof231:
		cs = 231
		goto _test_eof
	_test_eof232:
		cs = 232
		goto _test_eof
	_test_eof233:
		cs = 233
		goto _test_eof
	_test_eof234:
		cs = 234
		goto _test_eof
	_test_eof771:
		cs = 771
		goto _test_eof
	_test_eof235:
		cs = 235
		goto _test_eof
	_test_eof236:
		cs = 236
		goto _test_eof
	_test_eof237:
		cs = 237
		goto _test_eof
	_test_eof238:
		cs = 238
		goto _test_eof
	_test_eof239:
		cs = 239
		goto _test_eof
	_test_eof240:
		cs = 240
		goto _test_eof
	_test_eof241:
		cs = 241
		goto _test_eof
	_test_eof242:
		cs = 242
		goto _test_eof
	_test_eof243:
		cs = 243
		goto _test_eof
	_test_eof244:
		cs = 244
		goto _test_eof
	_test_eof245:
		cs = 245
		goto _test_eof
	_test_eof246:
		cs = 246
		goto _test_eof
	_test_eof772:
		cs = 772
		goto _test_eof
	_test_eof247:
		cs = 247
		goto _test_eof
	_test_eof248:
		cs = 248
		goto _test_eof
	_test_eof249:
		cs = 249
		goto _test_eof
	_test_eof773:
		cs = 773
		goto _test_eof
	_test_eof250:
		cs = 250
		goto _test_eof
	_test_eof251:
		cs = 251
		goto _test_eof
	_test_eof774:
		cs = 774
		goto _test_eof
	_test_eof252:
		cs = 252
		goto _test_eof
	_test_eof253:
		cs = 253
		goto _test_eof
	_test_eof254:
		cs = 254
		goto _test_eof
	_test_eof775:
		cs = 775
		goto _test_eof
	_test_eof776:
		cs = 776
		goto _test_eof
	_test_eof777:
		cs = 777
		goto _test_eof
	_test_eof778:
		cs = 778
		goto _test_eof
	_test_eof255:
		cs = 255
		goto _test_eof
	_test_eof256:
		cs = 256
		goto _test_eof
	_test_eof257:
		cs = 257
		goto _test_eof
	_test_eof258:
		cs = 258
		goto _test_eof
	_test_eof779:
		cs = 779
		goto _test_eof
	_test_eof259:
		cs = 259
		goto _test_eof
	_test_eof260:
		cs = 260
		goto _test_eof
	_test_eof261:
		cs = 261
		goto _test_eof
	_test_eof262:
		cs = 262
		goto _test_eof
	_test_eof263:
		cs = 263
		goto _test_eof
	_test_eof264:
		cs = 264
		goto _test_eof
	_test_eof265:
		cs = 265
		goto _test_eof
	_test_eof266:
		cs = 266
		goto _test_eof
	_test_eof267:
		cs = 267
		goto _test_eof
	_test_eof268:
		cs = 268
		goto _test_eof
	_test_eof269:
		cs = 269
		goto _test_eof
	_test_eof270:
		cs = 270
		goto _test_eof
	_test_eof780:
		cs = 780
		goto _test_eof
	_test_eof271:
		cs = 271
		goto _test_eof
	_test_eof272:
		cs = 272
		goto _test_eof
	_test_eof273:
		cs = 273
		goto _test_eof
	_test_eof274:
		cs = 274
		goto _test_eof
	_test_eof275:
		cs = 275
		goto _test_eof
	_test_eof276:
		cs = 276
		goto _test_eof
	_test_eof781:
		cs = 781
		goto _test_eof
	_test_eof277:
		cs = 277
		goto _test_eof
	_test_eof278:
		cs = 278
		goto _test_eof
	_test_eof279:
		cs = 279
		goto _test_eof
	_test_eof280:
		cs = 280
		goto _test_eof
	_test_eof281:
		cs = 281
		goto _test_eof
	_test_eof282:
		cs = 282
		goto _test_eof
	_test_eof283:
		cs = 283
		goto _test_eof
	_test_eof284:
		cs = 284
		goto _test_eof
	_test_eof782:
		cs = 782
		goto _test_eof
	_test_eof285:
		cs = 285
		goto _test_eof
	_test_eof286:
		cs = 286
		goto _test_eof
	_test_eof287:
		cs = 287
		goto _test_eof
	_test_eof288:
		cs = 288
		goto _test_eof
	_test_eof289:
		cs = 289
		goto _test_eof
	_test_eof290:
		cs = 290
		goto _test_eof
	_test_eof291:
		cs = 291
		goto _test_eof
	_test_eof292:
		cs = 292
		goto _test_eof
	_test_eof293:
		cs = 293
		goto _test_eof
	_test_eof294:
		cs = 294
		goto _test_eof
	_test_eof295:
		cs = 295
		goto _test_eof
	_test_eof296:
		cs = 296
		goto _test_eof
	_test_eof297:
		cs = 297
		goto _test_eof
	_test_eof298:
		cs = 298
		goto _test_eof
	_test_eof299:
		cs = 299
		goto _test_eof
	_test_eof300:
		cs = 300
		goto _test_eof
	_test_eof301:
		cs = 301
		goto _test_eof
	_test_eof302:
		cs = 302
		goto _test_eof
	_test_eof303:
		cs = 303
		goto _test_eof
	_test_eof304:
		cs = 304
		goto _test_eof
	_test_eof305:
		cs = 305
		goto _test_eof
	_test_eof306:
		cs = 306
		goto _test_eof
	_test_eof307:
		cs = 307
		goto _test_eof
	_test_eof308:
		cs = 308
		goto _test_eof
	_test_eof309:
		cs = 309
		goto _test_eof
	_test_eof310:
		cs = 310
		goto _test_eof
	_test_eof311:
		cs = 311
		goto _test_eof
	_test_eof312:
		cs = 312
		goto _test_eof
	_test_eof313:
		cs = 313
		goto _test_eof
	_test_eof314:
		cs = 314
		goto _test_eof
	_test_eof315:
		cs = 315
		goto _test_eof
	_test_eof316:
		cs = 316
		goto _test_eof
	_test_eof317:
		cs = 317
		goto _test_eof
	_test_eof318:
		cs = 318
		goto _test_eof
	_test_eof319:
		cs = 319
		goto _test_eof
	_test_eof320:
		cs = 320
		goto _test_eof
	_test_eof321:
		cs = 321
		goto _test_eof
	_test_eof322:
		cs = 322
		goto _test_eof
	_test_eof323:
		cs = 323
		goto _test_eof
	_test_eof324:
		cs = 324
		goto _test_eof
	_test_eof325:
		cs = 325
		goto _test_eof
	_test_eof326:
		cs = 326
		goto _test_eof
	_test_eof327:
		cs = 327
		goto _test_eof
	_test_eof328:
		cs = 328
		goto _test_eof
	_test_eof329:
		cs = 329
		goto _test_eof
	_test_eof330:
		cs = 330
		goto _test_eof
	_test_eof331:
		cs = 331
		goto _test_eof
	_test_eof332:
		cs = 332
		goto _test_eof
	_test_eof333:
		cs = 333
		goto _test_eof
	_test_eof334:
		cs = 334
		goto _test_eof
	_test_eof335:
		cs = 335
		goto _test_eof
	_test_eof336:
		cs = 336
		goto _test_eof
	_test_eof337:
		cs = 337
		goto _test_eof
	_test_eof338:
		cs = 338
		goto _test_eof
	_test_eof339:
		cs = 339
		goto _test_eof
	_test_eof340:
		cs = 340
		goto _test_eof
	_test_eof341:
		cs = 341
		goto _test_eof
	_test_eof342:
		cs = 342
		goto _test_eof
	_test_eof343:
		cs = 343
		goto _test_eof
	_test_eof344:
		cs = 344
		goto _test_eof
	_test_eof345:
		cs = 345
		goto _test_eof
	_test_eof346:
		cs = 346
		goto _test_eof
	_test_eof347:
		cs = 347
		goto _test_eof
	_test_eof348:
		cs = 348
		goto _test_eof
	_test_eof349:
		cs = 349
		goto _test_eof
	_test_eof350:
		cs = 350
		goto _test_eof
	_test_eof351:
		cs = 351
		goto _test_eof
	_test_eof352:
		cs = 352
		goto _test_eof
	_test_eof353:
		cs = 353
		goto _test_eof
	_test_eof354:
		cs = 354
		goto _test_eof
	_test_eof355:
		cs = 355
		goto _test_eof
	_test_eof356:
		cs = 356
		goto _test_eof
	_test_eof357:
		cs = 357
		goto _test_eof
	_test_eof358:
		cs = 358
		goto _test_eof
	_test_eof359:
		cs = 359
		goto _test_eof
	_test_eof360:
		cs = 360
		goto _test_eof
	_test_eof361:
		cs = 361
		goto _test_eof
	_test_eof362:
		cs = 362
		goto _test_eof
	_test_eof363:
		cs = 363
		goto _test_eof
	_test_eof364:
		cs = 364
		goto _test_eof
	_test_eof365:
		cs = 365
		goto _test_eof
	_test_eof366:
		cs = 366
		goto _test_eof
	_test_eof367:
		cs = 367
		goto _test_eof
	_test_eof368:
		cs = 368
		goto _test_eof
	_test_eof369:
		cs = 369
		goto _test_eof
	_test_eof370:
		cs = 370
		goto _test_eof
	_test_eof371:
		cs = 371
		goto _test_eof
	_test_eof372:
		cs = 372
		goto _test_eof
	_test_eof373:
		cs = 373
		goto _test_eof
	_test_eof374:
		cs = 374
		goto _test_eof
	_test_eof375:
		cs = 375
		goto _test_eof
	_test_eof376:
		cs = 376
		goto _test_eof
	_test_eof377:
		cs = 377
		goto _test_eof
	_test_eof378:
		cs = 378
		goto _test_eof
	_test_eof379:
		cs = 379
		goto _test_eof
	_test_eof380:
		cs = 380
		goto _test_eof
	_test_eof381:
		cs = 381
		goto _test_eof
	_test_eof382:
		cs = 382
		goto _test_eof
	_test_eof383:
		cs = 383
		goto _test_eof
	_test_eof384:
		cs = 384
		goto _test_eof
	_test_eof385:
		cs = 385
		goto _test_eof
	_test_eof386:
		cs = 386
		goto _test_eof
	_test_eof387:
		cs = 387
		goto _test_eof
	_test_eof388:
		cs = 388
		goto _test_eof
	_test_eof389:
		cs = 389
		goto _test_eof
	_test_eof390:
		cs = 390
		goto _test_eof
	_test_eof391:
		cs = 391
		goto _test_eof
	_test_eof392:
		cs = 392
		goto _test_eof
	_test_eof393:
		cs = 393
		goto _test_eof
	_test_eof394:
		cs = 394
		goto _test_eof
	_test_eof395:
		cs = 395
		goto _test_eof
	_test_eof396:
		cs = 396
		goto _test_eof
	_test_eof397:
		cs = 397
		goto _test_eof
	_test_eof398:
		cs = 398
		goto _test_eof
	_test_eof399:
		cs = 399
		goto _test_eof
	_test_eof400:
		cs = 400
		goto _test_eof
	_test_eof401:
		cs = 401
		goto _test_eof
	_test_eof402:
		cs = 402
		goto _test_eof
	_test_eof403:
		cs = 403
		goto _test_eof
	_test_eof404:
		cs = 404
		goto _test_eof
	_test_eof405:
		cs = 405
		goto _test_eof
	_test_eof406:
		cs = 406
		goto _test_eof
	_test_eof407:
		cs = 407
		goto _test_eof
	_test_eof408:
		cs = 408
		goto _test_eof
	_test_eof409:
		cs = 409
		goto _test_eof
	_test_eof410:
		cs = 410
		goto _test_eof
	_test_eof411:
		cs = 411
		goto _test_eof
	_test_eof412:
		cs = 412
		goto _test_eof
	_test_eof413:
		cs = 413
		goto _test_eof
	_test_eof414:
		cs = 414
		goto _test_eof
	_test_eof415:
		cs = 415
		goto _test_eof
	_test_eof416:
		cs = 416
		goto _test_eof
	_test_eof417:
		cs = 417
		goto _test_eof
	_test_eof418:
		cs = 418
		goto _test_eof
	_test_eof419:
		cs = 419
		goto _test_eof
	_test_eof420:
		cs = 420
		goto _test_eof
	_test_eof421:
		cs = 421
		goto _test_eof
	_test_eof422:
		cs = 422
		goto _test_eof
	_test_eof423:
		cs = 423
		goto _test_eof
	_test_eof424:
		cs = 424
		goto _test_eof
	_test_eof425:
		cs = 425
		goto _test_eof
	_test_eof426:
		cs = 426
		goto _test_eof
	_test_eof427:
		cs = 427
		goto _test_eof
	_test_eof428:
		cs = 428
		goto _test_eof
	_test_eof429:
		cs = 429
		goto _test_eof
	_test_eof430:
		cs = 430
		goto _test_eof
	_test_eof431:
		cs = 431
		goto _test_eof
	_test_eof432:
		cs = 432
		goto _test_eof
	_test_eof433:
		cs = 433
		goto _test_eof
	_test_eof434:
		cs = 434
		goto _test_eof
	_test_eof435:
		cs = 435
		goto _test_eof
	_test_eof436:
		cs = 436
		goto _test_eof
	_test_eof437:
		cs = 437
		goto _test_eof
	_test_eof438:
		cs = 438
		goto _test_eof
	_test_eof439:
		cs = 439
		goto _test_eof
	_test_eof440:
		cs = 440
		goto _test_eof
	_test_eof441:
		cs = 441
		goto _test_eof
	_test_eof442:
		cs = 442
		goto _test_eof
	_test_eof443:
		cs = 443
		goto _test_eof
	_test_eof444:
		cs = 444
		goto _test_eof
	_test_eof445:
		cs = 445
		goto _test_eof
	_test_eof446:
		cs = 446
		goto _test_eof
	_test_eof447:
		cs = 447
		goto _test_eof
	_test_eof448:
		cs = 448
		goto _test_eof
	_test_eof449:
		cs = 449
		goto _test_eof
	_test_eof450:
		cs = 450
		goto _test_eof
	_test_eof451:
		cs = 451
		goto _test_eof
	_test_eof452:
		cs = 452
		goto _test_eof
	_test_eof453:
		cs = 453
		goto _test_eof
	_test_eof454:
		cs = 454
		goto _test_eof
	_test_eof455:
		cs = 455
		goto _test_eof
	_test_eof456:
		cs = 456
		goto _test_eof
	_test_eof457:
		cs = 457
		goto _test_eof
	_test_eof458:
		cs = 458
		goto _test_eof
	_test_eof459:
		cs = 459
		goto _test_eof
	_test_eof460:
		cs = 460
		goto _test_eof
	_test_eof461:
		cs = 461
		goto _test_eof
	_test_eof462:
		cs = 462
		goto _test_eof
	_test_eof463:
		cs = 463
		goto _test_eof
	_test_eof464:
		cs = 464
		goto _test_eof
	_test_eof465:
		cs = 465
		goto _test_eof
	_test_eof466:
		cs = 466
		goto _test_eof
	_test_eof467:
		cs = 467
		goto _test_eof
	_test_eof468:
		cs = 468
		goto _test_eof
	_test_eof469:
		cs = 469
		goto _test_eof
	_test_eof470:
		cs = 470
		goto _test_eof
	_test_eof471:
		cs = 471
		goto _test_eof
	_test_eof472:
		cs = 472
		goto _test_eof
	_test_eof473:
		cs = 473
		goto _test_eof
	_test_eof474:
		cs = 474
		goto _test_eof
	_test_eof475:
		cs = 475
		goto _test_eof
	_test_eof476:
		cs = 476
		goto _test_eof
	_test_eof477:
		cs = 477
		goto _test_eof
	_test_eof478:
		cs = 478
		goto _test_eof
	_test_eof479:
		cs = 479
		goto _test_eof
	_test_eof480:
		cs = 480
		goto _test_eof
	_test_eof481:
		cs = 481
		goto _test_eof
	_test_eof482:
		cs = 482
		goto _test_eof
	_test_eof483:
		cs = 483
		goto _test_eof
	_test_eof484:
		cs = 484
		goto _test_eof
	_test_eof485:
		cs = 485
		goto _test_eof
	_test_eof486:
		cs = 486
		goto _test_eof
	_test_eof487:
		cs = 487
		goto _test_eof
	_test_eof488:
		cs = 488
		goto _test_eof
	_test_eof489:
		cs = 489
		goto _test_eof
	_test_eof490:
		cs = 490
		goto _test_eof
	_test_eof491:
		cs = 491
		goto _test_eof
	_test_eof492:
		cs = 492
		goto _test_eof
	_test_eof493:
		cs = 493
		goto _test_eof
	_test_eof494:
		cs = 494
		goto _test_eof
	_test_eof495:
		cs = 495
		goto _test_eof
	_test_eof496:
		cs = 496
		goto _test_eof
	_test_eof497:
		cs = 497
		goto _test_eof
	_test_eof498:
		cs = 498
		goto _test_eof
	_test_eof499:
		cs = 499
		goto _test_eof
	_test_eof500:
		cs = 500
		goto _test_eof
	_test_eof501:
		cs = 501
		goto _test_eof
	_test_eof502:
		cs = 502
		goto _test_eof
	_test_eof503:
		cs = 503
		goto _test_eof
	_test_eof504:
		cs = 504
		goto _test_eof
	_test_eof505:
		cs = 505
		goto _test_eof
	_test_eof506:
		cs = 506
		goto _test_eof
	_test_eof507:
		cs = 507
		goto _test_eof
	_test_eof508:
		cs = 508
		goto _test_eof
	_test_eof509:
		cs = 509
		goto _test_eof
	_test_eof510:
		cs = 510
		goto _test_eof
	_test_eof511:
		cs = 511
		goto _test_eof
	_test_eof512:
		cs = 512
		goto _test_eof
	_test_eof513:
		cs = 513
		goto _test_eof
	_test_eof514:
		cs = 514
		goto _test_eof
	_test_eof515:
		cs = 515
		goto _test_eof
	_test_eof516:
		cs = 516
		goto _test_eof
	_test_eof517:
		cs = 517
		goto _test_eof
	_test_eof518:
		cs = 518
		goto _test_eof
	_test_eof519:
		cs = 519
		goto _test_eof
	_test_eof520:
		cs = 520
		goto _test_eof
	_test_eof521:
		cs = 521
		goto _test_eof
	_test_eof522:
		cs = 522
		goto _test_eof
	_test_eof523:
		cs = 523
		goto _test_eof
	_test_eof524:
		cs = 524
		goto _test_eof
	_test_eof525:
		cs = 525
		goto _test_eof
	_test_eof526:
		cs = 526
		goto _test_eof
	_test_eof527:
		cs = 527
		goto _test_eof
	_test_eof528:
		cs = 528
		goto _test_eof
	_test_eof529:
		cs = 529
		goto _test_eof
	_test_eof530:
		cs = 530
		goto _test_eof
	_test_eof531:
		cs = 531
		goto _test_eof
	_test_eof532:
		cs = 532
		goto _test_eof
	_test_eof533:
		cs = 533
		goto _test_eof
	_test_eof534:
		cs = 534
		goto _test_eof
	_test_eof535:
		cs = 535
		goto _test_eof
	_test_eof536:
		cs = 536
		goto _test_eof
	_test_eof537:
		cs = 537
		goto _test_eof
	_test_eof538:
		cs = 538
		goto _test_eof
	_test_eof539:
		cs = 539
		goto _test_eof
	_test_eof540:
		cs = 540
		goto _test_eof
	_test_eof541:
		cs = 541
		goto _test_eof
	_test_eof542:
		cs = 542
		goto _test_eof
	_test_eof543:
		cs = 543
		goto _test_eof
	_test_eof544:
		cs = 544
		goto _test_eof
	_test_eof545:
		cs = 545
		goto _test_eof
	_test_eof546:
		cs = 546
		goto _test_eof
	_test_eof547:
		cs = 547
		goto _test_eof
	_test_eof548:
		cs = 548
		goto _test_eof
	_test_eof549:
		cs = 549
		goto _test_eof
	_test_eof550:
		cs = 550
		goto _test_eof
	_test_eof551:
		cs = 551
		goto _test_eof
	_test_eof552:
		cs = 552
		goto _test_eof
	_test_eof553:
		cs = 553
		goto _test_eof
	_test_eof554:
		cs = 554
		goto _test_eof
	_test_eof555:
		cs = 555
		goto _test_eof
	_test_eof556:
		cs = 556
		goto _test_eof
	_test_eof557:
		cs = 557
		goto _test_eof
	_test_eof558:
		cs = 558
		goto _test_eof
	_test_eof559:
		cs = 559
		goto _test_eof
	_test_eof560:
		cs = 560
		goto _test_eof
	_test_eof561:
		cs = 561
		goto _test_eof
	_test_eof562:
		cs = 562
		goto _test_eof
	_test_eof563:
		cs = 563
		goto _test_eof
	_test_eof564:
		cs = 564
		goto _test_eof
	_test_eof565:
		cs = 565
		goto _test_eof
	_test_eof566:
		cs = 566
		goto _test_eof
	_test_eof567:
		cs = 567
		goto _test_eof
	_test_eof568:
		cs = 568
		goto _test_eof
	_test_eof569:
		cs = 569
		goto _test_eof
	_test_eof570:
		cs = 570
		goto _test_eof
	_test_eof571:
		cs = 571
		goto _test_eof
	_test_eof572:
		cs = 572
		goto _test_eof
	_test_eof573:
		cs = 573
		goto _test_eof
	_test_eof574:
		cs = 574
		goto _test_eof
	_test_eof575:
		cs = 575
		goto _test_eof
	_test_eof576:
		cs = 576
		goto _test_eof
	_test_eof577:
		cs = 577
		goto _test_eof
	_test_eof578:
		cs = 578
		goto _test_eof
	_test_eof579:
		cs = 579
		goto _test_eof
	_test_eof580:
		cs = 580
		goto _test_eof
	_test_eof581:
		cs = 581
		goto _test_eof
	_test_eof582:
		cs = 582
		goto _test_eof
	_test_eof583:
		cs = 583
		goto _test_eof
	_test_eof584:
		cs = 584
		goto _test_eof
	_test_eof585:
		cs = 585
		goto _test_eof
	_test_eof586:
		cs = 586
		goto _test_eof
	_test_eof587:
		cs = 587
		goto _test_eof
	_test_eof588:
		cs = 588
		goto _test_eof
	_test_eof589:
		cs = 589
		goto _test_eof
	_test_eof590:
		cs = 590
		goto _test_eof
	_test_eof591:
		cs = 591
		goto _test_eof
	_test_eof592:
		cs = 592
		goto _test_eof
	_test_eof593:
		cs = 593
		goto _test_eof
	_test_eof594:
		cs = 594
		goto _test_eof
	_test_eof595:
		cs = 595
		goto _test_eof
	_test_eof596:
		cs = 596
		goto _test_eof
	_test_eof597:
		cs = 597
		goto _test_eof
	_test_eof598:
		cs = 598
		goto _test_eof
	_test_eof599:
		cs = 599
		goto _test_eof
	_test_eof600:
		cs = 600
		goto _test_eof
	_test_eof601:
		cs = 601
		goto _test_eof
	_test_eof602:
		cs = 602
		goto _test_eof
	_test_eof603:
		cs = 603
		goto _test_eof
	_test_eof604:
		cs = 604
		goto _test_eof
	_test_eof605:
		cs = 605
		goto _test_eof
	_test_eof606:
		cs = 606
		goto _test_eof
	_test_eof607:
		cs = 607
		goto _test_eof
	_test_eof608:
		cs = 608
		goto _test_eof
	_test_eof609:
		cs = 609
		goto _test_eof
	_test_eof610:
		cs = 610
		goto _test_eof
	_test_eof611:
		cs = 611
		goto _test_eof
	_test_eof612:
		cs = 612
		goto _test_eof
	_test_eof613:
		cs = 613
		goto _test_eof
	_test_eof614:
		cs = 614
		goto _test_eof
	_test_eof615:
		cs = 615
		goto _test_eof
	_test_eof616:
		cs = 616
		goto _test_eof
	_test_eof617:
		cs = 617
		goto _test_eof
	_test_eof618:
		cs = 618
		goto _test_eof
	_test_eof619:
		cs = 619
		goto _test_eof
	_test_eof620:
		cs = 620
		goto _test_eof
	_test_eof621:
		cs = 621
		goto _test_eof
	_test_eof622:
		cs = 622
		goto _test_eof
	_test_eof623:
		cs = 623
		goto _test_eof
	_test_eof624:
		cs = 624
		goto _test_eof
	_test_eof625:
		cs = 625
		goto _test_eof
	_test_eof626:
		cs = 626
		goto _test_eof
	_test_eof627:
		cs = 627
		goto _test_eof
	_test_eof628:
		cs = 628
		goto _test_eof
	_test_eof629:
		cs = 629
		goto _test_eof
	_test_eof630:
		cs = 630
		goto _test_eof
	_test_eof631:
		cs = 631
		goto _test_eof
	_test_eof632:
		cs = 632
		goto _test_eof
	_test_eof633:
		cs = 633
		goto _test_eof
	_test_eof634:
		cs = 634
		goto _test_eof
	_test_eof635:
		cs = 635
		goto _test_eof
	_test_eof636:
		cs = 636
		goto _test_eof
	_test_eof637:
		cs = 637
		goto _test_eof
	_test_eof638:
		cs = 638
		goto _test_eof
	_test_eof639:
		cs = 639
		goto _test_eof
	_test_eof640:
		cs = 640
		goto _test_eof
	_test_eof641:
		cs = 641
		goto _test_eof
	_test_eof642:
		cs = 642
		goto _test_eof
	_test_eof643:
		cs = 643
		goto _test_eof
	_test_eof644:
		cs = 644
		goto _test_eof
	_test_eof645:
		cs = 645
		goto _test_eof
	_test_eof646:
		cs = 646
		goto _test_eof
	_test_eof647:
		cs = 647
		goto _test_eof
	_test_eof648:
		cs = 648
		goto _test_eof
	_test_eof649:
		cs = 649
		goto _test_eof
	_test_eof650:
		cs = 650
		goto _test_eof
	_test_eof651:
		cs = 651
		goto _test_eof
	_test_eof652:
		cs = 652
		goto _test_eof
	_test_eof653:
		cs = 653
		goto _test_eof
	_test_eof654:
		cs = 654
		goto _test_eof
	_test_eof655:
		cs = 655
		goto _test_eof
	_test_eof656:
		cs = 656
		goto _test_eof
	_test_eof657:
		cs = 657
		goto _test_eof
	_test_eof658:
		cs = 658
		goto _test_eof
	_test_eof659:
		cs = 659
		goto _test_eof
	_test_eof660:
		cs = 660
		goto _test_eof
	_test_eof661:
		cs = 661
		goto _test_eof
	_test_eof662:
		cs = 662
		goto _test_eof
	_test_eof663:
		cs = 663
		goto _test_eof
	_test_eof664:
		cs = 664
		goto _test_eof
	_test_eof665:
		cs = 665
		goto _test_eof
	_test_eof666:
		cs = 666
		goto _test_eof
	_test_eof667:
		cs = 667
		goto _test_eof
	_test_eof668:
		cs = 668
		goto _test_eof
	_test_eof669:
		cs = 669
		goto _test_eof
	_test_eof670:
		cs = 670
		goto _test_eof
	_test_eof671:
		cs = 671
		goto _test_eof
	_test_eof672:
		cs = 672
		goto _test_eof
	_test_eof673:
		cs = 673
		goto _test_eof
	_test_eof674:
		cs = 674
		goto _test_eof
	_test_eof675:
		cs = 675
		goto _test_eof
	_test_eof676:
		cs = 676
		goto _test_eof
	_test_eof677:
		cs = 677
		goto _test_eof
	_test_eof678:
		cs = 678
		goto _test_eof
	_test_eof679:
		cs = 679
		goto _test_eof
	_test_eof680:
		cs = 680
		goto _test_eof
	_test_eof681:
		cs = 681
		goto _test_eof
	_test_eof682:
		cs = 682
		goto _test_eof
	_test_eof683:
		cs = 683
		goto _test_eof
	_test_eof684:
		cs = 684
		goto _test_eof
	_test_eof685:
		cs = 685
		goto _test_eof
	_test_eof686:
		cs = 686
		goto _test_eof
	_test_eof687:
		cs = 687
		goto _test_eof
	_test_eof688:
		cs = 688
		goto _test_eof
	_test_eof689:
		cs = 689
		goto _test_eof
	_test_eof690:
		cs = 690
		goto _test_eof
	_test_eof691:
		cs = 691
		goto _test_eof
	_test_eof692:
		cs = 692
		goto _test_eof
	_test_eof693:
		cs = 693
		goto _test_eof
	_test_eof694:
		cs = 694
		goto _test_eof
	_test_eof695:
		cs = 695
		goto _test_eof
	_test_eof696:
		cs = 696
		goto _test_eof
	_test_eof697:
		cs = 697
		goto _test_eof
	_test_eof698:
		cs = 698
		goto _test_eof
	_test_eof699:
		cs = 699
		goto _test_eof
	_test_eof700:
		cs = 700
		goto _test_eof
	_test_eof701:
		cs = 701
		goto _test_eof
	_test_eof702:
		cs = 702
		goto _test_eof
	_test_eof703:
		cs = 703
		goto _test_eof
	_test_eof704:
		cs = 704
		goto _test_eof
	_test_eof705:
		cs = 705
		goto _test_eof
	_test_eof706:
		cs = 706
		goto _test_eof
	_test_eof707:
		cs = 707
		goto _test_eof
	_test_eof708:
		cs = 708
		goto _test_eof
	_test_eof709:
		cs = 709
		goto _test_eof
	_test_eof710:
		cs = 710
		goto _test_eof
	_test_eof711:
		cs = 711
		goto _test_eof
	_test_eof712:
		cs = 712
		goto _test_eof
	_test_eof713:
		cs = 713
		goto _test_eof
	_test_eof714:
		cs = 714
		goto _test_eof
	_test_eof715:
		cs = 715
		goto _test_eof
	_test_eof716:
		cs = 716
		goto _test_eof
	_test_eof717:
		cs = 717
		goto _test_eof
	_test_eof718:
		cs = 718
		goto _test_eof
	_test_eof719:
		cs = 719
		goto _test_eof
	_test_eof720:
		cs = 720
		goto _test_eof
	_test_eof721:
		cs = 721
		goto _test_eof
	_test_eof722:
		cs = 722
		goto _test_eof
	_test_eof723:
		cs = 723
		goto _test_eof
	_test_eof724:
		cs = 724
		goto _test_eof
	_test_eof725:
		cs = 725
		goto _test_eof
	_test_eof726:
		cs = 726
		goto _test_eof
	_test_eof727:
		cs = 727
		goto _test_eof
	_test_eof728:
		cs = 728
		goto _test_eof
	_test_eof729:
		cs = 729
		goto _test_eof
	_test_eof730:
		cs = 730
		goto _test_eof
	_test_eof731:
		cs = 731
		goto _test_eof
	_test_eof732:
		cs = 732
		goto _test_eof
	_test_eof733:
		cs = 733
		goto _test_eof
	_test_eof734:
		cs = 734
		goto _test_eof
	_test_eof735:
		cs = 735
		goto _test_eof
	_test_eof736:
		cs = 736
		goto _test_eof
	_test_eof737:
		cs = 737
		goto _test_eof
	_test_eof738:
		cs = 738
		goto _test_eof
	_test_eof739:
		cs = 739
		goto _test_eof
	_test_eof740:
		cs = 740
		goto _test_eof
	_test_eof741:
		cs = 741
		goto _test_eof
	_test_eof742:
		cs = 742
		goto _test_eof
	_test_eof743:
		cs = 743
		goto _test_eof
	_test_eof744:
		cs = 744
		goto _test_eof
	_test_eof745:
		cs = 745
		goto _test_eof
	_test_eof746:
		cs = 746
		goto _test_eof
	_test_eof747:
		cs = 747
		goto _test_eof
	_test_eof748:
		cs = 748
		goto _test_eof
	_test_eof749:
		cs = 749
		goto _test_eof
	_test_eof750:
		cs = 750
		goto _test_eof
	_test_eof751:
		cs = 751
		goto _test_eof
	_test_eof752:
		cs = 752
		goto _test_eof
	_test_eof753:
		cs = 753
		goto _test_eof
	_test_eof754:
		cs = 754
		goto _test_eof
	_test_eof755:
		cs = 755
		goto _test_eof
	_test_eof756:
		cs = 756
		goto _test_eof
	_test_eof757:
		cs = 757
		goto _test_eof
	_test_eof758:
		cs = 758
		goto _test_eof
	_test_eof759:
		cs = 759
		goto _test_eof
	_test_eof760:
		cs = 760
		goto _test_eof
	_test_eof761:
		cs = 761
		goto _test_eof
	_test_eof762:
		cs = 762
		goto _test_eof
	_test_eof763:
		cs = 763
		goto _test_eof
	_test_eof764:
		cs = 764
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 281, 282, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 370, 371, 372, 373, 374, 375, 385, 386, 387, 388, 389, 390, 391, 392, 393, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422, 423, 424, 425, 426, 427, 428, 429, 430, 431, 432, 433, 434, 441, 442, 443, 444, 445, 446, 447, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 481, 482, 489, 490, 491, 492, 493, 494, 495, 496, 497, 498, 499, 500, 501, 502, 503, 504, 505, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 679, 680, 681, 682, 683, 684, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 695, 696, 697, 698, 699, 700, 701, 702, 703, 704, 705, 706, 707, 708, 709, 710, 711, 712, 713, 714, 715, 716, 717, 718, 719, 720, 721, 722, 723, 724, 725, 726, 727, 728, 729, 730, 731, 732, 733, 734, 740, 741, 742, 743, 744, 745, 746, 747, 748, 749, 750, 751, 752, 753, 754, 755, 756, 757, 758, 759, 760, 761, 762, 763:
//line sip.rl:158

				p--

				{
					goto st273
				}

			case 778:
//line sip.rl:201

				*addrp = addr
				addrp = &addr.Next
				addr = nil

//line msg_parse.go:17847
			}
		}

	_out:
		{
		}
	}

//line msg_parse.rl:41

	if cs < msg_first_final {
		if p == pe {
			return nil, MsgIncompleteError{data}
		} else {
			return nil, MsgParseError{Msg: data, Offset: p}
		}
	}

	if clen > 0 {
		if clen != len(data)-p {
			return nil, fmt.Errorf("Content-Length incorrect: %d != %d", clen, len(data)-p)
		}
		msg.Payload = &MiscPayload{T: ctype, D: data[p:]}
	}
	return msg, nil
}

func lookAheadWSP(data []byte, p, pe int) bool {
	return p+2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')
}

func lastAddr(addrp **Addr) **Addr {
	if *addrp == nil {
		return addrp
	} else {
		return lastAddr(&(*addrp).Next)
	}
}
