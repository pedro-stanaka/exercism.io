def is_letter_char(c):
    return ord(c) in range(ord('a'), ord('z'))


def is_isogram(param_str: str):
    if len(param_str) == 0:
        return True
    for c in param_str.lower():
        if is_letter_char(c):
            if param_str.lower().count(c) > 1:
                return False

    return True
