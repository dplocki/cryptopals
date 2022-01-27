def count_set_bits(number):
    result = 0

    while number:
        result += number & 1
        number >>= 1

    return result


def hamming_distance(a_input, b_input):
    return sum(count_set_bits(ord(a) ^ ord(b)) for a, b in zip(a_input, b_input))


assert hamming_distance('this is a test', 'wokka wokka!!!') == 37
