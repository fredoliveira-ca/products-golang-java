package com.fredoliveira.discountcalculator.app.utility;

import java.math.BigDecimal;

import lombok.experimental.UtilityClass;

import static java.util.Objects.nonNull;

@UtilityClass public class MoneyUtils {

  public Long toCents(BigDecimal amount) {
    if (nonNull(amount)) {
      return amount.multiply(BigDecimal.valueOf(100)).longValue();
    }
    return 0L;
  }

}
