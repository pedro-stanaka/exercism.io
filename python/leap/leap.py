def is_leap_year(year: int):
    if year % 4 == 0:
        if year % 100:
            return False
        return True
