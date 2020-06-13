package com.fredoliveira.discountcalculator.data;

import java.math.BigDecimal;

import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.service.DiscountFinder;

import static java.math.BigDecimal.valueOf;
import static java.time.LocalDate.now;

public class BirthdayDiscount implements DiscountFinder {

  private static final java.util.logging.Logger log = java.util.logging.Logger.getLogger(BirthdayDiscount.class.getName());
  private static final BigDecimal BIRTHDAY_DISCOUNT = valueOf(0.10);

  private final String userId;

  public BirthdayDiscount(String userId) {
    this.userId = userId;
  }

  @Override public BigDecimal getDiscount() {
    if (isUserBirthDay()) {
      return BIRTHDAY_DISCOUNT;
    }

    return BigDecimal.ZERO;
  }

  private boolean isUserBirthDay() {
    try {
      return now().isEqual(new FetchUserGrpc().fetchBy(userId).getDateOfBirth());
    } catch (Exception e) {
      log.info("gRPC user server is down.");
      return false;
    }
  }

}
