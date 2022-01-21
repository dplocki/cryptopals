import codecs
import binascii

def to_base64(hex_input: str) -> str:
    bytes_input = codecs.decode(hex_input, 'hex_codec')
    result = binascii.b2a_base64(bytes_input, newline=False).decode('utf-8')
    return result

assert to_base64('49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d') == 'SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t'
