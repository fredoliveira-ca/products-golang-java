package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.Month;

import lombok.AllArgsConstructor;
import lombok.Getter;

import static java.time.LocalDate.now;
import static java.time.Month.JUNE;

@Getter
@AllArgsConstructor
public enum Promotion {

  BIRTHDAY("desc", BigDecimal.valueOf(0.05), null, null),
  BLACK_FRIDAY("desc", BigDecimal.valueOf(0.10), JUNE, 11);

  private final String description;
  private final BigDecimal discount;
  private final Month month;
  private final Integer day;

  public LocalDate getPromotionDate() {
    return LocalDate.of(now().getYear(), this.getMonth(), this.getDay());
  }

}
