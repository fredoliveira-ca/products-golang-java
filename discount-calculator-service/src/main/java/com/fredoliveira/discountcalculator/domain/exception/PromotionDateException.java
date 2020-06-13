package com.fredoliveira.discountcalculator.domain.exception;

public class PromotionDateException extends RuntimeException {

  public PromotionDateException() {
    super("You need a date to use this type of promotion!");
  }

}
