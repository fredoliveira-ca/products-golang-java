package com.fredoliveira.discountcalculator.data;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.app.service.DiscountFinder;

import static com.fredoliveira.discountcalculator.domain.Promotion.BLACK_FRIDAY;
import static java.math.BigDecimal.valueOf;
import static java.time.LocalDate.now;

public class BlackFridayDiscount implements DiscountFinder {

  public static final BigDecimal BLACK_FRIDAY_DISCOUNT = valueOf(0.10);

  @Override public BigDecimal getDiscount() {
    if (isBlackFriday()) {
      return BLACK_FRIDAY_DISCOUNT;
    }
    return BigDecimal.ZERO;
  }

  private boolean isBlackFriday() {
    return now().isEqual(BLACK_FRIDAY.getPromotionDate());
  }

}
