package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.Month;

import com.fredoliveira.discountcalculator.domain.exception.PromotionDateException;
import lombok.AllArgsConstructor;
import lombok.Getter;

import static java.time.LocalDate.now;
import static java.time.Month.NOVEMBER;
import static java.util.Objects.nonNull;

@Getter
@AllArgsConstructor
public enum Promotion {

  BIRTHDAY("desc", BigDecimal.valueOf(0.05), null, null),
  BLACK_FRIDAY("desc", BigDecimal.valueOf(0.10), NOVEMBER, 25);

  private final String description;
  private final BigDecimal discount;
  private final Month month;
  private final Integer day;

  public LocalDate getPromotionDate() {
    if (nonNull(this.getDay()) && nonNull(this.getMonth())) {
      return LocalDate.of(now().getYear(), this.getMonth(), this.getDay());
    } else {
      throw new PromotionDateException();
    }
  }

}
