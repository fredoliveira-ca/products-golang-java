package com.fredoliveira.discountcalculator.app.service;

import com.fredoliveira.discountcalculator.app.grpc.user.FetchUserGrpc;
import com.fredoliveira.discountcalculator.domain.BirthdayDiscount;
import com.fredoliveira.discountcalculator.domain.BlackFridayDiscount;
import com.fredoliveira.discountcalculator.domain.Promotion;

public class DiscountStrategy {

  public DiscountType of(Promotion promotion, String userId, FetchUserGrpc fetchUserGrpc) {
    return switch (promotion) {
      case BIRTHDAY -> new BirthdayDiscount(userId, fetchUserGrpc);
      case BLACK_FRIDAY -> new BlackFridayDiscount();
    };
  }

}
