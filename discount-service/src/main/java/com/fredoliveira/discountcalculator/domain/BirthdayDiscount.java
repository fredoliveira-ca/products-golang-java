package com.fredoliveira.discountcalculator.domain;

import java.math.BigDecimal;
import java.util.logging.Logger;

import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.app.service.DiscountType;
import com.fredoliveira.discountcalculator.app.utility.DeLoreanMachine;
import lombok.RequiredArgsConstructor;

import static java.math.BigDecimal.valueOf;

@RequiredArgsConstructor public class BirthdayDiscount implements DiscountType {

  private static final Logger log = Logger.getLogger(BirthdayDiscount.class.getName());
  private static final BigDecimal BIRTHDAY_DISCOUNT = valueOf(0.05);

  private final String userId;
  private final FetchUserGrpc fetchUserGrpc;

  @Override public BigDecimal getDiscount() {
    if (isUserBirthDay()) {
      return BIRTHDAY_DISCOUNT;
    }
    return BigDecimal.ZERO;
  }

  private boolean isUserBirthDay() {
    try {
      return DeLoreanMachine.getToday().isEqual(fetchUserGrpc.fetchBy(userId).getDateOfBirth());
    } catch (Exception e) {
      log.info("gRPC user server is down.");
      return false;
    }
  }

}
