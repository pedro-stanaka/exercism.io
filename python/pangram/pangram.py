def is_pangram(validate: str) -> bool:
    """
    Check whether a string is a pangran, i.e the string contains all the letters in the alphabet.
    :param validate: The string to be checked.
    :return: Whether the string is a pangram or not.
    """
    alphabet = {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u',
                'v', 'w', 'x', 'y', 'z'}
    str_list = set(validate.lower())

    return alphabet.issubset(str_list)
