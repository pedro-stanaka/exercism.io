def is_leap_year(year: int):
    if year % 4 == 0:
        return year % 100 != 0
    return False
