import codecs
from operator import itemgetter
from collections import Counter


ENGLISH_MOST_FREQUENT_LETTERS = ' etaoinshrdlu'


def load_file(file_name):
    with open(file_name) as file:
        return file.read().strip()


def try_all_key_combinations(bytes_input):
    bytes_english_most_frequent_letters = {ord(c) for c in ENGLISH_MOST_FREQUENT_LETTERS}

    for key in range(255):
        message = bytes(a ^ key for a in bytes_input)
        counter = Counter(message)
        result = sum(frequent for character, frequent in counter.most_common() if character in bytes_english_most_frequent_letters)

        yield result, key, message


def try_all_lines(lines):
    for line in lines:
        bytes_input = codecs.decode(line, 'hex_codec')
        decoding_results = try_all_key_combinations(bytes_input)
        yield from decoding_results


# download file from https://cryptopals.com/static/challenge-data/4.txt
result = max(try_all_lines(load_file('set1/4.txt').splitlines()), key=itemgetter(0))
print(codecs.decode(result[2], 'ascii'))
