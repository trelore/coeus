package schemes

import "image/color"

// PBJ is a gradient color scheme from orange to purple.
var PBJ []color.Color

func init() {
	PBJ = []color.Color{
		color.RGBA{R: 0x29, G: 0xa, B: 0x59, A: 0xff},
		color.RGBA{R: 0x29, G: 0xa, B: 0x59, A: 0xff},
		color.RGBA{R: 0x2a, G: 0xa, B: 0x59, A: 0xff},
		color.RGBA{R: 0x2a, G: 0xa, B: 0x59, A: 0xff},
		color.RGBA{R: 0x2a, G: 0xa, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2b, G: 0xa, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2b, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2b, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2c, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2c, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2d, G: 0xa, B: 0x59, A: 0xff},
		color.RGBA{R: 0x2e, G: 0xa, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2e, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2f, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2f, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x2f, G: 0x9, B: 0x58, A: 0xff},
		color.RGBA{R: 0x30, G: 0x8, B: 0x58, A: 0xff},
		color.RGBA{R: 0x30, G: 0x8, B: 0x57, A: 0xff},
		color.RGBA{R: 0x31, G: 0x8, B: 0x57, A: 0xff},
		color.RGBA{R: 0x31, G: 0x8, B: 0x57, A: 0xff},
		color.RGBA{R: 0x31, G: 0x7, B: 0x57, A: 0xff},
		color.RGBA{R: 0x32, G: 0x7, B: 0x57, A: 0xff},
		color.RGBA{R: 0x32, G: 0x7, B: 0x57, A: 0xff},
		color.RGBA{R: 0x33, G: 0x7, B: 0x56, A: 0xff},
		color.RGBA{R: 0x33, G: 0x6, B: 0x56, A: 0xff},
		color.RGBA{R: 0x35, G: 0x7, B: 0x56, A: 0xff},
		color.RGBA{R: 0x35, G: 0x7, B: 0x56, A: 0xff},
		color.RGBA{R: 0x36, G: 0x7, B: 0x56, A: 0xff},
		color.RGBA{R: 0x36, G: 0x6, B: 0x55, A: 0xff},
		color.RGBA{R: 0x37, G: 0x6, B: 0x55, A: 0xff},
		color.RGBA{R: 0x37, G: 0x6, B: 0x55, A: 0xff},
		color.RGBA{R: 0x38, G: 0x5, B: 0x55, A: 0xff},
		color.RGBA{R: 0x38, G: 0x5, B: 0x55, A: 0xff},
		color.RGBA{R: 0x39, G: 0x5, B: 0x54, A: 0xff},
		color.RGBA{R: 0x39, G: 0x5, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3a, G: 0x4, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3b, G: 0x4, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3b, G: 0x5, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3c, G: 0x4, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3c, G: 0x4, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3d, G: 0x4, B: 0x54, A: 0xff},
		color.RGBA{R: 0x3d, G: 0x4, B: 0x53, A: 0xff},
		color.RGBA{R: 0x3e, G: 0x3, B: 0x53, A: 0xff},
		color.RGBA{R: 0x3f, G: 0x3, B: 0x53, A: 0xff},
		color.RGBA{R: 0x3f, G: 0x3, B: 0x53, A: 0xff},
		color.RGBA{R: 0x40, G: 0x3, B: 0x52, A: 0xff},
		color.RGBA{R: 0x40, G: 0x3, B: 0x52, A: 0xff},
		color.RGBA{R: 0x41, G: 0x3, B: 0x52, A: 0xff},
		color.RGBA{R: 0x42, G: 0x3, B: 0x52, A: 0xff},
		color.RGBA{R: 0x43, G: 0x4, B: 0x52, A: 0xff},
		color.RGBA{R: 0x44, G: 0x4, B: 0x52, A: 0xff},
		color.RGBA{R: 0x45, G: 0x4, B: 0x52, A: 0xff},
		color.RGBA{R: 0x45, G: 0x4, B: 0x51, A: 0xff},
		color.RGBA{R: 0x46, G: 0x4, B: 0x51, A: 0xff},
		color.RGBA{R: 0x47, G: 0x4, B: 0x51, A: 0xff},
		color.RGBA{R: 0x47, G: 0x4, B: 0x50, A: 0xff},
		color.RGBA{R: 0x48, G: 0x4, B: 0x50, A: 0xff},
		color.RGBA{R: 0x49, G: 0x4, B: 0x50, A: 0xff},
		color.RGBA{R: 0x49, G: 0x4, B: 0x4f, A: 0xff},
		color.RGBA{R: 0x4b, G: 0x5, B: 0x50, A: 0xff},
		color.RGBA{R: 0x4c, G: 0x5, B: 0x50, A: 0xff},
		color.RGBA{R: 0x4d, G: 0x5, B: 0x4f, A: 0xff},
		color.RGBA{R: 0x4d, G: 0x5, B: 0x4f, A: 0xff},
		color.RGBA{R: 0x4e, G: 0x5, B: 0x4f, A: 0xff},
		color.RGBA{R: 0x4f, G: 0x5, B: 0x4e, A: 0xff},
		color.RGBA{R: 0x50, G: 0x5, B: 0x4e, A: 0xff},
		color.RGBA{R: 0x50, G: 0x5, B: 0x4e, A: 0xff},
		color.RGBA{R: 0x50, G: 0x5, B: 0x4d, A: 0xff},
		color.RGBA{R: 0x51, G: 0x5, B: 0x4d, A: 0xff},
		color.RGBA{R: 0x53, G: 0x6, B: 0x4c, A: 0xff},
		color.RGBA{R: 0x53, G: 0x6, B: 0x4c, A: 0xff},
		color.RGBA{R: 0x54, G: 0x6, B: 0x4c, A: 0xff},
		color.RGBA{R: 0x55, G: 0x6, B: 0x4b, A: 0xff},
		color.RGBA{R: 0x56, G: 0x6, B: 0x4b, A: 0xff},
		color.RGBA{R: 0x57, G: 0x6, B: 0x4a, A: 0xff},
		color.RGBA{R: 0x58, G: 0x6, B: 0x4a, A: 0xff},
		color.RGBA{R: 0x58, G: 0x6, B: 0x49, A: 0xff},
		color.RGBA{R: 0x59, G: 0x6, B: 0x49, A: 0xff},
		color.RGBA{R: 0x5b, G: 0x7, B: 0x49, A: 0xff},
		color.RGBA{R: 0x5c, G: 0x7, B: 0x49, A: 0xff},
		color.RGBA{R: 0x5d, G: 0x7, B: 0x48, A: 0xff},
		color.RGBA{R: 0x5e, G: 0x7, B: 0x48, A: 0xff},
		color.RGBA{R: 0x5e, G: 0x7, B: 0x47, A: 0xff},
		color.RGBA{R: 0x5f, G: 0x7, B: 0x47, A: 0xff},
		color.RGBA{R: 0x60, G: 0x7, B: 0x46, A: 0xff},
		color.RGBA{R: 0x60, G: 0x7, B: 0x46, A: 0xff},
		color.RGBA{R: 0x61, G: 0x7, B: 0x45, A: 0xff},
		color.RGBA{R: 0x63, G: 0x9, B: 0x46, A: 0xff},
		color.RGBA{R: 0x64, G: 0x9, B: 0x45, A: 0xff},
		color.RGBA{R: 0x65, G: 0xa, B: 0x45, A: 0xff},
		color.RGBA{R: 0x66, G: 0xa, B: 0x44, A: 0xff},
		color.RGBA{R: 0x67, G: 0xb, B: 0x43, A: 0xff},
		color.RGBA{R: 0x68, G: 0xb, B: 0x43, A: 0xff},
		color.RGBA{R: 0x69, G: 0xc, B: 0x42, A: 0xff},
		color.RGBA{R: 0x6a, G: 0xd, B: 0x42, A: 0xff},
		color.RGBA{R: 0x6b, G: 0xe, B: 0x42, A: 0xff},
		color.RGBA{R: 0x6c, G: 0xf, B: 0x41, A: 0xff},
		color.RGBA{R: 0x6d, G: 0x10, B: 0x40, A: 0xff},
		color.RGBA{R: 0x6e, G: 0x10, B: 0x40, A: 0xff},
		color.RGBA{R: 0x6f, G: 0x11, B: 0x3f, A: 0xff},
		color.RGBA{R: 0x70, G: 0x12, B: 0x3e, A: 0xff},
		color.RGBA{R: 0x71, G: 0x12, B: 0x3d, A: 0xff},
		color.RGBA{R: 0x72, G: 0x13, B: 0x3d, A: 0xff},
		color.RGBA{R: 0x73, G: 0x14, B: 0x3c, A: 0xff},
		color.RGBA{R: 0x76, G: 0x16, B: 0x3c, A: 0xff},
		color.RGBA{R: 0x77, G: 0x16, B: 0x3b, A: 0xff},
		color.RGBA{R: 0x78, G: 0x16, B: 0x3a, A: 0xff},
		color.RGBA{R: 0x78, G: 0x17, B: 0x3a, A: 0xff},
		color.RGBA{R: 0x79, G: 0x18, B: 0x39, A: 0xff},
		color.RGBA{R: 0x7a, G: 0x19, B: 0x38, A: 0xff},
		color.RGBA{R: 0x7c, G: 0x1a, B: 0x37, A: 0xff},
		color.RGBA{R: 0x7d, G: 0x1b, B: 0x36, A: 0xff},
		color.RGBA{R: 0x7f, G: 0x1d, B: 0x36, A: 0xff},
		color.RGBA{R: 0x80, G: 0x1e, B: 0x36, A: 0xff},
		color.RGBA{R: 0x82, G: 0x1f, B: 0x35, A: 0xff},
		color.RGBA{R: 0x83, G: 0x20, B: 0x34, A: 0xff},
		color.RGBA{R: 0x84, G: 0x21, B: 0x33, A: 0xff},
		color.RGBA{R: 0x85, G: 0x22, B: 0x32, A: 0xff},
		color.RGBA{R: 0x86, G: 0x23, B: 0x31, A: 0xff},
		color.RGBA{R: 0x87, G: 0x24, B: 0x30, A: 0xff},
		color.RGBA{R: 0x89, G: 0x26, B: 0x30, A: 0xff},
		color.RGBA{R: 0x8a, G: 0x27, B: 0x2f, A: 0xff},
		color.RGBA{R: 0x8c, G: 0x28, B: 0x2e, A: 0xff},
		color.RGBA{R: 0x8d, G: 0x29, B: 0x2e, A: 0xff},
		color.RGBA{R: 0x8e, G: 0x2a, B: 0x2d, A: 0xff},
		color.RGBA{R: 0x8f, G: 0x2a, B: 0x2c, A: 0xff},
		color.RGBA{R: 0x90, G: 0x2b, B: 0x2b, A: 0xff},
		color.RGBA{R: 0x91, G: 0x2c, B: 0x2a, A: 0xff},
		color.RGBA{R: 0x92, G: 0x2d, B: 0x2a, A: 0xff},
		color.RGBA{R: 0x95, G: 0x2f, B: 0x29, A: 0xff},
		color.RGBA{R: 0x96, G: 0x30, B: 0x29, A: 0xff},
		color.RGBA{R: 0x97, G: 0x31, B: 0x28, A: 0xff},
		color.RGBA{R: 0x98, G: 0x32, B: 0x27, A: 0xff},
		color.RGBA{R: 0x99, G: 0x33, B: 0x26, A: 0xff},
		color.RGBA{R: 0x9a, G: 0x34, B: 0x26, A: 0xff},
		color.RGBA{R: 0x9b, G: 0x35, B: 0x25, A: 0xff},
		color.RGBA{R: 0x9d, G: 0x37, B: 0x24, A: 0xff},
		color.RGBA{R: 0x9f, G: 0x39, B: 0x24, A: 0xff},
		color.RGBA{R: 0xa0, G: 0x39, B: 0x23, A: 0xff},
		color.RGBA{R: 0xa0, G: 0x3a, B: 0x22, A: 0xff},
		color.RGBA{R: 0xa2, G: 0x3b, B: 0x21, A: 0xff},
		color.RGBA{R: 0xa3, G: 0x3c, B: 0x21, A: 0xff},
		color.RGBA{R: 0xa4, G: 0x3d, B: 0x20, A: 0xff},
		color.RGBA{R: 0xa5, G: 0x3e, B: 0x1f, A: 0xff},
		color.RGBA{R: 0xa7, G: 0x3f, B: 0x1e, A: 0xff},
		color.RGBA{R: 0xa8, G: 0x41, B: 0x1e, A: 0xff},
		color.RGBA{R: 0xa9, G: 0x42, B: 0x1d, A: 0xff},
		color.RGBA{R: 0xaa, G: 0x43, B: 0x1d, A: 0xff},
		color.RGBA{R: 0xac, G: 0x44, B: 0x1c, A: 0xff},
		color.RGBA{R: 0xad, G: 0x45, B: 0x1b, A: 0xff},
		color.RGBA{R: 0xae, G: 0x46, B: 0x1a, A: 0xff},
		color.RGBA{R: 0xaf, G: 0x47, B: 0x1a, A: 0xff},
		color.RGBA{R: 0xb0, G: 0x47, B: 0x19, A: 0xff},
		color.RGBA{R: 0xb2, G: 0x49, B: 0x19, A: 0xff},
		color.RGBA{R: 0xb3, G: 0x4a, B: 0x18, A: 0xff},
		color.RGBA{R: 0xb4, G: 0x4b, B: 0x18, A: 0xff},
		color.RGBA{R: 0xb5, G: 0x4c, B: 0x17, A: 0xff},
		color.RGBA{R: 0xb6, G: 0x4d, B: 0x17, A: 0xff},
		color.RGBA{R: 0xb7, G: 0x4e, B: 0x17, A: 0xff},
		color.RGBA{R: 0xb8, G: 0x4f, B: 0x16, A: 0xff},
		color.RGBA{R: 0xba, G: 0x50, B: 0x16, A: 0xff},
		color.RGBA{R: 0xbb, G: 0x51, B: 0x15, A: 0xff},
		color.RGBA{R: 0xbc, G: 0x52, B: 0x15, A: 0xff},
		color.RGBA{R: 0xbd, G: 0x53, B: 0x15, A: 0xff},
		color.RGBA{R: 0xbe, G: 0x53, B: 0x14, A: 0xff},
		color.RGBA{R: 0xbf, G: 0x54, B: 0x14, A: 0xff},
		color.RGBA{R: 0xc0, G: 0x55, B: 0x13, A: 0xff},
		color.RGBA{R: 0xc0, G: 0x56, B: 0x13, A: 0xff},
		color.RGBA{R: 0xc1, G: 0x57, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc2, G: 0x57, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc4, G: 0x59, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc4, G: 0x5a, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc5, G: 0x5a, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc6, G: 0x5a, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc7, G: 0x5b, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc8, G: 0x5c, B: 0x12, A: 0xff},
		color.RGBA{R: 0xc9, G: 0x5d, B: 0x12, A: 0xff},
		color.RGBA{R: 0xca, G: 0x5d, B: 0x12, A: 0xff},
		color.RGBA{R: 0xcb, G: 0x5e, B: 0x12, A: 0xff},
		color.RGBA{R: 0xcc, G: 0x60, B: 0x13, A: 0xff},
		color.RGBA{R: 0xcc, G: 0x60, B: 0x13, A: 0xff},
		color.RGBA{R: 0xcd, G: 0x61, B: 0x13, A: 0xff},
		color.RGBA{R: 0xce, G: 0x62, B: 0x13, A: 0xff},
		color.RGBA{R: 0xcf, G: 0x63, B: 0x13, A: 0xff},
		color.RGBA{R: 0xd0, G: 0x63, B: 0x13, A: 0xff},
		color.RGBA{R: 0xd1, G: 0x64, B: 0x13, A: 0xff},
		color.RGBA{R: 0xd2, G: 0x64, B: 0x13, A: 0xff},
		color.RGBA{R: 0xd3, G: 0x64, B: 0x13, A: 0xff},
		color.RGBA{R: 0xd4, G: 0x66, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd5, G: 0x67, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd6, G: 0x67, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd6, G: 0x68, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd7, G: 0x69, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd7, G: 0x69, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd8, G: 0x6a, B: 0x14, A: 0xff},
		color.RGBA{R: 0xd9, G: 0x6b, B: 0x14, A: 0xff},
		color.RGBA{R: 0xda, G: 0x6b, B: 0x14, A: 0xff},
		color.RGBA{R: 0xdb, G: 0x6c, B: 0x14, A: 0xff},
		color.RGBA{R: 0xdc, G: 0x6d, B: 0x15, A: 0xff},
		color.RGBA{R: 0xdd, G: 0x6d, B: 0x15, A: 0xff},
		color.RGBA{R: 0xde, G: 0x6e, B: 0x15, A: 0xff},
		color.RGBA{R: 0xde, G: 0x6f, B: 0x15, A: 0xff},
		color.RGBA{R: 0xdf, G: 0x6f, B: 0x15, A: 0xff},
		color.RGBA{R: 0xe0, G: 0x70, B: 0x15, A: 0xff},
		color.RGBA{R: 0xe1, G: 0x71, B: 0x15, A: 0xff},
		color.RGBA{R: 0xe2, G: 0x71, B: 0x15, A: 0xff},
		color.RGBA{R: 0xe3, G: 0x72, B: 0x15, A: 0xff},
		color.RGBA{R: 0xe3, G: 0x72, B: 0x15, A: 0xff},
		color.RGBA{R: 0xe4, G: 0x73, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe5, G: 0x74, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe5, G: 0x74, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe6, G: 0x75, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe7, G: 0x75, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe7, G: 0x76, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe8, G: 0x77, B: 0x16, A: 0xff},
		color.RGBA{R: 0xe9, G: 0x77, B: 0x16, A: 0xff},
		color.RGBA{R: 0xea, G: 0x78, B: 0x16, A: 0xff},
		color.RGBA{R: 0xea, G: 0x78, B: 0x16, A: 0xff},
		color.RGBA{R: 0xeb, G: 0x79, B: 0x16, A: 0xff},
		color.RGBA{R: 0xec, G: 0x79, B: 0x16, A: 0xff},
		color.RGBA{R: 0xed, G: 0x7a, B: 0x17, A: 0xff},
		color.RGBA{R: 0xed, G: 0x7a, B: 0x17, A: 0xff},
		color.RGBA{R: 0xee, G: 0x7b, B: 0x17, A: 0xff},
		color.RGBA{R: 0xef, G: 0x7c, B: 0x17, A: 0xff},
		color.RGBA{R: 0xef, G: 0x7c, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf0, G: 0x7d, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf0, G: 0x7d, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf1, G: 0x7e, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf1, G: 0x7e, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf2, G: 0x7f, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf3, G: 0x7f, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf3, G: 0x80, B: 0x17, A: 0xff},
		color.RGBA{R: 0xf4, G: 0x80, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf4, G: 0x80, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf5, G: 0x81, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf6, G: 0x81, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf6, G: 0x82, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf7, G: 0x82, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf7, G: 0x83, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf8, G: 0x83, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf9, G: 0x83, B: 0x18, A: 0xff},
		color.RGBA{R: 0xf9, G: 0x84, B: 0x18, A: 0xff},
		color.RGBA{R: 0xfa, G: 0x84, B: 0x18, A: 0xff},
		color.RGBA{R: 0xfa, G: 0x85, B: 0x18, A: 0xff},
		color.RGBA{R: 0xfa, G: 0x85, B: 0x18, A: 0xff},
		color.RGBA{R: 0xfa, G: 0x85, B: 0x18, A: 0xff},
		color.RGBA{R: 0xfb, G: 0x86, B: 0x18, A: 0xff},
		color.RGBA{R: 0xfb, G: 0x86, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfc, G: 0x87, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfc, G: 0x87, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfd, G: 0x87, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfd, G: 0x88, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfd, G: 0x88, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfe, G: 0x88, B: 0x19, A: 0xff},
		color.RGBA{R: 0xfe, G: 0x88, B: 0x19, A: 0xff},
		color.RGBA{R: 0xff, G: 0x89, B: 0x19, A: 0xff},
	}
}