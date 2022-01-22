import codecs

def fixed_xor(hex_input, hex_key):
    bytes_input = codecs.decode(hex_input, 'hex_codec')
    bytes_key = codecs.decode(hex_key, 'hex_codec')

    result = bytes(a ^ b for a, b in zip(bytes_input, bytes_key))

    return result.hex()

assert fixed_xor('1c0111001f010100061a024b53535009181c', '686974207468652062756c6c277320657965') == '746865206b696420646f6e277420706c6179'
