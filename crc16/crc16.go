package crc16

const CRC16_POLY = 0x1021

func Compute(data []byte) uint16 {
	crc := uint16(0x0000)

	for _, b := range data {
		crc ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ CRC16_POLY
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}
