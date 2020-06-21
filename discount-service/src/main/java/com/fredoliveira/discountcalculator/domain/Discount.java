package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.app.utility.MoneyUtils;
import lombok.Builder;
import lombok.Getter;

import static java.math.BigDecimal.ZERO;
import static java.math.BigDecimal.valueOf;

@Getter
@Builder
public class Discount {

  public static final BigDecimal LIMIT_DISCOUNT = valueOf(0.10);
  public static final BigDecimal DEFAULT_DISCOUNT = valueOf(0);

  private final BigDecimal percentage;
  private final Long valueInCents;

  public Discount calculate(Long priceInCents, BigDecimal discount) {
    if(priceInCents <= 0L) {
      return Discount.builder()
        .percentage(ZERO)
        .valueInCents(0L)
        .build();
    }

    return Discount.builder()
      .percentage(discount)
      .valueInCents(MoneyUtils.getDiscountValue(priceInCents, discount))
      .build();
  }

}
