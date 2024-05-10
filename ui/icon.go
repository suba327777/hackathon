package ui

import (
	"fmt"
	"strings"
)

var Mountain = `<path fill="#FFFFFF" fill-rule="evenodd" d="M18.4425,5.795625L12.450625,14.969875l-2.677,-4.09925L0,26.2045h32L18.4425,5.795625z M11.796875,21.410875l-2.978375,-2.08775l-5.275375,4.365125l6.246875,-9.801125l0.54825,0.839375l3.74825,6.4665L11.796875,21.410875z M17.4665,14.65075L11.19625,20.666375l-1.243375,-2.159875l5.065625,-7.756625l10.063875,15.149125L18.904125,20.855625l0.795625,4.248625L17.4665,14.65075z" />`
var Climber = `<svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 1819.000000 2771.000000" preserveAspectRatio="xMidYMid meet"><g transform="translate(0.000000,2771.000000) scale(0.100000,-0.100000)" fill="#FFFFFF"  fill-rule="evenodd"><path d="M8883 27695 c-705 -89 -1321 -472 -1714 -1064 -192 -289 -320 -620 -375 -966 -24 -157 -29 -463 -10 -629 109 -940 743 -1720 1638 -2015 272 -90 442 -116 753 -116 185 0 265 4 350 18 538 90 983 314 1350 682 369 368 591 809 682 1350 14 87 18 161 17 360 0 224 -3 265 -27 395 -62 346 -179 640 -366 921 -188 284 -408 502 -696 691 -278 183 -610 310 -950 364 -141 22 -507 27 -652 9z" /><path d="M2310 23095 c-178 -26 -319 -63 -478 -125 -684 -266 -1188 -898 -1291 -1620 -26 -179 -541 -6254 -541 -6379 0 -958 697 -1804 1646 -1996 168 -34 1693 -167 1914 -167 991 2 1859 748 2014 1732 27 167 546 6262 546 6405 0 680 -353 1323 -935 1708 -221 146 -501 256 -769 302 -48 8 -454 46 -901 85 -855 74 -1027 81 -1205 55z m1093 -1750 l748 -64 64 -29 c73 -35 136 -95 173 -167 20 -39 26 -68 29 -135 6 -118 -512 -6137 -533 -6201 -28 -87 -96 -165 -181 -208 -92 -46 -83 -47 -903 24 -421 36 -783 68 -805 71 -104 15 -209 94 -261 196 -21 41 -28 71 -32 133 -2 53 86 1117 258 3125 222 2575 265 3053 282 3097 24 65 100 150 166 186 55 30 135 46 202 40 25 -2 382 -33 793 -68z" /><path d="M8715 22383 c-22 -2 -86 -10 -143 -19 -954 -144 -1589 -814 -1786 -1884 -142 -773 -544 -5300 -686 -7720 -31 -519 -36 -945 -16 -1295 29 -515 36 -731 43 -1245 9 -796 -20 -1399 -93 -1925 -65 -468 -144 -736 -311 -1055 -490 -935 -2418 -4025 -2952 -4730 -226 -299 -351 -528 -425 -784 -101 -349 -62 -710 109 -1010 203 -356 533 -596 946 -688 136 -30 405 -32 541 -5 294 59 554 201 766 416 279 283 791 1018 1589 2281 892 1413 1610 2616 1951 3270 370 710 546 1395 636 2475 26 307 55 877 56 1077 l0 47 98 -102 c469 -492 1141 -1238 1411 -1567 185 -226 219 -277 263 -403 55 -155 136 -441 219 -766 38 -151 71 -276 72 -277 4 -4 1381 3491 1381 3507 2 45 -923 1061 -1648 1813 l-230 238 2 87 c18 628 193 3263 287 4311 26 289 110 1070 115 1070 3 0 59 -51 125 -112 459 -427 815 -740 930 -818 480 -325 1963 -878 4176 -1559 429 -132 451 -136 674 -136 192 0 272 12 424 62 399 129 731 452 869 844 59 166 77 275 77 464 0 182 -12 261 -62 419 -80 254 -251 501 -457 659 -173 134 -271 180 -581 277 -1698 529 -3009 982 -3470 1200 -92 43 -113 59 -255 192 -287 268 -842 810 -2286 2233 -341 336 -658 641 -704 678 -341 275 -718 434 -1129 476 -94 10 -430 12 -526 4z" /><path d="M15057 14758 c-242 -607 -5360 -13551 -5380 -13608 -27 -75 -31 -98 -31 -200 -1 -94 3 -129 22 -190 66 -210 216 -365 423 -435 73 -25 99 -28 199 -28 98 -1 127 3 195 26 44 15 109 43 144 64 78 46 196 170 233 244 34 69 5430 13710 5425 13715 -2 2 -271 96 -598 208 -327 113 -601 207 -610 211 -9 3 -19 0 -22 -7z" /><path d="M13901 6365 c-7 -16 -447 -1131 -978 -2477 l-966 -2448 33 -77 c167 -401 521 -701 960 -816 100 -26 138 -30 290 -34 193 -6 296 7 447 53 432 131 777 456 923 869 111 316 112 552 4 1305 -117 814 -426 2439 -668 3505 -31 141 -34 148 -45 120z" /></g></svg>`
var Tree = `<svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 2001.000000 2771.000000" preserveAspectRatio="xMidYMid meet"><g transform="translate(0.000000,2771.000000) scale(0.100000,-0.100000)" fill="#FFFFFF" fill-rule="evenodd"><path d="M9795 27704 c-1102 -121 -2565 -973 -4125 -2403 -1860 -1706 -3548 -3993 -4546 -6159 -634 -1376 -984 -2605 -1101 -3872 -25 -272 -25 -1085 0 -1400 78 -966 261 -1780 574 -2545 80 -196 255 -557 359 -742 1171 -2083 3402 -3399 6533 -3852 112 -17 204 -31 205 -33 2 -2 -215 -1470 -480 -3263 -266 -1793 -483 -3274 -484 -3290 0 -49 32 -98 80 -122 l44 -23 3148 0 c2637 0 3154 2 3181 14 42 17 85 68 92 109 3 19 -195 1382 -479 3302 -266 1799 -482 3271 -481 3273 2 2 94 17 205 33 2922 422 5054 1592 6285 3449 695 1047 1079 2272 1187 3785 20 279 17 1031 -5 1285 -106 1219 -452 2459 -1062 3805 -1150 2537 -3239 5218 -5430 6971 -1043 835 -2012 1380 -2830 1594 -232 60 -388 82 -620 85 -115 2 -228 1 -250 -1z m403 -1640 c552 -114 1422 -616 2322 -1339 1544 -1241 3112 -3085 4190 -4930 377 -645 738 -1377 983 -1995 395 -994 615 -1890 678 -2758 15 -215 6 -936 -15 -1162 -86 -918 -269 -1610 -602 -2280 -593 -1192 -1623 -2076 -3094 -2655 -616 -243 -1330 -435 -2075 -559 -248 -41 -506 -79 -510 -74 -1 2 -210 1409 -464 3128 -254 1719 -464 3137 -467 3152 -5 25 1 31 53 64 760 481 1412 1082 1941 1789 679 907 1142 1985 1337 3115 14 79 25 160 25 178 0 107 -107 175 -209 132 -27 -11 -55 -46 -146 -184 -795 -1211 -1849 -2202 -3117 -2931 l-183 -105 -8 33 c-4 17 -160 1069 -347 2337 -243 1654 -344 2315 -357 2340 -10 19 -34 45 -53 58 -70 47 -181 12 -211 -66 -5 -13 -270 -1790 -590 -3950 -319 -2160 -582 -3928 -584 -3930 -2 -2 -79 38 -172 88 -1119 614 -2089 1452 -2845 2462 -120 160 -146 189 -181 203 -97 39 -197 -29 -197 -134 0 -83 114 -514 214 -808 270 -795 671 -1526 1197 -2183 449 -562 1005 -1076 1582 -1462 64 -43 117 -84 117 -91 1 -15 -472 -3231 -476 -3235 -4 -6 -419 57 -654 99 -1374 245 -2551 703 -3431 1336 -424 305 -835 703 -1131 1098 -637 848 -1002 1939 -1077 3225 -40 668 -11 1181 100 1790 398 2176 1785 4733 3794 6990 407 457 972 1024 1400 1406 764 681 1530 1223 2180 1544 513 254 797 323 1083 264z" /></g></svg>`
var Triangle = `<svg  xmlns="http://www.w3.org/2000/svg" width="20" height="20" x="10" y="10" viewBox="0 0 264.58 264.58">
	  <path transform="matrix(.26458 0 0 .26458 -.001665 0)" d="m500 133c-10 0-20 5-25 15l-375 652c-30 50-25 75 25 75l375-0.26172 375 0.26172c50 0 55-25 25-75l-375-652c-5-10-15-15-25-15z" fill="#ADFF2F" stroke-width="3.7796"/>
  </svg>`
var Star = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" x="10" y="10" viewBox="0 0 1280.000000 1233.000000" preserveAspectRatio="xMidYMid meet">
    <g transform="translate(0.000000,1233.000000) scale(0.100000,-0.100000)"
    fill="#ffff00">
    <path d="M6310 11443 c-16 -16 -58 -80 -94 -143 -35 -63 -103 -182 -150 -265
    -94 -162 -262 -490 -346 -675 -81 -178 -564 -1227 -642 -1395 -114 -243 -226
    -460 -279 -540 -75 -112 -181 -215 -265 -258 -121 -62 -387 -125 -939 -222
    -82 -15 -222 -39 -310 -55 -88 -16 -263 -47 -390 -69 -1110 -196 -1285 -230
    -1623 -312 -241 -58 -288 -76 -298 -115 -10 -39 12 -72 139 -210 99 -106 183
    -206 262 -310 72 -95 309 -319 511 -481 21 -17 93 -76 160 -130 67 -54 207
    -167 311 -251 105 -84 203 -164 219 -177 16 -13 85 -69 154 -125 69 -56 134
    -109 145 -118 11 -9 68 -55 125 -102 58 -46 107 -87 110 -90 3 -3 34 -29 70
    -59 100 -84 300 -290 345 -356 112 -166 158 -314 157 -505 0 -154 -23 -305
    -108 -720 -36 -173 -69 -335 -74 -360 -5 -25 -36 -178 -70 -340 -137 -662
    -213 -1043 -239 -1198 -45 -261 -90 -686 -78 -732 20 -80 68 -83 247 -17 69
    25 179 62 245 83 66 20 161 52 210 71 85 33 436 204 462 225 7 5 52 31 100 57
    48 27 120 66 158 88 65 37 488 274 690 386 50 28 126 70 170 95 514 287 696
    371 915 424 80 19 308 16 395 -4 201 -49 324 -102 805 -350 354 -183 467 -242
    535 -278 39 -21 137 -72 218 -114 82 -42 165 -86 185 -98 84 -47 492 -248 552
    -271 119 -46 252 -85 423 -122 42 -9 123 -32 179 -51 179 -60 228 -43 228 82
    0 184 -62 610 -176 1202 -58 303 -123 644 -180 939 -197 1033 -204 1073 -211
    1263 -4 135 -3 151 21 233 26 88 96 240 116 252 6 4 14 15 18 25 14 37 186
    223 292 315 88 77 366 313 398 338 20 15 39 31 42 35 3 4 28 25 55 46 28 22
    55 44 60 50 6 6 33 29 60 50 28 22 58 47 67 55 9 9 23 21 30 26 29 23 124 101
    138 115 9 8 33 29 54 45 21 17 44 35 50 40 6 6 36 30 66 55 30 25 57 47 60 50
    3 3 30 25 60 50 58 47 105 88 151 130 15 14 59 52 98 84 78 67 282 273 351
    355 78 94 215 244 353 386 68 71 81 94 75 136 -5 39 -58 55 -288 93 -442 72
    -606 91 -2155 251 -711 73 -964 108 -1100 152 -176 56 -270 155 -431 453 -62
    116 -428 852 -601 1210 -194 402 -489 989 -560 1115 -172 305 -299 519 -329
    551 -38 41 -64 43 -104 7z"/>
    </g>
  </svg>`
var Sun = `<svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" x="10" y="10" viewBox="0 0 200 200"><path fill="#FDB86D" d="M95.783,198.255a5.964,5.964,0,0,1-1.75-4.3V168.069a5.967,5.967,0,1,1,11.935,0v25.882A5.968,5.968,0,0,1,100.088,200H100A5.969,5.969,0,0,1,95.783,198.255Zm66.425-27.6-18.3-18.3a5.967,5.967,0,1,1,8.438-8.438l18.3,18.3a5.967,5.967,0,1,1-8.438,8.438Zm-132.857.113a5.967,5.967,0,0,1,0-8.559l18.3-18.3a5.967,5.967,0,1,1,8.438,8.438l-18.3,18.3a5.967,5.967,0,0,1-8.434.121ZM54.131,100A45.869,45.869,0,1,1,100,145.869,45.869,45.869,0,0,1,54.131,100ZM162.1,100a5.967,5.967,0,0,1,5.968-5.967h25.882a5.968,5.968,0,1,1,0,11.934H168.069A5.968,5.968,0,0,1,162.1,100ZM6.05,105.966a5.967,5.967,0,1,1,0-11.933H31.932a5.967,5.967,0,1,1,0,11.935ZM143.914,56.094a5.966,5.966,0,0,1,0-8.438l18.3-18.3a5.967,5.967,0,1,1,8.438,8.438l-18.3,18.3a5.966,5.966,0,0,1-8.439,0Zm-96.258-.008-18.3-18.3a5.967,5.967,0,1,1,8.438-8.438l18.3,18.3a5.967,5.967,0,1,1-8.439,8.438ZM94.035,31.932V6.05q0-.039,0-.077a5.967,5.967,0,1,1,11.934,0q0,.038,0,.076V31.932a5.967,5.967,0,1,1-11.934,0Z" transform="translate(0 0)"/></svg>`
var Flag = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" x="10" y="10" viewBox="0 0 200 181.862"><path fill="#DA523A" d="M72.939,262.48a9.133,9.133,0,0,0-8.676,6.485L12.206,432.435a9.134,9.134,0,0,0,17.406,5.542l23.4-73.495,149.625,37.847-25.57-67.578,34.7-47.487L77.052,289l4.616-14.5a9.133,9.133,0,0,0-8.729-12.027Z" transform="translate(-11.774 -262.479)"/></svg>`

func changeIconColor(svgIcon, iconColor string) string {
	return strings.ReplaceAll(fmt.Sprintf("%s", svgIcon), `fill="#FFFFFF"`, fmt.Sprintf(`fill="%s"`, iconColor))
}

func generateTrianglePath(x, y int, triangleMountainColor string) string {
	triangle := strings.ReplaceAll(Triangle, `x="10" y="10"`, fmt.Sprintf(`x="%d" y="%d"`, x, y))
	triangle = strings.ReplaceAll(triangle, `fill="#ADFF2F"`, fmt.Sprintf(`fill="%s"`, triangleMountainColor))
	return triangle
}

func generateBonusIconPath(x, y int, svgIcon string) string {
	return strings.ReplaceAll(fmt.Sprintf("%s", svgIcon), `x="10" y="10"`, fmt.Sprintf(`x="%d" y="%d"`, x, y))
}
