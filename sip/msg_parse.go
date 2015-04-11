
//line msg_parse.rl:1
// -*-go-*-
//
// Ragel SIP Message Parser
//
// This file is compiled into Go code by the Ragel State Machine Compiler for
// the purpose of converting SIP messages into a Msg data structure. This
// machine works in tandem with the Ragel machine defined in uri_parse.rl.
//
// Perhaps it would have been better if the authors of this protocol had chosen
// to use a binary serialization format like protocol buffers. But instead they
// chose to create a plaintext protocol that looks similar to HTTP requests,
// but are phenomenally more complicated.
//
// SIP messages are quite insane.
//
//   o Whitespace can be used liberally in a variety of different ways.
//
//     - Via host:port can have whitespace, e.g. "host \t:  port"
//
//   o UTF-8 is supported in some places but not others.
//
//   o Headers can span multiple lines.
//
//   o Header values can contain comments, e.g. Message: lol (i'm (hidden))
//
//   o Header names are case-insensitive and have shorthand notation.
//
//   o There's ~50 standard headers, many of which have custom parsing rules.
//
//   o URIs can have ;params;like=this
//
//     - Params can belong either to a URI or Addr object, e.g. <sip:uri;param>
//       cf. <sip:uri>;param
//
//     - Addresses may omit angle brackets, in which case params belong to the
//       Addr object.
//
//     - URI params ;are=escaped%20like%22this but params belonging to Addr
//       ;are="escaped like\"this"
//
//     - Backslash escaping is not like C, e.g. \t\n -> tn
//
//     - Address display name can have whitespace without quotes, which is
//       collapsed. Quoted form is not collapsed.
//
//   o Via and address headers can be repeated in two ways: repeating the
//     header, using commas within a single header, or both.
//
// See: http://www.colm.net/files/ragel/ragel-guide-6.9.pdf
// See: http://zedshaw.com/archive/ragel-state-charts/

package sip

import (
	"errors"
	"fmt"
	"github.com/jart/gosip/sdp"
)


//line msg_parse.rl:61

//line msg_parse.go:66
const msg_start int = 1
const msg_first_final int = 739
const msg_error int = 0

const msg_en_via_param int = 34
const msg_en_via int = 69
const msg_en_addr_param int = 119
const msg_en_addr_angled int = 154
const msg_en_addr_uri int = 195
const msg_en_addr int = 222
const msg_en_value int = 229
const msg_en_xheader int = 239
const msg_en_header int = 246
const msg_en_main int = 1


//line msg_parse.rl:62

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

	
//line msg_parse.go:117
	{
	cs = msg_start
	}

//line msg_parse.go:122
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
	case 739:
		goto st_case_739
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
	case 740:
		goto st_case_740
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
	case 741:
		goto st_case_741
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
	case 742:
		goto st_case_742
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
	case 743:
		goto st_case_743
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
	case 744:
		goto st_case_744
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
	case 745:
		goto st_case_745
	case 213:
		goto st_case_213
	case 214:
		goto st_case_214
	case 215:
		goto st_case_215
	case 746:
		goto st_case_746
	case 216:
		goto st_case_216
	case 217:
		goto st_case_217
	case 747:
		goto st_case_747
	case 218:
		goto st_case_218
	case 219:
		goto st_case_219
	case 220:
		goto st_case_220
	case 748:
		goto st_case_748
	case 749:
		goto st_case_749
	case 750:
		goto st_case_750
	case 751:
		goto st_case_751
	case 221:
		goto st_case_221
	case 222:
		goto st_case_222
	case 223:
		goto st_case_223
	case 224:
		goto st_case_224
	case 752:
		goto st_case_752
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
	case 753:
		goto st_case_753
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
	case 754:
		goto st_case_754
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
	case 755:
		goto st_case_755
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
tr368:
//line msg_parse.rl:219

			p--

			{goto st239 }
		
	goto st0
//line msg_parse.go:1684
st_case_0:
	st0:
		cs = 0
		goto _out
tr0:
//line msg_parse.rl:103

			mark = p
		
	goto st2
	st2:
		p++
	st_case_2:
//line msg_parse.go:1698
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
//line msg_parse.rl:143

			msg.Method = string(data[mark:p])
		
	goto st3
	st3:
		p++
	st_case_3:
//line msg_parse.go:1743
		if data[p] == 32 {
			goto st0
		}
		goto tr5
tr5:
//line msg_parse.rl:103

			mark = p
		
	goto st4
	st4:
		p++
	st_case_4:
//line msg_parse.go:1757
		if data[p] == 32 {
			goto tr7
		}
		goto st4
tr7:
//line msg_parse.rl:155

			msg.Request, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st5
	st5:
		p++
	st_case_5:
//line msg_parse.go:1772
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
//line msg_parse.rl:147

			msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
		
	goto st10
	st10:
		p++
	st_case_10:
//line msg_parse.go:1814
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
//line msg_parse.rl:151

			msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
		
	goto st12
	st12:
		p++
	st_case_12:
//line msg_parse.go:1838
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
//line msg_parse.rl:164

			msg.Phrase = string(buf[0:amt])
		
	goto st13
	st13:
		p++
	st_case_13:
//line msg_parse.go:1862
		if data[p] == 10 {
			goto tr16
		}
		goto st0
tr16:
//line msg_parse.rl:201

			{goto st246 }
		
	goto st739
	st739:
		p++
	st_case_739:
//line msg_parse.go:1876
		goto st0
tr2:
//line msg_parse.rl:103

			mark = p
		
	goto st14
	st14:
		p++
	st_case_14:
//line msg_parse.go:1887
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
//line msg_parse.rl:147

			msg.VersionMajor = msg.VersionMajor * 10 + (data[p] - 0x30)
		
	goto st18
	st18:
		p++
	st_case_18:
//line msg_parse.go:2016
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
//line msg_parse.rl:151

			msg.VersionMinor = msg.VersionMinor * 10 + (data[p] - 0x30)
		
	goto st20
	st20:
		p++
	st_case_20:
//line msg_parse.go:2040
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
//line msg_parse.rl:160

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st22
	st22:
		p++
	st_case_22:
//line msg_parse.go:2064
		if 48 <= data[p] && data[p] <= 57 {
			goto tr25
		}
		goto st0
tr25:
//line msg_parse.rl:160

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st23
	st23:
		p++
	st_case_23:
//line msg_parse.go:2078
		if 48 <= data[p] && data[p] <= 57 {
			goto tr26
		}
		goto st0
tr26:
//line msg_parse.rl:160

			msg.Status = msg.Status * 10 + (int(data[p]) - 0x30)
		
	goto st24
	st24:
		p++
	st_case_24:
//line msg_parse.go:2092
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
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st26
tr35:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st26
tr44:
//line msg_parse.rl:133

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st26
	st26:
		p++
	st_case_26:
//line msg_parse.go:2182
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
//line msg_parse.rl:111

			amt = 0
		
	goto st27
	st27:
		p++
	st_case_27:
//line msg_parse.go:2254
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
//line msg_parse.rl:129

			hex = unhex(data[p]) * 16
		
	goto st28
	st28:
		p++
	st_case_28:
//line msg_parse.go:2277
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
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st29
tr37:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st29
	st29:
		p++
	st_case_29:
//line msg_parse.go:2312
		if 128 <= data[p] && data[p] <= 191 {
			goto tr35
		}
		goto st0
tr31:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st30
tr38:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st30
	st30:
		p++
	st_case_30:
//line msg_parse.go:2338
		if 128 <= data[p] && data[p] <= 191 {
			goto tr37
		}
		goto st0
tr32:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st31
tr39:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st31
	st31:
		p++
	st_case_31:
//line msg_parse.go:2364
		if 128 <= data[p] && data[p] <= 191 {
			goto tr38
		}
		goto st0
tr33:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st32
tr40:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st32
	st32:
		p++
	st_case_32:
//line msg_parse.go:2390
		if 128 <= data[p] && data[p] <= 191 {
			goto tr39
		}
		goto st0
tr34:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st33
tr41:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st33
	st33:
		p++
	st_case_33:
//line msg_parse.go:2416
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
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:103

			mark = p
		
	goto st35
	st35:
		p++
	st_case_35:
//line msg_parse.go:2470
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr46
		case 32:
			goto tr46
		case 33:
			goto st35
		case 37:
			goto st35
		case 39:
			goto st35
		case 44:
			goto tr48
		case 59:
			goto tr49
		case 61:
			goto tr50
		case 126:
			goto st35
		case 269:
			goto tr51
		case 525:
			goto tr52
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto st35
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st35
				}
			case _widec >= 65:
				goto st35
			}
		default:
			goto st35
		}
		goto st0
tr46:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st36
	st36:
		p++
	st_case_36:
//line msg_parse.go:2529
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st36
		case 32:
			goto st36
		case 44:
			goto st37
		case 59:
			goto st41
		case 61:
			goto st45
		case 525:
			goto st66
		}
		goto st0
tr48:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st37
	st37:
		p++
	st_case_37:
//line msg_parse.go:2561
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st37
		case 32:
			goto st37
		case 269:
			goto tr58
		case 525:
			goto st38
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr58
			}
		default:
			goto tr58
		}
		goto st0
tr58:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:194

			if via.Params == nil {
				via.Params = Params{}
			}
			via.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:168

			*viap = via
			viap = &via.Next
			via = nil
		
//line msg_parse.rl:209

			via = new(Via)
			{goto st69 }
		
	goto st740
tr62:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:194

			if via.Params == nil {
				via.Params = Params{}
			}
			via.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:214

			amt = 0  // Needed so ViaParam action works when there's no value.
			{goto st34 }
		
	goto st740
tr74:
//line msg_parse.rl:194

			if via.Params == nil {
				via.Params = Params{}
			}
			via.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:168

			*viap = via
			viap = &via.Next
			via = nil
		
//line msg_parse.rl:201

			{goto st246 }
		
	goto st740
	st740:
		p++
	st_case_740:
//line msg_parse.go:2654
		goto st0
	st38:
		p++
	st_case_38:
		if data[p] == 10 {
			goto st39
		}
		goto st0
	st39:
		p++
	st_case_39:
		switch data[p] {
		case 9:
			goto st40
		case 32:
			goto st40
		}
		goto st0
	st40:
		p++
	st_case_40:
		switch data[p] {
		case 9:
			goto st40
		case 32:
			goto st40
		}
		goto tr58
tr49:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st41
	st41:
		p++
	st_case_41:
//line msg_parse.go:2692
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
		case 269:
			goto tr62
		case 525:
			goto st42
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr62
			}
		default:
			goto tr62
		}
		goto st0
	st42:
		p++
	st_case_42:
		if data[p] == 10 {
			goto st43
		}
		goto st0
	st43:
		p++
	st_case_43:
		switch data[p] {
		case 9:
			goto st44
		case 32:
			goto st44
		}
		goto st0
	st44:
		p++
	st_case_44:
		switch data[p] {
		case 9:
			goto st44
		case 32:
			goto st44
		}
		goto tr62
tr50:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st45
	st45:
		p++
	st_case_45:
//line msg_parse.go:2755
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st45
		case 32:
			goto st45
		case 33:
			goto tr66
		case 34:
			goto st52
		case 37:
			goto tr66
		case 39:
			goto tr66
		case 93:
			goto tr66
		case 126:
			goto tr66
		case 525:
			goto st63
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr66
				}
			case _widec >= 42:
				goto tr66
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr66
				}
			case _widec >= 65:
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
tr66:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st46
	st46:
		p++
	st_case_46:
//line msg_parse.go:2816
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st47
		case 32:
			goto st47
		case 33:
			goto tr66
		case 37:
			goto tr66
		case 39:
			goto tr66
		case 44:
			goto st37
		case 59:
			goto st41
		case 93:
			goto tr66
		case 126:
			goto tr66
		case 269:
			goto st51
		case 525:
			goto st48
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr66
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr66
				}
			case _widec >= 65:
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
	st47:
		p++
	st_case_47:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st47
		case 32:
			goto st47
		case 44:
			goto st37
		case 59:
			goto st41
		case 525:
			goto st48
		}
		goto st0
	st48:
		p++
	st_case_48:
		if data[p] == 10 {
			goto st49
		}
		goto st0
	st49:
		p++
	st_case_49:
		switch data[p] {
		case 9:
			goto st50
		case 32:
			goto st50
		}
		goto st0
	st50:
		p++
	st_case_50:
		switch data[p] {
		case 9:
			goto st50
		case 32:
			goto st50
		case 44:
			goto st37
		case 59:
			goto st41
		}
		goto st0
tr51:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st51
	st51:
		p++
	st_case_51:
//line msg_parse.go:2929
		if data[p] == 10 {
			goto tr74
		}
		goto st0
	st52:
		p++
	st_case_52:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr75
		case 34:
			goto tr76
		case 92:
			goto tr77
		case 525:
			goto tr83
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr78
				}
			case _widec >= 32:
				goto tr75
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr80
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr82
				}
			default:
				goto tr81
			}
		default:
			goto tr79
		}
		goto st0
tr75:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st53
tr84:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st53
	st53:
		p++
	st_case_53:
//line msg_parse.go:3002
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr84
		case 34:
			goto st54
		case 92:
			goto st55
		case 525:
			goto tr92
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr87
				}
			case _widec >= 32:
				goto tr84
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr89
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr91
				}
			default:
				goto tr90
			}
		default:
			goto tr88
		}
		goto st0
tr76:
//line msg_parse.rl:111

			amt = 0
		
	goto st54
	st54:
		p++
	st_case_54:
//line msg_parse.go:3056
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st47
		case 32:
			goto st47
		case 44:
			goto st37
		case 59:
			goto st41
		case 269:
			goto st51
		case 525:
			goto st48
		}
		goto st0
tr77:
//line msg_parse.rl:111

			amt = 0
		
	goto st55
	st55:
		p++
	st_case_55:
//line msg_parse.go:3088
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr84
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr84
			}
		default:
			goto tr84
		}
		goto st0
tr78:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st56
tr87:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st56
	st56:
		p++
	st_case_56:
//line msg_parse.go:3123
		if 128 <= data[p] && data[p] <= 191 {
			goto tr84
		}
		goto st0
tr79:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st57
tr88:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st57
	st57:
		p++
	st_case_57:
//line msg_parse.go:3149
		if 128 <= data[p] && data[p] <= 191 {
			goto tr87
		}
		goto st0
tr80:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st58
tr89:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st58
	st58:
		p++
	st_case_58:
//line msg_parse.go:3175
		if 128 <= data[p] && data[p] <= 191 {
			goto tr88
		}
		goto st0
tr81:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st59
tr90:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st59
	st59:
		p++
	st_case_59:
//line msg_parse.go:3201
		if 128 <= data[p] && data[p] <= 191 {
			goto tr89
		}
		goto st0
tr82:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st60
tr91:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st60
	st60:
		p++
	st_case_60:
//line msg_parse.go:3227
		if 128 <= data[p] && data[p] <= 191 {
			goto tr90
		}
		goto st0
tr83:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st61
tr92:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st61
	st61:
		p++
	st_case_61:
//line msg_parse.go:3253
		if data[p] == 10 {
			goto tr93
		}
		goto st0
tr93:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st62
	st62:
		p++
	st_case_62:
//line msg_parse.go:3268
		switch data[p] {
		case 9:
			goto tr84
		case 32:
			goto tr84
		}
		goto st0
	st63:
		p++
	st_case_63:
		if data[p] == 10 {
			goto st64
		}
		goto st0
	st64:
		p++
	st_case_64:
		switch data[p] {
		case 9:
			goto st65
		case 32:
			goto st65
		}
		goto st0
	st65:
		p++
	st_case_65:
		switch data[p] {
		case 9:
			goto st65
		case 32:
			goto st65
		case 33:
			goto tr66
		case 34:
			goto st52
		case 37:
			goto tr66
		case 39:
			goto tr66
		case 93:
			goto tr66
		case 126:
			goto tr66
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr66
				}
			case data[p] >= 42:
				goto tr66
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr66
				}
			case data[p] >= 65:
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
tr52:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st66
	st66:
		p++
	st_case_66:
//line msg_parse.go:3346
		if data[p] == 10 {
			goto st67
		}
		goto st0
	st67:
		p++
	st_case_67:
		switch data[p] {
		case 9:
			goto st68
		case 32:
			goto st68
		}
		goto st0
	st68:
		p++
	st_case_68:
		switch data[p] {
		case 9:
			goto st68
		case 32:
			goto st68
		case 44:
			goto st37
		case 59:
			goto st41
		case 61:
			goto st45
		}
		goto st0
	st69:
		p++
	st_case_69:
		switch data[p] {
		case 33:
			goto tr98
		case 37:
			goto tr98
		case 39:
			goto tr98
		case 126:
			goto tr98
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr98
				}
			case data[p] >= 42:
				goto tr98
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr98
				}
			case data[p] >= 65:
				goto tr98
			}
		default:
			goto tr98
		}
		goto st0
tr98:
//line msg_parse.rl:103

			mark = p
		
	goto st70
	st70:
		p++
	st_case_70:
//line msg_parse.go:3422
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr99
		case 32:
			goto tr99
		case 33:
			goto st70
		case 37:
			goto st70
		case 39:
			goto st70
		case 47:
			goto tr101
		case 126:
			goto st70
		case 525:
			goto tr102
		}
		switch {
		case _widec < 45:
			if 42 <= _widec && _widec <= 43 {
				goto st70
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st70
				}
			case _widec >= 65:
				goto st70
			}
		default:
			goto st70
		}
		goto st0
tr99:
//line msg_parse.rl:174

			via.Protocol = string(data[mark:p])
		
	goto st71
	st71:
		p++
	st_case_71:
//line msg_parse.go:3475
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
		case 47:
			goto st72
		case 525:
			goto st116
		}
		goto st0
tr101:
//line msg_parse.rl:174

			via.Protocol = string(data[mark:p])
		
	goto st72
	st72:
		p++
	st_case_72:
//line msg_parse.go:3503
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st72
		case 32:
			goto st72
		case 33:
			goto tr106
		case 37:
			goto tr106
		case 39:
			goto tr106
		case 126:
			goto tr106
		case 525:
			goto st113
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr106
				}
			case _widec >= 42:
				goto tr106
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr106
				}
			case _widec >= 65:
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
tr106:
//line msg_parse.rl:103

			mark = p
		
	goto st73
	st73:
		p++
	st_case_73:
//line msg_parse.go:3559
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr108
		case 32:
			goto tr108
		case 33:
			goto st73
		case 37:
			goto st73
		case 39:
			goto st73
		case 47:
			goto tr110
		case 126:
			goto st73
		case 525:
			goto tr111
		}
		switch {
		case _widec < 45:
			if 42 <= _widec && _widec <= 43 {
				goto st73
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st73
				}
			case _widec >= 65:
				goto st73
			}
		default:
			goto st73
		}
		goto st0
tr108:
//line msg_parse.rl:178

			via.Version = string(data[mark:p])
		
	goto st74
	st74:
		p++
	st_case_74:
//line msg_parse.go:3612
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st74
		case 32:
			goto st74
		case 47:
			goto st75
		case 525:
			goto st110
		}
		goto st0
tr110:
//line msg_parse.rl:178

			via.Version = string(data[mark:p])
		
	goto st75
	st75:
		p++
	st_case_75:
//line msg_parse.go:3640
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
		case 33:
			goto tr115
		case 37:
			goto tr115
		case 39:
			goto tr115
		case 126:
			goto tr115
		case 525:
			goto st107
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr115
				}
			case _widec >= 42:
				goto tr115
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr115
				}
			case _widec >= 65:
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
tr115:
//line msg_parse.rl:103

			mark = p
		
	goto st76
	st76:
		p++
	st_case_76:
//line msg_parse.go:3696
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr117
		case 32:
			goto tr117
		case 33:
			goto st76
		case 37:
			goto st76
		case 39:
			goto st76
		case 126:
			goto st76
		case 525:
			goto tr119
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st76
				}
			case _widec >= 42:
				goto st76
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st76
				}
			case _widec >= 65:
				goto st76
			}
		default:
			goto st76
		}
		goto st0
tr117:
//line msg_parse.rl:182

			via.Transport = string(data[mark:p])
		
	goto st77
	st77:
		p++
	st_case_77:
//line msg_parse.go:3752
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st77
		case 32:
			goto st77
		case 91:
			goto st101
		case 525:
			goto st104
		}
		switch {
		case _widec < 48:
			if 45 <= _widec && _widec <= 46 {
				goto tr121
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 97 <= _widec && _widec <= 122 {
					goto tr121
				}
			case _widec >= 65:
				goto tr121
			}
		default:
			goto tr121
		}
		goto st0
tr121:
//line msg_parse.rl:103

			mark = p
		
	goto st78
	st78:
		p++
	st_case_78:
//line msg_parse.go:3797
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr124
		case 32:
			goto tr124
		case 44:
			goto tr125
		case 58:
			goto tr127
		case 59:
			goto tr128
		case 269:
			goto tr129
		case 525:
			goto tr130
		}
		switch {
		case _widec < 48:
			if 45 <= _widec && _widec <= 46 {
				goto st78
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 97 <= _widec && _widec <= 122 {
					goto st78
				}
			case _widec >= 65:
				goto st78
			}
		default:
			goto st78
		}
		goto st0
tr124:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st79
	st79:
		p++
	st_case_79:
//line msg_parse.go:3848
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
		case 44:
			goto st80
		case 58:
			goto st84
		case 59:
			goto st87
		case 525:
			goto st98
		}
		goto st0
tr125:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st80
	st80:
		p++
	st_case_80:
//line msg_parse.go:3880
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st80
		case 32:
			goto st80
		case 269:
			goto tr136
		case 525:
			goto st81
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr136
			}
		default:
			goto tr136
		}
		goto st0
tr136:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:168

			*viap = via
			viap = &via.Next
			via = nil
		
//line msg_parse.rl:209

			via = new(Via)
			{goto st69 }
		
	goto st741
tr145:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:214

			amt = 0  // Needed so ViaParam action works when there's no value.
			{goto st34 }
		
	goto st741
tr151:
//line msg_parse.rl:168

			*viap = via
			viap = &via.Next
			via = nil
		
//line msg_parse.rl:201

			{goto st246 }
		
	goto st741
	st741:
		p++
	st_case_741:
//line msg_parse.go:3952
		goto st0
	st81:
		p++
	st_case_81:
		if data[p] == 10 {
			goto st82
		}
		goto st0
	st82:
		p++
	st_case_82:
		switch data[p] {
		case 9:
			goto st83
		case 32:
			goto st83
		}
		goto st0
	st83:
		p++
	st_case_83:
		switch data[p] {
		case 9:
			goto st83
		case 32:
			goto st83
		}
		goto tr136
tr127:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st84
	st84:
		p++
	st_case_84:
//line msg_parse.go:3990
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st84
		case 32:
			goto st84
		case 525:
			goto st95
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr140
		}
		goto st0
tr140:
//line msg_parse.rl:190

			via.Port = via.Port * 10 + (uint16(data[p]) - 0x30)
		
	goto st85
	st85:
		p++
	st_case_85:
//line msg_parse.go:4019
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st86
		case 32:
			goto st86
		case 44:
			goto st80
		case 59:
			goto st87
		case 269:
			goto st94
		case 525:
			goto st91
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr140
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
			goto st86
		case 32:
			goto st86
		case 44:
			goto st80
		case 59:
			goto st87
		case 525:
			goto st91
		}
		goto st0
tr128:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st87
	st87:
		p++
	st_case_87:
//line msg_parse.go:4077
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st87
		case 32:
			goto st87
		case 269:
			goto tr145
		case 525:
			goto st88
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr145
			}
		default:
			goto tr145
		}
		goto st0
	st88:
		p++
	st_case_88:
		if data[p] == 10 {
			goto st89
		}
		goto st0
	st89:
		p++
	st_case_89:
		switch data[p] {
		case 9:
			goto st90
		case 32:
			goto st90
		}
		goto st0
	st90:
		p++
	st_case_90:
		switch data[p] {
		case 9:
			goto st90
		case 32:
			goto st90
		}
		goto tr145
	st91:
		p++
	st_case_91:
		if data[p] == 10 {
			goto st92
		}
		goto st0
	st92:
		p++
	st_case_92:
		switch data[p] {
		case 9:
			goto st93
		case 32:
			goto st93
		}
		goto st0
	st93:
		p++
	st_case_93:
		switch data[p] {
		case 9:
			goto st93
		case 32:
			goto st93
		case 44:
			goto st80
		case 59:
			goto st87
		}
		goto st0
tr129:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st94
	st94:
		p++
	st_case_94:
//line msg_parse.go:4171
		if data[p] == 10 {
			goto tr151
		}
		goto st0
	st95:
		p++
	st_case_95:
		if data[p] == 10 {
			goto st96
		}
		goto st0
	st96:
		p++
	st_case_96:
		switch data[p] {
		case 9:
			goto st97
		case 32:
			goto st97
		}
		goto st0
	st97:
		p++
	st_case_97:
		switch data[p] {
		case 9:
			goto st97
		case 32:
			goto st97
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr140
		}
		goto st0
tr130:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st98
	st98:
		p++
	st_case_98:
//line msg_parse.go:4215
		if data[p] == 10 {
			goto st99
		}
		goto st0
	st99:
		p++
	st_case_99:
		switch data[p] {
		case 9:
			goto st100
		case 32:
			goto st100
		}
		goto st0
	st100:
		p++
	st_case_100:
		switch data[p] {
		case 9:
			goto st100
		case 32:
			goto st100
		case 44:
			goto st80
		case 58:
			goto st84
		case 59:
			goto st87
		}
		goto st0
	st101:
		p++
	st_case_101:
		if data[p] == 46 {
			goto tr156
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto tr156
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr156
			}
		default:
			goto tr156
		}
		goto st0
tr156:
//line msg_parse.rl:103

			mark = p
		
	goto st102
	st102:
		p++
	st_case_102:
//line msg_parse.go:4274
		switch data[p] {
		case 46:
			goto st102
		case 93:
			goto tr158
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st102
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st102
			}
		default:
			goto st102
		}
		goto st0
tr158:
//line msg_parse.rl:186

			via.Host = string(data[mark:p])
		
	goto st103
	st103:
		p++
	st_case_103:
//line msg_parse.go:4303
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
		case 44:
			goto st80
		case 58:
			goto st84
		case 59:
			goto st87
		case 269:
			goto st94
		case 525:
			goto st98
		}
		goto st0
tr119:
//line msg_parse.rl:182

			via.Transport = string(data[mark:p])
		
	goto st104
	st104:
		p++
	st_case_104:
//line msg_parse.go:4337
		if data[p] == 10 {
			goto st105
		}
		goto st0
	st105:
		p++
	st_case_105:
		switch data[p] {
		case 9:
			goto st106
		case 32:
			goto st106
		}
		goto st0
	st106:
		p++
	st_case_106:
		switch data[p] {
		case 9:
			goto st106
		case 32:
			goto st106
		case 91:
			goto st101
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr121
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr121
				}
			case data[p] >= 65:
				goto tr121
			}
		default:
			goto tr121
		}
		goto st0
	st107:
		p++
	st_case_107:
		if data[p] == 10 {
			goto st108
		}
		goto st0
	st108:
		p++
	st_case_108:
		switch data[p] {
		case 9:
			goto st109
		case 32:
			goto st109
		}
		goto st0
	st109:
		p++
	st_case_109:
		switch data[p] {
		case 9:
			goto st109
		case 32:
			goto st109
		case 33:
			goto tr115
		case 37:
			goto tr115
		case 39:
			goto tr115
		case 126:
			goto tr115
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr115
				}
			case data[p] >= 42:
				goto tr115
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr115
				}
			case data[p] >= 65:
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
tr111:
//line msg_parse.rl:178

			via.Version = string(data[mark:p])
		
	goto st110
	st110:
		p++
	st_case_110:
//line msg_parse.go:4447
		if data[p] == 10 {
			goto st111
		}
		goto st0
	st111:
		p++
	st_case_111:
		switch data[p] {
		case 9:
			goto st112
		case 32:
			goto st112
		}
		goto st0
	st112:
		p++
	st_case_112:
		switch data[p] {
		case 9:
			goto st112
		case 32:
			goto st112
		case 47:
			goto st75
		}
		goto st0
	st113:
		p++
	st_case_113:
		if data[p] == 10 {
			goto st114
		}
		goto st0
	st114:
		p++
	st_case_114:
		switch data[p] {
		case 9:
			goto st115
		case 32:
			goto st115
		}
		goto st0
	st115:
		p++
	st_case_115:
		switch data[p] {
		case 9:
			goto st115
		case 32:
			goto st115
		case 33:
			goto tr106
		case 37:
			goto tr106
		case 39:
			goto tr106
		case 126:
			goto tr106
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr106
				}
			case data[p] >= 42:
				goto tr106
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr106
				}
			case data[p] >= 65:
				goto tr106
			}
		default:
			goto tr106
		}
		goto st0
tr102:
//line msg_parse.rl:174

			via.Protocol = string(data[mark:p])
		
	goto st116
	st116:
		p++
	st_case_116:
//line msg_parse.go:4540
		if data[p] == 10 {
			goto st117
		}
		goto st0
	st117:
		p++
	st_case_117:
		switch data[p] {
		case 9:
			goto st118
		case 32:
			goto st118
		}
		goto st0
	st118:
		p++
	st_case_118:
		switch data[p] {
		case 9:
			goto st118
		case 32:
			goto st118
		case 47:
			goto st72
		}
		goto st0
	st119:
		p++
	st_case_119:
		switch data[p] {
		case 33:
			goto tr169
		case 37:
			goto tr169
		case 39:
			goto tr169
		case 126:
			goto tr169
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr169
				}
			case data[p] >= 42:
				goto tr169
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr169
				}
			case data[p] >= 65:
				goto tr169
			}
		default:
			goto tr169
		}
		goto st0
tr169:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:103

			mark = p
		
	goto st120
	st120:
		p++
	st_case_120:
//line msg_parse.go:4616
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr170
		case 32:
			goto tr170
		case 33:
			goto st120
		case 37:
			goto st120
		case 39:
			goto st120
		case 44:
			goto tr172
		case 59:
			goto tr173
		case 61:
			goto tr174
		case 126:
			goto st120
		case 269:
			goto tr175
		case 525:
			goto tr176
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto st120
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st120
				}
			case _widec >= 65:
				goto st120
			}
		default:
			goto st120
		}
		goto st0
tr170:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st121
	st121:
		p++
	st_case_121:
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
			goto st121
		case 32:
			goto st121
		case 44:
			goto st122
		case 59:
			goto st126
		case 61:
			goto st130
		case 525:
			goto st151
		}
		goto st0
tr172:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st122
	st122:
		p++
	st_case_122:
//line msg_parse.go:4707
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st122
		case 32:
			goto st122
		case 269:
			goto tr182
		case 525:
			goto st123
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr182
			}
		default:
			goto tr182
		}
		goto st0
tr182:
//line msg_parse.rl:277

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st742
tr186:
//line msg_parse.rl:277

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st742
tr198:
//line msg_parse.rl:277

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:201

			{goto st246 }
		
	goto st742
	st742:
		p++
	st_case_742:
//line msg_parse.go:4798
		goto st0
	st123:
		p++
	st_case_123:
		if data[p] == 10 {
			goto st124
		}
		goto st0
	st124:
		p++
	st_case_124:
		switch data[p] {
		case 9:
			goto st125
		case 32:
			goto st125
		}
		goto st0
	st125:
		p++
	st_case_125:
		switch data[p] {
		case 9:
			goto st125
		case 32:
			goto st125
		}
		goto tr182
tr173:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st126
	st126:
		p++
	st_case_126:
//line msg_parse.go:4836
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st126
		case 32:
			goto st126
		case 269:
			goto tr186
		case 525:
			goto st127
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr186
			}
		default:
			goto tr186
		}
		goto st0
	st127:
		p++
	st_case_127:
		if data[p] == 10 {
			goto st128
		}
		goto st0
	st128:
		p++
	st_case_128:
		switch data[p] {
		case 9:
			goto st129
		case 32:
			goto st129
		}
		goto st0
	st129:
		p++
	st_case_129:
		switch data[p] {
		case 9:
			goto st129
		case 32:
			goto st129
		}
		goto tr186
tr174:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st130
	st130:
		p++
	st_case_130:
//line msg_parse.go:4899
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st130
		case 32:
			goto st130
		case 33:
			goto tr190
		case 34:
			goto st137
		case 37:
			goto tr190
		case 39:
			goto tr190
		case 93:
			goto tr190
		case 126:
			goto tr190
		case 525:
			goto st148
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr190
				}
			case _widec >= 42:
				goto tr190
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr190
				}
			case _widec >= 65:
				goto tr190
			}
		default:
			goto tr190
		}
		goto st0
tr190:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st131
	st131:
		p++
	st_case_131:
//line msg_parse.go:4960
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st132
		case 32:
			goto st132
		case 33:
			goto tr190
		case 37:
			goto tr190
		case 39:
			goto tr190
		case 44:
			goto st122
		case 59:
			goto st126
		case 93:
			goto tr190
		case 126:
			goto tr190
		case 269:
			goto st136
		case 525:
			goto st133
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr190
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr190
				}
			case _widec >= 65:
				goto tr190
			}
		default:
			goto tr190
		}
		goto st0
	st132:
		p++
	st_case_132:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st132
		case 32:
			goto st132
		case 44:
			goto st122
		case 59:
			goto st126
		case 525:
			goto st133
		}
		goto st0
	st133:
		p++
	st_case_133:
		if data[p] == 10 {
			goto st134
		}
		goto st0
	st134:
		p++
	st_case_134:
		switch data[p] {
		case 9:
			goto st135
		case 32:
			goto st135
		}
		goto st0
	st135:
		p++
	st_case_135:
		switch data[p] {
		case 9:
			goto st135
		case 32:
			goto st135
		case 44:
			goto st122
		case 59:
			goto st126
		}
		goto st0
tr175:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st136
	st136:
		p++
	st_case_136:
//line msg_parse.go:5073
		if data[p] == 10 {
			goto tr198
		}
		goto st0
	st137:
		p++
	st_case_137:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr199
		case 34:
			goto tr200
		case 92:
			goto tr201
		case 525:
			goto tr207
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr202
				}
			case _widec >= 32:
				goto tr199
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr204
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr206
				}
			default:
				goto tr205
			}
		default:
			goto tr203
		}
		goto st0
tr199:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st138
tr208:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st138
	st138:
		p++
	st_case_138:
//line msg_parse.go:5146
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr208
		case 34:
			goto st139
		case 92:
			goto st140
		case 525:
			goto tr216
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr211
				}
			case _widec >= 32:
				goto tr208
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr213
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr215
				}
			default:
				goto tr214
			}
		default:
			goto tr212
		}
		goto st0
tr200:
//line msg_parse.rl:111

			amt = 0
		
	goto st139
	st139:
		p++
	st_case_139:
//line msg_parse.go:5200
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st132
		case 32:
			goto st132
		case 44:
			goto st122
		case 59:
			goto st126
		case 269:
			goto st136
		case 525:
			goto st133
		}
		goto st0
tr201:
//line msg_parse.rl:111

			amt = 0
		
	goto st140
	st140:
		p++
	st_case_140:
//line msg_parse.go:5232
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr208
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr208
			}
		default:
			goto tr208
		}
		goto st0
tr202:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st141
tr211:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st141
	st141:
		p++
	st_case_141:
//line msg_parse.go:5267
		if 128 <= data[p] && data[p] <= 191 {
			goto tr208
		}
		goto st0
tr203:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st142
tr212:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st142
	st142:
		p++
	st_case_142:
//line msg_parse.go:5293
		if 128 <= data[p] && data[p] <= 191 {
			goto tr211
		}
		goto st0
tr204:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st143
tr213:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st143
	st143:
		p++
	st_case_143:
//line msg_parse.go:5319
		if 128 <= data[p] && data[p] <= 191 {
			goto tr212
		}
		goto st0
tr205:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st144
tr214:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st144
	st144:
		p++
	st_case_144:
//line msg_parse.go:5345
		if 128 <= data[p] && data[p] <= 191 {
			goto tr213
		}
		goto st0
tr206:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st145
tr215:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st145
	st145:
		p++
	st_case_145:
//line msg_parse.go:5371
		if 128 <= data[p] && data[p] <= 191 {
			goto tr214
		}
		goto st0
tr207:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st146
tr216:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st146
	st146:
		p++
	st_case_146:
//line msg_parse.go:5397
		if data[p] == 10 {
			goto tr217
		}
		goto st0
tr217:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st147
	st147:
		p++
	st_case_147:
//line msg_parse.go:5412
		switch data[p] {
		case 9:
			goto tr208
		case 32:
			goto tr208
		}
		goto st0
	st148:
		p++
	st_case_148:
		if data[p] == 10 {
			goto st149
		}
		goto st0
	st149:
		p++
	st_case_149:
		switch data[p] {
		case 9:
			goto st150
		case 32:
			goto st150
		}
		goto st0
	st150:
		p++
	st_case_150:
		switch data[p] {
		case 9:
			goto st150
		case 32:
			goto st150
		case 33:
			goto tr190
		case 34:
			goto st137
		case 37:
			goto tr190
		case 39:
			goto tr190
		case 93:
			goto tr190
		case 126:
			goto tr190
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr190
				}
			case data[p] >= 42:
				goto tr190
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr190
				}
			case data[p] >= 65:
				goto tr190
			}
		default:
			goto tr190
		}
		goto st0
tr176:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st151
	st151:
		p++
	st_case_151:
//line msg_parse.go:5490
		if data[p] == 10 {
			goto st152
		}
		goto st0
	st152:
		p++
	st_case_152:
		switch data[p] {
		case 9:
			goto st153
		case 32:
			goto st153
		}
		goto st0
	st153:
		p++
	st_case_153:
		switch data[p] {
		case 9:
			goto st153
		case 32:
			goto st153
		case 44:
			goto st122
		case 59:
			goto st126
		case 61:
			goto st130
		}
		goto st0
	st154:
		p++
	st_case_154:
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
		case 33:
			goto tr223
		case 34:
			goto tr224
		case 37:
			goto tr223
		case 39:
			goto tr223
		case 60:
			goto st156
		case 126:
			goto tr223
		case 525:
			goto st176
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr223
				}
			case _widec >= 42:
				goto tr223
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr223
				}
			case _widec >= 65:
				goto tr223
			}
		default:
			goto tr223
		}
		goto st0
tr281:
//line msg_parse.rl:260

			addr.Display = string(buf[0:amt])
		
	goto st155
	st155:
		p++
	st_case_155:
//line msg_parse.go:5583
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
		case 60:
			goto st156
		case 525:
			goto st176
		}
		goto st0
tr257:
//line msg_parse.rl:264
{
			end := p
			for end > mark && whitespacec(data[end - 1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
	goto st156
tr282:
//line msg_parse.rl:260

			addr.Display = string(buf[0:amt])
		
	goto st156
	st156:
		p++
	st_case_156:
//line msg_parse.go:5621
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr227
			}
		case data[p] >= 65:
			goto tr227
		}
		goto st0
tr227:
//line msg_parse.rl:103

			mark = p
		
	goto st157
	st157:
		p++
	st_case_157:
//line msg_parse.go:5640
		switch data[p] {
		case 43:
			goto st157
		case 58:
			goto st158
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st157
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st157
				}
			case data[p] >= 65:
				goto st157
			}
		default:
			goto st157
		}
		goto st0
	st158:
		p++
	st_case_158:
		switch data[p] {
		case 33:
			goto st159
		case 61:
			goto st159
		case 93:
			goto st159
		case 95:
			goto st159
		case 126:
			goto st159
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st159
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st159
			}
		default:
			goto st159
		}
		goto st0
	st159:
		p++
	st_case_159:
		switch data[p] {
		case 33:
			goto st159
		case 62:
			goto tr231
		case 93:
			goto st159
		case 95:
			goto st159
		case 126:
			goto st159
		}
		switch {
		case data[p] < 61:
			if 36 <= data[p] && data[p] <= 59 {
				goto st159
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st159
			}
		default:
			goto st159
		}
		goto st0
tr231:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st160
	st160:
		p++
	st_case_160:
//line msg_parse.go:5731
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
		case 44:
			goto st161
		case 59:
			goto st165
		case 269:
			goto st169
		case 525:
			goto st170
		}
		goto st0
	st161:
		p++
	st_case_161:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st161
		case 32:
			goto st161
		case 269:
			goto tr237
		case 525:
			goto st162
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr237
			}
		default:
			goto tr237
		}
		goto st0
tr237:
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st743
tr241:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st743
tr245:
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:201

			{goto st246 }
		
	goto st743
	st743:
		p++
	st_case_743:
//line msg_parse.go:5826
		goto st0
	st162:
		p++
	st_case_162:
		if data[p] == 10 {
			goto st163
		}
		goto st0
	st163:
		p++
	st_case_163:
		switch data[p] {
		case 9:
			goto st164
		case 32:
			goto st164
		}
		goto st0
	st164:
		p++
	st_case_164:
		switch data[p] {
		case 9:
			goto st164
		case 32:
			goto st164
		}
		goto tr237
	st165:
		p++
	st_case_165:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st165
		case 32:
			goto st165
		case 269:
			goto tr241
		case 525:
			goto st166
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr241
			}
		default:
			goto tr241
		}
		goto st0
	st166:
		p++
	st_case_166:
		if data[p] == 10 {
			goto st167
		}
		goto st0
	st167:
		p++
	st_case_167:
		switch data[p] {
		case 9:
			goto st168
		case 32:
			goto st168
		}
		goto st0
	st168:
		p++
	st_case_168:
		switch data[p] {
		case 9:
			goto st168
		case 32:
			goto st168
		}
		goto tr241
	st169:
		p++
	st_case_169:
		if data[p] == 10 {
			goto tr245
		}
		goto st0
	st170:
		p++
	st_case_170:
		if data[p] == 10 {
			goto st171
		}
		goto st0
	st171:
		p++
	st_case_171:
		switch data[p] {
		case 9:
			goto st172
		case 32:
			goto st172
		}
		goto st0
	st172:
		p++
	st_case_172:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st172
		case 32:
			goto st172
		case 44:
			goto st161
		case 59:
			goto st165
		case 269:
			goto st169
		case 525:
			goto st173
		}
		goto st0
	st173:
		p++
	st_case_173:
		if data[p] == 10 {
			goto st174
		}
		goto st0
	st174:
		p++
	st_case_174:
		switch data[p] {
		case 9:
			goto st175
		case 32:
			goto st175
		}
		goto st0
	st175:
		p++
	st_case_175:
		switch data[p] {
		case 9:
			goto st175
		case 32:
			goto st175
		case 44:
			goto st161
		case 59:
			goto st165
		}
		goto st0
tr262:
//line msg_parse.rl:264
{
			end := p
			for end > mark && whitespacec(data[end - 1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
	goto st176
tr283:
//line msg_parse.rl:260

			addr.Display = string(buf[0:amt])
		
	goto st176
	st176:
		p++
	st_case_176:
//line msg_parse.go:6010
		if data[p] == 10 {
			goto st177
		}
		goto st0
	st177:
		p++
	st_case_177:
		switch data[p] {
		case 9:
			goto st178
		case 32:
			goto st178
		}
		goto st0
	st178:
		p++
	st_case_178:
		switch data[p] {
		case 9:
			goto st178
		case 32:
			goto st178
		case 60:
			goto st156
		}
		goto st0
tr223:
//line msg_parse.rl:103

			mark = p
		
	goto st179
	st179:
		p++
	st_case_179:
//line msg_parse.go:6046
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st180
		case 32:
			goto st180
		case 33:
			goto st179
		case 37:
			goto st179
		case 39:
			goto st179
		case 126:
			goto st179
		case 525:
			goto st181
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st179
				}
			case _widec >= 42:
				goto st179
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st179
				}
			case _widec >= 65:
				goto st179
			}
		default:
			goto st179
		}
		goto st0
tr256:
//line msg_parse.rl:264
{
			end := p
			for end > mark && whitespacec(data[end - 1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
	goto st180
	st180:
		p++
	st_case_180:
//line msg_parse.go:6106
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
		case 32:
			goto tr256
		case 33:
			goto st179
		case 37:
			goto st179
		case 39:
			goto st179
		case 60:
			goto tr257
		case 126:
			goto st179
		case 525:
			goto tr258
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st179
				}
			case _widec >= 42:
				goto st179
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st179
				}
			case _widec >= 65:
				goto st179
			}
		default:
			goto st179
		}
		goto st0
tr258:
//line msg_parse.rl:264
{
			end := p
			for end > mark && whitespacec(data[end - 1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
	goto st181
	st181:
		p++
	st_case_181:
//line msg_parse.go:6168
		if data[p] == 10 {
			goto st182
		}
		goto st0
	st182:
		p++
	st_case_182:
		switch data[p] {
		case 9:
			goto st183
		case 32:
			goto st183
		}
		goto st0
tr261:
//line msg_parse.rl:264
{
			end := p
			for end > mark && whitespacec(data[end - 1]) {
				end--
			}
			addr.Display = string(data[mark:end])
		}
	goto st183
	st183:
		p++
	st_case_183:
//line msg_parse.go:6196
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr261
		case 32:
			goto tr261
		case 33:
			goto st179
		case 37:
			goto st179
		case 39:
			goto st179
		case 60:
			goto tr257
		case 126:
			goto st179
		case 525:
			goto tr262
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st179
				}
			case _widec >= 42:
				goto st179
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st179
				}
			case _widec >= 65:
				goto st179
			}
		default:
			goto st179
		}
		goto st0
tr224:
//line msg_parse.rl:111

			amt = 0
		
	goto st184
	st184:
		p++
	st_case_184:
//line msg_parse.go:6254
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr263
		case 34:
			goto tr264
		case 92:
			goto tr265
		case 525:
			goto tr271
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr266
				}
			case _widec >= 32:
				goto tr263
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr268
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr270
				}
			default:
				goto tr269
			}
		default:
			goto tr267
		}
		goto st0
tr263:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st185
tr272:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st185
	st185:
		p++
	st_case_185:
//line msg_parse.go:6320
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr272
		case 34:
			goto st186
		case 92:
			goto st187
		case 525:
			goto tr280
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr275
				}
			case _widec >= 32:
				goto tr272
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr277
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr279
				}
			default:
				goto tr278
			}
		default:
			goto tr276
		}
		goto st0
tr264:
//line msg_parse.rl:111

			amt = 0
		
	goto st186
	st186:
		p++
	st_case_186:
//line msg_parse.go:6374
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr281
		case 32:
			goto tr281
		case 60:
			goto tr282
		case 525:
			goto tr283
		}
		goto st0
tr265:
//line msg_parse.rl:111

			amt = 0
		
	goto st187
	st187:
		p++
	st_case_187:
//line msg_parse.go:6402
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr272
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr272
			}
		default:
			goto tr272
		}
		goto st0
tr266:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st188
tr275:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st188
	st188:
		p++
	st_case_188:
//line msg_parse.go:6437
		if 128 <= data[p] && data[p] <= 191 {
			goto tr272
		}
		goto st0
tr267:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st189
tr276:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st189
	st189:
		p++
	st_case_189:
//line msg_parse.go:6463
		if 128 <= data[p] && data[p] <= 191 {
			goto tr275
		}
		goto st0
tr268:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st190
tr277:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st190
	st190:
		p++
	st_case_190:
//line msg_parse.go:6489
		if 128 <= data[p] && data[p] <= 191 {
			goto tr276
		}
		goto st0
tr269:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st191
tr278:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st191
	st191:
		p++
	st_case_191:
//line msg_parse.go:6515
		if 128 <= data[p] && data[p] <= 191 {
			goto tr277
		}
		goto st0
tr270:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st192
tr279:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st192
	st192:
		p++
	st_case_192:
//line msg_parse.go:6541
		if 128 <= data[p] && data[p] <= 191 {
			goto tr278
		}
		goto st0
tr271:
//line msg_parse.rl:111

			amt = 0
		
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st193
tr280:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st193
	st193:
		p++
	st_case_193:
//line msg_parse.go:6567
		if data[p] == 10 {
			goto tr284
		}
		goto st0
tr284:
//line msg_parse.rl:115

			buf[amt] = data[p]
			amt++
		
	goto st194
	st194:
		p++
	st_case_194:
//line msg_parse.go:6582
		switch data[p] {
		case 9:
			goto tr272
		case 32:
			goto tr272
		}
		goto st0
	st195:
		p++
	st_case_195:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st196
			}
		case data[p] >= 65:
			goto st196
		}
		goto st0
	st196:
		p++
	st_case_196:
		switch data[p] {
		case 43:
			goto st196
		case 58:
			goto st197
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st196
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st196
				}
			case data[p] >= 65:
				goto st196
			}
		default:
			goto st196
		}
		goto st0
	st197:
		p++
	st_case_197:
		switch data[p] {
		case 33:
			goto st198
		case 61:
			goto st198
		case 93:
			goto st198
		case 95:
			goto st198
		case 126:
			goto st198
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st198
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st0
	st198:
		p++
	st_case_198:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr288
		case 32:
			goto tr288
		case 33:
			goto st198
		case 44:
			goto tr289
		case 59:
			goto tr290
		case 61:
			goto st198
		case 93:
			goto st198
		case 95:
			goto st198
		case 126:
			goto st198
		case 269:
			goto tr291
		case 525:
			goto tr292
		}
		switch {
		case _widec < 63:
			if 36 <= _widec && _widec <= 58 {
				goto st198
			}
		case _widec > 91:
			if 97 <= _widec && _widec <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st0
tr288:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st199
	st199:
		p++
	st_case_199:
//line msg_parse.go:6714
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
		case 44:
			goto st200
		case 59:
			goto st204
		case 525:
			goto st208
		}
		goto st0
	st200:
		p++
	st_case_200:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st200
		case 32:
			goto st200
		case 269:
			goto tr297
		case 525:
			goto st201
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr297
			}
		default:
			goto tr297
		}
		goto st0
tr297:
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st744
tr301:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st744
	st744:
		p++
	st_case_744:
//line msg_parse.go:6795
		goto st0
	st201:
		p++
	st_case_201:
		if data[p] == 10 {
			goto st202
		}
		goto st0
	st202:
		p++
	st_case_202:
		switch data[p] {
		case 9:
			goto st203
		case 32:
			goto st203
		}
		goto st0
	st203:
		p++
	st_case_203:
		switch data[p] {
		case 9:
			goto st203
		case 32:
			goto st203
		}
		goto tr297
	st204:
		p++
	st_case_204:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st204
		case 32:
			goto st204
		case 269:
			goto tr301
		case 525:
			goto st205
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr301
			}
		default:
			goto tr301
		}
		goto st0
	st205:
		p++
	st_case_205:
		if data[p] == 10 {
			goto st206
		}
		goto st0
	st206:
		p++
	st_case_206:
		switch data[p] {
		case 9:
			goto st207
		case 32:
			goto st207
		}
		goto st0
	st207:
		p++
	st_case_207:
		switch data[p] {
		case 9:
			goto st207
		case 32:
			goto st207
		}
		goto tr301
tr292:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st208
	st208:
		p++
	st_case_208:
//line msg_parse.go:6890
		if data[p] == 10 {
			goto st209
		}
		goto st0
	st209:
		p++
	st_case_209:
		switch data[p] {
		case 9:
			goto st210
		case 32:
			goto st210
		}
		goto st0
	st210:
		p++
	st_case_210:
		switch data[p] {
		case 9:
			goto st210
		case 32:
			goto st210
		case 44:
			goto st200
		case 59:
			goto st204
		}
		goto st0
tr289:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st211
	st211:
		p++
	st_case_211:
//line msg_parse.go:6929
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr307
		case 32:
			goto tr307
		case 33:
			goto tr308
		case 44:
			goto tr289
		case 59:
			goto tr309
		case 60:
			goto tr297
		case 62:
			goto tr297
		case 92:
			goto tr297
		case 94:
			goto tr297
		case 96:
			goto tr297
		case 126:
			goto tr308
		case 269:
			goto tr310
		case 525:
			goto tr311
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr297
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr297
				}
			case _widec >= 36:
				goto tr308
			}
		default:
			goto tr297
		}
		goto st0
tr307:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st212
	st212:
		p++
	st_case_212:
//line msg_parse.go:6993
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st212
		case 32:
			goto st212
		case 44:
			goto st200
		case 59:
			goto tr313
		case 269:
			goto tr297
		case 525:
			goto st213
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr297
			}
		default:
			goto tr297
		}
		goto st0
tr313:
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st745
	st745:
		p++
	st_case_745:
//line msg_parse.go:7044
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st204
		case 32:
			goto st204
		case 269:
			goto tr301
		case 525:
			goto st205
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr301
			}
		default:
			goto tr301
		}
		goto st0
tr311:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st213
	st213:
		p++
	st_case_213:
//line msg_parse.go:7081
		if data[p] == 10 {
			goto st214
		}
		goto st0
	st214:
		p++
	st_case_214:
		switch data[p] {
		case 9:
			goto st215
		case 32:
			goto st215
		}
		goto st0
	st215:
		p++
	st_case_215:
		switch data[p] {
		case 9:
			goto st215
		case 32:
			goto st215
		case 44:
			goto st200
		case 59:
			goto tr313
		}
		goto tr297
tr308:
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st746
tr318:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st746
	st746:
		p++
	st_case_746:
//line msg_parse.go:7141
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr288
		case 32:
			goto tr288
		case 33:
			goto st198
		case 44:
			goto tr289
		case 59:
			goto tr290
		case 61:
			goto st198
		case 93:
			goto st198
		case 95:
			goto st198
		case 126:
			goto st198
		case 269:
			goto tr291
		case 525:
			goto tr292
		}
		switch {
		case _widec < 63:
			if 36 <= _widec && _widec <= 58 {
				goto st198
			}
		case _widec > 91:
			if 97 <= _widec && _widec <= 122 {
				goto st198
			}
		default:
			goto st198
		}
		goto st0
tr290:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st216
	st216:
		p++
	st_case_216:
//line msg_parse.go:7196
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr317
		case 32:
			goto tr317
		case 33:
			goto tr318
		case 44:
			goto tr319
		case 59:
			goto tr290
		case 60:
			goto tr301
		case 62:
			goto tr301
		case 92:
			goto tr301
		case 94:
			goto tr301
		case 96:
			goto tr301
		case 126:
			goto tr318
		case 269:
			goto tr320
		case 525:
			goto tr321
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr301
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr301
				}
			case _widec >= 36:
				goto tr318
			}
		default:
			goto tr301
		}
		goto st0
tr317:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st217
	st217:
		p++
	st_case_217:
//line msg_parse.go:7260
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st217
		case 32:
			goto st217
		case 44:
			goto tr323
		case 59:
			goto st204
		case 269:
			goto tr301
		case 525:
			goto st218
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr301
			}
		default:
			goto tr301
		}
		goto st0
tr323:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st747
	st747:
		p++
	st_case_747:
//line msg_parse.go:7305
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st200
		case 32:
			goto st200
		case 269:
			goto tr297
		case 525:
			goto st201
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr297
			}
		default:
			goto tr297
		}
		goto st0
tr321:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st218
	st218:
		p++
	st_case_218:
//line msg_parse.go:7342
		if data[p] == 10 {
			goto st219
		}
		goto st0
	st219:
		p++
	st_case_219:
		switch data[p] {
		case 9:
			goto st220
		case 32:
			goto st220
		}
		goto st0
	st220:
		p++
	st_case_220:
		switch data[p] {
		case 9:
			goto st220
		case 32:
			goto st220
		case 44:
			goto tr323
		case 59:
			goto st204
		}
		goto tr301
tr319:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st748
	st748:
		p++
	st_case_748:
//line msg_parse.go:7390
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr307
		case 32:
			goto tr307
		case 33:
			goto tr308
		case 44:
			goto tr289
		case 59:
			goto tr309
		case 60:
			goto tr297
		case 62:
			goto tr297
		case 92:
			goto tr297
		case 94:
			goto tr297
		case 96:
			goto tr297
		case 126:
			goto tr308
		case 269:
			goto tr310
		case 525:
			goto tr311
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr297
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr297
				}
			case _widec >= 36:
				goto tr308
			}
		default:
			goto tr297
		}
		goto st0
tr309:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st749
	st749:
		p++
	st_case_749:
//line msg_parse.go:7469
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr317
		case 32:
			goto tr317
		case 33:
			goto tr318
		case 44:
			goto tr319
		case 59:
			goto tr290
		case 60:
			goto tr301
		case 62:
			goto tr301
		case 92:
			goto tr301
		case 94:
			goto tr301
		case 96:
			goto tr301
		case 126:
			goto tr318
		case 269:
			goto tr320
		case 525:
			goto tr321
		}
		switch {
		case _widec < 14:
			if _widec <= 12 {
				goto tr301
			}
		case _widec > 35:
			switch {
			case _widec > 122:
				if 123 <= _widec {
					goto tr301
				}
			case _widec >= 36:
				goto tr318
			}
		default:
			goto tr301
		}
		goto st0
tr310:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st750
tr320:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:252

			{goto st119 }
		
	goto st750
	st750:
		p++
	st_case_750:
//line msg_parse.go:7564
		if data[p] == 10 {
			goto tr327
		}
		goto st0
tr327:
//line msg_parse.rl:201

			{goto st246 }
		
	goto st751
	st751:
		p++
	st_case_751:
//line msg_parse.go:7578
		goto st0
tr291:
//line msg_parse.rl:272

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st221
	st221:
		p++
	st_case_221:
//line msg_parse.go:7590
		if data[p] == 10 {
			goto tr327
		}
		goto st0
	st222:
		p++
	st_case_222:
		switch data[p] {
		case 33:
			goto tr328
		case 34:
			goto tr329
		case 37:
			goto tr328
		case 39:
			goto tr328
		case 60:
			goto tr329
		case 126:
			goto tr328
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr328
				}
			case data[p] >= 42:
				goto tr328
			}
		case data[p] > 57:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr330
				}
			case data[p] > 96:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr330
				}
			default:
				goto tr328
			}
		default:
			goto tr328
		}
		goto st0
tr328:
//line msg_parse.rl:103

			mark = p
		
	goto st223
	st223:
		p++
	st_case_223:
//line msg_parse.go:7648
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st224
		case 32:
			goto st224
		case 33:
			goto st223
		case 37:
			goto st223
		case 39:
			goto st223
		case 126:
			goto st223
		case 525:
			goto st225
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st223
				}
			case _widec >= 42:
				goto st223
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st223
				}
			case _widec >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st0
	st224:
		p++
	st_case_224:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st224
		case 32:
			goto st224
		case 33:
			goto st223
		case 37:
			goto st223
		case 39:
			goto st223
		case 60:
			goto tr334
		case 126:
			goto st223
		case 525:
			goto st225
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st223
				}
			case _widec >= 42:
				goto st223
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st223
				}
			case _widec >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st0
tr329:
//line msg_parse.rl:256

			addr = new(Addr)
		
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:248

			{goto st154 }
		
	goto st752
tr334:
//line msg_parse.rl:256

			addr = new(Addr)
		
//line msg_parse.rl:107

			p = ( mark) - 1

		
//line msg_parse.rl:248

			{goto st154 }
		
	goto st752
tr338:
//line msg_parse.rl:256

			addr = new(Addr)
		
//line msg_parse.rl:107

			p = ( mark) - 1

		
//line msg_parse.rl:244

			{goto st195 }
		
	goto st752
	st752:
		p++
	st_case_752:
//line msg_parse.go:7794
		goto st0
	st225:
		p++
	st_case_225:
		if data[p] == 10 {
			goto st226
		}
		goto st0
	st226:
		p++
	st_case_226:
		switch data[p] {
		case 9:
			goto st227
		case 32:
			goto st227
		}
		goto st0
	st227:
		p++
	st_case_227:
		switch data[p] {
		case 9:
			goto st227
		case 32:
			goto st227
		case 33:
			goto st223
		case 37:
			goto st223
		case 39:
			goto st223
		case 60:
			goto tr334
		case 126:
			goto st223
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st223
				}
			case data[p] >= 42:
				goto st223
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st223
				}
			case data[p] >= 65:
				goto st223
			}
		default:
			goto st223
		}
		goto st0
tr330:
//line msg_parse.rl:103

			mark = p
		
	goto st228
	st228:
		p++
	st_case_228:
//line msg_parse.go:7864
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st224
		case 32:
			goto st224
		case 33:
			goto st223
		case 37:
			goto st223
		case 39:
			goto st223
		case 42:
			goto st223
		case 43:
			goto st228
		case 58:
			goto tr338
		case 126:
			goto st223
		case 525:
			goto st225
		}
		switch {
		case _widec < 65:
			switch {
			case _widec > 46:
				if 48 <= _widec && _widec <= 57 {
					goto st228
				}
			case _widec >= 45:
				goto st228
			}
		case _widec > 90:
			switch {
			case _widec > 96:
				if 97 <= _widec && _widec <= 122 {
					goto st228
				}
			case _widec >= 95:
				goto st223
			}
		default:
			goto st228
		}
		goto st0
	st229:
		p++
	st_case_229:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr339
		case 269:
			goto tr345
		case 525:
			goto tr346
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr340
				}
			case _widec >= 32:
				goto tr339
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr342
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr344
				}
			default:
				goto tr343
			}
		default:
			goto tr341
		}
		goto st0
tr339:
//line msg_parse.rl:103

			mark = p
		
	goto st230
	st230:
		p++
	st_case_230:
//line msg_parse.go:7971
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st230
		case 269:
			goto st236
		case 525:
			goto st237
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto st231
				}
			case _widec >= 32:
				goto st230
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto st233
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto st235
				}
			default:
				goto st234
			}
		default:
			goto st232
		}
		goto st0
tr340:
//line msg_parse.rl:103

			mark = p
		
	goto st231
	st231:
		p++
	st_case_231:
//line msg_parse.go:8023
		if 128 <= data[p] && data[p] <= 191 {
			goto st230
		}
		goto st0
tr341:
//line msg_parse.rl:103

			mark = p
		
	goto st232
	st232:
		p++
	st_case_232:
//line msg_parse.go:8037
		if 128 <= data[p] && data[p] <= 191 {
			goto st231
		}
		goto st0
tr342:
//line msg_parse.rl:103

			mark = p
		
	goto st233
	st233:
		p++
	st_case_233:
//line msg_parse.go:8051
		if 128 <= data[p] && data[p] <= 191 {
			goto st232
		}
		goto st0
tr343:
//line msg_parse.rl:103

			mark = p
		
	goto st234
	st234:
		p++
	st_case_234:
//line msg_parse.go:8065
		if 128 <= data[p] && data[p] <= 191 {
			goto st233
		}
		goto st0
tr344:
//line msg_parse.rl:103

			mark = p
		
	goto st235
	st235:
		p++
	st_case_235:
//line msg_parse.go:8079
		if 128 <= data[p] && data[p] <= 191 {
			goto st234
		}
		goto st0
tr345:
//line msg_parse.rl:103

			mark = p
		
	goto st236
	st236:
		p++
	st_case_236:
//line msg_parse.go:8093
		if data[p] == 10 {
			goto tr355
		}
		goto st0
tr355:
//line msg_parse.rl:228
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
//line msg_parse.rl:201

			{goto st246 }
		
	goto st753
	st753:
		p++
	st_case_753:
//line msg_parse.go:8119
		goto st0
tr346:
//line msg_parse.rl:103

			mark = p
		
	goto st237
	st237:
		p++
	st_case_237:
//line msg_parse.go:8130
		if data[p] == 10 {
			goto st238
		}
		goto st0
	st238:
		p++
	st_case_238:
		switch data[p] {
		case 9:
			goto st230
		case 32:
			goto st230
		}
		goto st0
	st239:
		p++
	st_case_239:
		switch data[p] {
		case 33:
			goto st240
		case 37:
			goto st240
		case 39:
			goto st240
		case 126:
			goto st240
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st240
				}
			case data[p] >= 42:
				goto st240
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st240
				}
			case data[p] >= 65:
				goto st240
			}
		default:
			goto st240
		}
		goto st0
	st240:
		p++
	st_case_240:
		switch data[p] {
		case 9:
			goto tr358
		case 32:
			goto tr358
		case 33:
			goto st240
		case 37:
			goto st240
		case 39:
			goto st240
		case 58:
			goto tr359
		case 126:
			goto st240
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto st240
				}
			case data[p] >= 42:
				goto st240
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto st240
				}
			case data[p] >= 65:
				goto st240
			}
		default:
			goto st240
		}
		goto st0
tr358:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st241
	st241:
		p++
	st_case_241:
//line msg_parse.go:8232
		switch data[p] {
		case 9:
			goto st241
		case 32:
			goto st241
		case 58:
			goto st242
		}
		goto st0
tr359:
//line msg_parse.rl:224

			name = string(data[mark:p])
		
	goto st242
	st242:
		p++
	st_case_242:
//line msg_parse.go:8251
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st242
		case 32:
			goto st242
		case 269:
			goto tr362
		case 525:
			goto st243
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr362
			}
		default:
			goto tr362
		}
		goto st0
tr362:
//line msg_parse.rl:573
value=nil
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:205

			{goto st229 }
		
	goto st754
	st754:
		p++
	st_case_754:
//line msg_parse.go:8294
		goto st0
	st243:
		p++
	st_case_243:
		if data[p] == 10 {
			goto st244
		}
		goto st0
	st244:
		p++
	st_case_244:
		switch data[p] {
		case 9:
			goto st245
		case 32:
			goto st245
		}
		goto st0
	st245:
		p++
	st_case_245:
		switch data[p] {
		case 9:
			goto st245
		case 32:
			goto st245
		}
		goto tr362
	st246:
		p++
	st_case_246:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 33:
			goto tr366
		case 37:
			goto tr366
		case 39:
			goto tr366
		case 126:
			goto tr366
		case 269:
			goto st738
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr366
				}
			case _widec >= 42:
				goto tr366
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr366
				}
			case _widec >= 65:
				goto tr366
			}
		default:
			goto tr366
		}
		goto st0
tr366:
//line msg_parse.rl:103

			mark = p
		
//line msg_parse.rl:95

			p--

		
	goto st247
	st247:
		p++
	st_case_247:
//line msg_parse.go:8382
		switch data[p] {
		case 65:
			goto st248
		case 66:
			goto st329
		case 67:
			goto st330
		case 68:
			goto st433
		case 69:
			goto st437
		case 70:
			goto st463
		case 73:
			goto st467
		case 75:
			goto st478
		case 76:
			goto st479
		case 77:
			goto st486
		case 79:
			goto st530
		case 80:
			goto st542
		case 82:
			goto st600
		case 83:
			goto st660
		case 84:
			goto st678
		case 85:
			goto st688
		case 86:
			goto st708
		case 87:
			goto st716
		case 97:
			goto st248
		case 98:
			goto st329
		case 99:
			goto st330
		case 100:
			goto st433
		case 101:
			goto st437
		case 102:
			goto st463
		case 105:
			goto st467
		case 107:
			goto st478
		case 108:
			goto st479
		case 109:
			goto st486
		case 111:
			goto st530
		case 112:
			goto st542
		case 114:
			goto st600
		case 115:
			goto st660
		case 116:
			goto st678
		case 117:
			goto st688
		case 118:
			goto st708
		case 119:
			goto st716
		}
		goto tr368
	st248:
		p++
	st_case_248:
		switch data[p] {
		case 9:
			goto tr387
		case 32:
			goto tr387
		case 58:
			goto tr388
		case 67:
			goto st254
		case 76:
			goto st283
		case 85:
			goto st302
		case 99:
			goto st254
		case 108:
			goto st283
		case 117:
			goto st302
		}
		goto tr368
tr387:
//line msg_parse.rl:491
value=&msg.AcceptContact
	goto st249
tr402:
//line msg_parse.rl:490
value=&msg.Accept
	goto st249
tr421:
//line msg_parse.rl:492
value=&msg.AcceptEncoding
	goto st249
tr430:
//line msg_parse.rl:493
value=&msg.AcceptLanguage
	goto st249
tr441:
//line msg_parse.rl:496
value=&msg.AlertInfo
	goto st249
tr445:
//line msg_parse.rl:494
value=&msg.Allow
	goto st249
tr454:
//line msg_parse.rl:495
value=&msg.AllowEvents
	goto st249
tr474:
//line msg_parse.rl:497
value=&msg.AuthenticationInfo
	goto st249
tr484:
//line msg_parse.rl:498
value=&msg.Authorization
	goto st249
tr486:
//line msg_parse.rl:515
value=&msg.ReferredBy
	goto st249
tr532:
//line msg_parse.rl:502
value=&msg.CallInfo
	goto st249
tr565:
//line msg_parse.rl:499
value=&msg.ContentDisposition
	goto st249
tr574:
//line msg_parse.rl:501
value=&msg.ContentEncoding
	goto st249
tr584:
//line msg_parse.rl:500
value=&msg.ContentLanguage
	goto st249
tr619:
//line msg_parse.rl:503
value=&msg.Date
	goto st249
tr632:
//line msg_parse.rl:504
value=&msg.ErrorInfo
	goto st249
tr637:
//line msg_parse.rl:505
value=&msg.Event
	goto st249
tr666:
//line msg_parse.rl:506
value=&msg.InReplyTo
	goto st249
tr668:
//line msg_parse.rl:520
value=&msg.Supported
	goto st249
tr707:
//line msg_parse.rl:508
value=&msg.MIMEVersion
	goto st249
tr735:
//line msg_parse.rl:509
value=&msg.Organization
	goto st249
tr765:
//line msg_parse.rl:510
value=&msg.Priority
	goto st249
tr784:
//line msg_parse.rl:511
value=&msg.ProxyAuthenticate
	goto st249
tr794:
//line msg_parse.rl:512
value=&msg.ProxyAuthorization
	goto st249
tr802:
//line msg_parse.rl:513
value=&msg.ProxyRequire
	goto st249
tr804:
//line msg_parse.rl:514
value=&msg.ReferTo
	goto st249
tr854:
//line msg_parse.rl:507
value=&msg.ReplyTo
	goto st249
tr860:
//line msg_parse.rl:516
value=&msg.Require
	goto st249
tr870:
//line msg_parse.rl:517
value=&msg.RetryAfter
	goto st249
tr877:
//line msg_parse.rl:519
value=&msg.Subject
	goto st249
tr885:
//line msg_parse.rl:518
value=&msg.Server
	goto st249
tr909:
//line msg_parse.rl:521
value=&msg.Timestamp
	goto st249
tr911:
//line msg_parse.rl:494
value=&msg.Allow
//line msg_parse.rl:495
value=&msg.AllowEvents
	goto st249
tr924:
//line msg_parse.rl:522
value=&msg.Unsupported
	goto st249
tr934:
//line msg_parse.rl:523
value=&msg.UserAgent
	goto st249
tr951:
//line msg_parse.rl:524
value=&msg.Warning
	goto st249
tr967:
//line msg_parse.rl:525
value=&msg.WWWAuthenticate
	goto st249
	st249:
		p++
	st_case_249:
//line msg_parse.go:8635
		switch data[p] {
		case 9:
			goto st249
		case 32:
			goto st249
		case 58:
			goto st250
		}
		goto st0
tr388:
//line msg_parse.rl:491
value=&msg.AcceptContact
	goto st250
tr404:
//line msg_parse.rl:490
value=&msg.Accept
	goto st250
tr422:
//line msg_parse.rl:492
value=&msg.AcceptEncoding
	goto st250
tr431:
//line msg_parse.rl:493
value=&msg.AcceptLanguage
	goto st250
tr442:
//line msg_parse.rl:496
value=&msg.AlertInfo
	goto st250
tr447:
//line msg_parse.rl:494
value=&msg.Allow
	goto st250
tr455:
//line msg_parse.rl:495
value=&msg.AllowEvents
	goto st250
tr475:
//line msg_parse.rl:497
value=&msg.AuthenticationInfo
	goto st250
tr485:
//line msg_parse.rl:498
value=&msg.Authorization
	goto st250
tr487:
//line msg_parse.rl:515
value=&msg.ReferredBy
	goto st250
tr533:
//line msg_parse.rl:502
value=&msg.CallInfo
	goto st250
tr566:
//line msg_parse.rl:499
value=&msg.ContentDisposition
	goto st250
tr575:
//line msg_parse.rl:501
value=&msg.ContentEncoding
	goto st250
tr585:
//line msg_parse.rl:500
value=&msg.ContentLanguage
	goto st250
tr620:
//line msg_parse.rl:503
value=&msg.Date
	goto st250
tr633:
//line msg_parse.rl:504
value=&msg.ErrorInfo
	goto st250
tr638:
//line msg_parse.rl:505
value=&msg.Event
	goto st250
tr667:
//line msg_parse.rl:506
value=&msg.InReplyTo
	goto st250
tr669:
//line msg_parse.rl:520
value=&msg.Supported
	goto st250
tr708:
//line msg_parse.rl:508
value=&msg.MIMEVersion
	goto st250
tr736:
//line msg_parse.rl:509
value=&msg.Organization
	goto st250
tr766:
//line msg_parse.rl:510
value=&msg.Priority
	goto st250
tr785:
//line msg_parse.rl:511
value=&msg.ProxyAuthenticate
	goto st250
tr795:
//line msg_parse.rl:512
value=&msg.ProxyAuthorization
	goto st250
tr803:
//line msg_parse.rl:513
value=&msg.ProxyRequire
	goto st250
tr805:
//line msg_parse.rl:514
value=&msg.ReferTo
	goto st250
tr855:
//line msg_parse.rl:507
value=&msg.ReplyTo
	goto st250
tr861:
//line msg_parse.rl:516
value=&msg.Require
	goto st250
tr871:
//line msg_parse.rl:517
value=&msg.RetryAfter
	goto st250
tr878:
//line msg_parse.rl:519
value=&msg.Subject
	goto st250
tr886:
//line msg_parse.rl:518
value=&msg.Server
	goto st250
tr910:
//line msg_parse.rl:521
value=&msg.Timestamp
	goto st250
tr912:
//line msg_parse.rl:494
value=&msg.Allow
//line msg_parse.rl:495
value=&msg.AllowEvents
	goto st250
tr925:
//line msg_parse.rl:522
value=&msg.Unsupported
	goto st250
tr935:
//line msg_parse.rl:523
value=&msg.UserAgent
	goto st250
tr952:
//line msg_parse.rl:524
value=&msg.Warning
	goto st250
tr968:
//line msg_parse.rl:525
value=&msg.WWWAuthenticate
	goto st250
	st250:
		p++
	st_case_250:
//line msg_parse.go:8798
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st250
		case 32:
			goto st250
		case 269:
			goto tr394
		case 525:
			goto st251
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr394
			}
		default:
			goto tr394
		}
		goto st0
tr509:
//line msg_parse.rl:201

			{goto st246 }
		
	goto st755
tr394:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:205

			{goto st229 }
		
	goto st755
tr544:
//line msg_parse.rl:575
value=nil
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:240

			{goto st222 }
		
	goto st755
tr939:
//line msg_parse.rl:95

			p--

		
//line msg_parse.rl:209

			via = new(Via)
			{goto st69 }
		
	goto st755
tr969:
//line msg_parse.rl:99

			{p++; cs = 755; goto _out }
		
	goto st755
	st755:
		p++
	st_case_755:
//line msg_parse.go:8876
		goto st0
	st251:
		p++
	st_case_251:
		if data[p] == 10 {
			goto st252
		}
		goto st0
	st252:
		p++
	st_case_252:
		switch data[p] {
		case 9:
			goto st253
		case 32:
			goto st253
		}
		goto st0
	st253:
		p++
	st_case_253:
		switch data[p] {
		case 9:
			goto st253
		case 32:
			goto st253
		}
		goto tr394
	st254:
		p++
	st_case_254:
		switch data[p] {
		case 67:
			goto st255
		case 99:
			goto st255
		}
		goto tr368
	st255:
		p++
	st_case_255:
		switch data[p] {
		case 69:
			goto st256
		case 101:
			goto st256
		}
		goto tr368
	st256:
		p++
	st_case_256:
		switch data[p] {
		case 80:
			goto st257
		case 112:
			goto st257
		}
		goto tr368
	st257:
		p++
	st_case_257:
		switch data[p] {
		case 84:
			goto st258
		case 116:
			goto st258
		}
		goto tr368
	st258:
		p++
	st_case_258:
		switch data[p] {
		case 9:
			goto tr402
		case 32:
			goto tr402
		case 45:
			goto st259
		case 58:
			goto tr404
		}
		goto tr368
	st259:
		p++
	st_case_259:
		switch data[p] {
		case 67:
			goto st260
		case 69:
			goto st267
		case 76:
			goto st275
		case 99:
			goto st260
		case 101:
			goto st267
		case 108:
			goto st275
		}
		goto tr368
	st260:
		p++
	st_case_260:
		switch data[p] {
		case 79:
			goto st261
		case 111:
			goto st261
		}
		goto tr368
	st261:
		p++
	st_case_261:
		switch data[p] {
		case 78:
			goto st262
		case 110:
			goto st262
		}
		goto tr368
	st262:
		p++
	st_case_262:
		switch data[p] {
		case 84:
			goto st263
		case 116:
			goto st263
		}
		goto tr368
	st263:
		p++
	st_case_263:
		switch data[p] {
		case 65:
			goto st264
		case 97:
			goto st264
		}
		goto tr368
	st264:
		p++
	st_case_264:
		switch data[p] {
		case 67:
			goto st265
		case 99:
			goto st265
		}
		goto tr368
	st265:
		p++
	st_case_265:
		switch data[p] {
		case 84:
			goto st266
		case 116:
			goto st266
		}
		goto tr368
	st266:
		p++
	st_case_266:
		switch data[p] {
		case 9:
			goto tr387
		case 32:
			goto tr387
		case 58:
			goto tr388
		}
		goto tr368
	st267:
		p++
	st_case_267:
		switch data[p] {
		case 78:
			goto st268
		case 110:
			goto st268
		}
		goto tr368
	st268:
		p++
	st_case_268:
		switch data[p] {
		case 67:
			goto st269
		case 99:
			goto st269
		}
		goto tr368
	st269:
		p++
	st_case_269:
		switch data[p] {
		case 79:
			goto st270
		case 111:
			goto st270
		}
		goto tr368
	st270:
		p++
	st_case_270:
		switch data[p] {
		case 68:
			goto st271
		case 100:
			goto st271
		}
		goto tr368
	st271:
		p++
	st_case_271:
		switch data[p] {
		case 73:
			goto st272
		case 105:
			goto st272
		}
		goto tr368
	st272:
		p++
	st_case_272:
		switch data[p] {
		case 78:
			goto st273
		case 110:
			goto st273
		}
		goto tr368
	st273:
		p++
	st_case_273:
		switch data[p] {
		case 71:
			goto st274
		case 103:
			goto st274
		}
		goto tr368
	st274:
		p++
	st_case_274:
		switch data[p] {
		case 9:
			goto tr421
		case 32:
			goto tr421
		case 58:
			goto tr422
		}
		goto tr368
	st275:
		p++
	st_case_275:
		switch data[p] {
		case 65:
			goto st276
		case 97:
			goto st276
		}
		goto tr368
	st276:
		p++
	st_case_276:
		switch data[p] {
		case 78:
			goto st277
		case 110:
			goto st277
		}
		goto tr368
	st277:
		p++
	st_case_277:
		switch data[p] {
		case 71:
			goto st278
		case 103:
			goto st278
		}
		goto tr368
	st278:
		p++
	st_case_278:
		switch data[p] {
		case 85:
			goto st279
		case 117:
			goto st279
		}
		goto tr368
	st279:
		p++
	st_case_279:
		switch data[p] {
		case 65:
			goto st280
		case 97:
			goto st280
		}
		goto tr368
	st280:
		p++
	st_case_280:
		switch data[p] {
		case 71:
			goto st281
		case 103:
			goto st281
		}
		goto tr368
	st281:
		p++
	st_case_281:
		switch data[p] {
		case 69:
			goto st282
		case 101:
			goto st282
		}
		goto tr368
	st282:
		p++
	st_case_282:
		switch data[p] {
		case 9:
			goto tr430
		case 32:
			goto tr430
		case 58:
			goto tr431
		}
		goto tr368
	st283:
		p++
	st_case_283:
		switch data[p] {
		case 69:
			goto st284
		case 76:
			goto st292
		case 101:
			goto st284
		case 108:
			goto st292
		}
		goto tr368
	st284:
		p++
	st_case_284:
		switch data[p] {
		case 82:
			goto st285
		case 114:
			goto st285
		}
		goto tr368
	st285:
		p++
	st_case_285:
		switch data[p] {
		case 84:
			goto st286
		case 116:
			goto st286
		}
		goto tr368
	st286:
		p++
	st_case_286:
		if data[p] == 45 {
			goto st287
		}
		goto tr368
	st287:
		p++
	st_case_287:
		switch data[p] {
		case 73:
			goto st288
		case 105:
			goto st288
		}
		goto tr368
	st288:
		p++
	st_case_288:
		switch data[p] {
		case 78:
			goto st289
		case 110:
			goto st289
		}
		goto tr368
	st289:
		p++
	st_case_289:
		switch data[p] {
		case 70:
			goto st290
		case 102:
			goto st290
		}
		goto tr368
	st290:
		p++
	st_case_290:
		switch data[p] {
		case 79:
			goto st291
		case 111:
			goto st291
		}
		goto tr368
	st291:
		p++
	st_case_291:
		switch data[p] {
		case 9:
			goto tr441
		case 32:
			goto tr441
		case 58:
			goto tr442
		}
		goto tr368
	st292:
		p++
	st_case_292:
		switch data[p] {
		case 79:
			goto st293
		case 111:
			goto st293
		}
		goto tr368
	st293:
		p++
	st_case_293:
		switch data[p] {
		case 87:
			goto st294
		case 119:
			goto st294
		}
		goto tr368
	st294:
		p++
	st_case_294:
		switch data[p] {
		case 9:
			goto tr445
		case 32:
			goto tr445
		case 45:
			goto st295
		case 58:
			goto tr447
		}
		goto tr368
	st295:
		p++
	st_case_295:
		switch data[p] {
		case 69:
			goto st296
		case 101:
			goto st296
		}
		goto tr368
	st296:
		p++
	st_case_296:
		switch data[p] {
		case 86:
			goto st297
		case 118:
			goto st297
		}
		goto tr368
	st297:
		p++
	st_case_297:
		switch data[p] {
		case 69:
			goto st298
		case 101:
			goto st298
		}
		goto tr368
	st298:
		p++
	st_case_298:
		switch data[p] {
		case 78:
			goto st299
		case 110:
			goto st299
		}
		goto tr368
	st299:
		p++
	st_case_299:
		switch data[p] {
		case 84:
			goto st300
		case 116:
			goto st300
		}
		goto tr368
	st300:
		p++
	st_case_300:
		switch data[p] {
		case 83:
			goto st301
		case 115:
			goto st301
		}
		goto tr368
	st301:
		p++
	st_case_301:
		switch data[p] {
		case 9:
			goto tr454
		case 32:
			goto tr454
		case 58:
			goto tr455
		}
		goto tr368
	st302:
		p++
	st_case_302:
		switch data[p] {
		case 84:
			goto st303
		case 116:
			goto st303
		}
		goto tr368
	st303:
		p++
	st_case_303:
		switch data[p] {
		case 72:
			goto st304
		case 104:
			goto st304
		}
		goto tr368
	st304:
		p++
	st_case_304:
		switch data[p] {
		case 69:
			goto st305
		case 79:
			goto st320
		case 101:
			goto st305
		case 111:
			goto st320
		}
		goto tr368
	st305:
		p++
	st_case_305:
		switch data[p] {
		case 78:
			goto st306
		case 110:
			goto st306
		}
		goto tr368
	st306:
		p++
	st_case_306:
		switch data[p] {
		case 84:
			goto st307
		case 116:
			goto st307
		}
		goto tr368
	st307:
		p++
	st_case_307:
		switch data[p] {
		case 73:
			goto st308
		case 105:
			goto st308
		}
		goto tr368
	st308:
		p++
	st_case_308:
		switch data[p] {
		case 67:
			goto st309
		case 99:
			goto st309
		}
		goto tr368
	st309:
		p++
	st_case_309:
		switch data[p] {
		case 65:
			goto st310
		case 97:
			goto st310
		}
		goto tr368
	st310:
		p++
	st_case_310:
		switch data[p] {
		case 84:
			goto st311
		case 116:
			goto st311
		}
		goto tr368
	st311:
		p++
	st_case_311:
		switch data[p] {
		case 73:
			goto st312
		case 105:
			goto st312
		}
		goto tr368
	st312:
		p++
	st_case_312:
		switch data[p] {
		case 79:
			goto st313
		case 111:
			goto st313
		}
		goto tr368
	st313:
		p++
	st_case_313:
		switch data[p] {
		case 78:
			goto st314
		case 110:
			goto st314
		}
		goto tr368
	st314:
		p++
	st_case_314:
		if data[p] == 45 {
			goto st315
		}
		goto tr368
	st315:
		p++
	st_case_315:
		switch data[p] {
		case 73:
			goto st316
		case 105:
			goto st316
		}
		goto tr368
	st316:
		p++
	st_case_316:
		switch data[p] {
		case 78:
			goto st317
		case 110:
			goto st317
		}
		goto tr368
	st317:
		p++
	st_case_317:
		switch data[p] {
		case 70:
			goto st318
		case 102:
			goto st318
		}
		goto tr368
	st318:
		p++
	st_case_318:
		switch data[p] {
		case 79:
			goto st319
		case 111:
			goto st319
		}
		goto tr368
	st319:
		p++
	st_case_319:
		switch data[p] {
		case 9:
			goto tr474
		case 32:
			goto tr474
		case 58:
			goto tr475
		}
		goto tr368
	st320:
		p++
	st_case_320:
		switch data[p] {
		case 82:
			goto st321
		case 114:
			goto st321
		}
		goto tr368
	st321:
		p++
	st_case_321:
		switch data[p] {
		case 73:
			goto st322
		case 105:
			goto st322
		}
		goto tr368
	st322:
		p++
	st_case_322:
		switch data[p] {
		case 90:
			goto st323
		case 122:
			goto st323
		}
		goto tr368
	st323:
		p++
	st_case_323:
		switch data[p] {
		case 65:
			goto st324
		case 97:
			goto st324
		}
		goto tr368
	st324:
		p++
	st_case_324:
		switch data[p] {
		case 84:
			goto st325
		case 116:
			goto st325
		}
		goto tr368
	st325:
		p++
	st_case_325:
		switch data[p] {
		case 73:
			goto st326
		case 105:
			goto st326
		}
		goto tr368
	st326:
		p++
	st_case_326:
		switch data[p] {
		case 79:
			goto st327
		case 111:
			goto st327
		}
		goto tr368
	st327:
		p++
	st_case_327:
		switch data[p] {
		case 78:
			goto st328
		case 110:
			goto st328
		}
		goto tr368
	st328:
		p++
	st_case_328:
		switch data[p] {
		case 9:
			goto tr484
		case 32:
			goto tr484
		case 58:
			goto tr485
		}
		goto tr368
	st329:
		p++
	st_case_329:
		switch data[p] {
		case 9:
			goto tr486
		case 32:
			goto tr486
		case 58:
			goto tr487
		}
		goto tr368
	st330:
		p++
	st_case_330:
		switch data[p] {
		case 9:
			goto st331
		case 32:
			goto st331
		case 58:
			goto st332
		case 65:
			goto st345
		case 79:
			goto st362
		case 83:
			goto st419
		case 97:
			goto st345
		case 111:
			goto st362
		case 115:
			goto st419
		}
		goto tr368
	st331:
		p++
	st_case_331:
		switch data[p] {
		case 9:
			goto st331
		case 32:
			goto st331
		case 58:
			goto st332
		}
		goto st0
	st332:
		p++
	st_case_332:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st332
		case 32:
			goto st332
		case 269:
			goto tr499
		case 525:
			goto st342
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr494
				}
			case _widec >= 33:
				goto tr493
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr496
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr498
				}
			default:
				goto tr497
			}
		default:
			goto tr495
		}
		goto st0
tr493:
//line msg_parse.rl:103

			mark = p
		
	goto st333
	st333:
		p++
	st_case_333:
//line msg_parse.go:9791
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st333
		case 269:
			goto tr507
		case 525:
			goto st340
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto st334
				}
			case _widec >= 32:
				goto st333
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto st336
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto st338
				}
			default:
				goto st337
			}
		default:
			goto st335
		}
		goto st0
tr494:
//line msg_parse.rl:103

			mark = p
		
	goto st334
	st334:
		p++
	st_case_334:
//line msg_parse.go:9843
		if 128 <= data[p] && data[p] <= 191 {
			goto st333
		}
		goto st0
tr495:
//line msg_parse.rl:103

			mark = p
		
	goto st335
	st335:
		p++
	st_case_335:
//line msg_parse.go:9857
		if 128 <= data[p] && data[p] <= 191 {
			goto st334
		}
		goto st0
tr496:
//line msg_parse.rl:103

			mark = p
		
	goto st336
	st336:
		p++
	st_case_336:
//line msg_parse.go:9871
		if 128 <= data[p] && data[p] <= 191 {
			goto st335
		}
		goto st0
tr497:
//line msg_parse.rl:103

			mark = p
		
	goto st337
	st337:
		p++
	st_case_337:
//line msg_parse.go:9885
		if 128 <= data[p] && data[p] <= 191 {
			goto st336
		}
		goto st0
tr498:
//line msg_parse.rl:103

			mark = p
		
	goto st338
	st338:
		p++
	st_case_338:
//line msg_parse.go:9899
		if 128 <= data[p] && data[p] <= 191 {
			goto st337
		}
		goto st0
tr499:
//line msg_parse.rl:103

			mark = p
		
//line msg_parse.rl:298

			ctype = string(data[mark:p])
		
	goto st339
tr507:
//line msg_parse.rl:298

			ctype = string(data[mark:p])
		
	goto st339
tr526:
//line msg_parse.rl:290

			msg.CallID = string(data[mark:p])
		
	goto st339
tr611:
//line msg_parse.rl:306

			msg.CSeqMethod = string(data[mark:p])
		
	goto st339
	st339:
		p++
	st_case_339:
//line msg_parse.go:9935
		if data[p] == 10 {
			goto tr509
		}
		goto st0
tr513:
//line msg_parse.rl:103

			mark = p
		
	goto st340
	st340:
		p++
	st_case_340:
//line msg_parse.go:9949
		if data[p] == 10 {
			goto st341
		}
		goto st0
	st341:
		p++
	st_case_341:
		switch data[p] {
		case 9:
			goto st333
		case 32:
			goto st333
		}
		goto st0
	st342:
		p++
	st_case_342:
		if data[p] == 10 {
			goto st343
		}
		goto st0
	st343:
		p++
	st_case_343:
		switch data[p] {
		case 9:
			goto st344
		case 32:
			goto st344
		}
		goto st0
	st344:
		p++
	st_case_344:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st344
		case 32:
			goto st344
		case 269:
			goto tr499
		case 525:
			goto tr513
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 127:
				if 192 <= _widec && _widec <= 223 {
					goto tr494
				}
			case _widec >= 33:
				goto tr493
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr496
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr498
				}
			default:
				goto tr497
			}
		default:
			goto tr495
		}
		goto st0
	st345:
		p++
	st_case_345:
		switch data[p] {
		case 76:
			goto st346
		case 108:
			goto st346
		}
		goto tr368
	st346:
		p++
	st_case_346:
		switch data[p] {
		case 76:
			goto st347
		case 108:
			goto st347
		}
		goto tr368
	st347:
		p++
	st_case_347:
		if data[p] == 45 {
			goto st348
		}
		goto tr368
	st348:
		p++
	st_case_348:
		switch data[p] {
		case 73:
			goto st349
		case 105:
			goto st349
		}
		goto tr368
	st349:
		p++
	st_case_349:
		switch data[p] {
		case 68:
			goto st350
		case 78:
			goto st359
		case 100:
			goto st350
		case 110:
			goto st359
		}
		goto tr368
	st350:
		p++
	st_case_350:
		switch data[p] {
		case 9:
			goto st351
		case 32:
			goto st351
		case 58:
			goto st352
		}
		goto tr368
	st351:
		p++
	st_case_351:
		switch data[p] {
		case 9:
			goto st351
		case 32:
			goto st351
		case 58:
			goto st352
		}
		goto st0
	st352:
		p++
	st_case_352:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st352
		case 32:
			goto st352
		case 37:
			goto tr522
		case 60:
			goto tr522
		case 525:
			goto st356
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto tr522
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto tr522
				}
			default:
				goto tr522
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto tr522
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto tr522
				}
			default:
				goto tr522
			}
		default:
			goto tr522
		}
		goto st0
tr522:
//line msg_parse.rl:103

			mark = p
		
	goto st353
	st353:
		p++
	st_case_353:
//line msg_parse.go:10165
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 37:
			goto st353
		case 60:
			goto st353
		case 64:
			goto st354
		case 269:
			goto tr526
		}
		switch {
		case _widec < 45:
			switch {
			case _widec > 34:
				if 39 <= _widec && _widec <= 43 {
					goto st353
				}
			case _widec >= 33:
				goto st353
			}
		case _widec > 58:
			switch {
			case _widec < 95:
				if 62 <= _widec && _widec <= 93 {
					goto st353
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto st353
				}
			default:
				goto st353
			}
		default:
			goto st353
		}
		goto st0
	st354:
		p++
	st_case_354:
		switch data[p] {
		case 37:
			goto st355
		case 60:
			goto st355
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto st355
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto st355
				}
			default:
				goto st355
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto st355
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto st355
				}
			default:
				goto st355
			}
		default:
			goto st355
		}
		goto st0
	st355:
		p++
	st_case_355:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 37:
			goto st355
		case 60:
			goto st355
		case 269:
			goto tr526
		}
		switch {
		case _widec < 62:
			switch {
			case _widec < 39:
				if 33 <= _widec && _widec <= 34 {
					goto st355
				}
			case _widec > 43:
				if 45 <= _widec && _widec <= 58 {
					goto st355
				}
			default:
				goto st355
			}
		case _widec > 63:
			switch {
			case _widec < 95:
				if 65 <= _widec && _widec <= 93 {
					goto st355
				}
			case _widec > 123:
				if 125 <= _widec && _widec <= 126 {
					goto st355
				}
			default:
				goto st355
			}
		default:
			goto st355
		}
		goto st0
	st356:
		p++
	st_case_356:
		if data[p] == 10 {
			goto st357
		}
		goto st0
	st357:
		p++
	st_case_357:
		switch data[p] {
		case 9:
			goto st358
		case 32:
			goto st358
		}
		goto st0
	st358:
		p++
	st_case_358:
		switch data[p] {
		case 9:
			goto st358
		case 32:
			goto st358
		case 37:
			goto tr522
		case 60:
			goto tr522
		}
		switch {
		case data[p] < 62:
			switch {
			case data[p] < 39:
				if 33 <= data[p] && data[p] <= 34 {
					goto tr522
				}
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 58 {
					goto tr522
				}
			default:
				goto tr522
			}
		case data[p] > 63:
			switch {
			case data[p] < 95:
				if 65 <= data[p] && data[p] <= 93 {
					goto tr522
				}
			case data[p] > 123:
				if 125 <= data[p] && data[p] <= 126 {
					goto tr522
				}
			default:
				goto tr522
			}
		default:
			goto tr522
		}
		goto st0
	st359:
		p++
	st_case_359:
		switch data[p] {
		case 70:
			goto st360
		case 102:
			goto st360
		}
		goto tr368
	st360:
		p++
	st_case_360:
		switch data[p] {
		case 79:
			goto st361
		case 111:
			goto st361
		}
		goto tr368
	st361:
		p++
	st_case_361:
		switch data[p] {
		case 9:
			goto tr532
		case 32:
			goto tr532
		case 58:
			goto tr533
		}
		goto tr368
	st362:
		p++
	st_case_362:
		switch data[p] {
		case 78:
			goto st363
		case 110:
			goto st363
		}
		goto tr368
	st363:
		p++
	st_case_363:
		switch data[p] {
		case 84:
			goto st364
		case 116:
			goto st364
		}
		goto tr368
	st364:
		p++
	st_case_364:
		switch data[p] {
		case 65:
			goto st365
		case 69:
			goto st373
		case 97:
			goto st365
		case 101:
			goto st373
		}
		goto tr368
	st365:
		p++
	st_case_365:
		switch data[p] {
		case 67:
			goto st366
		case 99:
			goto st366
		}
		goto tr368
	st366:
		p++
	st_case_366:
		switch data[p] {
		case 84:
			goto st367
		case 116:
			goto st367
		}
		goto tr368
	st367:
		p++
	st_case_367:
		switch data[p] {
		case 9:
			goto tr540
		case 32:
			goto tr540
		case 58:
			goto tr541
		}
		goto tr368
tr540:
//line msg_parse.rl:477
addrp=lastAddr(&msg.Contact)
	goto st368
tr651:
//line msg_parse.rl:478
addrp=lastAddr(&msg.From)
	goto st368
tr756:
//line msg_parse.rl:479
addrp=lastAddr(&msg.PAssertedIdentity)
	goto st368
tr823:
//line msg_parse.rl:480
addrp=lastAddr(&msg.RecordRoute)
	goto st368
tr847:
//line msg_parse.rl:481
addrp=lastAddr(&msg.RemotePartyID)
	goto st368
tr875:
//line msg_parse.rl:482
addrp=lastAddr(&msg.Route)
	goto st368
tr898:
//line msg_parse.rl:483
addrp=lastAddr(&msg.To)
	goto st368
	st368:
		p++
	st_case_368:
//line msg_parse.go:10489
		switch data[p] {
		case 9:
			goto st368
		case 32:
			goto st368
		case 58:
			goto st369
		}
		goto st0
tr541:
//line msg_parse.rl:477
addrp=lastAddr(&msg.Contact)
	goto st369
tr652:
//line msg_parse.rl:478
addrp=lastAddr(&msg.From)
	goto st369
tr757:
//line msg_parse.rl:479
addrp=lastAddr(&msg.PAssertedIdentity)
	goto st369
tr824:
//line msg_parse.rl:480
addrp=lastAddr(&msg.RecordRoute)
	goto st369
tr848:
//line msg_parse.rl:481
addrp=lastAddr(&msg.RemotePartyID)
	goto st369
tr876:
//line msg_parse.rl:482
addrp=lastAddr(&msg.Route)
	goto st369
tr899:
//line msg_parse.rl:483
addrp=lastAddr(&msg.To)
	goto st369
	st369:
		p++
	st_case_369:
//line msg_parse.go:10530
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st369
		case 32:
			goto st369
		case 269:
			goto tr544
		case 525:
			goto st370
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr544
			}
		default:
			goto tr544
		}
		goto st0
	st370:
		p++
	st_case_370:
		if data[p] == 10 {
			goto st371
		}
		goto st0
	st371:
		p++
	st_case_371:
		switch data[p] {
		case 9:
			goto st372
		case 32:
			goto st372
		}
		goto st0
	st372:
		p++
	st_case_372:
		switch data[p] {
		case 9:
			goto st372
		case 32:
			goto st372
		}
		goto tr544
	st373:
		p++
	st_case_373:
		switch data[p] {
		case 78:
			goto st374
		case 110:
			goto st374
		}
		goto tr368
	st374:
		p++
	st_case_374:
		switch data[p] {
		case 84:
			goto st375
		case 116:
			goto st375
		}
		goto tr368
	st375:
		p++
	st_case_375:
		if data[p] == 45 {
			goto st376
		}
		goto tr368
	st376:
		p++
	st_case_376:
		switch data[p] {
		case 68:
			goto st377
		case 69:
			goto st388
		case 76:
			goto st396
		case 84:
			goto st415
		case 100:
			goto st377
		case 101:
			goto st388
		case 108:
			goto st396
		case 116:
			goto st415
		}
		goto tr368
	st377:
		p++
	st_case_377:
		switch data[p] {
		case 73:
			goto st378
		case 105:
			goto st378
		}
		goto tr368
	st378:
		p++
	st_case_378:
		switch data[p] {
		case 83:
			goto st379
		case 115:
			goto st379
		}
		goto tr368
	st379:
		p++
	st_case_379:
		switch data[p] {
		case 80:
			goto st380
		case 112:
			goto st380
		}
		goto tr368
	st380:
		p++
	st_case_380:
		switch data[p] {
		case 79:
			goto st381
		case 111:
			goto st381
		}
		goto tr368
	st381:
		p++
	st_case_381:
		switch data[p] {
		case 83:
			goto st382
		case 115:
			goto st382
		}
		goto tr368
	st382:
		p++
	st_case_382:
		switch data[p] {
		case 73:
			goto st383
		case 105:
			goto st383
		}
		goto tr368
	st383:
		p++
	st_case_383:
		switch data[p] {
		case 84:
			goto st384
		case 116:
			goto st384
		}
		goto tr368
	st384:
		p++
	st_case_384:
		switch data[p] {
		case 73:
			goto st385
		case 105:
			goto st385
		}
		goto tr368
	st385:
		p++
	st_case_385:
		switch data[p] {
		case 79:
			goto st386
		case 111:
			goto st386
		}
		goto tr368
	st386:
		p++
	st_case_386:
		switch data[p] {
		case 78:
			goto st387
		case 110:
			goto st387
		}
		goto tr368
	st387:
		p++
	st_case_387:
		switch data[p] {
		case 9:
			goto tr565
		case 32:
			goto tr565
		case 58:
			goto tr566
		}
		goto tr368
	st388:
		p++
	st_case_388:
		switch data[p] {
		case 78:
			goto st389
		case 110:
			goto st389
		}
		goto tr368
	st389:
		p++
	st_case_389:
		switch data[p] {
		case 67:
			goto st390
		case 99:
			goto st390
		}
		goto tr368
	st390:
		p++
	st_case_390:
		switch data[p] {
		case 79:
			goto st391
		case 111:
			goto st391
		}
		goto tr368
	st391:
		p++
	st_case_391:
		switch data[p] {
		case 68:
			goto st392
		case 100:
			goto st392
		}
		goto tr368
	st392:
		p++
	st_case_392:
		switch data[p] {
		case 73:
			goto st393
		case 105:
			goto st393
		}
		goto tr368
	st393:
		p++
	st_case_393:
		switch data[p] {
		case 78:
			goto st394
		case 110:
			goto st394
		}
		goto tr368
	st394:
		p++
	st_case_394:
		switch data[p] {
		case 71:
			goto st395
		case 103:
			goto st395
		}
		goto tr368
	st395:
		p++
	st_case_395:
		switch data[p] {
		case 9:
			goto tr574
		case 32:
			goto tr574
		case 58:
			goto tr575
		}
		goto tr368
	st396:
		p++
	st_case_396:
		switch data[p] {
		case 65:
			goto st397
		case 69:
			goto st404
		case 97:
			goto st397
		case 101:
			goto st404
		}
		goto tr368
	st397:
		p++
	st_case_397:
		switch data[p] {
		case 78:
			goto st398
		case 110:
			goto st398
		}
		goto tr368
	st398:
		p++
	st_case_398:
		switch data[p] {
		case 71:
			goto st399
		case 103:
			goto st399
		}
		goto tr368
	st399:
		p++
	st_case_399:
		switch data[p] {
		case 85:
			goto st400
		case 117:
			goto st400
		}
		goto tr368
	st400:
		p++
	st_case_400:
		switch data[p] {
		case 65:
			goto st401
		case 97:
			goto st401
		}
		goto tr368
	st401:
		p++
	st_case_401:
		switch data[p] {
		case 71:
			goto st402
		case 103:
			goto st402
		}
		goto tr368
	st402:
		p++
	st_case_402:
		switch data[p] {
		case 69:
			goto st403
		case 101:
			goto st403
		}
		goto tr368
	st403:
		p++
	st_case_403:
		switch data[p] {
		case 9:
			goto tr584
		case 32:
			goto tr584
		case 58:
			goto tr585
		}
		goto tr368
	st404:
		p++
	st_case_404:
		switch data[p] {
		case 78:
			goto st405
		case 110:
			goto st405
		}
		goto tr368
	st405:
		p++
	st_case_405:
		switch data[p] {
		case 71:
			goto st406
		case 103:
			goto st406
		}
		goto tr368
	st406:
		p++
	st_case_406:
		switch data[p] {
		case 84:
			goto st407
		case 116:
			goto st407
		}
		goto tr368
	st407:
		p++
	st_case_407:
		switch data[p] {
		case 72:
			goto st408
		case 104:
			goto st408
		}
		goto tr368
	st408:
		p++
	st_case_408:
		switch data[p] {
		case 9:
			goto st409
		case 32:
			goto st409
		case 58:
			goto st410
		}
		goto tr368
	st409:
		p++
	st_case_409:
		switch data[p] {
		case 9:
			goto st409
		case 32:
			goto st409
		case 58:
			goto st410
		}
		goto st0
	st410:
		p++
	st_case_410:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st410
		case 32:
			goto st410
		case 525:
			goto st412
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr592
		}
		goto st0
tr592:
//line msg_parse.rl:533
clen=0
//line msg_parse.rl:294

			clen = clen * 10 + (int(data[p]) - 0x30)
		
	goto st411
tr594:
//line msg_parse.rl:294

			clen = clen * 10 + (int(data[p]) - 0x30)
		
	goto st411
	st411:
		p++
	st_case_411:
//line msg_parse.go:11016
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st339
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr594
		}
		goto st0
	st412:
		p++
	st_case_412:
		if data[p] == 10 {
			goto st413
		}
		goto st0
	st413:
		p++
	st_case_413:
		switch data[p] {
		case 9:
			goto st414
		case 32:
			goto st414
		}
		goto st0
	st414:
		p++
	st_case_414:
		switch data[p] {
		case 9:
			goto st414
		case 32:
			goto st414
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr592
		}
		goto st0
	st415:
		p++
	st_case_415:
		switch data[p] {
		case 89:
			goto st416
		case 121:
			goto st416
		}
		goto tr368
	st416:
		p++
	st_case_416:
		switch data[p] {
		case 80:
			goto st417
		case 112:
			goto st417
		}
		goto tr368
	st417:
		p++
	st_case_417:
		switch data[p] {
		case 69:
			goto st418
		case 101:
			goto st418
		}
		goto tr368
	st418:
		p++
	st_case_418:
		switch data[p] {
		case 9:
			goto st331
		case 32:
			goto st331
		case 58:
			goto st332
		}
		goto tr368
	st419:
		p++
	st_case_419:
		switch data[p] {
		case 69:
			goto st420
		case 101:
			goto st420
		}
		goto tr368
	st420:
		p++
	st_case_420:
		switch data[p] {
		case 81:
			goto st421
		case 113:
			goto st421
		}
		goto tr368
	st421:
		p++
	st_case_421:
		switch data[p] {
		case 9:
			goto st422
		case 32:
			goto st422
		case 58:
			goto st423
		}
		goto tr368
	st422:
		p++
	st_case_422:
		switch data[p] {
		case 9:
			goto st422
		case 32:
			goto st422
		case 58:
			goto st423
		}
		goto st0
	st423:
		p++
	st_case_423:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st423
		case 32:
			goto st423
		case 525:
			goto st430
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr605
		}
		goto st0
tr605:
//line msg_parse.rl:302

			msg.CSeq = msg.CSeq * 10 + (int(data[p]) - 0x30)
		
	goto st424
	st424:
		p++
	st_case_424:
//line msg_parse.go:11178
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st425
		case 32:
			goto st425
		case 525:
			goto st427
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr605
		}
		goto st0
	st425:
		p++
	st_case_425:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st425
		case 32:
			goto st425
		case 33:
			goto tr609
		case 37:
			goto tr609
		case 39:
			goto tr609
		case 126:
			goto tr609
		case 525:
			goto st427
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr609
				}
			case _widec >= 42:
				goto tr609
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr609
				}
			case _widec >= 65:
				goto tr609
			}
		default:
			goto tr609
		}
		goto st0
tr609:
//line msg_parse.rl:103

			mark = p
		
	goto st426
	st426:
		p++
	st_case_426:
//line msg_parse.go:11256
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 33:
			goto st426
		case 37:
			goto st426
		case 39:
			goto st426
		case 126:
			goto st426
		case 269:
			goto tr611
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto st426
				}
			case _widec >= 42:
				goto st426
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st426
				}
			case _widec >= 65:
				goto st426
			}
		default:
			goto st426
		}
		goto st0
	st427:
		p++
	st_case_427:
		if data[p] == 10 {
			goto st428
		}
		goto st0
	st428:
		p++
	st_case_428:
		switch data[p] {
		case 9:
			goto st429
		case 32:
			goto st429
		}
		goto st0
	st429:
		p++
	st_case_429:
		switch data[p] {
		case 9:
			goto st429
		case 32:
			goto st429
		case 33:
			goto tr609
		case 37:
			goto tr609
		case 39:
			goto tr609
		case 126:
			goto tr609
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr609
				}
			case data[p] >= 42:
				goto tr609
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr609
				}
			case data[p] >= 65:
				goto tr609
			}
		default:
			goto tr609
		}
		goto st0
	st430:
		p++
	st_case_430:
		if data[p] == 10 {
			goto st431
		}
		goto st0
	st431:
		p++
	st_case_431:
		switch data[p] {
		case 9:
			goto st432
		case 32:
			goto st432
		}
		goto st0
	st432:
		p++
	st_case_432:
		switch data[p] {
		case 9:
			goto st432
		case 32:
			goto st432
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr605
		}
		goto st0
	st433:
		p++
	st_case_433:
		switch data[p] {
		case 65:
			goto st434
		case 97:
			goto st434
		}
		goto tr368
	st434:
		p++
	st_case_434:
		switch data[p] {
		case 84:
			goto st435
		case 116:
			goto st435
		}
		goto tr368
	st435:
		p++
	st_case_435:
		switch data[p] {
		case 69:
			goto st436
		case 101:
			goto st436
		}
		goto tr368
	st436:
		p++
	st_case_436:
		switch data[p] {
		case 9:
			goto tr619
		case 32:
			goto tr619
		case 58:
			goto tr620
		}
		goto tr368
	st437:
		p++
	st_case_437:
		switch data[p] {
		case 9:
			goto tr574
		case 32:
			goto tr574
		case 58:
			goto tr575
		case 82:
			goto st438
		case 86:
			goto st447
		case 88:
			goto st451
		case 114:
			goto st438
		case 118:
			goto st447
		case 120:
			goto st451
		}
		goto tr368
	st438:
		p++
	st_case_438:
		switch data[p] {
		case 82:
			goto st439
		case 114:
			goto st439
		}
		goto tr368
	st439:
		p++
	st_case_439:
		switch data[p] {
		case 79:
			goto st440
		case 111:
			goto st440
		}
		goto tr368
	st440:
		p++
	st_case_440:
		switch data[p] {
		case 82:
			goto st441
		case 114:
			goto st441
		}
		goto tr368
	st441:
		p++
	st_case_441:
		if data[p] == 45 {
			goto st442
		}
		goto tr368
	st442:
		p++
	st_case_442:
		switch data[p] {
		case 73:
			goto st443
		case 105:
			goto st443
		}
		goto tr368
	st443:
		p++
	st_case_443:
		switch data[p] {
		case 78:
			goto st444
		case 110:
			goto st444
		}
		goto tr368
	st444:
		p++
	st_case_444:
		switch data[p] {
		case 70:
			goto st445
		case 102:
			goto st445
		}
		goto tr368
	st445:
		p++
	st_case_445:
		switch data[p] {
		case 79:
			goto st446
		case 111:
			goto st446
		}
		goto tr368
	st446:
		p++
	st_case_446:
		switch data[p] {
		case 9:
			goto tr632
		case 32:
			goto tr632
		case 58:
			goto tr633
		}
		goto tr368
	st447:
		p++
	st_case_447:
		switch data[p] {
		case 69:
			goto st448
		case 101:
			goto st448
		}
		goto tr368
	st448:
		p++
	st_case_448:
		switch data[p] {
		case 78:
			goto st449
		case 110:
			goto st449
		}
		goto tr368
	st449:
		p++
	st_case_449:
		switch data[p] {
		case 84:
			goto st450
		case 116:
			goto st450
		}
		goto tr368
	st450:
		p++
	st_case_450:
		switch data[p] {
		case 9:
			goto tr637
		case 32:
			goto tr637
		case 58:
			goto tr638
		}
		goto tr368
	st451:
		p++
	st_case_451:
		switch data[p] {
		case 80:
			goto st452
		case 112:
			goto st452
		}
		goto tr368
	st452:
		p++
	st_case_452:
		switch data[p] {
		case 73:
			goto st453
		case 105:
			goto st453
		}
		goto tr368
	st453:
		p++
	st_case_453:
		switch data[p] {
		case 82:
			goto st454
		case 114:
			goto st454
		}
		goto tr368
	st454:
		p++
	st_case_454:
		switch data[p] {
		case 69:
			goto st455
		case 101:
			goto st455
		}
		goto tr368
	st455:
		p++
	st_case_455:
		switch data[p] {
		case 83:
			goto st456
		case 115:
			goto st456
		}
		goto tr368
	st456:
		p++
	st_case_456:
		switch data[p] {
		case 9:
			goto st457
		case 32:
			goto st457
		case 58:
			goto st458
		}
		goto tr368
	st457:
		p++
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
		p++
	st_case_458:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st458
		case 32:
			goto st458
		case 525:
			goto st460
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr646
		}
		goto st0
tr646:
//line msg_parse.rl:536
msg.Expires=0
//line msg_parse.rl:310

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
	goto st459
tr648:
//line msg_parse.rl:310

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
	goto st459
	st459:
		p++
	st_case_459:
//line msg_parse.go:11696
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st339
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr648
		}
		goto st0
	st460:
		p++
	st_case_460:
		if data[p] == 10 {
			goto st461
		}
		goto st0
	st461:
		p++
	st_case_461:
		switch data[p] {
		case 9:
			goto st462
		case 32:
			goto st462
		}
		goto st0
	st462:
		p++
	st_case_462:
		switch data[p] {
		case 9:
			goto st462
		case 32:
			goto st462
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr646
		}
		goto st0
	st463:
		p++
	st_case_463:
		switch data[p] {
		case 9:
			goto tr651
		case 32:
			goto tr651
		case 58:
			goto tr652
		case 82:
			goto st464
		case 114:
			goto st464
		}
		goto tr368
	st464:
		p++
	st_case_464:
		switch data[p] {
		case 79:
			goto st465
		case 111:
			goto st465
		}
		goto tr368
	st465:
		p++
	st_case_465:
		switch data[p] {
		case 77:
			goto st466
		case 109:
			goto st466
		}
		goto tr368
	st466:
		p++
	st_case_466:
		switch data[p] {
		case 9:
			goto tr651
		case 32:
			goto tr651
		case 58:
			goto tr652
		}
		goto tr368
	st467:
		p++
	st_case_467:
		switch data[p] {
		case 9:
			goto st351
		case 32:
			goto st351
		case 58:
			goto st352
		case 78:
			goto st468
		case 110:
			goto st468
		}
		goto tr368
	st468:
		p++
	st_case_468:
		if data[p] == 45 {
			goto st469
		}
		goto tr368
	st469:
		p++
	st_case_469:
		switch data[p] {
		case 82:
			goto st470
		case 114:
			goto st470
		}
		goto tr368
	st470:
		p++
	st_case_470:
		switch data[p] {
		case 69:
			goto st471
		case 101:
			goto st471
		}
		goto tr368
	st471:
		p++
	st_case_471:
		switch data[p] {
		case 80:
			goto st472
		case 112:
			goto st472
		}
		goto tr368
	st472:
		p++
	st_case_472:
		switch data[p] {
		case 76:
			goto st473
		case 108:
			goto st473
		}
		goto tr368
	st473:
		p++
	st_case_473:
		switch data[p] {
		case 89:
			goto st474
		case 121:
			goto st474
		}
		goto tr368
	st474:
		p++
	st_case_474:
		if data[p] == 45 {
			goto st475
		}
		goto tr368
	st475:
		p++
	st_case_475:
		switch data[p] {
		case 84:
			goto st476
		case 116:
			goto st476
		}
		goto tr368
	st476:
		p++
	st_case_476:
		switch data[p] {
		case 79:
			goto st477
		case 111:
			goto st477
		}
		goto tr368
	st477:
		p++
	st_case_477:
		switch data[p] {
		case 9:
			goto tr666
		case 32:
			goto tr666
		case 58:
			goto tr667
		}
		goto tr368
	st478:
		p++
	st_case_478:
		switch data[p] {
		case 9:
			goto tr668
		case 32:
			goto tr668
		case 58:
			goto tr669
		}
		goto tr368
	st479:
		p++
	st_case_479:
		switch data[p] {
		case 9:
			goto st480
		case 32:
			goto st480
		case 58:
			goto st481
		}
		goto tr368
	st480:
		p++
	st_case_480:
		switch data[p] {
		case 9:
			goto st480
		case 32:
			goto st480
		case 58:
			goto st481
		}
		goto st0
	st481:
		p++
	st_case_481:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st481
		case 32:
			goto st481
		case 525:
			goto st483
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr672
		}
		goto st0
tr672:
//line msg_parse.rl:533
clen=0
//line msg_parse.rl:294

			clen = clen * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:536
msg.Expires=0
//line msg_parse.rl:310

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:537
msg.MaxForwards=0
//line msg_parse.rl:314

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:538
msg.MinExpires=0
//line msg_parse.rl:318

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st482
tr674:
//line msg_parse.rl:294

			clen = clen * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:310

			msg.Expires = msg.Expires * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:314

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
//line msg_parse.rl:318

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st482
	st482:
		p++
	st_case_482:
//line msg_parse.go:12006
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st339
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr674
		}
		goto st0
	st483:
		p++
	st_case_483:
		if data[p] == 10 {
			goto st484
		}
		goto st0
	st484:
		p++
	st_case_484:
		switch data[p] {
		case 9:
			goto st485
		case 32:
			goto st485
		}
		goto st0
	st485:
		p++
	st_case_485:
		switch data[p] {
		case 9:
			goto st485
		case 32:
			goto st485
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr672
		}
		goto st0
	st486:
		p++
	st_case_486:
		switch data[p] {
		case 9:
			goto tr540
		case 32:
			goto tr540
		case 58:
			goto tr541
		case 65:
			goto st487
		case 73:
			goto st504
		case 97:
			goto st487
		case 105:
			goto st504
		}
		goto tr368
	st487:
		p++
	st_case_487:
		switch data[p] {
		case 88:
			goto st488
		case 120:
			goto st488
		}
		goto tr368
	st488:
		p++
	st_case_488:
		if data[p] == 45 {
			goto st489
		}
		goto tr368
	st489:
		p++
	st_case_489:
		switch data[p] {
		case 70:
			goto st490
		case 102:
			goto st490
		}
		goto tr368
	st490:
		p++
	st_case_490:
		switch data[p] {
		case 79:
			goto st491
		case 111:
			goto st491
		}
		goto tr368
	st491:
		p++
	st_case_491:
		switch data[p] {
		case 82:
			goto st492
		case 114:
			goto st492
		}
		goto tr368
	st492:
		p++
	st_case_492:
		switch data[p] {
		case 87:
			goto st493
		case 119:
			goto st493
		}
		goto tr368
	st493:
		p++
	st_case_493:
		switch data[p] {
		case 65:
			goto st494
		case 97:
			goto st494
		}
		goto tr368
	st494:
		p++
	st_case_494:
		switch data[p] {
		case 82:
			goto st495
		case 114:
			goto st495
		}
		goto tr368
	st495:
		p++
	st_case_495:
		switch data[p] {
		case 68:
			goto st496
		case 100:
			goto st496
		}
		goto tr368
	st496:
		p++
	st_case_496:
		switch data[p] {
		case 83:
			goto st497
		case 115:
			goto st497
		}
		goto tr368
	st497:
		p++
	st_case_497:
		switch data[p] {
		case 9:
			goto st498
		case 32:
			goto st498
		case 58:
			goto st499
		}
		goto tr368
	st498:
		p++
	st_case_498:
		switch data[p] {
		case 9:
			goto st498
		case 32:
			goto st498
		case 58:
			goto st499
		}
		goto st0
	st499:
		p++
	st_case_499:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st499
		case 32:
			goto st499
		case 525:
			goto st501
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr691
		}
		goto st0
tr691:
//line msg_parse.rl:537
msg.MaxForwards=0
//line msg_parse.rl:314

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
	goto st500
tr693:
//line msg_parse.rl:314

			msg.MaxForwards = msg.MaxForwards * 10 + (int(data[p]) - 0x30)
		
	goto st500
	st500:
		p++
	st_case_500:
//line msg_parse.go:12231
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st339
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr693
		}
		goto st0
	st501:
		p++
	st_case_501:
		if data[p] == 10 {
			goto st502
		}
		goto st0
	st502:
		p++
	st_case_502:
		switch data[p] {
		case 9:
			goto st503
		case 32:
			goto st503
		}
		goto st0
	st503:
		p++
	st_case_503:
		switch data[p] {
		case 9:
			goto st503
		case 32:
			goto st503
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr691
		}
		goto st0
	st504:
		p++
	st_case_504:
		switch data[p] {
		case 77:
			goto st505
		case 78:
			goto st515
		case 109:
			goto st505
		case 110:
			goto st515
		}
		goto tr368
	st505:
		p++
	st_case_505:
		switch data[p] {
		case 69:
			goto st506
		case 101:
			goto st506
		}
		goto tr368
	st506:
		p++
	st_case_506:
		if data[p] == 45 {
			goto st507
		}
		goto tr368
	st507:
		p++
	st_case_507:
		switch data[p] {
		case 86:
			goto st508
		case 118:
			goto st508
		}
		goto tr368
	st508:
		p++
	st_case_508:
		switch data[p] {
		case 69:
			goto st509
		case 101:
			goto st509
		}
		goto tr368
	st509:
		p++
	st_case_509:
		switch data[p] {
		case 82:
			goto st510
		case 114:
			goto st510
		}
		goto tr368
	st510:
		p++
	st_case_510:
		switch data[p] {
		case 83:
			goto st511
		case 115:
			goto st511
		}
		goto tr368
	st511:
		p++
	st_case_511:
		switch data[p] {
		case 73:
			goto st512
		case 105:
			goto st512
		}
		goto tr368
	st512:
		p++
	st_case_512:
		switch data[p] {
		case 79:
			goto st513
		case 111:
			goto st513
		}
		goto tr368
	st513:
		p++
	st_case_513:
		switch data[p] {
		case 78:
			goto st514
		case 110:
			goto st514
		}
		goto tr368
	st514:
		p++
	st_case_514:
		switch data[p] {
		case 9:
			goto tr707
		case 32:
			goto tr707
		case 58:
			goto tr708
		}
		goto tr368
	st515:
		p++
	st_case_515:
		if data[p] == 45 {
			goto st516
		}
		goto tr368
	st516:
		p++
	st_case_516:
		switch data[p] {
		case 69:
			goto st517
		case 101:
			goto st517
		}
		goto tr368
	st517:
		p++
	st_case_517:
		switch data[p] {
		case 88:
			goto st518
		case 120:
			goto st518
		}
		goto tr368
	st518:
		p++
	st_case_518:
		switch data[p] {
		case 80:
			goto st519
		case 112:
			goto st519
		}
		goto tr368
	st519:
		p++
	st_case_519:
		switch data[p] {
		case 73:
			goto st520
		case 105:
			goto st520
		}
		goto tr368
	st520:
		p++
	st_case_520:
		switch data[p] {
		case 82:
			goto st521
		case 114:
			goto st521
		}
		goto tr368
	st521:
		p++
	st_case_521:
		switch data[p] {
		case 69:
			goto st522
		case 101:
			goto st522
		}
		goto tr368
	st522:
		p++
	st_case_522:
		switch data[p] {
		case 83:
			goto st523
		case 115:
			goto st523
		}
		goto tr368
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
		goto tr368
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
			goto tr719
		}
		goto st0
tr719:
//line msg_parse.rl:538
msg.MinExpires=0
//line msg_parse.rl:318

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st526
tr721:
//line msg_parse.rl:318

			msg.MinExpires = msg.MinExpires * 10 + (int(data[p]) - 0x30)
		
	goto st526
	st526:
		p++
	st_case_526:
//line msg_parse.go:12529
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		if _widec == 269 {
			goto st339
		}
		if 48 <= _widec && _widec <= 57 {
			goto tr721
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
			goto tr719
		}
		goto st0
	st530:
		p++
	st_case_530:
		switch data[p] {
		case 9:
			goto tr637
		case 32:
			goto tr637
		case 58:
			goto tr638
		case 82:
			goto st531
		case 114:
			goto st531
		}
		goto tr368
	st531:
		p++
	st_case_531:
		switch data[p] {
		case 71:
			goto st532
		case 103:
			goto st532
		}
		goto tr368
	st532:
		p++
	st_case_532:
		switch data[p] {
		case 65:
			goto st533
		case 97:
			goto st533
		}
		goto tr368
	st533:
		p++
	st_case_533:
		switch data[p] {
		case 78:
			goto st534
		case 110:
			goto st534
		}
		goto tr368
	st534:
		p++
	st_case_534:
		switch data[p] {
		case 73:
			goto st535
		case 105:
			goto st535
		}
		goto tr368
	st535:
		p++
	st_case_535:
		switch data[p] {
		case 90:
			goto st536
		case 122:
			goto st536
		}
		goto tr368
	st536:
		p++
	st_case_536:
		switch data[p] {
		case 65:
			goto st537
		case 97:
			goto st537
		}
		goto tr368
	st537:
		p++
	st_case_537:
		switch data[p] {
		case 84:
			goto st538
		case 116:
			goto st538
		}
		goto tr368
	st538:
		p++
	st_case_538:
		switch data[p] {
		case 73:
			goto st539
		case 105:
			goto st539
		}
		goto tr368
	st539:
		p++
	st_case_539:
		switch data[p] {
		case 79:
			goto st540
		case 111:
			goto st540
		}
		goto tr368
	st540:
		p++
	st_case_540:
		switch data[p] {
		case 78:
			goto st541
		case 110:
			goto st541
		}
		goto tr368
	st541:
		p++
	st_case_541:
		switch data[p] {
		case 9:
			goto tr735
		case 32:
			goto tr735
		case 58:
			goto tr736
		}
		goto tr368
	st542:
		p++
	st_case_542:
		switch data[p] {
		case 45:
			goto st543
		case 82:
			goto st561
		case 114:
			goto st561
		}
		goto tr368
	st543:
		p++
	st_case_543:
		switch data[p] {
		case 65:
			goto st544
		case 97:
			goto st544
		}
		goto tr368
	st544:
		p++
	st_case_544:
		switch data[p] {
		case 83:
			goto st545
		case 115:
			goto st545
		}
		goto tr368
	st545:
		p++
	st_case_545:
		switch data[p] {
		case 83:
			goto st546
		case 115:
			goto st546
		}
		goto tr368
	st546:
		p++
	st_case_546:
		switch data[p] {
		case 69:
			goto st547
		case 101:
			goto st547
		}
		goto tr368
	st547:
		p++
	st_case_547:
		switch data[p] {
		case 82:
			goto st548
		case 114:
			goto st548
		}
		goto tr368
	st548:
		p++
	st_case_548:
		switch data[p] {
		case 84:
			goto st549
		case 116:
			goto st549
		}
		goto tr368
	st549:
		p++
	st_case_549:
		switch data[p] {
		case 69:
			goto st550
		case 101:
			goto st550
		}
		goto tr368
	st550:
		p++
	st_case_550:
		switch data[p] {
		case 68:
			goto st551
		case 100:
			goto st551
		}
		goto tr368
	st551:
		p++
	st_case_551:
		if data[p] == 45 {
			goto st552
		}
		goto tr368
	st552:
		p++
	st_case_552:
		switch data[p] {
		case 73:
			goto st553
		case 105:
			goto st553
		}
		goto tr368
	st553:
		p++
	st_case_553:
		switch data[p] {
		case 68:
			goto st554
		case 100:
			goto st554
		}
		goto tr368
	st554:
		p++
	st_case_554:
		switch data[p] {
		case 69:
			goto st555
		case 101:
			goto st555
		}
		goto tr368
	st555:
		p++
	st_case_555:
		switch data[p] {
		case 78:
			goto st556
		case 110:
			goto st556
		}
		goto tr368
	st556:
		p++
	st_case_556:
		switch data[p] {
		case 84:
			goto st557
		case 116:
			goto st557
		}
		goto tr368
	st557:
		p++
	st_case_557:
		switch data[p] {
		case 73:
			goto st558
		case 105:
			goto st558
		}
		goto tr368
	st558:
		p++
	st_case_558:
		switch data[p] {
		case 84:
			goto st559
		case 116:
			goto st559
		}
		goto tr368
	st559:
		p++
	st_case_559:
		switch data[p] {
		case 89:
			goto st560
		case 121:
			goto st560
		}
		goto tr368
	st560:
		p++
	st_case_560:
		switch data[p] {
		case 9:
			goto tr756
		case 32:
			goto tr756
		case 58:
			goto tr757
		}
		goto tr368
	st561:
		p++
	st_case_561:
		switch data[p] {
		case 73:
			goto st562
		case 79:
			goto st568
		case 105:
			goto st562
		case 111:
			goto st568
		}
		goto tr368
	st562:
		p++
	st_case_562:
		switch data[p] {
		case 79:
			goto st563
		case 111:
			goto st563
		}
		goto tr368
	st563:
		p++
	st_case_563:
		switch data[p] {
		case 82:
			goto st564
		case 114:
			goto st564
		}
		goto tr368
	st564:
		p++
	st_case_564:
		switch data[p] {
		case 73:
			goto st565
		case 105:
			goto st565
		}
		goto tr368
	st565:
		p++
	st_case_565:
		switch data[p] {
		case 84:
			goto st566
		case 116:
			goto st566
		}
		goto tr368
	st566:
		p++
	st_case_566:
		switch data[p] {
		case 89:
			goto st567
		case 121:
			goto st567
		}
		goto tr368
	st567:
		p++
	st_case_567:
		switch data[p] {
		case 9:
			goto tr765
		case 32:
			goto tr765
		case 58:
			goto tr766
		}
		goto tr368
	st568:
		p++
	st_case_568:
		switch data[p] {
		case 88:
			goto st569
		case 120:
			goto st569
		}
		goto tr368
	st569:
		p++
	st_case_569:
		switch data[p] {
		case 89:
			goto st570
		case 121:
			goto st570
		}
		goto tr368
	st570:
		p++
	st_case_570:
		if data[p] == 45 {
			goto st571
		}
		goto tr368
	st571:
		p++
	st_case_571:
		switch data[p] {
		case 65:
			goto st572
		case 82:
			goto st593
		case 97:
			goto st572
		case 114:
			goto st593
		}
		goto tr368
	st572:
		p++
	st_case_572:
		switch data[p] {
		case 85:
			goto st573
		case 117:
			goto st573
		}
		goto tr368
	st573:
		p++
	st_case_573:
		switch data[p] {
		case 84:
			goto st574
		case 116:
			goto st574
		}
		goto tr368
	st574:
		p++
	st_case_574:
		switch data[p] {
		case 72:
			goto st575
		case 104:
			goto st575
		}
		goto tr368
	st575:
		p++
	st_case_575:
		switch data[p] {
		case 69:
			goto st576
		case 79:
			goto st584
		case 101:
			goto st576
		case 111:
			goto st584
		}
		goto tr368
	st576:
		p++
	st_case_576:
		switch data[p] {
		case 78:
			goto st577
		case 110:
			goto st577
		}
		goto tr368
	st577:
		p++
	st_case_577:
		switch data[p] {
		case 84:
			goto st578
		case 116:
			goto st578
		}
		goto tr368
	st578:
		p++
	st_case_578:
		switch data[p] {
		case 73:
			goto st579
		case 105:
			goto st579
		}
		goto tr368
	st579:
		p++
	st_case_579:
		switch data[p] {
		case 67:
			goto st580
		case 99:
			goto st580
		}
		goto tr368
	st580:
		p++
	st_case_580:
		switch data[p] {
		case 65:
			goto st581
		case 97:
			goto st581
		}
		goto tr368
	st581:
		p++
	st_case_581:
		switch data[p] {
		case 84:
			goto st582
		case 116:
			goto st582
		}
		goto tr368
	st582:
		p++
	st_case_582:
		switch data[p] {
		case 69:
			goto st583
		case 101:
			goto st583
		}
		goto tr368
	st583:
		p++
	st_case_583:
		switch data[p] {
		case 9:
			goto tr784
		case 32:
			goto tr784
		case 58:
			goto tr785
		}
		goto tr368
	st584:
		p++
	st_case_584:
		switch data[p] {
		case 82:
			goto st585
		case 114:
			goto st585
		}
		goto tr368
	st585:
		p++
	st_case_585:
		switch data[p] {
		case 73:
			goto st586
		case 105:
			goto st586
		}
		goto tr368
	st586:
		p++
	st_case_586:
		switch data[p] {
		case 90:
			goto st587
		case 122:
			goto st587
		}
		goto tr368
	st587:
		p++
	st_case_587:
		switch data[p] {
		case 65:
			goto st588
		case 97:
			goto st588
		}
		goto tr368
	st588:
		p++
	st_case_588:
		switch data[p] {
		case 84:
			goto st589
		case 116:
			goto st589
		}
		goto tr368
	st589:
		p++
	st_case_589:
		switch data[p] {
		case 73:
			goto st590
		case 105:
			goto st590
		}
		goto tr368
	st590:
		p++
	st_case_590:
		switch data[p] {
		case 79:
			goto st591
		case 111:
			goto st591
		}
		goto tr368
	st591:
		p++
	st_case_591:
		switch data[p] {
		case 78:
			goto st592
		case 110:
			goto st592
		}
		goto tr368
	st592:
		p++
	st_case_592:
		switch data[p] {
		case 9:
			goto tr794
		case 32:
			goto tr794
		case 58:
			goto tr795
		}
		goto tr368
	st593:
		p++
	st_case_593:
		switch data[p] {
		case 69:
			goto st594
		case 101:
			goto st594
		}
		goto tr368
	st594:
		p++
	st_case_594:
		switch data[p] {
		case 81:
			goto st595
		case 113:
			goto st595
		}
		goto tr368
	st595:
		p++
	st_case_595:
		switch data[p] {
		case 85:
			goto st596
		case 117:
			goto st596
		}
		goto tr368
	st596:
		p++
	st_case_596:
		switch data[p] {
		case 73:
			goto st597
		case 105:
			goto st597
		}
		goto tr368
	st597:
		p++
	st_case_597:
		switch data[p] {
		case 82:
			goto st598
		case 114:
			goto st598
		}
		goto tr368
	st598:
		p++
	st_case_598:
		switch data[p] {
		case 69:
			goto st599
		case 101:
			goto st599
		}
		goto tr368
	st599:
		p++
	st_case_599:
		switch data[p] {
		case 9:
			goto tr802
		case 32:
			goto tr802
		case 58:
			goto tr803
		}
		goto tr368
	st600:
		p++
	st_case_600:
		switch data[p] {
		case 9:
			goto tr804
		case 32:
			goto tr804
		case 58:
			goto tr805
		case 69:
			goto st601
		case 79:
			goto st656
		case 101:
			goto st601
		case 111:
			goto st656
		}
		goto tr368
	st601:
		p++
	st_case_601:
		switch data[p] {
		case 67:
			goto st602
		case 70:
			goto st612
		case 77:
			goto st623
		case 80:
			goto st636
		case 81:
			goto st642
		case 84:
			goto st647
		case 99:
			goto st602
		case 102:
			goto st612
		case 109:
			goto st623
		case 112:
			goto st636
		case 113:
			goto st642
		case 116:
			goto st647
		}
		goto tr368
	st602:
		p++
	st_case_602:
		switch data[p] {
		case 79:
			goto st603
		case 111:
			goto st603
		}
		goto tr368
	st603:
		p++
	st_case_603:
		switch data[p] {
		case 82:
			goto st604
		case 114:
			goto st604
		}
		goto tr368
	st604:
		p++
	st_case_604:
		switch data[p] {
		case 68:
			goto st605
		case 100:
			goto st605
		}
		goto tr368
	st605:
		p++
	st_case_605:
		if data[p] == 45 {
			goto st606
		}
		goto tr368
	st606:
		p++
	st_case_606:
		switch data[p] {
		case 82:
			goto st607
		case 114:
			goto st607
		}
		goto tr368
	st607:
		p++
	st_case_607:
		switch data[p] {
		case 79:
			goto st608
		case 111:
			goto st608
		}
		goto tr368
	st608:
		p++
	st_case_608:
		switch data[p] {
		case 85:
			goto st609
		case 117:
			goto st609
		}
		goto tr368
	st609:
		p++
	st_case_609:
		switch data[p] {
		case 84:
			goto st610
		case 116:
			goto st610
		}
		goto tr368
	st610:
		p++
	st_case_610:
		switch data[p] {
		case 69:
			goto st611
		case 101:
			goto st611
		}
		goto tr368
	st611:
		p++
	st_case_611:
		switch data[p] {
		case 9:
			goto tr823
		case 32:
			goto tr823
		case 58:
			goto tr824
		}
		goto tr368
	st612:
		p++
	st_case_612:
		switch data[p] {
		case 69:
			goto st613
		case 101:
			goto st613
		}
		goto tr368
	st613:
		p++
	st_case_613:
		switch data[p] {
		case 82:
			goto st614
		case 114:
			goto st614
		}
		goto tr368
	st614:
		p++
	st_case_614:
		switch data[p] {
		case 45:
			goto st615
		case 82:
			goto st618
		case 114:
			goto st618
		}
		goto tr368
	st615:
		p++
	st_case_615:
		switch data[p] {
		case 84:
			goto st616
		case 116:
			goto st616
		}
		goto tr368
	st616:
		p++
	st_case_616:
		switch data[p] {
		case 79:
			goto st617
		case 111:
			goto st617
		}
		goto tr368
	st617:
		p++
	st_case_617:
		switch data[p] {
		case 9:
			goto tr804
		case 32:
			goto tr804
		case 58:
			goto tr805
		}
		goto tr368
	st618:
		p++
	st_case_618:
		switch data[p] {
		case 69:
			goto st619
		case 101:
			goto st619
		}
		goto tr368
	st619:
		p++
	st_case_619:
		switch data[p] {
		case 68:
			goto st620
		case 100:
			goto st620
		}
		goto tr368
	st620:
		p++
	st_case_620:
		if data[p] == 45 {
			goto st621
		}
		goto tr368
	st621:
		p++
	st_case_621:
		switch data[p] {
		case 66:
			goto st622
		case 98:
			goto st622
		}
		goto tr368
	st622:
		p++
	st_case_622:
		switch data[p] {
		case 89:
			goto st329
		case 121:
			goto st329
		}
		goto tr368
	st623:
		p++
	st_case_623:
		switch data[p] {
		case 79:
			goto st624
		case 111:
			goto st624
		}
		goto tr368
	st624:
		p++
	st_case_624:
		switch data[p] {
		case 84:
			goto st625
		case 116:
			goto st625
		}
		goto tr368
	st625:
		p++
	st_case_625:
		switch data[p] {
		case 69:
			goto st626
		case 101:
			goto st626
		}
		goto tr368
	st626:
		p++
	st_case_626:
		if data[p] == 45 {
			goto st627
		}
		goto tr368
	st627:
		p++
	st_case_627:
		switch data[p] {
		case 80:
			goto st628
		case 112:
			goto st628
		}
		goto tr368
	st628:
		p++
	st_case_628:
		switch data[p] {
		case 65:
			goto st629
		case 97:
			goto st629
		}
		goto tr368
	st629:
		p++
	st_case_629:
		switch data[p] {
		case 82:
			goto st630
		case 114:
			goto st630
		}
		goto tr368
	st630:
		p++
	st_case_630:
		switch data[p] {
		case 84:
			goto st631
		case 116:
			goto st631
		}
		goto tr368
	st631:
		p++
	st_case_631:
		switch data[p] {
		case 89:
			goto st632
		case 121:
			goto st632
		}
		goto tr368
	st632:
		p++
	st_case_632:
		if data[p] == 45 {
			goto st633
		}
		goto tr368
	st633:
		p++
	st_case_633:
		switch data[p] {
		case 73:
			goto st634
		case 105:
			goto st634
		}
		goto tr368
	st634:
		p++
	st_case_634:
		switch data[p] {
		case 68:
			goto st635
		case 100:
			goto st635
		}
		goto tr368
	st635:
		p++
	st_case_635:
		switch data[p] {
		case 9:
			goto tr847
		case 32:
			goto tr847
		case 58:
			goto tr848
		}
		goto tr368
	st636:
		p++
	st_case_636:
		switch data[p] {
		case 76:
			goto st637
		case 108:
			goto st637
		}
		goto tr368
	st637:
		p++
	st_case_637:
		switch data[p] {
		case 89:
			goto st638
		case 121:
			goto st638
		}
		goto tr368
	st638:
		p++
	st_case_638:
		if data[p] == 45 {
			goto st639
		}
		goto tr368
	st639:
		p++
	st_case_639:
		switch data[p] {
		case 84:
			goto st640
		case 116:
			goto st640
		}
		goto tr368
	st640:
		p++
	st_case_640:
		switch data[p] {
		case 79:
			goto st641
		case 111:
			goto st641
		}
		goto tr368
	st641:
		p++
	st_case_641:
		switch data[p] {
		case 9:
			goto tr854
		case 32:
			goto tr854
		case 58:
			goto tr855
		}
		goto tr368
	st642:
		p++
	st_case_642:
		switch data[p] {
		case 85:
			goto st643
		case 117:
			goto st643
		}
		goto tr368
	st643:
		p++
	st_case_643:
		switch data[p] {
		case 73:
			goto st644
		case 105:
			goto st644
		}
		goto tr368
	st644:
		p++
	st_case_644:
		switch data[p] {
		case 82:
			goto st645
		case 114:
			goto st645
		}
		goto tr368
	st645:
		p++
	st_case_645:
		switch data[p] {
		case 69:
			goto st646
		case 101:
			goto st646
		}
		goto tr368
	st646:
		p++
	st_case_646:
		switch data[p] {
		case 9:
			goto tr860
		case 32:
			goto tr860
		case 58:
			goto tr861
		}
		goto tr368
	st647:
		p++
	st_case_647:
		switch data[p] {
		case 82:
			goto st648
		case 114:
			goto st648
		}
		goto tr368
	st648:
		p++
	st_case_648:
		switch data[p] {
		case 89:
			goto st649
		case 121:
			goto st649
		}
		goto tr368
	st649:
		p++
	st_case_649:
		if data[p] == 45 {
			goto st650
		}
		goto tr368
	st650:
		p++
	st_case_650:
		switch data[p] {
		case 65:
			goto st651
		case 97:
			goto st651
		}
		goto tr368
	st651:
		p++
	st_case_651:
		switch data[p] {
		case 70:
			goto st652
		case 102:
			goto st652
		}
		goto tr368
	st652:
		p++
	st_case_652:
		switch data[p] {
		case 84:
			goto st653
		case 116:
			goto st653
		}
		goto tr368
	st653:
		p++
	st_case_653:
		switch data[p] {
		case 69:
			goto st654
		case 101:
			goto st654
		}
		goto tr368
	st654:
		p++
	st_case_654:
		switch data[p] {
		case 82:
			goto st655
		case 114:
			goto st655
		}
		goto tr368
	st655:
		p++
	st_case_655:
		switch data[p] {
		case 9:
			goto tr870
		case 32:
			goto tr870
		case 58:
			goto tr871
		}
		goto tr368
	st656:
		p++
	st_case_656:
		switch data[p] {
		case 85:
			goto st657
		case 117:
			goto st657
		}
		goto tr368
	st657:
		p++
	st_case_657:
		switch data[p] {
		case 84:
			goto st658
		case 116:
			goto st658
		}
		goto tr368
	st658:
		p++
	st_case_658:
		switch data[p] {
		case 69:
			goto st659
		case 101:
			goto st659
		}
		goto tr368
	st659:
		p++
	st_case_659:
		switch data[p] {
		case 9:
			goto tr875
		case 32:
			goto tr875
		case 58:
			goto tr876
		}
		goto tr368
	st660:
		p++
	st_case_660:
		switch data[p] {
		case 9:
			goto tr877
		case 32:
			goto tr877
		case 58:
			goto tr878
		case 69:
			goto st661
		case 85:
			goto st666
		case 101:
			goto st661
		case 117:
			goto st666
		}
		goto tr368
	st661:
		p++
	st_case_661:
		switch data[p] {
		case 82:
			goto st662
		case 114:
			goto st662
		}
		goto tr368
	st662:
		p++
	st_case_662:
		switch data[p] {
		case 86:
			goto st663
		case 118:
			goto st663
		}
		goto tr368
	st663:
		p++
	st_case_663:
		switch data[p] {
		case 69:
			goto st664
		case 101:
			goto st664
		}
		goto tr368
	st664:
		p++
	st_case_664:
		switch data[p] {
		case 82:
			goto st665
		case 114:
			goto st665
		}
		goto tr368
	st665:
		p++
	st_case_665:
		switch data[p] {
		case 9:
			goto tr885
		case 32:
			goto tr885
		case 58:
			goto tr886
		}
		goto tr368
	st666:
		p++
	st_case_666:
		switch data[p] {
		case 66:
			goto st667
		case 80:
			goto st672
		case 98:
			goto st667
		case 112:
			goto st672
		}
		goto tr368
	st667:
		p++
	st_case_667:
		switch data[p] {
		case 74:
			goto st668
		case 106:
			goto st668
		}
		goto tr368
	st668:
		p++
	st_case_668:
		switch data[p] {
		case 69:
			goto st669
		case 101:
			goto st669
		}
		goto tr368
	st669:
		p++
	st_case_669:
		switch data[p] {
		case 67:
			goto st670
		case 99:
			goto st670
		}
		goto tr368
	st670:
		p++
	st_case_670:
		switch data[p] {
		case 84:
			goto st671
		case 116:
			goto st671
		}
		goto tr368
	st671:
		p++
	st_case_671:
		switch data[p] {
		case 9:
			goto tr877
		case 32:
			goto tr877
		case 58:
			goto tr878
		}
		goto tr368
	st672:
		p++
	st_case_672:
		switch data[p] {
		case 80:
			goto st673
		case 112:
			goto st673
		}
		goto tr368
	st673:
		p++
	st_case_673:
		switch data[p] {
		case 79:
			goto st674
		case 111:
			goto st674
		}
		goto tr368
	st674:
		p++
	st_case_674:
		switch data[p] {
		case 82:
			goto st675
		case 114:
			goto st675
		}
		goto tr368
	st675:
		p++
	st_case_675:
		switch data[p] {
		case 84:
			goto st676
		case 116:
			goto st676
		}
		goto tr368
	st676:
		p++
	st_case_676:
		switch data[p] {
		case 69:
			goto st677
		case 101:
			goto st677
		}
		goto tr368
	st677:
		p++
	st_case_677:
		switch data[p] {
		case 68:
			goto st478
		case 100:
			goto st478
		}
		goto tr368
	st678:
		p++
	st_case_678:
		switch data[p] {
		case 9:
			goto tr898
		case 32:
			goto tr898
		case 58:
			goto tr899
		case 73:
			goto st679
		case 79:
			goto st687
		case 105:
			goto st679
		case 111:
			goto st687
		}
		goto tr368
	st679:
		p++
	st_case_679:
		switch data[p] {
		case 77:
			goto st680
		case 109:
			goto st680
		}
		goto tr368
	st680:
		p++
	st_case_680:
		switch data[p] {
		case 69:
			goto st681
		case 101:
			goto st681
		}
		goto tr368
	st681:
		p++
	st_case_681:
		switch data[p] {
		case 83:
			goto st682
		case 115:
			goto st682
		}
		goto tr368
	st682:
		p++
	st_case_682:
		switch data[p] {
		case 84:
			goto st683
		case 116:
			goto st683
		}
		goto tr368
	st683:
		p++
	st_case_683:
		switch data[p] {
		case 65:
			goto st684
		case 97:
			goto st684
		}
		goto tr368
	st684:
		p++
	st_case_684:
		switch data[p] {
		case 77:
			goto st685
		case 109:
			goto st685
		}
		goto tr368
	st685:
		p++
	st_case_685:
		switch data[p] {
		case 80:
			goto st686
		case 112:
			goto st686
		}
		goto tr368
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
		}
		goto tr368
	st687:
		p++
	st_case_687:
		switch data[p] {
		case 9:
			goto tr898
		case 32:
			goto tr898
		case 58:
			goto tr899
		}
		goto tr368
	st688:
		p++
	st_case_688:
		switch data[p] {
		case 9:
			goto tr911
		case 32:
			goto tr911
		case 58:
			goto tr912
		case 78:
			goto st689
		case 83:
			goto st699
		case 110:
			goto st689
		case 115:
			goto st699
		}
		goto tr368
	st689:
		p++
	st_case_689:
		switch data[p] {
		case 83:
			goto st690
		case 115:
			goto st690
		}
		goto tr368
	st690:
		p++
	st_case_690:
		switch data[p] {
		case 85:
			goto st691
		case 117:
			goto st691
		}
		goto tr368
	st691:
		p++
	st_case_691:
		switch data[p] {
		case 80:
			goto st692
		case 112:
			goto st692
		}
		goto tr368
	st692:
		p++
	st_case_692:
		switch data[p] {
		case 80:
			goto st693
		case 112:
			goto st693
		}
		goto tr368
	st693:
		p++
	st_case_693:
		switch data[p] {
		case 79:
			goto st694
		case 111:
			goto st694
		}
		goto tr368
	st694:
		p++
	st_case_694:
		switch data[p] {
		case 82:
			goto st695
		case 114:
			goto st695
		}
		goto tr368
	st695:
		p++
	st_case_695:
		switch data[p] {
		case 84:
			goto st696
		case 116:
			goto st696
		}
		goto tr368
	st696:
		p++
	st_case_696:
		switch data[p] {
		case 69:
			goto st697
		case 101:
			goto st697
		}
		goto tr368
	st697:
		p++
	st_case_697:
		switch data[p] {
		case 68:
			goto st698
		case 100:
			goto st698
		}
		goto tr368
	st698:
		p++
	st_case_698:
		switch data[p] {
		case 9:
			goto tr924
		case 32:
			goto tr924
		case 58:
			goto tr925
		}
		goto tr368
	st699:
		p++
	st_case_699:
		switch data[p] {
		case 69:
			goto st700
		case 101:
			goto st700
		}
		goto tr368
	st700:
		p++
	st_case_700:
		switch data[p] {
		case 82:
			goto st701
		case 114:
			goto st701
		}
		goto tr368
	st701:
		p++
	st_case_701:
		if data[p] == 45 {
			goto st702
		}
		goto tr368
	st702:
		p++
	st_case_702:
		switch data[p] {
		case 65:
			goto st703
		case 97:
			goto st703
		}
		goto tr368
	st703:
		p++
	st_case_703:
		switch data[p] {
		case 71:
			goto st704
		case 103:
			goto st704
		}
		goto tr368
	st704:
		p++
	st_case_704:
		switch data[p] {
		case 69:
			goto st705
		case 101:
			goto st705
		}
		goto tr368
	st705:
		p++
	st_case_705:
		switch data[p] {
		case 78:
			goto st706
		case 110:
			goto st706
		}
		goto tr368
	st706:
		p++
	st_case_706:
		switch data[p] {
		case 84:
			goto st707
		case 116:
			goto st707
		}
		goto tr368
	st707:
		p++
	st_case_707:
		switch data[p] {
		case 9:
			goto tr934
		case 32:
			goto tr934
		case 58:
			goto tr935
		}
		goto tr368
	st708:
		p++
	st_case_708:
		switch data[p] {
		case 9:
			goto st709
		case 32:
			goto st709
		case 58:
			goto st710
		case 73:
			goto st714
		case 105:
			goto st714
		}
		goto tr368
	st709:
		p++
	st_case_709:
		switch data[p] {
		case 9:
			goto st709
		case 32:
			goto st709
		case 58:
			goto st710
		}
		goto st0
	st710:
		p++
	st_case_710:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  lookAheadWSP(data, p, pe)  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st710
		case 32:
			goto st710
		case 269:
			goto tr939
		case 525:
			goto st711
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr939
			}
		default:
			goto tr939
		}
		goto st0
	st711:
		p++
	st_case_711:
		if data[p] == 10 {
			goto st712
		}
		goto st0
	st712:
		p++
	st_case_712:
		switch data[p] {
		case 9:
			goto st713
		case 32:
			goto st713
		}
		goto st0
	st713:
		p++
	st_case_713:
		switch data[p] {
		case 9:
			goto st713
		case 32:
			goto st713
		}
		goto tr939
	st714:
		p++
	st_case_714:
		switch data[p] {
		case 65:
			goto st715
		case 97:
			goto st715
		}
		goto tr368
	st715:
		p++
	st_case_715:
		switch data[p] {
		case 9:
			goto st709
		case 32:
			goto st709
		case 58:
			goto st710
		}
		goto tr368
	st716:
		p++
	st_case_716:
		switch data[p] {
		case 65:
			goto st717
		case 87:
			goto st723
		case 97:
			goto st717
		case 119:
			goto st723
		}
		goto tr368
	st717:
		p++
	st_case_717:
		switch data[p] {
		case 82:
			goto st718
		case 114:
			goto st718
		}
		goto tr368
	st718:
		p++
	st_case_718:
		switch data[p] {
		case 78:
			goto st719
		case 110:
			goto st719
		}
		goto tr368
	st719:
		p++
	st_case_719:
		switch data[p] {
		case 73:
			goto st720
		case 105:
			goto st720
		}
		goto tr368
	st720:
		p++
	st_case_720:
		switch data[p] {
		case 78:
			goto st721
		case 110:
			goto st721
		}
		goto tr368
	st721:
		p++
	st_case_721:
		switch data[p] {
		case 71:
			goto st722
		case 103:
			goto st722
		}
		goto tr368
	st722:
		p++
	st_case_722:
		switch data[p] {
		case 9:
			goto tr951
		case 32:
			goto tr951
		case 58:
			goto tr952
		}
		goto tr368
	st723:
		p++
	st_case_723:
		switch data[p] {
		case 87:
			goto st724
		case 119:
			goto st724
		}
		goto tr368
	st724:
		p++
	st_case_724:
		if data[p] == 45 {
			goto st725
		}
		goto tr368
	st725:
		p++
	st_case_725:
		switch data[p] {
		case 65:
			goto st726
		case 97:
			goto st726
		}
		goto tr368
	st726:
		p++
	st_case_726:
		switch data[p] {
		case 85:
			goto st727
		case 117:
			goto st727
		}
		goto tr368
	st727:
		p++
	st_case_727:
		switch data[p] {
		case 84:
			goto st728
		case 116:
			goto st728
		}
		goto tr368
	st728:
		p++
	st_case_728:
		switch data[p] {
		case 72:
			goto st729
		case 104:
			goto st729
		}
		goto tr368
	st729:
		p++
	st_case_729:
		switch data[p] {
		case 69:
			goto st730
		case 101:
			goto st730
		}
		goto tr368
	st730:
		p++
	st_case_730:
		switch data[p] {
		case 78:
			goto st731
		case 110:
			goto st731
		}
		goto tr368
	st731:
		p++
	st_case_731:
		switch data[p] {
		case 84:
			goto st732
		case 116:
			goto st732
		}
		goto tr368
	st732:
		p++
	st_case_732:
		switch data[p] {
		case 73:
			goto st733
		case 105:
			goto st733
		}
		goto tr368
	st733:
		p++
	st_case_733:
		switch data[p] {
		case 67:
			goto st734
		case 99:
			goto st734
		}
		goto tr368
	st734:
		p++
	st_case_734:
		switch data[p] {
		case 65:
			goto st735
		case 97:
			goto st735
		}
		goto tr368
	st735:
		p++
	st_case_735:
		switch data[p] {
		case 84:
			goto st736
		case 116:
			goto st736
		}
		goto tr368
	st736:
		p++
	st_case_736:
		switch data[p] {
		case 69:
			goto st737
		case 101:
			goto st737
		}
		goto tr368
	st737:
		p++
	st_case_737:
		switch data[p] {
		case 9:
			goto tr967
		case 32:
			goto tr967
		case 58:
			goto tr968
		}
		goto tr368
	st738:
		p++
	st_case_738:
		if data[p] == 10 {
			goto tr969
		}
		goto st0
	st_out:

	if p == eof {
		switch cs {
		case 247, 248, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299, 300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 345, 346, 347, 348, 349, 350, 359, 360, 361, 362, 363, 364, 365, 366, 367, 373, 374, 375, 376, 377, 378, 379, 380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399, 400, 401, 402, 403, 404, 405, 406, 407, 408, 415, 416, 417, 418, 419, 420, 421, 433, 434, 435, 436, 437, 438, 439, 440, 441, 442, 443, 444, 445, 446, 447, 448, 449, 450, 451, 452, 453, 454, 455, 456, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 486, 487, 488, 489, 490, 491, 492, 493, 494, 495, 496, 497, 504, 505, 506, 507, 508, 509, 510, 511, 512, 513, 514, 515, 516, 517, 518, 519, 520, 521, 522, 523, 530, 531, 532, 533, 534, 535, 536, 537, 538, 539, 540, 541, 542, 543, 544, 545, 546, 547, 548, 549, 550, 551, 552, 553, 554, 555, 556, 557, 558, 559, 560, 561, 562, 563, 564, 565, 566, 567, 568, 569, 570, 571, 572, 573, 574, 575, 576, 577, 578, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 592, 593, 594, 595, 596, 597, 598, 599, 600, 601, 602, 603, 604, 605, 606, 607, 608, 609, 610, 611, 612, 613, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 630, 631, 632, 633, 634, 635, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 647, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 674, 675, 676, 677, 678, 679, 680, 681, 682, 683, 684, 685, 686, 687, 688, 689, 690, 691, 692, 693, 694, 695, 696, 697, 698, 699, 700, 701, 702, 703, 704, 705, 706, 707, 708, 714, 715, 716, 717, 718, 719, 720, 721, 722, 723, 724, 725, 726, 727, 728, 729, 730, 731, 732, 733, 734, 735, 736, 737:
//line msg_parse.rl:219

			p--

			{goto st239 }
		
		case 751:
//line msg_parse.rl:284

			*addrp = addr
			addrp = &addr.Next
			addr = nil
		
//line msg_parse.go:14807
		}
	}

	_out: {}
	}

//line msg_parse.rl:593


	if cs < msg_first_final {
		if p == pe {
			return nil, MsgIncompleteError{data}
		} else {
			return nil, MsgParseError{
				Msg: data,
				Offset: p,
			}
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
