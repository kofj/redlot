package redlot

func encodeHashKey(name, key []byte) (buf []byte) {
	buf = append(buf, typeHASH)
	buf = append(buf, uint32ToBytes(uint32(len(name)))...)
	buf = append(buf, name...)
	buf = append(buf, uint32ToBytes(uint32(len(key)))...)
	buf = append(buf, key...)
	return
}

func decodeHashKey(b []byte) (name, key []byte) {
	nameLen := bytesToUint32(b[1:5])
	name = b[5 : 5+nameLen]
	key = b[9+nameLen:]
	return
}

func encodeHsizeKey(name []byte) (buf []byte) {
	buf = append(buf, typeHSIZE)
	buf = append(buf, uint32ToBytes(uint32(len(name)))...)
	buf = append(buf, name...)
	return
}

func decodeHsizeKey(b []byte) (key []byte) {
	return b[5:]
}
