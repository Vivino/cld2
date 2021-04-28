package cld2

import "strings"

type Language interface {
	String() string
	Code() string
}

// Language is a single language.
// Note that the zero value is "ENGLISH".
type LanguageImpl uint16


func (l LanguageImpl) String() string {
	return strings.Replace(strings.Title(strings.ToLower(languageToCName[int(l)])), "_", " ", -1)
}

func (l LanguageImpl) Code() string {
	return languageToCode[int(l)]
}

var _ Language = LanguageImpl(1)

// NewLanguage supplies a safe way of returning a uint16
// to a Language.
// If an invalid id is supplied, UNKNOWN_LANGUAGE is returned.
func NewLanguage(id uint16) Language {
	if id >= uint16(NUM_LANGUAGES) {
		return UNKNOWN_LANGUAGE
	}
	return LanguageImpl(id)
}

// LanguageFromCode returns the language associated with the
// code. Returns UNKNOWN_LANGUAGE if the code isn't known.
func LanguageFromCode(code string) Language {
	if l, ok := codeToLanguage[code]; ok {
		return l
	}
	return UNKNOWN_LANGUAGE
}

// Copied from "generated_language.h"
const (
	ENGLISH                  LanguageImpl = 0   // en
	DANISH                   LanguageImpl = 1   // da
	DUTCH                    LanguageImpl = 2   // nl
	FINNISH                  LanguageImpl = 3   // fi
	FRENCH                   LanguageImpl = 4   // fr
	GERMAN                   LanguageImpl = 5   // de
	HEBREW                   LanguageImpl = 6   // iw
	ITALIAN                  LanguageImpl = 7   // it
	JAPANESE                 LanguageImpl = 8   // ja
	KOREAN                   LanguageImpl = 9   // ko
	NORWEGIAN                LanguageImpl = 10  // no
	POLISH                   LanguageImpl = 11  // pl
	PORTUGUESE               LanguageImpl = 12  // pt
	RUSSIAN                  LanguageImpl = 13  // ru
	SPANISH                  LanguageImpl = 14  // es
	SWEDISH                  LanguageImpl = 15  // sv
	CHINESE                  LanguageImpl = 16  // zh
	CZECH                    LanguageImpl = 17  // cs
	GREEK                    LanguageImpl = 18  // el
	ICELANDIC                LanguageImpl = 19  // is
	LATVIAN                  LanguageImpl = 20  // lv
	LITHUANIAN               LanguageImpl = 21  // lt
	ROMANIAN                 LanguageImpl = 22  // ro
	HUNGARIAN                LanguageImpl = 23  // hu
	ESTONIAN                 LanguageImpl = 24  // et
	TG_UNKNOWN_LANGUAGE      LanguageImpl = 25  // xxx
	UNKNOWN_LANGUAGE         LanguageImpl = 26  // un
	BULGARIAN                LanguageImpl = 27  // bg
	CROATIAN                 LanguageImpl = 28  // hr
	SERBIAN                  LanguageImpl = 29  // sr
	IRISH                    LanguageImpl = 30  // ga
	GALICIAN                 LanguageImpl = 31  // gl
	TAGALOG                  LanguageImpl = 32  // tl
	TURKISH                  LanguageImpl = 33  // tr
	UKRAINIAN                LanguageImpl = 34  // uk
	HINDI                    LanguageImpl = 35  // hi
	MACEDONIAN               LanguageImpl = 36  // mk
	BENGALI                  LanguageImpl = 37  // bn
	INDONESIAN               LanguageImpl = 38  // id
	LATIN                    LanguageImpl = 39  // la
	MALAY                    LanguageImpl = 40  // ms
	MALAYALAM                LanguageImpl = 41  // ml
	WELSH                    LanguageImpl = 42  // cy
	NEPALI                   LanguageImpl = 43  // ne
	TELUGU                   LanguageImpl = 44  // te
	ALBANIAN                 LanguageImpl = 45  // sq
	TAMIL                    LanguageImpl = 46  // ta
	BELARUSIAN               LanguageImpl = 47  // be
	JAVANESE                 LanguageImpl = 48  // jw
	OCCITAN                  LanguageImpl = 49  // oc
	URDU                     LanguageImpl = 50  // ur
	BIHARI                   LanguageImpl = 51  // bh
	GUJARATI                 LanguageImpl = 52  // gu
	THAI                     LanguageImpl = 53  // th
	ARABIC                   LanguageImpl = 54  // ar
	CATALAN                  LanguageImpl = 55  // ca
	ESPERANTO                LanguageImpl = 56  // eo
	BASQUE                   LanguageImpl = 57  // eu
	INTERLINGUA              LanguageImpl = 58  // ia
	KANNADA                  LanguageImpl = 59  // kn
	PUNJABI                  LanguageImpl = 60  // pa
	SCOTS_GAELIC             LanguageImpl = 61  // gd
	SWAHILI                  LanguageImpl = 62  // sw
	SLOVENIAN                LanguageImpl = 63  // sl
	MARATHI                  LanguageImpl = 64  // mr
	MALTESE                  LanguageImpl = 65  // mt
	VIETNAMESE               LanguageImpl = 66  // vi
	FRISIAN                  LanguageImpl = 67  // fy
	SLOVAK                   LanguageImpl = 68  // sk
	CHINESE_T                LanguageImpl = 69  // zh-Hant
	FAROESE                  LanguageImpl = 70  // fo
	SUNDANESE                LanguageImpl = 71  // su
	UZBEK                    LanguageImpl = 72  // uz
	AMHARIC                  LanguageImpl = 73  // am
	AZERBAIJANI              LanguageImpl = 74  // az
	GEORGIAN                 LanguageImpl = 75  // ka
	TIGRINYA                 LanguageImpl = 76  // ti
	PERSIAN                  LanguageImpl = 77  // fa
	BOSNIAN                  LanguageImpl = 78  // bs
	SINHALESE                LanguageImpl = 79  // si
	NORWEGIAN_N              LanguageImpl = 80  // nn
	X_81                     LanguageImpl = 81  //
	X_82                     LanguageImpl = 82  //
	XHOSA                    LanguageImpl = 83  // xh
	ZULU                     LanguageImpl = 84  // zu
	GUARANI                  LanguageImpl = 85  // gn
	SESOTHO                  LanguageImpl = 86  // st
	TURKMEN                  LanguageImpl = 87  // tk
	KYRGYZ                   LanguageImpl = 88  // ky
	BRETON                   LanguageImpl = 89  // br
	TWI                      LanguageImpl = 90  // tw
	YIDDISH                  LanguageImpl = 91  // yi
	X_92                     LanguageImpl = 92  //
	SOMALI                   LanguageImpl = 93  // so
	UIGHUR                   LanguageImpl = 94  // ug
	KURDISH                  LanguageImpl = 95  // ku
	MONGOLIAN                LanguageImpl = 96  // mn
	ARMENIAN                 LanguageImpl = 97  // hy
	LAOTHIAN                 LanguageImpl = 98  // lo
	SINDHI                   LanguageImpl = 99  // sd
	RHAETO_ROMANCE           LanguageImpl = 100 // rm
	AFRIKAANS                LanguageImpl = 101 // af
	LUXEMBOURGISH            LanguageImpl = 102 // lb
	BURMESE                  LanguageImpl = 103 // my
	KHMER                    LanguageImpl = 104 // km
	TIBETAN                  LanguageImpl = 105 // bo
	DHIVEHI                  LanguageImpl = 106 // dv
	CHEROKEE                 LanguageImpl = 107 // chr
	SYRIAC                   LanguageImpl = 108 // syr
	LIMBU                    LanguageImpl = 109 // lif
	ORIYA                    LanguageImpl = 110 // or
	ASSAMESE                 LanguageImpl = 111 // as
	CORSICAN                 LanguageImpl = 112 // co
	INTERLINGUE              LanguageImpl = 113 // ie
	KAZAKH                   LanguageImpl = 114 // kk
	LINGALA                  LanguageImpl = 115 // ln
	X_116                    LanguageImpl = 116 //
	PASHTO                   LanguageImpl = 117 // ps
	QUECHUA                  LanguageImpl = 118 // qu
	SHONA                    LanguageImpl = 119 // sn
	TAJIK                    LanguageImpl = 120 // tg
	TATAR                    LanguageImpl = 121 // tt
	TONGA                    LanguageImpl = 122 // to
	YORUBA                   LanguageImpl = 123 // yo
	X_124                    LanguageImpl = 124 //
	X_125                    LanguageImpl = 125 //
	X_126                    LanguageImpl = 126 //
	X_127                    LanguageImpl = 127 //
	MAORI                    LanguageImpl = 128 // mi
	WOLOF                    LanguageImpl = 129 // wo
	ABKHAZIAN                LanguageImpl = 130 // ab
	AFAR                     LanguageImpl = 131 // aa
	AYMARA                   LanguageImpl = 132 // ay
	BASHKIR                  LanguageImpl = 133 // ba
	BISLAMA                  LanguageImpl = 134 // bi
	DZONGKHA                 LanguageImpl = 135 // dz
	FIJIAN                   LanguageImpl = 136 // fj
	GREENLANDIC              LanguageImpl = 137 // kl
	HAUSA                    LanguageImpl = 138 // ha
	HAITIAN_CREOLE           LanguageImpl = 139 // ht
	INUPIAK                  LanguageImpl = 140 // ik
	INUKTITUT                LanguageImpl = 141 // iu
	KASHMIRI                 LanguageImpl = 142 // ks
	KINYARWANDA              LanguageImpl = 143 // rw
	MALAGASY                 LanguageImpl = 144 // mg
	NAURU                    LanguageImpl = 145 // na
	OROMO                    LanguageImpl = 146 // om
	RUNDI                    LanguageImpl = 147 // rn
	SAMOAN                   LanguageImpl = 148 // sm
	SANGO                    LanguageImpl = 149 // sg
	SANSKRIT                 LanguageImpl = 150 // sa
	SISWANT                  LanguageImpl = 151 // ss
	TSONGA                   LanguageImpl = 152 // ts
	TSWANA                   LanguageImpl = 153 // tn
	VOLAPUK                  LanguageImpl = 154 // vo
	ZHUANG                   LanguageImpl = 155 // za
	KHASI                    LanguageImpl = 156 // kha
	SCOTS                    LanguageImpl = 157 // sco
	GANDA                    LanguageImpl = 158 // lg
	MANX                     LanguageImpl = 159 // gv
	MONTENEGRIN              LanguageImpl = 160 // sr-ME
	AKAN                     LanguageImpl = 161 // ak
	IGBO                     LanguageImpl = 162 // ig
	MAURITIAN_CREOLE         LanguageImpl = 163 // mfe
	HAWAIIAN                 LanguageImpl = 164 // haw
	CEBUANO                  LanguageImpl = 165 // ceb
	EWE                      LanguageImpl = 166 // ee
	GA                       LanguageImpl = 167 // gaa
	HMONG                    LanguageImpl = 168 // blu
	KRIO                     LanguageImpl = 169 // kri
	LOZI                     LanguageImpl = 170 // loz
	LUBA_LULUA               LanguageImpl = 171 // lua
	LUO_KENYA_AND_TANZANIA   LanguageImpl = 172 // luo
	NEWARI                   LanguageImpl = 173 // new
	NYANJA                   LanguageImpl = 174 // ny
	OSSETIAN                 LanguageImpl = 175 // os
	PAMPANGA                 LanguageImpl = 176 // pam
	PEDI                     LanguageImpl = 177 // nso
	RAJASTHANI               LanguageImpl = 178 // raj
	SESELWA                  LanguageImpl = 179 // crs
	TUMBUKA                  LanguageImpl = 180 // tum
	VENDA                    LanguageImpl = 181 // ve
	WARAY_PHILIPPINES        LanguageImpl = 182 // war
	X_183                    LanguageImpl = 183 //
	X_184                    LanguageImpl = 184 //
	X_185                    LanguageImpl = 185 //
	X_186                    LanguageImpl = 186 //
	X_187                    LanguageImpl = 187 //
	X_188                    LanguageImpl = 188 //
	X_189                    LanguageImpl = 189 //
	X_190                    LanguageImpl = 190 //
	X_191                    LanguageImpl = 191 //
	X_192                    LanguageImpl = 192 //
	X_193                    LanguageImpl = 193 //
	X_194                    LanguageImpl = 194 //
	X_195                    LanguageImpl = 195 //
	X_196                    LanguageImpl = 196 //
	X_197                    LanguageImpl = 197 //
	X_198                    LanguageImpl = 198 //
	X_199                    LanguageImpl = 199 //
	X_200                    LanguageImpl = 200 //
	X_201                    LanguageImpl = 201 //
	X_202                    LanguageImpl = 202 //
	X_203                    LanguageImpl = 203 //
	X_204                    LanguageImpl = 204 //
	X_205                    LanguageImpl = 205 //
	X_206                    LanguageImpl = 206 //
	X_207                    LanguageImpl = 207 //
	X_208                    LanguageImpl = 208 //
	X_209                    LanguageImpl = 209 //
	X_210                    LanguageImpl = 210 //
	X_211                    LanguageImpl = 211 //
	X_212                    LanguageImpl = 212 //
	X_213                    LanguageImpl = 213 //
	X_214                    LanguageImpl = 214 //
	X_215                    LanguageImpl = 215 //
	X_216                    LanguageImpl = 216 //
	X_217                    LanguageImpl = 217 //
	X_218                    LanguageImpl = 218 //
	X_219                    LanguageImpl = 219 //
	X_220                    LanguageImpl = 220 //
	X_221                    LanguageImpl = 221 //
	X_222                    LanguageImpl = 222 //
	X_223                    LanguageImpl = 223 //
	X_224                    LanguageImpl = 224 //
	X_225                    LanguageImpl = 225 //
	X_226                    LanguageImpl = 226 //
	X_227                    LanguageImpl = 227 //
	X_228                    LanguageImpl = 228 //
	X_229                    LanguageImpl = 229 //
	X_230                    LanguageImpl = 230 //
	X_231                    LanguageImpl = 231 //
	X_232                    LanguageImpl = 232 //
	X_233                    LanguageImpl = 233 //
	X_234                    LanguageImpl = 234 //
	X_235                    LanguageImpl = 235 //
	X_236                    LanguageImpl = 236 //
	X_237                    LanguageImpl = 237 //
	X_238                    LanguageImpl = 238 //
	X_239                    LanguageImpl = 239 //
	X_240                    LanguageImpl = 240 //
	X_241                    LanguageImpl = 241 //
	X_242                    LanguageImpl = 242 //
	X_243                    LanguageImpl = 243 //
	X_244                    LanguageImpl = 244 //
	X_245                    LanguageImpl = 245 //
	X_246                    LanguageImpl = 246 //
	X_247                    LanguageImpl = 247 //
	X_248                    LanguageImpl = 248 //
	X_249                    LanguageImpl = 249 //
	X_250                    LanguageImpl = 250 //
	X_251                    LanguageImpl = 251 //
	X_252                    LanguageImpl = 252 //
	X_253                    LanguageImpl = 253 //
	X_254                    LanguageImpl = 254 //
	X_255                    LanguageImpl = 255 //
	X_256                    LanguageImpl = 256 //
	X_257                    LanguageImpl = 257 //
	X_258                    LanguageImpl = 258 //
	X_259                    LanguageImpl = 259 //
	X_260                    LanguageImpl = 260 //
	X_261                    LanguageImpl = 261 //
	X_262                    LanguageImpl = 262 //
	X_263                    LanguageImpl = 263 //
	X_264                    LanguageImpl = 264 //
	X_265                    LanguageImpl = 265 //
	X_266                    LanguageImpl = 266 //
	X_267                    LanguageImpl = 267 //
	X_268                    LanguageImpl = 268 //
	X_269                    LanguageImpl = 269 //
	X_270                    LanguageImpl = 270 //
	X_271                    LanguageImpl = 271 //
	X_272                    LanguageImpl = 272 //
	X_273                    LanguageImpl = 273 //
	X_274                    LanguageImpl = 274 //
	X_275                    LanguageImpl = 275 //
	X_276                    LanguageImpl = 276 //
	X_277                    LanguageImpl = 277 //
	X_278                    LanguageImpl = 278 //
	X_279                    LanguageImpl = 279 //
	X_280                    LanguageImpl = 280 //
	X_281                    LanguageImpl = 281 //
	X_282                    LanguageImpl = 282 //
	X_283                    LanguageImpl = 283 //
	X_284                    LanguageImpl = 284 //
	X_285                    LanguageImpl = 285 //
	X_286                    LanguageImpl = 286 //
	X_287                    LanguageImpl = 287 //
	X_288                    LanguageImpl = 288 //
	X_289                    LanguageImpl = 289 //
	X_290                    LanguageImpl = 290 //
	X_291                    LanguageImpl = 291 //
	X_292                    LanguageImpl = 292 //
	X_293                    LanguageImpl = 293 //
	X_294                    LanguageImpl = 294 //
	X_295                    LanguageImpl = 295 //
	X_296                    LanguageImpl = 296 //
	X_297                    LanguageImpl = 297 //
	X_298                    LanguageImpl = 298 //
	X_299                    LanguageImpl = 299 //
	X_300                    LanguageImpl = 300 //
	X_301                    LanguageImpl = 301 //
	X_302                    LanguageImpl = 302 //
	X_303                    LanguageImpl = 303 //
	X_304                    LanguageImpl = 304 //
	X_305                    LanguageImpl = 305 //
	X_306                    LanguageImpl = 306 //
	X_307                    LanguageImpl = 307 //
	X_308                    LanguageImpl = 308 //
	X_309                    LanguageImpl = 309 //
	X_310                    LanguageImpl = 310 //
	X_311                    LanguageImpl = 311 //
	X_312                    LanguageImpl = 312 //
	X_313                    LanguageImpl = 313 //
	X_314                    LanguageImpl = 314 //
	X_315                    LanguageImpl = 315 //
	X_316                    LanguageImpl = 316 //
	X_317                    LanguageImpl = 317 //
	X_318                    LanguageImpl = 318 //
	X_319                    LanguageImpl = 319 //
	X_320                    LanguageImpl = 320 //
	X_321                    LanguageImpl = 321 //
	X_322                    LanguageImpl = 322 //
	X_323                    LanguageImpl = 323 //
	X_324                    LanguageImpl = 324 //
	X_325                    LanguageImpl = 325 //
	X_326                    LanguageImpl = 326 //
	X_327                    LanguageImpl = 327 //
	X_328                    LanguageImpl = 328 //
	X_329                    LanguageImpl = 329 //
	X_330                    LanguageImpl = 330 //
	X_331                    LanguageImpl = 331 //
	X_332                    LanguageImpl = 332 //
	X_333                    LanguageImpl = 333 //
	X_334                    LanguageImpl = 334 //
	X_335                    LanguageImpl = 335 //
	X_336                    LanguageImpl = 336 //
	X_337                    LanguageImpl = 337 //
	X_338                    LanguageImpl = 338 //
	X_339                    LanguageImpl = 339 //
	X_340                    LanguageImpl = 340 //
	X_341                    LanguageImpl = 341 //
	X_342                    LanguageImpl = 342 //
	X_343                    LanguageImpl = 343 //
	X_344                    LanguageImpl = 344 //
	X_345                    LanguageImpl = 345 //
	X_346                    LanguageImpl = 346 //
	X_347                    LanguageImpl = 347 //
	X_348                    LanguageImpl = 348 //
	X_349                    LanguageImpl = 349 //
	X_350                    LanguageImpl = 350 //
	X_351                    LanguageImpl = 351 //
	X_352                    LanguageImpl = 352 //
	X_353                    LanguageImpl = 353 //
	X_354                    LanguageImpl = 354 //
	X_355                    LanguageImpl = 355 //
	X_356                    LanguageImpl = 356 //
	X_357                    LanguageImpl = 357 //
	X_358                    LanguageImpl = 358 //
	X_359                    LanguageImpl = 359 //
	X_360                    LanguageImpl = 360 //
	X_361                    LanguageImpl = 361 //
	X_362                    LanguageImpl = 362 //
	X_363                    LanguageImpl = 363 //
	X_364                    LanguageImpl = 364 //
	X_365                    LanguageImpl = 365 //
	X_366                    LanguageImpl = 366 //
	X_367                    LanguageImpl = 367 //
	X_368                    LanguageImpl = 368 //
	X_369                    LanguageImpl = 369 //
	X_370                    LanguageImpl = 370 //
	X_371                    LanguageImpl = 371 //
	X_372                    LanguageImpl = 372 //
	X_373                    LanguageImpl = 373 //
	X_374                    LanguageImpl = 374 //
	X_375                    LanguageImpl = 375 //
	X_376                    LanguageImpl = 376 //
	X_377                    LanguageImpl = 377 //
	X_378                    LanguageImpl = 378 //
	X_379                    LanguageImpl = 379 //
	X_380                    LanguageImpl = 380 //
	X_381                    LanguageImpl = 381 //
	X_382                    LanguageImpl = 382 //
	X_383                    LanguageImpl = 383 //
	X_384                    LanguageImpl = 384 //
	X_385                    LanguageImpl = 385 //
	X_386                    LanguageImpl = 386 //
	X_387                    LanguageImpl = 387 //
	X_388                    LanguageImpl = 388 //
	X_389                    LanguageImpl = 389 //
	X_390                    LanguageImpl = 390 //
	X_391                    LanguageImpl = 391 //
	X_392                    LanguageImpl = 392 //
	X_393                    LanguageImpl = 393 //
	X_394                    LanguageImpl = 394 //
	X_395                    LanguageImpl = 395 //
	X_396                    LanguageImpl = 396 //
	X_397                    LanguageImpl = 397 //
	X_398                    LanguageImpl = 398 //
	X_399                    LanguageImpl = 399 //
	X_400                    LanguageImpl = 400 //
	X_401                    LanguageImpl = 401 //
	X_402                    LanguageImpl = 402 //
	X_403                    LanguageImpl = 403 //
	X_404                    LanguageImpl = 404 //
	X_405                    LanguageImpl = 405 //
	X_406                    LanguageImpl = 406 //
	X_407                    LanguageImpl = 407 //
	X_408                    LanguageImpl = 408 //
	X_409                    LanguageImpl = 409 //
	X_410                    LanguageImpl = 410 //
	X_411                    LanguageImpl = 411 //
	X_412                    LanguageImpl = 412 //
	X_413                    LanguageImpl = 413 //
	X_414                    LanguageImpl = 414 //
	X_415                    LanguageImpl = 415 //
	X_416                    LanguageImpl = 416 //
	X_417                    LanguageImpl = 417 //
	X_418                    LanguageImpl = 418 //
	X_419                    LanguageImpl = 419 //
	X_420                    LanguageImpl = 420 //
	X_421                    LanguageImpl = 421 //
	X_422                    LanguageImpl = 422 //
	X_423                    LanguageImpl = 423 //
	X_424                    LanguageImpl = 424 //
	X_425                    LanguageImpl = 425 //
	X_426                    LanguageImpl = 426 //
	X_427                    LanguageImpl = 427 //
	X_428                    LanguageImpl = 428 //
	X_429                    LanguageImpl = 429 //
	X_430                    LanguageImpl = 430 //
	X_431                    LanguageImpl = 431 //
	X_432                    LanguageImpl = 432 //
	X_433                    LanguageImpl = 433 //
	X_434                    LanguageImpl = 434 //
	X_435                    LanguageImpl = 435 //
	X_436                    LanguageImpl = 436 //
	X_437                    LanguageImpl = 437 //
	X_438                    LanguageImpl = 438 //
	X_439                    LanguageImpl = 439 //
	X_440                    LanguageImpl = 440 //
	X_441                    LanguageImpl = 441 //
	X_442                    LanguageImpl = 442 //
	X_443                    LanguageImpl = 443 //
	X_444                    LanguageImpl = 444 //
	X_445                    LanguageImpl = 445 //
	X_446                    LanguageImpl = 446 //
	X_447                    LanguageImpl = 447 //
	X_448                    LanguageImpl = 448 //
	X_449                    LanguageImpl = 449 //
	X_450                    LanguageImpl = 450 //
	X_451                    LanguageImpl = 451 //
	X_452                    LanguageImpl = 452 //
	X_453                    LanguageImpl = 453 //
	X_454                    LanguageImpl = 454 //
	X_455                    LanguageImpl = 455 //
	X_456                    LanguageImpl = 456 //
	X_457                    LanguageImpl = 457 //
	X_458                    LanguageImpl = 458 //
	X_459                    LanguageImpl = 459 //
	X_460                    LanguageImpl = 460 //
	X_461                    LanguageImpl = 461 //
	X_462                    LanguageImpl = 462 //
	X_463                    LanguageImpl = 463 //
	X_464                    LanguageImpl = 464 //
	X_465                    LanguageImpl = 465 //
	X_466                    LanguageImpl = 466 //
	X_467                    LanguageImpl = 467 //
	X_468                    LanguageImpl = 468 //
	X_469                    LanguageImpl = 469 //
	X_470                    LanguageImpl = 470 //
	X_471                    LanguageImpl = 471 //
	X_472                    LanguageImpl = 472 //
	X_473                    LanguageImpl = 473 //
	X_474                    LanguageImpl = 474 //
	X_475                    LanguageImpl = 475 //
	X_476                    LanguageImpl = 476 //
	X_477                    LanguageImpl = 477 //
	X_478                    LanguageImpl = 478 //
	X_479                    LanguageImpl = 479 //
	X_480                    LanguageImpl = 480 //
	X_481                    LanguageImpl = 481 //
	X_482                    LanguageImpl = 482 //
	X_483                    LanguageImpl = 483 //
	X_484                    LanguageImpl = 484 //
	X_485                    LanguageImpl = 485 //
	X_486                    LanguageImpl = 486 //
	X_487                    LanguageImpl = 487 //
	X_488                    LanguageImpl = 488 //
	X_489                    LanguageImpl = 489 //
	X_490                    LanguageImpl = 490 //
	X_491                    LanguageImpl = 491 //
	X_492                    LanguageImpl = 492 //
	X_493                    LanguageImpl = 493 //
	X_494                    LanguageImpl = 494 //
	X_495                    LanguageImpl = 495 //
	X_496                    LanguageImpl = 496 //
	X_497                    LanguageImpl = 497 //
	X_498                    LanguageImpl = 498 //
	X_499                    LanguageImpl = 499 //
	X_500                    LanguageImpl = 500 //
	X_501                    LanguageImpl = 501 //
	X_502                    LanguageImpl = 502 //
	X_503                    LanguageImpl = 503 //
	X_504                    LanguageImpl = 504 //
	X_505                    LanguageImpl = 505 //
	NDEBELE                  LanguageImpl = 506 // nr
	X_BORK_BORK_BORK         LanguageImpl = 507 // zzb
	X_PIG_LATIN              LanguageImpl = 508 // zzp
	X_HACKER                 LanguageImpl = 509 // zzh
	X_KLINGON                LanguageImpl = 510 // tlh
	X_ELMER_FUDD             LanguageImpl = 511 // zze
	X_Common                 LanguageImpl = 512 // xx-Zyyy
	X_Latin                  LanguageImpl = 513 // xx-Latn
	X_Greek                  LanguageImpl = 514 // xx-Grek
	X_Cyrillic               LanguageImpl = 515 // xx-Cyrl
	X_Armenian               LanguageImpl = 516 // xx-Armn
	X_Hebrew                 LanguageImpl = 517 // xx-Hebr
	X_Arabic                 LanguageImpl = 518 // xx-Arab
	X_Syriac                 LanguageImpl = 519 // xx-Syrc
	X_Thaana                 LanguageImpl = 520 // xx-Thaa
	X_Devanagari             LanguageImpl = 521 // xx-Deva
	X_Bengali                LanguageImpl = 522 // xx-Beng
	X_Gurmukhi               LanguageImpl = 523 // xx-Guru
	X_Gujarati               LanguageImpl = 524 // xx-Gujr
	X_Oriya                  LanguageImpl = 525 // xx-Orya
	X_Tamil                  LanguageImpl = 526 // xx-Taml
	X_Telugu                 LanguageImpl = 527 // xx-Telu
	X_Kannada                LanguageImpl = 528 // xx-Knda
	X_Malayalam              LanguageImpl = 529 // xx-Mlym
	X_Sinhala                LanguageImpl = 530 // xx-Sinh
	X_Thai                   LanguageImpl = 531 // xx-Thai
	X_Lao                    LanguageImpl = 532 // xx-Laoo
	X_Tibetan                LanguageImpl = 533 // xx-Tibt
	X_Myanmar                LanguageImpl = 534 // xx-Mymr
	X_Georgian               LanguageImpl = 535 // xx-Geor
	X_Hangul                 LanguageImpl = 536 // xx-Hang
	X_Ethiopic               LanguageImpl = 537 // xx-Ethi
	X_Cherokee               LanguageImpl = 538 // xx-Cher
	X_Canadian_Aboriginal    LanguageImpl = 539 // xx-Cans
	X_Ogham                  LanguageImpl = 540 // xx-Ogam
	X_Runic                  LanguageImpl = 541 // xx-Runr
	X_Khmer                  LanguageImpl = 542 // xx-Khmr
	X_Mongolian              LanguageImpl = 543 // xx-Mong
	X_Hiragana               LanguageImpl = 544 // xx-Hira
	X_Katakana               LanguageImpl = 545 // xx-Kana
	X_Bopomofo               LanguageImpl = 546 // xx-Bopo
	X_Han                    LanguageImpl = 547 // xx-Hani
	X_Yi                     LanguageImpl = 548 // xx-Yiii
	X_Old_Italic             LanguageImpl = 549 // xx-Ital
	X_Gothic                 LanguageImpl = 550 // xx-Goth
	X_Deseret                LanguageImpl = 551 // xx-Dsrt
	X_Inherited              LanguageImpl = 552 // xx-Qaai
	X_Tagalog                LanguageImpl = 553 // xx-Tglg
	X_Hanunoo                LanguageImpl = 554 // xx-Hano
	X_Buhid                  LanguageImpl = 555 // xx-Buhd
	X_Tagbanwa               LanguageImpl = 556 // xx-Tagb
	X_Limbu                  LanguageImpl = 557 // xx-Limb
	X_Tai_Le                 LanguageImpl = 558 // xx-Tale
	X_Linear_B               LanguageImpl = 559 // xx-Linb
	X_Ugaritic               LanguageImpl = 560 // xx-Ugar
	X_Shavian                LanguageImpl = 561 // xx-Shaw
	X_Osmanya                LanguageImpl = 562 // xx-Osma
	X_Cypriot                LanguageImpl = 563 // xx-Cprt
	X_Braille                LanguageImpl = 564 // xx-Brai
	X_Buginese               LanguageImpl = 565 // xx-Bugi
	X_Coptic                 LanguageImpl = 566 // xx-Copt
	X_New_Tai_Lue            LanguageImpl = 567 // xx-Talu
	X_Glagolitic             LanguageImpl = 568 // xx-Glag
	X_Tifinagh               LanguageImpl = 569 // xx-Tfng
	X_Syloti_Nagri           LanguageImpl = 570 // xx-Sylo
	X_Old_Persian            LanguageImpl = 571 // xx-Xpeo
	X_Kharoshthi             LanguageImpl = 572 // xx-Khar
	X_Balinese               LanguageImpl = 573 // xx-Bali
	X_Cuneiform              LanguageImpl = 574 // xx-Xsux
	X_Phoenician             LanguageImpl = 575 // xx-Phnx
	X_Phags_Pa               LanguageImpl = 576 // xx-Phag
	X_Nko                    LanguageImpl = 577 // xx-Nkoo
	X_Sundanese              LanguageImpl = 578 // xx-Sund
	X_Lepcha                 LanguageImpl = 579 // xx-Lepc
	X_Ol_Chiki               LanguageImpl = 580 // xx-Olck
	X_Vai                    LanguageImpl = 581 // xx-Vaii
	X_Saurashtra             LanguageImpl = 582 // xx-Saur
	X_Kayah_Li               LanguageImpl = 583 // xx-Kali
	X_Rejang                 LanguageImpl = 584 // xx-Rjng
	X_Lycian                 LanguageImpl = 585 // xx-Lyci
	X_Carian                 LanguageImpl = 586 // xx-Cari
	X_Lydian                 LanguageImpl = 587 // xx-Lydi
	X_Cham                   LanguageImpl = 588 // xx-Cham
	X_Tai_Tham               LanguageImpl = 589 // xx-Lana
	X_Tai_Viet               LanguageImpl = 590 // xx-Tavt
	X_Avestan                LanguageImpl = 591 // xx-Avst
	X_Egyptian_Hieroglyphs   LanguageImpl = 592 // xx-Egyp
	X_Samaritan              LanguageImpl = 593 // xx-Samr
	X_Lisu                   LanguageImpl = 594 // xx-Lisu
	X_Bamum                  LanguageImpl = 595 // xx-Bamu
	X_Javanese               LanguageImpl = 596 // xx-Java
	X_Meetei_Mayek           LanguageImpl = 597 // xx-Mtei
	X_Imperial_Aramaic       LanguageImpl = 598 // xx-Armi
	X_Old_South_Arabian      LanguageImpl = 599 // xx-Sarb
	X_Inscriptional_Parthian LanguageImpl = 600 // xx-Prti
	X_Inscriptional_Pahlavi  LanguageImpl = 601 // xx-Phli
	X_Old_Turkic             LanguageImpl = 602 // xx-Orkh
	X_Kaithi                 LanguageImpl = 603 // xx-Kthi
	X_Batak                  LanguageImpl = 604 // xx-Batk
	X_Brahmi                 LanguageImpl = 605 // xx-Brah
	X_Mandaic                LanguageImpl = 606 // xx-Mand
	X_Chakma                 LanguageImpl = 607 // xx-Cakm
	X_Meroitic_Cursive       LanguageImpl = 608 // xx-Merc
	X_Meroitic_Hieroglyphs   LanguageImpl = 609 // xx-Mero
	X_Miao                   LanguageImpl = 610 // xx-Plrd
	X_Sharada                LanguageImpl = 611 // xx-Shrd
	X_Sora_Sompeng           LanguageImpl = 612 // xx-Sora
	X_Takri                  LanguageImpl = 613 // xx-Takr
	NUM_LANGUAGES            LanguageImpl = 614
)

var codeToLanguage = make(map[string]LanguageImpl)

func init() {
	for i, code := range languageToCode {
		if code != "" {
			codeToLanguage[code] = LanguageImpl(i)
		}
	}
}

var languageToCode = []string{
	"en",      // 0 ENGLISH
	"da",      // 1 DANISH
	"nl",      // 2 DUTCH
	"fi",      // 3 FINNISH
	"fr",      // 4 FRENCH
	"de",      // 5 GERMAN
	"iw",      // 6 HEBREW
	"it",      // 7 ITALIAN
	"ja",      // 8 Japanese
	"ko",      // 9 Korean
	"no",      // 10 NORWEGIAN
	"pl",      // 11 POLISH
	"pt",      // 12 PORTUGUESE
	"ru",      // 13 RUSSIAN
	"es",      // 14 SPANISH
	"sv",      // 15 SWEDISH
	"zh",      // 16 Chinese
	"cs",      // 17 CZECH
	"el",      // 18 GREEK
	"is",      // 19 ICELANDIC
	"lv",      // 20 LATVIAN
	"lt",      // 21 LITHUANIAN
	"ro",      // 22 ROMANIAN
	"hu",      // 23 HUNGARIAN
	"et",      // 24 ESTONIAN
	"xxx",     // 25 Ignore
	"un",      // 26 Unknown
	"bg",      // 27 BULGARIAN
	"hr",      // 28 CROATIAN
	"sr",      // 29 SERBIAN
	"ga",      // 30 IRISH
	"gl",      // 31 GALICIAN
	"tl",      // 32 TAGALOG
	"tr",      // 33 TURKISH
	"uk",      // 34 UKRAINIAN
	"hi",      // 35 HINDI
	"mk",      // 36 MACEDONIAN
	"bn",      // 37 BENGALI
	"id",      // 38 INDONESIAN
	"la",      // 39 LATIN
	"ms",      // 40 MALAY
	"ml",      // 41 MALAYALAM
	"cy",      // 42 WELSH
	"ne",      // 43 NEPALI
	"te",      // 44 TELUGU
	"sq",      // 45 ALBANIAN
	"ta",      // 46 TAMIL
	"be",      // 47 BELARUSIAN
	"jw",      // 48 JAVANESE
	"oc",      // 49 OCCITAN
	"ur",      // 50 URDU
	"bh",      // 51 BIHARI
	"gu",      // 52 GUJARATI
	"th",      // 53 THAI
	"ar",      // 54 ARABIC
	"ca",      // 55 CATALAN
	"eo",      // 56 ESPERANTO
	"eu",      // 57 BASQUE
	"ia",      // 58 INTERLINGUA
	"kn",      // 59 KANNADA
	"pa",      // 60 PUNJABI
	"gd",      // 61 SCOTS_GAELIC
	"sw",      // 62 SWAHILI
	"sl",      // 63 SLOVENIAN
	"mr",      // 64 MARATHI
	"mt",      // 65 MALTESE
	"vi",      // 66 VIETNAMESE
	"fy",      // 67 FRISIAN
	"sk",      // 68 SLOVAK
	"zh-Hant", // 69 ChineseT
	"fo",      // 70 FAROESE
	"su",      // 71 SUNDANESE
	"uz",      // 72 UZBEK
	"am",      // 73 AMHARIC
	"az",      // 74 AZERBAIJANI
	"ka",      // 75 GEORGIAN
	"ti",      // 76 TIGRINYA
	"fa",      // 77 PERSIAN
	"bs",      // 78 BOSNIAN
	"si",      // 79 SINHALESE
	"nn",      // 80 NORWEGIAN_N
	"",        // 81 81
	"",        // 82 82
	"xh",      // 83 XHOSA
	"zu",      // 84 ZULU
	"gn",      // 85 GUARANI
	"st",      // 86 SESOTHO
	"tk",      // 87 TURKMEN
	"ky",      // 88 KYRGYZ
	"br",      // 89 BRETON
	"tw",      // 90 TWI
	"yi",      // 91 YIDDISH
	"",        // 92 92
	"so",      // 93 SOMALI
	"ug",      // 94 UIGHUR
	"ku",      // 95 KURDISH
	"mn",      // 96 MONGOLIAN
	"hy",      // 97 ARMENIAN
	"lo",      // 98 LAOTHIAN
	"sd",      // 99 SINDHI
	"rm",      // 100 RHAETO_ROMANCE
	"af",      // 101 AFRIKAANS
	"lb",      // 102 LUXEMBOURGISH
	"my",      // 103 BURMESE
	"km",      // 104 KHMER
	"bo",      // 105 TIBETAN
	"dv",      // 106 DHIVEHI
	"chr",     // 107 CHEROKEE
	"syr",     // 108 SYRIAC
	"lif",     // 109 LIMBU
	"or",      // 110 ORIYA
	"as",      // 111 ASSAMESE
	"co",      // 112 CORSICAN
	"ie",      // 113 INTERLINGUE
	"kk",      // 114 KAZAKH
	"ln",      // 115 LINGALA
	"",        // 116 116
	"ps",      // 117 PASHTO
	"qu",      // 118 QUECHUA
	"sn",      // 119 SHONA
	"tg",      // 120 TAJIK
	"tt",      // 121 TATAR
	"to",      // 122 TONGA
	"yo",      // 123 YORUBA
	"",        // 124 124
	"",        // 125 125
	"",        // 126 126
	"",        // 127 127
	"mi",      // 128 MAORI
	"wo",      // 129 WOLOF
	"ab",      // 130 ABKHAZIAN
	"aa",      // 131 AFAR
	"ay",      // 132 AYMARA
	"ba",      // 133 BASHKIR
	"bi",      // 134 BISLAMA
	"dz",      // 135 DZONGKHA
	"fj",      // 136 FIJIAN
	"kl",      // 137 GREENLANDIC
	"ha",      // 138 HAUSA
	"ht",      // 139 HAITIAN_CREOLE
	"ik",      // 140 INUPIAK
	"iu",      // 141 INUKTITUT
	"ks",      // 142 KASHMIRI
	"rw",      // 143 KINYARWANDA
	"mg",      // 144 MALAGASY
	"na",      // 145 NAURU
	"om",      // 146 OROMO
	"rn",      // 147 RUNDI
	"sm",      // 148 SAMOAN
	"sg",      // 149 SANGO
	"sa",      // 150 SANSKRIT
	"ss",      // 151 SISWANT
	"ts",      // 152 TSONGA
	"tn",      // 153 TSWANA
	"vo",      // 154 VOLAPUK
	"za",      // 155 ZHUANG
	"kha",     // 156 KHASI
	"sco",     // 157 SCOTS
	"lg",      // 158 GANDA
	"gv",      // 159 MANX
	"sr-ME",   // 160 MONTENEGRIN
	"ak",      // 161 AKAN
	"ig",      // 162 IGBO
	"mfe",     // 163 MAURITIAN_CREOLE
	"haw",     // 164 HAWAIIAN
	"ceb",     // 165 CEBUANO
	"ee",      // 166 EWE
	"gaa",     // 167 GA
	"hmn",     // 168 HMONG
	"kri",     // 169 KRIO
	"loz",     // 170 LOZI
	"lua",     // 171 LUBA_LULUA
	"luo",     // 172 LUO_KENYA_AND_TANZANIA
	"new",     // 173 NEWARI
	"ny",      // 174 NYANJA
	"os",      // 175 OSSETIAN
	"pam",     // 176 PAMPANGA
	"nso",     // 177 PEDI
	"raj",     // 178 RAJASTHANI
	"crs",     // 179 SESELWA
	"tum",     // 180 TUMBUKA
	"ve",      // 181 VENDA
	"war",     // 182 WARAY_PHILIPPINES
	"",        // 183 183
	"",        // 184 184
	"",        // 185 185
	"",        // 186 186
	"",        // 187 187
	"",        // 188 188
	"",        // 189 189
	"",        // 190 190
	"",        // 191 191
	"",        // 192 192
	"",        // 193 193
	"",        // 194 194
	"",        // 195 195
	"",        // 196 196
	"",        // 197 197
	"",        // 198 198
	"",        // 199 199
	"",        // 200 200
	"",        // 201 201
	"",        // 202 202
	"",        // 203 203
	"",        // 204 204
	"",        // 205 205
	"",        // 206 206
	"",        // 207 207
	"",        // 208 208
	"",        // 209 209
	"",        // 210 210
	"",        // 211 211
	"",        // 212 212
	"",        // 213 213
	"",        // 214 214
	"",        // 215 215
	"",        // 216 216
	"",        // 217 217
	"",        // 218 218
	"",        // 219 219
	"",        // 220 220
	"",        // 221 221
	"",        // 222 222
	"",        // 223 223
	"",        // 224 224
	"",        // 225 225
	"",        // 226 226
	"",        // 227 227
	"",        // 228 228
	"",        // 229 229
	"",        // 230 230
	"",        // 231 231
	"",        // 232 232
	"",        // 233 233
	"",        // 234 234
	"",        // 235 235
	"",        // 236 236
	"",        // 237 237
	"",        // 238 238
	"",        // 239 239
	"",        // 240 240
	"",        // 241 241
	"",        // 242 242
	"",        // 243 243
	"",        // 244 244
	"",        // 245 245
	"",        // 246 246
	"",        // 247 247
	"",        // 248 248
	"",        // 249 249
	"",        // 250 250
	"",        // 251 251
	"",        // 252 252
	"",        // 253 253
	"",        // 254 254
	"",        // 255 255
	"",        // 256 256
	"",        // 257 257
	"",        // 258 258
	"",        // 259 259
	"",        // 260 260
	"",        // 261 261
	"",        // 262 262
	"",        // 263 263
	"",        // 264 264
	"",        // 265 265
	"",        // 266 266
	"",        // 267 267
	"",        // 268 268
	"",        // 269 269
	"",        // 270 270
	"",        // 271 271
	"",        // 272 272
	"",        // 273 273
	"",        // 274 274
	"",        // 275 275
	"",        // 276 276
	"",        // 277 277
	"",        // 278 278
	"",        // 279 279
	"",        // 280 280
	"",        // 281 281
	"",        // 282 282
	"",        // 283 283
	"",        // 284 284
	"",        // 285 285
	"",        // 286 286
	"",        // 287 287
	"",        // 288 288
	"",        // 289 289
	"",        // 290 290
	"",        // 291 291
	"",        // 292 292
	"",        // 293 293
	"",        // 294 294
	"",        // 295 295
	"",        // 296 296
	"",        // 297 297
	"",        // 298 298
	"",        // 299 299
	"",        // 300 300
	"",        // 301 301
	"",        // 302 302
	"",        // 303 303
	"",        // 304 304
	"",        // 305 305
	"",        // 306 306
	"",        // 307 307
	"",        // 308 308
	"",        // 309 309
	"",        // 310 310
	"",        // 311 311
	"",        // 312 312
	"",        // 313 313
	"",        // 314 314
	"",        // 315 315
	"",        // 316 316
	"",        // 317 317
	"",        // 318 318
	"",        // 319 319
	"",        // 320 320
	"",        // 321 321
	"",        // 322 322
	"",        // 323 323
	"",        // 324 324
	"",        // 325 325
	"",        // 326 326
	"",        // 327 327
	"",        // 328 328
	"",        // 329 329
	"",        // 330 330
	"",        // 331 331
	"",        // 332 332
	"",        // 333 333
	"",        // 334 334
	"",        // 335 335
	"",        // 336 336
	"",        // 337 337
	"",        // 338 338
	"",        // 339 339
	"",        // 340 340
	"",        // 341 341
	"",        // 342 342
	"",        // 343 343
	"",        // 344 344
	"",        // 345 345
	"",        // 346 346
	"",        // 347 347
	"",        // 348 348
	"",        // 349 349
	"",        // 350 350
	"",        // 351 351
	"",        // 352 352
	"",        // 353 353
	"",        // 354 354
	"",        // 355 355
	"",        // 356 356
	"",        // 357 357
	"",        // 358 358
	"",        // 359 359
	"",        // 360 360
	"",        // 361 361
	"",        // 362 362
	"",        // 363 363
	"",        // 364 364
	"",        // 365 365
	"",        // 366 366
	"",        // 367 367
	"",        // 368 368
	"",        // 369 369
	"",        // 370 370
	"",        // 371 371
	"",        // 372 372
	"",        // 373 373
	"",        // 374 374
	"",        // 375 375
	"",        // 376 376
	"",        // 377 377
	"",        // 378 378
	"",        // 379 379
	"",        // 380 380
	"",        // 381 381
	"",        // 382 382
	"",        // 383 383
	"",        // 384 384
	"",        // 385 385
	"",        // 386 386
	"",        // 387 387
	"",        // 388 388
	"",        // 389 389
	"",        // 390 390
	"",        // 391 391
	"",        // 392 392
	"",        // 393 393
	"",        // 394 394
	"",        // 395 395
	"",        // 396 396
	"",        // 397 397
	"",        // 398 398
	"",        // 399 399
	"",        // 400 400
	"",        // 401 401
	"",        // 402 402
	"",        // 403 403
	"",        // 404 404
	"",        // 405 405
	"",        // 406 406
	"",        // 407 407
	"",        // 408 408
	"",        // 409 409
	"",        // 410 410
	"",        // 411 411
	"",        // 412 412
	"",        // 413 413
	"",        // 414 414
	"",        // 415 415
	"",        // 416 416
	"",        // 417 417
	"",        // 418 418
	"",        // 419 419
	"",        // 420 420
	"",        // 421 421
	"",        // 422 422
	"",        // 423 423
	"",        // 424 424
	"",        // 425 425
	"",        // 426 426
	"",        // 427 427
	"",        // 428 428
	"",        // 429 429
	"",        // 430 430
	"",        // 431 431
	"",        // 432 432
	"",        // 433 433
	"",        // 434 434
	"",        // 435 435
	"",        // 436 436
	"",        // 437 437
	"",        // 438 438
	"",        // 439 439
	"",        // 440 440
	"",        // 441 441
	"",        // 442 442
	"",        // 443 443
	"",        // 444 444
	"",        // 445 445
	"",        // 446 446
	"",        // 447 447
	"",        // 448 448
	"",        // 449 449
	"",        // 450 450
	"",        // 451 451
	"",        // 452 452
	"",        // 453 453
	"",        // 454 454
	"",        // 455 455
	"",        // 456 456
	"",        // 457 457
	"",        // 458 458
	"",        // 459 459
	"",        // 460 460
	"",        // 461 461
	"",        // 462 462
	"",        // 463 463
	"",        // 464 464
	"",        // 465 465
	"",        // 466 466
	"",        // 467 467
	"",        // 468 468
	"",        // 469 469
	"",        // 470 470
	"",        // 471 471
	"",        // 472 472
	"",        // 473 473
	"",        // 474 474
	"",        // 475 475
	"",        // 476 476
	"",        // 477 477
	"",        // 478 478
	"",        // 479 479
	"",        // 480 480
	"",        // 481 481
	"",        // 482 482
	"",        // 483 483
	"",        // 484 484
	"",        // 485 485
	"",        // 486 486
	"",        // 487 487
	"",        // 488 488
	"",        // 489 489
	"",        // 490 490
	"",        // 491 491
	"",        // 492 492
	"",        // 493 493
	"",        // 494 494
	"",        // 495 495
	"",        // 496 496
	"",        // 497 497
	"",        // 498 498
	"",        // 499 499
	"",        // 500 500
	"",        // 501 501
	"",        // 502 502
	"",        // 503 503
	"",        // 504 504
	"",        // 505 505
	"nr",      // 506 NDEBELE
	"zzb",     // 507 X_BORK_BORK_BORK
	"zzp",     // 508 X_PIG_LATIN
	"zzh",     // 509 X_HACKER
	"tlh",     // 510 X_KLINGON
	"zze",     // 511 X_ELMER_FUDD
	"xx-Zyyy", // 512 X_Common
	"xx-Latn", // 513 X_Latin
	"xx-Grek", // 514 X_Greek
	"xx-Cyrl", // 515 X_Cyrillic
	"xx-Armn", // 516 X_Armenian
	"xx-Hebr", // 517 X_Hebrew
	"xx-Arab", // 518 X_Arabic
	"xx-Syrc", // 519 X_Syriac
	"xx-Thaa", // 520 X_Thaana
	"xx-Deva", // 521 X_Devanagari
	"xx-Beng", // 522 X_Bengali
	"xx-Guru", // 523 X_Gurmukhi
	"xx-Gujr", // 524 X_Gujarati
	"xx-Orya", // 525 X_Oriya
	"xx-Taml", // 526 X_Tamil
	"xx-Telu", // 527 X_Telugu
	"xx-Knda", // 528 X_Kannada
	"xx-Mlym", // 529 X_Malayalam
	"xx-Sinh", // 530 X_Sinhala
	"xx-Thai", // 531 X_Thai
	"xx-Laoo", // 532 X_Lao
	"xx-Tibt", // 533 X_Tibetan
	"xx-Mymr", // 534 X_Myanmar
	"xx-Geor", // 535 X_Georgian
	"xx-Hang", // 536 X_Hangul
	"xx-Ethi", // 537 X_Ethiopic
	"xx-Cher", // 538 X_Cherokee
	"xx-Cans", // 539 X_Canadian_Aboriginal
	"xx-Ogam", // 540 X_Ogham
	"xx-Runr", // 541 X_Runic
	"xx-Khmr", // 542 X_Khmer
	"xx-Mong", // 543 X_Mongolian
	"xx-Hira", // 544 X_Hiragana
	"xx-Kana", // 545 X_Katakana
	"xx-Bopo", // 546 X_Bopomofo
	"xx-Hani", // 547 X_Han
	"xx-Yiii", // 548 X_Yi
	"xx-Ital", // 549 X_Old_Italic
	"xx-Goth", // 550 X_Gothic
	"xx-Dsrt", // 551 X_Deseret
	"xx-Qaai", // 552 X_Inherited
	"xx-Tglg", // 553 X_Tagalog
	"xx-Hano", // 554 X_Hanunoo
	"xx-Buhd", // 555 X_Buhid
	"xx-Tagb", // 556 X_Tagbanwa
	"xx-Limb", // 557 X_Limbu
	"xx-Tale", // 558 X_Tai_Le
	"xx-Linb", // 559 X_Linear_B
	"xx-Ugar", // 560 X_Ugaritic
	"xx-Shaw", // 561 X_Shavian
	"xx-Osma", // 562 X_Osmanya
	"xx-Cprt", // 563 X_Cypriot
	"xx-Brai", // 564 X_Braille
	"xx-Bugi", // 565 X_Buginese
	"xx-Copt", // 566 X_Coptic
	"xx-Talu", // 567 X_New_Tai_Lue
	"xx-Glag", // 568 X_Glagolitic
	"xx-Tfng", // 569 X_Tifinagh
	"xx-Sylo", // 570 X_Syloti_Nagri
	"xx-Xpeo", // 571 X_Old_Persian
	"xx-Khar", // 572 X_Kharoshthi
	"xx-Bali", // 573 X_Balinese
	"xx-Xsux", // 574 X_Cuneiform
	"xx-Phnx", // 575 X_Phoenician
	"xx-Phag", // 576 X_Phags_Pa
	"xx-Nkoo", // 577 X_Nko
	"xx-Sund", // 578 X_Sundanese
	"xx-Lepc", // 579 X_Lepcha
	"xx-Olck", // 580 X_Ol_Chiki
	"xx-Vaii", // 581 X_Vai
	"xx-Saur", // 582 X_Saurashtra
	"xx-Kali", // 583 X_Kayah_Li
	"xx-Rjng", // 584 X_Rejang
	"xx-Lyci", // 585 X_Lycian
	"xx-Cari", // 586 X_Carian
	"xx-Lydi", // 587 X_Lydian
	"xx-Cham", // 588 X_Cham
	"xx-Lana", // 589 X_Tai_Tham
	"xx-Tavt", // 590 X_Tai_Viet
	"xx-Avst", // 591 X_Avestan
	"xx-Egyp", // 592 X_Egyptian_Hieroglyphs
	"xx-Samr", // 593 X_Samaritan
	"xx-Lisu", // 594 X_Lisu
	"xx-Bamu", // 595 X_Bamum
	"xx-Java", // 596 X_Javanese
	"xx-Mtei", // 597 X_Meetei_Mayek
	"xx-Armi", // 598 X_Imperial_Aramaic
	"xx-Sarb", // 599 X_Old_South_Arabian
	"xx-Prti", // 600 X_Inscriptional_Parthian
	"xx-Phli", // 601 X_Inscriptional_Pahlavi
	"xx-Orkh", // 602 X_Old_Turkic
	"xx-Kthi", // 603 X_Kaithi
	"xx-Batk", // 604 X_Batak
	"xx-Brah", // 605 X_Brahmi
	"xx-Mand", // 606 X_Mandaic
	"xx-Cakm", // 607 X_Chakma
	"xx-Merc", // 608 X_Meroitic_Cursive
	"xx-Mero", // 609 X_Meroitic_Hieroglyphs
	"xx-Plrd", // 610 X_Miao
	"xx-Shrd", // 611 X_Sharada
	"xx-Sora", // 612 X_Sora_Sompeng
	"xx-Takr", // 613 X_Takri
}

var languageToCName = []string{
	"ENGLISH",                  // 0 en
	"DANISH",                   // 1 da
	"DUTCH",                    // 2 nl
	"FINNISH",                  // 3 fi
	"FRENCH",                   // 4 fr
	"GERMAN",                   // 5 de
	"HEBREW",                   // 6 iw
	"ITALIAN",                  // 7 it
	"JAPANESE",                 // 8 ja
	"KOREAN",                   // 9 ko
	"NORWEGIAN",                // 10 no
	"POLISH",                   // 11 pl
	"PORTUGUESE",               // 12 pt
	"RUSSIAN",                  // 13 ru
	"SPANISH",                  // 14 es
	"SWEDISH",                  // 15 sv
	"CHINESE",                  // 16 zh
	"CZECH",                    // 17 cs
	"GREEK",                    // 18 el
	"ICELANDIC",                // 19 is
	"LATVIAN",                  // 20 lv
	"LITHUANIAN",               // 21 lt
	"ROMANIAN",                 // 22 ro
	"HUNGARIAN",                // 23 hu
	"ESTONIAN",                 // 24 et
	"TG_UNKNOWN_LANGUAGE",      // 25 xxx
	"UNKNOWN_LANGUAGE",         // 26 un
	"BULGARIAN",                // 27 bg
	"CROATIAN",                 // 28 hr
	"SERBIAN",                  // 29 sr
	"IRISH",                    // 30 ga
	"GALICIAN",                 // 31 gl
	"TAGALOG",                  // 32 tl
	"TURKISH",                  // 33 tr
	"UKRAINIAN",                // 34 uk
	"HINDI",                    // 35 hi
	"MACEDONIAN",               // 36 mk
	"BENGALI",                  // 37 bn
	"INDONESIAN",               // 38 id
	"LATIN",                    // 39 la
	"MALAY",                    // 40 ms
	"MALAYALAM",                // 41 ml
	"WELSH",                    // 42 cy
	"NEPALI",                   // 43 ne
	"TELUGU",                   // 44 te
	"ALBANIAN",                 // 45 sq
	"TAMIL",                    // 46 ta
	"BELARUSIAN",               // 47 be
	"JAVANESE",                 // 48 jw
	"OCCITAN",                  // 49 oc
	"URDU",                     // 50 ur
	"BIHARI",                   // 51 bh
	"GUJARATI",                 // 52 gu
	"THAI",                     // 53 th
	"ARABIC",                   // 54 ar
	"CATALAN",                  // 55 ca
	"ESPERANTO",                // 56 eo
	"BASQUE",                   // 57 eu
	"INTERLINGUA",              // 58 ia
	"KANNADA",                  // 59 kn
	"PUNJABI",                  // 60 pa
	"SCOTS_GAELIC",             // 61 gd
	"SWAHILI",                  // 62 sw
	"SLOVENIAN",                // 63 sl
	"MARATHI",                  // 64 mr
	"MALTESE",                  // 65 mt
	"VIETNAMESE",               // 66 vi
	"FRISIAN",                  // 67 fy
	"SLOVAK",                   // 68 sk
	"CHINESE_T",                // 69 zh-Hant
	"FAROESE",                  // 70 fo
	"SUNDANESE",                // 71 su
	"UZBEK",                    // 72 uz
	"AMHARIC",                  // 73 am
	"AZERBAIJANI",              // 74 az
	"GEORGIAN",                 // 75 ka
	"TIGRINYA",                 // 76 ti
	"PERSIAN",                  // 77 fa
	"BOSNIAN",                  // 78 bs
	"SINHALESE",                // 79 si
	"NORWEGIAN_N",              // 80 nn
	"X_81",                     // 81
	"X_82",                     // 82
	"XHOSA",                    // 83 xh
	"ZULU",                     // 84 zu
	"GUARANI",                  // 85 gn
	"SESOTHO",                  // 86 st
	"TURKMEN",                  // 87 tk
	"KYRGYZ",                   // 88 ky
	"BRETON",                   // 89 br
	"TWI",                      // 90 tw
	"YIDDISH",                  // 91 yi
	"X_92",                     // 92
	"SOMALI",                   // 93 so
	"UIGHUR",                   // 94 ug
	"KURDISH",                  // 95 ku
	"MONGOLIAN",                // 96 mn
	"ARMENIAN",                 // 97 hy
	"LAOTHIAN",                 // 98 lo
	"SINDHI",                   // 99 sd
	"RHAETO_ROMANCE",           // 100 rm
	"AFRIKAANS",                // 101 af
	"LUXEMBOURGISH",            // 102 lb
	"BURMESE",                  // 103 my
	"KHMER",                    // 104 km
	"TIBETAN",                  // 105 bo
	"DHIVEHI",                  // 106 dv
	"CHEROKEE",                 // 107 chr
	"SYRIAC",                   // 108 syr
	"LIMBU",                    // 109 lif
	"ORIYA",                    // 110 or
	"ASSAMESE",                 // 111 as
	"CORSICAN",                 // 112 co
	"INTERLINGUE",              // 113 ie
	"KAZAKH",                   // 114 kk
	"LINGALA",                  // 115 ln
	"X_116",                    // 116
	"PASHTO",                   // 117 ps
	"QUECHUA",                  // 118 qu
	"SHONA",                    // 119 sn
	"TAJIK",                    // 120 tg
	"TATAR",                    // 121 tt
	"TONGA",                    // 122 to
	"YORUBA",                   // 123 yo
	"X_124",                    // 124
	"X_125",                    // 125
	"X_126",                    // 126
	"X_127",                    // 127
	"MAORI",                    // 128 mi
	"WOLOF",                    // 129 wo
	"ABKHAZIAN",                // 130 ab
	"AFAR",                     // 131 aa
	"AYMARA",                   // 132 ay
	"BASHKIR",                  // 133 ba
	"BISLAMA",                  // 134 bi
	"DZONGKHA",                 // 135 dz
	"FIJIAN",                   // 136 fj
	"GREENLANDIC",              // 137 kl
	"HAUSA",                    // 138 ha
	"HAITIAN_CREOLE",           // 139 ht
	"INUPIAK",                  // 140 ik
	"INUKTITUT",                // 141 iu
	"KASHMIRI",                 // 142 ks
	"KINYARWANDA",              // 143 rw
	"MALAGASY",                 // 144 mg
	"NAURU",                    // 145 na
	"OROMO",                    // 146 om
	"RUNDI",                    // 147 rn
	"SAMOAN",                   // 148 sm
	"SANGO",                    // 149 sg
	"SANSKRIT",                 // 150 sa
	"SISWANT",                  // 151 ss
	"TSONGA",                   // 152 ts
	"TSWANA",                   // 153 tn
	"VOLAPUK",                  // 154 vo
	"ZHUANG",                   // 155 za
	"KHASI",                    // 156 kha
	"SCOTS",                    // 157 sco
	"GANDA",                    // 158 lg
	"MANX",                     // 159 gv
	"MONTENEGRIN",              // 160 sr-ME
	"AKAN",                     // 161 ak
	"IGBO",                     // 162 ig
	"MAURITIAN_CREOLE",         // 163 mfe
	"HAWAIIAN",                 // 164 haw
	"CEBUANO",                  // 165 ceb
	"EWE",                      // 166 ee
	"GA",                       // 167 gaa
	"HMONG",                    // 168 hmn
	"KRIO",                     // 169 kri
	"LOZI",                     // 170 loz
	"LUBA_LULUA",               // 171 lua
	"LUO_KENYA_AND_TANZANIA",   // 172 luo
	"NEWARI",                   // 173 new
	"NYANJA",                   // 174 ny
	"OSSETIAN",                 // 175 os
	"PAMPANGA",                 // 176 pam
	"PEDI",                     // 177 nso
	"RAJASTHANI",               // 178 raj
	"SESELWA",                  // 179 crs
	"TUMBUKA",                  // 180 tum
	"VENDA",                    // 181 ve
	"WARAY_PHILIPPINES",        // 182 war
	"X_183",                    // 183
	"X_184",                    // 184
	"X_185",                    // 185
	"X_186",                    // 186
	"X_187",                    // 187
	"X_188",                    // 188
	"X_189",                    // 189
	"X_190",                    // 190
	"X_191",                    // 191
	"X_192",                    // 192
	"X_193",                    // 193
	"X_194",                    // 194
	"X_195",                    // 195
	"X_196",                    // 196
	"X_197",                    // 197
	"X_198",                    // 198
	"X_199",                    // 199
	"X_200",                    // 200
	"X_201",                    // 201
	"X_202",                    // 202
	"X_203",                    // 203
	"X_204",                    // 204
	"X_205",                    // 205
	"X_206",                    // 206
	"X_207",                    // 207
	"X_208",                    // 208
	"X_209",                    // 209
	"X_210",                    // 210
	"X_211",                    // 211
	"X_212",                    // 212
	"X_213",                    // 213
	"X_214",                    // 214
	"X_215",                    // 215
	"X_216",                    // 216
	"X_217",                    // 217
	"X_218",                    // 218
	"X_219",                    // 219
	"X_220",                    // 220
	"X_221",                    // 221
	"X_222",                    // 222
	"X_223",                    // 223
	"X_224",                    // 224
	"X_225",                    // 225
	"X_226",                    // 226
	"X_227",                    // 227
	"X_228",                    // 228
	"X_229",                    // 229
	"X_230",                    // 230
	"X_231",                    // 231
	"X_232",                    // 232
	"X_233",                    // 233
	"X_234",                    // 234
	"X_235",                    // 235
	"X_236",                    // 236
	"X_237",                    // 237
	"X_238",                    // 238
	"X_239",                    // 239
	"X_240",                    // 240
	"X_241",                    // 241
	"X_242",                    // 242
	"X_243",                    // 243
	"X_244",                    // 244
	"X_245",                    // 245
	"X_246",                    // 246
	"X_247",                    // 247
	"X_248",                    // 248
	"X_249",                    // 249
	"X_250",                    // 250
	"X_251",                    // 251
	"X_252",                    // 252
	"X_253",                    // 253
	"X_254",                    // 254
	"X_255",                    // 255
	"X_256",                    // 256
	"X_257",                    // 257
	"X_258",                    // 258
	"X_259",                    // 259
	"X_260",                    // 260
	"X_261",                    // 261
	"X_262",                    // 262
	"X_263",                    // 263
	"X_264",                    // 264
	"X_265",                    // 265
	"X_266",                    // 266
	"X_267",                    // 267
	"X_268",                    // 268
	"X_269",                    // 269
	"X_270",                    // 270
	"X_271",                    // 271
	"X_272",                    // 272
	"X_273",                    // 273
	"X_274",                    // 274
	"X_275",                    // 275
	"X_276",                    // 276
	"X_277",                    // 277
	"X_278",                    // 278
	"X_279",                    // 279
	"X_280",                    // 280
	"X_281",                    // 281
	"X_282",                    // 282
	"X_283",                    // 283
	"X_284",                    // 284
	"X_285",                    // 285
	"X_286",                    // 286
	"X_287",                    // 287
	"X_288",                    // 288
	"X_289",                    // 289
	"X_290",                    // 290
	"X_291",                    // 291
	"X_292",                    // 292
	"X_293",                    // 293
	"X_294",                    // 294
	"X_295",                    // 295
	"X_296",                    // 296
	"X_297",                    // 297
	"X_298",                    // 298
	"X_299",                    // 299
	"X_300",                    // 300
	"X_301",                    // 301
	"X_302",                    // 302
	"X_303",                    // 303
	"X_304",                    // 304
	"X_305",                    // 305
	"X_306",                    // 306
	"X_307",                    // 307
	"X_308",                    // 308
	"X_309",                    // 309
	"X_310",                    // 310
	"X_311",                    // 311
	"X_312",                    // 312
	"X_313",                    // 313
	"X_314",                    // 314
	"X_315",                    // 315
	"X_316",                    // 316
	"X_317",                    // 317
	"X_318",                    // 318
	"X_319",                    // 319
	"X_320",                    // 320
	"X_321",                    // 321
	"X_322",                    // 322
	"X_323",                    // 323
	"X_324",                    // 324
	"X_325",                    // 325
	"X_326",                    // 326
	"X_327",                    // 327
	"X_328",                    // 328
	"X_329",                    // 329
	"X_330",                    // 330
	"X_331",                    // 331
	"X_332",                    // 332
	"X_333",                    // 333
	"X_334",                    // 334
	"X_335",                    // 335
	"X_336",                    // 336
	"X_337",                    // 337
	"X_338",                    // 338
	"X_339",                    // 339
	"X_340",                    // 340
	"X_341",                    // 341
	"X_342",                    // 342
	"X_343",                    // 343
	"X_344",                    // 344
	"X_345",                    // 345
	"X_346",                    // 346
	"X_347",                    // 347
	"X_348",                    // 348
	"X_349",                    // 349
	"X_350",                    // 350
	"X_351",                    // 351
	"X_352",                    // 352
	"X_353",                    // 353
	"X_354",                    // 354
	"X_355",                    // 355
	"X_356",                    // 356
	"X_357",                    // 357
	"X_358",                    // 358
	"X_359",                    // 359
	"X_360",                    // 360
	"X_361",                    // 361
	"X_362",                    // 362
	"X_363",                    // 363
	"X_364",                    // 364
	"X_365",                    // 365
	"X_366",                    // 366
	"X_367",                    // 367
	"X_368",                    // 368
	"X_369",                    // 369
	"X_370",                    // 370
	"X_371",                    // 371
	"X_372",                    // 372
	"X_373",                    // 373
	"X_374",                    // 374
	"X_375",                    // 375
	"X_376",                    // 376
	"X_377",                    // 377
	"X_378",                    // 378
	"X_379",                    // 379
	"X_380",                    // 380
	"X_381",                    // 381
	"X_382",                    // 382
	"X_383",                    // 383
	"X_384",                    // 384
	"X_385",                    // 385
	"X_386",                    // 386
	"X_387",                    // 387
	"X_388",                    // 388
	"X_389",                    // 389
	"X_390",                    // 390
	"X_391",                    // 391
	"X_392",                    // 392
	"X_393",                    // 393
	"X_394",                    // 394
	"X_395",                    // 395
	"X_396",                    // 396
	"X_397",                    // 397
	"X_398",                    // 398
	"X_399",                    // 399
	"X_400",                    // 400
	"X_401",                    // 401
	"X_402",                    // 402
	"X_403",                    // 403
	"X_404",                    // 404
	"X_405",                    // 405
	"X_406",                    // 406
	"X_407",                    // 407
	"X_408",                    // 408
	"X_409",                    // 409
	"X_410",                    // 410
	"X_411",                    // 411
	"X_412",                    // 412
	"X_413",                    // 413
	"X_414",                    // 414
	"X_415",                    // 415
	"X_416",                    // 416
	"X_417",                    // 417
	"X_418",                    // 418
	"X_419",                    // 419
	"X_420",                    // 420
	"X_421",                    // 421
	"X_422",                    // 422
	"X_423",                    // 423
	"X_424",                    // 424
	"X_425",                    // 425
	"X_426",                    // 426
	"X_427",                    // 427
	"X_428",                    // 428
	"X_429",                    // 429
	"X_430",                    // 430
	"X_431",                    // 431
	"X_432",                    // 432
	"X_433",                    // 433
	"X_434",                    // 434
	"X_435",                    // 435
	"X_436",                    // 436
	"X_437",                    // 437
	"X_438",                    // 438
	"X_439",                    // 439
	"X_440",                    // 440
	"X_441",                    // 441
	"X_442",                    // 442
	"X_443",                    // 443
	"X_444",                    // 444
	"X_445",                    // 445
	"X_446",                    // 446
	"X_447",                    // 447
	"X_448",                    // 448
	"X_449",                    // 449
	"X_450",                    // 450
	"X_451",                    // 451
	"X_452",                    // 452
	"X_453",                    // 453
	"X_454",                    // 454
	"X_455",                    // 455
	"X_456",                    // 456
	"X_457",                    // 457
	"X_458",                    // 458
	"X_459",                    // 459
	"X_460",                    // 460
	"X_461",                    // 461
	"X_462",                    // 462
	"X_463",                    // 463
	"X_464",                    // 464
	"X_465",                    // 465
	"X_466",                    // 466
	"X_467",                    // 467
	"X_468",                    // 468
	"X_469",                    // 469
	"X_470",                    // 470
	"X_471",                    // 471
	"X_472",                    // 472
	"X_473",                    // 473
	"X_474",                    // 474
	"X_475",                    // 475
	"X_476",                    // 476
	"X_477",                    // 477
	"X_478",                    // 478
	"X_479",                    // 479
	"X_480",                    // 480
	"X_481",                    // 481
	"X_482",                    // 482
	"X_483",                    // 483
	"X_484",                    // 484
	"X_485",                    // 485
	"X_486",                    // 486
	"X_487",                    // 487
	"X_488",                    // 488
	"X_489",                    // 489
	"X_490",                    // 490
	"X_491",                    // 491
	"X_492",                    // 492
	"X_493",                    // 493
	"X_494",                    // 494
	"X_495",                    // 495
	"X_496",                    // 496
	"X_497",                    // 497
	"X_498",                    // 498
	"X_499",                    // 499
	"X_500",                    // 500
	"X_501",                    // 501
	"X_502",                    // 502
	"X_503",                    // 503
	"X_504",                    // 504
	"X_505",                    // 505
	"NDEBELE",                  // 506 nr
	"X_BORK_BORK_BORK",         // 507 zzb
	"X_PIG_LATIN",              // 508 zzp
	"X_HACKER",                 // 509 zzh
	"X_KLINGON",                // 510 tlh
	"X_ELMER_FUDD",             // 511 zze
	"X_Common",                 // 512 xx-Zyyy
	"X_Latin",                  // 513 xx-Latn
	"X_Greek",                  // 514 xx-Grek
	"X_Cyrillic",               // 515 xx-Cyrl
	"X_Armenian",               // 516 xx-Armn
	"X_Hebrew",                 // 517 xx-Hebr
	"X_Arabic",                 // 518 xx-Arab
	"X_Syriac",                 // 519 xx-Syrc
	"X_Thaana",                 // 520 xx-Thaa
	"X_Devanagari",             // 521 xx-Deva
	"X_Bengali",                // 522 xx-Beng
	"X_Gurmukhi",               // 523 xx-Guru
	"X_Gujarati",               // 524 xx-Gujr
	"X_Oriya",                  // 525 xx-Orya
	"X_Tamil",                  // 526 xx-Taml
	"X_Telugu",                 // 527 xx-Telu
	"X_Kannada",                // 528 xx-Knda
	"X_Malayalam",              // 529 xx-Mlym
	"X_Sinhala",                // 530 xx-Sinh
	"X_Thai",                   // 531 xx-Thai
	"X_Lao",                    // 532 xx-Laoo
	"X_Tibetan",                // 533 xx-Tibt
	"X_Myanmar",                // 534 xx-Mymr
	"X_Georgian",               // 535 xx-Geor
	"X_Hangul",                 // 536 xx-Hang
	"X_Ethiopic",               // 537 xx-Ethi
	"X_Cherokee",               // 538 xx-Cher
	"X_Canadian_Aboriginal",    // 539 xx-Cans
	"X_Ogham",                  // 540 xx-Ogam
	"X_Runic",                  // 541 xx-Runr
	"X_Khmer",                  // 542 xx-Khmr
	"X_Mongolian",              // 543 xx-Mong
	"X_Hiragana",               // 544 xx-Hira
	"X_Katakana",               // 545 xx-Kana
	"X_Bopomofo",               // 546 xx-Bopo
	"X_Han",                    // 547 xx-Hani
	"X_Yi",                     // 548 xx-Yiii
	"X_Old_Italic",             // 549 xx-Ital
	"X_Gothic",                 // 550 xx-Goth
	"X_Deseret",                // 551 xx-Dsrt
	"X_Inherited",              // 552 xx-Qaai
	"X_Tagalog",                // 553 xx-Tglg
	"X_Hanunoo",                // 554 xx-Hano
	"X_Buhid",                  // 555 xx-Buhd
	"X_Tagbanwa",               // 556 xx-Tagb
	"X_Limbu",                  // 557 xx-Limb
	"X_Tai_Le",                 // 558 xx-Tale
	"X_Linear_B",               // 559 xx-Linb
	"X_Ugaritic",               // 560 xx-Ugar
	"X_Shavian",                // 561 xx-Shaw
	"X_Osmanya",                // 562 xx-Osma
	"X_Cypriot",                // 563 xx-Cprt
	"X_Braille",                // 564 xx-Brai
	"X_Buginese",               // 565 xx-Bugi
	"X_Coptic",                 // 566 xx-Copt
	"X_New_Tai_Lue",            // 567 xx-Talu
	"X_Glagolitic",             // 568 xx-Glag
	"X_Tifinagh",               // 569 xx-Tfng
	"X_Syloti_Nagri",           // 570 xx-Sylo
	"X_Old_Persian",            // 571 xx-Xpeo
	"X_Kharoshthi",             // 572 xx-Khar
	"X_Balinese",               // 573 xx-Bali
	"X_Cuneiform",              // 574 xx-Xsux
	"X_Phoenician",             // 575 xx-Phnx
	"X_Phags_Pa",               // 576 xx-Phag
	"X_Nko",                    // 577 xx-Nkoo
	"X_Sundanese",              // 578 xx-Sund
	"X_Lepcha",                 // 579 xx-Lepc
	"X_Ol_Chiki",               // 580 xx-Olck
	"X_Vai",                    // 581 xx-Vaii
	"X_Saurashtra",             // 582 xx-Saur
	"X_Kayah_Li",               // 583 xx-Kali
	"X_Rejang",                 // 584 xx-Rjng
	"X_Lycian",                 // 585 xx-Lyci
	"X_Carian",                 // 586 xx-Cari
	"X_Lydian",                 // 587 xx-Lydi
	"X_Cham",                   // 588 xx-Cham
	"X_Tai_Tham",               // 589 xx-Lana
	"X_Tai_Viet",               // 590 xx-Tavt
	"X_Avestan",                // 591 xx-Avst
	"X_Egyptian_Hieroglyphs",   // 592 xx-Egyp
	"X_Samaritan",              // 593 xx-Samr
	"X_Lisu",                   // 594 xx-Lisu
	"X_Bamum",                  // 595 xx-Bamu
	"X_Javanese",               // 596 xx-Java
	"X_Meetei_Mayek",           // 597 xx-Mtei
	"X_Imperial_Aramaic",       // 598 xx-Armi
	"X_Old_South_Arabian",      // 599 xx-Sarb
	"X_Inscriptional_Parthian", // 600 xx-Prti
	"X_Inscriptional_Pahlavi",  // 601 xx-Phli
	"X_Old_Turkic",             // 602 xx-Orkh
	"X_Kaithi",                 // 603 xx-Kthi
	"X_Batak",                  // 604 xx-Batk
	"X_Brahmi",                 // 605 xx-Brah
	"X_Mandaic",                // 606 xx-Mand
	"X_Chakma",                 // 607 xx-Cakm
	"X_Meroitic_Cursive",       // 608 xx-Merc
	"X_Meroitic_Hieroglyphs",   // 609 xx-Mero
	"X_Miao",                   // 610 xx-Plrd
	"X_Sharada",                // 611 xx-Shrd
	"X_Sora_Sompeng",           // 612 xx-Sora
	"X_Takri",                  // 613 xx-Takr
}
