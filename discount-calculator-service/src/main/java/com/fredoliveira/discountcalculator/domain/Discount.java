package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.app.utility.MoneyUtils;
import lombok.Builder;
import lombok.Getter;

import static java.math.BigDecimal.valueOf;

@Getter
@Builder
public class Discount {

  private final BigDecimal percentage;
  private final Long valueInCents;

  public Discount calculate(Long priceInCents, BigDecimal discount) {
    return Discount.builder()
      .percentage(discount)
      .valueInCents(getDiscountValue(priceInCents, discount))
      .build();
  }

  private Long getDiscountValue(Long priceInCents, BigDecimal discount) {
    return MoneyUtils.toCents(
      valueOf(priceInCents)
        .multiply(discount)
        .divide(valueOf(100))
    );
  }
}
