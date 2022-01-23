import codecs
from operator import itemgetter
from collections import Counter


ENGLISH_MOST_FREQUENT_LETTERS = ' etaoinshrdlu'


def try_all_key_combinations(bytes_input):
    bytes_english_most_frequent_letters = {ord(c) for c in ENGLISH_MOST_FREQUENT_LETTERS}

    for key in range(255):
        message = bytes(a ^ key for a in bytes_input)
        counter = Counter(message)
        result = sum(frequent for character, frequent in counter.most_common() if character in bytes_english_most_frequent_letters)

        yield result, key, message


def crack_single_xor_encryption(hex_input):
    bytes_input = codecs.decode(hex_input, 'hex_codec')
    decoding_results = try_all_key_combinations(bytes_input)
    best_candidate = max(decoding_results, key=itemgetter(0))

    try:
        return chr(best_candidate[1]), codecs.decode(best_candidate[2], 'ascii')
    except:
        raise Exception("message cannot be read as ASCII")


character, message = crack_single_xor_encryption('1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736')
print(f'key: "{character}" message: "{message}"')
