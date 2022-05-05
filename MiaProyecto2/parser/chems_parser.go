// Code generated from Chems.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import AD "MiaProyecto2/AdminDisks"

var ValSize string = ""
var ValFit string = ""
var ValUnit string = ""
var ValPath string = ""
var ValType string = ""
var ValName string = ""
var ValIdentificador string = ""
var ValUsuario string = ""
var ValPassword string = ""
var ValGrupo string = ""
var ValR bool = false
var ValCont string = ""

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 363,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 55, 10, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 74, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 102, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 113, 10, 7,
	12, 7, 14, 7, 116, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 148,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 155, 10, 9, 12, 9, 14, 9, 158,
	11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 188, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 198,
	10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	208, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5,
	13, 218, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 228, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 235, 10,
	15, 12, 15, 14, 15, 238, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 253, 10, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 260, 10, 17, 12, 17, 14, 17,
	263, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 5, 18, 277, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 7, 19, 284, 10, 19, 12, 19, 14, 19, 287, 11, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 297, 10, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 7, 21, 304, 10, 21, 12, 21, 14, 21, 307, 11, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 329,
	10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 335, 10, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 5, 24, 361, 10, 24, 3, 24, 2, 8, 12, 16, 28, 32, 36, 40, 25, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 2, 2, 2, 391, 2, 54, 3, 2, 2, 2, 4, 73, 3, 2, 2, 2, 6, 75, 3, 2,
	2, 2, 8, 101, 3, 2, 2, 2, 10, 103, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14,
	147, 3, 2, 2, 2, 16, 149, 3, 2, 2, 2, 18, 187, 3, 2, 2, 2, 20, 197, 3,
	2, 2, 2, 22, 207, 3, 2, 2, 2, 24, 217, 3, 2, 2, 2, 26, 227, 3, 2, 2, 2,
	28, 229, 3, 2, 2, 2, 30, 252, 3, 2, 2, 2, 32, 254, 3, 2, 2, 2, 34, 276,
	3, 2, 2, 2, 36, 278, 3, 2, 2, 2, 38, 296, 3, 2, 2, 2, 40, 298, 3, 2, 2,
	2, 42, 328, 3, 2, 2, 2, 44, 334, 3, 2, 2, 2, 46, 360, 3, 2, 2, 2, 48, 55,
	5, 4, 3, 2, 49, 55, 5, 6, 4, 2, 50, 55, 5, 8, 5, 2, 51, 55, 5, 10, 6, 2,
	52, 53, 7, 37, 2, 2, 53, 55, 8, 2, 1, 2, 54, 48, 3, 2, 2, 2, 54, 49, 3,
	2, 2, 2, 54, 50, 3, 2, 2, 2, 54, 51, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55,
	3, 3, 2, 2, 2, 56, 57, 7, 3, 2, 2, 57, 58, 5, 16, 9, 2, 58, 59, 8, 3, 1,
	2, 59, 74, 3, 2, 2, 2, 60, 61, 7, 4, 2, 2, 61, 62, 5, 20, 11, 2, 62, 63,
	8, 3, 1, 2, 63, 74, 3, 2, 2, 2, 64, 65, 7, 5, 2, 2, 65, 66, 5, 28, 15,
	2, 66, 67, 8, 3, 1, 2, 67, 74, 3, 2, 2, 2, 68, 69, 7, 6, 2, 2, 69, 70,
	5, 32, 17, 2, 70, 71, 8, 3, 1, 2, 71, 74, 3, 2, 2, 2, 72, 74, 7, 2, 2,
	3, 73, 56, 3, 2, 2, 2, 73, 60, 3, 2, 2, 2, 73, 64, 3, 2, 2, 2, 73, 68,
	3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 5, 3, 2, 2, 2, 75, 76, 7, 7, 2, 2,
	76, 77, 5, 36, 19, 2, 77, 78, 8, 4, 1, 2, 78, 7, 3, 2, 2, 2, 79, 80, 7,
	8, 2, 2, 80, 81, 5, 40, 21, 2, 81, 82, 8, 5, 1, 2, 82, 102, 3, 2, 2, 2,
	83, 84, 7, 9, 2, 2, 84, 102, 8, 5, 1, 2, 85, 86, 7, 10, 2, 2, 86, 87, 5,
	22, 12, 2, 87, 88, 8, 5, 1, 2, 88, 102, 3, 2, 2, 2, 89, 90, 7, 11, 2, 2,
	90, 91, 5, 24, 13, 2, 91, 92, 8, 5, 1, 2, 92, 102, 3, 2, 2, 2, 93, 94,
	7, 12, 2, 2, 94, 95, 5, 44, 23, 2, 95, 96, 8, 5, 1, 2, 96, 102, 3, 2, 2,
	2, 97, 98, 7, 13, 2, 2, 98, 99, 5, 26, 14, 2, 99, 100, 8, 5, 1, 2, 100,
	102, 3, 2, 2, 2, 101, 79, 3, 2, 2, 2, 101, 83, 3, 2, 2, 2, 101, 85, 3,
	2, 2, 2, 101, 89, 3, 2, 2, 2, 101, 93, 3, 2, 2, 2, 101, 97, 3, 2, 2, 2,
	102, 9, 3, 2, 2, 2, 103, 104, 7, 14, 2, 2, 104, 105, 5, 12, 7, 2, 105,
	106, 8, 6, 1, 2, 106, 11, 3, 2, 2, 2, 107, 108, 8, 7, 1, 2, 108, 109, 5,
	14, 8, 2, 109, 114, 3, 2, 2, 2, 110, 111, 12, 3, 2, 2, 111, 113, 5, 14,
	8, 2, 112, 110, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2,
	114, 115, 3, 2, 2, 2, 115, 13, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 117, 118,
	7, 18, 2, 2, 118, 119, 7, 46, 2, 2, 119, 120, 7, 42, 2, 2, 120, 148, 8,
	8, 1, 2, 121, 122, 7, 35, 2, 2, 122, 123, 7, 46, 2, 2, 123, 124, 7, 25,
	2, 2, 124, 148, 8, 8, 1, 2, 125, 126, 7, 35, 2, 2, 126, 127, 7, 46, 2,
	2, 127, 128, 7, 43, 2, 2, 128, 148, 8, 8, 1, 2, 129, 130, 7, 35, 2, 2,
	130, 131, 7, 46, 2, 2, 131, 132, 7, 45, 2, 2, 132, 148, 8, 8, 1, 2, 133,
	134, 7, 20, 2, 2, 134, 135, 7, 46, 2, 2, 135, 136, 7, 41, 2, 2, 136, 148,
	8, 8, 1, 2, 137, 138, 7, 21, 2, 2, 138, 139, 7, 46, 2, 2, 139, 140, 7,
	43, 2, 2, 140, 148, 8, 8, 1, 2, 141, 142, 7, 21, 2, 2, 142, 143, 7, 46,
	2, 2, 143, 144, 7, 45, 2, 2, 144, 148, 8, 8, 1, 2, 145, 146, 7, 34, 2,
	2, 146, 148, 8, 8, 1, 2, 147, 117, 3, 2, 2, 2, 147, 121, 3, 2, 2, 2, 147,
	125, 3, 2, 2, 2, 147, 129, 3, 2, 2, 2, 147, 133, 3, 2, 2, 2, 147, 137,
	3, 2, 2, 2, 147, 141, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 15, 3, 2,
	2, 2, 149, 150, 8, 9, 1, 2, 150, 151, 5, 18, 10, 2, 151, 156, 3, 2, 2,
	2, 152, 153, 12, 3, 2, 2, 153, 155, 5, 18, 10, 2, 154, 152, 3, 2, 2, 2,
	155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157,
	17, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 7, 18, 2, 2, 160, 161,
	7, 46, 2, 2, 161, 162, 7, 42, 2, 2, 162, 188, 8, 10, 1, 2, 163, 164, 7,
	19, 2, 2, 164, 165, 7, 46, 2, 2, 165, 166, 7, 25, 2, 2, 166, 188, 8, 10,
	1, 2, 167, 168, 7, 19, 2, 2, 168, 169, 7, 46, 2, 2, 169, 170, 7, 26, 2,
	2, 170, 188, 8, 10, 1, 2, 171, 172, 7, 19, 2, 2, 172, 173, 7, 46, 2, 2,
	173, 174, 7, 27, 2, 2, 174, 188, 8, 10, 1, 2, 175, 176, 7, 20, 2, 2, 176,
	177, 7, 46, 2, 2, 177, 178, 7, 41, 2, 2, 178, 188, 8, 10, 1, 2, 179, 180,
	7, 21, 2, 2, 180, 181, 7, 46, 2, 2, 181, 182, 7, 43, 2, 2, 182, 188, 8,
	10, 1, 2, 183, 184, 7, 21, 2, 2, 184, 185, 7, 46, 2, 2, 185, 186, 7, 45,
	2, 2, 186, 188, 8, 10, 1, 2, 187, 159, 3, 2, 2, 2, 187, 163, 3, 2, 2, 2,
	187, 167, 3, 2, 2, 2, 187, 171, 3, 2, 2, 2, 187, 175, 3, 2, 2, 2, 187,
	179, 3, 2, 2, 2, 187, 183, 3, 2, 2, 2, 188, 19, 3, 2, 2, 2, 189, 190, 7,
	21, 2, 2, 190, 191, 7, 46, 2, 2, 191, 192, 7, 43, 2, 2, 192, 198, 8, 11,
	1, 2, 193, 194, 7, 21, 2, 2, 194, 195, 7, 46, 2, 2, 195, 196, 7, 45, 2,
	2, 196, 198, 8, 11, 1, 2, 197, 189, 3, 2, 2, 2, 197, 193, 3, 2, 2, 2, 198,
	21, 3, 2, 2, 2, 199, 200, 7, 22, 2, 2, 200, 201, 7, 46, 2, 2, 201, 202,
	7, 43, 2, 2, 202, 208, 8, 12, 1, 2, 203, 204, 7, 22, 2, 2, 204, 205, 7,
	46, 2, 2, 205, 206, 7, 44, 2, 2, 206, 208, 8, 12, 1, 2, 207, 199, 3, 2,
	2, 2, 207, 203, 3, 2, 2, 2, 208, 23, 3, 2, 2, 2, 209, 210, 7, 22, 2, 2,
	210, 211, 7, 46, 2, 2, 211, 212, 7, 43, 2, 2, 212, 218, 8, 13, 1, 2, 213,
	214, 7, 22, 2, 2, 214, 215, 7, 46, 2, 2, 215, 216, 7, 44, 2, 2, 216, 218,
	8, 13, 1, 2, 217, 209, 3, 2, 2, 2, 217, 213, 3, 2, 2, 2, 218, 25, 3, 2,
	2, 2, 219, 220, 7, 22, 2, 2, 220, 221, 7, 46, 2, 2, 221, 222, 7, 43, 2,
	2, 222, 228, 8, 14, 1, 2, 223, 224, 7, 22, 2, 2, 224, 225, 7, 46, 2, 2,
	225, 226, 7, 44, 2, 2, 226, 228, 8, 14, 1, 2, 227, 219, 3, 2, 2, 2, 227,
	223, 3, 2, 2, 2, 228, 27, 3, 2, 2, 2, 229, 230, 8, 15, 1, 2, 230, 231,
	5, 30, 16, 2, 231, 236, 3, 2, 2, 2, 232, 233, 12, 3, 2, 2, 233, 235, 5,
	30, 16, 2, 234, 232, 3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2,
	2, 2, 236, 237, 3, 2, 2, 2, 237, 29, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2,
	239, 253, 5, 18, 10, 2, 240, 241, 7, 23, 2, 2, 241, 242, 7, 46, 2, 2, 242,
	243, 7, 41, 2, 2, 243, 253, 8, 16, 1, 2, 244, 245, 7, 22, 2, 2, 245, 246,
	7, 46, 2, 2, 246, 247, 7, 43, 2, 2, 247, 253, 8, 16, 1, 2, 248, 249, 7,
	22, 2, 2, 249, 250, 7, 46, 2, 2, 250, 251, 7, 44, 2, 2, 251, 253, 8, 16,
	1, 2, 252, 239, 3, 2, 2, 2, 252, 240, 3, 2, 2, 2, 252, 244, 3, 2, 2, 2,
	252, 248, 3, 2, 2, 2, 253, 31, 3, 2, 2, 2, 254, 255, 8, 17, 1, 2, 255,
	256, 5, 34, 18, 2, 256, 261, 3, 2, 2, 2, 257, 258, 12, 3, 2, 2, 258, 260,
	5, 34, 18, 2, 259, 257, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2, 261, 259, 3,
	2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 33, 3, 2, 2, 2, 263, 261, 3, 2, 2,
	2, 264, 265, 7, 21, 2, 2, 265, 266, 7, 46, 2, 2, 266, 267, 7, 43, 2, 2,
	267, 277, 8, 18, 1, 2, 268, 269, 7, 21, 2, 2, 269, 270, 7, 46, 2, 2, 270,
	271, 7, 45, 2, 2, 271, 277, 8, 18, 1, 2, 272, 273, 7, 22, 2, 2, 273, 274,
	7, 46, 2, 2, 274, 275, 7, 44, 2, 2, 275, 277, 8, 18, 1, 2, 276, 264, 3,
	2, 2, 2, 276, 268, 3, 2, 2, 2, 276, 272, 3, 2, 2, 2, 277, 35, 3, 2, 2,
	2, 278, 279, 8, 19, 1, 2, 279, 280, 5, 38, 20, 2, 280, 285, 3, 2, 2, 2,
	281, 282, 12, 3, 2, 2, 282, 284, 5, 38, 20, 2, 283, 281, 3, 2, 2, 2, 284,
	287, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 37, 3,
	2, 2, 2, 287, 285, 3, 2, 2, 2, 288, 289, 7, 24, 2, 2, 289, 290, 7, 46,
	2, 2, 290, 291, 7, 44, 2, 2, 291, 297, 8, 20, 1, 2, 292, 293, 7, 23, 2,
	2, 293, 294, 7, 46, 2, 2, 294, 295, 7, 41, 2, 2, 295, 297, 8, 20, 1, 2,
	296, 288, 3, 2, 2, 2, 296, 292, 3, 2, 2, 2, 297, 39, 3, 2, 2, 2, 298, 299,
	8, 21, 1, 2, 299, 300, 5, 42, 22, 2, 300, 305, 3, 2, 2, 2, 301, 302, 12,
	3, 2, 2, 302, 304, 5, 42, 22, 2, 303, 301, 3, 2, 2, 2, 304, 307, 3, 2,
	2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 41, 3, 2, 2, 2,
	307, 305, 3, 2, 2, 2, 308, 309, 7, 30, 2, 2, 309, 310, 7, 46, 2, 2, 310,
	311, 7, 44, 2, 2, 311, 329, 8, 22, 1, 2, 312, 313, 7, 30, 2, 2, 313, 314,
	7, 46, 2, 2, 314, 315, 7, 43, 2, 2, 315, 329, 8, 22, 1, 2, 316, 317, 7,
	31, 2, 2, 317, 318, 7, 46, 2, 2, 318, 319, 7, 44, 2, 2, 319, 329, 8, 22,
	1, 2, 320, 321, 7, 31, 2, 2, 321, 322, 7, 46, 2, 2, 322, 323, 7, 43, 2,
	2, 323, 329, 8, 22, 1, 2, 324, 325, 7, 24, 2, 2, 325, 326, 7, 46, 2, 2,
	326, 327, 7, 44, 2, 2, 327, 329, 8, 22, 1, 2, 328, 308, 3, 2, 2, 2, 328,
	312, 3, 2, 2, 2, 328, 316, 3, 2, 2, 2, 328, 320, 3, 2, 2, 2, 328, 324,
	3, 2, 2, 2, 329, 43, 3, 2, 2, 2, 330, 335, 5, 46, 24, 2, 331, 332, 5, 40,
	21, 2, 332, 333, 5, 46, 24, 2, 333, 335, 3, 2, 2, 2, 334, 330, 3, 2, 2,
	2, 334, 331, 3, 2, 2, 2, 335, 45, 3, 2, 2, 2, 336, 337, 7, 30, 2, 2, 337,
	338, 7, 46, 2, 2, 338, 339, 7, 44, 2, 2, 339, 361, 8, 24, 1, 2, 340, 341,
	7, 30, 2, 2, 341, 342, 7, 46, 2, 2, 342, 343, 7, 43, 2, 2, 343, 361, 8,
	24, 1, 2, 344, 345, 7, 32, 2, 2, 345, 346, 7, 46, 2, 2, 346, 347, 7, 44,
	2, 2, 347, 361, 8, 24, 1, 2, 348, 349, 7, 32, 2, 2, 349, 350, 7, 46, 2,
	2, 350, 351, 7, 43, 2, 2, 351, 361, 8, 24, 1, 2, 352, 353, 7, 33, 2, 2,
	353, 354, 7, 46, 2, 2, 354, 355, 7, 44, 2, 2, 355, 361, 8, 24, 1, 2, 356,
	357, 7, 33, 2, 2, 357, 358, 7, 46, 2, 2, 358, 359, 7, 43, 2, 2, 359, 361,
	8, 24, 1, 2, 360, 336, 3, 2, 2, 2, 360, 340, 3, 2, 2, 2, 360, 344, 3, 2,
	2, 2, 360, 348, 3, 2, 2, 2, 360, 352, 3, 2, 2, 2, 360, 356, 3, 2, 2, 2,
	361, 47, 3, 2, 2, 2, 23, 54, 73, 101, 114, 147, 156, 187, 197, 207, 217,
	227, 236, 252, 261, 276, 285, 296, 305, 328, 334, 360,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "'='",
}
var symbolicNames = []string{
	"", "MKDISK", "RMDISK", "FDISK", "MOUNT", "MKFS", "LOGIN", "LOGOUT", "MKGRP",
	"RMGRP", "MKUSER", "RMUSR", "MKFILE", "MKDIR", "EXEC", "REP", "SIZE", "FIT",
	"UNIT", "PATH", "NAME", "TYPE", "ID", "BF", "FF", "WF", "FAST", "FULL",
	"USR", "PWD", "PWD1", "GRP", "GR", "CONT", "GP", "PAUSA", "DISK", "TREE",
	"FILE", "LETRA", "ENTERO", "CADENA", "IDENTIFICADOR", "RUTA", "IGUAL",
	"WHITESPACE", "COMENTARIOS",
}

var ruleNames = []string{
	"start", "adminDiscos", "adminArchivos", "adminUsuarios", "adminCarpetas",
	"pmkfile", "parametrosmkfile", "pmkdisk", "parametrosmkdisk", "prmdisk",
	"pmkgrp", "prmkgrp", "prmkusr", "pfdisk", "parametrosfdisk", "pmount",
	"parametrosmount", "pmkfs", "parametrosmkfs", "plogin", "parametroslogin",
	"pmkuser", "parametrosmkuser",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Chems struct {
	*antlr.BaseParser
}

func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF           = antlr.TokenEOF
	ChemsMKDISK        = 1
	ChemsRMDISK        = 2
	ChemsFDISK         = 3
	ChemsMOUNT         = 4
	ChemsMKFS          = 5
	ChemsLOGIN         = 6
	ChemsLOGOUT        = 7
	ChemsMKGRP         = 8
	ChemsRMGRP         = 9
	ChemsMKUSER        = 10
	ChemsRMUSR         = 11
	ChemsMKFILE        = 12
	ChemsMKDIR         = 13
	ChemsEXEC          = 14
	ChemsREP           = 15
	ChemsSIZE          = 16
	ChemsFIT           = 17
	ChemsUNIT          = 18
	ChemsPATH          = 19
	ChemsNAME          = 20
	ChemsTYPE          = 21
	ChemsID            = 22
	ChemsBF            = 23
	ChemsFF            = 24
	ChemsWF            = 25
	ChemsFAST          = 26
	ChemsFULL          = 27
	ChemsUSR           = 28
	ChemsPWD           = 29
	ChemsPWD1          = 30
	ChemsGRP           = 31
	ChemsGR            = 32
	ChemsCONT          = 33
	ChemsGP            = 34
	ChemsPAUSA         = 35
	ChemsDISK          = 36
	ChemsTREE          = 37
	ChemsFILE          = 38
	ChemsLETRA         = 39
	ChemsENTERO        = 40
	ChemsCADENA        = 41
	ChemsIDENTIFICADOR = 42
	ChemsRUTA          = 43
	ChemsIGUAL         = 44
	ChemsWHITESPACE    = 45
	ChemsCOMENTARIOS   = 46
)

// Chems rules.
const (
	ChemsRULE_start            = 0
	ChemsRULE_adminDiscos      = 1
	ChemsRULE_adminArchivos    = 2
	ChemsRULE_adminUsuarios    = 3
	ChemsRULE_adminCarpetas    = 4
	ChemsRULE_pmkfile          = 5
	ChemsRULE_parametrosmkfile = 6
	ChemsRULE_pmkdisk          = 7
	ChemsRULE_parametrosmkdisk = 8
	ChemsRULE_prmdisk          = 9
	ChemsRULE_pmkgrp           = 10
	ChemsRULE_prmkgrp          = 11
	ChemsRULE_prmkusr          = 12
	ChemsRULE_pfdisk           = 13
	ChemsRULE_parametrosfdisk  = 14
	ChemsRULE_pmount           = 15
	ChemsRULE_parametrosmount  = 16
	ChemsRULE_pmkfs            = 17
	ChemsRULE_parametrosmkfs   = 18
	ChemsRULE_plogin           = 19
	ChemsRULE_parametroslogin  = 20
	ChemsRULE_pmkuser          = 21
	ChemsRULE_parametrosmkuser = 22
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AdminDiscos() IAdminDiscosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminDiscosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminDiscosContext)
}

func (s *StartContext) AdminArchivos() IAdminArchivosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminArchivosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminArchivosContext)
}

func (s *StartContext) AdminUsuarios() IAdminUsuariosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminUsuariosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminUsuariosContext)
}

func (s *StartContext) AdminCarpetas() IAdminCarpetasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdminCarpetasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdminCarpetasContext)
}

func (s *StartContext) PAUSA() antlr.TerminalNode {
	return s.GetToken(ChemsPAUSA, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsEOF, ChemsMKDISK, ChemsRMDISK, ChemsFDISK, ChemsMOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.AdminDiscos()
		}

	case ChemsMKFS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.AdminArchivos()
		}

	case ChemsLOGIN, ChemsLOGOUT, ChemsMKGRP, ChemsRMGRP, ChemsMKUSER, ChemsRMUSR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.AdminUsuarios()
		}

	case ChemsMKFILE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.AdminCarpetas()
		}

	case ChemsPAUSA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.Match(ChemsPAUSA)
		}
		AD.Pausa()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAdminDiscosContext is an interface to support dynamic dispatch.
type IAdminDiscosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdminDiscosContext differentiates from other interfaces.
	IsAdminDiscosContext()
}

type AdminDiscosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminDiscosContext() *AdminDiscosContext {
	var p = new(AdminDiscosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_adminDiscos
	return p
}

func (*AdminDiscosContext) IsAdminDiscosContext() {}

func NewAdminDiscosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminDiscosContext {
	var p = new(AdminDiscosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_adminDiscos

	return p
}

func (s *AdminDiscosContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminDiscosContext) MKDISK() antlr.TerminalNode {
	return s.GetToken(ChemsMKDISK, 0)
}

func (s *AdminDiscosContext) Pmkdisk() IPmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkdiskContext)
}

func (s *AdminDiscosContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(ChemsRMDISK, 0)
}

func (s *AdminDiscosContext) Prmdisk() IPrmdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrmdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrmdiskContext)
}

func (s *AdminDiscosContext) FDISK() antlr.TerminalNode {
	return s.GetToken(ChemsFDISK, 0)
}

func (s *AdminDiscosContext) Pfdisk() IPfdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPfdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPfdiskContext)
}

func (s *AdminDiscosContext) MOUNT() antlr.TerminalNode {
	return s.GetToken(ChemsMOUNT, 0)
}

func (s *AdminDiscosContext) Pmount() IPmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmountContext)
}

func (s *AdminDiscosContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChemsEOF, 0)
}

func (s *AdminDiscosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminDiscosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminDiscosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAdminDiscos(s)
	}
}

func (s *AdminDiscosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAdminDiscos(s)
	}
}

func (p *Chems) AdminDiscos() (localctx IAdminDiscosContext) {
	localctx = NewAdminDiscosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_adminDiscos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(ChemsMKDISK)
		}
		{
			p.SetState(55)
			p.pmkdisk(0)
		}
		AD.Mkdisk(ValSize, ValFit, ValUnit, ValPath)
		ValSize = ""
		ValFit = ""
		ValUnit = ""
		ValPath = ""

	case ChemsRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Match(ChemsRMDISK)
		}
		{
			p.SetState(59)
			p.Prmdisk()
		}
		AD.Rmdisk(ValPath)
		ValPath = ""

	case ChemsFDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Match(ChemsFDISK)
		}
		{
			p.SetState(63)
			p.pfdisk(0)
		}
		AD.Fdisk(ValSize, ValFit, ValUnit, ValPath, ValType, ValName)
		ValSize = ""
		ValFit = ""
		ValUnit = ""
		ValPath = ""
		ValType = ""
		ValName = ""

	case ChemsMOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(66)
			p.Match(ChemsMOUNT)
		}
		{
			p.SetState(67)
			p.pmount(0)
		}
		AD.Mount(ValPath, ValName)
		ValPath = ""
		ValName = ""

	case ChemsEOF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(70)
			p.Match(ChemsEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAdminArchivosContext is an interface to support dynamic dispatch.
type IAdminArchivosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdminArchivosContext differentiates from other interfaces.
	IsAdminArchivosContext()
}

type AdminArchivosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminArchivosContext() *AdminArchivosContext {
	var p = new(AdminArchivosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_adminArchivos
	return p
}

func (*AdminArchivosContext) IsAdminArchivosContext() {}

func NewAdminArchivosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminArchivosContext {
	var p = new(AdminArchivosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_adminArchivos

	return p
}

func (s *AdminArchivosContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminArchivosContext) MKFS() antlr.TerminalNode {
	return s.GetToken(ChemsMKFS, 0)
}

func (s *AdminArchivosContext) Pmkfs() IPmkfsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkfsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkfsContext)
}

func (s *AdminArchivosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminArchivosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminArchivosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAdminArchivos(s)
	}
}

func (s *AdminArchivosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAdminArchivos(s)
	}
}

func (p *Chems) AdminArchivos() (localctx IAdminArchivosContext) {
	localctx = NewAdminArchivosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_adminArchivos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(ChemsMKFS)
	}
	{
		p.SetState(74)
		p.pmkfs(0)
	}
	AD.Mkfs(ValIdentificador, ValType)
	ValIdentificador = ""
	ValType = ""

	return localctx
}

// IAdminUsuariosContext is an interface to support dynamic dispatch.
type IAdminUsuariosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdminUsuariosContext differentiates from other interfaces.
	IsAdminUsuariosContext()
}

type AdminUsuariosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminUsuariosContext() *AdminUsuariosContext {
	var p = new(AdminUsuariosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_adminUsuarios
	return p
}

func (*AdminUsuariosContext) IsAdminUsuariosContext() {}

func NewAdminUsuariosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminUsuariosContext {
	var p = new(AdminUsuariosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_adminUsuarios

	return p
}

func (s *AdminUsuariosContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminUsuariosContext) LOGIN() antlr.TerminalNode {
	return s.GetToken(ChemsLOGIN, 0)
}

func (s *AdminUsuariosContext) Plogin() IPloginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPloginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPloginContext)
}

func (s *AdminUsuariosContext) LOGOUT() antlr.TerminalNode {
	return s.GetToken(ChemsLOGOUT, 0)
}

func (s *AdminUsuariosContext) MKGRP() antlr.TerminalNode {
	return s.GetToken(ChemsMKGRP, 0)
}

func (s *AdminUsuariosContext) Pmkgrp() IPmkgrpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkgrpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkgrpContext)
}

func (s *AdminUsuariosContext) RMGRP() antlr.TerminalNode {
	return s.GetToken(ChemsRMGRP, 0)
}

func (s *AdminUsuariosContext) Prmkgrp() IPrmkgrpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrmkgrpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrmkgrpContext)
}

func (s *AdminUsuariosContext) MKUSER() antlr.TerminalNode {
	return s.GetToken(ChemsMKUSER, 0)
}

func (s *AdminUsuariosContext) Pmkuser() IPmkuserContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkuserContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkuserContext)
}

func (s *AdminUsuariosContext) RMUSR() antlr.TerminalNode {
	return s.GetToken(ChemsRMUSR, 0)
}

func (s *AdminUsuariosContext) Prmkusr() IPrmkusrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrmkusrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrmkusrContext)
}

func (s *AdminUsuariosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminUsuariosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminUsuariosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAdminUsuarios(s)
	}
}

func (s *AdminUsuariosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAdminUsuarios(s)
	}
}

func (p *Chems) AdminUsuarios() (localctx IAdminUsuariosContext) {
	localctx = NewAdminUsuariosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_adminUsuarios)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsLOGIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(ChemsLOGIN)
		}
		{
			p.SetState(78)
			p.plogin(0)
		}
		AD.Login(ValUsuario, ValPassword, ValIdentificador)
		ValUsuario = ""
		ValPassword = ""
		ValIdentificador = ""

	case ChemsLOGOUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(ChemsLOGOUT)
		}
		AD.LogOut()

	case ChemsMKGRP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.Match(ChemsMKGRP)
		}
		{
			p.SetState(84)
			p.Pmkgrp()
		}
		AD.Mkgrp(ValName)
		ValName = ""

	case ChemsRMGRP:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Match(ChemsRMGRP)
		}
		{
			p.SetState(88)
			p.Prmkgrp()
		}
		AD.RMkgrp(ValName)
		ValName = ""

	case ChemsMKUSER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.Match(ChemsMKUSER)
		}
		{
			p.SetState(92)
			p.Pmkuser()
		}
		AD.Mkusr(ValUsuario, ValPassword, ValGrupo)
		ValUsuario = ""
		ValPassword = ""
		ValGrupo = ""

	case ChemsRMUSR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(ChemsRMUSR)
		}
		{
			p.SetState(96)
			p.Prmkusr()
		}
		AD.RMkusr(ValName)
		ValName = ""

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAdminCarpetasContext is an interface to support dynamic dispatch.
type IAdminCarpetasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdminCarpetasContext differentiates from other interfaces.
	IsAdminCarpetasContext()
}

type AdminCarpetasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminCarpetasContext() *AdminCarpetasContext {
	var p = new(AdminCarpetasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_adminCarpetas
	return p
}

func (*AdminCarpetasContext) IsAdminCarpetasContext() {}

func NewAdminCarpetasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminCarpetasContext {
	var p = new(AdminCarpetasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_adminCarpetas

	return p
}

func (s *AdminCarpetasContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminCarpetasContext) MKFILE() antlr.TerminalNode {
	return s.GetToken(ChemsMKFILE, 0)
}

func (s *AdminCarpetasContext) Pmkfile() IPmkfileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkfileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkfileContext)
}

func (s *AdminCarpetasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminCarpetasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminCarpetasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAdminCarpetas(s)
	}
}

func (s *AdminCarpetasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAdminCarpetas(s)
	}
}

func (p *Chems) AdminCarpetas() (localctx IAdminCarpetasContext) {
	localctx = NewAdminCarpetasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_adminCarpetas)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(ChemsMKFILE)
	}
	{
		p.SetState(102)
		p.pmkfile(0)
	}
	AD.MkFile(ValPath, ValR, ValSize, ValCont)
	ValPath = ""
	ValR = false
	ValSize = ""
	ValCont = ""

	return localctx
}

// IPmkfileContext is an interface to support dynamic dispatch.
type IPmkfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmkfileContext differentiates from other interfaces.
	IsPmkfileContext()
}

type PmkfileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmkfileContext() *PmkfileContext {
	var p = new(PmkfileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmkfile
	return p
}

func (*PmkfileContext) IsPmkfileContext() {}

func NewPmkfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmkfileContext {
	var p = new(PmkfileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmkfile

	return p
}

func (s *PmkfileContext) GetParser() antlr.Parser { return s.parser }

func (s *PmkfileContext) Parametrosmkfile() IParametrosmkfileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkfileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkfileContext)
}

func (s *PmkfileContext) Pmkfile() IPmkfileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkfileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkfileContext)
}

func (s *PmkfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmkfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmkfileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmkfile(s)
	}
}

func (s *PmkfileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmkfile(s)
	}
}

func (p *Chems) Pmkfile() (localctx IPmkfileContext) {
	return p.pmkfile(0)
}

func (p *Chems) pmkfile(_p int) (localctx IPmkfileContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPmkfileContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPmkfileContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ChemsRULE_pmkfile, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Parametrosmkfile()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPmkfileContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pmkfile)
			p.SetState(108)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(109)
				p.Parametrosmkfile()
			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosmkfileContext is an interface to support dynamic dispatch.
type IParametrosmkfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Get_LETRA returns the _LETRA token.
	Get_LETRA() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// Set_LETRA sets the _LETRA token.
	Set_LETRA(antlr.Token)

	// IsParametrosmkfileContext differentiates from other interfaces.
	IsParametrosmkfileContext()
}

type ParametrosmkfileContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_ENTERO antlr.Token
	_CADENA antlr.Token
	_RUTA   antlr.Token
	_LETRA  antlr.Token
}

func NewEmptyParametrosmkfileContext() *ParametrosmkfileContext {
	var p = new(ParametrosmkfileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmkfile
	return p
}

func (*ParametrosmkfileContext) IsParametrosmkfileContext() {}

func NewParametrosmkfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmkfileContext {
	var p = new(ParametrosmkfileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmkfile

	return p
}

func (s *ParametrosmkfileContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmkfileContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *ParametrosmkfileContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosmkfileContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *ParametrosmkfileContext) Get_LETRA() antlr.Token { return s._LETRA }

func (s *ParametrosmkfileContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ParametrosmkfileContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosmkfileContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *ParametrosmkfileContext) Set_LETRA(v antlr.Token) { s._LETRA = v }

func (s *ParametrosmkfileContext) SIZE() antlr.TerminalNode {
	return s.GetToken(ChemsSIZE, 0)
}

func (s *ParametrosmkfileContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmkfileContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(ChemsENTERO, 0)
}

func (s *ParametrosmkfileContext) CONT() antlr.TerminalNode {
	return s.GetToken(ChemsCONT, 0)
}

func (s *ParametrosmkfileContext) BF() antlr.TerminalNode {
	return s.GetToken(ChemsBF, 0)
}

func (s *ParametrosmkfileContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosmkfileContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *ParametrosmkfileContext) UNIT() antlr.TerminalNode {
	return s.GetToken(ChemsUNIT, 0)
}

func (s *ParametrosmkfileContext) LETRA() antlr.TerminalNode {
	return s.GetToken(ChemsLETRA, 0)
}

func (s *ParametrosmkfileContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *ParametrosmkfileContext) GR() antlr.TerminalNode {
	return s.GetToken(ChemsGR, 0)
}

func (s *ParametrosmkfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmkfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmkfileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmkfile(s)
	}
}

func (s *ParametrosmkfileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmkfile(s)
	}
}

func (p *Chems) Parametrosmkfile() (localctx IParametrosmkfileContext) {
	localctx = NewParametrosmkfileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChemsRULE_parametrosmkfile)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(ChemsSIZE)
		}
		{
			p.SetState(116)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(117)

			var _m = p.Match(ChemsENTERO)

			localctx.(*ParametrosmkfileContext)._ENTERO = _m
		}

		ValSize = (func() string {
			if localctx.(*ParametrosmkfileContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_ENTERO().GetText()
			}
		}())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Match(ChemsCONT)
		}
		{
			p.SetState(120)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(121)
			p.Match(ChemsBF)
		}
		ValCont = ""

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.Match(ChemsCONT)
		}
		{
			p.SetState(124)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(125)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkfileContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmkfileContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmkfileContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValCont = str

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.Match(ChemsCONT)
		}
		{
			p.SetState(128)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(129)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmkfileContext)._RUTA = _m
		}
		ValCont = (func() string {
			if localctx.(*ParametrosmkfileContext).Get_RUTA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_RUTA().GetText()
			}
		}())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Match(ChemsUNIT)
		}
		{
			p.SetState(132)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(133)

			var _m = p.Match(ChemsLETRA)

			localctx.(*ParametrosmkfileContext)._LETRA = _m
		}
		ValUnit = (func() string {
			if localctx.(*ParametrosmkfileContext).Get_LETRA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_LETRA().GetText()
			}
		}())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(135)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(136)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(137)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkfileContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmkfileContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmkfileContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValPath = str

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(139)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(140)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(141)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmkfileContext)._RUTA = _m
		}
		ValPath = (func() string {
			if localctx.(*ParametrosmkfileContext).Get_RUTA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfileContext).Get_RUTA().GetText()
			}
		}())

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(143)
			p.Match(ChemsGR)
		}
		ValR = true

	}

	return localctx
}

// IPmkdiskContext is an interface to support dynamic dispatch.
type IPmkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmkdiskContext differentiates from other interfaces.
	IsPmkdiskContext()
}

type PmkdiskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmkdiskContext() *PmkdiskContext {
	var p = new(PmkdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmkdisk
	return p
}

func (*PmkdiskContext) IsPmkdiskContext() {}

func NewPmkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmkdiskContext {
	var p = new(PmkdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmkdisk

	return p
}

func (s *PmkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *PmkdiskContext) Parametrosmkdisk() IParametrosmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkdiskContext)
}

func (s *PmkdiskContext) Pmkdisk() IPmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkdiskContext)
}

func (s *PmkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmkdisk(s)
	}
}

func (s *PmkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmkdisk(s)
	}
}

func (p *Chems) Pmkdisk() (localctx IPmkdiskContext) {
	return p.pmkdisk(0)
}

func (p *Chems) pmkdisk(_p int) (localctx IPmkdiskContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPmkdiskContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPmkdiskContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ChemsRULE_pmkdisk, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Parametrosmkdisk()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPmkdiskContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pmkdisk)
			p.SetState(150)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(151)
				p.Parametrosmkdisk()
			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosmkdiskContext is an interface to support dynamic dispatch.
type IParametrosmkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_LETRA returns the _LETRA token.
	Get_LETRA() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_LETRA sets the _LETRA token.
	Set_LETRA(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// IsParametrosmkdiskContext differentiates from other interfaces.
	IsParametrosmkdiskContext()
}

type ParametrosmkdiskContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_ENTERO antlr.Token
	_LETRA  antlr.Token
	_CADENA antlr.Token
	_RUTA   antlr.Token
}

func NewEmptyParametrosmkdiskContext() *ParametrosmkdiskContext {
	var p = new(ParametrosmkdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmkdisk
	return p
}

func (*ParametrosmkdiskContext) IsParametrosmkdiskContext() {}

func NewParametrosmkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmkdiskContext {
	var p = new(ParametrosmkdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmkdisk

	return p
}

func (s *ParametrosmkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmkdiskContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *ParametrosmkdiskContext) Get_LETRA() antlr.Token { return s._LETRA }

func (s *ParametrosmkdiskContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosmkdiskContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *ParametrosmkdiskContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *ParametrosmkdiskContext) Set_LETRA(v antlr.Token) { s._LETRA = v }

func (s *ParametrosmkdiskContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosmkdiskContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *ParametrosmkdiskContext) SIZE() antlr.TerminalNode {
	return s.GetToken(ChemsSIZE, 0)
}

func (s *ParametrosmkdiskContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmkdiskContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(ChemsENTERO, 0)
}

func (s *ParametrosmkdiskContext) FIT() antlr.TerminalNode {
	return s.GetToken(ChemsFIT, 0)
}

func (s *ParametrosmkdiskContext) BF() antlr.TerminalNode {
	return s.GetToken(ChemsBF, 0)
}

func (s *ParametrosmkdiskContext) FF() antlr.TerminalNode {
	return s.GetToken(ChemsFF, 0)
}

func (s *ParametrosmkdiskContext) WF() antlr.TerminalNode {
	return s.GetToken(ChemsWF, 0)
}

func (s *ParametrosmkdiskContext) UNIT() antlr.TerminalNode {
	return s.GetToken(ChemsUNIT, 0)
}

func (s *ParametrosmkdiskContext) LETRA() antlr.TerminalNode {
	return s.GetToken(ChemsLETRA, 0)
}

func (s *ParametrosmkdiskContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *ParametrosmkdiskContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosmkdiskContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *ParametrosmkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmkdisk(s)
	}
}

func (s *ParametrosmkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmkdisk(s)
	}
}

func (p *Chems) Parametrosmkdisk() (localctx IParametrosmkdiskContext) {
	localctx = NewParametrosmkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_parametrosmkdisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(ChemsSIZE)
		}
		{
			p.SetState(158)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(159)

			var _m = p.Match(ChemsENTERO)

			localctx.(*ParametrosmkdiskContext)._ENTERO = _m
		}
		ValSize = (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_ENTERO().GetText()
			}
		}())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(162)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(163)
			p.Match(ChemsBF)
		}
		ValFit = "b"

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(166)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(167)
			p.Match(ChemsFF)
		}
		ValFit = "f"

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(169)
			p.Match(ChemsFIT)
		}
		{
			p.SetState(170)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(171)
			p.Match(ChemsWF)
		}
		ValFit = "w"

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Match(ChemsUNIT)
		}
		{
			p.SetState(174)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(175)

			var _m = p.Match(ChemsLETRA)

			localctx.(*ParametrosmkdiskContext)._LETRA = _m
		}
		ValUnit = (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_LETRA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_LETRA().GetText()
			}
		}())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(177)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(178)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(179)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkdiskContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmkdiskContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValPath = str

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(181)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(182)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(183)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmkdiskContext)._RUTA = _m
		}
		ValPath = (func() string {
			if localctx.(*ParametrosmkdiskContext).Get_RUTA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkdiskContext).Get_RUTA().GetText()
			}
		}())

	}

	return localctx
}

// IPrmdiskContext is an interface to support dynamic dispatch.
type IPrmdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// IsPrmdiskContext differentiates from other interfaces.
	IsPrmdiskContext()
}

type PrmdiskContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_CADENA antlr.Token
	_RUTA   antlr.Token
}

func NewEmptyPrmdiskContext() *PrmdiskContext {
	var p = new(PrmdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_prmdisk
	return p
}

func (*PrmdiskContext) IsPrmdiskContext() {}

func NewPrmdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrmdiskContext {
	var p = new(PrmdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_prmdisk

	return p
}

func (s *PrmdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *PrmdiskContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrmdiskContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *PrmdiskContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrmdiskContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *PrmdiskContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *PrmdiskContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *PrmdiskContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *PrmdiskContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *PrmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrmdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrmdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrmdisk(s)
	}
}

func (s *PrmdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrmdisk(s)
	}
}

func (p *Chems) Prmdisk() (localctx IPrmdiskContext) {
	localctx = NewPrmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChemsRULE_prmdisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(188)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(189)

			var _m = p.Match(ChemsCADENA)

			localctx.(*PrmdiskContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*PrmdiskContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrmdiskContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrmdiskContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrmdiskContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValPath = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(192)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(193)

			var _m = p.Match(ChemsRUTA)

			localctx.(*PrmdiskContext)._RUTA = _m
		}
		ValPath = (func() string {
			if localctx.(*PrmdiskContext).Get_RUTA() == nil {
				return ""
			} else {
				return localctx.(*PrmdiskContext).Get_RUTA().GetText()
			}
		}())

	}

	return localctx
}

// IPmkgrpContext is an interface to support dynamic dispatch.
type IPmkgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// IsPmkgrpContext differentiates from other interfaces.
	IsPmkgrpContext()
}

type PmkgrpContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_CADENA        antlr.Token
	_IDENTIFICADOR antlr.Token
}

func NewEmptyPmkgrpContext() *PmkgrpContext {
	var p = new(PmkgrpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmkgrp
	return p
}

func (*PmkgrpContext) IsPmkgrpContext() {}

func NewPmkgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmkgrpContext {
	var p = new(PmkgrpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmkgrp

	return p
}

func (s *PmkgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *PmkgrpContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PmkgrpContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *PmkgrpContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PmkgrpContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *PmkgrpContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *PmkgrpContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *PmkgrpContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *PmkgrpContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *PmkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmkgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmkgrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmkgrp(s)
	}
}

func (s *PmkgrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmkgrp(s)
	}
}

func (p *Chems) Pmkgrp() (localctx IPmkgrpContext) {
	localctx = NewPmkgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChemsRULE_pmkgrp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(198)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(199)

			var _m = p.Match(ChemsCADENA)

			localctx.(*PmkgrpContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*PmkgrpContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PmkgrpContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PmkgrpContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PmkgrpContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValName = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(202)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(203)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*PmkgrpContext)._IDENTIFICADOR = _m
		}
		ValName = (func() string {
			if localctx.(*PmkgrpContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*PmkgrpContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	}

	return localctx
}

// IPrmkgrpContext is an interface to support dynamic dispatch.
type IPrmkgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// IsPrmkgrpContext differentiates from other interfaces.
	IsPrmkgrpContext()
}

type PrmkgrpContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_CADENA        antlr.Token
	_IDENTIFICADOR antlr.Token
}

func NewEmptyPrmkgrpContext() *PrmkgrpContext {
	var p = new(PrmkgrpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_prmkgrp
	return p
}

func (*PrmkgrpContext) IsPrmkgrpContext() {}

func NewPrmkgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrmkgrpContext {
	var p = new(PrmkgrpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_prmkgrp

	return p
}

func (s *PrmkgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrmkgrpContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrmkgrpContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *PrmkgrpContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrmkgrpContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *PrmkgrpContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *PrmkgrpContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *PrmkgrpContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *PrmkgrpContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *PrmkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrmkgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrmkgrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrmkgrp(s)
	}
}

func (s *PrmkgrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrmkgrp(s)
	}
}

func (p *Chems) Prmkgrp() (localctx IPrmkgrpContext) {
	localctx = NewPrmkgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChemsRULE_prmkgrp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(208)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(209)

			var _m = p.Match(ChemsCADENA)

			localctx.(*PrmkgrpContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*PrmkgrpContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrmkgrpContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrmkgrpContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrmkgrpContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValName = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(212)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(213)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*PrmkgrpContext)._IDENTIFICADOR = _m
		}
		ValName = (func() string {
			if localctx.(*PrmkgrpContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*PrmkgrpContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	}

	return localctx
}

// IPrmkusrContext is an interface to support dynamic dispatch.
type IPrmkusrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// IsPrmkusrContext differentiates from other interfaces.
	IsPrmkusrContext()
}

type PrmkusrContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_CADENA        antlr.Token
	_IDENTIFICADOR antlr.Token
}

func NewEmptyPrmkusrContext() *PrmkusrContext {
	var p = new(PrmkusrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_prmkusr
	return p
}

func (*PrmkusrContext) IsPrmkusrContext() {}

func NewPrmkusrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrmkusrContext {
	var p = new(PrmkusrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_prmkusr

	return p
}

func (s *PrmkusrContext) GetParser() antlr.Parser { return s.parser }

func (s *PrmkusrContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrmkusrContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *PrmkusrContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrmkusrContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *PrmkusrContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *PrmkusrContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *PrmkusrContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *PrmkusrContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *PrmkusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrmkusrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrmkusrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrmkusr(s)
	}
}

func (s *PrmkusrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrmkusr(s)
	}
}

func (p *Chems) Prmkusr() (localctx IPrmkusrContext) {
	localctx = NewPrmkusrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ChemsRULE_prmkusr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(218)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(219)

			var _m = p.Match(ChemsCADENA)

			localctx.(*PrmkusrContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*PrmkusrContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrmkusrContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrmkusrContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrmkusrContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValName = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(222)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(223)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*PrmkusrContext)._IDENTIFICADOR = _m
		}
		ValName = (func() string {
			if localctx.(*PrmkusrContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*PrmkusrContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	}

	return localctx
}

// IPfdiskContext is an interface to support dynamic dispatch.
type IPfdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPfdiskContext differentiates from other interfaces.
	IsPfdiskContext()
}

type PfdiskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPfdiskContext() *PfdiskContext {
	var p = new(PfdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pfdisk
	return p
}

func (*PfdiskContext) IsPfdiskContext() {}

func NewPfdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PfdiskContext {
	var p = new(PfdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pfdisk

	return p
}

func (s *PfdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *PfdiskContext) Parametrosfdisk() IParametrosfdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosfdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosfdiskContext)
}

func (s *PfdiskContext) Pfdisk() IPfdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPfdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPfdiskContext)
}

func (s *PfdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PfdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PfdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPfdisk(s)
	}
}

func (s *PfdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPfdisk(s)
	}
}

func (p *Chems) Pfdisk() (localctx IPfdiskContext) {
	return p.pfdisk(0)
}

func (p *Chems) pfdisk(_p int) (localctx IPfdiskContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPfdiskContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPfdiskContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ChemsRULE_pfdisk, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Parametrosfdisk()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPfdiskContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pfdisk)
			p.SetState(230)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(231)
				p.Parametrosfdisk()
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosfdiskContext is an interface to support dynamic dispatch.
type IParametrosfdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LETRA returns the _LETRA token.
	Get_LETRA() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Set_LETRA sets the _LETRA token.
	Set_LETRA(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// IsParametrosfdiskContext differentiates from other interfaces.
	IsParametrosfdiskContext()
}

type ParametrosfdiskContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_LETRA         antlr.Token
	_CADENA        antlr.Token
	_IDENTIFICADOR antlr.Token
}

func NewEmptyParametrosfdiskContext() *ParametrosfdiskContext {
	var p = new(ParametrosfdiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosfdisk
	return p
}

func (*ParametrosfdiskContext) IsParametrosfdiskContext() {}

func NewParametrosfdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosfdiskContext {
	var p = new(ParametrosfdiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosfdisk

	return p
}

func (s *ParametrosfdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosfdiskContext) Get_LETRA() antlr.Token { return s._LETRA }

func (s *ParametrosfdiskContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosfdiskContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *ParametrosfdiskContext) Set_LETRA(v antlr.Token) { s._LETRA = v }

func (s *ParametrosfdiskContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosfdiskContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *ParametrosfdiskContext) Parametrosmkdisk() IParametrosmkdiskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkdiskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkdiskContext)
}

func (s *ParametrosfdiskContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ChemsTYPE, 0)
}

func (s *ParametrosfdiskContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosfdiskContext) LETRA() antlr.TerminalNode {
	return s.GetToken(ChemsLETRA, 0)
}

func (s *ParametrosfdiskContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *ParametrosfdiskContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosfdiskContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *ParametrosfdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosfdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosfdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosfdisk(s)
	}
}

func (s *ParametrosfdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosfdisk(s)
	}
}

func (p *Chems) Parametrosfdisk() (localctx IParametrosfdiskContext) {
	localctx = NewParametrosfdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChemsRULE_parametrosfdisk)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Parametrosmkdisk()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Match(ChemsTYPE)
		}
		{
			p.SetState(239)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(240)

			var _m = p.Match(ChemsLETRA)

			localctx.(*ParametrosfdiskContext)._LETRA = _m
		}
		ValType = (func() string {
			if localctx.(*ParametrosfdiskContext).Get_LETRA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosfdiskContext).Get_LETRA().GetText()
			}
		}())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(243)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(244)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosfdiskContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosfdiskContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosfdiskContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosfdiskContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosfdiskContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValName = str

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(246)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(247)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(248)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosfdiskContext)._IDENTIFICADOR = _m
		}
		ValName = (func() string {
			if localctx.(*ParametrosfdiskContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosfdiskContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	}

	return localctx
}

// IPmountContext is an interface to support dynamic dispatch.
type IPmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmountContext differentiates from other interfaces.
	IsPmountContext()
}

type PmountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmountContext() *PmountContext {
	var p = new(PmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmount
	return p
}

func (*PmountContext) IsPmountContext() {}

func NewPmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmountContext {
	var p = new(PmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmount

	return p
}

func (s *PmountContext) GetParser() antlr.Parser { return s.parser }

func (s *PmountContext) Parametrosmount() IParametrosmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmountContext)
}

func (s *PmountContext) Pmount() IPmountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmountContext)
}

func (s *PmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmount(s)
	}
}

func (s *PmountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmount(s)
	}
}

func (p *Chems) Pmount() (localctx IPmountContext) {
	return p.pmount(0)
}

func (p *Chems) pmount(_p int) (localctx IPmountContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPmountContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPmountContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, ChemsRULE_pmount, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Parametrosmount()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPmountContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pmount)
			p.SetState(255)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(256)
				p.Parametrosmount()
			}

		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosmountContext is an interface to support dynamic dispatch.
type IParametrosmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_RUTA returns the _RUTA token.
	Get_RUTA() antlr.Token

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_RUTA sets the _RUTA token.
	Set_RUTA(antlr.Token)

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// IsParametrosmountContext differentiates from other interfaces.
	IsParametrosmountContext()
}

type ParametrosmountContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_CADENA        antlr.Token
	_RUTA          antlr.Token
	_IDENTIFICADOR antlr.Token
}

func NewEmptyParametrosmountContext() *ParametrosmountContext {
	var p = new(ParametrosmountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmount
	return p
}

func (*ParametrosmountContext) IsParametrosmountContext() {}

func NewParametrosmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmountContext {
	var p = new(ParametrosmountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmount

	return p
}

func (s *ParametrosmountContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmountContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosmountContext) Get_RUTA() antlr.Token { return s._RUTA }

func (s *ParametrosmountContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *ParametrosmountContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosmountContext) Set_RUTA(v antlr.Token) { s._RUTA = v }

func (s *ParametrosmountContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *ParametrosmountContext) PATH() antlr.TerminalNode {
	return s.GetToken(ChemsPATH, 0)
}

func (s *ParametrosmountContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmountContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosmountContext) RUTA() antlr.TerminalNode {
	return s.GetToken(ChemsRUTA, 0)
}

func (s *ParametrosmountContext) NAME() antlr.TerminalNode {
	return s.GetToken(ChemsNAME, 0)
}

func (s *ParametrosmountContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *ParametrosmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmount(s)
	}
}

func (s *ParametrosmountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmount(s)
	}
}

func (p *Chems) Parametrosmount() (localctx IParametrosmountContext) {
	localctx = NewParametrosmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChemsRULE_parametrosmount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(263)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(264)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmountContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmountContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmountContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmountContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmountContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValPath = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(ChemsPATH)
		}
		{
			p.SetState(267)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(268)

			var _m = p.Match(ChemsRUTA)

			localctx.(*ParametrosmountContext)._RUTA = _m
		}
		ValPath = (func() string {
			if localctx.(*ParametrosmountContext).Get_RUTA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmountContext).Get_RUTA().GetText()
			}
		}())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.Match(ChemsNAME)
		}
		{
			p.SetState(271)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(272)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosmountContext)._IDENTIFICADOR = _m
		}
		ValName = (func() string {
			if localctx.(*ParametrosmountContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmountContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	}

	return localctx
}

// IPmkfsContext is an interface to support dynamic dispatch.
type IPmkfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmkfsContext differentiates from other interfaces.
	IsPmkfsContext()
}

type PmkfsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmkfsContext() *PmkfsContext {
	var p = new(PmkfsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmkfs
	return p
}

func (*PmkfsContext) IsPmkfsContext() {}

func NewPmkfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmkfsContext {
	var p = new(PmkfsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmkfs

	return p
}

func (s *PmkfsContext) GetParser() antlr.Parser { return s.parser }

func (s *PmkfsContext) Parametrosmkfs() IParametrosmkfsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkfsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkfsContext)
}

func (s *PmkfsContext) Pmkfs() IPmkfsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmkfsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmkfsContext)
}

func (s *PmkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmkfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmkfsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmkfs(s)
	}
}

func (s *PmkfsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmkfs(s)
	}
}

func (p *Chems) Pmkfs() (localctx IPmkfsContext) {
	return p.pmkfs(0)
}

func (p *Chems) pmkfs(_p int) (localctx IPmkfsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPmkfsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPmkfsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, ChemsRULE_pmkfs, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Parametrosmkfs()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPmkfsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_pmkfs)
			p.SetState(279)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(280)
				p.Parametrosmkfs()
			}

		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosmkfsContext is an interface to support dynamic dispatch.
type IParametrosmkfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Get_LETRA returns the _LETRA token.
	Get_LETRA() antlr.Token

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// Set_LETRA sets the _LETRA token.
	Set_LETRA(antlr.Token)

	// IsParametrosmkfsContext differentiates from other interfaces.
	IsParametrosmkfsContext()
}

type ParametrosmkfsContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_IDENTIFICADOR antlr.Token
	_LETRA         antlr.Token
}

func NewEmptyParametrosmkfsContext() *ParametrosmkfsContext {
	var p = new(ParametrosmkfsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmkfs
	return p
}

func (*ParametrosmkfsContext) IsParametrosmkfsContext() {}

func NewParametrosmkfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmkfsContext {
	var p = new(ParametrosmkfsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmkfs

	return p
}

func (s *ParametrosmkfsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmkfsContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *ParametrosmkfsContext) Get_LETRA() antlr.Token { return s._LETRA }

func (s *ParametrosmkfsContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *ParametrosmkfsContext) Set_LETRA(v antlr.Token) { s._LETRA = v }

func (s *ParametrosmkfsContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *ParametrosmkfsContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmkfsContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *ParametrosmkfsContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ChemsTYPE, 0)
}

func (s *ParametrosmkfsContext) LETRA() antlr.TerminalNode {
	return s.GetToken(ChemsLETRA, 0)
}

func (s *ParametrosmkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmkfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmkfsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmkfs(s)
	}
}

func (s *ParametrosmkfsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmkfs(s)
	}
}

func (p *Chems) Parametrosmkfs() (localctx IParametrosmkfsContext) {
	localctx = NewParametrosmkfsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ChemsRULE_parametrosmkfs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(294)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Match(ChemsID)
		}
		{
			p.SetState(287)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(288)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosmkfsContext)._IDENTIFICADOR = _m
		}
		ValIdentificador = (func() string {
			if localctx.(*ParametrosmkfsContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfsContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	case ChemsTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)
			p.Match(ChemsTYPE)
		}
		{
			p.SetState(291)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(292)

			var _m = p.Match(ChemsLETRA)

			localctx.(*ParametrosmkfsContext)._LETRA = _m
		}
		ValType = (func() string {
			if localctx.(*ParametrosmkfsContext).Get_LETRA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkfsContext).Get_LETRA().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPloginContext is an interface to support dynamic dispatch.
type IPloginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPloginContext differentiates from other interfaces.
	IsPloginContext()
}

type PloginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPloginContext() *PloginContext {
	var p = new(PloginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_plogin
	return p
}

func (*PloginContext) IsPloginContext() {}

func NewPloginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PloginContext {
	var p = new(PloginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_plogin

	return p
}

func (s *PloginContext) GetParser() antlr.Parser { return s.parser }

func (s *PloginContext) Parametroslogin() IParametrosloginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosloginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosloginContext)
}

func (s *PloginContext) Plogin() IPloginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPloginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPloginContext)
}

func (s *PloginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PloginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PloginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPlogin(s)
	}
}

func (s *PloginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPlogin(s)
	}
}

func (p *Chems) Plogin() (localctx IPloginContext) {
	return p.plogin(0)
}

func (p *Chems) plogin(_p int) (localctx IPloginContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPloginContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPloginContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, ChemsRULE_plogin, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Parametroslogin()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPloginContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_plogin)
			p.SetState(299)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(300)
				p.Parametroslogin()
			}

		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IParametrosloginContext is an interface to support dynamic dispatch.
type IParametrosloginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// IsParametrosloginContext differentiates from other interfaces.
	IsParametrosloginContext()
}

type ParametrosloginContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_IDENTIFICADOR antlr.Token
	_CADENA        antlr.Token
}

func NewEmptyParametrosloginContext() *ParametrosloginContext {
	var p = new(ParametrosloginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametroslogin
	return p
}

func (*ParametrosloginContext) IsParametrosloginContext() {}

func NewParametrosloginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosloginContext {
	var p = new(ParametrosloginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametroslogin

	return p
}

func (s *ParametrosloginContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosloginContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *ParametrosloginContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosloginContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *ParametrosloginContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosloginContext) USR() antlr.TerminalNode {
	return s.GetToken(ChemsUSR, 0)
}

func (s *ParametrosloginContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosloginContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *ParametrosloginContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosloginContext) PWD() antlr.TerminalNode {
	return s.GetToken(ChemsPWD, 0)
}

func (s *ParametrosloginContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *ParametrosloginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosloginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosloginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametroslogin(s)
	}
}

func (s *ParametrosloginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametroslogin(s)
	}
}

func (p *Chems) Parametroslogin() (localctx IParametrosloginContext) {
	localctx = NewParametrosloginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ChemsRULE_parametroslogin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.Match(ChemsUSR)
		}
		{
			p.SetState(307)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(308)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosloginContext)._IDENTIFICADOR = _m
		}
		ValUsuario = (func() string {
			if localctx.(*ParametrosloginContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Match(ChemsUSR)
		}
		{
			p.SetState(311)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(312)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosloginContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosloginContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosloginContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValUsuario = str

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(314)
			p.Match(ChemsPWD)
		}
		{
			p.SetState(315)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(316)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosloginContext)._IDENTIFICADOR = _m
		}
		ValPassword = (func() string {
			if localctx.(*ParametrosloginContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(318)
			p.Match(ChemsPWD)
		}
		{
			p.SetState(319)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(320)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosloginContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosloginContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosloginContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValPassword = str

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(322)
			p.Match(ChemsID)
		}
		{
			p.SetState(323)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(324)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosloginContext)._IDENTIFICADOR = _m
		}
		ValIdentificador = (func() string {
			if localctx.(*ParametrosloginContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosloginContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	}

	return localctx
}

// IPmkuserContext is an interface to support dynamic dispatch.
type IPmkuserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmkuserContext differentiates from other interfaces.
	IsPmkuserContext()
}

type PmkuserContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmkuserContext() *PmkuserContext {
	var p = new(PmkuserContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_pmkuser
	return p
}

func (*PmkuserContext) IsPmkuserContext() {}

func NewPmkuserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmkuserContext {
	var p = new(PmkuserContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_pmkuser

	return p
}

func (s *PmkuserContext) GetParser() antlr.Parser { return s.parser }

func (s *PmkuserContext) Parametrosmkuser() IParametrosmkuserContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametrosmkuserContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametrosmkuserContext)
}

func (s *PmkuserContext) Plogin() IPloginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPloginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPloginContext)
}

func (s *PmkuserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmkuserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmkuserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPmkuser(s)
	}
}

func (s *PmkuserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPmkuser(s)
	}
}

func (p *Chems) Pmkuser() (localctx IPmkuserContext) {
	localctx = NewPmkuserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ChemsRULE_pmkuser)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Parametrosmkuser()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.plogin(0)
		}
		{
			p.SetState(330)
			p.Parametrosmkuser()
		}

	}

	return localctx
}

// IParametrosmkuserContext is an interface to support dynamic dispatch.
type IParametrosmkuserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFICADOR returns the _IDENTIFICADOR token.
	Get_IDENTIFICADOR() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Set_IDENTIFICADOR sets the _IDENTIFICADOR token.
	Set_IDENTIFICADOR(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// IsParametrosmkuserContext differentiates from other interfaces.
	IsParametrosmkuserContext()
}

type ParametrosmkuserContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_IDENTIFICADOR antlr.Token
	_CADENA        antlr.Token
}

func NewEmptyParametrosmkuserContext() *ParametrosmkuserContext {
	var p = new(ParametrosmkuserContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_parametrosmkuser
	return p
}

func (*ParametrosmkuserContext) IsParametrosmkuserContext() {}

func NewParametrosmkuserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametrosmkuserContext {
	var p = new(ParametrosmkuserContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_parametrosmkuser

	return p
}

func (s *ParametrosmkuserContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametrosmkuserContext) Get_IDENTIFICADOR() antlr.Token { return s._IDENTIFICADOR }

func (s *ParametrosmkuserContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ParametrosmkuserContext) Set_IDENTIFICADOR(v antlr.Token) { s._IDENTIFICADOR = v }

func (s *ParametrosmkuserContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ParametrosmkuserContext) USR() antlr.TerminalNode {
	return s.GetToken(ChemsUSR, 0)
}

func (s *ParametrosmkuserContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *ParametrosmkuserContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(ChemsIDENTIFICADOR, 0)
}

func (s *ParametrosmkuserContext) CADENA() antlr.TerminalNode {
	return s.GetToken(ChemsCADENA, 0)
}

func (s *ParametrosmkuserContext) PWD1() antlr.TerminalNode {
	return s.GetToken(ChemsPWD1, 0)
}

func (s *ParametrosmkuserContext) GRP() antlr.TerminalNode {
	return s.GetToken(ChemsGRP, 0)
}

func (s *ParametrosmkuserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametrosmkuserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametrosmkuserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterParametrosmkuser(s)
	}
}

func (s *ParametrosmkuserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitParametrosmkuser(s)
	}
}

func (p *Chems) Parametrosmkuser() (localctx IParametrosmkuserContext) {
	localctx = NewParametrosmkuserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ChemsRULE_parametrosmkuser)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.Match(ChemsUSR)
		}
		{
			p.SetState(335)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(336)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosmkuserContext)._IDENTIFICADOR = _m
		}
		ValUsuario = (func() string {
			if localctx.(*ParametrosmkuserContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(338)
			p.Match(ChemsUSR)
		}
		{
			p.SetState(339)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(340)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkuserContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmkuserContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmkuserContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValUsuario = str

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(342)
			p.Match(ChemsPWD1)
		}
		{
			p.SetState(343)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(344)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosmkuserContext)._IDENTIFICADOR = _m
		}
		ValPassword = (func() string {
			if localctx.(*ParametrosmkuserContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(346)
			p.Match(ChemsPWD1)
		}
		{
			p.SetState(347)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(348)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkuserContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmkuserContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmkuserContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValPassword = str

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(350)
			p.Match(ChemsGRP)
		}
		{
			p.SetState(351)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(352)

			var _m = p.Match(ChemsIDENTIFICADOR)

			localctx.(*ParametrosmkuserContext)._IDENTIFICADOR = _m
		}
		ValGrupo = (func() string {
			if localctx.(*ParametrosmkuserContext).Get_IDENTIFICADOR() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_IDENTIFICADOR().GetText()
			}
		}())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(354)
			p.Match(ChemsGRP)
		}
		{
			p.SetState(355)
			p.Match(ChemsIGUAL)
		}
		{
			p.SetState(356)

			var _m = p.Match(ChemsCADENA)

			localctx.(*ParametrosmkuserContext)._CADENA = _m
		}
		str := (func() string {
			if localctx.(*ParametrosmkuserContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ParametrosmkuserContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ParametrosmkuserContext).Get_CADENA().GetText()
			}
		}()))-1]
		ValGrupo = str

	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *PmkfileContext = nil
		if localctx != nil {
			t = localctx.(*PmkfileContext)
		}
		return p.Pmkfile_Sempred(t, predIndex)

	case 7:
		var t *PmkdiskContext = nil
		if localctx != nil {
			t = localctx.(*PmkdiskContext)
		}
		return p.Pmkdisk_Sempred(t, predIndex)

	case 13:
		var t *PfdiskContext = nil
		if localctx != nil {
			t = localctx.(*PfdiskContext)
		}
		return p.Pfdisk_Sempred(t, predIndex)

	case 15:
		var t *PmountContext = nil
		if localctx != nil {
			t = localctx.(*PmountContext)
		}
		return p.Pmount_Sempred(t, predIndex)

	case 17:
		var t *PmkfsContext = nil
		if localctx != nil {
			t = localctx.(*PmkfsContext)
		}
		return p.Pmkfs_Sempred(t, predIndex)

	case 19:
		var t *PloginContext = nil
		if localctx != nil {
			t = localctx.(*PloginContext)
		}
		return p.Plogin_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Pmkfile_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Pmkdisk_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Pfdisk_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Pmount_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Pmkfs_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Plogin_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
