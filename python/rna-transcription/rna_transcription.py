def is_valid(dna: str):
    """
    Chec if a given string is a valid DNA sequence.
    :param dna: The DNA strain.
    :return: True if is valid, False otherwise
    """
    return set(dna.upper()).issubset({'A', 'C', 'G', 'T'})


def to_rna(dna: str):
    if len(dna) == 0 or not is_valid(dna):
        return ""

    return dna \
        .upper() \
        .replace('G', '9') \
        .replace('C', 'G') \
        .replace('9', 'C') \
        .replace('A', '9') \
        .replace('T', 'A') \
        .replace('9', 'U')
