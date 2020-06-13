package com.fredoliveira.discountcalculator.app.service;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.domain.Discount;
import com.fredoliveira.discountcalculator.domain.Promotion;

import static com.fredoliveira.discountcalculator.app.service.DiscountStrategy.of;
import static java.math.BigDecimal.valueOf;
import static java.util.Arrays.stream;

public class DiscountCalculator {

  public static final BigDecimal LIMIT_DISCOUNT = valueOf(0.10);
  public static final BigDecimal DEFAULT_DISCOUNT = valueOf(0);

  public Discount calculateDiscount(Long priceInCents, String userId) {
    final var totalDiscount = stream(Promotion.values())
        .map(promotion -> of(promotion, userId).getDiscount())
        .reduce(BigDecimal.ZERO, BigDecimal::add);

    var discount = DEFAULT_DISCOUNT;
    if (totalDiscount.compareTo(DEFAULT_DISCOUNT) > 0) {
      discount = LIMIT_DISCOUNT.max(totalDiscount).min(LIMIT_DISCOUNT);
    }

    return Discount.builder().build()
        .calculate(priceInCents, discount);
  }

}
