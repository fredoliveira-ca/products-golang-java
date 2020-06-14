package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.app.service.DiscountType;
import com.fredoliveira.discountcalculator.app.utility.DeLoreanMachine;

import static com.fredoliveira.discountcalculator.domain.Promotion.BLACK_FRIDAY;
import static java.math.BigDecimal.valueOf;

public class BlackFridayDiscount implements DiscountType {

  public static final BigDecimal BLACK_FRIDAY_DISCOUNT = valueOf(0.10);

  @Override
  public BigDecimal getDiscount() {
    if (isBlackFriday()) {
      return BLACK_FRIDAY_DISCOUNT;
    }
    return BigDecimal.ZERO;
  }

  private boolean isBlackFriday() {
    return DeLoreanMachine.getToday().isEqual(BLACK_FRIDAY.getPromotionDate());
  }

}
